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

// ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling Range of instance sizes to which your cluster can scale.
type ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling struct {
	Compute *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScalingCompute `json:"compute,omitempty"`
}

// NewApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling instantiates a new ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling() *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling {
	this := ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling{}
	return &this
}

// NewApiAtlasFreeProviderSettingsViewManualAllOfAutoScalingWithDefaults instantiates a new ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasFreeProviderSettingsViewManualAllOfAutoScalingWithDefaults() *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling {
	this := ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) GetCompute() ApiAtlasFreeProviderSettingsViewManualAllOfAutoScalingCompute {
	if o == nil || isNil(o.Compute) {
		var ret ApiAtlasFreeProviderSettingsViewManualAllOfAutoScalingCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) GetComputeOk() (*ApiAtlasFreeProviderSettingsViewManualAllOfAutoScalingCompute, bool) {
	if o == nil || isNil(o.Compute) {
    return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) HasCompute() bool {
	if o != nil && !isNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given ApiAtlasFreeProviderSettingsViewManualAllOfAutoScalingCompute and assigns it to the Compute field.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) SetCompute(v ApiAtlasFreeProviderSettingsViewManualAllOfAutoScalingCompute) {
	o.Compute = &v
}

func (o ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling struct {
	value *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling
	isSet bool
}

func (v NullableApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) Get() *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling {
	return v.value
}

func (v *NullableApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) Set(val *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling(val *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) *NullableApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling {
	return &NullableApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling{value: val, isSet: true}
}

func (v NullableApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


