/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Information about snapshot commands. <p/> This object holds all the information a regular ApiCommand object provides, and adds specific information about the results of a snapshot command. <p/> Depending on the type of the service where the snapshot command was run, a different result property will be populated.
type ApiSnapshotCommand struct {
	// The command ID.
	Id float64 `json:"id,omitempty"`
	// The command name.
	Name string `json:"name,omitempty"`
	// The start time.
	StartTime string `json:"startTime,omitempty"`
	// The end time, if the command is finished.
	EndTime string `json:"endTime,omitempty"`
	// Whether the command is currently active.
	Active bool `json:"active,omitempty"`
	// If the command is finished, whether it was successful.
	Success bool `json:"success,omitempty"`
	// If the command is finished, the result message.
	ResultMessage string `json:"resultMessage,omitempty"`
	// URL to the command's downloadable result data, if any exists.
	ResultDataUrl string          `json:"resultDataUrl,omitempty"`
	ClusterRef    *ApiClusterRef  `json:"clusterRef,omitempty"`
	ServiceRef    *ApiServiceRef  `json:"serviceRef,omitempty"`
	RoleRef       *ApiRoleRef     `json:"roleRef,omitempty"`
	HostRef       *ApiHostRef     `json:"hostRef,omitempty"`
	Parent        *ApiCommand     `json:"parent,omitempty"`
	Children      *ApiCommandList `json:"children,omitempty"`
	// If the command can be retried. Available since V11
	CanRetry    bool                    `json:"canRetry,omitempty"`
	HbaseResult *ApiHBaseSnapshotResult `json:"hbaseResult,omitempty"`
	HdfsResult  *ApiHdfsSnapshotResult  `json:"hdfsResult,omitempty"`
}
