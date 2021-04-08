// Copyright (c) 2020 Tigera, Inc. All rights reserved.

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

package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindKubeControllersConfiguration     = "KubeControllersConfiguration"
	KindKubeControllersConfigurationList = "KubeControllersConfigurationList"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeControllersConfiguration contains the configuration for Calico Kubernetes Controllers.
type KubeControllersConfiguration struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the KubeControllersConfiguration.
	Spec KubeControllersConfigurationSpec `json:"spec,omitempty"`
	// Status of the KubeControllersConfiguration
	Status KubeControllersConfigurationStatus `json:"status,omitempty"`
}

// KubeControllersConfigurationSpec contains the values of the Kubernetes controllers configuration.
type KubeControllersConfigurationSpec struct {
	// LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]
	LogSeverityScreen string `json:"logSeverityScreen,omitempty" validate:"omitempty,logLevel"`

	// HealthChecks enables or disables support for health checks [Default: Enabled]
	HealthChecks string `json:"healthChecks,omitempty" validate:"omitempty,oneof=Enabled Disabled"`

	// EtcdV3CompactionPeriod is the period between etcdv3 compaction requests. Set to 0 to disable. [Default: 10m]
	EtcdV3CompactionPeriod *metav1.Duration `json:"etcdV3CompactionPeriod,omitempty" validate:"omitempty"`

	// PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to. Set to 0 to disable. [Default: 9094]
	PrometheusMetricsPort *int `json:"prometheusMetricsPort,omitempty"`

	// Controllers enables and configures individual Kubernetes controllers
	Controllers ControllersConfig `json:"controllers"`
}

// ControllersConfig enables and configures individual Kubernetes controllers
type ControllersConfig struct {
	// Node enables and configures the node controller. Enabled by default, set to nil to disable.
	Node *NodeControllerConfig `json:"node,omitempty"`

	// Policy enables and configures the policy controller. Enabled by default, set to nil to disable.
	Policy *PolicyControllerConfig `json:"policy,omitempty"`

	// WorkloadEndpoint enables and configures the workload endpoint controller. Enabled by default, set to nil to disable.
	WorkloadEndpoint *WorkloadEndpointControllerConfig `json:"workloadEndpoint,omitempty"`

	// ServiceAccount enables and configures the service account controller. Enabled by default, set to nil to disable.
	ServiceAccount *ServiceAccountControllerConfig `json:"serviceAccount,omitempty"`

	// Namespace enables and configures the namespace controller. Enabled by default, set to nil to disable.
	Namespace *NamespaceControllerConfig `json:"namespace,omitempty"`

	// RouteReflector enables and configures the route reflector controller. Disabled by default, set value to enable.
	RouteReflector *RouteReflectorControllerConfig `json:"routereflector,omitempty"`
}

// NodeControllerConfig configures the node controller, which automatically cleans up configuration
// for nodes that no longer exist. Optionally, it can create host endpoints for all Kubernetes nodes.
type NodeControllerConfig struct {
	// ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]
	ReconcilerPeriod *metav1.Duration `json:"reconcilerPeriod,omitempty" validate:"omitempty"`

	// SyncLabels controls whether to copy Kubernetes node labels to Calico nodes. [Default: Enabled]
	SyncLabels string `json:"syncLabels,omitempty" validate:"omitempty,oneof=Enabled Disabled"`

	// HostEndpoint controls syncing nodes to host endpoints. Disabled by default, set to nil to disable.
	HostEndpoint *AutoHostEndpointConfig `json:"hostEndpoint,omitempty"`
}

type AutoHostEndpointConfig struct {
	// AutoCreate enables automatic creation of host endpoints for every node. [Default: Disabled]
	AutoCreate string `json:"autoCreate,omitempty" validate:"omitempty,oneof=Enabled Disabled"`
}

// PolicyControllerConfig configures the network policy controller, which syncs Kubernetes policies
// to Calico policies (only used for etcdv3 datastore).
type PolicyControllerConfig struct {
	// ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]
	ReconcilerPeriod *metav1.Duration `json:"reconcilerPeriod,omitempty" validate:"omitempty"`
}

// WorkloadEndpointControllerConfig configures the workload endpoint controller, which syncs Kubernetes
// labels to Calico workload endpoints (only used for etcdv3 datastore).
type WorkloadEndpointControllerConfig struct {
	// ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]
	ReconcilerPeriod *metav1.Duration `json:"reconcilerPeriod,omitempty" validate:"omitempty"`
}

// ServiceAccountControllerConfig configures the service account controller, which syncs Kubernetes
// service accounts to Calico profiles (only used for etcdv3 datastore).
type ServiceAccountControllerConfig struct {
	// ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]
	ReconcilerPeriod *metav1.Duration `json:"reconcilerPeriod,omitempty" validate:"omitempty"`
}

