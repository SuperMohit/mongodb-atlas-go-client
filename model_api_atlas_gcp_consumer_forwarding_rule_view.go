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

// ApiAtlasGCPConsumerForwardingRuleView struct for ApiAtlasGCPConsumerForwardingRuleView
type ApiAtlasGCPConsumerForwardingRuleView struct {
	// Human-readable label that identifies the Google Cloud consumer forwarding rule that you created.
	EndpointName *string `json:"endpointName,omitempty"`
	// One Private Internet Protocol version 4 (IPv4) address to which this Google Cloud consumer forwarding rule resolves.
	IpAddress *string `json:"ipAddress,omitempty"`
	// State of the MongoDB Cloud endpoint group when MongoDB Cloud received this request.
	Status *string `json:"status,omitempty"`
}

// NewApiAtlasGCPConsumerForwardingRuleView instantiates a new ApiAtlasGCPConsumerForwardingRuleView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasGCPConsumerForwardingRuleView() *ApiAtlasGCPConsumerForwardingRuleView {
	this := ApiAtlasGCPConsumerForwardingRuleView{}
	return &this
}

// NewApiAtlasGCPConsumerForwardingRuleViewWithDefaults instantiates a new ApiAtlasGCPConsumerForwardingRuleView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasGCPConsumerForwardingRuleViewWithDefaults() *ApiAtlasGCPConsumerForwardingRuleView {
	this := ApiAtlasGCPConsumerForwardingRuleView{}
	return &this
}

// GetEndpointName returns the EndpointName field value if set, zero value otherwise.
func (o *ApiAtlasGCPConsumerForwardingRuleView) GetEndpointName() string {
	if o == nil || isNil(o.EndpointName) {
		var ret string
		return ret
	}
	return *o.EndpointName
}

// GetEndpointNameOk returns a tuple with the EndpointName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGCPConsumerForwardingRuleView) GetEndpointNameOk() (*string, bool) {
	if o == nil || isNil(o.EndpointName) {
    return nil, false
	}
	return o.EndpointName, true
}

// HasEndpointName returns a boolean if a field has been set.
func (o *ApiAtlasGCPConsumerForwardingRuleView) HasEndpointName() bool {
	if o != nil && !isNil(o.EndpointName) {
		return true
	}

	return false
}

// SetEndpointName gets a reference to the given string and assigns it to the EndpointName field.
func (o *ApiAtlasGCPConsumerForwardingRuleView) SetEndpointName(v string) {
	o.EndpointName = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ApiAtlasGCPConsumerForwardingRuleView) GetIpAddress() string {
	if o == nil || isNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGCPConsumerForwardingRuleView) GetIpAddressOk() (*string, bool) {
	if o == nil || isNil(o.IpAddress) {
    return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ApiAtlasGCPConsumerForwardingRuleView) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ApiAtlasGCPConsumerForwardingRuleView) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiAtlasGCPConsumerForwardingRuleView) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGCPConsumerForwardingRuleView) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiAtlasGCPConsumerForwardingRuleView) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiAtlasGCPConsumerForwardingRuleView) SetStatus(v string) {
	o.Status = &v
}

func (o ApiAtlasGCPConsumerForwardingRuleView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EndpointName) {
		toSerialize["endpointName"] = o.EndpointName
	}
	if !isNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasGCPConsumerForwardingRuleView struct {
	value *ApiAtlasGCPConsumerForwardingRuleView
	isSet bool
}

func (v NullableApiAtlasGCPConsumerForwardingRuleView) Get() *ApiAtlasGCPConsumerForwardingRuleView {
	return v.value
}

func (v *NullableApiAtlasGCPConsumerForwardingRuleView) Set(val *ApiAtlasGCPConsumerForwardingRuleView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasGCPConsumerForwardingRuleView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasGCPConsumerForwardingRuleView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasGCPConsumerForwardingRuleView(val *ApiAtlasGCPConsumerForwardingRuleView) *NullableApiAtlasGCPConsumerForwardingRuleView {
	return &NullableApiAtlasGCPConsumerForwardingRuleView{value: val, isSet: true}
}

func (v NullableApiAtlasGCPConsumerForwardingRuleView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasGCPConsumerForwardingRuleView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


