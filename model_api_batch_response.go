/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// A batch response, comprised of one or more response elements.
type ApiBatchResponse struct {
	Items []ApiBatchResponseElement `json:"items,omitempty"`
	// Read-only. True if every response element's ApiBatchResponseElement#getStatusCode() is in the range [200, 300), false otherwise.
	Success bool `json:"success,omitempty"`
}
