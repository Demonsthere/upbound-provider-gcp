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

type AccessLevelConditionDevicePolicyInitParameters struct {

	// A list of allowed device management levels.
	// An empty list allows all management levels.
	// Each value may be one of: MANAGEMENT_UNSPECIFIED, NONE, BASIC, COMPLETE.
	AllowedDeviceManagementLevels []*string `json:"allowedDeviceManagementLevels,omitempty" tf:"allowed_device_management_levels,omitempty"`

	// A list of allowed encryptions statuses.
	// An empty list allows all statuses.
	// Each value may be one of: ENCRYPTION_UNSPECIFIED, ENCRYPTION_UNSUPPORTED, UNENCRYPTED, ENCRYPTED.
	AllowedEncryptionStatuses []*string `json:"allowedEncryptionStatuses,omitempty" tf:"allowed_encryption_statuses,omitempty"`

	// A list of allowed OS versions.
	// An empty list allows all types and all versions.
	// Structure is documented below.
	OsConstraints []DevicePolicyOsConstraintsInitParameters `json:"osConstraints,omitempty" tf:"os_constraints,omitempty"`

	// Whether the device needs to be approved by the customer admin.
	RequireAdminApproval *bool `json:"requireAdminApproval,omitempty" tf:"require_admin_approval,omitempty"`

	// Whether the device needs to be corp owned.
	RequireCorpOwned *bool `json:"requireCorpOwned,omitempty" tf:"require_corp_owned,omitempty"`

	// Whether or not screenlock is required for the DevicePolicy
	// to be true. Defaults to false.
	RequireScreenLock *bool `json:"requireScreenLock,omitempty" tf:"require_screen_lock,omitempty"`
}

type AccessLevelConditionDevicePolicyObservation struct {

	// A list of allowed device management levels.
	// An empty list allows all management levels.
	// Each value may be one of: MANAGEMENT_UNSPECIFIED, NONE, BASIC, COMPLETE.
	AllowedDeviceManagementLevels []*string `json:"allowedDeviceManagementLevels,omitempty" tf:"allowed_device_management_levels,omitempty"`

	// A list of allowed encryptions statuses.
	// An empty list allows all statuses.
	// Each value may be one of: ENCRYPTION_UNSPECIFIED, ENCRYPTION_UNSUPPORTED, UNENCRYPTED, ENCRYPTED.
	AllowedEncryptionStatuses []*string `json:"allowedEncryptionStatuses,omitempty" tf:"allowed_encryption_statuses,omitempty"`

	// A list of allowed OS versions.
	// An empty list allows all types and all versions.
	// Structure is documented below.
	OsConstraints []DevicePolicyOsConstraintsObservation `json:"osConstraints,omitempty" tf:"os_constraints,omitempty"`

	// Whether the device needs to be approved by the customer admin.
	RequireAdminApproval *bool `json:"requireAdminApproval,omitempty" tf:"require_admin_approval,omitempty"`

	// Whether the device needs to be corp owned.
	RequireCorpOwned *bool `json:"requireCorpOwned,omitempty" tf:"require_corp_owned,omitempty"`

	// Whether or not screenlock is required for the DevicePolicy
	// to be true. Defaults to false.
	RequireScreenLock *bool `json:"requireScreenLock,omitempty" tf:"require_screen_lock,omitempty"`
}

type AccessLevelConditionDevicePolicyParameters struct {

	// A list of allowed device management levels.
	// An empty list allows all management levels.
	// Each value may be one of: MANAGEMENT_UNSPECIFIED, NONE, BASIC, COMPLETE.
	AllowedDeviceManagementLevels []*string `json:"allowedDeviceManagementLevels,omitempty" tf:"allowed_device_management_levels,omitempty"`

	// A list of allowed encryptions statuses.
	// An empty list allows all statuses.
	// Each value may be one of: ENCRYPTION_UNSPECIFIED, ENCRYPTION_UNSUPPORTED, UNENCRYPTED, ENCRYPTED.
	AllowedEncryptionStatuses []*string `json:"allowedEncryptionStatuses,omitempty" tf:"allowed_encryption_statuses,omitempty"`

	// A list of allowed OS versions.
	// An empty list allows all types and all versions.
	// Structure is documented below.
	OsConstraints []DevicePolicyOsConstraintsParameters `json:"osConstraints,omitempty" tf:"os_constraints,omitempty"`

	// Whether the device needs to be approved by the customer admin.
	RequireAdminApproval *bool `json:"requireAdminApproval,omitempty" tf:"require_admin_approval,omitempty"`

	// Whether the device needs to be corp owned.
	RequireCorpOwned *bool `json:"requireCorpOwned,omitempty" tf:"require_corp_owned,omitempty"`

	// Whether or not screenlock is required for the DevicePolicy
	// to be true. Defaults to false.
	RequireScreenLock *bool `json:"requireScreenLock,omitempty" tf:"require_screen_lock,omitempty"`
}

