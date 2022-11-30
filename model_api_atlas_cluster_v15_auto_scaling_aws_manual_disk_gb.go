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

// ApiAtlasClusterV15AutoScalingAWSManualDiskGB Automatic cluster storage settings that apply to this cluster.
type ApiAtlasClusterV15AutoScalingAWSManualDiskGB struct {
	// Flag that indicates whether this cluster enables disk auto-scaling. The maximum memory allowed for the selected cluster tier and the oplog size can limit storage auto-scaling.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewApiAtlasClusterV15AutoScalingAWSManualDiskGB instantiates a new ApiAtlasClusterV15AutoScalingAWSManualDiskGB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasClusterV15AutoScalingAWSManualDiskGB() *ApiAtlasClusterV15AutoScalingAWSManualDiskGB {
	this := ApiAtlasClusterV15AutoScalingAWSManualDiskGB{}
	return &this
}

// NewApiAtlasClusterV15AutoScalingAWSManualDiskGBWithDefaults instantiates a new ApiAtlasClusterV15AutoScalingAWSManualDiskGB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasClusterV15AutoScalingAWSManualDiskGBWithDefaults() *ApiAtlasClusterV15AutoScalingAWSManualDiskGB {
	this := ApiAtlasClusterV15AutoScalingAWSManualDiskGB{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiAtlasClusterV15AutoScalingAWSManualDiskGB) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterV15AutoScalingAWSManualDiskGB) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiAtlasClusterV15AutoScalingAWSManualDiskGB) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiAtlasClusterV15AutoScalingAWSManualDiskGB) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ApiAtlasClusterV15AutoScalingAWSManualDiskGB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasClusterV15AutoScalingAWSManualDiskGB struct {
	value *ApiAtlasClusterV15AutoScalingAWSManualDiskGB
	isSet bool
}

func (v NullableApiAtlasClusterV15AutoScalingAWSManualDiskGB) Get() *ApiAtlasClusterV15AutoScalingAWSManualDiskGB {
	return v.value
}

func (v *NullableApiAtlasClusterV15AutoScalingAWSManualDiskGB) Set(val *ApiAtlasClusterV15AutoScalingAWSManualDiskGB) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasClusterV15AutoScalingAWSManualDiskGB) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasClusterV15AutoScalingAWSManualDiskGB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasClusterV15AutoScalingAWSManualDiskGB(val *ApiAtlasClusterV15AutoScalingAWSManualDiskGB) *NullableApiAtlasClusterV15AutoScalingAWSManualDiskGB {
	return &NullableApiAtlasClusterV15AutoScalingAWSManualDiskGB{value: val, isSet: true}
}

func (v NullableApiAtlasClusterV15AutoScalingAWSManualDiskGB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasClusterV15AutoScalingAWSManualDiskGB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


