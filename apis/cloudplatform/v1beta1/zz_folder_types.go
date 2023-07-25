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

type FolderInitParameters struct {

	// The folder’s display name.
	// A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`
}

type FolderObservation struct {

	// Timestamp when the Folder was created. Assigned by the server.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The folder’s display name.
	// A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The lifecycle state of the folder such as ACTIVE or DELETE_REQUESTED.
	LifecycleState *string `json:"lifecycleState,omitempty" tf:"lifecycle_state,omitempty"`

	// The resource name of the Folder. Its format is folders/{folder_id}.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The resource name of the parent Folder or Organization.
	// Must be of the form folders/{folder_id} or organizations/{org_id}.
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`
}

type FolderParameters struct {

	// The folder’s display name.
	// A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The resource name of the parent Folder or Organization.
	// Must be of the form folders/{folder_id} or organizations/{org_id}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Folder
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Reference to a Folder in cloudplatform to populate parent.
	// +kubebuilder:validation:Optional
	ParentRef *v1.Reference `json:"parentRef,omitempty" tf:"-"`

	// Selector for a Folder in cloudplatform to populate parent.
	// +kubebuilder:validation:Optional
	ParentSelector *v1.Selector `json:"parentSelector,omitempty" tf:"-"`
}

// FolderSpec defines the desired state of Folder
type FolderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider FolderInitParameters `json:"initProvider,omitempty"`
}

// FolderStatus defines the observed state of Folder.
type FolderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Folder is the Schema for the Folders API. Allows management of a Google Cloud Platform folder.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Folder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || has(self.initProvider.displayName)",message="displayName is a required parameter"
	Spec   FolderSpec   `json:"spec"`
	Status FolderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderList contains a list of Folders
type FolderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Folder `json:"items"`
}

// Repository type metadata.
var (
	Folder_Kind             = "Folder"
	Folder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Folder_Kind}.String()
	Folder_KindAPIVersion   = Folder_Kind + "." + CRDGroupVersion.String()
	Folder_GroupVersionKind = CRDGroupVersion.WithKind(Folder_Kind)
)

func init() {
	SchemeBuilder.Register(&Folder{}, &FolderList{})
}
