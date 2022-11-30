# ApiGroupInvitationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud sent the invitation. This parameter expresses its value in ISO 8601 format in UTC. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud expires the invitation. This parameter expresses its value in ISO 8601 format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**GroupName** | Pointer to **string** | Human-readable label that identifies the project to which you invited the MongoDB Cloud user. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the invitation. | [optional] [readonly] 
**InviterUsername** | Pointer to **string** | Email address of the MongoDB Cloud user who sent the invitation. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Roles** | Pointer to **[]string** | One or more organization or project level roles to assign to the MongoDB Cloud user. | [optional] 
**Username** | Pointer to **string** | Email address of the MongoDB Cloud user invited to join the project. | [optional] [readonly] 

## Methods

### NewApiGroupInvitationView

`func NewApiGroupInvitationView(links []Link, ) *ApiGroupInvitationView`

NewApiGroupInvitationView instantiates a new ApiGroupInvitationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiGroupInvitationViewWithDefaults

`func NewApiGroupInvitationViewWithDefaults() *ApiGroupInvitationView`

NewApiGroupInvitationViewWithDefaults instantiates a new ApiGroupInvitationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ApiGroupInvitationView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiGroupInvitationView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiGroupInvitationView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiGroupInvitationView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiGroupInvitationView) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiGroupInvitationView) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiGroupInvitationView) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiGroupInvitationView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiGroupInvitationView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiGroupInvitationView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiGroupInvitationView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiGroupInvitationView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *ApiGroupInvitationView) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ApiGroupInvitationView) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ApiGroupInvitationView) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ApiGroupInvitationView) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetId

`func (o *ApiGroupInvitationView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiGroupInvitationView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiGroupInvitationView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiGroupInvitationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInviterUsername

`func (o *ApiGroupInvitationView) GetInviterUsername() string`

GetInviterUsername returns the InviterUsername field if non-nil, zero value otherwise.

### GetInviterUsernameOk

`func (o *ApiGroupInvitationView) GetInviterUsernameOk() (*string, bool)`

GetInviterUsernameOk returns a tuple with the InviterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviterUsername

`func (o *ApiGroupInvitationView) SetInviterUsername(v string)`

SetInviterUsername sets InviterUsername field to given value.

### HasInviterUsername

`func (o *ApiGroupInvitationView) HasInviterUsername() bool`

HasInviterUsername returns a boolean if a field has been set.

### GetLinks

`func (o *ApiGroupInvitationView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiGroupInvitationView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiGroupInvitationView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetRoles

`func (o *ApiGroupInvitationView) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiGroupInvitationView) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiGroupInvitationView) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiGroupInvitationView) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUsername

`func (o *ApiGroupInvitationView) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiGroupInvitationView) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiGroupInvitationView) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiGroupInvitationView) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


