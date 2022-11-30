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

// ApiDatabaseUserSCRAMViewAllOf struct for ApiDatabaseUserSCRAMViewAllOf
type ApiDatabaseUserSCRAMViewAllOf struct {
	// Alphanumeric string that authenticates this database user against the database specified in `databaseName`. To authenticate with SCRAM-SHA, you must specify this parameter.
	Password *string `json:"password,omitempty"`
	// Human-readable label that represents the user that authenticates to MongoDB. This must be formatted as a [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name. 
	Username *string `json:"username,omitempty"`
}

// NewApiDatabaseUserSCRAMViewAllOf instantiates a new ApiDatabaseUserSCRAMViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDatabaseUserSCRAMViewAllOf() *ApiDatabaseUserSCRAMViewAllOf {
	this := ApiDatabaseUserSCRAMViewAllOf{}
	return &this
}

// NewApiDatabaseUserSCRAMViewAllOfWithDefaults instantiates a new ApiDatabaseUserSCRAMViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDatabaseUserSCRAMViewAllOfWithDefaults() *ApiDatabaseUserSCRAMViewAllOf {
	this := ApiDatabaseUserSCRAMViewAllOf{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiDatabaseUserSCRAMViewAllOf) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserSCRAMViewAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiDatabaseUserSCRAMViewAllOf) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApiDatabaseUserSCRAMViewAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiDatabaseUserSCRAMViewAllOf) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDatabaseUserSCRAMViewAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiDatabaseUserSCRAMViewAllOf) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiDatabaseUserSCRAMViewAllOf) SetUsername(v string) {
	o.Username = &v
}

func (o ApiDatabaseUserSCRAMViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableApiDatabaseUserSCRAMViewAllOf struct {
	value *ApiDatabaseUserSCRAMViewAllOf
	isSet bool
}

func (v NullableApiDatabaseUserSCRAMViewAllOf) Get() *ApiDatabaseUserSCRAMViewAllOf {
	return v.value
}

func (v *NullableApiDatabaseUserSCRAMViewAllOf) Set(val *ApiDatabaseUserSCRAMViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDatabaseUserSCRAMViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDatabaseUserSCRAMViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDatabaseUserSCRAMViewAllOf(val *ApiDatabaseUserSCRAMViewAllOf) *NullableApiDatabaseUserSCRAMViewAllOf {
	return &NullableApiDatabaseUserSCRAMViewAllOf{value: val, isSet: true}
}

func (v NullableApiDatabaseUserSCRAMViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDatabaseUserSCRAMViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


