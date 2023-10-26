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

type BackupInitParameters struct {

	// User-provided description of the backup.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
	// Structure is documented below.
	EncryptionConfig []EncryptionConfigInitParameters `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// User-defined labels for the alloydb backup.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type BackupObservation struct {

	// The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Time the Backup was created in UTC.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// User-provided description of the backup.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
	// Structure is documented below.
	EncryptionConfig []EncryptionConfigObservation `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// EncryptionInfo describes the encryption information of a cluster or a backup.
	// Structure is documented below.
	EncryptionInfo []EncryptionInfoObservation `json:"encryptionInfo,omitempty" tf:"encryption_info,omitempty"`

	// A hash of the resource.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/backups/{{backup_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// User-defined labels for the alloydb backup.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location where the alloydb backup should reside.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// If true, indicates that the service is actively updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.
	Reconciling *bool `json:"reconciling,omitempty" tf:"reconciling,omitempty"`

	// The current state of the backup.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// Time the Backup was updated in UTC.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type BackupParameters struct {

	// The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/alloydb/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in alloydb to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in alloydb to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// User-provided description of the backup.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EncryptionConfig []EncryptionConfigParameters `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// User-defined labels for the alloydb backup.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location where the alloydb backup should reside.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type EncryptionConfigInitParameters struct {

	// The fully-qualified resource name of the KMS key. Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type EncryptionConfigObservation struct {

	// The fully-qualified resource name of the KMS key. Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type EncryptionConfigParameters struct {

	// The fully-qualified resource name of the KMS key. Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type EncryptionInfoInitParameters struct {
}

type EncryptionInfoObservation struct {

	// (Output)
	// Output only. Type of encryption.
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`

	// (Output)
	// Output only. Cloud KMS key versions that are being used to protect the database or the backup.
	KMSKeyVersions []*string `json:"kmsKeyVersions,omitempty" tf:"kms_key_versions,omitempty"`
}

type EncryptionInfoParameters struct {
}

// BackupSpec defines the desired state of Backup
type BackupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupParameters `json:"forProvider"`
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
	InitProvider BackupInitParameters `json:"initProvider,omitempty"`
}

// BackupStatus defines the observed state of Backup.
type BackupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Backup is the Schema for the Backups API. An AlloyDB Backup.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupSpec   `json:"spec"`
	Status            BackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupList contains a list of Backups
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

// Repository type metadata.
var (
	Backup_Kind             = "Backup"
	Backup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Backup_Kind}.String()
	Backup_KindAPIVersion   = Backup_Kind + "." + CRDGroupVersion.String()
	Backup_GroupVersionKind = CRDGroupVersion.WithKind(Backup_Kind)
)

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
