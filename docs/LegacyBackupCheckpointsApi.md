# \LegacyBackupCheckpointsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnAllLegacyBackupCheckpoints**](LegacyBackupCheckpointsApi.md#ReturnAllLegacyBackupCheckpoints) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backupCheckpoints | Return All Legacy Backup Checkpoints
[**ReturnOneLegacyBackupCheckpoint**](LegacyBackupCheckpointsApi.md#ReturnOneLegacyBackupCheckpoint) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backupCheckpoints/{checkpointId} | Return One Legacy Backup Checkpoint



## ReturnAllLegacyBackupCheckpoints

> []ApiAtlasCheckpointView ReturnAllLegacyBackupCheckpoints(ctx, groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Legacy Backup Checkpoints



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster that contains the checkpoints that you want to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyBackupCheckpointsApi.ReturnAllLegacyBackupCheckpoints(context.Background(), groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupCheckpointsApi.ReturnAllLegacyBackupCheckpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllLegacyBackupCheckpoints`: []ApiAtlasCheckpointView
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupCheckpointsApi.ReturnAllLegacyBackupCheckpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the checkpoints that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllLegacyBackupCheckpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiAtlasCheckpointView**](ApiAtlasCheckpointView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneLegacyBackupCheckpoint

> ApiAtlasCheckpointView ReturnOneLegacyBackupCheckpoint(ctx, groupId, checkpointId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return One Legacy Backup Checkpoint



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
    checkpointId := "checkpointId_example" // string | Unique 24-hexadecimal digit string that identifies the checkpoint.
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster that contains the checkpoints that you want to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyBackupCheckpointsApi.ReturnOneLegacyBackupCheckpoint(context.Background(), groupId, checkpointId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupCheckpointsApi.ReturnOneLegacyBackupCheckpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneLegacyBackupCheckpoint`: ApiAtlasCheckpointView
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupCheckpointsApi.ReturnOneLegacyBackupCheckpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**checkpointId** | **string** | Unique 24-hexadecimal digit string that identifies the checkpoint. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that contains the checkpoints that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneLegacyBackupCheckpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasCheckpointView**](ApiAtlasCheckpointView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

