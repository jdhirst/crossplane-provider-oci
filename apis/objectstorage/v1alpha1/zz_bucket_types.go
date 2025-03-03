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

type BucketInitParameters struct {

	// (Updatable) The type of public access enabled on this bucket. A bucket is set to NoPublicAccess by default, which only allows an authenticated caller to access the bucket and its contents. When ObjectRead is enabled on the bucket, public access is allowed for the GetObject, HeadObject, and ListObjects operations. When ObjectReadWithoutList is enabled on the bucket, public access is allowed for the GetObject and HeadObject operations.
	AccessType *string `json:"accessType,omitempty" tf:"access_type,omitempty"`

	// (Updatable) Set the auto tiering status on the bucket. By default, a bucket is created with auto tiering Disabled. Use this option to enable auto tiering during bucket creation. Objects in a bucket with auto tiering set to InfrequentAccess are transitioned automatically between the 'Standard' and 'InfrequentAccess' tiers based on the access pattern of the objects.
	AutoTiering *string `json:"autoTiering,omitempty" tf:"auto_tiering,omitempty"`

	// (Updatable) The ID of the compartment in which to create the bucket.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The OCID of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/kms/v1alpha1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// (Updatable) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods. Bucket names must be unique within an Object Storage namespace. Avoid entering confidential information. example: Example: my-new-bucket1
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Object Storage namespace used for the request.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (Updatable) Whether or not events are emitted for object state changes in this bucket. By default, objectEventsEnabled is set to false. Set objectEventsEnabled to true to emit events for object state changes. For more information about events, see Overview of Events.
	ObjectEventsEnabled *bool `json:"objectEventsEnabled,omitempty" tf:"object_events_enabled,omitempty"`

	// (Updatable) Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds. Note that a maximum of 100 rules are supported on a bucket.
	RetentionRules []RetentionRulesInitParameters `json:"retentionRules,omitempty" tf:"retention_rules,omitempty"`

	// The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created.
	StorageTier *string `json:"storageTier,omitempty" tf:"storage_tier,omitempty"`

	// (Updatable) Set the versioning status on the bucket. By default, a bucket is created with versioning Disabled. Use this option to enable versioning during bucket creation. Objects in a version enabled bucket are protected from overwrites and deletions. Previous versions of the same object will be available in the bucket. Allowed Create values: Enabled, Disabled. Allowed Update values: Enabled, Suspended.
	Versioning *string `json:"versioning,omitempty" tf:"versioning,omitempty"`
}

