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

// ApiAtlasDLSIngestionSinkView Atlas Data Lake Storage as the destination for a Data Lake Pipeline.
type ApiAtlasDLSIngestionSinkView struct {
	// Target cloud provider for this Data Lake Pipeline.
	MetadataProvider *string `json:"metadataProvider,omitempty"`
	// Target cloud provider region for this Data Lake Pipeline.
	MetadataRegion *string `json:"metadataRegion,omitempty"`
	// Ordered fields used to physically organize data in the destination.
	PartitionFields []ApiAtlasPartitionFieldView `json:"partitionFields,omitempty"`
}

// NewApiAtlasDLSIngestionSinkView instantiates a new ApiAtlasDLSIngestionSinkView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasDLSIngestionSinkView() *ApiAtlasDLSIngestionSinkView {
	this := ApiAtlasDLSIngestionSinkView{}
	return &this
}

// NewApiAtlasDLSIngestionSinkViewWithDefaults instantiates a new ApiAtlasDLSIngestionSinkView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasDLSIngestionSinkViewWithDefaults() *ApiAtlasDLSIngestionSinkView {
	this := ApiAtlasDLSIngestionSinkView{}
	return &this
}

// GetMetadataProvider returns the MetadataProvider field value if set, zero value otherwise.
func (o *ApiAtlasDLSIngestionSinkView) GetMetadataProvider() string {
	if o == nil || isNil(o.MetadataProvider) {
		var ret string
		return ret
	}
	return *o.MetadataProvider
}

// GetMetadataProviderOk returns a tuple with the MetadataProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDLSIngestionSinkView) GetMetadataProviderOk() (*string, bool) {
	if o == nil || isNil(o.MetadataProvider) {
    return nil, false
	}
	return o.MetadataProvider, true
}

// HasMetadataProvider returns a boolean if a field has been set.
func (o *ApiAtlasDLSIngestionSinkView) HasMetadataProvider() bool {
	if o != nil && !isNil(o.MetadataProvider) {
		return true
	}

	return false
}

// SetMetadataProvider gets a reference to the given string and assigns it to the MetadataProvider field.
func (o *ApiAtlasDLSIngestionSinkView) SetMetadataProvider(v string) {
	o.MetadataProvider = &v
}

// GetMetadataRegion returns the MetadataRegion field value if set, zero value otherwise.
func (o *ApiAtlasDLSIngestionSinkView) GetMetadataRegion() string {
	if o == nil || isNil(o.MetadataRegion) {
		var ret string
		return ret
	}
	return *o.MetadataRegion
}

// GetMetadataRegionOk returns a tuple with the MetadataRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDLSIngestionSinkView) GetMetadataRegionOk() (*string, bool) {
	if o == nil || isNil(o.MetadataRegion) {
    return nil, false
	}
	return o.MetadataRegion, true
}

// HasMetadataRegion returns a boolean if a field has been set.
func (o *ApiAtlasDLSIngestionSinkView) HasMetadataRegion() bool {
	if o != nil && !isNil(o.MetadataRegion) {
		return true
	}

	return false
}

// SetMetadataRegion gets a reference to the given string and assigns it to the MetadataRegion field.
func (o *ApiAtlasDLSIngestionSinkView) SetMetadataRegion(v string) {
	o.MetadataRegion = &v
}

// GetPartitionFields returns the PartitionFields field value if set, zero value otherwise.
func (o *ApiAtlasDLSIngestionSinkView) GetPartitionFields() []ApiAtlasPartitionFieldView {
	if o == nil || isNil(o.PartitionFields) {
		var ret []ApiAtlasPartitionFieldView
		return ret
	}
	return o.PartitionFields
}

// GetPartitionFieldsOk returns a tuple with the PartitionFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDLSIngestionSinkView) GetPartitionFieldsOk() ([]ApiAtlasPartitionFieldView, bool) {
	if o == nil || isNil(o.PartitionFields) {
    return nil, false
	}
	return o.PartitionFields, true
}

// HasPartitionFields returns a boolean if a field has been set.
func (o *ApiAtlasDLSIngestionSinkView) HasPartitionFields() bool {
	if o != nil && !isNil(o.PartitionFields) {
		return true
	}

	return false
}

// SetPartitionFields gets a reference to the given []ApiAtlasPartitionFieldView and assigns it to the PartitionFields field.
func (o *ApiAtlasDLSIngestionSinkView) SetPartitionFields(v []ApiAtlasPartitionFieldView) {
	o.PartitionFields = v
}

func (o ApiAtlasDLSIngestionSinkView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MetadataProvider) {
		toSerialize["metadataProvider"] = o.MetadataProvider
	}
	if !isNil(o.MetadataRegion) {
		toSerialize["metadataRegion"] = o.MetadataRegion
	}
	if !isNil(o.PartitionFields) {
		toSerialize["partitionFields"] = o.PartitionFields
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasDLSIngestionSinkView struct {
	value *ApiAtlasDLSIngestionSinkView
	isSet bool
}

func (v NullableApiAtlasDLSIngestionSinkView) Get() *ApiAtlasDLSIngestionSinkView {
	return v.value
}

func (v *NullableApiAtlasDLSIngestionSinkView) Set(val *ApiAtlasDLSIngestionSinkView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasDLSIngestionSinkView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasDLSIngestionSinkView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasDLSIngestionSinkView(val *ApiAtlasDLSIngestionSinkView) *NullableApiAtlasDLSIngestionSinkView {
	return &NullableApiAtlasDLSIngestionSinkView{value: val, isSet: true}
}

func (v NullableApiAtlasDLSIngestionSinkView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasDLSIngestionSinkView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


