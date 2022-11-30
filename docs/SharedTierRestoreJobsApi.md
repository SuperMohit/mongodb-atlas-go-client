# \SharedTierRestoreJobsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneRestoreJobFromOneM2OrM5Cluster**](SharedTierRestoreJobsApi.md#CreateOneRestoreJobFromOneM2OrM5Cluster) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/tenant/restore | Create One Restore Job from One M2 or M5 Cluster
[**ReturnAllRestoreJobsForOneM2OrM5Cluster**](SharedTierRestoreJobsApi.md#ReturnAllRestoreJobsForOneM2OrM5Cluster) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores | Return All Restore Jobs for One M2 or M5 Cluster
[**ReturnOneRestoreJobForOneM2OrM5Cluster**](SharedTierRestoreJobsApi.md#ReturnOneRestoreJobForOneM2OrM5Cluster) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores/{restoreId} | Return One Restore Job for One M2 or M5 Cluster



## CreateOneRestoreJobFromOneM2OrM5Cluster

> ApiAtlasTenantRestoreView CreateOneRestoreJobFromOneM2OrM5Cluster(ctx, clusterName, groupId).ApiAtlasTenantRestoreView(apiAtlasTenantRestoreView).Envelope(envelope).Pretty(pretty).Execute()

Create One Restore Job from One M2 or M5 Cluster



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
    apiAtlasTenantRestoreView := *openapiclient.NewApiAtlasTenantRestoreView([]openapiclient.Link{*openapiclient.NewLink()}, "78e18e587716102f711a5bca", "TargetDeploymentItemName_example") // ApiAtlasTenantRestoreView | The restore job details.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedTierRestoreJobsApi.CreateOneRestoreJobFromOneM2OrM5Cluster(context.Background(), clusterName, groupId).ApiAtlasTenantRestoreView(apiAtlasTenantRestoreView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedTierRestoreJobsApi.CreateOneRestoreJobFromOneM2OrM5Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneRestoreJobFromOneM2OrM5Cluster`: ApiAtlasTenantRestoreView
    fmt.Fprintf(os.Stdout, "Response from `SharedTierRestoreJobsApi.CreateOneRestoreJobFromOneM2OrM5Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneRestoreJobFromOneM2OrM5ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasTenantRestoreView** | [**ApiAtlasTenantRestoreView**](ApiAtlasTenantRestoreView.md) | The restore job details. | 
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


## ReturnAllRestoreJobsForOneM2OrM5Cluster

> PaginatedTenantRestoreView ReturnAllRestoreJobsForOneM2OrM5Cluster(ctx, clusterName, groupId).Envelope(envelope).Pretty(pretty).Execute()

Return All Restore Jobs for One M2 or M5 Cluster



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
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedTierRestoreJobsApi.ReturnAllRestoreJobsForOneM2OrM5Cluster(context.Background(), clusterName, groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedTierRestoreJobsApi.ReturnAllRestoreJobsForOneM2OrM5Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllRestoreJobsForOneM2OrM5Cluster`: PaginatedTenantRestoreView
    fmt.Fprintf(os.Stdout, "Response from `SharedTierRestoreJobsApi.ReturnAllRestoreJobsForOneM2OrM5Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllRestoreJobsForOneM2OrM5ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedTenantRestoreView**](PaginatedTenantRestoreView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneRestoreJobForOneM2OrM5Cluster

> ApiAtlasTenantRestoreView ReturnOneRestoreJobForOneM2OrM5Cluster(ctx, clusterName, groupId, restoreId).Envelope(envelope).Pretty(pretty).Execute()

Return One Restore Job for One M2 or M5 Cluster



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
    restoreId := "restoreId_example" // string | Unique 24-hexadecimal digit string that identifies the restore job to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharedTierRestoreJobsApi.ReturnOneRestoreJobForOneM2OrM5Cluster(context.Background(), clusterName, groupId, restoreId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedTierRestoreJobsApi.ReturnOneRestoreJobForOneM2OrM5Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneRestoreJobForOneM2OrM5Cluster`: ApiAtlasTenantRestoreView
    fmt.Fprintf(os.Stdout, "Response from `SharedTierRestoreJobsApi.ReturnOneRestoreJobForOneM2OrM5Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**restoreId** | **string** | Unique 24-hexadecimal digit string that identifies the restore job to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneRestoreJobForOneM2OrM5ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasTenantRestoreView**](ApiAtlasTenantRestoreView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

