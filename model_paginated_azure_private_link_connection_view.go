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

// PaginatedAzurePrivateLinkConnectionView Group of Private Endpoint Service settings.
type PaginatedAzurePrivateLinkConnectionView struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []ApiAtlasAzurePrivateLinkConnectionView `json:"results"`
	// Number of documents returned in this response.
	TotalCount int32 `json:"totalCount"`
}

// NewPaginatedAzurePrivateLinkConnectionView instantiates a new PaginatedAzurePrivateLinkConnectionView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAzurePrivateLinkConnectionView(links []Link, results []ApiAtlasAzurePrivateLinkConnectionView, totalCount int32) *PaginatedAzurePrivateLinkConnectionView {
	this := PaginatedAzurePrivateLinkConnectionView{}
	this.Links = links
	this.Results = results
	this.TotalCount = totalCount
	return &this
}

// NewPaginatedAzurePrivateLinkConnectionViewWithDefaults instantiates a new PaginatedAzurePrivateLinkConnectionView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAzurePrivateLinkConnectionViewWithDefaults() *PaginatedAzurePrivateLinkConnectionView {
	this := PaginatedAzurePrivateLinkConnectionView{}
	return &this
}

// GetLinks returns the Links field value
func (o *PaginatedAzurePrivateLinkConnectionView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PaginatedAzurePrivateLinkConnectionView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *PaginatedAzurePrivateLinkConnectionView) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *PaginatedAzurePrivateLinkConnectionView) GetResults() []ApiAtlasAzurePrivateLinkConnectionView {
	if o == nil {
		var ret []ApiAtlasAzurePrivateLinkConnectionView
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAzurePrivateLinkConnectionView) GetResultsOk() ([]ApiAtlasAzurePrivateLinkConnectionView, bool) {
	if o == nil {
    return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAzurePrivateLinkConnectionView) SetResults(v []ApiAtlasAzurePrivateLinkConnectionView) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value
func (o *PaginatedAzurePrivateLinkConnectionView) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PaginatedAzurePrivateLinkConnectionView) GetTotalCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PaginatedAzurePrivateLinkConnectionView) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o PaginatedAzurePrivateLinkConnectionView) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedAzurePrivateLinkConnectionView struct {
	value *PaginatedAzurePrivateLinkConnectionView
	isSet bool
}

func (v NullablePaginatedAzurePrivateLinkConnectionView) Get() *PaginatedAzurePrivateLinkConnectionView {
	return v.value
}

func (v *NullablePaginatedAzurePrivateLinkConnectionView) Set(val *PaginatedAzurePrivateLinkConnectionView) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAzurePrivateLinkConnectionView) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAzurePrivateLinkConnectionView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAzurePrivateLinkConnectionView(val *PaginatedAzurePrivateLinkConnectionView) *NullablePaginatedAzurePrivateLinkConnectionView {
	return &NullablePaginatedAzurePrivateLinkConnectionView{value: val, isSet: true}
}

func (v NullablePaginatedAzurePrivateLinkConnectionView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAzurePrivateLinkConnectionView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


