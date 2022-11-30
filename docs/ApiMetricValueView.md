# ApiMetricValueView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **float32** | Amount of the **metricName** recorded at the time of the event. This value triggered the alert. | [optional] [readonly] 
**Units** | Pointer to **string** | Element used to express the quantity in **currentValue.number**. This can be an element of time, storage capacity, and the like. This metric triggered the alert. | [optional] [readonly] 

## Methods

### NewApiMetricValueView

`func NewApiMetricValueView() *ApiMetricValueView`

NewApiMetricValueView instantiates a new ApiMetricValueView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMetricValueViewWithDefaults

`func NewApiMetricValueViewWithDefaults() *ApiMetricValueView`

NewApiMetricValueViewWithDefaults instantiates a new ApiMetricValueView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ApiMetricValueView) GetNumber() float32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ApiMetricValueView) GetNumberOk() (*float32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ApiMetricValueView) SetNumber(v float32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ApiMetricValueView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetUnits

`func (o *ApiMetricValueView) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ApiMetricValueView) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ApiMetricValueView) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ApiMetricValueView) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


