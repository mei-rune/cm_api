# ApiImpalaTenantUtilization

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantName** | **string** | Name of the tenant. | [optional] [default to null]
**TotalQueries** | **float64** | Total number of queries submitted to Impala. | [optional] [default to null]
**SuccessfulQueries** | **float64** | Number of queries that finished successfully. | [optional] [default to null]
**OomQueries** | **float64** | Number of queries that failed due to insufficient memory. | [optional] [default to null]
**TimeOutQueries** | **float64** | Number of queries that timed out while waiting for resources in a pool. | [optional] [default to null]
**RejectedQueries** | **float64** | Number of queries that were rejected by Impala because the pool was full. | [optional] [default to null]
**AvgWaitTimeInQueue** | **float64** | Average time, in milliseconds, spent by a query in an Impala pool while waiting for resources. | [optional] [default to null]
**PeakAllocationTimestampMS** | **float64** | The time when Impala reserved the maximum amount of memory for queries. | [optional] [default to null]
**MaxAllocatedMemory** | **float64** | The maximum memory (in bytes) that was reserved by Impala for executing queries. | [optional] [default to null]
**MaxAllocatedMemoryPercentage** | **float64** | The maximum percentage of memory that was reserved by Impala for executing queries. | [optional] [default to null]
**UtilizedAtMaxAllocated** | **float64** | The amount of memory (in bytes) used by Impala for running queries at the time when maximum memory was reserved. | [optional] [default to null]
**UtilizedAtMaxAllocatedPercentage** | **float64** | The percentage of memory used by Impala for running queries at the time when maximum memory was reserved. | [optional] [default to null]
**PeakUsageTimestampMS** | **float64** | The time when Impala used the maximum amount of memory for queries. | [optional] [default to null]
**MaxUtilizedMemory** | **float64** | The maximum memory (in bytes) that was used by Impala for executing queries. | [optional] [default to null]
**MaxUtilizedMemoryPercentage** | **float64** | The maximum percentage of memory that was used by Impala for executing queries. | [optional] [default to null]
**AllocatedAtMaxUtilized** | **float64** | The amount of memory (in bytes) reserved by Impala at the time when it was using the maximum memory for executing queries. | [optional] [default to null]
**AllocatedAtMaxUtilizedPercentage** | **float64** | The percentage of memory reserved by Impala at the time when it was using the maximum memory for executing queries. | [optional] [default to null]
**DistributionUtilizedByImpalaDaemon** | [***ApiImpalaUtilizationHistogram**](ApiImpalaUtilizationHistogram.md) |  | [optional] [default to null]
**DistributionAllocatedByImpalaDaemon** | [***ApiImpalaUtilizationHistogram**](ApiImpalaUtilizationHistogram.md) |  | [optional] [default to null]
**AvgSpilledMemory** | **float64** | Average spill per query. | [optional] [default to null]
**MaxSpilledMemory** | **float64** | Maximum spill per query. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

