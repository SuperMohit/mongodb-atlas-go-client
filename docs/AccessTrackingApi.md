# \AccessTrackingApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnDatabaseAccessHistoryForOneClusterUsingItsClusterName**](AccessTrackingApi.md#ReturnDatabaseAccessHistoryForOneClusterUsingItsClusterName) | **Get** /api/atlas/v1.0/groups/{groupId}/dbAccessHistory/clusters/{clusterName} | Return Database Access History for One Cluster using Its Cluster Name
[**ReturnDatabaseAccessHistoryForOneClusterUsingItsHostname**](AccessTrackingApi.md#ReturnDatabaseAccessHistoryForOneClusterUsingItsHostname) | **Get** /api/atlas/v1.0/groups/{groupId}/dbAccessHistory/processes/{hostname} | Return Database Access History for One Cluster using Its Hostname



## ReturnDatabaseAccessHistoryForOneClusterUsingItsClusterName

> ApiMongoDBAccessLogsListView ReturnDatabaseAccessHistoryForOneClusterUsingItsClusterName(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).AuthResult(authResult).End(end).IpAddress(ipAddress).NLogs(nLogs).Start(start).Execute()

Return Database Access History for One Cluster using Its Cluster Name



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    authResult := true // bool | Flag that indicates whether the response returns the successful authentication attempts only. (optional)
    end := "end_example" // string | Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses the ISO 8601 timestamp format in UTC. (optional)
    ipAddress := "ipAddress_example" // string | One Internet Protocol address that attempted to authenticate with the database. (optional)
    nLogs := int64(789) // int64 | Maximum number of lines from the log to return. (optional) (default to 20000)
    start := time.Now() // time.Time | Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses the ISO 8601 timestamp format in UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTrackingApi.ReturnDatabaseAccessHistoryForOneClusterUsingItsClusterName(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).AuthResult(authResult).End(end).IpAddress(ipAddress).NLogs(nLogs).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTrackingApi.ReturnDatabaseAccessHistoryForOneClusterUsingItsClusterName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnDatabaseAccessHistoryForOneClusterUsingItsClusterName`: ApiMongoDBAccessLogsListView
    fmt.Fprintf(os.Stdout, "Response from `AccessTrackingApi.ReturnDatabaseAccessHistoryForOneClusterUsingItsClusterName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnDatabaseAccessHistoryForOneClusterUsingItsClusterNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **authResult** | **bool** | Flag that indicates whether the response returns the successful authentication attempts only. | 
 **end** | **string** | Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses the ISO 8601 timestamp format in UTC. | 
 **ipAddress** | **string** | One Internet Protocol address that attempted to authenticate with the database. | 
 **nLogs** | **int64** | Maximum number of lines from the log to return. | [default to 20000]
 **start** | **time.Time** | Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses the ISO 8601 timestamp format in UTC. | 

### Return type

[**ApiMongoDBAccessLogsListView**](ApiMongoDBAccessLogsListView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnDatabaseAccessHistoryForOneClusterUsingItsHostname

> ApiMongoDBAccessLogsListView ReturnDatabaseAccessHistoryForOneClusterUsingItsHostname(ctx, groupId, hostname).Envelope(envelope).Pretty(pretty).AuthResult(authResult).End(end).IpAddress(ipAddress).NLogs(nLogs).Start(start).Execute()

Return Database Access History for One Cluster using Its Hostname



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    hostname := "hostname_example" // string | Fully qualified domain name or IP address of the MongoDB host that stores the log files that you want to download.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    authResult := true // bool | Flag that indicates whether the response returns the successful authentication attempts only. (optional)
    end := time.Now() // time.Time | Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses the ISO 8601 timestamp format in UTC. (optional)
    ipAddress := "ipAddress_example" // string | One Internet Protocol address that attempted to authenticate with the database. (optional)
    nLogs := int32(56) // int32 | Maximum number of lines from the log to return. (optional) (default to 20000)
    start := time.Now() // time.Time | Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses the ISO 8601 timestamp format in UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTrackingApi.ReturnDatabaseAccessHistoryForOneClusterUsingItsHostname(context.Background(), groupId, hostname).Envelope(envelope).Pretty(pretty).AuthResult(authResult).End(end).IpAddress(ipAddress).NLogs(nLogs).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTrackingApi.ReturnDatabaseAccessHistoryForOneClusterUsingItsHostname``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnDatabaseAccessHistoryForOneClusterUsingItsHostname`: ApiMongoDBAccessLogsListView
    fmt.Fprintf(os.Stdout, "Response from `AccessTrackingApi.ReturnDatabaseAccessHistoryForOneClusterUsingItsHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**hostname** | **string** | Fully qualified domain name or IP address of the MongoDB host that stores the log files that you want to download. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnDatabaseAccessHistoryForOneClusterUsingItsHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **authResult** | **bool** | Flag that indicates whether the response returns the successful authentication attempts only. | 
 **end** | **time.Time** | Date and time when to stop retrieving database history. If you specify **end**, you must also specify **start**. This parameter uses the ISO 8601 timestamp format in UTC. | 
 **ipAddress** | **string** | One Internet Protocol address that attempted to authenticate with the database. | 
 **nLogs** | **int32** | Maximum number of lines from the log to return. | [default to 20000]
 **start** | **time.Time** | Date and time when MongoDB Cloud begins retrieving database history. If you specify **start**, you must also specify **end**. This parameter uses the ISO 8601 timestamp format in UTC. | 

### Return type

[**ApiMongoDBAccessLogsListView**](ApiMongoDBAccessLogsListView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

