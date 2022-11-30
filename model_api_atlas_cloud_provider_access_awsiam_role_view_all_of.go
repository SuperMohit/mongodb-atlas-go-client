/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf struct for ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf
type ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf struct {
	// Amazon Resource Name that identifies the Amazon Web Services (AWS) user account that MongoDB Cloud uses when it assumes the Identity and Access Management (IAM) role.
	AtlasAWSAccountArn *string `json:"atlasAWSAccountArn,omitempty"`
	// Unique external ID that MongoDB Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account.
	AtlasAssumedRoleExternalId *string `json:"atlasAssumedRoleExternalId,omitempty"`
	// Date and time when someone authorized this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	AuthorizedDate *time.Time `json:"authorizedDate,omitempty"`
	// Date and time when someone created this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// List that contains application features associated with this Amazon Web Services (AWS) Identity and Access Management (IAM) role.
	FeatureUsages []ApiAtlasCloudProviderAccessFeatureUsageView `json:"featureUsages,omitempty"`
	// Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that MongoDB Cloud assumes when it accesses resources in your AWS account.
	IamAssumedRoleArn *string `json:"iamAssumedRoleArn,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the role.
	RoleId *string `json:"roleId,omitempty"`
}

// NewApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf instantiates a new ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf() *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf {
	this := ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf{}
	return &this
}

// NewApiAtlasCloudProviderAccessAWSIAMRoleViewAllOfWithDefaults instantiates a new ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCloudProviderAccessAWSIAMRoleViewAllOfWithDefaults() *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf {
	this := ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf{}
	return &this
}

// GetAtlasAWSAccountArn returns the AtlasAWSAccountArn field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAtlasAWSAccountArn() string {
	if o == nil || isNil(o.AtlasAWSAccountArn) {
		var ret string
		return ret
	}
	return *o.AtlasAWSAccountArn
}

// GetAtlasAWSAccountArnOk returns a tuple with the AtlasAWSAccountArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAtlasAWSAccountArnOk() (*string, bool) {
	if o == nil || isNil(o.AtlasAWSAccountArn) {
    return nil, false
	}
	return o.AtlasAWSAccountArn, true
}

// HasAtlasAWSAccountArn returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasAtlasAWSAccountArn() bool {
	if o != nil && !isNil(o.AtlasAWSAccountArn) {
		return true
	}

	return false
}

// SetAtlasAWSAccountArn gets a reference to the given string and assigns it to the AtlasAWSAccountArn field.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetAtlasAWSAccountArn(v string) {
	o.AtlasAWSAccountArn = &v
}

// GetAtlasAssumedRoleExternalId returns the AtlasAssumedRoleExternalId field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAtlasAssumedRoleExternalId() string {
	if o == nil || isNil(o.AtlasAssumedRoleExternalId) {
		var ret string
		return ret
	}
	return *o.AtlasAssumedRoleExternalId
}

// GetAtlasAssumedRoleExternalIdOk returns a tuple with the AtlasAssumedRoleExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAtlasAssumedRoleExternalIdOk() (*string, bool) {
	if o == nil || isNil(o.AtlasAssumedRoleExternalId) {
    return nil, false
	}
	return o.AtlasAssumedRoleExternalId, true
}

// HasAtlasAssumedRoleExternalId returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasAtlasAssumedRoleExternalId() bool {
	if o != nil && !isNil(o.AtlasAssumedRoleExternalId) {
		return true
	}

	return false
}

// SetAtlasAssumedRoleExternalId gets a reference to the given string and assigns it to the AtlasAssumedRoleExternalId field.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetAtlasAssumedRoleExternalId(v string) {
	o.AtlasAssumedRoleExternalId = &v
}

// GetAuthorizedDate returns the AuthorizedDate field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAuthorizedDate() time.Time {
	if o == nil || isNil(o.AuthorizedDate) {
		var ret time.Time
		return ret
	}
	return *o.AuthorizedDate
}

// GetAuthorizedDateOk returns a tuple with the AuthorizedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAuthorizedDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.AuthorizedDate) {
    return nil, false
	}
	return o.AuthorizedDate, true
}

// HasAuthorizedDate returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasAuthorizedDate() bool {
	if o != nil && !isNil(o.AuthorizedDate) {
		return true
	}

	return false
}

// SetAuthorizedDate gets a reference to the given time.Time and assigns it to the AuthorizedDate field.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetAuthorizedDate(v time.Time) {
	o.AuthorizedDate = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetCreatedDate() time.Time {
	if o == nil || isNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedDate) {
    return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasCreatedDate() bool {
	if o != nil && !isNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetFeatureUsages returns the FeatureUsages field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetFeatureUsages() []ApiAtlasCloudProviderAccessFeatureUsageView {
	if o == nil || isNil(o.FeatureUsages) {
		var ret []ApiAtlasCloudProviderAccessFeatureUsageView
		return ret
	}
	return o.FeatureUsages
}

// GetFeatureUsagesOk returns a tuple with the FeatureUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetFeatureUsagesOk() ([]ApiAtlasCloudProviderAccessFeatureUsageView, bool) {
	if o == nil || isNil(o.FeatureUsages) {
    return nil, false
	}
	return o.FeatureUsages, true
}

// HasFeatureUsages returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasFeatureUsages() bool {
	if o != nil && !isNil(o.FeatureUsages) {
		return true
	}

	return false
}

// SetFeatureUsages gets a reference to the given []ApiAtlasCloudProviderAccessFeatureUsageView and assigns it to the FeatureUsages field.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetFeatureUsages(v []ApiAtlasCloudProviderAccessFeatureUsageView) {
	o.FeatureUsages = v
}

// GetIamAssumedRoleArn returns the IamAssumedRoleArn field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetIamAssumedRoleArn() string {
	if o == nil || isNil(o.IamAssumedRoleArn) {
		var ret string
		return ret
	}
	return *o.IamAssumedRoleArn
}

// GetIamAssumedRoleArnOk returns a tuple with the IamAssumedRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetIamAssumedRoleArnOk() (*string, bool) {
	if o == nil || isNil(o.IamAssumedRoleArn) {
    return nil, false
	}
	return o.IamAssumedRoleArn, true
}

// HasIamAssumedRoleArn returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasIamAssumedRoleArn() bool {
	if o != nil && !isNil(o.IamAssumedRoleArn) {
		return true
	}

	return false
}

// SetIamAssumedRoleArn gets a reference to the given string and assigns it to the IamAssumedRoleArn field.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetIamAssumedRoleArn(v string) {
	o.IamAssumedRoleArn = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetRoleId() string {
	if o == nil || isNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetRoleIdOk() (*string, bool) {
	if o == nil || isNil(o.RoleId) {
    return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasRoleId() bool {
	if o != nil && !isNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetRoleId(v string) {
	o.RoleId = &v
}

func (o ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AtlasAWSAccountArn) {
		toSerialize["atlasAWSAccountArn"] = o.AtlasAWSAccountArn
	}
	if !isNil(o.AtlasAssumedRoleExternalId) {
		toSerialize["atlasAssumedRoleExternalId"] = o.AtlasAssumedRoleExternalId
	}
	if !isNil(o.AuthorizedDate) {
		toSerialize["authorizedDate"] = o.AuthorizedDate
	}
	if !isNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !isNil(o.FeatureUsages) {
		toSerialize["featureUsages"] = o.FeatureUsages
	}
	if !isNil(o.IamAssumedRoleArn) {
		toSerialize["iamAssumedRoleArn"] = o.IamAssumedRoleArn
	}
	if !isNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf struct {
	value *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf
	isSet bool
}

func (v NullableApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) Get() *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf {
	return v.value
}

func (v *NullableApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) Set(val *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf(val *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) *NullableApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf {
	return &NullableApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf{value: val, isSet: true}
}

func (v NullableApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


