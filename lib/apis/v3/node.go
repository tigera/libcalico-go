// Copyright (c) 2017,2020 Tigera, Inc. All rights reserved.

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

	"github.com/projectcalico/libcalico-go/lib/numorstring"
)

const (
	KindNode     = "Node"
	KindNodeList = "NodeList"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Node contains information about a Node resource.
type Node struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the Node.
	Spec NodeSpec `json:"spec,omitempty"`
	// Status of the Node.
	Status NodeStatus `json:"status,omitempty"`
}

// NodeSpec contains the specification for a Node resource.
type NodeSpec struct {
	// BGP configuration for this node.
	BGP *NodeBGPSpec `json:"bgp,omitempty" validate:"omitempty"`

	// IPv4VXLANTunnelAddr is the IPv4 address of the VXLAN tunnel.
	IPv4VXLANTunnelAddr string `json:"ipv4VXLANTunnelAddr,omitempty" validate:"omitempty,ipv4"`

	// VXLANTunnelMACAddr is the MAC address of the VXLAN tunnel.
	VXLANTunnelMACAddr string `json:"vxlanTunnelMACAddr,omitempty" validate:"omitempty,mac"`

	// OrchRefs for this node.
	OrchRefs []OrchRef `json:"orchRefs,omitempty" validate:"omitempty"`

	// Wireguard configuration for this node.
	Wireguard *NodeWireguardSpec `json:"wireguard,omitempty" validate:"omitempty"`

	// Addresses lists all known addresses and their types. Addresses are
	// possibly on different ISO/OSI levels, that is the address can be IP or MAC
	// or something else, denoted by its type.
	Addresses []NodeAddress `json:"addresses,omitempty" validate:"omitempty"`
}

type NodeStatus struct {
	// WireguardPublicKey is the Wireguard public-key for this node.
	// wireguardPublicKey validates if the string is a valid base64 encoded key.
	WireguardPublicKey string `json:"wireguardPublicKey,omitempty" validate:"omitempty,wireguardPublicKey"`

	// PodCIDR is a reflection of the Kubernetes node's spec.PodCIDRs field.
	PodCIDRs []string `json:"podCIDRs,omitempty" validate:"omitempty"`
}

// OrchRef is used to correlate a Calico node to its corresponding representation in a given orchestrator
type OrchRef struct {
	// NodeName represents the name for this node according to the orchestrator.
	NodeName string `json:"nodeName,omitempty" validate:"omitempty"`
	// Orchestrator represents the orchestrator using this node.
	Orchestrator string `json:"orchestrator"`
}

// NodeBGPSpec contains the specification for the Node BGP configuration.
type NodeBGPSpec struct {
	// The AS Number of the node.  If this is not specified, the global
	// default value will be used.
	ASNumber *numorstring.ASNumber `json:"asNumber,omitempty"`
	// IPv4Address is the IPv4 address and network of this node.  The IPv4 address
	// should always be specified if you are using BGP.
	IPv4Address string `json:"ipv4Address,omitempty" validate:"omitempty,cidrv4"`
	// IPv6Address is the IPv6 address and network of this node.  Not required if you
	// are not using BGP or you do not require IPv6 routing.
	IPv6Address string `json:"ipv6Address,omitempty" validate:"omitempty,cidrv6"`
	// IPv4IPIPTunnelAddr is the IPv4 address of the IP in IP tunnel.
	IPv4IPIPTunnelAddr string `json:"ipv4IPIPTunnelAddr,omitempty" validate:"omitempty,ipv4"`
	// RouteReflectorClusterID enables this node as a route reflector within the given
	// cluster.
	RouteReflectorClusterID string `json:"routeReflectorClusterID,omitempty" validate:"omitempty,ipv4"`
}

// NodeWireguardSpec contains the specification for the Node wireguard configuration.
type NodeWireguardSpec struct {
	// InterfaceIPv4Address is the IPv4 address for the Wireguard interface.
	InterfaceIPv4Address string `json:"interfaceIPv4Address,omitempty" validate:"omitempty,ipv4"`
}

// NodeAddressType represents a node address type as a string
type NodeAddressType string

const (
	NodeAddressIP                NodeAddressType = "IP"
	NodeAddressVxlanTunnelIP     NodeAddressType = "VxlanTunnelIP"
	NodeAddressVxlanMAC          NodeAddressType = "VxlanMAC"
	NodeAddressIPIPTunnelIP      NodeAddressType = "IPIPTunnelIP"
	NodeAddressWireguardTunnelIP NodeAddressType = "WireguardTunnelIP"
	NodeAddressK8sInternalIP     NodeAddressType = "K8sInternalIP"
	NodeAddressK8sExternalIP     NodeAddressType = "K8sExternalIP"
)

// NodeAddress represents an address assigned to the node annotated with its type.
type NodeAddress struct {
	// Type denotes what sort of address it is.
	Type NodeAddressType `json:"type"`
	// Address is a string rpresentation of the actual address.
	Address string `json:"address"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeList contains a list of Node resources.
type NodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Node `json:"items"`
}

// NewNode creates a new (zeroed) Node struct with the TypeMetadata initialised to the current
// version.
func NewNode() *Node {
	return &Node{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindNode,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewNodeList creates a new (zeroed) NodeList struct with the TypeMetadata initialised to the current
// version.
func NewNodeList() *NodeList {
	return &NodeList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindNodeList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
