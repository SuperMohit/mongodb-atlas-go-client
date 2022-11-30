# \NetworkPeeringContainersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneNewNetworkPeeringContainer**](NetworkPeeringContainersApi.md#CreateOneNewNetworkPeeringContainer) | **Post** /api/atlas/v1.0/groups/{groupId}/containers | Create One New Network Peering Container
[**RemoveOneNetworkPeeringContainer**](NetworkPeeringContainersApi.md#RemoveOneNetworkPeeringContainer) | **Delete** /api/atlas/v1.0/groups/{groupId}/containers/{containerId} | Remove One Network Peering Container
[**ReturnAllNetworkPeeringContainersInOneProject**](NetworkPeeringContainersApi.md#ReturnAllNetworkPeeringContainersInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/containers/all | Return All Network Peering Containers in One Project
[**ReturnAllNetworkPeeringContainersInOneProjectForOneCloudProvider**](NetworkPeeringContainersApi.md#ReturnAllNetworkPeeringContainersInOneProjectForOneCloudProvider) | **Get** /api/atlas/v1.0/groups/{groupId}/containers | Return All Network Peering Containers in One Project for One Cloud Provider
[**ReturnOneNetworkPeeringContainer**](NetworkPeeringContainersApi.md#ReturnOneNetworkPeeringContainer) | **Get** /api/atlas/v1.0/groups/{groupId}/containers/{containerId} | Return One Network Peering Container
[**UpdateOneNetworkPeeringContainer**](NetworkPeeringContainersApi.md#UpdateOneNetworkPeeringContainer) | **Patch** /api/atlas/v1.0/groups/{groupId}/containers/{containerId} | Update One Network Peering Container



## CreateOneNewNetworkPeeringContainer

> ApiAtlasCloudProviderContainerView CreateOneNewNetworkPeeringContainer(ctx, groupId).ApiAtlasCloudProviderContainerView(apiAtlasCloudProviderContainerView).Envelope(envelope).Pretty(pretty).Execute()

Create One New Network Peering Container



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
    apiAtlasCloudProviderContainerView := *openapiclient.NewApiAtlasCloudProviderContainerView() // ApiAtlasCloudProviderContainerView | Creates one new network peering container in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPeeringContainersApi.CreateOneNewNetworkPeeringContainer(context.Background(), groupId).ApiAtlasCloudProviderContainerView(apiAtlasCloudProviderContainerView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringContainersApi.CreateOneNewNetworkPeeringContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneNewNetworkPeeringContainer`: ApiAtlasCloudProviderContainerView
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringContainersApi.CreateOneNewNetworkPeeringContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneNewNetworkPeeringContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasCloudProviderContainerView** | [**ApiAtlasCloudProviderContainerView**](ApiAtlasCloudProviderContainerView.md) | Creates one new network peering container in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasCloudProviderContainerView**](ApiAtlasCloudProviderContainerView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneNetworkPeeringContainer

> RemoveOneNetworkPeeringContainer(ctx, groupId, containerId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Network Peering Container



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
    containerId := "2f3ed24096c4b70db8993cb7" // string | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPeeringContainersApi.RemoveOneNetworkPeeringContainer(context.Background(), groupId, containerId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringContainersApi.RemoveOneNetworkPeeringContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**containerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneNetworkPeeringContainerRequest struct via the builder pattern


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


## ReturnAllNetworkPeeringContainersInOneProject

> PaginatedCloudProviderContainerView ReturnAllNetworkPeeringContainersInOneProject(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Network Peering Containers in One Project



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
    resp, r, err := apiClient.NetworkPeeringContainersApi.ReturnAllNetworkPeeringContainersInOneProject(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringContainersApi.ReturnAllNetworkPeeringContainersInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllNetworkPeeringContainersInOneProject`: PaginatedCloudProviderContainerView
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringContainersApi.ReturnAllNetworkPeeringContainersInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllNetworkPeeringContainersInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedCloudProviderContainerView**](PaginatedCloudProviderContainerView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllNetworkPeeringContainersInOneProjectForOneCloudProvider

> PaginatedCloudProviderContainerView ReturnAllNetworkPeeringContainersInOneProjectForOneCloudProvider(ctx, groupId).ProviderName(providerName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Network Peering Containers in One Project for One Cloud Provider



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
    providerName := "providerName_example" // string | Cloud service provider that serves the desired network peering containers. (default to "AWS")
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPeeringContainersApi.ReturnAllNetworkPeeringContainersInOneProjectForOneCloudProvider(context.Background(), groupId).ProviderName(providerName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringContainersApi.ReturnAllNetworkPeeringContainersInOneProjectForOneCloudProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllNetworkPeeringContainersInOneProjectForOneCloudProvider`: PaginatedCloudProviderContainerView
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringContainersApi.ReturnAllNetworkPeeringContainersInOneProjectForOneCloudProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllNetworkPeeringContainersInOneProjectForOneCloudProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerName** | **string** | Cloud service provider that serves the desired network peering containers. | [default to &quot;AWS&quot;]
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedCloudProviderContainerView**](PaginatedCloudProviderContainerView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneNetworkPeeringContainer

> ApiAtlasCloudProviderContainerView ReturnOneNetworkPeeringContainer(ctx, groupId, containerId).Envelope(envelope).Pretty(pretty).Execute()

Return One Network Peering Container



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
    containerId := "2f3ed24096c4b70db8993cb7" // string | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPeeringContainersApi.ReturnOneNetworkPeeringContainer(context.Background(), groupId, containerId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringContainersApi.ReturnOneNetworkPeeringContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneNetworkPeeringContainer`: ApiAtlasCloudProviderContainerView
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringContainersApi.ReturnOneNetworkPeeringContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**containerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneNetworkPeeringContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasCloudProviderContainerView**](ApiAtlasCloudProviderContainerView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneNetworkPeeringContainer

> ApiAtlasCloudProviderContainerView UpdateOneNetworkPeeringContainer(ctx, groupId, containerId).ApiAtlasCloudProviderContainerView(apiAtlasCloudProviderContainerView).Envelope(envelope).Pretty(pretty).Execute()

Update One Network Peering Container



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
    containerId := "2f3ed24096c4b70db8993cb7" // string | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.
    apiAtlasCloudProviderContainerView := *openapiclient.NewApiAtlasCloudProviderContainerView() // ApiAtlasCloudProviderContainerView | Updates the network details and labels of one specified network peering container in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPeeringContainersApi.UpdateOneNetworkPeeringContainer(context.Background(), groupId, containerId).ApiAtlasCloudProviderContainerView(apiAtlasCloudProviderContainerView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringContainersApi.UpdateOneNetworkPeeringContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneNetworkPeeringContainer`: ApiAtlasCloudProviderContainerView
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringContainersApi.UpdateOneNetworkPeeringContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**containerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneNetworkPeeringContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasCloudProviderContainerView** | [**ApiAtlasCloudProviderContainerView**](ApiAtlasCloudProviderContainerView.md) | Updates the network details and labels of one specified network peering container in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasCloudProviderContainerView**](ApiAtlasCloudProviderContainerView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

