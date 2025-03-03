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

type LogSavedSearchInitParameters struct {

	// (Updatable) The OCID of the compartment that the resource belongs to.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Description for this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) The search query that is saved.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type LogSavedSearchObservation struct {

	// (Updatable) The OCID of the compartment that the resource belongs to.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Description for this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The OCID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Updatable) The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) The search query that is saved.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// The state of the LogSavedSearch
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Time the resource was created.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// Time the resource was last modified.
	TimeLastModified *string `json:"timeLastModified,omitempty" tf:"time_last_modified,omitempty"`
}

type LogSavedSearchParameters struct {

	// (Updatable) The OCID of the compartment that the resource belongs to.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Description for this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) The search query that is saved.
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

// LogSavedSearchSpec defines the desired state of LogSavedSearch
type LogSavedSearchSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogSavedSearchParameters `json:"forProvider"`
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
	InitProvider LogSavedSearchInitParameters `json:"initProvider,omitempty"`
}

// LogSavedSearchStatus defines the observed state of LogSavedSearch.
type LogSavedSearchStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogSavedSearchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LogSavedSearch is the Schema for the LogSavedSearchs API. Provides the Log Saved Search resource in Oracle Cloud Infrastructure Logging service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type LogSavedSearch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.query) || (has(self.initProvider) && has(self.initProvider.query))",message="spec.forProvider.query is a required parameter"
	Spec   LogSavedSearchSpec   `json:"spec"`
	Status LogSavedSearchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogSavedSearchList contains a list of LogSavedSearchs
type LogSavedSearchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogSavedSearch `json:"items"`
}

// Repository type metadata.
var (
	LogSavedSearch_Kind             = "LogSavedSearch"
	LogSavedSearch_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogSavedSearch_Kind}.String()
	LogSavedSearch_KindAPIVersion   = LogSavedSearch_Kind + "." + CRDGroupVersion.String()
	LogSavedSearch_GroupVersionKind = CRDGroupVersion.WithKind(LogSavedSearch_Kind)
)

func init() {
	SchemeBuilder.Register(&LogSavedSearch{}, &LogSavedSearchList{})
}
