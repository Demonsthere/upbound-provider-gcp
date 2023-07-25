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

type DatasetIAMBindingConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type DatasetIAMBindingConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type DatasetIAMBindingConditionParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type DatasetIAMBindingInitParameters struct {
	Condition []DatasetIAMBindingConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type DatasetIAMBindingObservation struct {
	Condition []DatasetIAMBindingConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type DatasetIAMBindingParameters struct {
	Condition []DatasetIAMBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	// +kubebuilder:validation:Optional
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// Reference to a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDRef *v1.Reference `json:"datasetIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDSelector *v1.Selector `json:"datasetIdSelector,omitempty" tf:"-"`

	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// DatasetIAMBindingSpec defines the desired state of DatasetIAMBinding
type DatasetIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasetIAMBindingParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DatasetIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// DatasetIAMBindingStatus defines the observed state of DatasetIAMBinding.
type DatasetIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasetIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetIAMBinding is the Schema for the DatasetIAMBindings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DatasetIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.members) || has(self.initProvider.members)",message="members is a required parameter"
	Spec   DatasetIAMBindingSpec   `json:"spec"`
	Status DatasetIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetIAMBindingList contains a list of DatasetIAMBindings
type DatasetIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasetIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	DatasetIAMBinding_Kind             = "DatasetIAMBinding"
	DatasetIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatasetIAMBinding_Kind}.String()
	DatasetIAMBinding_KindAPIVersion   = DatasetIAMBinding_Kind + "." + CRDGroupVersion.String()
	DatasetIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(DatasetIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&DatasetIAMBinding{}, &DatasetIAMBindingList{})
}