type BucketObservation struct {

	// (Updatable) The type of public access enabled on this bucket. A bucket is set to NoPublicAccess by default, which only allows an authenticated caller to access the bucket and its contents. When ObjectRead is enabled on the bucket, public access is allowed for the GetObject, HeadObject, and ListObjects operations. When ObjectReadWithoutList is enabled on the bucket, public access is allowed for the GetObject and HeadObject operations.
	AccessType *string `json:"accessType,omitempty" tf:"access_type,omitempty"`

	// The approximate number of objects in the bucket. Count statistics are reported periodically. You will see a lag between what is displayed and the actual object count.
	ApproximateCount *string `json:"approximateCount,omitempty" tf:"approximate_count,omitempty"`

	// The approximate total size in bytes of all objects in the bucket. Size statistics are reported periodically. You will see a lag between what is displayed and the actual size of the bucket.
	ApproximateSize *string `json:"approximateSize,omitempty" tf:"approximate_size,omitempty"`

	// (Updatable) Set the auto tiering status on the bucket. By default, a bucket is created with auto tiering Disabled. Use this option to enable auto tiering during bucket creation. Objects in a bucket with auto tiering set to InfrequentAccess are transitioned automatically between the 'Standard' and 'InfrequentAccess' tiers based on the access pattern of the objects.
	AutoTiering *string `json:"autoTiering,omitempty" tf:"auto_tiering,omitempty"`

	// The OCID of the bucket which is a Oracle assigned unique identifier for this resource type (bucket). bucket_id cannot be used for bucket lookup.
	BucketID *string `json:"bucketId,omitempty" tf:"bucket_id,omitempty"`

	// (Updatable) The ID of the compartment in which to create the bucket.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// The OCID of the user who created the bucket.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: {"Operations.CostCenter": "42"}
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// The entity tag (ETag) for the bucket.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether or not this bucket is read only. By default, isReadOnly is set to false. This will be set to 'true' when this bucket is configured as a destination in a replication policy.
	IsReadOnly *bool `json:"isReadOnly,omitempty" tf:"is_read_only,omitempty"`

	// (Updatable) The OCID of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// (Updatable) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods. Bucket names must be unique within an Object Storage namespace. Avoid entering confidential information. example: Example: my-new-bucket1
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Object Storage namespace used for the request.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (Updatable) Whether or not events are emitted for object state changes in this bucket. By default, objectEventsEnabled is set to false. Set objectEventsEnabled to true to emit events for object state changes. For more information about events, see Overview of Events.
	ObjectEventsEnabled *bool `json:"objectEventsEnabled,omitempty" tf:"object_events_enabled,omitempty"`

	// The entity tag (ETag) for the live object lifecycle policy on the bucket.
	ObjectLifecyclePolicyEtag *string `json:"objectLifecyclePolicyEtag,omitempty" tf:"object_lifecycle_policy_etag,omitempty"`

	// Whether or not this bucket is a replication source. By default, replicationEnabled is set to false. This will be set to 'true' when you create a replication policy for the bucket.
	ReplicationEnabled *bool `json:"replicationEnabled,omitempty" tf:"replication_enabled,omitempty"`

	// (Updatable) Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds. Note that a maximum of 100 rules are supported on a bucket.
	RetentionRules []RetentionRulesObservation `json:"retentionRules,omitempty" tf:"retention_rules,omitempty"`

	// The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created.
	StorageTier *string `json:"storageTier,omitempty" tf:"storage_tier,omitempty"`

	// The date and time that the retention rule was created as per RFC3339.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// (Updatable) Set the versioning status on the bucket. By default, a bucket is created with versioning Disabled. Use this option to enable versioning during bucket creation. Objects in a version enabled bucket are protected from overwrites and deletions. Previous versions of the same object will be available in the bucket. Allowed Create values: Enabled, Disabled. Allowed Update values: Enabled, Suspended.
	Versioning *string `json:"versioning,omitempty" tf:"versioning,omitempty"`
}

type BucketParameters struct {

	// (Updatable) The type of public access enabled on this bucket. A bucket is set to NoPublicAccess by default, which only allows an authenticated caller to access the bucket and its contents. When ObjectRead is enabled on the bucket, public access is allowed for the GetObject, HeadObject, and ListObjects operations. When ObjectReadWithoutList is enabled on the bucket, public access is allowed for the GetObject and HeadObject operations.
	// +kubebuilder:validation:Optional
	AccessType *string `json:"accessType,omitempty" tf:"access_type,omitempty"`

	// (Updatable) Set the auto tiering status on the bucket. By default, a bucket is created with auto tiering Disabled. Use this option to enable auto tiering during bucket creation. Objects in a bucket with auto tiering set to InfrequentAccess are transitioned automatically between the 'Standard' and 'InfrequentAccess' tiers based on the access pattern of the objects.
	// +kubebuilder:validation:Optional
	AutoTiering *string `json:"autoTiering,omitempty" tf:"auto_tiering,omitempty"`

	// (Updatable) The ID of the compartment in which to create the bucket.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The OCID of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/kms/v1alpha1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// (Updatable) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods. Bucket names must be unique within an Object Storage namespace. Avoid entering confidential information. example: Example: my-new-bucket1
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Object Storage namespace used for the request.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (Updatable) Whether or not events are emitted for object state changes in this bucket. By default, objectEventsEnabled is set to false. Set objectEventsEnabled to true to emit events for object state changes. For more information about events, see Overview of Events.
	// +kubebuilder:validation:Optional
	ObjectEventsEnabled *bool `json:"objectEventsEnabled,omitempty" tf:"object_events_enabled,omitempty"`

	// (Updatable) Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds. Note that a maximum of 100 rules are supported on a bucket.
	// +kubebuilder:validation:Optional
	RetentionRules []RetentionRulesParameters `json:"retentionRules,omitempty" tf:"retention_rules,omitempty"`

	// The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created.
	// +kubebuilder:validation:Optional
	StorageTier *string `json:"storageTier,omitempty" tf:"storage_tier,omitempty"`

	// (Updatable) Set the versioning status on the bucket. By default, a bucket is created with versioning Disabled. Use this option to enable versioning during bucket creation. Objects in a version enabled bucket are protected from overwrites and deletions. Previous versions of the same object will be available in the bucket. Allowed Create values: Enabled, Disabled. Allowed Update values: Enabled, Suspended.
	// +kubebuilder:validation:Optional
	Versioning *string `json:"versioning,omitempty" tf:"versioning,omitempty"`
}

