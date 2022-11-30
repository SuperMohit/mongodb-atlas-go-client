# \X509AuthenticationForDatabaseUsersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneX509CertificateForOneMongodbUser**](X509AuthenticationForDatabaseUsersApi.md#CreateOneX509CertificateForOneMongodbUser) | **Post** /api/atlas/v1.0/groups/{groupId}/databaseUsers/{username}/certs | Create One X.509 Certificate for One MongoDB User
[**DisableCustomerManagedX509**](X509AuthenticationForDatabaseUsersApi.md#DisableCustomerManagedX509) | **Delete** /api/atlas/v1.0/groups/{groupId}/userSecurity/customerX509 | Disable Customer-Managed X.509
[**ReturnAllX509CertificatesAssignedToOneMongodbUser**](X509AuthenticationForDatabaseUsersApi.md#ReturnAllX509CertificatesAssignedToOneMongodbUser) | **Get** /api/atlas/v1.0/groups/{groupId}/databaseUsers/{username}/certs | Return All X.509 Certificates Assigned to One MongoDB User



## CreateOneX509CertificateForOneMongodbUser

> ApiAtlasUserCertView CreateOneX509CertificateForOneMongodbUser(ctx, groupId, username).ApiAtlasUserCertView(apiAtlasUserCertView).Envelope(envelope).Pretty(pretty).Execute()

Create One X.509 Certificate for One MongoDB User



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
    username := "username_example" // string | Human-readable label that represents the MongoDB database user account for whom to create a certificate.
    apiAtlasUserCertView := *openapiclient.NewApiAtlasUserCertView([]openapiclient.Link{*openapiclient.NewLink()}) // ApiAtlasUserCertView | Generates one X.509 certificate for the specified MongoDB user.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.X509AuthenticationForDatabaseUsersApi.CreateOneX509CertificateForOneMongodbUser(context.Background(), groupId, username).ApiAtlasUserCertView(apiAtlasUserCertView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `X509AuthenticationForDatabaseUsersApi.CreateOneX509CertificateForOneMongodbUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneX509CertificateForOneMongodbUser`: ApiAtlasUserCertView
    fmt.Fprintf(os.Stdout, "Response from `X509AuthenticationForDatabaseUsersApi.CreateOneX509CertificateForOneMongodbUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**username** | **string** | Human-readable label that represents the MongoDB database user account for whom to create a certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneX509CertificateForOneMongodbUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasUserCertView** | [**ApiAtlasUserCertView**](ApiAtlasUserCertView.md) | Generates one X.509 certificate for the specified MongoDB user. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasUserCertView**](ApiAtlasUserCertView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableCustomerManagedX509

> ApiAtlasUserSecurityView DisableCustomerManagedX509(ctx, groupId).Envelope(envelope).Execute()

Disable Customer-Managed X.509



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
    resp, r, err := apiClient.X509AuthenticationForDatabaseUsersApi.DisableCustomerManagedX509(context.Background(), groupId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `X509AuthenticationForDatabaseUsersApi.DisableCustomerManagedX509``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableCustomerManagedX509`: ApiAtlasUserSecurityView
    fmt.Fprintf(os.Stdout, "Response from `X509AuthenticationForDatabaseUsersApi.DisableCustomerManagedX509`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableCustomerManagedX509Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

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


## ReturnAllX509CertificatesAssignedToOneMongodbUser

> PaginatedUserCertView ReturnAllX509CertificatesAssignedToOneMongodbUser(ctx, groupId, username).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All X.509 Certificates Assigned to One MongoDB User



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
    username := "username_example" // string | Human-readable label that represents the MongoDB database user account whose certificates you want to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.X509AuthenticationForDatabaseUsersApi.ReturnAllX509CertificatesAssignedToOneMongodbUser(context.Background(), groupId, username).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `X509AuthenticationForDatabaseUsersApi.ReturnAllX509CertificatesAssignedToOneMongodbUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllX509CertificatesAssignedToOneMongodbUser`: PaginatedUserCertView
    fmt.Fprintf(os.Stdout, "Response from `X509AuthenticationForDatabaseUsersApi.ReturnAllX509CertificatesAssignedToOneMongodbUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**username** | **string** | Human-readable label that represents the MongoDB database user account whose certificates you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllX509CertificatesAssignedToOneMongodbUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedUserCertView**](PaginatedUserCertView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

