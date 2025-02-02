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

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this InstanceIAMMember.
func (mg *InstanceIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceNameRef,
		Selector:     mg.Spec.ForProvider.InstanceNameSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.InstanceNameRef,
		Selector:     mg.Spec.InitProvider.InstanceNameSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceName")
	}
	mg.Spec.InitProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RuntimeIAMMember.
func (mg *RuntimeIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RuntimeName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RuntimeNameRef,
		Selector:     mg.Spec.ForProvider.RuntimeNameSelector,
		To: reference.To{
			List:    &RuntimeList{},
			Managed: &Runtime{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RuntimeName")
	}
	mg.Spec.ForProvider.RuntimeName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RuntimeNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RuntimeName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RuntimeNameRef,
		Selector:     mg.Spec.InitProvider.RuntimeNameSelector,
		To: reference.To{
			List:    &RuntimeList{},
			Managed: &Runtime{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RuntimeName")
	}
	mg.Spec.InitProvider.RuntimeName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RuntimeNameRef = rsp.ResolvedReference

	return nil
}
