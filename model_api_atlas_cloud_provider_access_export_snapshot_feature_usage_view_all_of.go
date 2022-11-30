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

// ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf struct for ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf
type ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf struct {
	FeatureId *ApiAtlasCloudProviderAccessFeatureUsageExportSnapshotFeatureIdView `json:"featureId,omitempty"`
}

// NewApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf instantiates a new ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf() *ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf {
	this := ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf{}
	return &this
}

// NewApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOfWithDefaults instantiates a new ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOfWithDefaults() *ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf {
	this := ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) GetFeatureId() ApiAtlasCloudProviderAccessFeatureUsageExportSnapshotFeatureIdView {
	if o == nil || isNil(o.FeatureId) {
		var ret ApiAtlasCloudProviderAccessFeatureUsageExportSnapshotFeatureIdView
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) GetFeatureIdOk() (*ApiAtlasCloudProviderAccessFeatureUsageExportSnapshotFeatureIdView, bool) {
	if o == nil || isNil(o.FeatureId) {
    return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) HasFeatureId() bool {
	if o != nil && !isNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given ApiAtlasCloudProviderAccessFeatureUsageExportSnapshotFeatureIdView and assigns it to the FeatureId field.
func (o *ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) SetFeatureId(v ApiAtlasCloudProviderAccessFeatureUsageExportSnapshotFeatureIdView) {
	o.FeatureId = &v
}

func (o ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FeatureId) {
		toSerialize["featureId"] = o.FeatureId
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf struct {
	value *ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf
	isSet bool
}

func (v NullableApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) Get() *ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf {
	return v.value
}

func (v *NullableApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) Set(val *ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf(val *ApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) *NullableApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf {
	return &NullableApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf{value: val, isSet: true}
}

func (v NullableApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasCloudProviderAccessExportSnapshotFeatureUsageViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

