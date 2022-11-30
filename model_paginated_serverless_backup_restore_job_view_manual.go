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

// PaginatedServerlessBackupRestoreJobViewManual struct for PaginatedServerlessBackupRestoreJobViewManual
type PaginatedServerlessBackupRestoreJobViewManual struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// List of returned documents that MongoDB Cloud provides when completing this request.
	Results []ApiAtlasServerlessBackupRestoreJobViewManual `json:"results"`
	// Number of documents returned in this response.
	TotalCount float32 `json:"totalCount"`
}

// NewPaginatedServerlessBackupRestoreJobViewManual instantiates a new PaginatedServerlessBackupRestoreJobViewManual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedServerlessBackupRestoreJobViewManual(links []Link, results []ApiAtlasServerlessBackupRestoreJobViewManual, totalCount float32) *PaginatedServerlessBackupRestoreJobViewManual {
	this := PaginatedServerlessBackupRestoreJobViewManual{}
	this.Links = links
	this.Results = results
	this.TotalCount = totalCount
	return &this
}

// NewPaginatedServerlessBackupRestoreJobViewManualWithDefaults instantiates a new PaginatedServerlessBackupRestoreJobViewManual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedServerlessBackupRestoreJobViewManualWithDefaults() *PaginatedServerlessBackupRestoreJobViewManual {
	this := PaginatedServerlessBackupRestoreJobViewManual{}
	return &this
}

// GetLinks returns the Links field value
func (o *PaginatedServerlessBackupRestoreJobViewManual) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PaginatedServerlessBackupRestoreJobViewManual) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *PaginatedServerlessBackupRestoreJobViewManual) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *PaginatedServerlessBackupRestoreJobViewManual) GetResults() []ApiAtlasServerlessBackupRestoreJobViewManual {
	if o == nil {
		var ret []ApiAtlasServerlessBackupRestoreJobViewManual
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedServerlessBackupRestoreJobViewManual) GetResultsOk() ([]ApiAtlasServerlessBackupRestoreJobViewManual, bool) {
	if o == nil {
    return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedServerlessBackupRestoreJobViewManual) SetResults(v []ApiAtlasServerlessBackupRestoreJobViewManual) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value
func (o *PaginatedServerlessBackupRestoreJobViewManual) GetTotalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PaginatedServerlessBackupRestoreJobViewManual) GetTotalCountOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PaginatedServerlessBackupRestoreJobViewManual) SetTotalCount(v float32) {
	o.TotalCount = v
}

func (o PaginatedServerlessBackupRestoreJobViewManual) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedServerlessBackupRestoreJobViewManual struct {
	value *PaginatedServerlessBackupRestoreJobViewManual
	isSet bool
}

func (v NullablePaginatedServerlessBackupRestoreJobViewManual) Get() *PaginatedServerlessBackupRestoreJobViewManual {
	return v.value
}

func (v *NullablePaginatedServerlessBackupRestoreJobViewManual) Set(val *PaginatedServerlessBackupRestoreJobViewManual) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedServerlessBackupRestoreJobViewManual) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedServerlessBackupRestoreJobViewManual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedServerlessBackupRestoreJobViewManual(val *PaginatedServerlessBackupRestoreJobViewManual) *NullablePaginatedServerlessBackupRestoreJobViewManual {
	return &NullablePaginatedServerlessBackupRestoreJobViewManual{value: val, isSet: true}
}

func (v NullablePaginatedServerlessBackupRestoreJobViewManual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedServerlessBackupRestoreJobViewManual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


