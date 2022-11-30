/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// SourceOrgView struct for SourceOrgView
type SourceOrgView struct {
	// The ISO-8601-formatted timestamp when Cloud Manager or Ops Manager created the link.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Error message pertaining to the organization sync. Returns null if there are no errors.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The ISO-8601-formatted timestamp when Cloud Manager or Ops Manager last synced with Atlas.
	LastSyncedAt *time.Time `json:"lastSyncedAt,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization that contains the data you want to migrate.
	SourceOrgId string `json:"sourceOrgId"`
	TargetOrg *TargetOrgView `json:"targetOrg,omitempty"`
}

// NewSourceOrgView instantiates a new SourceOrgView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceOrgView(sourceOrgId string) *SourceOrgView {
	this := SourceOrgView{}
	this.SourceOrgId = sourceOrgId
	return &this
}

// NewSourceOrgViewWithDefaults instantiates a new SourceOrgView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceOrgViewWithDefaults() *SourceOrgView {
	this := SourceOrgView{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SourceOrgView) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOrgView) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SourceOrgView) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SourceOrgView) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SourceOrgView) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOrgView) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
    return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SourceOrgView) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SourceOrgView) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetLastSyncedAt returns the LastSyncedAt field value if set, zero value otherwise.
func (o *SourceOrgView) GetLastSyncedAt() time.Time {
	if o == nil || isNil(o.LastSyncedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastSyncedAt
}

// GetLastSyncedAtOk returns a tuple with the LastSyncedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOrgView) GetLastSyncedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastSyncedAt) {
    return nil, false
	}
	return o.LastSyncedAt, true
}

// HasLastSyncedAt returns a boolean if a field has been set.
func (o *SourceOrgView) HasLastSyncedAt() bool {
	if o != nil && !isNil(o.LastSyncedAt) {
		return true
	}

	return false
}

// SetLastSyncedAt gets a reference to the given time.Time and assigns it to the LastSyncedAt field.
func (o *SourceOrgView) SetLastSyncedAt(v time.Time) {
	o.LastSyncedAt = &v
}

// GetSourceOrgId returns the SourceOrgId field value
func (o *SourceOrgView) GetSourceOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceOrgId
}

// GetSourceOrgIdOk returns a tuple with the SourceOrgId field value
// and a boolean to check if the value has been set.
func (o *SourceOrgView) GetSourceOrgIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SourceOrgId, true
}

// SetSourceOrgId sets field value
func (o *SourceOrgView) SetSourceOrgId(v string) {
	o.SourceOrgId = v
}

// GetTargetOrg returns the TargetOrg field value if set, zero value otherwise.
func (o *SourceOrgView) GetTargetOrg() TargetOrgView {
	if o == nil || isNil(o.TargetOrg) {
		var ret TargetOrgView
		return ret
	}
	return *o.TargetOrg
}

// GetTargetOrgOk returns a tuple with the TargetOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOrgView) GetTargetOrgOk() (*TargetOrgView, bool) {
	if o == nil || isNil(o.TargetOrg) {
    return nil, false
	}
	return o.TargetOrg, true
}

// HasTargetOrg returns a boolean if a field has been set.
func (o *SourceOrgView) HasTargetOrg() bool {
	if o != nil && !isNil(o.TargetOrg) {
		return true
	}

	return false
}

// SetTargetOrg gets a reference to the given TargetOrgView and assigns it to the TargetOrg field.
func (o *SourceOrgView) SetTargetOrg(v TargetOrgView) {
	o.TargetOrg = &v
}

func (o SourceOrgView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !isNil(o.LastSyncedAt) {
		toSerialize["lastSyncedAt"] = o.LastSyncedAt
	}
	if true {
		toSerialize["sourceOrgId"] = o.SourceOrgId
	}
	if !isNil(o.TargetOrg) {
		toSerialize["targetOrg"] = o.TargetOrg
	}
	return json.Marshal(toSerialize)
}

type NullableSourceOrgView struct {
	value *SourceOrgView
	isSet bool
}

func (v NullableSourceOrgView) Get() *SourceOrgView {
	return v.value
}

func (v *NullableSourceOrgView) Set(val *SourceOrgView) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceOrgView) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceOrgView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceOrgView(val *SourceOrgView) *NullableSourceOrgView {
	return &NullableSourceOrgView{value: val, isSet: true}
}

func (v NullableSourceOrgView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceOrgView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

