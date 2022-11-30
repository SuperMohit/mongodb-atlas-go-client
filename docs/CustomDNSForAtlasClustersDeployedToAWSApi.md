# \CustomDNSForAtlasClustersDeployedToAWSApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnOneCustomDnsConfigurationForAtlasClustersOnAws**](CustomDNSForAtlasClustersDeployedToAWSApi.md#ReturnOneCustomDnsConfigurationForAtlasClustersOnAws) | **Get** /api/atlas/v1.0/groups/{groupId}/awsCustomDNS | Return One Custom DNS Configuration for Atlas Clusters on AWS
[**ToggleOneStateOfOneCustomDnsConfigurationForAtlasClustersOnAws**](CustomDNSForAtlasClustersDeployedToAWSApi.md#ToggleOneStateOfOneCustomDnsConfigurationForAtlasClustersOnAws) | **Patch** /api/atlas/v1.0/groups/{groupId}/awsCustomDNS | Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS



## ReturnOneCustomDnsConfigurationForAtlasClustersOnAws

> AWSCustomDNSEnabledView ReturnOneCustomDnsConfigurationForAtlasClustersOnAws(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Return One Custom DNS Configuration for Atlas Clusters on AWS



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
    resp, r, err := apiClient.CustomDNSForAtlasClustersDeployedToAWSApi.ReturnOneCustomDnsConfigurationForAtlasClustersOnAws(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDNSForAtlasClustersDeployedToAWSApi.ReturnOneCustomDnsConfigurationForAtlasClustersOnAws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneCustomDnsConfigurationForAtlasClustersOnAws`: AWSCustomDNSEnabledView
    fmt.Fprintf(os.Stdout, "Response from `CustomDNSForAtlasClustersDeployedToAWSApi.ReturnOneCustomDnsConfigurationForAtlasClustersOnAws`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneCustomDnsConfigurationForAtlasClustersOnAwsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**AWSCustomDNSEnabledView**](AWSCustomDNSEnabledView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleOneStateOfOneCustomDnsConfigurationForAtlasClustersOnAws

> AWSCustomDNSEnabledView ToggleOneStateOfOneCustomDnsConfigurationForAtlasClustersOnAws(ctx, groupId).AWSCustomDNSEnabledView(aWSCustomDNSEnabledView).Envelope(envelope).Pretty(pretty).Execute()

Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS



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
    aWSCustomDNSEnabledView := *openapiclient.NewAWSCustomDNSEnabledView(false) // AWSCustomDNSEnabledView | Enables or disables the custom DNS configuration for AWS clusters in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDNSForAtlasClustersDeployedToAWSApi.ToggleOneStateOfOneCustomDnsConfigurationForAtlasClustersOnAws(context.Background(), groupId).AWSCustomDNSEnabledView(aWSCustomDNSEnabledView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDNSForAtlasClustersDeployedToAWSApi.ToggleOneStateOfOneCustomDnsConfigurationForAtlasClustersOnAws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleOneStateOfOneCustomDnsConfigurationForAtlasClustersOnAws`: AWSCustomDNSEnabledView
    fmt.Fprintf(os.Stdout, "Response from `CustomDNSForAtlasClustersDeployedToAWSApi.ToggleOneStateOfOneCustomDnsConfigurationForAtlasClustersOnAws`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleOneStateOfOneCustomDnsConfigurationForAtlasClustersOnAwsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aWSCustomDNSEnabledView** | [**AWSCustomDNSEnabledView**](AWSCustomDNSEnabledView.md) | Enables or disables the custom DNS configuration for AWS clusters in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**AWSCustomDNSEnabledView**](AWSCustomDNSEnabledView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

