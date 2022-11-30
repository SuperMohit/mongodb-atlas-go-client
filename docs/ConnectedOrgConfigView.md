# ConnectedOrgConfigView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainAllowList** | Pointer to **[]string** | Approved domains that restrict users who can join the organization based on their email address. | [optional] 
**DomainRestrictionEnabled** | **bool** | Value that indicates whether domain restriction is enabled for this connected org. | 
**IdentityProviderId** | **string** | Unique 20-hexadecimal digit string that identifies the identity provider that this connected org config is associated with. | 
**OrgId** | **string** | Unique 24-hexadecimal digit string that identifies the connected organization configuration. | [readonly] 
**PostAuthRoleGrants** | Pointer to **[]string** | Atlas roles that are granted to a user in this organization after authenticating. | [optional] 
**RoleMappings** | Pointer to [**[]RoleMappingView**](RoleMappingView.md) | Role mappings that are configured in this organization. | [optional] 
**UserConflicts** | Pointer to [**[]FederatedUserView**](FederatedUserView.md) | List that contains the users who have an email address that doesn&#39;t match any domain on the allowed list. | [optional] 

## Methods

### NewConnectedOrgConfigView

`func NewConnectedOrgConfigView(domainRestrictionEnabled bool, identityProviderId string, orgId string, ) *ConnectedOrgConfigView`

NewConnectedOrgConfigView instantiates a new ConnectedOrgConfigView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedOrgConfigViewWithDefaults

`func NewConnectedOrgConfigViewWithDefaults() *ConnectedOrgConfigView`

NewConnectedOrgConfigViewWithDefaults instantiates a new ConnectedOrgConfigView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainAllowList

`func (o *ConnectedOrgConfigView) GetDomainAllowList() []string`

GetDomainAllowList returns the DomainAllowList field if non-nil, zero value otherwise.

### GetDomainAllowListOk

`func (o *ConnectedOrgConfigView) GetDomainAllowListOk() (*[]string, bool)`

GetDomainAllowListOk returns a tuple with the DomainAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAllowList

`func (o *ConnectedOrgConfigView) SetDomainAllowList(v []string)`

SetDomainAllowList sets DomainAllowList field to given value.

### HasDomainAllowList

`func (o *ConnectedOrgConfigView) HasDomainAllowList() bool`

HasDomainAllowList returns a boolean if a field has been set.

### GetDomainRestrictionEnabled

`func (o *ConnectedOrgConfigView) GetDomainRestrictionEnabled() bool`

GetDomainRestrictionEnabled returns the DomainRestrictionEnabled field if non-nil, zero value otherwise.

### GetDomainRestrictionEnabledOk

`func (o *ConnectedOrgConfigView) GetDomainRestrictionEnabledOk() (*bool, bool)`

GetDomainRestrictionEnabledOk returns a tuple with the DomainRestrictionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainRestrictionEnabled

`func (o *ConnectedOrgConfigView) SetDomainRestrictionEnabled(v bool)`

SetDomainRestrictionEnabled sets DomainRestrictionEnabled field to given value.


### GetIdentityProviderId

`func (o *ConnectedOrgConfigView) GetIdentityProviderId() string`

GetIdentityProviderId returns the IdentityProviderId field if non-nil, zero value otherwise.

### GetIdentityProviderIdOk

`func (o *ConnectedOrgConfigView) GetIdentityProviderIdOk() (*string, bool)`

GetIdentityProviderIdOk returns a tuple with the IdentityProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderId

`func (o *ConnectedOrgConfigView) SetIdentityProviderId(v string)`

SetIdentityProviderId sets IdentityProviderId field to given value.


### GetOrgId

`func (o *ConnectedOrgConfigView) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ConnectedOrgConfigView) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ConnectedOrgConfigView) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetPostAuthRoleGrants

`func (o *ConnectedOrgConfigView) GetPostAuthRoleGrants() []string`

GetPostAuthRoleGrants returns the PostAuthRoleGrants field if non-nil, zero value otherwise.

### GetPostAuthRoleGrantsOk

`func (o *ConnectedOrgConfigView) GetPostAuthRoleGrantsOk() (*[]string, bool)`

GetPostAuthRoleGrantsOk returns a tuple with the PostAuthRoleGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAuthRoleGrants

`func (o *ConnectedOrgConfigView) SetPostAuthRoleGrants(v []string)`

SetPostAuthRoleGrants sets PostAuthRoleGrants field to given value.

### HasPostAuthRoleGrants

`func (o *ConnectedOrgConfigView) HasPostAuthRoleGrants() bool`

HasPostAuthRoleGrants returns a boolean if a field has been set.

### GetRoleMappings

`func (o *ConnectedOrgConfigView) GetRoleMappings() []RoleMappingView`

GetRoleMappings returns the RoleMappings field if non-nil, zero value otherwise.

### GetRoleMappingsOk

`func (o *ConnectedOrgConfigView) GetRoleMappingsOk() (*[]RoleMappingView, bool)`

GetRoleMappingsOk returns a tuple with the RoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleMappings

`func (o *ConnectedOrgConfigView) SetRoleMappings(v []RoleMappingView)`

SetRoleMappings sets RoleMappings field to given value.

### HasRoleMappings

`func (o *ConnectedOrgConfigView) HasRoleMappings() bool`

HasRoleMappings returns a boolean if a field has been set.

### GetUserConflicts

`func (o *ConnectedOrgConfigView) GetUserConflicts() []FederatedUserView`

GetUserConflicts returns the UserConflicts field if non-nil, zero value otherwise.

### GetUserConflictsOk

`func (o *ConnectedOrgConfigView) GetUserConflictsOk() (*[]FederatedUserView, bool)`

GetUserConflictsOk returns a tuple with the UserConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserConflicts

`func (o *ConnectedOrgConfigView) SetUserConflicts(v []FederatedUserView)`

SetUserConflicts sets UserConflicts field to given value.

### HasUserConflicts

`func (o *ConnectedOrgConfigView) HasUserConflicts() bool`

HasUserConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


