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

// ApiAtlasDiskBackupOnDemandSnapshotRequestView struct for ApiAtlasDiskBackupOnDemandSnapshotRequestView
type ApiAtlasDiskBackupOnDemandSnapshotRequestView struct {
	// Human-readable phrase or sentence that explains the purpose of the snapshot. The resource returns this parameter when `\"status\" : \"onDemand\"`.
	Description *string `json:"description,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// Number of days that MongoDB Cloud should retain the on-demand snapshot. Must be at least **1**
	RetentionInDays *int32 `json:"retentionInDays,omitempty"`
}

// NewApiAtlasDiskBackupOnDemandSnapshotRequestView instantiates a new ApiAtlasDiskBackupOnDemandSnapshotRequestView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasDiskBackupOnDemandSnapshotRequestView(links []Link) *ApiAtlasDiskBackupOnDemandSnapshotRequestView {
	this := ApiAtlasDiskBackupOnDemandSnapshotRequestView{}
	this.Links = links
	return &this
}

// NewApiAtlasDiskBackupOnDemandSnapshotRequestViewWithDefaults instantiates a new ApiAtlasDiskBackupOnDemandSnapshotRequestView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasDiskBackupOnDemandSnapshotRequestViewWithDefaults() *ApiAtlasDiskBackupOnDemandSnapshotRequestView {
	this := ApiAtlasDiskBackupOnDemandSnapshotRequestView{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) SetDescription(v string) {
	o.Description = &v
}

// GetLinks returns the Links field value
func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) SetLinks(v []Link) {
	o.Links = v
}

// GetRetentionInDays returns the RetentionInDays field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetRetentionInDays() int32 {
	if o == nil || isNil(o.RetentionInDays) {
		var ret int32
		return ret
	}
	return *o.RetentionInDays
}

// GetRetentionInDaysOk returns a tuple with the RetentionInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetRetentionInDaysOk() (*int32, bool) {
	if o == nil || isNil(o.RetentionInDays) {
    return nil, false
	}
	return o.RetentionInDays, true
}

// HasRetentionInDays returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) HasRetentionInDays() bool {
	if o != nil && !isNil(o.RetentionInDays) {
		return true
	}

	return false
}

// SetRetentionInDays gets a reference to the given int32 and assigns it to the RetentionInDays field.
func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) SetRetentionInDays(v int32) {
	o.RetentionInDays = &v
}

func (o ApiAtlasDiskBackupOnDemandSnapshotRequestView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.RetentionInDays) {
		toSerialize["retentionInDays"] = o.RetentionInDays
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasDiskBackupOnDemandSnapshotRequestView struct {
	value *ApiAtlasDiskBackupOnDemandSnapshotRequestView
	isSet bool
}

func (v NullableApiAtlasDiskBackupOnDemandSnapshotRequestView) Get() *ApiAtlasDiskBackupOnDemandSnapshotRequestView {
	return v.value
}

func (v *NullableApiAtlasDiskBackupOnDemandSnapshotRequestView) Set(val *ApiAtlasDiskBackupOnDemandSnapshotRequestView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasDiskBackupOnDemandSnapshotRequestView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasDiskBackupOnDemandSnapshotRequestView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasDiskBackupOnDemandSnapshotRequestView(val *ApiAtlasDiskBackupOnDemandSnapshotRequestView) *NullableApiAtlasDiskBackupOnDemandSnapshotRequestView {
	return &NullableApiAtlasDiskBackupOnDemandSnapshotRequestView{value: val, isSet: true}
}

func (v NullableApiAtlasDiskBackupOnDemandSnapshotRequestView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasDiskBackupOnDemandSnapshotRequestView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

