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

// PemFileInfo struct for PemFileInfo
type PemFileInfo struct {
	Certificates []X509Certificate `json:"certificates,omitempty"`
	FileName *string `json:"fileName,omitempty"`
}

// NewPemFileInfo instantiates a new PemFileInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPemFileInfo() *PemFileInfo {
	this := PemFileInfo{}
	return &this
}

// NewPemFileInfoWithDefaults instantiates a new PemFileInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPemFileInfoWithDefaults() *PemFileInfo {
	this := PemFileInfo{}
	return &this
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *PemFileInfo) GetCertificates() []X509Certificate {
	if o == nil || isNil(o.Certificates) {
		var ret []X509Certificate
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PemFileInfo) GetCertificatesOk() ([]X509Certificate, bool) {
	if o == nil || isNil(o.Certificates) {
    return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *PemFileInfo) HasCertificates() bool {
	if o != nil && !isNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []X509Certificate and assigns it to the Certificates field.
func (o *PemFileInfo) SetCertificates(v []X509Certificate) {
	o.Certificates = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *PemFileInfo) GetFileName() string {
	if o == nil || isNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PemFileInfo) GetFileNameOk() (*string, bool) {
	if o == nil || isNil(o.FileName) {
    return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *PemFileInfo) HasFileName() bool {
	if o != nil && !isNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *PemFileInfo) SetFileName(v string) {
	o.FileName = &v
}

func (o PemFileInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Certificates) {
		toSerialize["certificates"] = o.Certificates
	}
	if !isNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	return json.Marshal(toSerialize)
}

type NullablePemFileInfo struct {
	value *PemFileInfo
	isSet bool
}

func (v NullablePemFileInfo) Get() *PemFileInfo {
	return v.value
}

func (v *NullablePemFileInfo) Set(val *PemFileInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePemFileInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePemFileInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePemFileInfo(val *PemFileInfo) *NullablePemFileInfo {
	return &NullablePemFileInfo{value: val, isSet: true}
}

func (v NullablePemFileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePemFileInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


