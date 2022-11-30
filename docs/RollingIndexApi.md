# \RollingIndexApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneRollingIndex**](RollingIndexApi.md#CreateOneRollingIndex) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/index | Create One Rolling Index



## CreateOneRollingIndex

> CreateOneRollingIndex(ctx, groupId, clusterName).ApiIndexRequestView(apiIndexRequestView).Envelope(envelope).Pretty(pretty).Execute()

Create One Rolling Index



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster on which MongoDB Cloud creates an index.
    apiIndexRequestView := *openapiclient.NewApiIndexRequestView("Collection_example", "Db_example", []map[string]string{map[string]string{"key": "Inner_example"}}) // ApiIndexRequestView | Rolling index to create on the specified cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RollingIndexApi.CreateOneRollingIndex(context.Background(), groupId, clusterName).ApiIndexRequestView(apiIndexRequestView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RollingIndexApi.CreateOneRollingIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster on which MongoDB Cloud creates an index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneRollingIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiIndexRequestView** | [**ApiIndexRequestView**](ApiIndexRequestView.md) | Rolling index to create on the specified cluster. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

 (empty response body)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

