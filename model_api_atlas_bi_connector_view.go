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

// ApiAtlasBiConnectorView Settings needed to configure the MongoDB Connector for Business Intelligence for this cluster.
type ApiAtlasBiConnectorView struct {
	// Flag that indicates whether MongoDB Connector for Business Intelligence is enabled on the specified cluster.
	Enabled *bool `json:"enabled,omitempty"`
	// Data source node designated for the MongoDB Connector for Business Intelligence on MongoDB Cloud. The MongoDB Connector for Business Intelligence on MongoDB Cloud reads data from the primary, secondary, or analytics node based on your read preferences. Defaults to `ANALYTICS` node, or `SECONDARY` if there are no `ANALYTICS` nodes.
	ReadPreference *string `json:"readPreference,omitempty"`
}

// NewApiAtlasBiConnectorView instantiates a new ApiAtlasBiConnectorView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasBiConnectorView() *ApiAtlasBiConnectorView {
	this := ApiAtlasBiConnectorView{}
	return &this
}

// NewApiAtlasBiConnectorViewWithDefaults instantiates a new ApiAtlasBiConnectorView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasBiConnectorViewWithDefaults() *ApiAtlasBiConnectorView {
	this := ApiAtlasBiConnectorView{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiAtlasBiConnectorView) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasBiConnectorView) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiAtlasBiConnectorView) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiAtlasBiConnectorView) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetReadPreference returns the ReadPreference field value if set, zero value otherwise.
func (o *ApiAtlasBiConnectorView) GetReadPreference() string {
	if o == nil || isNil(o.ReadPreference) {
		var ret string
		return ret
	}
	return *o.ReadPreference
}

// GetReadPreferenceOk returns a tuple with the ReadPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasBiConnectorView) GetReadPreferenceOk() (*string, bool) {
	if o == nil || isNil(o.ReadPreference) {
    return nil, false
	}
	return o.ReadPreference, true
}

// HasReadPreference returns a boolean if a field has been set.
func (o *ApiAtlasBiConnectorView) HasReadPreference() bool {
	if o != nil && !isNil(o.ReadPreference) {
		return true
	}

	return false
}

// SetReadPreference gets a reference to the given string and assigns it to the ReadPreference field.
func (o *ApiAtlasBiConnectorView) SetReadPreference(v string) {
	o.ReadPreference = &v
}

func (o ApiAtlasBiConnectorView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.ReadPreference) {
		toSerialize["readPreference"] = o.ReadPreference
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasBiConnectorView struct {
	value *ApiAtlasBiConnectorView
	isSet bool
}

func (v NullableApiAtlasBiConnectorView) Get() *ApiAtlasBiConnectorView {
	return v.value
}

func (v *NullableApiAtlasBiConnectorView) Set(val *ApiAtlasBiConnectorView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasBiConnectorView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasBiConnectorView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasBiConnectorView(val *ApiAtlasBiConnectorView) *NullableApiAtlasBiConnectorView {
	return &NullableApiAtlasBiConnectorView{value: val, isSet: true}
}

func (v NullableApiAtlasBiConnectorView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasBiConnectorView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


