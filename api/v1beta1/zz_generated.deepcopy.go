//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 Teppei Sudo.

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
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudInit) DeepCopyInto(out *CloudInit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudInit.
func (in *CloudInit) DeepCopy() *CloudInit {
	if in == nil {
		return nil
	}
	out := new(CloudInit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxCluster) DeepCopyInto(out *ProxmoxCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxCluster.
func (in *ProxmoxCluster) DeepCopy() *ProxmoxCluster {
	if in == nil {
		return nil
	}
	out := new(ProxmoxCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxClusterList) DeepCopyInto(out *ProxmoxClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProxmoxCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxClusterList.
func (in *ProxmoxClusterList) DeepCopy() *ProxmoxClusterList {
	if in == nil {
		return nil
	}
	out := new(ProxmoxClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxClusterSpec) DeepCopyInto(out *ProxmoxClusterSpec) {
	*out = *in
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	if in.CredentialsRef != nil {
		in, out := &in.CredentialsRef, &out.CredentialsRef
		*out = new(ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxClusterSpec.
func (in *ProxmoxClusterSpec) DeepCopy() *ProxmoxClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ProxmoxClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxClusterStatus) DeepCopyInto(out *ProxmoxClusterStatus) {
	*out = *in
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(apiv1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxClusterStatus.
func (in *ProxmoxClusterStatus) DeepCopy() *ProxmoxClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ProxmoxClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachine) DeepCopyInto(out *ProxmoxMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachine.
func (in *ProxmoxMachine) DeepCopy() *ProxmoxMachine {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachineList) DeepCopyInto(out *ProxmoxMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProxmoxMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachineList.
func (in *ProxmoxMachineList) DeepCopy() *ProxmoxMachineList {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxmoxMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachineSpec) DeepCopyInto(out *ProxmoxMachineSpec) {
	*out = *in
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	out.CloudInit = in.CloudInit
	if in.FailureDomain != nil {
		in, out := &in.FailureDomain, &out.FailureDomain
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachineSpec.
func (in *ProxmoxMachineSpec) DeepCopy() *ProxmoxMachineSpec {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxmoxMachineStatus) DeepCopyInto(out *ProxmoxMachineStatus) {
	*out = *in
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]apiv1beta1.MachineAddress, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InstanceStatus != nil {
		in, out := &in.InstanceStatus, &out.InstanceStatus
		*out = new(InstanceStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxmoxMachineStatus.
func (in *ProxmoxMachineStatus) DeepCopy() *ProxmoxMachineStatus {
	if in == nil {
		return nil
	}
	out := new(ProxmoxMachineStatus)
	in.DeepCopyInto(out)
	return out
}