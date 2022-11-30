# \EventsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnAllEventsFromOneOrganization**](EventsApi.md#ReturnAllEventsFromOneOrganization) | **Get** /api/atlas/v1.0/orgs/{orgId}/events | Return All Events from One Organization
[**ReturnAllEventsFromOneProject**](EventsApi.md#ReturnAllEventsFromOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/events | Return All Events from One Project
[**ReturnOneEventFromOneOrganization**](EventsApi.md#ReturnOneEventFromOneOrganization) | **Get** /api/atlas/v1.0/orgs/{orgId}/events/{eventId} | Return One Event from One Organization
[**ReturnOneEventFromOneProject**](EventsApi.md#ReturnOneEventFromOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/events/{eventId} | Return One Event from One Project



## ReturnAllEventsFromOneOrganization

> []ApiEventView ReturnAllEventsFromOneOrganization(ctx, orgId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).EventType(eventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()

Return All Events from One Organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    orgId := "4888442a3354817a7320eb61" // string | Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    eventType := "eventType_example" // string | Category of incident recorded at this moment in time. (optional)
    includeRaw := true // bool | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. (optional) (default to false)
    maxDate := time.Now() // time.Time | Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC. (optional)
    minDate := time.Now() // time.Time | Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ReturnAllEventsFromOneOrganization(context.Background(), orgId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).EventType(eventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ReturnAllEventsFromOneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllEventsFromOneOrganization`: []ApiEventView
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ReturnAllEventsFromOneOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllEventsFromOneOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **eventType** | **string** | Category of incident recorded at this moment in time. | 
 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]
 **maxDate** | **time.Time** | Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC. | 
 **minDate** | **time.Time** | Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC. | 

### Return type

[**[]ApiEventView**](ApiEventView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllEventsFromOneProject

> PaginatedEventView ReturnAllEventsFromOneProject(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).ClusterNames(clusterNames).EventType(eventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()

Return All Events from One Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    clusterNames := []string{"Inner_example"} // []string | Human-readable label that identifies the cluster. (optional)
    eventType := "eventType_example" // string | Category of incident recorded at this moment in time. (optional)
    includeRaw := true // bool | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. (optional) (default to false)
    maxDate := time.Now() // time.Time | Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC. (optional)
    minDate := time.Now() // time.Time | Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ReturnAllEventsFromOneProject(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).ClusterNames(clusterNames).EventType(eventType).IncludeRaw(includeRaw).MaxDate(maxDate).MinDate(minDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ReturnAllEventsFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllEventsFromOneProject`: PaginatedEventView
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ReturnAllEventsFromOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllEventsFromOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **clusterNames** | **[]string** | Human-readable label that identifies the cluster. | 
 **eventType** | **string** | Category of incident recorded at this moment in time. | 
 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]
 **maxDate** | **time.Time** | Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC. | 
 **minDate** | **time.Time** | Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC. | 

### Return type

[**PaginatedEventView**](PaginatedEventView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneEventFromOneOrganization

> ApiEventView ReturnOneEventFromOneOrganization(ctx, orgId, eventId).Envelope(envelope).Pretty(pretty).IncludeRaw(includeRaw).Execute()

Return One Event from One Organization



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
    eventId := "eventId_example" // string | Unique 24-hexadecimal digit string that identifies the event that you want to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    includeRaw := true // bool | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ReturnOneEventFromOneOrganization(context.Background(), orgId, eventId).Envelope(envelope).Pretty(pretty).IncludeRaw(includeRaw).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ReturnOneEventFromOneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneEventFromOneOrganization`: ApiEventView
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ReturnOneEventFromOneOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**eventId** | **string** | Unique 24-hexadecimal digit string that identifies the event that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneEventFromOneOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]

### Return type

[**ApiEventView**](ApiEventView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneEventFromOneProject

> ApiEventView ReturnOneEventFromOneProject(ctx, groupId, eventId).Envelope(envelope).Pretty(pretty).IncludeRaw(includeRaw).Execute()

Return One Event from One Project



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
    eventId := "eventId_example" // string | Unique 24-hexadecimal digit string that identifies the event that you want to return.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    includeRaw := true // bool | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ReturnOneEventFromOneProject(context.Background(), groupId, eventId).Envelope(envelope).Pretty(pretty).IncludeRaw(includeRaw).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ReturnOneEventFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneEventFromOneProject`: ApiEventView
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ReturnOneEventFromOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**eventId** | **string** | Unique 24-hexadecimal digit string that identifies the event that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneEventFromOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **includeRaw** | **bool** | Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event. | [default to false]

### Return type

[**ApiEventView**](ApiEventView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

