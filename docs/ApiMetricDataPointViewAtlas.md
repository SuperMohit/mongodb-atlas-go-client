# ApiMetricDataPointViewAtlas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | Date and time when this data point occurred. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Value** | Pointer to **float32** | Value that comprises this data point. | [optional] [readonly] 

## Methods

### NewApiMetricDataPointViewAtlas

`func NewApiMetricDataPointViewAtlas() *ApiMetricDataPointViewAtlas`

NewApiMetricDataPointViewAtlas instantiates a new ApiMetricDataPointViewAtlas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMetricDataPointViewAtlasWithDefaults

`func NewApiMetricDataPointViewAtlasWithDefaults() *ApiMetricDataPointViewAtlas`

NewApiMetricDataPointViewAtlasWithDefaults instantiates a new ApiMetricDataPointViewAtlas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApiMetricDataPointViewAtlas) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiMetricDataPointViewAtlas) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiMetricDataPointViewAtlas) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiMetricDataPointViewAtlas) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *ApiMetricDataPointViewAtlas) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiMetricDataPointViewAtlas) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiMetricDataPointViewAtlas) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApiMetricDataPointViewAtlas) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


