# ApiCommand

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** | The command ID. | [optional] [default to null]
**Name** | **string** | The command name. | [optional] [default to null]
**StartTime** | **string** | The start time. | [optional] [default to null]
**EndTime** | **string** | The end time, if the command is finished. | [optional] [default to null]
**Active** | **bool** | Whether the command is currently active. | [optional] [default to null]
**Success** | **bool** | If the command is finished, whether it was successful. | [optional] [default to null]
**ResultMessage** | **string** | If the command is finished, the result message. | [optional] [default to null]
**ResultDataUrl** | **string** | URL to the command&#x27;s downloadable result data, if any exists. | [optional] [default to null]
**ClusterRef** | [***ApiClusterRef**](ApiClusterRef.md) |  | [optional] [default to null]
**ServiceRef** | [***ApiServiceRef**](ApiServiceRef.md) |  | [optional] [default to null]
**RoleRef** | [***ApiRoleRef**](ApiRoleRef.md) |  | [optional] [default to null]
**HostRef** | [***ApiHostRef**](ApiHostRef.md) |  | [optional] [default to null]
**Parent** | [***ApiCommand**](ApiCommand.md) |  | [optional] [default to null]
**Children** | [***ApiCommandList**](ApiCommandList.md) |  | [optional] [default to null]
**CanRetry** | **bool** | If the command can be retried. Available since V11 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

