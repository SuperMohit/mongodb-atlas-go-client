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

// ApiAtlasManagedNamespacesView List that contains a namespace for a Global Cluster. MongoDB Cloud manages this cluster.
type ApiAtlasManagedNamespacesView struct {
	// Human-readable label of the collection to manage for this Global Cluster.
	Collection string `json:"collection"`
	// Database parameter used to divide the *collection* into shards. Global clusters require a compound shard key. This compound shard key combines the location parameter and the user-selected custom key.
	CustomShardKey string `json:"customShardKey"`
	// Human-readable label of the database to manage for this Global Cluster.
	Db string `json:"db"`
	// Flag that indicates whether someone hashed the custom shard key for the specified collection. If you set this value to `false`, MongoDB Cloud uses ranged sharding.
	IsCustomShardKeyHashed *bool `json:"isCustomShardKeyHashed,omitempty"`
	// Flag that indicates whether someone [hashed](https://www.mongodb.com/docs/manual/reference/method/sh.shardCollection/#hashed-shard-keys) the custom shard key. If this parameter returns `false`, this cluster uses [ranged sharding](https://www.mongodb.com/docs/manual/core/ranged-sharding/).
	IsShardKeyUnique *bool `json:"isShardKeyUnique,omitempty"`
	// Minimum number of chunks to create initially when sharding an empty collection with a [hashed shard key](https://www.mongodb.com/docs/manual/core/hashed-sharding/).
	NumInitialChunks *int64 `json:"numInitialChunks,omitempty"`
	// Flag that indicates whether MongoDB Cloud should create and distribute initial chunks for an empty or non-existing collection. MongoDB Cloud distributes data based on the defined zones and zone ranges for the collection.
	PresplitHashedZones *bool `json:"presplitHashedZones,omitempty"`
}

// NewApiAtlasManagedNamespacesView instantiates a new ApiAtlasManagedNamespacesView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasManagedNamespacesView(collection string, customShardKey string, db string) *ApiAtlasManagedNamespacesView {
	this := ApiAtlasManagedNamespacesView{}
	this.Collection = collection
	this.CustomShardKey = customShardKey
	this.Db = db
	var isCustomShardKeyHashed bool = false
	this.IsCustomShardKeyHashed = &isCustomShardKeyHashed
	var isShardKeyUnique bool = false
	this.IsShardKeyUnique = &isShardKeyUnique
	var presplitHashedZones bool = false
	this.PresplitHashedZones = &presplitHashedZones
	return &this
}

// NewApiAtlasManagedNamespacesViewWithDefaults instantiates a new ApiAtlasManagedNamespacesView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasManagedNamespacesViewWithDefaults() *ApiAtlasManagedNamespacesView {
	this := ApiAtlasManagedNamespacesView{}
	var isCustomShardKeyHashed bool = false
	this.IsCustomShardKeyHashed = &isCustomShardKeyHashed
	var isShardKeyUnique bool = false
	this.IsShardKeyUnique = &isShardKeyUnique
	var presplitHashedZones bool = false
	this.PresplitHashedZones = &presplitHashedZones
	return &this
}

// GetCollection returns the Collection field value
func (o *ApiAtlasManagedNamespacesView) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasManagedNamespacesView) GetCollectionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *ApiAtlasManagedNamespacesView) SetCollection(v string) {
	o.Collection = v
}

// GetCustomShardKey returns the CustomShardKey field value
func (o *ApiAtlasManagedNamespacesView) GetCustomShardKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomShardKey
}

// GetCustomShardKeyOk returns a tuple with the CustomShardKey field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasManagedNamespacesView) GetCustomShardKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CustomShardKey, true
}

// SetCustomShardKey sets field value
func (o *ApiAtlasManagedNamespacesView) SetCustomShardKey(v string) {
	o.CustomShardKey = v
}

// GetDb returns the Db field value
func (o *ApiAtlasManagedNamespacesView) GetDb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Db
}

// GetDbOk returns a tuple with the Db field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasManagedNamespacesView) GetDbOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Db, true
}

// SetDb sets field value
func (o *ApiAtlasManagedNamespacesView) SetDb(v string) {
	o.Db = v
}

// GetIsCustomShardKeyHashed returns the IsCustomShardKeyHashed field value if set, zero value otherwise.
func (o *ApiAtlasManagedNamespacesView) GetIsCustomShardKeyHashed() bool {
	if o == nil || isNil(o.IsCustomShardKeyHashed) {
		var ret bool
		return ret
	}
	return *o.IsCustomShardKeyHashed
}

// GetIsCustomShardKeyHashedOk returns a tuple with the IsCustomShardKeyHashed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasManagedNamespacesView) GetIsCustomShardKeyHashedOk() (*bool, bool) {
	if o == nil || isNil(o.IsCustomShardKeyHashed) {
    return nil, false
	}
	return o.IsCustomShardKeyHashed, true
}

// HasIsCustomShardKeyHashed returns a boolean if a field has been set.
func (o *ApiAtlasManagedNamespacesView) HasIsCustomShardKeyHashed() bool {
	if o != nil && !isNil(o.IsCustomShardKeyHashed) {
		return true
	}

	return false
}

