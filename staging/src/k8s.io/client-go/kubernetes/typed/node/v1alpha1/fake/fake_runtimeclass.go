/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "k8s.io/api/node/v1alpha1"
	nodev1alpha1 "k8s.io/client-go/applyconfigurations/node/v1alpha1"
	gentype "k8s.io/client-go/gentype"
	typednodev1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
)

// fakeRuntimeClasses implements RuntimeClassInterface
type fakeRuntimeClasses struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.RuntimeClass, *v1alpha1.RuntimeClassList, *nodev1alpha1.RuntimeClassApplyConfiguration]
	Fake *FakeNodeV1alpha1
}

func newFakeRuntimeClasses(fake *FakeNodeV1alpha1) typednodev1alpha1.RuntimeClassInterface {
	return &fakeRuntimeClasses{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.RuntimeClass, *v1alpha1.RuntimeClassList, *nodev1alpha1.RuntimeClassApplyConfiguration](
			fake.Fake,
			"",
			v1alpha1.SchemeGroupVersion.WithResource("runtimeclasses"),
			v1alpha1.SchemeGroupVersion.WithKind("RuntimeClass"),
			func() *v1alpha1.RuntimeClass { return &v1alpha1.RuntimeClass{} },
			func() *v1alpha1.RuntimeClassList { return &v1alpha1.RuntimeClassList{} },
			func(dst, src *v1alpha1.RuntimeClassList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.RuntimeClassList) []*v1alpha1.RuntimeClass {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.RuntimeClassList, items []*v1alpha1.RuntimeClass) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
