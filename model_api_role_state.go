/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// ApiRoleState : Represents the configured run state of a role.
type ApiRoleState string

// List of ApiRoleState
const (
	HISTORY_NOT_AVAILABLE_ApiRoleState ApiRoleState = "HISTORY_NOT_AVAILABLE"
	UNKNOWN_ApiRoleState               ApiRoleState = "UNKNOWN"
	STARTING_ApiRoleState              ApiRoleState = "STARTING"
	STARTED_ApiRoleState               ApiRoleState = "STARTED"
	BUSY_ApiRoleState                  ApiRoleState = "BUSY"
	STOPPING_ApiRoleState              ApiRoleState = "STOPPING"
	STOPPED_ApiRoleState               ApiRoleState = "STOPPED"
	NA_ApiRoleState                    ApiRoleState = "NA"
)