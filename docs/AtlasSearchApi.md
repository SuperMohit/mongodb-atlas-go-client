# \AtlasSearchApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneAtlasSearchIndex**](AtlasSearchApi.md#CreateOneAtlasSearchIndex) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes | Create One Atlas Search Index
[**RemoveOneAtlasSearchIndex**](AtlasSearchApi.md#RemoveOneAtlasSearchIndex) | **Delete** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Remove One Atlas Search Index
[**ReturnAllAtlasSearchIndexesForOneCollection**](AtlasSearchApi.md#ReturnAllAtlasSearchIndexesForOneCollection) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection
[**ReturnAllUserDefinedAnalyzersForOneCluster**](AtlasSearchApi.md#ReturnAllUserDefinedAnalyzersForOneCluster) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/analyzers | Return All User-Defined Analyzers for One Cluster
[**ReturnOneAtlasSearchIndex**](AtlasSearchApi.md#ReturnOneAtlasSearchIndex) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Return One Atlas Search Index
[**UpdateAllUserDefinedAnalyzersForOneCluster**](AtlasSearchApi.md#UpdateAllUserDefinedAnalyzersForOneCluster) | **Put** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/analyzers | Update All User-Defined Analyzers for One Cluster
[**UpdateOneAtlasSearchIndex**](AtlasSearchApi.md#UpdateOneAtlasSearchIndex) | **Patch** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Update One Atlas Search Index



## CreateOneAtlasSearchIndex

> ApiAtlasFTSIndexView CreateOneAtlasSearchIndex(ctx, groupId, clusterName).ApiAtlasFTSIndexView(apiAtlasFTSIndexView).Envelope(envelope).Pretty(pretty).Execute()

Create One Atlas Search Index



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
    clusterName := "clusterName_example" // string | Name of the cluster that contains the collection on which to create an Atlas Search index.
    apiAtlasFTSIndexView := *openapiclient.NewApiAtlasFTSIndexView("CollectionName_example", "Database_example", "Name_example") // ApiAtlasFTSIndexView | Creates one Atlas Search index on the specified collection.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AtlasSearchApi.CreateOneAtlasSearchIndex(context.Background(), groupId, clusterName).ApiAtlasFTSIndexView(apiAtlasFTSIndexView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.CreateOneAtlasSearchIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneAtlasSearchIndex`: ApiAtlasFTSIndexView
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.CreateOneAtlasSearchIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Name of the cluster that contains the collection on which to create an Atlas Search index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneAtlasSearchIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasFTSIndexView** | [**ApiAtlasFTSIndexView**](ApiAtlasFTSIndexView.md) | Creates one Atlas Search index on the specified collection. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasFTSIndexView**](ApiAtlasFTSIndexView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneAtlasSearchIndex

> RemoveOneAtlasSearchIndex(ctx, groupId, clusterName, indexId).Envelope(envelope).Pretty(pretty).Execute()

Remove One Atlas Search Index



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
    clusterName := "clusterName_example" // string | Name of the cluster that contains the database and collection with one or more Application Search indexes.
    indexId := "indexId_example" // string | Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AtlasSearchApi.RemoveOneAtlasSearchIndex(context.Background(), groupId, clusterName, indexId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.RemoveOneAtlasSearchIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Name of the cluster that contains the database and collection with one or more Application Search indexes. | 
**indexId** | **string** | Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneAtlasSearchIndexRequest struct via the builder pattern


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


## ReturnAllAtlasSearchIndexesForOneCollection

> []ApiAtlasFTSIndexView ReturnAllAtlasSearchIndexesForOneCollection(ctx, groupId, clusterName, collectionName, databaseName).Envelope(envelope).Pretty(pretty).Execute()

Return All Atlas Search Indexes for One Collection



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
    clusterName := "clusterName_example" // string | Name of the cluster that contains the collection with one or more Atlas Search indexes.
    collectionName := "collectionName_example" // string | Name of the collection that contains one or more Atlas Search indexes.
    databaseName := "databaseName_example" // string | Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AtlasSearchApi.ReturnAllAtlasSearchIndexesForOneCollection(context.Background(), groupId, clusterName, collectionName, databaseName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.ReturnAllAtlasSearchIndexesForOneCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllAtlasSearchIndexesForOneCollection`: []ApiAtlasFTSIndexView
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.ReturnAllAtlasSearchIndexesForOneCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 
**collectionName** | **string** | Name of the collection that contains one or more Atlas Search indexes. | 
**databaseName** | **string** | Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllAtlasSearchIndexesForOneCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiAtlasFTSIndexView**](ApiAtlasFTSIndexView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllUserDefinedAnalyzersForOneCluster

> []ApiAtlasFTSAnalyzersViewManual ReturnAllUserDefinedAnalyzersForOneCluster(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()

Return All User-Defined Analyzers for One Cluster



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
    clusterName := "clusterName_example" // string | Name of the cluster that contains the collection with one or more Atlas Search indexes.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AtlasSearchApi.ReturnAllUserDefinedAnalyzersForOneCluster(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.ReturnAllUserDefinedAnalyzersForOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllUserDefinedAnalyzersForOneCluster`: []ApiAtlasFTSAnalyzersViewManual
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.ReturnAllUserDefinedAnalyzersForOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllUserDefinedAnalyzersForOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiAtlasFTSAnalyzersViewManual**](ApiAtlasFTSAnalyzersViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneAtlasSearchIndex

> ApiAtlasFTSIndexView ReturnOneAtlasSearchIndex(ctx, groupId, clusterName, indexId).Envelope(envelope).Pretty(pretty).Execute()

Return One Atlas Search Index



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
    clusterName := "clusterName_example" // string | Name of the cluster that contains the collection with one or more Atlas Search indexes.
    indexId := "indexId_example" // string | Unique 24-hexadecimal digit string that identifies the Application Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AtlasSearchApi.ReturnOneAtlasSearchIndex(context.Background(), groupId, clusterName, indexId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.ReturnOneAtlasSearchIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneAtlasSearchIndex`: ApiAtlasFTSIndexView
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.ReturnOneAtlasSearchIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 
**indexId** | **string** | Unique 24-hexadecimal digit string that identifies the Application Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneAtlasSearchIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasFTSIndexView**](ApiAtlasFTSIndexView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAllUserDefinedAnalyzersForOneCluster

> []ApiAtlasFTSAnalyzersViewManual UpdateAllUserDefinedAnalyzersForOneCluster(ctx, groupId, clusterName).ApiAtlasFTSAnalyzersViewManual(apiAtlasFTSAnalyzersViewManual).Envelope(envelope).Pretty(pretty).Execute()

Update All User-Defined Analyzers for One Cluster



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
    clusterName := "clusterName_example" // string | Name of the cluster that contains the collection with one or more Atlas Search indexes.
    apiAtlasFTSAnalyzersViewManual := []openapiclient.ApiAtlasFTSAnalyzersViewManual{*openapiclient.NewApiAtlasFTSAnalyzersViewManual("Name_example", openapiclient.ApiAtlasFTSAnalyzersViewManual_tokenizer{TokenizeredgeGram: openapiclient.NewTokenizeredgeGram(int32(123), int32(123), "Type_example")})} // []ApiAtlasFTSAnalyzersViewManual | Updates one or more Atlas Search custom analyzers on the specified cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AtlasSearchApi.UpdateAllUserDefinedAnalyzersForOneCluster(context.Background(), groupId, clusterName).ApiAtlasFTSAnalyzersViewManual(apiAtlasFTSAnalyzersViewManual).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.UpdateAllUserDefinedAnalyzersForOneCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAllUserDefinedAnalyzersForOneCluster`: []ApiAtlasFTSAnalyzersViewManual
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.UpdateAllUserDefinedAnalyzersForOneCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Name of the cluster that contains the collection with one or more Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAllUserDefinedAnalyzersForOneClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasFTSAnalyzersViewManual** | [**[]ApiAtlasFTSAnalyzersViewManual**](ApiAtlasFTSAnalyzersViewManual.md) | Updates one or more Atlas Search custom analyzers on the specified cluster. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiAtlasFTSAnalyzersViewManual**](ApiAtlasFTSAnalyzersViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneAtlasSearchIndex

> ApiAtlasFTSIndexView UpdateOneAtlasSearchIndex(ctx, groupId, clusterName, indexId).ApiAtlasFTSIndexView(apiAtlasFTSIndexView).Envelope(envelope).Pretty(pretty).Execute()

Update One Atlas Search Index



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
    clusterName := "clusterName_example" // string | Name of the cluster that contains the collection whose Atlas Search index to update.
    indexId := "indexId_example" // string | Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
    apiAtlasFTSIndexView := *openapiclient.NewApiAtlasFTSIndexView("CollectionName_example", "Database_example", "Name_example") // ApiAtlasFTSIndexView | Details to update on the Atlas Search index.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AtlasSearchApi.UpdateOneAtlasSearchIndex(context.Background(), groupId, clusterName, indexId).ApiAtlasFTSIndexView(apiAtlasFTSIndexView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AtlasSearchApi.UpdateOneAtlasSearchIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneAtlasSearchIndex`: ApiAtlasFTSIndexView
    fmt.Fprintf(os.Stdout, "Response from `AtlasSearchApi.UpdateOneAtlasSearchIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Name of the cluster that contains the collection whose Atlas Search index to update. | 
**indexId** | **string** | Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneAtlasSearchIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiAtlasFTSIndexView** | [**ApiAtlasFTSIndexView**](ApiAtlasFTSIndexView.md) | Details to update on the Atlas Search index. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasFTSIndexView**](ApiAtlasFTSIndexView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

