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

type VnicAttachmentCreateVnicDetailsObservation struct {
}

type VnicAttachmentCreateVnicDetailsParameters struct {

	// Whether the VNIC should be assigned a DNS record. If set to false, no DNS record registion for the VNIC; if set to true, DNS record will be registered. Example: true
	// +kubebuilder:validation:Optional
	AssignPrivateDNSRecord *bool `json:"assignPrivateDnsRecord,omitempty" tf:"assign_private_dns_record,omitempty"`

	// Whether the VNIC should be assigned a public IP address. Defaults to whether the subnet is public or private. If not set and the VNIC is being created in a private subnet (that is, where prohibitPublicIpOnVnic = true in the Subnet), then no public IP address is assigned. If not set and the subnet is public (prohibitPublicIpOnVnic = false), then a public IP address is assigned. If set to true and prohibitPublicIpOnVnic = true, an error is returned.
	// +kubebuilder:validation:Optional
	AssignPublicIP *string `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, bminstance1 in FQDN bminstance1.subnet123.vcn1.oraclevcn.com). Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123. The value appears in the Vnic object and also the PrivateIp object returned by ListPrivateIps and GetPrivateIp.
	// +kubebuilder:validation:Optional
	HostnameLabel *string `json:"hostnameLabel,omitempty" tf:"hostname_label,omitempty"`

	// (Updatable) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see NetworkSecurityGroup.
	// +kubebuilder:validation:Optional
	NsgIds []*string `json:"nsgIds,omitempty" tf:"nsg_ids,omitempty"`

	// A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This is the VNIC's primary private IP address. The value appears in the Vnic object and also the PrivateIp object returned by ListPrivateIps and GetPrivateIp.
	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// (Updatable) Whether the source/destination check is disabled on the VNIC. Defaults to false, which means the check is performed. For information about why you would skip the source/destination check, see Using a Private IP as a Route Target.
	// +kubebuilder:validation:Optional
	SkipSourceDestCheck *bool `json:"skipSourceDestCheck,omitempty" tf:"skip_source_dest_check,omitempty"`

	// The OCID of the subnet to create the VNIC in. When launching an instance, use this subnetId instead of the deprecated subnetId in LaunchInstanceDetails. At least one of them is required; if you provide both, the values must match.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Provide this attribute only if you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN. The value is the OCID of the VLAN. See Vlan.
	// +kubebuilder:validation:Optional
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type VnicAttachmentObservation struct {

	// The availability domain of the instance.  Example: Uocm:PHX-AD-1
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain,omitempty"`

	// The OCID of the compartment the VNIC attachment is in, which is the same compartment the instance is in.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// The OCID of the VNIC attachment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The current state of the VNIC attachment.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The OCID of the subnet to create the VNIC in. When launching an instance, use this subnetId instead of the deprecated subnetId in LaunchInstanceDetails. At least one of them is required; if you provide both, the values must match.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The date and time the VNIC attachment was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// Provide this attribute only if you are an Oracle Cloud VMware Solution customer and creating a secondary VNIC in a VLAN. The value is the OCID of the VLAN. See Vlan.
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// The Oracle-assigned VLAN tag of the attached VNIC. Available after the attachment process is complete.
	VlanTag *float64 `json:"vlanTag,omitempty" tf:"vlan_tag,omitempty"`

	// The OCID of the VNIC. Available after the attachment process is complete.
	VnicID *string `json:"vnicId,omitempty" tf:"vnic_id,omitempty"`
}

type VnicAttachmentParameters struct {

	// (Updatable) Contains properties for a VNIC. You use this object when creating the primary VNIC during instance launch or when creating a secondary VNIC. For more information about VNICs, see Virtual Network Interface Cards (VNICs).
	// +kubebuilder:validation:Required
	CreateVnicDetails []VnicAttachmentCreateVnicDetailsParameters `json:"createVnicDetails" tf:"create_vnic_details,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The OCID of the instance.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Which physical network interface card (NIC) the VNIC will use. Defaults to 0. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see Virtual Network Interface Cards (VNICs).
	// +kubebuilder:validation:Optional
	NicIndex *float64 `json:"nicIndex,omitempty" tf:"nic_index,omitempty"`
}

// VnicAttachmentSpec defines the desired state of VnicAttachment
type VnicAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VnicAttachmentParameters `json:"forProvider"`
}

// VnicAttachmentStatus defines the observed state of VnicAttachment.
type VnicAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VnicAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VnicAttachment is the Schema for the VnicAttachments API. Provides the Vnic Attachment resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type VnicAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VnicAttachmentSpec   `json:"spec"`
	Status            VnicAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VnicAttachmentList contains a list of VnicAttachments
type VnicAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VnicAttachment `json:"items"`
}

// Repository type metadata.
var (
	VnicAttachment_Kind             = "VnicAttachment"
	VnicAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VnicAttachment_Kind}.String()
	VnicAttachment_KindAPIVersion   = VnicAttachment_Kind + "." + CRDGroupVersion.String()
	VnicAttachment_GroupVersionKind = CRDGroupVersion.WithKind(VnicAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&VnicAttachment{}, &VnicAttachmentList{})
}
