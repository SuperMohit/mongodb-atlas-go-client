# \DataFederationPrivateNetworksApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneDataFederationPrivateEndpointForOneProject**](DataFederationPrivateNetworksApi.md#CreateOneDataFederationPrivateEndpointForOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/privateNetworkSettings/endpointIds | Create One Federated Database Instance and Online Archive Private Endpoint for One Project
[**RemoveOneDataFederationPrivateEndpointFromOneProject**](DataFederationPrivateNetworksApi.md#RemoveOneDataFederationPrivateEndpointFromOneProject) | **Delete** /api/atlas/v1.0/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Remove One Federated Database Instance and Online Archive Private Endpoint from One Project
[**ReturnAllDataFederationPrivateEndpointsInOneProject**](DataFederationPrivateNetworksApi.md#ReturnAllDataFederationPrivateEndpointsInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/privateNetworkSettings/endpointIds | Return All Federated Database Instance and Online Archive Private Endpoints in One Project
[**ReturnOneDataFederationPrivateEndpointInOneProject**](DataFederationPrivateNetworksApi.md#ReturnOneDataFederationPrivateEndpointInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Return One Federated Database Instance and Online Archive Private Endpoint in One Project



## CreateOneDataFederationPrivateEndpointForOneProject

> []ApiAtlasPrivateNetworkEndpointIdEntryView CreateOneDataFederationPrivateEndpointForOneProject(ctx, groupId).ApiAtlasPrivateNetworkEndpointIdEntryView(apiAtlasPrivateNetworkEndpointIdEntryView).Envelope(envelope).Pretty(pretty).Execute()

Create One Federated Database Instance and Online Archive Private Endpoint for One Project



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
    apiAtlasPrivateNetworkEndpointIdEntryView := *openapiclient.NewApiAtlasPrivateNetworkEndpointIdEntryView("vpce-3bf78b0ddee411ba1") // ApiAtlasPrivateNetworkEndpointIdEntryView | Private endpoint for Federated Database Instances and Online Archives to add to the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataFederationPrivateNetworksApi.CreateOneDataFederationPrivateEndpointForOneProject(context.Background(), groupId).ApiAtlasPrivateNetworkEndpointIdEntryView(apiAtlasPrivateNetworkEndpointIdEntryView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationPrivateNetworksApi.CreateOneDataFederationPrivateEndpointForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneDataFederationPrivateEndpointForOneProject`: []ApiAtlasPrivateNetworkEndpointIdEntryView
    fmt.Fprintf(os.Stdout, "Response from `DataFederationPrivateNetworksApi.CreateOneDataFederationPrivateEndpointForOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneDataFederationPrivateEndpointForOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasPrivateNetworkEndpointIdEntryView** | [**ApiAtlasPrivateNetworkEndpointIdEntryView**](ApiAtlasPrivateNetworkEndpointIdEntryView.md) | Private endpoint for Federated Database Instances and Online Archives to add to the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiAtlasPrivateNetworkEndpointIdEntryView**](ApiAtlasPrivateNetworkEndpointIdEntryView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneDataFederationPrivateEndpointFromOneProject

> RemoveOneDataFederationPrivateEndpointFromOneProject(ctx, groupId, endpointId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Federated Database Instance and Online Archive Private Endpoint from One Project



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
    endpointId := "endpointId_example" // string | Unique 22-character alphanumeric string that identifies the private endpoint to remove. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataFederationPrivateNetworksApi.RemoveOneDataFederationPrivateEndpointFromOneProject(context.Background(), groupId, endpointId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationPrivateNetworksApi.RemoveOneDataFederationPrivateEndpointFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**endpointId** | **string** | Unique 22-character alphanumeric string that identifies the private endpoint to remove. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneDataFederationPrivateEndpointFromOneProjectRequest struct via the builder pattern


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


## ReturnAllDataFederationPrivateEndpointsInOneProject

> []ApiAtlasPrivateNetworkEndpointIdEntryView ReturnAllDataFederationPrivateEndpointsInOneProject(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Return All Federated Database Instance and Online Archive Private Endpoints in One Project



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
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataFederationPrivateNetworksApi.ReturnAllDataFederationPrivateEndpointsInOneProject(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationPrivateNetworksApi.ReturnAllDataFederationPrivateEndpointsInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllDataFederationPrivateEndpointsInOneProject`: []ApiAtlasPrivateNetworkEndpointIdEntryView
    fmt.Fprintf(os.Stdout, "Response from `DataFederationPrivateNetworksApi.ReturnAllDataFederationPrivateEndpointsInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllDataFederationPrivateEndpointsInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiAtlasPrivateNetworkEndpointIdEntryView**](ApiAtlasPrivateNetworkEndpointIdEntryView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneDataFederationPrivateEndpointInOneProject

> ApiAtlasPrivateNetworkEndpointIdEntryView ReturnOneDataFederationPrivateEndpointInOneProject(ctx, groupId, endpointId).Envelope(envelope).Pretty(pretty).Execute()

Return One Federated Database Instance and Online Archive Private Endpoint in One Project



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
    endpointId := "endpointId_example" // string | Unique 22-character alphanumeric string that identifies the private endpoint to return. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataFederationPrivateNetworksApi.ReturnOneDataFederationPrivateEndpointInOneProject(context.Background(), groupId, endpointId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationPrivateNetworksApi.ReturnOneDataFederationPrivateEndpointInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneDataFederationPrivateEndpointInOneProject`: ApiAtlasPrivateNetworkEndpointIdEntryView
    fmt.Fprintf(os.Stdout, "Response from `DataFederationPrivateNetworksApi.ReturnOneDataFederationPrivateEndpointInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**endpointId** | **string** | Unique 22-character alphanumeric string that identifies the private endpoint to return. Atlas Data Federation supports AWS private endpoints using the AWS PrivateLink feature. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneDataFederationPrivateEndpointInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasPrivateNetworkEndpointIdEntryView**](ApiAtlasPrivateNetworkEndpointIdEntryView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

