# \EncryptionAtRestUsingCustomerKeyManagementApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject**](EncryptionAtRestUsingCustomerKeyManagementApi.md#ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/encryptionAtRest | Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project
[**UpdateConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject**](EncryptionAtRestUsingCustomerKeyManagementApi.md#UpdateConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject) | **Patch** /api/atlas/v1.0/groups/{groupId}/encryptionAtRest | Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project



## ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject

> ApiAtlasEncryptionAtRestView ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject(ctx, groupId).Envelope(envelope).Pretty(pretty).Execute()

Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project



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
    resp, r, err := apiClient.EncryptionAtRestUsingCustomerKeyManagementApi.ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject(context.Background(), groupId).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestUsingCustomerKeyManagementApi.ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject`: ApiAtlasEncryptionAtRestView
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestUsingCustomerKeyManagementApi.ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasEncryptionAtRestView**](ApiAtlasEncryptionAtRestView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject

> ApiAtlasEncryptionAtRestView UpdateConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject(ctx, groupId).ApiAtlasEncryptionAtRestView(apiAtlasEncryptionAtRestView).Envelope(envelope).Pretty(pretty).Execute()

Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project



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
    apiAtlasEncryptionAtRestView := *openapiclient.NewApiAtlasEncryptionAtRestView() // ApiAtlasEncryptionAtRestView | Required parameters depend on whether someone has enabled Encryption at Rest using Customer Key Management:  If you have enabled Encryption at Rest using Customer Key Management (CMK), Atlas requires all of the parameters for the desired encryption provider.  - To use AWS Key Management Service (KMS), MongoDB Cloud requires all the fields in the **awsKms** object. - To use Azure Key Vault, MongoDB Cloud requires all the fields in the **azureKeyVault** object. - To use Google Cloud Key Management Service (KMS), MongoDB Cloud requires all the fields in the **googleCloudKms** object.  If you enabled Encryption at Rest using Customer Key  Management, administrators can pass only the changed fields for the **awsKms**, **azureKeyVault**, or **googleCloudKms** object to update the configuration to this endpoint.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EncryptionAtRestUsingCustomerKeyManagementApi.UpdateConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject(context.Background(), groupId).ApiAtlasEncryptionAtRestView(apiAtlasEncryptionAtRestView).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestUsingCustomerKeyManagementApi.UpdateConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject`: ApiAtlasEncryptionAtRestView
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestUsingCustomerKeyManagementApi.UpdateConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasEncryptionAtRestView** | [**ApiAtlasEncryptionAtRestView**](ApiAtlasEncryptionAtRestView.md) | Required parameters depend on whether someone has enabled Encryption at Rest using Customer Key Management:  If you have enabled Encryption at Rest using Customer Key Management (CMK), Atlas requires all of the parameters for the desired encryption provider.  - To use AWS Key Management Service (KMS), MongoDB Cloud requires all the fields in the **awsKms** object. - To use Azure Key Vault, MongoDB Cloud requires all the fields in the **azureKeyVault** object. - To use Google Cloud Key Management Service (KMS), MongoDB Cloud requires all the fields in the **googleCloudKms** object.  If you enabled Encryption at Rest using Customer Key  Management, administrators can pass only the changed fields for the **awsKms**, **azureKeyVault**, or **googleCloudKms** object to update the configuration to this endpoint. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasEncryptionAtRestView**](ApiAtlasEncryptionAtRestView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

