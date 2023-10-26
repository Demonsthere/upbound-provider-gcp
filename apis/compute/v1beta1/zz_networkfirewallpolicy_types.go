// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NetworkFirewallPolicyInitParameters struct {

	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type NetworkFirewallPolicyObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Fingerprint of the resource. This field is used internally during updates of this resource.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/firewallPolicies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier for the resource. This identifier is defined by the server.
	NetworkFirewallPolicyID *string `json:"networkFirewallPolicyId,omitempty" tf:"network_firewall_policy_id,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount *float64 `json:"ruleTupleCount,omitempty" tf:"rule_tuple_count,omitempty"`

	// Server-defined URL for the resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Server-defined URL for this resource with the resource id.
	SelfLinkWithID *string `json:"selfLinkWithId,omitempty" tf:"self_link_with_id,omitempty"`
}

type NetworkFirewallPolicyParameters struct {

	// An optional description of this resource. Provide this property when you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// NetworkFirewallPolicySpec defines the desired state of NetworkFirewallPolicy
type NetworkFirewallPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkFirewallPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider NetworkFirewallPolicyInitParameters `json:"initProvider,omitempty"`
}

// NetworkFirewallPolicyStatus defines the observed state of NetworkFirewallPolicy.
type NetworkFirewallPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkFirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkFirewallPolicy is the Schema for the NetworkFirewallPolicys API. The Compute NetworkFirewallPolicy resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NetworkFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkFirewallPolicySpec   `json:"spec"`
	Status            NetworkFirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkFirewallPolicyList contains a list of NetworkFirewallPolicys
type NetworkFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkFirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	NetworkFirewallPolicy_Kind             = "NetworkFirewallPolicy"
	NetworkFirewallPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkFirewallPolicy_Kind}.String()
	NetworkFirewallPolicy_KindAPIVersion   = NetworkFirewallPolicy_Kind + "." + CRDGroupVersion.String()
	NetworkFirewallPolicy_GroupVersionKind = CRDGroupVersion.WithKind(NetworkFirewallPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkFirewallPolicy{}, &NetworkFirewallPolicyList{})
}
