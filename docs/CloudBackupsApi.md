# \CloudBackupsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeExpirationDateForOneCloudBackup**](CloudBackupsApi.md#ChangeExpirationDateForOneCloudBackup) | **Patch** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Change Expiration Date for One Cloud Backup
[**RemoveOneReplicaSetCloudBackup**](CloudBackupsApi.md#RemoveOneReplicaSetCloudBackup) | **Delete** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Remove One Replica Set Cloud Backup
[**RemoveOneShardedClusterCloudBackup**](CloudBackupsApi.md#RemoveOneShardedClusterCloudBackup) | **Delete** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Remove One Sharded Cluster Cloud Backup
[**ReturnAllReplicaSetCloudBackups**](CloudBackupsApi.md#ReturnAllReplicaSetCloudBackups) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Return All Replica Set Cloud Backups
[**ReturnAllShardedClusterCloudBackups**](CloudBackupsApi.md#ReturnAllShardedClusterCloudBackups) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedClusters | Return All Sharded Cluster Cloud Backups
[**ReturnAllSnapshotsOfOneServerlessInstance**](CloudBackupsApi.md#ReturnAllSnapshotsOfOneServerlessInstance) | **Get** /api/atlas/v1.0/groups/{groupId}/serverless/{clusterName}/backup/snapshots | Return All Snapshots of One Serverless Instance
[**ReturnOneReplicaSetCloudBackup**](CloudBackupsApi.md#ReturnOneReplicaSetCloudBackup) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Return One Replica Set Cloud Backup
[**ReturnOneShardedClusterCloudBackup**](CloudBackupsApi.md#ReturnOneShardedClusterCloudBackup) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Return One Sharded Cluster Cloud Backup
[**ReturnOneSnapshotOfOneServerlessInstance**](CloudBackupsApi.md#ReturnOneSnapshotOfOneServerlessInstance) | **Get** /api/atlas/v1.0/groups/{groupId}/serverless/{clusterName}/backup/snapshots/{snapshotId} | Return One Snapshot of One Serverless Instance
[**TakeOneOnDemandSnapshot**](CloudBackupsApi.md#TakeOneOnDemandSnapshot) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Take One On-Demand Snapshot



## ChangeExpirationDateForOneCloudBackup

> ApiAtlasDiskBackupReplicaSetView ChangeExpirationDateForOneCloudBackup(ctx, groupId, clusterName, snapshotId).ApiAtlasSnapshotRetentionView(apiAtlasSnapshotRetentionView).Envelope(envelope).Pretty(pretty).Execute()

Change Expiration Date for One Cloud Backup



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
    apiAtlasSnapshotRetentionView := *openapiclient.NewApiAtlasSnapshotRetentionView([]openapiclient.Link{*openapiclient.NewLink()}, "RetentionUnit_example", int32(5)) // ApiAtlasSnapshotRetentionView | Changes the expiration date for one cloud backup snapshot for one cluster in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupsApi.ChangeExpirationDateForOneCloudBackup(context.Background(), groupId, clusterName, snapshotId).ApiAtlasSnapshotRetentionView(apiAtlasSnapshotRetentionView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ChangeExpirationDateForOneCloudBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeExpirationDateForOneCloudBackup`: ApiAtlasDiskBackupReplicaSetView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ChangeExpirationDateForOneCloudBackup`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiChangeExpirationDateForOneCloudBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiAtlasSnapshotRetentionView** | [**ApiAtlasSnapshotRetentionView**](ApiAtlasSnapshotRetentionView.md) | Changes the expiration date for one cloud backup snapshot for one cluster in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDiskBackupReplicaSetView**](ApiAtlasDiskBackupReplicaSetView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneReplicaSetCloudBackup

> RemoveOneReplicaSetCloudBackup(ctx, groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Replica Set Cloud Backup



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
    resp, r, err := apiClient.CloudBackupsApi.RemoveOneReplicaSetCloudBackup(context.Background(), groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.RemoveOneReplicaSetCloudBackup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveOneReplicaSetCloudBackupRequest struct via the builder pattern


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


## RemoveOneShardedClusterCloudBackup

> RemoveOneShardedClusterCloudBackup(ctx, groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Sharded Cluster Cloud Backup



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
    resp, r, err := apiClient.CloudBackupsApi.RemoveOneShardedClusterCloudBackup(context.Background(), groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.RemoveOneShardedClusterCloudBackup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveOneShardedClusterCloudBackupRequest struct via the builder pattern


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


## ReturnAllReplicaSetCloudBackups

> PaginatedCloudBackupReplicaSetView ReturnAllReplicaSetCloudBackups(ctx, groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Replica Set Cloud Backups



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
    resp, r, err := apiClient.CloudBackupsApi.ReturnAllReplicaSetCloudBackups(context.Background(), groupId, clusterName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ReturnAllReplicaSetCloudBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllReplicaSetCloudBackups`: PaginatedCloudBackupReplicaSetView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ReturnAllReplicaSetCloudBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllReplicaSetCloudBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedCloudBackupReplicaSetView**](PaginatedCloudBackupReplicaSetView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllShardedClusterCloudBackups

> PaginatedCloudBackupShardedClusterSnapshotView ReturnAllShardedClusterCloudBackups(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return All Sharded Cluster Cloud Backups



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
    resp, r, err := apiClient.CloudBackupsApi.ReturnAllShardedClusterCloudBackups(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ReturnAllShardedClusterCloudBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllShardedClusterCloudBackups`: PaginatedCloudBackupShardedClusterSnapshotView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ReturnAllShardedClusterCloudBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllShardedClusterCloudBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedCloudBackupShardedClusterSnapshotView**](PaginatedCloudBackupShardedClusterSnapshotView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllSnapshotsOfOneServerlessInstance

> []PaginatedServerlessBackupSnapshotViewManual ReturnAllSnapshotsOfOneServerlessInstance(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Snapshots of One Serverless Instance



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
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupsApi.ReturnAllSnapshotsOfOneServerlessInstance(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ReturnAllSnapshotsOfOneServerlessInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllSnapshotsOfOneServerlessInstance`: []PaginatedServerlessBackupSnapshotViewManual
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ReturnAllSnapshotsOfOneServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the serverless instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllSnapshotsOfOneServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**[]PaginatedServerlessBackupSnapshotViewManual**](PaginatedServerlessBackupSnapshotViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneReplicaSetCloudBackup

> ApiAtlasDiskBackupReplicaSetView ReturnOneReplicaSetCloudBackup(ctx, groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()

Return One Replica Set Cloud Backup



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
    resp, r, err := apiClient.CloudBackupsApi.ReturnOneReplicaSetCloudBackup(context.Background(), groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ReturnOneReplicaSetCloudBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneReplicaSetCloudBackup`: ApiAtlasDiskBackupReplicaSetView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ReturnOneReplicaSetCloudBackup`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReturnOneReplicaSetCloudBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDiskBackupReplicaSetView**](ApiAtlasDiskBackupReplicaSetView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneShardedClusterCloudBackup

> ApiAtlasDiskBackupShardedClusterSnapshotView ReturnOneShardedClusterCloudBackup(ctx, groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()

Return One Sharded Cluster Cloud Backup



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
    resp, r, err := apiClient.CloudBackupsApi.ReturnOneShardedClusterCloudBackup(context.Background(), groupId, clusterName, snapshotId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ReturnOneShardedClusterCloudBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneShardedClusterCloudBackup`: ApiAtlasDiskBackupShardedClusterSnapshotView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ReturnOneShardedClusterCloudBackup`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReturnOneShardedClusterCloudBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDiskBackupShardedClusterSnapshotView**](ApiAtlasDiskBackupShardedClusterSnapshotView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneSnapshotOfOneServerlessInstance

> ApiAtlasServerlessBackupSnapshotViewManual ReturnOneSnapshotOfOneServerlessInstance(ctx, groupId, clusterName, snapshotId).Envelope(envelope).Execute()

Return One Snapshot of One Serverless Instance



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
    snapshotId := "78e18e587716102f711a5bca" // string | Unique 24-hexadecimal digit string that identifies the desired snapshot.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupsApi.ReturnOneSnapshotOfOneServerlessInstance(context.Background(), groupId, clusterName, snapshotId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.ReturnOneSnapshotOfOneServerlessInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneSnapshotOfOneServerlessInstance`: ApiAtlasServerlessBackupSnapshotViewManual
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.ReturnOneSnapshotOfOneServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the serverless instance. | 
**snapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the desired snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneSnapshotOfOneServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiAtlasServerlessBackupSnapshotViewManual**](ApiAtlasServerlessBackupSnapshotViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TakeOneOnDemandSnapshot

> ApiAtlasDiskBackupSnapshotView TakeOneOnDemandSnapshot(ctx, groupId, clusterName).ApiAtlasDiskBackupOnDemandSnapshotRequestView(apiAtlasDiskBackupOnDemandSnapshotRequestView).Envelope(envelope).Pretty(pretty).Execute()

Take One On-Demand Snapshot



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
    apiAtlasDiskBackupOnDemandSnapshotRequestView := *openapiclient.NewApiAtlasDiskBackupOnDemandSnapshotRequestView([]openapiclient.Link{*openapiclient.NewLink()}) // ApiAtlasDiskBackupOnDemandSnapshotRequestView | Takes one on-demand snapshot.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupsApi.TakeOneOnDemandSnapshot(context.Background(), groupId, clusterName).ApiAtlasDiskBackupOnDemandSnapshotRequestView(apiAtlasDiskBackupOnDemandSnapshotRequestView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupsApi.TakeOneOnDemandSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TakeOneOnDemandSnapshot`: ApiAtlasDiskBackupSnapshotView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupsApi.TakeOneOnDemandSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTakeOneOnDemandSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasDiskBackupOnDemandSnapshotRequestView** | [**ApiAtlasDiskBackupOnDemandSnapshotRequestView**](ApiAtlasDiskBackupOnDemandSnapshotRequestView.md) | Takes one on-demand snapshot. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDiskBackupSnapshotView**](ApiAtlasDiskBackupSnapshotView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

