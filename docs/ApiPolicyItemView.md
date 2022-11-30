# ApiPolicyItemView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrequencyInterval** | **int32** | Number that indicates the frequency interval for a set of snapshots. A value of &#x60;1&#x60; specifies the first instance of the corresponding &#x60;frequencyType&#x60;.  - In a monthly policy item, &#x60;1&#x60; indicates that the monthly snapshot occurs on the first day of the month and &#x60;40&#x60; indicates the last day of the month.  - In a weekly policy item, &#x60;1&#x60; indicates that the weekly snapshot occurs on Monday and &#x60;7&#x60; indicates Sunday.  - In an hourly policy item, you can set the frequency interval to &#x60;1&#x60;, &#x60;2&#x60;, &#x60;4&#x60;, &#x60;6&#x60;, &#x60;8&#x60;, or &#x60;12&#x60;. For hourly policy items for NVMe clusters, MongoDB Cloud only accepts &#x60;12&#x60; as the frequency interval value. | 
**FrequencyType** | **string** | Human-readable label that identifies the frequency type associated with the backup policy. | 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this backup policy item. | [optional] 
**RetentionUnit** | **string** | Unit of time in which MongoDB Cloud measures snapshot retention. | 
**RetentionValue** | **int32** | Duration in days, weeks, or months that MongoDB Cloud retains the snapshot. For less frequent policy items, MongoDB Cloud requires that you specify a value greater than or equal to the value specified for more frequent policy items.  For example: If the hourly policy item specifies a retention of two days, you must specify two days or greater for the retention of the weekly policy item. | 

## Methods

### NewApiPolicyItemView

`func NewApiPolicyItemView(frequencyInterval int32, frequencyType string, retentionUnit string, retentionValue int32, ) *ApiPolicyItemView`

NewApiPolicyItemView instantiates a new ApiPolicyItemView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyItemViewWithDefaults

`func NewApiPolicyItemViewWithDefaults() *ApiPolicyItemView`

NewApiPolicyItemViewWithDefaults instantiates a new ApiPolicyItemView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequencyInterval

`func (o *ApiPolicyItemView) GetFrequencyInterval() int32`

GetFrequencyInterval returns the FrequencyInterval field if non-nil, zero value otherwise.

### GetFrequencyIntervalOk

`func (o *ApiPolicyItemView) GetFrequencyIntervalOk() (*int32, bool)`

GetFrequencyIntervalOk returns a tuple with the FrequencyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyInterval

`func (o *ApiPolicyItemView) SetFrequencyInterval(v int32)`

SetFrequencyInterval sets FrequencyInterval field to given value.


### GetFrequencyType

`func (o *ApiPolicyItemView) GetFrequencyType() string`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *ApiPolicyItemView) GetFrequencyTypeOk() (*string, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *ApiPolicyItemView) SetFrequencyType(v string)`

SetFrequencyType sets FrequencyType field to given value.


### GetId

`func (o *ApiPolicyItemView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiPolicyItemView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiPolicyItemView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiPolicyItemView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRetentionUnit

`func (o *ApiPolicyItemView) GetRetentionUnit() string`

GetRetentionUnit returns the RetentionUnit field if non-nil, zero value otherwise.

### GetRetentionUnitOk

`func (o *ApiPolicyItemView) GetRetentionUnitOk() (*string, bool)`

GetRetentionUnitOk returns a tuple with the RetentionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionUnit

`func (o *ApiPolicyItemView) SetRetentionUnit(v string)`

SetRetentionUnit sets RetentionUnit field to given value.


### GetRetentionValue

`func (o *ApiPolicyItemView) GetRetentionValue() int32`

GetRetentionValue returns the RetentionValue field if non-nil, zero value otherwise.

### GetRetentionValueOk

`func (o *ApiPolicyItemView) GetRetentionValueOk() (*int32, bool)`

GetRetentionValueOk returns a tuple with the RetentionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionValue

`func (o *ApiPolicyItemView) SetRetentionValue(v int32)`

SetRetentionValue sets RetentionValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

