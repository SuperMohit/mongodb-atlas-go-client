# ApiMetricThresholdView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | Pointer to **string** | Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**. | [optional] 
**Mode** | Pointer to **string** | MongoDB Cloud computes the current metric value as an average. | [optional] [default to "AVERAGE"]
**Operator** | Pointer to **string** | Comparison operator to apply when checking the current metric value. | [optional] 
**Threshold** | Pointer to **float64** | Value of metric that, when exceeded, triggers an alert. | [optional] 
**Units** | Pointer to **string** | Element used to express the quantity. This can be an element of time, storage capacity, and the like. | [optional] 

## Methods

### NewApiMetricThresholdView

`func NewApiMetricThresholdView() *ApiMetricThresholdView`

NewApiMetricThresholdView instantiates a new ApiMetricThresholdView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMetricThresholdViewWithDefaults

`func NewApiMetricThresholdViewWithDefaults() *ApiMetricThresholdView`

NewApiMetricThresholdViewWithDefaults instantiates a new ApiMetricThresholdView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *ApiMetricThresholdView) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ApiMetricThresholdView) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ApiMetricThresholdView) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *ApiMetricThresholdView) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMode

`func (o *ApiMetricThresholdView) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApiMetricThresholdView) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApiMetricThresholdView) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ApiMetricThresholdView) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOperator

`func (o *ApiMetricThresholdView) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ApiMetricThresholdView) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ApiMetricThresholdView) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ApiMetricThresholdView) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetThreshold

`func (o *ApiMetricThresholdView) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ApiMetricThresholdView) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ApiMetricThresholdView) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ApiMetricThresholdView) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUnits

`func (o *ApiMetricThresholdView) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ApiMetricThresholdView) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ApiMetricThresholdView) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ApiMetricThresholdView) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


