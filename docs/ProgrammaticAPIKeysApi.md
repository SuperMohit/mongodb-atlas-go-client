# \ProgrammaticAPIKeysApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignOneOrganizationApiKeyToOneProject**](ProgrammaticAPIKeysApi.md#AssignOneOrganizationApiKeyToOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/apiKeys/{apiUserId} | Assign One Organization API Key to One Project
[**CreateAccessListEntriesForOneOrganizationApiKey**](ProgrammaticAPIKeysApi.md#CreateAccessListEntriesForOneOrganizationApiKey) | **Post** /api/atlas/v1.0/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Create Access List Entries for One Organization API Key
[**CreateAndAssignOneOrganizationApiKeyToOneProject**](ProgrammaticAPIKeysApi.md#CreateAndAssignOneOrganizationApiKeyToOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/apiKeys | Create and Assign One Organization API Key to One Project
[**CreateOneOrganizationApiKey**](ProgrammaticAPIKeysApi.md#CreateOneOrganizationApiKey) | **Post** /api/atlas/v1.0/orgs/{orgId}/apiKeys | Create One Organization API Key
[**RemoveOneAccessListEntryForOneOrganizationApiKey**](ProgrammaticAPIKeysApi.md#RemoveOneAccessListEntryForOneOrganizationApiKey) | **Delete** /api/atlas/v1.0/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Remove One Access List Entry for One Organization API Key
[**RemoveOneOrganizationApiKey**](ProgrammaticAPIKeysApi.md#RemoveOneOrganizationApiKey) | **Delete** /api/atlas/v1.0/orgs/{orgId}/apiKeys/{apiUserId} | Remove One Organization API Key
[**ReturnAllAccessListEntriesForOneOrganizationApiKey**](ProgrammaticAPIKeysApi.md#ReturnAllAccessListEntriesForOneOrganizationApiKey) | **Get** /api/atlas/v1.0/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Return All Access List Entries for One Organization API Key
[**ReturnAllOrganizationApiKeys**](ProgrammaticAPIKeysApi.md#ReturnAllOrganizationApiKeys) | **Get** /api/atlas/v1.0/orgs/{orgId}/apiKeys | Return All Organization API Keys
[**ReturnAllOrganizationApiKeysAssignedToOneProject**](ProgrammaticAPIKeysApi.md#ReturnAllOrganizationApiKeysAssignedToOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/apiKeys | Return All Organization API Keys Assigned to One Project
[**ReturnOneAccessListEntryForOneOrganizationApiKey**](ProgrammaticAPIKeysApi.md#ReturnOneAccessListEntryForOneOrganizationApiKey) | **Get** /api/atlas/v1.0/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Return One Access List Entry for One Organization API Key
[**ReturnOneOrganizationApiKey**](ProgrammaticAPIKeysApi.md#ReturnOneOrganizationApiKey) | **Get** /api/atlas/v1.0/orgs/{orgId}/apiKeys/{apiUserId} | Return One Organization API Key
[**SetApiUserGroupRoles**](ProgrammaticAPIKeysApi.md#SetApiUserGroupRoles) | **Patch** /api/atlas/v1.0/groups/{groupId}/apiKeys/{apiUserId} | Update Roles of One Organization API Key to One Project
[**UnassignOneOrganizationApiKeyFromOneProject**](ProgrammaticAPIKeysApi.md#UnassignOneOrganizationApiKeyFromOneProject) | **Delete** /api/atlas/v1.0/groups/{groupId}/apiKeys/{apiUserId} | Unassign One Organization API Key from One Project
[**UpdateOneOrganizationApiKey**](ProgrammaticAPIKeysApi.md#UpdateOneOrganizationApiKey) | **Patch** /api/atlas/v1.0/orgs/{orgId}/apiKeys/{apiUserId} | Update One Organization API Key



## AssignOneOrganizationApiKeyToOneProject

> ApiApiUserView AssignOneOrganizationApiKeyToOneProject(ctx, groupId, apiUserId).ApiUserRoleAssignment(apiUserRoleAssignment).Envelope(envelope).Pretty(pretty).Execute()

Assign One Organization API Key to One Project



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
    apiUserId := "apiUserId_example" // string | Unique 24-hexadecimal digit string that identifies this organization API key that you want to assign to one project.
    apiUserRoleAssignment := []openapiclient.ApiUserRoleAssignment{*openapiclient.NewApiUserRoleAssignment()} // []ApiUserRoleAssignment | Organization API key to be assigned to the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.AssignOneOrganizationApiKeyToOneProject(context.Background(), groupId, apiUserId).ApiUserRoleAssignment(apiUserRoleAssignment).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.AssignOneOrganizationApiKeyToOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignOneOrganizationApiKeyToOneProject`: ApiApiUserView
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.AssignOneOrganizationApiKeyToOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key that you want to assign to one project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOneOrganizationApiKeyToOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiUserRoleAssignment** | [**[]ApiUserRoleAssignment**](ApiUserRoleAssignment.md) | Organization API key to be assigned to the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiApiUserView**](ApiApiUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccessListEntriesForOneOrganizationApiKey

> ApiUserAccessListView CreateAccessListEntriesForOneOrganizationApiKey(ctx, orgId, apiUserId).ApiUserAccessListView(apiUserAccessListView).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Create Access List Entries for One Organization API Key



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
    apiUserId := "apiUserId_example" // string | Unique 24-hexadecimal digit string that identifies this organization API key for which you want to create a new access list entry.
    apiUserAccessListView := []openapiclient.ApiUserAccessListView{*openapiclient.NewApiUserAccessListView([]openapiclient.Link{*openapiclient.NewLink()})} // []ApiUserAccessListView | Access list entries to be created for the specified organization API key.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.CreateAccessListEntriesForOneOrganizationApiKey(context.Background(), orgId, apiUserId).ApiUserAccessListView(apiUserAccessListView).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.CreateAccessListEntriesForOneOrganizationApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessListEntriesForOneOrganizationApiKey`: ApiUserAccessListView
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.CreateAccessListEntriesForOneOrganizationApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key for which you want to create a new access list entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessListEntriesForOneOrganizationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiUserAccessListView** | [**[]ApiUserAccessListView**](ApiUserAccessListView.md) | Access list entries to be created for the specified organization API key. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiUserAccessListView**](ApiUserAccessListView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAndAssignOneOrganizationApiKeyToOneProject

> ApiApiUserView CreateAndAssignOneOrganizationApiKeyToOneProject(ctx, groupId).ApiCreateApiKeyView(apiCreateApiKeyView).Envelope(envelope).Pretty(pretty).Execute()

Create and Assign One Organization API Key to One Project



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
    apiCreateApiKeyView := *openapiclient.NewApiCreateApiKeyView() // ApiCreateApiKeyView | Organization API key to be created and assigned to the specified project. This request requires a minimum of one of the two body parameters.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.CreateAndAssignOneOrganizationApiKeyToOneProject(context.Background(), groupId).ApiCreateApiKeyView(apiCreateApiKeyView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.CreateAndAssignOneOrganizationApiKeyToOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAndAssignOneOrganizationApiKeyToOneProject`: ApiApiUserView
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.CreateAndAssignOneOrganizationApiKeyToOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAndAssignOneOrganizationApiKeyToOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiCreateApiKeyView** | [**ApiCreateApiKeyView**](ApiCreateApiKeyView.md) | Organization API key to be created and assigned to the specified project. This request requires a minimum of one of the two body parameters. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiApiUserView**](ApiApiUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOneOrganizationApiKey

> ApiApiUserView CreateOneOrganizationApiKey(ctx, orgId).ApiCreateApiKeyView(apiCreateApiKeyView).Envelope(envelope).Pretty(pretty).Execute()

Create One Organization API Key



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
    apiCreateApiKeyView := *openapiclient.NewApiCreateApiKeyView() // ApiCreateApiKeyView | Organization API Key to be created. This request requires a minimum of one of the two body parameters.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.CreateOneOrganizationApiKey(context.Background(), orgId).ApiCreateApiKeyView(apiCreateApiKeyView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.CreateOneOrganizationApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneOrganizationApiKey`: ApiApiUserView
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.CreateOneOrganizationApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneOrganizationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiCreateApiKeyView** | [**ApiCreateApiKeyView**](ApiCreateApiKeyView.md) | Organization API Key to be created. This request requires a minimum of one of the two body parameters. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiApiUserView**](ApiApiUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneAccessListEntryForOneOrganizationApiKey

> RemoveOneAccessListEntryForOneOrganizationApiKey(ctx, orgId, apiUserId, ipAddress).Envelope(envelope).Pretty(pretty).Execute()

Remove One Access List Entry for One Organization API Key



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
    apiUserId := "apiUserId_example" // string | Unique 24-hexadecimal digit string that identifies this organization API key for which you want to remove access list entries.
    ipAddress := "192.0.2.0%2F24" // string | One IP address or multiple IP addresses represented as one CIDR block to limit requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.RemoveOneAccessListEntryForOneOrganizationApiKey(context.Background(), orgId, apiUserId, ipAddress).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.RemoveOneAccessListEntryForOneOrganizationApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key for which you want to remove access list entries. | 
**ipAddress** | **string** | One IP address or multiple IP addresses represented as one CIDR block to limit requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneAccessListEntryForOneOrganizationApiKeyRequest struct via the builder pattern


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


## RemoveOneOrganizationApiKey

> RemoveOneOrganizationApiKey(ctx, orgId, apiUserId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Organization API Key



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
    apiUserId := "apiUserId_example" // string | Unique 24-hexadecimal digit string that identifies this organization API key.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.RemoveOneOrganizationApiKey(context.Background(), orgId, apiUserId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.RemoveOneOrganizationApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneOrganizationApiKeyRequest struct via the builder pattern


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


## ReturnAllAccessListEntriesForOneOrganizationApiKey

> []ApiUserAccessListView ReturnAllAccessListEntriesForOneOrganizationApiKey(ctx, orgId, apiUserId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Access List Entries for One Organization API Key



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
    apiUserId := "apiUserId_example" // string | Unique 24-hexadecimal digit string that identifies this organization API key for which you want to return access list entries.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.ReturnAllAccessListEntriesForOneOrganizationApiKey(context.Background(), orgId, apiUserId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.ReturnAllAccessListEntriesForOneOrganizationApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllAccessListEntriesForOneOrganizationApiKey`: []ApiUserAccessListView
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.ReturnAllAccessListEntriesForOneOrganizationApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key for which you want to return access list entries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllAccessListEntriesForOneOrganizationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiUserAccessListView**](ApiUserAccessListView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllOrganizationApiKeys

> []ApiApiUserView ReturnAllOrganizationApiKeys(ctx, orgId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Organization API Keys



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.ReturnAllOrganizationApiKeys(context.Background(), orgId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.ReturnAllOrganizationApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllOrganizationApiKeys`: []ApiApiUserView
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.ReturnAllOrganizationApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllOrganizationApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiApiUserView**](ApiApiUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllOrganizationApiKeysAssignedToOneProject

> []ApiApiUserView ReturnAllOrganizationApiKeysAssignedToOneProject(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Organization API Keys Assigned to One Project



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
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.ReturnAllOrganizationApiKeysAssignedToOneProject(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.ReturnAllOrganizationApiKeysAssignedToOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllOrganizationApiKeysAssignedToOneProject`: []ApiApiUserView
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.ReturnAllOrganizationApiKeysAssignedToOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllOrganizationApiKeysAssignedToOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiApiUserView**](ApiApiUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneAccessListEntryForOneOrganizationApiKey

> []ApiUserAccessListView ReturnOneAccessListEntryForOneOrganizationApiKey(ctx, orgId, ipAddress, apiUserId).Envelope(envelope).Pretty(pretty).Execute()

Return One Access List Entry for One Organization API Key



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
    ipAddress := "192.0.2.0%2F24" // string | One IP address or multiple IP addresses represented as one CIDR block to limit  requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as  192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.
    apiUserId := "apiUserId_example" // string | Unique 24-hexadecimal digit string that identifies this organization API key for  which you want to return access list entries.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.ReturnOneAccessListEntryForOneOrganizationApiKey(context.Background(), orgId, ipAddress, apiUserId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.ReturnOneAccessListEntryForOneOrganizationApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneAccessListEntryForOneOrganizationApiKey`: []ApiUserAccessListView
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.ReturnOneAccessListEntryForOneOrganizationApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**ipAddress** | **string** | One IP address or multiple IP addresses represented as one CIDR block to limit  requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as  192.0.2.0/24, use the URL-encoded value %2F for the forward slash /. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key for  which you want to return access list entries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneAccessListEntryForOneOrganizationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiUserAccessListView**](ApiUserAccessListView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneOrganizationApiKey

> ApiApiUserView ReturnOneOrganizationApiKey(ctx, orgId, apiUserId).Envelope(envelope).Pretty(pretty).Execute()

Return One Organization API Key



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
    apiUserId := "apiUserId_example" // string | Unique 24-hexadecimal digit string that identifies this organization API key that  you want to update.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.ReturnOneOrganizationApiKey(context.Background(), orgId, apiUserId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.ReturnOneOrganizationApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneOrganizationApiKey`: ApiApiUserView
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.ReturnOneOrganizationApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key that  you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneOrganizationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiApiUserView**](ApiApiUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetApiUserGroupRoles

> ApiApiUserView SetApiUserGroupRoles(ctx, groupId, apiUserId).ApiCreateApiKeyView(apiCreateApiKeyView).PageNum(pageNum).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).Pretty(pretty).Execute()

Update Roles of One Organization API Key to One Project



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
    apiUserId := "apiUserId_example" // string | Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.
    apiCreateApiKeyView := *openapiclient.NewApiCreateApiKeyView() // ApiCreateApiKeyView | Organization API Key to be updated. This request requires a minimum of one of the two body parameters.
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.SetApiUserGroupRoles(context.Background(), groupId, apiUserId).ApiCreateApiKeyView(apiCreateApiKeyView).PageNum(pageNum).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.SetApiUserGroupRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetApiUserGroupRoles`: ApiApiUserView
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.SetApiUserGroupRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetApiUserGroupRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiCreateApiKeyView** | [**ApiCreateApiKeyView**](ApiCreateApiKeyView.md) | Organization API Key to be updated. This request requires a minimum of one of the two body parameters. | 
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiApiUserView**](ApiApiUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignOneOrganizationApiKeyFromOneProject

> UnassignOneOrganizationApiKeyFromOneProject(ctx, groupId, apiUserId).Envelope(envelope).Pretty(pretty).Execute()

Unassign One Organization API Key from One Project



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
    apiUserId := "apiUserId_example" // string | Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.UnassignOneOrganizationApiKeyFromOneProject(context.Background(), groupId, apiUserId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.UnassignOneOrganizationApiKeyFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignOneOrganizationApiKeyFromOneProjectRequest struct via the builder pattern


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


## UpdateOneOrganizationApiKey

> ApiApiUserView UpdateOneOrganizationApiKey(ctx, orgId, apiUserId).ApiApiUserView(apiApiUserView).Envelope(envelope).Pretty(pretty).Execute()

Update One Organization API Key



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
    apiUserId := "apiUserId_example" // string | Unique 24-hexadecimal digit string that identifies this organization API key you  want to update.
    apiApiUserView := *openapiclient.NewApiApiUserView([]openapiclient.Link{*openapiclient.NewLink()}) // ApiApiUserView | Organization API key to be updated. This request requires a minimum of one of the two body parameters.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProgrammaticAPIKeysApi.UpdateOneOrganizationApiKey(context.Background(), orgId, apiUserId).ApiApiUserView(apiApiUserView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticAPIKeysApi.UpdateOneOrganizationApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneOrganizationApiKey`: ApiApiUserView
    fmt.Fprintf(os.Stdout, "Response from `ProgrammaticAPIKeysApi.UpdateOneOrganizationApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**apiUserId** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key you  want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneOrganizationApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiApiUserView** | [**ApiApiUserView**](ApiApiUserView.md) | Organization API key to be updated. This request requires a minimum of one of the two body parameters. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiApiUserView**](ApiApiUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

