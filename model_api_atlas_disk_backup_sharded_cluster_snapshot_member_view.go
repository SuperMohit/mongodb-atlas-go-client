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

// ApiAtlasDiskBackupShardedClusterSnapshotMemberView List that includes the snapshots and the cloud provider that stores the snapshots. The resource returns this parameter when `\"type\" : \"SHARDED_CLUSTER\"`.
type ApiAtlasDiskBackupShardedClusterSnapshotMemberView struct {
	// Human-readable label that identifies the cloud provider that stores this snapshot. The resource returns this parameter when `\"type\": \"replicaSet\".`
	CloudProvider string `json:"cloudProvider"`
	// Unique 24-hexadecimal digit string that identifies the snapshot.
	Id string `json:"id"`
	// Human-readable label that identifies the shard or config host from which MongoDB Cloud took this snapshot.
	ReplicaSetName string `json:"replicaSetName"`
}

// NewApiAtlasDiskBackupShardedClusterSnapshotMemberView instantiates a new ApiAtlasDiskBackupShardedClusterSnapshotMemberView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasDiskBackupShardedClusterSnapshotMemberView(cloudProvider string, id string, replicaSetName string) *ApiAtlasDiskBackupShardedClusterSnapshotMemberView {
	this := ApiAtlasDiskBackupShardedClusterSnapshotMemberView{}
	this.CloudProvider = cloudProvider
	this.Id = id
	this.ReplicaSetName = replicaSetName
	return &this
}

// NewApiAtlasDiskBackupShardedClusterSnapshotMemberViewWithDefaults instantiates a new ApiAtlasDiskBackupShardedClusterSnapshotMemberView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasDiskBackupShardedClusterSnapshotMemberViewWithDefaults() *ApiAtlasDiskBackupShardedClusterSnapshotMemberView {
	this := ApiAtlasDiskBackupShardedClusterSnapshotMemberView{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetCloudProviderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetId returns the Id field value
func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) SetId(v string) {
	o.Id = v
}

// GetReplicaSetName returns the ReplicaSetName field value
func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetReplicaSetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetReplicaSetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReplicaSetName, true
}

// SetReplicaSetName sets field value
func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) SetReplicaSetName(v string) {
	o.ReplicaSetName = v
}

func (o ApiAtlasDiskBackupShardedClusterSnapshotMemberView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["replicaSetName"] = o.ReplicaSetName
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasDiskBackupShardedClusterSnapshotMemberView struct {
	value *ApiAtlasDiskBackupShardedClusterSnapshotMemberView
	isSet bool
}

func (v NullableApiAtlasDiskBackupShardedClusterSnapshotMemberView) Get() *ApiAtlasDiskBackupShardedClusterSnapshotMemberView {
	return v.value
}

func (v *NullableApiAtlasDiskBackupShardedClusterSnapshotMemberView) Set(val *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasDiskBackupShardedClusterSnapshotMemberView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasDiskBackupShardedClusterSnapshotMemberView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasDiskBackupShardedClusterSnapshotMemberView(val *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) *NullableApiAtlasDiskBackupShardedClusterSnapshotMemberView {
	return &NullableApiAtlasDiskBackupShardedClusterSnapshotMemberView{value: val, isSet: true}
}

func (v NullableApiAtlasDiskBackupShardedClusterSnapshotMemberView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasDiskBackupShardedClusterSnapshotMemberView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


