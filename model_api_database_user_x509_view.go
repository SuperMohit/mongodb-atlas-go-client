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

// ApiDatabaseUserX509View Details to create one new MongoDB database user who authenticates with a MongoDB database using X.509 certificates. 
type ApiDatabaseUserX509View struct {
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
	// Human-readable label that represents the user that authenticates to MongoDB. This must be formatted as a [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name. 
	Username string `json:"username"`
	// X.509 method that MongoDB Cloud uses to authenticate the database user. - For MongoDB Cloud-managed X.509, specify `MANAGED`. - For self-managed X.509, specify `CUSTOMER`. Users created with the `CUSTOMER` method require a Common Name (CN) in the **username** parameter. You must create externally authenticated users on the `$external` database.
	X509Type string `json:"x509Type"`
}

// NewApiDatabaseUserX509View instantiates a new ApiDatabaseUserX509View object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDatabaseUserX509View(databaseName string, groupId string, roles []ApiAtlasRoleView, username string, x509Type string) *ApiDatabaseUserX509View {
	this := ApiDatabaseUserX509View{}
	this.DatabaseName = databaseName
	this.GroupId = groupId
	this.Roles = roles
	this.Username = username
	this.X509Type = x509Type
	return &this
}

// NewApiDatabaseUserX509ViewWithDefaults instantiates a new ApiDatabaseUserX509View object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDatabaseUserX509ViewWithDefaults() *ApiDatabaseUserX509View {
	this := ApiDatabaseUserX509View{}
	var databaseName string = "admin"
	this.DatabaseName = databaseName
	var x509Type string = "NONE"
	this.X509Type = x509Type
	return &this
}

// GetDatabaseName returns the DatabaseName field value
func (o *ApiDatabaseUserX509View) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserX509View) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value
func (o *ApiDatabaseUserX509View) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetDeleteAfterDate returns the DeleteAfterDate field value if set, zero value otherwise.
func (o *ApiDatabaseUserX509View) GetDeleteAfterDate() string {
	if o == nil || isNil(o.DeleteAfterDate) {
		var ret string
		return ret
	}
	return *o.DeleteAfterDate
}

// GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserX509View) GetDeleteAfterDateOk() (*string, bool) {
	if o == nil || isNil(o.DeleteAfterDate) {
    return nil, false
	}
	return o.DeleteAfterDate, true
}

// HasDeleteAfterDate returns a boolean if a field has been set.
func (o *ApiDatabaseUserX509View) HasDeleteAfterDate() bool {
	if o != nil && !isNil(o.DeleteAfterDate) {
		return true
	}

	return false
}

// SetDeleteAfterDate gets a reference to the given string and assigns it to the DeleteAfterDate field.
func (o *ApiDatabaseUserX509View) SetDeleteAfterDate(v string) {
	o.DeleteAfterDate = &v
}

// GetGroupId returns the GroupId field value
func (o *ApiDatabaseUserX509View) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserX509View) GetGroupIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ApiDatabaseUserX509View) SetGroupId(v string) {
	o.GroupId = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ApiDatabaseUserX509View) GetLabels() []ApiAtlasNDSLabelView {
	if o == nil || isNil(o.Labels) {
		var ret []ApiAtlasNDSLabelView
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserX509View) GetLabelsOk() ([]ApiAtlasNDSLabelView, bool) {
	if o == nil || isNil(o.Labels) {
    return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ApiDatabaseUserX509View) HasLabels() bool {
	if o != nil && !isNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []ApiAtlasNDSLabelView and assigns it to the Labels field.
func (o *ApiDatabaseUserX509View) SetLabels(v []ApiAtlasNDSLabelView) {
	o.Labels = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiDatabaseUserX509View) GetLinks() []Link {
	if o == nil || isNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserX509View) GetLinksOk() ([]Link, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiDatabaseUserX509View) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ApiDatabaseUserX509View) SetLinks(v []Link) {
	o.Links = v
}

// GetRoles returns the Roles field value
func (o *ApiDatabaseUserX509View) GetRoles() []ApiAtlasRoleView {
	if o == nil {
		var ret []ApiAtlasRoleView
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserX509View) GetRolesOk() ([]ApiAtlasRoleView, bool) {
	if o == nil {
    return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *ApiDatabaseUserX509View) SetRoles(v []ApiAtlasRoleView) {
	o.Roles = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ApiDatabaseUserX509View) GetScopes() []ApiAtlasUserScopeView {
	if o == nil || isNil(o.Scopes) {
		var ret []ApiAtlasUserScopeView
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserX509View) GetScopesOk() ([]ApiAtlasUserScopeView, bool) {
	if o == nil || isNil(o.Scopes) {
    return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ApiDatabaseUserX509View) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []ApiAtlasUserScopeView and assigns it to the Scopes field.
func (o *ApiDatabaseUserX509View) SetScopes(v []ApiAtlasUserScopeView) {
	o.Scopes = v
}

// GetUsername returns the Username field value
func (o *ApiDatabaseUserX509View) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserX509View) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ApiDatabaseUserX509View) SetUsername(v string) {
	o.Username = v
}

// GetX509Type returns the X509Type field value
func (o *ApiDatabaseUserX509View) GetX509Type() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X509Type
}

// GetX509TypeOk returns a tuple with the X509Type field value
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserX509View) GetX509TypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.X509Type, true
}

// SetX509Type sets field value
func (o *ApiDatabaseUserX509View) SetX509Type(v string) {
	o.X509Type = v
}

func (o ApiDatabaseUserX509View) MarshalJSON() ([]byte, error) {
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
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["x509Type"] = o.X509Type
	}
	return json.Marshal(toSerialize)
}

type NullableApiDatabaseUserX509View struct {
	value *ApiDatabaseUserX509View
	isSet bool
}

func (v NullableApiDatabaseUserX509View) Get() *ApiDatabaseUserX509View {
	return v.value
}

func (v *NullableApiDatabaseUserX509View) Set(val *ApiDatabaseUserX509View) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDatabaseUserX509View) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDatabaseUserX509View) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDatabaseUserX509View(val *ApiDatabaseUserX509View) *NullableApiDatabaseUserX509View {
	return &NullableApiDatabaseUserX509View{value: val, isSet: true}
}

func (v NullableApiDatabaseUserX509View) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDatabaseUserX509View) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


