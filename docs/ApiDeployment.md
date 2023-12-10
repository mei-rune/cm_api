# ApiDeployment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | Readonly. This timestamp is provided when you request a deployment and is not required (or even read) when creating a deployment. This timestamp is useful if you have multiple deployments saved and want to determine which one to use as a restore point. | [optional] [default to null]
**Clusters** | [**[]ApiCluster**](ApiCluster.md) | List of clusters in the system including their services, roles and complete config values. | [optional] [default to null]
**Hosts** | [**[]ApiHost**](ApiHost.md) | List of hosts in the system | [optional] [default to null]
**Users** | [**[]ApiUser**](ApiUser.md) | List of all users in the system | [optional] [default to null]
**VersionInfo** | [***ApiVersionInfo**](ApiVersionInfo.md) |  | [optional] [default to null]
**ManagementService** | [***ApiService**](ApiService.md) |  | [optional] [default to null]
**ManagerSettings** | [***ApiConfigList**](ApiConfigList.md) |  | [optional] [default to null]
**AllHostsConfig** | [***ApiConfigList**](ApiConfigList.md) |  | [optional] [default to null]
**Peers** | [**[]ApiCmPeer**](ApiCmPeer.md) | The list of peers configured in Cloudera Manager. Available since API v3. | [optional] [default to null]
**HostTemplates** | [***ApiHostTemplateList**](ApiHostTemplateList.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

