# ApiAtlasSnapshotRetentionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**RetentionUnit** | **string** | Quantity of time in which MongoDB Cloud measures snapshot retention. | 
**RetentionValue** | **int32** | Number that indicates the amount of days, weeks, or months that MongoDB Cloud retains the snapshot. For less frequent policy items, MongoDB Cloud requires that you specify a value greater than or equal to the value specified for more frequent policy items. If the hourly policy item specifies a retention of two days, specify two days or greater for the retention of the weekly policy item. | 

## Methods

### NewApiAtlasSnapshotRetentionView

`func NewApiAtlasSnapshotRetentionView(links []Link, retentionUnit string, retentionValue int32, ) *ApiAtlasSnapshotRetentionView`

NewApiAtlasSnapshotRetentionView instantiates a new ApiAtlasSnapshotRetentionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasSnapshotRetentionViewWithDefaults

`func NewApiAtlasSnapshotRetentionViewWithDefaults() *ApiAtlasSnapshotRetentionView`

NewApiAtlasSnapshotRetentionViewWithDefaults instantiates a new ApiAtlasSnapshotRetentionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApiAtlasSnapshotRetentionView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasSnapshotRetentionView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasSnapshotRetentionView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetRetentionUnit

`func (o *ApiAtlasSnapshotRetentionView) GetRetentionUnit() string`

GetRetentionUnit returns the RetentionUnit field if non-nil, zero value otherwise.

### GetRetentionUnitOk

`func (o *ApiAtlasSnapshotRetentionView) GetRetentionUnitOk() (*string, bool)`

GetRetentionUnitOk returns a tuple with the RetentionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionUnit

`func (o *ApiAtlasSnapshotRetentionView) SetRetentionUnit(v string)`

SetRetentionUnit sets RetentionUnit field to given value.


### GetRetentionValue

`func (o *ApiAtlasSnapshotRetentionView) GetRetentionValue() int32`

GetRetentionValue returns the RetentionValue field if non-nil, zero value otherwise.

### GetRetentionValueOk

`func (o *ApiAtlasSnapshotRetentionView) GetRetentionValueOk() (*int32, bool)`

GetRetentionValueOk returns a tuple with the RetentionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionValue

`func (o *ApiAtlasSnapshotRetentionView) SetRetentionValue(v int32)`

SetRetentionValue sets RetentionValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


