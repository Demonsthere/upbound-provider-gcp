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
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Function.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Function) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.BuildConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("artifact.gcp.upbound.io", "v1beta1", "RegistryRepository", "RegistryRepositoryList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BuildConfig[i3].DockerRepository),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.BuildConfig[i3].DockerRepositoryRef,
				Selector:     mg.Spec.ForProvider.BuildConfig[i3].DockerRepositorySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.BuildConfig[i3].DockerRepository")
		}
		mg.Spec.ForProvider.BuildConfig[i3].DockerRepository = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.BuildConfig[i3].DockerRepositoryRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.BuildConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.BuildConfig[i3].Source); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta1", "Bucket", "BucketList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Bucket),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].BucketRef,
						Selector:     mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].BucketSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Bucket")
				}
				mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].BucketRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.BuildConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.BuildConfig[i3].Source); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta1", "BucketObject", "BucketObjectList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Object),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].ObjectRef,
						Selector:     mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].ObjectSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Object")
				}
				mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Object = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.BuildConfig[i3].Source[i4].StorageSource[i5].ObjectRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.BuildConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudbuild.gcp.upbound.io", "v1beta1", "WorkerPool", "WorkerPoolList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BuildConfig[i3].WorkerPool),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.BuildConfig[i3].WorkerPoolRef,
				Selector:     mg.Spec.ForProvider.BuildConfig[i3].WorkerPoolSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.BuildConfig[i3].WorkerPool")
		}
		mg.Spec.ForProvider.BuildConfig[i3].WorkerPool = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.BuildConfig[i3].WorkerPoolRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.EventTrigger); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.EventTrigger[i3].EventFilters); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta1", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventTrigger[i3].EventFilters[i4].Value),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.EventTrigger[i3].EventFilters[i4].ValueRef,
					Selector:     mg.Spec.ForProvider.EventTrigger[i3].EventFilters[i4].ValueSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EventTrigger[i3].EventFilters[i4].Value")
			}
			mg.Spec.ForProvider.EventTrigger[i3].EventFilters[i4].Value = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EventTrigger[i3].EventFilters[i4].ValueRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.EventTrigger); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta1", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventTrigger[i3].PubsubTopic),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.EventTrigger[i3].PubsubTopicRef,
				Selector:     mg.Spec.ForProvider.EventTrigger[i3].PubsubTopicSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EventTrigger[i3].PubsubTopic")
		}
		mg.Spec.ForProvider.EventTrigger[i3].PubsubTopic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EventTrigger[i3].PubsubTopicRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.EventTrigger); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventTrigger[i3].ServiceAccountEmail),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.ForProvider.EventTrigger[i3].ServiceAccountEmailRef,
				Selector:     mg.Spec.ForProvider.EventTrigger[i3].ServiceAccountEmailSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EventTrigger[i3].ServiceAccountEmail")
		}
		mg.Spec.ForProvider.EventTrigger[i3].ServiceAccountEmail = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EventTrigger[i3].ServiceAccountEmailRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ServiceConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ServiceConfig[i3].SecretEnvironmentVariables); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta1", "Secret", "SecretList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].Secret),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].SecretRef,
					Selector:     mg.Spec.ForProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].SecretSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].Secret")
			}
			mg.Spec.ForProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].Secret = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].SecretRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ServiceConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ServiceConfig[i3].SecretVolumes); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta1", "Secret", "SecretList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceConfig[i3].SecretVolumes[i4].Secret),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.ServiceConfig[i3].SecretVolumes[i4].SecretRef,
					Selector:     mg.Spec.ForProvider.ServiceConfig[i3].SecretVolumes[i4].SecretSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ServiceConfig[i3].SecretVolumes[i4].Secret")
			}
			mg.Spec.ForProvider.ServiceConfig[i3].SecretVolumes[i4].Secret = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ServiceConfig[i3].SecretVolumes[i4].SecretRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ServiceConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceConfig[i3].ServiceAccountEmail),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.ForProvider.ServiceConfig[i3].ServiceAccountEmailRef,
				Selector:     mg.Spec.ForProvider.ServiceConfig[i3].ServiceAccountEmailSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ServiceConfig[i3].ServiceAccountEmail")
		}
		mg.Spec.ForProvider.ServiceConfig[i3].ServiceAccountEmail = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ServiceConfig[i3].ServiceAccountEmailRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.BuildConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("artifact.gcp.upbound.io", "v1beta1", "RegistryRepository", "RegistryRepositoryList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BuildConfig[i3].DockerRepository),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.BuildConfig[i3].DockerRepositoryRef,
				Selector:     mg.Spec.InitProvider.BuildConfig[i3].DockerRepositorySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.BuildConfig[i3].DockerRepository")
		}
		mg.Spec.InitProvider.BuildConfig[i3].DockerRepository = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.BuildConfig[i3].DockerRepositoryRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.BuildConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.BuildConfig[i3].Source); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta1", "Bucket", "BucketList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Bucket),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].BucketRef,
						Selector:     mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].BucketSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Bucket")
				}
				mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].BucketRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.BuildConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.BuildConfig[i3].Source); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta1", "BucketObject", "BucketObjectList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Object),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].ObjectRef,
						Selector:     mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].ObjectSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Object")
				}
				mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].Object = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.BuildConfig[i3].Source[i4].StorageSource[i5].ObjectRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.BuildConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudbuild.gcp.upbound.io", "v1beta1", "WorkerPool", "WorkerPoolList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BuildConfig[i3].WorkerPool),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.BuildConfig[i3].WorkerPoolRef,
				Selector:     mg.Spec.InitProvider.BuildConfig[i3].WorkerPoolSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.BuildConfig[i3].WorkerPool")
		}
		mg.Spec.InitProvider.BuildConfig[i3].WorkerPool = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.BuildConfig[i3].WorkerPoolRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.EventTrigger); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.EventTrigger[i3].EventFilters); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("storage.gcp.upbound.io", "v1beta1", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventTrigger[i3].EventFilters[i4].Value),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.EventTrigger[i3].EventFilters[i4].ValueRef,
					Selector:     mg.Spec.InitProvider.EventTrigger[i3].EventFilters[i4].ValueSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.EventTrigger[i3].EventFilters[i4].Value")
			}
			mg.Spec.InitProvider.EventTrigger[i3].EventFilters[i4].Value = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.EventTrigger[i3].EventFilters[i4].ValueRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.EventTrigger); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta1", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventTrigger[i3].PubsubTopic),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.EventTrigger[i3].PubsubTopicRef,
				Selector:     mg.Spec.InitProvider.EventTrigger[i3].PubsubTopicSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EventTrigger[i3].PubsubTopic")
		}
		mg.Spec.InitProvider.EventTrigger[i3].PubsubTopic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EventTrigger[i3].PubsubTopicRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.EventTrigger); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventTrigger[i3].ServiceAccountEmail),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.InitProvider.EventTrigger[i3].ServiceAccountEmailRef,
				Selector:     mg.Spec.InitProvider.EventTrigger[i3].ServiceAccountEmailSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EventTrigger[i3].ServiceAccountEmail")
		}
		mg.Spec.InitProvider.EventTrigger[i3].ServiceAccountEmail = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EventTrigger[i3].ServiceAccountEmailRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ServiceConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ServiceConfig[i3].SecretEnvironmentVariables); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta1", "Secret", "SecretList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].Secret),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].SecretRef,
					Selector:     mg.Spec.InitProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].SecretSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].Secret")
			}
			mg.Spec.InitProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].Secret = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.ServiceConfig[i3].SecretEnvironmentVariables[i4].SecretRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ServiceConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ServiceConfig[i3].SecretVolumes); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("secretmanager.gcp.upbound.io", "v1beta1", "Secret", "SecretList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceConfig[i3].SecretVolumes[i4].Secret),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.ServiceConfig[i3].SecretVolumes[i4].SecretRef,
					Selector:     mg.Spec.InitProvider.ServiceConfig[i3].SecretVolumes[i4].SecretSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.ServiceConfig[i3].SecretVolumes[i4].Secret")
			}
			mg.Spec.InitProvider.ServiceConfig[i3].SecretVolumes[i4].Secret = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.ServiceConfig[i3].SecretVolumes[i4].SecretRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ServiceConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceConfig[i3].ServiceAccountEmail),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.InitProvider.ServiceConfig[i3].ServiceAccountEmailRef,
				Selector:     mg.Spec.InitProvider.ServiceConfig[i3].ServiceAccountEmailSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ServiceConfig[i3].ServiceAccountEmail")
		}
		mg.Spec.InitProvider.ServiceConfig[i3].ServiceAccountEmail = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ServiceConfig[i3].ServiceAccountEmailRef = rsp.ResolvedReference

	}

	return nil
}
