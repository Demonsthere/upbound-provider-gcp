// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AutoscalingPolicyInitParameters struct {

	// Basic algorithm for autoscaling.
	// Structure is documented below.
	BasicAlgorithm *BasicAlgorithmInitParameters `json:"basicAlgorithm,omitempty" tf:"basic_algorithm,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Describes how the autoscaler will operate for secondary workers.
	// Structure is documented below.
	SecondaryWorkerConfig *SecondaryWorkerConfigInitParameters `json:"secondaryWorkerConfig,omitempty" tf:"secondary_worker_config,omitempty"`

	// Describes how the autoscaler will operate for primary workers.
	// Structure is documented below.
	WorkerConfig *WorkerConfigInitParameters `json:"workerConfig,omitempty" tf:"worker_config,omitempty"`
}

type AutoscalingPolicyObservation struct {

	// Basic algorithm for autoscaling.
	// Structure is documented below.
	BasicAlgorithm *BasicAlgorithmObservation `json:"basicAlgorithm,omitempty" tf:"basic_algorithm,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The  location where the autoscaling policy should reside.
	// The default value is global.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The "resource name" of the autoscaling policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Describes how the autoscaler will operate for secondary workers.
	// Structure is documented below.
	SecondaryWorkerConfig *SecondaryWorkerConfigObservation `json:"secondaryWorkerConfig,omitempty" tf:"secondary_worker_config,omitempty"`

	// Describes how the autoscaler will operate for primary workers.
	// Structure is documented below.
	WorkerConfig *WorkerConfigObservation `json:"workerConfig,omitempty" tf:"worker_config,omitempty"`
}

type AutoscalingPolicyParameters struct {

	// Basic algorithm for autoscaling.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BasicAlgorithm *BasicAlgorithmParameters `json:"basicAlgorithm,omitempty" tf:"basic_algorithm,omitempty"`

	// The  location where the autoscaling policy should reside.
	// The default value is global.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Describes how the autoscaler will operate for secondary workers.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SecondaryWorkerConfig *SecondaryWorkerConfigParameters `json:"secondaryWorkerConfig,omitempty" tf:"secondary_worker_config,omitempty"`

	// Describes how the autoscaler will operate for primary workers.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	WorkerConfig *WorkerConfigParameters `json:"workerConfig,omitempty" tf:"worker_config,omitempty"`
}

type BasicAlgorithmInitParameters struct {

	// Duration between scaling events. A scaling period starts after the
	// update operation from the previous event has completed.
	// Bounds: [2m, 1d]. Default: 2m.
	CooldownPeriod *string `json:"cooldownPeriod,omitempty" tf:"cooldown_period,omitempty"`

	// YARN autoscaling configuration.
	// Structure is documented below.
	YarnConfig *YarnConfigInitParameters `json:"yarnConfig,omitempty" tf:"yarn_config,omitempty"`
}

type BasicAlgorithmObservation struct {

	// Duration between scaling events. A scaling period starts after the
	// update operation from the previous event has completed.
	// Bounds: [2m, 1d]. Default: 2m.
	CooldownPeriod *string `json:"cooldownPeriod,omitempty" tf:"cooldown_period,omitempty"`

	// YARN autoscaling configuration.
	// Structure is documented below.
	YarnConfig *YarnConfigObservation `json:"yarnConfig,omitempty" tf:"yarn_config,omitempty"`
}

type BasicAlgorithmParameters struct {

	// Duration between scaling events. A scaling period starts after the
	// update operation from the previous event has completed.
	// Bounds: [2m, 1d]. Default: 2m.
	// +kubebuilder:validation:Optional
	CooldownPeriod *string `json:"cooldownPeriod,omitempty" tf:"cooldown_period,omitempty"`

	// YARN autoscaling configuration.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	YarnConfig *YarnConfigParameters `json:"yarnConfig" tf:"yarn_config,omitempty"`
}

