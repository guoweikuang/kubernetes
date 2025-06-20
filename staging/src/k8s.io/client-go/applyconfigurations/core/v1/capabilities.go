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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// CapabilitiesApplyConfiguration represents a declarative configuration of the Capabilities type for use
// with apply.
type CapabilitiesApplyConfiguration struct {
	Add  []corev1.Capability `json:"add,omitempty"`
	Drop []corev1.Capability `json:"drop,omitempty"`
}

// CapabilitiesApplyConfiguration constructs a declarative configuration of the Capabilities type for use with
// apply.
func Capabilities() *CapabilitiesApplyConfiguration {
	return &CapabilitiesApplyConfiguration{}
}
func (b CapabilitiesApplyConfiguration) IsApplyConfiguration() {}

// WithAdd adds the given value to the Add field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Add field.
func (b *CapabilitiesApplyConfiguration) WithAdd(values ...corev1.Capability) *CapabilitiesApplyConfiguration {
	for i := range values {
		b.Add = append(b.Add, values[i])
	}
	return b
}

// WithDrop adds the given value to the Drop field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Drop field.
func (b *CapabilitiesApplyConfiguration) WithDrop(values ...corev1.Capability) *CapabilitiesApplyConfiguration {
	for i := range values {
		b.Drop = append(b.Drop, values[i])
	}
	return b
}
