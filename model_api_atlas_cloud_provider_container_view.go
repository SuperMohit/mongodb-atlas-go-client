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

// ApiAtlasCloudProviderContainerView Collection of settings that configures the network container for a virtual private connection on Amazon Web Services.
type ApiAtlasCloudProviderContainerView struct {
	// Unique 24-hexadecimal digit string that identifies the network peering container.
	Id *string `json:"id,omitempty"`
	// Cloud service provider that serves the requested network peering containers.
	ProviderName *string `json:"providerName,omitempty"`
	// Flag that indicates whether MongoDB Cloud clusters exist in the specified network peering container.
	Provisioned *bool `json:"provisioned,omitempty"`
}

// NewApiAtlasCloudProviderContainerView instantiates a new ApiAtlasCloudProviderContainerView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCloudProviderContainerView() *ApiAtlasCloudProviderContainerView {
	this := ApiAtlasCloudProviderContainerView{}
	return &this
}

// NewApiAtlasCloudProviderContainerViewWithDefaults instantiates a new ApiAtlasCloudProviderContainerView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCloudProviderContainerViewWithDefaults() *ApiAtlasCloudProviderContainerView {
	this := ApiAtlasCloudProviderContainerView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderContainerView) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderContainerView) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderContainerView) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAtlasCloudProviderContainerView) SetId(v string) {
	o.Id = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderContainerView) GetProviderName() string {
	if o == nil || isNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderContainerView) GetProviderNameOk() (*string, bool) {
	if o == nil || isNil(o.ProviderName) {
    return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderContainerView) HasProviderName() bool {
	if o != nil && !isNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *ApiAtlasCloudProviderContainerView) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetProvisioned returns the Provisioned field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderContainerView) GetProvisioned() bool {
	if o == nil || isNil(o.Provisioned) {
		var ret bool
		return ret
	}
	return *o.Provisioned
}

// GetProvisionedOk returns a tuple with the Provisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderContainerView) GetProvisionedOk() (*bool, bool) {
	if o == nil || isNil(o.Provisioned) {
    return nil, false
	}
	return o.Provisioned, true
}

// HasProvisioned returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderContainerView) HasProvisioned() bool {
	if o != nil && !isNil(o.Provisioned) {
		return true
	}

	return false
}

// SetProvisioned gets a reference to the given bool and assigns it to the Provisioned field.
func (o *ApiAtlasCloudProviderContainerView) SetProvisioned(v bool) {
	o.Provisioned = &v
}

func (o ApiAtlasCloudProviderContainerView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	if !isNil(o.Provisioned) {
		toSerialize["provisioned"] = o.Provisioned
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasCloudProviderContainerView struct {
	value *ApiAtlasCloudProviderContainerView
	isSet bool
}

func (v NullableApiAtlasCloudProviderContainerView) Get() *ApiAtlasCloudProviderContainerView {
	return v.value
}

func (v *NullableApiAtlasCloudProviderContainerView) Set(val *ApiAtlasCloudProviderContainerView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasCloudProviderContainerView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasCloudProviderContainerView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasCloudProviderContainerView(val *ApiAtlasCloudProviderContainerView) *NullableApiAtlasCloudProviderContainerView {
	return &NullableApiAtlasCloudProviderContainerView{value: val, isSet: true}
}

func (v NullableApiAtlasCloudProviderContainerView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasCloudProviderContainerView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


