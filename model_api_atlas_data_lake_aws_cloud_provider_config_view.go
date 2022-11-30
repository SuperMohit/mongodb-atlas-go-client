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

// ApiAtlasDataLakeAWSCloudProviderConfigView Name of the cloud service that hosts the data lake's data stores.
type ApiAtlasDataLakeAWSCloudProviderConfigView struct {
	// Unique identifier associated with the Identity and Access Management (IAM) role that the data lake assumes when accessing the data stores.
	ExternalId *string `json:"externalId,omitempty"`
	// Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that the data lake assumes when accessing data stores.
	IamAssumedRoleARN *string `json:"iamAssumedRoleARN,omitempty"`
	// Amazon Resource Name (ARN) of the user that the data lake assumes when accessing data stores.
	IamUserARN *string `json:"iamUserARN,omitempty"`
	// Unique identifier of the role that the data lake can use to access the data stores.Required if specifying cloudProviderConfig.
	RoleId string `json:"roleId"`
	// Name of the S3 data bucket that the provided role ID is authorized to access.Required if specifying cloudProviderConfig.
	TestS3Bucket string `json:"testS3Bucket"`
}

// NewApiAtlasDataLakeAWSCloudProviderConfigView instantiates a new ApiAtlasDataLakeAWSCloudProviderConfigView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasDataLakeAWSCloudProviderConfigView(roleId string, testS3Bucket string) *ApiAtlasDataLakeAWSCloudProviderConfigView {
	this := ApiAtlasDataLakeAWSCloudProviderConfigView{}
	this.RoleId = roleId
	this.TestS3Bucket = testS3Bucket
	return &this
}

// NewApiAtlasDataLakeAWSCloudProviderConfigViewWithDefaults instantiates a new ApiAtlasDataLakeAWSCloudProviderConfigView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasDataLakeAWSCloudProviderConfigViewWithDefaults() *ApiAtlasDataLakeAWSCloudProviderConfigView {
	this := ApiAtlasDataLakeAWSCloudProviderConfigView{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) GetExternalId() string {
	if o == nil || isNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) GetExternalIdOk() (*string, bool) {
	if o == nil || isNil(o.ExternalId) {
    return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) HasExternalId() bool {
	if o != nil && !isNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetIamAssumedRoleARN returns the IamAssumedRoleARN field value if set, zero value otherwise.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) GetIamAssumedRoleARN() string {
	if o == nil || isNil(o.IamAssumedRoleARN) {
		var ret string
		return ret
	}
	return *o.IamAssumedRoleARN
}

// GetIamAssumedRoleARNOk returns a tuple with the IamAssumedRoleARN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) GetIamAssumedRoleARNOk() (*string, bool) {
	if o == nil || isNil(o.IamAssumedRoleARN) {
    return nil, false
	}
	return o.IamAssumedRoleARN, true
}

// HasIamAssumedRoleARN returns a boolean if a field has been set.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) HasIamAssumedRoleARN() bool {
	if o != nil && !isNil(o.IamAssumedRoleARN) {
		return true
	}

	return false
}

// SetIamAssumedRoleARN gets a reference to the given string and assigns it to the IamAssumedRoleARN field.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) SetIamAssumedRoleARN(v string) {
	o.IamAssumedRoleARN = &v
}

// GetIamUserARN returns the IamUserARN field value if set, zero value otherwise.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) GetIamUserARN() string {
	if o == nil || isNil(o.IamUserARN) {
		var ret string
		return ret
	}
	return *o.IamUserARN
}

// GetIamUserARNOk returns a tuple with the IamUserARN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) GetIamUserARNOk() (*string, bool) {
	if o == nil || isNil(o.IamUserARN) {
    return nil, false
	}
	return o.IamUserARN, true
}

// HasIamUserARN returns a boolean if a field has been set.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) HasIamUserARN() bool {
	if o != nil && !isNil(o.IamUserARN) {
		return true
	}

	return false
}

// SetIamUserARN gets a reference to the given string and assigns it to the IamUserARN field.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) SetIamUserARN(v string) {
	o.IamUserARN = &v
}

// GetRoleId returns the RoleId field value
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) GetRoleIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) SetRoleId(v string) {
	o.RoleId = v
}

// GetTestS3Bucket returns the TestS3Bucket field value
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) GetTestS3Bucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestS3Bucket
}

// GetTestS3BucketOk returns a tuple with the TestS3Bucket field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) GetTestS3BucketOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TestS3Bucket, true
}

// SetTestS3Bucket sets field value
func (o *ApiAtlasDataLakeAWSCloudProviderConfigView) SetTestS3Bucket(v string) {
	o.TestS3Bucket = v
}

func (o ApiAtlasDataLakeAWSCloudProviderConfigView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !isNil(o.IamAssumedRoleARN) {
		toSerialize["iamAssumedRoleARN"] = o.IamAssumedRoleARN
	}
	if !isNil(o.IamUserARN) {
		toSerialize["iamUserARN"] = o.IamUserARN
	}
	if true {
		toSerialize["roleId"] = o.RoleId
	}
	if true {
		toSerialize["testS3Bucket"] = o.TestS3Bucket
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasDataLakeAWSCloudProviderConfigView struct {
	value *ApiAtlasDataLakeAWSCloudProviderConfigView
	isSet bool
}

func (v NullableApiAtlasDataLakeAWSCloudProviderConfigView) Get() *ApiAtlasDataLakeAWSCloudProviderConfigView {
	return v.value
}

func (v *NullableApiAtlasDataLakeAWSCloudProviderConfigView) Set(val *ApiAtlasDataLakeAWSCloudProviderConfigView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasDataLakeAWSCloudProviderConfigView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasDataLakeAWSCloudProviderConfigView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasDataLakeAWSCloudProviderConfigView(val *ApiAtlasDataLakeAWSCloudProviderConfigView) *NullableApiAtlasDataLakeAWSCloudProviderConfigView {
	return &NullableApiAtlasDataLakeAWSCloudProviderConfigView{value: val, isSet: true}
}

func (v NullableApiAtlasDataLakeAWSCloudProviderConfigView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasDataLakeAWSCloudProviderConfigView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

