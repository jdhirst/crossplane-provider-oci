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

type VolumeBackupPolicyAssignmentInitParameters struct {

	// The OCID of the volume or volume group to assign the policy to.
	// +crossplane:generate:reference:type=Volume
	AssetID *string `json:"assetId,omitempty" tf:"asset_id,omitempty"`

	// Reference to a Volume to populate assetId.
	// +kubebuilder:validation:Optional
	AssetIDRef *v1.Reference `json:"assetIdRef,omitempty" tf:"-"`

	// Selector for a Volume to populate assetId.
	// +kubebuilder:validation:Optional
	AssetIDSelector *v1.Selector `json:"assetIdSelector,omitempty" tf:"-"`

	// The OCID of the volume backup policy to assign to the volume.
	// +crossplane:generate:reference:type=VolumeBackupPolicy
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a VolumeBackupPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a VolumeBackupPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`

	// The OCID of the Vault service key which is the master encryption key for the block / boot volume cross region backups, which will be used in the destination region to encrypt the backup's encryption keys. For more information about the Vault service and encryption keys, see Overview of Vault service and Using Keys.
	XrcKMSKeyID *string `json:"xrcKmsKeyId,omitempty" tf:"xrc_kms_key_id,omitempty"`
}

type VolumeBackupPolicyAssignmentObservation struct {

	// The OCID of the volume or volume group to assign the policy to.
	AssetID *string `json:"assetId,omitempty" tf:"asset_id,omitempty"`

	// The OCID of the volume backup policy assignment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The OCID of the volume backup policy to assign to the volume.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// The date and time the volume backup policy was assigned to the volume. The format is defined by RFC3339.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The OCID of the Vault service key which is the master encryption key for the block / boot volume cross region backups, which will be used in the destination region to encrypt the backup's encryption keys. For more information about the Vault service and encryption keys, see Overview of Vault service and Using Keys.
	XrcKMSKeyID *string `json:"xrcKmsKeyId,omitempty" tf:"xrc_kms_key_id,omitempty"`
}

type VolumeBackupPolicyAssignmentParameters struct {

	// The OCID of the volume or volume group to assign the policy to.
	// +crossplane:generate:reference:type=Volume
	// +kubebuilder:validation:Optional
	AssetID *string `json:"assetId,omitempty" tf:"asset_id,omitempty"`

	// Reference to a Volume to populate assetId.
	// +kubebuilder:validation:Optional
	AssetIDRef *v1.Reference `json:"assetIdRef,omitempty" tf:"-"`

	// Selector for a Volume to populate assetId.
	// +kubebuilder:validation:Optional
	AssetIDSelector *v1.Selector `json:"assetIdSelector,omitempty" tf:"-"`

	// The OCID of the volume backup policy to assign to the volume.
	// +crossplane:generate:reference:type=VolumeBackupPolicy
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a VolumeBackupPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a VolumeBackupPolicy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`

	// The OCID of the Vault service key which is the master encryption key for the block / boot volume cross region backups, which will be used in the destination region to encrypt the backup's encryption keys. For more information about the Vault service and encryption keys, see Overview of Vault service and Using Keys.
	// +kubebuilder:validation:Optional
	XrcKMSKeyID *string `json:"xrcKmsKeyId,omitempty" tf:"xrc_kms_key_id,omitempty"`
}

// VolumeBackupPolicyAssignmentSpec defines the desired state of VolumeBackupPolicyAssignment
type VolumeBackupPolicyAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeBackupPolicyAssignmentParameters `json:"forProvider"`
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
	InitProvider VolumeBackupPolicyAssignmentInitParameters `json:"initProvider,omitempty"`
}

// VolumeBackupPolicyAssignmentStatus defines the observed state of VolumeBackupPolicyAssignment.
type VolumeBackupPolicyAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeBackupPolicyAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VolumeBackupPolicyAssignment is the Schema for the VolumeBackupPolicyAssignments API. Provides the Volume Backup Policy Assignment resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type VolumeBackupPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeBackupPolicyAssignmentSpec   `json:"spec"`
	Status            VolumeBackupPolicyAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeBackupPolicyAssignmentList contains a list of VolumeBackupPolicyAssignments
type VolumeBackupPolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeBackupPolicyAssignment `json:"items"`
}

// Repository type metadata.
var (
	VolumeBackupPolicyAssignment_Kind             = "VolumeBackupPolicyAssignment"
	VolumeBackupPolicyAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeBackupPolicyAssignment_Kind}.String()
	VolumeBackupPolicyAssignment_KindAPIVersion   = VolumeBackupPolicyAssignment_Kind + "." + CRDGroupVersion.String()
	VolumeBackupPolicyAssignment_GroupVersionKind = CRDGroupVersion.WithKind(VolumeBackupPolicyAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeBackupPolicyAssignment{}, &VolumeBackupPolicyAssignmentList{})
}
