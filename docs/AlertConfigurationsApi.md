# \AlertConfigurationsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneAlertConfigurationInOneProject**](AlertConfigurationsApi.md#CreateOneAlertConfigurationInOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/alertConfigs | Create One Alert Configuration in One Project
[**RemoveOneAlertConfigurationFromOneProject**](AlertConfigurationsApi.md#RemoveOneAlertConfigurationFromOneProject) | **Delete** /api/atlas/v1.0/groups/{groupId}/alertConfigs/{alertConfigId} | Remove One Alert Configuration from One Project
[**ReturnAlertConfigMatchersFieldNames**](AlertConfigurationsApi.md#ReturnAlertConfigMatchersFieldNames) | **Get** /api/atlas/v1.0/alertConfigs/matchers/fieldNames | Get All Alert Configuration Matchers Field Names
[**ReturnAllAlertConfigurationsForOneProject**](AlertConfigurationsApi.md#ReturnAllAlertConfigurationsForOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/alertConfigs | Return All Alert Configurations for One Project
[**ReturnAllOpenAlertsForAlertConfiguration**](AlertConfigurationsApi.md#ReturnAllOpenAlertsForAlertConfiguration) | **Get** /api/atlas/v1.0/groups/{groupId}/alertConfigs/{alertConfigId}/alerts | Return All Open Alerts for Alert Configuration
[**ReturnOneAlertConfigurationFromOneProject**](AlertConfigurationsApi.md#ReturnOneAlertConfigurationFromOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/alertConfigs/{alertConfigId} | Return One Alert Configuration from One Project
[**ToggleOneStateOfOneAlertConfigurationInOneProject**](AlertConfigurationsApi.md#ToggleOneStateOfOneAlertConfigurationInOneProject) | **Patch** /api/atlas/v1.0/groups/{groupId}/alertConfigs/{alertConfigId} | Toggle One State of One Alert Configuration in One Project
[**UpdateOneAlertConfigurationForOneProject**](AlertConfigurationsApi.md#UpdateOneAlertConfigurationForOneProject) | **Put** /api/atlas/v1.0/groups/{groupId}/alertConfigs/{alertConfigId} | Update One Alert Configuration for One Project



## CreateOneAlertConfigurationInOneProject

> ApiAlertConfigView CreateOneAlertConfigurationInOneProject(ctx, groupId).ApiAlertConfigView(apiAlertConfigView).Envelope(envelope).Pretty(pretty).Execute()

Create One Alert Configuration in One Project



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
    apiAlertConfigView := *openapiclient.NewApiAlertConfigView("EventTypeName_example", []openapiclient.Link{*openapiclient.NewLink()}) // ApiAlertConfigView | Creates one alert configuration for the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigurationsApi.CreateOneAlertConfigurationInOneProject(context.Background(), groupId).ApiAlertConfigView(apiAlertConfigView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.CreateOneAlertConfigurationInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneAlertConfigurationInOneProject`: ApiAlertConfigView
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.CreateOneAlertConfigurationInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneAlertConfigurationInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAlertConfigView** | [**ApiAlertConfigView**](ApiAlertConfigView.md) | Creates one alert configuration for the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAlertConfigView**](ApiAlertConfigView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneAlertConfigurationFromOneProject

> RemoveOneAlertConfigurationFromOneProject(ctx, groupId, alertConfigId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Alert Configuration from One Project



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
    alertConfigId := "adddde5af4ee4bd0b519375f" // string | Unique 24-hexadecimal digit string that identifies the alert configuration.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigurationsApi.RemoveOneAlertConfigurationFromOneProject(context.Background(), groupId, alertConfigId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.RemoveOneAlertConfigurationFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**alertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneAlertConfigurationFromOneProjectRequest struct via the builder pattern


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


## ReturnAlertConfigMatchersFieldNames

> []string ReturnAlertConfigMatchersFieldNames(ctx).Envelope(envelope).Pretty(pretty).Execute()

Get All Alert Configuration Matchers Field Names



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
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigurationsApi.ReturnAlertConfigMatchersFieldNames(context.Background()).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.ReturnAlertConfigMatchersFieldNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAlertConfigMatchersFieldNames`: []string
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.ReturnAlertConfigMatchersFieldNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnAlertConfigMatchersFieldNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

**[]string**

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllAlertConfigurationsForOneProject

> PaginatedAlertConfigView ReturnAllAlertConfigurationsForOneProject(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Alert Configurations for One Project



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
    resp, r, err := apiClient.AlertConfigurationsApi.ReturnAllAlertConfigurationsForOneProject(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.ReturnAllAlertConfigurationsForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllAlertConfigurationsForOneProject`: PaginatedAlertConfigView
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.ReturnAllAlertConfigurationsForOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllAlertConfigurationsForOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedAlertConfigView**](PaginatedAlertConfigView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllOpenAlertsForAlertConfiguration

> PaginatedAlertView ReturnAllOpenAlertsForAlertConfiguration(ctx, groupId, alertConfigId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Open Alerts for Alert Configuration



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
    alertConfigId := "adddde5af4ee4bd0b519375f" // string | Unique 24-hexadecimal digit string that identifies the alert configuration.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigurationsApi.ReturnAllOpenAlertsForAlertConfiguration(context.Background(), groupId, alertConfigId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.ReturnAllOpenAlertsForAlertConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllOpenAlertsForAlertConfiguration`: PaginatedAlertView
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.ReturnAllOpenAlertsForAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**alertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllOpenAlertsForAlertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedAlertView**](PaginatedAlertView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneAlertConfigurationFromOneProject

> ApiAlertConfigView ReturnOneAlertConfigurationFromOneProject(ctx, groupId, alertConfigId).Envelope(envelope).Pretty(pretty).Execute()

Return One Alert Configuration from One Project



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
    alertConfigId := "adddde5af4ee4bd0b519375f" // string | Unique 24-hexadecimal digit string that identifies the alert configuration.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigurationsApi.ReturnOneAlertConfigurationFromOneProject(context.Background(), groupId, alertConfigId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.ReturnOneAlertConfigurationFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneAlertConfigurationFromOneProject`: ApiAlertConfigView
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.ReturnOneAlertConfigurationFromOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**alertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneAlertConfigurationFromOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAlertConfigView**](ApiAlertConfigView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleOneStateOfOneAlertConfigurationInOneProject

> ApiAlertConfigView ToggleOneStateOfOneAlertConfigurationInOneProject(ctx, groupId, alertConfigId).Body(body).Envelope(envelope).Pretty(pretty).Execute()

Toggle One State of One Alert Configuration in One Project



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
    alertConfigId := "adddde5af4ee4bd0b519375f" // string | Unique 24-hexadecimal digit string that identifies the alert configuration that triggered this alert.
    body := true // bool | Enables or disables the specified alert configuration in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigurationsApi.ToggleOneStateOfOneAlertConfigurationInOneProject(context.Background(), groupId, alertConfigId).Body(body).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.ToggleOneStateOfOneAlertConfigurationInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleOneStateOfOneAlertConfigurationInOneProject`: ApiAlertConfigView
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.ToggleOneStateOfOneAlertConfigurationInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**alertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration that triggered this alert. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleOneStateOfOneAlertConfigurationInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **bool** | Enables or disables the specified alert configuration in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAlertConfigView**](ApiAlertConfigView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneAlertConfigurationForOneProject

> ApiAlertConfigView UpdateOneAlertConfigurationForOneProject(ctx, groupId, alertConfigId).ApiAlertConfigView(apiAlertConfigView).Envelope(envelope).Pretty(pretty).Execute()

Update One Alert Configuration for One Project



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
    alertConfigId := "adddde5af4ee4bd0b519375f" // string | Unique 24-hexadecimal digit string that identifies the alert configuration.
    apiAlertConfigView := *openapiclient.NewApiAlertConfigView("EventTypeName_example", []openapiclient.Link{*openapiclient.NewLink()}) // ApiAlertConfigView | Updates one alert configuration in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigurationsApi.UpdateOneAlertConfigurationForOneProject(context.Background(), groupId, alertConfigId).ApiAlertConfigView(apiAlertConfigView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigurationsApi.UpdateOneAlertConfigurationForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneAlertConfigurationForOneProject`: ApiAlertConfigView
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigurationsApi.UpdateOneAlertConfigurationForOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**alertConfigId** | **string** | Unique 24-hexadecimal digit string that identifies the alert configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneAlertConfigurationForOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAlertConfigView** | [**ApiAlertConfigView**](ApiAlertConfigView.md) | Updates one alert configuration in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAlertConfigView**](ApiAlertConfigView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

