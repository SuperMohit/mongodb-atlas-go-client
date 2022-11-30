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

// ApiAddUserToTeamView struct for ApiAddUserToTeamView
type ApiAddUserToTeamView struct {
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user.
	Id string `json:"id"`
}

// NewApiAddUserToTeamView instantiates a new ApiAddUserToTeamView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAddUserToTeamView(id string) *ApiAddUserToTeamView {
	this := ApiAddUserToTeamView{}
	this.Id = id
	return &this
}

// NewApiAddUserToTeamViewWithDefaults instantiates a new ApiAddUserToTeamView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAddUserToTeamViewWithDefaults() *ApiAddUserToTeamView {
	this := ApiAddUserToTeamView{}
	return &this
}

// GetId returns the Id field value
func (o *ApiAddUserToTeamView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiAddUserToTeamView) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiAddUserToTeamView) SetId(v string) {
	o.Id = v
}

func (o ApiAddUserToTeamView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableApiAddUserToTeamView struct {
	value *ApiAddUserToTeamView
	isSet bool
}

func (v NullableApiAddUserToTeamView) Get() *ApiAddUserToTeamView {
	return v.value
}

func (v *NullableApiAddUserToTeamView) Set(val *ApiAddUserToTeamView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAddUserToTeamView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAddUserToTeamView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAddUserToTeamView(val *ApiAddUserToTeamView) *NullableApiAddUserToTeamView {
	return &NullableApiAddUserToTeamView{value: val, isSet: true}
}

func (v NullableApiAddUserToTeamView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAddUserToTeamView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


