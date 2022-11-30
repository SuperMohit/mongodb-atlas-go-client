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

// ApiAtlasCreateAzureEndpointRequestView Group of Private Endpoint settings.
type ApiAtlasCreateAzureEndpointRequestView struct {
	// Unique string that identifies the private endpoint's network interface that someone added to this private endpoint service.
	Id *string `json:"id,omitempty"`
	// IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.
	PrivateEndpointIPAddress *string `json:"privateEndpointIPAddress,omitempty"`
}

// NewApiAtlasCreateAzureEndpointRequestView instantiates a new ApiAtlasCreateAzureEndpointRequestView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCreateAzureEndpointRequestView() *ApiAtlasCreateAzureEndpointRequestView {
	this := ApiAtlasCreateAzureEndpointRequestView{}
	return &this
}

// NewApiAtlasCreateAzureEndpointRequestViewWithDefaults instantiates a new ApiAtlasCreateAzureEndpointRequestView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCreateAzureEndpointRequestViewWithDefaults() *ApiAtlasCreateAzureEndpointRequestView {
	this := ApiAtlasCreateAzureEndpointRequestView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAtlasCreateAzureEndpointRequestView) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCreateAzureEndpointRequestView) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAtlasCreateAzureEndpointRequestView) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAtlasCreateAzureEndpointRequestView) SetId(v string) {
	o.Id = &v
}

// GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field value if set, zero value otherwise.
func (o *ApiAtlasCreateAzureEndpointRequestView) GetPrivateEndpointIPAddress() string {
	if o == nil || isNil(o.PrivateEndpointIPAddress) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointIPAddress
}

// GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCreateAzureEndpointRequestView) GetPrivateEndpointIPAddressOk() (*string, bool) {
	if o == nil || isNil(o.PrivateEndpointIPAddress) {
    return nil, false
	}
	return o.PrivateEndpointIPAddress, true
}

// HasPrivateEndpointIPAddress returns a boolean if a field has been set.
func (o *ApiAtlasCreateAzureEndpointRequestView) HasPrivateEndpointIPAddress() bool {
	if o != nil && !isNil(o.PrivateEndpointIPAddress) {
		return true
	}

	return false
}

// SetPrivateEndpointIPAddress gets a reference to the given string and assigns it to the PrivateEndpointIPAddress field.
func (o *ApiAtlasCreateAzureEndpointRequestView) SetPrivateEndpointIPAddress(v string) {
	o.PrivateEndpointIPAddress = &v
}

func (o ApiAtlasCreateAzureEndpointRequestView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.PrivateEndpointIPAddress) {
		toSerialize["privateEndpointIPAddress"] = o.PrivateEndpointIPAddress
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasCreateAzureEndpointRequestView struct {
	value *ApiAtlasCreateAzureEndpointRequestView
	isSet bool
}

func (v NullableApiAtlasCreateAzureEndpointRequestView) Get() *ApiAtlasCreateAzureEndpointRequestView {
	return v.value
}

func (v *NullableApiAtlasCreateAzureEndpointRequestView) Set(val *ApiAtlasCreateAzureEndpointRequestView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasCreateAzureEndpointRequestView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasCreateAzureEndpointRequestView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasCreateAzureEndpointRequestView(val *ApiAtlasCreateAzureEndpointRequestView) *NullableApiAtlasCreateAzureEndpointRequestView {
	return &NullableApiAtlasCreateAzureEndpointRequestView{value: val, isSet: true}
}

func (v NullableApiAtlasCreateAzureEndpointRequestView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasCreateAzureEndpointRequestView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


