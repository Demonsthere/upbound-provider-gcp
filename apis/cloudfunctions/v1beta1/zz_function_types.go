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

type EventTriggerInitParameters struct {

	// The type of event to observe. For example: "google.storage.object.finalize".
	// See the documentation on calling Cloud Functions for a
	// full reference of accepted triggers.
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// Specifies policy for failed executions. Structure is documented below.
	FailurePolicy []FailurePolicyInitParameters `json:"failurePolicy,omitempty" tf:"failure_policy,omitempty"`

	// Required. The name or partial URI of the resource from
	// which to observe events. For example, "myBucket" or "projects/my-project/topics/my-topic"
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`
}

type EventTriggerObservation struct {

	// The type of event to observe. For example: "google.storage.object.finalize".
	// See the documentation on calling Cloud Functions for a
	// full reference of accepted triggers.
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// Specifies policy for failed executions. Structure is documented below.
	FailurePolicy []FailurePolicyObservation `json:"failurePolicy,omitempty" tf:"failure_policy,omitempty"`

	// Required. The name or partial URI of the resource from
	// which to observe events. For example, "myBucket" or "projects/my-project/topics/my-topic"
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`
}

type EventTriggerParameters struct {

	// The type of event to observe. For example: "google.storage.object.finalize".
	// See the documentation on calling Cloud Functions for a
	// full reference of accepted triggers.
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// Specifies policy for failed executions. Structure is documented below.
	FailurePolicy []FailurePolicyParameters `json:"failurePolicy,omitempty" tf:"failure_policy,omitempty"`

	// Required. The name or partial URI of the resource from
	// which to observe events. For example, "myBucket" or "projects/my-project/topics/my-topic"
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`
}

type FailurePolicyInitParameters struct {

	// Whether the function should be retried on failure. Defaults to false.
	Retry *bool `json:"retry,omitempty" tf:"retry,omitempty"`
}

type FailurePolicyObservation struct {

	// Whether the function should be retried on failure. Defaults to false.
	Retry *bool `json:"retry,omitempty" tf:"retry,omitempty"`
}

type FailurePolicyParameters struct {

	// Whether the function should be retried on failure. Defaults to false.
	Retry *bool `json:"retry,omitempty" tf:"retry,omitempty"`
}

type FunctionInitParameters struct {

	// Memory (in MB), available to the function. Default value is 256. Possible values include 128, 256, 512, 1024, etc.
	AvailableMemoryMb *float64 `json:"availableMemoryMb,omitempty" tf:"available_memory_mb,omitempty"`

	// A set of key/value environment variable pairs available during build time.
	BuildEnvironmentVariables map[string]string `json:"buildEnvironmentVariables,omitempty" tf:"build_environment_variables,omitempty"`

	BuildWorkerPool *string `json:"buildWorkerPool,omitempty" tf:"build_worker_pool,omitempty"`

	// Description of the function.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY.
	DockerRegistry *string `json:"dockerRegistry,omitempty" tf:"docker_registry,omitempty"`

	// User managed repository created in Artifact Registry optionally with a customer managed encryption key. If specified, deployments will use Artifact Registry. This is the repository to which the function docker image will be pushed after it is built by Cloud Build. If unspecified, Container Registry will be used by default, unless specified otherwise by other means.
	DockerRepository *string `json:"dockerRepository,omitempty" tf:"docker_repository,omitempty"`

	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint *string `json:"entryPoint,omitempty" tf:"entry_point,omitempty"`

	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with trigger_http.
	EventTrigger []EventTriggerInitParameters `json:"eventTrigger,omitempty" tf:"event_trigger,omitempty"`

	// The security level for the function. The following options are available:
	HTTPSTriggerSecurityLevel *string `json:"httpsTriggerSecurityLevel,omitempty" tf:"https_trigger_security_level,omitempty"`

	// URL which triggers function execution. Returned only if trigger_http is used.
	HTTPSTriggerURL *string `json:"httpsTriggerUrl,omitempty" tf:"https_trigger_url,omitempty"`

	// String value that controls what traffic can reach the function. Allowed values are ALLOW_ALL, ALLOW_INTERNAL_AND_GCLB and ALLOW_INTERNAL_ONLY. Check ingress documentation to see the impact of each settings value. Changes to this field will recreate the cloud function.
	IngressSettings *string `json:"ingressSettings,omitempty" tf:"ingress_settings,omitempty"`

	// Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources. It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
	// If specified, you must also provide an artifact registry repository using the docker_repository field that was created with the same KMS crypto key. Before deploying, please complete all pre-requisites described in https://cloud.google.com/functions/docs/securing/cmek#granting_service_accounts_access_to_the_key
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances *float64 `json:"maxInstances,omitempty" tf:"max_instances,omitempty"`

	// The limit on the minimum number of function instances that may coexist at a given time.
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Project of the function. If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The runtime in which the function is going to run.
	// Eg. "nodejs16", "python39", "dotnet3", "go116", "java11", "ruby30", "php74", etc. Check the official doc for the up-to-date list.
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Secret environment variables configuration. Structure is documented below.
	SecretEnvironmentVariables []SecretEnvironmentVariablesInitParameters `json:"secretEnvironmentVariables,omitempty" tf:"secret_environment_variables,omitempty"`

	// Secret volumes configuration. Structure is documented below.
	SecretVolumes []SecretVolumesInitParameters `json:"secretVolumes,omitempty" tf:"secret_volumes,omitempty"`

	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside source_archive_bucket or source_archive_object. Structure is documented below. It must match the pattern projects/{project}/locations/{location}/repositories/{repository}.*
	SourceRepository []SourceRepositoryInitParameters `json:"sourceRepository,omitempty" tf:"source_repository,omitempty"`

	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as https_trigger_url. Cannot be used with event_trigger.
	TriggerHTTP *bool `json:"triggerHttp,omitempty" tf:"trigger_http,omitempty"`

	// The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is projects/*/locations/*/connectors/*.
	VPCConnector *string `json:"vpcConnector,omitempty" tf:"vpc_connector,omitempty"`

	// The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are ALL_TRAFFIC and PRIVATE_RANGES_ONLY. Defaults to PRIVATE_RANGES_ONLY. If unset, this field preserves the previously set value.
	VPCConnectorEgressSettings *string `json:"vpcConnectorEgressSettings,omitempty" tf:"vpc_connector_egress_settings,omitempty"`
}

type FunctionObservation struct {

	// Memory (in MB), available to the function. Default value is 256. Possible values include 128, 256, 512, 1024, etc.
	AvailableMemoryMb *float64 `json:"availableMemoryMb,omitempty" tf:"available_memory_mb,omitempty"`

	// A set of key/value environment variable pairs available during build time.
	BuildEnvironmentVariables map[string]string `json:"buildEnvironmentVariables,omitempty" tf:"build_environment_variables,omitempty"`

	BuildWorkerPool *string `json:"buildWorkerPool,omitempty" tf:"build_worker_pool,omitempty"`

	// Description of the function.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY.
	DockerRegistry *string `json:"dockerRegistry,omitempty" tf:"docker_registry,omitempty"`

	// User managed repository created in Artifact Registry optionally with a customer managed encryption key. If specified, deployments will use Artifact Registry. This is the repository to which the function docker image will be pushed after it is built by Cloud Build. If unspecified, Container Registry will be used by default, unless specified otherwise by other means.
	DockerRepository *string `json:"dockerRepository,omitempty" tf:"docker_repository,omitempty"`

	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint *string `json:"entryPoint,omitempty" tf:"entry_point,omitempty"`

	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with trigger_http.
	EventTrigger []EventTriggerObservation `json:"eventTrigger,omitempty" tf:"event_trigger,omitempty"`

	// The security level for the function. The following options are available:
	HTTPSTriggerSecurityLevel *string `json:"httpsTriggerSecurityLevel,omitempty" tf:"https_trigger_security_level,omitempty"`

	// URL which triggers function execution. Returned only if trigger_http is used.
	HTTPSTriggerURL *string `json:"httpsTriggerUrl,omitempty" tf:"https_trigger_url,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// String value that controls what traffic can reach the function. Allowed values are ALLOW_ALL, ALLOW_INTERNAL_AND_GCLB and ALLOW_INTERNAL_ONLY. Check ingress documentation to see the impact of each settings value. Changes to this field will recreate the cloud function.
	IngressSettings *string `json:"ingressSettings,omitempty" tf:"ingress_settings,omitempty"`

	// Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources. It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
	// If specified, you must also provide an artifact registry repository using the docker_repository field that was created with the same KMS crypto key. Before deploying, please complete all pre-requisites described in https://cloud.google.com/functions/docs/securing/cmek#granting_service_accounts_access_to_the_key
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances *float64 `json:"maxInstances,omitempty" tf:"max_instances,omitempty"`

	// The limit on the minimum number of function instances that may coexist at a given time.
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Project of the function. If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region of function. If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The runtime in which the function is going to run.
	// Eg. "nodejs16", "python39", "dotnet3", "go116", "java11", "ruby30", "php74", etc. Check the official doc for the up-to-date list.
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Secret environment variables configuration. Structure is documented below.
	SecretEnvironmentVariables []SecretEnvironmentVariablesObservation `json:"secretEnvironmentVariables,omitempty" tf:"secret_environment_variables,omitempty"`

	// Secret volumes configuration. Structure is documented below.
	SecretVolumes []SecretVolumesObservation `json:"secretVolumes,omitempty" tf:"secret_volumes,omitempty"`

	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket *string `json:"sourceArchiveBucket,omitempty" tf:"source_archive_bucket,omitempty"`

	// The source archive object (file) in archive bucket.
	SourceArchiveObject *string `json:"sourceArchiveObject,omitempty" tf:"source_archive_object,omitempty"`

	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside source_archive_bucket or source_archive_object. Structure is documented below. It must match the pattern projects/{project}/locations/{location}/repositories/{repository}.*
	SourceRepository []SourceRepositoryObservation `json:"sourceRepository,omitempty" tf:"source_repository,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as https_trigger_url. Cannot be used with event_trigger.
	TriggerHTTP *bool `json:"triggerHttp,omitempty" tf:"trigger_http,omitempty"`

	// The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is projects/*/locations/*/connectors/*.
	VPCConnector *string `json:"vpcConnector,omitempty" tf:"vpc_connector,omitempty"`

	// The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are ALL_TRAFFIC and PRIVATE_RANGES_ONLY. Defaults to PRIVATE_RANGES_ONLY. If unset, this field preserves the previously set value.
	VPCConnectorEgressSettings *string `json:"vpcConnectorEgressSettings,omitempty" tf:"vpc_connector_egress_settings,omitempty"`
}

type FunctionParameters struct {

	// Memory (in MB), available to the function. Default value is 256. Possible values include 128, 256, 512, 1024, etc.
	AvailableMemoryMb *float64 `json:"availableMemoryMb,omitempty" tf:"available_memory_mb,omitempty"`

	// A set of key/value environment variable pairs available during build time.
	BuildEnvironmentVariables map[string]string `json:"buildEnvironmentVariables,omitempty" tf:"build_environment_variables,omitempty"`

	BuildWorkerPool *string `json:"buildWorkerPool,omitempty" tf:"build_worker_pool,omitempty"`

	// Description of the function.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY.
	DockerRegistry *string `json:"dockerRegistry,omitempty" tf:"docker_registry,omitempty"`

	// User managed repository created in Artifact Registry optionally with a customer managed encryption key. If specified, deployments will use Artifact Registry. This is the repository to which the function docker image will be pushed after it is built by Cloud Build. If unspecified, Container Registry will be used by default, unless specified otherwise by other means.
	DockerRepository *string `json:"dockerRepository,omitempty" tf:"docker_repository,omitempty"`

	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint *string `json:"entryPoint,omitempty" tf:"entry_point,omitempty"`

	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with trigger_http.
	EventTrigger []EventTriggerParameters `json:"eventTrigger,omitempty" tf:"event_trigger,omitempty"`

	// The security level for the function. The following options are available:
	HTTPSTriggerSecurityLevel *string `json:"httpsTriggerSecurityLevel,omitempty" tf:"https_trigger_security_level,omitempty"`

	// URL which triggers function execution. Returned only if trigger_http is used.
	HTTPSTriggerURL *string `json:"httpsTriggerUrl,omitempty" tf:"https_trigger_url,omitempty"`

	// String value that controls what traffic can reach the function. Allowed values are ALLOW_ALL, ALLOW_INTERNAL_AND_GCLB and ALLOW_INTERNAL_ONLY. Check ingress documentation to see the impact of each settings value. Changes to this field will recreate the cloud function.
	IngressSettings *string `json:"ingressSettings,omitempty" tf:"ingress_settings,omitempty"`

	// Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources. It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
	// If specified, you must also provide an artifact registry repository using the docker_repository field that was created with the same KMS crypto key. Before deploying, please complete all pre-requisites described in https://cloud.google.com/functions/docs/securing/cmek#granting_service_accounts_access_to_the_key
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances *float64 `json:"maxInstances,omitempty" tf:"max_instances,omitempty"`

	// The limit on the minimum number of function instances that may coexist at a given time.
	MinInstances *float64 `json:"minInstances,omitempty" tf:"min_instances,omitempty"`

	// Project of the function. If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region of function. If it is not provided, the provider region is used.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The runtime in which the function is going to run.
	// Eg. "nodejs16", "python39", "dotnet3", "go116", "java11", "ruby30", "php74", etc. Check the official doc for the up-to-date list.
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Secret environment variables configuration. Structure is documented below.
	SecretEnvironmentVariables []SecretEnvironmentVariablesParameters `json:"secretEnvironmentVariables,omitempty" tf:"secret_environment_variables,omitempty"`

	// Secret volumes configuration. Structure is documented below.
	SecretVolumes []SecretVolumesParameters `json:"secretVolumes,omitempty" tf:"secret_volumes,omitempty"`

	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// The GCS bucket containing the zip archive which contains the function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	SourceArchiveBucket *string `json:"sourceArchiveBucket,omitempty" tf:"source_archive_bucket,omitempty"`

	// Reference to a Bucket in storage to populate sourceArchiveBucket.
	// +kubebuilder:validation:Optional
	SourceArchiveBucketRef *v1.Reference `json:"sourceArchiveBucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate sourceArchiveBucket.
	// +kubebuilder:validation:Optional
	SourceArchiveBucketSelector *v1.Selector `json:"sourceArchiveBucketSelector,omitempty" tf:"-"`

	// The source archive object (file) in archive bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.BucketObject
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	SourceArchiveObject *string `json:"sourceArchiveObject,omitempty" tf:"source_archive_object,omitempty"`

	// Reference to a BucketObject in storage to populate sourceArchiveObject.
	// +kubebuilder:validation:Optional
	SourceArchiveObjectRef *v1.Reference `json:"sourceArchiveObjectRef,omitempty" tf:"-"`

	// Selector for a BucketObject in storage to populate sourceArchiveObject.
	// +kubebuilder:validation:Optional
	SourceArchiveObjectSelector *v1.Selector `json:"sourceArchiveObjectSelector,omitempty" tf:"-"`

	// Represents parameters related to source repository where a function is hosted.
	// Cannot be set alongside source_archive_bucket or source_archive_object. Structure is documented below. It must match the pattern projects/{project}/locations/{location}/repositories/{repository}.*
	SourceRepository []SourceRepositoryParameters `json:"sourceRepository,omitempty" tf:"source_repository,omitempty"`

	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as https_trigger_url. Cannot be used with event_trigger.
	TriggerHTTP *bool `json:"triggerHttp,omitempty" tf:"trigger_http,omitempty"`

	// The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is projects/*/locations/*/connectors/*.
	VPCConnector *string `json:"vpcConnector,omitempty" tf:"vpc_connector,omitempty"`

	// The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are ALL_TRAFFIC and PRIVATE_RANGES_ONLY. Defaults to PRIVATE_RANGES_ONLY. If unset, this field preserves the previously set value.
	VPCConnectorEgressSettings *string `json:"vpcConnectorEgressSettings,omitempty" tf:"vpc_connector_egress_settings,omitempty"`
}

