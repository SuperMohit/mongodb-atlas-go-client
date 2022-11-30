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

// ApiMongoDBAccessLogsListView struct for ApiMongoDBAccessLogsListView
type ApiMongoDBAccessLogsListView struct {
	// Authentication attempt, one per object, made against the cluster.
	AccessLogs []ApiMongoDBAccessLogsView `json:"accessLogs,omitempty"`
}

// NewApiMongoDBAccessLogsListView instantiates a new ApiMongoDBAccessLogsListView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMongoDBAccessLogsListView() *ApiMongoDBAccessLogsListView {
	this := ApiMongoDBAccessLogsListView{}
	return &this
}

// NewApiMongoDBAccessLogsListViewWithDefaults instantiates a new ApiMongoDBAccessLogsListView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMongoDBAccessLogsListViewWithDefaults() *ApiMongoDBAccessLogsListView {
	this := ApiMongoDBAccessLogsListView{}
	return &this
}

// GetAccessLogs returns the AccessLogs field value if set, zero value otherwise.
func (o *ApiMongoDBAccessLogsListView) GetAccessLogs() []ApiMongoDBAccessLogsView {
	if o == nil || isNil(o.AccessLogs) {
		var ret []ApiMongoDBAccessLogsView
		return ret
	}
	return o.AccessLogs
}

// GetAccessLogsOk returns a tuple with the AccessLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMongoDBAccessLogsListView) GetAccessLogsOk() ([]ApiMongoDBAccessLogsView, bool) {
	if o == nil || isNil(o.AccessLogs) {
    return nil, false
	}
	return o.AccessLogs, true
}

// HasAccessLogs returns a boolean if a field has been set.
func (o *ApiMongoDBAccessLogsListView) HasAccessLogs() bool {
	if o != nil && !isNil(o.AccessLogs) {
		return true
	}

	return false
}

// SetAccessLogs gets a reference to the given []ApiMongoDBAccessLogsView and assigns it to the AccessLogs field.
func (o *ApiMongoDBAccessLogsListView) SetAccessLogs(v []ApiMongoDBAccessLogsView) {
	o.AccessLogs = v
}

func (o ApiMongoDBAccessLogsListView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessLogs) {
		toSerialize["accessLogs"] = o.AccessLogs
	}
	return json.Marshal(toSerialize)
}

type NullableApiMongoDBAccessLogsListView struct {
	value *ApiMongoDBAccessLogsListView
	isSet bool
}

func (v NullableApiMongoDBAccessLogsListView) Get() *ApiMongoDBAccessLogsListView {
	return v.value
}

func (v *NullableApiMongoDBAccessLogsListView) Set(val *ApiMongoDBAccessLogsListView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMongoDBAccessLogsListView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMongoDBAccessLogsListView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMongoDBAccessLogsListView(val *ApiMongoDBAccessLogsListView) *NullableApiMongoDBAccessLogsListView {
	return &NullableApiMongoDBAccessLogsListView{value: val, isSet: true}
}

func (v NullableApiMongoDBAccessLogsListView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMongoDBAccessLogsListView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

