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

// ApiAtlasDiskBackupExportJobView struct for ApiAtlasDiskBackupExportJobView
type ApiAtlasDiskBackupExportJobView struct {
	// Information on the export job for each replica set in the sharded cluster.
	Components []ApiAtlasDiskBackupBaseRestoreMemberView `json:"components,omitempty"`
	// Date and time when someone created this export job. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Collection of key-value pairs that represent custom data for the metadata file that MongoDB Cloud uploads to the bucket when the export job finishes.
	CustomData []ApiAtlasLabelView `json:"customData,omitempty"`
	// One or more Uniform Resource Locators (URLs) that point to the compressed snapshot files for manual download. MongoDB Cloud returns this parameter when `\"deliveryType\" : \"download\"`.
	DeliveryUrl []string `json:"deliveryUrl,omitempty"`
	// Unique 24-hexadecimal character string that identifies the AWS bucket to which MongoDB Cloud exports the Cloud Backup snapshot.
	ExportBucketId string `json:"exportBucketId"`
	ExportStatus *ApiExportStatusView `json:"exportStatus,omitempty"`
	// Date and time when this export job completed. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// Unique 24-hexadecimal character string that identifies the restore job.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// Full path on the cloud provider bucket to the folder where the snapshot is exported.
	Prefix *string `json:"prefix,omitempty"`
	// Unique 24-hexadecimal character string that identifies the snapshot.
	SnapshotId *string `json:"snapshotId,omitempty"`
	// State of the export job.
	State *string `json:"state,omitempty"`
}

// NewApiAtlasDiskBackupExportJobView instantiates a new ApiAtlasDiskBackupExportJobView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasDiskBackupExportJobView(exportBucketId string, links []Link) *ApiAtlasDiskBackupExportJobView {
	this := ApiAtlasDiskBackupExportJobView{}
	this.ExportBucketId = exportBucketId
	this.Links = links
	return &this
}

// NewApiAtlasDiskBackupExportJobViewWithDefaults instantiates a new ApiAtlasDiskBackupExportJobView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasDiskBackupExportJobViewWithDefaults() *ApiAtlasDiskBackupExportJobView {
	this := ApiAtlasDiskBackupExportJobView{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupExportJobView) GetComponents() []ApiAtlasDiskBackupBaseRestoreMemberView {
	if o == nil || isNil(o.Components) {
		var ret []ApiAtlasDiskBackupBaseRestoreMemberView
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetComponentsOk() ([]ApiAtlasDiskBackupBaseRestoreMemberView, bool) {
	if o == nil || isNil(o.Components) {
    return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupExportJobView) HasComponents() bool {
	if o != nil && !isNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ApiAtlasDiskBackupBaseRestoreMemberView and assigns it to the Components field.
func (o *ApiAtlasDiskBackupExportJobView) SetComponents(v []ApiAtlasDiskBackupBaseRestoreMemberView) {
	o.Components = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupExportJobView) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupExportJobView) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ApiAtlasDiskBackupExportJobView) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupExportJobView) GetCustomData() []ApiAtlasLabelView {
	if o == nil || isNil(o.CustomData) {
		var ret []ApiAtlasLabelView
		return ret
	}
	return o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetCustomDataOk() ([]ApiAtlasLabelView, bool) {
	if o == nil || isNil(o.CustomData) {
    return nil, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupExportJobView) HasCustomData() bool {
	if o != nil && !isNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given []ApiAtlasLabelView and assigns it to the CustomData field.
func (o *ApiAtlasDiskBackupExportJobView) SetCustomData(v []ApiAtlasLabelView) {
	o.CustomData = v
}

// GetDeliveryUrl returns the DeliveryUrl field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupExportJobView) GetDeliveryUrl() []string {
	if o == nil || isNil(o.DeliveryUrl) {
		var ret []string
		return ret
	}
	return o.DeliveryUrl
}

// GetDeliveryUrlOk returns a tuple with the DeliveryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetDeliveryUrlOk() ([]string, bool) {
	if o == nil || isNil(o.DeliveryUrl) {
    return nil, false
	}
	return o.DeliveryUrl, true
}

// HasDeliveryUrl returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupExportJobView) HasDeliveryUrl() bool {
	if o != nil && !isNil(o.DeliveryUrl) {
		return true
	}

	return false
}

// SetDeliveryUrl gets a reference to the given []string and assigns it to the DeliveryUrl field.
func (o *ApiAtlasDiskBackupExportJobView) SetDeliveryUrl(v []string) {
	o.DeliveryUrl = v
}

// GetExportBucketId returns the ExportBucketId field value
func (o *ApiAtlasDiskBackupExportJobView) GetExportBucketId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExportBucketId
}

