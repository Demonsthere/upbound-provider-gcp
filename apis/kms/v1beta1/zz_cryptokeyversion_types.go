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

type AttestationInitParameters struct {
}

type AttestationObservation struct {

	// The certificate chains needed to validate the attestation
	// Structure is documented below.
	CertChains []CertChainsObservation `json:"certChains,omitempty" tf:"cert_chains,omitempty"`

	// (Output)
	// The attestation data provided by the HSM when the key operation was performed.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level and EXTERNAL_VPC protection levels.
	// Structure is documented below.
	ExternalProtectionLevelOptions []ExternalProtectionLevelOptionsObservation `json:"externalProtectionLevelOptions,omitempty" tf:"external_protection_level_options,omitempty"`

	// (Output)
	// The format of the attestation data.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`
}

type AttestationParameters struct {
}

type CertChainsInitParameters struct {
}

type CertChainsObservation struct {

	// Cavium certificate chain corresponding to the attestation.
	CaviumCerts *string `json:"caviumCerts,omitempty" tf:"cavium_certs,omitempty"`

	// Google card certificate chain corresponding to the attestation.
	GoogleCardCerts *string `json:"googleCardCerts,omitempty" tf:"google_card_certs,omitempty"`

	// Google partition certificate chain corresponding to the attestation.
	GooglePartitionCerts *string `json:"googlePartitionCerts,omitempty" tf:"google_partition_certs,omitempty"`
}

type CertChainsParameters struct {
}

type CryptoKeyVersionInitParameters struct {

	// The current state of the CryptoKeyVersion.
	// Possible values are: PENDING_GENERATION, ENABLED, DISABLED, DESTROYED, DESTROY_SCHEDULED, PENDING_IMPORT, IMPORT_FAILED.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type CryptoKeyVersionObservation struct {

	// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
	// Only provided for key versions with protectionLevel HSM.
	// Structure is documented below.
	Attestation []AttestationObservation `json:"attestation,omitempty" tf:"attestation,omitempty"`

	// The name of the cryptoKey associated with the CryptoKeyVersions.
	// Format: 'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'
	CryptoKey *string `json:"cryptoKey,omitempty" tf:"crypto_key,omitempty"`

	// The time this CryptoKeyVersion key material was generated
	GenerateTime *string `json:"generateTime,omitempty" tf:"generate_time,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource name for this CryptoKeyVersion.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
	ProtectionLevel *string `json:"protectionLevel,omitempty" tf:"protection_level,omitempty"`

	// The current state of the CryptoKeyVersion.
	// Possible values are: PENDING_GENERATION, ENABLED, DISABLED, DESTROYED, DESTROY_SCHEDULED, PENDING_IMPORT, IMPORT_FAILED.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type CryptoKeyVersionParameters struct {

	// The name of the cryptoKey associated with the CryptoKeyVersions.
	// Format: 'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CryptoKey *string `json:"cryptoKey,omitempty" tf:"crypto_key,omitempty"`

	// Reference to a CryptoKey in kms to populate cryptoKey.
	// +kubebuilder:validation:Optional
	CryptoKeyRef *v1.Reference `json:"cryptoKeyRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate cryptoKey.
	// +kubebuilder:validation:Optional
	CryptoKeySelector *v1.Selector `json:"cryptoKeySelector,omitempty" tf:"-"`

	// The current state of the CryptoKeyVersion.
	// Possible values are: PENDING_GENERATION, ENABLED, DISABLED, DESTROYED, DESTROY_SCHEDULED, PENDING_IMPORT, IMPORT_FAILED.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ExternalProtectionLevelOptionsInitParameters struct {
}

type ExternalProtectionLevelOptionsObservation struct {

	// The path to the external key material on the EKM when using EkmConnection e.g., "v0/my/key". Set this field instead of externalKeyUri when using an EkmConnection.
	EkmConnectionKeyPath *string `json:"ekmConnectionKeyPath,omitempty" tf:"ekm_connection_key_path,omitempty"`

	// The URI for an external resource that this CryptoKeyVersion represents.
	ExternalKeyURI *string `json:"externalKeyUri,omitempty" tf:"external_key_uri,omitempty"`
}

type ExternalProtectionLevelOptionsParameters struct {
}

// CryptoKeyVersionSpec defines the desired state of CryptoKeyVersion
type CryptoKeyVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CryptoKeyVersionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider CryptoKeyVersionInitParameters `json:"initProvider,omitempty"`
}

// CryptoKeyVersionStatus defines the observed state of CryptoKeyVersion.
type CryptoKeyVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CryptoKeyVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CryptoKeyVersion is the Schema for the CryptoKeyVersions API. A
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type CryptoKeyVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CryptoKeyVersionSpec   `json:"spec"`
	Status            CryptoKeyVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CryptoKeyVersionList contains a list of CryptoKeyVersions
type CryptoKeyVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CryptoKeyVersion `json:"items"`
}

// Repository type metadata.
var (
	CryptoKeyVersion_Kind             = "CryptoKeyVersion"
	CryptoKeyVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CryptoKeyVersion_Kind}.String()
	CryptoKeyVersion_KindAPIVersion   = CryptoKeyVersion_Kind + "." + CRDGroupVersion.String()
	CryptoKeyVersion_GroupVersionKind = CRDGroupVersion.WithKind(CryptoKeyVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&CryptoKeyVersion{}, &CryptoKeyVersionList{})
}
