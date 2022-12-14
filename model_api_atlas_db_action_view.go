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

// ApiAtlasDBActionView Privilege action that the role grants.
type ApiAtlasDBActionView struct {
	// Human-readable label that identifies the privilege action.
	Action string `json:"action"`
	// List of resources on which you grant the action.
	Resources []ApiAtlasDBResourceView `json:"resources"`
}

// NewApiAtlasDBActionView instantiates a new ApiAtlasDBActionView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasDBActionView(action string, resources []ApiAtlasDBResourceView) *ApiAtlasDBActionView {
	this := ApiAtlasDBActionView{}
	this.Action = action
	this.Resources = resources
	return &this
}

// NewApiAtlasDBActionViewWithDefaults instantiates a new ApiAtlasDBActionView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasDBActionViewWithDefaults() *ApiAtlasDBActionView {
	this := ApiAtlasDBActionView{}
	return &this
}

// GetAction returns the Action field value
func (o *ApiAtlasDBActionView) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDBActionView) GetActionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ApiAtlasDBActionView) SetAction(v string) {
	o.Action = v
}

// GetResources returns the Resources field value
func (o *ApiAtlasDBActionView) GetResources() []ApiAtlasDBResourceView {
	if o == nil {
		var ret []ApiAtlasDBResourceView
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDBActionView) GetResourcesOk() ([]ApiAtlasDBResourceView, bool) {
	if o == nil {
    return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *ApiAtlasDBActionView) SetResources(v []ApiAtlasDBResourceView) {
	o.Resources = v
}

func (o ApiAtlasDBActionView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasDBActionView struct {
	value *ApiAtlasDBActionView
	isSet bool
}

func (v NullableApiAtlasDBActionView) Get() *ApiAtlasDBActionView {
	return v.value
}

func (v *NullableApiAtlasDBActionView) Set(val *ApiAtlasDBActionView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasDBActionView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasDBActionView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasDBActionView(val *ApiAtlasDBActionView) *NullableApiAtlasDBActionView {
	return &NullableApiAtlasDBActionView{value: val, isSet: true}
}

func (v NullableApiAtlasDBActionView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasDBActionView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


