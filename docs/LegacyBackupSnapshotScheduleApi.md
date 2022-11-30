# \LegacyBackupSnapshotScheduleApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnOneSnapshotSchedule**](LegacyBackupSnapshotScheduleApi.md#ReturnOneSnapshotSchedule) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Return One Snapshot Schedule
[**UpdateSnapshotScheduleForOneCluster**](LegacyBackupSnapshotScheduleApi.md#UpdateSnapshotScheduleForOneCluster) | **Patch** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Update Snapshot Schedule for One Cluster



## ReturnOneSnapshotSchedule

> ApiAtlasSnapshotScheduleView ReturnOneSnapshotSchedule(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return One Snapshot Schedule



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
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyBackupSnapshotScheduleApi.ReturnOneSnapshotSchedule(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupSnapshotScheduleApi.ReturnOneSnapshotSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneSnapshotSchedule`: ApiAtlasSnapshotScheduleView
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupSnapshotScheduleApi.ReturnOneSnapshotSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneSnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasSnapshotScheduleView**](ApiAtlasSnapshotScheduleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshotScheduleForOneCluster

> ApiAtlasSnapshotScheduleView UpdateSnapshotScheduleForOneCluster(ctx, groupId, clusterName).ApiAtlasSnapshotScheduleView(apiAtlasSnapshotScheduleView).Envelope(envelope).Pretty(pretty).Execute()

Update Snapshot Schedule for One Cluster



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
    apiAtlasSnapshotScheduleView := *openapiclient.NewApiAtlasSnapshotScheduleView(int32(123), "5faec72692cee15b51eb22a4", int32(123), "32b6e34b3d91647abb20e7b8", []openapiclient.Link{*openapiclient.NewLink()}, int32(123), int32(123), int32(123), int32(123), int32(123)) // ApiAtlasSnapshotScheduleView | Update the snapshot schedule for one cluster in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyBackupSnapshotScheduleApi.UpdateSnapshotScheduleForOneCluster(context.Background(), groupId, clusterName).ApiAtlasSnapshotScheduleView(apiAtlasSnapshotScheduleView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupSnapshotScheduleApi.UpdateSnapshotScheduleForOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSnapshotScheduleForOneCluster`: ApiAtlasSnapshotScheduleView
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupSnapshotScheduleApi.UpdateSnapshotScheduleForOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnapshotScheduleForOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasSnapshotScheduleView** | [**ApiAtlasSnapshotScheduleView**](ApiAtlasSnapshotScheduleView.md) | Update the snapshot schedule for one cluster in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasSnapshotScheduleView**](ApiAtlasSnapshotScheduleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

