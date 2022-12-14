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

// ApiAtlasUserScopeView Range of resources available to this database user.
type ApiAtlasUserScopeView struct {
	// Human-readable label that identifies the cluster or MongoDB Atlas Data Lake that this database user can access.
	Name string `json:"name"`
	// Category of resource that this database user can access.
	Type string `json:"type"`
}

// NewApiAtlasUserScopeView instantiates a new ApiAtlasUserScopeView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasUserScopeView(name string, type_ string) *ApiAtlasUserScopeView {
	this := ApiAtlasUserScopeView{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewApiAtlasUserScopeViewWithDefaults instantiates a new ApiAtlasUserScopeView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasUserScopeViewWithDefaults() *ApiAtlasUserScopeView {
	this := ApiAtlasUserScopeView{}
	return &this
}

// GetName returns the Name field value
func (o *ApiAtlasUserScopeView) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasUserScopeView) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiAtlasUserScopeView) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ApiAtlasUserScopeView) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasUserScopeView) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiAtlasUserScopeView) SetType(v string) {
	o.Type = v
}

func (o ApiAtlasUserScopeView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasUserScopeView struct {
	value *ApiAtlasUserScopeView
	isSet bool
}

func (v NullableApiAtlasUserScopeView) Get() *ApiAtlasUserScopeView {
	return v.value
}

func (v *NullableApiAtlasUserScopeView) Set(val *ApiAtlasUserScopeView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasUserScopeView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasUserScopeView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasUserScopeView(val *ApiAtlasUserScopeView) *NullableApiAtlasUserScopeView {
	return &NullableApiAtlasUserScopeView{value: val, isSet: true}
}

func (v NullableApiAtlasUserScopeView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasUserScopeView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


