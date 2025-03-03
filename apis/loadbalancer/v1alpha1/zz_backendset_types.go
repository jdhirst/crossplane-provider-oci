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

type BackendSetBackendInitParameters struct {
}

type BackendSetBackendObservation struct {

	// (Updatable) Whether the load balancer should treat this server as a backup unit. If true, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// (Updatable) Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: false
	Drain *bool `json:"drain,omitempty" tf:"drain,omitempty"`

	// (Updatable) The IP address of the backend server.  Example: 10.0.0.3
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// (Updatable) The maximum number of simultaneous connections the load balancer can make to the backend. If this is not set then the maximum number of simultaneous connections the load balancer can make to the backend is unlimited.  Example: 300
	MaxConnections *float64 `json:"maxConnections,omitempty" tf:"max_connections,omitempty"`

	// A friendly name for the backend set. It must be unique and it cannot be changed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: false
	Offline *bool `json:"offline,omitempty" tf:"offline,omitempty"`

	// (Updatable) The communication port for the backend server.  Example: 8080
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (Updatable) The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see How Load Balancing Policies Work.  Example: 3
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type BackendSetBackendParameters struct {
}

type BackendSetInitParameters struct {

	// (Updatable) The maximum number of simultaneous connections the load balancer can make to any backend in the backend set unless the backend has its own maxConnections setting. If this is not set then the number of simultaneous connections the load balancer can make to any backend in the backend set unless the backend has its own maxConnections setting is unlimited.  Example: 300
	BackendMaxConnections *float64 `json:"backendMaxConnections,omitempty" tf:"backend_max_connections,omitempty"`

	// (Updatable) The health check policy's configuration details.
	HealthChecker []HealthCheckerInitParameters `json:"healthChecker,omitempty" tf:"health_checker,omitempty"`

	// (Updatable) The configuration details for implementing load balancer cookie session persistence (LB cookie stickiness).
	LBCookieSessionPersistenceConfiguration []LBCookieSessionPersistenceConfigurationInitParameters `json:"lbCookieSessionPersistenceConfiguration,omitempty" tf:"lb_cookie_session_persistence_configuration,omitempty"`

	// The OCID of the load balancer on which to add a backend set.
	// +crossplane:generate:reference:type=LoadBalancer
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Reference to a LoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// A friendly name for the backend set. It must be unique and it cannot be changed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) The load balancer policy for the backend set. To get a list of available policies, use the ListPolicies operation.  Example: LEAST_CONNECTIONS
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// (Updatable) The load balancer's SSL handling configuration details.
	SSLConfiguration []SSLConfigurationInitParameters `json:"sslConfiguration,omitempty" tf:"ssl_configuration,omitempty"`

	// (Updatable) The configuration details for implementing session persistence based on a user-specified cookie name (application cookie stickiness).
	SessionPersistenceConfiguration []SessionPersistenceConfigurationInitParameters `json:"sessionPersistenceConfiguration,omitempty" tf:"session_persistence_configuration,omitempty"`
}

type BackendSetObservation struct {

	// (Updatable)
	Backend []BackendSetBackendObservation `json:"backend,omitempty" tf:"backend,omitempty"`

	// (Updatable) The maximum number of simultaneous connections the load balancer can make to any backend in the backend set unless the backend has its own maxConnections setting. If this is not set then the number of simultaneous connections the load balancer can make to any backend in the backend set unless the backend has its own maxConnections setting is unlimited.  Example: 300
	BackendMaxConnections *float64 `json:"backendMaxConnections,omitempty" tf:"backend_max_connections,omitempty"`

	// (Updatable) The health check policy's configuration details.
	HealthChecker []HealthCheckerObservation `json:"healthChecker,omitempty" tf:"health_checker,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Updatable) The configuration details for implementing load balancer cookie session persistence (LB cookie stickiness).
	LBCookieSessionPersistenceConfiguration []LBCookieSessionPersistenceConfigurationObservation `json:"lbCookieSessionPersistenceConfiguration,omitempty" tf:"lb_cookie_session_persistence_configuration,omitempty"`

	// The OCID of the load balancer on which to add a backend set.
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// A friendly name for the backend set. It must be unique and it cannot be changed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) The load balancer policy for the backend set. To get a list of available policies, use the ListPolicies operation.  Example: LEAST_CONNECTIONS
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// (Updatable) The load balancer's SSL handling configuration details.
	SSLConfiguration []SSLConfigurationObservation `json:"sslConfiguration,omitempty" tf:"ssl_configuration,omitempty"`

	// (Updatable) The configuration details for implementing session persistence based on a user-specified cookie name (application cookie stickiness).
	SessionPersistenceConfiguration []SessionPersistenceConfigurationObservation `json:"sessionPersistenceConfiguration,omitempty" tf:"session_persistence_configuration,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type BackendSetParameters struct {

	// (Updatable) The maximum number of simultaneous connections the load balancer can make to any backend in the backend set unless the backend has its own maxConnections setting. If this is not set then the number of simultaneous connections the load balancer can make to any backend in the backend set unless the backend has its own maxConnections setting is unlimited.  Example: 300
	// +kubebuilder:validation:Optional
	BackendMaxConnections *float64 `json:"backendMaxConnections,omitempty" tf:"backend_max_connections,omitempty"`

	// (Updatable) The health check policy's configuration details.
	// +kubebuilder:validation:Optional
	HealthChecker []HealthCheckerParameters `json:"healthChecker,omitempty" tf:"health_checker,omitempty"`

	// (Updatable) The configuration details for implementing load balancer cookie session persistence (LB cookie stickiness).
	// +kubebuilder:validation:Optional
	LBCookieSessionPersistenceConfiguration []LBCookieSessionPersistenceConfigurationParameters `json:"lbCookieSessionPersistenceConfiguration,omitempty" tf:"lb_cookie_session_persistence_configuration,omitempty"`

	// The OCID of the load balancer on which to add a backend set.
	// +crossplane:generate:reference:type=LoadBalancer
	// +kubebuilder:validation:Optional
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Reference to a LoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// A friendly name for the backend set. It must be unique and it cannot be changed.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) The load balancer policy for the backend set. To get a list of available policies, use the ListPolicies operation.  Example: LEAST_CONNECTIONS
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// (Updatable) The load balancer's SSL handling configuration details.
	// +kubebuilder:validation:Optional
	SSLConfiguration []SSLConfigurationParameters `json:"sslConfiguration,omitempty" tf:"ssl_configuration,omitempty"`

	// (Updatable) The configuration details for implementing session persistence based on a user-specified cookie name (application cookie stickiness).
	// +kubebuilder:validation:Optional
	SessionPersistenceConfiguration []SessionPersistenceConfigurationParameters `json:"sessionPersistenceConfiguration,omitempty" tf:"session_persistence_configuration,omitempty"`
}

type HealthCheckerInitParameters struct {

	// (Updatable) The interval between health checks, in milliseconds.  Example: 10000
	IntervalMs *float64 `json:"intervalMs,omitempty" tf:"interval_ms,omitempty"`

	// (Updatable) Specifies if health checks should always be done using plain text instead of depending on whether or not the associated backend set is using SSL.
	IsForcePlainText *bool `json:"isForcePlainText,omitempty" tf:"is_force_plain_text,omitempty"`

	// (Updatable) The communication port for the backend server.  Example: 8080
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (Updatable) The protocol the health check must use; either HTTP or TCP.  Example: HTTP
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Updatable) A regular expression for parsing the response body from the backend server.  Example: ^((?!false).|\s)*$
	ResponseBodyRegex *string `json:"responseBodyRegex,omitempty" tf:"response_body_regex,omitempty"`

	// (Updatable) The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state.  Example: 3
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// (Updatable) The status code a healthy backend server should return.  Example: 200
	ReturnCode *float64 `json:"returnCode,omitempty" tf:"return_code,omitempty"`

	// (Updatable) The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period.  Example: 3000
	TimeoutInMillis *float64 `json:"timeoutInMillis,omitempty" tf:"timeout_in_millis,omitempty"`

	// (Updatable) The path against which to run the health check.  Example: /healthcheck
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`
}

type HealthCheckerObservation struct {

	// (Updatable) The interval between health checks, in milliseconds.  Example: 10000
	IntervalMs *float64 `json:"intervalMs,omitempty" tf:"interval_ms,omitempty"`

	// (Updatable) Specifies if health checks should always be done using plain text instead of depending on whether or not the associated backend set is using SSL.
	IsForcePlainText *bool `json:"isForcePlainText,omitempty" tf:"is_force_plain_text,omitempty"`

	// (Updatable) The communication port for the backend server.  Example: 8080
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (Updatable) The protocol the health check must use; either HTTP or TCP.  Example: HTTP
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Updatable) A regular expression for parsing the response body from the backend server.  Example: ^((?!false).|\s)*$
	ResponseBodyRegex *string `json:"responseBodyRegex,omitempty" tf:"response_body_regex,omitempty"`

	// (Updatable) The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state.  Example: 3
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// (Updatable) The status code a healthy backend server should return.  Example: 200
	ReturnCode *float64 `json:"returnCode,omitempty" tf:"return_code,omitempty"`

	// (Updatable) The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period.  Example: 3000
	TimeoutInMillis *float64 `json:"timeoutInMillis,omitempty" tf:"timeout_in_millis,omitempty"`

	// (Updatable) The path against which to run the health check.  Example: /healthcheck
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`
}

