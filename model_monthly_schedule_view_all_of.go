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

// MonthlyScheduleViewAllOf struct for MonthlyScheduleViewAllOf
type MonthlyScheduleViewAllOf struct {
	// Day of the month when the scheduled archive starts.
	DayOfMonth *int32 `json:"dayOfMonth,omitempty"`
	// Hour of the day when the scheduled window to run one online archive ends.
	EndHour *int32 `json:"endHour,omitempty"`
	// Minute of the hour when the scheduled window to run one online archive ends.
	EndMinute *int32 `json:"endMinute,omitempty"`
	// Hour of the day when the when the scheduled window to run one online archive starts.
	StartHour *int32 `json:"startHour,omitempty"`
	// Minute of the hour when the scheduled window to run one online archive starts.
	StartMinute *int32 `json:"startMinute,omitempty"`
}

// NewMonthlyScheduleViewAllOf instantiates a new MonthlyScheduleViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonthlyScheduleViewAllOf() *MonthlyScheduleViewAllOf {
	this := MonthlyScheduleViewAllOf{}
	return &this
}

// NewMonthlyScheduleViewAllOfWithDefaults instantiates a new MonthlyScheduleViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonthlyScheduleViewAllOfWithDefaults() *MonthlyScheduleViewAllOf {
	this := MonthlyScheduleViewAllOf{}
	return &this
}

// GetDayOfMonth returns the DayOfMonth field value if set, zero value otherwise.
func (o *MonthlyScheduleViewAllOf) GetDayOfMonth() int32 {
	if o == nil || isNil(o.DayOfMonth) {
		var ret int32
		return ret
	}
	return *o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyScheduleViewAllOf) GetDayOfMonthOk() (*int32, bool) {
	if o == nil || isNil(o.DayOfMonth) {
    return nil, false
	}
	return o.DayOfMonth, true
}

// HasDayOfMonth returns a boolean if a field has been set.
func (o *MonthlyScheduleViewAllOf) HasDayOfMonth() bool {
	if o != nil && !isNil(o.DayOfMonth) {
		return true
	}

	return false
}

// SetDayOfMonth gets a reference to the given int32 and assigns it to the DayOfMonth field.
func (o *MonthlyScheduleViewAllOf) SetDayOfMonth(v int32) {
	o.DayOfMonth = &v
}

// GetEndHour returns the EndHour field value if set, zero value otherwise.
func (o *MonthlyScheduleViewAllOf) GetEndHour() int32 {
	if o == nil || isNil(o.EndHour) {
		var ret int32
		return ret
	}
	return *o.EndHour
}

// GetEndHourOk returns a tuple with the EndHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyScheduleViewAllOf) GetEndHourOk() (*int32, bool) {
	if o == nil || isNil(o.EndHour) {
    return nil, false
	}
	return o.EndHour, true
}

// HasEndHour returns a boolean if a field has been set.
func (o *MonthlyScheduleViewAllOf) HasEndHour() bool {
	if o != nil && !isNil(o.EndHour) {
		return true
	}

	return false
}

// SetEndHour gets a reference to the given int32 and assigns it to the EndHour field.
func (o *MonthlyScheduleViewAllOf) SetEndHour(v int32) {
	o.EndHour = &v
}

// GetEndMinute returns the EndMinute field value if set, zero value otherwise.
func (o *MonthlyScheduleViewAllOf) GetEndMinute() int32 {
	if o == nil || isNil(o.EndMinute) {
		var ret int32
		return ret
	}
	return *o.EndMinute
}

// GetEndMinuteOk returns a tuple with the EndMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyScheduleViewAllOf) GetEndMinuteOk() (*int32, bool) {
	if o == nil || isNil(o.EndMinute) {
    return nil, false
	}
	return o.EndMinute, true
}

// HasEndMinute returns a boolean if a field has been set.
func (o *MonthlyScheduleViewAllOf) HasEndMinute() bool {
	if o != nil && !isNil(o.EndMinute) {
		return true
	}

	return false
}

// SetEndMinute gets a reference to the given int32 and assigns it to the EndMinute field.
func (o *MonthlyScheduleViewAllOf) SetEndMinute(v int32) {
	o.EndMinute = &v
}

// GetStartHour returns the StartHour field value if set, zero value otherwise.
func (o *MonthlyScheduleViewAllOf) GetStartHour() int32 {
	if o == nil || isNil(o.StartHour) {
		var ret int32
		return ret
	}
	return *o.StartHour
}

// GetStartHourOk returns a tuple with the StartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyScheduleViewAllOf) GetStartHourOk() (*int32, bool) {
	if o == nil || isNil(o.StartHour) {
    return nil, false
	}
	return o.StartHour, true
}

// HasStartHour returns a boolean if a field has been set.
func (o *MonthlyScheduleViewAllOf) HasStartHour() bool {
	if o != nil && !isNil(o.StartHour) {
		return true
	}

	return false
}

// SetStartHour gets a reference to the given int32 and assigns it to the StartHour field.
func (o *MonthlyScheduleViewAllOf) SetStartHour(v int32) {
	o.StartHour = &v
}

// GetStartMinute returns the StartMinute field value if set, zero value otherwise.
func (o *MonthlyScheduleViewAllOf) GetStartMinute() int32 {
	if o == nil || isNil(o.StartMinute) {
		var ret int32
		return ret
	}
	return *o.StartMinute
}

// GetStartMinuteOk returns a tuple with the StartMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyScheduleViewAllOf) GetStartMinuteOk() (*int32, bool) {
	if o == nil || isNil(o.StartMinute) {
    return nil, false
	}
	return o.StartMinute, true
}

// HasStartMinute returns a boolean if a field has been set.
func (o *MonthlyScheduleViewAllOf) HasStartMinute() bool {
	if o != nil && !isNil(o.StartMinute) {
		return true
	}

	return false
}

// SetStartMinute gets a reference to the given int32 and assigns it to the StartMinute field.
func (o *MonthlyScheduleViewAllOf) SetStartMinute(v int32) {
	o.StartMinute = &v
}

func (o MonthlyScheduleViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DayOfMonth) {
		toSerialize["dayOfMonth"] = o.DayOfMonth
	}
	if !isNil(o.EndHour) {
		toSerialize["endHour"] = o.EndHour
	}
	if !isNil(o.EndMinute) {
		toSerialize["endMinute"] = o.EndMinute
	}
	if !isNil(o.StartHour) {
		toSerialize["startHour"] = o.StartHour
	}
	if !isNil(o.StartMinute) {
		toSerialize["startMinute"] = o.StartMinute
	}
	return json.Marshal(toSerialize)
}

type NullableMonthlyScheduleViewAllOf struct {
	value *MonthlyScheduleViewAllOf
	isSet bool
}

func (v NullableMonthlyScheduleViewAllOf) Get() *MonthlyScheduleViewAllOf {
	return v.value
}

func (v *NullableMonthlyScheduleViewAllOf) Set(val *MonthlyScheduleViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMonthlyScheduleViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMonthlyScheduleViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonthlyScheduleViewAllOf(val *MonthlyScheduleViewAllOf) *NullableMonthlyScheduleViewAllOf {
	return &NullableMonthlyScheduleViewAllOf{value: val, isSet: true}
}

func (v NullableMonthlyScheduleViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonthlyScheduleViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


