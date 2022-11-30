/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ApiBSONTimestampView BSON timestamp that indicates when the checkpoint token entry in the oplog occurred.
type ApiBSONTimestampView struct {
	// Date and time when the oplog recorded this database operation. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Date *time.Time `json:"date,omitempty"`
	// Order of the database operation that the oplog recorded at specific date and time.
	Increment *int32 `json:"increment,omitempty"`
}

// NewApiBSONTimestampView instantiates a new ApiBSONTimestampView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBSONTimestampView() *ApiBSONTimestampView {
	this := ApiBSONTimestampView{}
	return &this
}

// NewApiBSONTimestampViewWithDefaults instantiates a new ApiBSONTimestampView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBSONTimestampViewWithDefaults() *ApiBSONTimestampView {
	this := ApiBSONTimestampView{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ApiBSONTimestampView) GetDate() time.Time {
	if o == nil || isNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBSONTimestampView) GetDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.Date) {
    return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ApiBSONTimestampView) HasDate() bool {
	if o != nil && !isNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *ApiBSONTimestampView) SetDate(v time.Time) {
	o.Date = &v
}

// GetIncrement returns the Increment field value if set, zero value otherwise.
func (o *ApiBSONTimestampView) GetIncrement() int32 {
	if o == nil || isNil(o.Increment) {
		var ret int32
		return ret
	}
	return *o.Increment
}

// GetIncrementOk returns a tuple with the Increment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBSONTimestampView) GetIncrementOk() (*int32, bool) {
	if o == nil || isNil(o.Increment) {
    return nil, false
	}
	return o.Increment, true
}

// HasIncrement returns a boolean if a field has been set.
func (o *ApiBSONTimestampView) HasIncrement() bool {
	if o != nil && !isNil(o.Increment) {
		return true
	}

	return false
}

// SetIncrement gets a reference to the given int32 and assigns it to the Increment field.
func (o *ApiBSONTimestampView) SetIncrement(v int32) {
	o.Increment = &v
}

func (o ApiBSONTimestampView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !isNil(o.Increment) {
		toSerialize["increment"] = o.Increment
	}
	return json.Marshal(toSerialize)
}

type NullableApiBSONTimestampView struct {
	value *ApiBSONTimestampView
	isSet bool
}

func (v NullableApiBSONTimestampView) Get() *ApiBSONTimestampView {
	return v.value
}

func (v *NullableApiBSONTimestampView) Set(val *ApiBSONTimestampView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBSONTimestampView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBSONTimestampView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBSONTimestampView(val *ApiBSONTimestampView) *NullableApiBSONTimestampView {
	return &NullableApiBSONTimestampView{value: val, isSet: true}
}

func (v NullableApiBSONTimestampView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBSONTimestampView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


