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

// ApiAtlasNDSLabelView Human-readable labels applied to this MongoDB Cloud component.
type ApiAtlasNDSLabelView struct {
	// Key applied to tag and categorize this component.
	Key *string `json:"key,omitempty"`
	// Value set to the Key applied to tag and categorize this component.
	Value *string `json:"value,omitempty"`
}

// NewApiAtlasNDSLabelView instantiates a new ApiAtlasNDSLabelView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasNDSLabelView() *ApiAtlasNDSLabelView {
	this := ApiAtlasNDSLabelView{}
	return &this
}

// NewApiAtlasNDSLabelViewWithDefaults instantiates a new ApiAtlasNDSLabelView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasNDSLabelViewWithDefaults() *ApiAtlasNDSLabelView {
	this := ApiAtlasNDSLabelView{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ApiAtlasNDSLabelView) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasNDSLabelView) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ApiAtlasNDSLabelView) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ApiAtlasNDSLabelView) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiAtlasNDSLabelView) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasNDSLabelView) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiAtlasNDSLabelView) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApiAtlasNDSLabelView) SetValue(v string) {
	o.Value = &v
}

func (o ApiAtlasNDSLabelView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasNDSLabelView struct {
	value *ApiAtlasNDSLabelView
	isSet bool
}

func (v NullableApiAtlasNDSLabelView) Get() *ApiAtlasNDSLabelView {
	return v.value
}

func (v *NullableApiAtlasNDSLabelView) Set(val *ApiAtlasNDSLabelView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasNDSLabelView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasNDSLabelView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasNDSLabelView(val *ApiAtlasNDSLabelView) *NullableApiAtlasNDSLabelView {
	return &NullableApiAtlasNDSLabelView{value: val, isSet: true}
}

func (v NullableApiAtlasNDSLabelView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasNDSLabelView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


