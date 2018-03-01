// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package custom

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
			in.(*GlobalFelixConfigList).DeepCopyInto(out.(*GlobalFelixConfigList))
			return nil
		}, InType: reflect.TypeOf(&GlobalFelixConfigList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GlobalFelixConfigSpec).DeepCopyInto(out.(*GlobalFelixConfigSpec))
			return nil
		}, InType: reflect.TypeOf(&GlobalFelixConfigSpec{})},
	)
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
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
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

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalFelixConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalFelixConfigList) DeepCopyInto(out *GlobalFelixConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
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

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalFelixConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
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