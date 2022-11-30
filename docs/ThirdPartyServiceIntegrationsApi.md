# \ThirdPartyServiceIntegrationsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureOneThirdPartyServiceIntegration**](ThirdPartyServiceIntegrationsApi.md#ConfigureOneThirdPartyServiceIntegration) | **Post** /api/atlas/v1.0/groups/{groupId}/integrations/{integrationType} | Configure One Third-Party Service Integration
[**RemoveOneThirdPartyServiceIntegration**](ThirdPartyServiceIntegrationsApi.md#RemoveOneThirdPartyServiceIntegration) | **Delete** /api/atlas/v1.0/groups/{groupId}/integrations/{integrationType} | Remove One Third-Party Service Integration
[**ReturnAllActiveThirdPartyServiceIntegrations**](ThirdPartyServiceIntegrationsApi.md#ReturnAllActiveThirdPartyServiceIntegrations) | **Get** /api/atlas/v1.0/groups/{groupId}/integrations | Return All Active Third-Party Service Integrations
[**ReturnOneThirdPartyServiceIntegration**](ThirdPartyServiceIntegrationsApi.md#ReturnOneThirdPartyServiceIntegration) | **Get** /api/atlas/v1.0/groups/{groupId}/integrations/{integrationType} | Return One Third-Party Service Integration
[**UpdateOneThirdPartyServiceIntegration**](ThirdPartyServiceIntegrationsApi.md#UpdateOneThirdPartyServiceIntegration) | **Put** /api/atlas/v1.0/groups/{groupId}/integrations/{integrationType} | Update One Third-Party Service Integration



## ConfigureOneThirdPartyServiceIntegration

> PaginatedIntegrationView ConfigureOneThirdPartyServiceIntegration(ctx, integrationType, groupId).ApiIntegrationViewManual(apiIntegrationViewManual).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Configure One Third-Party Service Integration



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
    integrationType := "integrationType_example" // string | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    apiIntegrationViewManual := *openapiclient.NewApiIntegrationViewManual() // ApiIntegrationViewManual | Third-party integration configuration that you want to configure for your project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThirdPartyServiceIntegrationsApi.ConfigureOneThirdPartyServiceIntegration(context.Background(), integrationType, groupId).ApiIntegrationViewManual(apiIntegrationViewManual).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyServiceIntegrationsApi.ConfigureOneThirdPartyServiceIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigureOneThirdPartyServiceIntegration`: PaginatedIntegrationView
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyServiceIntegrationsApi.ConfigureOneThirdPartyServiceIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationType** | **string** | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureOneThirdPartyServiceIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiIntegrationViewManual** | [**ApiIntegrationViewManual**](ApiIntegrationViewManual.md) | Third-party integration configuration that you want to configure for your project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedIntegrationView**](PaginatedIntegrationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneThirdPartyServiceIntegration

> RemoveOneThirdPartyServiceIntegration(ctx, integrationType, groupId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Third-Party Service Integration



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
    integrationType := "integrationType_example" // string | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThirdPartyServiceIntegrationsApi.RemoveOneThirdPartyServiceIntegration(context.Background(), integrationType, groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyServiceIntegrationsApi.RemoveOneThirdPartyServiceIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationType** | **string** | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneThirdPartyServiceIntegrationRequest struct via the builder pattern


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


## ReturnAllActiveThirdPartyServiceIntegrations

> PaginatedIntegrationView ReturnAllActiveThirdPartyServiceIntegrations(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Active Third-Party Service Integrations



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
    resp, r, err := apiClient.ThirdPartyServiceIntegrationsApi.ReturnAllActiveThirdPartyServiceIntegrations(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyServiceIntegrationsApi.ReturnAllActiveThirdPartyServiceIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllActiveThirdPartyServiceIntegrations`: PaginatedIntegrationView
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyServiceIntegrationsApi.ReturnAllActiveThirdPartyServiceIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllActiveThirdPartyServiceIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedIntegrationView**](PaginatedIntegrationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneThirdPartyServiceIntegration

> ApiIntegrationViewManual ReturnOneThirdPartyServiceIntegration(ctx, groupId, integrationType).Envelope(envelope).Pretty(pretty).Execute()

Return One Third-Party Service Integration



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
    integrationType := "integrationType_example" // string | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThirdPartyServiceIntegrationsApi.ReturnOneThirdPartyServiceIntegration(context.Background(), groupId, integrationType).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyServiceIntegrationsApi.ReturnOneThirdPartyServiceIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneThirdPartyServiceIntegration`: ApiIntegrationViewManual
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyServiceIntegrationsApi.ReturnOneThirdPartyServiceIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**integrationType** | **string** | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneThirdPartyServiceIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiIntegrationViewManual**](ApiIntegrationViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneThirdPartyServiceIntegration

> PaginatedIntegrationView UpdateOneThirdPartyServiceIntegration(ctx, integrationType, groupId).ApiIntegrationViewManual(apiIntegrationViewManual).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Update One Third-Party Service Integration



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
    integrationType := "integrationType_example" // string | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    apiIntegrationViewManual := *openapiclient.NewApiIntegrationViewManual() // ApiIntegrationViewManual | Third-party integration configuration that you want to configure for your project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThirdPartyServiceIntegrationsApi.UpdateOneThirdPartyServiceIntegration(context.Background(), integrationType, groupId).ApiIntegrationViewManual(apiIntegrationViewManual).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyServiceIntegrationsApi.UpdateOneThirdPartyServiceIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneThirdPartyServiceIntegration`: PaginatedIntegrationView
    fmt.Fprintf(os.Stdout, "Response from `ThirdPartyServiceIntegrationsApi.UpdateOneThirdPartyServiceIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationType** | **string** | Human-readable label that identifies the service which you want to integrate with MongoDB Cloud. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneThirdPartyServiceIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiIntegrationViewManual** | [**ApiIntegrationViewManual**](ApiIntegrationViewManual.md) | Third-party integration configuration that you want to configure for your project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedIntegrationView**](PaginatedIntegrationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

