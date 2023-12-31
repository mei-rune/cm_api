/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Represents MapReduce2 information for a YARN application.
type ApiMr2AppInformation struct {
	// The state of the job. This is only set on completed jobs. This can take on the following values: \"NEW\", \"INITED\", \"RUNNING\", \"SUCCEEDED\", \"FAILED\", \"KILLED\", \"ERROR\".
	JobState string `json:"jobState,omitempty"`
}
