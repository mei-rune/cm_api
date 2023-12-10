/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Arguments used when enabling HDFS automatic failover.
type ApiHdfsFailoverArguments struct {
	// Nameservice for which to enable automatic failover.
	Nameservice      string         `json:"nameservice,omitempty"`
	ZooKeeperService *ApiServiceRef `json:"zooKeeperService,omitempty"`
	// Name of the failover controller to create for the active NameNode.
	ActiveFCName string `json:"activeFCName,omitempty"`
	// Name of the failover controller to create for the stand-by NameNode.
	StandByFCName string `json:"standByFCName,omitempty"`
}
