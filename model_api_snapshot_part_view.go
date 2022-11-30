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

// ApiSnapshotPartView Characteristics that identify this snapshot.
type ApiSnapshotPartView struct {
	// Unique 24-hexadecimal digit string that identifies the cluster with the snapshots you want to return.
	ClusterId *string `json:"clusterId,omitempty"`
	// Human-readable label that identifies the method of compression for the snapshot.
	CompressionSetting *string `json:"compressionSetting,omitempty"`
	// Total size of the data stored on each node in the cluster. This parameter expresses its value in bytes.
	DataSizeBytes *int64 `json:"dataSizeBytes,omitempty"`
	// Flag that indicates whether someone encrypted this snapshot.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty"`
	// Number that indicates the total size of the data files in bytes.
	FileSizeBytes *int64 `json:"fileSizeBytes,omitempty"`
	// Unique string that identifies the Key Management Interoperability (KMIP) master key used to encrypt the snapshot data. The resource returns this parameter when `\"parts.encryptionEnabled\" : true`.
	MasterKeyUUID *string `json:"masterKeyUUID,omitempty"`
	// Number that indicates the version of MongoDB that the replica set primary ran when MongoDB Cloud created the snapshot.
	MongodVersion *string `json:"mongodVersion,omitempty"`
	// Human-readable label that identifies the replica set.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// Number that indicates the total size of space allocated for document storage.
	StorageSizeBytes *int64 `json:"storageSizeBytes,omitempty"`
	// Human-readable label that identifies the type of server from which MongoDB Cloud took this snapshot.
	TypeName *string `json:"typeName,omitempty"`
}

// NewApiSnapshotPartView instantiates a new ApiSnapshotPartView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSnapshotPartView() *ApiSnapshotPartView {
	this := ApiSnapshotPartView{}
	return &this
}

// NewApiSnapshotPartViewWithDefaults instantiates a new ApiSnapshotPartView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSnapshotPartViewWithDefaults() *ApiSnapshotPartView {
	this := ApiSnapshotPartView{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *ApiSnapshotPartView) GetClusterId() string {
	if o == nil || isNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSnapshotPartView) GetClusterIdOk() (*string, bool) {
	if o == nil || isNil(o.ClusterId) {
    return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *ApiSnapshotPartView) HasClusterId() bool {
	if o != nil && !isNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *ApiSnapshotPartView) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetCompressionSetting returns the CompressionSetting field value if set, zero value otherwise.
func (o *ApiSnapshotPartView) GetCompressionSetting() string {
	if o == nil || isNil(o.CompressionSetting) {
		var ret string
		return ret
	}
	return *o.CompressionSetting
}

// GetCompressionSettingOk returns a tuple with the CompressionSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSnapshotPartView) GetCompressionSettingOk() (*string, bool) {
	if o == nil || isNil(o.CompressionSetting) {
    return nil, false
	}
	return o.CompressionSetting, true
}

// HasCompressionSetting returns a boolean if a field has been set.
func (o *ApiSnapshotPartView) HasCompressionSetting() bool {
	if o != nil && !isNil(o.CompressionSetting) {
		return true
	}

	return false
}

// SetCompressionSetting gets a reference to the given string and assigns it to the CompressionSetting field.
func (o *ApiSnapshotPartView) SetCompressionSetting(v string) {
	o.CompressionSetting = &v
}

// GetDataSizeBytes returns the DataSizeBytes field value if set, zero value otherwise.
func (o *ApiSnapshotPartView) GetDataSizeBytes() int64 {
	if o == nil || isNil(o.DataSizeBytes) {
		var ret int64
		return ret
	}
	return *o.DataSizeBytes
}

// GetDataSizeBytesOk returns a tuple with the DataSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSnapshotPartView) GetDataSizeBytesOk() (*int64, bool) {
	if o == nil || isNil(o.DataSizeBytes) {
    return nil, false
	}
	return o.DataSizeBytes, true
}

// HasDataSizeBytes returns a boolean if a field has been set.
func (o *ApiSnapshotPartView) HasDataSizeBytes() bool {
	if o != nil && !isNil(o.DataSizeBytes) {
		return true
	}

	return false
}

// SetDataSizeBytes gets a reference to the given int64 and assigns it to the DataSizeBytes field.
func (o *ApiSnapshotPartView) SetDataSizeBytes(v int64) {
	o.DataSizeBytes = &v
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise.
func (o *ApiSnapshotPartView) GetEncryptionEnabled() bool {
	if o == nil || isNil(o.EncryptionEnabled) {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSnapshotPartView) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.EncryptionEnabled) {
    return nil, false
	}
	return o.EncryptionEnabled, true
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *ApiSnapshotPartView) HasEncryptionEnabled() bool {
	if o != nil && !isNil(o.EncryptionEnabled) {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given bool and assigns it to the EncryptionEnabled field.
func (o *ApiSnapshotPartView) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled = &v
}

// GetFileSizeBytes returns the FileSizeBytes field value if set, zero value otherwise.
func (o *ApiSnapshotPartView) GetFileSizeBytes() int64 {
	if o == nil || isNil(o.FileSizeBytes) {
		var ret int64
		return ret
	}
	return *o.FileSizeBytes
}

// GetFileSizeBytesOk returns a tuple with the FileSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSnapshotPartView) GetFileSizeBytesOk() (*int64, bool) {
	if o == nil || isNil(o.FileSizeBytes) {
    return nil, false
	}
	return o.FileSizeBytes, true
}

// HasFileSizeBytes returns a boolean if a field has been set.
func (o *ApiSnapshotPartView) HasFileSizeBytes() bool {
	if o != nil && !isNil(o.FileSizeBytes) {
		return true
	}

	return false
}

