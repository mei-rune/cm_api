/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// A dashboard definition. Dashboards are composed of tsquery-based charts.
type ApiDashboard struct {
	// Returns the dashboard name.
	Name string `json:"name,omitempty"`
	// Returns the json structure for the dashboard. This should be treated as an opaque blob.
	Json string `json:"json,omitempty"`
}
