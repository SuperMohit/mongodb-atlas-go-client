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

// ApiMongoDBAccessLogsView Authentication attempt, one per object, made against the cluster.
type ApiMongoDBAccessLogsView struct {
	// Flag that indicates whether the response should return successful authentication attempts only.
	AuthResult *bool `json:"authResult,omitempty"`
	// Database against which someone attempted to authenticate.
	AuthSource *string `json:"authSource,omitempty"`
	// Reason that the authentication failed. Null if authentication succeeded. 
	FailureReason *string `json:"failureReason,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	GroupId *string `json:"groupId,omitempty"`
	// Human-readable label that identifies the hostname of the target node that received the authentication attempt.
	Hostname *string `json:"hostname,omitempty"`
	// Internet Protocol address that attempted to authenticate with the database.
	IpAddress *string `json:"ipAddress,omitempty"`
	// Text of the host log concerning the authentication attempt.
	LogLine *string `json:"logLine,omitempty"`
	// Date and time when someone made this authentication attempt. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	Timestamp *string `json:"timestamp,omitempty"`
	// Username used to authenticate against the database.
	Username *string `json:"username,omitempty"`
}

// NewApiMongoDBAccessLogsView instantiates a new ApiMongoDBAccessLogsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMongoDBAccessLogsView() *ApiMongoDBAccessLogsView {
	this := ApiMongoDBAccessLogsView{}
	return &this
}

// NewApiMongoDBAccessLogsViewWithDefaults instantiates a new ApiMongoDBAccessLogsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMongoDBAccessLogsViewWithDefaults() *ApiMongoDBAccessLogsView {
	this := ApiMongoDBAccessLogsView{}
	return &this
}

// GetAuthResult returns the AuthResult field value if set, zero value otherwise.
func (o *ApiMongoDBAccessLogsView) GetAuthResult() bool {
	if o == nil || isNil(o.AuthResult) {
		var ret bool
		return ret
	}
	return *o.AuthResult
}

// GetAuthResultOk returns a tuple with the AuthResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMongoDBAccessLogsView) GetAuthResultOk() (*bool, bool) {
	if o == nil || isNil(o.AuthResult) {
    return nil, false
	}
	return o.AuthResult, true
}

// HasAuthResult returns a boolean if a field has been set.
func (o *ApiMongoDBAccessLogsView) HasAuthResult() bool {
	if o != nil && !isNil(o.AuthResult) {
		return true
	}

	return false
}

// SetAuthResult gets a reference to the given bool and assigns it to the AuthResult field.
func (o *ApiMongoDBAccessLogsView) SetAuthResult(v bool) {
	o.AuthResult = &v
}

// GetAuthSource returns the AuthSource field value if set, zero value otherwise.
func (o *ApiMongoDBAccessLogsView) GetAuthSource() string {
	if o == nil || isNil(o.AuthSource) {
		var ret string
		return ret
	}
	return *o.AuthSource
}

// GetAuthSourceOk returns a tuple with the AuthSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMongoDBAccessLogsView) GetAuthSourceOk() (*string, bool) {
	if o == nil || isNil(o.AuthSource) {
    return nil, false
	}
	return o.AuthSource, true
}

// HasAuthSource returns a boolean if a field has been set.
func (o *ApiMongoDBAccessLogsView) HasAuthSource() bool {
	if o != nil && !isNil(o.AuthSource) {
		return true
	}

	return false
}

// SetAuthSource gets a reference to the given string and assigns it to the AuthSource field.
func (o *ApiMongoDBAccessLogsView) SetAuthSource(v string) {
	o.AuthSource = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *ApiMongoDBAccessLogsView) GetFailureReason() string {
	if o == nil || isNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMongoDBAccessLogsView) GetFailureReasonOk() (*string, bool) {
	if o == nil || isNil(o.FailureReason) {
    return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *ApiMongoDBAccessLogsView) HasFailureReason() bool {
	if o != nil && !isNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *ApiMongoDBAccessLogsView) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApiMongoDBAccessLogsView) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMongoDBAccessLogsView) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiMongoDBAccessLogsView) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiMongoDBAccessLogsView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApiMongoDBAccessLogsView) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMongoDBAccessLogsView) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
    return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApiMongoDBAccessLogsView) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApiMongoDBAccessLogsView) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ApiMongoDBAccessLogsView) GetIpAddress() string {
	if o == nil || isNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMongoDBAccessLogsView) GetIpAddressOk() (*string, bool) {
	if o == nil || isNil(o.IpAddress) {
    return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ApiMongoDBAccessLogsView) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ApiMongoDBAccessLogsView) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLogLine returns the LogLine field value if set, zero value otherwise.
func (o *ApiMongoDBAccessLogsView) GetLogLine() string {
	if o == nil || isNil(o.LogLine) {
		var ret string
		return ret
	}
	return *o.LogLine
}

// GetLogLineOk returns a tuple with the LogLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMongoDBAccessLogsView) GetLogLineOk() (*string, bool) {
	if o == nil || isNil(o.LogLine) {
    return nil, false
	}
	return o.LogLine, true
}

// HasLogLine returns a boolean if a field has been set.
func (o *ApiMongoDBAccessLogsView) HasLogLine() bool {
	if o != nil && !isNil(o.LogLine) {
		return true
	}

	return false
}

// SetLogLine gets a reference to the given string and assigns it to the LogLine field.
func (o *ApiMongoDBAccessLogsView) SetLogLine(v string) {
	o.LogLine = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ApiMongoDBAccessLogsView) GetTimestamp() string {
	if o == nil || isNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMongoDBAccessLogsView) GetTimestampOk() (*string, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ApiMongoDBAccessLogsView) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ApiMongoDBAccessLogsView) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiMongoDBAccessLogsView) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMongoDBAccessLogsView) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiMongoDBAccessLogsView) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiMongoDBAccessLogsView) SetUsername(v string) {
	o.Username = &v
}

func (o ApiMongoDBAccessLogsView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthResult) {
		toSerialize["authResult"] = o.AuthResult
	}
	if !isNil(o.AuthSource) {
		toSerialize["authSource"] = o.AuthSource
	}
	if !isNil(o.FailureReason) {
		toSerialize["failureReason"] = o.FailureReason
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !isNil(o.LogLine) {
		toSerialize["logLine"] = o.LogLine
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableApiMongoDBAccessLogsView struct {
	value *ApiMongoDBAccessLogsView
	isSet bool
}

func (v NullableApiMongoDBAccessLogsView) Get() *ApiMongoDBAccessLogsView {
	return v.value
}

func (v *NullableApiMongoDBAccessLogsView) Set(val *ApiMongoDBAccessLogsView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMongoDBAccessLogsView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMongoDBAccessLogsView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMongoDBAccessLogsView(val *ApiMongoDBAccessLogsView) *NullableApiMongoDBAccessLogsView {
	return &NullableApiMongoDBAccessLogsView{value: val, isSet: true}
}

func (v NullableApiMongoDBAccessLogsView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMongoDBAccessLogsView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