// SetIsCustomShardKeyHashed gets a reference to the given bool and assigns it to the IsCustomShardKeyHashed field.
func (o *ApiAtlasManagedNamespacesView) SetIsCustomShardKeyHashed(v bool) {
	o.IsCustomShardKeyHashed = &v
}

// GetIsShardKeyUnique returns the IsShardKeyUnique field value if set, zero value otherwise.
func (o *ApiAtlasManagedNamespacesView) GetIsShardKeyUnique() bool {
	if o == nil || isNil(o.IsShardKeyUnique) {
		var ret bool
		return ret
	}
	return *o.IsShardKeyUnique
}

// GetIsShardKeyUniqueOk returns a tuple with the IsShardKeyUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasManagedNamespacesView) GetIsShardKeyUniqueOk() (*bool, bool) {
	if o == nil || isNil(o.IsShardKeyUnique) {
    return nil, false
	}
	return o.IsShardKeyUnique, true
}

// HasIsShardKeyUnique returns a boolean if a field has been set.
func (o *ApiAtlasManagedNamespacesView) HasIsShardKeyUnique() bool {
	if o != nil && !isNil(o.IsShardKeyUnique) {
		return true
	}

	return false
}

// SetIsShardKeyUnique gets a reference to the given bool and assigns it to the IsShardKeyUnique field.
func (o *ApiAtlasManagedNamespacesView) SetIsShardKeyUnique(v bool) {
	o.IsShardKeyUnique = &v
}

// GetNumInitialChunks returns the NumInitialChunks field value if set, zero value otherwise.
func (o *ApiAtlasManagedNamespacesView) GetNumInitialChunks() int64 {
	if o == nil || isNil(o.NumInitialChunks) {
		var ret int64
		return ret
	}
	return *o.NumInitialChunks
}

// GetNumInitialChunksOk returns a tuple with the NumInitialChunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasManagedNamespacesView) GetNumInitialChunksOk() (*int64, bool) {
	if o == nil || isNil(o.NumInitialChunks) {
    return nil, false
	}
	return o.NumInitialChunks, true
}

// HasNumInitialChunks returns a boolean if a field has been set.
func (o *ApiAtlasManagedNamespacesView) HasNumInitialChunks() bool {
	if o != nil && !isNil(o.NumInitialChunks) {
		return true
	}

	return false
}

// SetNumInitialChunks gets a reference to the given int64 and assigns it to the NumInitialChunks field.
func (o *ApiAtlasManagedNamespacesView) SetNumInitialChunks(v int64) {
	o.NumInitialChunks = &v
}

// GetPresplitHashedZones returns the PresplitHashedZones field value if set, zero value otherwise.
func (o *ApiAtlasManagedNamespacesView) GetPresplitHashedZones() bool {
	if o == nil || isNil(o.PresplitHashedZones) {
		var ret bool
		return ret
	}
	return *o.PresplitHashedZones
}

// GetPresplitHashedZonesOk returns a tuple with the PresplitHashedZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasManagedNamespacesView) GetPresplitHashedZonesOk() (*bool, bool) {
	if o == nil || isNil(o.PresplitHashedZones) {
    return nil, false
	}
	return o.PresplitHashedZones, true
}

// HasPresplitHashedZones returns a boolean if a field has been set.
func (o *ApiAtlasManagedNamespacesView) HasPresplitHashedZones() bool {
	if o != nil && !isNil(o.PresplitHashedZones) {
		return true
	}

	return false
}

// SetPresplitHashedZones gets a reference to the given bool and assigns it to the PresplitHashedZones field.
func (o *ApiAtlasManagedNamespacesView) SetPresplitHashedZones(v bool) {
	o.PresplitHashedZones = &v
}

func (o ApiAtlasManagedNamespacesView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["collection"] = o.Collection
	}
	if true {
		toSerialize["customShardKey"] = o.CustomShardKey
	}
	if true {
		toSerialize["db"] = o.Db
	}
	if !isNil(o.IsCustomShardKeyHashed) {
		toSerialize["isCustomShardKeyHashed"] = o.IsCustomShardKeyHashed
	}
	if !isNil(o.IsShardKeyUnique) {
		toSerialize["isShardKeyUnique"] = o.IsShardKeyUnique
	}
	if !isNil(o.NumInitialChunks) {
		toSerialize["numInitialChunks"] = o.NumInitialChunks
	}
	if !isNil(o.PresplitHashedZones) {
		toSerialize["presplitHashedZones"] = o.PresplitHashedZones
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasManagedNamespacesView struct {
	value *ApiAtlasManagedNamespacesView
	isSet bool
}

func (v NullableApiAtlasManagedNamespacesView) Get() *ApiAtlasManagedNamespacesView {
	return v.value
}

func (v *NullableApiAtlasManagedNamespacesView) Set(val *ApiAtlasManagedNamespacesView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasManagedNamespacesView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasManagedNamespacesView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasManagedNamespacesView(val *ApiAtlasManagedNamespacesView) *NullableApiAtlasManagedNamespacesView {
	return &NullableApiAtlasManagedNamespacesView{value: val, isSet: true}
}

func (v NullableApiAtlasManagedNamespacesView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasManagedNamespacesView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


