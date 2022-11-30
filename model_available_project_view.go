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

// AvailableProjectView struct for AvailableProjectView
type AvailableProjectView struct {
	// List of clusters that can be migrated to MongoDB Cloud.
	Deployments []AvailableDeploymentView `json:"deployments"`
	MigrationHosts []string `json:"migrationHosts"`
	// Human-readable label that identifies this project.
	Name string `json:"name"`
	// Unique 24-hexadecimal digit string that identifies the project to be migrated.
	ProjectId string `json:"projectId"`
}

// NewAvailableProjectView instantiates a new AvailableProjectView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableProjectView(deployments []AvailableDeploymentView, migrationHosts []string, name string, projectId string) *AvailableProjectView {
	this := AvailableProjectView{}
	this.Deployments = deployments
	this.MigrationHosts = migrationHosts
	this.Name = name
	this.ProjectId = projectId
	return &this
}

// NewAvailableProjectViewWithDefaults instantiates a new AvailableProjectView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableProjectViewWithDefaults() *AvailableProjectView {
	this := AvailableProjectView{}
	return &this
}

// GetDeployments returns the Deployments field value
func (o *AvailableProjectView) GetDeployments() []AvailableDeploymentView {
	if o == nil {
		var ret []AvailableDeploymentView
		return ret
	}

	return o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value
// and a boolean to check if the value has been set.
func (o *AvailableProjectView) GetDeploymentsOk() ([]AvailableDeploymentView, bool) {
	if o == nil {
    return nil, false
	}
	return o.Deployments, true
}

// SetDeployments sets field value
func (o *AvailableProjectView) SetDeployments(v []AvailableDeploymentView) {
	o.Deployments = v
}

// GetMigrationHosts returns the MigrationHosts field value
func (o *AvailableProjectView) GetMigrationHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MigrationHosts
}

// GetMigrationHostsOk returns a tuple with the MigrationHosts field value
// and a boolean to check if the value has been set.
func (o *AvailableProjectView) GetMigrationHostsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MigrationHosts, true
}

// SetMigrationHosts sets field value
func (o *AvailableProjectView) SetMigrationHosts(v []string) {
	o.MigrationHosts = v
}

// GetName returns the Name field value
func (o *AvailableProjectView) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AvailableProjectView) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AvailableProjectView) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *AvailableProjectView) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AvailableProjectView) GetProjectIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AvailableProjectView) SetProjectId(v string) {
	o.ProjectId = v
}

func (o AvailableProjectView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deployments"] = o.Deployments
	}
	if true {
		toSerialize["migrationHosts"] = o.MigrationHosts
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["projectId"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableAvailableProjectView struct {
	value *AvailableProjectView
	isSet bool
}

func (v NullableAvailableProjectView) Get() *AvailableProjectView {
	return v.value
}

func (v *NullableAvailableProjectView) Set(val *AvailableProjectView) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableProjectView) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableProjectView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableProjectView(val *AvailableProjectView) *NullableAvailableProjectView {
	return &NullableAvailableProjectView{value: val, isSet: true}
}

func (v NullableAvailableProjectView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableProjectView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

