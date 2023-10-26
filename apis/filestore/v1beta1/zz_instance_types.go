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

type FileSharesInitParameters struct {

	// File share capacity in GiB. This must be at least 1024 GiB
	// for the standard tier, or 2560 GiB for the premium tier.
	CapacityGb *float64 `json:"capacityGb,omitempty" tf:"capacity_gb,omitempty"`

	// Nfs Export Options. There is a limit of 10 export options per file share.
	// Structure is documented below.
	NFSExportOptions []NFSExportOptionsInitParameters `json:"nfsExportOptions,omitempty" tf:"nfs_export_options,omitempty"`

	// The name of the fileshare (16 characters or less)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FileSharesObservation struct {

	// File share capacity in GiB. This must be at least 1024 GiB
	// for the standard tier, or 2560 GiB for the premium tier.
	CapacityGb *float64 `json:"capacityGb,omitempty" tf:"capacity_gb,omitempty"`

	// Nfs Export Options. There is a limit of 10 export options per file share.
	// Structure is documented below.
	NFSExportOptions []NFSExportOptionsObservation `json:"nfsExportOptions,omitempty" tf:"nfs_export_options,omitempty"`

	// The name of the fileshare (16 characters or less)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Output)
	// The resource name of the backup, in the format
	// projects/{projectId}/locations/{locationId}/backups/{backupId},
	// that this file share has been restored from.
	SourceBackup *string `json:"sourceBackup,omitempty" tf:"source_backup,omitempty"`
}

