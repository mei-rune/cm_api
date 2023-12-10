# {{classname}}

All URIs are relative to */api/v32*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelImpalaQuery**](ImpalaQueriesResourceApi.md#CancelImpalaQuery) | **Post** /clusters/{clusterName}/services/{serviceName}/impalaQueries/{queryId}/cancel | Cancels an Impala Query.
[**GetImpalaQueries**](ImpalaQueriesResourceApi.md#GetImpalaQueries) | **Get** /clusters/{clusterName}/services/{serviceName}/impalaQueries | Returns a list of queries that satisfy the filter.
[**GetImpalaQueryAttributes**](ImpalaQueriesResourceApi.md#GetImpalaQueryAttributes) | **Get** /clusters/{clusterName}/services/{serviceName}/impalaQueries/attributes | Returns the list of all attributes that the Service Monitor can associate with Impala queries.
[**GetQueryDetails**](ImpalaQueriesResourceApi.md#GetQueryDetails) | **Get** /clusters/{clusterName}/services/{serviceName}/impalaQueries/{queryId} | Returns details about the query.

# **CancelImpalaQuery**
> ApiImpalaCancelResponse CancelImpalaQuery(ctx, clusterName, queryId, serviceName)
Cancels an Impala Query.

Cancels an Impala Query. <p> Available since API v4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **queryId** | **string**| The queryId to cancel | 
  **serviceName** | **string**| The name of the service | 

### Return type

[**ApiImpalaCancelResponse**](ApiImpalaCancelResponse.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImpalaQueries**
> ApiImpalaQueryResponse GetImpalaQueries(ctx, clusterName, serviceName, optional)
Returns a list of queries that satisfy the filter.

Returns a list of queries that satisfy the filter <p> Available since API v4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**| The name of the service | 
 **optional** | ***ImpalaQueriesResourceApiGetImpalaQueriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImpalaQueriesResourceApiGetImpalaQueriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| A filter to apply to the queries. A basic filter tests the value of an attribute and looks something like &#x27;rowsFetched &#x3D; 1&#x27; or &#x27;user &#x3D; root&#x27;. Multiple basic filters can be combined into a complex expression using standard and / or boolean logic and parenthesis. An example of a complex filter is: &#x27;query_duration &gt; 5s and (user &#x3D; root or user &#x3D; myUserName)&#x27;. An example of expected full query string in requested URL is: &#x27;?filter&#x3D;(query_duration &gt; 5s and (user &#x3D; root or user &#x3D; myUserName))&#x27;. | 
 **from** | **optional.String**| Start of the period to query in ISO 8601 format (defaults to 5 minutes before the &#x27;to&#x27; time). | 
 **limit** | **optional.Int32**| The maximum number of queries to return. Queries will be returned in the following order: &lt;ul&gt; &lt;li&gt; All executing queries, ordered from longest to shortest running &lt;/li&gt; &lt;li&gt; All completed queries order by end time descending. &lt;/li&gt; &lt;/ul&gt; | [default to 100]
 **offset** | **optional.Int32**| The offset to start returning queries from. This is useful for paging through lists of queries. Note that this has non-deterministic behavior if executing queries are included in the response because they can disappear from the list while paging. To exclude executing queries from the response and a &#x27;executing &#x3D; false&#x27; clause to your filter. | [default to 0]
 **to** | **optional.String**| End of the period to query in ISO 8601 format (defaults to current time). | [default to now]

### Return type

[**ApiImpalaQueryResponse**](ApiImpalaQueryResponse.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImpalaQueryAttributes**
> ApiImpalaQueryAttributeList GetImpalaQueryAttributes(ctx, clusterName, serviceName)
Returns the list of all attributes that the Service Monitor can associate with Impala queries.

Returns the list of all attributes that the Service Monitor can associate with Impala queries. <p> Examples of attributes include the user who issued the query and the number of HDFS bytes read by the query. <p> These attributes can be used to search for specific Impala queries through the getImpalaQueries API. For example the 'user' attribute could be used in the search 'user = root'. If the attribute is numeric it can also be used as a metric in a tsquery (ie, 'select hdfs_bytes_read from IMPALA_QUERIES'). <p> Note that this response is identical for all Impala services. <p> Available since API v6.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **serviceName** | **string**|  | 

### Return type

[**ApiImpalaQueryAttributeList**](ApiImpalaQueryAttributeList.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQueryDetails**
> ApiImpalaQueryDetailsResponse GetQueryDetails(ctx, clusterName, queryId, serviceName, optional)
Returns details about the query.

Returns details about the query. Not all queries have details, check the detailsAvailable field from the getQueries response. <p> Available since API v4.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**|  | 
  **queryId** | **string**| The queryId to get information about | 
  **serviceName** | **string**|  | 
 **optional** | ***ImpalaQueriesResourceApiGetQueryDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImpalaQueriesResourceApiGetQueryDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **format** | **optional.String**| There are two valid format parameters: &lt;ul&gt; &lt;li&gt; &#x27;text&#x27;: this is a text based, human readable representation of the Impala runtime profile. &lt;/li&gt; &lt;li&gt; &#x27;thrift_encoded&#x27;: this a compact-thrift, base64 encoded representation of the Impala RuntimeProfile.thrift object. See the Impala documentation for more details. &lt;/li&gt; &lt;/ul&gt; | [default to text]

### Return type

[**ApiImpalaQueryDetailsResponse**](ApiImpalaQueryDetailsResponse.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

