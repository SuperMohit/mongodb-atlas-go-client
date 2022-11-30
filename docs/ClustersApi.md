# \ClustersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckStatusOfClusterSampleDatasetRequest**](ClustersApi.md#CheckStatusOfClusterSampleDatasetRequest) | **Get** /api/atlas/v1.0/groups/{groupId}/sampleDatasetLoad/{sampleDatasetId} | Check Status of Cluster Sample Dataset Request
[**CreateOneCluster**](ClustersApi.md#CreateOneCluster) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters | Create One Cluster
[**LoadSampleDatasetRequestIntoCluster**](ClustersApi.md#LoadSampleDatasetRequestIntoCluster) | **Post** /api/atlas/v1.0/groups/{groupId}/sampleDatasetLoad/{name} | Load Sample Dataset Request into Cluster
[**RemoveOneCluster**](ClustersApi.md#RemoveOneCluster) | **Delete** /api/atlas/v1.0/groups/{groupId}/clusters/{name} | Remove One Cluster
[**ReturnAllAuthorizedClustersInAllProjects**](ClustersApi.md#ReturnAllAuthorizedClustersInAllProjects) | **Get** /api/atlas/v1.0/clusters | Return All Authorized Clusters in All Projects
[**ReturnAllCloudProviderRegions**](ClustersApi.md#ReturnAllCloudProviderRegions) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/provider/regions | Return All Cloud Provider Regions
[**ReturnAllClusters**](ClustersApi.md#ReturnAllClusters) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters | Return All Clusters in One Project
[**ReturnOneAdvancedConfigurationOptionsForOneCluster**](ClustersApi.md#ReturnOneAdvancedConfigurationOptionsForOneCluster) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/processArgs | Return One Advanced Configuration Options for One Cluster
[**ReturnOneCluster**](ClustersApi.md#ReturnOneCluster) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{name} | Return One Cluster in One Project
[**ReturnOneStatusOfAllClusterOperations**](ClustersApi.md#ReturnOneStatusOfAllClusterOperations) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/status | Return Status of All Cluster Operations
[**TestFailoverForOneCluster**](ClustersApi.md#TestFailoverForOneCluster) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/restartPrimaries | Test Failover for One Cluster
[**UpdateAdvancedConfigurationOptionsForOneCluster**](ClustersApi.md#UpdateAdvancedConfigurationOptionsForOneCluster) | **Patch** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/processArgs | Update Advanced Configuration Options for One Cluster
[**UpdateConfigurationOfOneCluster**](ClustersApi.md#UpdateConfigurationOfOneCluster) | **Patch** /api/atlas/v1.0/groups/{groupId}/clusters/{name} | Update Configuration of One Cluster
[**UpgradeOneTenantCluster**](ClustersApi.md#UpgradeOneTenantCluster) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/tenantUpgrade | Upgrade One Shared-tier Cluster



## CheckStatusOfClusterSampleDatasetRequest

> ApiAtlasSampleDatasetStatusView CheckStatusOfClusterSampleDatasetRequest(ctx, groupId, sampleDatasetId).Envelope(envelope).Execute()

Check Status of Cluster Sample Dataset Request



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
    sampleDatasetId := "sampleDatasetId_example" // string | Unique 24-hexadecimal digit string that identifies the loaded sample dataset.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.CheckStatusOfClusterSampleDatasetRequest(context.Background(), groupId, sampleDatasetId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CheckStatusOfClusterSampleDatasetRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckStatusOfClusterSampleDatasetRequest`: ApiAtlasSampleDatasetStatusView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.CheckStatusOfClusterSampleDatasetRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**sampleDatasetId** | **string** | Unique 24-hexadecimal digit string that identifies the loaded sample dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckStatusOfClusterSampleDatasetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiAtlasSampleDatasetStatusView**](ApiAtlasSampleDatasetStatusView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOneCluster

> ApiAtlasLegacyClusterDescriptionView CreateOneCluster(ctx, groupId).ApiAtlasLegacyClusterDescriptionView(apiAtlasLegacyClusterDescriptionView).Envelope(envelope).Pretty(pretty).Execute()

Create One Cluster



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
    apiAtlasLegacyClusterDescriptionView := *openapiclient.NewApiAtlasLegacyClusterDescriptionView([]openapiclient.Link{*openapiclient.NewLink()}) // ApiAtlasLegacyClusterDescriptionView | Details to create for one cluster in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.CreateOneCluster(context.Background(), groupId).ApiAtlasLegacyClusterDescriptionView(apiAtlasLegacyClusterDescriptionView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CreateOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneCluster`: ApiAtlasLegacyClusterDescriptionView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.CreateOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasLegacyClusterDescriptionView** | [**ApiAtlasLegacyClusterDescriptionView**](ApiAtlasLegacyClusterDescriptionView.md) | Details to create for one cluster in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasLegacyClusterDescriptionView**](ApiAtlasLegacyClusterDescriptionView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadSampleDatasetRequestIntoCluster

> []ApiAtlasSampleDatasetStatusView LoadSampleDatasetRequestIntoCluster(ctx, groupId, name).ApiAtlasSampleDatasetStatusView(apiAtlasSampleDatasetStatusView).Envelope(envelope).Execute()

Load Sample Dataset Request into Cluster



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
    name := "name_example" // string | Human-readable label that identifies the cluster into which you load the sample dataset.
    apiAtlasSampleDatasetStatusView := *openapiclient.NewApiAtlasSampleDatasetStatusView() // ApiAtlasSampleDatasetStatusView | Cluster into which to load the sample dataset.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.LoadSampleDatasetRequestIntoCluster(context.Background(), groupId, name).ApiAtlasSampleDatasetStatusView(apiAtlasSampleDatasetStatusView).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.LoadSampleDatasetRequestIntoCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadSampleDatasetRequestIntoCluster`: []ApiAtlasSampleDatasetStatusView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.LoadSampleDatasetRequestIntoCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**name** | **string** | Human-readable label that identifies the cluster into which you load the sample dataset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadSampleDatasetRequestIntoClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasSampleDatasetStatusView** | [**ApiAtlasSampleDatasetStatusView**](ApiAtlasSampleDatasetStatusView.md) | Cluster into which to load the sample dataset. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**[]ApiAtlasSampleDatasetStatusView**](ApiAtlasSampleDatasetStatusView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneCluster

> RemoveOneCluster(ctx, groupId, name).Envelope(envelope).Pretty(pretty).Execute()

Remove One Cluster



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
    name := "name_example" // string | Human-readable label that identifies the cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.RemoveOneCluster(context.Background(), groupId, name).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.RemoveOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**name** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneClusterRequest struct via the builder pattern


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


## ReturnAllAuthorizedClustersInAllProjects

> PaginatedOrgGroupView ReturnAllAuthorizedClustersInAllProjects(ctx).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Authorized Clusters in All Projects



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
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ReturnAllAuthorizedClustersInAllProjects(context.Background()).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ReturnAllAuthorizedClustersInAllProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllAuthorizedClustersInAllProjects`: PaginatedOrgGroupView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ReturnAllAuthorizedClustersInAllProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllAuthorizedClustersInAllProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedOrgGroupView**](PaginatedOrgGroupView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllCloudProviderRegions

> []ApiAtlasProviderRegionsView ReturnAllCloudProviderRegions(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Providers(providers).Tier(tier).Execute()

Return All Cloud Provider Regions



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
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    providers := []string{"Inner_example"} // []string | Cloud providers whose regions to retrieve. When you specify multiple providers, the response can return only tiers and regions that support multi-cloud clusters. (optional)
    tier := "tier_example" // string | Cluster tier for which to retrieve the regions. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ReturnAllCloudProviderRegions(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Providers(providers).Tier(tier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ReturnAllCloudProviderRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllCloudProviderRegions`: []ApiAtlasProviderRegionsView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ReturnAllCloudProviderRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllCloudProviderRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **providers** | **[]string** | Cloud providers whose regions to retrieve. When you specify multiple providers, the response can return only tiers and regions that support multi-cloud clusters. | 
 **tier** | **string** | Cluster tier for which to retrieve the regions. | 

### Return type

[**[]ApiAtlasProviderRegionsView**](ApiAtlasProviderRegionsView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllClusters

> PaginatedLegacyClusterView ReturnAllClusters(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Clusters in One Project



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
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ReturnAllClusters(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ReturnAllClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllClusters`: PaginatedLegacyClusterView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ReturnAllClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedLegacyClusterView**](PaginatedLegacyClusterView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneAdvancedConfigurationOptionsForOneCluster

> ApiAtlasClusterDescriptionProcessArgsView ReturnOneAdvancedConfigurationOptionsForOneCluster(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return One Advanced Configuration Options for One Cluster



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
    resp, r, err := apiClient.ClustersApi.ReturnOneAdvancedConfigurationOptionsForOneCluster(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ReturnOneAdvancedConfigurationOptionsForOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneAdvancedConfigurationOptionsForOneCluster`: ApiAtlasClusterDescriptionProcessArgsView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ReturnOneAdvancedConfigurationOptionsForOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneAdvancedConfigurationOptionsForOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasClusterDescriptionProcessArgsView**](ApiAtlasClusterDescriptionProcessArgsView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneCluster

> ApiAtlasLegacyClusterDescriptionView ReturnOneCluster(ctx, groupId, name).Envelope(envelope).Pretty(pretty).Execute()

Return One Cluster in One Project



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
    name := "name_example" // string | Human-readable label that identifies the cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ReturnOneCluster(context.Background(), groupId, name).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ReturnOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneCluster`: ApiAtlasLegacyClusterDescriptionView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ReturnOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**name** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasLegacyClusterDescriptionView**](ApiAtlasLegacyClusterDescriptionView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneStatusOfAllClusterOperations

> ApiAtlasClusterStatusView ReturnOneStatusOfAllClusterOperations(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return Status of All Cluster Operations



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
    resp, r, err := apiClient.ClustersApi.ReturnOneStatusOfAllClusterOperations(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ReturnOneStatusOfAllClusterOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneStatusOfAllClusterOperations`: ApiAtlasClusterStatusView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ReturnOneStatusOfAllClusterOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneStatusOfAllClusterOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasClusterStatusView**](ApiAtlasClusterStatusView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestFailoverForOneCluster

> TestFailoverForOneCluster(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Test Failover for One Cluster



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
    resp, r, err := apiClient.ClustersApi.TestFailoverForOneCluster(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.TestFailoverForOneCluster``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiTestFailoverForOneClusterRequest struct via the builder pattern


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


## UpdateAdvancedConfigurationOptionsForOneCluster

> ApiAtlasClusterDescriptionProcessArgsView UpdateAdvancedConfigurationOptionsForOneCluster(ctx, groupId, clusterName).ApiAtlasClusterDescriptionProcessArgsView(apiAtlasClusterDescriptionProcessArgsView).Envelope(envelope).Pretty(pretty).Execute()

Update Advanced Configuration Options for One Cluster



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
    apiAtlasClusterDescriptionProcessArgsView := *openapiclient.NewApiAtlasClusterDescriptionProcessArgsView() // ApiAtlasClusterDescriptionProcessArgsView | Advanced configuration details to add for one cluster in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.UpdateAdvancedConfigurationOptionsForOneCluster(context.Background(), groupId, clusterName).ApiAtlasClusterDescriptionProcessArgsView(apiAtlasClusterDescriptionProcessArgsView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateAdvancedConfigurationOptionsForOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAdvancedConfigurationOptionsForOneCluster`: ApiAtlasClusterDescriptionProcessArgsView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateAdvancedConfigurationOptionsForOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdvancedConfigurationOptionsForOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasClusterDescriptionProcessArgsView** | [**ApiAtlasClusterDescriptionProcessArgsView**](ApiAtlasClusterDescriptionProcessArgsView.md) | Advanced configuration details to add for one cluster in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasClusterDescriptionProcessArgsView**](ApiAtlasClusterDescriptionProcessArgsView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationOfOneCluster

> ApiAtlasLegacyClusterDescriptionView UpdateConfigurationOfOneCluster(ctx, groupId, name).ApiAtlasLegacyClusterDescriptionView(apiAtlasLegacyClusterDescriptionView).Envelope(envelope).Pretty(pretty).Execute()

Update Configuration of One Cluster



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
    name := "name_example" // string | Human-readable label that identifies the cluster.
    apiAtlasLegacyClusterDescriptionView := *openapiclient.NewApiAtlasLegacyClusterDescriptionView([]openapiclient.Link{*openapiclient.NewLink()}) // ApiAtlasLegacyClusterDescriptionView | Details to update for one cluster in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.UpdateConfigurationOfOneCluster(context.Background(), groupId, name).ApiAtlasLegacyClusterDescriptionView(apiAtlasLegacyClusterDescriptionView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateConfigurationOfOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigurationOfOneCluster`: ApiAtlasLegacyClusterDescriptionView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateConfigurationOfOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**name** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationOfOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasLegacyClusterDescriptionView** | [**ApiAtlasLegacyClusterDescriptionView**](ApiAtlasLegacyClusterDescriptionView.md) | Details to update for one cluster in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasLegacyClusterDescriptionView**](ApiAtlasLegacyClusterDescriptionView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeOneTenantCluster

> ApiAtlasLegacyClusterDescriptionView UpgradeOneTenantCluster(ctx, groupId).ApiAtlasLegacyClusterDescriptionView(apiAtlasLegacyClusterDescriptionView).Envelope(envelope).Pretty(pretty).Execute()

Upgrade One Shared-tier Cluster



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
    apiAtlasLegacyClusterDescriptionView := *openapiclient.NewApiAtlasLegacyClusterDescriptionView([]openapiclient.Link{*openapiclient.NewLink()}) // ApiAtlasLegacyClusterDescriptionView | Details of the tenant upgrade in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.UpgradeOneTenantCluster(context.Background(), groupId).ApiAtlasLegacyClusterDescriptionView(apiAtlasLegacyClusterDescriptionView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpgradeOneTenantCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeOneTenantCluster`: ApiAtlasLegacyClusterDescriptionView
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpgradeOneTenantCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeOneTenantClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasLegacyClusterDescriptionView** | [**ApiAtlasLegacyClusterDescriptionView**](ApiAtlasLegacyClusterDescriptionView.md) | Details of the tenant upgrade in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasLegacyClusterDescriptionView**](ApiAtlasLegacyClusterDescriptionView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

