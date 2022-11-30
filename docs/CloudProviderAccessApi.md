# \CloudProviderAccessApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeOneCloudProviderAccessRole**](CloudProviderAccessApi.md#AuthorizeOneCloudProviderAccessRole) | **Patch** /api/atlas/v1.0/groups/{groupId}/cloudProviderAccess/{roleId} | Authorize One Cloud Provider Access Role
[**CreateOneCloudProviderAccessRole**](CloudProviderAccessApi.md#CreateOneCloudProviderAccessRole) | **Post** /api/atlas/v1.0/groups/{groupId}/cloudProviderAccess | Create One Cloud Provider Access Role
[**DeauthorizeOneCloudProviderAccessRole**](CloudProviderAccessApi.md#DeauthorizeOneCloudProviderAccessRole) | **Delete** /api/atlas/v1.0/groups/{groupId}/cloudProviderAccess/{cloudProvider}/{roleId} | Deauthorize One Cloud Provider Access Role
[**ReturnAllCloudProviderAccessRoles**](CloudProviderAccessApi.md#ReturnAllCloudProviderAccessRoles) | **Get** /api/atlas/v1.0/groups/{groupId}/cloudProviderAccess | Return All Cloud Provider Access Roles
[**ReturnCloudProviderAccessRole**](CloudProviderAccessApi.md#ReturnCloudProviderAccessRole) | **Get** /api/atlas/v1.0/groups/{groupId}/cloudProviderAccess/{roleId} | Return specified Cloud Provider Access Role



## AuthorizeOneCloudProviderAccessRole

> ApiAtlasCloudProviderAccessRoleView AuthorizeOneCloudProviderAccessRole(ctx, groupId, roleId).ApiAtlasCloudProviderAccessRoleView(apiAtlasCloudProviderAccessRoleView).Envelope(envelope).Pretty(pretty).Execute()

Authorize One Cloud Provider Access Role



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
    roleId := "roleId_example" // string | Unique 24-hexadecimal digit string that identifies the role.
    apiAtlasCloudProviderAccessRoleView := *openapiclient.NewApiAtlasCloudProviderAccessRoleView("ProviderName_example") // ApiAtlasCloudProviderAccessRoleView | Grants access to the specified project for the specified AWS IAM role.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderAccessApi.AuthorizeOneCloudProviderAccessRole(context.Background(), groupId, roleId).ApiAtlasCloudProviderAccessRoleView(apiAtlasCloudProviderAccessRoleView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.AuthorizeOneCloudProviderAccessRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizeOneCloudProviderAccessRole`: ApiAtlasCloudProviderAccessRoleView
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.AuthorizeOneCloudProviderAccessRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**roleId** | **string** | Unique 24-hexadecimal digit string that identifies the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeOneCloudProviderAccessRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasCloudProviderAccessRoleView** | [**ApiAtlasCloudProviderAccessRoleView**](ApiAtlasCloudProviderAccessRoleView.md) | Grants access to the specified project for the specified AWS IAM role. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasCloudProviderAccessRoleView**](ApiAtlasCloudProviderAccessRoleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOneCloudProviderAccessRole

> ApiAtlasCloudProviderAccessRoleView CreateOneCloudProviderAccessRole(ctx, groupId).ApiAtlasCloudProviderAccessRoleView(apiAtlasCloudProviderAccessRoleView).Envelope(envelope).Pretty(pretty).Execute()

Create One Cloud Provider Access Role



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
    apiAtlasCloudProviderAccessRoleView := *openapiclient.NewApiAtlasCloudProviderAccessRoleView("ProviderName_example") // ApiAtlasCloudProviderAccessRoleView | Creates one AWS IAM role.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderAccessApi.CreateOneCloudProviderAccessRole(context.Background(), groupId).ApiAtlasCloudProviderAccessRoleView(apiAtlasCloudProviderAccessRoleView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.CreateOneCloudProviderAccessRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneCloudProviderAccessRole`: ApiAtlasCloudProviderAccessRoleView
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.CreateOneCloudProviderAccessRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneCloudProviderAccessRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasCloudProviderAccessRoleView** | [**ApiAtlasCloudProviderAccessRoleView**](ApiAtlasCloudProviderAccessRoleView.md) | Creates one AWS IAM role. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasCloudProviderAccessRoleView**](ApiAtlasCloudProviderAccessRoleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeauthorizeOneCloudProviderAccessRole

> DeauthorizeOneCloudProviderAccessRole(ctx, groupId, cloudProvider, roleId).Envelope(envelope).Pretty(pretty).Execute()

Deauthorize One Cloud Provider Access Role



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
    cloudProvider := "cloudProvider_example" // string | Human-readable label that identifies the cloud provider of the role to deauthorize.
    roleId := "roleId_example" // string | Unique 24-hexadecimal digit string that identifies the role.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderAccessApi.DeauthorizeOneCloudProviderAccessRole(context.Background(), groupId, cloudProvider, roleId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.DeauthorizeOneCloudProviderAccessRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**cloudProvider** | **string** | Human-readable label that identifies the cloud provider of the role to deauthorize. | 
**roleId** | **string** | Unique 24-hexadecimal digit string that identifies the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeauthorizeOneCloudProviderAccessRoleRequest struct via the builder pattern


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


## ReturnAllCloudProviderAccessRoles

> ApiAtlasCloudProviderAccessView ReturnAllCloudProviderAccessRoles(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Return All Cloud Provider Access Roles



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
    resp, r, err := apiClient.CloudProviderAccessApi.ReturnAllCloudProviderAccessRoles(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.ReturnAllCloudProviderAccessRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllCloudProviderAccessRoles`: ApiAtlasCloudProviderAccessView
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.ReturnAllCloudProviderAccessRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllCloudProviderAccessRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasCloudProviderAccessView**](ApiAtlasCloudProviderAccessView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnCloudProviderAccessRole

> ApiAtlasCloudProviderAccessView ReturnCloudProviderAccessRole(ctx, groupId, roleId).Envelope(envelope).Pretty(pretty).Execute()

Return specified Cloud Provider Access Role



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
    roleId := "roleId_example" // string | Unique 24-hexadecimal digit string that identifies the role.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderAccessApi.ReturnCloudProviderAccessRole(context.Background(), groupId, roleId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAccessApi.ReturnCloudProviderAccessRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnCloudProviderAccessRole`: ApiAtlasCloudProviderAccessView
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAccessApi.ReturnCloudProviderAccessRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**roleId** | **string** | Unique 24-hexadecimal digit string that identifies the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnCloudProviderAccessRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasCloudProviderAccessView**](ApiAtlasCloudProviderAccessView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