type SecondaryWorkerConfigInitParameters struct {

	// Maximum number of instances for this group. Note that by default, clusters will not use
	// secondary workers. Required for secondary workers if the minimum secondary instances is set.
	// Bounds: [minInstances, ). Defaults to 0.
	MaxInstances *float64 `json:"maxInstances,omitempty" tf:"max_instances,omitempty"`

	// Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Weight for the instance group, which is used to determine the fraction of total workers
	// in the cluster from this instance group. For example, if primary workers have weight 2,
	// and secondary workers have weight 1, the cluster will have approximately 2 primary workers
	// for each secondary worker.
	// The cluster may not reach the specified balance if constrained by min/max bounds or other
	// autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
	// primary workers will be added. The cluster can also be out of balance when created.
	// If weight is not set on any instance group, the cluster will default to equal weight for
	// all groups: the cluster will attempt to maintain an equal number of workers in each group
	// within the configured size bounds for each group. If weight is set for one group only,
	// the cluster will default to zero weight on the unset group. For example if weight is set
	// only on primary workers, the cluster will use primary workers only and no secondary workers.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type SecondaryWorkerConfigObservation struct {

	// Maximum number of instances for this group. Note that by default, clusters will not use
	// secondary workers. Required for secondary workers if the minimum secondary instances is set.
	// Bounds: [minInstances, ). Defaults to 0.
	MaxInstances *float64 `json:"maxInstances,omitempty" tf:"max_instances,omitempty"`

	// Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Weight for the instance group, which is used to determine the fraction of total workers
	// in the cluster from this instance group. For example, if primary workers have weight 2,
	// and secondary workers have weight 1, the cluster will have approximately 2 primary workers
	// for each secondary worker.
	// The cluster may not reach the specified balance if constrained by min/max bounds or other
	// autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
	// primary workers will be added. The cluster can also be out of balance when created.
	// If weight is not set on any instance group, the cluster will default to equal weight for
	// all groups: the cluster will attempt to maintain an equal number of workers in each group
	// within the configured size bounds for each group. If weight is set for one group only,
	// the cluster will default to zero weight on the unset group. For example if weight is set
	// only on primary workers, the cluster will use primary workers only and no secondary workers.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type SecondaryWorkerConfigParameters struct {

	// Maximum number of instances for this group. Note that by default, clusters will not use
	// secondary workers. Required for secondary workers if the minimum secondary instances is set.
	// Bounds: [minInstances, ). Defaults to 0.
	// +kubebuilder:validation:Optional
	MaxInstances *float64 `json:"maxInstances,omitempty" tf:"max_instances,omitempty"`

	// Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.
	// +kubebuilder:validation:Optional
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Weight for the instance group, which is used to determine the fraction of total workers
	// in the cluster from this instance group. For example, if primary workers have weight 2,
	// and secondary workers have weight 1, the cluster will have approximately 2 primary workers
	// for each secondary worker.
	// The cluster may not reach the specified balance if constrained by min/max bounds or other
	// autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
	// primary workers will be added. The cluster can also be out of balance when created.
	// If weight is not set on any instance group, the cluster will default to equal weight for
	// all groups: the cluster will attempt to maintain an equal number of workers in each group
	// within the configured size bounds for each group. If weight is set for one group only,
	// the cluster will default to zero weight on the unset group. For example if weight is set
	// only on primary workers, the cluster will use primary workers only and no secondary workers.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type WorkerConfigInitParameters struct {

	// Maximum number of instances for this group.
	MaxInstances *float64 `json:"maxInstances,omitempty" tf:"max_instances,omitempty"`

	// Minimum number of instances for this group. Bounds: [2, maxInstances]. Defaults to 2.
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Weight for the instance group, which is used to determine the fraction of total workers
	// in the cluster from this instance group. For example, if primary workers have weight 2,
	// and secondary workers have weight 1, the cluster will have approximately 2 primary workers
	// for each secondary worker.
	// The cluster may not reach the specified balance if constrained by min/max bounds or other
	// autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
	// primary workers will be added. The cluster can also be out of balance when created.
	// If weight is not set on any instance group, the cluster will default to equal weight for
	// all groups: the cluster will attempt to maintain an equal number of workers in each group
	// within the configured size bounds for each group. If weight is set for one group only,
	// the cluster will default to zero weight on the unset group. For example if weight is set
	// only on primary workers, the cluster will use primary workers only and no secondary workers.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type WorkerConfigObservation struct {

	// Maximum number of instances for this group.
	MaxInstances *float64 `json:"maxInstances,omitempty" tf:"max_instances,omitempty"`

	// Minimum number of instances for this group. Bounds: [2, maxInstances]. Defaults to 2.
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Weight for the instance group, which is used to determine the fraction of total workers
	// in the cluster from this instance group. For example, if primary workers have weight 2,
	// and secondary workers have weight 1, the cluster will have approximately 2 primary workers
	// for each secondary worker.
	// The cluster may not reach the specified balance if constrained by min/max bounds or other
	// autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
	// primary workers will be added. The cluster can also be out of balance when created.
	// If weight is not set on any instance group, the cluster will default to equal weight for
	// all groups: the cluster will attempt to maintain an equal number of workers in each group
	// within the configured size bounds for each group. If weight is set for one group only,
	// the cluster will default to zero weight on the unset group. For example if weight is set
	// only on primary workers, the cluster will use primary workers only and no secondary workers.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type WorkerConfigParameters struct {

	// Maximum number of instances for this group.
	// +kubebuilder:validation:Optional
	MaxInstances *float64 `json:"maxInstances" tf:"max_instances,omitempty"`

	// Minimum number of instances for this group. Bounds: [2, maxInstances]. Defaults to 2.
	// +kubebuilder:validation:Optional
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Weight for the instance group, which is used to determine the fraction of total workers
	// in the cluster from this instance group. For example, if primary workers have weight 2,
	// and secondary workers have weight 1, the cluster will have approximately 2 primary workers
	// for each secondary worker.
	// The cluster may not reach the specified balance if constrained by min/max bounds or other
	// autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
	// primary workers will be added. The cluster can also be out of balance when created.
	// If weight is not set on any instance group, the cluster will default to equal weight for
	// all groups: the cluster will attempt to maintain an equal number of workers in each group
	// within the configured size bounds for each group. If weight is set for one group only,
	// the cluster will default to zero weight on the unset group. For example if weight is set
	// only on primary workers, the cluster will use primary workers only and no secondary workers.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type YarnConfigInitParameters struct {

	// Timeout for YARN graceful decommissioning of Node Managers. Specifies the
	// duration to wait for jobs to complete before forcefully removing workers
	// (and potentially interrupting jobs). Only applicable to downscaling operations.
	// Bounds: [0s, 1d].
	GracefulDecommissionTimeout *string `json:"gracefulDecommissionTimeout,omitempty" tf:"graceful_decommission_timeout,omitempty"`

	// Fraction of average pending memory in the last cooldown period for which to
	// remove workers. A scale-down factor of 1 will result in scaling down so that there
	// is no available memory remaining after the update (more aggressive scaling).
	// A scale-down factor of 0 disables removing workers, which can be beneficial for
	// autoscaling a single job.
	// Bounds: [0.0, 1.0].
	ScaleDownFactor *float64 `json:"scaleDownFactor,omitempty" tf:"scale_down_factor,omitempty"`

	// Minimum scale-down threshold as a fraction of total cluster size before scaling occurs.
	// For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must
	// recommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0
	// means the autoscaler will scale down on any recommended change.
	// Bounds: [0.0, 1.0]. Default: 0.0.
	ScaleDownMinWorkerFraction *float64 `json:"scaleDownMinWorkerFraction,omitempty" tf:"scale_down_min_worker_fraction,omitempty"`

	// Fraction of average pending memory in the last cooldown period for which to
	// add workers. A scale-up factor of 1.0 will result in scaling up so that there
	// is no pending memory remaining after the update (more aggressive scaling).
	// A scale-up factor closer to 0 will result in a smaller magnitude of scaling up
	// (less aggressive scaling).
	// Bounds: [0.0, 1.0].
	ScaleUpFactor *float64 `json:"scaleUpFactor,omitempty" tf:"scale_up_factor,omitempty"`

	// Minimum scale-up threshold as a fraction of total cluster size before scaling
	// occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler
	// must recommend at least a 2-worker scale-up for the cluster to scale. A threshold of
	// 0 means the autoscaler will scale up on any recommended change.
	// Bounds: [0.0, 1.0]. Default: 0.0.
	ScaleUpMinWorkerFraction *float64 `json:"scaleUpMinWorkerFraction,omitempty" tf:"scale_up_min_worker_fraction,omitempty"`
}

