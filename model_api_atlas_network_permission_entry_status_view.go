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

// ApiAtlasNetworkPermissionEntryStatusView struct for ApiAtlasNetworkPermissionEntryStatusView
type ApiAtlasNetworkPermissionEntryStatusView struct {
	// State of the access list entry when MongoDB Cloud made this request.  | Status | Activity | |---|---| | `ACTIVE` | This access list entry applies to all relevant cloud providers. | | `PENDING` | MongoDB Cloud has started to add access list entry. This access list entry may not apply to all cloud providers at the time of this request. | | `FAILED` | MongoDB Cloud didn't succeed in adding this access list entry. | 
	STATUS string `json:"STATUS"`
}

// NewApiAtlasNetworkPermissionEntryStatusView instantiates a new ApiAtlasNetworkPermissionEntryStatusView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasNetworkPermissionEntryStatusView(sTATUS string) *ApiAtlasNetworkPermissionEntryStatusView {
	this := ApiAtlasNetworkPermissionEntryStatusView{}
	this.STATUS = sTATUS
	return &this
}

// NewApiAtlasNetworkPermissionEntryStatusViewWithDefaults instantiates a new ApiAtlasNetworkPermissionEntryStatusView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasNetworkPermissionEntryStatusViewWithDefaults() *ApiAtlasNetworkPermissionEntryStatusView {
	this := ApiAtlasNetworkPermissionEntryStatusView{}
	return &this
}

// GetSTATUS returns the STATUS field value
func (o *ApiAtlasNetworkPermissionEntryStatusView) GetSTATUS() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasNetworkPermissionEntryStatusView) GetSTATUSOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.STATUS, true
}

// SetSTATUS sets field value
func (o *ApiAtlasNetworkPermissionEntryStatusView) SetSTATUS(v string) {
	o.STATUS = v
}

func (o ApiAtlasNetworkPermissionEntryStatusView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["STATUS"] = o.STATUS
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasNetworkPermissionEntryStatusView struct {
	value *ApiAtlasNetworkPermissionEntryStatusView
	isSet bool
}

func (v NullableApiAtlasNetworkPermissionEntryStatusView) Get() *ApiAtlasNetworkPermissionEntryStatusView {
	return v.value
}

func (v *NullableApiAtlasNetworkPermissionEntryStatusView) Set(val *ApiAtlasNetworkPermissionEntryStatusView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasNetworkPermissionEntryStatusView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasNetworkPermissionEntryStatusView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasNetworkPermissionEntryStatusView(val *ApiAtlasNetworkPermissionEntryStatusView) *NullableApiAtlasNetworkPermissionEntryStatusView {
	return &NullableApiAtlasNetworkPermissionEntryStatusView{value: val, isSet: true}
}

func (v NullableApiAtlasNetworkPermissionEntryStatusView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasNetworkPermissionEntryStatusView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