type FileSharesParameters struct {

	// File share capacity in GiB. This must be at least 1024 GiB
	// for the standard tier, or 2560 GiB for the premium tier.
	// +kubebuilder:validation:Optional
	CapacityGb *float64 `json:"capacityGb" tf:"capacity_gb,omitempty"`

	// Nfs Export Options. There is a limit of 10 export options per file share.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NFSExportOptions []NFSExportOptionsParameters `json:"nfsExportOptions,omitempty" tf:"nfs_export_options,omitempty"`

	// The name of the fileshare (16 characters or less)
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type InstanceInitParameters struct {

	// A description of the instance.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares []FileSharesInitParameters `json:"fileShares,omitempty" tf:"file_shares,omitempty"`

	// Resource labels to represent user-provided metadata.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks []NetworksInitParameters `json:"networks,omitempty" tf:"networks,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The service tier of the instance.
	// Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD and ENTERPRISE
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// The name of the Filestore zone of the instance.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type InstanceObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// A description of the instance.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Server-specified ETag for the instance resource to prevent
	// simultaneous updates from overwriting each other.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	FileShares []FileSharesObservation `json:"fileShares,omitempty" tf:"file_shares,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/instances/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// KMS key name used for data encryption.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Resource labels to represent user-provided metadata.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	Networks []NetworksObservation `json:"networks,omitempty" tf:"networks,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The service tier of the instance.
	// Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD and ENTERPRISE
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// The name of the Filestore zone of the instance.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type InstanceParameters struct {

	// A description of the instance.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// File system shares on the instance. For this version, only a
	// single file share is supported.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	FileShares []FileSharesParameters `json:"fileShares,omitempty" tf:"file_shares,omitempty"`

	// KMS key name used for data encryption.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameRef *v1.Reference `json:"kmsKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameSelector *v1.Selector `json:"kmsKeyNameSelector,omitempty" tf:"-"`

	// Resource labels to represent user-provided metadata.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Networks []NetworksParameters `json:"networks,omitempty" tf:"networks,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The service tier of the instance.
	// Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD and ENTERPRISE
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// The name of the Filestore zone of the instance.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NFSExportOptionsInitParameters struct {

	// Either READ_ONLY, for allowing only read requests on the exported directory,
	// or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
	// Default value is READ_WRITE.
	// Possible values are: READ_ONLY, READ_WRITE.
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// An integer representing the anonymous group id with a default value of 65534.
	// Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
	// if this field is specified for other squashMode settings.
	AnonGID *float64 `json:"anonGid,omitempty" tf:"anon_gid,omitempty"`

	// An integer representing the anonymous user id with a default value of 65534.
	// Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
	// if this field is specified for other squashMode settings.
	AnonUID *float64 `json:"anonUid,omitempty" tf:"anon_uid,omitempty"`

	// List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
	// Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
	// The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`

	// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
	// for not allowing root access. The default is NO_ROOT_SQUASH.
	// Default value is NO_ROOT_SQUASH.
	// Possible values are: NO_ROOT_SQUASH, ROOT_SQUASH.
	SquashMode *string `json:"squashMode,omitempty" tf:"squash_mode,omitempty"`
}

type NFSExportOptionsObservation struct {

	// Either READ_ONLY, for allowing only read requests on the exported directory,
	// or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
	// Default value is READ_WRITE.
	// Possible values are: READ_ONLY, READ_WRITE.
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// An integer representing the anonymous group id with a default value of 65534.
	// Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
	// if this field is specified for other squashMode settings.
	AnonGID *float64 `json:"anonGid,omitempty" tf:"anon_gid,omitempty"`

	// An integer representing the anonymous user id with a default value of 65534.
	// Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
	// if this field is specified for other squashMode settings.
	AnonUID *float64 `json:"anonUid,omitempty" tf:"anon_uid,omitempty"`

	// List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
	// Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
	// The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`

	// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
	// for not allowing root access. The default is NO_ROOT_SQUASH.
	// Default value is NO_ROOT_SQUASH.
	// Possible values are: NO_ROOT_SQUASH, ROOT_SQUASH.
	SquashMode *string `json:"squashMode,omitempty" tf:"squash_mode,omitempty"`
}

type NFSExportOptionsParameters struct {

	// Either READ_ONLY, for allowing only read requests on the exported directory,
	// or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
	// Default value is READ_WRITE.
	// Possible values are: READ_ONLY, READ_WRITE.
	// +kubebuilder:validation:Optional
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// An integer representing the anonymous group id with a default value of 65534.
	// Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
	// if this field is specified for other squashMode settings.
	// +kubebuilder:validation:Optional
	AnonGID *float64 `json:"anonGid,omitempty" tf:"anon_gid,omitempty"`

	// An integer representing the anonymous user id with a default value of 65534.
	// Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
	// if this field is specified for other squashMode settings.
	// +kubebuilder:validation:Optional
	AnonUID *float64 `json:"anonUid,omitempty" tf:"anon_uid,omitempty"`

	// List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
	// Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
	// The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
	// +kubebuilder:validation:Optional
	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`

	// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
	// for not allowing root access. The default is NO_ROOT_SQUASH.
	// Default value is NO_ROOT_SQUASH.
	// Possible values are: NO_ROOT_SQUASH, ROOT_SQUASH.
	// +kubebuilder:validation:Optional
	SquashMode *string `json:"squashMode,omitempty" tf:"squash_mode,omitempty"`
}

type NetworksInitParameters struct {

	// The network connect mode of the Filestore instance.
	// If not provided, the connect mode defaults to
	// DIRECT_PEERING.
	// Default value is DIRECT_PEERING.
	// Possible values are: DIRECT_PEERING, PRIVATE_SERVICE_ACCESS.
	ConnectMode *string `json:"connectMode,omitempty" tf:"connect_mode,omitempty"`

	// IP versions for which the instance has
	// IP addresses assigned.
	// Each value may be one of: ADDRESS_MODE_UNSPECIFIED, MODE_IPV4, MODE_IPV6.
	Modes []*string `json:"modes,omitempty" tf:"modes,omitempty"`

	// The name of the GCE VPC network to which the
	// instance is connected.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// A /29 CIDR block that identifies the range of IP
	// addresses reserved for this instance.
	ReservedIPRange *string `json:"reservedIpRange,omitempty" tf:"reserved_ip_range,omitempty"`
}

type NetworksObservation struct {

	// The network connect mode of the Filestore instance.
	// If not provided, the connect mode defaults to
	// DIRECT_PEERING.
	// Default value is DIRECT_PEERING.
	// Possible values are: DIRECT_PEERING, PRIVATE_SERVICE_ACCESS.
	ConnectMode *string `json:"connectMode,omitempty" tf:"connect_mode,omitempty"`

	// (Output)
	// A list of IPv4 or IPv6 addresses.
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// IP versions for which the instance has
	// IP addresses assigned.
	// Each value may be one of: ADDRESS_MODE_UNSPECIFIED, MODE_IPV4, MODE_IPV6.
	Modes []*string `json:"modes,omitempty" tf:"modes,omitempty"`

	// The name of the GCE VPC network to which the
	// instance is connected.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// A /29 CIDR block that identifies the range of IP
	// addresses reserved for this instance.
	ReservedIPRange *string `json:"reservedIpRange,omitempty" tf:"reserved_ip_range,omitempty"`
}

type NetworksParameters struct {

	// The network connect mode of the Filestore instance.
	// If not provided, the connect mode defaults to
	// DIRECT_PEERING.
	// Default value is DIRECT_PEERING.
	// Possible values are: DIRECT_PEERING, PRIVATE_SERVICE_ACCESS.
	// +kubebuilder:validation:Optional
	ConnectMode *string `json:"connectMode,omitempty" tf:"connect_mode,omitempty"`

	// IP versions for which the instance has
	// IP addresses assigned.
	// Each value may be one of: ADDRESS_MODE_UNSPECIFIED, MODE_IPV4, MODE_IPV6.
	// +kubebuilder:validation:Optional
	Modes []*string `json:"modes" tf:"modes,omitempty"`

	// The name of the GCE VPC network to which the
	// instance is connected.
	// +kubebuilder:validation:Optional
	Network *string `json:"network" tf:"network,omitempty"`

	// A /29 CIDR block that identifies the range of IP
	// addresses reserved for this instance.
	// +kubebuilder:validation:Optional
	ReservedIPRange *string `json:"reservedIpRange,omitempty" tf:"reserved_ip_range,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
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
	InitProvider InstanceInitParameters `json:"initProvider,omitempty"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. A Google Cloud Filestore instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fileShares) || (has(self.initProvider) && has(self.initProvider.fileShares))",message="spec.forProvider.fileShares is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networks) || (has(self.initProvider) && has(self.initProvider.networks))",message="spec.forProvider.networks is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tier) || (has(self.initProvider) && has(self.initProvider.tier))",message="spec.forProvider.tier is a required parameter"
	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
