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

// PaginatedAzurePeerNetworkView Group of Network Peering connection settings.
type PaginatedAzurePeerNetworkView struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []ApiAtlasAzurePeerNetworkView `json:"results"`
	// Number of documents returned in this response.
	TotalCount int32 `json:"totalCount"`
}

// NewPaginatedAzurePeerNetworkView instantiates a new PaginatedAzurePeerNetworkView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAzurePeerNetworkView(links []Link, results []ApiAtlasAzurePeerNetworkView, totalCount int32) *PaginatedAzurePeerNetworkView {
	this := PaginatedAzurePeerNetworkView{}
	this.Links = links
	this.Results = results
	this.TotalCount = totalCount
	return &this
}

// NewPaginatedAzurePeerNetworkViewWithDefaults instantiates a new PaginatedAzurePeerNetworkView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAzurePeerNetworkViewWithDefaults() *PaginatedAzurePeerNetworkView {
	this := PaginatedAzurePeerNetworkView{}
	return &this
}

// GetLinks returns the Links field value
func (o *PaginatedAzurePeerNetworkView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PaginatedAzurePeerNetworkView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *PaginatedAzurePeerNetworkView) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *PaginatedAzurePeerNetworkView) GetResults() []ApiAtlasAzurePeerNetworkView {
	if o == nil {
		var ret []ApiAtlasAzurePeerNetworkView
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAzurePeerNetworkView) GetResultsOk() ([]ApiAtlasAzurePeerNetworkView, bool) {
	if o == nil {
    return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAzurePeerNetworkView) SetResults(v []ApiAtlasAzurePeerNetworkView) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value
func (o *PaginatedAzurePeerNetworkView) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PaginatedAzurePeerNetworkView) GetTotalCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PaginatedAzurePeerNetworkView) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o PaginatedAzurePeerNetworkView) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedAzurePeerNetworkView struct {
	value *PaginatedAzurePeerNetworkView
	isSet bool
}

func (v NullablePaginatedAzurePeerNetworkView) Get() *PaginatedAzurePeerNetworkView {
	return v.value
}

func (v *NullablePaginatedAzurePeerNetworkView) Set(val *PaginatedAzurePeerNetworkView) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAzurePeerNetworkView) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAzurePeerNetworkView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAzurePeerNetworkView(val *PaginatedAzurePeerNetworkView) *NullablePaginatedAzurePeerNetworkView {
	return &NullablePaginatedAzurePeerNetworkView{value: val, isSet: true}
}

func (v NullablePaginatedAzurePeerNetworkView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAzurePeerNetworkView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


