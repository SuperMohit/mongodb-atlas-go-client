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

// ApiDeleteCopiedBackupsView Deleted copy setting whose backup copies need to also be deleted.
type ApiDeleteCopiedBackupsView struct {
	// Human-readable label that identifies the cloud provider for the deleted copy setting whose backup copies you want to delete.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Target region for the deleted copy setting whose backup copies you want to delete. Please supply the 'Atlas Region' which can be found under https://www.mongodb.com/docs/atlas/reference/cloud-providers/ 'regions' link
	RegionName *string `json:"regionName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica setclusters, there is only one zone in the cluster. To find the Replication Spec Id, do a GET request to Return One Cluster in One Project and consult the replicationSpecs array https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/returnOneCluster
	ReplicationSpecId *string `json:"replicationSpecId,omitempty"`
}

// NewApiDeleteCopiedBackupsView instantiates a new ApiDeleteCopiedBackupsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDeleteCopiedBackupsView() *ApiDeleteCopiedBackupsView {
	this := ApiDeleteCopiedBackupsView{}
	return &this
}

// NewApiDeleteCopiedBackupsViewWithDefaults instantiates a new ApiDeleteCopiedBackupsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDeleteCopiedBackupsViewWithDefaults() *ApiDeleteCopiedBackupsView {
	this := ApiDeleteCopiedBackupsView{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *ApiDeleteCopiedBackupsView) GetCloudProvider() string {
	if o == nil || isNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDeleteCopiedBackupsView) GetCloudProviderOk() (*string, bool) {
	if o == nil || isNil(o.CloudProvider) {
    return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ApiDeleteCopiedBackupsView) HasCloudProvider() bool {
	if o != nil && !isNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ApiDeleteCopiedBackupsView) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *ApiDeleteCopiedBackupsView) GetRegionName() string {
	if o == nil || isNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDeleteCopiedBackupsView) GetRegionNameOk() (*string, bool) {
	if o == nil || isNil(o.RegionName) {
    return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ApiDeleteCopiedBackupsView) HasRegionName() bool {
	if o != nil && !isNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ApiDeleteCopiedBackupsView) SetRegionName(v string) {
	o.RegionName = &v
}

// GetReplicationSpecId returns the ReplicationSpecId field value if set, zero value otherwise.
func (o *ApiDeleteCopiedBackupsView) GetReplicationSpecId() string {
	if o == nil || isNil(o.ReplicationSpecId) {
		var ret string
		return ret
	}
	return *o.ReplicationSpecId
}

// GetReplicationSpecIdOk returns a tuple with the ReplicationSpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDeleteCopiedBackupsView) GetReplicationSpecIdOk() (*string, bool) {
	if o == nil || isNil(o.ReplicationSpecId) {
    return nil, false
	}
	return o.ReplicationSpecId, true
}

// HasReplicationSpecId returns a boolean if a field has been set.
func (o *ApiDeleteCopiedBackupsView) HasReplicationSpecId() bool {
	if o != nil && !isNil(o.ReplicationSpecId) {
		return true
	}

	return false
}

// SetReplicationSpecId gets a reference to the given string and assigns it to the ReplicationSpecId field.
func (o *ApiDeleteCopiedBackupsView) SetReplicationSpecId(v string) {
	o.ReplicationSpecId = &v
}

func (o ApiDeleteCopiedBackupsView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !isNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !isNil(o.ReplicationSpecId) {
		toSerialize["replicationSpecId"] = o.ReplicationSpecId
	}
	return json.Marshal(toSerialize)
}

type NullableApiDeleteCopiedBackupsView struct {
	value *ApiDeleteCopiedBackupsView
	isSet bool
}

func (v NullableApiDeleteCopiedBackupsView) Get() *ApiDeleteCopiedBackupsView {
	return v.value
}

func (v *NullableApiDeleteCopiedBackupsView) Set(val *ApiDeleteCopiedBackupsView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDeleteCopiedBackupsView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDeleteCopiedBackupsView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDeleteCopiedBackupsView(val *ApiDeleteCopiedBackupsView) *NullableApiDeleteCopiedBackupsView {
	return &NullableApiDeleteCopiedBackupsView{value: val, isSet: true}
}

func (v NullableApiDeleteCopiedBackupsView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDeleteCopiedBackupsView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


