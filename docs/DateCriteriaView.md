# DateCriteriaView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateField** | Pointer to **string** | Indexed database parameter that stores the date that determines when data moves to the online archive. MongoDB Cloud archives the data when the current date exceeds the date in this database parameter plus the number of days specified through the **expireAfterDays** parameter. Set this parameter when you set &#x60;\&quot;criteria.type\&quot; : \&quot;DATE\&quot;&#x60;. | [optional] 
**DateFormat** | Pointer to **string** | Syntax used to write the date after which data moves to the online archive. Date can be expressed as ISO 8601 or Epoch timestamps. The Epoch timestamp can be expressed as nanoseconds, milliseconds, or seconds. Set this parameter when **\&quot;criteria.type\&quot; : \&quot;DATE\&quot;**. You must set **\&quot;criteria.type\&quot; : \&quot;DATE\&quot;** if **\&quot;collectionType\&quot;: \&quot;TIMESERIES\&quot;**. | [optional] [default to "ISODATE"]
**ExpireAfterDays** | Pointer to **int32** | Number of days after the value in the **criteria.dateField** when MongoDB Cloud archives data in the specified cluster. Set this parameter when you set **\&quot;criteria.type\&quot; : \&quot;DATE\&quot;**. | [optional] 

## Methods

### NewDateCriteriaView

`func NewDateCriteriaView() *DateCriteriaView`

NewDateCriteriaView instantiates a new DateCriteriaView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateCriteriaViewWithDefaults

`func NewDateCriteriaViewWithDefaults() *DateCriteriaView`

NewDateCriteriaViewWithDefaults instantiates a new DateCriteriaView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateField

`func (o *DateCriteriaView) GetDateField() string`

GetDateField returns the DateField field if non-nil, zero value otherwise.

### GetDateFieldOk

`func (o *DateCriteriaView) GetDateFieldOk() (*string, bool)`

GetDateFieldOk returns a tuple with the DateField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateField

`func (o *DateCriteriaView) SetDateField(v string)`

SetDateField sets DateField field to given value.

### HasDateField

`func (o *DateCriteriaView) HasDateField() bool`

HasDateField returns a boolean if a field has been set.

### GetDateFormat

`func (o *DateCriteriaView) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *DateCriteriaView) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *DateCriteriaView) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *DateCriteriaView) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetExpireAfterDays

`func (o *DateCriteriaView) GetExpireAfterDays() int32`

GetExpireAfterDays returns the ExpireAfterDays field if non-nil, zero value otherwise.

### GetExpireAfterDaysOk

`func (o *DateCriteriaView) GetExpireAfterDaysOk() (*int32, bool)`

GetExpireAfterDaysOk returns a tuple with the ExpireAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfterDays

`func (o *DateCriteriaView) SetExpireAfterDays(v int32)`

SetExpireAfterDays sets ExpireAfterDays field to given value.

### HasExpireAfterDays

`func (o *DateCriteriaView) HasExpireAfterDays() bool`

HasExpireAfterDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


