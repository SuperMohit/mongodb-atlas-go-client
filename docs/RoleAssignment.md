# RoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **map[string]interface{}** | List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.  This parameter returns an empty object if no custom zones exist. | [optional] [readonly] 
**OrgId** | Pointer to **map[string]interface{}** | List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.  This parameter returns an empty object if no custom zones exist. | [optional] [readonly] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewRoleAssignment

`func NewRoleAssignment() *RoleAssignment`

NewRoleAssignment instantiates a new RoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignmentWithDefaults

`func NewRoleAssignmentWithDefaults() *RoleAssignment`

NewRoleAssignmentWithDefaults instantiates a new RoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *RoleAssignment) GetGroupId() map[string]interface{}`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *RoleAssignment) GetGroupIdOk() (*map[string]interface{}, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *RoleAssignment) SetGroupId(v map[string]interface{})`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *RoleAssignment) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetOrgId

`func (o *RoleAssignment) GetOrgId() map[string]interface{}`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RoleAssignment) GetOrgIdOk() (*map[string]interface{}, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RoleAssignment) SetOrgId(v map[string]interface{})`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *RoleAssignment) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRole

`func (o *RoleAssignment) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleAssignment) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleAssignment) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleAssignment) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


