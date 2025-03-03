/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessLoggingInitParameters struct {

	// (Updatable) Determines if the logging configuration is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

type AccessLoggingObservation struct {

	// (Updatable) Determines if the logging configuration is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

type AccessLoggingParameters struct {

	// (Updatable) Determines if the logging configuration is enabled.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

type ClientValidationInitParameters struct {

	// (Updatable) A list of alternate names to verify the subject identity in the certificate presented by the client.
	SubjectAlternateNames []*string `json:"subjectAlternateNames,omitempty" tf:"subject_alternate_names,omitempty"`

	// (Updatable) Resource representing the CA bundle.
	TrustedCABundle []TrustedCABundleInitParameters `json:"trustedCaBundle,omitempty" tf:"trusted_ca_bundle,omitempty"`
}

type ClientValidationObservation struct {

	// (Updatable) A list of alternate names to verify the subject identity in the certificate presented by the client.
	SubjectAlternateNames []*string `json:"subjectAlternateNames,omitempty" tf:"subject_alternate_names,omitempty"`

	// (Updatable) Resource representing the CA bundle.
	TrustedCABundle []TrustedCABundleObservation `json:"trustedCaBundle,omitempty" tf:"trusted_ca_bundle,omitempty"`
}

type ClientValidationParameters struct {

	// (Updatable) A list of alternate names to verify the subject identity in the certificate presented by the client.
	// +kubebuilder:validation:Optional
	SubjectAlternateNames []*string `json:"subjectAlternateNames,omitempty" tf:"subject_alternate_names,omitempty"`

	// (Updatable) Resource representing the CA bundle.
	// +kubebuilder:validation:Optional
	TrustedCABundle []TrustedCABundleParameters `json:"trustedCaBundle,omitempty" tf:"trusted_ca_bundle,omitempty"`
}

type HostsInitParameters struct {

	// (Updatable) Hostnames of the host. Applicable only for HTTP and TLS_PASSTHROUGH listeners. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", ".example.com", ".com".
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`

	// (Updatable) The listeners for the ingress gateway.
	Listeners []ListenersInitParameters `json:"listeners,omitempty" tf:"listeners,omitempty"`

	// (Updatable) A user-friendly name for the host. The name must be unique within the same ingress gateway. This name can be used in the ingress gateway route table resource to attach a route to this host.  Example: MyExampleHost
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type HostsObservation struct {

	// (Updatable) Hostnames of the host. Applicable only for HTTP and TLS_PASSTHROUGH listeners. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", ".example.com", ".com".
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`

	// (Updatable) The listeners for the ingress gateway.
	Listeners []ListenersObservation `json:"listeners,omitempty" tf:"listeners,omitempty"`

	// (Updatable) A user-friendly name for the host. The name must be unique within the same ingress gateway. This name can be used in the ingress gateway route table resource to attach a route to this host.  Example: MyExampleHost
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type HostsParameters struct {

	// (Updatable) Hostnames of the host. Applicable only for HTTP and TLS_PASSTHROUGH listeners. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", ".example.com", ".com".
	// +kubebuilder:validation:Optional
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`

	// (Updatable) The listeners for the ingress gateway.
	// +kubebuilder:validation:Optional
	Listeners []ListenersParameters `json:"listeners" tf:"listeners,omitempty"`

	// (Updatable) A user-friendly name for the host. The name must be unique within the same ingress gateway. This name can be used in the ingress gateway route table resource to attach a route to this host.  Example: MyExampleHost
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type IngressGatewayInitParameters struct {

	// (Updatable) This configuration determines if logging is enabled and where the logs will be output.
	AccessLogging []AccessLoggingInitParameters `json:"accessLogging,omitempty" tf:"access_logging,omitempty"`

	// (Updatable) The OCID of the compartment.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: {"foo-namespace.bar-key": "value"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: This is my new resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: {"bar-key": "value"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) An array of hostnames and their listener configuration that this gateway will bind to.
	Hosts []HostsInitParameters `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// The OCID of the service mesh in which this ingress gateway is created.
	// +crossplane:generate:reference:type=Mesh
	MeshID *string `json:"meshId,omitempty" tf:"mesh_id,omitempty"`

	// Reference to a Mesh to populate meshId.
	// +kubebuilder:validation:Optional
	MeshIDRef *v1.Reference `json:"meshIdRef,omitempty" tf:"-"`

	// Selector for a Mesh to populate meshId.
	// +kubebuilder:validation:Optional
	MeshIDSelector *v1.Selector `json:"meshIdSelector,omitempty" tf:"-"`

	// (Updatable) Mutual TLS settings used when sending requests to virtual services within the mesh.
	Mtls []MtlsInitParameters `json:"mtls,omitempty" tf:"mtls,omitempty"`

	// (Updatable) A user-friendly name for the host. The name must be unique within the same ingress gateway. This name can be used in the ingress gateway route table resource to attach a route to this host.  Example: MyExampleHost
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IngressGatewayObservation struct {

	// (Updatable) This configuration determines if logging is enabled and where the logs will be output.
	AccessLogging []AccessLoggingObservation `json:"accessLogging,omitempty" tf:"access_logging,omitempty"`

	// (Updatable) The OCID of the compartment.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: {"foo-namespace.bar-key": "value"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: This is my new resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: {"bar-key": "value"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) An array of hostnames and their listener configuration that this gateway will bind to.
	Hosts []HostsObservation `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// Unique identifier that is immutable on creation.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details,omitempty"`

	// The OCID of the service mesh in which this ingress gateway is created.
	MeshID *string `json:"meshId,omitempty" tf:"mesh_id,omitempty"`

	// (Updatable) Mutual TLS settings used when sending requests to virtual services within the mesh.
	Mtls []MtlsObservation `json:"mtls,omitempty" tf:"mtls,omitempty"`

	// (Updatable) A user-friendly name for the host. The name must be unique within the same ingress gateway. This name can be used in the ingress gateway route table resource to attach a route to this host.  Example: MyExampleHost
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The current state of the Resource.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. Example: {"orcl-cloud.free-tier-retained": "true"}
	// +mapType=granular
	SystemTags map[string]*string `json:"systemTags,omitempty" tf:"system_tags,omitempty"`

	// The time when this resource was created in an RFC3339 formatted datetime string.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The time when this resource was updated in an RFC3339 formatted datetime string.
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`
}

type IngressGatewayParameters struct {

	// (Updatable) This configuration determines if logging is enabled and where the logs will be output.
	// +kubebuilder:validation:Optional
	AccessLogging []AccessLoggingParameters `json:"accessLogging,omitempty" tf:"access_logging,omitempty"`

	// (Updatable) The OCID of the compartment.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: {"foo-namespace.bar-key": "value"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: This is my new resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: {"bar-key": "value"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) An array of hostnames and their listener configuration that this gateway will bind to.
	// +kubebuilder:validation:Optional
	Hosts []HostsParameters `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// The OCID of the service mesh in which this ingress gateway is created.
	// +crossplane:generate:reference:type=Mesh
	// +kubebuilder:validation:Optional
	MeshID *string `json:"meshId,omitempty" tf:"mesh_id,omitempty"`

	// Reference to a Mesh to populate meshId.
	// +kubebuilder:validation:Optional
	MeshIDRef *v1.Reference `json:"meshIdRef,omitempty" tf:"-"`

	// Selector for a Mesh to populate meshId.
	// +kubebuilder:validation:Optional
	MeshIDSelector *v1.Selector `json:"meshIdSelector,omitempty" tf:"-"`

	// (Updatable) Mutual TLS settings used when sending requests to virtual services within the mesh.
	// +kubebuilder:validation:Optional
	Mtls []MtlsParameters `json:"mtls,omitempty" tf:"mtls,omitempty"`

	// (Updatable) A user-friendly name for the host. The name must be unique within the same ingress gateway. This name can be used in the ingress gateway route table resource to attach a route to this host.  Example: MyExampleHost
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ListenersInitParameters struct {

	// (Updatable) Port on which ingress gateway is listening.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (Updatable) Type of protocol used.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Updatable) TLS enforcement config for the ingress listener.
	TLS []TLSInitParameters `json:"tls,omitempty" tf:"tls,omitempty"`
}

type ListenersObservation struct {

	// (Updatable) Port on which ingress gateway is listening.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (Updatable) Type of protocol used.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Updatable) TLS enforcement config for the ingress listener.
	TLS []TLSObservation `json:"tls,omitempty" tf:"tls,omitempty"`
}

type ListenersParameters struct {

	// (Updatable) Port on which ingress gateway is listening.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// (Updatable) Type of protocol used.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// (Updatable) TLS enforcement config for the ingress listener.
	// +kubebuilder:validation:Optional
	TLS []TLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`
}

type MtlsInitParameters struct {

	// (Updatable) The number of days the mTLS certificate is valid.  This value should be less than the Maximum Validity Duration  for Certificates (Days) setting on the Certificate Authority associated with this Mesh.  The certificate will be automatically renewed after 2/3 of the validity period, so a certificate with a maximum validity of 45 days will be renewed every 30 days.
	MaximumValidity *float64 `json:"maximumValidity,omitempty" tf:"maximum_validity,omitempty"`
}

type MtlsObservation struct {

	// (Applicable when type=OCI_CERTIFICATES) (Updatable) The OCID of the leaf certificate resource.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// (Updatable) The number of days the mTLS certificate is valid.  This value should be less than the Maximum Validity Duration  for Certificates (Days) setting on the Certificate Authority associated with this Mesh.  The certificate will be automatically renewed after 2/3 of the validity period, so a certificate with a maximum validity of 45 days will be renewed every 30 days.
	MaximumValidity *float64 `json:"maximumValidity,omitempty" tf:"maximum_validity,omitempty"`
}

type MtlsParameters struct {

	// (Updatable) The number of days the mTLS certificate is valid.  This value should be less than the Maximum Validity Duration  for Certificates (Days) setting on the Certificate Authority associated with this Mesh.  The certificate will be automatically renewed after 2/3 of the validity period, so a certificate with a maximum validity of 45 days will be renewed every 30 days.
	// +kubebuilder:validation:Optional
	MaximumValidity *float64 `json:"maximumValidity,omitempty" tf:"maximum_validity,omitempty"`
}

type ServerCertificateInitParameters struct {

	// (Applicable when type=OCI_CERTIFICATES) (Updatable) The OCID of the leaf certificate resource.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// (Applicable when type=LOCAL_FILE) (Updatable) Name of the secret. For Kubernetes this will be the name of an opaque Kubernetes secret with key ca.crt. For other platforms the secret must be mounted at: /etc/oci/secrets/${secretName}/ca.crt
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`

	// (Updatable) Type of certificate.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServerCertificateObservation struct {

	// (Applicable when type=OCI_CERTIFICATES) (Updatable) The OCID of the leaf certificate resource.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// (Applicable when type=LOCAL_FILE) (Updatable) Name of the secret. For Kubernetes this will be the name of an opaque Kubernetes secret with key ca.crt. For other platforms the secret must be mounted at: /etc/oci/secrets/${secretName}/ca.crt
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`

	// (Updatable) Type of certificate.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServerCertificateParameters struct {

	// (Applicable when type=OCI_CERTIFICATES) (Updatable) The OCID of the leaf certificate resource.
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// (Applicable when type=LOCAL_FILE) (Updatable) Name of the secret. For Kubernetes this will be the name of an opaque Kubernetes secret with key ca.crt. For other platforms the secret must be mounted at: /etc/oci/secrets/${secretName}/ca.crt
	// +kubebuilder:validation:Optional
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`

	// (Updatable) Type of certificate.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type TLSInitParameters struct {

	// (Updatable) Resource representing the TLS configuration used for validating client certificates.
	ClientValidation []ClientValidationInitParameters `json:"clientValidation,omitempty" tf:"client_validation,omitempty"`

	// (Updatable) DISABLED: Connection can only be plaintext. PERMISSIVE: Connection can be either plaintext or TLS/mTLS. If the clientValidation.trustedCaBundle property is configured for the listener, mTLS is performed and the client's certificates are validated by the gateway. TLS: Connection can only be TLS.  MUTUAL_TLS: Connection can only be MTLS.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// (Updatable) Resource representing the location of the TLS certificate.
	ServerCertificate []ServerCertificateInitParameters `json:"serverCertificate,omitempty" tf:"server_certificate,omitempty"`
}

type TLSObservation struct {

	// (Updatable) Resource representing the TLS configuration used for validating client certificates.
	ClientValidation []ClientValidationObservation `json:"clientValidation,omitempty" tf:"client_validation,omitempty"`

	// (Updatable) DISABLED: Connection can only be plaintext. PERMISSIVE: Connection can be either plaintext or TLS/mTLS. If the clientValidation.trustedCaBundle property is configured for the listener, mTLS is performed and the client's certificates are validated by the gateway. TLS: Connection can only be TLS.  MUTUAL_TLS: Connection can only be MTLS.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// (Updatable) Resource representing the location of the TLS certificate.
	ServerCertificate []ServerCertificateObservation `json:"serverCertificate,omitempty" tf:"server_certificate,omitempty"`
}

type TLSParameters struct {

	// (Updatable) Resource representing the TLS configuration used for validating client certificates.
	// +kubebuilder:validation:Optional
	ClientValidation []ClientValidationParameters `json:"clientValidation,omitempty" tf:"client_validation,omitempty"`

	// (Updatable) DISABLED: Connection can only be plaintext. PERMISSIVE: Connection can be either plaintext or TLS/mTLS. If the clientValidation.trustedCaBundle property is configured for the listener, mTLS is performed and the client's certificates are validated by the gateway. TLS: Connection can only be TLS.  MUTUAL_TLS: Connection can only be MTLS.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// (Updatable) Resource representing the location of the TLS certificate.
	// +kubebuilder:validation:Optional
	ServerCertificate []ServerCertificateParameters `json:"serverCertificate,omitempty" tf:"server_certificate,omitempty"`
}

type TrustedCABundleInitParameters struct {

	// (Applicable when type=OCI_CERTIFICATES) (Updatable) The OCID of the CA Bundle resource.
	CABundleID *string `json:"caBundleId,omitempty" tf:"ca_bundle_id,omitempty"`

	// (Applicable when type=LOCAL_FILE) (Updatable) Name of the secret. For Kubernetes this will be the name of an opaque Kubernetes secret with key ca.crt. For other platforms the secret must be mounted at: /etc/oci/secrets/${secretName}/ca.crt
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`

	// (Updatable) Type of certificate.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TrustedCABundleObservation struct {

	// (Applicable when type=OCI_CERTIFICATES) (Updatable) The OCID of the CA Bundle resource.
	CABundleID *string `json:"caBundleId,omitempty" tf:"ca_bundle_id,omitempty"`

	// (Applicable when type=LOCAL_FILE) (Updatable) Name of the secret. For Kubernetes this will be the name of an opaque Kubernetes secret with key ca.crt. For other platforms the secret must be mounted at: /etc/oci/secrets/${secretName}/ca.crt
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`

	// (Updatable) Type of certificate.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TrustedCABundleParameters struct {

	// (Applicable when type=OCI_CERTIFICATES) (Updatable) The OCID of the CA Bundle resource.
	// +kubebuilder:validation:Optional
	CABundleID *string `json:"caBundleId,omitempty" tf:"ca_bundle_id,omitempty"`

	// (Applicable when type=LOCAL_FILE) (Updatable) Name of the secret. For Kubernetes this will be the name of an opaque Kubernetes secret with key ca.crt. For other platforms the secret must be mounted at: /etc/oci/secrets/${secretName}/ca.crt
	// +kubebuilder:validation:Optional
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`

	// (Updatable) Type of certificate.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// IngressGatewaySpec defines the desired state of IngressGateway
type IngressGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IngressGatewayParameters `json:"forProvider"`
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
	InitProvider IngressGatewayInitParameters `json:"initProvider,omitempty"`
}

// IngressGatewayStatus defines the observed state of IngressGateway.
type IngressGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IngressGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IngressGateway is the Schema for the IngressGateways API. Provides the Ingress Gateway resource in Oracle Cloud Infrastructure Service Mesh service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type IngressGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hosts) || (has(self.initProvider) && has(self.initProvider.hosts))",message="spec.forProvider.hosts is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   IngressGatewaySpec   `json:"spec"`
	Status IngressGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IngressGatewayList contains a list of IngressGateways
type IngressGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IngressGateway `json:"items"`
}

// Repository type metadata.
var (
	IngressGateway_Kind             = "IngressGateway"
	IngressGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IngressGateway_Kind}.String()
	IngressGateway_KindAPIVersion   = IngressGateway_Kind + "." + CRDGroupVersion.String()
	IngressGateway_GroupVersionKind = CRDGroupVersion.WithKind(IngressGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&IngressGateway{}, &IngressGatewayList{})
}
