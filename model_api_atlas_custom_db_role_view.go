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

// ApiAtlasCustomDBRoleView struct for ApiAtlasCustomDBRoleView
type ApiAtlasCustomDBRoleView struct {
	// List of the individual privilege actions that the role grants.
	Actions []ApiAtlasDBActionView `json:"actions,omitempty"`
	// List of the built-in roles that this custom role inherits.
	InheritedRoles []ApiAtlasInheritedRoleView `json:"inheritedRoles,omitempty"`
	// Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.
	RoleName string `json:"roleName"`
}

// NewApiAtlasCustomDBRoleView instantiates a new ApiAtlasCustomDBRoleView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCustomDBRoleView(roleName string) *ApiAtlasCustomDBRoleView {
	this := ApiAtlasCustomDBRoleView{}
	this.RoleName = roleName
	return &this
}

// NewApiAtlasCustomDBRoleViewWithDefaults instantiates a new ApiAtlasCustomDBRoleView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCustomDBRoleViewWithDefaults() *ApiAtlasCustomDBRoleView {
	this := ApiAtlasCustomDBRoleView{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ApiAtlasCustomDBRoleView) GetActions() []ApiAtlasDBActionView {
	if o == nil || isNil(o.Actions) {
		var ret []ApiAtlasDBActionView
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCustomDBRoleView) GetActionsOk() ([]ApiAtlasDBActionView, bool) {
	if o == nil || isNil(o.Actions) {
    return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ApiAtlasCustomDBRoleView) HasActions() bool {
	if o != nil && !isNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []ApiAtlasDBActionView and assigns it to the Actions field.
func (o *ApiAtlasCustomDBRoleView) SetActions(v []ApiAtlasDBActionView) {
	o.Actions = v
}

// GetInheritedRoles returns the InheritedRoles field value if set, zero value otherwise.
func (o *ApiAtlasCustomDBRoleView) GetInheritedRoles() []ApiAtlasInheritedRoleView {
	if o == nil || isNil(o.InheritedRoles) {
		var ret []ApiAtlasInheritedRoleView
		return ret
	}
	return o.InheritedRoles
}

// GetInheritedRolesOk returns a tuple with the InheritedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCustomDBRoleView) GetInheritedRolesOk() ([]ApiAtlasInheritedRoleView, bool) {
	if o == nil || isNil(o.InheritedRoles) {
    return nil, false
	}
	return o.InheritedRoles, true
}

// HasInheritedRoles returns a boolean if a field has been set.
func (o *ApiAtlasCustomDBRoleView) HasInheritedRoles() bool {
	if o != nil && !isNil(o.InheritedRoles) {
		return true
	}

	return false
}

// SetInheritedRoles gets a reference to the given []ApiAtlasInheritedRoleView and assigns it to the InheritedRoles field.
func (o *ApiAtlasCustomDBRoleView) SetInheritedRoles(v []ApiAtlasInheritedRoleView) {
	o.InheritedRoles = v
}

// GetRoleName returns the RoleName field value
func (o *ApiAtlasCustomDBRoleView) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasCustomDBRoleView) GetRoleNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *ApiAtlasCustomDBRoleView) SetRoleName(v string) {
	o.RoleName = v
}

func (o ApiAtlasCustomDBRoleView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !isNil(o.InheritedRoles) {
		toSerialize["inheritedRoles"] = o.InheritedRoles
	}
	if true {
		toSerialize["roleName"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasCustomDBRoleView struct {
	value *ApiAtlasCustomDBRoleView
	isSet bool
}

func (v NullableApiAtlasCustomDBRoleView) Get() *ApiAtlasCustomDBRoleView {
	return v.value
}

func (v *NullableApiAtlasCustomDBRoleView) Set(val *ApiAtlasCustomDBRoleView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasCustomDBRoleView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasCustomDBRoleView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasCustomDBRoleView(val *ApiAtlasCustomDBRoleView) *NullableApiAtlasCustomDBRoleView {
	return &NullableApiAtlasCustomDBRoleView{value: val, isSet: true}
}

func (v NullableApiAtlasCustomDBRoleView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasCustomDBRoleView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

