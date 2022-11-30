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

// ApiAtlasServerlessBackupRestoreJobViewManual struct for ApiAtlasServerlessBackupRestoreJobViewManual
type ApiAtlasServerlessBackupRestoreJobViewManual struct {
	// Flag that indicates whether someone canceled this restore job.
	Cancelled *bool `json:"cancelled,omitempty"`
	// Human-readable label that categorizes the restore job to create.
	DeliveryType string `json:"deliveryType"`
	// One or more Uniform Resource Locators (URLs) that point to the compressed snapshot files for manual download. MongoDB Cloud returns this parameter when `\"deliveryType\" : \"download\"`.\"
	DeliveryUrl []string `json:"deliveryUrl,omitempty"`
	DesiredTimestamp *ApiBSONTimestampView `json:"desiredTimestamp,omitempty"`
	// Flag that indicates whether the restore job expired.
	Expired *bool `json:"expired,omitempty"`
	// Date and time when the restore job expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Flag that indicates whether the restore job failed.
	Failed *bool `json:"failed,omitempty"`
	// Date and time when the restore job completed. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// Unique 24-hexadecimal character string that identifies the restore job.
	Id *string `json:"id,omitempty"`
	// Unique 24-hexadecimal character string that identifies the snapshot.
	SnapshotId *string `json:"snapshotId,omitempty"`
	// Human-readable label that identifies the target cluster to which the restore job restores the snapshot. The resource returns this parameter when `\"deliveryType\":` `\"automated\"`.
	TargetClusterName string `json:"targetClusterName"`
	// Unique 24-hexadecimal digit string that identifies the target project for the specified **targetClusterName**.
	TargetGroupId string `json:"targetGroupId"`
	// Date and time when MongoDB Cloud took the snapshot associated with **snapshotId**. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewApiAtlasServerlessBackupRestoreJobViewManual instantiates a new ApiAtlasServerlessBackupRestoreJobViewManual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasServerlessBackupRestoreJobViewManual(deliveryType string, targetClusterName string, targetGroupId string) *ApiAtlasServerlessBackupRestoreJobViewManual {
	this := ApiAtlasServerlessBackupRestoreJobViewManual{}
	this.DeliveryType = deliveryType
	this.TargetClusterName = targetClusterName
	this.TargetGroupId = targetGroupId
	return &this
}

// NewApiAtlasServerlessBackupRestoreJobViewManualWithDefaults instantiates a new ApiAtlasServerlessBackupRestoreJobViewManual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasServerlessBackupRestoreJobViewManualWithDefaults() *ApiAtlasServerlessBackupRestoreJobViewManual {
	this := ApiAtlasServerlessBackupRestoreJobViewManual{}
	return &this
}

// GetCancelled returns the Cancelled field value if set, zero value otherwise.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetCancelled() bool {
	if o == nil || isNil(o.Cancelled) {
		var ret bool
		return ret
	}
	return *o.Cancelled
}

// GetCancelledOk returns a tuple with the Cancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetCancelledOk() (*bool, bool) {
	if o == nil || isNil(o.Cancelled) {
    return nil, false
	}
	return o.Cancelled, true
}

// HasCancelled returns a boolean if a field has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasCancelled() bool {
	if o != nil && !isNil(o.Cancelled) {
		return true
	}

	return false
}

// SetCancelled gets a reference to the given bool and assigns it to the Cancelled field.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetCancelled(v bool) {
	o.Cancelled = &v
}

// GetDeliveryType returns the DeliveryType field value
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDeliveryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryType
}

// GetDeliveryTypeOk returns a tuple with the DeliveryType field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDeliveryTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DeliveryType, true
}

// SetDeliveryType sets field value
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetDeliveryType(v string) {
	o.DeliveryType = v
}

// GetDeliveryUrl returns the DeliveryUrl field value if set, zero value otherwise.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDeliveryUrl() []string {
	if o == nil || isNil(o.DeliveryUrl) {
		var ret []string
		return ret
	}
	return o.DeliveryUrl
}

// GetDeliveryUrlOk returns a tuple with the DeliveryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDeliveryUrlOk() ([]string, bool) {
	if o == nil || isNil(o.DeliveryUrl) {
    return nil, false
	}
	return o.DeliveryUrl, true
}

// HasDeliveryUrl returns a boolean if a field has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasDeliveryUrl() bool {
	if o != nil && !isNil(o.DeliveryUrl) {
		return true
	}

	return false
}

// SetDeliveryUrl gets a reference to the given []string and assigns it to the DeliveryUrl field.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetDeliveryUrl(v []string) {
	o.DeliveryUrl = v
}

