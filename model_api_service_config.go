/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Service and role type configuration.
type ApiServiceConfig struct {
	Items []ApiConfig `json:"items,omitempty"`
	// List of role type configurations. Only available up to API v2.
	RoleTypeConfigs []ApiRoleTypeConfig `json:"roleTypeConfigs,omitempty"`
}
