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

type ExternalMastersObservation struct {
}

type ExternalMastersParameters struct {

	// (Updatable) The server's IP address (IPv4 or IPv6).
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// (Updatable) The server's port. Port value must be a value of 53, otherwise omit the port value.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (Updatable) The OCID of the TSIG key.
	// +kubebuilder:validation:Optional
	TsigKeyID *string `json:"tsigKeyId,omitempty" tf:"tsig_key_id,omitempty"`
}

type NameserversObservation struct {

	// The hostname of the nameserver.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`
}

type NameserversParameters struct {
}

type ZoneObservation struct {

	// The OCID of the zone.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed.
	IsProtected *bool `json:"isProtected,omitempty" tf:"is_protected,omitempty"`

	// The authoritative nameservers for the zone.
	Nameservers []NameserversObservation `json:"nameservers,omitempty" tf:"nameservers,omitempty"`

	// The canonical absolute URL of the resource.
	Self *string `json:"self,omitempty" tf:"self,omitempty"`

	// The current serial of the zone. As seen in the zone's SOA record.
	Serial *float64 `json:"serial,omitempty" tf:"serial,omitempty"`

	// The current state of the zone resource.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// Version is the never-repeating, totally-orderable, version of the zone, from which the serial field of the zone's SOA record is derived.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ZoneParameters struct {

	// (Updatable) The OCID of the compartment the resource belongs to.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) External master servers for the zone. externalMasters becomes a required parameter when the zoneType value is SECONDARY.
	// +kubebuilder:validation:Optional
	ExternalMasters []ExternalMastersParameters `json:"externalMasters,omitempty" tf:"external_masters,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The name of the zone.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies to operate only on resources that have a matching DNS scope.
	// This value will be null for zones in the global DNS and PRIVATE when creating a private zone.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The OCID of the private view containing the zone. This value will be null for zones in the global DNS, which are publicly resolvable and not part of a private view.
	// +crossplane:generate:reference:type=View
	// +kubebuilder:validation:Optional
	ViewID *string `json:"viewId,omitempty" tf:"view_id,omitempty"`

	// Reference to a View to populate viewId.
	// +kubebuilder:validation:Optional
	ViewIDRef *v1.Reference `json:"viewIdRef,omitempty" tf:"-"`

	// Selector for a View to populate viewId.
	// +kubebuilder:validation:Optional
	ViewIDSelector *v1.Selector `json:"viewIdSelector,omitempty" tf:"-"`

	// The type of the zone. Must be either PRIMARY or SECONDARY. SECONDARY is only supported for GLOBAL zones.
	// +kubebuilder:validation:Required
	ZoneType *string `json:"zoneType" tf:"zone_type,omitempty"`
}

// ZoneSpec defines the desired state of Zone
type ZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneParameters `json:"forProvider"`
}

// ZoneStatus defines the observed state of Zone.
type ZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Zone is the Schema for the Zones API. Provides the Zone resource in Oracle Cloud Infrastructure DNS service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ZoneSpec   `json:"spec"`
	Status            ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneList contains a list of Zones
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zone `json:"items"`
}

// Repository type metadata.
var (
	Zone_Kind             = "Zone"
	Zone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Zone_Kind}.String()
	Zone_KindAPIVersion   = Zone_Kind + "." + CRDGroupVersion.String()
	Zone_GroupVersionKind = CRDGroupVersion.WithKind(Zone_Kind)
)

func init() {
	SchemeBuilder.Register(&Zone{}, &ZoneList{})
}
