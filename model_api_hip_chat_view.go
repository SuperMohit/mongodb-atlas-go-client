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

// ApiHipChatView struct for ApiHipChatView
type ApiHipChatView struct {
}

// NewApiHipChatView instantiates a new ApiHipChatView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHipChatView() *ApiHipChatView {
	this := ApiHipChatView{}
	return &this
}

// NewApiHipChatViewWithDefaults instantiates a new ApiHipChatView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHipChatViewWithDefaults() *ApiHipChatView {
	this := ApiHipChatView{}
	return &this
}

func (o ApiHipChatView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableApiHipChatView struct {
	value *ApiHipChatView
	isSet bool
}

func (v NullableApiHipChatView) Get() *ApiHipChatView {
	return v.value
}

func (v *NullableApiHipChatView) Set(val *ApiHipChatView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHipChatView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHipChatView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHipChatView(val *ApiHipChatView) *NullableApiHipChatView {
	return &NullableApiHipChatView{value: val, isSet: true}
}

func (v NullableApiHipChatView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHipChatView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


