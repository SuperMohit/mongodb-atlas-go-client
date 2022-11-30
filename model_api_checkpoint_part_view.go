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

// ApiCheckpointPartView Metadata contained in one document that describes the complete snapshot taken for this node.
type ApiCheckpointPartView struct {
	// Human-readable label that identifies the replica set to which this checkpoint applies.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// Human-readable label that identifies the shard to which this checkpoint applies.
	ShardName *string `json:"shardName,omitempty"`
	// Flag that indicates whether the token exists.
	TokenDiscovered *bool `json:"tokenDiscovered,omitempty"`
	TokenTimestamp *ApiBSONTimestampView `json:"tokenTimestamp,omitempty"`
	// Human-readable label that identifies the type of host that the part represents.
	TypeName *string `json:"typeName,omitempty"`
}

// NewApiCheckpointPartView instantiates a new ApiCheckpointPartView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCheckpointPartView() *ApiCheckpointPartView {
	this := ApiCheckpointPartView{}
	return &this
}

// NewApiCheckpointPartViewWithDefaults instantiates a new ApiCheckpointPartView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCheckpointPartViewWithDefaults() *ApiCheckpointPartView {
	this := ApiCheckpointPartView{}
	return &this
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise.
func (o *ApiCheckpointPartView) GetReplicaSetName() string {
	if o == nil || isNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCheckpointPartView) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || isNil(o.ReplicaSetName) {
    return nil, false
	}
	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *ApiCheckpointPartView) HasReplicaSetName() bool {
	if o != nil && !isNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *ApiCheckpointPartView) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

// GetShardName returns the ShardName field value if set, zero value otherwise.
func (o *ApiCheckpointPartView) GetShardName() string {
	if o == nil || isNil(o.ShardName) {
		var ret string
		return ret
	}
	return *o.ShardName
}

// GetShardNameOk returns a tuple with the ShardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCheckpointPartView) GetShardNameOk() (*string, bool) {
	if o == nil || isNil(o.ShardName) {
    return nil, false
	}
	return o.ShardName, true
}

// HasShardName returns a boolean if a field has been set.
func (o *ApiCheckpointPartView) HasShardName() bool {
	if o != nil && !isNil(o.ShardName) {
		return true
	}

	return false
}

// SetShardName gets a reference to the given string and assigns it to the ShardName field.
func (o *ApiCheckpointPartView) SetShardName(v string) {
	o.ShardName = &v
}

// GetTokenDiscovered returns the TokenDiscovered field value if set, zero value otherwise.
func (o *ApiCheckpointPartView) GetTokenDiscovered() bool {
	if o == nil || isNil(o.TokenDiscovered) {
		var ret bool
		return ret
	}
	return *o.TokenDiscovered
}

// GetTokenDiscoveredOk returns a tuple with the TokenDiscovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCheckpointPartView) GetTokenDiscoveredOk() (*bool, bool) {
	if o == nil || isNil(o.TokenDiscovered) {
    return nil, false
	}
	return o.TokenDiscovered, true
}

// HasTokenDiscovered returns a boolean if a field has been set.
func (o *ApiCheckpointPartView) HasTokenDiscovered() bool {
	if o != nil && !isNil(o.TokenDiscovered) {
		return true
	}

	return false
}

// SetTokenDiscovered gets a reference to the given bool and assigns it to the TokenDiscovered field.
func (o *ApiCheckpointPartView) SetTokenDiscovered(v bool) {
	o.TokenDiscovered = &v
}

// GetTokenTimestamp returns the TokenTimestamp field value if set, zero value otherwise.
func (o *ApiCheckpointPartView) GetTokenTimestamp() ApiBSONTimestampView {
	if o == nil || isNil(o.TokenTimestamp) {
		var ret ApiBSONTimestampView
		return ret
	}
	return *o.TokenTimestamp
}

// GetTokenTimestampOk returns a tuple with the TokenTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCheckpointPartView) GetTokenTimestampOk() (*ApiBSONTimestampView, bool) {
	if o == nil || isNil(o.TokenTimestamp) {
    return nil, false
	}
	return o.TokenTimestamp, true
}

// HasTokenTimestamp returns a boolean if a field has been set.
func (o *ApiCheckpointPartView) HasTokenTimestamp() bool {
	if o != nil && !isNil(o.TokenTimestamp) {
		return true
	}

	return false
}

// SetTokenTimestamp gets a reference to the given ApiBSONTimestampView and assigns it to the TokenTimestamp field.
func (o *ApiCheckpointPartView) SetTokenTimestamp(v ApiBSONTimestampView) {
	o.TokenTimestamp = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *ApiCheckpointPartView) GetTypeName() string {
	if o == nil || isNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCheckpointPartView) GetTypeNameOk() (*string, bool) {
	if o == nil || isNil(o.TypeName) {
    return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *ApiCheckpointPartView) HasTypeName() bool {
	if o != nil && !isNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *ApiCheckpointPartView) SetTypeName(v string) {
	o.TypeName = &v
}

func (o ApiCheckpointPartView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReplicaSetName) {
		toSerialize["replicaSetName"] = o.ReplicaSetName
	}
	if !isNil(o.ShardName) {
		toSerialize["shardName"] = o.ShardName
	}
	if !isNil(o.TokenDiscovered) {
		toSerialize["tokenDiscovered"] = o.TokenDiscovered
	}
	if !isNil(o.TokenTimestamp) {
		toSerialize["tokenTimestamp"] = o.TokenTimestamp
	}
	if !isNil(o.TypeName) {
		toSerialize["typeName"] = o.TypeName
	}
	return json.Marshal(toSerialize)
}

type NullableApiCheckpointPartView struct {
	value *ApiCheckpointPartView
	isSet bool
}

func (v NullableApiCheckpointPartView) Get() *ApiCheckpointPartView {
	return v.value
}

func (v *NullableApiCheckpointPartView) Set(val *ApiCheckpointPartView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCheckpointPartView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCheckpointPartView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCheckpointPartView(val *ApiCheckpointPartView) *NullableApiCheckpointPartView {
	return &NullableApiCheckpointPartView{value: val, isSet: true}
}

func (v NullableApiCheckpointPartView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCheckpointPartView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


