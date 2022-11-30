# \MaintenanceWindowsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomateDeferralOfMaintenanceForOneProject**](MaintenanceWindowsApi.md#AutomateDeferralOfMaintenanceForOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/maintenanceWindow/autoDefer | Toggle Automatic Deferral of Maintenance for One Project
[**DeferOneMaintenanceWindowForOneProject**](MaintenanceWindowsApi.md#DeferOneMaintenanceWindowForOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/maintenanceWindow/defer | Defer One Maintenance Window for One Project
[**ResetOneMaintenanceWindowForOneProject**](MaintenanceWindowsApi.md#ResetOneMaintenanceWindowForOneProject) | **Delete** /api/atlas/v1.0/groups/{groupId}/maintenanceWindow | Reset One Maintenance Window for One Project
[**ReturnOneMaintenanceWindowForOneProject**](MaintenanceWindowsApi.md#ReturnOneMaintenanceWindowForOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/maintenanceWindow | Return One Maintenance Window for One Project
[**UpdateMaintenanceWindowForOneProject**](MaintenanceWindowsApi.md#UpdateMaintenanceWindowForOneProject) | **Patch** /api/atlas/v1.0/groups/{groupId}/maintenanceWindow | Update Maintenance Window for One Project



## AutomateDeferralOfMaintenanceForOneProject

> AutomateDeferralOfMaintenanceForOneProject(ctx, groupId).Envelope(envelope).Execute()

Toggle Automatic Deferral of Maintenance for One Project



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceWindowsApi.AutomateDeferralOfMaintenanceForOneProject(context.Background(), groupId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.AutomateDeferralOfMaintenanceForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomateDeferralOfMaintenanceForOneProjectRequest struct via the builder pattern


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


## DeferOneMaintenanceWindowForOneProject

> DeferOneMaintenanceWindowForOneProject(ctx, groupId).Envelope(envelope).Execute()

Defer One Maintenance Window for One Project



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceWindowsApi.DeferOneMaintenanceWindowForOneProject(context.Background(), groupId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.DeferOneMaintenanceWindowForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeferOneMaintenanceWindowForOneProjectRequest struct via the builder pattern


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


## ResetOneMaintenanceWindowForOneProject

> ResetOneMaintenanceWindowForOneProject(ctx, groupId).Envelope(envelope).Execute()

Reset One Maintenance Window for One Project



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceWindowsApi.ResetOneMaintenanceWindowForOneProject(context.Background(), groupId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.ResetOneMaintenanceWindowForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetOneMaintenanceWindowForOneProjectRequest struct via the builder pattern


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


## ReturnOneMaintenanceWindowForOneProject

> ApiAtlasGroupMaintenanceWindowView ReturnOneMaintenanceWindowForOneProject(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Return One Maintenance Window for One Project



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
    resp, r, err := apiClient.MaintenanceWindowsApi.ReturnOneMaintenanceWindowForOneProject(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.ReturnOneMaintenanceWindowForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneMaintenanceWindowForOneProject`: ApiAtlasGroupMaintenanceWindowView
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.ReturnOneMaintenanceWindowForOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneMaintenanceWindowForOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasGroupMaintenanceWindowView**](ApiAtlasGroupMaintenanceWindowView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMaintenanceWindowForOneProject

> UpdateMaintenanceWindowForOneProject(ctx, groupId).ApiAtlasGroupMaintenanceWindowView(apiAtlasGroupMaintenanceWindowView).Envelope(envelope).Execute()

Update Maintenance Window for One Project



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
    apiAtlasGroupMaintenanceWindowView := *openapiclient.NewApiAtlasGroupMaintenanceWindowView(int32(123), int32(123)) // ApiAtlasGroupMaintenanceWindowView | Updates the maintenance window for the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceWindowsApi.UpdateMaintenanceWindowForOneProject(context.Background(), groupId).ApiAtlasGroupMaintenanceWindowView(apiAtlasGroupMaintenanceWindowView).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.UpdateMaintenanceWindowForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMaintenanceWindowForOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasGroupMaintenanceWindowView** | [**ApiAtlasGroupMaintenanceWindowView**](ApiAtlasGroupMaintenanceWindowView.md) | Updates the maintenance window for the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

 (empty response body)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

