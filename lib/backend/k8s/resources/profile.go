// Copyright (c) 2016-2017 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resources

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
	kapiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/projectcalico/libcalico-go/lib/backend/api"
	"github.com/projectcalico/libcalico-go/lib/backend/k8s/conversion"
	"github.com/projectcalico/libcalico-go/lib/backend/model"
	cerrors "github.com/projectcalico/libcalico-go/lib/errors"
)

func NewProfileClient(c *kubernetes.Clientset) K8sResourceClient {
	return &profileClient{
		clientSet: c,
	}
}

// Implements the api.Client interface for Profiles.
type profileClient struct {
	clientSet *kubernetes.Clientset
	converter conversion.Converter
}

func (c *profileClient) Create(ctx context.Context, kvp *model.KVPair) (*model.KVPair, error) {
	log.Warn("Operation Create is not supported on Profile type")
	return nil, cerrors.ErrorOperationNotSupported{
		Identifier: kvp.Key,
		Operation:  "Create",
	}
}

func (c *profileClient) Update(ctx context.Context, kvp *model.KVPair) (*model.KVPair, error) {
	log.Warn("Operation Update is not supported on Profile type")
	return nil, cerrors.ErrorOperationNotSupported{
		Identifier: kvp.Key,
		Operation:  "Update",
	}
}

func (c *profileClient) Delete(ctx context.Context, key model.Key, revision string) (*model.KVPair, error) {
	log.Warn("Operation Delete is not supported on Profile type")
	return nil, cerrors.ErrorOperationNotSupported{
		Identifier: key,
		Operation:  "Delete",
	}
}

func (c *profileClient) getSaKv(sa *kapiv1.ServiceAccount) (*model.KVPair, error) {
	kvPair, err := c.converter.ServiceAccountToProfile(sa)
	if err != nil {
		return nil, err
	}

	kvPair.Revision = joinRev("", kvPair.Revision)
	return kvPair, nil
}

func (c *profileClient) getServiceAccount(ctx context.Context, namespace, name, revision string) (*model.KVPair, error) {
	serviceAccountName, err := c.converter.ProfileNameToServiceAccount(name)
	if err != nil {
		return nil, err
	}

	serviceAccount, err := c.clientSet.CoreV1().ServiceAccounts(namespace).Get(serviceAccountName, metav1.GetOptions{ResourceVersion: revision})
	if err != nil {
		return nil, err
	}

	return c.getSaKv(serviceAccount)
}

func (c *profileClient) getNsKv(ns *kapiv1.Namespace) (*model.KVPair, error) {
	kvPair, err := c.converter.NamespaceToProfile(ns)
	if err != nil {
		return nil, err
	}

	kvPair.Revision = joinRev(kvPair.Revision, "")
	return kvPair, nil
}

func (c *profileClient) getNamespace(ctx context.Context, name, revision string) (*model.KVPair, error) {
	namespaceName, err := c.converter.ProfileNameToNamespace(name)
	if err != nil {
		return nil, err
	}

	namespace, err := c.clientSet.CoreV1().Namespaces().Get(namespaceName, metav1.GetOptions{ResourceVersion: revision})
	if err != nil {
		return nil, err
	}

	return c.getNsKv(namespace)
}

func (c *profileClient) Get(ctx context.Context, key model.Key, revision string) (*model.KVPair, error) {
	log.Debug("Received Get request on Profile type")
	rk := key.(model.ResourceKey)
	if rk.Name == "" {
		return nil, fmt.Errorf("Profile key missing name: %+v", rk)
	}

	nsRev, saRev, err := splitRev(revision)
	if err != nil {
		return nil, err
	}

	if nsRev != "" && saRev != "" {
		return nil, fmt.Errorf("Invalid revision in get %s", revision)
	}

	if nsRev != "" {
		return c.getNamespace(ctx, rk.Name, nsRev)
	}

	if saRev != "" {
		return c.getServiceAccount(ctx, kapiv1.NamespaceAll, rk.Name, saRev)
	}

	return nil, fmt.Errorf("Revision %s invalid", revision)
}

