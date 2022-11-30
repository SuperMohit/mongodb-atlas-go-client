# \MonitoringAndLogsApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadLogsForOneClusterHostInOneProject**](MonitoringAndLogsApi.md#DownloadLogsForOneClusterHostInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/clusters/{hostName}/logs/{logName}.gz | Download Logs for One Cluster Host in One Project
[**GetFTSIndexMeasurementsForNamespaceAndIndexName**](MonitoringAndLogsApi.md#GetFTSIndexMeasurementsForNamespaceAndIndexName) | **Get** /api/atlas/v1.0/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/{indexName}/measurements | Return Atlas Search Metrics for One Index in One Specified Namespace
[**GetFTSMetricTypes**](MonitoringAndLogsApi.md#GetFTSMetricTypes) | **Get** /api/atlas/v1.0/groups/{groupId}/hosts/{processId}/fts/metrics | Return All Atlas Search Metric Types for One Process
[**GetFTSNamespaceIndexMeasurements**](MonitoringAndLogsApi.md#GetFTSNamespaceIndexMeasurements) | **Get** /api/atlas/v1.0/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/measurements | Return All Atlas Search Index Metrics for One Namespace
[**GetFTSNonIndexMeasurements**](MonitoringAndLogsApi.md#GetFTSNonIndexMeasurements) | **Get** /api/atlas/v1.0/groups/{groupId}/hosts/{processId}/fts/metrics/measurements | Return Atlas Search Hardware and Status Metrics
[**GetHostLogs**](MonitoringAndLogsApi.md#GetHostLogs) | **Get** /api/atlas/v1.5/groups/{groupId}/clusters/{hostName}/logs/{logName}.gz | Download Logs for One Multi-Cloud Cluster Host in One Project
[**ReturnAllMongodbProcessesInOneProject**](MonitoringAndLogsApi.md#ReturnAllMongodbProcessesInOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/processes | Return All MongoDB Processes in One Project
[**ReturnAvailableDatabasesForOneMongodbProcess**](MonitoringAndLogsApi.md#ReturnAvailableDatabasesForOneMongodbProcess) | **Get** /api/atlas/v1.0/groups/{groupId}/processes/{processId}/databases | Return Available Databases for One MongoDB Process
[**ReturnAvailableDisksForOneMongodbProcess**](MonitoringAndLogsApi.md#ReturnAvailableDisksForOneMongodbProcess) | **Get** /api/atlas/v1.0/groups/{groupId}/processes/{processId}/disks | Return Available Disks for One MongoDB Process
[**ReturnDatabasesForOneMongodbProcess**](MonitoringAndLogsApi.md#ReturnDatabasesForOneMongodbProcess) | **Get** /api/atlas/v1.0/groups/{groupId}/processes/{processId}/databases/{databaseName} | Return One Database for a MongoDB Process
[**ReturnMeasurementsForOneMongodbProcess**](MonitoringAndLogsApi.md#ReturnMeasurementsForOneMongodbProcess) | **Get** /api/atlas/v1.0/groups/{groupId}/processes/{processId}/measurements | Return Measurements for One MongoDB Process
[**ReturnMeasurementsOfOneDatabaseForOneMongodbProcess**](MonitoringAndLogsApi.md#ReturnMeasurementsOfOneDatabaseForOneMongodbProcess) | **Get** /api/atlas/v1.0/groups/{groupId}/processes/{processId}/databases/{databaseName}/measurements | Return Measurements of One Database for One MongoDB Process
[**ReturnMeasurementsOfOneDisk**](MonitoringAndLogsApi.md#ReturnMeasurementsOfOneDisk) | **Get** /api/atlas/v1.0/groups/{groupId}/processes/{processId}/disks/{partitionName} | Return Measurements of One Disk
[**ReturnMeasurementsOfOneDiskForOneMongodbProcess**](MonitoringAndLogsApi.md#ReturnMeasurementsOfOneDiskForOneMongodbProcess) | **Get** /api/atlas/v1.0/groups/{groupId}/processes/{processId}/disks/{partitionName}/measurements | Return Measurements of One Disk for One MongoDB Process
[**ReturnOneMongodbProcessById**](MonitoringAndLogsApi.md#ReturnOneMongodbProcessById) | **Get** /api/atlas/v1.0/groups/{groupId}/processes/{processId} | Return One MongoDB Process by ID



## DownloadLogsForOneClusterHostInOneProject

> *os.File DownloadLogsForOneClusterHostInOneProject(ctx, groupId, hostName, logName).EndDate(endDate).StartDate(startDate).Execute()

Download Logs for One Cluster Host in One Project



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
    hostName := "hostName_example" // string | Human-readable label that identifies the host that stores the log files that you want to download.
    logName := "logName_example" // string | Human-readable label that identifies the log file that you want to return. To return audit logs, enable *Database Auditing* for the specified project.
    endDate := int64(789) // int64 | Date and time when the period specifies the inclusive ending point for the range of log messages to retrieve. This parameter expresses its value in the number of seconds that have elapsed since the UNIX epoch. (optional)
    startDate := int64(789) // int64 | Date and time when the period specifies the inclusive starting point for the range of log messages to retrieve. This parameter expresses its value in the number of seconds that have elapsed since the UNIX epoch. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.DownloadLogsForOneClusterHostInOneProject(context.Background(), groupId, hostName, logName).EndDate(endDate).StartDate(startDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.DownloadLogsForOneClusterHostInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadLogsForOneClusterHostInOneProject`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.DownloadLogsForOneClusterHostInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**hostName** | **string** | Human-readable label that identifies the host that stores the log files that you want to download. | 
**logName** | **string** | Human-readable label that identifies the log file that you want to return. To return audit logs, enable *Database Auditing* for the specified project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLogsForOneClusterHostInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **endDate** | **int64** | Date and time when the period specifies the inclusive ending point for the range of log messages to retrieve. This parameter expresses its value in the number of seconds that have elapsed since the UNIX epoch. | 
 **startDate** | **int64** | Date and time when the period specifies the inclusive starting point for the range of log messages to retrieve. This parameter expresses its value in the number of seconds that have elapsed since the UNIX epoch. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFTSIndexMeasurementsForNamespaceAndIndexName

> ApiMeasurementsIndexesView GetFTSIndexMeasurementsForNamespaceAndIndexName(ctx, processId, indexName, databaseName, collectionName, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Envelope(envelope).Execute()

Return Atlas Search Metrics for One Index in One Specified Namespace



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
    processId := "my.host.name.com:27017" // string | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
    indexName := "myindex" // string | Human-readable label that identifies the index.
    databaseName := "mydb" // string | Human-readable label that identifies the database.
    collectionName := "mycoll" // string | Human-readable label that identifies the collection.
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    granularity := "PT1M" // string | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.
    metrics := []string{"Metrics_example"} // []string | List that contains the measurements that MongoDB Atlas reports for the associated data series.
    period := "PT10H" // string | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. (optional)
    start := time.Now() // time.Time | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. (optional)
    end := time.Now() // time.Time | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. (optional)
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.GetFTSIndexMeasurementsForNamespaceAndIndexName(context.Background(), processId, indexName, databaseName, collectionName, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetFTSIndexMeasurementsForNamespaceAndIndexName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFTSIndexMeasurementsForNamespaceAndIndexName`: ApiMeasurementsIndexesView
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetFTSIndexMeasurementsForNamespaceAndIndexName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **string** | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests. | 
**indexName** | **string** | Human-readable label that identifies the index. | 
**databaseName** | **string** | Human-readable label that identifies the database. | 
**collectionName** | **string** | Human-readable label that identifies the collection. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFTSIndexMeasurementsForNamespaceAndIndexNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **granularity** | **string** | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. | 
 **metrics** | **[]string** | List that contains the measurements that MongoDB Atlas reports for the associated data series. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiMeasurementsIndexesView**](ApiMeasurementsIndexesView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFTSMetricTypes

> ApiFTSMetricsView GetFTSMetricTypes(ctx, processId, groupId).Envelope(envelope).Execute()

Return All Atlas Search Metric Types for One Process



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
    processId := "my.host.name.com:27017" // string | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.GetFTSMetricTypes(context.Background(), processId, groupId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetFTSMetricTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFTSMetricTypes`: ApiFTSMetricsView
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetFTSMetricTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **string** | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFTSMetricTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiFTSMetricsView**](ApiFTSMetricsView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFTSNamespaceIndexMeasurements

> ApiMeasurementsIndexesView GetFTSNamespaceIndexMeasurements(ctx, processId, databaseName, collectionName, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Envelope(envelope).Execute()

Return All Atlas Search Index Metrics for One Namespace



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
    processId := "my.host.name.com:27017" // string | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
    databaseName := "mydb" // string | Human-readable label that identifies the database.
    collectionName := "mycoll" // string | Human-readable label that identifies the collection.
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    granularity := "PT1M" // string | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.
    metrics := []string{"Metrics_example"} // []string | List that contains the measurements that MongoDB Atlas reports for the associated data series.
    period := "PT10H" // string | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. (optional)
    start := time.Now() // time.Time | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. (optional)
    end := time.Now() // time.Time | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. (optional)
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.GetFTSNamespaceIndexMeasurements(context.Background(), processId, databaseName, collectionName, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetFTSNamespaceIndexMeasurements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFTSNamespaceIndexMeasurements`: ApiMeasurementsIndexesView
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetFTSNamespaceIndexMeasurements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **string** | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests. | 
**databaseName** | **string** | Human-readable label that identifies the database. | 
**collectionName** | **string** | Human-readable label that identifies the collection. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFTSNamespaceIndexMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **granularity** | **string** | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. | 
 **metrics** | **[]string** | List that contains the measurements that MongoDB Atlas reports for the associated data series. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiMeasurementsIndexesView**](ApiMeasurementsIndexesView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFTSNonIndexMeasurements

> ApiMeasurementsNonIndexView GetFTSNonIndexMeasurements(ctx, processId, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Envelope(envelope).Execute()

Return Atlas Search Hardware and Status Metrics



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
    processId := "my.host.name.com:27017" // string | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    granularity := "PT1M" // string | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC.
    metrics := []string{"Metrics_example"} // []string | List that contains the metrics that you want MongoDB Atlas to report for the associated data series. If you don't set this parameter, this resource returns all hardware and status metrics for the associated data series.
    period := "PT10H" // string | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. (optional)
    start := time.Now() // time.Time | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. (optional)
    end := time.Now() // time.Time | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. (optional)
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.GetFTSNonIndexMeasurements(context.Background(), processId, groupId).Granularity(granularity).Metrics(metrics).Period(period).Start(start).End(end).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetFTSNonIndexMeasurements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFTSNonIndexMeasurements`: ApiMeasurementsNonIndexView
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetFTSNonIndexMeasurements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **string** | Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFTSNonIndexMeasurementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **granularity** | **string** | Duration that specifies the interval at which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. | 
 **metrics** | **[]string** | List that contains the metrics that you want MongoDB Atlas to report for the associated data series. If you don&#39;t set this parameter, this resource returns all hardware and status metrics for the associated data series. | 
 **period** | **string** | Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**. | 
 **start** | **time.Time** | Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **end** | **time.Time** | Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiMeasurementsNonIndexView**](ApiMeasurementsNonIndexView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostLogs

> *os.File GetHostLogs(ctx, groupId, hostName, logName).EndDate(endDate).StartDate(startDate).Execute()

Download Logs for One Multi-Cloud Cluster Host in One Project



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
    hostName := "hostName_example" // string | Fully qualified domain name or IP address of the MongoDB host that stores the log files that you want to download.
    logName := "logName_example" // string | Human-readable label of the log file that you want to return. You can return audit logs only if you enable Database Auditing for the specified project.
    endDate := int64(789) // int64 | Date and time when the period specifies the inclusive ending point for the range of log messages to retrieve. This parameter expresses its value in the number of seconds that have elapsed since the UNIX epoch. (optional)
    startDate := int64(789) // int64 | Date and time when the period specifies the inclusive starting point for the range of log messages to retrieve. This parameter expresses its value in the number of seconds that have elapsed since the UNIX epoch. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.GetHostLogs(context.Background(), groupId, hostName, logName).EndDate(endDate).StartDate(startDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.GetHostLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostLogs`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.GetHostLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**hostName** | **string** | Fully qualified domain name or IP address of the MongoDB host that stores the log files that you want to download. | 
**logName** | **string** | Human-readable label of the log file that you want to return. You can return audit logs only if you enable Database Auditing for the specified project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **endDate** | **int64** | Date and time when the period specifies the inclusive ending point for the range of log messages to retrieve. This parameter expresses its value in the number of seconds that have elapsed since the UNIX epoch. | 
 **startDate** | **int64** | Date and time when the period specifies the inclusive starting point for the range of log messages to retrieve. This parameter expresses its value in the number of seconds that have elapsed since the UNIX epoch. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllMongodbProcessesInOneProject

> PaginatedHostViewAtlas ReturnAllMongodbProcessesInOneProject(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All MongoDB Processes in One Project



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
    resp, r, err := apiClient.MonitoringAndLogsApi.ReturnAllMongodbProcessesInOneProject(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ReturnAllMongodbProcessesInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllMongodbProcessesInOneProject`: PaginatedHostViewAtlas
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ReturnAllMongodbProcessesInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllMongodbProcessesInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedHostViewAtlas**](PaginatedHostViewAtlas.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAvailableDatabasesForOneMongodbProcess

> PaginatedDatabaseView ReturnAvailableDatabasesForOneMongodbProcess(ctx, groupId, processId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return Available Databases for One MongoDB Process



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
    processId := "mongodb.example.com:27017" // string | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.ReturnAvailableDatabasesForOneMongodbProcess(context.Background(), groupId, processId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ReturnAvailableDatabasesForOneMongodbProcess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAvailableDatabasesForOneMongodbProcess`: PaginatedDatabaseView
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ReturnAvailableDatabasesForOneMongodbProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAvailableDatabasesForOneMongodbProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedDatabaseView**](PaginatedDatabaseView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAvailableDisksForOneMongodbProcess

> PaginatedDiskPartitionView ReturnAvailableDisksForOneMongodbProcess(ctx, groupId, processId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return Available Disks for One MongoDB Process



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
    processId := "mongodb.example.com:27017" // string | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.ReturnAvailableDisksForOneMongodbProcess(context.Background(), groupId, processId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ReturnAvailableDisksForOneMongodbProcess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAvailableDisksForOneMongodbProcess`: PaginatedDiskPartitionView
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ReturnAvailableDisksForOneMongodbProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAvailableDisksForOneMongodbProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedDiskPartitionView**](PaginatedDiskPartitionView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnDatabasesForOneMongodbProcess

> ApiDatabaseView ReturnDatabasesForOneMongodbProcess(ctx, groupId, databaseName, processId).Envelope(envelope).Pretty(pretty).Execute()

Return One Database for a MongoDB Process



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
    databaseName := "databaseName_example" // string | Human-readable label that identifies the database that the specified MongoDB process serves.
    processId := "mongodb.example.com:27017" // string | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.ReturnDatabasesForOneMongodbProcess(context.Background(), groupId, databaseName, processId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ReturnDatabasesForOneMongodbProcess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnDatabasesForOneMongodbProcess`: ApiDatabaseView
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ReturnDatabasesForOneMongodbProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**databaseName** | **string** | Human-readable label that identifies the database that the specified MongoDB process serves. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnDatabasesForOneMongodbProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiDatabaseView**](ApiDatabaseView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnMeasurementsForOneMongodbProcess

> ApiMeasurementsGeneralViewAtlas ReturnMeasurementsForOneMongodbProcess(ctx, groupId, processId).Envelope(envelope).Pretty(pretty).M(m).Period(period).Execute()

Return Measurements for One MongoDB Process



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
    processId := "mongodb.example.com:27017" // string | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    m := []string{"M_example"} // []string | One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for `m`, repeat the `m` parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements. (optional)
    period := time.Now() // time.Time | Date and time that indicates how far in the past to query. You can't set this value with **start** and **end** in the same request. This parameter expresses its value in the ISO 8601 duration format in UTC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.ReturnMeasurementsForOneMongodbProcess(context.Background(), groupId, processId).Envelope(envelope).Pretty(pretty).M(m).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ReturnMeasurementsForOneMongodbProcess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnMeasurementsForOneMongodbProcess`: ApiMeasurementsGeneralViewAtlas
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ReturnMeasurementsForOneMongodbProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnMeasurementsForOneMongodbProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **m** | **[]string** | One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements. | 
 **period** | **time.Time** | Date and time that indicates how far in the past to query. You can&#39;t set this value with **start** and **end** in the same request. This parameter expresses its value in the ISO 8601 duration format in UTC | 

### Return type

[**ApiMeasurementsGeneralViewAtlas**](ApiMeasurementsGeneralViewAtlas.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnMeasurementsOfOneDatabaseForOneMongodbProcess

> ApiMeasurementsGeneralViewAtlas ReturnMeasurementsOfOneDatabaseForOneMongodbProcess(ctx, groupId, databaseName, processId).Envelope(envelope).Pretty(pretty).M(m).Execute()

Return Measurements of One Database for One MongoDB Process



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
    databaseName := "databaseName_example" // string | Human-readable label that identifies the database that the specified MongoDB process serves.
    processId := "mongodb.example.com:27017" // string | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    m := []string{"M_example"} // []string | One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for `m`, repeat the `m` parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.ReturnMeasurementsOfOneDatabaseForOneMongodbProcess(context.Background(), groupId, databaseName, processId).Envelope(envelope).Pretty(pretty).M(m).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ReturnMeasurementsOfOneDatabaseForOneMongodbProcess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnMeasurementsOfOneDatabaseForOneMongodbProcess`: ApiMeasurementsGeneralViewAtlas
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ReturnMeasurementsOfOneDatabaseForOneMongodbProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**databaseName** | **string** | Human-readable label that identifies the database that the specified MongoDB process serves. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnMeasurementsOfOneDatabaseForOneMongodbProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **m** | **[]string** | One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements. | 

### Return type

[**ApiMeasurementsGeneralViewAtlas**](ApiMeasurementsGeneralViewAtlas.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnMeasurementsOfOneDisk

> ApiDiskPartitionView ReturnMeasurementsOfOneDisk(ctx, partitionName, groupId, processId).Envelope(envelope).Execute()

Return Measurements of One Disk



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
    partitionName := "partitionName_example" // string | Human-readable label of the disk or partition to which the measurements apply.
    groupId := "32b6e34b3d91647abb20e7b8" // string | Unique 24-hexadecimal digit string that identifies your project.
    processId := "mongodb.example.com:27017" // string | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.ReturnMeasurementsOfOneDisk(context.Background(), partitionName, groupId, processId).Envelope(envelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ReturnMeasurementsOfOneDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnMeasurementsOfOneDisk`: ApiDiskPartitionView
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ReturnMeasurementsOfOneDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Human-readable label of the disk or partition to which the measurements apply. | 
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnMeasurementsOfOneDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]

### Return type

[**ApiDiskPartitionView**](ApiDiskPartitionView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnMeasurementsOfOneDiskForOneMongodbProcess

> ApiMeasurementsGeneralViewAtlas ReturnMeasurementsOfOneDiskForOneMongodbProcess(ctx, groupId, partitionName, processId).Envelope(envelope).Pretty(pretty).M(m).Execute()

Return Measurements of One Disk for One MongoDB Process



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
    partitionName := "partitionName_example" // string | Human-readable label of the disk or partition to which the measurements apply.
    processId := "mongodb.example.com:27017" // string | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)
    m := []string{"M_example"} // []string | One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for `m`, repeat the `m` parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.ReturnMeasurementsOfOneDiskForOneMongodbProcess(context.Background(), groupId, partitionName, processId).Envelope(envelope).Pretty(pretty).M(m).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ReturnMeasurementsOfOneDiskForOneMongodbProcess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnMeasurementsOfOneDiskForOneMongodbProcess`: ApiMeasurementsGeneralViewAtlas
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ReturnMeasurementsOfOneDiskForOneMongodbProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**partitionName** | **string** | Human-readable label of the disk or partition to which the measurements apply. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnMeasurementsOfOneDiskForOneMongodbProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]
 **m** | **[]string** | One or more types of measurement to request for this MongoDB process. If omitted, the resource returns all measurements. To specify multiple values for &#x60;m&#x60;, repeat the &#x60;m&#x60; parameter for each value. Specify measurements that apply to the specified host. MongoDB Cloud returns an error if you specified any invalid measurements. | 

### Return type

[**ApiMeasurementsGeneralViewAtlas**](ApiMeasurementsGeneralViewAtlas.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneMongodbProcessById

> ApiHostViewAtlas ReturnOneMongodbProcessById(ctx, groupId, processId).Envelope(envelope).Pretty(pretty).Execute()

Return One MongoDB Process by ID



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
    processId := "mongodb.example.com:27017" // string | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndLogsApi.ReturnOneMongodbProcessById(context.Background(), groupId, processId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndLogsApi.ReturnOneMongodbProcessById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneMongodbProcessById`: ApiHostViewAtlas
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndLogsApi.ReturnOneMongodbProcessById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**processId** | **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneMongodbProcessByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiHostViewAtlas**](ApiHostViewAtlas.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

