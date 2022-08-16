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

type GuestOsFeaturesObservation struct {
}

type GuestOsFeaturesParameters struct {

	// The type of supported feature. Read Enabling guest operating system features to see a list of available options.
	// Possible values are MULTI_IP_SUBNET, SECURE_BOOT, SEV_CAPABLE, UEFI_COMPATIBLE, VIRTIO_SCSI_MULTIQUEUE, WINDOWS, and GVNIC.
	// The type of supported feature. Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options. Possible values: ["MULTI_IP_SUBNET", "SECURE_BOOT", "SEV_CAPABLE", "UEFI_COMPATIBLE", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "GVNIC"]
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ImageObservation struct {

	// Size of the image tar.gz archive stored in Google Cloud Storage .
	// Size of the image tar.gz archive stored in Google Cloud Storage (in
	// bytes).
	ArchiveSizeBytes *float64 `json:"archiveSizeBytes,omitempty" tf:"archive_size_bytes,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/images/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The fingerprint used for optimistic locking of this resource. Used
	// internally during updates.
	// The fingerprint used for optimistic locking of this resource. Used
	// internally during updates.
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ImageParameters struct {

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Size of the image when restored onto a persistent disk .
	// Size of the image when restored onto a persistent disk (in GB).
	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	// The name of the image family to which this image belongs. You can
	// create disks by specifying an image family instead of a specific
	// image name. The image family always returns its latest image that is
	// not deprecated. The name of the image family must comply with
	// RFC1035.
	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// Structure is documented below.
	// A list of features to enable on the guest operating system.
	// Applicable only for bootable images.
	// +kubebuilder:validation:Optional
	GuestOsFeatures []GuestOsFeaturesParameters `json:"guestOsFeatures,omitempty" tf:"guest_os_features,omitempty"`

	// Labels to apply to this Image.
	// Labels to apply to this Image.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Any applicable license URI.
	// Any applicable license URI.
	// +kubebuilder:validation:Optional
	Licenses []*string `json:"licenses,omitempty" tf:"licenses,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The parameters of the raw disk image.
	// Structure is documented below.
	// The parameters of the raw disk image.
	// +kubebuilder:validation:Optional
	RawDisk []RawDiskParameters `json:"rawDisk,omitempty" tf:"raw_disk,omitempty"`

	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	// The source disk to create this image based on.
	// You must provide either this property or the
	// rawDisk.source property but not both to create an image.
	// +kubebuilder:validation:Optional
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`

	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following:
	//
	// * The selfLink URL
	// * This property
	// * The rawDisk.source URL
	// * The sourceDisk URL
	// +kubebuilder:validation:Optional
	SourceImage *string `json:"sourceImage,omitempty" tf:"source_image,omitempty"`

	// URL of the source snapshot used to create this image.
	// In order to create an image, you must provide the full or partial URL of one of the following:
	// URL of the source snapshot used to create this image.
	//
	// In order to create an image, you must provide the full or partial URL of one of the following:
	//
	// * The selfLink URL
	// * This property
	// * The sourceImage URL
	// * The rawDisk.source URL
	// * The sourceDisk URL
	// +kubebuilder:validation:Optional
	SourceSnapshot *string `json:"sourceSnapshot,omitempty" tf:"source_snapshot,omitempty"`
}

type RawDiskObservation struct {
}

type RawDiskParameters struct {

	// The format used to encode and transmit the block device, which
	// should be TAR. This is just a container and transmission format
	// and not a runtime format. Provided by the client when the disk
	// image is created.
	// Default value is TAR.
	// Possible values are TAR.
	// The format used to encode and transmit the block device, which
	// should be TAR. This is just a container and transmission format
	// and not a runtime format. Provided by the client when the disk
	// image is created. Default value: "TAR" Possible values: ["TAR"]
	// +kubebuilder:validation:Optional
	ContainerType *string `json:"containerType,omitempty" tf:"container_type,omitempty"`

	// An optional SHA1 checksum of the disk image before unpackaging.
	// This is provided by the client when the disk image is created.
	// An optional SHA1 checksum of the disk image before unpackaging.
	// This is provided by the client when the disk image is created.
	// +kubebuilder:validation:Optional
	Sha1 *string `json:"sha1,omitempty" tf:"sha1,omitempty"`

	// The full Google Cloud Storage URL where disk storage is stored
	// You must provide either this property or the sourceDisk property
	// but not both.
	// The full Google Cloud Storage URL where disk storage is stored
	// You must provide either this property or the sourceDisk property
	// but not both.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageParameters `json:"forProvider"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Image is the Schema for the Images API. Represents an Image resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpec   `json:"spec"`
	Status            ImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

// Repository type metadata.
var (
	Image_Kind             = "Image"
	Image_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Image_Kind}.String()
	Image_KindAPIVersion   = Image_Kind + "." + CRDGroupVersion.String()
	Image_GroupVersionKind = CRDGroupVersion.WithKind(Image_Kind)
)

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}