type SecretEnvironmentVariablesInitParameters struct {

	// Name of the environment variable.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret. If not set, it will be populated with the function's project, assuming that the secret exists in the same project as of the function.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// ID of the secret in secret manager (not the full resource name).
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// Version of the secret (version number or the string "latest"). It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new clones start.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SecretEnvironmentVariablesObservation struct {

	// Name of the environment variable.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret. If not set, it will be populated with the function's project, assuming that the secret exists in the same project as of the function.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// ID of the secret in secret manager (not the full resource name).
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// Version of the secret (version number or the string "latest"). It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new clones start.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SecretEnvironmentVariablesParameters struct {

	// Name of the environment variable.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret. If not set, it will be populated with the function's project, assuming that the secret exists in the same project as of the function.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// ID of the secret in secret manager (not the full resource name).
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// Version of the secret (version number or the string "latest"). It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new clones start.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SecretVolumesInitParameters struct {

	// The path within the container to mount the secret volume. For example, setting the mount_path as "/etc/secrets" would mount the secret value files under the "/etc/secrets" directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount paths: "/etc/secrets" Restricted mount paths: "/cloudsql", "/dev/log", "/pod", "/proc", "/var/log".
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`

	// Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret. If not set, it will be populated with the function's project, assuming that the secret exists in the same project as of the function.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// ID of the secret in secret manager (not the full resource name).
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// List of secret versions to mount for this secret. If empty, the "latest" version of the secret will be made available in a file named after the secret under the mount point. Structure is documented below.
	Versions []VersionsInitParameters `json:"versions,omitempty" tf:"versions,omitempty"`
}

type SecretVolumesObservation struct {

	// The path within the container to mount the secret volume. For example, setting the mount_path as "/etc/secrets" would mount the secret value files under the "/etc/secrets" directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount paths: "/etc/secrets" Restricted mount paths: "/cloudsql", "/dev/log", "/pod", "/proc", "/var/log".
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`

	// Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret. If not set, it will be populated with the function's project, assuming that the secret exists in the same project as of the function.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// ID of the secret in secret manager (not the full resource name).
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// List of secret versions to mount for this secret. If empty, the "latest" version of the secret will be made available in a file named after the secret under the mount point. Structure is documented below.
	Versions []VersionsObservation `json:"versions,omitempty" tf:"versions,omitempty"`
}