type HealthCheckerParameters struct {

	// (Updatable) The interval between health checks, in milliseconds.  Example: 10000
	// +kubebuilder:validation:Optional
	IntervalMs *float64 `json:"intervalMs,omitempty" tf:"interval_ms,omitempty"`

	// (Updatable) Specifies if health checks should always be done using plain text instead of depending on whether or not the associated backend set is using SSL.
	// +kubebuilder:validation:Optional
	IsForcePlainText *bool `json:"isForcePlainText,omitempty" tf:"is_force_plain_text,omitempty"`

	// (Updatable) The communication port for the backend server.  Example: 8080
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (Updatable) The protocol the health check must use; either HTTP or TCP.  Example: HTTP
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// (Updatable) A regular expression for parsing the response body from the backend server.  Example: ^((?!false).|\s)*$
	// +kubebuilder:validation:Optional
	ResponseBodyRegex *string `json:"responseBodyRegex,omitempty" tf:"response_body_regex,omitempty"`

	// (Updatable) The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state.  Example: 3
	// +kubebuilder:validation:Optional
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// (Updatable) The status code a healthy backend server should return.  Example: 200
	// +kubebuilder:validation:Optional
	ReturnCode *float64 `json:"returnCode,omitempty" tf:"return_code,omitempty"`

	// (Updatable) The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period.  Example: 3000
	// +kubebuilder:validation:Optional
	TimeoutInMillis *float64 `json:"timeoutInMillis,omitempty" tf:"timeout_in_millis,omitempty"`

	// (Updatable) The path against which to run the health check.  Example: /healthcheck
	// +kubebuilder:validation:Optional
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`
}

type LBCookieSessionPersistenceConfigurationInitParameters struct {

	// (Updatable) The name of the cookie inserted by the load balancer. If this field is not configured, the cookie name defaults to "X-Oracle-BMC-LBS-Route".  Example: example_cookie
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// (Updatable) Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: false
	DisableFallback *bool `json:"disableFallback,omitempty" tf:"disable_fallback,omitempty"`

	// (Updatable) The domain in which the cookie is valid. The Set-cookie header inserted by the load balancer contains a domain attribute with the specified value.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (Updatable) Whether the Set-cookie header should contain the HttpOnly attribute. If true, the Set-cookie header inserted by the load balancer contains the HttpOnly attribute, which limits the scope of the cookie to HTTP requests. This attribute directs the client or browser to omit the cookie when providing access to cookies through non-HTTP APIs. For example, it restricts the cookie from JavaScript channels.  Example: true
	IsHTTPOnly *bool `json:"isHttpOnly,omitempty" tf:"is_http_only,omitempty"`

	// (Updatable) Whether the Set-cookie header should contain the Secure attribute. If true, the Set-cookie header inserted by the load balancer contains the Secure attribute, which directs the client or browser to send the cookie only using a secure protocol.
	IsSecure *bool `json:"isSecure,omitempty" tf:"is_secure,omitempty"`

	// (Updatable) The amount of time the cookie remains valid. The Set-cookie header inserted by the load balancer contains a Max-Age attribute with the specified value.
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`

	// (Updatable) The path in which the cookie is valid. The Set-cookie header inserted by the load balancer contains a Path attribute with the specified value.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type LBCookieSessionPersistenceConfigurationObservation struct {

	// (Updatable) The name of the cookie inserted by the load balancer. If this field is not configured, the cookie name defaults to "X-Oracle-BMC-LBS-Route".  Example: example_cookie
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// (Updatable) Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: false
	DisableFallback *bool `json:"disableFallback,omitempty" tf:"disable_fallback,omitempty"`

	// (Updatable) The domain in which the cookie is valid. The Set-cookie header inserted by the load balancer contains a domain attribute with the specified value.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (Updatable) Whether the Set-cookie header should contain the HttpOnly attribute. If true, the Set-cookie header inserted by the load balancer contains the HttpOnly attribute, which limits the scope of the cookie to HTTP requests. This attribute directs the client or browser to omit the cookie when providing access to cookies through non-HTTP APIs. For example, it restricts the cookie from JavaScript channels.  Example: true
	IsHTTPOnly *bool `json:"isHttpOnly,omitempty" tf:"is_http_only,omitempty"`

	// (Updatable) Whether the Set-cookie header should contain the Secure attribute. If true, the Set-cookie header inserted by the load balancer contains the Secure attribute, which directs the client or browser to send the cookie only using a secure protocol.
	IsSecure *bool `json:"isSecure,omitempty" tf:"is_secure,omitempty"`

	// (Updatable) The amount of time the cookie remains valid. The Set-cookie header inserted by the load balancer contains a Max-Age attribute with the specified value.
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`

	// (Updatable) The path in which the cookie is valid. The Set-cookie header inserted by the load balancer contains a Path attribute with the specified value.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type LBCookieSessionPersistenceConfigurationParameters struct {

	// (Updatable) The name of the cookie inserted by the load balancer. If this field is not configured, the cookie name defaults to "X-Oracle-BMC-LBS-Route".  Example: example_cookie
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// (Updatable) Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: false
	// +kubebuilder:validation:Optional
	DisableFallback *bool `json:"disableFallback,omitempty" tf:"disable_fallback,omitempty"`

	// (Updatable) The domain in which the cookie is valid. The Set-cookie header inserted by the load balancer contains a domain attribute with the specified value.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (Updatable) Whether the Set-cookie header should contain the HttpOnly attribute. If true, the Set-cookie header inserted by the load balancer contains the HttpOnly attribute, which limits the scope of the cookie to HTTP requests. This attribute directs the client or browser to omit the cookie when providing access to cookies through non-HTTP APIs. For example, it restricts the cookie from JavaScript channels.  Example: true
	// +kubebuilder:validation:Optional
	IsHTTPOnly *bool `json:"isHttpOnly,omitempty" tf:"is_http_only,omitempty"`

	// (Updatable) Whether the Set-cookie header should contain the Secure attribute. If true, the Set-cookie header inserted by the load balancer contains the Secure attribute, which directs the client or browser to send the cookie only using a secure protocol.
	// +kubebuilder:validation:Optional
	IsSecure *bool `json:"isSecure,omitempty" tf:"is_secure,omitempty"`

	// (Updatable) The amount of time the cookie remains valid. The Set-cookie header inserted by the load balancer contains a Max-Age attribute with the specified value.
	// +kubebuilder:validation:Optional
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`

	// (Updatable) The path in which the cookie is valid. The Set-cookie header inserted by the load balancer contains a Path attribute with the specified value.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type SSLConfigurationInitParameters struct {

	// (Updatable) Ids for Oracle Cloud Infrastructure certificates service certificates. Currently only a single Id may be passed.  Example: [ocid1.certificate.oc1.us-ashburn-1.amaaaaaaav3bgsaa5o2q7rh5nfmkkukfkogasqhk6af2opufhjlqg7m6jqzq]
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// (Updatable) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: example_certificate_bundle
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// (Updatable) The name of the cipher suite to use for HTTPS or SSL connections.
	CipherSuiteName *string `json:"cipherSuiteName,omitempty" tf:"cipher_suite_name,omitempty"`

	// (Updatable) A list of SSL protocols the load balancer must support for HTTPS or SSL connections.
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// (Updatable) When this attribute is set to ENABLED, the system gives preference to the server ciphers over the client ciphers.
	ServerOrderPreference *string `json:"serverOrderPreference,omitempty" tf:"server_order_preference,omitempty"`

	// (Updatable) Ids for Oracle Cloud Infrastructure certificates service CA or CA bundles for the load balancer to trust.  Example: [ocid1.cabundle.oc1.us-ashburn-1.amaaaaaaav3bgsaagl4zzyqdop5i2vuwoqewdvauuw34llqa74otq2jdsfyq]
	TrustedCertificateAuthorityIds []*string `json:"trustedCertificateAuthorityIds,omitempty" tf:"trusted_certificate_authority_ids,omitempty"`

	// (Updatable) The maximum depth for peer certificate chain verification.  Example: 3
	VerifyDepth *float64 `json:"verifyDepth,omitempty" tf:"verify_depth,omitempty"`

	// (Updatable) Whether the load balancer listener should verify peer certificates.  Example: true
	VerifyPeerCertificate *bool `json:"verifyPeerCertificate,omitempty" tf:"verify_peer_certificate,omitempty"`
}

