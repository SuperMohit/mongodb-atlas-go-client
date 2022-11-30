# \NetworkPeeringConnectionsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneNewNetworkPeeringConnection**](NetworkPeeringConnectionsApi.md#CreateOneNewNetworkPeeringConnection) | **Post** /api/atlas/v1.0/groups/{groupId}/peers | Create One New Network Peering Connection
[**RemoveOneExistingNetworkPeeringConnection**](NetworkPeeringConnectionsApi.md#RemoveOneExistingNetworkPeeringConnection) | **Delete** /api/atlas/v1.0/groups/{groupId}/peers/{peerId} | Remove One Existing Network Peering Connection
[**ReturnAllNetworkPeeringConnectionsInOneProject**](NetworkPeeringConnectionsApi.md#ReturnAllNetworkPeeringConnectionsInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/peers | Return All Network Peering Connections in One Project
[**ReturnOneNetworkPeeringConnectionInOneProject**](NetworkPeeringConnectionsApi.md#ReturnOneNetworkPeeringConnectionInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/peers/{peerId} | Return One Network Peering Connection in One Project
[**UpdateOneNewNetworkPeeringConnection**](NetworkPeeringConnectionsApi.md#UpdateOneNewNetworkPeeringConnection) | **Patch** /api/atlas/v1.0/groups/{groupId}/peers/{peerId} | Update One New Network Peering Connection



## CreateOneNewNetworkPeeringConnection

> CreateOneNewNetworkPeeringConnection200Response CreateOneNewNetworkPeeringConnection(ctx, groupId).ApiAtlasContainerPeerViewRequest(apiAtlasContainerPeerViewRequest).Envelope(envelope).Pretty(pretty).Execute()

Create One New Network Peering Connection



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
    apiAtlasContainerPeerViewRequest := openapiclient.ApiAtlasContainerPeerViewRequest{ApiAtlasAWSPeerVpcRequestView: openapiclient.NewApiAtlasAWSPeerVpcRequestView("2f3ed24096c4b70db8993cb7", "ProviderName_example", "AccepterRegionName_example", "AwsAccountId_example", "RouteTableCidrBlock_example", "VpcId_example")} // ApiAtlasContainerPeerViewRequest | Create one network peering connection.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPeeringConnectionsApi.CreateOneNewNetworkPeeringConnection(context.Background(), groupId).ApiAtlasContainerPeerViewRequest(apiAtlasContainerPeerViewRequest).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringConnectionsApi.CreateOneNewNetworkPeeringConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneNewNetworkPeeringConnection`: CreateOneNewNetworkPeeringConnection200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringConnectionsApi.CreateOneNewNetworkPeeringConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneNewNetworkPeeringConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasContainerPeerViewRequest** | [**ApiAtlasContainerPeerViewRequest**](ApiAtlasContainerPeerViewRequest.md) | Create one network peering connection. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**CreateOneNewNetworkPeeringConnection200Response**](CreateOneNewNetworkPeeringConnection200Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneExistingNetworkPeeringConnection

> RemoveOneExistingNetworkPeeringConnection(ctx, groupId, peerId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Existing Network Peering Connection



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
    peerId := "peerId_example" // string | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to delete.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPeeringConnectionsApi.RemoveOneExistingNetworkPeeringConnection(context.Background(), groupId, peerId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringConnectionsApi.RemoveOneExistingNetworkPeeringConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**peerId** | **string** | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneExistingNetworkPeeringConnectionRequest struct via the builder pattern


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


## ReturnAllNetworkPeeringConnectionsInOneProject

> ReturnAllNetworkPeeringConnectionsInOneProject200Response ReturnAllNetworkPeeringConnectionsInOneProject(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).ProviderName(providerName).Execute()

Return All Network Peering Connections in One Project



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
    providerName := "providerName_example" // string | Cloud service provider to use for this VPC peering connection. (optional) (default to "AWS")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPeeringConnectionsApi.ReturnAllNetworkPeeringConnectionsInOneProject(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).ProviderName(providerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringConnectionsApi.ReturnAllNetworkPeeringConnectionsInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllNetworkPeeringConnectionsInOneProject`: ReturnAllNetworkPeeringConnectionsInOneProject200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringConnectionsApi.ReturnAllNetworkPeeringConnectionsInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllNetworkPeeringConnectionsInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **providerName** | **string** | Cloud service provider to use for this VPC peering connection. | [default to &quot;AWS&quot;]

### Return type

[**ReturnAllNetworkPeeringConnectionsInOneProject200Response**](ReturnAllNetworkPeeringConnectionsInOneProject200Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneNetworkPeeringConnectionInOneProject

> ReturnOneNetworkPeeringConnectionInOneProject200Response ReturnOneNetworkPeeringConnectionInOneProject(ctx, groupId, peerId).Envelope(envelope).Pretty(pretty).Execute()

Return One Network Peering Connection in One Project



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
    peerId := "peerId_example" // string | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to retrieve.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPeeringConnectionsApi.ReturnOneNetworkPeeringConnectionInOneProject(context.Background(), groupId, peerId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringConnectionsApi.ReturnOneNetworkPeeringConnectionInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneNetworkPeeringConnectionInOneProject`: ReturnOneNetworkPeeringConnectionInOneProject200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringConnectionsApi.ReturnOneNetworkPeeringConnectionInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**peerId** | **string** | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneNetworkPeeringConnectionInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ReturnOneNetworkPeeringConnectionInOneProject200Response**](ReturnOneNetworkPeeringConnectionInOneProject200Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneNewNetworkPeeringConnection

> ReturnOneNetworkPeeringConnectionInOneProject200Response UpdateOneNewNetworkPeeringConnection(ctx, groupId, peerId).ApiAtlasContainerPeerViewRequest(apiAtlasContainerPeerViewRequest).Envelope(envelope).Pretty(pretty).Execute()

Update One New Network Peering Connection



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
    peerId := "peerId_example" // string | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to update.
    apiAtlasContainerPeerViewRequest := openapiclient.ApiAtlasContainerPeerViewRequest{ApiAtlasAWSPeerVpcRequestView: openapiclient.NewApiAtlasAWSPeerVpcRequestView("2f3ed24096c4b70db8993cb7", "ProviderName_example", "AccepterRegionName_example", "AwsAccountId_example", "RouteTableCidrBlock_example", "VpcId_example")} // ApiAtlasContainerPeerViewRequest | Modify one network peering connection.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPeeringConnectionsApi.UpdateOneNewNetworkPeeringConnection(context.Background(), groupId, peerId).ApiAtlasContainerPeerViewRequest(apiAtlasContainerPeerViewRequest).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringConnectionsApi.UpdateOneNewNetworkPeeringConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneNewNetworkPeeringConnection`: ReturnOneNetworkPeeringConnectionInOneProject200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringConnectionsApi.UpdateOneNewNetworkPeeringConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**peerId** | **string** | Unique 24-hexadecimal digit string that identifies the network peering connection that you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneNewNetworkPeeringConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasContainerPeerViewRequest** | [**ApiAtlasContainerPeerViewRequest**](ApiAtlasContainerPeerViewRequest.md) | Modify one network peering connection. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ReturnOneNetworkPeeringConnectionInOneProject200Response**](ReturnOneNetworkPeeringConnectionInOneProject200Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

