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

type AppProfileInitParameters struct {

	// Long form description of the use case for this app profile.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings *bool `json:"ignoreWarnings,omitempty" tf:"ignore_warnings,omitempty"`

	MultiClusterRoutingClusterIds []*string `json:"multiClusterRoutingClusterIds,omitempty" tf:"multi_cluster_routing_cluster_ids,omitempty"`

	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	MultiClusterRoutingUseAny *bool `json:"multiClusterRoutingUseAny,omitempty" tf:"multi_cluster_routing_use_any,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Use a single-cluster routing policy.
	// Structure is documented below.
	SingleClusterRouting []SingleClusterRoutingInitParameters `json:"singleClusterRouting,omitempty" tf:"single_cluster_routing,omitempty"`
}

type AppProfileObservation struct {

	// Long form description of the use case for this app profile.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings *bool `json:"ignoreWarnings,omitempty" tf:"ignore_warnings,omitempty"`

	// The name of the instance to create the app profile within.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	MultiClusterRoutingClusterIds []*string `json:"multiClusterRoutingClusterIds,omitempty" tf:"multi_cluster_routing_cluster_ids,omitempty"`

	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	MultiClusterRoutingUseAny *bool `json:"multiClusterRoutingUseAny,omitempty" tf:"multi_cluster_routing_use_any,omitempty"`

	// The unique name of the requested app profile. Values are of the form projects/<project>/instances/<instance>/appProfiles/<appProfileId>.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Use a single-cluster routing policy.
	// Structure is documented below.
	SingleClusterRouting []SingleClusterRoutingObservation `json:"singleClusterRouting,omitempty" tf:"single_cluster_routing,omitempty"`
}

type AppProfileParameters struct {

	// Long form description of the use case for this app profile.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If true, ignore safety checks when deleting/updating the app profile.
	IgnoreWarnings *bool `json:"ignoreWarnings,omitempty" tf:"ignore_warnings,omitempty"`

	// The name of the instance to create the app profile within.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigtable/v1beta1.Instance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a Instance in bigtable to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a Instance in bigtable to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	MultiClusterRoutingClusterIds []*string `json:"multiClusterRoutingClusterIds,omitempty" tf:"multi_cluster_routing_cluster_ids,omitempty"`

	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	MultiClusterRoutingUseAny *bool `json:"multiClusterRoutingUseAny,omitempty" tf:"multi_cluster_routing_use_any,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Use a single-cluster routing policy.
	// Structure is documented below.
	SingleClusterRouting []SingleClusterRoutingParameters `json:"singleClusterRouting,omitempty" tf:"single_cluster_routing,omitempty"`
}

type SingleClusterRoutingInitParameters struct {

	// If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
	// It is unsafe to send these requests to the same table/row/column in multiple clusters.
	AllowTransactionalWrites *bool `json:"allowTransactionalWrites,omitempty" tf:"allow_transactional_writes,omitempty"`

	// The cluster to which read/write requests should be routed.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`
}

type SingleClusterRoutingObservation struct {

	// If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
	// It is unsafe to send these requests to the same table/row/column in multiple clusters.
	AllowTransactionalWrites *bool `json:"allowTransactionalWrites,omitempty" tf:"allow_transactional_writes,omitempty"`

	// The cluster to which read/write requests should be routed.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`
}

type SingleClusterRoutingParameters struct {

	// If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
	// It is unsafe to send these requests to the same table/row/column in multiple clusters.
	AllowTransactionalWrites *bool `json:"allowTransactionalWrites,omitempty" tf:"allow_transactional_writes,omitempty"`

	// The cluster to which read/write requests should be routed.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`
}

// AppProfileSpec defines the desired state of AppProfile
type AppProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppProfileParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider AppProfileInitParameters `json:"initProvider,omitempty"`
}

// AppProfileStatus defines the observed state of AppProfile.
type AppProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppProfile is the Schema for the AppProfiles API. App profile is a configuration object describing how Cloud Bigtable should treat traffic from a particular end user application.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AppProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppProfileSpec   `json:"spec"`
	Status            AppProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppProfileList contains a list of AppProfiles
type AppProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppProfile `json:"items"`
}

// Repository type metadata.
var (
	AppProfile_Kind             = "AppProfile"
	AppProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppProfile_Kind}.String()
	AppProfile_KindAPIVersion   = AppProfile_Kind + "." + CRDGroupVersion.String()
	AppProfile_GroupVersionKind = CRDGroupVersion.WithKind(AppProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&AppProfile{}, &AppProfileList{})
}
