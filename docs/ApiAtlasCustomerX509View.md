# ApiAtlasCustomerX509View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cas** | Pointer to **string** | Concatenated list of customer certificate authority (CA) certificates needed to authenticate database users. MongoDB Cloud expects this as a PEM-formatted certificate. | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 

## Methods

### NewApiAtlasCustomerX509View

`func NewApiAtlasCustomerX509View(links []Link, ) *ApiAtlasCustomerX509View`

NewApiAtlasCustomerX509View instantiates a new ApiAtlasCustomerX509View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCustomerX509ViewWithDefaults

`func NewApiAtlasCustomerX509ViewWithDefaults() *ApiAtlasCustomerX509View`

NewApiAtlasCustomerX509ViewWithDefaults instantiates a new ApiAtlasCustomerX509View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCas

`func (o *ApiAtlasCustomerX509View) GetCas() string`

GetCas returns the Cas field if non-nil, zero value otherwise.

### GetCasOk

`func (o *ApiAtlasCustomerX509View) GetCasOk() (*string, bool)`

GetCasOk returns a tuple with the Cas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCas

`func (o *ApiAtlasCustomerX509View) SetCas(v string)`

SetCas sets Cas field to given value.

### HasCas

`func (o *ApiAtlasCustomerX509View) HasCas() bool`

HasCas returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasCustomerX509View) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasCustomerX509View) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasCustomerX509View) SetLinks(v []Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


