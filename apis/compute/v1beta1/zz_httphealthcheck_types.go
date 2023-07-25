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

type HTTPHealthCheckInitParameters_2 struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the public IP on behalf of which this
	// health check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HTTPHealthCheckObservation_2 struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the public IP on behalf of which this
	// health check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/httpHealthChecks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HTTPHealthCheckParameters_2 struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the public IP on behalf of which this
	// health check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

// HTTPHealthCheckSpec defines the desired state of HTTPHealthCheck
type HTTPHealthCheckSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HTTPHealthCheckParameters_2 `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider HTTPHealthCheckInitParameters_2 `json:"initProvider,omitempty"`
}

// HTTPHealthCheckStatus defines the observed state of HTTPHealthCheck.
type HTTPHealthCheckStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HTTPHealthCheckObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPHealthCheck is the Schema for the HTTPHealthChecks API. An HttpHealthCheck resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type HTTPHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HTTPHealthCheckSpec   `json:"spec"`
	Status            HTTPHealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPHealthCheckList contains a list of HTTPHealthChecks
type HTTPHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPHealthCheck `json:"items"`
}

// Repository type metadata.
var (
	HTTPHealthCheck_Kind             = "HTTPHealthCheck"
	HTTPHealthCheck_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HTTPHealthCheck_Kind}.String()
	HTTPHealthCheck_KindAPIVersion   = HTTPHealthCheck_Kind + "." + CRDGroupVersion.String()
	HTTPHealthCheck_GroupVersionKind = CRDGroupVersion.WithKind(HTTPHealthCheck_Kind)
)

func init() {
	SchemeBuilder.Register(&HTTPHealthCheck{}, &HTTPHealthCheckList{})
}
