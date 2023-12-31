/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Arguments used for the Cloudera Manager level performance inspector. Network diagnostics will be run from every host in sourceHostList to every host in targetHostList.
type ApiHostsPerfInspectorArgs struct {
	SourceHostList *ApiHostNameList          `json:"sourceHostList,omitempty"`
	TargetHostList *ApiHostNameList          `json:"targetHostList,omitempty"`
	PingArgs       *ApiPerfInspectorPingArgs `json:"pingArgs,omitempty"`
}
