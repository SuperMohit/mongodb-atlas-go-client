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

// ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf struct for ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf
type ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf struct {
	FeatureId *ApiAtlasCloudProviderAccessFeatureUsageDataLakeFeatureIdView `json:"featureId,omitempty"`
}

// NewApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf instantiates a new ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf() *ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf {
	this := ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf{}
	return &this
}

// NewApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOfWithDefaults instantiates a new ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOfWithDefaults() *ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf {
	this := ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) GetFeatureId() ApiAtlasCloudProviderAccessFeatureUsageDataLakeFeatureIdView {
	if o == nil || isNil(o.FeatureId) {
		var ret ApiAtlasCloudProviderAccessFeatureUsageDataLakeFeatureIdView
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) GetFeatureIdOk() (*ApiAtlasCloudProviderAccessFeatureUsageDataLakeFeatureIdView, bool) {
	if o == nil || isNil(o.FeatureId) {
    return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) HasFeatureId() bool {
	if o != nil && !isNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given ApiAtlasCloudProviderAccessFeatureUsageDataLakeFeatureIdView and assigns it to the FeatureId field.
func (o *ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) SetFeatureId(v ApiAtlasCloudProviderAccessFeatureUsageDataLakeFeatureIdView) {
	o.FeatureId = &v
}

func (o ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FeatureId) {
		toSerialize["featureId"] = o.FeatureId
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf struct {
	value *ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf
	isSet bool
}

func (v NullableApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) Get() *ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf {
	return v.value
}

func (v *NullableApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) Set(val *ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf(val *ApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) *NullableApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf {
	return &NullableApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf{value: val, isSet: true}
}

func (v NullableApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasCloudProviderAccessDataLakeFeatureUsageViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


