# \CloudBackupScheduleApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoveAllCloudBackupSchedules**](CloudBackupScheduleApi.md#RemoveAllCloudBackupSchedules) | **Delete** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/schedule | Remove All Cloud Backup Schedules
[**ReturnOneCloudBackupSchedule**](CloudBackupScheduleApi.md#ReturnOneCloudBackupSchedule) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/schedule | Return One Cloud Backup Schedule
[**UpdateCloudBackupBackupPolicyForOneCluster**](CloudBackupScheduleApi.md#UpdateCloudBackupBackupPolicyForOneCluster) | **Patch** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/schedule | Update Cloud Backup Schedule for One Cluster



## RemoveAllCloudBackupSchedules

> ApiAtlasDiskBackupSnapshotScheduleView RemoveAllCloudBackupSchedules(ctx, groupId, clusterName).Envelope(envelope).Execute()

Remove All Cloud Backup Schedules



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
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupScheduleApi.RemoveAllCloudBackupSchedules(context.Background(), groupId, clusterName).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupScheduleApi.RemoveAllCloudBackupSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveAllCloudBackupSchedules`: ApiAtlasDiskBackupSnapshotScheduleView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupScheduleApi.RemoveAllCloudBackupSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAllCloudBackupSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiAtlasDiskBackupSnapshotScheduleView**](ApiAtlasDiskBackupSnapshotScheduleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneCloudBackupSchedule

> ApiAtlasDiskBackupSnapshotScheduleView ReturnOneCloudBackupSchedule(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return One Cloud Backup Schedule



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
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupScheduleApi.ReturnOneCloudBackupSchedule(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupScheduleApi.ReturnOneCloudBackupSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneCloudBackupSchedule`: ApiAtlasDiskBackupSnapshotScheduleView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupScheduleApi.ReturnOneCloudBackupSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneCloudBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDiskBackupSnapshotScheduleView**](ApiAtlasDiskBackupSnapshotScheduleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudBackupBackupPolicyForOneCluster

> ApiAtlasDiskBackupSnapshotScheduleView UpdateCloudBackupBackupPolicyForOneCluster(ctx, groupId, clusterName).ApiAtlasDiskBackupSnapshotScheduleView(apiAtlasDiskBackupSnapshotScheduleView).Envelope(envelope).Pretty(pretty).Execute()

Update Cloud Backup Schedule for One Cluster



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
    apiAtlasDiskBackupSnapshotScheduleView := *openapiclient.NewApiAtlasDiskBackupSnapshotScheduleView([]openapiclient.Link{*openapiclient.NewLink()}, []openapiclient.ApiPolicyView{*openapiclient.NewApiPolicyView([]openapiclient.ApiPolicyItemView{*openapiclient.NewApiPolicyItemView(int32(123), "FrequencyType_example", "RetentionUnit_example", int32(123))})}) // ApiAtlasDiskBackupSnapshotScheduleView | Updates the cloud backup schedule for one cluster within the specified project.  **Note**: In the request body, provide only the fields that you want to update.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupScheduleApi.UpdateCloudBackupBackupPolicyForOneCluster(context.Background(), groupId, clusterName).ApiAtlasDiskBackupSnapshotScheduleView(apiAtlasDiskBackupSnapshotScheduleView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupScheduleApi.UpdateCloudBackupBackupPolicyForOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCloudBackupBackupPolicyForOneCluster`: ApiAtlasDiskBackupSnapshotScheduleView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupScheduleApi.UpdateCloudBackupBackupPolicyForOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudBackupBackupPolicyForOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasDiskBackupSnapshotScheduleView** | [**ApiAtlasDiskBackupSnapshotScheduleView**](ApiAtlasDiskBackupSnapshotScheduleView.md) | Updates the cloud backup schedule for one cluster within the specified project.  **Note**: In the request body, provide only the fields that you want to update. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDiskBackupSnapshotScheduleView**](ApiAtlasDiskBackupSnapshotScheduleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

