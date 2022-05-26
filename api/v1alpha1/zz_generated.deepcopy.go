//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
func (in *AgentNode) DeepCopyInto(out *AgentNode) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentNode.
func (in *AgentNode) DeepCopy() *AgentNode {
	if in == nil {
		return nil
	}
	out := new(AgentNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionalStatus) DeepCopyInto(out *ConditionalStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionalStatus.
func (in *ConditionalStatus) DeepCopy() *ConditionalStatus {
	if in == nil {
		return nil
	}
	out := new(ConditionalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservability) DeepCopyInto(out *NodeObservability) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservability.
func (in *NodeObservability) DeepCopy() *NodeObservability {
	if in == nil {
		return nil
	}
	out := new(NodeObservability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeObservability) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityDebug) DeepCopyInto(out *NodeObservabilityDebug) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityDebug.
func (in *NodeObservabilityDebug) DeepCopy() *NodeObservabilityDebug {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityDebug)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityList) DeepCopyInto(out *NodeObservabilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeObservability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityList.
func (in *NodeObservabilityList) DeepCopy() *NodeObservabilityList {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeObservabilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityMachineConfig) DeepCopyInto(out *NodeObservabilityMachineConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityMachineConfig.
func (in *NodeObservabilityMachineConfig) DeepCopy() *NodeObservabilityMachineConfig {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityMachineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeObservabilityMachineConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityMachineConfigList) DeepCopyInto(out *NodeObservabilityMachineConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeObservabilityMachineConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityMachineConfigList.
func (in *NodeObservabilityMachineConfigList) DeepCopy() *NodeObservabilityMachineConfigList {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityMachineConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeObservabilityMachineConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityMachineConfigSpec) DeepCopyInto(out *NodeObservabilityMachineConfigSpec) {
	*out = *in
	out.Debug = in.Debug
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityMachineConfigSpec.
func (in *NodeObservabilityMachineConfigSpec) DeepCopy() *NodeObservabilityMachineConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityMachineConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityMachineConfigStatus) DeepCopyInto(out *NodeObservabilityMachineConfigStatus) {
	*out = *in
	in.ConditionalStatus.DeepCopyInto(&out.ConditionalStatus)
	in.LastReconcile.DeepCopyInto(&out.LastReconcile)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityMachineConfigStatus.
func (in *NodeObservabilityMachineConfigStatus) DeepCopy() *NodeObservabilityMachineConfigStatus {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityMachineConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityRef) DeepCopyInto(out *NodeObservabilityRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityRef.
func (in *NodeObservabilityRef) DeepCopy() *NodeObservabilityRef {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityRun) DeepCopyInto(out *NodeObservabilityRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityRun.
func (in *NodeObservabilityRun) DeepCopy() *NodeObservabilityRun {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeObservabilityRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityRunList) DeepCopyInto(out *NodeObservabilityRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeObservabilityRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityRunList.
func (in *NodeObservabilityRunList) DeepCopy() *NodeObservabilityRunList {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeObservabilityRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityRunSpec) DeepCopyInto(out *NodeObservabilityRunSpec) {
	*out = *in
	if in.NodeObservabilityRef != nil {
		in, out := &in.NodeObservabilityRef, &out.NodeObservabilityRef
		*out = new(NodeObservabilityRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityRunSpec.
func (in *NodeObservabilityRunSpec) DeepCopy() *NodeObservabilityRunSpec {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityRunStatus) DeepCopyInto(out *NodeObservabilityRunStatus) {
	*out = *in
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.FinishedTimestamp != nil {
		in, out := &in.FinishedTimestamp, &out.FinishedTimestamp
		*out = (*in).DeepCopy()
	}
	if in.Agents != nil {
		in, out := &in.Agents, &out.Agents
		*out = make([]AgentNode, len(*in))
		copy(*out, *in)
	}
	if in.FailedAgents != nil {
		in, out := &in.FailedAgents, &out.FailedAgents
		*out = make([]AgentNode, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityRunStatus.
func (in *NodeObservabilityRunStatus) DeepCopy() *NodeObservabilityRunStatus {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilitySpec) DeepCopyInto(out *NodeObservabilitySpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilitySpec.
func (in *NodeObservabilitySpec) DeepCopy() *NodeObservabilitySpec {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservabilityStatus) DeepCopyInto(out *NodeObservabilityStatus) {
	*out = *in
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservabilityStatus.
func (in *NodeObservabilityStatus) DeepCopy() *NodeObservabilityStatus {
	if in == nil {
		return nil
	}
	out := new(NodeObservabilityStatus)
	in.DeepCopyInto(out)
	return out
}
