/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

// Detailed information about a Hive replication job.
type ApiHiveReplicationResult struct {
	// Phase the replication is in. <p/> If the replication job is still active, this will contain a string describing the current phase. This will be one of: EXPORT, DATA or IMPORT, for, respectively, exporting the source metastore information, replicating table data (if configured), and importing metastore information in the target. <p/> This value will not be present if the replication is not active. <p/> Available since API v4.
	Phase string `json:"phase,omitempty"`
	// Number of tables that were successfully replicated. Available since API v4.
	TableCount float64 `json:"tableCount,omitempty"`
	// The list of tables successfully replicated. <p/> Since API v4, this is only available in the full view.
	Tables []ApiHiveTable `json:"tables,omitempty"`
	// Number of impala UDFs that were successfully replicated. Available since API v6.
	ImpalaUDFCount float64 `json:"impalaUDFCount,omitempty"`
	// Number of hive UDFs that were successfully replicated. Available since API v14.
	HiveUDFCount float64 `json:"hiveUDFCount,omitempty"`
	// The list of Impala UDFs successfully replicated. Available since API v6 in the full view.
	ImpalaUDFs []ApiImpalaUdf `json:"impalaUDFs,omitempty"`
	// The list of Impala UDFs successfully replicated. Available since API v6 in the full view.
	HiveUDFs []ApiHiveUdf `json:"hiveUDFs,omitempty"`
	// Number of errors detected during replication job. Available since API v4.
	ErrorCount float64 `json:"errorCount,omitempty"`
	// List of errors encountered during replication. <p/> Since API v4, this is only available in the full view.
	Errors                []ApiHiveReplicationError `json:"errors,omitempty"`
	DataReplicationResult *ApiHdfsReplicationResult `json:"dataReplicationResult,omitempty"`
	// Whether this was a dry run.
	DryRun bool `json:"dryRun,omitempty"`
	// Name of the of proxy user, if any. Available since API v11.
	RunAsUser string `json:"runAsUser,omitempty"`
	// Name of the source proxy user, if any. Available since API v18.
	RunOnSourceAsUser string `json:"runOnSourceAsUser,omitempty"`
	// Whether stats are available to display or not. Available since API v19.
	StatsAvailable bool `json:"statsAvailable,omitempty"`
	// Number of Db's Imported/Exported. Available since API v19.
	DbProcessed float64 `json:"dbProcessed,omitempty"`
	// Number of Tables Imported/Exported. Available since API v19.
	TableProcessed float64 `json:"tableProcessed,omitempty"`
	// Number of Partitions Imported/Exported. Available since API v19.
	PartitionProcessed float64 `json:"partitionProcessed,omitempty"`
	// Number of Functions Imported/Exported. Available since API v19.
	FunctionProcessed float64 `json:"functionProcessed,omitempty"`
	// Number of Indexes Imported/Exported. Available since API v19.
	IndexProcessed float64 `json:"indexProcessed,omitempty"`
	// Number of Table and Partitions Statistics Imported/Exported. Available since API v19.
	StatsProcessed float64 `json:"statsProcessed,omitempty"`
	// Number of Db's Expected. Available since API v19.
	DbExpected float64 `json:"dbExpected,omitempty"`
	// Number of Tables Expected. Available since API v19.
	TableExpected float64 `json:"tableExpected,omitempty"`
	// Number of Partitions Expected. Available since API v19.
	PartitionExpected float64 `json:"partitionExpected,omitempty"`
	// Number of Functions Expected. Available since API v19.
	FunctionExpected float64 `json:"functionExpected,omitempty"`
	// Number of Indexes Expected. Available since API v19.
	IndexExpected float64 `json:"indexExpected,omitempty"`
	// Number of Table and Partition Statistics Expected. Available since API v19.
	StatsExpected float64 `json:"statsExpected,omitempty"`
}
