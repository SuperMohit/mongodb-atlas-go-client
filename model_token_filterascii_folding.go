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

// TokenFilterasciiFolding Filter that converts alphabetic, numeric, and symbolic unicode characters that are not in the Basic Latin Unicode block to their ASCII equivalents, if available.
type TokenFilterasciiFolding struct {
	// Value that indicates whether to include or omit the original tokens in the output of the token filter.  Choose `include` if you want to support queries on both the original tokens as well as the converted forms.   Choose `omit` if you want to query only on the converted forms of the original tokens.
	OriginalTokens *string `json:"originalTokens,omitempty"`
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFilterasciiFolding instantiates a new TokenFilterasciiFolding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFilterasciiFolding(type_ string) *TokenFilterasciiFolding {
	this := TokenFilterasciiFolding{}
	var originalTokens string = "omit"
	this.OriginalTokens = &originalTokens
	this.Type = type_
	return &this
}

// NewTokenFilterasciiFoldingWithDefaults instantiates a new TokenFilterasciiFolding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFilterasciiFoldingWithDefaults() *TokenFilterasciiFolding {
	this := TokenFilterasciiFolding{}
	var originalTokens string = "omit"
	this.OriginalTokens = &originalTokens
	return &this
}

// GetOriginalTokens returns the OriginalTokens field value if set, zero value otherwise.
func (o *TokenFilterasciiFolding) GetOriginalTokens() string {
	if o == nil || isNil(o.OriginalTokens) {
		var ret string
		return ret
	}
	return *o.OriginalTokens
}

// GetOriginalTokensOk returns a tuple with the OriginalTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenFilterasciiFolding) GetOriginalTokensOk() (*string, bool) {
	if o == nil || isNil(o.OriginalTokens) {
    return nil, false
	}
	return o.OriginalTokens, true
}

// HasOriginalTokens returns a boolean if a field has been set.
func (o *TokenFilterasciiFolding) HasOriginalTokens() bool {
	if o != nil && !isNil(o.OriginalTokens) {
		return true
	}

	return false
}

// SetOriginalTokens gets a reference to the given string and assigns it to the OriginalTokens field.
func (o *TokenFilterasciiFolding) SetOriginalTokens(v string) {
	o.OriginalTokens = &v
}

// GetType returns the Type field value
func (o *TokenFilterasciiFolding) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFilterasciiFolding) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFilterasciiFolding) SetType(v string) {
	o.Type = v
}

func (o TokenFilterasciiFolding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OriginalTokens) {
		toSerialize["originalTokens"] = o.OriginalTokens
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTokenFilterasciiFolding struct {
	value *TokenFilterasciiFolding
	isSet bool
}

func (v NullableTokenFilterasciiFolding) Get() *TokenFilterasciiFolding {
	return v.value
}

func (v *NullableTokenFilterasciiFolding) Set(val *TokenFilterasciiFolding) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFilterasciiFolding) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFilterasciiFolding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFilterasciiFolding(val *TokenFilterasciiFolding) *NullableTokenFilterasciiFolding {
	return &NullableTokenFilterasciiFolding{value: val, isSet: true}
}

func (v NullableTokenFilterasciiFolding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFilterasciiFolding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


