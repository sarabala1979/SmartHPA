//go:build !ignore_autogenerated

/*
Copyright 2025.

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
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurrentTrigger) DeepCopyInto(out *CurrentTrigger) {
	*out = *in
	if in.HPAConfig != nil {
		in, out := &in.HPAConfig, &out.HPAConfig
		*out = new(HPAConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurrentTrigger.
func (in *CurrentTrigger) DeepCopy() *CurrentTrigger {
	if in == nil {
		return nil
	}
	out := new(CurrentTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CurrentTriggers) DeepCopyInto(out *CurrentTriggers) {
	{
		in := &in
		*out = make(CurrentTriggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurrentTriggers.
func (in CurrentTriggers) DeepCopy() CurrentTriggers {
	if in == nil {
		return nil
	}
	out := new(CurrentTriggers)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPAConfig) DeepCopyInto(out *HPAConfig) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	if in.DesiredReplicas != nil {
		in, out := &in.DesiredReplicas, &out.DesiredReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPAConfig.
func (in *HPAConfig) DeepCopy() *HPAConfig {
	if in == nil {
		return nil
	}
	out := new(HPAConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPAObjectReference) DeepCopyInto(out *HPAObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPAObjectReference.
func (in *HPAObjectReference) DeepCopy() *HPAObjectReference {
	if in == nil {
		return nil
	}
	out := new(HPAObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPASpecTemplate) DeepCopyInto(out *HPASpecTemplate) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(HPAConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPASpecTemplate.
func (in *HPASpecTemplate) DeepCopy() *HPASpecTemplate {
	if in == nil {
		return nil
	}
	out := new(HPASpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Interval) DeepCopyInto(out *Interval) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Interval.
func (in *Interval) DeepCopy() *Interval {
	if in == nil {
		return nil
	}
	out := new(Interval)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartHorizontalPodAutoscaler) DeepCopyInto(out *SmartHorizontalPodAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartHorizontalPodAutoscaler.
func (in *SmartHorizontalPodAutoscaler) DeepCopy() *SmartHorizontalPodAutoscaler {
	if in == nil {
		return nil
	}
	out := new(SmartHorizontalPodAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SmartHorizontalPodAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartHorizontalPodAutoscalerList) DeepCopyInto(out *SmartHorizontalPodAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SmartHorizontalPodAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartHorizontalPodAutoscalerList.
func (in *SmartHorizontalPodAutoscalerList) DeepCopy() *SmartHorizontalPodAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(SmartHorizontalPodAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SmartHorizontalPodAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartHorizontalPodAutoscalerSpec) DeepCopyInto(out *SmartHorizontalPodAutoscalerSpec) {
	*out = *in
	if in.SmartRecommendation != nil {
		in, out := &in.SmartRecommendation, &out.SmartRecommendation
		*out = new(SmartRecommendation)
		(*in).DeepCopyInto(*out)
	}
	if in.HPAObjectRef != nil {
		in, out := &in.HPAObjectRef, &out.HPAObjectRef
		*out = new(HPAObjectReference)
		**out = **in
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make(Triggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartHorizontalPodAutoscalerSpec.
func (in *SmartHorizontalPodAutoscalerSpec) DeepCopy() *SmartHorizontalPodAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(SmartHorizontalPodAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartHorizontalPodAutoscalerStatus) DeepCopyInto(out *SmartHorizontalPodAutoscalerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HPAObjectRef != nil {
		in, out := &in.HPAObjectRef, &out.HPAObjectRef
		*out = new(HPAObjectReference)
		**out = **in
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make(Triggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartHorizontalPodAutoscalerStatus.
func (in *SmartHorizontalPodAutoscalerStatus) DeepCopy() *SmartHorizontalPodAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(SmartHorizontalPodAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartRecommendation) DeepCopyInto(out *SmartRecommendation) {
	*out = *in
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartRecommendation.
func (in *SmartRecommendation) DeepCopy() *SmartRecommendation {
	if in == nil {
		return nil
	}
	out := new(SmartRecommendation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(Interval)
		**out = **in
	}
	if in.StartHPAConfig != nil {
		in, out := &in.StartHPAConfig, &out.StartHPAConfig
		*out = new(HPAConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EndHPAConfig != nil {
		in, out := &in.EndHPAConfig, &out.EndHPAConfig
		*out = new(HPAConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Triggers) DeepCopyInto(out *Triggers) {
	{
		in := &in
		*out = make(Triggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Triggers.
func (in Triggers) DeepCopy() Triggers {
	if in == nil {
		return nil
	}
	out := new(Triggers)
	in.DeepCopyInto(out)
	return *out
}