func (c *profileClient) List(ctx context.Context, list model.ListInterface, revision string) (*model.KVPairList, error) {
	log.Debug("Received List request on Profile type")
	nl := list.(model.ResourceListOptions)
	kvps := []*model.KVPair{}

	// If a name is specified, then do an exact lookup.
	if nl.Name != "" {
		kvp, err := c.Get(ctx, model.ResourceKey{Name: nl.Name, Kind: nl.Kind}, revision)
		if err != nil {
			if _, ok := err.(cerrors.ErrorResourceDoesNotExist); !ok {
				return nil, err
			}
			return &model.KVPairList{
				KVPairs:  kvps,
				Revision: revision,
			}, nil
		}

		kvps = append(kvps, kvp)
		return &model.KVPairList{
			KVPairs:  []*model.KVPair{kvp},
			Revision: revision,
		}, nil
	}

	var nsRev, saRev string
	var err error
	// First call to list the revision may be empty
	// call split only if revision is not null.
	if revision != "" {
		nsRev, saRev, err = splitRev(revision)
		if err != nil {
			return nil, err
		}
	}

	// Otherwise, enumerate all.
	namespaces, err := c.clientSet.CoreV1().Namespaces().List(metav1.ListOptions{ResourceVersion: nsRev})
	if err != nil {
		return nil, K8sErrorToCalico(err, nl)
	}

	// For each Namespace, return a profile.
	for _, ns := range namespaces.Items {
		kvp, err := c.getNsKv(&ns)
		if err != nil {
			log.Errorf("Unable to convert k8s Namespace to Calico Profile: Namespace=%s: %v", ns.Name, err)
			continue
		}
		kvps = append(kvps, kvp)
	}

	// Enumerate all SA
	var serviceaccounts *kapiv1.ServiceAccountList
	// TBD: narrow down to only to the required namespace
	serviceaccounts, err = c.clientSet.CoreV1().ServiceAccounts(kapiv1.NamespaceAll).List(metav1.ListOptions{ResourceVersion: saRev})
	if err != nil {
		return nil, K8sErrorToCalico(err, nl)
	}

	for _, sa := range serviceaccounts.Items {
		kvp, err := c.getSaKv(&sa)
		if err != nil {
			log.Errorf("Unable to convert k8s service account to Calico Profile: %s: %v", sa.Name, err)
			continue
		}
		log.Debug("Converted k8s sa to Calico profile ", sa.Name)
		kvps = append(kvps, kvp)
	}
	return &model.KVPairList{
		KVPairs:  kvps,
		Revision: namespaces.ResourceVersion + "/" + serviceaccounts.ResourceVersion,
	}, nil
}

func (c *profileClient) EnsureInitialized() error {
	return nil
}

func (c *profileClient) Watch(ctx context.Context, list model.ListInterface, revision string) (api.WatchInterface, error) {
	resl := list.(model.ResourceListOptions)
	if len(resl.Name) != 0 {
		return nil, fmt.Errorf("cannot watch specific resource instance: %s", list.(model.ResourceListOptions).Name)
	}

	nsRev, saRev, err := splitRev(revision)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"NSPROFRev": nsRev,
		"SAPROFRev": saRev,
	}).Info("Watching two resources at individual revisions")

	nsWatch, err := c.clientSet.CoreV1().Namespaces().Watch(metav1.ListOptions{ResourceVersion: nsRev})
	if err != nil {
		return nil, K8sErrorToCalico(err, list)
	}
	converter := func(r Resource) (*model.KVPair, error) {
		k8sNamespace, ok := r.(*kapiv1.Namespace)
		if !ok {
			return nil, errors.New("profile conversion with incorrect k8s resource type")
		}
		return c.converter.NamespaceToProfile(k8sNamespace)
	}

	nsWatcher := newK8sWatcherConverter(ctx, "Profile-NS", converter, nsWatch)

	// Watch all service accounts in ALL namespaces
	saWatch, err := c.clientSet.CoreV1().ServiceAccounts(kapiv1.NamespaceAll).Watch(metav1.ListOptions{ResourceVersion: saRev})
	if err != nil {
		nsWatch.Stop()
		return nil, err
	}

	converterSA := func(r Resource) (*model.KVPair, error) {
		k8sServiceAccount, ok := r.(*kapiv1.ServiceAccount)
		if !ok {
			nsWatch.Stop()
			return nil, errors.New("Profile converion with incorrect k8s resource type")
		}
		return c.converter.ServiceAccountToProfile(k8sServiceAccount)
	}

	saWatcher := newK8sWatcherConverter(ctx, "Profile-SA", converterSA, saWatch)

	return newProfileWatcher(ctx, nsWatcher, saWatcher), nil

}

