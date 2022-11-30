# \FederatedAuthenticationApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoleMapping**](FederatedAuthenticationApi.md#AddRoleMapping) | **Post** /api/atlas/v1.0/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Add One Role Mapping to One Organization
[**DeleteFederationApp**](FederatedAuthenticationApi.md#DeleteFederationApp) | **Delete** /api/atlas/v1.0/federationSettings/{federationSettingsId} | Delete the federation settings instance.
[**DeleteRoleMapping**](FederatedAuthenticationApi.md#DeleteRoleMapping) | **Delete** /api/atlas/v1.0/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Remove One Role Mapping from One Organization
[**GetAllConnectedOrgConfigs**](FederatedAuthenticationApi.md#GetAllConnectedOrgConfigs) | **Get** /api/atlas/v1.0/federationSettings/{federationSettingsId}/connectedOrgConfigs | Return All Connected Org Configs from the Federation
[**GetAllIdentityProviders**](FederatedAuthenticationApi.md#GetAllIdentityProviders) | **Get** /api/atlas/v1.0/federationSettings/{federationSettingsId}/identityProviders | Return all identity providers from the specified federation.
[**GetAllRoleMappings**](FederatedAuthenticationApi.md#GetAllRoleMappings) | **Get** /api/atlas/v1.0/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Return All Role Mappings from One Organization
[**GetConnectedOrgConfig**](FederatedAuthenticationApi.md#GetConnectedOrgConfig) | **Get** /api/atlas/v1.0/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Return One Org Config Connected to One Federation
[**GetFederationSettings**](FederatedAuthenticationApi.md#GetFederationSettings) | **Get** /api/atlas/v1.0/orgs/{orgId}/federationSettings | Return Federation Settings for One Organization
[**GetIdentityProvider**](FederatedAuthenticationApi.md#GetIdentityProvider) | **Get** /api/atlas/v1.0/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Return one identity provider from the specified federation.
[**GetIdentityProviderMetadata**](FederatedAuthenticationApi.md#GetIdentityProviderMetadata) | **Get** /api/atlas/v1.0/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/metadata.xml | Return the metadata of one identity provider in the specified federation.
[**GetRoleMapping**](FederatedAuthenticationApi.md#GetRoleMapping) | **Get** /api/atlas/v1.0/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Return One Role Mapping from One Organization
[**RemoveConnectedOrgConfig**](FederatedAuthenticationApi.md#RemoveConnectedOrgConfig) | **Delete** /api/atlas/v1.0/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Remove One Org Config Connected to One Federation
[**UpdateConnectedOrgConfig**](FederatedAuthenticationApi.md#UpdateConnectedOrgConfig) | **Patch** /api/atlas/v1.0/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Update one connected organization configuration in the specified federation.
[**UpdateIdentityProvider**](FederatedAuthenticationApi.md#UpdateIdentityProvider) | **Patch** /api/atlas/v1.0/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Update the identity provider.
[**UpdateRoleMapping**](FederatedAuthenticationApi.md#UpdateRoleMapping) | **Put** /api/atlas/v1.0/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Update One Role Mapping in One Organization



## AddRoleMapping

> RoleMappingView AddRoleMapping(ctx, federationSettingsId, orgId).RoleMappingView(roleMappingView).Envelope(envelope).Execute()

Add One Role Mapping to One Organization



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    orgId := "4888442a3354817a7320eb61" // string | Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
    roleMappingView := *openapiclient.NewRoleMappingView("ExternalGroupName_example", []openapiclient.RoleAssignment{*openapiclient.NewRoleAssignment()}) // RoleMappingView | The role mapping that you want to create.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.AddRoleMapping(context.Background(), federationSettingsId, orgId).RoleMappingView(roleMappingView).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.AddRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleMapping`: RoleMappingView
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.AddRoleMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleMappingView** | [**RoleMappingView**](RoleMappingView.md) | The role mapping that you want to create. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**RoleMappingView**](RoleMappingView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFederationApp

> DeleteFederationApp(ctx, federationSettingsId).Execute()

Delete the federation settings instance.



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.DeleteFederationApp(context.Background(), federationSettingsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.DeleteFederationApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFederationAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteRoleMapping

> DeleteRoleMapping(ctx, federationSettingsId, id, orgId).Envelope(envelope).Execute()

Remove One Role Mapping from One Organization



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    id := "4faca5b098702de1952f4d09" // string | Unique 24-hexadecimal digit string that identifies the role mapping that you want to remove.
    orgId := "4888442a3354817a7320eb61" // string | Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.DeleteRoleMapping(context.Background(), federationSettingsId, id, orgId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.DeleteRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**id** | **string** | Unique 24-hexadecimal digit string that identifies the role mapping that you want to remove. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleMappingRequest struct via the builder pattern


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


## GetAllConnectedOrgConfigs

> []ConnectedOrgConfigView GetAllConnectedOrgConfigs(ctx, federationSettingsId).Envelope(envelope).Execute()

Return All Connected Org Configs from the Federation



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.GetAllConnectedOrgConfigs(context.Background(), federationSettingsId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetAllConnectedOrgConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConnectedOrgConfigs`: []ConnectedOrgConfigView
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetAllConnectedOrgConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConnectedOrgConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**[]ConnectedOrgConfigView**](ConnectedOrgConfigView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllIdentityProviders

> []IdentityProviderView GetAllIdentityProviders(ctx, federationSettingsId).Envelope(envelope).Execute()

Return all identity providers from the specified federation.



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.GetAllIdentityProviders(context.Background(), federationSettingsId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetAllIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllIdentityProviders`: []IdentityProviderView
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetAllIdentityProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**[]IdentityProviderView**](IdentityProviderView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRoleMappings

> []RoleMappingView GetAllRoleMappings(ctx, federationSettingsId, orgId).Envelope(envelope).Execute()

Return All Role Mappings from One Organization



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    orgId := "4888442a3354817a7320eb61" // string | Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.GetAllRoleMappings(context.Background(), federationSettingsId, orgId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetAllRoleMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRoleMappings`: []RoleMappingView
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetAllRoleMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRoleMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**[]RoleMappingView**](RoleMappingView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectedOrgConfig

> ConnectedOrgConfigView GetConnectedOrgConfig(ctx, federationSettingsId, orgId).Envelope(envelope).Execute()

Return One Org Config Connected to One Federation



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    orgId := "4888442a3354817a7320eb61" // string | Unique 24-hexadecimal digit string that identifies the connected organization configuration to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.GetConnectedOrgConfig(context.Background(), federationSettingsId, orgId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetConnectedOrgConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectedOrgConfig`: ConnectedOrgConfigView
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetConnectedOrgConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the connected organization configuration to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectedOrgConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ConnectedOrgConfigView**](ConnectedOrgConfigView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFederationSettings

> OrgFederationSettingsView GetFederationSettings(ctx, orgId).Envelope(envelope).Pretty(pretty).Execute()

Return Federation Settings for One Organization



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
    orgId := "4888442a3354817a7320eb61" // string | Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.GetFederationSettings(context.Background(), orgId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetFederationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFederationSettings`: OrgFederationSettingsView
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetFederationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFederationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**OrgFederationSettingsView**](OrgFederationSettingsView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProvider

> IdentityProviderView GetIdentityProvider(ctx, federationSettingsId, identityProviderId).Envelope(envelope).Execute()

Return one identity provider from the specified federation.



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    identityProviderId := "c2777a9eca931f29fc2f" // string | Unique 20-hexadecimal digit string that identifies the identity provider.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.GetIdentityProvider(context.Background(), federationSettingsId, identityProviderId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProvider`: IdentityProviderView
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**identityProviderId** | **string** | Unique 20-hexadecimal digit string that identifies the identity provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**IdentityProviderView**](IdentityProviderView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProviderMetadata

> string GetIdentityProviderMetadata(ctx, federationSettingsId, identityProviderId).Execute()

Return the metadata of one identity provider in the specified federation.



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    identityProviderId := "c2777a9eca931f29fc2f" // string | Unique 20-hexadecimal digit string that identifies the identity provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.GetIdentityProviderMetadata(context.Background(), federationSettingsId, identityProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetIdentityProviderMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProviderMetadata`: string
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetIdentityProviderMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**identityProviderId** | **string** | Unique 20-hexadecimal digit string that identifies the identity provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleMapping

> RoleMappingView GetRoleMapping(ctx, federationSettingsId, id, orgId).Envelope(envelope).Execute()

Return One Role Mapping from One Organization



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    id := "4faca5b098702de1952f4d09" // string | Unique 24-hexadecimal digit string that identifies the role mapping that you want to return.
    orgId := "4888442a3354817a7320eb61" // string | Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.GetRoleMapping(context.Background(), federationSettingsId, id, orgId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.GetRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleMapping`: RoleMappingView
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.GetRoleMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**id** | **string** | Unique 24-hexadecimal digit string that identifies the role mapping that you want to return. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**RoleMappingView**](RoleMappingView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveConnectedOrgConfig

> RemoveConnectedOrgConfig(ctx, federationSettingsId, orgId).Envelope(envelope).Execute()

Remove One Org Config Connected to One Federation



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    orgId := "4888442a3354817a7320eb61" // string | Unique 24-hexadecimal digit string that identifies the connected organization configuration to remove.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.RemoveConnectedOrgConfig(context.Background(), federationSettingsId, orgId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.RemoveConnectedOrgConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the connected organization configuration to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveConnectedOrgConfigRequest struct via the builder pattern


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


## UpdateConnectedOrgConfig

> ConnectedOrgConfigView UpdateConnectedOrgConfig(ctx, federationSettingsId, orgId).ConnectedOrgConfigView(connectedOrgConfigView).Envelope(envelope).Execute()

Update one connected organization configuration in the specified federation.



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    orgId := "4888442a3354817a7320eb61" // string | Unique 24-hexadecimal digit string that identifies the connected organization configuration to update.
    connectedOrgConfigView := *openapiclient.NewConnectedOrgConfigView(false, "IdentityProviderId_example", "4888442a3354817a7320eb61") // ConnectedOrgConfigView | The connected organization configuration that you want to update.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.UpdateConnectedOrgConfig(context.Background(), federationSettingsId, orgId).ConnectedOrgConfigView(connectedOrgConfigView).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.UpdateConnectedOrgConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectedOrgConfig`: ConnectedOrgConfigView
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.UpdateConnectedOrgConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the connected organization configuration to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectedOrgConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectedOrgConfigView** | [**ConnectedOrgConfigView**](ConnectedOrgConfigView.md) | The connected organization configuration that you want to update. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ConnectedOrgConfigView**](ConnectedOrgConfigView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityProvider

> IdentityProviderView UpdateIdentityProvider(ctx, federationSettingsId, identityProviderId).IdentityProviderUpdate(identityProviderUpdate).Envelope(envelope).Execute()

Update the identity provider.



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    identityProviderId := "c2777a9eca931f29fc2f" // string | Unique 20-hexadecimal digit string that identifies the identity provider.
    identityProviderUpdate := *openapiclient.NewIdentityProviderUpdate(false) // IdentityProviderUpdate | The identity provider that you want to update.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.UpdateIdentityProvider(context.Background(), federationSettingsId, identityProviderId).IdentityProviderUpdate(identityProviderUpdate).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.UpdateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentityProvider`: IdentityProviderView
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.UpdateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**identityProviderId** | **string** | Unique 20-hexadecimal digit string that identifies the identity provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identityProviderUpdate** | [**IdentityProviderUpdate**](IdentityProviderUpdate.md) | The identity provider that you want to update. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**IdentityProviderView**](IdentityProviderView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleMapping

> RoleMappingView UpdateRoleMapping(ctx, federationSettingsId, id, orgId).RoleMappingView(roleMappingView).Envelope(envelope).Execute()

Update One Role Mapping in One Organization



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
    federationSettingsId := "55fa922fb343282757d9554e" // string | Unique 24-hexadecimal digit string that identifies your federation.
    id := "4faca5b098702de1952f4d09" // string | Unique 24-hexadecimal digit string that identifies the role mapping that you want to update.
    orgId := "4888442a3354817a7320eb61" // string | Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
    roleMappingView := *openapiclient.NewRoleMappingView("ExternalGroupName_example", []openapiclient.RoleAssignment{*openapiclient.NewRoleAssignment()}) // RoleMappingView | The role mapping that you want to update.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederatedAuthenticationApi.UpdateRoleMapping(context.Background(), federationSettingsId, id, orgId).RoleMappingView(roleMappingView).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederatedAuthenticationApi.UpdateRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoleMapping`: RoleMappingView
    fmt.Fprintf(os.Stdout, "Response from `FederatedAuthenticationApi.UpdateRoleMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**federationSettingsId** | **string** | Unique 24-hexadecimal digit string that identifies your federation. | 
**id** | **string** | Unique 24-hexadecimal digit string that identifies the role mapping that you want to update. | 
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleMappingView** | [**RoleMappingView**](RoleMappingView.md) | The role mapping that you want to update. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**RoleMappingView**](RoleMappingView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

