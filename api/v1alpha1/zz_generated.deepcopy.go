//go:build !ignore_autogenerated

/*
Copyright 2024.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Discovery) DeepCopyInto(out *Discovery) {
	*out = *in
	out.S3 = in.S3
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Discovery.
func (in *Discovery) DeepCopy() *Discovery {
	if in == nil {
		return nil
	}
	out := new(Discovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IRSA) DeepCopyInto(out *IRSA) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IRSA.
func (in *IRSA) DeepCopy() *IRSA {
	if in == nil {
		return nil
	}
	out := new(IRSA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IRSA) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IRSAList) DeepCopyInto(out *IRSAList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IRSA, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IRSAList.
func (in *IRSAList) DeepCopy() *IRSAList {
	if in == nil {
		return nil
	}
	out := new(IRSAList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IRSAList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IRSANamespacedNameWithTags) DeepCopyInto(out *IRSANamespacedNameWithTags) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IRSANamespacedNameWithTags.
func (in *IRSANamespacedNameWithTags) DeepCopy() *IRSANamespacedNameWithTags {
	if in == nil {
		return nil
	}
	out := new(IRSANamespacedNameWithTags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IRSAServiceAccount) DeepCopyInto(out *IRSAServiceAccount) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IRSAServiceAccount.
func (in *IRSAServiceAccount) DeepCopy() *IRSAServiceAccount {
	if in == nil {
		return nil
	}
	out := new(IRSAServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IRSASetup) DeepCopyInto(out *IRSASetup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IRSASetup.
func (in *IRSASetup) DeepCopy() *IRSASetup {
	if in == nil {
		return nil
	}
	out := new(IRSASetup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IRSASetup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IRSASetupList) DeepCopyInto(out *IRSASetupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IRSASetup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IRSASetupList.
func (in *IRSASetupList) DeepCopy() *IRSASetupList {
	if in == nil {
		return nil
	}
	out := new(IRSASetupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IRSASetupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IRSASetupSpec) DeepCopyInto(out *IRSASetupSpec) {
	*out = *in
	out.Discovery = in.Discovery
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IRSASetupSpec.
func (in *IRSASetupSpec) DeepCopy() *IRSASetupSpec {
	if in == nil {
		return nil
	}
	out := new(IRSASetupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IRSASetupStatus) DeepCopyInto(out *IRSASetupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IRSASetupStatus.
func (in *IRSASetupStatus) DeepCopy() *IRSASetupStatus {
	if in == nil {
		return nil
	}
	out := new(IRSASetupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IRSASpec) DeepCopyInto(out *IRSASpec) {
	*out = *in
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
	out.IamRole = in.IamRole
	if in.IamPolicies != nil {
		in, out := &in.IamPolicies, &out.IamPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IRSASpec.
func (in *IRSASpec) DeepCopy() *IRSASpec {
	if in == nil {
		return nil
	}
	out := new(IRSASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IRSAStatus) DeepCopyInto(out *IRSAStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccounts != nil {
		in, out := &in.ServiceAccounts, &out.ServiceAccounts
		*out = make(StatusServiceAccountList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IRSAStatus.
func (in *IRSAStatus) DeepCopy() *IRSAStatus {
	if in == nil {
		return nil
	}
	out := new(IRSAStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamRole) DeepCopyInto(out *IamRole) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamRole.
func (in *IamRole) DeepCopy() *IamRole {
	if in == nil {
		return nil
	}
	out := new(IamRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Discovery) DeepCopyInto(out *S3Discovery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Discovery.
func (in *S3Discovery) DeepCopy() *S3Discovery {
	if in == nil {
		return nil
	}
	out := new(S3Discovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in StatusServiceAccountList) DeepCopyInto(out *StatusServiceAccountList) {
	{
		in := &in
		*out = make(StatusServiceAccountList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusServiceAccountList.
func (in StatusServiceAccountList) DeepCopy() StatusServiceAccountList {
	if in == nil {
		return nil
	}
	out := new(StatusServiceAccountList)
	in.DeepCopyInto(out)
	return *out
}
