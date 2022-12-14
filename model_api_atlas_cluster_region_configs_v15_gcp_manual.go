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

// ApiAtlasClusterRegionConfigsV15GCPManual Hardware specifications for nodes set for a given region. Each **regionConfigs** object describes the region's priority in elections and the number and type of MongoDB nodes that MongoDB Cloud deploys to the region. Each **regionConfigs** object must have either an **analyticsSpecs** object, **electableSpecs** object, or **readOnlySpecs** object. Tenant clusters only require **electableSpecs. Dedicated** clusters can specify any of these specifications, but must have at least one **electableSpecs** object within a **replicationSpec**. Every hardware specification must use the same **instanceSize**.  **Example:**  If you set `\\\"replicationSpecs[n].regionConfigs[m].analyticsSpecs.instanceSize\\\" : \\\"M30\\\"`, set `\\\"replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize\\\" : `\\\"M30\\\"` if you have electable nodes and `\\\"replicationSpecs[n].regionConfigs[m].readOnlySpecs.instanceSize\\\" : `\\\"M30\\\"` if you have read-only nodes.\",
type ApiAtlasClusterRegionConfigsV15GCPManual struct {
	AnalyticsAutoScaling *ApiAtlasClusterV15AutoScalingGCPManual `json:"analyticsAutoScaling,omitempty"`
	AnalyticsSpecs *ApiAtlasClusterHardwareSpecsV15GCPManual `json:"analyticsSpecs,omitempty"`
	AutoScaling *ApiAtlasClusterV15AutoScalingGCPManual `json:"autoScaling,omitempty"`
	ElectableSpecs *ApiAtlasClusterHardwareSpecsV15GCPManual `json:"electableSpecs,omitempty"`
	// Precedence that MongoDB Cloud gives this region when a primary election occurs. If your **regionConfigs** has only **readOnlySpecs**, **analyticsSpecs**, or both, set this value to `0`. If you have multiple **regionConfigs** objects, set priorities in descending order starting from `7`.  Example: If you have three regions, their priorities would be `7`, `6`, and `5` respectively. If you added two more regions for supporting electable nodes, the priorities of those regions would be `4` and `3` respectively. 
	Priority *int32 `json:"priority,omitempty"`
	// Cloud service provider on which MongoDB Cloud provisions the MongoDB nodes. 
	ProviderName *string `json:"providerName,omitempty"`
	ReadOnlySpecs *ApiAtlasClusterHardwareSpecsV15GCPManual `json:"readOnlySpecs,omitempty"`
	// Physical location where MongoDB Cloud deploys your Google Compute-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts.
	RegionName *string `json:"regionName,omitempty"`
}

// NewApiAtlasClusterRegionConfigsV15GCPManual instantiates a new ApiAtlasClusterRegionConfigsV15GCPManual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasClusterRegionConfigsV15GCPManual() *ApiAtlasClusterRegionConfigsV15GCPManual {
	this := ApiAtlasClusterRegionConfigsV15GCPManual{}
	return &this
}

// NewApiAtlasClusterRegionConfigsV15GCPManualWithDefaults instantiates a new ApiAtlasClusterRegionConfigsV15GCPManual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasClusterRegionConfigsV15GCPManualWithDefaults() *ApiAtlasClusterRegionConfigsV15GCPManual {
	this := ApiAtlasClusterRegionConfigsV15GCPManual{}
	return &this
}

// GetAnalyticsAutoScaling returns the AnalyticsAutoScaling field value if set, zero value otherwise.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetAnalyticsAutoScaling() ApiAtlasClusterV15AutoScalingGCPManual {
	if o == nil || isNil(o.AnalyticsAutoScaling) {
		var ret ApiAtlasClusterV15AutoScalingGCPManual
		return ret
	}
	return *o.AnalyticsAutoScaling
}

// GetAnalyticsAutoScalingOk returns a tuple with the AnalyticsAutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetAnalyticsAutoScalingOk() (*ApiAtlasClusterV15AutoScalingGCPManual, bool) {
	if o == nil || isNil(o.AnalyticsAutoScaling) {
    return nil, false
	}
	return o.AnalyticsAutoScaling, true
}

// HasAnalyticsAutoScaling returns a boolean if a field has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) HasAnalyticsAutoScaling() bool {
	if o != nil && !isNil(o.AnalyticsAutoScaling) {
		return true
	}

	return false
}

// SetAnalyticsAutoScaling gets a reference to the given ApiAtlasClusterV15AutoScalingGCPManual and assigns it to the AnalyticsAutoScaling field.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) SetAnalyticsAutoScaling(v ApiAtlasClusterV15AutoScalingGCPManual) {
	o.AnalyticsAutoScaling = &v
}

// GetAnalyticsSpecs returns the AnalyticsSpecs field value if set, zero value otherwise.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetAnalyticsSpecs() ApiAtlasClusterHardwareSpecsV15GCPManual {
	if o == nil || isNil(o.AnalyticsSpecs) {
		var ret ApiAtlasClusterHardwareSpecsV15GCPManual
		return ret
	}
	return *o.AnalyticsSpecs
}

// GetAnalyticsSpecsOk returns a tuple with the AnalyticsSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetAnalyticsSpecsOk() (*ApiAtlasClusterHardwareSpecsV15GCPManual, bool) {
	if o == nil || isNil(o.AnalyticsSpecs) {
    return nil, false
	}
	return o.AnalyticsSpecs, true
}

// HasAnalyticsSpecs returns a boolean if a field has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) HasAnalyticsSpecs() bool {
	if o != nil && !isNil(o.AnalyticsSpecs) {
		return true
	}

	return false
}

