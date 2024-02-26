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

type VPNTunnelInitParameters struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IKE protocol version to use when establishing the VPN tunnel with
	// peer VPN gateway.
	// Acceptable IKE versions are 1 or 2. Default version is 2.
	IkeVersion *float64 `json:"ikeVersion,omitempty" tf:"ike_version,omitempty"`

	// Labels to apply to this VpnTunnel.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Local traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example 192.168.0.0/16. The ranges should be disjoint.
	// Only IPv4 is supported.
	// +listType=set
	LocalTrafficSelector []*string `json:"localTrafficSelector,omitempty" tf:"local_traffic_selector,omitempty"`

	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	// +crossplane:generate:reference:type=ExternalVPNGateway
	PeerExternalGateway *string `json:"peerExternalGateway,omitempty" tf:"peer_external_gateway,omitempty"`

	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface *float64 `json:"peerExternalGatewayInterface,omitempty" tf:"peer_external_gateway_interface,omitempty"`

	// Reference to a ExternalVPNGateway to populate peerExternalGateway.
	// +kubebuilder:validation:Optional
	PeerExternalGatewayRef *v1.Reference `json:"peerExternalGatewayRef,omitempty" tf:"-"`

	// Selector for a ExternalVPNGateway to populate peerExternalGateway.
	// +kubebuilder:validation:Optional
	PeerExternalGatewaySelector *v1.Selector `json:"peerExternalGatewaySelector,omitempty" tf:"-"`

	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
	// If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
	// ID in the peer GCP VPN gateway.
	// This field must reference a google_compute_ha_vpn_gateway resource.
	PeerGCPGateway *string `json:"peerGcpGateway,omitempty" tf:"peer_gcp_gateway,omitempty"`

	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIP *string `json:"peerIp,omitempty" tf:"peer_ip,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Remote traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example 192.168.0.0/16. The ranges should be disjoint.
	// Only IPv4 is supported.
	// +listType=set
	RemoteTrafficSelector []*string `json:"remoteTrafficSelector,omitempty" tf:"remote_traffic_selector,omitempty"`

	// URL of router resource to be used for dynamic routing.
	// +crossplane:generate:reference:type=Router
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// Reference to a Router to populate router.
	// +kubebuilder:validation:Optional
	RouterRef *v1.Reference `json:"routerRef,omitempty" tf:"-"`

	// Selector for a Router to populate router.
	// +kubebuilder:validation:Optional
	RouterSelector *v1.Selector `json:"routerSelector,omitempty" tf:"-"`

	// URL of the Target VPN gateway with which this VPN tunnel is
	// associated.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.VPNGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TargetVPNGateway *string `json:"targetVpnGateway,omitempty" tf:"target_vpn_gateway,omitempty"`

	// Reference to a VPNGateway in compute to populate targetVpnGateway.
	// +kubebuilder:validation:Optional
	TargetVPNGatewayRef *v1.Reference `json:"targetVpnGatewayRef,omitempty" tf:"-"`

	// Selector for a VPNGateway in compute to populate targetVpnGateway.
	// +kubebuilder:validation:Optional
	TargetVPNGatewaySelector *v1.Selector `json:"targetVpnGatewaySelector,omitempty" tf:"-"`

	// URL of the VPN gateway with which this VPN tunnel is associated.
	// This must be used if a High Availability VPN gateway resource is created.
	// This field must reference a google_compute_ha_vpn_gateway resource.
	// +crossplane:generate:reference:type=HaVPNGateway
	VPNGateway *string `json:"vpnGateway,omitempty" tf:"vpn_gateway,omitempty"`

	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VPNGatewayInterface *float64 `json:"vpnGatewayInterface,omitempty" tf:"vpn_gateway_interface,omitempty"`

	// Reference to a HaVPNGateway to populate vpnGateway.
	// +kubebuilder:validation:Optional
	VPNGatewayRef *v1.Reference `json:"vpnGatewayRef,omitempty" tf:"-"`

	// Selector for a HaVPNGateway to populate vpnGateway.
	// +kubebuilder:validation:Optional
	VPNGatewaySelector *v1.Selector `json:"vpnGatewaySelector,omitempty" tf:"-"`
}

type VPNTunnelObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Detailed status message for the VPN tunnel.
	DetailedStatus *string `json:"detailedStatus,omitempty" tf:"detailed_status,omitempty"`

	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IKE protocol version to use when establishing the VPN tunnel with
	// peer VPN gateway.
	// Acceptable IKE versions are 1 or 2. Default version is 2.
	IkeVersion *float64 `json:"ikeVersion,omitempty" tf:"ike_version,omitempty"`

	// The fingerprint used for optimistic locking of this resource.  Used
	// internally during updates.
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	// Labels to apply to this VpnTunnel.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Local traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example 192.168.0.0/16. The ranges should be disjoint.
	// Only IPv4 is supported.
	// +listType=set
	LocalTrafficSelector []*string `json:"localTrafficSelector,omitempty" tf:"local_traffic_selector,omitempty"`

	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway *string `json:"peerExternalGateway,omitempty" tf:"peer_external_gateway,omitempty"`

	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface *float64 `json:"peerExternalGatewayInterface,omitempty" tf:"peer_external_gateway_interface,omitempty"`

	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
	// If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
	// ID in the peer GCP VPN gateway.
	// This field must reference a google_compute_ha_vpn_gateway resource.
	PeerGCPGateway *string `json:"peerGcpGateway,omitempty" tf:"peer_gcp_gateway,omitempty"`

	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIP *string `json:"peerIp,omitempty" tf:"peer_ip,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region where the tunnel is located. If unset, is set to the region of target_vpn_gateway.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Remote traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example 192.168.0.0/16. The ranges should be disjoint.
	// Only IPv4 is supported.
	// +listType=set
	RemoteTrafficSelector []*string `json:"remoteTrafficSelector,omitempty" tf:"remote_traffic_selector,omitempty"`

	// URL of router resource to be used for dynamic routing.
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Hash of the shared secret.
	SharedSecretHash *string `json:"sharedSecretHash,omitempty" tf:"shared_secret_hash,omitempty"`

	// URL of the Target VPN gateway with which this VPN tunnel is
	// associated.
	TargetVPNGateway *string `json:"targetVpnGateway,omitempty" tf:"target_vpn_gateway,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// The unique identifier for the resource. This identifier is defined by the server.
	TunnelID *string `json:"tunnelId,omitempty" tf:"tunnel_id,omitempty"`

	// URL of the VPN gateway with which this VPN tunnel is associated.
	// This must be used if a High Availability VPN gateway resource is created.
	// This field must reference a google_compute_ha_vpn_gateway resource.
	VPNGateway *string `json:"vpnGateway,omitempty" tf:"vpn_gateway,omitempty"`

	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VPNGatewayInterface *float64 `json:"vpnGatewayInterface,omitempty" tf:"vpn_gateway_interface,omitempty"`
}

type VPNTunnelParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IKE protocol version to use when establishing the VPN tunnel with
	// peer VPN gateway.
	// Acceptable IKE versions are 1 or 2. Default version is 2.
	// +kubebuilder:validation:Optional
	IkeVersion *float64 `json:"ikeVersion,omitempty" tf:"ike_version,omitempty"`

	// Labels to apply to this VpnTunnel.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Local traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example 192.168.0.0/16. The ranges should be disjoint.
	// Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	// +listType=set
	LocalTrafficSelector []*string `json:"localTrafficSelector,omitempty" tf:"local_traffic_selector,omitempty"`

	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	// +crossplane:generate:reference:type=ExternalVPNGateway
	// +kubebuilder:validation:Optional
	PeerExternalGateway *string `json:"peerExternalGateway,omitempty" tf:"peer_external_gateway,omitempty"`

	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	// +kubebuilder:validation:Optional
	PeerExternalGatewayInterface *float64 `json:"peerExternalGatewayInterface,omitempty" tf:"peer_external_gateway_interface,omitempty"`

	// Reference to a ExternalVPNGateway to populate peerExternalGateway.
	// +kubebuilder:validation:Optional
	PeerExternalGatewayRef *v1.Reference `json:"peerExternalGatewayRef,omitempty" tf:"-"`

	// Selector for a ExternalVPNGateway to populate peerExternalGateway.
	// +kubebuilder:validation:Optional
	PeerExternalGatewaySelector *v1.Selector `json:"peerExternalGatewaySelector,omitempty" tf:"-"`

	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
	// If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
	// ID in the peer GCP VPN gateway.
	// This field must reference a google_compute_ha_vpn_gateway resource.
	// +kubebuilder:validation:Optional
	PeerGCPGateway *string `json:"peerGcpGateway,omitempty" tf:"peer_gcp_gateway,omitempty"`

	// IP address of the peer VPN gateway. Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	PeerIP *string `json:"peerIp,omitempty" tf:"peer_ip,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region where the tunnel is located. If unset, is set to the region of target_vpn_gateway.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// Remote traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example 192.168.0.0/16. The ranges should be disjoint.
	// Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	// +listType=set
	RemoteTrafficSelector []*string `json:"remoteTrafficSelector,omitempty" tf:"remote_traffic_selector,omitempty"`

	// URL of router resource to be used for dynamic routing.
	// +crossplane:generate:reference:type=Router
	// +kubebuilder:validation:Optional
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// Reference to a Router to populate router.
	// +kubebuilder:validation:Optional
	RouterRef *v1.Reference `json:"routerRef,omitempty" tf:"-"`

	// Selector for a Router to populate router.
	// +kubebuilder:validation:Optional
	RouterSelector *v1.Selector `json:"routerSelector,omitempty" tf:"-"`

	// Shared secret used to set the secure session between the Cloud VPN
	// gateway and the peer VPN gateway.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	SharedSecretSecretRef v1.SecretKeySelector `json:"sharedSecretSecretRef" tf:"-"`

	// URL of the Target VPN gateway with which this VPN tunnel is
	// associated.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.VPNGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetVPNGateway *string `json:"targetVpnGateway,omitempty" tf:"target_vpn_gateway,omitempty"`

	// Reference to a VPNGateway in compute to populate targetVpnGateway.
	// +kubebuilder:validation:Optional
	TargetVPNGatewayRef *v1.Reference `json:"targetVpnGatewayRef,omitempty" tf:"-"`

	// Selector for a VPNGateway in compute to populate targetVpnGateway.
	// +kubebuilder:validation:Optional
	TargetVPNGatewaySelector *v1.Selector `json:"targetVpnGatewaySelector,omitempty" tf:"-"`

	// URL of the VPN gateway with which this VPN tunnel is associated.
	// This must be used if a High Availability VPN gateway resource is created.
	// This field must reference a google_compute_ha_vpn_gateway resource.
	// +crossplane:generate:reference:type=HaVPNGateway
	// +kubebuilder:validation:Optional
	VPNGateway *string `json:"vpnGateway,omitempty" tf:"vpn_gateway,omitempty"`

	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	// +kubebuilder:validation:Optional
	VPNGatewayInterface *float64 `json:"vpnGatewayInterface,omitempty" tf:"vpn_gateway_interface,omitempty"`

	// Reference to a HaVPNGateway to populate vpnGateway.
	// +kubebuilder:validation:Optional
	VPNGatewayRef *v1.Reference `json:"vpnGatewayRef,omitempty" tf:"-"`

	// Selector for a HaVPNGateway to populate vpnGateway.
	// +kubebuilder:validation:Optional
	VPNGatewaySelector *v1.Selector `json:"vpnGatewaySelector,omitempty" tf:"-"`
}

// VPNTunnelSpec defines the desired state of VPNTunnel
type VPNTunnelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNTunnelParameters `json:"forProvider"`
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
	InitProvider VPNTunnelInitParameters `json:"initProvider,omitempty"`
}

// VPNTunnelStatus defines the observed state of VPNTunnel.
type VPNTunnelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNTunnelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPNTunnel is the Schema for the VPNTunnels API. VPN tunnel resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type VPNTunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sharedSecretSecretRef)",message="spec.forProvider.sharedSecretSecretRef is a required parameter"
	Spec   VPNTunnelSpec   `json:"spec"`
	Status VPNTunnelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNTunnelList contains a list of VPNTunnels
type VPNTunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNTunnel `json:"items"`
}

// Repository type metadata.
var (
	VPNTunnel_Kind             = "VPNTunnel"
	VPNTunnel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNTunnel_Kind}.String()
	VPNTunnel_KindAPIVersion   = VPNTunnel_Kind + "." + CRDGroupVersion.String()
	VPNTunnel_GroupVersionKind = CRDGroupVersion.WithKind(VPNTunnel_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNTunnel{}, &VPNTunnelList{})
}
