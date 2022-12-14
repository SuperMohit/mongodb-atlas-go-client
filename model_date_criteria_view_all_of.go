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

// DateCriteriaViewAllOf struct for DateCriteriaViewAllOf
type DateCriteriaViewAllOf struct {
	// Indexed database parameter that stores the date that determines when data moves to the online archive. MongoDB Cloud archives the data when the current date exceeds the date in this database parameter plus the number of days specified through the **expireAfterDays** parameter. Set this parameter when you set `\"criteria.type\" : \"DATE\"`.
	DateField *string `json:"dateField,omitempty"`
	// Syntax used to write the date after which data moves to the online archive. Date can be expressed as ISO 8601 or Epoch timestamps. The Epoch timestamp can be expressed as nanoseconds, milliseconds, or seconds. Set this parameter when **\"criteria.type\" : \"DATE\"**. You must set **\"criteria.type\" : \"DATE\"** if **\"collectionType\": \"TIMESERIES\"**.
	DateFormat *string `json:"dateFormat,omitempty"`
	// Number of days after the value in the **criteria.dateField** when MongoDB Cloud archives data in the specified cluster. Set this parameter when you set **\"criteria.type\" : \"DATE\"**.
	ExpireAfterDays *int32 `json:"expireAfterDays,omitempty"`
}

// NewDateCriteriaViewAllOf instantiates a new DateCriteriaViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateCriteriaViewAllOf() *DateCriteriaViewAllOf {
	this := DateCriteriaViewAllOf{}
	var dateFormat string = "ISODATE"
	this.DateFormat = &dateFormat
	return &this
}

// NewDateCriteriaViewAllOfWithDefaults instantiates a new DateCriteriaViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateCriteriaViewAllOfWithDefaults() *DateCriteriaViewAllOf {
	this := DateCriteriaViewAllOf{}
	var dateFormat string = "ISODATE"
	this.DateFormat = &dateFormat
	return &this
}

// GetDateField returns the DateField field value if set, zero value otherwise.
func (o *DateCriteriaViewAllOf) GetDateField() string {
	if o == nil || isNil(o.DateField) {
		var ret string
		return ret
	}
	return *o.DateField
}

// GetDateFieldOk returns a tuple with the DateField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateCriteriaViewAllOf) GetDateFieldOk() (*string, bool) {
	if o == nil || isNil(o.DateField) {
    return nil, false
	}
	return o.DateField, true
}

// HasDateField returns a boolean if a field has been set.
func (o *DateCriteriaViewAllOf) HasDateField() bool {
	if o != nil && !isNil(o.DateField) {
		return true
	}

	return false
}

// SetDateField gets a reference to the given string and assigns it to the DateField field.
func (o *DateCriteriaViewAllOf) SetDateField(v string) {
	o.DateField = &v
}

// GetDateFormat returns the DateFormat field value if set, zero value otherwise.
func (o *DateCriteriaViewAllOf) GetDateFormat() string {
	if o == nil || isNil(o.DateFormat) {
		var ret string
		return ret
	}
	return *o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateCriteriaViewAllOf) GetDateFormatOk() (*string, bool) {
	if o == nil || isNil(o.DateFormat) {
    return nil, false
	}
	return o.DateFormat, true
}

// HasDateFormat returns a boolean if a field has been set.
func (o *DateCriteriaViewAllOf) HasDateFormat() bool {
	if o != nil && !isNil(o.DateFormat) {
		return true
	}

	return false
}

// SetDateFormat gets a reference to the given string and assigns it to the DateFormat field.
func (o *DateCriteriaViewAllOf) SetDateFormat(v string) {
	o.DateFormat = &v
}

// GetExpireAfterDays returns the ExpireAfterDays field value if set, zero value otherwise.
func (o *DateCriteriaViewAllOf) GetExpireAfterDays() int32 {
	if o == nil || isNil(o.ExpireAfterDays) {
		var ret int32
		return ret
	}
	return *o.ExpireAfterDays
}

// GetExpireAfterDaysOk returns a tuple with the ExpireAfterDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateCriteriaViewAllOf) GetExpireAfterDaysOk() (*int32, bool) {
	if o == nil || isNil(o.ExpireAfterDays) {
    return nil, false
	}
	return o.ExpireAfterDays, true
}

// HasExpireAfterDays returns a boolean if a field has been set.
func (o *DateCriteriaViewAllOf) HasExpireAfterDays() bool {
	if o != nil && !isNil(o.ExpireAfterDays) {
		return true
	}

	return false
}

// SetExpireAfterDays gets a reference to the given int32 and assigns it to the ExpireAfterDays field.
func (o *DateCriteriaViewAllOf) SetExpireAfterDays(v int32) {
	o.ExpireAfterDays = &v
}

func (o DateCriteriaViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DateField) {
		toSerialize["dateField"] = o.DateField
	}
	if !isNil(o.DateFormat) {
		toSerialize["dateFormat"] = o.DateFormat
	}
	if !isNil(o.ExpireAfterDays) {
		toSerialize["expireAfterDays"] = o.ExpireAfterDays
	}
	return json.Marshal(toSerialize)
}

type NullableDateCriteriaViewAllOf struct {
	value *DateCriteriaViewAllOf
	isSet bool
}

func (v NullableDateCriteriaViewAllOf) Get() *DateCriteriaViewAllOf {
	return v.value
}

func (v *NullableDateCriteriaViewAllOf) Set(val *DateCriteriaViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDateCriteriaViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDateCriteriaViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateCriteriaViewAllOf(val *DateCriteriaViewAllOf) *NullableDateCriteriaViewAllOf {
	return &NullableDateCriteriaViewAllOf{value: val, isSet: true}
}

func (v NullableDateCriteriaViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateCriteriaViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


