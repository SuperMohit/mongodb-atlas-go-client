# \DataLakePipelinesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneDataLakePipeline**](DataLakePipelinesApi.md#CreateOneDataLakePipeline) | **Post** /api/atlas/v1.0/groups/{groupId}/pipelines | Create One Data Lake Pipeline
[**DeletePipelineRunDataset**](DataLakePipelinesApi.md#DeletePipelineRunDataset) | **Delete** /api/atlas/v1.0/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Delete Pipeline Run Dataset
[**PauseOnePipelineInOneProject**](DataLakePipelinesApi.md#PauseOnePipelineInOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/pipelines/{pipelineName}/pause | Pause One Data Lake Pipeline
[**RemoveOneDataLakePipeline**](DataLakePipelinesApi.md#RemoveOneDataLakePipeline) | **Delete** /api/atlas/v1.0/groups/{groupId}/pipelines/{pipelineName} | Remove One Data Lake Pipeline
[**ResumeOnePipelineInOneProject**](DataLakePipelinesApi.md#ResumeOnePipelineInOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/pipelines/{pipelineName}/resume | Resume One Data Lake Pipeline
[**ReturnAllDataLakePipelineRunsFromOneProject**](DataLakePipelinesApi.md#ReturnAllDataLakePipelineRunsFromOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/pipelines/{pipelineName}/runs | Return All Data Lake Pipeline Runs from One Project
[**ReturnAllDataLakePipelinesFromOneProject**](DataLakePipelinesApi.md#ReturnAllDataLakePipelinesFromOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/pipelines | Return All Data Lake Pipelines from One Project
[**ReturnAvailableSchedulesForPipeline**](DataLakePipelinesApi.md#ReturnAvailableSchedulesForPipeline) | **Get** /api/atlas/v1.0/groups/{groupId}/pipelines/{pipelineName}/availableSchedules | Return Available Ingestion Schedules for One Data Lake Pipeline
[**ReturnAvailableSnapshotsForPipeline**](DataLakePipelinesApi.md#ReturnAvailableSnapshotsForPipeline) | **Get** /api/atlas/v1.0/groups/{groupId}/pipelines/{pipelineName}/availableSnapshots | Return Available Backup Snapshots for One Data Lake Pipeline
[**ReturnOnePipelineInOneProject**](DataLakePipelinesApi.md#ReturnOnePipelineInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/pipelines/{pipelineName} | Return One Data Lake Pipeline
[**ReturnOnePipelineRunInOneProject**](DataLakePipelinesApi.md#ReturnOnePipelineRunInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Return One Data Lake Pipeline Run
[**TriggerOneOnDemandSnapshotIngestion**](DataLakePipelinesApi.md#TriggerOneOnDemandSnapshotIngestion) | **Post** /api/atlas/v1.0/groups/{groupId}/pipelines/{pipelineName}/trigger | Trigger on demand snapshot ingestion
[**UpdateOneDataLakePipeline**](DataLakePipelinesApi.md#UpdateOneDataLakePipeline) | **Patch** /api/atlas/v1.0/groups/{groupId}/pipelines/{pipelineName} | Update One Data Lake Pipeline



## CreateOneDataLakePipeline

> ApiAtlasIngestionPipelineView CreateOneDataLakePipeline(ctx, groupId).ApiAtlasIngestionPipelineView(apiAtlasIngestionPipelineView).Envelope(envelope).Execute()

Create One Data Lake Pipeline



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
    apiAtlasIngestionPipelineView := *openapiclient.NewApiAtlasIngestionPipelineView() // ApiAtlasIngestionPipelineView | Creates one Data Lake Pipeline.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.CreateOneDataLakePipeline(context.Background(), groupId).ApiAtlasIngestionPipelineView(apiAtlasIngestionPipelineView).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.CreateOneDataLakePipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneDataLakePipeline`: ApiAtlasIngestionPipelineView
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesApi.CreateOneDataLakePipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneDataLakePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasIngestionPipelineView** | [**ApiAtlasIngestionPipelineView**](ApiAtlasIngestionPipelineView.md) | Creates one Data Lake Pipeline. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiAtlasIngestionPipelineView**](ApiAtlasIngestionPipelineView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipelineRunDataset

> DeletePipelineRunDataset(ctx, groupId, pipelineName, pipelineRunId).Envelope(envelope).Pretty(pretty).Execute()

Delete Pipeline Run Dataset



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
    pipelineName := "pipelineName_example" // string | Human-readable label that identifies the Data Lake Pipeline.
    pipelineRunId := "6300603d5fce7d9fd1543170" // string | Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.DeletePipelineRunDataset(context.Background(), groupId, pipelineName, pipelineRunId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.DeletePipelineRunDataset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 
**pipelineRunId** | **string** | Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineRunDatasetRequest struct via the builder pattern


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


## PauseOnePipelineInOneProject

> ApiAtlasIngestionPipelineView PauseOnePipelineInOneProject(ctx, groupId, pipelineName).Envelope(envelope).Pretty(pretty).Execute()

Pause One Data Lake Pipeline



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
    pipelineName := "pipelineName_example" // string | Human-readable label that identifies the Data Lake Pipeline.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.PauseOnePipelineInOneProject(context.Background(), groupId, pipelineName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.PauseOnePipelineInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseOnePipelineInOneProject`: ApiAtlasIngestionPipelineView
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesApi.PauseOnePipelineInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseOnePipelineInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasIngestionPipelineView**](ApiAtlasIngestionPipelineView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneDataLakePipeline

> RemoveOneDataLakePipeline(ctx, groupId, pipelineName).Envelope(envelope).Execute()

Remove One Data Lake Pipeline



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
    pipelineName := "pipelineName_example" // string | Human-readable label that identifies the Data Lake Pipeline.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.RemoveOneDataLakePipeline(context.Background(), groupId, pipelineName).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.RemoveOneDataLakePipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneDataLakePipelineRequest struct via the builder pattern


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


## ResumeOnePipelineInOneProject

> ApiAtlasIngestionPipelineView ResumeOnePipelineInOneProject(ctx, groupId, pipelineName).Envelope(envelope).Pretty(pretty).Execute()

Resume One Data Lake Pipeline



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
    pipelineName := "pipelineName_example" // string | Human-readable label that identifies the Data Lake Pipeline.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.ResumeOnePipelineInOneProject(context.Background(), groupId, pipelineName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.ResumeOnePipelineInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeOnePipelineInOneProject`: ApiAtlasIngestionPipelineView
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesApi.ResumeOnePipelineInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeOnePipelineInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasIngestionPipelineView**](ApiAtlasIngestionPipelineView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllDataLakePipelineRunsFromOneProject

> PaginatedPipelineRunView ReturnAllDataLakePipelineRunsFromOneProject(ctx, groupId, pipelineName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).CreatedBefore(createdBefore).Execute()

Return All Data Lake Pipeline Runs from One Project



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
    pipelineName := "pipelineName_example" // string | Human-readable label that identifies the Data Lake Pipeline.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    createdBefore := time.Now() // time.Time | If specified, Atlas returns only Data Lake Pipeline runs initiated before this time and date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.ReturnAllDataLakePipelineRunsFromOneProject(context.Background(), groupId, pipelineName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).CreatedBefore(createdBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.ReturnAllDataLakePipelineRunsFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllDataLakePipelineRunsFromOneProject`: PaginatedPipelineRunView
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesApi.ReturnAllDataLakePipelineRunsFromOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllDataLakePipelineRunsFromOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **createdBefore** | **time.Time** | If specified, Atlas returns only Data Lake Pipeline runs initiated before this time and date. | 

### Return type

[**PaginatedPipelineRunView**](PaginatedPipelineRunView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllDataLakePipelinesFromOneProject

> []ApiAtlasIngestionPipelineView ReturnAllDataLakePipelinesFromOneProject(ctx, groupId).Envelope(envelope).Execute()

Return All Data Lake Pipelines from One Project



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.ReturnAllDataLakePipelinesFromOneProject(context.Background(), groupId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.ReturnAllDataLakePipelinesFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllDataLakePipelinesFromOneProject`: []ApiAtlasIngestionPipelineView
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesApi.ReturnAllDataLakePipelinesFromOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllDataLakePipelinesFromOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**[]ApiAtlasIngestionPipelineView**](ApiAtlasIngestionPipelineView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAvailableSchedulesForPipeline

> []ApiPolicyItemView ReturnAvailableSchedulesForPipeline(ctx, groupId, pipelineName).Envelope(envelope).Pretty(pretty).Execute()

Return Available Ingestion Schedules for One Data Lake Pipeline



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
    pipelineName := "pipelineName_example" // string | Human-readable label that identifies the Data Lake Pipeline.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.ReturnAvailableSchedulesForPipeline(context.Background(), groupId, pipelineName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.ReturnAvailableSchedulesForPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAvailableSchedulesForPipeline`: []ApiPolicyItemView
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesApi.ReturnAvailableSchedulesForPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAvailableSchedulesForPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**[]ApiPolicyItemView**](ApiPolicyItemView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAvailableSnapshotsForPipeline

> PaginatedBackupSnapshotView ReturnAvailableSnapshotsForPipeline(ctx, groupId, pipelineName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).CompletedAfter(completedAfter).Execute()

Return Available Backup Snapshots for One Data Lake Pipeline



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
    pipelineName := "pipelineName_example" // string | Human-readable label that identifies the Data Lake Pipeline.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    completedAfter := time.Now() // time.Time | Date and time after which MongoDB Cloud created the snapshot. If specified, MongoDB Cloud returns available backup snapshots created after this time and date only. This parameter expresses its value in the ISO 8601 timestamp format in UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.ReturnAvailableSnapshotsForPipeline(context.Background(), groupId, pipelineName).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).CompletedAfter(completedAfter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.ReturnAvailableSnapshotsForPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAvailableSnapshotsForPipeline`: PaginatedBackupSnapshotView
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesApi.ReturnAvailableSnapshotsForPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAvailableSnapshotsForPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **completedAfter** | **time.Time** | Date and time after which MongoDB Cloud created the snapshot. If specified, MongoDB Cloud returns available backup snapshots created after this time and date only. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | 

### Return type

[**PaginatedBackupSnapshotView**](PaginatedBackupSnapshotView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOnePipelineInOneProject

> ApiAtlasIngestionPipelineView ReturnOnePipelineInOneProject(ctx, groupId, pipelineName).Envelope(envelope).Pretty(pretty).Execute()

Return One Data Lake Pipeline



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
    pipelineName := "pipelineName_example" // string | Human-readable label that identifies the Data Lake Pipeline.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.ReturnOnePipelineInOneProject(context.Background(), groupId, pipelineName).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.ReturnOnePipelineInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOnePipelineInOneProject`: ApiAtlasIngestionPipelineView
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesApi.ReturnOnePipelineInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOnePipelineInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasIngestionPipelineView**](ApiAtlasIngestionPipelineView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOnePipelineRunInOneProject

> ApiAtlasIngestionPipelineRunView ReturnOnePipelineRunInOneProject(ctx, groupId, pipelineName, pipelineRunId).Envelope(envelope).Pretty(pretty).Execute()

Return One Data Lake Pipeline Run



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
    pipelineName := "pipelineName_example" // string | Human-readable label that identifies the Data Lake Pipeline.
    pipelineRunId := "6300603d5fce7d9fd1543170" // string | Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.ReturnOnePipelineRunInOneProject(context.Background(), groupId, pipelineName, pipelineRunId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.ReturnOnePipelineRunInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOnePipelineRunInOneProject`: ApiAtlasIngestionPipelineRunView
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesApi.ReturnOnePipelineRunInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 
**pipelineRunId** | **string** | Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOnePipelineRunInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasIngestionPipelineRunView**](ApiAtlasIngestionPipelineRunView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerOneOnDemandSnapshotIngestion

> ApiAtlasIngestionPipelineRunView TriggerOneOnDemandSnapshotIngestion(ctx, groupId, pipelineName).ApiAtlasTriggerIngestionRequestView(apiAtlasTriggerIngestionRequestView).Envelope(envelope).Pretty(pretty).Execute()

Trigger on demand snapshot ingestion



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
    pipelineName := "pipelineName_example" // string | Human-readable label that identifies the Data Lake Pipeline.
    apiAtlasTriggerIngestionRequestView := *openapiclient.NewApiAtlasTriggerIngestionRequestView("78e18e587716102f711a5bca") // ApiAtlasTriggerIngestionRequestView | Triggers a single ingestion run of a snapshot.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.TriggerOneOnDemandSnapshotIngestion(context.Background(), groupId, pipelineName).ApiAtlasTriggerIngestionRequestView(apiAtlasTriggerIngestionRequestView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.TriggerOneOnDemandSnapshotIngestion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TriggerOneOnDemandSnapshotIngestion`: ApiAtlasIngestionPipelineRunView
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesApi.TriggerOneOnDemandSnapshotIngestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerOneOnDemandSnapshotIngestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasTriggerIngestionRequestView** | [**ApiAtlasTriggerIngestionRequestView**](ApiAtlasTriggerIngestionRequestView.md) | Triggers a single ingestion run of a snapshot. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasIngestionPipelineRunView**](ApiAtlasIngestionPipelineRunView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneDataLakePipeline

> ApiAtlasIngestionPipelineView UpdateOneDataLakePipeline(ctx, groupId, pipelineName).ApiAtlasIngestionPipelineView(apiAtlasIngestionPipelineView).Envelope(envelope).Execute()

Update One Data Lake Pipeline



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
    pipelineName := "pipelineName_example" // string | Human-readable label that identifies the Data Lake Pipeline.
    apiAtlasIngestionPipelineView := *openapiclient.NewApiAtlasIngestionPipelineView() // ApiAtlasIngestionPipelineView | Updates one Data Lake Pipeline.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataLakePipelinesApi.UpdateOneDataLakePipeline(context.Background(), groupId, pipelineName).ApiAtlasIngestionPipelineView(apiAtlasIngestionPipelineView).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataLakePipelinesApi.UpdateOneDataLakePipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneDataLakePipeline`: ApiAtlasIngestionPipelineView
    fmt.Fprintf(os.Stdout, "Response from `DataLakePipelinesApi.UpdateOneDataLakePipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**pipelineName** | **string** | Human-readable label that identifies the Data Lake Pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneDataLakePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAtlasIngestionPipelineView** | [**ApiAtlasIngestionPipelineView**](ApiAtlasIngestionPipelineView.md) | Updates one Data Lake Pipeline. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiAtlasIngestionPipelineView**](ApiAtlasIngestionPipelineView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