type AccessLevelConditionInitParameters struct {

	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy []AccessLevelConditionDevicePolicyInitParameters `json:"devicePolicy,omitempty" tf:"device_policy,omitempty"`

	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IPSubnetworks []*string `json:"ipSubnetworks,omitempty" tf:"ip_subnetworks,omitempty"`

	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: user:{emailid}, serviceAccount:{emailid}
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels []*string `json:"requiredAccessLevels,omitempty" tf:"required_access_levels,omitempty"`
}

type AccessLevelConditionObservation struct {

	// The name of the Access Level to add this condition to.
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`

	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy []AccessLevelConditionDevicePolicyObservation `json:"devicePolicy,omitempty" tf:"device_policy,omitempty"`

	// an identifier for the resource with format {{access_level}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IPSubnetworks []*string `json:"ipSubnetworks,omitempty" tf:"ip_subnetworks,omitempty"`

	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: user:{emailid}, serviceAccount:{emailid}
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels []*string `json:"requiredAccessLevels,omitempty" tf:"required_access_levels,omitempty"`
}

type AccessLevelConditionParameters struct {

	// The name of the Access Level to add this condition to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/accesscontextmanager/v1beta1.AccessLevel
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`

	// Reference to a AccessLevel in accesscontextmanager to populate accessLevel.
	// +kubebuilder:validation:Optional
	AccessLevelRef *v1.Reference `json:"accessLevelRef,omitempty" tf:"-"`

	// Selector for a AccessLevel in accesscontextmanager to populate accessLevel.
	// +kubebuilder:validation:Optional
	AccessLevelSelector *v1.Selector `json:"accessLevelSelector,omitempty" tf:"-"`

	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// Structure is documented below.
	DevicePolicy []AccessLevelConditionDevicePolicyParameters `json:"devicePolicy,omitempty" tf:"device_policy,omitempty"`

	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IPSubnetworks []*string `json:"ipSubnetworks,omitempty" tf:"ip_subnetworks,omitempty"`

	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: user:{emailid}, serviceAccount:{emailid}
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	RequiredAccessLevels []*string `json:"requiredAccessLevels,omitempty" tf:"required_access_levels,omitempty"`
}

type DevicePolicyOsConstraintsInitParameters struct {

	// The minimum allowed OS version. If not set, any version
	// of this OS satisfies the constraint.
	// Format: "major.minor.patch" such as "10.5.301", "9.2.1".
	MinimumVersion *string `json:"minimumVersion,omitempty" tf:"minimum_version,omitempty"`

	// The operating system type of the device.
	// Possible values are: OS_UNSPECIFIED, DESKTOP_MAC, DESKTOP_WINDOWS, DESKTOP_LINUX, DESKTOP_CHROME_OS, ANDROID, IOS.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`
}

type DevicePolicyOsConstraintsObservation struct {

	// The minimum allowed OS version. If not set, any version
	// of this OS satisfies the constraint.
	// Format: "major.minor.patch" such as "10.5.301", "9.2.1".
	MinimumVersion *string `json:"minimumVersion,omitempty" tf:"minimum_version,omitempty"`

	// The operating system type of the device.
	// Possible values are: OS_UNSPECIFIED, DESKTOP_MAC, DESKTOP_WINDOWS, DESKTOP_LINUX, DESKTOP_CHROME_OS, ANDROID, IOS.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`
}

type DevicePolicyOsConstraintsParameters struct {

	// The minimum allowed OS version. If not set, any version
	// of this OS satisfies the constraint.
	// Format: "major.minor.patch" such as "10.5.301", "9.2.1".
	MinimumVersion *string `json:"minimumVersion,omitempty" tf:"minimum_version,omitempty"`

	// The operating system type of the device.
	// Possible values are: OS_UNSPECIFIED, DESKTOP_MAC, DESKTOP_WINDOWS, DESKTOP_LINUX, DESKTOP_CHROME_OS, ANDROID, IOS.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`
}

// AccessLevelConditionSpec defines the desired state of AccessLevelCondition
type AccessLevelConditionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessLevelConditionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider AccessLevelConditionInitParameters `json:"initProvider,omitempty"`
}

// AccessLevelConditionStatus defines the observed state of AccessLevelCondition.
type AccessLevelConditionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessLevelConditionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccessLevelCondition is the Schema for the AccessLevelConditions API. Allows configuring a single access level condition to be appended to an access level's conditions.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AccessLevelCondition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessLevelConditionSpec   `json:"spec"`
	Status            AccessLevelConditionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessLevelConditionList contains a list of AccessLevelConditions
type AccessLevelConditionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessLevelCondition `json:"items"`
}

// Repository type metadata.
var (
	AccessLevelCondition_Kind             = "AccessLevelCondition"
	AccessLevelCondition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessLevelCondition_Kind}.String()
	AccessLevelCondition_KindAPIVersion   = AccessLevelCondition_Kind + "." + CRDGroupVersion.String()
	AccessLevelCondition_GroupVersionKind = CRDGroupVersion.WithKind(AccessLevelCondition_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessLevelCondition{}, &AccessLevelConditionList{})
}
