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

type NetworkPeeringRoutesConfigInitParameters struct {

	// Whether to export the custom routes to the peer network.
	ExportCustomRoutes *bool `json:"exportCustomRoutes,omitempty" tf:"export_custom_routes,omitempty"`

	// Whether to import the custom routes to the peer network.
	ImportCustomRoutes *bool `json:"importCustomRoutes,omitempty" tf:"import_custom_routes,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type NetworkPeeringRoutesConfigObservation struct {

	// Whether to export the custom routes to the peer network.
	ExportCustomRoutes *bool `json:"exportCustomRoutes,omitempty" tf:"export_custom_routes,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to import the custom routes to the peer network.
	ImportCustomRoutes *bool `json:"importCustomRoutes,omitempty" tf:"import_custom_routes,omitempty"`

	// The name of the primary network for the peering.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Name of the peering.
	Peering *string `json:"peering,omitempty" tf:"peering,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type NetworkPeeringRoutesConfigParameters struct {

	// Whether to export the custom routes to the peer network.
	ExportCustomRoutes *bool `json:"exportCustomRoutes,omitempty" tf:"export_custom_routes,omitempty"`

	// Whether to import the custom routes to the peer network.
	ImportCustomRoutes *bool `json:"importCustomRoutes,omitempty" tf:"import_custom_routes,omitempty"`

	// The name of the primary network for the peering.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// Name of the peering.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.NetworkPeering
	// +kubebuilder:validation:Optional
	Peering *string `json:"peering,omitempty" tf:"peering,omitempty"`

	// Reference to a NetworkPeering in compute to populate peering.
	// +kubebuilder:validation:Optional
	PeeringRef *v1.Reference `json:"peeringRef,omitempty" tf:"-"`

	// Selector for a NetworkPeering in compute to populate peering.
	// +kubebuilder:validation:Optional
	PeeringSelector *v1.Selector `json:"peeringSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// NetworkPeeringRoutesConfigSpec defines the desired state of NetworkPeeringRoutesConfig
type NetworkPeeringRoutesConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkPeeringRoutesConfigParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider NetworkPeeringRoutesConfigInitParameters `json:"initProvider,omitempty"`
}

// NetworkPeeringRoutesConfigStatus defines the observed state of NetworkPeeringRoutesConfig.
type NetworkPeeringRoutesConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkPeeringRoutesConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkPeeringRoutesConfig is the Schema for the NetworkPeeringRoutesConfigs API. Manage a network peering's route settings without managing the peering as a whole.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NetworkPeeringRoutesConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.exportCustomRoutes) || has(self.initProvider.exportCustomRoutes)",message="exportCustomRoutes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.importCustomRoutes) || has(self.initProvider.importCustomRoutes)",message="importCustomRoutes is a required parameter"
	Spec   NetworkPeeringRoutesConfigSpec   `json:"spec"`
	Status NetworkPeeringRoutesConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkPeeringRoutesConfigList contains a list of NetworkPeeringRoutesConfigs
type NetworkPeeringRoutesConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkPeeringRoutesConfig `json:"items"`
}

// Repository type metadata.
var (
	NetworkPeeringRoutesConfig_Kind             = "NetworkPeeringRoutesConfig"
	NetworkPeeringRoutesConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkPeeringRoutesConfig_Kind}.String()
	NetworkPeeringRoutesConfig_KindAPIVersion   = NetworkPeeringRoutesConfig_Kind + "." + CRDGroupVersion.String()
	NetworkPeeringRoutesConfig_GroupVersionKind = CRDGroupVersion.WithKind(NetworkPeeringRoutesConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkPeeringRoutesConfig{}, &NetworkPeeringRoutesConfigList{})
}
