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

type SharedVPCHostProjectInitParameters struct {
}

type SharedVPCHostProjectObservation struct {

	// an identifier for the resource with format {{project}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project that will serve as a Shared VPC host project
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type SharedVPCHostProjectParameters struct {

	// The ID of the project that will serve as a Shared VPC host project
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractProjectID()
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`
}

// SharedVPCHostProjectSpec defines the desired state of SharedVPCHostProject
type SharedVPCHostProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SharedVPCHostProjectParameters `json:"forProvider"`
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
	InitProvider SharedVPCHostProjectInitParameters `json:"initProvider,omitempty"`
}

// SharedVPCHostProjectStatus defines the observed state of SharedVPCHostProject.
type SharedVPCHostProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SharedVPCHostProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SharedVPCHostProject is the Schema for the SharedVPCHostProjects API. Enables the Google Compute Engine Shared VPC feature for a project, assigning it as a host project.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type SharedVPCHostProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SharedVPCHostProjectSpec   `json:"spec"`
	Status            SharedVPCHostProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedVPCHostProjectList contains a list of SharedVPCHostProjects
type SharedVPCHostProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedVPCHostProject `json:"items"`
}

// Repository type metadata.
var (
	SharedVPCHostProject_Kind             = "SharedVPCHostProject"
	SharedVPCHostProject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SharedVPCHostProject_Kind}.String()
	SharedVPCHostProject_KindAPIVersion   = SharedVPCHostProject_Kind + "." + CRDGroupVersion.String()
	SharedVPCHostProject_GroupVersionKind = CRDGroupVersion.WithKind(SharedVPCHostProject_Kind)
)

func init() {
	SchemeBuilder.Register(&SharedVPCHostProject{}, &SharedVPCHostProjectList{})
}