type YarnConfigObservation struct {

	// Timeout for YARN graceful decommissioning of Node Managers. Specifies the
	// duration to wait for jobs to complete before forcefully removing workers
	// (and potentially interrupting jobs). Only applicable to downscaling operations.
	// Bounds: [0s, 1d].
	GracefulDecommissionTimeout *string `json:"gracefulDecommissionTimeout,omitempty" tf:"graceful_decommission_timeout,omitempty"`

	// Fraction of average pending memory in the last cooldown period for which to
	// remove workers. A scale-down factor of 1 will result in scaling down so that there
	// is no available memory remaining after the update (more aggressive scaling).
	// A scale-down factor of 0 disables removing workers, which can be beneficial for
	// autoscaling a single job.
	// Bounds: [0.0, 1.0].
	ScaleDownFactor *float64 `json:"scaleDownFactor,omitempty" tf:"scale_down_factor,omitempty"`

	// Minimum scale-down threshold as a fraction of total cluster size before scaling occurs.
	// For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must
	// recommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0
	// means the autoscaler will scale down on any recommended change.
	// Bounds: [0.0, 1.0]. Default: 0.0.
	ScaleDownMinWorkerFraction *float64 `json:"scaleDownMinWorkerFraction,omitempty" tf:"scale_down_min_worker_fraction,omitempty"`

	// Fraction of average pending memory in the last cooldown period for which to
	// add workers. A scale-up factor of 1.0 will result in scaling up so that there
	// is no pending memory remaining after the update (more aggressive scaling).
	// A scale-up factor closer to 0 will result in a smaller magnitude of scaling up
	// (less aggressive scaling).
	// Bounds: [0.0, 1.0].
	ScaleUpFactor *float64 `json:"scaleUpFactor,omitempty" tf:"scale_up_factor,omitempty"`

	// Minimum scale-up threshold as a fraction of total cluster size before scaling
	// occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler
	// must recommend at least a 2-worker scale-up for the cluster to scale. A threshold of
	// 0 means the autoscaler will scale up on any recommended change.
	// Bounds: [0.0, 1.0]. Default: 0.0.
	ScaleUpMinWorkerFraction *float64 `json:"scaleUpMinWorkerFraction,omitempty" tf:"scale_up_min_worker_fraction,omitempty"`
}

