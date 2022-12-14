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

// ApiAtlasServerlessTenantEndpointCreateView struct for ApiAtlasServerlessTenantEndpointCreateView
type ApiAtlasServerlessTenantEndpointCreateView struct {
	// Human-readable comment associated with the private endpoint.
	Comment *string `json:"comment,omitempty"`
}

// NewApiAtlasServerlessTenantEndpointCreateView instantiates a new ApiAtlasServerlessTenantEndpointCreateView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasServerlessTenantEndpointCreateView() *ApiAtlasServerlessTenantEndpointCreateView {
	this := ApiAtlasServerlessTenantEndpointCreateView{}
	return &this
}

// NewApiAtlasServerlessTenantEndpointCreateViewWithDefaults instantiates a new ApiAtlasServerlessTenantEndpointCreateView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasServerlessTenantEndpointCreateViewWithDefaults() *ApiAtlasServerlessTenantEndpointCreateView {
	this := ApiAtlasServerlessTenantEndpointCreateView{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApiAtlasServerlessTenantEndpointCreateView) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessTenantEndpointCreateView) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApiAtlasServerlessTenantEndpointCreateView) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApiAtlasServerlessTenantEndpointCreateView) SetComment(v string) {
	o.Comment = &v
}

func (o ApiAtlasServerlessTenantEndpointCreateView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasServerlessTenantEndpointCreateView struct {
	value *ApiAtlasServerlessTenantEndpointCreateView
	isSet bool
}

func (v NullableApiAtlasServerlessTenantEndpointCreateView) Get() *ApiAtlasServerlessTenantEndpointCreateView {
	return v.value
}

func (v *NullableApiAtlasServerlessTenantEndpointCreateView) Set(val *ApiAtlasServerlessTenantEndpointCreateView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasServerlessTenantEndpointCreateView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasServerlessTenantEndpointCreateView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasServerlessTenantEndpointCreateView(val *ApiAtlasServerlessTenantEndpointCreateView) *NullableApiAtlasServerlessTenantEndpointCreateView {
	return &NullableApiAtlasServerlessTenantEndpointCreateView{value: val, isSet: true}
}

func (v NullableApiAtlasServerlessTenantEndpointCreateView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasServerlessTenantEndpointCreateView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


