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

// ApiAtlasTenantSnapshotView struct for ApiAtlasTenantSnapshotView
type ApiAtlasTenantSnapshotView struct {
	// Date and time when the download link no longer works. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Expiration *time.Time `json:"expiration,omitempty"`
	// Date and time when MongoDB Cloud completed writing this snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	FinishTime *time.Time `json:"finishTime,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the restore job.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// MongoDB host version that the snapshot runs.
	MongoDBVersion *string `json:"mongoDBVersion,omitempty"`
	// Date and time when MongoDB Cloud will take the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	ScheduledTime *time.Time `json:"scheduledTime,omitempty"`
	// Date and time when MongoDB Cloud began taking the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	StartTime *time.Time `json:"startTime,omitempty"`
	// Phase of the workflow for this snapshot at the time this resource made this request.
	Status *string `json:"status,omitempty"`
}

// NewApiAtlasTenantSnapshotView instantiates a new ApiAtlasTenantSnapshotView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasTenantSnapshotView(links []Link) *ApiAtlasTenantSnapshotView {
	this := ApiAtlasTenantSnapshotView{}
	this.Links = links
	return &this
}

// NewApiAtlasTenantSnapshotViewWithDefaults instantiates a new ApiAtlasTenantSnapshotView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasTenantSnapshotViewWithDefaults() *ApiAtlasTenantSnapshotView {
	this := ApiAtlasTenantSnapshotView{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ApiAtlasTenantSnapshotView) GetExpiration() time.Time {
	if o == nil || isNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasTenantSnapshotView) GetExpirationOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiration) {
    return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ApiAtlasTenantSnapshotView) HasExpiration() bool {
	if o != nil && !isNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *ApiAtlasTenantSnapshotView) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetFinishTime returns the FinishTime field value if set, zero value otherwise.
func (o *ApiAtlasTenantSnapshotView) GetFinishTime() time.Time {
	if o == nil || isNil(o.FinishTime) {
		var ret time.Time
		return ret
	}
	return *o.FinishTime
}

// GetFinishTimeOk returns a tuple with the FinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasTenantSnapshotView) GetFinishTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.FinishTime) {
    return nil, false
	}
	return o.FinishTime, true
}

// HasFinishTime returns a boolean if a field has been set.
func (o *ApiAtlasTenantSnapshotView) HasFinishTime() bool {
	if o != nil && !isNil(o.FinishTime) {
		return true
	}

	return false
}

// SetFinishTime gets a reference to the given time.Time and assigns it to the FinishTime field.
func (o *ApiAtlasTenantSnapshotView) SetFinishTime(v time.Time) {
	o.FinishTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAtlasTenantSnapshotView) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasTenantSnapshotView) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAtlasTenantSnapshotView) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAtlasTenantSnapshotView) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value
func (o *ApiAtlasTenantSnapshotView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasTenantSnapshotView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ApiAtlasTenantSnapshotView) SetLinks(v []Link) {
	o.Links = v
}

// GetMongoDBVersion returns the MongoDBVersion field value if set, zero value otherwise.
func (o *ApiAtlasTenantSnapshotView) GetMongoDBVersion() string {
	if o == nil || isNil(o.MongoDBVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBVersion
}

// GetMongoDBVersionOk returns a tuple with the MongoDBVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasTenantSnapshotView) GetMongoDBVersionOk() (*string, bool) {
	if o == nil || isNil(o.MongoDBVersion) {
    return nil, false
	}
	return o.MongoDBVersion, true
}

// HasMongoDBVersion returns a boolean if a field has been set.
func (o *ApiAtlasTenantSnapshotView) HasMongoDBVersion() bool {
	if o != nil && !isNil(o.MongoDBVersion) {
		return true
	}

	return false
}

// SetMongoDBVersion gets a reference to the given string and assigns it to the MongoDBVersion field.
func (o *ApiAtlasTenantSnapshotView) SetMongoDBVersion(v string) {
	o.MongoDBVersion = &v
}

// GetScheduledTime returns the ScheduledTime field value if set, zero value otherwise.
func (o *ApiAtlasTenantSnapshotView) GetScheduledTime() time.Time {
	if o == nil || isNil(o.ScheduledTime) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledTime
}

// GetScheduledTimeOk returns a tuple with the ScheduledTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasTenantSnapshotView) GetScheduledTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ScheduledTime) {
    return nil, false
	}
	return o.ScheduledTime, true
}

// HasScheduledTime returns a boolean if a field has been set.
func (o *ApiAtlasTenantSnapshotView) HasScheduledTime() bool {
	if o != nil && !isNil(o.ScheduledTime) {
		return true
	}

	return false
}

// SetScheduledTime gets a reference to the given time.Time and assigns it to the ScheduledTime field.
func (o *ApiAtlasTenantSnapshotView) SetScheduledTime(v time.Time) {
	o.ScheduledTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApiAtlasTenantSnapshotView) GetStartTime() time.Time {
	if o == nil || isNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasTenantSnapshotView) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTime) {
    return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApiAtlasTenantSnapshotView) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApiAtlasTenantSnapshotView) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiAtlasTenantSnapshotView) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasTenantSnapshotView) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiAtlasTenantSnapshotView) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiAtlasTenantSnapshotView) SetStatus(v string) {
	o.Status = &v
}

func (o ApiAtlasTenantSnapshotView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !isNil(o.FinishTime) {
		toSerialize["finishTime"] = o.FinishTime
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.MongoDBVersion) {
		toSerialize["mongoDBVersion"] = o.MongoDBVersion
	}
	if !isNil(o.ScheduledTime) {
		toSerialize["scheduledTime"] = o.ScheduledTime
	}
	if !isNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasTenantSnapshotView struct {
	value *ApiAtlasTenantSnapshotView
	isSet bool
}

func (v NullableApiAtlasTenantSnapshotView) Get() *ApiAtlasTenantSnapshotView {
	return v.value
}

func (v *NullableApiAtlasTenantSnapshotView) Set(val *ApiAtlasTenantSnapshotView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasTenantSnapshotView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasTenantSnapshotView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasTenantSnapshotView(val *ApiAtlasTenantSnapshotView) *NullableApiAtlasTenantSnapshotView {
	return &NullableApiAtlasTenantSnapshotView{value: val, isSet: true}
}

func (v NullableApiAtlasTenantSnapshotView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasTenantSnapshotView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

