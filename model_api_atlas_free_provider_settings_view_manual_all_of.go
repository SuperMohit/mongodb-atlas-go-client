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

// ApiAtlasFreeProviderSettingsViewManualAllOf struct for ApiAtlasFreeProviderSettingsViewManualAllOf
type ApiAtlasFreeProviderSettingsViewManualAllOf struct {
	AutoScaling *ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling `json:"autoScaling,omitempty"`
	// Cloud service provider on which MongoDB Cloud provisioned the multi-tenant host. The resource returns this parameter when **providerSettings.providerName** is `TENANT` and **providerSetting.instanceSizeName** is `M2` or `M5`.
	BackingProviderName *string `json:"backingProviderName,omitempty"`
	// Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. You must set **providerSettings.providerName** to `TENANT` and specify the cloud service provider in **providerSettings.backingProviderName**.
	InstanceSizeName *string `json:"instanceSizeName,omitempty"`
	// Human-readable label that identifies the geographic location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/). For multi-region clusters, see **replicationSpec.{region}**.
	RegionName *string `json:"regionName,omitempty"`
}

// NewApiAtlasFreeProviderSettingsViewManualAllOf instantiates a new ApiAtlasFreeProviderSettingsViewManualAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasFreeProviderSettingsViewManualAllOf() *ApiAtlasFreeProviderSettingsViewManualAllOf {
	this := ApiAtlasFreeProviderSettingsViewManualAllOf{}
	return &this
}

// NewApiAtlasFreeProviderSettingsViewManualAllOfWithDefaults instantiates a new ApiAtlasFreeProviderSettingsViewManualAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasFreeProviderSettingsViewManualAllOfWithDefaults() *ApiAtlasFreeProviderSettingsViewManualAllOf {
	this := ApiAtlasFreeProviderSettingsViewManualAllOf{}
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) GetAutoScaling() ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling {
	if o == nil || isNil(o.AutoScaling) {
		var ret ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) GetAutoScalingOk() (*ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling, bool) {
	if o == nil || isNil(o.AutoScaling) {
    return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) HasAutoScaling() bool {
	if o != nil && !isNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling and assigns it to the AutoScaling field.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) SetAutoScaling(v ApiAtlasFreeProviderSettingsViewManualAllOfAutoScaling) {
	o.AutoScaling = &v
}

// GetBackingProviderName returns the BackingProviderName field value if set, zero value otherwise.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) GetBackingProviderName() string {
	if o == nil || isNil(o.BackingProviderName) {
		var ret string
		return ret
	}
	return *o.BackingProviderName
}

// GetBackingProviderNameOk returns a tuple with the BackingProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) GetBackingProviderNameOk() (*string, bool) {
	if o == nil || isNil(o.BackingProviderName) {
    return nil, false
	}
	return o.BackingProviderName, true
}

// HasBackingProviderName returns a boolean if a field has been set.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) HasBackingProviderName() bool {
	if o != nil && !isNil(o.BackingProviderName) {
		return true
	}

	return false
}

// SetBackingProviderName gets a reference to the given string and assigns it to the BackingProviderName field.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) SetBackingProviderName(v string) {
	o.BackingProviderName = &v
}

// GetInstanceSizeName returns the InstanceSizeName field value if set, zero value otherwise.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) GetInstanceSizeName() string {
	if o == nil || isNil(o.InstanceSizeName) {
		var ret string
		return ret
	}
	return *o.InstanceSizeName
}

// GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) GetInstanceSizeNameOk() (*string, bool) {
	if o == nil || isNil(o.InstanceSizeName) {
    return nil, false
	}
	return o.InstanceSizeName, true
}

// HasInstanceSizeName returns a boolean if a field has been set.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) HasInstanceSizeName() bool {
	if o != nil && !isNil(o.InstanceSizeName) {
		return true
	}

	return false
}

// SetInstanceSizeName gets a reference to the given string and assigns it to the InstanceSizeName field.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) SetInstanceSizeName(v string) {
	o.InstanceSizeName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) GetRegionName() string {
	if o == nil || isNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) GetRegionNameOk() (*string, bool) {
	if o == nil || isNil(o.RegionName) {
    return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) HasRegionName() bool {
	if o != nil && !isNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ApiAtlasFreeProviderSettingsViewManualAllOf) SetRegionName(v string) {
	o.RegionName = &v
}

func (o ApiAtlasFreeProviderSettingsViewManualAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AutoScaling) {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if !isNil(o.BackingProviderName) {
		toSerialize["backingProviderName"] = o.BackingProviderName
	}
	if !isNil(o.InstanceSizeName) {
		toSerialize["instanceSizeName"] = o.InstanceSizeName
	}
	if !isNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasFreeProviderSettingsViewManualAllOf struct {
	value *ApiAtlasFreeProviderSettingsViewManualAllOf
	isSet bool
}

func (v NullableApiAtlasFreeProviderSettingsViewManualAllOf) Get() *ApiAtlasFreeProviderSettingsViewManualAllOf {
	return v.value
}

func (v *NullableApiAtlasFreeProviderSettingsViewManualAllOf) Set(val *ApiAtlasFreeProviderSettingsViewManualAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasFreeProviderSettingsViewManualAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasFreeProviderSettingsViewManualAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasFreeProviderSettingsViewManualAllOf(val *ApiAtlasFreeProviderSettingsViewManualAllOf) *NullableApiAtlasFreeProviderSettingsViewManualAllOf {
	return &NullableApiAtlasFreeProviderSettingsViewManualAllOf{value: val, isSet: true}
}

func (v NullableApiAtlasFreeProviderSettingsViewManualAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasFreeProviderSettingsViewManualAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


