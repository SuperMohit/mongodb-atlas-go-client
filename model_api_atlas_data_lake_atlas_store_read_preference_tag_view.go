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

// ApiAtlasDataLakeAtlasStoreReadPreferenceTagView List that contains [tag sets](https://docs.mongodb.com/manual/core/read-preference-tags/) or tag specification documents. If specified, Atlas Data Lake routes read requests to replica set member or members that are associated with the specified tags.
type ApiAtlasDataLakeAtlasStoreReadPreferenceTagView struct {
	// Human-readable label of the tag.
	Name *string `json:"name,omitempty"`
	// Value of the tag.
	Value *string `json:"value,omitempty"`
}

// NewApiAtlasDataLakeAtlasStoreReadPreferenceTagView instantiates a new ApiAtlasDataLakeAtlasStoreReadPreferenceTagView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasDataLakeAtlasStoreReadPreferenceTagView() *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView {
	this := ApiAtlasDataLakeAtlasStoreReadPreferenceTagView{}
	return &this
}

// NewApiAtlasDataLakeAtlasStoreReadPreferenceTagViewWithDefaults instantiates a new ApiAtlasDataLakeAtlasStoreReadPreferenceTagView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasDataLakeAtlasStoreReadPreferenceTagViewWithDefaults() *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView {
	this := ApiAtlasDataLakeAtlasStoreReadPreferenceTagView{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView) SetValue(v string) {
	o.Value = &v
}

func (o ApiAtlasDataLakeAtlasStoreReadPreferenceTagView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasDataLakeAtlasStoreReadPreferenceTagView struct {
	value *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView
	isSet bool
}

func (v NullableApiAtlasDataLakeAtlasStoreReadPreferenceTagView) Get() *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView {
	return v.value
}

func (v *NullableApiAtlasDataLakeAtlasStoreReadPreferenceTagView) Set(val *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasDataLakeAtlasStoreReadPreferenceTagView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasDataLakeAtlasStoreReadPreferenceTagView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasDataLakeAtlasStoreReadPreferenceTagView(val *ApiAtlasDataLakeAtlasStoreReadPreferenceTagView) *NullableApiAtlasDataLakeAtlasStoreReadPreferenceTagView {
	return &NullableApiAtlasDataLakeAtlasStoreReadPreferenceTagView{value: val, isSet: true}
}

func (v NullableApiAtlasDataLakeAtlasStoreReadPreferenceTagView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasDataLakeAtlasStoreReadPreferenceTagView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


