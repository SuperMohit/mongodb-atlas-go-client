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

// TokenFilternGram Filter that tokenizes input into n-grams of configured sizes. You can't use this token filter in synonym or autocomplete mapping definitions.
type TokenFilternGram struct {
	// Value that specifies the maximum length of generated n-grams. This value must be greater than or equal to **minGram**.
	MaxGram int32 `json:"maxGram"`
	// Value that specifies the minimum length of generated n-grams. This value must be less than or equal to **maxGram**.
	MinGram int32 `json:"minGram"`
	// Value that indicates whether to index tokens shorter than **minGram** or longer than **maxGram**.
	TermNotInBounds *string `json:"termNotInBounds,omitempty"`
	// Human-readable label that identifies this token filter type.
	Type string `json:"type"`
}

// NewTokenFilternGram instantiates a new TokenFilternGram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFilternGram(maxGram int32, minGram int32, type_ string) *TokenFilternGram {
	this := TokenFilternGram{}
	this.MaxGram = maxGram
	this.MinGram = minGram
	var termNotInBounds string = "omit"
	this.TermNotInBounds = &termNotInBounds
	this.Type = type_
	return &this
}

// NewTokenFilternGramWithDefaults instantiates a new TokenFilternGram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFilternGramWithDefaults() *TokenFilternGram {
	this := TokenFilternGram{}
	var termNotInBounds string = "omit"
	this.TermNotInBounds = &termNotInBounds
	return &this
}

// GetMaxGram returns the MaxGram field value
func (o *TokenFilternGram) GetMaxGram() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxGram
}

// GetMaxGramOk returns a tuple with the MaxGram field value
// and a boolean to check if the value has been set.
func (o *TokenFilternGram) GetMaxGramOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MaxGram, true
}

// SetMaxGram sets field value
func (o *TokenFilternGram) SetMaxGram(v int32) {
	o.MaxGram = v
}

// GetMinGram returns the MinGram field value
func (o *TokenFilternGram) GetMinGram() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinGram
}

// GetMinGramOk returns a tuple with the MinGram field value
// and a boolean to check if the value has been set.
func (o *TokenFilternGram) GetMinGramOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MinGram, true
}

// SetMinGram sets field value
func (o *TokenFilternGram) SetMinGram(v int32) {
	o.MinGram = v
}

// GetTermNotInBounds returns the TermNotInBounds field value if set, zero value otherwise.
func (o *TokenFilternGram) GetTermNotInBounds() string {
	if o == nil || isNil(o.TermNotInBounds) {
		var ret string
		return ret
	}
	return *o.TermNotInBounds
}

// GetTermNotInBoundsOk returns a tuple with the TermNotInBounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenFilternGram) GetTermNotInBoundsOk() (*string, bool) {
	if o == nil || isNil(o.TermNotInBounds) {
    return nil, false
	}
	return o.TermNotInBounds, true
}

// HasTermNotInBounds returns a boolean if a field has been set.
func (o *TokenFilternGram) HasTermNotInBounds() bool {
	if o != nil && !isNil(o.TermNotInBounds) {
		return true
	}

	return false
}

// SetTermNotInBounds gets a reference to the given string and assigns it to the TermNotInBounds field.
func (o *TokenFilternGram) SetTermNotInBounds(v string) {
	o.TermNotInBounds = &v
}

// GetType returns the Type field value
func (o *TokenFilternGram) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenFilternGram) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenFilternGram) SetType(v string) {
	o.Type = v
}

func (o TokenFilternGram) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["maxGram"] = o.MaxGram
	}
	if true {
		toSerialize["minGram"] = o.MinGram
	}
	if !isNil(o.TermNotInBounds) {
		toSerialize["termNotInBounds"] = o.TermNotInBounds
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTokenFilternGram struct {
	value *TokenFilternGram
	isSet bool
}

func (v NullableTokenFilternGram) Get() *TokenFilternGram {
	return v.value
}

func (v *NullableTokenFilternGram) Set(val *TokenFilternGram) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFilternGram) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFilternGram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFilternGram(val *TokenFilternGram) *NullableTokenFilternGram {
	return &NullableTokenFilternGram{value: val, isSet: true}
}

func (v NullableTokenFilternGram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFilternGram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


