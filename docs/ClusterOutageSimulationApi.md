# \ClusterOutageSimulationApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndOutageSimulation**](ClusterOutageSimulationApi.md#EndOutageSimulation) | **Delete** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/outageSimulation | End an Outage Simulation
[**GetOutageSimulation**](ClusterOutageSimulationApi.md#GetOutageSimulation) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/outageSimulation | Return One Outage Simulation
[**StartOutageSimulation**](ClusterOutageSimulationApi.md#StartOutageSimulation) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/outageSimulation | Start an Outage Simulation



## EndOutageSimulation

> ApiAtlasClusterOutageSimulationView EndOutageSimulation(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

End an Outage Simulation



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster that is undergoing outage simulation.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterOutageSimulationApi.EndOutageSimulation(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterOutageSimulationApi.EndOutageSimulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndOutageSimulation`: ApiAtlasClusterOutageSimulationView
    fmt.Fprintf(os.Stdout, "Response from `ClusterOutageSimulationApi.EndOutageSimulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that is undergoing outage simulation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndOutageSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasClusterOutageSimulationView**](ApiAtlasClusterOutageSimulationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutageSimulation

> ApiAtlasClusterOutageSimulationView GetOutageSimulation(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return One Outage Simulation



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster that is undergoing outage simulation.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterOutageSimulationApi.GetOutageSimulation(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterOutageSimulationApi.GetOutageSimulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutageSimulation`: ApiAtlasClusterOutageSimulationView
    fmt.Fprintf(os.Stdout, "Response from `ClusterOutageSimulationApi.GetOutageSimulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster that is undergoing outage simulation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutageSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasClusterOutageSimulationView**](ApiAtlasClusterOutageSimulationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartOutageSimulation

> ApiAtlasClusterOutageSimulationView StartOutageSimulation(ctx, groupId, clusterName).ApiAtlasClusterOutageSimulationView(apiAtlasClusterOutageSimulationView).Envelope(envelope).Pretty(pretty).Execute()

Start an Outage Simulation



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster to undergo an outage simulation.
    apiAtlasClusterOutageSimulationView := *openapiclient.NewApiAtlasClusterOutageSimulationView() // ApiAtlasClusterOutageSimulationView | Describes the outage simulation.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterOutageSimulationApi.StartOutageSimulation(context.Background(), groupId, clusterName).ApiAtlasClusterOutageSimulationView(apiAtlasClusterOutageSimulationView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterOutageSimulationApi.StartOutageSimulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartOutageSimulation`: ApiAtlasClusterOutageSimulationView
    fmt.Fprintf(os.Stdout, "Response from `ClusterOutageSimulationApi.StartOutageSimulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster to undergo an outage simulation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartOutageSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasClusterOutageSimulationView** | [**ApiAtlasClusterOutageSimulationView**](ApiAtlasClusterOutageSimulationView.md) | Describes the outage simulation. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasClusterOutageSimulationView**](ApiAtlasClusterOutageSimulationView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