// GetExportBucketIdOk returns a tuple with the ExportBucketId field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetExportBucketIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExportBucketId, true
}

// SetExportBucketId sets field value
func (o *ApiAtlasDiskBackupExportJobView) SetExportBucketId(v string) {
	o.ExportBucketId = v
}

// GetExportStatus returns the ExportStatus field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupExportJobView) GetExportStatus() ApiExportStatusView {
	if o == nil || isNil(o.ExportStatus) {
		var ret ApiExportStatusView
		return ret
	}
	return *o.ExportStatus
}

// GetExportStatusOk returns a tuple with the ExportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetExportStatusOk() (*ApiExportStatusView, bool) {
	if o == nil || isNil(o.ExportStatus) {
    return nil, false
	}
	return o.ExportStatus, true
}

// HasExportStatus returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupExportJobView) HasExportStatus() bool {
	if o != nil && !isNil(o.ExportStatus) {
		return true
	}

	return false
}

// SetExportStatus gets a reference to the given ApiExportStatusView and assigns it to the ExportStatus field.
func (o *ApiAtlasDiskBackupExportJobView) SetExportStatus(v ApiExportStatusView) {
	o.ExportStatus = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupExportJobView) GetFinishedAt() time.Time {
	if o == nil || isNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.FinishedAt) {
    return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupExportJobView) HasFinishedAt() bool {
	if o != nil && !isNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *ApiAtlasDiskBackupExportJobView) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupExportJobView) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupExportJobView) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAtlasDiskBackupExportJobView) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value
func (o *ApiAtlasDiskBackupExportJobView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ApiAtlasDiskBackupExportJobView) SetLinks(v []Link) {
	o.Links = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupExportJobView) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupExportJobView) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *ApiAtlasDiskBackupExportJobView) SetPrefix(v string) {
	o.Prefix = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupExportJobView) GetSnapshotId() string {
	if o == nil || isNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetSnapshotIdOk() (*string, bool) {
	if o == nil || isNil(o.SnapshotId) {
    return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupExportJobView) HasSnapshotId() bool {
	if o != nil && !isNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *ApiAtlasDiskBackupExportJobView) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupExportJobView) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupExportJobView) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupExportJobView) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ApiAtlasDiskBackupExportJobView) SetState(v string) {
	o.State = &v
}

func (o ApiAtlasDiskBackupExportJobView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CustomData) {
		toSerialize["customData"] = o.CustomData
	}
	if !isNil(o.DeliveryUrl) {
		toSerialize["deliveryUrl"] = o.DeliveryUrl
	}
	if true {
		toSerialize["exportBucketId"] = o.ExportBucketId
	}
	if !isNil(o.ExportStatus) {
		toSerialize["exportStatus"] = o.ExportStatus
	}
	if !isNil(o.FinishedAt) {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.SnapshotId) {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasDiskBackupExportJobView struct {
	value *ApiAtlasDiskBackupExportJobView
	isSet bool
}

func (v NullableApiAtlasDiskBackupExportJobView) Get() *ApiAtlasDiskBackupExportJobView {
	return v.value
}

func (v *NullableApiAtlasDiskBackupExportJobView) Set(val *ApiAtlasDiskBackupExportJobView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasDiskBackupExportJobView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasDiskBackupExportJobView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasDiskBackupExportJobView(val *ApiAtlasDiskBackupExportJobView) *NullableApiAtlasDiskBackupExportJobView {
	return &NullableApiAtlasDiskBackupExportJobView{value: val, isSet: true}
}

func (v NullableApiAtlasDiskBackupExportJobView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasDiskBackupExportJobView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

