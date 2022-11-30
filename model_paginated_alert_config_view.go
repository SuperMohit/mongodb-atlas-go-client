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

// PaginatedAlertConfigView struct for PaginatedAlertConfigView
type PaginatedAlertConfigView struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	Results []ApiAlertConfigView `json:"results"`
	// Number of documents returned in this response.
	TotalCount int32 `json:"totalCount"`
}

// NewPaginatedAlertConfigView instantiates a new PaginatedAlertConfigView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAlertConfigView(links []Link, results []ApiAlertConfigView, totalCount int32) *PaginatedAlertConfigView {
	this := PaginatedAlertConfigView{}
	this.Links = links
	this.Results = results
	this.TotalCount = totalCount
	return &this
}

// NewPaginatedAlertConfigViewWithDefaults instantiates a new PaginatedAlertConfigView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAlertConfigViewWithDefaults() *PaginatedAlertConfigView {
	this := PaginatedAlertConfigView{}
	return &this
}

// GetLinks returns the Links field value
func (o *PaginatedAlertConfigView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PaginatedAlertConfigView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *PaginatedAlertConfigView) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *PaginatedAlertConfigView) GetResults() []ApiAlertConfigView {
	if o == nil {
		var ret []ApiAlertConfigView
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAlertConfigView) GetResultsOk() ([]ApiAlertConfigView, bool) {
	if o == nil {
    return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAlertConfigView) SetResults(v []ApiAlertConfigView) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value
func (o *PaginatedAlertConfigView) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PaginatedAlertConfigView) GetTotalCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PaginatedAlertConfigView) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o PaginatedAlertConfigView) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedAlertConfigView struct {
	value *PaginatedAlertConfigView
	isSet bool
}

func (v NullablePaginatedAlertConfigView) Get() *PaginatedAlertConfigView {
	return v.value
}

func (v *NullablePaginatedAlertConfigView) Set(val *PaginatedAlertConfigView) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAlertConfigView) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAlertConfigView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAlertConfigView(val *PaginatedAlertConfigView) *NullablePaginatedAlertConfigView {
	return &NullablePaginatedAlertConfigView{value: val, isSet: true}
}

func (v NullablePaginatedAlertConfigView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAlertConfigView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


