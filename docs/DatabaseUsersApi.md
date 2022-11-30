# \DatabaseUsersApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOneDatabaseUserInOneProject**](DatabaseUsersApi.md#CreateOneDatabaseUserInOneProject) | **Post** /api/atlas/v1.0/groups/{groupId}/databaseUsers | Create One Database User in One Project
[**RemoveOneDatabaseUserFromOneProject**](DatabaseUsersApi.md#RemoveOneDatabaseUserFromOneProject) | **Delete** /api/atlas/v1.0/groups/{groupId}/databaseUsers/{databaseName}/{username} | Remove One Database User from One Project
[**ReturnAllDatabaseUsersFromOneProject**](DatabaseUsersApi.md#ReturnAllDatabaseUsersFromOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/databaseUsers | Return All Database Users from One Project
[**ReturnOneDatabaseUserFromOneProject**](DatabaseUsersApi.md#ReturnOneDatabaseUserFromOneProject) | **Get** /api/atlas/v1.0/groups/{groupId}/databaseUsers/{databaseName}/{username} | Return One Database User from One Project
[**UpdateOneDatabaseUserInOneProject**](DatabaseUsersApi.md#UpdateOneDatabaseUserInOneProject) | **Patch** /api/atlas/v1.0/groups/{groupId}/databaseUsers/{databaseName}/{username} | Update One Database User in One Project



## CreateOneDatabaseUserInOneProject

> ApiAtlasDatabaseUserViewManual CreateOneDatabaseUserInOneProject(ctx, groupId).ApiAtlasDatabaseUserViewManual(apiAtlasDatabaseUserViewManual).Envelope(envelope).Pretty(pretty).Execute()

Create One Database User in One Project



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
    apiAtlasDatabaseUserViewManual := openapiclient.ApiAtlasDatabaseUserViewManual{ApiDatabaseUserIAMView: openapiclient.NewApiDatabaseUserIAMView("DatabaseName_example", "32b6e34b3d91647abb20e7b8", []openapiclient.ApiAtlasRoleView{*openapiclient.NewApiAtlasRoleView("CollectionName_example", "DatabaseName_example", "RoleName_example")}, "AwsIAMType_example", "arn:aws:iam::358363220050:user/mongodb-aws-iam-auth-pat-bloggs")} // ApiAtlasDatabaseUserViewManual | Creates one database user in the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseUsersApi.CreateOneDatabaseUserInOneProject(context.Background(), groupId).ApiAtlasDatabaseUserViewManual(apiAtlasDatabaseUserViewManual).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseUsersApi.CreateOneDatabaseUserInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOneDatabaseUserInOneProject`: ApiAtlasDatabaseUserViewManual
    fmt.Fprintf(os.Stdout, "Response from `DatabaseUsersApi.CreateOneDatabaseUserInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOneDatabaseUserInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiAtlasDatabaseUserViewManual** | [**ApiAtlasDatabaseUserViewManual**](ApiAtlasDatabaseUserViewManual.md) | Creates one database user in the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDatabaseUserViewManual**](ApiAtlasDatabaseUserViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOneDatabaseUserFromOneProject

> RemoveOneDatabaseUserFromOneProject(ctx, groupId, databaseName, username).Envelope(envelope).Pretty(pretty).Execute()

Remove One Database User from One Project



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
    databaseName := "databaseName_example" // string | Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be `$external`. If the user authenticates with SCRAM-SHA, this value should be `admin`.
    username := "SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs orx.509 or LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM" // string | Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsType | ROLE | <abbr title=\"Amazon Resource Name\">ARN</abbr> | | AWS IAM | awsType | USER | <abbr title=\"Amazon Resource Name\">ARN</abbr> | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string | 
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseUsersApi.RemoveOneDatabaseUserFromOneProject(context.Background(), groupId, databaseName, username).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseUsersApi.RemoveOneDatabaseUserFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**databaseName** | **string** | Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA, this value should be &#x60;admin&#x60;. | 
**username** | **string** | Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | AWS IAM | awsType | USER | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOneDatabaseUserFromOneProjectRequest struct via the builder pattern


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


## ReturnAllDatabaseUsersFromOneProject

> PaginatedDatabaseUserView ReturnAllDatabaseUsersFromOneProject(ctx, groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()

Return All Database Users from One Project



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
    resp, r, err := apiClient.DatabaseUsersApi.ReturnAllDatabaseUsersFromOneProject(context.Background(), groupId).Envelope(envelope).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseUsersApi.ReturnAllDatabaseUsersFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnAllDatabaseUsersFromOneProject`: PaginatedDatabaseUserView
    fmt.Fprintf(os.Stdout, "Response from `DatabaseUsersApi.ReturnAllDatabaseUsersFromOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnAllDatabaseUsersFromOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int32** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int32** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**PaginatedDatabaseUserView**](PaginatedDatabaseUserView.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnOneDatabaseUserFromOneProject

> ApiAtlasDatabaseUserViewManual ReturnOneDatabaseUserFromOneProject(ctx, groupId, databaseName, username).Envelope(envelope).Pretty(pretty).Execute()

Return One Database User from One Project



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
    databaseName := "databaseName_example" // string | Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be `$external`. If the user authenticates with SCRAM-SHA, this value should be `admin`.
    username := "SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs orx.509 or LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM" // string | Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsType | ROLE | <abbr title=\"Amazon Resource Name\">ARN</abbr> | | AWS IAM | awsType | USER | <abbr title=\"Amazon Resource Name\">ARN</abbr> | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string | 
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseUsersApi.ReturnOneDatabaseUserFromOneProject(context.Background(), groupId, databaseName, username).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseUsersApi.ReturnOneDatabaseUserFromOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnOneDatabaseUserFromOneProject`: ApiAtlasDatabaseUserViewManual
    fmt.Fprintf(os.Stdout, "Response from `DatabaseUsersApi.ReturnOneDatabaseUserFromOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**databaseName** | **string** | Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA, this value should be &#x60;admin&#x60;. | 
**username** | **string** | Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | AWS IAM | awsType | USER | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnOneDatabaseUserFromOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDatabaseUserViewManual**](ApiAtlasDatabaseUserViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneDatabaseUserInOneProject

> ApiAtlasDatabaseUserViewManual UpdateOneDatabaseUserInOneProject(ctx, groupId, databaseName, username).ApiAtlasDatabaseUserViewManual(apiAtlasDatabaseUserViewManual).Envelope(envelope).Pretty(pretty).Execute()

Update One Database User in One Project



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
    databaseName := "databaseName_example" // string | Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be `$external`. If the user authenticates with SCRAM-SHA, this value should be `admin`.
    username := "SCRAM-SHA: dylan or AWS IAM: arn:aws:iam::123456789012:user/sales/enterprise/DylanBloggs orx.509 or LDAP: CN=Dylan Bloggs,OU=Enterprise,OU=Sales,DC=Example,DC=COM" // string | Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsType | ROLE | <abbr title=\"Amazon Resource Name\">ARN</abbr> | | AWS IAM | awsType | USER | <abbr title=\"Amazon Resource Name\">ARN</abbr> | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string | 
    apiAtlasDatabaseUserViewManual := openapiclient.ApiAtlasDatabaseUserViewManual{ApiDatabaseUserIAMView: openapiclient.NewApiDatabaseUserIAMView("DatabaseName_example", "32b6e34b3d91647abb20e7b8", []openapiclient.ApiAtlasRoleView{*openapiclient.NewApiAtlasRoleView("CollectionName_example", "DatabaseName_example", "RoleName_example")}, "AwsIAMType_example", "arn:aws:iam::358363220050:user/mongodb-aws-iam-auth-pat-bloggs")} // ApiAtlasDatabaseUserViewManual | Updates one database user that belongs to the specified project.
    envelope := false // bool | Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. (optional) (default to false)
    pretty := false // bool | Flag that indicates whether the response body should be in the prettyprint format. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseUsersApi.UpdateOneDatabaseUserInOneProject(context.Background(), groupId, databaseName, username).ApiAtlasDatabaseUserViewManual(apiAtlasDatabaseUserViewManual).Envelope(envelope).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseUsersApi.UpdateOneDatabaseUserInOneProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneDatabaseUserInOneProject`: ApiAtlasDatabaseUserViewManual
    fmt.Fprintf(os.Stdout, "Response from `DatabaseUsersApi.UpdateOneDatabaseUserInOneProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Unique 24-hexadecimal digit string that identifies your project. | 
**databaseName** | **string** | Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA, this value should be &#x60;admin&#x60;. | 
**username** | **string** | Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | AWS IAM | awsType | USER | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneDatabaseUserInOneProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiAtlasDatabaseUserViewManual** | [**ApiAtlasDatabaseUserViewManual**](ApiAtlasDatabaseUserViewManual.md) | Updates one database user that belongs to the specified project. | 
 **envelope** | **bool** | Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body. | [default to false]
 **pretty** | **bool** | Flag that indicates whether the response body should be in the prettyprint format. | [default to false]

### Return type

[**ApiAtlasDatabaseUserViewManual**](ApiAtlasDatabaseUserViewManual.md)

### Authorization

[DigestAuth](../README.md#DigestAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

