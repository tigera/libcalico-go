// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package crd

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BGPPeer).DeepCopyInto(out.(*BGPPeer))
			return nil
		}, InType: reflect.TypeOf(&BGPPeer{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BGPPeerCopy).DeepCopyInto(out.(*BGPPeerCopy))
			return nil
		}, InType: reflect.TypeOf(&BGPPeerCopy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BGPPeerList).DeepCopyInto(out.(*BGPPeerList))
			return nil
		}, InType: reflect.TypeOf(&BGPPeerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BGPPeerListCopy).DeepCopyInto(out.(*BGPPeerListCopy))
			return nil
		}, InType: reflect.TypeOf(&BGPPeerListCopy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BGPPeerSpec).DeepCopyInto(out.(*BGPPeerSpec))
			return nil
		}, InType: reflect.TypeOf(&BGPPeerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalBGPConfig).DeepCopyInto(out.(*GlobalBGPConfig))
			return nil
		}, InType: reflect.TypeOf(&GlobalBGPConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalBGPConfigList).DeepCopyInto(out.(*GlobalBGPConfigList))
			return nil
		}, InType: reflect.TypeOf(&GlobalBGPConfigList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalBGPConfigSpec).DeepCopyInto(out.(*GlobalBGPConfigSpec))
			return nil
		}, InType: reflect.TypeOf(&GlobalBGPConfigSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalFelixConfig).DeepCopyInto(out.(*GlobalFelixConfig))
			return nil
		}, InType: reflect.TypeOf(&GlobalFelixConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalFelixConfigCopy).DeepCopyInto(out.(*GlobalFelixConfigCopy))
			return nil
		}, InType: reflect.TypeOf(&GlobalFelixConfigCopy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalFelixConfigList).DeepCopyInto(out.(*GlobalFelixConfigList))
			return nil
		}, InType: reflect.TypeOf(&GlobalFelixConfigList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalFelixConfigListCopy).DeepCopyInto(out.(*GlobalFelixConfigListCopy))
			return nil
		}, InType: reflect.TypeOf(&GlobalFelixConfigListCopy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalFelixConfigSpec).DeepCopyInto(out.(*GlobalFelixConfigSpec))
			return nil
		}, InType: reflect.TypeOf(&GlobalFelixConfigSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalNetworkPolicy).DeepCopyInto(out.(*GlobalNetworkPolicy))
			return nil
		}, InType: reflect.TypeOf(&GlobalNetworkPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalNetworkPolicyCopy).DeepCopyInto(out.(*GlobalNetworkPolicyCopy))
			return nil
		}, InType: reflect.TypeOf(&GlobalNetworkPolicyCopy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalNetworkPolicyList).DeepCopyInto(out.(*GlobalNetworkPolicyList))
			return nil
		}, InType: reflect.TypeOf(&GlobalNetworkPolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalNetworkPolicyListCopy).DeepCopyInto(out.(*GlobalNetworkPolicyListCopy))
			return nil
		}, InType: reflect.TypeOf(&GlobalNetworkPolicyListCopy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IPPool).DeepCopyInto(out.(*IPPool))
			return nil
		}, InType: reflect.TypeOf(&IPPool{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IPPoolCopy).DeepCopyInto(out.(*IPPoolCopy))
			return nil
		}, InType: reflect.TypeOf(&IPPoolCopy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IPPoolList).DeepCopyInto(out.(*IPPoolList))
			return nil
		}, InType: reflect.TypeOf(&IPPoolList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IPPoolListCopy).DeepCopyInto(out.(*IPPoolListCopy))
			return nil
		}, InType: reflect.TypeOf(&IPPoolListCopy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IPPoolSpec).DeepCopyInto(out.(*IPPoolSpec))
			return nil
		}, InType: reflect.TypeOf(&IPPoolSpec{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeer) DeepCopyInto(out *BGPPeer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeer.
func (in *BGPPeer) DeepCopy() *BGPPeer {
	if in == nil {
		return nil
	}
	out := new(BGPPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeerCopy) DeepCopyInto(out *BGPPeerCopy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeerCopy.
func (in *BGPPeerCopy) DeepCopy() *BGPPeerCopy {
	if in == nil {
		return nil
	}
	out := new(BGPPeerCopy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeerList) DeepCopyInto(out *BGPPeerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Metadata = in.Metadata
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BGPPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeerList.
func (in *BGPPeerList) DeepCopy() *BGPPeerList {
	if in == nil {
		return nil
	}
	out := new(BGPPeerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeerListCopy) DeepCopyInto(out *BGPPeerListCopy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Metadata = in.Metadata
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BGPPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeerListCopy.
func (in *BGPPeerListCopy) DeepCopy() *BGPPeerListCopy {
	if in == nil {
		return nil
	}
	out := new(BGPPeerListCopy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeerSpec) DeepCopyInto(out *BGPPeerSpec) {
	*out = *in
	out.BGPPeerSpec = in.BGPPeerSpec
	in.PeerIP.DeepCopyInto(&out.PeerIP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeerSpec.
func (in *BGPPeerSpec) DeepCopy() *BGPPeerSpec {
	if in == nil {
		return nil
	}
	out := new(BGPPeerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalBGPConfig) DeepCopyInto(out *GlobalBGPConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalBGPConfig.
func (in *GlobalBGPConfig) DeepCopy() *GlobalBGPConfig {
	if in == nil {
		return nil
	}
	out := new(GlobalBGPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalBGPConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalBGPConfigList) DeepCopyInto(out *GlobalBGPConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalBGPConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalBGPConfigList.
func (in *GlobalBGPConfigList) DeepCopy() *GlobalBGPConfigList {
	if in == nil {
		return nil
	}
	out := new(GlobalBGPConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalBGPConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalBGPConfigSpec) DeepCopyInto(out *GlobalBGPConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalBGPConfigSpec.
func (in *GlobalBGPConfigSpec) DeepCopy() *GlobalBGPConfigSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalBGPConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalFelixConfig) DeepCopyInto(out *GlobalFelixConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Metadata.DeepCopyInto(&out.Metadata)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalFelixConfig.
func (in *GlobalFelixConfig) DeepCopy() *GlobalFelixConfig {
	if in == nil {
		return nil
	}
	out := new(GlobalFelixConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalFelixConfigCopy) DeepCopyInto(out *GlobalFelixConfigCopy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Metadata.DeepCopyInto(&out.Metadata)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalFelixConfigCopy.
func (in *GlobalFelixConfigCopy) DeepCopy() *GlobalFelixConfigCopy {
	if in == nil {
		return nil
	}
	out := new(GlobalFelixConfigCopy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalFelixConfigList) DeepCopyInto(out *GlobalFelixConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Metadata = in.Metadata
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalFelixConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalFelixConfigList.
func (in *GlobalFelixConfigList) DeepCopy() *GlobalFelixConfigList {
	if in == nil {
		return nil
	}
	out := new(GlobalFelixConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalFelixConfigListCopy) DeepCopyInto(out *GlobalFelixConfigListCopy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Metadata = in.Metadata
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalFelixConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalFelixConfigListCopy.
func (in *GlobalFelixConfigListCopy) DeepCopy() *GlobalFelixConfigListCopy {
	if in == nil {
		return nil
	}
	out := new(GlobalFelixConfigListCopy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalFelixConfigSpec) DeepCopyInto(out *GlobalFelixConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalFelixConfigSpec.
func (in *GlobalFelixConfigSpec) DeepCopy() *GlobalFelixConfigSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalFelixConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalNetworkPolicy) DeepCopyInto(out *GlobalNetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalNetworkPolicy.
func (in *GlobalNetworkPolicy) DeepCopy() *GlobalNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(GlobalNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalNetworkPolicyCopy) DeepCopyInto(out *GlobalNetworkPolicyCopy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalNetworkPolicyCopy.
func (in *GlobalNetworkPolicyCopy) DeepCopy() *GlobalNetworkPolicyCopy {
	if in == nil {
		return nil
	}
	out := new(GlobalNetworkPolicyCopy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalNetworkPolicyList) DeepCopyInto(out *GlobalNetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Metadata = in.Metadata
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalNetworkPolicyList.
func (in *GlobalNetworkPolicyList) DeepCopy() *GlobalNetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(GlobalNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalNetworkPolicyListCopy) DeepCopyInto(out *GlobalNetworkPolicyListCopy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Metadata = in.Metadata
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalNetworkPolicyListCopy.
func (in *GlobalNetworkPolicyListCopy) DeepCopy() *GlobalNetworkPolicyListCopy {
	if in == nil {
		return nil
	}
	out := new(GlobalNetworkPolicyListCopy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPool) DeepCopyInto(out *IPPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPool.
func (in *IPPool) DeepCopy() *IPPool {
	if in == nil {
		return nil
	}
	out := new(IPPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolCopy) DeepCopyInto(out *IPPoolCopy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolCopy.
func (in *IPPoolCopy) DeepCopy() *IPPoolCopy {
	if in == nil {
		return nil
	}
	out := new(IPPoolCopy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolList) DeepCopyInto(out *IPPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Metadata = in.Metadata
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolList.
func (in *IPPoolList) DeepCopy() *IPPoolList {
	if in == nil {
		return nil
	}
	out := new(IPPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolListCopy) DeepCopyInto(out *IPPoolListCopy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Metadata = in.Metadata
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolListCopy.
func (in *IPPoolListCopy) DeepCopy() *IPPoolListCopy {
	if in == nil {
		return nil
	}
	out := new(IPPoolListCopy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolSpec) DeepCopyInto(out *IPPoolSpec) {
	*out = *in
	in.IPPoolSpec.DeepCopyInto(&out.IPPoolSpec)
	in.CIDR.DeepCopyInto(&out.CIDR)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolSpec.
func (in *IPPoolSpec) DeepCopy() *IPPoolSpec {
	if in == nil {
		return nil
	}
	out := new(IPPoolSpec)
	in.DeepCopyInto(out)
	return out
}
