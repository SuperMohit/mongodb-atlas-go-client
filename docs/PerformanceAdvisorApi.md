# \PerformanceAdvisorApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableManagedSlowOperationThreshold**](PerformanceAdvisorApi.md#DisableManagedSlowOperationThreshold) | **Delete** /api/atlas/v1.0/groups/{groupId}/managedSlowMs/disable | Disable Managed Slow Operation Threshold
[**EnableManagedSlowOperationThreshold**](PerformanceAdvisorApi.md#EnableManagedSlowOperationThreshold) | **Post** /api/atlas/v1.0/groups/{groupId}/managedSlowMs/enable | Enable Managed Slow Operation Threshold
[**ReturnAllNamespacesForOneHost**](PerformanceAdvisorApi.md#ReturnAllNamespacesForOneHost) | **Get** /api/public/v1.0/groups/{groupId}/hosts/{hostId}/performanceAdvisor/namespaces | Return All Namespaces for One Host
[**ReturnAllNamespacesForOneHost1**](PerformanceAdvisorApi.md#ReturnAllNamespacesForOneHost1) | **Get** /api/atlas/v1.0/groups/{groupId}/processes/{processId}/performanceAdvisor/namespaces | Return All Namespaces for One Host
[**ReturnSlowQueries**](PerformanceAdvisorApi.md#ReturnSlowQueries) | **Get** /api/public/v1.0/groups/{groupId}/hosts/{hostId}/performanceAdvisor/slowQueryLogs | Return Slow Queries
[**ReturnSlowQueries1**](PerformanceAdvisorApi.md#ReturnSlowQueries1) | **Get** /api/atlas/v1.0/groups/{groupId}/processes/{processId}/performanceAdvisor/slowQueryLogs | Return Slow Queries
[**ReturnSuggestedIndexes**](PerformanceAdvisorApi.md#ReturnSuggestedIndexes) | **Get** /api/public/v1.0/groups/{groupId}/hosts/{hostId}/performanceAdvisor/suggestedIndexes | Return Suggested Indexes
[**ReturnSuggestedIndexes1**](PerformanceAdvisorApi.md#ReturnSuggestedIndexes1) | **Get** /api/atlas/v1.0/groups/{groupId}/processes/{processId}/performanceAdvisor/suggestedIndexes | Return Suggested Indexes



## DisableManagedSlowOperationThreshold

> DisableManagedSlowOperationThreshold(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Disable Managed Slow Operation Threshold



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PerformanceAdvisorApi.DisableManagedSlowOperationThreshold(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.DisableManagedSlowOperationThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableManagedSlowOperationThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

 (empty response body)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableManagedSlowOperationThreshold

> EnableManagedSlowOperationThreshold(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Enable Managed Slow Operation Threshold



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PerformanceAdvisorApi.EnableManagedSlowOperationThreshold(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.EnableManagedSlowOperationThreshold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableManagedSlowOperationThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

 (empty response body)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllNamespacesForOneHost

> []ApiNamespacesView ReturnAllNamespacesForOneHost(ctx, groupId, hostId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Duration(duration).Since(since).Execute()

Return All Namespaces for One Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    hostId := "hostId_example" // string | Unique identifier of the host for the MongoDB process.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    duration := int64(789) // int64 | Length of time expressed in milliseconds during which the query finds suggested indexes among the managed namespaces in the cluster. (optional)
    since := int64(789) // int64 | Timestamp in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time) from which the query retrieves the suggested indexes. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PerformanceAdvisorApi.ReturnAllNamespacesForOneHost(context.Background(), groupId, hostId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Duration(duration).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ReturnAllNamespacesForOneHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllNamespacesForOneHost`: []ApiNamespacesView
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ReturnAllNamespacesForOneHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**hostId** | **string** | Unique identifier of the host for the MongoDB process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllNamespacesForOneHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **duration** | **int64** | Length of time expressed in milliseconds during which the query finds suggested indexes among the managed namespaces in the cluster. | 
 **since** | **int64** | Timestamp in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time) from which the query retrieves the suggested indexes. | 

### Return type

[**[]ApiNamespacesView**](ApiNamespacesView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllNamespacesForOneHost1

> []ApiNamespacesView ReturnAllNamespacesForOneHost1(ctx, groupId, processId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Duration(duration).Since(since).Execute()

Return All Namespaces for One Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    processId := "processId_example" // string | Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    duration := int64(789) // int64 | Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don't specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. (optional)
    since := int64(789) // int64 | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don't specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PerformanceAdvisorApi.ReturnAllNamespacesForOneHost1(context.Background(), groupId, processId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Duration(duration).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ReturnAllNamespacesForOneHost1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllNamespacesForOneHost1`: []ApiNamespacesView
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ReturnAllNamespacesForOneHost1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**processId** | **string** | Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllNamespacesForOneHost1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **duration** | **int64** | Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. | 
 **since** | **int64** | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. | 

### Return type

[**[]ApiNamespacesView**](ApiNamespacesView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnSlowQueries

> []ApiPerformanceAdvisorSlowQueryListView ReturnSlowQueries(ctx, groupId, hostId).Envelope(envelope).Pretty(pretty).Duration(duration).Namespaces(namespaces).NLogs(nLogs).Since(since).Execute()

Return Slow Queries



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    hostId := "hostId_example" // string | Unique identifier of the host for the MongoDB process.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    duration := int64(789) // int64 | Length of time expressed in milliseconds during which the query finds suggested indexes among the managed namespaces in the cluster. (optional)
    namespaces := []string{"Inner_example"} // []string | Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as `.`: `<database>.<collection>`. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (`&`) between each namespace. Omit this parameter to return results for all namespaces. (optional)
    nLogs := int64(789) // int64 | Maximum number of lines from the log to return. (optional) (default to 20000)
    since := int64(789) // int64 | Timestamp in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time) from which the query retrieves the suggested indexes. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PerformanceAdvisorApi.ReturnSlowQueries(context.Background(), groupId, hostId).Envelope(envelope).Pretty(pretty).Duration(duration).Namespaces(namespaces).NLogs(nLogs).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ReturnSlowQueries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnSlowQueries`: []ApiPerformanceAdvisorSlowQueryListView
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ReturnSlowQueries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**hostId** | **string** | Unique identifier of the host for the MongoDB process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnSlowQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **duration** | **int64** | Length of time expressed in milliseconds during which the query finds suggested indexes among the managed namespaces in the cluster. | 
 **namespaces** | **[]string** | Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces. | 
 **nLogs** | **int64** | Maximum number of lines from the log to return. | [default to 20000]
 **since** | **int64** | Timestamp in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time) from which the query retrieves the suggested indexes. | 

### Return type

[**[]ApiPerformanceAdvisorSlowQueryListView**](ApiPerformanceAdvisorSlowQueryListView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnSlowQueries1

> []ApiPerformanceAdvisorSlowQueryListView ReturnSlowQueries1(ctx, groupId, processId).Envelope(envelope).Pretty(pretty).Duration(duration).Namespaces(namespaces).NLogs(nLogs).Since(since).Execute()

Return Slow Queries



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    processId := "processId_example" // string | Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    duration := float32(8.14) // float32 | Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don't specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. (optional)
    namespaces := []string{"Inner_example"} // []string | Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as `.`: `<database>.<collection>`. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (`&`) between each namespace. Omit this parameter to return results for all namespaces. (optional)
    nLogs := int64(789) // int64 | Maximum number of lines from the log to return. (optional) (default to 20000)
    since := float32(8.14) // float32 | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don't specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PerformanceAdvisorApi.ReturnSlowQueries1(context.Background(), groupId, processId).Envelope(envelope).Pretty(pretty).Duration(duration).Namespaces(namespaces).NLogs(nLogs).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ReturnSlowQueries1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnSlowQueries1`: []ApiPerformanceAdvisorSlowQueryListView
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ReturnSlowQueries1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**processId** | **string** | Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnSlowQueries1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **duration** | **float32** | Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. | 
 **namespaces** | **[]string** | Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces. | 
 **nLogs** | **int64** | Maximum number of lines from the log to return. | [default to 20000]
 **since** | **float32** | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. | 

### Return type

[**[]ApiPerformanceAdvisorSlowQueryListView**](ApiPerformanceAdvisorSlowQueryListView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnSuggestedIndexes

> ApiPerformanceAdvisorResponseView ReturnSuggestedIndexes(ctx, groupId, hostId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Duration(duration).Namespaces(namespaces).NExamples(nExamples).NIndexes(nIndexes).Since(since).Execute()

Return Suggested Indexes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    hostId := "hostId_example" // string | Unique identifier of the host for the MongoDB process.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    duration := int64(789) // int64 | Length of time expressed in milliseconds during which the query finds suggested indexes among the managed namespaces in the cluster. (optional)
    namespaces := []string{"Inner_example"} // []string | Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as `.`: `<database>.<collection>`. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (`&`) between each namespace. Omit this parameter to return results for all namespaces. (optional)
    nExamples := int64(789) // int64 | Maximum number of example queries that benefit from the suggested index. (optional) (default to 5)
    nIndexes := int64(789) // int64 | Number that indicates the maximum indexes to suggest. (optional)
    since := int64(789) // int64 | Timestamp in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time) from which the query retrieves the suggested indexes. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PerformanceAdvisorApi.ReturnSuggestedIndexes(context.Background(), groupId, hostId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Duration(duration).Namespaces(namespaces).NExamples(nExamples).NIndexes(nIndexes).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ReturnSuggestedIndexes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnSuggestedIndexes`: ApiPerformanceAdvisorResponseView
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ReturnSuggestedIndexes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**hostId** | **string** | Unique identifier of the host for the MongoDB process. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnSuggestedIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **duration** | **int64** | Length of time expressed in milliseconds during which the query finds suggested indexes among the managed namespaces in the cluster. | 
 **namespaces** | **[]string** | Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces. | 
 **nExamples** | **int64** | Maximum number of example queries that benefit from the suggested index. | [default to 5]
 **nIndexes** | **int64** | Number that indicates the maximum indexes to suggest. | 
 **since** | **int64** | Timestamp in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time) from which the query retrieves the suggested indexes. | 

### Return type

[**ApiPerformanceAdvisorResponseView**](ApiPerformanceAdvisorResponseView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnSuggestedIndexes1

> ApiPerformanceAdvisorResponseView ReturnSuggestedIndexes1(ctx, groupId, processId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Duration(duration).Namespaces(namespaces).NExamples(nExamples).NIndexes(nIndexes).Since(since).Execute()

Return Suggested Indexes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    processId := "processId_example" // string | Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    duration := float32(8.14) // float32 | Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don't specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. (optional)
    namespaces := []string{"Inner_example"} // []string | Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as `.`: `<database>.<collection>`. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (`&`) between each namespace. Omit this parameter to return results for all namespaces. (optional)
    nExamples := int64(789) // int64 | Maximum number of example queries that benefit from the suggested index. (optional) (default to 5)
    nIndexes := int64(789) // int64 | Number that indicates the maximum indexes to suggest. (optional)
    since := float32(8.14) // float32 | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don't specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PerformanceAdvisorApi.ReturnSuggestedIndexes1(context.Background(), groupId, processId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Duration(duration).Namespaces(namespaces).NExamples(nExamples).NIndexes(nIndexes).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.ReturnSuggestedIndexes1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnSuggestedIndexes1`: ApiPerformanceAdvisorResponseView
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.ReturnSuggestedIndexes1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**processId** | **string** | Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnSuggestedIndexes1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **duration** | **float32** | Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours. | 
 **namespaces** | **[]string** | Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces. | 
 **nExamples** | **int64** | Maximum number of example queries that benefit from the suggested index. | [default to 5]
 **nIndexes** | **int64** | Number that indicates the maximum indexes to suggest. | 
 **since** | **float32** | Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours. | 

### Return type

[**ApiPerformanceAdvisorResponseView**](ApiPerformanceAdvisorResponseView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

