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

type NATAddressInitParameters struct {
}

type NATAddressObservation struct {

	// an identifier for the resource with format {{instance_id}}/natAddresses/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The allocated NAT IP address.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The Apigee instance associated with the Apigee environment,
	// in the format organizations/{{org_name}}/instances/{{instance_name}}.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// State of the NAT IP address.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type NATAddressParameters struct {

	// The Apigee instance associated with the Apigee environment,
	// in the format organizations/{{org_name}}/instances/{{instance_name}}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/apigee/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in apigee to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in apigee to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`
}

// NATAddressSpec defines the desired state of NATAddress
type NATAddressSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NATAddressParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider NATAddressInitParameters `json:"initProvider,omitempty"`
}

// NATAddressStatus defines the observed state of NATAddress.
type NATAddressStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NATAddressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NATAddress is the Schema for the NATAddresss API. Apigee NAT (network address translation) address.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NATAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NATAddressSpec   `json:"spec"`
	Status            NATAddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NATAddressList contains a list of NATAddresss
type NATAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NATAddress `json:"items"`
}

// Repository type metadata.
var (
	NATAddress_Kind             = "NATAddress"
	NATAddress_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NATAddress_Kind}.String()
	NATAddress_KindAPIVersion   = NATAddress_Kind + "." + CRDGroupVersion.String()
	NATAddress_GroupVersionKind = CRDGroupVersion.WithKind(NATAddress_Kind)
)

func init() {
	SchemeBuilder.Register(&NATAddress{}, &NATAddressList{})
}
