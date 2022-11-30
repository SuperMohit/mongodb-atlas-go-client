# ApiOrganizationInvitationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud sent the invitation. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Date and time when the invitation from MongoDB Cloud expires. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this organization. | [optional] [readonly] 
**InviterUsername** | Pointer to **string** | Email address of the MongoDB Cloud user who sent the invitation to join the organization. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization. | [optional] [readonly] 
**OrgName** | **string** | Human-readable label that identifies this organization. | 
**Roles** | Pointer to **[]string** | One or more organization or project level roles to assign to the MongoDB Cloud user. | [optional] 
**TeamIds** | Pointer to **[]string** | List of unique 24-hexadecimal digit strings that identifies each team. | [optional] [readonly] 
**Username** | Pointer to **string** | Email address of the MongoDB Cloud user invited to join the organization. | [optional] 

## Methods

### NewApiOrganizationInvitationView

`func NewApiOrganizationInvitationView(links []Link, orgName string, ) *ApiOrganizationInvitationView`

NewApiOrganizationInvitationView instantiates a new ApiOrganizationInvitationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOrganizationInvitationViewWithDefaults

`func NewApiOrganizationInvitationViewWithDefaults() *ApiOrganizationInvitationView`

NewApiOrganizationInvitationViewWithDefaults instantiates a new ApiOrganizationInvitationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ApiOrganizationInvitationView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiOrganizationInvitationView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiOrganizationInvitationView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiOrganizationInvitationView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiOrganizationInvitationView) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiOrganizationInvitationView) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiOrganizationInvitationView) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiOrganizationInvitationView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *ApiOrganizationInvitationView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiOrganizationInvitationView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiOrganizationInvitationView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiOrganizationInvitationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInviterUsername

`func (o *ApiOrganizationInvitationView) GetInviterUsername() string`

GetInviterUsername returns the InviterUsername field if non-nil, zero value otherwise.

### GetInviterUsernameOk

`func (o *ApiOrganizationInvitationView) GetInviterUsernameOk() (*string, bool)`

GetInviterUsernameOk returns a tuple with the InviterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviterUsername

`func (o *ApiOrganizationInvitationView) SetInviterUsername(v string)`

SetInviterUsername sets InviterUsername field to given value.

### HasInviterUsername

`func (o *ApiOrganizationInvitationView) HasInviterUsername() bool`

HasInviterUsername returns a boolean if a field has been set.

### GetLinks

`func (o *ApiOrganizationInvitationView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiOrganizationInvitationView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiOrganizationInvitationView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetOrgId

`func (o *ApiOrganizationInvitationView) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ApiOrganizationInvitationView) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ApiOrganizationInvitationView) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ApiOrganizationInvitationView) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *ApiOrganizationInvitationView) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ApiOrganizationInvitationView) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ApiOrganizationInvitationView) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetRoles

`func (o *ApiOrganizationInvitationView) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiOrganizationInvitationView) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiOrganizationInvitationView) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiOrganizationInvitationView) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTeamIds

`func (o *ApiOrganizationInvitationView) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *ApiOrganizationInvitationView) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *ApiOrganizationInvitationView) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *ApiOrganizationInvitationView) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetUsername

`func (o *ApiOrganizationInvitationView) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiOrganizationInvitationView) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiOrganizationInvitationView) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiOrganizationInvitationView) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


