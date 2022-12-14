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

// ApiFTSMetricView Measurement of one Atlas Search status when MongoDB Atlas received this request.
type ApiFTSMetricView struct {
	// Human-readable label that identifies this Atlas Search hardware, status, or index measurement.
	MetricName string `json:"metricName"`
	// Unit of measurement that applies to this Atlas Search metric.
	Units string `json:"units"`
}

// NewApiFTSMetricView instantiates a new ApiFTSMetricView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiFTSMetricView(metricName string, units string) *ApiFTSMetricView {
	this := ApiFTSMetricView{}
	this.MetricName = metricName
	this.Units = units
	return &this
}

// NewApiFTSMetricViewWithDefaults instantiates a new ApiFTSMetricView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiFTSMetricViewWithDefaults() *ApiFTSMetricView {
	this := ApiFTSMetricView{}
	return &this
}

// GetMetricName returns the MetricName field value
func (o *ApiFTSMetricView) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *ApiFTSMetricView) GetMetricNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *ApiFTSMetricView) SetMetricName(v string) {
	o.MetricName = v
}

// GetUnits returns the Units field value
func (o *ApiFTSMetricView) GetUnits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *ApiFTSMetricView) GetUnitsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *ApiFTSMetricView) SetUnits(v string) {
	o.Units = v
}

func (o ApiFTSMetricView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metricName"] = o.MetricName
	}
	if true {
		toSerialize["units"] = o.Units
	}
	return json.Marshal(toSerialize)
}

type NullableApiFTSMetricView struct {
	value *ApiFTSMetricView
	isSet bool
}

func (v NullableApiFTSMetricView) Get() *ApiFTSMetricView {
	return v.value
}

func (v *NullableApiFTSMetricView) Set(val *ApiFTSMetricView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFTSMetricView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFTSMetricView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFTSMetricView(val *ApiFTSMetricView) *NullableApiFTSMetricView {
	return &NullableApiFTSMetricView{value: val, isSet: true}
}

func (v NullableApiFTSMetricView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFTSMetricView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


