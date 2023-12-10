# ApiMrUsageReportRow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimePeriod** | **string** | The time period over which this report is generated. | [optional] [default to null]
**User** | **string** | The user being reported. | [optional] [default to null]
**Group** | **string** | The group this user belongs to. | [optional] [default to null]
**CpuSec** | **float64** | Amount of CPU time (in seconds) taken up this user&#x27;s MapReduce jobs. | [optional] [default to null]
**MemoryBytes** | **float64** | The sum of physical memory used (collected as a snapshot) by this user&#x27;s MapReduce jobs. | [optional] [default to null]
**JobCount** | **float64** | Number of jobs. | [optional] [default to null]
**TaskCount** | **float64** | Number of tasks. | [optional] [default to null]
**DurationSec** | **float64** | Total duration of this user&#x27;s MapReduce jobs. | [optional] [default to null]
**FailedMaps** | **float64** | Failed maps of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]
**TotalMaps** | **float64** | Total maps of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]
**FailedReduces** | **float64** | Failed reduces of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]
**TotalReduces** | **float64** | Total reduces of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]
**MapInputBytes** | **float64** | Map input bytes of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]
**MapOutputBytes** | **float64** | Map output bytes of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]
**HdfsBytesRead** | **float64** | HDFS bytes read of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]
**HdfsBytesWritten** | **float64** | HDFS bytes written of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]
**LocalBytesRead** | **float64** | Local bytes read of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]
**LocalBytesWritten** | **float64** | Local bytes written of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]
**DataLocalMaps** | **float64** | Data local maps of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]
**RackLocalMaps** | **float64** | Rack local maps of this user&#x27;s MapReduce jobs. Available since v11. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

