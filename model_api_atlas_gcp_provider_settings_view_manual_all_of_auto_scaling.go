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

// ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling Range of instance sizes to which your cluster can scale.
type ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling struct {
	Compute *ApiAtlasGCPProviderSettingsViewManualAllOfAutoScalingCompute `json:"compute,omitempty"`
}

// NewApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling instantiates a new ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling() *ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling {
	this := ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling{}
	return &this
}

// NewApiAtlasGCPProviderSettingsViewManualAllOfAutoScalingWithDefaults instantiates a new ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasGCPProviderSettingsViewManualAllOfAutoScalingWithDefaults() *ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling {
	this := ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) GetCompute() ApiAtlasGCPProviderSettingsViewManualAllOfAutoScalingCompute {
	if o == nil || isNil(o.Compute) {
		var ret ApiAtlasGCPProviderSettingsViewManualAllOfAutoScalingCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) GetComputeOk() (*ApiAtlasGCPProviderSettingsViewManualAllOfAutoScalingCompute, bool) {
	if o == nil || isNil(o.Compute) {
    return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) HasCompute() bool {
	if o != nil && !isNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given ApiAtlasGCPProviderSettingsViewManualAllOfAutoScalingCompute and assigns it to the Compute field.
func (o *ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) SetCompute(v ApiAtlasGCPProviderSettingsViewManualAllOfAutoScalingCompute) {
	o.Compute = &v
}

func (o ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling struct {
	value *ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling
	isSet bool
}

func (v NullableApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) Get() *ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling {
	return v.value
}

func (v *NullableApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) Set(val *ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling(val *ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) *NullableApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling {
	return &NullableApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling{value: val, isSet: true}
}

func (v NullableApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


