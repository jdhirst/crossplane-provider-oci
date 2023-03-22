/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this BalancerBackend
func (mg *BalancerBackend) GetTerraformResourceType() string {
	return "oci_load_balancer_backend"
}

// GetConnectionDetailsMapping for this BalancerBackend
func (tr *BalancerBackend) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BalancerBackend
func (tr *BalancerBackend) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BalancerBackend
func (tr *BalancerBackend) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BalancerBackend
func (tr *BalancerBackend) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BalancerBackend
func (tr *BalancerBackend) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BalancerBackend
func (tr *BalancerBackend) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BalancerBackend using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BalancerBackend) LateInitialize(attrs []byte) (bool, error) {
	params := &BalancerBackendParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BalancerBackend) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BalancerBackendSet
func (mg *BalancerBackendSet) GetTerraformResourceType() string {
	return "oci_load_balancer_backend_set"
}

// GetConnectionDetailsMapping for this BalancerBackendSet
func (tr *BalancerBackendSet) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BalancerBackendSet
func (tr *BalancerBackendSet) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BalancerBackendSet
func (tr *BalancerBackendSet) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BalancerBackendSet
func (tr *BalancerBackendSet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BalancerBackendSet
func (tr *BalancerBackendSet) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BalancerBackendSet
func (tr *BalancerBackendSet) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BalancerBackendSet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BalancerBackendSet) LateInitialize(attrs []byte) (bool, error) {
	params := &BalancerBackendSetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BalancerBackendSet) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BalancerCertificate
func (mg *BalancerCertificate) GetTerraformResourceType() string {
	return "oci_load_balancer_certificate"
}

// GetConnectionDetailsMapping for this BalancerCertificate
func (tr *BalancerCertificate) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"passphrase": "spec.forProvider.passphraseSecretRef", "private_key": "spec.forProvider.privateKeySecretRef"}
}

// GetObservation of this BalancerCertificate
func (tr *BalancerCertificate) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BalancerCertificate
func (tr *BalancerCertificate) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BalancerCertificate
func (tr *BalancerCertificate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BalancerCertificate
func (tr *BalancerCertificate) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BalancerCertificate
func (tr *BalancerCertificate) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BalancerCertificate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BalancerCertificate) LateInitialize(attrs []byte) (bool, error) {
	params := &BalancerCertificateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BalancerCertificate) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BalancerHostname
func (mg *BalancerHostname) GetTerraformResourceType() string {
	return "oci_load_balancer_hostname"
}

// GetConnectionDetailsMapping for this BalancerHostname
func (tr *BalancerHostname) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BalancerHostname
func (tr *BalancerHostname) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BalancerHostname
func (tr *BalancerHostname) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BalancerHostname
func (tr *BalancerHostname) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BalancerHostname
func (tr *BalancerHostname) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BalancerHostname
func (tr *BalancerHostname) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BalancerHostname using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BalancerHostname) LateInitialize(attrs []byte) (bool, error) {
	params := &BalancerHostnameParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BalancerHostname) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BalancerListener
func (mg *BalancerListener) GetTerraformResourceType() string {
	return "oci_load_balancer_listener"
}

// GetConnectionDetailsMapping for this BalancerListener
func (tr *BalancerListener) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BalancerListener
func (tr *BalancerListener) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BalancerListener
func (tr *BalancerListener) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BalancerListener
func (tr *BalancerListener) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BalancerListener
func (tr *BalancerListener) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BalancerListener
func (tr *BalancerListener) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BalancerListener using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BalancerListener) LateInitialize(attrs []byte) (bool, error) {
	params := &BalancerListenerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BalancerListener) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BalancerLoadBalancer
func (mg *BalancerLoadBalancer) GetTerraformResourceType() string {
	return "oci_load_balancer_load_balancer"
}

