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

// ApiDatabaseUserSCRAMView Details to create one new MongoDB database user who authenticates with a MongoDB database using a Lightweight Directory Access Protocol (LDAP) host. 
type ApiDatabaseUserSCRAMView struct {
	// MongoDB database against which the MongoDB database user authenticates. MongoDB database users must provide both a username and authentication database to log into MongoDB.
	DatabaseName string `json:"databaseName"`
	// Date and time when MongoDB Cloud deletes the user. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request.
	DeleteAfterDate *string `json:"deleteAfterDate,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project.
	GroupId string `json:"groupId"`
	// List that contains the key-value pairs for tagging and categorizing the MongoDB database user. The labels that you define do not appear in the console.
	Labels []ApiAtlasNDSLabelView `json:"labels,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// List that provides the pairings of one role with one applicable database.
	Roles []ApiAtlasRoleView `json:"roles"`
	// List that contains clusters and MongoDB Atlas Data Lakes that this database user can access. If omitted, MongoDB Cloud grants the database user access to all the clusters and MongoDB Atlas Data Lakes in the project.
	Scopes []ApiAtlasUserScopeView `json:"scopes,omitempty"`
	// Alphanumeric string that authenticates this database user against the database specified in `databaseName`. To authenticate with SCRAM-SHA, you must specify this parameter.
	Password string `json:"password"`
	// Human-readable label that represents the user that authenticates to MongoDB. This must be formatted as a [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name. 
	Username string `json:"username"`
}

// NewApiDatabaseUserSCRAMView instantiates a new ApiDatabaseUserSCRAMView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDatabaseUserSCRAMView(databaseName string, groupId string, roles []ApiAtlasRoleView, password string, username string) *ApiDatabaseUserSCRAMView {
	this := ApiDatabaseUserSCRAMView{}
	this.DatabaseName = databaseName
	this.GroupId = groupId
	this.Roles = roles
	this.Password = password
	this.Username = username
	return &this
}

// NewApiDatabaseUserSCRAMViewWithDefaults instantiates a new ApiDatabaseUserSCRAMView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDatabaseUserSCRAMViewWithDefaults() *ApiDatabaseUserSCRAMView {
	this := ApiDatabaseUserSCRAMView{}
	var databaseName string = "admin"
	this.DatabaseName = databaseName
	return &this
}

// GetDatabaseName returns the DatabaseName field value
func (o *ApiDatabaseUserSCRAMView) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserSCRAMView) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value
func (o *ApiDatabaseUserSCRAMView) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetDeleteAfterDate returns the DeleteAfterDate field value if set, zero value otherwise.
func (o *ApiDatabaseUserSCRAMView) GetDeleteAfterDate() string {
	if o == nil || isNil(o.DeleteAfterDate) {
		var ret string
		return ret
	}
	return *o.DeleteAfterDate
}

// GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserSCRAMView) GetDeleteAfterDateOk() (*string, bool) {
	if o == nil || isNil(o.DeleteAfterDate) {
    return nil, false
	}
	return o.DeleteAfterDate, true
}

// HasDeleteAfterDate returns a boolean if a field has been set.
func (o *ApiDatabaseUserSCRAMView) HasDeleteAfterDate() bool {
	if o != nil && !isNil(o.DeleteAfterDate) {
		return true
	}

	return false
}

// SetDeleteAfterDate gets a reference to the given string and assigns it to the DeleteAfterDate field.
func (o *ApiDatabaseUserSCRAMView) SetDeleteAfterDate(v string) {
	o.DeleteAfterDate = &v
}

// GetGroupId returns the GroupId field value
func (o *ApiDatabaseUserSCRAMView) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserSCRAMView) GetGroupIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ApiDatabaseUserSCRAMView) SetGroupId(v string) {
	o.GroupId = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ApiDatabaseUserSCRAMView) GetLabels() []ApiAtlasNDSLabelView {
	if o == nil || isNil(o.Labels) {
		var ret []ApiAtlasNDSLabelView
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserSCRAMView) GetLabelsOk() ([]ApiAtlasNDSLabelView, bool) {
	if o == nil || isNil(o.Labels) {
    return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ApiDatabaseUserSCRAMView) HasLabels() bool {
	if o != nil && !isNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []ApiAtlasNDSLabelView and assigns it to the Labels field.
func (o *ApiDatabaseUserSCRAMView) SetLabels(v []ApiAtlasNDSLabelView) {
	o.Labels = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiDatabaseUserSCRAMView) GetLinks() []Link {
	if o == nil || isNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserSCRAMView) GetLinksOk() ([]Link, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiDatabaseUserSCRAMView) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ApiDatabaseUserSCRAMView) SetLinks(v []Link) {
	o.Links = v
}

// GetRoles returns the Roles field value
func (o *ApiDatabaseUserSCRAMView) GetRoles() []ApiAtlasRoleView {
	if o == nil {
		var ret []ApiAtlasRoleView
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserSCRAMView) GetRolesOk() ([]ApiAtlasRoleView, bool) {
	if o == nil {
    return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *ApiDatabaseUserSCRAMView) SetRoles(v []ApiAtlasRoleView) {
	o.Roles = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ApiDatabaseUserSCRAMView) GetScopes() []ApiAtlasUserScopeView {
	if o == nil || isNil(o.Scopes) {
		var ret []ApiAtlasUserScopeView
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserSCRAMView) GetScopesOk() ([]ApiAtlasUserScopeView, bool) {
	if o == nil || isNil(o.Scopes) {
    return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ApiDatabaseUserSCRAMView) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []ApiAtlasUserScopeView and assigns it to the Scopes field.
func (o *ApiDatabaseUserSCRAMView) SetScopes(v []ApiAtlasUserScopeView) {
	o.Scopes = v
}

// GetPassword returns the Password field value
func (o *ApiDatabaseUserSCRAMView) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserSCRAMView) GetPasswordOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ApiDatabaseUserSCRAMView) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *ApiDatabaseUserSCRAMView) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserSCRAMView) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ApiDatabaseUserSCRAMView) SetUsername(v string) {
	o.Username = v
}

func (o ApiDatabaseUserSCRAMView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if !isNil(o.DeleteAfterDate) {
		toSerialize["deleteAfterDate"] = o.DeleteAfterDate
	}
	if true {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !isNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableApiDatabaseUserSCRAMView struct {
	value *ApiDatabaseUserSCRAMView
	isSet bool
}

func (v NullableApiDatabaseUserSCRAMView) Get() *ApiDatabaseUserSCRAMView {
	return v.value
}

func (v *NullableApiDatabaseUserSCRAMView) Set(val *ApiDatabaseUserSCRAMView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDatabaseUserSCRAMView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDatabaseUserSCRAMView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDatabaseUserSCRAMView(val *ApiDatabaseUserSCRAMView) *NullableApiDatabaseUserSCRAMView {
	return &NullableApiDatabaseUserSCRAMView{value: val, isSet: true}
}

func (v NullableApiDatabaseUserSCRAMView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDatabaseUserSCRAMView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


