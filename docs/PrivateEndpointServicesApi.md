# \PrivateEndpointServicesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOnePrivateEndpointForOneProvider**](PrivateEndpointServicesApi.md#CreateOnePrivateEndpointForOneProvider) | **Post** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint | Create One Private Endpoint for One Provider
[**CreateOnePrivateEndpointServiceForOneProvider**](PrivateEndpointServicesApi.md#CreateOnePrivateEndpointServiceForOneProvider) | **Post** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/endpointService | Create One Private Endpoint Service for One Provider
[**RemoveOnePrivateEndpointForOneProvider**](PrivateEndpointServicesApi.md#RemoveOnePrivateEndpointForOneProvider) | **Delete** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Remove One Private Endpoint for One Provider
[**RemoveOnePrivateEndpointServiceForOneProvider**](PrivateEndpointServicesApi.md#RemoveOnePrivateEndpointServiceForOneProvider) | **Delete** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Remove One Private Endpoint Service for One Provider
[**ReturnAllPrivateEndpointServicesForOneProvider**](PrivateEndpointServicesApi.md#ReturnAllPrivateEndpointServicesForOneProvider) | **Get** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService | Return All Private Endpoint Services for One Provider
[**ReturnOnePrivateEndpointForOneProvider**](PrivateEndpointServicesApi.md#ReturnOnePrivateEndpointForOneProvider) | **Get** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Return One Private Endpoint for One Provider
[**ReturnOnePrivateEndpointServiceForOneProvider**](PrivateEndpointServicesApi.md#ReturnOnePrivateEndpointServiceForOneProvider) | **Get** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Return One Private Endpoint Service for One Provider
[**ReturnRegionalizedPrivateEndpointStatus**](PrivateEndpointServicesApi.md#ReturnRegionalizedPrivateEndpointStatus) | **Get** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/regionalMode | Return Regionalized Private Endpoint Status
[**ToggleRegionalizedPrivateEndpointStatus**](PrivateEndpointServicesApi.md#ToggleRegionalizedPrivateEndpointStatus) | **Patch** /api/atlas/v1.0/groups/{groupId}/privateEndpoint/regionalMode | Toggle Regionalized Private Endpoint Status



## CreateOnePrivateEndpointForOneProvider

> CreateOnePrivateEndpointForOneProvider200Response CreateOnePrivateEndpointForOneProvider(ctx, groupId, cloudProvider, endpointServiceId).CreateOnePrivateEndpointForOneProviderRequest(createOnePrivateEndpointForOneProviderRequest).Envelope(envelope).Pretty(pretty).Execute()

Create One Private Endpoint for One Provider



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
    cloudProvider := "cloudProvider_example" // string | Cloud service provider that manages this private endpoint. (default to "AWS")
    endpointServiceId := "endpointServiceId_example" // string | Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint.
    createOnePrivateEndpointForOneProviderRequest := openapiclient.createOnePrivateEndpointForOneProvider_request{ApiAtlasCreateAWSEndpointRequestView: openapiclient.NewApiAtlasCreateAWSEndpointRequestView()} // CreateOnePrivateEndpointForOneProviderRequest | Creates one private resource endpoint for the specified cloud service provider.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateEndpointServicesApi.CreateOnePrivateEndpointForOneProvider(context.Background(), groupId, cloudProvider, endpointServiceId).CreateOnePrivateEndpointForOneProviderRequest(createOnePrivateEndpointForOneProviderRequest).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.CreateOnePrivateEndpointForOneProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOnePrivateEndpointForOneProvider`: CreateOnePrivateEndpointForOneProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.CreateOnePrivateEndpointForOneProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint. | [default to &quot;AWS&quot;]
**endpointServiceId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOnePrivateEndpointForOneProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createOnePrivateEndpointForOneProviderRequest** | [**CreateOnePrivateEndpointForOneProviderRequest**](CreateOnePrivateEndpointForOneProviderRequest.md) | Creates one private resource endpoint for the specified cloud service provider. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**CreateOnePrivateEndpointForOneProvider200Response**](CreateOnePrivateEndpointForOneProvider200Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOnePrivateEndpointServiceForOneProvider

> CreateOnePrivateEndpointServiceForOneProvider200Response CreateOnePrivateEndpointServiceForOneProvider(ctx, groupId).ApiAtlasCreateEndpointServiceRequestView(apiAtlasCreateEndpointServiceRequestView).Envelope(envelope).Pretty(pretty).Execute()

Create One Private Endpoint Service for One Provider



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
    apiAtlasCreateEndpointServiceRequestView := *openapiclient.NewApiAtlasCreateEndpointServiceRequestView("ProviderName_example", "Region_example") // ApiAtlasCreateEndpointServiceRequestView | Creates one private resource service for the specified cloud service provider.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateEndpointServicesApi.CreateOnePrivateEndpointServiceForOneProvider(context.Background(), groupId).ApiAtlasCreateEndpointServiceRequestView(apiAtlasCreateEndpointServiceRequestView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.CreateOnePrivateEndpointServiceForOneProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOnePrivateEndpointServiceForOneProvider`: CreateOnePrivateEndpointServiceForOneProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.CreateOnePrivateEndpointServiceForOneProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOnePrivateEndpointServiceForOneProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasCreateEndpointServiceRequestView** | [**ApiAtlasCreateEndpointServiceRequestView**](ApiAtlasCreateEndpointServiceRequestView.md) | Creates one private resource service for the specified cloud service provider. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**CreateOnePrivateEndpointServiceForOneProvider200Response**](CreateOnePrivateEndpointServiceForOneProvider200Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOnePrivateEndpointForOneProvider

> RemoveOnePrivateEndpointForOneProvider(ctx, groupId, cloudProvider, endpointId, endpointServiceId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Private Endpoint for One Provider



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
    cloudProvider := "cloudProvider_example" // string | Cloud service provider that manages this private endpoint. (default to "AWS")
    endpointId := "endpointId_example" // string | Unique string that identifies the private endpoint you want to delete. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.
    endpointServiceId := "endpointServiceId_example" // string | Unique 24-hexadecimal digit string that identifies the private endpoint service from which you want to delete a private endpoint.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateEndpointServicesApi.RemoveOnePrivateEndpointForOneProvider(context.Background(), groupId, cloudProvider, endpointId, endpointServiceId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.RemoveOnePrivateEndpointForOneProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint. | [default to &quot;AWS&quot;]
**endpointId** | **string** | Unique string that identifies the private endpoint you want to delete. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints. | 
**endpointServiceId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint service from which you want to delete a private endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOnePrivateEndpointForOneProviderRequest struct via the builder pattern


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


## RemoveOnePrivateEndpointServiceForOneProvider

> RemoveOnePrivateEndpointServiceForOneProvider(ctx, groupId, cloudProvider, endpointServiceId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Private Endpoint Service for One Provider



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
    cloudProvider := "cloudProvider_example" // string | Cloud service provider that manages this private endpoint service. (default to "AWS")
    endpointServiceId := "endpointServiceId_example" // string | Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to delete.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateEndpointServicesApi.RemoveOnePrivateEndpointServiceForOneProvider(context.Background(), groupId, cloudProvider, endpointServiceId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.RemoveOnePrivateEndpointServiceForOneProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint service. | [default to &quot;AWS&quot;]
**endpointServiceId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOnePrivateEndpointServiceForOneProviderRequest struct via the builder pattern


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


## ReturnAllPrivateEndpointServicesForOneProvider

> ReturnAllPrivateEndpointServicesForOneProvider200Response ReturnAllPrivateEndpointServicesForOneProvider(ctx, groupId, cloudProvider).Envelope(envelope).Pretty(pretty).Execute()

Return All Private Endpoint Services for One Provider



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
    cloudProvider := "cloudProvider_example" // string | Cloud service provider that manages this private endpoint service. (default to "AWS")
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateEndpointServicesApi.ReturnAllPrivateEndpointServicesForOneProvider(context.Background(), groupId, cloudProvider).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ReturnAllPrivateEndpointServicesForOneProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllPrivateEndpointServicesForOneProvider`: ReturnAllPrivateEndpointServicesForOneProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ReturnAllPrivateEndpointServicesForOneProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint service. | [default to &quot;AWS&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllPrivateEndpointServicesForOneProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ReturnAllPrivateEndpointServicesForOneProvider200Response**](ReturnAllPrivateEndpointServicesForOneProvider200Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOnePrivateEndpointForOneProvider

> ReturnOnePrivateEndpointForOneProvider200Response ReturnOnePrivateEndpointForOneProvider(ctx, groupId, cloudProvider, endpointId, endpointServiceId).Envelope(envelope).Pretty(pretty).Execute()

Return One Private Endpoint for One Provider



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
    cloudProvider := "cloudProvider_example" // string | Cloud service provider that manages this private endpoint. (default to "AWS")
    endpointId := "endpointId_example" // string | Unique string that identifies the private endpoint you want to return. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.
    endpointServiceId := "endpointServiceId_example" // string | Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to return a private endpoint.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateEndpointServicesApi.ReturnOnePrivateEndpointForOneProvider(context.Background(), groupId, cloudProvider, endpointId, endpointServiceId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ReturnOnePrivateEndpointForOneProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOnePrivateEndpointForOneProvider`: ReturnOnePrivateEndpointForOneProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ReturnOnePrivateEndpointForOneProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint. | [default to &quot;AWS&quot;]
**endpointId** | **string** | Unique string that identifies the private endpoint you want to return. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints. | 
**endpointServiceId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to return a private endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOnePrivateEndpointForOneProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ReturnOnePrivateEndpointForOneProvider200Response**](ReturnOnePrivateEndpointForOneProvider200Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOnePrivateEndpointServiceForOneProvider

> CreateOnePrivateEndpointServiceForOneProvider200Response ReturnOnePrivateEndpointServiceForOneProvider(ctx, groupId, cloudProvider, endpointServiceId).Envelope(envelope).Pretty(pretty).Execute()

Return One Private Endpoint Service for One Provider



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
    cloudProvider := "cloudProvider_example" // string | Cloud service provider that manages this private endpoint service. (default to "AWS")
    endpointServiceId := "endpointServiceId_example" // string | Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateEndpointServicesApi.ReturnOnePrivateEndpointServiceForOneProvider(context.Background(), groupId, cloudProvider, endpointServiceId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ReturnOnePrivateEndpointServiceForOneProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOnePrivateEndpointServiceForOneProvider`: CreateOnePrivateEndpointServiceForOneProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ReturnOnePrivateEndpointServiceForOneProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**cloudProvider** | **string** | Cloud service provider that manages this private endpoint service. | [default to &quot;AWS&quot;]
**endpointServiceId** | **string** | Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOnePrivateEndpointServiceForOneProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**CreateOnePrivateEndpointServiceForOneProvider200Response**](CreateOnePrivateEndpointServiceForOneProvider200Response.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnRegionalizedPrivateEndpointStatus

> ProjectSettingItemView ReturnRegionalizedPrivateEndpointStatus(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Return Regionalized Private Endpoint Status



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
    resp, r, err := apiClient.PrivateEndpointServicesApi.ReturnRegionalizedPrivateEndpointStatus(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ReturnRegionalizedPrivateEndpointStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnRegionalizedPrivateEndpointStatus`: ProjectSettingItemView
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ReturnRegionalizedPrivateEndpointStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnRegionalizedPrivateEndpointStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ProjectSettingItemView**](ProjectSettingItemView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleRegionalizedPrivateEndpointStatus

> ProjectSettingItemView ToggleRegionalizedPrivateEndpointStatus(ctx, groupId).ProjectSettingItemView(projectSettingItemView).Envelope(envelope).Pretty(pretty).Execute()

Toggle Regionalized Private Endpoint Status



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
    projectSettingItemView := *openapiclient.NewProjectSettingItemView(false) // ProjectSettingItemView | Enables or disables the ability can create multiple private resources per region in all cloud service providers in one project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateEndpointServicesApi.ToggleRegionalizedPrivateEndpointStatus(context.Background(), groupId).ProjectSettingItemView(projectSettingItemView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ToggleRegionalizedPrivateEndpointStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleRegionalizedPrivateEndpointStatus`: ProjectSettingItemView
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ToggleRegionalizedPrivateEndpointStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleRegionalizedPrivateEndpointStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectSettingItemView** | [**ProjectSettingItemView**](ProjectSettingItemView.md) | Enables or disables the ability can create multiple private resources per region in all cloud service providers in one project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ProjectSettingItemView**](ProjectSettingItemView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

