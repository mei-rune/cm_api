/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// This contains information about the host or host range on which provided host template will be applied.
type ApiClusterTemplateHostInfo struct {
	HostName            string   `json:"hostName,omitempty"`
	HostNameRange       string   `json:"hostNameRange,omitempty"`
	RackId              string   `json:"rackId,omitempty"`
	HostTemplateRefName string   `json:"hostTemplateRefName,omitempty"`
	RoleRefNames        []string `json:"roleRefNames,omitempty"`
}
