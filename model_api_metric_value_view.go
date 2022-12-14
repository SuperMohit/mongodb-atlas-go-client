/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApiMetricValueView Value of the metric that triggered the alert. The resource returns this parameter for alerts of events impacting hosts.
type ApiMetricValueView struct {
	// Amount of the **metricName** recorded at the time of the event. This value triggered the alert.
	Number *float32 `json:"number,omitempty"`
	// Element used to express the quantity in **currentValue.number**. This can be an element of time, storage capacity, and the like. This metric triggered the alert.
	Units *string `json:"units,omitempty"`
}

// NewApiMetricValueView instantiates a new ApiMetricValueView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMetricValueView() *ApiMetricValueView {
	this := ApiMetricValueView{}
	return &this
}

// NewApiMetricValueViewWithDefaults instantiates a new ApiMetricValueView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMetricValueViewWithDefaults() *ApiMetricValueView {
	this := ApiMetricValueView{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ApiMetricValueView) GetNumber() float32 {
	if o == nil || isNil(o.Number) {
		var ret float32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricValueView) GetNumberOk() (*float32, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ApiMetricValueView) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given float32 and assigns it to the Number field.
func (o *ApiMetricValueView) SetNumber(v float32) {
	o.Number = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *ApiMetricValueView) GetUnits() string {
	if o == nil || isNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMetricValueView) GetUnitsOk() (*string, bool) {
	if o == nil || isNil(o.Units) {
    return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *ApiMetricValueView) HasUnits() bool {
	if o != nil && !isNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *ApiMetricValueView) SetUnits(v string) {
	o.Units = &v
}

func (o ApiMetricValueView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !isNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	return json.Marshal(toSerialize)
}

type NullableApiMetricValueView struct {
	value *ApiMetricValueView
	isSet bool
}

func (v NullableApiMetricValueView) Get() *ApiMetricValueView {
	return v.value
}

func (v *NullableApiMetricValueView) Set(val *ApiMetricValueView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMetricValueView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMetricValueView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMetricValueView(val *ApiMetricValueView) *NullableApiMetricValueView {
	return &NullableApiMetricValueView{value: val, isSet: true}
}

func (v NullableApiMetricValueView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMetricValueView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


