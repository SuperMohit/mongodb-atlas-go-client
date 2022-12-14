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

// ApiAtlasGroupView struct for ApiAtlasGroupView
type ApiAtlasGroupView struct {
	// Quantity of MongoDB Cloud clusters deployed in this project.
	ClusterCount int64 `json:"clusterCount"`
	// Date and time when MongoDB Cloud created this project. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Created time.Time `json:"created"`
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud project.
	Id string `json:"id"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// Human-readable label that identifies the project included in the MongoDB Cloud organization.
	Name string `json:"name"`
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud organization to which the project belongs.
	OrgId string `json:"orgId"`
	// Flag that indicates whether to create the project with default alert settings.
	WithDefaultAlertsSettings *bool `json:"withDefaultAlertsSettings,omitempty"`
}

// NewApiAtlasGroupView instantiates a new ApiAtlasGroupView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasGroupView(clusterCount int64, created time.Time, id string, links []Link, name string, orgId string) *ApiAtlasGroupView {
	this := ApiAtlasGroupView{}
	this.ClusterCount = clusterCount
	this.Created = created
	this.Id = id
	this.Links = links
	this.Name = name
	this.OrgId = orgId
	return &this
}

// NewApiAtlasGroupViewWithDefaults instantiates a new ApiAtlasGroupView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasGroupViewWithDefaults() *ApiAtlasGroupView {
	this := ApiAtlasGroupView{}
	return &this
}

// GetClusterCount returns the ClusterCount field value
func (o *ApiAtlasGroupView) GetClusterCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupView) GetClusterCountOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClusterCount, true
}

// SetClusterCount sets field value
func (o *ApiAtlasGroupView) SetClusterCount(v int64) {
	o.ClusterCount = v
}

// GetCreated returns the Created field value
func (o *ApiAtlasGroupView) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupView) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ApiAtlasGroupView) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *ApiAtlasGroupView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupView) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiAtlasGroupView) SetId(v string) {
	o.Id = v
}

// GetLinks returns the Links field value
func (o *ApiAtlasGroupView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ApiAtlasGroupView) SetLinks(v []Link) {
	o.Links = v
}

// GetName returns the Name field value
func (o *ApiAtlasGroupView) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupView) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiAtlasGroupView) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value
func (o *ApiAtlasGroupView) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupView) GetOrgIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ApiAtlasGroupView) SetOrgId(v string) {
	o.OrgId = v
}

// GetWithDefaultAlertsSettings returns the WithDefaultAlertsSettings field value if set, zero value otherwise.
func (o *ApiAtlasGroupView) GetWithDefaultAlertsSettings() bool {
	if o == nil || isNil(o.WithDefaultAlertsSettings) {
		var ret bool
		return ret
	}
	return *o.WithDefaultAlertsSettings
}

// GetWithDefaultAlertsSettingsOk returns a tuple with the WithDefaultAlertsSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupView) GetWithDefaultAlertsSettingsOk() (*bool, bool) {
	if o == nil || isNil(o.WithDefaultAlertsSettings) {
    return nil, false
	}
	return o.WithDefaultAlertsSettings, true
}

// HasWithDefaultAlertsSettings returns a boolean if a field has been set.
func (o *ApiAtlasGroupView) HasWithDefaultAlertsSettings() bool {
	if o != nil && !isNil(o.WithDefaultAlertsSettings) {
		return true
	}

	return false
}

// SetWithDefaultAlertsSettings gets a reference to the given bool and assigns it to the WithDefaultAlertsSettings field.
func (o *ApiAtlasGroupView) SetWithDefaultAlertsSettings(v bool) {
	o.WithDefaultAlertsSettings = &v
}

func (o ApiAtlasGroupView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clusterCount"] = o.ClusterCount
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["orgId"] = o.OrgId
	}
	if !isNil(o.WithDefaultAlertsSettings) {
		toSerialize["withDefaultAlertsSettings"] = o.WithDefaultAlertsSettings
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasGroupView struct {
	value *ApiAtlasGroupView
	isSet bool
}

func (v NullableApiAtlasGroupView) Get() *ApiAtlasGroupView {
	return v.value
}

func (v *NullableApiAtlasGroupView) Set(val *ApiAtlasGroupView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasGroupView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasGroupView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasGroupView(val *ApiAtlasGroupView) *NullableApiAtlasGroupView {
	return &NullableApiAtlasGroupView{value: val, isSet: true}
}

func (v NullableApiAtlasGroupView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasGroupView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


