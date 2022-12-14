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

// PaginatedIntegrationView Collection of settings that describe third-party integrations.
type PaginatedIntegrationView struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// List of returned documents that MongoDB Cloud provides when completing this request.
	Results []ApiIntegrationViewManual `json:"results"`
	// Number of documents returned in this response.
	TotalCount *float32 `json:"totalCount,omitempty"`
}

// NewPaginatedIntegrationView instantiates a new PaginatedIntegrationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedIntegrationView(links []Link, results []ApiIntegrationViewManual) *PaginatedIntegrationView {
	this := PaginatedIntegrationView{}
	this.Links = links
	this.Results = results
	return &this
}

// NewPaginatedIntegrationViewWithDefaults instantiates a new PaginatedIntegrationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedIntegrationViewWithDefaults() *PaginatedIntegrationView {
	this := PaginatedIntegrationView{}
	return &this
}

// GetLinks returns the Links field value
func (o *PaginatedIntegrationView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PaginatedIntegrationView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *PaginatedIntegrationView) SetLinks(v []Link) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *PaginatedIntegrationView) GetResults() []ApiIntegrationViewManual {
	if o == nil {
		var ret []ApiIntegrationViewManual
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedIntegrationView) GetResultsOk() ([]ApiIntegrationViewManual, bool) {
	if o == nil {
    return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedIntegrationView) SetResults(v []ApiIntegrationViewManual) {
	o.Results = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PaginatedIntegrationView) GetTotalCount() float32 {
	if o == nil || isNil(o.TotalCount) {
		var ret float32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedIntegrationView) GetTotalCountOk() (*float32, bool) {
	if o == nil || isNil(o.TotalCount) {
    return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedIntegrationView) HasTotalCount() bool {
	if o != nil && !isNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given float32 and assigns it to the TotalCount field.
func (o *PaginatedIntegrationView) SetTotalCount(v float32) {
	o.TotalCount = &v
}

func (o PaginatedIntegrationView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["links"] = o.Links
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if !isNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedIntegrationView struct {
	value *PaginatedIntegrationView
	isSet bool
}

func (v NullablePaginatedIntegrationView) Get() *PaginatedIntegrationView {
	return v.value
}

func (v *NullablePaginatedIntegrationView) Set(val *PaginatedIntegrationView) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedIntegrationView) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedIntegrationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedIntegrationView(val *PaginatedIntegrationView) *NullablePaginatedIntegrationView {
	return &NullablePaginatedIntegrationView{value: val, isSet: true}
}

func (v NullablePaginatedIntegrationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedIntegrationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


