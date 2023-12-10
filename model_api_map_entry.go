/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Models a map entry, with a key and a value. By forming a list of these entries you can have the equivalent of Map&lt;String, String&gt; (since JAX-B doesn't support maps).
type ApiMapEntry struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}