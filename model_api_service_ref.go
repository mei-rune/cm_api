/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// A serviceRef references a service. It is identified by the \"serviceName\", \"clusterName\" (name of the cluster which the service belongs to) and an optional \"peerName\" (to reference a remote service i.e. services managed by other CM instances). To operate on the service object, use the API with those fields as parameters.
type ApiServiceRef struct {
	// The name of the CM peer corresponding to the remote CM that manages the referenced service. This should only be set when referencing a remote service.
	PeerName string `json:"peerName,omitempty"`
	// The enclosing cluster for this service.
	ClusterName string `json:"clusterName,omitempty"`
	// The service name.
	ServiceName        string `json:"serviceName,omitempty"`
	ServiceDisplayName string `json:"serviceDisplayName,omitempty"`
	// The service type. This is available since version 32
	ServiceType string `json:"serviceType,omitempty"`
}
