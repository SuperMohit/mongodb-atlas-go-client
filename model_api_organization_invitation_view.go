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

// ApiOrganizationInvitationView struct for ApiOrganizationInvitationView
type ApiOrganizationInvitationView struct {
	// Date and time when MongoDB Cloud sent the invitation. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Date and time when the invitation from MongoDB Cloud expires. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this organization.
	Id *string `json:"id,omitempty"`
	// Email address of the MongoDB Cloud user who sent the invitation to join the organization.
	InviterUsername *string `json:"inviterUsername,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// Unique 24-hexadecimal digit string that identifies the organization.
	OrgId *string `json:"orgId,omitempty"`
	// Human-readable label that identifies this organization.
	OrgName string `json:"orgName"`
	// One or more organization or project level roles to assign to the MongoDB Cloud user.
	Roles []string `json:"roles,omitempty"`
	// List of unique 24-hexadecimal digit strings that identifies each team.
	TeamIds []string `json:"teamIds,omitempty"`
	// Email address of the MongoDB Cloud user invited to join the organization.
	Username *string `json:"username,omitempty"`
}

// NewApiOrganizationInvitationView instantiates a new ApiOrganizationInvitationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiOrganizationInvitationView(links []Link, orgName string) *ApiOrganizationInvitationView {
	this := ApiOrganizationInvitationView{}
	this.Links = links
	this.OrgName = orgName
	return &this
}

// NewApiOrganizationInvitationViewWithDefaults instantiates a new ApiOrganizationInvitationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiOrganizationInvitationViewWithDefaults() *ApiOrganizationInvitationView {
	this := ApiOrganizationInvitationView{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApiOrganizationInvitationView) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationInvitationView) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApiOrganizationInvitationView) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ApiOrganizationInvitationView) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ApiOrganizationInvitationView) GetExpiresAt() time.Time {
	if o == nil || isNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationInvitationView) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiresAt) {
    return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ApiOrganizationInvitationView) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *ApiOrganizationInvitationView) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiOrganizationInvitationView) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationInvitationView) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiOrganizationInvitationView) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiOrganizationInvitationView) SetId(v string) {
	o.Id = &v
}

// GetInviterUsername returns the InviterUsername field value if set, zero value otherwise.
func (o *ApiOrganizationInvitationView) GetInviterUsername() string {
	if o == nil || isNil(o.InviterUsername) {
		var ret string
		return ret
	}
	return *o.InviterUsername
}

// GetInviterUsernameOk returns a tuple with the InviterUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationInvitationView) GetInviterUsernameOk() (*string, bool) {
	if o == nil || isNil(o.InviterUsername) {
    return nil, false
	}
	return o.InviterUsername, true
}

// HasInviterUsername returns a boolean if a field has been set.
func (o *ApiOrganizationInvitationView) HasInviterUsername() bool {
	if o != nil && !isNil(o.InviterUsername) {
		return true
	}

	return false
}

// SetInviterUsername gets a reference to the given string and assigns it to the InviterUsername field.
func (o *ApiOrganizationInvitationView) SetInviterUsername(v string) {
	o.InviterUsername = &v
}

// GetLinks returns the Links field value
func (o *ApiOrganizationInvitationView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiOrganizationInvitationView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ApiOrganizationInvitationView) SetLinks(v []Link) {
	o.Links = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ApiOrganizationInvitationView) GetOrgId() string {
	if o == nil || isNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationInvitationView) GetOrgIdOk() (*string, bool) {
	if o == nil || isNil(o.OrgId) {
    return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ApiOrganizationInvitationView) HasOrgId() bool {
	if o != nil && !isNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ApiOrganizationInvitationView) SetOrgId(v string) {
	o.OrgId = &v
}

// GetOrgName returns the OrgName field value
func (o *ApiOrganizationInvitationView) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *ApiOrganizationInvitationView) GetOrgNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *ApiOrganizationInvitationView) SetOrgName(v string) {
	o.OrgName = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApiOrganizationInvitationView) GetRoles() []string {
	if o == nil || isNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationInvitationView) GetRolesOk() ([]string, bool) {
	if o == nil || isNil(o.Roles) {
    return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApiOrganizationInvitationView) HasRoles() bool {
	if o != nil && !isNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *ApiOrganizationInvitationView) SetRoles(v []string) {
	o.Roles = v
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise.
func (o *ApiOrganizationInvitationView) GetTeamIds() []string {
	if o == nil || isNil(o.TeamIds) {
		var ret []string
		return ret
	}
	return o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationInvitationView) GetTeamIdsOk() ([]string, bool) {
	if o == nil || isNil(o.TeamIds) {
    return nil, false
	}
	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *ApiOrganizationInvitationView) HasTeamIds() bool {
	if o != nil && !isNil(o.TeamIds) {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []string and assigns it to the TeamIds field.
func (o *ApiOrganizationInvitationView) SetTeamIds(v []string) {
	o.TeamIds = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiOrganizationInvitationView) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOrganizationInvitationView) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiOrganizationInvitationView) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiOrganizationInvitationView) SetUsername(v string) {
	o.Username = &v
}

func (o ApiOrganizationInvitationView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.InviterUsername) {
		toSerialize["inviterUsername"] = o.InviterUsername
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if true {
		toSerialize["orgName"] = o.OrgName
	}
	if !isNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !isNil(o.TeamIds) {
		toSerialize["teamIds"] = o.TeamIds
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableApiOrganizationInvitationView struct {
	value *ApiOrganizationInvitationView
	isSet bool
}

func (v NullableApiOrganizationInvitationView) Get() *ApiOrganizationInvitationView {
	return v.value
}

func (v *NullableApiOrganizationInvitationView) Set(val *ApiOrganizationInvitationView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiOrganizationInvitationView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiOrganizationInvitationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiOrganizationInvitationView(val *ApiOrganizationInvitationView) *NullableApiOrganizationInvitationView {
	return &NullableApiOrganizationInvitationView{value: val, isSet: true}
}

func (v NullableApiOrganizationInvitationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiOrganizationInvitationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

