/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/oracle/provider-oci/apis/identity/v1alpha1"
	v1alpha11 "github.com/oracle/provider-oci/apis/kms/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Secret.
func (mg *Secret) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KeyIDRef,
		Selector:     mg.Spec.ForProvider.KeyIDSelector,
		To: reference.To{
			List:    &v1alpha11.KeyList{},
			Managed: &v1alpha11.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyID")
	}
	mg.Spec.ForProvider.KeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VaultIDRef,
		Selector:     mg.Spec.ForProvider.VaultIDSelector,
		To: reference.To{
			List:    &v1alpha11.VaultList{},
			Managed: &v1alpha11.Vault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultID")
	}
	mg.Spec.ForProvider.VaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.CompartmentIDRef,
		Selector:     mg.Spec.InitProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CompartmentID")
	}
	mg.Spec.InitProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.KeyIDRef,
		Selector:     mg.Spec.InitProvider.KeyIDSelector,
		To: reference.To{
			List:    &v1alpha11.KeyList{},
			Managed: &v1alpha11.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyID")
	}
	mg.Spec.InitProvider.KeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VaultID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VaultIDRef,
		Selector:     mg.Spec.InitProvider.VaultIDSelector,
		To: reference.To{
			List:    &v1alpha11.VaultList{},
			Managed: &v1alpha11.Vault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VaultID")
	}
	mg.Spec.InitProvider.VaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VaultIDRef = rsp.ResolvedReference

	return nil
}
