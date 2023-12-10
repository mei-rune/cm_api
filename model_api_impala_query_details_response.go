/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// A query details response.
type ApiImpalaQueryDetailsResponse struct {
	// The details for this query. Two formats are supported: <ul> <li> 'text': this is a text based, human readable representation of the Impala runtime profile. </li> <li> 'thrift_encoded': this a compact-thrift, base64 encoded representation of the impala RuntimeProfile.thrift object. See the Impala documentation for more details. </li> </ul>
	Details string `json:"details,omitempty"`
}