type SSLConfigurationObservation struct {

	// (Updatable) Ids for Oracle Cloud Infrastructure certificates service certificates. Currently only a single Id may be passed.  Example: [ocid1.certificate.oc1.us-ashburn-1.amaaaaaaav3bgsaa5o2q7rh5nfmkkukfkogasqhk6af2opufhjlqg7m6jqzq]
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// (Updatable) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: example_certificate_bundle
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// (Updatable) The name of the cipher suite to use for HTTPS or SSL connections.
	CipherSuiteName *string `json:"cipherSuiteName,omitempty" tf:"cipher_suite_name,omitempty"`

	// (Updatable) A list of SSL protocols the load balancer must support for HTTPS or SSL connections.
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// (Updatable) When this attribute is set to ENABLED, the system gives preference to the server ciphers over the client ciphers.
	ServerOrderPreference *string `json:"serverOrderPreference,omitempty" tf:"server_order_preference,omitempty"`

	// (Updatable) Ids for Oracle Cloud Infrastructure certificates service CA or CA bundles for the load balancer to trust.  Example: [ocid1.cabundle.oc1.us-ashburn-1.amaaaaaaav3bgsaagl4zzyqdop5i2vuwoqewdvauuw34llqa74otq2jdsfyq]
	TrustedCertificateAuthorityIds []*string `json:"trustedCertificateAuthorityIds,omitempty" tf:"trusted_certificate_authority_ids,omitempty"`

	// (Updatable) The maximum depth for peer certificate chain verification.  Example: 3
	VerifyDepth *float64 `json:"verifyDepth,omitempty" tf:"verify_depth,omitempty"`

	// (Updatable) Whether the load balancer listener should verify peer certificates.  Example: true
	VerifyPeerCertificate *bool `json:"verifyPeerCertificate,omitempty" tf:"verify_peer_certificate,omitempty"`
}

