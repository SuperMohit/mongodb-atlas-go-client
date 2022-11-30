# \LegacyBackupSnapshotsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeOneLegacyBackupSnapshotExpiration**](LegacyBackupSnapshotsApi.md#ChangeOneLegacyBackupSnapshotExpiration) | **Patch** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Change One Legacy Backup Snapshot Expiration
[**RemoveOneLegacyBackupSnapshot**](LegacyBackupSnapshotsApi.md#RemoveOneLegacyBackupSnapshot) | **Delete** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Remove One Legacy Backup Snapshot
[**ReturnAllLegacyBackupSnapshots**](LegacyBackupSnapshotsApi.md#ReturnAllLegacyBackupSnapshots) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/snapshots | Return All Legacy Backup Snapshots
[**ReturnOneLegacyBackupSnapshot**](LegacyBackupSnapshotsApi.md#ReturnOneLegacyBackupSnapshot) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Return One Legacy Backup Snapshot



## ChangeOneLegacyBackupSnapshotExpiration

> ApiSnapshotView ChangeOneLegacyBackupSnapshotExpiration(ctx, groupId, clusterName, snapshotId).ApiSnapshotView(apiSnapshotView).Envelope(envelope).Pretty(pretty).Execute()

Change One Legacy Backup Snapshot Expiration



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
    apiSnapshotView := *openapiclient.NewApiSnapshotView([]openapiclient.Link{*openapiclient.NewLink()}) // ApiSnapshotView | Changes One Legacy Backup Snapshot Expiration.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyBackupSnapshotsApi.ChangeOneLegacyBackupSnapshotExpiration(context.Background(), groupId, clusterName, snapshotId).ApiSnapshotView(apiSnapshotView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupSnapshotsApi.ChangeOneLegacyBackupSnapshotExpiration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeOneLegacyBackupSnapshotExpiration`: ApiSnapshotView
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupSnapshotsApi.ChangeOneLegacyBackupSnapshotExpiration`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiChangeOneLegacyBackupSnapshotExpirationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiSnapshotView** | [**ApiSnapshotView**](ApiSnapshotView.md) | Changes One Legacy Backup Snapshot Expiration. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiSnapshotView**](ApiSnapshotView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneLegacyBackupSnapshot

> RemoveOneLegacyBackupSnapshot(ctx, groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Legacy Backup Snapshot



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
    resp, r, err := apiClient.LegacyBackupSnapshotsApi.RemoveOneLegacyBackupSnapshot(context.Background(), groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupSnapshotsApi.RemoveOneLegacyBackupSnapshot``: %v\n", err)
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
**snapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the desired snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneLegacyBackupSnapshotRequest struct via the builder pattern


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


## ReturnAllLegacyBackupSnapshots

> PaginatedSnapshotView ReturnAllLegacyBackupSnapshots(ctx, groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Legacy Backup Snapshots



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
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyBackupSnapshotsApi.ReturnAllLegacyBackupSnapshots(context.Background(), groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupSnapshotsApi.ReturnAllLegacyBackupSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllLegacyBackupSnapshots`: PaginatedSnapshotView
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupSnapshotsApi.ReturnAllLegacyBackupSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllLegacyBackupSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedSnapshotView**](PaginatedSnapshotView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneLegacyBackupSnapshot

> ApiSnapshotView ReturnOneLegacyBackupSnapshot(ctx, groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()

Return One Legacy Backup Snapshot



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
    resp, r, err := apiClient.LegacyBackupSnapshotsApi.ReturnOneLegacyBackupSnapshot(context.Background(), groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyBackupSnapshotsApi.ReturnOneLegacyBackupSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneLegacyBackupSnapshot`: ApiSnapshotView
    fmt.Fprintf(os.Stdout, "Response from `LegacyBackupSnapshotsApi.ReturnOneLegacyBackupSnapshot`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReturnOneLegacyBackupSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiSnapshotView**](ApiSnapshotView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

