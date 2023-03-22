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

type BalancerPathRouteSetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type BalancerPathRouteSetParameters struct {

	// The OCID of the load balancer to add the path route set to.
	// +crossplane:generate:reference:type=BalancerLoadBalancer
	// +kubebuilder:validation:Optional
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Reference to a BalancerLoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// Selector for a BalancerLoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// The name for this set of path route rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: example_path_route_set
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// (Updatable) The set of path route rules.
	// +kubebuilder:validation:Required
	PathRoutes []PathRoutesParameters `json:"pathRoutes" tf:"path_routes,omitempty"`
}

type PathMatchTypeObservation struct {
}

type PathMatchTypeParameters struct {

	// (Updatable) Specifies how the load balancing service compares a PathRoute object's path string against the incoming URI.
	// +kubebuilder:validation:Required
	MatchType *string `json:"matchType" tf:"match_type,omitempty"`
}

type PathRoutesObservation struct {
}

type PathRoutesParameters struct {

	// (Updatable) The name of the target backend set for requests where the incoming URI matches the specified path.  Example: example_backend_set
	// +crossplane:generate:reference:type=BalancerBackendSet
	// +kubebuilder:validation:Optional
	BackendSetName *string `json:"backendSetName,omitempty" tf:"backend_set_name,omitempty"`

	// Reference to a BalancerBackendSet to populate backendSetName.
	// +kubebuilder:validation:Optional
	BackendSetNameRef *v1.Reference `json:"backendSetNameRef,omitempty" tf:"-"`

	// Selector for a BalancerBackendSet to populate backendSetName.
	// +kubebuilder:validation:Optional
	BackendSetNameSelector *v1.Selector `json:"backendSetNameSelector,omitempty" tf:"-"`

	// (Updatable) The path string to match against the incoming URI path.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// (Updatable) The type of matching to apply to incoming URIs.
	// +kubebuilder:validation:Required
	PathMatchType []PathMatchTypeParameters `json:"pathMatchType" tf:"path_match_type,omitempty"`
}

// BalancerPathRouteSetSpec defines the desired state of BalancerPathRouteSet
type BalancerPathRouteSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BalancerPathRouteSetParameters `json:"forProvider"`
}

// BalancerPathRouteSetStatus defines the observed state of BalancerPathRouteSet.
type BalancerPathRouteSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BalancerPathRouteSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerPathRouteSet is the Schema for the BalancerPathRouteSets API. Provides the Path Route Set resource in Oracle Cloud Infrastructure Load Balancer service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type BalancerPathRouteSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BalancerPathRouteSetSpec   `json:"spec"`
	Status            BalancerPathRouteSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerPathRouteSetList contains a list of BalancerPathRouteSets
type BalancerPathRouteSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BalancerPathRouteSet `json:"items"`
}

// Repository type metadata.
var (
	BalancerPathRouteSet_Kind             = "BalancerPathRouteSet"
	BalancerPathRouteSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BalancerPathRouteSet_Kind}.String()
	BalancerPathRouteSet_KindAPIVersion   = BalancerPathRouteSet_Kind + "." + CRDGroupVersion.String()
	BalancerPathRouteSet_GroupVersionKind = CRDGroupVersion.WithKind(BalancerPathRouteSet_Kind)
)

func init() {
	SchemeBuilder.Register(&BalancerPathRouteSet{}, &BalancerPathRouteSetList{})
}
