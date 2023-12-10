# ApiPerfInspectorPingArgs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PingTimeoutSecs** | **float64** | Timeout in seconds for the ping request to each target host. If not specified, defaults to 10 seconds. Must be a value between 1 and 3600 seconds, inclusive. | [optional] [default to null]
**PingCount** | **float64** | Number of iterations of the ping request to each target host. If not specified, defaults to 10 count. | [optional] [default to null]
**PingPacketSizeBytes** | **float64** | Packet size in bytes for each ping request. If not specified, defaults to 56 bytes. Must be a value between 1 and 65507 bytes, inclusive. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

