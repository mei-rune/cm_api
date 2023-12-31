# ApiYarnTenantUtilization

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantName** | **string** | Name of the tenant. | [optional] [default to null]
**AvgYarnCpuAllocation** | **float64** | Average number of VCores allocated to YARN applications of the tenant. | [optional] [default to null]
**AvgYarnCpuUtilization** | **float64** | Average number of VCores used by YARN applications of the tenant. | [optional] [default to null]
**AvgYarnCpuUnusedCapacity** | **float64** | Average unused VCores of the tenant. | [optional] [default to null]
**AvgYarnCpuSteadyFairShare** | **float64** | Average steady fair share VCores. | [optional] [default to null]
**AvgYarnPoolAllocatedCpuDuringContention** | **float64** | Average allocated Vcores with pending containers. | [optional] [default to null]
**AvgYarnPoolFairShareCpuDuringContention** | **float64** | Average fair share VCores with pending containers. | [optional] [default to null]
**AvgYarnPoolSteadyFairShareCpuDuringContention** | **float64** | Average steady fair share VCores with pending containers. | [optional] [default to null]
**AvgYarnContainerWaitRatio** | **float64** | Average percentage of pending containers for the pool during periods of contention. | [optional] [default to null]
**AvgYarnMemoryAllocation** | **float64** | Average memory allocated to YARN applications of the tenant. | [optional] [default to null]
**AvgYarnMemoryUtilization** | **float64** | Average memory used by YARN applications of the tenant. | [optional] [default to null]
**AvgYarnMemoryUnusedCapacity** | **float64** | Average unused memory of the tenant. | [optional] [default to null]
**AvgYarnMemorySteadyFairShare** | **float64** | Average steady fair share memory. | [optional] [default to null]
**AvgYarnPoolAllocatedMemoryDuringContention** | **float64** | Average allocated memory with pending containers. | [optional] [default to null]
**AvgYarnPoolFairShareMemoryDuringContention** | **float64** | Average fair share memory with pending containers. | [optional] [default to null]
**AvgYarnPoolSteadyFairShareMemoryDuringContention** | **float64** | Average steady fair share memory with pending containers. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

