# \ServerlessInstancesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneServerlessInstanceInOneProject**](ServerlessInstancesApi.md#CreateOneServerlessInstanceInOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/serverless | Create One Serverless Instance in One Project
[**RemoveOneServerlessInstanceFromOneProject**](ServerlessInstancesApi.md#RemoveOneServerlessInstanceFromOneProject) | **Delete** /api/atlas/v1.0/groups/{groupId}/serverless/{name} | Remove One Serverless Instance from One Project
[**ReturnAllServerlessInstancesFromOneProject**](ServerlessInstancesApi.md#ReturnAllServerlessInstancesFromOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/serverless | Return All Serverless Instances from One Project
[**ReturnOneServerlessInstanceFromOneProject**](ServerlessInstancesApi.md#ReturnOneServerlessInstanceFromOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/serverless/{name} | Return One Serverless Instance from One Project
[**UpdateOneServerlessInstanceInOneProject**](ServerlessInstancesApi.md#UpdateOneServerlessInstanceInOneProject) | **Patch** /api/atlas/v1.0/groups/{groupId}/serverless/{name} | Update One Serverless Instance in One Project



## CreateOneServerlessInstanceInOneProject

> ApiAtlasServerlessClusterDescriptionViewManual CreateOneServerlessInstanceInOneProject(ctx, groupId).ApiAtlasServerlessClusterDescriptionViewManual(apiAtlasServerlessClusterDescriptionViewManual).Envelope(envelope).Pretty(pretty).Execute()

Create One Serverless Instance in One Project



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
    apiAtlasServerlessClusterDescriptionViewManual := *openapiclient.NewApiAtlasServerlessClusterDescriptionViewManual() // ApiAtlasServerlessClusterDescriptionViewManual | Create One Serverless Instance in One Project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerlessInstancesApi.CreateOneServerlessInstanceInOneProject(context.Background(), groupId).ApiAtlasServerlessClusterDescriptionViewManual(apiAtlasServerlessClusterDescriptionViewManual).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessInstancesApi.CreateOneServerlessInstanceInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneServerlessInstanceInOneProject`: ApiAtlasServerlessClusterDescriptionViewManual
    fmt.Fprintf(os.Stdout, "Response from `ServerlessInstancesApi.CreateOneServerlessInstanceInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneServerlessInstanceInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasServerlessClusterDescriptionViewManual** | [**ApiAtlasServerlessClusterDescriptionViewManual**](ApiAtlasServerlessClusterDescriptionViewManual.md) | Create One Serverless Instance in One Project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasServerlessClusterDescriptionViewManual**](ApiAtlasServerlessClusterDescriptionViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneServerlessInstanceFromOneProject

> RemoveOneServerlessInstanceFromOneProject(ctx, groupId, name).Envelope(envelope).Pretty(pretty).Execute()

Remove One Serverless Instance from One Project



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
    name := "name_example" // string | Human-readable label that identifies the serverless instance.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerlessInstancesApi.RemoveOneServerlessInstanceFromOneProject(context.Background(), groupId, name).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessInstancesApi.RemoveOneServerlessInstanceFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**name** | **string** | Human-readable label that identifies the serverless instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneServerlessInstanceFromOneProjectRequest struct via the builder pattern


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


## ReturnAllServerlessInstancesFromOneProject

> PaginatedServerlessInstancesView ReturnAllServerlessInstancesFromOneProject(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Serverless Instances from One Project



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
    resp, r, err := apiClient.ServerlessInstancesApi.ReturnAllServerlessInstancesFromOneProject(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessInstancesApi.ReturnAllServerlessInstancesFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllServerlessInstancesFromOneProject`: PaginatedServerlessInstancesView
    fmt.Fprintf(os.Stdout, "Response from `ServerlessInstancesApi.ReturnAllServerlessInstancesFromOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllServerlessInstancesFromOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedServerlessInstancesView**](PaginatedServerlessInstancesView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneServerlessInstanceFromOneProject

> ApiAtlasServerlessClusterDescriptionViewManual ReturnOneServerlessInstanceFromOneProject(ctx, groupId, name).Envelope(envelope).Pretty(pretty).Execute()

Return One Serverless Instance from One Project



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
    name := "name_example" // string | Human-readable label that identifies the serverless instance.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerlessInstancesApi.ReturnOneServerlessInstanceFromOneProject(context.Background(), groupId, name).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessInstancesApi.ReturnOneServerlessInstanceFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneServerlessInstanceFromOneProject`: ApiAtlasServerlessClusterDescriptionViewManual
    fmt.Fprintf(os.Stdout, "Response from `ServerlessInstancesApi.ReturnOneServerlessInstanceFromOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**name** | **string** | Human-readable label that identifies the serverless instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneServerlessInstanceFromOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasServerlessClusterDescriptionViewManual**](ApiAtlasServerlessClusterDescriptionViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneServerlessInstanceInOneProject

> ApiAtlasServerlessClusterDescriptionPatchView UpdateOneServerlessInstanceInOneProject(ctx, groupId, name).ApiAtlasServerlessClusterDescriptionPatchView(apiAtlasServerlessClusterDescriptionPatchView).Envelope(envelope).Pretty(pretty).Execute()

Update One Serverless Instance in One Project



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
    name := "name_example" // string | Human-readable label that identifies the serverless instance.
    apiAtlasServerlessClusterDescriptionPatchView := *openapiclient.NewApiAtlasServerlessClusterDescriptionPatchView() // ApiAtlasServerlessClusterDescriptionPatchView | Update One Serverless Instance in One Project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerlessInstancesApi.UpdateOneServerlessInstanceInOneProject(context.Background(), groupId, name).ApiAtlasServerlessClusterDescriptionPatchView(apiAtlasServerlessClusterDescriptionPatchView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerlessInstancesApi.UpdateOneServerlessInstanceInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneServerlessInstanceInOneProject`: ApiAtlasServerlessClusterDescriptionPatchView
    fmt.Fprintf(os.Stdout, "Response from `ServerlessInstancesApi.UpdateOneServerlessInstanceInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**name** | **string** | Human-readable label that identifies the serverless instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneServerlessInstanceInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasServerlessClusterDescriptionPatchView** | [**ApiAtlasServerlessClusterDescriptionPatchView**](ApiAtlasServerlessClusterDescriptionPatchView.md) | Update One Serverless Instance in One Project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasServerlessClusterDescriptionPatchView**](ApiAtlasServerlessClusterDescriptionPatchView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

