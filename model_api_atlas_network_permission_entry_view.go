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

// ApiAtlasNetworkPermissionEntryView struct for ApiAtlasNetworkPermissionEntryView
type ApiAtlasNetworkPermissionEntryView struct {
	// Unique string of the Amazon Web Services (AWS) security group that you want to add to the project's IP access list. Your IP access list entry can be one **awsSecurityGroup**, one **cidrBlock**, or one **ipAddress**. You must configure Virtual Private Connection (VPC) peering for your project before you can add an AWS security group to an IP access list. You cannot set AWS security groups as temporary access list entries. Don't set this parameter if you set **cidrBlock** or **ipAddress**.
	AwsSecurityGroup *string `json:"awsSecurityGroup,omitempty"`
	// Range of IP addresses in Classless Inter-Domain Routing (CIDR) notation that you want to add to the project's IP access list. Your IP access list entry can be one **awsSecurityGroup**, one **cidrBlock**, or one **ipAddress**. Don't set this parameter if you set **awsSecurityGroup** or **ipAddress**.
	CidrBlock *string `json:"cidrBlock,omitempty"`
	// Remark that explains the purpose or scope of this IP access list entry.
	Comment *string `json:"comment,omitempty"`
	// Date and time after which MongoDB Cloud deletes the temporary access list entry. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. The date must be later than the current date but no later than one week after you submit this request. The resource returns this parameter if you specified an expiration date when creating this IP access list entry.
	DeleteAfterDate *time.Time `json:"deleteAfterDate,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project that contains the IP access list to which you want to add one or more entries.
	GroupId *string `json:"groupId,omitempty"`
	// IP address that you want to add to the project's IP access list. Your IP access list entry can be one **awsSecurityGroup**, one **cidrBlock**, or one **ipAddress**. Don't set this parameter if you set **awsSecurityGroup** or **cidrBlock**.
	IpAddress *string `json:"ipAddress,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links"`
}

// NewApiAtlasNetworkPermissionEntryView instantiates a new ApiAtlasNetworkPermissionEntryView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasNetworkPermissionEntryView(links []Link) *ApiAtlasNetworkPermissionEntryView {
	this := ApiAtlasNetworkPermissionEntryView{}
	this.Links = links
	return &this
}

// NewApiAtlasNetworkPermissionEntryViewWithDefaults instantiates a new ApiAtlasNetworkPermissionEntryView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasNetworkPermissionEntryViewWithDefaults() *ApiAtlasNetworkPermissionEntryView {
	this := ApiAtlasNetworkPermissionEntryView{}
	return &this
}

// GetAwsSecurityGroup returns the AwsSecurityGroup field value if set, zero value otherwise.
func (o *ApiAtlasNetworkPermissionEntryView) GetAwsSecurityGroup() string {
	if o == nil || isNil(o.AwsSecurityGroup) {
		var ret string
		return ret
	}
	return *o.AwsSecurityGroup
}

// GetAwsSecurityGroupOk returns a tuple with the AwsSecurityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasNetworkPermissionEntryView) GetAwsSecurityGroupOk() (*string, bool) {
	if o == nil || isNil(o.AwsSecurityGroup) {
    return nil, false
	}
	return o.AwsSecurityGroup, true
}

// HasAwsSecurityGroup returns a boolean if a field has been set.
func (o *ApiAtlasNetworkPermissionEntryView) HasAwsSecurityGroup() bool {
	if o != nil && !isNil(o.AwsSecurityGroup) {
		return true
	}

	return false
}

// SetAwsSecurityGroup gets a reference to the given string and assigns it to the AwsSecurityGroup field.
func (o *ApiAtlasNetworkPermissionEntryView) SetAwsSecurityGroup(v string) {
	o.AwsSecurityGroup = &v
}

// GetCidrBlock returns the CidrBlock field value if set, zero value otherwise.
func (o *ApiAtlasNetworkPermissionEntryView) GetCidrBlock() string {
	if o == nil || isNil(o.CidrBlock) {
		var ret string
		return ret
	}
	return *o.CidrBlock
}

// GetCidrBlockOk returns a tuple with the CidrBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasNetworkPermissionEntryView) GetCidrBlockOk() (*string, bool) {
	if o == nil || isNil(o.CidrBlock) {
    return nil, false
	}
	return o.CidrBlock, true
}

// HasCidrBlock returns a boolean if a field has been set.
func (o *ApiAtlasNetworkPermissionEntryView) HasCidrBlock() bool {
	if o != nil && !isNil(o.CidrBlock) {
		return true
	}

	return false
}

// SetCidrBlock gets a reference to the given string and assigns it to the CidrBlock field.
func (o *ApiAtlasNetworkPermissionEntryView) SetCidrBlock(v string) {
	o.CidrBlock = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApiAtlasNetworkPermissionEntryView) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasNetworkPermissionEntryView) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApiAtlasNetworkPermissionEntryView) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApiAtlasNetworkPermissionEntryView) SetComment(v string) {
	o.Comment = &v
}

// GetDeleteAfterDate returns the DeleteAfterDate field value if set, zero value otherwise.
func (o *ApiAtlasNetworkPermissionEntryView) GetDeleteAfterDate() time.Time {
	if o == nil || isNil(o.DeleteAfterDate) {
		var ret time.Time
		return ret
	}
	return *o.DeleteAfterDate
}

// GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasNetworkPermissionEntryView) GetDeleteAfterDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.DeleteAfterDate) {
    return nil, false
	}
	return o.DeleteAfterDate, true
}

// HasDeleteAfterDate returns a boolean if a field has been set.
func (o *ApiAtlasNetworkPermissionEntryView) HasDeleteAfterDate() bool {
	if o != nil && !isNil(o.DeleteAfterDate) {
		return true
	}

	return false
}

// SetDeleteAfterDate gets a reference to the given time.Time and assigns it to the DeleteAfterDate field.
func (o *ApiAtlasNetworkPermissionEntryView) SetDeleteAfterDate(v time.Time) {
	o.DeleteAfterDate = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApiAtlasNetworkPermissionEntryView) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasNetworkPermissionEntryView) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiAtlasNetworkPermissionEntryView) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiAtlasNetworkPermissionEntryView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ApiAtlasNetworkPermissionEntryView) GetIpAddress() string {
	if o == nil || isNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasNetworkPermissionEntryView) GetIpAddressOk() (*string, bool) {
	if o == nil || isNil(o.IpAddress) {
    return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ApiAtlasNetworkPermissionEntryView) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ApiAtlasNetworkPermissionEntryView) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLinks returns the Links field value
func (o *ApiAtlasNetworkPermissionEntryView) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasNetworkPermissionEntryView) GetLinksOk() ([]Link, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ApiAtlasNetworkPermissionEntryView) SetLinks(v []Link) {
	o.Links = v
}

func (o ApiAtlasNetworkPermissionEntryView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AwsSecurityGroup) {
		toSerialize["awsSecurityGroup"] = o.AwsSecurityGroup
	}
	if !isNil(o.CidrBlock) {
		toSerialize["cidrBlock"] = o.CidrBlock
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.DeleteAfterDate) {
		toSerialize["deleteAfterDate"] = o.DeleteAfterDate
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if true {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasNetworkPermissionEntryView struct {
	value *ApiAtlasNetworkPermissionEntryView
	isSet bool
}

func (v NullableApiAtlasNetworkPermissionEntryView) Get() *ApiAtlasNetworkPermissionEntryView {
	return v.value
}

func (v *NullableApiAtlasNetworkPermissionEntryView) Set(val *ApiAtlasNetworkPermissionEntryView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasNetworkPermissionEntryView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasNetworkPermissionEntryView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasNetworkPermissionEntryView(val *ApiAtlasNetworkPermissionEntryView) *NullableApiAtlasNetworkPermissionEntryView {
	return &NullableApiAtlasNetworkPermissionEntryView{value: val, isSet: true}
}

func (v NullableApiAtlasNetworkPermissionEntryView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasNetworkPermissionEntryView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


