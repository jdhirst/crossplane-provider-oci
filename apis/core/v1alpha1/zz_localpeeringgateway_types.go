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

type LocalPeeringGatewayInitParameters struct {

	// (Updatable) The OCID of the compartment containing the local peering gateway (LPG).
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

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// should only be specified in one of the LPGs.
	PeerID *string `json:"peerId,omitempty" tf:"peer_id,omitempty"`

	// (Updatable) The OCID of the route table the LPG will use.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The OCID of the VCN the LPG belongs to.
	// +crossplane:generate:reference:type=Vcn
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`

	// Reference to a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDRef *v1.Reference `json:"vcnIdRef,omitempty" tf:"-"`

	// Selector for a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDSelector *v1.Selector `json:"vcnIdSelector,omitempty" tf:"-"`
}

type LocalPeeringGatewayObservation struct {

	// (Updatable) The OCID of the compartment containing the local peering gateway (LPG).
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The LPG's Oracle ID (OCID).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the VCN at the other end of the peering is in a different tenancy.  Example: false
	IsCrossTenancyPeering *bool `json:"isCrossTenancyPeering,omitempty" tf:"is_cross_tenancy_peering,omitempty"`

	// The smallest aggregate CIDR that contains all the CIDR routes advertised by the VCN at the other end of the peering from this LPG. See peerAdvertisedCidrDetails for the individual CIDRs. The value is null if the LPG is not peered.  Example: 192.168.0.0/16, or if aggregated with 172.16.0.0/24 then 128.0.0.0/1
	PeerAdvertisedCidr *string `json:"peerAdvertisedCidr,omitempty" tf:"peer_advertised_cidr,omitempty"`

	// The specific ranges of IP addresses available on or via the VCN at the other end of the peering from this LPG. The value is null if the LPG is not peered. You can use these as destination CIDRs for route rules to route a subnet's traffic to this LPG.  Example: [192.168.0.0/16, 172.16.0.0/24]
	PeerAdvertisedCidrDetails []*string `json:"peerAdvertisedCidrDetails,omitempty" tf:"peer_advertised_cidr_details,omitempty"`

	// should only be specified in one of the LPGs.
	PeerID *string `json:"peerId,omitempty" tf:"peer_id,omitempty"`

	// Whether the LPG is peered with another LPG. NEW means the LPG has not yet been peered. PENDING means the peering is being established. REVOKED means the LPG at the other end of the peering has been deleted.
	PeeringStatus *string `json:"peeringStatus,omitempty" tf:"peering_status,omitempty"`

	// Additional information regarding the peering status, if applicable.
	PeeringStatusDetails *string `json:"peeringStatusDetails,omitempty" tf:"peering_status_details,omitempty"`

	// (Updatable) The OCID of the route table the LPG will use.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The LPG's current lifecycle state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the LPG was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The OCID of the VCN the LPG belongs to.
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`
}

type LocalPeeringGatewayParameters struct {

	// (Updatable) The OCID of the compartment containing the local peering gateway (LPG).
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

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// should only be specified in one of the LPGs.
	// +kubebuilder:validation:Optional
	PeerID *string `json:"peerId,omitempty" tf:"peer_id,omitempty"`

	// (Updatable) The OCID of the route table the LPG will use.
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The OCID of the VCN the LPG belongs to.
	// +crossplane:generate:reference:type=Vcn
	// +kubebuilder:validation:Optional
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`

	// Reference to a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDRef *v1.Reference `json:"vcnIdRef,omitempty" tf:"-"`

	// Selector for a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDSelector *v1.Selector `json:"vcnIdSelector,omitempty" tf:"-"`
}

// LocalPeeringGatewaySpec defines the desired state of LocalPeeringGateway
type LocalPeeringGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocalPeeringGatewayParameters `json:"forProvider"`
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
	InitProvider LocalPeeringGatewayInitParameters `json:"initProvider,omitempty"`
}

// LocalPeeringGatewayStatus defines the observed state of LocalPeeringGateway.
type LocalPeeringGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocalPeeringGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LocalPeeringGateway is the Schema for the LocalPeeringGateways API. Provides the Local Peering Gateway resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type LocalPeeringGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocalPeeringGatewaySpec   `json:"spec"`
	Status            LocalPeeringGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocalPeeringGatewayList contains a list of LocalPeeringGateways
type LocalPeeringGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocalPeeringGateway `json:"items"`
}

// Repository type metadata.
var (
	LocalPeeringGateway_Kind             = "LocalPeeringGateway"
	LocalPeeringGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LocalPeeringGateway_Kind}.String()
	LocalPeeringGateway_KindAPIVersion   = LocalPeeringGateway_Kind + "." + CRDGroupVersion.String()
	LocalPeeringGateway_GroupVersionKind = CRDGroupVersion.WithKind(LocalPeeringGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&LocalPeeringGateway{}, &LocalPeeringGatewayList{})
}
