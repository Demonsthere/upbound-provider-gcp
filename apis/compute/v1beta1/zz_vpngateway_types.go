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

type VPNGatewayInitParameters struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type VPNGatewayObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique identifier for the resource.
	GatewayID *float64 `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The network this VPN gateway is accepting traffic for.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region this gateway should sit in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type VPNGatewayParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The network this VPN gateway is accepting traffic for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region this gateway should sit in.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

// VPNGatewaySpec defines the desired state of VPNGateway
type VPNGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNGatewayParameters `json:"forProvider"`
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
	InitProvider VPNGatewayInitParameters `json:"initProvider,omitempty"`
}

// VPNGatewayStatus defines the observed state of VPNGateway.
type VPNGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGateway is the Schema for the VPNGateways API. Represents a VPN gateway running in GCP.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type VPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPNGatewaySpec   `json:"spec"`
	Status            VPNGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayList contains a list of VPNGateways
type VPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNGateway `json:"items"`
}

// Repository type metadata.
var (
	VPNGateway_Kind             = "VPNGateway"
	VPNGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNGateway_Kind}.String()
	VPNGateway_KindAPIVersion   = VPNGateway_Kind + "." + CRDGroupVersion.String()
	VPNGateway_GroupVersionKind = CRDGroupVersion.WithKind(VPNGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNGateway{}, &VPNGatewayList{})
}
