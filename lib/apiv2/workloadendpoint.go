// Copyright (c) 2017 Tigera, Inc. All rights reserved.

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

package apiv2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	KindWorkloadEndpoint     = "WorkloadEndpoint"
	KindWorkloadEndpointList = "WorkloadEndpointList"
)

// WorkloadEndpoint contains information about a WorkloadEndpoint resource that is a peer of a Calico
// compute node.
type WorkloadEndpoint struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	Metadata metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the WorkloadEndpoint.
	Spec WorkloadEndpointSpec `json:"spec,omitempty"`
}

// WorkloadEndpointMetadata contains the specification for a WorkloadEndpoint resource.
type WorkloadEndpointSpec struct {
	// The name of the orchestrator.
	Orchestrator string `json:"orchestrator,omitempty" validate:"omitempty,namespacedname"`
	// The name of the workload.
	Workload string `json:"workload,omitempty" validate:"omitempty,namespacedname"`
	// The node name identifying the Calico node instance.
	Node string `json:"node,omitempty" validate:"omitempty,name"`
	// This field is not an index field of the WorkloadEndpoint resource.
	ContainerID string `json:"containerID,omitempty" validate:"omitempty,name"`
	// This field is not an index field of the WorkloadEndpoint resource.
	PodID string `json:"podID,omitempty" validate:"omitempty,name"`
	// IPNetworks is a list of subnets allocated to this endpoint. IP packets will only be
	// allowed to leave this interface if they come from an address in one of these subnets.
	// Currently only /32 for IPv4 and /128 for IPv6 networks are supported.
	IPNetworks []string `json:"ipNetworks,omitempty" validate:"omitempty,dive,cidr"`
	// IPNATs is a list of 1:1 NAT mappings to apply to the endpoint. Inbound connections
	// to the external IP will be forwarded to the internal IP. Connections initiated from the
	// internal IP will not have their source address changed, except when an endpoint attempts
	// to connect one of its own external IPs. Each internal IP must be associated with the same
	// endpoint via the configured IPNetworks.
	IPNATs []IPNAT `json:"ipNATs,omitempty" validate:"omitempty,dive"`
	// IPv4Gateway is the gateway IPv4 address for traffic from the workload.
	IPv4Gateway string `json:"ipv4Gateway,omitempty" validate:"omitempty,ipv4"`
	// IPv6Gateway is the gateway IPv6 address for traffic from the workload.
	IPv6Gateway string `json:"ipv6Gateway,omitempty" validate:"omitempty,ipv6"`
	// A list of security Profile resources that apply to this endpoint. Each profile is
	// applied in the order that they appear in this list.  Profile rules are applied
	// after the selector-based security policy.
	Profiles []string `json:"profiles,omitempty" validate:"omitempty,dive,namespacedname"`
	// InterfaceName the name of the Linux interface on the host: for example, tap80.
	InterfaceName string `json:"interfaceName,omitempty" validate:"interface"`
	// MAC is the MAC address of the endpoint interface.
	MAC string `json:"mac,omitempty" validate:"omitempty,mac"`
}

// IPNat contains a single NAT mapping for a WorkloadEndpoint resource.
type IPNAT struct {
	// The internal IP address which must be associated with the owning endpoint via the
	// configured IPNetworks for the endpoint.
	InternalIP string `json:"internalIP" validate:"omitempty,ip"`
	// The external IP address.
	ExternalIP string `json:"externalIP" validate:"omitempty,ip"`
}

// WorkloadEndpointList contains a list of WorkloadEndpoint resources.
type WorkloadEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ListMeta    `json:"metadata"`
	Items           []WorkloadEndpoint `json:"items"`
}

// NewWorkloadEndpoint creates a new (zeroed) WorkloadEndpoint struct with the TypeMetadata initialised to the current
// version.
func NewWorkloadEndpoint() *WorkloadEndpoint {
	return &WorkloadEndpoint{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindWorkloadEndpoint,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewWorkloadEndpointList creates a new (zeroed) WorkloadEndpointList struct with the TypeMetadata initialised to the current
// version.
func NewWorkloadEndpointList() *WorkloadEndpointList {
	return &WorkloadEndpointList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindWorkloadEndpointList,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// GetObjectKind returns the kind of this object.  Required to satisfy Object interface
func (e *WorkloadEndpoint) GetObjectKind() schema.ObjectKind {
	return &e.TypeMeta
}

// GetObjectMeta returns the object metadata of this object. Required to satisfy ObjectMetaAccessor interface
func (e *WorkloadEndpoint) GetObjectMeta() metav1.Object {
	return &e.Metadata
}

// GetObjectKind returns the kind of this object. Required to satisfy Object interface
func (el *WorkloadEndpointList) GetObjectKind() schema.ObjectKind {
	return &el.TypeMeta
}

// GetListMeta returns the list metadata of this object. Required to satisfy ListMetaAccessor interface
func (el *WorkloadEndpointList) GetListMeta() metav1.List {
	return &el.Metadata
}
