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

// ApiAtlasDBResourceView Namespace to which this database user has access.
type ApiAtlasDBResourceView struct {
	// Flag that indicates whether to grant the action on the cluster resource. If `true`, MongoDB Cloud ignores the **actions.resources.collection** and **actions.resources.db** parameters.
	Cluster bool `json:"cluster"`
	// Human-readable label that identifies the collection on which you grant the action to one MongoDB user. If you don't set this parameter, you grant the action to all collections in the database specified in the **actions.resources.db** parameter. If you set `\"actions.resources.cluster\" : true`, MongoDB Cloud ignores this parameter.
	Collection string `json:"collection"`
	// Human-readable label that identifies the database on which you grant the action to one MongoDB user. If you set `\"actions.resources.cluster\" : true`, MongoDB Cloud ignores this parameter.
	Db string `json:"db"`
}

// NewApiAtlasDBResourceView instantiates a new ApiAtlasDBResourceView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasDBResourceView(cluster bool, collection string, db string) *ApiAtlasDBResourceView {
	this := ApiAtlasDBResourceView{}
	this.Cluster = cluster
	this.Collection = collection
	this.Db = db
	return &this
}

// NewApiAtlasDBResourceViewWithDefaults instantiates a new ApiAtlasDBResourceView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasDBResourceViewWithDefaults() *ApiAtlasDBResourceView {
	this := ApiAtlasDBResourceView{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *ApiAtlasDBResourceView) GetCluster() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDBResourceView) GetClusterOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *ApiAtlasDBResourceView) SetCluster(v bool) {
	o.Cluster = v
}

// GetCollection returns the Collection field value
func (o *ApiAtlasDBResourceView) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDBResourceView) GetCollectionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *ApiAtlasDBResourceView) SetCollection(v string) {
	o.Collection = v
}

// GetDb returns the Db field value
func (o *ApiAtlasDBResourceView) GetDb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Db
}

// GetDbOk returns a tuple with the Db field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDBResourceView) GetDbOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Db, true
}

// SetDb sets field value
func (o *ApiAtlasDBResourceView) SetDb(v string) {
	o.Db = v
}

func (o ApiAtlasDBResourceView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	if true {
		toSerialize["collection"] = o.Collection
	}
	if true {
		toSerialize["db"] = o.Db
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasDBResourceView struct {
	value *ApiAtlasDBResourceView
	isSet bool
}

func (v NullableApiAtlasDBResourceView) Get() *ApiAtlasDBResourceView {
	return v.value
}

func (v *NullableApiAtlasDBResourceView) Set(val *ApiAtlasDBResourceView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasDBResourceView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasDBResourceView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasDBResourceView(val *ApiAtlasDBResourceView) *NullableApiAtlasDBResourceView {
	return &NullableApiAtlasDBResourceView{value: val, isSet: true}
}

func (v NullableApiAtlasDBResourceView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasDBResourceView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

