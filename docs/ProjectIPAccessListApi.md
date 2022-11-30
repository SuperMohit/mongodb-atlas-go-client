# \ProjectIPAccessListApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEntriesToProjectIpAccessList**](ProjectIPAccessListApi.md#AddEntriesToProjectIpAccessList) | **Post** /api/atlas/v1.0/groups/{groupId}/accessList | Add Entries to Project IP Access List
[**GetAtlasNetworkPermissionEntryStatus**](ProjectIPAccessListApi.md#GetAtlasNetworkPermissionEntryStatus) | **Get** /api/atlas/v1.0/groups/{groupId}/accessList/{entryValue}/status | Return Status of One Project IP Access List Entry
[**RemoveOneEntryFromOneProjectIpAccessList**](ProjectIPAccessListApi.md#RemoveOneEntryFromOneProjectIpAccessList) | **Delete** /api/atlas/v1.0/groups/{groupId}/accessList/{entryValue} | Remove One Entry from One Project IP Access List
[**ReturnOneProjectIpAccessListEntry**](ProjectIPAccessListApi.md#ReturnOneProjectIpAccessListEntry) | **Get** /api/atlas/v1.0/groups/{groupId}/accessList/{entryValue} | Return One Project IP Access List Entry
[**ReturnProjectIpAccessList**](ProjectIPAccessListApi.md#ReturnProjectIpAccessList) | **Get** /api/atlas/v1.0/groups/{groupId}/accessList | Return Project IP Access List



## AddEntriesToProjectIpAccessList

> PaginatedNetworkAccessView AddEntriesToProjectIpAccessList(ctx, groupId).ApiAtlasNetworkPermissionEntryView(apiAtlasNetworkPermissionEntryView).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Add Entries to Project IP Access List



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
    apiAtlasNetworkPermissionEntryView := []openapiclient.ApiAtlasNetworkPermissionEntryView{*openapiclient.NewApiAtlasNetworkPermissionEntryView([]openapiclient.Link{*openapiclient.NewLink()})} // []ApiAtlasNetworkPermissionEntryView | One or more access list entries to add to the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectIPAccessListApi.AddEntriesToProjectIpAccessList(context.Background(), groupId).ApiAtlasNetworkPermissionEntryView(apiAtlasNetworkPermissionEntryView).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectIPAccessListApi.AddEntriesToProjectIpAccessList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEntriesToProjectIpAccessList`: PaginatedNetworkAccessView
    fmt.Fprintf(os.Stdout, "Response from `ProjectIPAccessListApi.AddEntriesToProjectIpAccessList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEntriesToProjectIpAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasNetworkPermissionEntryView** | [**[]ApiAtlasNetworkPermissionEntryView**](ApiAtlasNetworkPermissionEntryView.md) | One or more access list entries to add to the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedNetworkAccessView**](PaginatedNetworkAccessView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAtlasNetworkPermissionEntryStatus

> ApiAtlasNetworkPermissionEntryStatusView GetAtlasNetworkPermissionEntryStatus(ctx, groupId, entryValue).Envelope(envelope).Pretty(pretty).Execute()

Return Status of One Project IP Access List Entry



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
    entryValue := "IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8" // string | Network address or cloud provider security construct that identifies which project access list entry to be verified.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectIPAccessListApi.GetAtlasNetworkPermissionEntryStatus(context.Background(), groupId, entryValue).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectIPAccessListApi.GetAtlasNetworkPermissionEntryStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAtlasNetworkPermissionEntryStatus`: ApiAtlasNetworkPermissionEntryStatusView
    fmt.Fprintf(os.Stdout, "Response from `ProjectIPAccessListApi.GetAtlasNetworkPermissionEntryStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**entryValue** | **string** | Network address or cloud provider security construct that identifies which project access list entry to be verified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAtlasNetworkPermissionEntryStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasNetworkPermissionEntryStatusView**](ApiAtlasNetworkPermissionEntryStatusView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneEntryFromOneProjectIpAccessList

> RemoveOneEntryFromOneProjectIpAccessList(ctx, groupId, entryValue).Envelope(envelope).Pretty(pretty).Execute()

Remove One Entry from One Project IP Access List



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
    entryValue := "IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8" // string | Access list entry that you want to remove from the project's IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (`/`) with its URL-encoded value (`%2F`). When you remove an entry from the IP access list, existing connections from the removed address or addresses may remain open for a variable amount of time. The amount of time it takes MongoDB Cloud to close the connection depends upon several factors, including:  - how your application established the connection, - how MongoDB Cloud or the driver using the address behaves, and - which protocol (like TCP or UDP) the connection uses.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectIPAccessListApi.RemoveOneEntryFromOneProjectIpAccessList(context.Background(), groupId, entryValue).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectIPAccessListApi.RemoveOneEntryFromOneProjectIpAccessList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**entryValue** | **string** | Access list entry that you want to remove from the project&#39;s IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (&#x60;/&#x60;) with its URL-encoded value (&#x60;%2F&#x60;). When you remove an entry from the IP access list, existing connections from the removed address or addresses may remain open for a variable amount of time. The amount of time it takes MongoDB Cloud to close the connection depends upon several factors, including:  - how your application established the connection, - how MongoDB Cloud or the driver using the address behaves, and - which protocol (like TCP or UDP) the connection uses. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneEntryFromOneProjectIpAccessListRequest struct via the builder pattern


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


## ReturnOneProjectIpAccessListEntry

> ApiAtlasNetworkPermissionEntryView ReturnOneProjectIpAccessListEntry(ctx, groupId, entryValue).Envelope(envelope).Pretty(pretty).Execute()

Return One Project IP Access List Entry



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
    entryValue := "IPv4: 192.0.2.0%2F24 or IPv6: 2001:db8:85a3:8d3:1319:8a2e:370:7348 or IPv4 CIDR: 198.51.100.0%2f24 or IPv6 CIDR: 2001:db8::%2f58 or AWS SG: sg-903004f8" // string | Access list entry that you want to return from the project's IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (`/`) with its URL-encoded value (`%2F`).
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectIPAccessListApi.ReturnOneProjectIpAccessListEntry(context.Background(), groupId, entryValue).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectIPAccessListApi.ReturnOneProjectIpAccessListEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneProjectIpAccessListEntry`: ApiAtlasNetworkPermissionEntryView
    fmt.Fprintf(os.Stdout, "Response from `ProjectIPAccessListApi.ReturnOneProjectIpAccessListEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**entryValue** | **string** | Access list entry that you want to return from the project&#39;s IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (&#x60;/&#x60;) with its URL-encoded value (&#x60;%2F&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneProjectIpAccessListEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasNetworkPermissionEntryView**](ApiAtlasNetworkPermissionEntryView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnProjectIpAccessList

> PaginatedNetworkAccessView ReturnProjectIpAccessList(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return Project IP Access List



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
    resp, r, err := apiClient.ProjectIPAccessListApi.ReturnProjectIpAccessList(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectIPAccessListApi.ReturnProjectIpAccessList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnProjectIpAccessList`: PaginatedNetworkAccessView
    fmt.Fprintf(os.Stdout, "Response from `ProjectIPAccessListApi.ReturnProjectIpAccessList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnProjectIpAccessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedNetworkAccessView**](PaginatedNetworkAccessView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

