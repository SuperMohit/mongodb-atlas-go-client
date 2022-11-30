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

// ApiApiUserView struct for ApiApiUserView
type ApiApiUserView struct {
	// Purpose or explanation provided when someone created this organization API key.
	Desc *string `json:"desc,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this organization API key assigned to this project.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
	// Redacted private key returned for this organization API key. This key displays unredacted when first created.
	PrivateKey *string `json:"privateKey,omitempty"`
	// Public API key value set for the specified organization API key.
	PublicKey *string `json:"publicKey,omitempty"`
	// List that contains the roles that the API key needs to have. All roles you provide must be valid for the specified project or organization. Each request must include a minimum of one valid role. The resource returns all project and organization roles assigned to the API key.
	Roles []ApiRoleAssignmentView `json:"roles,omitempty"`
}

// NewApiApiUserView instantiates a new ApiApiUserView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiApiUserView(links []Link) *ApiApiUserView {
	this := ApiApiUserView{}
	this.Links = links
	return &this
}

// NewApiApiUserViewWithDefaults instantiates a new ApiApiUserView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiApiUserViewWithDefaults() *ApiApiUserView {
	this := ApiApiUserView{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *ApiApiUserView) GetDesc() string {
	if o == nil || isNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApiUserView) GetDescOk() (*string, bool) {
	if o == nil || isNil(o.Desc) {
    return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *ApiApiUserView) HasDesc() bool {
	if o != nil && !isNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *ApiApiUserView) SetDesc(v string) {
	o.Desc = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiApiUserView) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApiUserView) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiApiUserView) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiApiUserView) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value
func (o *ApiApiUserView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiApiUserView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ApiApiUserView) SetLinks(v []Link) {
	o.Links = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *ApiApiUserView) GetPrivateKey() string {
	if o == nil || isNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApiUserView) GetPrivateKeyOk() (*string, bool) {
	if o == nil || isNil(o.PrivateKey) {
    return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *ApiApiUserView) HasPrivateKey() bool {
	if o != nil && !isNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *ApiApiUserView) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *ApiApiUserView) GetPublicKey() string {
	if o == nil || isNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApiUserView) GetPublicKeyOk() (*string, bool) {
	if o == nil || isNil(o.PublicKey) {
    return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *ApiApiUserView) HasPublicKey() bool {
	if o != nil && !isNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *ApiApiUserView) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApiApiUserView) GetRoles() []ApiRoleAssignmentView {
	if o == nil || isNil(o.Roles) {
		var ret []ApiRoleAssignmentView
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApiUserView) GetRolesOk() ([]ApiRoleAssignmentView, bool) {
	if o == nil || isNil(o.Roles) {
    return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApiApiUserView) HasRoles() bool {
	if o != nil && !isNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []ApiRoleAssignmentView and assigns it to the Roles field.
func (o *ApiApiUserView) SetRoles(v []ApiRoleAssignmentView) {
	o.Roles = v
}

func (o ApiApiUserView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if !isNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	if !isNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableApiApiUserView struct {
	value *ApiApiUserView
	isSet bool
}

func (v NullableApiApiUserView) Get() *ApiApiUserView {
	return v.value
}

func (v *NullableApiApiUserView) Set(val *ApiApiUserView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiApiUserView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiApiUserView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiApiUserView(val *ApiApiUserView) *NullableApiApiUserView {
	return &NullableApiApiUserView{value: val, isSet: true}
}

func (v NullableApiApiUserView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiApiUserView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


