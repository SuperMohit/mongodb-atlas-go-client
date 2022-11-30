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

// ApiAtlasOnDemandCpsSnapshotSourceView On-Demand Cloud Provider Snapshots as Source for a Data Lake Pipeline.
type ApiAtlasOnDemandCpsSnapshotSourceView struct {
	// Human-readable name that identifies the cluster.
	ClusterName *string `json:"clusterName,omitempty"`
	// Human-readable name that identifies the collection.
	CollectionName *string `json:"collectionName,omitempty"`
	// Human-readable name that identifies the database.
	DatabaseName *string `json:"databaseName,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	GroupId *string `json:"groupId,omitempty"`
}

// NewApiAtlasOnDemandCpsSnapshotSourceView instantiates a new ApiAtlasOnDemandCpsSnapshotSourceView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasOnDemandCpsSnapshotSourceView() *ApiAtlasOnDemandCpsSnapshotSourceView {
	this := ApiAtlasOnDemandCpsSnapshotSourceView{}
	return &this
}

// NewApiAtlasOnDemandCpsSnapshotSourceViewWithDefaults instantiates a new ApiAtlasOnDemandCpsSnapshotSourceView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasOnDemandCpsSnapshotSourceViewWithDefaults() *ApiAtlasOnDemandCpsSnapshotSourceView {
	this := ApiAtlasOnDemandCpsSnapshotSourceView{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) GetClusterName() string {
	if o == nil || isNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) GetClusterNameOk() (*string, bool) {
	if o == nil || isNil(o.ClusterName) {
    return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) HasClusterName() bool {
	if o != nil && !isNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCollectionName returns the CollectionName field value if set, zero value otherwise.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) GetCollectionName() string {
	if o == nil || isNil(o.CollectionName) {
		var ret string
		return ret
	}
	return *o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) GetCollectionNameOk() (*string, bool) {
	if o == nil || isNil(o.CollectionName) {
    return nil, false
	}
	return o.CollectionName, true
}

// HasCollectionName returns a boolean if a field has been set.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) HasCollectionName() bool {
	if o != nil && !isNil(o.CollectionName) {
		return true
	}

	return false
}

// SetCollectionName gets a reference to the given string and assigns it to the CollectionName field.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) SetCollectionName(v string) {
	o.CollectionName = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) GetDatabaseName() string {
	if o == nil || isNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) GetDatabaseNameOk() (*string, bool) {
	if o == nil || isNil(o.DatabaseName) {
    return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) HasDatabaseName() bool {
	if o != nil && !isNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiAtlasOnDemandCpsSnapshotSourceView) SetGroupId(v string) {
	o.GroupId = &v
}

func (o ApiAtlasOnDemandCpsSnapshotSourceView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClusterName) {
		toSerialize["clusterName"] = o.ClusterName
	}
	if !isNil(o.CollectionName) {
		toSerialize["collectionName"] = o.CollectionName
	}
	if !isNil(o.DatabaseName) {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasOnDemandCpsSnapshotSourceView struct {
	value *ApiAtlasOnDemandCpsSnapshotSourceView
	isSet bool
}

func (v NullableApiAtlasOnDemandCpsSnapshotSourceView) Get() *ApiAtlasOnDemandCpsSnapshotSourceView {
	return v.value
}

func (v *NullableApiAtlasOnDemandCpsSnapshotSourceView) Set(val *ApiAtlasOnDemandCpsSnapshotSourceView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasOnDemandCpsSnapshotSourceView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasOnDemandCpsSnapshotSourceView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasOnDemandCpsSnapshotSourceView(val *ApiAtlasOnDemandCpsSnapshotSourceView) *NullableApiAtlasOnDemandCpsSnapshotSourceView {
	return &NullableApiAtlasOnDemandCpsSnapshotSourceView{value: val, isSet: true}
}

func (v NullableApiAtlasOnDemandCpsSnapshotSourceView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasOnDemandCpsSnapshotSourceView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

