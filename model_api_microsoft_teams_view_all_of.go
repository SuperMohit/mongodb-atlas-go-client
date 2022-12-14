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

// ApiMicrosoftTeamsViewAllOf struct for ApiMicrosoftTeamsViewAllOf
type ApiMicrosoftTeamsViewAllOf struct {
	// Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications.
	MicrosoftTeamsWebhookUrl *string `json:"microsoftTeamsWebhookUrl,omitempty"`
}

// NewApiMicrosoftTeamsViewAllOf instantiates a new ApiMicrosoftTeamsViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMicrosoftTeamsViewAllOf() *ApiMicrosoftTeamsViewAllOf {
	this := ApiMicrosoftTeamsViewAllOf{}
	return &this
}

// NewApiMicrosoftTeamsViewAllOfWithDefaults instantiates a new ApiMicrosoftTeamsViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMicrosoftTeamsViewAllOfWithDefaults() *ApiMicrosoftTeamsViewAllOf {
	this := ApiMicrosoftTeamsViewAllOf{}
	return &this
}

// GetMicrosoftTeamsWebhookUrl returns the MicrosoftTeamsWebhookUrl field value if set, zero value otherwise.
func (o *ApiMicrosoftTeamsViewAllOf) GetMicrosoftTeamsWebhookUrl() string {
	if o == nil || isNil(o.MicrosoftTeamsWebhookUrl) {
		var ret string
		return ret
	}
	return *o.MicrosoftTeamsWebhookUrl
}

// GetMicrosoftTeamsWebhookUrlOk returns a tuple with the MicrosoftTeamsWebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMicrosoftTeamsViewAllOf) GetMicrosoftTeamsWebhookUrlOk() (*string, bool) {
	if o == nil || isNil(o.MicrosoftTeamsWebhookUrl) {
    return nil, false
	}
	return o.MicrosoftTeamsWebhookUrl, true
}

// HasMicrosoftTeamsWebhookUrl returns a boolean if a field has been set.
func (o *ApiMicrosoftTeamsViewAllOf) HasMicrosoftTeamsWebhookUrl() bool {
	if o != nil && !isNil(o.MicrosoftTeamsWebhookUrl) {
		return true
	}

	return false
}

// SetMicrosoftTeamsWebhookUrl gets a reference to the given string and assigns it to the MicrosoftTeamsWebhookUrl field.
func (o *ApiMicrosoftTeamsViewAllOf) SetMicrosoftTeamsWebhookUrl(v string) {
	o.MicrosoftTeamsWebhookUrl = &v
}

func (o ApiMicrosoftTeamsViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MicrosoftTeamsWebhookUrl) {
		toSerialize["microsoftTeamsWebhookUrl"] = o.MicrosoftTeamsWebhookUrl
	}
	return json.Marshal(toSerialize)
}

type NullableApiMicrosoftTeamsViewAllOf struct {
	value *ApiMicrosoftTeamsViewAllOf
	isSet bool
}

func (v NullableApiMicrosoftTeamsViewAllOf) Get() *ApiMicrosoftTeamsViewAllOf {
	return v.value
}

func (v *NullableApiMicrosoftTeamsViewAllOf) Set(val *ApiMicrosoftTeamsViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMicrosoftTeamsViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMicrosoftTeamsViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMicrosoftTeamsViewAllOf(val *ApiMicrosoftTeamsViewAllOf) *NullableApiMicrosoftTeamsViewAllOf {
	return &NullableApiMicrosoftTeamsViewAllOf{value: val, isSet: true}
}

func (v NullableApiMicrosoftTeamsViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMicrosoftTeamsViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


