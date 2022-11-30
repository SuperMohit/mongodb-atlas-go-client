# \CloudBackupExportApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneCloudBackupSnapshotExportJob**](CloudBackupExportApi.md#CreateOneCloudBackupSnapshotExportJob) | **Post** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/exports | Create One Cloud Backup Snapshot Export Job
[**GrantAccessToAwsS3BucketForCloudBackupSnapshotExports**](CloudBackupExportApi.md#GrantAccessToAwsS3BucketForCloudBackupSnapshotExports) | **Post** /api/atlas/v1.0/groups/{groupId}/backup/exportBuckets | Grant Access to AWS S3 Bucket for Cloud Backup Snapshot Exports
[**ReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExports**](CloudBackupExportApi.md#ReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExports) | **Get** /api/atlas/v1.0/groups/{groupId}/backup/exportBuckets | Return All AWS S3 Buckets Used for Cloud Backup Snapshot Exports
[**ReturnAllCloudBackupSnapshotExportJobs**](CloudBackupExportApi.md#ReturnAllCloudBackupSnapshotExportJobs) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/exports | Return All Cloud Backup Snapshot Export Jobs
[**ReturnOneAwsS3BucketUsedForCloudBackupSnapshotExports**](CloudBackupExportApi.md#ReturnOneAwsS3BucketUsedForCloudBackupSnapshotExports) | **Get** /api/atlas/v1.0/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Return One AWS S3 Bucket Used for Cloud Backup Snapshot Exports
[**ReturnOneCloudBackupSnapshotExportJob**](CloudBackupExportApi.md#ReturnOneCloudBackupSnapshotExportJob) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/exports/{exportId} | Return One Cloud Backup Snapshot Export Job
[**RevokeAccessToAwsS3BucketForCloudBackupSnapshotExports**](CloudBackupExportApi.md#RevokeAccessToAwsS3BucketForCloudBackupSnapshotExports) | **Delete** /api/atlas/v1.0/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Revoke Access to AWS S3 Bucket for Cloud Backup Snapshot Exports



## CreateOneCloudBackupSnapshotExportJob

> ApiAtlasDiskBackupExportJobView CreateOneCloudBackupSnapshotExportJob(ctx, groupId, clusterName).Envelope(envelope).ApiAtlasDiskBackupExportJobRequestView(apiAtlasDiskBackupExportJobRequestView).Execute()

Create One Cloud Backup Snapshot Export Job



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    apiAtlasDiskBackupExportJobRequestView := *openapiclient.NewApiAtlasDiskBackupExportJobRequestView("d705ef7f03974e8363ffee5d", []openapiclient.Link{*openapiclient.NewLink()}, "78e18e587716102f711a5bca") // ApiAtlasDiskBackupExportJobRequestView |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupExportApi.CreateOneCloudBackupSnapshotExportJob(context.Background(), groupId, clusterName).Envelope(envelope).ApiAtlasDiskBackupExportJobRequestView(apiAtlasDiskBackupExportJobRequestView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupExportApi.CreateOneCloudBackupSnapshotExportJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneCloudBackupSnapshotExportJob`: ApiAtlasDiskBackupExportJobView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupExportApi.CreateOneCloudBackupSnapshotExportJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneCloudBackupSnapshotExportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **apiAtlasDiskBackupExportJobRequestView** | [**ApiAtlasDiskBackupExportJobRequestView**](ApiAtlasDiskBackupExportJobRequestView.md) |  | 

### Return type

[**ApiAtlasDiskBackupExportJobView**](ApiAtlasDiskBackupExportJobView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantAccessToAwsS3BucketForCloudBackupSnapshotExports

> ApiAtlasDiskBackupSnapshotAWSExportBucketView GrantAccessToAwsS3BucketForCloudBackupSnapshotExports(ctx, groupId).ApiAtlasDiskBackupSnapshotAWSExportBucketView(apiAtlasDiskBackupSnapshotAWSExportBucketView).Envelope(envelope).Pretty(pretty).Execute()

Grant Access to AWS S3 Bucket for Cloud Backup Snapshot Exports



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
    apiAtlasDiskBackupSnapshotAWSExportBucketView := *openapiclient.NewApiAtlasDiskBackupSnapshotAWSExportBucketView([]openapiclient.Link{*openapiclient.NewLink()}) // ApiAtlasDiskBackupSnapshotAWSExportBucketView | Grants MongoDB Cloud access to the specified AWS S3 bucket.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupExportApi.GrantAccessToAwsS3BucketForCloudBackupSnapshotExports(context.Background(), groupId).ApiAtlasDiskBackupSnapshotAWSExportBucketView(apiAtlasDiskBackupSnapshotAWSExportBucketView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupExportApi.GrantAccessToAwsS3BucketForCloudBackupSnapshotExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GrantAccessToAwsS3BucketForCloudBackupSnapshotExports`: ApiAtlasDiskBackupSnapshotAWSExportBucketView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupExportApi.GrantAccessToAwsS3BucketForCloudBackupSnapshotExports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantAccessToAwsS3BucketForCloudBackupSnapshotExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasDiskBackupSnapshotAWSExportBucketView** | [**ApiAtlasDiskBackupSnapshotAWSExportBucketView**](ApiAtlasDiskBackupSnapshotAWSExportBucketView.md) | Grants MongoDB Cloud access to the specified AWS S3 bucket. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDiskBackupSnapshotAWSExportBucketView**](ApiAtlasDiskBackupSnapshotAWSExportBucketView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExports

> PaginatedBackupSnapshotExportBucketView ReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExports(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Return All AWS S3 Buckets Used for Cloud Backup Snapshot Exports



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupExportApi.ReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExports(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupExportApi.ReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExports`: PaginatedBackupSnapshotExportBucketView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupExportApi.ReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedBackupSnapshotExportBucketView**](PaginatedBackupSnapshotExportBucketView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllCloudBackupSnapshotExportJobs

> []ApiAtlasDiskBackupExportJobView ReturnAllCloudBackupSnapshotExportJobs(ctx, groupId, clusterName).Envelope(envelope).Pretty(pretty).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Cloud Backup Snapshot Export Jobs



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupExportApi.ReturnAllCloudBackupSnapshotExportJobs(context.Background(), groupId, clusterName).Envelope(envelope).Pretty(pretty).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupExportApi.ReturnAllCloudBackupSnapshotExportJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllCloudBackupSnapshotExportJobs`: []ApiAtlasDiskBackupExportJobView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupExportApi.ReturnAllCloudBackupSnapshotExportJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllCloudBackupSnapshotExportJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**[]ApiAtlasDiskBackupExportJobView**](ApiAtlasDiskBackupExportJobView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneAwsS3BucketUsedForCloudBackupSnapshotExports

> ApiAtlasDiskBackupSnapshotAWSExportBucketView ReturnOneAwsS3BucketUsedForCloudBackupSnapshotExports(ctx, groupId, exportBucketId).Envelope(envelope).Execute()

Return One AWS S3 Bucket Used for Cloud Backup Snapshot Exports



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
    exportBucketId := "exportBucketId_example" // string | Unique string that identifies the AWS S3 bucket to which you export your snapshots.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupExportApi.ReturnOneAwsS3BucketUsedForCloudBackupSnapshotExports(context.Background(), groupId, exportBucketId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupExportApi.ReturnOneAwsS3BucketUsedForCloudBackupSnapshotExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneAwsS3BucketUsedForCloudBackupSnapshotExports`: ApiAtlasDiskBackupSnapshotAWSExportBucketView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupExportApi.ReturnOneAwsS3BucketUsedForCloudBackupSnapshotExports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**exportBucketId** | **string** | Unique string that identifies the AWS S3 bucket to which you export your snapshots. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneAwsS3BucketUsedForCloudBackupSnapshotExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiAtlasDiskBackupSnapshotAWSExportBucketView**](ApiAtlasDiskBackupSnapshotAWSExportBucketView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneCloudBackupSnapshotExportJob

> ApiAtlasDiskBackupExportJobView ReturnOneCloudBackupSnapshotExportJob(ctx, groupId, clusterName, exportId).Envelope(envelope).Execute()

Return One Cloud Backup Snapshot Export Job



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
    clusterName := "clusterName_example" // string | Human-readable label that identifies the cluster.
    exportId := "exportId_example" // string | Unique string that identifies the AWS S3 bucket to which you export your snapshots.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupExportApi.ReturnOneCloudBackupSnapshotExportJob(context.Background(), groupId, clusterName, exportId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupExportApi.ReturnOneCloudBackupSnapshotExportJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneCloudBackupSnapshotExportJob`: ApiAtlasDiskBackupExportJobView
    fmt.Fprintf(os.Stdout, "Response from `CloudBackupExportApi.ReturnOneCloudBackupSnapshotExportJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**clusterName** | **string** | Human-readable label that identifies the cluster. | 
**exportId** | **string** | Unique string that identifies the AWS S3 bucket to which you export your snapshots. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneCloudBackupSnapshotExportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiAtlasDiskBackupExportJobView**](ApiAtlasDiskBackupExportJobView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAccessToAwsS3BucketForCloudBackupSnapshotExports

> RevokeAccessToAwsS3BucketForCloudBackupSnapshotExports(ctx, groupId, exportBucketId).Envelope(envelope).Execute()

Revoke Access to AWS S3 Bucket for Cloud Backup Snapshot Exports



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
    exportBucketId := "exportBucketId_example" // string | Unique string that identifies the AWS S3 bucket to which you export your snapshots.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudBackupExportApi.RevokeAccessToAwsS3BucketForCloudBackupSnapshotExports(context.Background(), groupId, exportBucketId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudBackupExportApi.RevokeAccessToAwsS3BucketForCloudBackupSnapshotExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**exportBucketId** | **string** | Unique string that identifies the AWS S3 bucket to which you export your snapshots. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAccessToAwsS3BucketForCloudBackupSnapshotExportsRequest struct via the builder pattern


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

