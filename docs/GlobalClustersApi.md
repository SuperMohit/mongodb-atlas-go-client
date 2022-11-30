# \GlobalClustersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOneCustomZoneMappingToOneGlobalCluster**](GlobalClustersApi.md#AddOneCustomZoneMappingToOneGlobalCluster) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Add Custom Zone Mappings to One Global Cluster
[**CreateOneManagedNamespaceInOneGlobalCluster**](GlobalClustersApi.md#CreateOneManagedNamespaceInOneGlobalCluster) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Create One Managed Namespace in One Global Cluster
[**DeleteAllCustomZoneMappings**](GlobalClustersApi.md#DeleteAllCustomZoneMappings) | **Delete** /api/atlas/v1.5/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Remove All Custom Zone Mappings from One Global Multi-Cloud Cluster
[**DeleteManagedNamespace**](GlobalClustersApi.md#DeleteManagedNamespace) | **Delete** /api/atlas/v1.5/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Remove One Managed Namespace from One Global Multi-Cloud Cluster
[**RemoveAllCustomZoneMappingsFromOneGlobalCluster**](GlobalClustersApi.md#RemoveAllCustomZoneMappingsFromOneGlobalCluster) | **Delete** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Remove All Custom Zone Mappings from One Global Cluster
[**RemoveOneManagedNamespaceFromOneGlobalCluster**](GlobalClustersApi.md#RemoveOneManagedNamespaceFromOneGlobalCluster) | **Delete** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Remove One Managed Namespace from One Global Cluster
[**ReturnAllGlobalClustersData**](GlobalClustersApi.md#ReturnAllGlobalClustersData) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/globalWrites | Return All Global Clusters Data



## AddOneCustomZoneMappingToOneGlobalCluster

> ApiAtlasGeoShardingView AddOneCustomZoneMappingToOneGlobalCluster(ctx, groupId, clusterName).ApiAtlasCustomZoneMappingsView(apiAtlasCustomZoneMappingsView).Envelope(envelope).Pretty(pretty).Execute()

Add Custom Zone Mappings to One Global Cluster



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies this global cluster.
    apiAtlasCustomZoneMappingsView := *openapiclient.NewApiAtlasCustomZoneMappingsView([]openapiclient.ApiAtlasZoneMappingView{*openapiclient.NewApiAtlasZoneMappingView("Location_example", "Zone_example")}) // ApiAtlasCustomZoneMappingsView | Custom zone mapping to add to the specified global cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalClustersApi.AddOneCustomZoneMappingToOneGlobalCluster(context.Background(), groupId, clusterName).ApiAtlasCustomZoneMappingsView(apiAtlasCustomZoneMappingsView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.AddOneCustomZoneMappingToOneGlobalCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOneCustomZoneMappingToOneGlobalCluster`: ApiAtlasGeoShardingView
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.AddOneCustomZoneMappingToOneGlobalCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies this global cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOneCustomZoneMappingToOneGlobalClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasCustomZoneMappingsView** | [**ApiAtlasCustomZoneMappingsView**](ApiAtlasCustomZoneMappingsView.md) | Custom zone mapping to add to the specified global cluster. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasGeoShardingView**](ApiAtlasGeoShardingView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOneManagedNamespaceInOneGlobalCluster

> ApiAtlasGeoShardingView CreateOneManagedNamespaceInOneGlobalCluster(ctx, groupId, clusterName).ApiAtlasManagedNamespacesView(apiAtlasManagedNamespacesView).Envelope(envelope).Pretty(pretty).Execute()

Create One Managed Namespace in One Global Cluster



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies this global cluster.
    apiAtlasManagedNamespacesView := *openapiclient.NewApiAtlasManagedNamespacesView("Collection_example", "CustomShardKey_example", "Db_example") // ApiAtlasManagedNamespacesView | Managed namespace to create within the specified global cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalClustersApi.CreateOneManagedNamespaceInOneGlobalCluster(context.Background(), groupId, clusterName).ApiAtlasManagedNamespacesView(apiAtlasManagedNamespacesView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.CreateOneManagedNamespaceInOneGlobalCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneManagedNamespaceInOneGlobalCluster`: ApiAtlasGeoShardingView
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.CreateOneManagedNamespaceInOneGlobalCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies this global cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneManagedNamespaceInOneGlobalClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasManagedNamespacesView** | [**ApiAtlasManagedNamespacesView**](ApiAtlasManagedNamespacesView.md) | Managed namespace to create within the specified global cluster. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasGeoShardingView**](ApiAtlasGeoShardingView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllCustomZoneMappings

> ApiAtlasGeoShardingView DeleteAllCustomZoneMappings(ctx, groupId, clusterName).Envelope(envelope).Execute()

Remove All Custom Zone Mappings from One Global Multi-Cloud Cluster



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies this advanced cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalClustersApi.DeleteAllCustomZoneMappings(context.Background(), groupId, clusterName).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.DeleteAllCustomZoneMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllCustomZoneMappings`: ApiAtlasGeoShardingView
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.DeleteAllCustomZoneMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies this advanced cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllCustomZoneMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiAtlasGeoShardingView**](ApiAtlasGeoShardingView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteManagedNamespace

> GeoShardingView DeleteManagedNamespace(ctx, clusterName, groupId).Envelope(envelope).Pretty(pretty).Db(db).Collection(collection).Execute()

Remove One Managed Namespace from One Global Multi-Cloud Cluster



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies this advanced cluster.
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    db := "db_example" // string | Human-readable label that identifies the database that contains the collection. (optional)
    collection := "collection_example" // string | Human-readable label that identifies the collection associated with the managed namespace. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalClustersApi.DeleteManagedNamespace(context.Background(), clusterName, groupId).Envelope(envelope).Pretty(pretty).Db(db).Collection(collection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.DeleteManagedNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteManagedNamespace`: GeoShardingView
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.DeleteManagedNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Human-readable label that identifies this advanced cluster. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **db** | **string** | Human-readable label that identifies the database that contains the collection. | 
 **collection** | **string** | Human-readable label that identifies the collection associated with the managed namespace. | 

### Return type

[**GeoShardingView**](GeoShardingView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAllCustomZoneMappingsFromOneGlobalCluster

> ApiAtlasGeoShardingView RemoveAllCustomZoneMappingsFromOneGlobalCluster(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Remove All Custom Zone Mappings from One Global Cluster



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies this global cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalClustersApi.RemoveAllCustomZoneMappingsFromOneGlobalCluster(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.RemoveAllCustomZoneMappingsFromOneGlobalCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveAllCustomZoneMappingsFromOneGlobalCluster`: ApiAtlasGeoShardingView
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.RemoveAllCustomZoneMappingsFromOneGlobalCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies this global cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAllCustomZoneMappingsFromOneGlobalClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasGeoShardingView**](ApiAtlasGeoShardingView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneManagedNamespaceFromOneGlobalCluster

> ApiAtlasGeoShardingView RemoveOneManagedNamespaceFromOneGlobalCluster(ctx, groupId, clusterName).Envelope(envelope).Db(db).Collection(collection).Execute()

Remove One Managed Namespace from One Global Cluster



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies this global cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    db := "db_example" // string | Human-readable label that identifies the database that contains the collection. (optional)
    collection := "collection_example" // string | Human-readable label that identifies the collection associated with the managed namespace. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalClustersApi.RemoveOneManagedNamespaceFromOneGlobalCluster(context.Background(), groupId, clusterName).Envelope(envelope).Db(db).Collection(collection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.RemoveOneManagedNamespaceFromOneGlobalCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveOneManagedNamespaceFromOneGlobalCluster`: ApiAtlasGeoShardingView
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.RemoveOneManagedNamespaceFromOneGlobalCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies this global cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneManagedNamespaceFromOneGlobalClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **db** | **string** | Human-readable label that identifies the database that contains the collection. | 
 **collection** | **string** | Human-readable label that identifies the collection associated with the managed namespace. | 

### Return type

[**ApiAtlasGeoShardingView**](ApiAtlasGeoShardingView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllGlobalClustersData

> ApiAtlasGeoShardingView ReturnAllGlobalClustersData(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return All Global Clusters Data



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies this global cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalClustersApi.ReturnAllGlobalClustersData(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalClustersApi.ReturnAllGlobalClustersData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllGlobalClustersData`: ApiAtlasGeoShardingView
    fmt.Fprintf(os.Stdout, "Response from `GlobalClustersApi.ReturnAllGlobalClustersData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies this global cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllGlobalClustersDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasGeoShardingView**](ApiAtlasGeoShardingView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

