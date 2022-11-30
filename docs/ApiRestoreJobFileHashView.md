# ApiRestoreJobFileHashView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | Human-readable label that identifies the hashed file. | [optional] [readonly] 
**Hash** | Pointer to **string** | Hashed checksum that maps to the restore file. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**TypeName** | Pointer to **string** | Human-readable label that identifies the hashing algorithm used to compute the hash value. | [optional] [readonly] 

## Methods

### NewApiRestoreJobFileHashView

`func NewApiRestoreJobFileHashView(links []Link, ) *ApiRestoreJobFileHashView`

NewApiRestoreJobFileHashView instantiates a new ApiRestoreJobFileHashView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRestoreJobFileHashViewWithDefaults

`func NewApiRestoreJobFileHashViewWithDefaults() *ApiRestoreJobFileHashView`

NewApiRestoreJobFileHashViewWithDefaults instantiates a new ApiRestoreJobFileHashView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *ApiRestoreJobFileHashView) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ApiRestoreJobFileHashView) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ApiRestoreJobFileHashView) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ApiRestoreJobFileHashView) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetHash

`func (o *ApiRestoreJobFileHashView) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiRestoreJobFileHashView) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiRestoreJobFileHashView) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiRestoreJobFileHashView) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetLinks

`func (o *ApiRestoreJobFileHashView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiRestoreJobFileHashView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiRestoreJobFileHashView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetTypeName

`func (o *ApiRestoreJobFileHashView) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *ApiRestoreJobFileHashView) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *ApiRestoreJobFileHashView) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *ApiRestoreJobFileHashView) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