type SecretVolumesParameters struct {

	// The path within the container to mount the secret volume. For example, setting the mount_path as "/etc/secrets" would mount the secret value files under the "/etc/secrets" directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount paths: "/etc/secrets" Restricted mount paths: "/cloudsql", "/dev/log", "/pod", "/proc", "/var/log".
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`

	// Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret. If not set, it will be populated with the function's project, assuming that the secret exists in the same project as of the function.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// ID of the secret in secret manager (not the full resource name).
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// List of secret versions to mount for this secret. If empty, the "latest" version of the secret will be made available in a file named after the secret under the mount point. Structure is documented below.
	Versions []VersionsParameters `json:"versions,omitempty" tf:"versions,omitempty"`
}

type SourceRepositoryInitParameters struct {

	// The URL pointing to the hosted repository where the function is defined. There are supported Cloud Source Repository URLs in the following formats:
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SourceRepositoryObservation struct {

	// The URL pointing to the hosted repository where the function was defined at the time of deployment.
	DeployedURL *string `json:"deployedUrl,omitempty" tf:"deployed_url,omitempty"`

	// The URL pointing to the hosted repository where the function is defined. There are supported Cloud Source Repository URLs in the following formats:
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SourceRepositoryParameters struct {

	// The URL pointing to the hosted repository where the function is defined. There are supported Cloud Source Repository URLs in the following formats:
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type VersionsInitParameters struct {

	// Relative path of the file under the mount path where the secret value for this version will be fetched and made available. For example, setting the mount_path as "/etc/secrets" and path as "/secret_foo" would mount the secret value file at "/etc/secrets/secret_foo".
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Version of the secret (version number or the string "latest"). It is preferable to use "latest" version with secret volumes as secret value changes are reflected immediately.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type VersionsObservation struct {

	// Relative path of the file under the mount path where the secret value for this version will be fetched and made available. For example, setting the mount_path as "/etc/secrets" and path as "/secret_foo" would mount the secret value file at "/etc/secrets/secret_foo".
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Version of the secret (version number or the string "latest"). It is preferable to use "latest" version with secret volumes as secret value changes are reflected immediately.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type VersionsParameters struct {

	// Relative path of the file under the mount path where the secret value for this version will be fetched and made available. For example, setting the mount_path as "/etc/secrets" and path as "/secret_foo" would mount the secret value file at "/etc/secrets/secret_foo".
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Version of the secret (version number or the string "latest"). It is preferable to use "latest" version with secret volumes as secret value changes are reflected immediately.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider FunctionInitParameters `json:"initProvider,omitempty"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Function is the Schema for the Functions API. Creates a new Cloud Function.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.runtime) || has(self.initProvider.runtime)",message="runtime is a required parameter"
	Spec   FunctionSpec   `json:"spec"`
	Status FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	Function_Kind             = "Function"
	Function_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Function_Kind}.String()
	Function_KindAPIVersion   = Function_Kind + "." + CRDGroupVersion.String()
	Function_GroupVersionKind = CRDGroupVersion.WithKind(Function_Kind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