type SSLConfigurationParameters struct {

	// (Updatable) Ids for Oracle Cloud Infrastructure certificates service certificates. Currently only a single Id may be passed.  Example: [ocid1.certificate.oc1.us-ashburn-1.amaaaaaaav3bgsaa5o2q7rh5nfmkkukfkogasqhk6af2opufhjlqg7m6jqzq]
	// +kubebuilder:validation:Optional
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// (Updatable) A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: example_certificate_bundle
	// +kubebuilder:validation:Optional
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// (Updatable) The name of the cipher suite to use for HTTPS or SSL connections.
	// +kubebuilder:validation:Optional
	CipherSuiteName *string `json:"cipherSuiteName,omitempty" tf:"cipher_suite_name,omitempty"`

	// (Updatable) A list of SSL protocols the load balancer must support for HTTPS or SSL connections.
	// +kubebuilder:validation:Optional
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// (Updatable) When this attribute is set to ENABLED, the system gives preference to the server ciphers over the client ciphers.
	// +kubebuilder:validation:Optional
	ServerOrderPreference *string `json:"serverOrderPreference,omitempty" tf:"server_order_preference,omitempty"`

	// (Updatable) Ids for Oracle Cloud Infrastructure certificates service CA or CA bundles for the load balancer to trust.  Example: [ocid1.cabundle.oc1.us-ashburn-1.amaaaaaaav3bgsaagl4zzyqdop5i2vuwoqewdvauuw34llqa74otq2jdsfyq]
	// +kubebuilder:validation:Optional
	TrustedCertificateAuthorityIds []*string `json:"trustedCertificateAuthorityIds,omitempty" tf:"trusted_certificate_authority_ids,omitempty"`

	// (Updatable) The maximum depth for peer certificate chain verification.  Example: 3
	// +kubebuilder:validation:Optional
	VerifyDepth *float64 `json:"verifyDepth,omitempty" tf:"verify_depth,omitempty"`

	// (Updatable) Whether the load balancer listener should verify peer certificates.  Example: true
	// +kubebuilder:validation:Optional
	VerifyPeerCertificate *bool `json:"verifyPeerCertificate,omitempty" tf:"verify_peer_certificate,omitempty"`
}