// GetDesiredTimestamp returns the DesiredTimestamp field value if set, zero value otherwise.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDesiredTimestamp() ApiBSONTimestampView {
	if o == nil || isNil(o.DesiredTimestamp) {
		var ret ApiBSONTimestampView
		return ret
	}
	return *o.DesiredTimestamp
}

// GetDesiredTimestampOk returns a tuple with the DesiredTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDesiredTimestampOk() (*ApiBSONTimestampView, bool) {
	if o == nil || isNil(o.DesiredTimestamp) {
    return nil, false
	}
	return o.DesiredTimestamp, true
}

// HasDesiredTimestamp returns a boolean if a field has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasDesiredTimestamp() bool {
	if o != nil && !isNil(o.DesiredTimestamp) {
		return true
	}

	return false
}

// SetDesiredTimestamp gets a reference to the given ApiBSONTimestampView and assigns it to the DesiredTimestamp field.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetDesiredTimestamp(v ApiBSONTimestampView) {
	o.DesiredTimestamp = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetExpired() bool {
	if o == nil || isNil(o.Expired) {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetExpiredOk() (*bool, bool) {
	if o == nil || isNil(o.Expired) {
    return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasExpired() bool {
	if o != nil && !isNil(o.Expired) {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetExpired(v bool) {
	o.Expired = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetExpiresAt() time.Time {
	if o == nil || isNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiresAt) {
    return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetFailed() bool {
	if o == nil || isNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetFailedOk() (*bool, bool) {
	if o == nil || isNil(o.Failed) {
    return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasFailed() bool {
	if o != nil && !isNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetFailed(v bool) {
	o.Failed = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetFinishedAt() time.Time {
	if o == nil || isNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.FinishedAt) {
    return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasFinishedAt() bool {
	if o != nil && !isNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetId(v string) {
	o.Id = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetSnapshotId() string {
	if o == nil || isNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetSnapshotIdOk() (*string, bool) {
	if o == nil || isNil(o.SnapshotId) {
    return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasSnapshotId() bool {
	if o != nil && !isNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetTargetClusterName returns the TargetClusterName field value
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTargetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetClusterName
}

// GetTargetClusterNameOk returns a tuple with the TargetClusterName field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTargetClusterNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TargetClusterName, true
}

// SetTargetClusterName sets field value
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetTargetClusterName(v string) {
	o.TargetClusterName = v
}

// GetTargetGroupId returns the TargetGroupId field value
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTargetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetGroupId
}

// GetTargetGroupIdOk returns a tuple with the TargetGroupId field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTargetGroupIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TargetGroupId, true
}

// SetTargetGroupId sets field value
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetTargetGroupId(v string) {
	o.TargetGroupId = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o ApiAtlasServerlessBackupRestoreJobViewManual) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Cancelled) {
		toSerialize["cancelled"] = o.Cancelled
	}
	if true {
		toSerialize["deliveryType"] = o.DeliveryType
	}
	if !isNil(o.DeliveryUrl) {
		toSerialize["deliveryUrl"] = o.DeliveryUrl
	}
	if !isNil(o.DesiredTimestamp) {
		toSerialize["desiredTimestamp"] = o.DesiredTimestamp
	}
	if !isNil(o.Expired) {
		toSerialize["expired"] = o.Expired
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !isNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !isNil(o.FinishedAt) {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.SnapshotId) {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if true {
		toSerialize["targetClusterName"] = o.TargetClusterName
	}
	if true {
		toSerialize["targetGroupId"] = o.TargetGroupId
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasServerlessBackupRestoreJobViewManual struct {
	value *ApiAtlasServerlessBackupRestoreJobViewManual
	isSet bool
}

func (v NullableApiAtlasServerlessBackupRestoreJobViewManual) Get() *ApiAtlasServerlessBackupRestoreJobViewManual {
	return v.value
}

func (v *NullableApiAtlasServerlessBackupRestoreJobViewManual) Set(val *ApiAtlasServerlessBackupRestoreJobViewManual) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasServerlessBackupRestoreJobViewManual) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasServerlessBackupRestoreJobViewManual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasServerlessBackupRestoreJobViewManual(val *ApiAtlasServerlessBackupRestoreJobViewManual) *NullableApiAtlasServerlessBackupRestoreJobViewManual {
	return &NullableApiAtlasServerlessBackupRestoreJobViewManual{value: val, isSet: true}
}

func (v NullableApiAtlasServerlessBackupRestoreJobViewManual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasServerlessBackupRestoreJobViewManual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


