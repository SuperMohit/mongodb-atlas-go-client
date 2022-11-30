# \OnlineArchiveApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneOnlineArchive**](OnlineArchiveApi.md#CreateOneOnlineArchive) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/onlineArchives | Create One Online Archive
[**DownloadOnlineArchiveQueryLogs**](OnlineArchiveApi.md#DownloadOnlineArchiveQueryLogs) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/onlineArchives/queryLogs.gz | Download Online Archive Query Logs
[**RemoveOneOnlineArchive**](OnlineArchiveApi.md#RemoveOneOnlineArchive) | **Delete** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Remove One Online Archive
[**ReturnAllOnlineArchivesForOneCluster**](OnlineArchiveApi.md#ReturnAllOnlineArchivesForOneCluster) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/onlineArchives | Return All Online Archives for One Cluster
[**ReturnOneOnlineArchive**](OnlineArchiveApi.md#ReturnOneOnlineArchive) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Return One Online Archive
[**UpdateOneOnlineArchive**](OnlineArchiveApi.md#UpdateOneOnlineArchive) | **Patch** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Update One Online Archive



## CreateOneOnlineArchive

> ApiAtlasOnlineArchiveView CreateOneOnlineArchive(ctx, groupId, clusterName).ApiAtlasOnlineArchiveView(apiAtlasOnlineArchiveView).Envelope(envelope).Pretty(pretty).Execute()

Create One Online Archive



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster that contains the collection for which you want to create one online archive.
    apiAtlasOnlineArchiveView := *openapiclient.NewApiAtlasOnlineArchiveView(*openapiclient.NewCriteriaView()) // ApiAtlasOnlineArchiveView | Creates one online archive.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnlineArchiveApi.CreateOneOnlineArchive(context.Background(), groupId, clusterName).ApiAtlasOnlineArchiveView(apiAtlasOnlineArchiveView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.CreateOneOnlineArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneOnlineArchive`: ApiAtlasOnlineArchiveView
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.CreateOneOnlineArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection for which you want to create one online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneOnlineArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasOnlineArchiveView** | [**ApiAtlasOnlineArchiveView**](ApiAtlasOnlineArchiveView.md) | Creates one online archive. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasOnlineArchiveView**](ApiAtlasOnlineArchiveView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadOnlineArchiveQueryLogs

> *os.File DownloadOnlineArchiveQueryLogs(ctx, groupId, clusterName).StartDate(startDate).EndDate(endDate).ArchiveOnly(archiveOnly).Execute()

Download Online Archive Query Logs



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster that contains the collection for which you want to return the query logs from one online archive.
    startDate := int64(1636481348) // int64 | Date and time that specifies the starting point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time). (optional)
    endDate := int64(1636481348) // int64 | Date and time that specifies the end point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time). (optional)
    archiveOnly := true // bool | Flag that indicates whether to download logs for queries against your online archive only or both your online archive and cluster. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnlineArchiveApi.DownloadOnlineArchiveQueryLogs(context.Background(), groupId, clusterName).StartDate(startDate).EndDate(endDate).ArchiveOnly(archiveOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.DownloadOnlineArchiveQueryLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadOnlineArchiveQueryLogs`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.DownloadOnlineArchiveQueryLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection for which you want to return the query logs from one online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOnlineArchiveQueryLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **int64** | Date and time that specifies the starting point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time). | 
 **endDate** | **int64** | Date and time that specifies the end point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time). | 
 **archiveOnly** | **bool** | Flag that indicates whether to download logs for queries against your online archive only or both your online archive and cluster. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/gzip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneOnlineArchive

> RemoveOneOnlineArchive(ctx, groupId, archiveId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Remove One Online Archive



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
    archiveId := "archiveId_example" // string | Unique 24-hexadecimal digit string that identifies the online archive to delete.
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnlineArchiveApi.RemoveOneOnlineArchive(context.Background(), groupId, archiveId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.RemoveOneOnlineArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**archiveId** | **string** | Unique 24-hexadecimal digit string that identifies the online archive to delete. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneOnlineArchiveRequest struct via the builder pattern


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


## ReturnAllOnlineArchivesForOneCluster

> PaginatedOnlineArchiveView ReturnAllOnlineArchivesForOneCluster(ctx, groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Online Archives for One Cluster



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster that contains the collection for which you want to return the online archives.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnlineArchiveApi.ReturnAllOnlineArchivesForOneCluster(context.Background(), groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.ReturnAllOnlineArchivesForOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllOnlineArchivesForOneCluster`: PaginatedOnlineArchiveView
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.ReturnAllOnlineArchivesForOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the collection for which you want to return the online archives. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllOnlineArchivesForOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedOnlineArchiveView**](PaginatedOnlineArchiveView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneOnlineArchive

> ApiAtlasOnlineArchiveView ReturnOneOnlineArchive(ctx, groupId, archiveId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return One Online Archive



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
    archiveId := "archiveId_example" // string | Unique 24-hexadecimal digit string that identifies the online archive to return.
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnlineArchiveApi.ReturnOneOnlineArchive(context.Background(), groupId, archiveId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.ReturnOneOnlineArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneOnlineArchive`: ApiAtlasOnlineArchiveView
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.ReturnOneOnlineArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**archiveId** | **string** | Unique 24-hexadecimal digit string that identifies the online archive to return. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneOnlineArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasOnlineArchiveView**](ApiAtlasOnlineArchiveView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneOnlineArchive

> ApiAtlasOnlineArchiveView UpdateOneOnlineArchive(ctx, groupId, archiveId, clusterName).ApiAtlasOnlineArchiveView(apiAtlasOnlineArchiveView).Envelope(envelope).Pretty(pretty).Execute()

Update One Online Archive



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
    archiveId := "archiveId_example" // string | Unique 24-hexadecimal digit string that identifies the online archive to update.
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.
    apiAtlasOnlineArchiveView := *openapiclient.NewApiAtlasOnlineArchiveView(*openapiclient.NewCriteriaView()) // ApiAtlasOnlineArchiveView | Updates, pauses, or resumes one online archive.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnlineArchiveApi.UpdateOneOnlineArchive(context.Background(), groupId, archiveId, clusterName).ApiAtlasOnlineArchiveView(apiAtlasOnlineArchiveView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnlineArchiveApi.UpdateOneOnlineArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneOnlineArchive`: ApiAtlasOnlineArchiveView
    fmt.Fprintf(os.Stdout, "Response from `OnlineArchiveApi.UpdateOneOnlineArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**archiveId** | **string** | Unique 24-hexadecimal digit string that identifies the online archive to update. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneOnlineArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiAtlasOnlineArchiveView** | [**ApiAtlasOnlineArchiveView**](ApiAtlasOnlineArchiveView.md) | Updates, pauses, or resumes one online archive. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasOnlineArchiveView**](ApiAtlasOnlineArchiveView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

