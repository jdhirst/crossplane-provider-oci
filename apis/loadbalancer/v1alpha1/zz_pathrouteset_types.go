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

type PathMatchTypeInitParameters struct {

	// (Updatable) Specifies how the load balancing service compares a PathRoute object's path string against the incoming URI.
	MatchType *string `json:"matchType,omitempty" tf:"match_type,omitempty"`
}

type PathMatchTypeObservation struct {

	// (Updatable) Specifies how the load balancing service compares a PathRoute object's path string against the incoming URI.
	MatchType *string `json:"matchType,omitempty" tf:"match_type,omitempty"`
}

type PathMatchTypeParameters struct {

	// (Updatable) Specifies how the load balancing service compares a PathRoute object's path string against the incoming URI.
	// +kubebuilder:validation:Optional
	MatchType *string `json:"matchType" tf:"match_type,omitempty"`
}

type PathRouteSetInitParameters struct {

	// The OCID of the load balancer to add the path route set to.
	// +crossplane:generate:reference:type=LoadBalancer
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Reference to a LoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// The name for this set of path route rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: example_path_route_set
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) The set of path route rules.
	PathRoutes []PathRoutesInitParameters `json:"pathRoutes,omitempty" tf:"path_routes,omitempty"`
}

type PathRouteSetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The OCID of the load balancer to add the path route set to.
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// The name for this set of path route rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: example_path_route_set
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) The set of path route rules.
	PathRoutes []PathRoutesObservation `json:"pathRoutes,omitempty" tf:"path_routes,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type PathRouteSetParameters struct {

	// The OCID of the load balancer to add the path route set to.
	// +crossplane:generate:reference:type=LoadBalancer
	// +kubebuilder:validation:Optional
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Reference to a LoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// The name for this set of path route rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: example_path_route_set
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) The set of path route rules.
	// +kubebuilder:validation:Optional
	PathRoutes []PathRoutesParameters `json:"pathRoutes,omitempty" tf:"path_routes,omitempty"`
}

type PathRoutesInitParameters struct {

	// (Updatable) The name of the target backend set for requests where the incoming URI matches the specified path.  Example: example_backend_set
	// +crossplane:generate:reference:type=BackendSet
	BackendSetName *string `json:"backendSetName,omitempty" tf:"backend_set_name,omitempty"`

	// Reference to a BackendSet to populate backendSetName.
	// +kubebuilder:validation:Optional
	BackendSetNameRef *v1.Reference `json:"backendSetNameRef,omitempty" tf:"-"`

	// Selector for a BackendSet to populate backendSetName.
	// +kubebuilder:validation:Optional
	BackendSetNameSelector *v1.Selector `json:"backendSetNameSelector,omitempty" tf:"-"`

	// (Updatable) The path string to match against the incoming URI path.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (Updatable) The type of matching to apply to incoming URIs.
	PathMatchType []PathMatchTypeInitParameters `json:"pathMatchType,omitempty" tf:"path_match_type,omitempty"`
}

type PathRoutesObservation struct {

	// (Updatable) The name of the target backend set for requests where the incoming URI matches the specified path.  Example: example_backend_set
	BackendSetName *string `json:"backendSetName,omitempty" tf:"backend_set_name,omitempty"`

	// (Updatable) The path string to match against the incoming URI path.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (Updatable) The type of matching to apply to incoming URIs.
	PathMatchType []PathMatchTypeObservation `json:"pathMatchType,omitempty" tf:"path_match_type,omitempty"`
}

type PathRoutesParameters struct {

	// (Updatable) The name of the target backend set for requests where the incoming URI matches the specified path.  Example: example_backend_set
	// +crossplane:generate:reference:type=BackendSet
	// +kubebuilder:validation:Optional
	BackendSetName *string `json:"backendSetName,omitempty" tf:"backend_set_name,omitempty"`

	// Reference to a BackendSet to populate backendSetName.
	// +kubebuilder:validation:Optional
	BackendSetNameRef *v1.Reference `json:"backendSetNameRef,omitempty" tf:"-"`

	// Selector for a BackendSet to populate backendSetName.
	// +kubebuilder:validation:Optional
	BackendSetNameSelector *v1.Selector `json:"backendSetNameSelector,omitempty" tf:"-"`

	// (Updatable) The path string to match against the incoming URI path.
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`

	// (Updatable) The type of matching to apply to incoming URIs.
	// +kubebuilder:validation:Optional
	PathMatchType []PathMatchTypeParameters `json:"pathMatchType" tf:"path_match_type,omitempty"`
}

// PathRouteSetSpec defines the desired state of PathRouteSet
type PathRouteSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PathRouteSetParameters `json:"forProvider"`
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
	InitProvider PathRouteSetInitParameters `json:"initProvider,omitempty"`
}

// PathRouteSetStatus defines the observed state of PathRouteSet.
type PathRouteSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PathRouteSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PathRouteSet is the Schema for the PathRouteSets API. Provides the Path Route Set resource in Oracle Cloud Infrastructure Load Balancer service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type PathRouteSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pathRoutes) || (has(self.initProvider) && has(self.initProvider.pathRoutes))",message="spec.forProvider.pathRoutes is a required parameter"
	Spec   PathRouteSetSpec   `json:"spec"`
	Status PathRouteSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PathRouteSetList contains a list of PathRouteSets
type PathRouteSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PathRouteSet `json:"items"`
}

// Repository type metadata.
var (
	PathRouteSet_Kind             = "PathRouteSet"
	PathRouteSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PathRouteSet_Kind}.String()
	PathRouteSet_KindAPIVersion   = PathRouteSet_Kind + "." + CRDGroupVersion.String()
	PathRouteSet_GroupVersionKind = CRDGroupVersion.WithKind(PathRouteSet_Kind)
)

func init() {
	SchemeBuilder.Register(&PathRouteSet{}, &PathRouteSetList{})
}
