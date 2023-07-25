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

type AssetStatusInitParameters struct {
}

type AssetStatusObservation struct {
	ActiveAssets *float64 `json:"activeAssets,omitempty" tf:"active_assets,omitempty"`

	SecurityPolicyApplyingAssets *float64 `json:"securityPolicyApplyingAssets,omitempty" tf:"security_policy_applying_assets,omitempty"`

	// Output only. The time when the lake was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type AssetStatusParameters struct {
}

type LakeInitParameters struct {

	// Optional. Description of the lake.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Optional. User friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Optional. User-defined labels for the lake.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Optional. Settings to manage lake and Dataproc Metastore service instance association.
	Metastore []MetastoreInitParameters `json:"metastore,omitempty" tf:"metastore,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type LakeObservation struct {

	// Output only. Aggregated status of the underlying assets of the lake.
	AssetStatus []AssetStatusObservation `json:"assetStatus,omitempty" tf:"asset_status,omitempty"`

	// Output only. The time when the lake was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Optional. Description of the lake.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Optional. User friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/lakes/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional. User-defined labels for the lake.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Optional. Settings to manage lake and Dataproc Metastore service instance association.
	Metastore []MetastoreObservation `json:"metastore,omitempty" tf:"metastore,omitempty"`

	// Output only. Metastore status of the lake.
	MetastoreStatus []MetastoreStatusObservation `json:"metastoreStatus,omitempty" tf:"metastore_status,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Output only. Service account associated with this lake. This service account must be authorized to access or operate on resources managed by the lake.
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// Output only. Current state of the lake. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Output only. System generated globally unique ID for the lake. This ID will be different if the lake is deleted and re-created with the same name.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// Output only. The time when the lake was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type LakeParameters struct {

	// Optional. Description of the lake.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Optional. User friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Optional. User-defined labels for the lake.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Optional. Settings to manage lake and Dataproc Metastore service instance association.
	Metastore []MetastoreParameters `json:"metastore,omitempty" tf:"metastore,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type MetastoreInitParameters struct {

	// Optional. A relative reference to the Dataproc Metastore (https://cloud.google.com/dataproc-metastore/docs) service associated with the lake: projects/{project_id}/locations/{location_id}/services/{service_id}
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type MetastoreObservation struct {

	// Optional. A relative reference to the Dataproc Metastore (https://cloud.google.com/dataproc-metastore/docs) service associated with the lake: projects/{project_id}/locations/{location_id}/services/{service_id}
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type MetastoreParameters struct {

	// Optional. A relative reference to the Dataproc Metastore (https://cloud.google.com/dataproc-metastore/docs) service associated with the lake: projects/{project_id}/locations/{location_id}/services/{service_id}
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type MetastoreStatusInitParameters struct {
}

type MetastoreStatusObservation struct {
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Output only. Current state of the lake. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Output only. The time when the lake was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type MetastoreStatusParameters struct {
}

// LakeSpec defines the desired state of Lake
type LakeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LakeParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider LakeInitParameters `json:"initProvider,omitempty"`
}

// LakeStatus defines the observed state of Lake.
type LakeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LakeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Lake is the Schema for the Lakes API. The Dataplex Lake resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Lake struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LakeSpec   `json:"spec"`
	Status            LakeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LakeList contains a list of Lakes
type LakeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Lake `json:"items"`
}

// Repository type metadata.
var (
	Lake_Kind             = "Lake"
	Lake_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Lake_Kind}.String()
	Lake_KindAPIVersion   = Lake_Kind + "." + CRDGroupVersion.String()
	Lake_GroupVersionKind = CRDGroupVersion.WithKind(Lake_Kind)
)

func init() {
	SchemeBuilder.Register(&Lake{}, &LakeList{})
}
