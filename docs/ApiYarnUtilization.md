# ApiYarnUtilization

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgCpuUtilization** | **float64** | Average number of VCores used by YARN applications during the report window. | [optional] [default to null]
**MaxCpuUtilization** | **float64** | Maximum number of VCores used by YARN applications during the report window. | [optional] [default to null]
**AvgCpuDailyPeak** | **float64** | Average daily peak VCores used by YARN applications during the report window. The number is computed by first finding the maximum resource consumption per day and then taking their mean. | [optional] [default to null]
**MaxCpuUtilizationTimestampMs** | **float64** | Timestamp corresponds to maximum number of VCores used by YARN applications during the report window. | [optional] [default to null]
**AvgCpuUtilizationPercentage** | **float64** | Average percentage of VCores used by YARN applications during the report window. | [optional] [default to null]
**MaxCpuUtilizationPercentage** | **float64** | Maximum percentage of VCores used by YARN applications during the report window. | [optional] [default to null]
**AvgCpuDailyPeakPercentage** | **float64** | Average daily peak percentage of VCores used by YARN applications during the report window. | [optional] [default to null]
**AvgMemoryUtilization** | **float64** | Average memory used by YARN applications during the report window. | [optional] [default to null]
**MaxMemoryUtilization** | **float64** | Maximum memory used by YARN applications during the report window. | [optional] [default to null]
**AvgMemoryDailyPeak** | **float64** | Average daily peak memory used by YARN applications during the report window. The number is computed by first finding the maximum resource consumption per day and then taking their mean. | [optional] [default to null]
**MaxMemoryUtilizationTimestampMs** | **float64** | Timestamp corresponds to maximum memory used by YARN applications during the report window. | [optional] [default to null]
**AvgMemoryUtilizationPercentage** | **float64** | Average percentage memory used by YARN applications during the report window. | [optional] [default to null]
**MaxMemoryUtilizationPercentage** | **float64** | Maximum percentage of memory used by YARN applications during the report window. | [optional] [default to null]
**AvgMemoryDailyPeakPercentage** | **float64** | Average daily peak percentage of memory used by YARN applications during the report window. | [optional] [default to null]
**TenantUtilizations** | [***ApiYarnTenantUtilizationList**](ApiYarnTenantUtilizationList.md) |  | [optional] [default to null]
**ErrorMessage** | **string** | error message of utilization report. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

