# ApiAtlasDataLakeHTTPStoreView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInsecure** | Pointer to **bool** | Flag that validates the scheme in the specified URLs. If &#x60;true&#x60;, allows insecure &#x60;HTTP&#x60; scheme, doesn&#39;t verify the server&#39;s certificate chain and hostname, and accepts any certificate with any hostname presented by the server. If &#x60;false&#x60;, allows secure &#x60;HTTPS&#x60; scheme only. | [optional] [default to false]
**DefaultFormat** | Pointer to **string** | Default format that Data Lake assumes if it encounters a file without an extension while searching the &#x60;storeName&#x60;. If omitted, Data Lake attempts to detect the file type by processing a few bytes of the file. The specified format only applies to the URLs specified in the **databases.[n].collections.[n].dataSources** object. | [optional] 
**Urls** | Pointer to **[]string** | Comma-separated list of publicly accessible HTTP URLs where data is stored. You can&#39;t specify URLs that require authentication. | [optional] 

## Methods

### NewApiAtlasDataLakeHTTPStoreView

`func NewApiAtlasDataLakeHTTPStoreView() *ApiAtlasDataLakeHTTPStoreView`

NewApiAtlasDataLakeHTTPStoreView instantiates a new ApiAtlasDataLakeHTTPStoreView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDataLakeHTTPStoreViewWithDefaults

`func NewApiAtlasDataLakeHTTPStoreViewWithDefaults() *ApiAtlasDataLakeHTTPStoreView`

NewApiAtlasDataLakeHTTPStoreViewWithDefaults instantiates a new ApiAtlasDataLakeHTTPStoreView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInsecure

`func (o *ApiAtlasDataLakeHTTPStoreView) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *ApiAtlasDataLakeHTTPStoreView) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *ApiAtlasDataLakeHTTPStoreView) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *ApiAtlasDataLakeHTTPStoreView) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetDefaultFormat

`func (o *ApiAtlasDataLakeHTTPStoreView) GetDefaultFormat() string`

GetDefaultFormat returns the DefaultFormat field if non-nil, zero value otherwise.

### GetDefaultFormatOk

`func (o *ApiAtlasDataLakeHTTPStoreView) GetDefaultFormatOk() (*string, bool)`

GetDefaultFormatOk returns a tuple with the DefaultFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFormat

`func (o *ApiAtlasDataLakeHTTPStoreView) SetDefaultFormat(v string)`

SetDefaultFormat sets DefaultFormat field to given value.

### HasDefaultFormat

`func (o *ApiAtlasDataLakeHTTPStoreView) HasDefaultFormat() bool`

HasDefaultFormat returns a boolean if a field has been set.

### GetUrls

`func (o *ApiAtlasDataLakeHTTPStoreView) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *ApiAtlasDataLakeHTTPStoreView) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *ApiAtlasDataLakeHTTPStoreView) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *ApiAtlasDataLakeHTTPStoreView) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