type SessionPersistenceConfigurationInitParameters struct {

	// (Updatable) The name of the cookie inserted by the load balancer. If this field is not configured, the cookie name defaults to "X-Oracle-BMC-LBS-Route".  Example: example_cookie
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// (Updatable) Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: false
	DisableFallback *bool `json:"disableFallback,omitempty" tf:"disable_fallback,omitempty"`
}

type SessionPersistenceConfigurationObservation struct {

	// (Updatable) The name of the cookie inserted by the load balancer. If this field is not configured, the cookie name defaults to "X-Oracle-BMC-LBS-Route".  Example: example_cookie
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// (Updatable) Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: false
	DisableFallback *bool `json:"disableFallback,omitempty" tf:"disable_fallback,omitempty"`
}

type SessionPersistenceConfigurationParameters struct {

	// (Updatable) The name of the cookie inserted by the load balancer. If this field is not configured, the cookie name defaults to "X-Oracle-BMC-LBS-Route".  Example: example_cookie
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName" tf:"cookie_name,omitempty"`

	// (Updatable) Whether the load balancer is prevented from directing traffic from a persistent session client to a different backend server if the original server is unavailable. Defaults to false.  Example: false
	// +kubebuilder:validation:Optional
	DisableFallback *bool `json:"disableFallback,omitempty" tf:"disable_fallback,omitempty"`
}

// BackendSetSpec defines the desired state of BackendSet
type BackendSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendSetParameters `json:"forProvider"`
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
	InitProvider BackendSetInitParameters `json:"initProvider,omitempty"`
}

// BackendSetStatus defines the observed state of BackendSet.
type BackendSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackendSet is the Schema for the BackendSets API. Provides the Backend Set resource in Oracle Cloud Infrastructure Load Balancer service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type BackendSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.healthChecker) || (has(self.initProvider) && has(self.initProvider.healthChecker))",message="spec.forProvider.healthChecker is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   BackendSetSpec   `json:"spec"`
	Status BackendSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendSetList contains a list of BackendSets
type BackendSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendSet `json:"items"`
}

// Repository type metadata.
var (
	BackendSet_Kind             = "BackendSet"
	BackendSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackendSet_Kind}.String()
	BackendSet_KindAPIVersion   = BackendSet_Kind + "." + CRDGroupVersion.String()
	BackendSet_GroupVersionKind = CRDGroupVersion.WithKind(BackendSet_Kind)
)

func init() {
	SchemeBuilder.Register(&BackendSet{}, &BackendSetList{})
}
