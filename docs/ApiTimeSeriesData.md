# ApiTimeSeriesData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | The timestamp for this time series data point. Note that the timestamp reflects coordinated universal time (UTC) and not necessarily the server&#x27;s time zone. The rest API formats the UTC timestamp as an ISO-8061 string. | [optional] [default to null]
**Value** | **float64** | The value of the time series data. | [optional] [default to null]
**Type_** | **string** | The type of the time series data. | [optional] [default to null]
**AggregateStatistics** | [***ApiTimeSeriesAggregateStatistics**](ApiTimeSeriesAggregateStatistics.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

