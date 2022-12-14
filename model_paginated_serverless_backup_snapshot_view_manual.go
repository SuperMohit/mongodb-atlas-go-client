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

// PaginatedServerlessBackupSnapshotViewManual struct for PaginatedServerlessBackupSnapshotViewManual
type PaginatedServerlessBackupSnapshotViewManual struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.\"
	Links []Link `json:"links"`
	// List of returned documents that MongoDB Cloud provides when completing this request.
	Results []ApiAtlasServerlessBackupSnapshotViewManual `json:"results"`
	// Number of documents returned in this response.
	TotalCount float32 `json:"totalCount"`
}

// NewPaginatedServerlessBackupSnapshotViewManual instantiates a new PaginatedServerlessBackupSnapshotViewManual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedServerlessBackupSnapshotViewManual(links []Link, results []ApiAtlasServerlessBackupSnapshotViewManual, totalCount float32) *PaginatedServerlessBackupSnapshotViewManual {
	this := PaginatedServerlessBackupSnapshotViewManual{}
	this.Links = links
	this.Results = results
	this.TotalCount = totalCount
	return &this
}

// NewPaginatedServerlessBackupSnapshotViewManualWithDefaults instantiates a new PaginatedServerlessBackupSnapshotViewManual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedServerlessBackupSnapshotViewManualWithDefaults() *PaginatedServerlessBackupSnapshotViewManual {
	this := PaginatedServerlessBackupSnapshotViewManual{}
	return &this
}

// GetLinks returns the Links field value
func (o *PaginatedServerlessBackupSnapshotViewManual) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PaginatedServerlessBackupSnapshotViewManual) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *PaginatedServerlessBackupSnapshotViewManual) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *PaginatedServerlessBackupSnapshotViewManual) GetResults() []ApiAtlasServerlessBackupSnapshotViewManual {
	if o == nil {
		var ret []ApiAtlasServerlessBackupSnapshotViewManual
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedServerlessBackupSnapshotViewManual) GetResultsOk() ([]ApiAtlasServerlessBackupSnapshotViewManual, bool) {
	if o == nil {
    return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedServerlessBackupSnapshotViewManual) SetResults(v []ApiAtlasServerlessBackupSnapshotViewManual) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value
func (o *PaginatedServerlessBackupSnapshotViewManual) GetTotalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PaginatedServerlessBackupSnapshotViewManual) GetTotalCountOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PaginatedServerlessBackupSnapshotViewManual) SetTotalCount(v float32) {
	o.TotalCount = v
}

func (o PaginatedServerlessBackupSnapshotViewManual) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["links"] = o.Links
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if true {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedServerlessBackupSnapshotViewManual struct {
	value *PaginatedServerlessBackupSnapshotViewManual
	isSet bool
}

func (v NullablePaginatedServerlessBackupSnapshotViewManual) Get() *PaginatedServerlessBackupSnapshotViewManual {
	return v.value
}

func (v *NullablePaginatedServerlessBackupSnapshotViewManual) Set(val *PaginatedServerlessBackupSnapshotViewManual) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedServerlessBackupSnapshotViewManual) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedServerlessBackupSnapshotViewManual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedServerlessBackupSnapshotViewManual(val *PaginatedServerlessBackupSnapshotViewManual) *NullablePaginatedServerlessBackupSnapshotViewManual {
	return &NullablePaginatedServerlessBackupSnapshotViewManual{value: val, isSet: true}
}

func (v NullablePaginatedServerlessBackupSnapshotViewManual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedServerlessBackupSnapshotViewManual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


