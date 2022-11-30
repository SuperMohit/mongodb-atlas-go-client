# ApiDatabaseView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | Pointer to **string** | Human-readable label that identifies the database that the specified MongoDB process serves. | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 

## Methods

### NewApiDatabaseView

`func NewApiDatabaseView(links []Link, ) *ApiDatabaseView`

NewApiDatabaseView instantiates a new ApiDatabaseView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDatabaseViewWithDefaults

`func NewApiDatabaseViewWithDefaults() *ApiDatabaseView`

NewApiDatabaseViewWithDefaults instantiates a new ApiDatabaseView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *ApiDatabaseView) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ApiDatabaseView) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ApiDatabaseView) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *ApiDatabaseView) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetLinks

`func (o *ApiDatabaseView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiDatabaseView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiDatabaseView) SetLinks(v []Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


