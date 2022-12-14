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

// ApiCreateApiKeyView struct for ApiCreateApiKeyView
type ApiCreateApiKeyView struct {
	// Purpose or explanation provided when someone created this organization API key.
	Desc *string `json:"desc,omitempty"`
	// List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this organization or project.
	Roles []string `json:"roles,omitempty"`
}

// NewApiCreateApiKeyView instantiates a new ApiCreateApiKeyView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCreateApiKeyView() *ApiCreateApiKeyView {
	this := ApiCreateApiKeyView{}
	return &this
}

// NewApiCreateApiKeyViewWithDefaults instantiates a new ApiCreateApiKeyView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCreateApiKeyViewWithDefaults() *ApiCreateApiKeyView {
	this := ApiCreateApiKeyView{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *ApiCreateApiKeyView) GetDesc() string {
	if o == nil || isNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCreateApiKeyView) GetDescOk() (*string, bool) {
	if o == nil || isNil(o.Desc) {
    return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *ApiCreateApiKeyView) HasDesc() bool {
	if o != nil && !isNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *ApiCreateApiKeyView) SetDesc(v string) {
	o.Desc = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApiCreateApiKeyView) GetRoles() []string {
	if o == nil || isNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCreateApiKeyView) GetRolesOk() ([]string, bool) {
	if o == nil || isNil(o.Roles) {
    return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApiCreateApiKeyView) HasRoles() bool {
	if o != nil && !isNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *ApiCreateApiKeyView) SetRoles(v []string) {
	o.Roles = v
}

func (o ApiCreateApiKeyView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !isNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableApiCreateApiKeyView struct {
	value *ApiCreateApiKeyView
	isSet bool
}

func (v NullableApiCreateApiKeyView) Get() *ApiCreateApiKeyView {
	return v.value
}

func (v *NullableApiCreateApiKeyView) Set(val *ApiCreateApiKeyView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCreateApiKeyView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCreateApiKeyView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCreateApiKeyView(val *ApiCreateApiKeyView) *NullableApiCreateApiKeyView {
	return &NullableApiCreateApiKeyView{value: val, isSet: true}
}

func (v NullableApiCreateApiKeyView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCreateApiKeyView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


