# ApiHiveReplicationArguments

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceService** | [***ApiServiceRef**](ApiServiceRef.md) |  | [optional] [default to null]
**TableFilters** | [**[]ApiHiveTable**](ApiHiveTable.md) | Filters for tables to include in the replication. Optional. If not provided, include all tables in all databases. | [optional] [default to null]
**ExportDir** | **string** | Directory, in the HDFS service where the target Hive service&#x27;s data is stored, where the export file will be saved. Optional. If not provided, Cloudera Manager will pick a directory for storing the data. | [optional] [default to null]
**Force** | **bool** | Whether to force overwriting of mismatched tables. | [optional] [default to null]
**ReplicateData** | **bool** | Whether to replicate table data stored in HDFS. &lt;p/&gt; If set, the \&quot;hdfsArguments\&quot; property must be set to configure the HDFS replication job. | [optional] [default to null]
**HdfsArguments** | [***ApiHdfsReplicationArguments**](ApiHdfsReplicationArguments.md) |  | [optional] [default to null]
**ReplicateImpalaMetadata** | **bool** | Whether to replicate the impala metadata. (i.e. the metadata for impala UDFs and their corresponding binaries in HDFS). | [optional] [default to null]
**RunInvalidateMetadata** | **bool** | Whether to run invalidate metadata query or not | [optional] [default to null]
**DryRun** | **bool** | Whether to perform a dry run. Defaults to false. | [optional] [default to null]
**NumThreads** | **float64** | Number of threads to use in multi-threaded export/import phase | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

