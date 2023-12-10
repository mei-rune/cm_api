# ApiYarnApplication

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedMB** | **float64** | The sum of memory in MB allocated to the application&#x27;s running containers Available since v12. | [optional] [default to null]
**AllocatedVCores** | **float64** | The sum of virtual cores allocated to the application&#x27;s running containers Available since v12. | [optional] [default to null]
**RunningContainers** | **float64** | The number of containers currently running for the application Available since v12. | [optional] [default to null]
**ApplicationTags** | **[]string** | List of YARN application tags. Available since v12. | [optional] [default to null]
**AllocatedMemorySeconds** | **float64** | Allocated memory to the application in units of mb-secs. Available since v12. | [optional] [default to null]
**AllocatedVcoreSeconds** | **float64** | Allocated vcore-secs to the application. Available since v12. | [optional] [default to null]
**ApplicationId** | **string** | The application id. | [optional] [default to null]
**Name** | **string** | The name of the application. | [optional] [default to null]
**StartTime** | **string** | The time the application was submitted. | [optional] [default to null]
**EndTime** | **string** | The time the application finished. If the application hasn&#x27;t finished this will return null. | [optional] [default to null]
**User** | **string** | The user who submitted the application. | [optional] [default to null]
**Pool** | **string** | The pool the application was submitted to. | [optional] [default to null]
**Progress** | **float64** | The progress, as a percentage, the application has made. This is only set if the application is currently executing. | [optional] [default to null]
**Attributes** | **map[string]string** | A map of additional application attributes which is generated by Cloudera Manager. For example MR2 job counters are exposed as key/value pairs here. For more details see the Cloudera Manager documentation. | [optional] [default to null]
**Mr2AppInformation** | [***ApiMr2AppInformation**](ApiMr2AppInformation.md) |  | [optional] [default to null]
**State** | **string** |  | [optional] [default to null]
**ContainerUsedMemorySeconds** | **float64** | Actual memory (in MB-secs) used by containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics. Available since v12. | [optional] [default to null]
**ContainerUsedMemoryMax** | **float64** | Maximum memory used by containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics Available since v16 | [optional] [default to null]
**ContainerUsedCpuSeconds** | **float64** | Actual CPU (in percent-secs) used by containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics. Available since v12. | [optional] [default to null]
**ContainerUsedVcoreSeconds** | **float64** | Actual VCore-secs used by containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics. Available since v12. | [optional] [default to null]
**ContainerAllocatedMemorySeconds** | **float64** | Total memory (in mb-secs) allocated to containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics. Available since v12. | [optional] [default to null]
**ContainerAllocatedVcoreSeconds** | **float64** | Total vcore-secs allocated to containers launched by the YARN application. Computed by running a MapReduce job from Cloudera Service Monitor to aggregate YARN usage metrics. Available since v12. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
