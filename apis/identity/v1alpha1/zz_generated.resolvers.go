/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AuthenticationPolicy.
func (mg *AuthenticationPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &CompartmentList{},
			Managed: &Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.CompartmentIDRef,
		Selector:     mg.Spec.InitProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &CompartmentList{},
			Managed: &Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CompartmentID")
	}
	mg.Spec.InitProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CompartmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this IdentityProvider.
func (mg *IdentityProvider) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &CompartmentList{},
			Managed: &Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.CompartmentIDRef,
		Selector:     mg.Spec.InitProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &CompartmentList{},
			Managed: &Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CompartmentID")
	}
	mg.Spec.InitProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CompartmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &CompartmentList{},
			Managed: &Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.CompartmentIDRef,
		Selector:     mg.Spec.InitProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &CompartmentList{},
			Managed: &Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CompartmentID")
	}
	mg.Spec.InitProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CompartmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Tag.
func (mg *Tag) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TagNamespaceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TagNamespaceIDRef,
		Selector:     mg.Spec.ForProvider.TagNamespaceIDSelector,
		To: reference.To{
			List:    &TagNamespaceList{},
			Managed: &TagNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TagNamespaceID")
	}
	mg.Spec.ForProvider.TagNamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TagNamespaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TagNamespaceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.TagNamespaceIDRef,
		Selector:     mg.Spec.InitProvider.TagNamespaceIDSelector,
		To: reference.To{
			List:    &TagNamespaceList{},
			Managed: &TagNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TagNamespaceID")
	}
	mg.Spec.InitProvider.TagNamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TagNamespaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TagNamespace.
func (mg *TagNamespace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &CompartmentList{},
			Managed: &Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.CompartmentIDRef,
		Selector:     mg.Spec.InitProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &CompartmentList{},
			Managed: &Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CompartmentID")
	}
	mg.Spec.InitProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CompartmentIDRef = rsp.ResolvedReference

	return nil
}
