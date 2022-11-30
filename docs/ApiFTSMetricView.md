# ApiFTSMetricView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | **string** | Human-readable label that identifies this Atlas Search hardware, status, or index measurement. | [readonly] 
**Units** | **string** | Unit of measurement that applies to this Atlas Search metric. | [readonly] 

## Methods

### NewApiFTSMetricView

`func NewApiFTSMetricView(metricName string, units string, ) *ApiFTSMetricView`

NewApiFTSMetricView instantiates a new ApiFTSMetricView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFTSMetricViewWithDefaults

`func NewApiFTSMetricViewWithDefaults() *ApiFTSMetricView`

NewApiFTSMetricViewWithDefaults instantiates a new ApiFTSMetricView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *ApiFTSMetricView) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ApiFTSMetricView) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ApiFTSMetricView) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetUnits

`func (o *ApiFTSMetricView) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ApiFTSMetricView) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ApiFTSMetricView) SetUnits(v string)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


