# \CloudMigrationServiceApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvailableProjects**](CloudMigrationServiceApi.md#AvailableProjects) | **Get** /api/atlas/v1.0/orgs/{orgId}/liveMigrations/availableProjects | Return All Projects Available for Migration
[**CreateLinkToken**](CloudMigrationServiceApi.md#CreateLinkToken) | **Post** /api/atlas/v1.0/orgs/{orgId}/liveMigrations/linkTokens | Create One Link-Token
[**CreatePushMigration**](CloudMigrationServiceApi.md#CreatePushMigration) | **Post** /api/atlas/v1.0/groups/{groupId}/liveMigrations | Migrate One Local Managed Cluster to MongoDB Atlas
[**CutoverOneMigration**](CloudMigrationServiceApi.md#CutoverOneMigration) | **Put** /api/atlas/v1.0/groups/{groupId}/liveMigrations/{liveMigrationId}/cutover | Cut Over the Migrated Cluster
[**DeleteOrgLink**](CloudMigrationServiceApi.md#DeleteOrgLink) | **Delete** /api/public/v1.0/orgs/{orgId}/liveExport/migrationLink | Remove One Link between Organizations
[**DeleteOrgLink1**](CloudMigrationServiceApi.md#DeleteOrgLink1) | **Delete** /api/atlas/v1.0/orgs/{orgId}/liveMigrations/linkTokens | Remove One Link-Token
[**MigrationLink**](CloudMigrationServiceApi.md#MigrationLink) | **Post** /api/public/v1.0/orgs/{orgId}/liveExport/migrationLink | Link the Local Organization with MongoDB Atlas
[**MigrationLinkStatus**](CloudMigrationServiceApi.md#MigrationLinkStatus) | **Get** /api/public/v1.0/orgs/{orgId}/liveExport/migrationLink/status | Return Status of the Organization Link
[**ReturnOnePushMigration**](CloudMigrationServiceApi.md#ReturnOnePushMigration) | **Get** /api/atlas/v1.0/groups/{groupId}/liveMigrations/{liveMigrationId} | Return One Migration Job
[**ReturnOneValidationJob**](CloudMigrationServiceApi.md#ReturnOneValidationJob) | **Get** /api/atlas/v1.0/groups/{groupId}/liveMigrations/validate/{validationId} | Return One Migration Validation Job
[**ValidateOneMigration**](CloudMigrationServiceApi.md#ValidateOneMigration) | **Post** /api/atlas/v1.0/groups/{groupId}/liveMigrations/validate | Validate One Migration Request



## AvailableProjects

> []AvailableProjectView AvailableProjects(ctx, orgId).Envelope(envelope).Pretty(pretty).Execute()

Return All Projects Available for Migration



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
    resp, r, err := apiClient.CloudMigrationServiceApi.AvailableProjects(context.Background(), orgId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.AvailableProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvailableProjects`: []AvailableProjectView
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.AvailableProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAvailableProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]AvailableProjectView**](AvailableProjectView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLinkToken

> TargetOrgView CreateLinkToken(ctx, orgId).Envelope(envelope).Pretty(pretty).TargetOrgRequestView(targetOrgRequestView).Execute()

Create One Link-Token



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
    targetOrgRequestView := *openapiclient.NewTargetOrgRequestView() // TargetOrgRequestView |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudMigrationServiceApi.CreateLinkToken(context.Background(), orgId).Envelope(envelope).Pretty(pretty).TargetOrgRequestView(targetOrgRequestView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.CreateLinkToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLinkToken`: TargetOrgView
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.CreateLinkToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **targetOrgRequestView** | [**TargetOrgRequestView**](TargetOrgRequestView.md) |  | 

### Return type

[**TargetOrgView**](TargetOrgView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePushMigration

> LiveMigrationResponseView CreatePushMigration(ctx, groupId).LiveMigrationRequestView(liveMigrationRequestView).Envelope(envelope).Pretty(pretty).Execute()

Migrate One Local Managed Cluster to MongoDB Atlas



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
    liveMigrationRequestView := *openapiclient.NewLiveMigrationRequestView(*openapiclient.NewDestination("ClusterName_example", "9b43a5b329223c3a1591a678", "PUBLIC"), false, []string{"vm001.example.com"}, *openapiclient.NewSource("ClusterName_example", "9b43a5b329223c3a1591a678", false, false)) // LiveMigrationRequestView | One migration to be created.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudMigrationServiceApi.CreatePushMigration(context.Background(), groupId).LiveMigrationRequestView(liveMigrationRequestView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.CreatePushMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePushMigration`: LiveMigrationResponseView
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.CreatePushMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePushMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveMigrationRequestView** | [**LiveMigrationRequestView**](LiveMigrationRequestView.md) | One migration to be created. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**LiveMigrationResponseView**](LiveMigrationResponseView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CutoverOneMigration

> ValidationView CutoverOneMigration(ctx, groupId, liveMigrationId).Envelope(envelope).Pretty(pretty).Execute()

Cut Over the Migrated Cluster



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
    liveMigrationId := "6296fb4c7c7aa997cf94e9a8" // string | Unique 24-hexadecimal digit string that identifies the migration.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudMigrationServiceApi.CutoverOneMigration(context.Background(), groupId, liveMigrationId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.CutoverOneMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CutoverOneMigration`: ValidationView
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.CutoverOneMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**liveMigrationId** | **string** | Unique 24-hexadecimal digit string that identifies the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCutoverOneMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ValidationView**](ValidationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgLink

> DeleteOrgLink(ctx, orgId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Link between Organizations



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
    resp, r, err := apiClient.CloudMigrationServiceApi.DeleteOrgLink(context.Background(), orgId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.DeleteOrgLink``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrgLinkRequest struct via the builder pattern


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


## DeleteOrgLink1

> DeleteOrgLink1(ctx, orgId).Envelope(envelope).Execute()

Remove One Link-Token



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
    resp, r, err := apiClient.CloudMigrationServiceApi.DeleteOrgLink1(context.Background(), orgId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.DeleteOrgLink1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrgLink1Request struct via the builder pattern


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


## MigrationLink

> SourceOrgView MigrationLink(ctx, orgId).Envelope(envelope).Pretty(pretty).TargetOrgView(targetOrgView).Execute()

Link the Local Organization with MongoDB Atlas



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
    targetOrgView := *openapiclient.NewTargetOrgView("LinkToken_example") // TargetOrgView |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudMigrationServiceApi.MigrationLink(context.Background(), orgId).Envelope(envelope).Pretty(pretty).TargetOrgView(targetOrgView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.MigrationLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationLink`: SourceOrgView
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.MigrationLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **targetOrgView** | [**TargetOrgView**](TargetOrgView.md) |  | 

### Return type

[**SourceOrgView**](SourceOrgView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationLinkStatus

> MigrationLinkStatusView MigrationLinkStatus(ctx, orgId).Envelope(envelope).Pretty(pretty).Execute()

Return Status of the Organization Link



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
    resp, r, err := apiClient.CloudMigrationServiceApi.MigrationLinkStatus(context.Background(), orgId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.MigrationLinkStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationLinkStatus`: MigrationLinkStatusView
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.MigrationLinkStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationLinkStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**MigrationLinkStatusView**](MigrationLinkStatusView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOnePushMigration

> LiveMigrationResponseView ReturnOnePushMigration(ctx, groupId, liveMigrationId).Envelope(envelope).Pretty(pretty).Execute()

Return One Migration Job



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
    liveMigrationId := "6296fb4c7c7aa997cf94e9a8" // string | Unique 24-hexadecimal digit string that identifies the migration.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudMigrationServiceApi.ReturnOnePushMigration(context.Background(), groupId, liveMigrationId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.ReturnOnePushMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOnePushMigration`: LiveMigrationResponseView
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.ReturnOnePushMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**liveMigrationId** | **string** | Unique 24-hexadecimal digit string that identifies the migration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOnePushMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**LiveMigrationResponseView**](LiveMigrationResponseView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneValidationJob

> ValidationView ReturnOneValidationJob(ctx, groupId, validationId).Envelope(envelope).Execute()

Return One Migration Validation Job



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
    validationId := "507f1f77bcf86cd799439011" // string | Unique 24-hexadecimal digit string that identifies the validation job.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudMigrationServiceApi.ReturnOneValidationJob(context.Background(), groupId, validationId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.ReturnOneValidationJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneValidationJob`: ValidationView
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.ReturnOneValidationJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**validationId** | **string** | Unique 24-hexadecimal digit string that identifies the validation job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneValidationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ValidationView**](ValidationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateOneMigration

> ValidationView ValidateOneMigration(ctx, groupId).LiveMigrationRequestView(liveMigrationRequestView).Envelope(envelope).Pretty(pretty).Execute()

Validate One Migration Request



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
    liveMigrationRequestView := *openapiclient.NewLiveMigrationRequestView(*openapiclient.NewDestination("ClusterName_example", "9b43a5b329223c3a1591a678", "PUBLIC"), false, []string{"vm001.example.com"}, *openapiclient.NewSource("ClusterName_example", "9b43a5b329223c3a1591a678", false, false)) // LiveMigrationRequestView | One migration to be validated.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudMigrationServiceApi.ValidateOneMigration(context.Background(), groupId).LiveMigrationRequestView(liveMigrationRequestView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudMigrationServiceApi.ValidateOneMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateOneMigration`: ValidationView
    fmt.Fprintf(os.Stdout, "Response from `CloudMigrationServiceApi.ValidateOneMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateOneMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveMigrationRequestView** | [**LiveMigrationRequestView**](LiveMigrationRequestView.md) | One migration to be validated. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ValidationView**](ValidationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

