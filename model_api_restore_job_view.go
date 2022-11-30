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

// ApiRestoreJobView struct for ApiRestoreJobView
type ApiRestoreJobView struct {
	// Unique 24-hexadecimal digit string that identifies the batch to which this restore job belongs. This parameter exists only for a sharded cluster restore.
	BatchId *string `json:"batchId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the sharded cluster checkpoint. The checkpoint represents the point in time back to which you want to restore you data. This parameter applies when `\"delivery.methodName\" : \"AUTOMATED_RESTORE\"`. Use this parameter with sharded clusters only.  - If you set **checkpointId**, you can't set **oplogInc**, **oplogTs**, **snapshotId**, or **pointInTimeUTCMillis**. - If you provide this parameter, this endpoint restores all data up to this checkpoint to the database you specify in the **delivery** object.
	CheckpointId *string `json:"checkpointId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the cluster with the snapshot you want to return. This parameter returns for restore clusters.
	ClusterId *string `json:"clusterId,omitempty"`
	// Human-readable label that identifies the cluster containing the snapshots you want to retrieve.
	ClusterName *string `json:"clusterName,omitempty"`
	// Date and time when someone requested this restore job. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	Delivery ApiRestoreJobDeliveryView `json:"delivery"`
	// Flag that indicates whether someone encrypted the data in the restored snapshot.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project that owns the snapshots.
	GroupId *string `json:"groupId,omitempty"`
	// List that contains documents mapping each restore file to a hashed checksum. This parameter applies after you download the corresponding **delivery.url**. If `\"methodName\" : \"HTTP\"`, this list contains one object that represents the hash of the **.tar.gz** file.
	Hashes []ApiRestoreJobFileHashView `json:"hashes,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the restore job.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// Universally Unique Identifier (UUID) that identifies the Key Management Interoperability (KMIP) master key used to encrypt the snapshot data. This parameter applies only when `\"encryptionEnabled\" : \"true\"`.
	MasterKeyUUID *string `json:"masterKeyUUID,omitempty"`
	// Thirty-two-bit incrementing ordinal that represents operations within a given second. When paired with **oplogTs**, this represents the point in time to which MongoDB Cloud restores your data. This parameter applies when `\"delivery.methodName\" : \"AUTOMATED_RESTORE\"`.  - If you set **oplogInc**, you must set **oplogTs**, and can't set **checkpointId**, **snapshotId**, or **pointInTimeUTCMillis**. - If you provide this parameter, this endpoint restores all data up to and including this Oplog timestamp to the database you specified in the **delivery** object.
	OplogInc *float32 `json:"oplogInc,omitempty"`
	// Date and time from which you want to restore this snapshot. This parameter expresses its value in ISO 8601 format in UTC. This represents the first part of an Oplog timestamp. When paired with **oplogInc**, they represent the last database operation to which you want to restore your data. This parameter applies when `\"delivery.methodName\" : \"AUTOMATED_RESTORE\"`. Run a query against **local.oplog.rs** on your replica set to find the desired timestamp.  - If you set **oplogTs**, you must set **oplogInc**, and you can't set **checkpointId**, **snapshotId**, or **pointInTimeUTCMillis**. - If you provide this parameter, this endpoint restores all data up to and including this Oplog timestamp to the database you specified in the **delivery** object.
	OplogTs *string `json:"oplogTs,omitempty"`
	// Timestamp from which you want to restore this snapshot. This parameter expresses its value in the number of milliseconds elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time). This timestamp must fall within the last 24 hours of the current time. This parameter applies when `\"delivery.methodName\" : \"AUTOMATED_RESTORE\"`.  - If you provide this parameter, this endpoint restores all data up to this point in time to the database you specified in the **delivery** object. - If you set **pointInTimeUTCMillis**, you can't set **oplogInc**, **oplogTs**, **snapshotId**, or **checkpointId**.
	PointInTimeUTCMillis *int64 `json:"pointInTimeUTCMillis,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the snapshot to restore. If you set **snapshotId**, you can't set **oplogInc**, **oplogTs**, **pointInTimeUTCMillis**, or **checkpointId**.
	SnapshotId *string `json:"snapshotId,omitempty"`
	// Human-readable label that identifies the status of the downloadable file at the time of the request.
	StatusName *string `json:"statusName,omitempty"`
	Timestamp *ApiBSONTimestampView `json:"timestamp,omitempty"`
}

// NewApiRestoreJobView instantiates a new ApiRestoreJobView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRestoreJobView(delivery ApiRestoreJobDeliveryView, links []Link) *ApiRestoreJobView {
	this := ApiRestoreJobView{}
	this.Delivery = delivery
	this.Links = links
	return &this
}

// NewApiRestoreJobViewWithDefaults instantiates a new ApiRestoreJobView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRestoreJobViewWithDefaults() *ApiRestoreJobView {
	this := ApiRestoreJobView{}
	return &this
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetBatchId() string {
	if o == nil || isNil(o.BatchId) {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetBatchIdOk() (*string, bool) {
	if o == nil || isNil(o.BatchId) {
    return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasBatchId() bool {
	if o != nil && !isNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *ApiRestoreJobView) SetBatchId(v string) {
	o.BatchId = &v
}

// GetCheckpointId returns the CheckpointId field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetCheckpointId() string {
	if o == nil || isNil(o.CheckpointId) {
		var ret string
		return ret
	}
	return *o.CheckpointId
}

// GetCheckpointIdOk returns a tuple with the CheckpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetCheckpointIdOk() (*string, bool) {
	if o == nil || isNil(o.CheckpointId) {
    return nil, false
	}
	return o.CheckpointId, true
}

// HasCheckpointId returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasCheckpointId() bool {
	if o != nil && !isNil(o.CheckpointId) {
		return true
	}

	return false
}

// SetCheckpointId gets a reference to the given string and assigns it to the CheckpointId field.
func (o *ApiRestoreJobView) SetCheckpointId(v string) {
	o.CheckpointId = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetClusterId() string {
	if o == nil || isNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetClusterIdOk() (*string, bool) {
	if o == nil || isNil(o.ClusterId) {
    return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasClusterId() bool {
	if o != nil && !isNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *ApiRestoreJobView) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetClusterName() string {
	if o == nil || isNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetClusterNameOk() (*string, bool) {
	if o == nil || isNil(o.ClusterName) {
    return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasClusterName() bool {
	if o != nil && !isNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ApiRestoreJobView) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
    return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ApiRestoreJobView) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDelivery returns the Delivery field value
func (o *ApiRestoreJobView) GetDelivery() ApiRestoreJobDeliveryView {
	if o == nil {
		var ret ApiRestoreJobDeliveryView
		return ret
	}

	return o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetDeliveryOk() (*ApiRestoreJobDeliveryView, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Delivery, true
}

// SetDelivery sets field value
func (o *ApiRestoreJobView) SetDelivery(v ApiRestoreJobDeliveryView) {
	o.Delivery = v
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetEncryptionEnabled() bool {
	if o == nil || isNil(o.EncryptionEnabled) {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.EncryptionEnabled) {
    return nil, false
	}
	return o.EncryptionEnabled, true
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasEncryptionEnabled() bool {
	if o != nil && !isNil(o.EncryptionEnabled) {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given bool and assigns it to the EncryptionEnabled field.
func (o *ApiRestoreJobView) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiRestoreJobView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetHashes returns the Hashes field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetHashes() []ApiRestoreJobFileHashView {
	if o == nil || isNil(o.Hashes) {
		var ret []ApiRestoreJobFileHashView
		return ret
	}
	return o.Hashes
}

// GetHashesOk returns a tuple with the Hashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetHashesOk() ([]ApiRestoreJobFileHashView, bool) {
	if o == nil || isNil(o.Hashes) {
    return nil, false
	}
	return o.Hashes, true
}

// HasHashes returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasHashes() bool {
	if o != nil && !isNil(o.Hashes) {
		return true
	}

	return false
}

// SetHashes gets a reference to the given []ApiRestoreJobFileHashView and assigns it to the Hashes field.
func (o *ApiRestoreJobView) SetHashes(v []ApiRestoreJobFileHashView) {
	o.Hashes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiRestoreJobView) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value
func (o *ApiRestoreJobView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ApiRestoreJobView) SetLinks(v []Link) {
	o.Links = v
}

// GetMasterKeyUUID returns the MasterKeyUUID field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetMasterKeyUUID() string {
	if o == nil || isNil(o.MasterKeyUUID) {
		var ret string
		return ret
	}
	return *o.MasterKeyUUID
}

// GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetMasterKeyUUIDOk() (*string, bool) {
	if o == nil || isNil(o.MasterKeyUUID) {
    return nil, false
	}
	return o.MasterKeyUUID, true
}

// HasMasterKeyUUID returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasMasterKeyUUID() bool {
	if o != nil && !isNil(o.MasterKeyUUID) {
		return true
	}

	return false
}

// SetMasterKeyUUID gets a reference to the given string and assigns it to the MasterKeyUUID field.
func (o *ApiRestoreJobView) SetMasterKeyUUID(v string) {
	o.MasterKeyUUID = &v
}

// GetOplogInc returns the OplogInc field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetOplogInc() float32 {
	if o == nil || isNil(o.OplogInc) {
		var ret float32
		return ret
	}
	return *o.OplogInc
}

// GetOplogIncOk returns a tuple with the OplogInc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetOplogIncOk() (*float32, bool) {
	if o == nil || isNil(o.OplogInc) {
    return nil, false
	}
	return o.OplogInc, true
}

// HasOplogInc returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasOplogInc() bool {
	if o != nil && !isNil(o.OplogInc) {
		return true
	}

	return false
}

// SetOplogInc gets a reference to the given float32 and assigns it to the OplogInc field.
func (o *ApiRestoreJobView) SetOplogInc(v float32) {
	o.OplogInc = &v
}

// GetOplogTs returns the OplogTs field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetOplogTs() string {
	if o == nil || isNil(o.OplogTs) {
		var ret string
		return ret
	}
	return *o.OplogTs
}

// GetOplogTsOk returns a tuple with the OplogTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetOplogTsOk() (*string, bool) {
	if o == nil || isNil(o.OplogTs) {
    return nil, false
	}
	return o.OplogTs, true
}

// HasOplogTs returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasOplogTs() bool {
	if o != nil && !isNil(o.OplogTs) {
		return true
	}

	return false
}

// SetOplogTs gets a reference to the given string and assigns it to the OplogTs field.
func (o *ApiRestoreJobView) SetOplogTs(v string) {
	o.OplogTs = &v
}

// GetPointInTimeUTCMillis returns the PointInTimeUTCMillis field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetPointInTimeUTCMillis() int64 {
	if o == nil || isNil(o.PointInTimeUTCMillis) {
		var ret int64
		return ret
	}
	return *o.PointInTimeUTCMillis
}

// GetPointInTimeUTCMillisOk returns a tuple with the PointInTimeUTCMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetPointInTimeUTCMillisOk() (*int64, bool) {
	if o == nil || isNil(o.PointInTimeUTCMillis) {
    return nil, false
	}
	return o.PointInTimeUTCMillis, true
}

// HasPointInTimeUTCMillis returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasPointInTimeUTCMillis() bool {
	if o != nil && !isNil(o.PointInTimeUTCMillis) {
		return true
	}

	return false
}

// SetPointInTimeUTCMillis gets a reference to the given int64 and assigns it to the PointInTimeUTCMillis field.
func (o *ApiRestoreJobView) SetPointInTimeUTCMillis(v int64) {
	o.PointInTimeUTCMillis = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetSnapshotId() string {
	if o == nil || isNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetSnapshotIdOk() (*string, bool) {
	if o == nil || isNil(o.SnapshotId) {
    return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasSnapshotId() bool {
	if o != nil && !isNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *ApiRestoreJobView) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetStatusName returns the StatusName field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetStatusName() string {
	if o == nil || isNil(o.StatusName) {
		var ret string
		return ret
	}
	return *o.StatusName
}

// GetStatusNameOk returns a tuple with the StatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetStatusNameOk() (*string, bool) {
	if o == nil || isNil(o.StatusName) {
    return nil, false
	}
	return o.StatusName, true
}

// HasStatusName returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasStatusName() bool {
	if o != nil && !isNil(o.StatusName) {
		return true
	}

	return false
}

// SetStatusName gets a reference to the given string and assigns it to the StatusName field.
func (o *ApiRestoreJobView) SetStatusName(v string) {
	o.StatusName = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ApiRestoreJobView) GetTimestamp() ApiBSONTimestampView {
	if o == nil || isNil(o.Timestamp) {
		var ret ApiBSONTimestampView
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobView) GetTimestampOk() (*ApiBSONTimestampView, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ApiRestoreJobView) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given ApiBSONTimestampView and assigns it to the Timestamp field.
func (o *ApiRestoreJobView) SetTimestamp(v ApiBSONTimestampView) {
	o.Timestamp = &v
}

func (o ApiRestoreJobView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BatchId) {
		toSerialize["batchId"] = o.BatchId
	}
	if !isNil(o.CheckpointId) {
		toSerialize["checkpointId"] = o.CheckpointId
	}
	if !isNil(o.ClusterId) {
		toSerialize["clusterId"] = o.ClusterId
	}
	if !isNil(o.ClusterName) {
		toSerialize["clusterName"] = o.ClusterName
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["delivery"] = o.Delivery
	}
	if !isNil(o.EncryptionEnabled) {
		toSerialize["encryptionEnabled"] = o.EncryptionEnabled
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.Hashes) {
		toSerialize["hashes"] = o.Hashes
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.MasterKeyUUID) {
		toSerialize["masterKeyUUID"] = o.MasterKeyUUID
	}
	if !isNil(o.OplogInc) {
		toSerialize["oplogInc"] = o.OplogInc
	}
	if !isNil(o.OplogTs) {
		toSerialize["oplogTs"] = o.OplogTs
	}
	if !isNil(o.PointInTimeUTCMillis) {
		toSerialize["pointInTimeUTCMillis"] = o.PointInTimeUTCMillis
	}
	if !isNil(o.SnapshotId) {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if !isNil(o.StatusName) {
		toSerialize["statusName"] = o.StatusName
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableApiRestoreJobView struct {
	value *ApiRestoreJobView
	isSet bool
}

func (v NullableApiRestoreJobView) Get() *ApiRestoreJobView {
	return v.value
}

func (v *NullableApiRestoreJobView) Set(val *ApiRestoreJobView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRestoreJobView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRestoreJobView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRestoreJobView(val *ApiRestoreJobView) *NullableApiRestoreJobView {
	return &NullableApiRestoreJobView{value: val, isSet: true}
}

func (v NullableApiRestoreJobView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRestoreJobView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


