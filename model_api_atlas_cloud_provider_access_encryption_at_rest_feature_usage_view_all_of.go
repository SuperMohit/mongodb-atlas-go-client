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

// ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf struct for ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf
type ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf struct {
	// Object that contains the identifying characteristics of the Amazon Web Services (AWS) Key Management Service (KMS). This field always returns a null value.
	FeatureId map[string]interface{} `json:"featureId,omitempty"`
}

// NewApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf instantiates a new ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf() *ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf {
	this := ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf{}
	return &this
}

// NewApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOfWithDefaults instantiates a new ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOfWithDefaults() *ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf {
	this := ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) GetFeatureId() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) GetFeatureIdOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.FeatureId) {
    return map[string]interface{}{}, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) HasFeatureId() bool {
	if o != nil && isNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given map[string]interface{} and assigns it to the FeatureId field.
func (o *ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) SetFeatureId(v map[string]interface{}) {
	o.FeatureId = v
}

func (o ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf struct {
	value *ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf
	isSet bool
}

func (v NullableApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) Get() *ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf {
	return v.value
}

func (v *NullableApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) Set(val *ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf(val *ApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) *NullableApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf {
	return &NullableApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf{value: val, isSet: true}
}

func (v NullableApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasCloudProviderAccessEncryptionAtRestFeatureUsageViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


