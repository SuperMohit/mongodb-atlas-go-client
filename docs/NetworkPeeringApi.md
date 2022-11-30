# \NetworkPeeringApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableConnectViaPeeringOnlyModeForOneProject**](NetworkPeeringApi.md#DisableConnectViaPeeringOnlyModeForOneProject) | **Patch** /api/atlas/v1.0/groups/{groupId}/privateIpMode | Disable Connect via Peering Only Mode for One Project
[**VerifyConnectViaPeeringOnlyModeForOneProject**](NetworkPeeringApi.md#VerifyConnectViaPeeringOnlyModeForOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/privateIpMode | Verify Connect via Peering Only Mode for One Project



## DisableConnectViaPeeringOnlyModeForOneProject

> PrivateIPModeView DisableConnectViaPeeringOnlyModeForOneProject(ctx, groupId).PrivateIPModeView(privateIPModeView).Envelope(envelope).Pretty(pretty).Execute()

Disable Connect via Peering Only Mode for One Project



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
    privateIPModeView := *openapiclient.NewPrivateIPModeView(false) // PrivateIPModeView | Disables Connect via Peering Only mode for the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkPeeringApi.DisableConnectViaPeeringOnlyModeForOneProject(context.Background(), groupId).PrivateIPModeView(privateIPModeView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.DisableConnectViaPeeringOnlyModeForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableConnectViaPeeringOnlyModeForOneProject`: PrivateIPModeView
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.DisableConnectViaPeeringOnlyModeForOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableConnectViaPeeringOnlyModeForOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateIPModeView** | [**PrivateIPModeView**](PrivateIPModeView.md) | Disables Connect via Peering Only mode for the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PrivateIPModeView**](PrivateIPModeView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyConnectViaPeeringOnlyModeForOneProject

> PrivateIPModeView VerifyConnectViaPeeringOnlyModeForOneProject(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Verify Connect via Peering Only Mode for One Project



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
    resp, r, err := apiClient.NetworkPeeringApi.VerifyConnectViaPeeringOnlyModeForOneProject(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkPeeringApi.VerifyConnectViaPeeringOnlyModeForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyConnectViaPeeringOnlyModeForOneProject`: PrivateIPModeView
    fmt.Fprintf(os.Stdout, "Response from `NetworkPeeringApi.VerifyConnectViaPeeringOnlyModeForOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyConnectViaPeeringOnlyModeForOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PrivateIPModeView**](PrivateIPModeView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