// SetAnalyticsSpecs gets a reference to the given ApiAtlasClusterHardwareSpecsV15GCPManual and assigns it to the AnalyticsSpecs field.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) SetAnalyticsSpecs(v ApiAtlasClusterHardwareSpecsV15GCPManual) {
	o.AnalyticsSpecs = &v
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetAutoScaling() ApiAtlasClusterV15AutoScalingGCPManual {
	if o == nil || isNil(o.AutoScaling) {
		var ret ApiAtlasClusterV15AutoScalingGCPManual
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetAutoScalingOk() (*ApiAtlasClusterV15AutoScalingGCPManual, bool) {
	if o == nil || isNil(o.AutoScaling) {
    return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) HasAutoScaling() bool {
	if o != nil && !isNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given ApiAtlasClusterV15AutoScalingGCPManual and assigns it to the AutoScaling field.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) SetAutoScaling(v ApiAtlasClusterV15AutoScalingGCPManual) {
	o.AutoScaling = &v
}

// GetElectableSpecs returns the ElectableSpecs field value if set, zero value otherwise.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetElectableSpecs() ApiAtlasClusterHardwareSpecsV15GCPManual {
	if o == nil || isNil(o.ElectableSpecs) {
		var ret ApiAtlasClusterHardwareSpecsV15GCPManual
		return ret
	}
	return *o.ElectableSpecs
}

// GetElectableSpecsOk returns a tuple with the ElectableSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetElectableSpecsOk() (*ApiAtlasClusterHardwareSpecsV15GCPManual, bool) {
	if o == nil || isNil(o.ElectableSpecs) {
    return nil, false
	}
	return o.ElectableSpecs, true
}

// HasElectableSpecs returns a boolean if a field has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) HasElectableSpecs() bool {
	if o != nil && !isNil(o.ElectableSpecs) {
		return true
	}

	return false
}

// SetElectableSpecs gets a reference to the given ApiAtlasClusterHardwareSpecsV15GCPManual and assigns it to the ElectableSpecs field.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) SetElectableSpecs(v ApiAtlasClusterHardwareSpecsV15GCPManual) {
	o.ElectableSpecs = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetPriority() int32 {
	if o == nil || isNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) SetPriority(v int32) {
	o.Priority = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetProviderName() string {
	if o == nil || isNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetProviderNameOk() (*string, bool) {
	if o == nil || isNil(o.ProviderName) {
    return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) HasProviderName() bool {
	if o != nil && !isNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetReadOnlySpecs returns the ReadOnlySpecs field value if set, zero value otherwise.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetReadOnlySpecs() ApiAtlasClusterHardwareSpecsV15GCPManual {
	if o == nil || isNil(o.ReadOnlySpecs) {
		var ret ApiAtlasClusterHardwareSpecsV15GCPManual
		return ret
	}
	return *o.ReadOnlySpecs
}

// GetReadOnlySpecsOk returns a tuple with the ReadOnlySpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetReadOnlySpecsOk() (*ApiAtlasClusterHardwareSpecsV15GCPManual, bool) {
	if o == nil || isNil(o.ReadOnlySpecs) {
    return nil, false
	}
	return o.ReadOnlySpecs, true
}

// HasReadOnlySpecs returns a boolean if a field has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) HasReadOnlySpecs() bool {
	if o != nil && !isNil(o.ReadOnlySpecs) {
		return true
	}

	return false
}

// SetReadOnlySpecs gets a reference to the given ApiAtlasClusterHardwareSpecsV15GCPManual and assigns it to the ReadOnlySpecs field.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) SetReadOnlySpecs(v ApiAtlasClusterHardwareSpecsV15GCPManual) {
	o.ReadOnlySpecs = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetRegionName() string {
	if o == nil || isNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) GetRegionNameOk() (*string, bool) {
	if o == nil || isNil(o.RegionName) {
    return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) HasRegionName() bool {
	if o != nil && !isNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ApiAtlasClusterRegionConfigsV15GCPManual) SetRegionName(v string) {
	o.RegionName = &v
}

func (o ApiAtlasClusterRegionConfigsV15GCPManual) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AnalyticsAutoScaling) {
		toSerialize["analyticsAutoScaling"] = o.AnalyticsAutoScaling
	}
	if !isNil(o.AnalyticsSpecs) {
		toSerialize["analyticsSpecs"] = o.AnalyticsSpecs
	}
	if !isNil(o.AutoScaling) {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if !isNil(o.ElectableSpecs) {
		toSerialize["electableSpecs"] = o.ElectableSpecs
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	if !isNil(o.ReadOnlySpecs) {
		toSerialize["readOnlySpecs"] = o.ReadOnlySpecs
	}
	if !isNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasClusterRegionConfigsV15GCPManual struct {
	value *ApiAtlasClusterRegionConfigsV15GCPManual
	isSet bool
}

func (v NullableApiAtlasClusterRegionConfigsV15GCPManual) Get() *ApiAtlasClusterRegionConfigsV15GCPManual {
	return v.value
}

func (v *NullableApiAtlasClusterRegionConfigsV15GCPManual) Set(val *ApiAtlasClusterRegionConfigsV15GCPManual) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasClusterRegionConfigsV15GCPManual) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasClusterRegionConfigsV15GCPManual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasClusterRegionConfigsV15GCPManual(val *ApiAtlasClusterRegionConfigsV15GCPManual) *NullableApiAtlasClusterRegionConfigsV15GCPManual {
	return &NullableApiAtlasClusterRegionConfigsV15GCPManual{value: val, isSet: true}
}

func (v NullableApiAtlasClusterRegionConfigsV15GCPManual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasClusterRegionConfigsV15GCPManual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