type DurationInitParameters struct {

	// (Updatable) The timeAmount is interpreted in units defined by the timeUnit parameter, and is calculated in relation to each object's Last-Modified timestamp.
	TimeAmount *string `json:"timeAmount,omitempty" tf:"time_amount,omitempty"`

	// (Updatable) The unit that should be used to interpret timeAmount.
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DurationObservation struct {

	// (Updatable) The timeAmount is interpreted in units defined by the timeUnit parameter, and is calculated in relation to each object's Last-Modified timestamp.
	TimeAmount *string `json:"timeAmount,omitempty" tf:"time_amount,omitempty"`

	// (Updatable) The unit that should be used to interpret timeAmount.
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`
}

type DurationParameters struct {

	// (Updatable) The timeAmount is interpreted in units defined by the timeUnit parameter, and is calculated in relation to each object's Last-Modified timestamp.
	// +kubebuilder:validation:Optional
	TimeAmount *string `json:"timeAmount" tf:"time_amount,omitempty"`

	// (Updatable) The unit that should be used to interpret timeAmount.
	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit" tf:"time_unit,omitempty"`
}

type RetentionRulesInitParameters struct {

	// A user-specified name for the retention rule. Names can be helpful in identifying retention rules. The name should be unique. This attribute is a forcenew attribute
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable)
	Duration []DurationInitParameters `json:"duration,omitempty" tf:"duration,omitempty"`

	// (Updatable) The date and time as per RFC 3339 after which this rule is locked and can only be deleted by deleting the bucket. Once a rule is locked, only increases in the duration are allowed and no other properties can be changed. This property cannot be updated for rules that are in a locked state. Specifying it when a duration is not specified is considered an error.
	TimeRuleLocked *string `json:"timeRuleLocked,omitempty" tf:"time_rule_locked,omitempty"`
}

type RetentionRulesObservation struct {

	// A user-specified name for the retention rule. Names can be helpful in identifying retention rules. The name should be unique. This attribute is a forcenew attribute
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable)
	Duration []DurationObservation `json:"duration,omitempty" tf:"duration,omitempty"`

	// Unique identifier for the retention rule.
	RetentionRuleID *string `json:"retentionRuleId,omitempty" tf:"retention_rule_id,omitempty"`

	// The date and time that the retention rule was created as per RFC3339.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The date and time that the retention rule was modified as per RFC3339.
	TimeModified *string `json:"timeModified,omitempty" tf:"time_modified,omitempty"`

	// (Updatable) The date and time as per RFC 3339 after which this rule is locked and can only be deleted by deleting the bucket. Once a rule is locked, only increases in the duration are allowed and no other properties can be changed. This property cannot be updated for rules that are in a locked state. Specifying it when a duration is not specified is considered an error.
	TimeRuleLocked *string `json:"timeRuleLocked,omitempty" tf:"time_rule_locked,omitempty"`
}

type RetentionRulesParameters struct {

	// A user-specified name for the retention rule. Names can be helpful in identifying retention rules. The name should be unique. This attribute is a forcenew attribute
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// (Updatable)
	// +kubebuilder:validation:Optional
	Duration []DurationParameters `json:"duration,omitempty" tf:"duration,omitempty"`

	// (Updatable) The date and time as per RFC 3339 after which this rule is locked and can only be deleted by deleting the bucket. Once a rule is locked, only increases in the duration are allowed and no other properties can be changed. This property cannot be updated for rules that are in a locked state. Specifying it when a duration is not specified is considered an error.
	// +kubebuilder:validation:Optional
	TimeRuleLocked *string `json:"timeRuleLocked,omitempty" tf:"time_rule_locked,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
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
	InitProvider BucketInitParameters `json:"initProvider,omitempty"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Bucket is the Schema for the Buckets API. Provides the Bucket resource in Oracle Cloud Infrastructure Object Storage service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.__namespace__) || (has(self.initProvider) && has(self.initProvider.__namespace__))",message="spec.forProvider.namespace is a required parameter"
	Spec   BucketSpec   `json:"spec"`
	Status BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
