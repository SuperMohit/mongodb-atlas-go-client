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

// AutoExportPolicyView Policy for automatically exporting cloud backup snapshots.
type AutoExportPolicyView struct {
	// Unique 24-hexadecimal character string that identifies the AWS Bucket.
	ExportBucketId *string `json:"exportBucketId,omitempty"`
	// Human-readable label that indicates the rate at which the export policy item occurs.
	FrequencyType *string `json:"frequencyType,omitempty"`
}

// NewAutoExportPolicyView instantiates a new AutoExportPolicyView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoExportPolicyView() *AutoExportPolicyView {
	this := AutoExportPolicyView{}
	return &this
}

// NewAutoExportPolicyViewWithDefaults instantiates a new AutoExportPolicyView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoExportPolicyViewWithDefaults() *AutoExportPolicyView {
	this := AutoExportPolicyView{}
	return &this
}

// GetExportBucketId returns the ExportBucketId field value if set, zero value otherwise.
func (o *AutoExportPolicyView) GetExportBucketId() string {
	if o == nil || isNil(o.ExportBucketId) {
		var ret string
		return ret
	}
	return *o.ExportBucketId
}

// GetExportBucketIdOk returns a tuple with the ExportBucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoExportPolicyView) GetExportBucketIdOk() (*string, bool) {
	if o == nil || isNil(o.ExportBucketId) {
    return nil, false
	}
	return o.ExportBucketId, true
}

// HasExportBucketId returns a boolean if a field has been set.
func (o *AutoExportPolicyView) HasExportBucketId() bool {
	if o != nil && !isNil(o.ExportBucketId) {
		return true
	}

	return false
}

// SetExportBucketId gets a reference to the given string and assigns it to the ExportBucketId field.
func (o *AutoExportPolicyView) SetExportBucketId(v string) {
	o.ExportBucketId = &v
}

// GetFrequencyType returns the FrequencyType field value if set, zero value otherwise.
func (o *AutoExportPolicyView) GetFrequencyType() string {
	if o == nil || isNil(o.FrequencyType) {
		var ret string
		return ret
	}
	return *o.FrequencyType
}

// GetFrequencyTypeOk returns a tuple with the FrequencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoExportPolicyView) GetFrequencyTypeOk() (*string, bool) {
	if o == nil || isNil(o.FrequencyType) {
    return nil, false
	}
	return o.FrequencyType, true
}

// HasFrequencyType returns a boolean if a field has been set.
func (o *AutoExportPolicyView) HasFrequencyType() bool {
	if o != nil && !isNil(o.FrequencyType) {
		return true
	}

	return false
}

// SetFrequencyType gets a reference to the given string and assigns it to the FrequencyType field.
func (o *AutoExportPolicyView) SetFrequencyType(v string) {
	o.FrequencyType = &v
}

func (o AutoExportPolicyView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExportBucketId) {
		toSerialize["exportBucketId"] = o.ExportBucketId
	}
	if !isNil(o.FrequencyType) {
		toSerialize["frequencyType"] = o.FrequencyType
	}
	return json.Marshal(toSerialize)
}

type NullableAutoExportPolicyView struct {
	value *AutoExportPolicyView
	isSet bool
}

func (v NullableAutoExportPolicyView) Get() *AutoExportPolicyView {
	return v.value
}

func (v *NullableAutoExportPolicyView) Set(val *AutoExportPolicyView) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoExportPolicyView) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoExportPolicyView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoExportPolicyView(val *AutoExportPolicyView) *NullableAutoExportPolicyView {
	return &NullableAutoExportPolicyView{value: val, isSet: true}
}

func (v NullableAutoExportPolicyView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoExportPolicyView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