func newProfileWatcher(ctx context.Context, k8sWatchNS, k8sWatchSA api.WatchInterface) api.WatchInterface {
	ctx, cancel := context.WithCancel(ctx)
	wc := &profileWatcher{
		k8sNSWatch: k8sWatchNS,
		k8sSAWatch: k8sWatchSA,
		context:    ctx,
		cancel:     cancel,
		resultChan: make(chan api.WatchEvent, resultsBufSize),
	}
	go wc.processProfileEvents()
	return wc
}

type profileWatcher struct {
	converter  ConvertK8sResourceToKVPair
	k8sNSWatch api.WatchInterface
	k8sSAWatch api.WatchInterface
	k8sNSRev   string
	k8sSARev   string
	context    context.Context
	cancel     context.CancelFunc
	resultChan chan api.WatchEvent
	terminated uint32
}

// Stop stops the watcher and releases associated resources.
// This calls through the context cancel function.
func (pw *profileWatcher) Stop() {
	pw.cancel()
	pw.k8sNSWatch.Stop()
	pw.k8sSAWatch.Stop()
}

// ResultChan returns a channel used to receive WatchEvents.
func (pw *profileWatcher) ResultChan() <-chan api.WatchEvent {
	return pw.resultChan
}

// HasTerminated returns true when the watcher has completed termination processing.
func (pw *profileWatcher) HasTerminated() bool {
	terminated := atomic.LoadUint32(&pw.terminated) != 0

	if pw.k8sNSWatch != nil {
		terminated = terminated && pw.k8sNSWatch.HasTerminated()
	}
	if pw.k8sSAWatch != nil {
		terminated = terminated && pw.k8sSAWatch.HasTerminated()
	}

	return terminated
}

// Loop to process the events stream from the underlying k8s Watcher and convert them to
// backend KVPs.
func (pw *profileWatcher) processProfileEvents() {
	log.Info("Watcher process started for profile.")
	defer func() {
		log.Info("Profile watcher process terminated")
		pw.Stop()
		close(pw.resultChan)
		atomic.AddUint32(&pw.terminated, 1)
	}()

	var e api.WatchEvent
	var isNsEvent bool
	var value interface{}
	for {
		select {
		case e = <-pw.k8sNSWatch.ResultChan():
			log.Debug("Processing Namespace event")
			isNsEvent = true

		case e = <-pw.k8sSAWatch.ResultChan():
			log.Debug("Processing SeriveAccount event")
			isNsEvent = false

		case <-pw.context.Done(): //user cancel
			log.Info("Process watcher done event in kdd client")
			return
		}

		switch e.Type {
		case api.WatchModified, api.WatchAdded:
			value = e.New.Value
		case api.WatchDeleted:
			value = e.Old.Value
		}
		oma, ok := value.(metav1.ObjectMetaAccessor)

		if !ok {
			log.Error("Resource returned from watch does not implement ObjectMetaAccessor interface")
			return
		}
		if isNsEvent {
			pw.k8sNSRev = oma.GetObjectMeta().GetResourceVersion()
		} else {
			pw.k8sSARev = oma.GetObjectMeta().GetResourceVersion()
		}
		oma.GetObjectMeta().SetResourceVersion(joinRev(pw.k8sNSRev, pw.k8sSARev))

		// Send the processed event.
		select {
		case pw.resultChan <- e:
			// If this is an error event. check to see if it's a terminating one.
			// If so, terminate this watcher.
			if e.Type == api.WatchError {
				log.WithError(e.Error).Debug("Kubernetes event converted to backend watcher error event")
				if _, ok := e.Error.(cerrors.ErrorWatchTerminated); ok {
					log.Info("Watch terminated event")
					return
				}
			}

		case <-pw.context.Done():
			log.Info("Process watcher done event during watch event in kdd client")
			return
		}
	}
}

// joinRev combines the revesion of namespace and serviceaccount
func joinRev(nsRev, saRev string) string {
	return nsRev + "/" + saRev
}

// splitRev takes a revision string and splits it into the namespace and serviceaccount componenents
func splitRev(rev string) (nsRev, saRev string, err error) {
	if rev == "" {
		return "", "", fmt.Errorf("Invalid nil string to splitRev")
	}

	revs := strings.Split(rev, "/")
	if len(revs) != 2 {
		return "", "", fmt.Errorf("Invalid rev %s", rev)
	}

	return revs[0], revs[1], nil
}
