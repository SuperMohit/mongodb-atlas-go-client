# \LDAPConfigurationApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoveOneCurrentLdapUserToDnMapping**](LDAPConfigurationApi.md#RemoveOneCurrentLdapUserToDnMapping) | **Delete** /api/atlas/v1.0/groups/{groupId}/userSecurity/ldap/userToDNMapping | Remove the Current LDAP User to DN Mapping
[**ReturnOneCurrentLdapConfiguration**](LDAPConfigurationApi.md#ReturnOneCurrentLdapConfiguration) | **Get** /api/atlas/v1.0/groups/{groupId}/userSecurity | Return the Current LDAP or X.509 Configuration
[**ReturnOneStatusOfOneVerifyLdapConfigurationRequest**](LDAPConfigurationApi.md#ReturnOneStatusOfOneVerifyLdapConfigurationRequest) | **Get** /api/atlas/v1.0/groups/{groupId}/userSecurity/ldap/verify/{requestId} | Return the Status of One Verify LDAP Configuration Request
[**SaveOneLdapConfiguration**](LDAPConfigurationApi.md#SaveOneLdapConfiguration) | **Patch** /api/atlas/v1.0/groups/{groupId}/userSecurity | Edit the LDAP or X.509 Configuration
[**VerifyOneLdapConfigurationInOneProject**](LDAPConfigurationApi.md#VerifyOneLdapConfigurationInOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/userSecurity/ldap/verify | Verify the LDAP Configuration in One Project



## RemoveOneCurrentLdapUserToDnMapping

> RemoveOneCurrentLdapUserToDnMapping(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Remove the Current LDAP User to DN Mapping



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
    resp, r, err := apiClient.LDAPConfigurationApi.RemoveOneCurrentLdapUserToDnMapping(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.RemoveOneCurrentLdapUserToDnMapping``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveOneCurrentLdapUserToDnMappingRequest struct via the builder pattern


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


## ReturnOneCurrentLdapConfiguration

> ApiAtlasUserSecurityView ReturnOneCurrentLdapConfiguration(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Return the Current LDAP or X.509 Configuration



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
    resp, r, err := apiClient.LDAPConfigurationApi.ReturnOneCurrentLdapConfiguration(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.ReturnOneCurrentLdapConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneCurrentLdapConfiguration`: ApiAtlasUserSecurityView
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.ReturnOneCurrentLdapConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneCurrentLdapConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasUserSecurityView**](ApiAtlasUserSecurityView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneStatusOfOneVerifyLdapConfigurationRequest

> ApiAtlasNDSLDAPVerifyConnectivityJobRequestView ReturnOneStatusOfOneVerifyLdapConfigurationRequest(ctx, groupId, requestId).Envelope(envelope).Pretty(pretty).Execute()

Return the Status of One Verify LDAP Configuration Request



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
    requestId := "requestId_example" // string | Unique string that identifies the request to verify an <abbr title=\"Lightweight Directory Access Protocol\">LDAP</abbr> configuration.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LDAPConfigurationApi.ReturnOneStatusOfOneVerifyLdapConfigurationRequest(context.Background(), groupId, requestId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.ReturnOneStatusOfOneVerifyLdapConfigurationRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneStatusOfOneVerifyLdapConfigurationRequest`: ApiAtlasNDSLDAPVerifyConnectivityJobRequestView
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.ReturnOneStatusOfOneVerifyLdapConfigurationRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**requestId** | **string** | Unique string that identifies the request to verify an &lt;abbr title&#x3D;\&quot;Lightweight Directory Access Protocol\&quot;&gt;LDAP&lt;/abbr&gt; configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneStatusOfOneVerifyLdapConfigurationRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasNDSLDAPVerifyConnectivityJobRequestView**](ApiAtlasNDSLDAPVerifyConnectivityJobRequestView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveOneLdapConfiguration

> ApiAtlasUserSecurityView SaveOneLdapConfiguration(ctx, groupId).ApiAtlasUserSecurityView(apiAtlasUserSecurityView).Envelope(envelope).Pretty(pretty).Execute()

Edit the LDAP or X.509 Configuration



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
    apiAtlasUserSecurityView := *openapiclient.NewApiAtlasUserSecurityView([]openapiclient.Link{*openapiclient.NewLink()}) // ApiAtlasUserSecurityView | Updates the LDAP configuration for the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LDAPConfigurationApi.SaveOneLdapConfiguration(context.Background(), groupId).ApiAtlasUserSecurityView(apiAtlasUserSecurityView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.SaveOneLdapConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveOneLdapConfiguration`: ApiAtlasUserSecurityView
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.SaveOneLdapConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveOneLdapConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasUserSecurityView** | [**ApiAtlasUserSecurityView**](ApiAtlasUserSecurityView.md) | Updates the LDAP configuration for the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasUserSecurityView**](ApiAtlasUserSecurityView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyOneLdapConfigurationInOneProject

> ApiAtlasNDSLDAPVerifyConnectivityJobRequestView VerifyOneLdapConfigurationInOneProject(ctx, groupId).ApiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView(apiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView).Envelope(envelope).Pretty(pretty).Execute()

Verify the LDAP Configuration in One Project



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
    apiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView := *openapiclient.NewApiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView("BindPassword_example", "CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com", "Hostname_example", []openapiclient.Link{*openapiclient.NewLink()}, int32(123)) // ApiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView | The LDAP configuration for the specified project that you want to verify.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LDAPConfigurationApi.VerifyOneLdapConfigurationInOneProject(context.Background(), groupId).ApiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView(apiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPConfigurationApi.VerifyOneLdapConfigurationInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyOneLdapConfigurationInOneProject`: ApiAtlasNDSLDAPVerifyConnectivityJobRequestView
    fmt.Fprintf(os.Stdout, "Response from `LDAPConfigurationApi.VerifyOneLdapConfigurationInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyOneLdapConfigurationInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView** | [**ApiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView**](ApiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView.md) | The LDAP configuration for the specified project that you want to verify. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasNDSLDAPVerifyConnectivityJobRequestView**](ApiAtlasNDSLDAPVerifyConnectivityJobRequestView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

