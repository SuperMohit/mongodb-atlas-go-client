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

// ApiAtlasRegionSpecView Physical location where MongoDB Cloud provisions cluster nodes.
type ApiAtlasRegionSpecView struct {
	// Number of analytics nodes in the region. Analytics nodes handle analytic data such as reporting queries from MongoDB Connector for Business Intelligence on MongoDB Cloud. Analytics nodes are read-only, and can never become the primary. Use **replicationSpecs[n].{region}.analyticsNodes** instead.
	AnalyticsNodes *int32 `json:"analyticsNodes,omitempty"`
	// Number of electable nodes to deploy in the specified region. Electable nodes can become the primary and can facilitate local reads. Use **replicationSpecs[n].{region}.electableNodes** instead.
	ElectableNodes *int32 `json:"electableNodes,omitempty"`
	// Number that indicates the election priority of the region. To identify the Preferred Region of the cluster, set this parameter to `7`. The primary node runs in the **Preferred Region**. To identify a read-only region, set this parameter to `0`.
	Priority *int32 `json:"priority,omitempty"`
	// Number of read-only nodes in the region. Read-only nodes can never become the primary member, but can facilitate local reads. Use **replicationSpecs[n].{region}.readOnlyNodes** instead.
	ReadOnlyNodes *int32 `json:"readOnlyNodes,omitempty"`
}

// NewApiAtlasRegionSpecView instantiates a new ApiAtlasRegionSpecView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasRegionSpecView() *ApiAtlasRegionSpecView {
	this := ApiAtlasRegionSpecView{}
	return &this
}

// NewApiAtlasRegionSpecViewWithDefaults instantiates a new ApiAtlasRegionSpecView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasRegionSpecViewWithDefaults() *ApiAtlasRegionSpecView {
	this := ApiAtlasRegionSpecView{}
	return &this
}

// GetAnalyticsNodes returns the AnalyticsNodes field value if set, zero value otherwise.
func (o *ApiAtlasRegionSpecView) GetAnalyticsNodes() int32 {
	if o == nil || isNil(o.AnalyticsNodes) {
		var ret int32
		return ret
	}
	return *o.AnalyticsNodes
}

// GetAnalyticsNodesOk returns a tuple with the AnalyticsNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasRegionSpecView) GetAnalyticsNodesOk() (*int32, bool) {
	if o == nil || isNil(o.AnalyticsNodes) {
    return nil, false
	}
	return o.AnalyticsNodes, true
}

// HasAnalyticsNodes returns a boolean if a field has been set.
func (o *ApiAtlasRegionSpecView) HasAnalyticsNodes() bool {
	if o != nil && !isNil(o.AnalyticsNodes) {
		return true
	}

	return false
}

// SetAnalyticsNodes gets a reference to the given int32 and assigns it to the AnalyticsNodes field.
func (o *ApiAtlasRegionSpecView) SetAnalyticsNodes(v int32) {
	o.AnalyticsNodes = &v
}

// GetElectableNodes returns the ElectableNodes field value if set, zero value otherwise.
func (o *ApiAtlasRegionSpecView) GetElectableNodes() int32 {
	if o == nil || isNil(o.ElectableNodes) {
		var ret int32
		return ret
	}
	return *o.ElectableNodes
}

// GetElectableNodesOk returns a tuple with the ElectableNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasRegionSpecView) GetElectableNodesOk() (*int32, bool) {
	if o == nil || isNil(o.ElectableNodes) {
    return nil, false
	}
	return o.ElectableNodes, true
}

// HasElectableNodes returns a boolean if a field has been set.
func (o *ApiAtlasRegionSpecView) HasElectableNodes() bool {
	if o != nil && !isNil(o.ElectableNodes) {
		return true
	}

	return false
}

// SetElectableNodes gets a reference to the given int32 and assigns it to the ElectableNodes field.
func (o *ApiAtlasRegionSpecView) SetElectableNodes(v int32) {
	o.ElectableNodes = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ApiAtlasRegionSpecView) GetPriority() int32 {
	if o == nil || isNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasRegionSpecView) GetPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ApiAtlasRegionSpecView) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ApiAtlasRegionSpecView) SetPriority(v int32) {
	o.Priority = &v
}

// GetReadOnlyNodes returns the ReadOnlyNodes field value if set, zero value otherwise.
func (o *ApiAtlasRegionSpecView) GetReadOnlyNodes() int32 {
	if o == nil || isNil(o.ReadOnlyNodes) {
		var ret int32
		return ret
	}
	return *o.ReadOnlyNodes
}

// GetReadOnlyNodesOk returns a tuple with the ReadOnlyNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasRegionSpecView) GetReadOnlyNodesOk() (*int32, bool) {
	if o == nil || isNil(o.ReadOnlyNodes) {
    return nil, false
	}
	return o.ReadOnlyNodes, true
}

// HasReadOnlyNodes returns a boolean if a field has been set.
func (o *ApiAtlasRegionSpecView) HasReadOnlyNodes() bool {
	if o != nil && !isNil(o.ReadOnlyNodes) {
		return true
	}

	return false
}

// SetReadOnlyNodes gets a reference to the given int32 and assigns it to the ReadOnlyNodes field.
func (o *ApiAtlasRegionSpecView) SetReadOnlyNodes(v int32) {
	o.ReadOnlyNodes = &v
}

func (o ApiAtlasRegionSpecView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AnalyticsNodes) {
		toSerialize["analyticsNodes"] = o.AnalyticsNodes
	}
	if !isNil(o.ElectableNodes) {
		toSerialize["electableNodes"] = o.ElectableNodes
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.ReadOnlyNodes) {
		toSerialize["readOnlyNodes"] = o.ReadOnlyNodes
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasRegionSpecView struct {
	value *ApiAtlasRegionSpecView
	isSet bool
}

func (v NullableApiAtlasRegionSpecView) Get() *ApiAtlasRegionSpecView {
	return v.value
}

func (v *NullableApiAtlasRegionSpecView) Set(val *ApiAtlasRegionSpecView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasRegionSpecView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasRegionSpecView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasRegionSpecView(val *ApiAtlasRegionSpecView) *NullableApiAtlasRegionSpecView {
	return &NullableApiAtlasRegionSpecView{value: val, isSet: true}
}

func (v NullableApiAtlasRegionSpecView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasRegionSpecView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


