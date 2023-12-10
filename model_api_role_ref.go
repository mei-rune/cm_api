/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// A roleRef references a role. Each role is identified by its \"roleName\", the \"serviceName\" for the service it belongs to, and the \"clusterName\" in which the service resides. To operate on the role object, use the API with the those fields as parameters.
type ApiRoleRef struct {
	ClusterName string `json:"clusterName,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
	RoleName    string `json:"roleName,omitempty"`
}
