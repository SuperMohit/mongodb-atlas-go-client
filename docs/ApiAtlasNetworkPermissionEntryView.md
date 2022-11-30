# ApiAtlasNetworkPermissionEntryView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsSecurityGroup** | Pointer to **string** | Unique string of the Amazon Web Services (AWS) security group that you want to add to the project&#39;s IP access list. Your IP access list entry can be one **awsSecurityGroup**, one **cidrBlock**, or one **ipAddress**. You must configure Virtual Private Connection (VPC) peering for your project before you can add an AWS security group to an IP access list. You cannot set AWS security groups as temporary access list entries. Don&#39;t set this parameter if you set **cidrBlock** or **ipAddress**. | [optional] 
**CidrBlock** | Pointer to **string** | Range of IP addresses in Classless Inter-Domain Routing (CIDR) notation that you want to add to the project&#39;s IP access list. Your IP access list entry can be one **awsSecurityGroup**, one **cidrBlock**, or one **ipAddress**. Don&#39;t set this parameter if you set **awsSecurityGroup** or **ipAddress**. | [optional] 
**Comment** | Pointer to **string** | Remark that explains the purpose or scope of this IP access list entry. | [optional] 
**DeleteAfterDate** | Pointer to **time.Time** | Date and time after which MongoDB Cloud deletes the temporary access list entry. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. The date must be later than the current date but no later than one week after you submit this request. The resource returns this parameter if you specified an expiration date when creating this IP access list entry. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that contains the IP access list to which you want to add one or more entries. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | IP address that you want to add to the project&#39;s IP access list. Your IP access list entry can be one **awsSecurityGroup**, one **cidrBlock**, or one **ipAddress**. Don&#39;t set this parameter if you set **awsSecurityGroup** or **cidrBlock**. | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 

## Methods

### NewApiAtlasNetworkPermissionEntryView

`func NewApiAtlasNetworkPermissionEntryView(links []Link, ) *ApiAtlasNetworkPermissionEntryView`

NewApiAtlasNetworkPermissionEntryView instantiates a new ApiAtlasNetworkPermissionEntryView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasNetworkPermissionEntryViewWithDefaults

`func NewApiAtlasNetworkPermissionEntryViewWithDefaults() *ApiAtlasNetworkPermissionEntryView`

NewApiAtlasNetworkPermissionEntryViewWithDefaults instantiates a new ApiAtlasNetworkPermissionEntryView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsSecurityGroup

`func (o *ApiAtlasNetworkPermissionEntryView) GetAwsSecurityGroup() string`

GetAwsSecurityGroup returns the AwsSecurityGroup field if non-nil, zero value otherwise.

### GetAwsSecurityGroupOk

`func (o *ApiAtlasNetworkPermissionEntryView) GetAwsSecurityGroupOk() (*string, bool)`

GetAwsSecurityGroupOk returns a tuple with the AwsSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecurityGroup

`func (o *ApiAtlasNetworkPermissionEntryView) SetAwsSecurityGroup(v string)`

SetAwsSecurityGroup sets AwsSecurityGroup field to given value.

### HasAwsSecurityGroup

`func (o *ApiAtlasNetworkPermissionEntryView) HasAwsSecurityGroup() bool`

HasAwsSecurityGroup returns a boolean if a field has been set.

### GetCidrBlock

`func (o *ApiAtlasNetworkPermissionEntryView) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *ApiAtlasNetworkPermissionEntryView) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *ApiAtlasNetworkPermissionEntryView) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *ApiAtlasNetworkPermissionEntryView) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetComment

`func (o *ApiAtlasNetworkPermissionEntryView) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiAtlasNetworkPermissionEntryView) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiAtlasNetworkPermissionEntryView) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiAtlasNetworkPermissionEntryView) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeleteAfterDate

`func (o *ApiAtlasNetworkPermissionEntryView) GetDeleteAfterDate() time.Time`

GetDeleteAfterDate returns the DeleteAfterDate field if non-nil, zero value otherwise.

### GetDeleteAfterDateOk

`func (o *ApiAtlasNetworkPermissionEntryView) GetDeleteAfterDateOk() (*time.Time, bool)`

GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfterDate

`func (o *ApiAtlasNetworkPermissionEntryView) SetDeleteAfterDate(v time.Time)`

SetDeleteAfterDate sets DeleteAfterDate field to given value.

### HasDeleteAfterDate

`func (o *ApiAtlasNetworkPermissionEntryView) HasDeleteAfterDate() bool`

HasDeleteAfterDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiAtlasNetworkPermissionEntryView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasNetworkPermissionEntryView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasNetworkPermissionEntryView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAtlasNetworkPermissionEntryView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIpAddress

`func (o *ApiAtlasNetworkPermissionEntryView) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ApiAtlasNetworkPermissionEntryView) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ApiAtlasNetworkPermissionEntryView) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ApiAtlasNetworkPermissionEntryView) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasNetworkPermissionEntryView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasNetworkPermissionEntryView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasNetworkPermissionEntryView) SetLinks(v []Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


