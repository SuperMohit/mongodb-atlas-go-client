# \SharedTierSnapshotsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadOneM2OrM5ClusterSnapshot**](SharedTierSnapshotsApi.md#DownloadOneM2OrM5ClusterSnapshot) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/tenant/download | Download One M2 or M5 Cluster Snapshot
[**ReturnAllSnapshotsForOneM2OrM5Cluster**](SharedTierSnapshotsApi.md#ReturnAllSnapshotsForOneM2OrM5Cluster) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots | Return All Snapshots for One M2 or M5 Cluster
[**ReturnOneSnapshotForOneM2OrM5Cluster**](SharedTierSnapshotsApi.md#ReturnOneSnapshotForOneM2OrM5Cluster) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots/{snapshotId} | Return One Snapshot for One M2 or M5 Cluster



## DownloadOneM2OrM5ClusterSnapshot

> ApiAtlasTenantRestoreView DownloadOneM2OrM5ClusterSnapshot(ctx, clusterName, groupId).ApiAtlasTenantRestoreView(apiAtlasTenantRestoreView).Envelope(envelope).Pretty(pretty).Execute()

Download One M2 or M5 Cluster Snapshot



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster.
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    apiAtlasTenantRestoreView := *openapiclient.NewApiAtlasTenantRestoreView([]openapiclient.Link{*openapiclient.NewLink()}, "78e18e587716102f711a5bca", "TargetDeploymentItemName_example") // ApiAtlasTenantRestoreView | Snapshot to be downloaded.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedTierSnapshotsApi.DownloadOneM2OrM5ClusterSnapshot(context.Background(), clusterName, groupId).ApiAtlasTenantRestoreView(apiAtlasTenantRestoreView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedTierSnapshotsApi.DownloadOneM2OrM5ClusterSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadOneM2OrM5ClusterSnapshot`: ApiAtlasTenantRestoreView
    fmt.Fprintf(os.Stdout, "Response from `SharedTierSnapshotsApi.DownloadOneM2OrM5ClusterSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOneM2OrM5ClusterSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasTenantRestoreView** | [**ApiAtlasTenantRestoreView**](ApiAtlasTenantRestoreView.md) | Snapshot to be downloaded. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasTenantRestoreView**](ApiAtlasTenantRestoreView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllSnapshotsForOneM2OrM5Cluster

> PaginatedTenantSnapshotView ReturnAllSnapshotsForOneM2OrM5Cluster(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return All Snapshots for One M2 or M5 Cluster



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
    resp, r, err := apiClient.SharedTierSnapshotsApi.ReturnAllSnapshotsForOneM2OrM5Cluster(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedTierSnapshotsApi.ReturnAllSnapshotsForOneM2OrM5Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllSnapshotsForOneM2OrM5Cluster`: PaginatedTenantSnapshotView
    fmt.Fprintf(os.Stdout, "Response from `SharedTierSnapshotsApi.ReturnAllSnapshotsForOneM2OrM5Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllSnapshotsForOneM2OrM5ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedTenantSnapshotView**](PaginatedTenantSnapshotView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneSnapshotForOneM2OrM5Cluster

> ApiAtlasTenantSnapshotView ReturnOneSnapshotForOneM2OrM5Cluster(ctx, groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()

Return One Snapshot for One M2 or M5 Cluster



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
    snapshotId := "snapshotId_example" // string | Unique 24-hexadecimal digit string that identifies the desired snapshot.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedTierSnapshotsApi.ReturnOneSnapshotForOneM2OrM5Cluster(context.Background(), groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedTierSnapshotsApi.ReturnOneSnapshotForOneM2OrM5Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneSnapshotForOneM2OrM5Cluster`: ApiAtlasTenantSnapshotView
    fmt.Fprintf(os.Stdout, "Response from `SharedTierSnapshotsApi.ReturnOneSnapshotForOneM2OrM5Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**snapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the desired snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneSnapshotForOneM2OrM5ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasTenantSnapshotView**](ApiAtlasTenantSnapshotView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

