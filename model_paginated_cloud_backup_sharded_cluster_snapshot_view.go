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

// PaginatedCloudBackupShardedClusterSnapshotView struct for PaginatedCloudBackupShardedClusterSnapshotView
type PaginatedCloudBackupShardedClusterSnapshotView struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []ApiAtlasDiskBackupShardedClusterSnapshotView `json:"results"`
	// Number of documents returned in this response.
	TotalCount int32 `json:"totalCount"`
}

// NewPaginatedCloudBackupShardedClusterSnapshotView instantiates a new PaginatedCloudBackupShardedClusterSnapshotView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCloudBackupShardedClusterSnapshotView(links []Link, results []ApiAtlasDiskBackupShardedClusterSnapshotView, totalCount int32) *PaginatedCloudBackupShardedClusterSnapshotView {
	this := PaginatedCloudBackupShardedClusterSnapshotView{}
	this.Links = links
	this.Results = results
	this.TotalCount = totalCount
	return &this
}

// NewPaginatedCloudBackupShardedClusterSnapshotViewWithDefaults instantiates a new PaginatedCloudBackupShardedClusterSnapshotView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCloudBackupShardedClusterSnapshotViewWithDefaults() *PaginatedCloudBackupShardedClusterSnapshotView {
	this := PaginatedCloudBackupShardedClusterSnapshotView{}
	return &this
}

// GetLinks returns the Links field value
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *PaginatedCloudBackupShardedClusterSnapshotView) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetResults() []ApiAtlasDiskBackupShardedClusterSnapshotView {
	if o == nil {
		var ret []ApiAtlasDiskBackupShardedClusterSnapshotView
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetResultsOk() ([]ApiAtlasDiskBackupShardedClusterSnapshotView, bool) {
	if o == nil {
    return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedCloudBackupShardedClusterSnapshotView) SetResults(v []ApiAtlasDiskBackupShardedClusterSnapshotView) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PaginatedCloudBackupShardedClusterSnapshotView) GetTotalCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PaginatedCloudBackupShardedClusterSnapshotView) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o PaginatedCloudBackupShardedClusterSnapshotView) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedCloudBackupShardedClusterSnapshotView struct {
	value *PaginatedCloudBackupShardedClusterSnapshotView
	isSet bool
}

func (v NullablePaginatedCloudBackupShardedClusterSnapshotView) Get() *PaginatedCloudBackupShardedClusterSnapshotView {
	return v.value
}

func (v *NullablePaginatedCloudBackupShardedClusterSnapshotView) Set(val *PaginatedCloudBackupShardedClusterSnapshotView) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCloudBackupShardedClusterSnapshotView) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCloudBackupShardedClusterSnapshotView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCloudBackupShardedClusterSnapshotView(val *PaginatedCloudBackupShardedClusterSnapshotView) *NullablePaginatedCloudBackupShardedClusterSnapshotView {
	return &NullablePaginatedCloudBackupShardedClusterSnapshotView{value: val, isSet: true}
}

func (v NullablePaginatedCloudBackupShardedClusterSnapshotView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCloudBackupShardedClusterSnapshotView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


