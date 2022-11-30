# \DataFederationApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneFederatedDatabaseInOneProject**](DataFederationApi.md#CreateOneFederatedDatabaseInOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/dataFederation | Create One Federated Database in One Project
[**DownloadQueryLogsForOneFederatedDatabase**](DataFederationApi.md#DownloadQueryLogsForOneFederatedDatabase) | **Get** /api/atlas/v1.0/groups/{groupId}/dataFederation/{tenantName}/queryLogs.gz | Download Query Logs for One Federated Database
[**RemoveOneFederatedDatabaseFromOneProject**](DataFederationApi.md#RemoveOneFederatedDatabaseFromOneProject) | **Delete** /api/atlas/v1.0/groups/{groupId}/dataFederation/{tenantName} | Remove One Federated Database from One Project
[**ReturnAllFederatedDatabasesInOneProject**](DataFederationApi.md#ReturnAllFederatedDatabasesInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/dataFederation | Return All Federated Databases in One Project
[**ReturnOneFederatedDatabaseInOneProject**](DataFederationApi.md#ReturnOneFederatedDatabaseInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/dataFederation/{tenantName} | Return One Federated Database in One Project
[**UpdateOneFederatedDatabaseInOneProject**](DataFederationApi.md#UpdateOneFederatedDatabaseInOneProject) | **Patch** /api/atlas/v1.0/groups/{groupId}/dataFederation/{tenantName} | Update One Federated Database in One Project



## CreateOneFederatedDatabaseInOneProject

> ApiAtlasDataLakeTenantView CreateOneFederatedDatabaseInOneProject(ctx, groupId).ApiAtlasDataLakeTenantView(apiAtlasDataLakeTenantView).Envelope(envelope).Pretty(pretty).Execute()

Create One Federated Database in One Project



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
    apiAtlasDataLakeTenantView := *openapiclient.NewApiAtlasDataLakeTenantView() // ApiAtlasDataLakeTenantView | Details to create one Federated Database in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataFederationApi.CreateOneFederatedDatabaseInOneProject(context.Background(), groupId).ApiAtlasDataLakeTenantView(apiAtlasDataLakeTenantView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.CreateOneFederatedDatabaseInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneFederatedDatabaseInOneProject`: ApiAtlasDataLakeTenantView
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.CreateOneFederatedDatabaseInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneFederatedDatabaseInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasDataLakeTenantView** | [**ApiAtlasDataLakeTenantView**](ApiAtlasDataLakeTenantView.md) | Details to create one Federated Database in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDataLakeTenantView**](ApiAtlasDataLakeTenantView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadQueryLogsForOneFederatedDatabase

> *os.File DownloadQueryLogsForOneFederatedDatabase(ctx, groupId, tenantName).EndDate(endDate).StartDate(startDate).Execute()

Download Query Logs for One Federated Database



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
    tenantName := "tenantName_example" // string | Human-readable label that identifies the Federated Database for which you want to download query logs.
    endDate := int64(1636481348) // int64 | Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. (optional)
    startDate := int64(1636466948) // int64 | Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataFederationApi.DownloadQueryLogsForOneFederatedDatabase(context.Background(), groupId, tenantName).EndDate(endDate).StartDate(startDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.DownloadQueryLogsForOneFederatedDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadQueryLogsForOneFederatedDatabase`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.DownloadQueryLogsForOneFederatedDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**tenantName** | **string** | Human-readable label that identifies the Federated Database for which you want to download query logs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadQueryLogsForOneFederatedDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **int64** | Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. | 
 **startDate** | **int64** | Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/gzip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneFederatedDatabaseFromOneProject

> RemoveOneFederatedDatabaseFromOneProject(ctx, groupId, tenantName).Envelope(envelope).Pretty(pretty).Execute()

Remove One Federated Database from One Project



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
    tenantName := "tenantName_example" // string | Human-readable label that identifies the Federated Database to remove.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataFederationApi.RemoveOneFederatedDatabaseFromOneProject(context.Background(), groupId, tenantName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.RemoveOneFederatedDatabaseFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**tenantName** | **string** | Human-readable label that identifies the Federated Database to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneFederatedDatabaseFromOneProjectRequest struct via the builder pattern


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


## ReturnAllFederatedDatabasesInOneProject

> []ApiAtlasDataLakeTenantView ReturnAllFederatedDatabasesInOneProject(ctx, groupId).Envelope(envelope).Pretty(pretty).Type_(type_).Execute()

Return All Federated Databases in One Project



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
    type_ := "type__example" // string | Type of Federated Databases to return. (optional) (default to "USER")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataFederationApi.ReturnAllFederatedDatabasesInOneProject(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.ReturnAllFederatedDatabasesInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllFederatedDatabasesInOneProject`: []ApiAtlasDataLakeTenantView
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.ReturnAllFederatedDatabasesInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllFederatedDatabasesInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **type_** | **string** | Type of Federated Databases to return. | [default to &quot;USER&quot;]

### Return type

[**[]ApiAtlasDataLakeTenantView**](ApiAtlasDataLakeTenantView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneFederatedDatabaseInOneProject

> ApiAtlasDataLakeTenantView ReturnOneFederatedDatabaseInOneProject(ctx, groupId, tenantName).Envelope(envelope).Execute()

Return One Federated Database in One Project



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
    tenantName := "tenantName_example" // string | Human-readable label that identifies the Federated Database to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataFederationApi.ReturnOneFederatedDatabaseInOneProject(context.Background(), groupId, tenantName).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.ReturnOneFederatedDatabaseInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneFederatedDatabaseInOneProject`: ApiAtlasDataLakeTenantView
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.ReturnOneFederatedDatabaseInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**tenantName** | **string** | Human-readable label that identifies the Federated Database to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneFederatedDatabaseInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiAtlasDataLakeTenantView**](ApiAtlasDataLakeTenantView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneFederatedDatabaseInOneProject

> ApiAtlasDataLakeTenantView UpdateOneFederatedDatabaseInOneProject(ctx, groupId, tenantName).SkipRoleValidation(skipRoleValidation).ApiAtlasDataLakeTenantView(apiAtlasDataLakeTenantView).Envelope(envelope).Pretty(pretty).Execute()

Update One Federated Database in One Project



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
    tenantName := "tenantName_example" // string | Human-readable label that identifies the Federated Database to update.
    skipRoleValidation := true // bool | Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.
    apiAtlasDataLakeTenantView := *openapiclient.NewApiAtlasDataLakeTenantView() // ApiAtlasDataLakeTenantView | Details of one Federated Database to update in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataFederationApi.UpdateOneFederatedDatabaseInOneProject(context.Background(), groupId, tenantName).SkipRoleValidation(skipRoleValidation).ApiAtlasDataLakeTenantView(apiAtlasDataLakeTenantView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataFederationApi.UpdateOneFederatedDatabaseInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneFederatedDatabaseInOneProject`: ApiAtlasDataLakeTenantView
    fmt.Fprintf(os.Stdout, "Response from `DataFederationApi.UpdateOneFederatedDatabaseInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**tenantName** | **string** | Human-readable label that identifies the Federated Database to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneFederatedDatabaseInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipRoleValidation** | **bool** | Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check. | 
 **apiAtlasDataLakeTenantView** | [**ApiAtlasDataLakeTenantView**](ApiAtlasDataLakeTenantView.md) | Details of one Federated Database to update in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDataLakeTenantView**](ApiAtlasDataLakeTenantView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