// SetFileSizeBytes gets a reference to the given int64 and assigns it to the FileSizeBytes field.
func (o *ApiSnapshotPartView) SetFileSizeBytes(v int64) {
	o.FileSizeBytes = &v
}

// GetMasterKeyUUID returns the MasterKeyUUID field value if set, zero value otherwise.
func (o *ApiSnapshotPartView) GetMasterKeyUUID() string {
	if o == nil || isNil(o.MasterKeyUUID) {
		var ret string
		return ret
	}
	return *o.MasterKeyUUID
}

// GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSnapshotPartView) GetMasterKeyUUIDOk() (*string, bool) {
	if o == nil || isNil(o.MasterKeyUUID) {
    return nil, false
	}
	return o.MasterKeyUUID, true
}

// HasMasterKeyUUID returns a boolean if a field has been set.
func (o *ApiSnapshotPartView) HasMasterKeyUUID() bool {
	if o != nil && !isNil(o.MasterKeyUUID) {
		return true
	}

	return false
}

// SetMasterKeyUUID gets a reference to the given string and assigns it to the MasterKeyUUID field.
func (o *ApiSnapshotPartView) SetMasterKeyUUID(v string) {
	o.MasterKeyUUID = &v
}

// GetMongodVersion returns the MongodVersion field value if set, zero value otherwise.
func (o *ApiSnapshotPartView) GetMongodVersion() string {
	if o == nil || isNil(o.MongodVersion) {
		var ret string
		return ret
	}
	return *o.MongodVersion
}

// GetMongodVersionOk returns a tuple with the MongodVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSnapshotPartView) GetMongodVersionOk() (*string, bool) {
	if o == nil || isNil(o.MongodVersion) {
    return nil, false
	}
	return o.MongodVersion, true
}

// HasMongodVersion returns a boolean if a field has been set.
func (o *ApiSnapshotPartView) HasMongodVersion() bool {
	if o != nil && !isNil(o.MongodVersion) {
		return true
	}

	return false
}

// SetMongodVersion gets a reference to the given string and assigns it to the MongodVersion field.
func (o *ApiSnapshotPartView) SetMongodVersion(v string) {
	o.MongodVersion = &v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise.
func (o *ApiSnapshotPartView) GetReplicaSetName() string {
	if o == nil || isNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSnapshotPartView) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || isNil(o.ReplicaSetName) {
    return nil, false
	}
	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *ApiSnapshotPartView) HasReplicaSetName() bool {
	if o != nil && !isNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *ApiSnapshotPartView) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

// GetStorageSizeBytes returns the StorageSizeBytes field value if set, zero value otherwise.
func (o *ApiSnapshotPartView) GetStorageSizeBytes() int64 {
	if o == nil || isNil(o.StorageSizeBytes) {
		var ret int64
		return ret
	}
	return *o.StorageSizeBytes
}

// GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSnapshotPartView) GetStorageSizeBytesOk() (*int64, bool) {
	if o == nil || isNil(o.StorageSizeBytes) {
    return nil, false
	}
	return o.StorageSizeBytes, true
}

// HasStorageSizeBytes returns a boolean if a field has been set.
func (o *ApiSnapshotPartView) HasStorageSizeBytes() bool {
	if o != nil && !isNil(o.StorageSizeBytes) {
		return true
	}

	return false
}

// SetStorageSizeBytes gets a reference to the given int64 and assigns it to the StorageSizeBytes field.
func (o *ApiSnapshotPartView) SetStorageSizeBytes(v int64) {
	o.StorageSizeBytes = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *ApiSnapshotPartView) GetTypeName() string {
	if o == nil || isNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSnapshotPartView) GetTypeNameOk() (*string, bool) {
	if o == nil || isNil(o.TypeName) {
    return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *ApiSnapshotPartView) HasTypeName() bool {
	if o != nil && !isNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *ApiSnapshotPartView) SetTypeName(v string) {
	o.TypeName = &v
}

func (o ApiSnapshotPartView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClusterId) {
		toSerialize["clusterId"] = o.ClusterId
	}
	if !isNil(o.CompressionSetting) {
		toSerialize["compressionSetting"] = o.CompressionSetting
	}
	if !isNil(o.DataSizeBytes) {
		toSerialize["dataSizeBytes"] = o.DataSizeBytes
	}
	if !isNil(o.EncryptionEnabled) {
		toSerialize["encryptionEnabled"] = o.EncryptionEnabled
	}
	if !isNil(o.FileSizeBytes) {
		toSerialize["fileSizeBytes"] = o.FileSizeBytes
	}
	if !isNil(o.MasterKeyUUID) {
		toSerialize["masterKeyUUID"] = o.MasterKeyUUID
	}
	if !isNil(o.MongodVersion) {
		toSerialize["mongodVersion"] = o.MongodVersion
	}
	if !isNil(o.ReplicaSetName) {
		toSerialize["replicaSetName"] = o.ReplicaSetName
	}
	if !isNil(o.StorageSizeBytes) {
		toSerialize["storageSizeBytes"] = o.StorageSizeBytes
	}
	if !isNil(o.TypeName) {
		toSerialize["typeName"] = o.TypeName
	}
	return json.Marshal(toSerialize)
}

type NullableApiSnapshotPartView struct {
	value *ApiSnapshotPartView
	isSet bool
}

func (v NullableApiSnapshotPartView) Get() *ApiSnapshotPartView {
	return v.value
}

func (v *NullableApiSnapshotPartView) Set(val *ApiSnapshotPartView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSnapshotPartView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSnapshotPartView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSnapshotPartView(val *ApiSnapshotPartView) *NullableApiSnapshotPartView {
	return &NullableApiSnapshotPartView{value: val, isSet: true}
}

func (v NullableApiSnapshotPartView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSnapshotPartView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


