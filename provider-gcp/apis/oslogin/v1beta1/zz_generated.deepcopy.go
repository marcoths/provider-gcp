//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHPublicKey) DeepCopyInto(out *SSHPublicKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHPublicKey.
func (in *SSHPublicKey) DeepCopy() *SSHPublicKey {
	if in == nil {
		return nil
	}
	out := new(SSHPublicKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHPublicKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHPublicKeyList) DeepCopyInto(out *SSHPublicKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SSHPublicKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHPublicKeyList.
func (in *SSHPublicKeyList) DeepCopy() *SSHPublicKeyList {
	if in == nil {
		return nil
	}
	out := new(SSHPublicKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHPublicKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHPublicKeyObservation) DeepCopyInto(out *SSHPublicKeyObservation) {
	*out = *in
	if in.Fingerprint != nil {
		in, out := &in.Fingerprint, &out.Fingerprint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHPublicKeyObservation.
func (in *SSHPublicKeyObservation) DeepCopy() *SSHPublicKeyObservation {
	if in == nil {
		return nil
	}
	out := new(SSHPublicKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHPublicKeyParameters) DeepCopyInto(out *SSHPublicKeyParameters) {
	*out = *in
	if in.ExpirationTimeUsec != nil {
		in, out := &in.ExpirationTimeUsec, &out.ExpirationTimeUsec
		*out = new(string)
		**out = **in
	}
	out.KeySecretRef = in.KeySecretRef
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHPublicKeyParameters.
func (in *SSHPublicKeyParameters) DeepCopy() *SSHPublicKeyParameters {
	if in == nil {
		return nil
	}
	out := new(SSHPublicKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHPublicKeySpec) DeepCopyInto(out *SSHPublicKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHPublicKeySpec.
func (in *SSHPublicKeySpec) DeepCopy() *SSHPublicKeySpec {
	if in == nil {
		return nil
	}
	out := new(SSHPublicKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHPublicKeyStatus) DeepCopyInto(out *SSHPublicKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHPublicKeyStatus.
func (in *SSHPublicKeyStatus) DeepCopy() *SSHPublicKeyStatus {
	if in == nil {
		return nil
	}
	out := new(SSHPublicKeyStatus)
	in.DeepCopyInto(out)
	return out
}
