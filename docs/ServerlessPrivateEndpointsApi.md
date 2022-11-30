# \ServerlessPrivateEndpointsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOnePrivateEndpointForOneServerlessInstance**](ServerlessPrivateEndpointsApi.md#CreateOnePrivateEndpointForOneServerlessInstance) | **Post** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Create One Private Endpoint for One Serverless Instance
[**RemoveOnePrivateEndpointFromOneServerlessInstance**](ServerlessPrivateEndpointsApi.md#RemoveOnePrivateEndpointFromOneServerlessInstance) | **Delete** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Remove One Private Endpoint for One Serverless Instance
[**ReturnAllPrivateEndpointsForOneServerlessInstance**](ServerlessPrivateEndpointsApi.md#ReturnAllPrivateEndpointsForOneServerlessInstance) | **Get** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Return All Private Endpoints for One Serverless Instance
[**ReturnOnePrivateEndpointForOneServerlessInstance**](ServerlessPrivateEndpointsApi.md#ReturnOnePrivateEndpointForOneServerlessInstance) | **Get** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Return One Private Endpoint for One Serverless Instance
[**UpdateOnePrivateEndpointForOneServerlessInstance**](ServerlessPrivateEndpointsApi.md#UpdateOnePrivateEndpointForOneServerlessInstance) | **Patch** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Update One Private Endpoint for One Serverless Instance



## CreateOnePrivateEndpointForOneServerlessInstance

> CreateOnePrivateEndpointForOneServerlessInstance201Response CreateOnePrivateEndpointForOneServerlessInstance(ctx, groupId, instanceName).Envelope(envelope).ApiAtlasServerlessTenantEndpointCreateView(apiAtlasServerlessTenantEndpointCreateView).Execute()

Create One Private Endpoint for One Serverless Instance



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
    instanceName := "instanceName_example" // string | Human-readable label that identifies the serverless instance for which the tenant endpoint will be created.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    apiAtlasServerlessTenantEndpointCreateView := *openapiclient.NewApiAtlasServerlessTenantEndpointCreateView() // ApiAtlasServerlessTenantEndpointCreateView |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerlessPrivateEndpointsApi.CreateOnePrivateEndpointForOneServerlessInstance(context.Background(), groupId, instanceName).Envelope(envelope).ApiAtlasServerlessTenantEndpointCreateView(apiAtlasServerlessTenantEndpointCreateView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessPrivateEndpointsApi.CreateOnePrivateEndpointForOneServerlessInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOnePrivateEndpointForOneServerlessInstance`: CreateOnePrivateEndpointForOneServerlessInstance201Response
    fmt.Fprintf(os.Stdout, "Response from `ServerlessPrivateEndpointsApi.CreateOnePrivateEndpointForOneServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**instanceName** | **string** | Human-readable label that identifies the serverless instance for which the tenant endpoint will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOnePrivateEndpointForOneServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **apiAtlasServerlessTenantEndpointCreateView** | [**ApiAtlasServerlessTenantEndpointCreateView**](ApiAtlasServerlessTenantEndpointCreateView.md) |  | 

### Return type

[**CreateOnePrivateEndpointForOneServerlessInstance201Response**](CreateOnePrivateEndpointForOneServerlessInstance201Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOnePrivateEndpointFromOneServerlessInstance

> RemoveOnePrivateEndpointFromOneServerlessInstance(ctx, groupId, instanceName, endpointId).Envelope(envelope).Execute()

Remove One Private Endpoint for One Serverless Instance



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
    instanceName := "instanceName_example" // string | Human-readable label that identifies the serverless instance from which the tenant endpoint will be removed.
    endpointId := "endpointId_example" // string | Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be removed.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerlessPrivateEndpointsApi.RemoveOnePrivateEndpointFromOneServerlessInstance(context.Background(), groupId, instanceName, endpointId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessPrivateEndpointsApi.RemoveOnePrivateEndpointFromOneServerlessInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**instanceName** | **string** | Human-readable label that identifies the serverless instance from which the tenant endpoint will be removed. | 
**endpointId** | **string** | Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOnePrivateEndpointFromOneServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

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


## ReturnAllPrivateEndpointsForOneServerlessInstance

> []ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner ReturnAllPrivateEndpointsForOneServerlessInstance(ctx, groupId, instanceName).Envelope(envelope).Execute()

Return All Private Endpoints for One Serverless Instance



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
    instanceName := "instanceName_example" // string | Human-readable label that identifies the serverless instance associated with the tenant endpoint.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerlessPrivateEndpointsApi.ReturnAllPrivateEndpointsForOneServerlessInstance(context.Background(), groupId, instanceName).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessPrivateEndpointsApi.ReturnAllPrivateEndpointsForOneServerlessInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllPrivateEndpointsForOneServerlessInstance`: []ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ServerlessPrivateEndpointsApi.ReturnAllPrivateEndpointsForOneServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**instanceName** | **string** | Human-readable label that identifies the serverless instance associated with the tenant endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllPrivateEndpointsForOneServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**[]ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner**](ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOnePrivateEndpointForOneServerlessInstance

> CreateOnePrivateEndpointForOneServerlessInstance201Response ReturnOnePrivateEndpointForOneServerlessInstance(ctx, groupId, instanceName, endpointId).Envelope(envelope).Execute()

Return One Private Endpoint for One Serverless Instance



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
    instanceName := "instanceName_example" // string | Human-readable label that identifies the serverless instance associated with the tenant endpoint.
    endpointId := "endpointId_example" // string | Unique 24-hexadecimal digit string that identifies the tenant endpoint.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerlessPrivateEndpointsApi.ReturnOnePrivateEndpointForOneServerlessInstance(context.Background(), groupId, instanceName, endpointId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessPrivateEndpointsApi.ReturnOnePrivateEndpointForOneServerlessInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOnePrivateEndpointForOneServerlessInstance`: CreateOnePrivateEndpointForOneServerlessInstance201Response
    fmt.Fprintf(os.Stdout, "Response from `ServerlessPrivateEndpointsApi.ReturnOnePrivateEndpointForOneServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**instanceName** | **string** | Human-readable label that identifies the serverless instance associated with the tenant endpoint. | 
**endpointId** | **string** | Unique 24-hexadecimal digit string that identifies the tenant endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOnePrivateEndpointForOneServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**CreateOnePrivateEndpointForOneServerlessInstance201Response**](CreateOnePrivateEndpointForOneServerlessInstance201Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOnePrivateEndpointForOneServerlessInstance

> CreateOnePrivateEndpointForOneServerlessInstance201Response UpdateOnePrivateEndpointForOneServerlessInstance(ctx, groupId, instanceName, endpointId).Envelope(envelope).ApiAtlasServerlessTenantEndpointUpdateView(apiAtlasServerlessTenantEndpointUpdateView).Execute()

Update One Private Endpoint for One Serverless Instance



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
    instanceName := "instanceName_example" // string | Human-readable label that identifies the serverless instance associated with the tenant endpoint that will be updated.
    endpointId := "endpointId_example" // string | Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be updated.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    apiAtlasServerlessTenantEndpointUpdateView := *openapiclient.NewApiAtlasServerlessTenantEndpointUpdateView("ProviderName_example") // ApiAtlasServerlessTenantEndpointUpdateView |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerlessPrivateEndpointsApi.UpdateOnePrivateEndpointForOneServerlessInstance(context.Background(), groupId, instanceName, endpointId).Envelope(envelope).ApiAtlasServerlessTenantEndpointUpdateView(apiAtlasServerlessTenantEndpointUpdateView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessPrivateEndpointsApi.UpdateOnePrivateEndpointForOneServerlessInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOnePrivateEndpointForOneServerlessInstance`: CreateOnePrivateEndpointForOneServerlessInstance201Response
    fmt.Fprintf(os.Stdout, "Response from `ServerlessPrivateEndpointsApi.UpdateOnePrivateEndpointForOneServerlessInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**instanceName** | **string** | Human-readable label that identifies the serverless instance associated with the tenant endpoint that will be updated. | 
**endpointId** | **string** | Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOnePrivateEndpointForOneServerlessInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **apiAtlasServerlessTenantEndpointUpdateView** | [**ApiAtlasServerlessTenantEndpointUpdateView**](ApiAtlasServerlessTenantEndpointUpdateView.md) |  | 

### Return type

[**CreateOnePrivateEndpointForOneServerlessInstance201Response**](CreateOnePrivateEndpointForOneServerlessInstance201Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

