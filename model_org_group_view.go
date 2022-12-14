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

// OrgGroupView struct for OrgGroupView
type OrgGroupView struct {
	// Settings that describe the clusters in each project that the API key is authorized to view.
	Clusters []ClusterView `json:"clusters,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	GroupId *string `json:"groupId,omitempty"`
	// Human-readable label that identifies the project.
	GroupName *string `json:"groupName,omitempty"`
	// Unique 24-hexadecimal character string that identifies the organization that contains the project.
	OrgId *string `json:"orgId,omitempty"`
	// Human-readable label that identifies the organization that contains the project.
	OrgName *string `json:"orgName,omitempty"`
	// Human-readable label that indicates the plan type.
	PlanType *string `json:"planType,omitempty"`
	// List of human-readable labels that categorize the specified project. MongoDB Cloud returns an empty array.
	Tags []string `json:"tags,omitempty"`
}

// NewOrgGroupView instantiates a new OrgGroupView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgGroupView() *OrgGroupView {
	this := OrgGroupView{}
	return &this
}

// NewOrgGroupViewWithDefaults instantiates a new OrgGroupView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgGroupViewWithDefaults() *OrgGroupView {
	this := OrgGroupView{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *OrgGroupView) GetClusters() []ClusterView {
	if o == nil || isNil(o.Clusters) {
		var ret []ClusterView
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupView) GetClustersOk() ([]ClusterView, bool) {
	if o == nil || isNil(o.Clusters) {
    return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *OrgGroupView) HasClusters() bool {
	if o != nil && !isNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []ClusterView and assigns it to the Clusters field.
func (o *OrgGroupView) SetClusters(v []ClusterView) {
	o.Clusters = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *OrgGroupView) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupView) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *OrgGroupView) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *OrgGroupView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *OrgGroupView) GetGroupName() string {
	if o == nil || isNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupView) GetGroupNameOk() (*string, bool) {
	if o == nil || isNil(o.GroupName) {
    return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *OrgGroupView) HasGroupName() bool {
	if o != nil && !isNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *OrgGroupView) SetGroupName(v string) {
	o.GroupName = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *OrgGroupView) GetOrgId() string {
	if o == nil || isNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupView) GetOrgIdOk() (*string, bool) {
	if o == nil || isNil(o.OrgId) {
    return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *OrgGroupView) HasOrgId() bool {
	if o != nil && !isNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *OrgGroupView) SetOrgId(v string) {
	o.OrgId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *OrgGroupView) GetOrgName() string {
	if o == nil || isNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupView) GetOrgNameOk() (*string, bool) {
	if o == nil || isNil(o.OrgName) {
    return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *OrgGroupView) HasOrgName() bool {
	if o != nil && !isNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *OrgGroupView) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPlanType returns the PlanType field value if set, zero value otherwise.
func (o *OrgGroupView) GetPlanType() string {
	if o == nil || isNil(o.PlanType) {
		var ret string
		return ret
	}
	return *o.PlanType
}

// GetPlanTypeOk returns a tuple with the PlanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupView) GetPlanTypeOk() (*string, bool) {
	if o == nil || isNil(o.PlanType) {
    return nil, false
	}
	return o.PlanType, true
}

// HasPlanType returns a boolean if a field has been set.
func (o *OrgGroupView) HasPlanType() bool {
	if o != nil && !isNil(o.PlanType) {
		return true
	}

	return false
}

// SetPlanType gets a reference to the given string and assigns it to the PlanType field.
func (o *OrgGroupView) SetPlanType(v string) {
	o.PlanType = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OrgGroupView) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupView) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OrgGroupView) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *OrgGroupView) SetTags(v []string) {
	o.Tags = v
}

func (o OrgGroupView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !isNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !isNil(o.OrgName) {
		toSerialize["orgName"] = o.OrgName
	}
	if !isNil(o.PlanType) {
		toSerialize["planType"] = o.PlanType
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableOrgGroupView struct {
	value *OrgGroupView
	isSet bool
}

func (v NullableOrgGroupView) Get() *OrgGroupView {
	return v.value
}

func (v *NullableOrgGroupView) Set(val *OrgGroupView) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgGroupView) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgGroupView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgGroupView(val *OrgGroupView) *NullableOrgGroupView {
	return &NullableOrgGroupView{value: val, isSet: true}
}

func (v NullableOrgGroupView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgGroupView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


