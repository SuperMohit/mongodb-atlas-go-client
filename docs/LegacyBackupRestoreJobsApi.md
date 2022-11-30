# \LegacyBackupRestoreJobsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneLegacyBackupRestoreJob**](LegacyBackupRestoreJobsApi.md#CreateOneLegacyBackupRestoreJob) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/restoreJobs | Create One Legacy Backup Restore Job
[**ReturnAllLegacyBackupRestoreJobs**](LegacyBackupRestoreJobsApi.md#ReturnAllLegacyBackupRestoreJobs) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/restoreJobs | Return All Legacy Backup Restore Jobs
[**ReturnOneLegacyBackupRestoreJob**](LegacyBackupRestoreJobsApi.md#ReturnOneLegacyBackupRestoreJob) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/restoreJobs/{jobId} | Return One Legacy Backup Restore Job



## CreateOneLegacyBackupRestoreJob

> PaginatedRestoreJobView CreateOneLegacyBackupRestoreJob(ctx, groupId, clusterName).ApiRestoreJobView(apiRestoreJobView).Envelope(envelope).Pretty(pretty).Execute()

Create One Legacy Backup Restore Job



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster with the snapshot you want to return.
    apiRestoreJobView := *openapiclient.NewApiRestoreJobView(*openapiclient.NewApiRestoreJobDeliveryView("MethodName_example"), []openapiclient.Link{*openapiclient.NewLink()}) // ApiRestoreJobView | Legacy backup to restore to one cluster in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyBackupRestoreJobsApi.CreateOneLegacyBackupRestoreJob(context.Background(), groupId, clusterName).ApiRestoreJobView(apiRestoreJobView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupRestoreJobsApi.CreateOneLegacyBackupRestoreJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneLegacyBackupRestoreJob`: PaginatedRestoreJobView
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupRestoreJobsApi.CreateOneLegacyBackupRestoreJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneLegacyBackupRestoreJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiRestoreJobView** | [**ApiRestoreJobView**](ApiRestoreJobView.md) | Legacy backup to restore to one cluster in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedRestoreJobView**](PaginatedRestoreJobView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllLegacyBackupRestoreJobs

> PaginatedRestoreJobView ReturnAllLegacyBackupRestoreJobs(ctx, groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).BatchId(batchId).Execute()

Return All Legacy Backup Restore Jobs



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster with the snapshot you want to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    batchId := "batchId_example" // string | Unique 24-hexadecimal digit string that identifies the batch of restore jobs to return. Timestamp in ISO 8601 date and time format in UTC when creating a restore job for a sharded cluster, Application creates a separate job for each shard, plus another for the config host. Each of these jobs comprise one batch. A restore job for a replica set can't be part of a batch. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyBackupRestoreJobsApi.ReturnAllLegacyBackupRestoreJobs(context.Background(), groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).BatchId(batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupRestoreJobsApi.ReturnAllLegacyBackupRestoreJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllLegacyBackupRestoreJobs`: PaginatedRestoreJobView
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupRestoreJobsApi.ReturnAllLegacyBackupRestoreJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllLegacyBackupRestoreJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **batchId** | **string** | Unique 24-hexadecimal digit string that identifies the batch of restore jobs to return. Timestamp in ISO 8601 date and time format in UTC when creating a restore job for a sharded cluster, Application creates a separate job for each shard, plus another for the config host. Each of these jobs comprise one batch. A restore job for a replica set can&#39;t be part of a batch. | 

### Return type

[**PaginatedRestoreJobView**](PaginatedRestoreJobView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneLegacyBackupRestoreJob

> ApiRestoreJobView ReturnOneLegacyBackupRestoreJob(ctx, groupId, clusterName, jobId).Envelope(envelope).Pretty(pretty).Execute()

Return One Legacy Backup Restore Job



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster with the snapshot you want to return.
    jobId := "jobId_example" // string | Unique 24-hexadecimal digit string that identifies the restore job.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyBackupRestoreJobsApi.ReturnOneLegacyBackupRestoreJob(context.Background(), groupId, clusterName, jobId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupRestoreJobsApi.ReturnOneLegacyBackupRestoreJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneLegacyBackupRestoreJob`: ApiRestoreJobView
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupRestoreJobsApi.ReturnOneLegacyBackupRestoreJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | 
**jobId** | **string** | Unique 24-hexadecimal digit string that identifies the restore job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneLegacyBackupRestoreJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiRestoreJobView**](ApiRestoreJobView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

