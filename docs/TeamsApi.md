# \TeamsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOneOrMoreTeamsToOneProject**](TeamsApi.md#AddOneOrMoreTeamsToOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/teams | Add One or More Teams to One Project
[**AssignOneOrganizationUserToOneTeam**](TeamsApi.md#AssignOneOrganizationUserToOneTeam) | **Post** /api/atlas/v1.0/orgs/{orgId}/teams/{teamId}/users | Assign MongoDB Cloud Users from One Organization to One Team
[**CreateOneTeamInOneOrganization**](TeamsApi.md#CreateOneTeamInOneOrganization) | **Post** /api/atlas/v1.0/orgs/{orgId}/teams | Create One Team in One Organization
[**RemoveOneMongoDBCloudUserFromOneTeam**](TeamsApi.md#RemoveOneMongoDBCloudUserFromOneTeam) | **Delete** /api/atlas/v1.0/orgs/{orgId}/teams/{teamId}/users/{userId} | Remove One MongoDB Cloud User from One Team
[**RemoveOneTeamFromOneOrganization**](TeamsApi.md#RemoveOneTeamFromOneOrganization) | **Delete** /api/atlas/v1.0/orgs/{orgId}/teams/{teamId} | Remove One Team from One Organization
[**RemoveOneTeamFromOneProject**](TeamsApi.md#RemoveOneTeamFromOneProject) | **Delete** /api/atlas/v1.0/groups/{groupId}/teams/{teamId} | Remove One Team from One Project
[**RenameOneTeam**](TeamsApi.md#RenameOneTeam) | **Patch** /api/atlas/v1.0/orgs/{orgId}/teams/{teamId} | Rename One Team
[**ReturnAllMongoDBCloudUsersAssignedToOneTeam**](TeamsApi.md#ReturnAllMongoDBCloudUsersAssignedToOneTeam) | **Get** /api/atlas/v1.0/orgs/{orgId}/teams/{teamId}/users | Return All MongoDB Cloud Users Assigned to One Team
[**ReturnAllTeams**](TeamsApi.md#ReturnAllTeams) | **Get** /api/atlas/v1.0/groups/{groupId}/teams | Return All Teams in One Project
[**ReturnAllTeamsInOneOrganization**](TeamsApi.md#ReturnAllTeamsInOneOrganization) | **Get** /api/atlas/v1.0/orgs/{orgId}/teams | Return All Teams in One Organization
[**ReturnOneTeamUsingItsId**](TeamsApi.md#ReturnOneTeamUsingItsId) | **Get** /api/atlas/v1.0/orgs/{orgId}/teams/{teamId} | Return One Team using its ID
[**ReturnOneTeamUsingItsName**](TeamsApi.md#ReturnOneTeamUsingItsName) | **Get** /api/atlas/v1.0/orgs/{orgId}/teams/byName/{teamName} | Return One Team using its Name
[**UpdateTeamRolesInOneProject**](TeamsApi.md#UpdateTeamRolesInOneProject) | **Patch** /api/atlas/v1.0/groups/{groupId}/teams/{teamId} | Update Team Roles in One Project



## AddOneOrMoreTeamsToOneProject

> PaginatedTeamRoleView AddOneOrMoreTeamsToOneProject(ctx, groupId).ApiTeamRoleView(apiTeamRoleView).Envelope(envelope).Pretty(pretty).Execute()

Add One or More Teams to One Project



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
    apiTeamRoleView := *openapiclient.NewApiTeamRoleView([]openapiclient.Link{*openapiclient.NewLink()}) // ApiTeamRoleView | Team to add to the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.AddOneOrMoreTeamsToOneProject(context.Background(), groupId).ApiTeamRoleView(apiTeamRoleView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.AddOneOrMoreTeamsToOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOneOrMoreTeamsToOneProject`: PaginatedTeamRoleView
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.AddOneOrMoreTeamsToOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOneOrMoreTeamsToOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiTeamRoleView** | [**ApiTeamRoleView**](ApiTeamRoleView.md) | Team to add to the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedTeamRoleView**](PaginatedTeamRoleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignOneOrganizationUserToOneTeam

> []PaginatedAppUserView AssignOneOrganizationUserToOneTeam(ctx, orgId, teamId).ApiAddUserToTeamView(apiAddUserToTeamView).Envelope(envelope).Pretty(pretty).Execute()

Assign MongoDB Cloud Users from One Organization to One Team



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
    teamId := "teamId_example" // string | Unique 24-hexadecimal character string that identifies the team to which you want to add MongoDB Cloud users.
    apiAddUserToTeamView := *openapiclient.NewApiAddUserToTeamView("887821d389de9fe084df2bdf") // ApiAddUserToTeamView | One or more MongoDB Cloud users that you want to add to the specified team.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.AssignOneOrganizationUserToOneTeam(context.Background(), orgId, teamId).ApiAddUserToTeamView(apiAddUserToTeamView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.AssignOneOrganizationUserToOneTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignOneOrganizationUserToOneTeam`: []PaginatedAppUserView
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.AssignOneOrganizationUserToOneTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**teamId** | **string** | Unique 24-hexadecimal character string that identifies the team to which you want to add MongoDB Cloud users. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOneOrganizationUserToOneTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAddUserToTeamView** | [**ApiAddUserToTeamView**](ApiAddUserToTeamView.md) | One or more MongoDB Cloud users that you want to add to the specified team. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]PaginatedAppUserView**](PaginatedAppUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOneTeamInOneOrganization

> ApiTeamView CreateOneTeamInOneOrganization(ctx, orgId).ApiTeamView(apiTeamView).Envelope(envelope).Pretty(pretty).Execute()

Create One Team in One Organization



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
    apiTeamView := *openapiclient.NewApiTeamView("e254e4ade2206f4520d4107d", []openapiclient.Link{*openapiclient.NewLink()}, "Name_example", []string{"Usernames_example"}) // ApiTeamView | Team that you want to create in the specified organization.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.CreateOneTeamInOneOrganization(context.Background(), orgId).ApiTeamView(apiTeamView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.CreateOneTeamInOneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneTeamInOneOrganization`: ApiTeamView
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.CreateOneTeamInOneOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneTeamInOneOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiTeamView** | [**ApiTeamView**](ApiTeamView.md) | Team that you want to create in the specified organization. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiTeamView**](ApiTeamView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneMongoDBCloudUserFromOneTeam

> RemoveOneMongoDBCloudUserFromOneTeam(ctx, orgId, teamId, userId).Envelope(envelope).Pretty(pretty).Execute()

Remove One MongoDB Cloud User from One Team



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
    teamId := "teamId_example" // string | Unique 24-hexadecimal digit string that identifies the team from which you want to remove one database application user.
    userId := "userId_example" // string | Unique 24-hexadecimal digit string that identifies MongoDB Cloud user that you want to remove from the specified team.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.RemoveOneMongoDBCloudUserFromOneTeam(context.Background(), orgId, teamId, userId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.RemoveOneMongoDBCloudUserFromOneTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**teamId** | **string** | Unique 24-hexadecimal digit string that identifies the team from which you want to remove one database application user. | 
**userId** | **string** | Unique 24-hexadecimal digit string that identifies MongoDB Cloud user that you want to remove from the specified team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneMongoDBCloudUserFromOneTeamRequest struct via the builder pattern


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


## RemoveOneTeamFromOneOrganization

> RemoveOneTeamFromOneOrganization(ctx, orgId, teamId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Team from One Organization



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
    teamId := "teamId_example" // string | Unique 24-hexadecimal digit string that identifies the team that you want to delete.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.RemoveOneTeamFromOneOrganization(context.Background(), orgId, teamId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.RemoveOneTeamFromOneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**teamId** | **string** | Unique 24-hexadecimal digit string that identifies the team that you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneTeamFromOneOrganizationRequest struct via the builder pattern


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


## RemoveOneTeamFromOneProject

> RemoveOneTeamFromOneProject(ctx, groupId, teamId).Envelope(envelope).Execute()

Remove One Team from One Project



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
    teamId := "teamId_example" // string | Unique 24-hexadecimal digit string that identifies the team that you want to remove from the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.RemoveOneTeamFromOneProject(context.Background(), groupId, teamId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.RemoveOneTeamFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**teamId** | **string** | Unique 24-hexadecimal digit string that identifies the team that you want to remove from the specified project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneTeamFromOneProjectRequest struct via the builder pattern


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


## RenameOneTeam

> ApiTeamResponseView RenameOneTeam(ctx, orgId, teamId).ApiTeamView(apiTeamView).Envelope(envelope).Pretty(pretty).Execute()

Rename One Team



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
    teamId := "teamId_example" // string | Unique 24-hexadecimal digit string that identifies the team that you want to rename.
    apiTeamView := *openapiclient.NewApiTeamView("e254e4ade2206f4520d4107d", []openapiclient.Link{*openapiclient.NewLink()}, "Name_example", []string{"Usernames_example"}) // ApiTeamView | Details to update on the specified team.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.RenameOneTeam(context.Background(), orgId, teamId).ApiTeamView(apiTeamView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.RenameOneTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameOneTeam`: ApiTeamResponseView
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.RenameOneTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**teamId** | **string** | Unique 24-hexadecimal digit string that identifies the team that you want to rename. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameOneTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiTeamView** | [**ApiTeamView**](ApiTeamView.md) | Details to update on the specified team. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiTeamResponseView**](ApiTeamResponseView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllMongoDBCloudUsersAssignedToOneTeam

> PaginatedAppUserView ReturnAllMongoDBCloudUsersAssignedToOneTeam(ctx, orgId, teamId).Envelope(envelope).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All MongoDB Cloud Users Assigned to One Team



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
    teamId := "teamId_example" // string | Unique 24-hexadecimal digit string that identifies the team whose application users you want to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.ReturnAllMongoDBCloudUsersAssignedToOneTeam(context.Background(), orgId, teamId).Envelope(envelope).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ReturnAllMongoDBCloudUsersAssignedToOneTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllMongoDBCloudUsersAssignedToOneTeam`: PaginatedAppUserView
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ReturnAllMongoDBCloudUsersAssignedToOneTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**teamId** | **string** | Unique 24-hexadecimal digit string that identifies the team whose application users you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllMongoDBCloudUsersAssignedToOneTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
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


## ReturnAllTeams

> PaginatedTeamRoleView ReturnAllTeams(ctx, groupId).Envelope(envelope).Pretty(pretty).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Teams in One Project



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
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.ReturnAllTeams(context.Background(), groupId).Envelope(envelope).Pretty(pretty).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ReturnAllTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllTeams`: PaginatedTeamRoleView
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ReturnAllTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedTeamRoleView**](PaginatedTeamRoleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllTeamsInOneOrganization

> PaginatedTeamView ReturnAllTeamsInOneOrganization(ctx, orgId).Envelope(envelope).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Pretty(pretty).Execute()

Return All Teams in One Organization



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
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.ReturnAllTeamsInOneOrganization(context.Background(), orgId).Envelope(envelope).ItemsPerPage(itemsPerPage).IncludeCount(includeCount).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ReturnAllTeamsInOneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllTeamsInOneOrganization`: PaginatedTeamView
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ReturnAllTeamsInOneOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllTeamsInOneOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedTeamView**](PaginatedTeamView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneTeamUsingItsId

> ApiTeamResponseView ReturnOneTeamUsingItsId(ctx, orgId, teamId).Envelope(envelope).Pretty(pretty).Execute()

Return One Team using its ID



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
    teamId := "teamId_example" // string | Unique 24-hexadecimal digit string that identifies the team whose information you want to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.ReturnOneTeamUsingItsId(context.Background(), orgId, teamId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ReturnOneTeamUsingItsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneTeamUsingItsId`: ApiTeamResponseView
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ReturnOneTeamUsingItsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**teamId** | **string** | Unique 24-hexadecimal digit string that identifies the team whose information you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneTeamUsingItsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiTeamResponseView**](ApiTeamResponseView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneTeamUsingItsName

> ApiTeamResponseView ReturnOneTeamUsingItsName(ctx, orgId, teamName).Envelope(envelope).Pretty(pretty).Execute()

Return One Team using its Name



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
    teamName := "teamName_example" // string | Name of the team whose information you want to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.ReturnOneTeamUsingItsName(context.Background(), orgId, teamName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ReturnOneTeamUsingItsName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneTeamUsingItsName`: ApiTeamResponseView
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ReturnOneTeamUsingItsName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**teamName** | **string** | Name of the team whose information you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneTeamUsingItsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiTeamResponseView**](ApiTeamResponseView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamRolesInOneProject

> PaginatedTeamRoleView UpdateTeamRolesInOneProject(ctx, groupId, teamId).ApiTeamRoleView(apiTeamRoleView).Envelope(envelope).Pretty(pretty).Execute()

Update Team Roles in One Project



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
    teamId := "teamId_example" // string | Unique 24-hexadecimal digit string that identifies the team for which you want to update roles.
    apiTeamRoleView := *openapiclient.NewApiTeamRoleView([]openapiclient.Link{*openapiclient.NewLink()}) // ApiTeamRoleView | The project roles assigned to the specified team.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.UpdateTeamRolesInOneProject(context.Background(), groupId, teamId).ApiTeamRoleView(apiTeamRoleView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateTeamRolesInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamRolesInOneProject`: PaginatedTeamRoleView
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateTeamRolesInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**teamId** | **string** | Unique 24-hexadecimal digit string that identifies the team for which you want to update roles. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamRolesInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiTeamRoleView** | [**ApiTeamRoleView**](ApiTeamRoleView.md) | The project roles assigned to the specified team. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedTeamRoleView**](PaginatedTeamRoleView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

