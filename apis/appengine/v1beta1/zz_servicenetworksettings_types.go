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

type NetworkSettingsInitParameters struct {

	// The ingress settings for version or service.
	// Default value is INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED.
	// Possible values are: INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED, INGRESS_TRAFFIC_ALLOWED_ALL, INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY, INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB.
	IngressTrafficAllowed *string `json:"ingressTrafficAllowed,omitempty" tf:"ingress_traffic_allowed,omitempty"`
}

type NetworkSettingsObservation struct {

	// The ingress settings for version or service.
	// Default value is INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED.
	// Possible values are: INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED, INGRESS_TRAFFIC_ALLOWED_ALL, INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY, INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB.
	IngressTrafficAllowed *string `json:"ingressTrafficAllowed,omitempty" tf:"ingress_traffic_allowed,omitempty"`
}

type NetworkSettingsParameters struct {

	// The ingress settings for version or service.
	// Default value is INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED.
	// Possible values are: INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED, INGRESS_TRAFFIC_ALLOWED_ALL, INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY, INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB.
	// +kubebuilder:validation:Optional
	IngressTrafficAllowed *string `json:"ingressTrafficAllowed,omitempty" tf:"ingress_traffic_allowed,omitempty"`
}

type ServiceNetworkSettingsInitParameters struct {

	// Ingress settings for this service. Will apply to all versions.
	// Structure is documented below.
	NetworkSettings []NetworkSettingsInitParameters `json:"networkSettings,omitempty" tf:"network_settings,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ServiceNetworkSettingsObservation struct {

	// an identifier for the resource with format apps/{{project}}/services/{{service}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Ingress settings for this service. Will apply to all versions.
	// Structure is documented below.
	NetworkSettings []NetworkSettingsObservation `json:"networkSettings,omitempty" tf:"network_settings,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the service these settings apply to.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ServiceNetworkSettingsParameters struct {

	// Ingress settings for this service. Will apply to all versions.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NetworkSettings []NetworkSettingsParameters `json:"networkSettings,omitempty" tf:"network_settings,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the service these settings apply to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/appengine/v1beta1.StandardAppVersion
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("service",false)
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// Reference to a StandardAppVersion in appengine to populate service.
	// +kubebuilder:validation:Optional
	ServiceRef *v1.Reference `json:"serviceRef,omitempty" tf:"-"`

	// Selector for a StandardAppVersion in appengine to populate service.
	// +kubebuilder:validation:Optional
	ServiceSelector *v1.Selector `json:"serviceSelector,omitempty" tf:"-"`
}

// ServiceNetworkSettingsSpec defines the desired state of ServiceNetworkSettings
type ServiceNetworkSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceNetworkSettingsParameters `json:"forProvider"`
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
	InitProvider ServiceNetworkSettingsInitParameters `json:"initProvider,omitempty"`
}

// ServiceNetworkSettingsStatus defines the observed state of ServiceNetworkSettings.
type ServiceNetworkSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceNetworkSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceNetworkSettings is the Schema for the ServiceNetworkSettingss API. A NetworkSettings resource is a container for ingress settings for a version or service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ServiceNetworkSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkSettings) || (has(self.initProvider) && has(self.initProvider.networkSettings))",message="spec.forProvider.networkSettings is a required parameter"
	Spec   ServiceNetworkSettingsSpec   `json:"spec"`
	Status ServiceNetworkSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceNetworkSettingsList contains a list of ServiceNetworkSettingss
type ServiceNetworkSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceNetworkSettings `json:"items"`
}

// Repository type metadata.
var (
	ServiceNetworkSettings_Kind             = "ServiceNetworkSettings"
	ServiceNetworkSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceNetworkSettings_Kind}.String()
	ServiceNetworkSettings_KindAPIVersion   = ServiceNetworkSettings_Kind + "." + CRDGroupVersion.String()
	ServiceNetworkSettings_GroupVersionKind = CRDGroupVersion.WithKind(ServiceNetworkSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceNetworkSettings{}, &ServiceNetworkSettingsList{})
}
