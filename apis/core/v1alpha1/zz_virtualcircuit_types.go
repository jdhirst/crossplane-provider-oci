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

type CrossConnectMappingsInitParameters struct {

	// (Updatable) The key for BGP MD5 authentication. Only applicable if your system requires MD5 authentication. If empty or not set (null), that means you don't use BGP MD5 authentication.
	BGPMd5AuthKey *string `json:"bgpMd5AuthKey,omitempty" tf:"bgp_md5auth_key,omitempty"`

	// (Updatable) The OCID of the cross-connect or cross-connect group for this mapping. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider).
	CrossConnectOrCrossConnectGroupID *string `json:"crossConnectOrCrossConnectGroupId,omitempty" tf:"cross_connect_or_cross_connect_group_id,omitempty"`

	// (Updatable) The BGP IPv4 address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IPv4 address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IPv4 address of the provider's edge router. Must use a subnet mask from /28 to /31.
	CustomerBGPPeeringIP *string `json:"customerBgpPeeringIp,omitempty" tf:"customer_bgp_peering_ip,omitempty"`

	// (Updatable) IPv6 is currently supported only in the Government Cloud. The BGP IPv6 address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IPv6 address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IPv6 address of the provider's edge router. Only subnet masks from /64 up to /127 are allowed.
	CustomerBGPPeeringIPv6 *string `json:"customerBgpPeeringIpv6,omitempty" tf:"customer_bgp_peering_ipv6,omitempty"`

	// (Updatable) The IPv4 address for Oracle's end of the BGP session. Must use a subnet mask from /28 to /31. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.
	OracleBGPPeeringIP *string `json:"oracleBgpPeeringIp,omitempty" tf:"oracle_bgp_peering_ip,omitempty"`

	// (Updatable) IPv6 is currently supported only in the Government Cloud. The IPv6 address for Oracle's end of the BGP session.  Only subnet masks from /64 up to /127 are allowed. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.
	OracleBGPPeeringIPv6 *string `json:"oracleBgpPeeringIpv6,omitempty" tf:"oracle_bgp_peering_ipv6,omitempty"`

	// (Updatable) The number of the specific VLAN (on the cross-connect or cross-connect group) that is assigned to this virtual circuit. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider).  Example: 200
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type CrossConnectMappingsObservation struct {

	// (Updatable) The key for BGP MD5 authentication. Only applicable if your system requires MD5 authentication. If empty or not set (null), that means you don't use BGP MD5 authentication.
	BGPMd5AuthKey *string `json:"bgpMd5AuthKey,omitempty" tf:"bgp_md5auth_key,omitempty"`

	// (Updatable) The OCID of the cross-connect or cross-connect group for this mapping. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider).
	CrossConnectOrCrossConnectGroupID *string `json:"crossConnectOrCrossConnectGroupId,omitempty" tf:"cross_connect_or_cross_connect_group_id,omitempty"`

	// (Updatable) The BGP IPv4 address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IPv4 address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IPv4 address of the provider's edge router. Must use a subnet mask from /28 to /31.
	CustomerBGPPeeringIP *string `json:"customerBgpPeeringIp,omitempty" tf:"customer_bgp_peering_ip,omitempty"`

	// (Updatable) IPv6 is currently supported only in the Government Cloud. The BGP IPv6 address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IPv6 address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IPv6 address of the provider's edge router. Only subnet masks from /64 up to /127 are allowed.
	CustomerBGPPeeringIPv6 *string `json:"customerBgpPeeringIpv6,omitempty" tf:"customer_bgp_peering_ipv6,omitempty"`

	// (Updatable) The IPv4 address for Oracle's end of the BGP session. Must use a subnet mask from /28 to /31. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.
	OracleBGPPeeringIP *string `json:"oracleBgpPeeringIp,omitempty" tf:"oracle_bgp_peering_ip,omitempty"`

	// (Updatable) IPv6 is currently supported only in the Government Cloud. The IPv6 address for Oracle's end of the BGP session.  Only subnet masks from /64 up to /127 are allowed. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.
	OracleBGPPeeringIPv6 *string `json:"oracleBgpPeeringIpv6,omitempty" tf:"oracle_bgp_peering_ipv6,omitempty"`

	// (Updatable) The number of the specific VLAN (on the cross-connect or cross-connect group) that is assigned to this virtual circuit. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider).  Example: 200
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type CrossConnectMappingsParameters struct {

	// (Updatable) The key for BGP MD5 authentication. Only applicable if your system requires MD5 authentication. If empty or not set (null), that means you don't use BGP MD5 authentication.
	// +kubebuilder:validation:Optional
	BGPMd5AuthKey *string `json:"bgpMd5AuthKey,omitempty" tf:"bgp_md5auth_key,omitempty"`

	// (Updatable) The OCID of the cross-connect or cross-connect group for this mapping. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider).
	// +kubebuilder:validation:Optional
	CrossConnectOrCrossConnectGroupID *string `json:"crossConnectOrCrossConnectGroupId,omitempty" tf:"cross_connect_or_cross_connect_group_id,omitempty"`

	// (Updatable) The BGP IPv4 address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IPv4 address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IPv4 address of the provider's edge router. Must use a subnet mask from /28 to /31.
	// +kubebuilder:validation:Optional
	CustomerBGPPeeringIP *string `json:"customerBgpPeeringIp,omitempty" tf:"customer_bgp_peering_ip,omitempty"`

	// (Updatable) IPv6 is currently supported only in the Government Cloud. The BGP IPv6 address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IPv6 address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IPv6 address of the provider's edge router. Only subnet masks from /64 up to /127 are allowed.
	// +kubebuilder:validation:Optional
	CustomerBGPPeeringIPv6 *string `json:"customerBgpPeeringIpv6,omitempty" tf:"customer_bgp_peering_ipv6,omitempty"`

	// (Updatable) The IPv4 address for Oracle's end of the BGP session. Must use a subnet mask from /28 to /31. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.
	// +kubebuilder:validation:Optional
	OracleBGPPeeringIP *string `json:"oracleBgpPeeringIp,omitempty" tf:"oracle_bgp_peering_ip,omitempty"`

	// (Updatable) IPv6 is currently supported only in the Government Cloud. The IPv6 address for Oracle's end of the BGP session.  Only subnet masks from /64 up to /127 are allowed. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.
	// +kubebuilder:validation:Optional
	OracleBGPPeeringIPv6 *string `json:"oracleBgpPeeringIpv6,omitempty" tf:"oracle_bgp_peering_ipv6,omitempty"`

	// (Updatable) The number of the specific VLAN (on the cross-connect or cross-connect group) that is assigned to this virtual circuit. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider).  Example: 200
	// +kubebuilder:validation:Optional
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type PublicPrefixesInitParameters struct {

	// (Updatable) An individual public IP prefix (CIDR) to add to the public virtual circuit. All prefix sizes are allowed.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`
}

type PublicPrefixesObservation struct {

	// (Updatable) An individual public IP prefix (CIDR) to add to the public virtual circuit. All prefix sizes are allowed.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`
}

type PublicPrefixesParameters struct {

	// (Updatable) An individual public IP prefix (CIDR) to add to the public virtual circuit. All prefix sizes are allowed.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock" tf:"cidr_block,omitempty"`
}

type VirtualCircuitInitParameters struct {

	// (Updatable) Set to ENABLED (the default) to activate the BGP session of the virtual circuit, set to DISABLED to deactivate the virtual circuit.
	BGPAdminState *string `json:"bgpAdminState,omitempty" tf:"bgp_admin_state,omitempty"`

	// (Updatable) The provisioned data rate of the connection. To get a list of the available bandwidth levels (that is, shapes), see ListFastConnectProviderServiceVirtualCircuitBandwidthShapes.  Example: 10 Gbps
	BandwidthShapeName *string `json:"bandwidthShapeName,omitempty" tf:"bandwidth_shape_name,omitempty"`

	// (Updatable) The OCID of the compartment to contain the virtual circuit.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) Create a CrossConnectMapping for each cross-connect or cross-connect group this virtual circuit will run on.
	CrossConnectMappings []CrossConnectMappingsInitParameters `json:"crossConnectMappings,omitempty" tf:"cross_connect_mappings,omitempty"`

	// (Updatable) Your BGP ASN (either public or private). Provide this value only if there's a BGP session that goes from your edge router to Oracle. Otherwise, leave this empty or null. Can be a 2-byte or 4-byte ASN. Uses "asplain" format.  Example: 12345 (2-byte) or 1587232876 (4-byte)
	CustomerAsn *string `json:"customerAsn,omitempty" tf:"customer_asn,omitempty"`

	// (Updatable) Deprecated. Instead use customerAsn. If you specify values for both, the request will be rejected.
	CustomerBGPAsn *float64 `json:"customerBgpAsn,omitempty" tf:"customer_bgp_asn,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) For private virtual circuits only. The OCID of the dynamic routing gateway (DRG) that this virtual circuit uses.
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// (Updatable) The layer 3 IP MTU to use with this virtual circuit.
	IPMtu *string `json:"ipMtu,omitempty" tf:"ip_mtu,omitempty"`

	// (Updatable) Set to true to enable BFD for IPv4 BGP peering, or set to false to disable BFD. If this is not set, the default is false.
	IsBfdEnabled *bool `json:"isBfdEnabled,omitempty" tf:"is_bfd_enabled,omitempty"`

	// (Updatable) Set to true for the virtual circuit to carry only encrypted traffic, or set to false for the virtual circuit to carry unencrypted traffic. If this is not set, the default is false.
	IsTransportMode *bool `json:"isTransportMode,omitempty" tf:"is_transport_mode,omitempty"`

	// The OCID of the service offered by the provider (if you're connecting via a provider). To get a list of the available service offerings, see ListFastConnectProviderServices.
	ProviderServiceID *string `json:"providerServiceId,omitempty" tf:"provider_service_id,omitempty"`

	// (Updatable) The service key name offered by the provider (if the customer is connecting via a provider).
	ProviderServiceKeyName *string `json:"providerServiceKeyName,omitempty" tf:"provider_service_key_name,omitempty"`

	// (Updatable) For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to advertise across the connection.
	PublicPrefixes []PublicPrefixesInitParameters `json:"publicPrefixes,omitempty" tf:"public_prefixes,omitempty"`

	// The Oracle Cloud Infrastructure region where this virtual circuit is located. Example: phx
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// (Updatable) The routing policy sets how routing information about the Oracle cloud is shared over a public virtual circuit. Policies available are: ORACLE_SERVICE_NETWORK, REGIONAL, MARKET_LEVEL, and GLOBAL. See Route Filtering for details. By default, routing information is shared for all routes in the same market.
	RoutingPolicy []*string `json:"routingPolicy,omitempty" tf:"routing_policy,omitempty"`

	// The type of IP addresses used in this virtual circuit. PRIVATE means RFC 1918 addresses (10.0.0.0/8, 172.16/12, and 192.168/16).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VirtualCircuitObservation struct {

	// (Updatable) Set to ENABLED (the default) to activate the BGP session of the virtual circuit, set to DISABLED to deactivate the virtual circuit.
	BGPAdminState *string `json:"bgpAdminState,omitempty" tf:"bgp_admin_state,omitempty"`

	// The state of the Ipv6 BGP session associated with the virtual circuit.
	BGPIpv6SessionState *string `json:"bgpIpv6SessionState,omitempty" tf:"bgp_ipv6session_state,omitempty"`

	// Deprecated. Instead use the information in FastConnectProviderService.
	BGPManagement *string `json:"bgpManagement,omitempty" tf:"bgp_management,omitempty"`

	// The state of the Ipv4 BGP session associated with the virtual circuit.
	BGPSessionState *string `json:"bgpSessionState,omitempty" tf:"bgp_session_state,omitempty"`

	// (Updatable) The provisioned data rate of the connection. To get a list of the available bandwidth levels (that is, shapes), see ListFastConnectProviderServiceVirtualCircuitBandwidthShapes.  Example: 10 Gbps
	BandwidthShapeName *string `json:"bandwidthShapeName,omitempty" tf:"bandwidth_shape_name,omitempty"`

	// (Updatable) The OCID of the compartment to contain the virtual circuit.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) Create a CrossConnectMapping for each cross-connect or cross-connect group this virtual circuit will run on.
	CrossConnectMappings []CrossConnectMappingsObservation `json:"crossConnectMappings,omitempty" tf:"cross_connect_mappings,omitempty"`

	// (Updatable) Your BGP ASN (either public or private). Provide this value only if there's a BGP session that goes from your edge router to Oracle. Otherwise, leave this empty or null. Can be a 2-byte or 4-byte ASN. Uses "asplain" format.  Example: 12345 (2-byte) or 1587232876 (4-byte)
	CustomerAsn *string `json:"customerAsn,omitempty" tf:"customer_asn,omitempty"`

	// (Updatable) Deprecated. Instead use customerAsn. If you specify values for both, the request will be rejected.
	CustomerBGPAsn *float64 `json:"customerBgpAsn,omitempty" tf:"customer_bgp_asn,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) For private virtual circuits only. The OCID of the dynamic routing gateway (DRG) that this virtual circuit uses.
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// The virtual circuit's Oracle ID (OCID).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Updatable) The layer 3 IP MTU to use with this virtual circuit.
	IPMtu *string `json:"ipMtu,omitempty" tf:"ip_mtu,omitempty"`

	// (Updatable) Set to true to enable BFD for IPv4 BGP peering, or set to false to disable BFD. If this is not set, the default is false.
	IsBfdEnabled *bool `json:"isBfdEnabled,omitempty" tf:"is_bfd_enabled,omitempty"`

	// (Updatable) Set to true for the virtual circuit to carry only encrypted traffic, or set to false for the virtual circuit to carry unencrypted traffic. If this is not set, the default is false.
	IsTransportMode *bool `json:"isTransportMode,omitempty" tf:"is_transport_mode,omitempty"`

	// The Oracle BGP ASN.
	OracleBGPAsn *float64 `json:"oracleBgpAsn,omitempty" tf:"oracle_bgp_asn,omitempty"`

	// The OCID of the service offered by the provider (if you're connecting via a provider). To get a list of the available service offerings, see ListFastConnectProviderServices.
	ProviderServiceID *string `json:"providerServiceId,omitempty" tf:"provider_service_id,omitempty"`

	// (Updatable) The service key name offered by the provider (if the customer is connecting via a provider).
	ProviderServiceKeyName *string `json:"providerServiceKeyName,omitempty" tf:"provider_service_key_name,omitempty"`

	// The provider's state in relation to this virtual circuit (if the customer is connecting via a provider). ACTIVE means the provider has provisioned the virtual circuit from their end. INACTIVE means the provider has not yet provisioned the virtual circuit, or has de-provisioned it.
	ProviderState *string `json:"providerState,omitempty" tf:"provider_state,omitempty"`

	// (Updatable) For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to advertise across the connection.
	PublicPrefixes []PublicPrefixesObservation `json:"publicPrefixes,omitempty" tf:"public_prefixes,omitempty"`

	// Provider-supplied reference information about this virtual circuit (if the customer is connecting via a provider).
	ReferenceComment *string `json:"referenceComment,omitempty" tf:"reference_comment,omitempty"`

	// The Oracle Cloud Infrastructure region where this virtual circuit is located. Example: phx
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// (Updatable) The routing policy sets how routing information about the Oracle cloud is shared over a public virtual circuit. Policies available are: ORACLE_SERVICE_NETWORK, REGIONAL, MARKET_LEVEL, and GLOBAL. See Route Filtering for details. By default, routing information is shared for all routes in the same market.
	RoutingPolicy []*string `json:"routingPolicy,omitempty" tf:"routing_policy,omitempty"`

	// Provider service type.
	ServiceType *string `json:"serviceType,omitempty" tf:"service_type,omitempty"`

	// The virtual circuit's current state. For information about the different states, see FastConnect Overview.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the virtual circuit was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The type of IP addresses used in this virtual circuit. PRIVATE means RFC 1918 addresses (10.0.0.0/8, 172.16/12, and 192.168/16).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// This resource provides redundancy level details for the virtual circuit. For more about redundancy, see FastConnect Redundancy Best Practices.
	VirtualCircuitRedundancyMetadata []VirtualCircuitRedundancyMetadataObservation `json:"virtualCircuitRedundancyMetadata,omitempty" tf:"virtual_circuit_redundancy_metadata,omitempty"`
}

type VirtualCircuitParameters struct {

	// (Updatable) Set to ENABLED (the default) to activate the BGP session of the virtual circuit, set to DISABLED to deactivate the virtual circuit.
	// +kubebuilder:validation:Optional
	BGPAdminState *string `json:"bgpAdminState,omitempty" tf:"bgp_admin_state,omitempty"`

	// (Updatable) The provisioned data rate of the connection. To get a list of the available bandwidth levels (that is, shapes), see ListFastConnectProviderServiceVirtualCircuitBandwidthShapes.  Example: 10 Gbps
	// +kubebuilder:validation:Optional
	BandwidthShapeName *string `json:"bandwidthShapeName,omitempty" tf:"bandwidth_shape_name,omitempty"`

	// (Updatable) The OCID of the compartment to contain the virtual circuit.
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) Create a CrossConnectMapping for each cross-connect or cross-connect group this virtual circuit will run on.
	// +kubebuilder:validation:Optional
	CrossConnectMappings []CrossConnectMappingsParameters `json:"crossConnectMappings,omitempty" tf:"cross_connect_mappings,omitempty"`

	// (Updatable) Your BGP ASN (either public or private). Provide this value only if there's a BGP session that goes from your edge router to Oracle. Otherwise, leave this empty or null. Can be a 2-byte or 4-byte ASN. Uses "asplain" format.  Example: 12345 (2-byte) or 1587232876 (4-byte)
	// +kubebuilder:validation:Optional
	CustomerAsn *string `json:"customerAsn,omitempty" tf:"customer_asn,omitempty"`

	// (Updatable) Deprecated. Instead use customerAsn. If you specify values for both, the request will be rejected.
	// +kubebuilder:validation:Optional
	CustomerBGPAsn *float64 `json:"customerBgpAsn,omitempty" tf:"customer_bgp_asn,omitempty"`

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

	// (Updatable) For private virtual circuits only. The OCID of the dynamic routing gateway (DRG) that this virtual circuit uses.
	// +kubebuilder:validation:Optional
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// (Updatable) The layer 3 IP MTU to use with this virtual circuit.
	// +kubebuilder:validation:Optional
	IPMtu *string `json:"ipMtu,omitempty" tf:"ip_mtu,omitempty"`

	// (Updatable) Set to true to enable BFD for IPv4 BGP peering, or set to false to disable BFD. If this is not set, the default is false.
	// +kubebuilder:validation:Optional
	IsBfdEnabled *bool `json:"isBfdEnabled,omitempty" tf:"is_bfd_enabled,omitempty"`

	// (Updatable) Set to true for the virtual circuit to carry only encrypted traffic, or set to false for the virtual circuit to carry unencrypted traffic. If this is not set, the default is false.
	// +kubebuilder:validation:Optional
	IsTransportMode *bool `json:"isTransportMode,omitempty" tf:"is_transport_mode,omitempty"`

	// The OCID of the service offered by the provider (if you're connecting via a provider). To get a list of the available service offerings, see ListFastConnectProviderServices.
	// +kubebuilder:validation:Optional
	ProviderServiceID *string `json:"providerServiceId,omitempty" tf:"provider_service_id,omitempty"`

	// (Updatable) The service key name offered by the provider (if the customer is connecting via a provider).
	// +kubebuilder:validation:Optional
	ProviderServiceKeyName *string `json:"providerServiceKeyName,omitempty" tf:"provider_service_key_name,omitempty"`

	// (Updatable) For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to advertise across the connection.
	// +kubebuilder:validation:Optional
	PublicPrefixes []PublicPrefixesParameters `json:"publicPrefixes,omitempty" tf:"public_prefixes,omitempty"`

	// The Oracle Cloud Infrastructure region where this virtual circuit is located. Example: phx
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// (Updatable) The routing policy sets how routing information about the Oracle cloud is shared over a public virtual circuit. Policies available are: ORACLE_SERVICE_NETWORK, REGIONAL, MARKET_LEVEL, and GLOBAL. See Route Filtering for details. By default, routing information is shared for all routes in the same market.
	// +kubebuilder:validation:Optional
	RoutingPolicy []*string `json:"routingPolicy,omitempty" tf:"routing_policy,omitempty"`

	// The type of IP addresses used in this virtual circuit. PRIVATE means RFC 1918 addresses (10.0.0.0/8, 172.16/12, and 192.168/16).
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VirtualCircuitRedundancyMetadataInitParameters struct {
}

type VirtualCircuitRedundancyMetadataObservation struct {

	// The configured redundancy level of the virtual circuit.
	ConfiguredRedundancyLevel *string `json:"configuredRedundancyLevel,omitempty" tf:"configured_redundancy_level,omitempty"`

	// Indicates if the configured level is met for IPv4 BGP redundancy.
	Ipv4BgpSessionRedundancyStatus *string `json:"ipv4bgpSessionRedundancyStatus,omitempty" tf:"ipv4bgp_session_redundancy_status,omitempty"`

	// Indicates if the configured level is met for IPv6 BGP redundancy.
	Ipv6BgpSessionRedundancyStatus *string `json:"ipv6bgpSessionRedundancyStatus,omitempty" tf:"ipv6bgp_session_redundancy_status,omitempty"`
}

type VirtualCircuitRedundancyMetadataParameters struct {
}

// VirtualCircuitSpec defines the desired state of VirtualCircuit
type VirtualCircuitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualCircuitParameters `json:"forProvider"`
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
	InitProvider VirtualCircuitInitParameters `json:"initProvider,omitempty"`
}

// VirtualCircuitStatus defines the observed state of VirtualCircuit.
type VirtualCircuitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualCircuitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VirtualCircuit is the Schema for the VirtualCircuits API. Provides the Virtual Circuit resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type VirtualCircuit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.compartmentId) || (has(self.initProvider) && has(self.initProvider.compartmentId))",message="spec.forProvider.compartmentId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   VirtualCircuitSpec   `json:"spec"`
	Status VirtualCircuitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualCircuitList contains a list of VirtualCircuits
type VirtualCircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualCircuit `json:"items"`
}

// Repository type metadata.
var (
	VirtualCircuit_Kind             = "VirtualCircuit"
	VirtualCircuit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualCircuit_Kind}.String()
	VirtualCircuit_KindAPIVersion   = VirtualCircuit_Kind + "." + CRDGroupVersion.String()
	VirtualCircuit_GroupVersionKind = CRDGroupVersion.WithKind(VirtualCircuit_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualCircuit{}, &VirtualCircuitList{})
}
