# \OrganizationsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOneOrganizationInvitation**](OrganizationsApi.md#CancelOneOrganizationInvitation) | **Delete** /api/atlas/v1.0/orgs/{orgId}/invites/{invitationId} | Cancel One Organization Invitation
[**InviteOneMongoDBUserToJoinOneAtlasOrganization**](OrganizationsApi.md#InviteOneMongoDBUserToJoinOneAtlasOrganization) | **Post** /api/atlas/v1.0/orgs/{orgId}/invites | Invite One MongoDB Cloud User to Join One Atlas Organization
[**RemoveOneOrganization**](OrganizationsApi.md#RemoveOneOrganization) | **Delete** /api/atlas/v1.0/orgs/{orgId} | Remove One Organization
[**RenameOneOrganization**](OrganizationsApi.md#RenameOneOrganization) | **Patch** /api/atlas/v1.0/orgs/{orgId} | Rename One Organization
[**ReturnAllOrganizationAtlasUsers**](OrganizationsApi.md#ReturnAllOrganizationAtlasUsers) | **Get** /api/atlas/v1.0/orgs/{orgId}/users | Return All MongoDB Cloud Users in One Organization
[**ReturnAllOrganizationInvitations**](OrganizationsApi.md#ReturnAllOrganizationInvitations) | **Get** /api/atlas/v1.0/orgs/{orgId}/invites | Return All Organization Invitations
[**ReturnAllOrganizations**](OrganizationsApi.md#ReturnAllOrganizations) | **Get** /api/atlas/v1.0/orgs | Return All Organizations
[**ReturnOneOrMoreProjectsInOneOrganization**](OrganizationsApi.md#ReturnOneOrMoreProjectsInOneOrganization) | **Get** /api/atlas/v1.0/orgs/{orgId}/groups | Return One or More Projects in One Organization
[**ReturnOneOrganization**](OrganizationsApi.md#ReturnOneOrganization) | **Get** /api/atlas/v1.0/orgs/{orgId} | Return One Organization
[**ReturnOneOrganizationInvitation**](OrganizationsApi.md#ReturnOneOrganizationInvitation) | **Get** /api/atlas/v1.0/orgs/{orgId}/invites/{invitationId} | Return One Organization Invitation
[**UpdateOneOrganizationInvitation**](OrganizationsApi.md#UpdateOneOrganizationInvitation) | **Patch** /api/atlas/v1.0/orgs/{orgId}/invites | Update One Organization Invitation
[**UpdateOneOrganizationInvitationByInvitationId**](OrganizationsApi.md#UpdateOneOrganizationInvitationByInvitationId) | **Patch** /api/atlas/v1.0/orgs/{orgId}/invites/{invitationId} | Update One Organization Invitation by Invitation ID



## CancelOneOrganizationInvitation

> CancelOneOrganizationInvitation(ctx, orgId, invitationId).Envelope(envelope).Pretty(pretty).Execute()

Cancel One Organization Invitation



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
    invitationId := "invitationId_example" // string | Unique 24-hexadecimal digit string that identifies the invitation.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CancelOneOrganizationInvitation(context.Background(), orgId, invitationId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CancelOneOrganizationInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**invitationId** | **string** | Unique 24-hexadecimal digit string that identifies the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOneOrganizationInvitationRequest struct via the builder pattern


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


## InviteOneMongoDBUserToJoinOneAtlasOrganization

> ApiOrganizationInvitationView InviteOneMongoDBUserToJoinOneAtlasOrganization(ctx, orgId).ApiOrganizationInvitationRequestView(apiOrganizationInvitationRequestView).Envelope(envelope).Pretty(pretty).Execute()

Invite One MongoDB Cloud User to Join One Atlas Organization



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
    apiOrganizationInvitationRequestView := *openapiclient.NewApiOrganizationInvitationRequestView() // ApiOrganizationInvitationRequestView | Invites one MongoDB Cloud user to join the specified organization.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.InviteOneMongoDBUserToJoinOneAtlasOrganization(context.Background(), orgId).ApiOrganizationInvitationRequestView(apiOrganizationInvitationRequestView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.InviteOneMongoDBUserToJoinOneAtlasOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InviteOneMongoDBUserToJoinOneAtlasOrganization`: ApiOrganizationInvitationView
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.InviteOneMongoDBUserToJoinOneAtlasOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteOneMongoDBUserToJoinOneAtlasOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiOrganizationInvitationRequestView** | [**ApiOrganizationInvitationRequestView**](ApiOrganizationInvitationRequestView.md) | Invites one MongoDB Cloud user to join the specified organization. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiOrganizationInvitationView**](ApiOrganizationInvitationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneOrganization

> RemoveOneOrganization(ctx, orgId).Envelope(envelope).Execute()

Remove One Organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.RemoveOneOrganization(context.Background(), orgId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.RemoveOneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneOrganizationRequest struct via the builder pattern


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


## RenameOneOrganization

> ApiOrganizationView RenameOneOrganization(ctx, orgId).ApiOrganizationView(apiOrganizationView).Envelope(envelope).Pretty(pretty).Execute()

Rename One Organization



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
    apiOrganizationView := *openapiclient.NewApiOrganizationView([]openapiclient.Link{*openapiclient.NewLink()}, "Name_example") // ApiOrganizationView | Details to update on the specified organization.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.RenameOneOrganization(context.Background(), orgId).ApiOrganizationView(apiOrganizationView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.RenameOneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameOneOrganization`: ApiOrganizationView
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.RenameOneOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameOneOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiOrganizationView** | [**ApiOrganizationView**](ApiOrganizationView.md) | Details to update on the specified organization. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiOrganizationView**](ApiOrganizationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllOrganizationAtlasUsers

> PaginatedAppUserView ReturnAllOrganizationAtlasUsers(ctx, orgId).Envelope(envelope).Pretty(pretty).Execute()

Return All MongoDB Cloud Users in One Organization



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
    resp, r, err := apiClient.OrganizationsApi.ReturnAllOrganizationAtlasUsers(context.Background(), orgId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ReturnAllOrganizationAtlasUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllOrganizationAtlasUsers`: PaginatedAppUserView
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ReturnAllOrganizationAtlasUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllOrganizationAtlasUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedAppUserView**](PaginatedAppUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllOrganizationInvitations

> PaginatedOrganizationInvitationView ReturnAllOrganizationInvitations(ctx, orgId).Envelope(envelope).Pretty(pretty).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Username(username).Execute()

Return All Organization Invitations



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
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    username := "username_example" // string | Email address of the user account invited to this organization. If you exclude this parameter, this resource returns all pending invitations. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ReturnAllOrganizationInvitations(context.Background(), orgId).Envelope(envelope).Pretty(pretty).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ReturnAllOrganizationInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllOrganizationInvitations`: PaginatedOrganizationInvitationView
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ReturnAllOrganizationInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllOrganizationInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **username** | **string** | Email address of the user account invited to this organization. If you exclude this parameter, this resource returns all pending invitations. | 

### Return type

[**PaginatedOrganizationInvitationView**](PaginatedOrganizationInvitationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllOrganizations

> PaginatedOrganizationView ReturnAllOrganizations(ctx).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Name(name).Execute()

Return All Organizations



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
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    name := "name_example" // string | Human-readable label of the organization to use to filter the returned list. Performs a case-insensitive search for an organization that starts with the specified name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ReturnAllOrganizations(context.Background()).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ReturnAllOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllOrganizations`: PaginatedOrganizationView
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ReturnAllOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **name** | **string** | Human-readable label of the organization to use to filter the returned list. Performs a case-insensitive search for an organization that starts with the specified name. | 

### Return type

[**PaginatedOrganizationView**](PaginatedOrganizationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneOrMoreProjectsInOneOrganization

> PaginatedAtlasGroupView ReturnOneOrMoreProjectsInOneOrganization(ctx, orgId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Name(name).Execute()

Return One or More Projects in One Organization



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
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    name := "name_example" // string | Human-readable label of the project to use to filter the returned list. Performs a case-insensitive search for a project within the organization which is prefixed by the specified name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ReturnOneOrMoreProjectsInOneOrganization(context.Background(), orgId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ReturnOneOrMoreProjectsInOneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneOrMoreProjectsInOneOrganization`: PaginatedAtlasGroupView
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ReturnOneOrMoreProjectsInOneOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneOrMoreProjectsInOneOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **name** | **string** | Human-readable label of the project to use to filter the returned list. Performs a case-insensitive search for a project within the organization which is prefixed by the specified name. | 

### Return type

[**PaginatedAtlasGroupView**](PaginatedAtlasGroupView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneOrganization

> ApiOrganizationView ReturnOneOrganization(ctx, orgId).Envelope(envelope).Pretty(pretty).Execute()

Return One Organization



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
    resp, r, err := apiClient.OrganizationsApi.ReturnOneOrganization(context.Background(), orgId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ReturnOneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneOrganization`: ApiOrganizationView
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ReturnOneOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiOrganizationView**](ApiOrganizationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneOrganizationInvitation

> ApiOrganizationInvitationView ReturnOneOrganizationInvitation(ctx, orgId, invitationId).Envelope(envelope).Execute()

Return One Organization Invitation



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
    invitationId := "invitationId_example" // string | Unique 24-hexadecimal digit string that identifies the invitation.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ReturnOneOrganizationInvitation(context.Background(), orgId, invitationId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ReturnOneOrganizationInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneOrganizationInvitation`: ApiOrganizationInvitationView
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ReturnOneOrganizationInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**invitationId** | **string** | Unique 24-hexadecimal digit string that identifies the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneOrganizationInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiOrganizationInvitationView**](ApiOrganizationInvitationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneOrganizationInvitation

> ApiOrganizationInvitationView UpdateOneOrganizationInvitation(ctx, orgId).ApiOrganizationInvitationRequestView(apiOrganizationInvitationRequestView).Envelope(envelope).Pretty(pretty).Execute()

Update One Organization Invitation



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
    apiOrganizationInvitationRequestView := *openapiclient.NewApiOrganizationInvitationRequestView() // ApiOrganizationInvitationRequestView | Updates the details of one pending invitation to the specified organization.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOneOrganizationInvitation(context.Background(), orgId).ApiOrganizationInvitationRequestView(apiOrganizationInvitationRequestView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOneOrganizationInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneOrganizationInvitation`: ApiOrganizationInvitationView
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOneOrganizationInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneOrganizationInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiOrganizationInvitationRequestView** | [**ApiOrganizationInvitationRequestView**](ApiOrganizationInvitationRequestView.md) | Updates the details of one pending invitation to the specified organization. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiOrganizationInvitationView**](ApiOrganizationInvitationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneOrganizationInvitationByInvitationId

> ApiOrganizationInvitationUpdateRequestView UpdateOneOrganizationInvitationByInvitationId(ctx, orgId, invitationId).ApiOrganizationInvitationUpdateRequestView(apiOrganizationInvitationUpdateRequestView).Envelope(envelope).Pretty(pretty).Execute()

Update One Organization Invitation by Invitation ID



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
    invitationId := "invitationId_example" // string | Unique 24-hexadecimal digit string that identifies the invitation.
    apiOrganizationInvitationUpdateRequestView := *openapiclient.NewApiOrganizationInvitationUpdateRequestView() // ApiOrganizationInvitationUpdateRequestView | Updates the details of one pending invitation to the specified organization.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOneOrganizationInvitationByInvitationId(context.Background(), orgId, invitationId).ApiOrganizationInvitationUpdateRequestView(apiOrganizationInvitationUpdateRequestView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOneOrganizationInvitationByInvitationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneOrganizationInvitationByInvitationId`: ApiOrganizationInvitationUpdateRequestView
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOneOrganizationInvitationByInvitationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**invitationId** | **string** | Unique 24-hexadecimal digit string that identifies the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneOrganizationInvitationByInvitationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiOrganizationInvitationUpdateRequestView** | [**ApiOrganizationInvitationUpdateRequestView**](ApiOrganizationInvitationUpdateRequestView.md) | Updates the details of one pending invitation to the specified organization. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiOrganizationInvitationUpdateRequestView**](ApiOrganizationInvitationUpdateRequestView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