// GetConnectionDetailsMapping for this BalancerLoadBalancer
func (tr *BalancerLoadBalancer) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BalancerLoadBalancer
func (tr *BalancerLoadBalancer) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BalancerLoadBalancer
func (tr *BalancerLoadBalancer) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BalancerLoadBalancer
func (tr *BalancerLoadBalancer) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BalancerLoadBalancer
func (tr *BalancerLoadBalancer) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BalancerLoadBalancer
func (tr *BalancerLoadBalancer) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BalancerLoadBalancer using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BalancerLoadBalancer) LateInitialize(attrs []byte) (bool, error) {
	params := &BalancerLoadBalancerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BalancerLoadBalancer) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BalancerLoadBalancerRoutingPolicy
func (mg *BalancerLoadBalancerRoutingPolicy) GetTerraformResourceType() string {
	return "oci_load_balancer_load_balancer_routing_policy"
}

// GetConnectionDetailsMapping for this BalancerLoadBalancerRoutingPolicy
func (tr *BalancerLoadBalancerRoutingPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BalancerLoadBalancerRoutingPolicy
func (tr *BalancerLoadBalancerRoutingPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BalancerLoadBalancerRoutingPolicy
func (tr *BalancerLoadBalancerRoutingPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BalancerLoadBalancerRoutingPolicy
func (tr *BalancerLoadBalancerRoutingPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BalancerLoadBalancerRoutingPolicy
func (tr *BalancerLoadBalancerRoutingPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BalancerLoadBalancerRoutingPolicy
func (tr *BalancerLoadBalancerRoutingPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BalancerLoadBalancerRoutingPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BalancerLoadBalancerRoutingPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &BalancerLoadBalancerRoutingPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BalancerLoadBalancerRoutingPolicy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BalancerPathRouteSet
func (mg *BalancerPathRouteSet) GetTerraformResourceType() string {
	return "oci_load_balancer_path_route_set"
}

// GetConnectionDetailsMapping for this BalancerPathRouteSet
func (tr *BalancerPathRouteSet) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BalancerPathRouteSet
func (tr *BalancerPathRouteSet) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BalancerPathRouteSet
func (tr *BalancerPathRouteSet) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BalancerPathRouteSet
func (tr *BalancerPathRouteSet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BalancerPathRouteSet
func (tr *BalancerPathRouteSet) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BalancerPathRouteSet
func (tr *BalancerPathRouteSet) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BalancerPathRouteSet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BalancerPathRouteSet) LateInitialize(attrs []byte) (bool, error) {
	params := &BalancerPathRouteSetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BalancerPathRouteSet) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BalancerRuleSet
func (mg *BalancerRuleSet) GetTerraformResourceType() string {
	return "oci_load_balancer_rule_set"
}

// GetConnectionDetailsMapping for this BalancerRuleSet
func (tr *BalancerRuleSet) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BalancerRuleSet
func (tr *BalancerRuleSet) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BalancerRuleSet
func (tr *BalancerRuleSet) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BalancerRuleSet
func (tr *BalancerRuleSet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BalancerRuleSet
func (tr *BalancerRuleSet) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BalancerRuleSet
func (tr *BalancerRuleSet) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BalancerRuleSet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BalancerRuleSet) LateInitialize(attrs []byte) (bool, error) {
	params := &BalancerRuleSetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BalancerRuleSet) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BalancerSSLCipherSuite
func (mg *BalancerSSLCipherSuite) GetTerraformResourceType() string {
	return "oci_load_balancer_ssl_cipher_suite"
}

// GetConnectionDetailsMapping for this BalancerSSLCipherSuite
func (tr *BalancerSSLCipherSuite) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BalancerSSLCipherSuite
func (tr *BalancerSSLCipherSuite) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BalancerSSLCipherSuite
func (tr *BalancerSSLCipherSuite) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BalancerSSLCipherSuite
func (tr *BalancerSSLCipherSuite) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BalancerSSLCipherSuite
func (tr *BalancerSSLCipherSuite) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BalancerSSLCipherSuite
func (tr *BalancerSSLCipherSuite) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BalancerSSLCipherSuite using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BalancerSSLCipherSuite) LateInitialize(attrs []byte) (bool, error) {
	params := &BalancerSSLCipherSuiteParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BalancerSSLCipherSuite) GetTerraformSchemaVersion() int {
	return 0
}
