/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// ApiClusterVersion : The CDH version of the cluster.
type ApiClusterVersion string

// List of ApiClusterVersion
const (
	CDH3_ApiClusterVersion     ApiClusterVersion = "CDH3"
	CDH3U4_X_ApiClusterVersion ApiClusterVersion = "CDH3u4X"
	CDH4_ApiClusterVersion     ApiClusterVersion = "CDH4"
	CDH5_ApiClusterVersion     ApiClusterVersion = "CDH5"
	CDH6_ApiClusterVersion     ApiClusterVersion = "CDH6"
	UNKNOWN_ApiClusterVersion  ApiClusterVersion = "UNKNOWN"
)
