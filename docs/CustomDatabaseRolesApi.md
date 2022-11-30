# \CustomDatabaseRolesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneCustomRole**](CustomDatabaseRolesApi.md#CreateOneCustomRole) | **Post** /api/atlas/v1.0/groups/{groupId}/customDBRoles/roles | Create One Custom Role
[**RemoveOneCustomRoleFromOneProject**](CustomDatabaseRolesApi.md#RemoveOneCustomRoleFromOneProject) | **Delete** /api/atlas/v1.0/groups/{groupId}/customDBRoles/roles/{roleName} | Remove One Custom Role from One Project
[**ReturnAllCustomRolesInOneProject**](CustomDatabaseRolesApi.md#ReturnAllCustomRolesInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/customDBRoles/roles | Return All Custom Roles in One Project
[**ReturnOneCustomRoleInOneProject**](CustomDatabaseRolesApi.md#ReturnOneCustomRoleInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/customDBRoles/roles/{roleName} | Return One Custom Role in One Project
[**UpdateOneCustomRoleInOneProject**](CustomDatabaseRolesApi.md#UpdateOneCustomRoleInOneProject) | **Patch** /api/atlas/v1.0/groups/{groupId}/customDBRoles/roles/{roleName} | Update One Custom Role in One Project



## CreateOneCustomRole

> ApiAtlasCustomDBRoleView CreateOneCustomRole(ctx, groupId).ApiAtlasCustomDBRoleView(apiAtlasCustomDBRoleView).Envelope(envelope).Pretty(pretty).Execute()

Create One Custom Role



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
    apiAtlasCustomDBRoleView := *openapiclient.NewApiAtlasCustomDBRoleView("RoleName_example") // ApiAtlasCustomDBRoleView | Creates one custom role in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDatabaseRolesApi.CreateOneCustomRole(context.Background(), groupId).ApiAtlasCustomDBRoleView(apiAtlasCustomDBRoleView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.CreateOneCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneCustomRole`: ApiAtlasCustomDBRoleView
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.CreateOneCustomRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasCustomDBRoleView** | [**ApiAtlasCustomDBRoleView**](ApiAtlasCustomDBRoleView.md) | Creates one custom role in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasCustomDBRoleView**](ApiAtlasCustomDBRoleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneCustomRoleFromOneProject

> RemoveOneCustomRoleFromOneProject(ctx, groupId, roleName).Envelope(envelope).Pretty(pretty).Execute()

Remove One Custom Role from One Project



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
    roleName := "roleName_example" // string | Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDatabaseRolesApi.RemoveOneCustomRoleFromOneProject(context.Background(), groupId, roleName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.RemoveOneCustomRoleFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**roleName** | **string** | Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneCustomRoleFromOneProjectRequest struct via the builder pattern


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


## ReturnAllCustomRolesInOneProject

> []ApiAtlasCustomDBRoleView ReturnAllCustomRolesInOneProject(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Return All Custom Roles in One Project



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
    resp, r, err := apiClient.CustomDatabaseRolesApi.ReturnAllCustomRolesInOneProject(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.ReturnAllCustomRolesInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllCustomRolesInOneProject`: []ApiAtlasCustomDBRoleView
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.ReturnAllCustomRolesInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllCustomRolesInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiAtlasCustomDBRoleView**](ApiAtlasCustomDBRoleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneCustomRoleInOneProject

> ApiAtlasCustomDBRoleView ReturnOneCustomRoleInOneProject(ctx, groupId, roleName).Envelope(envelope).Pretty(pretty).Execute()

Return One Custom Role in One Project



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
    roleName := "roleName_example" // string | Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDatabaseRolesApi.ReturnOneCustomRoleInOneProject(context.Background(), groupId, roleName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.ReturnOneCustomRoleInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneCustomRoleInOneProject`: ApiAtlasCustomDBRoleView
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.ReturnOneCustomRoleInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**roleName** | **string** | Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneCustomRoleInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasCustomDBRoleView**](ApiAtlasCustomDBRoleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneCustomRoleInOneProject

> ApiAtlasCustomDBRoleView UpdateOneCustomRoleInOneProject(ctx, groupId, roleName).ApiAtlasUpdateCustomDBRoleView(apiAtlasUpdateCustomDBRoleView).Envelope(envelope).Pretty(pretty).Execute()

Update One Custom Role in One Project



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
    roleName := "roleName_example" // string | Human-readable label that identifies the role for the request. This name must beunique for this custom role in this project.
    apiAtlasUpdateCustomDBRoleView := *openapiclient.NewApiAtlasUpdateCustomDBRoleView() // ApiAtlasUpdateCustomDBRoleView | Updates one custom role in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDatabaseRolesApi.UpdateOneCustomRoleInOneProject(context.Background(), groupId, roleName).ApiAtlasUpdateCustomDBRoleView(apiAtlasUpdateCustomDBRoleView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDatabaseRolesApi.UpdateOneCustomRoleInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneCustomRoleInOneProject`: ApiAtlasCustomDBRoleView
    fmt.Fprintf(os.Stdout, "Response from `CustomDatabaseRolesApi.UpdateOneCustomRoleInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**roleName** | **string** | Human-readable label that identifies the role for the request. This name must beunique for this custom role in this project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneCustomRoleInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasUpdateCustomDBRoleView** | [**ApiAtlasUpdateCustomDBRoleView**](ApiAtlasUpdateCustomDBRoleView.md) | Updates one custom role in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasCustomDBRoleView**](ApiAtlasCustomDBRoleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

