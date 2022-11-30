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

// PaginatedAWSPeerVpcView Group of Network Peering connection settings.
type PaginatedAWSPeerVpcView struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []ApiAtlasAWSPeerVpcView `json:"results"`
	// Number of documents returned in this response.
	TotalCount int32 `json:"totalCount"`
}

// NewPaginatedAWSPeerVpcView instantiates a new PaginatedAWSPeerVpcView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAWSPeerVpcView(links []Link, results []ApiAtlasAWSPeerVpcView, totalCount int32) *PaginatedAWSPeerVpcView {
	this := PaginatedAWSPeerVpcView{}
	this.Links = links
	this.Results = results
	this.TotalCount = totalCount
	return &this
}

// NewPaginatedAWSPeerVpcViewWithDefaults instantiates a new PaginatedAWSPeerVpcView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAWSPeerVpcViewWithDefaults() *PaginatedAWSPeerVpcView {
	this := PaginatedAWSPeerVpcView{}
	return &this
}

// GetLinks returns the Links field value
func (o *PaginatedAWSPeerVpcView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PaginatedAWSPeerVpcView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *PaginatedAWSPeerVpcView) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *PaginatedAWSPeerVpcView) GetResults() []ApiAtlasAWSPeerVpcView {
	if o == nil {
		var ret []ApiAtlasAWSPeerVpcView
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAWSPeerVpcView) GetResultsOk() ([]ApiAtlasAWSPeerVpcView, bool) {
	if o == nil {
    return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAWSPeerVpcView) SetResults(v []ApiAtlasAWSPeerVpcView) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value
func (o *PaginatedAWSPeerVpcView) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PaginatedAWSPeerVpcView) GetTotalCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PaginatedAWSPeerVpcView) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o PaginatedAWSPeerVpcView) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedAWSPeerVpcView struct {
	value *PaginatedAWSPeerVpcView
	isSet bool
}

func (v NullablePaginatedAWSPeerVpcView) Get() *PaginatedAWSPeerVpcView {
	return v.value
}

func (v *NullablePaginatedAWSPeerVpcView) Set(val *PaginatedAWSPeerVpcView) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAWSPeerVpcView) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAWSPeerVpcView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAWSPeerVpcView(val *PaginatedAWSPeerVpcView) *NullablePaginatedAWSPeerVpcView {
	return &NullablePaginatedAWSPeerVpcView{value: val, isSet: true}
}

func (v NullablePaginatedAWSPeerVpcView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAWSPeerVpcView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


