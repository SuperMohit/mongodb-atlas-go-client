# \InvoicesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnAllInvoicesForOneOrganization**](InvoicesApi.md#ReturnAllInvoicesForOneOrganization) | **Get** /api/atlas/v1.0/orgs/{orgId}/invoices | Return All Invoices for One Organization
[**ReturnAllPendingInvoicesForOneOrganization**](InvoicesApi.md#ReturnAllPendingInvoicesForOneOrganization) | **Get** /api/atlas/v1.0/orgs/{orgId}/invoices/pending | Return All Pending Invoices for One Organization
[**ReturnOneOrganizationInvoice**](InvoicesApi.md#ReturnOneOrganizationInvoice) | **Get** /api/atlas/v1.0/orgs/{orgId}/invoices/{invoiceId} | Return One Organization Invoice
[**ReturnOneOrganizationInvoiceAsCsv**](InvoicesApi.md#ReturnOneOrganizationInvoiceAsCsv) | **Get** /api/atlas/v1.0/orgs/{orgId}/invoices/{invoiceId}/csv | Return One Organization Invoice as CSV



## ReturnAllInvoicesForOneOrganization

> PaginatedApiInvoiceView ReturnAllInvoicesForOneOrganization(ctx, orgId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Invoices for One Organization



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
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    includeCount := true // bool | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. (optional) (default to true)
    itemsPerPage := int32(100) // int32 | Number of items that the response returns per page. (optional) (default to 100)
    pageNum := int32(1) // int32 | Number of the page that displays the current set of the total objects that the response returns. (optional) (default to 1)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesApi.ReturnAllInvoicesForOneOrganization(context.Background(), orgId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.ReturnAllInvoicesForOneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllInvoicesForOneOrganization`: PaginatedApiInvoiceView
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.ReturnAllInvoicesForOneOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllInvoicesForOneOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedApiInvoiceView**](PaginatedApiInvoiceView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnAllPendingInvoicesForOneOrganization

> PaginatedApiInvoiceView ReturnAllPendingInvoicesForOneOrganization(ctx, orgId).Envelope(envelope).Pretty(pretty).Execute()

Return All Pending Invoices for One Organization



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
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesApi.ReturnAllPendingInvoicesForOneOrganization(context.Background(), orgId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.ReturnAllPendingInvoicesForOneOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllPendingInvoicesForOneOrganization`: PaginatedApiInvoiceView
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.ReturnAllPendingInvoicesForOneOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllPendingInvoicesForOneOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedApiInvoiceView**](PaginatedApiInvoiceView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneOrganizationInvoice

> ApiInvoiceView ReturnOneOrganizationInvoice(ctx, orgId, invoiceId).Accept(accept).Envelope(envelope).Pretty(pretty).Execute()

Return One Organization Invoice



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
    invoiceId := "invoiceId_example" // string | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
    accept := "accept_example" // string | Indicates which content types the client is able to understand. Expressed as MIME types. (optional) (default to "application/json")
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesApi.ReturnOneOrganizationInvoice(context.Background(), orgId, invoiceId).Accept(accept).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.ReturnOneOrganizationInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneOrganizationInvoice`: ApiInvoiceView
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.ReturnOneOrganizationInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**invoiceId** | **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneOrganizationInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** | Indicates which content types the client is able to understand. Expressed as MIME types. | [default to &quot;application/json&quot;]
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiInvoiceView**](ApiInvoiceView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneOrganizationInvoiceAsCsv

> ReturnOneOrganizationInvoiceAsCsv(ctx, orgId, invoiceId).Envelope(envelope).Pretty(pretty).Execute()

Return One Organization Invoice as CSV



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
    invoiceId := "invoiceId_example" // string | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesApi.ReturnOneOrganizationInvoiceAsCsv(context.Background(), orgId, invoiceId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.ReturnOneOrganizationInvoiceAsCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. | 
**invoiceId** | **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneOrganizationInvoiceAsCsvRequest struct via the builder pattern


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
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