// NamespaceControllerConfig configures the service account controller, which syncs Kubernetes
// service accounts to Calico profiles (only used for etcdv3 datastore).
type NamespaceControllerConfig struct {
	// ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]
	ReconcilerPeriod *metav1.Duration `json:"reconcilerPeriod,omitempty" validate:"omitempty"`
}

// RouteReflectorControllerConfig configures the route reflector controller, which scales Calico
// route reflector topology.
type RouteReflectorControllerConfig struct {
	// ReconcilerPeriod is the period to perform reconciliation with the Calico datastore. [Default: 5m]
	ReconcilerPeriod *metav1.Duration `json:"reconcilerPeriod,omitempty" validate:"omitempty"`

	// TopologyType deines the type of topology, which can be [single, multi]. [Default: multi]
	// +kubebuilder:default=multi
	// +kubebuilder:validation:Pattern=`^(single|multi)$`
	TopologyType *string `json:"topologyType,omitempty" validate:"omitempty,oneof=single multi"`

	// ClusterID is the Route Reflector cluster id. Multi cluster topology uses zeros as wildcard. [Default: 224.0.0.0]
	// +kubebuilder:default="224.0.0.0"
	// +kubebuilder:validation:Pattern=`^\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}\b$`
	ClusterID *string `json:"clusterId,omitempty" validate:"omitempty"`

	// Min the minimum number of Route Refletors. [Default: 3]
	// +kubebuilder:default=3
	// +kubebuilder:validation:Minimum:=1
	Min *int `json:"min,omitempty" validate:"omitempty"`

	// Max the maxium number of Route Refletors. [Default: 10]
	// +kubebuilder:default=10
	// +kubebuilder:validation:Minimum:=3
	Max *int `json:"max,omitempty" validate:"omitempty"`

	// Ratio defines the ration of Route Reflectors and clients. [Default: 0.005]
	Ratio *float32 `json:"ratio,omitempty" validate:"omitempty"`

	// ReflectorsPerNode the number of route reflectors per client. Single cluster topology ignores. [Default: 3]
	// +kubebuilder:default=3
	// +kubebuilder:validation:Minimum:=1
	RouteReflectorsPerNode *int `json:"routeReflectorsPerNode,omitempty" validate:"omitempty"`

	// RouteReflectorLabelKey label key of Route Reflector selector. [Default: calico-route-reflector]
	// +kubebuilder:default=calico-route-reflector
	RouteReflectorLabelKey *string `json:"routeReflectorLabelKey,omitempty" validate:"omitempty"`

	// RouteReflectorLabelValue label value of Route Reflector selector.
	RouteReflectorLabelValue *string `json:"routeReflectorLabelValue,omitempty" validate:"omitempty"`

	// ZoneLabel zone label on Kubernetes nodes. [Default: failure-domain.beta.kubernetes.io/zone]
	// +kubebuilder:default=failure-domain.beta.kubernetes.io/zone
	ZoneLabel *string `json:"zoneLabel,omitempty" validate:"omitempty"`

	// HostnameLabel hostname label on Kubernetes nodes. [Default: kubernetes.io/hostname]
	// +kubebuilder:default=kubernetes.io/hostname
	HostnameLabel *string `json:"hostnameLabel,omitempty" validate:"omitempty"`

	// IncompatibleLabels List of node labels to disallow Route Reflector selection.
	IncompatibleLabels *string `json:"incompatibleLabels,omitempty" validate:"omitempty"`
}

// KubeControllersConfigurationStatus represents the status of the configuration. It's useful for admins to
// be able to see the actual config that was applied, which can be modified by environment variables on the
// kube-controllers process.
type KubeControllersConfigurationStatus struct {
	// RunningConfig contains the effective config that is running in the kube-controllers pod, after
	// merging the API resource with any environment variables.
	RunningConfig KubeControllersConfigurationSpec `json:"runningConfig,omitempty"`

	// EnvironmentVars contains the environment variables on the kube-controllers that influenced
	// the RunningConfig.
	EnvironmentVars map[string]string `json:"environmentVars,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeControllersConfigurationList contains a list of KubeControllersConfiguration resources.
type KubeControllersConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KubeControllersConfiguration `json:"items"`
}

// New KubeControllersConfiguration creates a new (zeroed) KubeControllersConfiguration struct with
// the TypeMetadata initialized to the current version.
func NewKubeControllersConfiguration() *KubeControllersConfiguration {
	return &KubeControllersConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindKubeControllersConfiguration,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewKubeControllersConfigurationList creates a new (zeroed) KubeControllersConfigurationList struct with the TypeMetadata
// initialized to the current version.
func NewKubeControllersConfigurationList() *KubeControllersConfigurationList {
	return &KubeControllersConfigurationList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindKubeControllersConfigurationList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
