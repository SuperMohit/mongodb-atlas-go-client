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

// ApiAtlasDiskBackupCopySettingView Copy setting item in the desired backup policy.
type ApiAtlasDiskBackupCopySettingView struct {
	// Human-readable label that identifies the cloud provider that stores the snapshot copy.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// List that describes which types of snapshots to copy.
	Frequencies []string `json:"frequencies,omitempty"`
	// Target region to copy snapshots belonging to replicationSpecId to. Please supply the 'Atlas Region' which can be found under https://www.mongodb.com/docs/atlas/reference/cloud-providers/ 'regions' link
	RegionName *string `json:"regionName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Replication Spec Id, do a GET request to Return One Cluster in One Project and consult the replicationSpecs array https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/returnOneCluster
	ReplicationSpecId *string `json:"replicationSpecId,omitempty"`
	// Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores.
	ShouldCopyOplogs *bool `json:"shouldCopyOplogs,omitempty"`
}

// NewApiAtlasDiskBackupCopySettingView instantiates a new ApiAtlasDiskBackupCopySettingView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasDiskBackupCopySettingView() *ApiAtlasDiskBackupCopySettingView {
	this := ApiAtlasDiskBackupCopySettingView{}
	return &this
}

// NewApiAtlasDiskBackupCopySettingViewWithDefaults instantiates a new ApiAtlasDiskBackupCopySettingView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasDiskBackupCopySettingViewWithDefaults() *ApiAtlasDiskBackupCopySettingView {
	this := ApiAtlasDiskBackupCopySettingView{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupCopySettingView) GetCloudProvider() string {
	if o == nil || isNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupCopySettingView) GetCloudProviderOk() (*string, bool) {
	if o == nil || isNil(o.CloudProvider) {
    return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupCopySettingView) HasCloudProvider() bool {
	if o != nil && !isNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ApiAtlasDiskBackupCopySettingView) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetFrequencies returns the Frequencies field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupCopySettingView) GetFrequencies() []string {
	if o == nil || isNil(o.Frequencies) {
		var ret []string
		return ret
	}
	return o.Frequencies
}

// GetFrequenciesOk returns a tuple with the Frequencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupCopySettingView) GetFrequenciesOk() ([]string, bool) {
	if o == nil || isNil(o.Frequencies) {
    return nil, false
	}
	return o.Frequencies, true
}

// HasFrequencies returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupCopySettingView) HasFrequencies() bool {
	if o != nil && !isNil(o.Frequencies) {
		return true
	}

	return false
}

// SetFrequencies gets a reference to the given []string and assigns it to the Frequencies field.
func (o *ApiAtlasDiskBackupCopySettingView) SetFrequencies(v []string) {
	o.Frequencies = v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupCopySettingView) GetRegionName() string {
	if o == nil || isNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupCopySettingView) GetRegionNameOk() (*string, bool) {
	if o == nil || isNil(o.RegionName) {
    return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupCopySettingView) HasRegionName() bool {
	if o != nil && !isNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ApiAtlasDiskBackupCopySettingView) SetRegionName(v string) {
	o.RegionName = &v
}

// GetReplicationSpecId returns the ReplicationSpecId field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupCopySettingView) GetReplicationSpecId() string {
	if o == nil || isNil(o.ReplicationSpecId) {
		var ret string
		return ret
	}
	return *o.ReplicationSpecId
}

// GetReplicationSpecIdOk returns a tuple with the ReplicationSpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupCopySettingView) GetReplicationSpecIdOk() (*string, bool) {
	if o == nil || isNil(o.ReplicationSpecId) {
    return nil, false
	}
	return o.ReplicationSpecId, true
}

// HasReplicationSpecId returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupCopySettingView) HasReplicationSpecId() bool {
	if o != nil && !isNil(o.ReplicationSpecId) {
		return true
	}

	return false
}

// SetReplicationSpecId gets a reference to the given string and assigns it to the ReplicationSpecId field.
func (o *ApiAtlasDiskBackupCopySettingView) SetReplicationSpecId(v string) {
	o.ReplicationSpecId = &v
}

// GetShouldCopyOplogs returns the ShouldCopyOplogs field value if set, zero value otherwise.
func (o *ApiAtlasDiskBackupCopySettingView) GetShouldCopyOplogs() bool {
	if o == nil || isNil(o.ShouldCopyOplogs) {
		var ret bool
		return ret
	}
	return *o.ShouldCopyOplogs
}

// GetShouldCopyOplogsOk returns a tuple with the ShouldCopyOplogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDiskBackupCopySettingView) GetShouldCopyOplogsOk() (*bool, bool) {
	if o == nil || isNil(o.ShouldCopyOplogs) {
    return nil, false
	}
	return o.ShouldCopyOplogs, true
}

// HasShouldCopyOplogs returns a boolean if a field has been set.
func (o *ApiAtlasDiskBackupCopySettingView) HasShouldCopyOplogs() bool {
	if o != nil && !isNil(o.ShouldCopyOplogs) {
		return true
	}

	return false
}

// SetShouldCopyOplogs gets a reference to the given bool and assigns it to the ShouldCopyOplogs field.
func (o *ApiAtlasDiskBackupCopySettingView) SetShouldCopyOplogs(v bool) {
	o.ShouldCopyOplogs = &v
}

func (o ApiAtlasDiskBackupCopySettingView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !isNil(o.Frequencies) {
		toSerialize["frequencies"] = o.Frequencies
	}
	if !isNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !isNil(o.ReplicationSpecId) {
		toSerialize["replicationSpecId"] = o.ReplicationSpecId
	}
	if !isNil(o.ShouldCopyOplogs) {
		toSerialize["shouldCopyOplogs"] = o.ShouldCopyOplogs
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasDiskBackupCopySettingView struct {
	value *ApiAtlasDiskBackupCopySettingView
	isSet bool
}

func (v NullableApiAtlasDiskBackupCopySettingView) Get() *ApiAtlasDiskBackupCopySettingView {
	return v.value
}

func (v *NullableApiAtlasDiskBackupCopySettingView) Set(val *ApiAtlasDiskBackupCopySettingView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasDiskBackupCopySettingView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasDiskBackupCopySettingView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasDiskBackupCopySettingView(val *ApiAtlasDiskBackupCopySettingView) *NullableApiAtlasDiskBackupCopySettingView {
	return &NullableApiAtlasDiskBackupCopySettingView{value: val, isSet: true}
}

func (v NullableApiAtlasDiskBackupCopySettingView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasDiskBackupCopySettingView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

