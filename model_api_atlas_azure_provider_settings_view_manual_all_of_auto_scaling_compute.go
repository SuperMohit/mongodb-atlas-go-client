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

// ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute Collection of settings that configures how a cluster might scale its cluster tier and whether the cluster can scale down. Cluster tier auto-scaling is unavailable for clusters using Low CPU or NVME storage classes.
type ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute struct {
	// Maximum instance size to which your cluster can automatically scale.
	MaxInstanceSize *string `json:"maxInstanceSize,omitempty"`
	// Minimum instance size to which your cluster can automatically scale.
	MinInstanceSize *string `json:"minInstanceSize,omitempty"`
}

// NewApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute instantiates a new ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute() *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute {
	this := ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute{}
	return &this
}

// NewApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingComputeWithDefaults instantiates a new ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingComputeWithDefaults() *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute {
	this := ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute{}
	return &this
}

// GetMaxInstanceSize returns the MaxInstanceSize field value if set, zero value otherwise.
func (o *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) GetMaxInstanceSize() string {
	if o == nil || isNil(o.MaxInstanceSize) {
		var ret string
		return ret
	}
	return *o.MaxInstanceSize
}

// GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) GetMaxInstanceSizeOk() (*string, bool) {
	if o == nil || isNil(o.MaxInstanceSize) {
    return nil, false
	}
	return o.MaxInstanceSize, true
}

// HasMaxInstanceSize returns a boolean if a field has been set.
func (o *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) HasMaxInstanceSize() bool {
	if o != nil && !isNil(o.MaxInstanceSize) {
		return true
	}

	return false
}

// SetMaxInstanceSize gets a reference to the given string and assigns it to the MaxInstanceSize field.
func (o *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) SetMaxInstanceSize(v string) {
	o.MaxInstanceSize = &v
}

// GetMinInstanceSize returns the MinInstanceSize field value if set, zero value otherwise.
func (o *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) GetMinInstanceSize() string {
	if o == nil || isNil(o.MinInstanceSize) {
		var ret string
		return ret
	}
	return *o.MinInstanceSize
}

// GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) GetMinInstanceSizeOk() (*string, bool) {
	if o == nil || isNil(o.MinInstanceSize) {
    return nil, false
	}
	return o.MinInstanceSize, true
}

// HasMinInstanceSize returns a boolean if a field has been set.
func (o *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) HasMinInstanceSize() bool {
	if o != nil && !isNil(o.MinInstanceSize) {
		return true
	}

	return false
}

// SetMinInstanceSize gets a reference to the given string and assigns it to the MinInstanceSize field.
func (o *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) SetMinInstanceSize(v string) {
	o.MinInstanceSize = &v
}

func (o ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxInstanceSize) {
		toSerialize["maxInstanceSize"] = o.MaxInstanceSize
	}
	if !isNil(o.MinInstanceSize) {
		toSerialize["minInstanceSize"] = o.MinInstanceSize
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute struct {
	value *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute
	isSet bool
}

func (v NullableApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) Get() *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute {
	return v.value
}

func (v *NullableApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) Set(val *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute(val *ApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) *NullableApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute {
	return &NullableApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute{value: val, isSet: true}
}

func (v NullableApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasAzureProviderSettingsViewManualAllOfAutoScalingCompute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

