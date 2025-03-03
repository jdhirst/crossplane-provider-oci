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

type ApplicationInitParameters struct {

	// (Updatable) The OCID of the compartment to create the application within.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Application configuration. These values are passed on to the function as environment variables, functions may override application configuration. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: {"MY_FUNCTION_CONFIG": "ConfVal"}
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// The display name of the application. The display name must be unique within the compartment containing the application. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) Define the image signature verification policy for an application.
	ImagePolicyConfig []ImagePolicyConfigInitParameters `json:"imagePolicyConfig,omitempty" tf:"image_policy_config,omitempty"`

	// (Updatable) The OCIDs of the Network Security Groups to add the application to.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.NetworkSecurityGroup
	// +listType=set
	NetworkSecurityGroupIds []*string `json:"networkSecurityGroupIds,omitempty" tf:"network_security_group_ids,omitempty"`

	// References to NetworkSecurityGroup in core to populate networkSecurityGroupIds.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIdsRefs []v1.Reference `json:"networkSecurityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of NetworkSecurityGroup in core to populate networkSecurityGroupIds.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIdsSelector *v1.Selector `json:"networkSecurityGroupIdsSelector,omitempty" tf:"-"`

	// Valid values are GENERIC_X86, GENERIC_ARM and GENERIC_X86_ARM. Default is GENERIC_X86. Setting this to GENERIC_X86, will run the functions in the application on X86 processor architecture. Setting this to GENERIC_ARM, will run the functions in the application on ARM processor architecture. When set to GENERIC_X86_ARM, functions in the application are run on either X86 or ARM processor architecture. Accepted values are: GENERIC_X86, GENERIC_ARM, GENERIC_X86_ARM
	Shape *string `json:"shape,omitempty" tf:"shape,omitempty"`

	// The OCIDs of the subnets in which to run functions in the application.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.Subnet
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// References to Subnet in core to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in core to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`

	// (Updatable) A syslog URL to which to send all function logs. Supports tcp, udp, and tcp+tls. The syslog URL must be reachable from all of the subnets configured for the application. Note: If you enable the Oracle Cloud Infrastructure Logging service for this application, the syslogUrl value is ignored. Function logs are sent to the Oracle Cloud Infrastructure Logging service, and not to the syslog URL.  Example: tcp://logserver.myserver:1234
	SyslogURL *string `json:"syslogUrl,omitempty" tf:"syslog_url,omitempty"`

	// (Updatable) Define the tracing configuration for an application.
	TraceConfig []TraceConfigInitParameters `json:"traceConfig,omitempty" tf:"trace_config,omitempty"`
}

type ApplicationObservation struct {

	// (Updatable) The OCID of the compartment to create the application within.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) Application configuration. These values are passed on to the function as environment variables, functions may override application configuration. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: {"MY_FUNCTION_CONFIG": "ConfVal"}
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// The display name of the application. The display name must be unique within the compartment containing the application. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The OCID of the application.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Updatable) Define the image signature verification policy for an application.
	ImagePolicyConfig []ImagePolicyConfigObservation `json:"imagePolicyConfig,omitempty" tf:"image_policy_config,omitempty"`

	// (Updatable) The OCIDs of the Network Security Groups to add the application to.
	// +listType=set
	NetworkSecurityGroupIds []*string `json:"networkSecurityGroupIds,omitempty" tf:"network_security_group_ids,omitempty"`

	// Valid values are GENERIC_X86, GENERIC_ARM and GENERIC_X86_ARM. Default is GENERIC_X86. Setting this to GENERIC_X86, will run the functions in the application on X86 processor architecture. Setting this to GENERIC_ARM, will run the functions in the application on ARM processor architecture. When set to GENERIC_X86_ARM, functions in the application are run on either X86 or ARM processor architecture. Accepted values are: GENERIC_X86, GENERIC_ARM, GENERIC_X86_ARM
	Shape *string `json:"shape,omitempty" tf:"shape,omitempty"`

	// The current state of the application.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The OCIDs of the subnets in which to run functions in the application.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// (Updatable) A syslog URL to which to send all function logs. Supports tcp, udp, and tcp+tls. The syslog URL must be reachable from all of the subnets configured for the application. Note: If you enable the Oracle Cloud Infrastructure Logging service for this application, the syslogUrl value is ignored. Function logs are sent to the Oracle Cloud Infrastructure Logging service, and not to the syslog URL.  Example: tcp://logserver.myserver:1234
	SyslogURL *string `json:"syslogUrl,omitempty" tf:"syslog_url,omitempty"`

	// The time the application was created, expressed in RFC 3339 timestamp format.  Example: 2018-09-12T22:47:12.613Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The time the application was updated, expressed in RFC 3339 timestamp format. Example: 2018-09-12T22:47:12.613Z
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`

	// (Updatable) Define the tracing configuration for an application.
	TraceConfig []TraceConfigObservation `json:"traceConfig,omitempty" tf:"trace_config,omitempty"`
}

type ApplicationParameters struct {

	// (Updatable) The OCID of the compartment to create the application within.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Application configuration. These values are passed on to the function as environment variables, functions may override application configuration. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: {"MY_FUNCTION_CONFIG": "ConfVal"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// The display name of the application. The display name must be unique within the compartment containing the application. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) Define the image signature verification policy for an application.
	// +kubebuilder:validation:Optional
	ImagePolicyConfig []ImagePolicyConfigParameters `json:"imagePolicyConfig,omitempty" tf:"image_policy_config,omitempty"`

	// (Updatable) The OCIDs of the Network Security Groups to add the application to.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.NetworkSecurityGroup
	// +kubebuilder:validation:Optional
	// +listType=set
	NetworkSecurityGroupIds []*string `json:"networkSecurityGroupIds,omitempty" tf:"network_security_group_ids,omitempty"`

	// References to NetworkSecurityGroup in core to populate networkSecurityGroupIds.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIdsRefs []v1.Reference `json:"networkSecurityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of NetworkSecurityGroup in core to populate networkSecurityGroupIds.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIdsSelector *v1.Selector `json:"networkSecurityGroupIdsSelector,omitempty" tf:"-"`

	// Valid values are GENERIC_X86, GENERIC_ARM and GENERIC_X86_ARM. Default is GENERIC_X86. Setting this to GENERIC_X86, will run the functions in the application on X86 processor architecture. Setting this to GENERIC_ARM, will run the functions in the application on ARM processor architecture. When set to GENERIC_X86_ARM, functions in the application are run on either X86 or ARM processor architecture. Accepted values are: GENERIC_X86, GENERIC_ARM, GENERIC_X86_ARM
	// +kubebuilder:validation:Optional
	Shape *string `json:"shape,omitempty" tf:"shape,omitempty"`

	// The OCIDs of the subnets in which to run functions in the application.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// References to Subnet in core to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in core to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`

	// (Updatable) A syslog URL to which to send all function logs. Supports tcp, udp, and tcp+tls. The syslog URL must be reachable from all of the subnets configured for the application. Note: If you enable the Oracle Cloud Infrastructure Logging service for this application, the syslogUrl value is ignored. Function logs are sent to the Oracle Cloud Infrastructure Logging service, and not to the syslog URL.  Example: tcp://logserver.myserver:1234
	// +kubebuilder:validation:Optional
	SyslogURL *string `json:"syslogUrl,omitempty" tf:"syslog_url,omitempty"`

	// (Updatable) Define the tracing configuration for an application.
	// +kubebuilder:validation:Optional
	TraceConfig []TraceConfigParameters `json:"traceConfig,omitempty" tf:"trace_config,omitempty"`
}

