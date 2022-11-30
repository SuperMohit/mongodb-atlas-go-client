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

// ApiAtlasClusterHardwareSpecsV15TenantManual struct for ApiAtlasClusterHardwareSpecsV15TenantManual
type ApiAtlasClusterHardwareSpecsV15TenantManual struct {
	// Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size. If you deploy a Global Cluster, you must choose a instance size of `M30` or greater.
	InstanceSize *string `json:"instanceSize,omitempty"`
	// Number of read-only nodes for MongoDB Cloud deploys to the region. Read-only nodes can never become the primary, but can enable local reads.
	NodeCount *int32 `json:"nodeCount,omitempty"`
}

// NewApiAtlasClusterHardwareSpecsV15TenantManual instantiates a new ApiAtlasClusterHardwareSpecsV15TenantManual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasClusterHardwareSpecsV15TenantManual() *ApiAtlasClusterHardwareSpecsV15TenantManual {
	this := ApiAtlasClusterHardwareSpecsV15TenantManual{}
	return &this
}

// NewApiAtlasClusterHardwareSpecsV15TenantManualWithDefaults instantiates a new ApiAtlasClusterHardwareSpecsV15TenantManual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasClusterHardwareSpecsV15TenantManualWithDefaults() *ApiAtlasClusterHardwareSpecsV15TenantManual {
	this := ApiAtlasClusterHardwareSpecsV15TenantManual{}
	return &this
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise.
func (o *ApiAtlasClusterHardwareSpecsV15TenantManual) GetInstanceSize() string {
	if o == nil || isNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterHardwareSpecsV15TenantManual) GetInstanceSizeOk() (*string, bool) {
	if o == nil || isNil(o.InstanceSize) {
    return nil, false
	}
	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *ApiAtlasClusterHardwareSpecsV15TenantManual) HasInstanceSize() bool {
	if o != nil && !isNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *ApiAtlasClusterHardwareSpecsV15TenantManual) SetInstanceSize(v string) {
	o.InstanceSize = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *ApiAtlasClusterHardwareSpecsV15TenantManual) GetNodeCount() int32 {
	if o == nil || isNil(o.NodeCount) {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterHardwareSpecsV15TenantManual) GetNodeCountOk() (*int32, bool) {
	if o == nil || isNil(o.NodeCount) {
    return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *ApiAtlasClusterHardwareSpecsV15TenantManual) HasNodeCount() bool {
	if o != nil && !isNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *ApiAtlasClusterHardwareSpecsV15TenantManual) SetNodeCount(v int32) {
	o.NodeCount = &v
}

func (o ApiAtlasClusterHardwareSpecsV15TenantManual) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InstanceSize) {
		toSerialize["instanceSize"] = o.InstanceSize
	}
	if !isNil(o.NodeCount) {
		toSerialize["nodeCount"] = o.NodeCount
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasClusterHardwareSpecsV15TenantManual struct {
	value *ApiAtlasClusterHardwareSpecsV15TenantManual
	isSet bool
}

func (v NullableApiAtlasClusterHardwareSpecsV15TenantManual) Get() *ApiAtlasClusterHardwareSpecsV15TenantManual {
	return v.value
}

func (v *NullableApiAtlasClusterHardwareSpecsV15TenantManual) Set(val *ApiAtlasClusterHardwareSpecsV15TenantManual) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasClusterHardwareSpecsV15TenantManual) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasClusterHardwareSpecsV15TenantManual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasClusterHardwareSpecsV15TenantManual(val *ApiAtlasClusterHardwareSpecsV15TenantManual) *NullableApiAtlasClusterHardwareSpecsV15TenantManual {
	return &NullableApiAtlasClusterHardwareSpecsV15TenantManual{value: val, isSet: true}
}

func (v NullableApiAtlasClusterHardwareSpecsV15TenantManual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasClusterHardwareSpecsV15TenantManual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


