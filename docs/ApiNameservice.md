# ApiNameservice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the nameservice. | [optional] [default to null]
**Active** | [***ApiRoleRef**](ApiRoleRef.md) |  | [optional] [default to null]
**ActiveFailoverController** | [***ApiRoleRef**](ApiRoleRef.md) |  | [optional] [default to null]
**StandBy** | [***ApiRoleRef**](ApiRoleRef.md) |  | [optional] [default to null]
**StandByFailoverController** | [***ApiRoleRef**](ApiRoleRef.md) |  | [optional] [default to null]
**Secondary** | [***ApiRoleRef**](ApiRoleRef.md) |  | [optional] [default to null]
**MountPoints** | **[]string** | Mount points assigned to this nameservice in a federation. | [optional] [default to null]
**HealthSummary** | [***ApiHealthSummary**](ApiHealthSummary.md) |  | [optional] [default to null]
**HealthChecks** | [**[]ApiHealthCheck**](ApiHealthCheck.md) | Requires \&quot;full\&quot; view. List of health checks performed on the nameservice. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

