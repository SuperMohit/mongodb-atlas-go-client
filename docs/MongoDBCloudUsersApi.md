# \MongoDBCloudUsersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneCloudUser**](MongoDBCloudUsersApi.md#CreateOneCloudUser) | **Post** /api/atlas/v1.0/users | Create One MongoDB Cloud User
[**ReturnOneCloudUserUsingItsId**](MongoDBCloudUsersApi.md#ReturnOneCloudUserUsingItsId) | **Get** /api/atlas/v1.0/users/{userId} | Return One MongoDB Cloud User using Its ID
[**ReturnOneCloudUserUsingItsUsername**](MongoDBCloudUsersApi.md#ReturnOneCloudUserUsingItsUsername) | **Get** /api/atlas/v1.0/users/byName/{userName} | Return One MongoDB Cloud User using Their Username



## CreateOneCloudUser

> ApiAppUserView CreateOneCloudUser(ctx).ApiAppUserView(apiAppUserView).Envelope(envelope).Pretty(pretty).Execute()

Create One MongoDB Cloud User



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
    apiAppUserView := *openapiclient.NewApiAppUserView("Country_example", "EmailAddress_example", "FirstName_example", "LastName_example", []openapiclient.Link{*openapiclient.NewLink()}, "MobileNumber_example", "Password_example", []openapiclient.ApiRoleAssignmentView{*openapiclient.NewApiRoleAssignmentView()}, "Username_example") // ApiAppUserView | MongoDB Cloud user account to create.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MongoDBCloudUsersApi.CreateOneCloudUser(context.Background()).ApiAppUserView(apiAppUserView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.CreateOneCloudUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneCloudUser`: ApiAppUserView
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.CreateOneCloudUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneCloudUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAppUserView** | [**ApiAppUserView**](ApiAppUserView.md) | MongoDB Cloud user account to create. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAppUserView**](ApiAppUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneCloudUserUsingItsId

> ApiAppUserView ReturnOneCloudUserUsingItsId(ctx, userId).Envelope(envelope).Pretty(pretty).Execute()

Return One MongoDB Cloud User using Its ID



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
    userId := "userId_example" // string | Unique 24-hexadecimal digit string that identifies this user.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MongoDBCloudUsersApi.ReturnOneCloudUserUsingItsId(context.Background(), userId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.ReturnOneCloudUserUsingItsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneCloudUserUsingItsId`: ApiAppUserView
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.ReturnOneCloudUserUsingItsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Unique 24-hexadecimal digit string that identifies this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneCloudUserUsingItsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAppUserView**](ApiAppUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneCloudUserUsingItsUsername

> ApiAppUserView ReturnOneCloudUserUsingItsUsername(ctx, userName).Envelope(envelope).Pretty(pretty).Execute()

Return One MongoDB Cloud User using Their Username



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
    userName := "userName_example" // string | Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MongoDBCloudUsersApi.ReturnOneCloudUserUsingItsUsername(context.Background(), userName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MongoDBCloudUsersApi.ReturnOneCloudUserUsingItsUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneCloudUserUsingItsUsername`: ApiAppUserView
    fmt.Fprintf(os.Stdout, "Response from `MongoDBCloudUsersApi.ReturnOneCloudUserUsingItsUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneCloudUserUsingItsUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAppUserView**](ApiAppUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

