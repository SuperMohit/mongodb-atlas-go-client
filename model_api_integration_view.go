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

// ApiIntegrationView Collection of settings that describe third-party integrations.
type ApiIntegrationView struct {
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
}

// NewApiIntegrationView instantiates a new ApiIntegrationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiIntegrationView() *ApiIntegrationView {
	this := ApiIntegrationView{}
	return &this
}

// NewApiIntegrationViewWithDefaults instantiates a new ApiIntegrationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiIntegrationViewWithDefaults() *ApiIntegrationView {
	this := ApiIntegrationView{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiIntegrationView) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIntegrationView) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiIntegrationView) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiIntegrationView) SetType(v string) {
	o.Type = &v
}

func (o ApiIntegrationView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableApiIntegrationView struct {
	value *ApiIntegrationView
	isSet bool
}

func (v NullableApiIntegrationView) Get() *ApiIntegrationView {
	return v.value
}

func (v *NullableApiIntegrationView) Set(val *ApiIntegrationView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiIntegrationView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiIntegrationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiIntegrationView(val *ApiIntegrationView) *NullableApiIntegrationView {
	return &NullableApiIntegrationView{value: val, isSet: true}
}

func (v NullableApiIntegrationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiIntegrationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


