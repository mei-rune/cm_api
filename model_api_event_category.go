/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

type ApiEventCategory string

// List of ApiEventCategory
const (
	UNKNOWN_ApiEventCategory        ApiEventCategory = "UNKNOWN"
	HEALTH_EVENT_ApiEventCategory   ApiEventCategory = "HEALTH_EVENT"
	LOG_EVENT_ApiEventCategory      ApiEventCategory = "LOG_EVENT"
	AUDIT_EVENT_ApiEventCategory    ApiEventCategory = "AUDIT_EVENT"
	ACTIVITY_EVENT_ApiEventCategory ApiEventCategory = "ACTIVITY_EVENT"
	HBASE_ApiEventCategory          ApiEventCategory = "HBASE"
	SYSTEM_ApiEventCategory         ApiEventCategory = "SYSTEM"
)