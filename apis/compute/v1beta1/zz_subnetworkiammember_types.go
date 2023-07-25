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

type SubnetworkIAMMemberConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SubnetworkIAMMemberConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SubnetworkIAMMemberConditionParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SubnetworkIAMMemberInitParameters struct {
	Condition []SubnetworkIAMMemberConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type SubnetworkIAMMemberObservation struct {
	Condition []SubnetworkIAMMemberConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`
}

type SubnetworkIAMMemberParameters struct {
	Condition []SubnetworkIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +crossplane:generate:reference:type=Subnetwork
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Reference to a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// Selector for a Subnetwork to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`
}

// SubnetworkIAMMemberSpec defines the desired state of SubnetworkIAMMember
type SubnetworkIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetworkIAMMemberParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SubnetworkIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// SubnetworkIAMMemberStatus defines the observed state of SubnetworkIAMMember.
type SubnetworkIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetworkIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetworkIAMMember is the Schema for the SubnetworkIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type SubnetworkIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.member) || has(self.initProvider.member)",message="member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || has(self.initProvider.region)",message="region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || has(self.initProvider.role)",message="role is a required parameter"
	Spec   SubnetworkIAMMemberSpec   `json:"spec"`
	Status SubnetworkIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetworkIAMMemberList contains a list of SubnetworkIAMMembers
type SubnetworkIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetworkIAMMember `json:"items"`
}

// Repository type metadata.
var (
	SubnetworkIAMMember_Kind             = "SubnetworkIAMMember"
	SubnetworkIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetworkIAMMember_Kind}.String()
	SubnetworkIAMMember_KindAPIVersion   = SubnetworkIAMMember_Kind + "." + CRDGroupVersion.String()
	SubnetworkIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(SubnetworkIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetworkIAMMember{}, &SubnetworkIAMMemberList{})
}