type ImagePolicyConfigInitParameters struct {

	// (Updatable) Define if image signature verification policy is enabled for the application.
	IsPolicyEnabled *bool `json:"isPolicyEnabled,omitempty" tf:"is_policy_enabled,omitempty"`

	// (Updatable) A list of KMS key details.
	KeyDetails []KeyDetailsInitParameters `json:"keyDetails,omitempty" tf:"key_details,omitempty"`
}

type ImagePolicyConfigObservation struct {

	// (Updatable) Define if image signature verification policy is enabled for the application.
	IsPolicyEnabled *bool `json:"isPolicyEnabled,omitempty" tf:"is_policy_enabled,omitempty"`

	// (Updatable) A list of KMS key details.
	KeyDetails []KeyDetailsObservation `json:"keyDetails,omitempty" tf:"key_details,omitempty"`
}

type ImagePolicyConfigParameters struct {

	// (Updatable) Define if image signature verification policy is enabled for the application.
	// +kubebuilder:validation:Optional
	IsPolicyEnabled *bool `json:"isPolicyEnabled" tf:"is_policy_enabled,omitempty"`

	// (Updatable) A list of KMS key details.
	// +kubebuilder:validation:Optional
	KeyDetails []KeyDetailsParameters `json:"keyDetails,omitempty" tf:"key_details,omitempty"`
}

type KeyDetailsInitParameters struct {

	// (Updatable) The OCIDs of the KMS key that will be used to verify the image signature.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type KeyDetailsObservation struct {

	// (Updatable) The OCIDs of the KMS key that will be used to verify the image signature.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`
}

type KeyDetailsParameters struct {

	// (Updatable) The OCIDs of the KMS key that will be used to verify the image signature.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId" tf:"kms_key_id,omitempty"`
}

type TraceConfigInitParameters struct {

	// (Updatable) The OCID of the collector (e.g. an APM Domain) trace events will be sent to.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// (Updatable) Define if tracing is enabled for the resource.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

type TraceConfigObservation struct {

	// (Updatable) The OCID of the collector (e.g. an APM Domain) trace events will be sent to.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// (Updatable) Define if tracing is enabled for the resource.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

type TraceConfigParameters struct {

	// (Updatable) The OCID of the collector (e.g. an APM Domain) trace events will be sent to.
	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// (Updatable) Define if tracing is enabled for the resource.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
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
	InitProvider ApplicationInitParameters `json:"initProvider,omitempty"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Application is the Schema for the Applications API. Provides the Application resource in Oracle Cloud Infrastructure Functions service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	Spec   ApplicationSpec   `json:"spec"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
