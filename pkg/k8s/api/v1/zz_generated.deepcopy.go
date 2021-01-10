// +build !ignore_autogenerated

/*
Copyright 2020-2021 AccuKnox.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapabilitiesType) DeepCopyInto(out *CapabilitiesType) {
	*out = *in
	if in.MatchCapabilities != nil {
		in, out := &in.MatchCapabilities, &out.MatchCapabilities
		*out = make([]MatchCapabilitiesType, len(*in))
		copy(*out, *in)
	}
	if in.MatchOperations != nil {
		in, out := &in.MatchOperations, &out.MatchOperations
		*out = make([]MatchOperationsType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapabilitiesType.
func (in *CapabilitiesType) DeepCopy() *CapabilitiesType {
	if in == nil {
		return nil
	}
	out := new(CapabilitiesType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileDirectoryType) DeepCopyInto(out *FileDirectoryType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileDirectoryType.
func (in *FileDirectoryType) DeepCopy() *FileDirectoryType {
	if in == nil {
		return nil
	}
	out := new(FileDirectoryType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilePathType) DeepCopyInto(out *FilePathType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilePathType.
func (in *FilePathType) DeepCopy() *FilePathType {
	if in == nil {
		return nil
	}
	out := new(FilePathType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilePatternType) DeepCopyInto(out *FilePatternType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilePatternType.
func (in *FilePatternType) DeepCopy() *FilePatternType {
	if in == nil {
		return nil
	}
	out := new(FilePatternType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileType) DeepCopyInto(out *FileType) {
	*out = *in
	if in.MatchPaths != nil {
		in, out := &in.MatchPaths, &out.MatchPaths
		*out = make([]FilePathType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchDirectories != nil {
		in, out := &in.MatchDirectories, &out.MatchDirectories
		*out = make([]FileDirectoryType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchPatterns != nil {
		in, out := &in.MatchPatterns, &out.MatchPatterns
		*out = make([]FilePatternType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileType.
func (in *FileType) DeepCopy() *FileType {
	if in == nil {
		return nil
	}
	out := new(FileType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorPolicy) DeepCopyInto(out *KubeArmorPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorPolicy.
func (in *KubeArmorPolicy) DeepCopy() *KubeArmorPolicy {
	if in == nil {
		return nil
	}
	out := new(KubeArmorPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeArmorPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorPolicyList) DeepCopyInto(out *KubeArmorPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeArmorPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorPolicyList.
func (in *KubeArmorPolicyList) DeepCopy() *KubeArmorPolicyList {
	if in == nil {
		return nil
	}
	out := new(KubeArmorPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeArmorPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorPolicySpec) DeepCopyInto(out *KubeArmorPolicySpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.Process.DeepCopyInto(&out.Process)
	in.File.DeepCopyInto(&out.File)
	in.Capabilities.DeepCopyInto(&out.Capabilities)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorPolicySpec.
func (in *KubeArmorPolicySpec) DeepCopy() *KubeArmorPolicySpec {
	if in == nil {
		return nil
	}
	out := new(KubeArmorPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorPolicyStatus) DeepCopyInto(out *KubeArmorPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorPolicyStatus.
func (in *KubeArmorPolicyStatus) DeepCopy() *KubeArmorPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(KubeArmorPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchSourceType) DeepCopyInto(out *MatchSourceType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchSourceType.
func (in *MatchSourceType) DeepCopy() *MatchSourceType {
	if in == nil {
		return nil
	}
	out := new(MatchSourceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessDirectoryType) DeepCopyInto(out *ProcessDirectoryType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessDirectoryType.
func (in *ProcessDirectoryType) DeepCopy() *ProcessDirectoryType {
	if in == nil {
		return nil
	}
	out := new(ProcessDirectoryType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessPathType) DeepCopyInto(out *ProcessPathType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessPathType.
func (in *ProcessPathType) DeepCopy() *ProcessPathType {
	if in == nil {
		return nil
	}
	out := new(ProcessPathType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessPatternType) DeepCopyInto(out *ProcessPatternType) {
	*out = *in
	if in.FromSource != nil {
		in, out := &in.FromSource, &out.FromSource
		*out = make([]MatchSourceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessPatternType.
func (in *ProcessPatternType) DeepCopy() *ProcessPatternType {
	if in == nil {
		return nil
	}
	out := new(ProcessPatternType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessType) DeepCopyInto(out *ProcessType) {
	*out = *in
	if in.MatchPaths != nil {
		in, out := &in.MatchPaths, &out.MatchPaths
		*out = make([]ProcessPathType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchDirectories != nil {
		in, out := &in.MatchDirectories, &out.MatchDirectories
		*out = make([]ProcessDirectoryType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchPatterns != nil {
		in, out := &in.MatchPatterns, &out.MatchPatterns
		*out = make([]ProcessPatternType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessType.
func (in *ProcessType) DeepCopy() *ProcessType {
	if in == nil {
		return nil
	}
	out := new(ProcessType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectorType) DeepCopyInto(out *SelectorType) {
	*out = *in
	if in.MatchNames != nil {
		in, out := &in.MatchNames, &out.MatchNames
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectorType.
func (in *SelectorType) DeepCopy() *SelectorType {
	if in == nil {
		return nil
	}
	out := new(SelectorType)
	in.DeepCopyInto(out)
	return out
}
