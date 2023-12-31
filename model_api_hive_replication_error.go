/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// A Hive replication error.
type ApiHiveReplicationError struct {
	// Name of the database.
	Database string `json:"database,omitempty"`
	// Name of the table.
	TableName string `json:"tableName,omitempty"`
	// UDF signature, includes the UDF name and parameter types.
	ImpalaUDF string `json:"impalaUDF,omitempty"`
	// Hive UDF signature, includes the UDF name and parameter types.
	HiveUDF string `json:"hiveUDF,omitempty"`
	// Description of the error.
	Error_ string `json:"error,omitempty"`
}
