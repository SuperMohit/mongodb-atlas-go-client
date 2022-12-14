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

// ApiAtlasAuditLogView struct for ApiAtlasAuditLogView
type ApiAtlasAuditLogView struct {
	// Flag that indicates whether someone set auditing to track successful authentications. This only applies to the `\"atype\" : \"authCheck\"` audit filter. Setting this parameter to `true` degrades cluster performance.
	AuditAuthorizationSuccess bool `json:"auditAuthorizationSuccess"`
	// JSON document that specifies which events to record. Escape any characters that may prevent parsing, such as single or double quotes, using a backslash (`\\`).
	AuditFilter string `json:"auditFilter"`
	// Human-readable label that displays how to configure the audit filter.
	ConfigurationType *string `json:"configurationType,omitempty"`
	// Flag that indicates whether someone enabled database auditing for the specified project.
	Enabled bool `json:"enabled"`
}

// NewApiAtlasAuditLogView instantiates a new ApiAtlasAuditLogView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasAuditLogView(auditAuthorizationSuccess bool, auditFilter string, enabled bool) *ApiAtlasAuditLogView {
	this := ApiAtlasAuditLogView{}
	this.AuditAuthorizationSuccess = auditAuthorizationSuccess
	this.AuditFilter = auditFilter
	this.Enabled = enabled
	return &this
}

// NewApiAtlasAuditLogViewWithDefaults instantiates a new ApiAtlasAuditLogView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasAuditLogViewWithDefaults() *ApiAtlasAuditLogView {
	this := ApiAtlasAuditLogView{}
	var auditAuthorizationSuccess bool = false
	this.AuditAuthorizationSuccess = auditAuthorizationSuccess
	var enabled bool = false
	this.Enabled = enabled
	return &this
}

// GetAuditAuthorizationSuccess returns the AuditAuthorizationSuccess field value
func (o *ApiAtlasAuditLogView) GetAuditAuthorizationSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AuditAuthorizationSuccess
}

// GetAuditAuthorizationSuccessOk returns a tuple with the AuditAuthorizationSuccess field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasAuditLogView) GetAuditAuthorizationSuccessOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuditAuthorizationSuccess, true
}

// SetAuditAuthorizationSuccess sets field value
func (o *ApiAtlasAuditLogView) SetAuditAuthorizationSuccess(v bool) {
	o.AuditAuthorizationSuccess = v
}

// GetAuditFilter returns the AuditFilter field value
func (o *ApiAtlasAuditLogView) GetAuditFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditFilter
}

// GetAuditFilterOk returns a tuple with the AuditFilter field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasAuditLogView) GetAuditFilterOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuditFilter, true
}

// SetAuditFilter sets field value
func (o *ApiAtlasAuditLogView) SetAuditFilter(v string) {
	o.AuditFilter = v
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise.
func (o *ApiAtlasAuditLogView) GetConfigurationType() string {
	if o == nil || isNil(o.ConfigurationType) {
		var ret string
		return ret
	}
	return *o.ConfigurationType
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasAuditLogView) GetConfigurationTypeOk() (*string, bool) {
	if o == nil || isNil(o.ConfigurationType) {
    return nil, false
	}
	return o.ConfigurationType, true
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *ApiAtlasAuditLogView) HasConfigurationType() bool {
	if o != nil && !isNil(o.ConfigurationType) {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given string and assigns it to the ConfigurationType field.
func (o *ApiAtlasAuditLogView) SetConfigurationType(v string) {
	o.ConfigurationType = &v
}

// GetEnabled returns the Enabled field value
func (o *ApiAtlasAuditLogView) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasAuditLogView) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ApiAtlasAuditLogView) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ApiAtlasAuditLogView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auditAuthorizationSuccess"] = o.AuditAuthorizationSuccess
	}
	if true {
		toSerialize["auditFilter"] = o.AuditFilter
	}
	if !isNil(o.ConfigurationType) {
		toSerialize["configurationType"] = o.ConfigurationType
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasAuditLogView struct {
	value *ApiAtlasAuditLogView
	isSet bool
}

func (v NullableApiAtlasAuditLogView) Get() *ApiAtlasAuditLogView {
	return v.value
}

func (v *NullableApiAtlasAuditLogView) Set(val *ApiAtlasAuditLogView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasAuditLogView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasAuditLogView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasAuditLogView(val *ApiAtlasAuditLogView) *NullableApiAtlasAuditLogView {
	return &NullableApiAtlasAuditLogView{value: val, isSet: true}
}

func (v NullableApiAtlasAuditLogView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasAuditLogView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


