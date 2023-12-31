/*
 * Cloudera Manager API
 *
 * <h1>Cloudera Manager API v32</h1>       <p>Introduced in Cloudera Manager 6.2.0</p>       <p><a href=\"http://www.cloudera.com/documentation.html\">Cloudera Product Documentation</a></p>
 *
 * API version: 6.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cm_api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type NameservicesResourceApiService service

/*
NameservicesResourceApiService Fetch metric readings for a particular nameservice.
Fetch metric readings for a particular nameservice. &lt;p&gt; By default, this call will look up all metrics available. If only specific metrics are desired, use the &lt;i&gt;metrics&lt;/i&gt; parameter. &lt;p&gt; By default, the returned results correspond to a 5 minute window based on the provided end time (which defaults to the current server time). The &lt;i&gt;from&lt;/i&gt; and &lt;i&gt;to&lt;/i&gt; parameters can be used to control the window being queried. A maximum window of 3 hours is enforced. &lt;p&gt; When requesting a \&quot;full\&quot; view, aside from the extended properties of the returned metric data, the collection will also contain information about all metrics available, even if no readings are available in the requested window.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterName
 * @param nameservice The nameservice.
 * @param serviceName The service name.
 * @param optional nil or *NameservicesResourceApiGetMetricsOpts - Optional Parameters:
     * @param "From" (optional.String) -  Start of the period to query.
     * @param "Metrics" (optional.Interface of []string) -  Filter for which metrics to query.
     * @param "To" (optional.String) -  End of the period to query.
     * @param "View" (optional.String) -  The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;.
@return ApiMetricList
*/

type NameservicesResourceApiGetMetricsOpts struct {
	From    optional.String
	Metrics optional.Interface
	To      optional.String
	View    optional.String
}

func (a *NameservicesResourceApiService) GetMetrics(ctx context.Context, clusterName string, serviceName string, nameservice string, localVarOptionals *NameservicesResourceApiGetMetricsOpts) (ApiMetricList, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ApiMetricList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{clusterName}/services/{serviceName}/nameservices/{nameservice}/metrics"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nameservice"+"}", fmt.Sprintf("%v", nameservice), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceName"+"}", fmt.Sprintf("%v", serviceName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.From.IsSet() {
		localVarQueryParams.Add("from", parameterToString(localVarOptionals.From.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Metrics.IsSet() {
		localVarQueryParams.Add("metrics", parameterToString(localVarOptionals.Metrics.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.To.IsSet() {
		localVarQueryParams.Add("to", parameterToString(localVarOptionals.To.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.View.IsSet() {
		localVarQueryParams.Add("view", parameterToString(localVarOptionals.View.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ApiMetricList
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
NameservicesResourceApiService List the nameservices of an HDFS service.
List the nameservices of an HDFS service.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterName
 * @param serviceName The service name.
 * @param optional nil or *NameservicesResourceApiListNameservicesOpts - Optional Parameters:
     * @param "View" (optional.String) -  The view of the data to materialize, either \&quot;summary\&quot; or \&quot;full\&quot;.
@return ApiNameserviceList
*/

type NameservicesResourceApiListNameservicesOpts struct {
	View optional.String
}

func (a *NameservicesResourceApiService) ListNameservices(ctx context.Context, clusterName string, serviceName string, localVarOptionals *NameservicesResourceApiListNameservicesOpts) (ApiNameserviceList, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ApiNameserviceList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{clusterName}/services/{serviceName}/nameservices"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceName"+"}", fmt.Sprintf("%v", serviceName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.View.IsSet() {
		localVarQueryParams.Add("view", parameterToString(localVarOptionals.View.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ApiNameserviceList
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
NameservicesResourceApiService Retrieve information about a nameservice.
Retrieve information about a nameservice.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterName
 * @param nameservice The nameservice to retrieve.
 * @param serviceName The service name.
 * @param optional nil or *NameservicesResourceApiReadNameserviceOpts - Optional Parameters:
     * @param "View" (optional.String) -  The view to materialize. Defaults to &#x27;full&#x27;.
@return ApiNameservice
*/

type NameservicesResourceApiReadNameserviceOpts struct {
	View optional.String
}

func (a *NameservicesResourceApiService) ReadNameservice(ctx context.Context, clusterName string, serviceName string, nameservice string, localVarOptionals *NameservicesResourceApiReadNameserviceOpts) (ApiNameservice, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ApiNameservice
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{clusterName}/services/{serviceName}/nameservices/{nameservice}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", fmt.Sprintf("%v", clusterName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nameservice"+"}", fmt.Sprintf("%v", nameservice), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceName"+"}", fmt.Sprintf("%v", serviceName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.View.IsSet() {
		localVarQueryParams.Add("view", parameterToString(localVarOptionals.View.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ApiNameservice
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
