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

// ApiAtlasPipelineRunStatsView Runtime statistics for this Data Lake Pipeline run.
type ApiAtlasPipelineRunStatsView struct {
	// Total data size in bytes exported for this pipeline run.
	BytesExported *int64 `json:"bytesExported,omitempty"`
	// Number of docs ingested for a this pipeline run.
	NumDocs *int64 `json:"numDocs,omitempty"`
}

// NewApiAtlasPipelineRunStatsView instantiates a new ApiAtlasPipelineRunStatsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasPipelineRunStatsView() *ApiAtlasPipelineRunStatsView {
	this := ApiAtlasPipelineRunStatsView{}
	return &this
}

// NewApiAtlasPipelineRunStatsViewWithDefaults instantiates a new ApiAtlasPipelineRunStatsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasPipelineRunStatsViewWithDefaults() *ApiAtlasPipelineRunStatsView {
	this := ApiAtlasPipelineRunStatsView{}
	return &this
}

// GetBytesExported returns the BytesExported field value if set, zero value otherwise.
func (o *ApiAtlasPipelineRunStatsView) GetBytesExported() int64 {
	if o == nil || isNil(o.BytesExported) {
		var ret int64
		return ret
	}
	return *o.BytesExported
}

// GetBytesExportedOk returns a tuple with the BytesExported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasPipelineRunStatsView) GetBytesExportedOk() (*int64, bool) {
	if o == nil || isNil(o.BytesExported) {
    return nil, false
	}
	return o.BytesExported, true
}

// HasBytesExported returns a boolean if a field has been set.
func (o *ApiAtlasPipelineRunStatsView) HasBytesExported() bool {
	if o != nil && !isNil(o.BytesExported) {
		return true
	}

	return false
}

// SetBytesExported gets a reference to the given int64 and assigns it to the BytesExported field.
func (o *ApiAtlasPipelineRunStatsView) SetBytesExported(v int64) {
	o.BytesExported = &v
}

// GetNumDocs returns the NumDocs field value if set, zero value otherwise.
func (o *ApiAtlasPipelineRunStatsView) GetNumDocs() int64 {
	if o == nil || isNil(o.NumDocs) {
		var ret int64
		return ret
	}
	return *o.NumDocs
}

// GetNumDocsOk returns a tuple with the NumDocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasPipelineRunStatsView) GetNumDocsOk() (*int64, bool) {
	if o == nil || isNil(o.NumDocs) {
    return nil, false
	}
	return o.NumDocs, true
}

// HasNumDocs returns a boolean if a field has been set.
func (o *ApiAtlasPipelineRunStatsView) HasNumDocs() bool {
	if o != nil && !isNil(o.NumDocs) {
		return true
	}

	return false
}

// SetNumDocs gets a reference to the given int64 and assigns it to the NumDocs field.
func (o *ApiAtlasPipelineRunStatsView) SetNumDocs(v int64) {
	o.NumDocs = &v
}

func (o ApiAtlasPipelineRunStatsView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BytesExported) {
		toSerialize["bytesExported"] = o.BytesExported
	}
	if !isNil(o.NumDocs) {
		toSerialize["numDocs"] = o.NumDocs
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasPipelineRunStatsView struct {
	value *ApiAtlasPipelineRunStatsView
	isSet bool
}

func (v NullableApiAtlasPipelineRunStatsView) Get() *ApiAtlasPipelineRunStatsView {
	return v.value
}

func (v *NullableApiAtlasPipelineRunStatsView) Set(val *ApiAtlasPipelineRunStatsView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasPipelineRunStatsView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasPipelineRunStatsView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasPipelineRunStatsView(val *ApiAtlasPipelineRunStatsView) *NullableApiAtlasPipelineRunStatsView {
	return &NullableApiAtlasPipelineRunStatsView{value: val, isSet: true}
}

func (v NullableApiAtlasPipelineRunStatsView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasPipelineRunStatsView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


