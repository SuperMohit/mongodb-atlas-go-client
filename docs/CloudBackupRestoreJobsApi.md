# \CloudBackupRestoreJobsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOneRestoreJobOfOneCluster**](CloudBackupRestoreJobsApi.md#CancelOneRestoreJobOfOneCluster) | **Delete** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Cancel One Restore Job of One Cluster
[**RestoreOneSnapshotOfOneCluster**](CloudBackupRestoreJobsApi.md#RestoreOneSnapshotOfOneCluster) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Cluster
[**RestoreOneSnapshotOfOneCluster1**](CloudBackupRestoreJobsApi.md#RestoreOneSnapshotOfOneCluster1) | **Post** /api/atlas/v1.0/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Serverless Instance
[**ReturnAllRestoreJobsForOneCluster**](CloudBackupRestoreJobsApi.md#ReturnAllRestoreJobsForOneCluster) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Cluster
[**ReturnAllRestoreJobsForOneServerlessInstance**](CloudBackupRestoreJobsApi.md#ReturnAllRestoreJobsForOneServerlessInstance) | **Get** /api/atlas/v1.0/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Serverless Instance
[**ReturnOneRestoreJobForOneServerlessInstance**](CloudBackupRestoreJobsApi.md#ReturnOneRestoreJobForOneServerlessInstance) | **Get** /api/atlas/v1.0/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Serverless Instance
[**ReturnOneRestoreJobOfOneCluster**](CloudBackupRestoreJobsApi.md#ReturnOneRestoreJobOfOneCluster) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job of One Cluster



## CancelOneRestoreJobOfOneCluster

> CancelOneRestoreJobOfOneCluster(ctx, groupId, clusterName, restoreJobId).Envelope(envelope).Pretty(pretty).Execute()

Cancel One Restore Job of One Cluster



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster.
    restoreJobId := "restoreJobId_example" // string | Unique 24-hexadecimal digit string that identifies the restore job to remove.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupRestoreJobsApi.CancelOneRestoreJobOfOneCluster(context.Background(), groupId, clusterName, restoreJobId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupRestoreJobsApi.CancelOneRestoreJobOfOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**restoreJobId** | **string** | Unique 24-hexadecimal digit string that identifies the restore job to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOneRestoreJobOfOneClusterRequest struct via the builder pattern


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


## RestoreOneSnapshotOfOneCluster

> ApiAtlasDiskBackupRestoreJobView RestoreOneSnapshotOfOneCluster(ctx, groupId, clusterName).ApiAtlasDiskBackupRestoreJobView(apiAtlasDiskBackupRestoreJobView).Envelope(envelope).Pretty(pretty).Execute()

Restore One Snapshot of One Cluster



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster.
    apiAtlasDiskBackupRestoreJobView := *openapiclient.NewApiAtlasDiskBackupRestoreJobView("DeliveryType_example", []openapiclient.Link{*openapiclient.NewLink()}, "TargetClusterName_example", "ca4d8a18b96b317422974eec") // ApiAtlasDiskBackupRestoreJobView | Restores one snapshot of one cluster from the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupRestoreJobsApi.RestoreOneSnapshotOfOneCluster(context.Background(), groupId, clusterName).ApiAtlasDiskBackupRestoreJobView(apiAtlasDiskBackupRestoreJobView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupRestoreJobsApi.RestoreOneSnapshotOfOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreOneSnapshotOfOneCluster`: ApiAtlasDiskBackupRestoreJobView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupRestoreJobsApi.RestoreOneSnapshotOfOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreOneSnapshotOfOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasDiskBackupRestoreJobView** | [**ApiAtlasDiskBackupRestoreJobView**](ApiAtlasDiskBackupRestoreJobView.md) | Restores one snapshot of one cluster from the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDiskBackupRestoreJobView**](ApiAtlasDiskBackupRestoreJobView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreOneSnapshotOfOneCluster1

> ApiAtlasServerlessBackupRestoreJobViewManual RestoreOneSnapshotOfOneCluster1(ctx, groupId, clusterName).ApiAtlasServerlessBackupRestoreJobViewManual(apiAtlasServerlessBackupRestoreJobViewManual).Envelope(envelope).Pretty(pretty).Execute()

Restore One Snapshot of One Serverless Instance



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the serverless instance whose snapshot you want to restore.
    apiAtlasServerlessBackupRestoreJobViewManual := *openapiclient.NewApiAtlasServerlessBackupRestoreJobViewManual("DeliveryType_example", "TargetClusterName_example", "ca4d8a18b96b317422974eec") // ApiAtlasServerlessBackupRestoreJobViewManual | Restores one snapshot of one serverless instance from the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupRestoreJobsApi.RestoreOneSnapshotOfOneCluster1(context.Background(), groupId, clusterName).ApiAtlasServerlessBackupRestoreJobViewManual(apiAtlasServerlessBackupRestoreJobViewManual).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupRestoreJobsApi.RestoreOneSnapshotOfOneCluster1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreOneSnapshotOfOneCluster1`: ApiAtlasServerlessBackupRestoreJobViewManual
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupRestoreJobsApi.RestoreOneSnapshotOfOneCluster1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the serverless instance whose snapshot you want to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreOneSnapshotOfOneCluster1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasServerlessBackupRestoreJobViewManual** | [**ApiAtlasServerlessBackupRestoreJobViewManual**](ApiAtlasServerlessBackupRestoreJobViewManual.md) | Restores one snapshot of one serverless instance from the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasServerlessBackupRestoreJobViewManual**](ApiAtlasServerlessBackupRestoreJobViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllRestoreJobsForOneCluster

> PaginatedCloudBackupRestoreJobView ReturnAllRestoreJobsForOneCluster(ctx, groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Restore Jobs for One Cluster



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster with the restore jobs you want to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupRestoreJobsApi.ReturnAllRestoreJobsForOneCluster(context.Background(), groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupRestoreJobsApi.ReturnAllRestoreJobsForOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllRestoreJobsForOneCluster`: PaginatedCloudBackupRestoreJobView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupRestoreJobsApi.ReturnAllRestoreJobsForOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the restore jobs you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllRestoreJobsForOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedCloudBackupRestoreJobView**](PaginatedCloudBackupRestoreJobView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllRestoreJobsForOneServerlessInstance

> PaginatedServerlessBackupRestoreJobViewManual ReturnAllRestoreJobsForOneServerlessInstance(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return All Restore Jobs for One Serverless Instance



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the serverless instance.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupRestoreJobsApi.ReturnAllRestoreJobsForOneServerlessInstance(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupRestoreJobsApi.ReturnAllRestoreJobsForOneServerlessInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllRestoreJobsForOneServerlessInstance`: PaginatedServerlessBackupRestoreJobViewManual
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupRestoreJobsApi.ReturnAllRestoreJobsForOneServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the serverless instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllRestoreJobsForOneServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedServerlessBackupRestoreJobViewManual**](PaginatedServerlessBackupRestoreJobViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneRestoreJobForOneServerlessInstance

> ApiAtlasServerlessBackupRestoreJobViewManual ReturnOneRestoreJobForOneServerlessInstance(ctx, groupId, clusterName, restoreJobId).Envelope(envelope).Pretty(pretty).Execute()

Return One Restore Job for One Serverless Instance



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the serverless instance.
    restoreJobId := "restoreJobId_example" // string | Unique 24-hexadecimal digit string that identifies the restore job to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupRestoreJobsApi.ReturnOneRestoreJobForOneServerlessInstance(context.Background(), groupId, clusterName, restoreJobId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupRestoreJobsApi.ReturnOneRestoreJobForOneServerlessInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneRestoreJobForOneServerlessInstance`: ApiAtlasServerlessBackupRestoreJobViewManual
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupRestoreJobsApi.ReturnOneRestoreJobForOneServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the serverless instance. | 
**restoreJobId** | **string** | Unique 24-hexadecimal digit string that identifies the restore job to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneRestoreJobForOneServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasServerlessBackupRestoreJobViewManual**](ApiAtlasServerlessBackupRestoreJobViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneRestoreJobOfOneCluster

> ApiAtlasDiskBackupRestoreJobView ReturnOneRestoreJobOfOneCluster(ctx, groupId, clusterName, restoreJobId).Envelope(envelope).Pretty(pretty).Execute()

Return One Restore Job of One Cluster



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster with the restore jobs you want to return.
    restoreJobId := "restoreJobId_example" // string | Unique 24-hexadecimal digit string that identifies the restore job to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupRestoreJobsApi.ReturnOneRestoreJobOfOneCluster(context.Background(), groupId, clusterName, restoreJobId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupRestoreJobsApi.ReturnOneRestoreJobOfOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneRestoreJobOfOneCluster`: ApiAtlasDiskBackupRestoreJobView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupRestoreJobsApi.ReturnOneRestoreJobOfOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the restore jobs you want to return. | 
**restoreJobId** | **string** | Unique 24-hexadecimal digit string that identifies the restore job to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneRestoreJobOfOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDiskBackupRestoreJobView**](ApiAtlasDiskBackupRestoreJobView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