type YarnConfigParameters struct {

	// Timeout for YARN graceful decommissioning of Node Managers. Specifies the
	// duration to wait for jobs to complete before forcefully removing workers
	// (and potentially interrupting jobs). Only applicable to downscaling operations.
	// Bounds: [0s, 1d].
	// +kubebuilder:validation:Optional
	GracefulDecommissionTimeout *string `json:"gracefulDecommissionTimeout" tf:"graceful_decommission_timeout,omitempty"`

	// Fraction of average pending memory in the last cooldown period for which to
	// remove workers. A scale-down factor of 1 will result in scaling down so that there
	// is no available memory remaining after the update (more aggressive scaling).
	// A scale-down factor of 0 disables removing workers, which can be beneficial for
	// autoscaling a single job.
	// Bounds: [0.0, 1.0].
	// +kubebuilder:validation:Optional
	ScaleDownFactor *float64 `json:"scaleDownFactor" tf:"scale_down_factor,omitempty"`

	// Minimum scale-down threshold as a fraction of total cluster size before scaling occurs.
	// For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must
	// recommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0
	// means the autoscaler will scale down on any recommended change.
	// Bounds: [0.0, 1.0]. Default: 0.0.
	// +kubebuilder:validation:Optional
	ScaleDownMinWorkerFraction *float64 `json:"scaleDownMinWorkerFraction,omitempty" tf:"scale_down_min_worker_fraction,omitempty"`

	// Fraction of average pending memory in the last cooldown period for which to
	// add workers. A scale-up factor of 1.0 will result in scaling up so that there
	// is no pending memory remaining after the update (more aggressive scaling).
	// A scale-up factor closer to 0 will result in a smaller magnitude of scaling up
	// (less aggressive scaling).
	// Bounds: [0.0, 1.0].
	// +kubebuilder:validation:Optional
	ScaleUpFactor *float64 `json:"scaleUpFactor" tf:"scale_up_factor,omitempty"`

	// Minimum scale-up threshold as a fraction of total cluster size before scaling
	// occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler
	// must recommend at least a 2-worker scale-up for the cluster to scale. A threshold of
	// 0 means the autoscaler will scale up on any recommended change.
	// Bounds: [0.0, 1.0]. Default: 0.0.
	// +kubebuilder:validation:Optional
	ScaleUpMinWorkerFraction *float64 `json:"scaleUpMinWorkerFraction,omitempty" tf:"scale_up_min_worker_fraction,omitempty"`
}

// AutoscalingPolicySpec defines the desired state of AutoscalingPolicy
type AutoscalingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutoscalingPolicyParameters `json:"forProvider"`
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
	InitProvider AutoscalingPolicyInitParameters `json:"initProvider,omitempty"`
}

// AutoscalingPolicyStatus defines the observed state of AutoscalingPolicy.
type AutoscalingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutoscalingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AutoscalingPolicy is the Schema for the AutoscalingPolicys API. Describes an autoscaling policy for Dataproc cluster autoscaler.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AutoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingPolicySpec   `json:"spec"`
	Status            AutoscalingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingPolicyList contains a list of AutoscalingPolicys
type AutoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingPolicy `json:"items"`
}

// Repository type metadata.
var (
	AutoscalingPolicy_Kind             = "AutoscalingPolicy"
	AutoscalingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutoscalingPolicy_Kind}.String()
	AutoscalingPolicy_KindAPIVersion   = AutoscalingPolicy_Kind + "." + CRDGroupVersion.String()
	AutoscalingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(AutoscalingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&AutoscalingPolicy{}, &AutoscalingPolicyList{})
}