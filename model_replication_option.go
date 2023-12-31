/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// ReplicationOption : This will decide how cloud replication will take place
type ReplicationOption string

// List of ReplicationOption
const (
	METADATA_ONLY_ReplicationOption      ReplicationOption = "METADATA_ONLY"
	METADATA_AND_DATA_ReplicationOption  ReplicationOption = "METADATA_AND_DATA"
	KEEP_DATA_IN_CLOUD_ReplicationOption ReplicationOption = "KEEP_DATA_IN_CLOUD"
)
