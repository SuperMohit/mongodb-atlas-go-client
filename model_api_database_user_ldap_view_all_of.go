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

// ApiDatabaseUserLDAPViewAllOf struct for ApiDatabaseUserLDAPViewAllOf
type ApiDatabaseUserLDAPViewAllOf struct {
	// Part of the Lightweight Directory Access Protocol (LDAP) record that the database uses to authenticate this database user on the LDAP host.
	LdapAuthType *string `json:"ldapAuthType,omitempty"`
	// Human-readable label that represents the user that authenticates to MongoDB. This must be formatted as a [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name. 
	Username *string `json:"username,omitempty"`
}

// NewApiDatabaseUserLDAPViewAllOf instantiates a new ApiDatabaseUserLDAPViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDatabaseUserLDAPViewAllOf() *ApiDatabaseUserLDAPViewAllOf {
	this := ApiDatabaseUserLDAPViewAllOf{}
	var ldapAuthType string = "NONE"
	this.LdapAuthType = &ldapAuthType
	return &this
}

// NewApiDatabaseUserLDAPViewAllOfWithDefaults instantiates a new ApiDatabaseUserLDAPViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDatabaseUserLDAPViewAllOfWithDefaults() *ApiDatabaseUserLDAPViewAllOf {
	this := ApiDatabaseUserLDAPViewAllOf{}
	var ldapAuthType string = "NONE"
	this.LdapAuthType = &ldapAuthType
	return &this
}

// GetLdapAuthType returns the LdapAuthType field value if set, zero value otherwise.
func (o *ApiDatabaseUserLDAPViewAllOf) GetLdapAuthType() string {
	if o == nil || isNil(o.LdapAuthType) {
		var ret string
		return ret
	}
	return *o.LdapAuthType
}

// GetLdapAuthTypeOk returns a tuple with the LdapAuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserLDAPViewAllOf) GetLdapAuthTypeOk() (*string, bool) {
	if o == nil || isNil(o.LdapAuthType) {
    return nil, false
	}
	return o.LdapAuthType, true
}

// HasLdapAuthType returns a boolean if a field has been set.
func (o *ApiDatabaseUserLDAPViewAllOf) HasLdapAuthType() bool {
	if o != nil && !isNil(o.LdapAuthType) {
		return true
	}

	return false
}

// SetLdapAuthType gets a reference to the given string and assigns it to the LdapAuthType field.
func (o *ApiDatabaseUserLDAPViewAllOf) SetLdapAuthType(v string) {
	o.LdapAuthType = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiDatabaseUserLDAPViewAllOf) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserLDAPViewAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiDatabaseUserLDAPViewAllOf) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiDatabaseUserLDAPViewAllOf) SetUsername(v string) {
	o.Username = &v
}

func (o ApiDatabaseUserLDAPViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LdapAuthType) {
		toSerialize["ldapAuthType"] = o.LdapAuthType
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableApiDatabaseUserLDAPViewAllOf struct {
	value *ApiDatabaseUserLDAPViewAllOf
	isSet bool
}

func (v NullableApiDatabaseUserLDAPViewAllOf) Get() *ApiDatabaseUserLDAPViewAllOf {
	return v.value
}

func (v *NullableApiDatabaseUserLDAPViewAllOf) Set(val *ApiDatabaseUserLDAPViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDatabaseUserLDAPViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDatabaseUserLDAPViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDatabaseUserLDAPViewAllOf(val *ApiDatabaseUserLDAPViewAllOf) *NullableApiDatabaseUserLDAPViewAllOf {
	return &NullableApiDatabaseUserLDAPViewAllOf{value: val, isSet: true}
}

func (v NullableApiDatabaseUserLDAPViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDatabaseUserLDAPViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


