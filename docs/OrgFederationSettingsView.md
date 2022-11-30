# OrgFederationSettingsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FederatedDomains** | **[]string** | List of domains associated with the organization&#39;s identity provider. | 
**HasRoleMappings** | Pointer to **bool** | Flag that indicates whether this organization has role mappings configured. | [optional] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this federation. | [optional] [readonly] 
**IdentityProviderId** | Pointer to **string** | Unique 20-hexadecimal digit string that identifies the identity provider connected to this organization. | [optional] 
**IdentityProviderStatus** | Pointer to **string** | String enum that indicates whether the identity provider is active. | [optional] 

## Methods

### NewOrgFederationSettingsView

`func NewOrgFederationSettingsView(federatedDomains []string, ) *OrgFederationSettingsView`

NewOrgFederationSettingsView instantiates a new OrgFederationSettingsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgFederationSettingsViewWithDefaults

`func NewOrgFederationSettingsViewWithDefaults() *OrgFederationSettingsView`

NewOrgFederationSettingsViewWithDefaults instantiates a new OrgFederationSettingsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFederatedDomains

`func (o *OrgFederationSettingsView) GetFederatedDomains() []string`

GetFederatedDomains returns the FederatedDomains field if non-nil, zero value otherwise.

### GetFederatedDomainsOk

`func (o *OrgFederationSettingsView) GetFederatedDomainsOk() (*[]string, bool)`

GetFederatedDomainsOk returns a tuple with the FederatedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedDomains

`func (o *OrgFederationSettingsView) SetFederatedDomains(v []string)`

SetFederatedDomains sets FederatedDomains field to given value.


### GetHasRoleMappings

`func (o *OrgFederationSettingsView) GetHasRoleMappings() bool`

GetHasRoleMappings returns the HasRoleMappings field if non-nil, zero value otherwise.

### GetHasRoleMappingsOk

`func (o *OrgFederationSettingsView) GetHasRoleMappingsOk() (*bool, bool)`

GetHasRoleMappingsOk returns a tuple with the HasRoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRoleMappings

`func (o *OrgFederationSettingsView) SetHasRoleMappings(v bool)`

SetHasRoleMappings sets HasRoleMappings field to given value.

### HasHasRoleMappings

`func (o *OrgFederationSettingsView) HasHasRoleMappings() bool`

HasHasRoleMappings returns a boolean if a field has been set.

### GetId

`func (o *OrgFederationSettingsView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgFederationSettingsView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgFederationSettingsView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrgFederationSettingsView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityProviderId

`func (o *OrgFederationSettingsView) GetIdentityProviderId() string`

GetIdentityProviderId returns the IdentityProviderId field if non-nil, zero value otherwise.

### GetIdentityProviderIdOk

`func (o *OrgFederationSettingsView) GetIdentityProviderIdOk() (*string, bool)`

GetIdentityProviderIdOk returns a tuple with the IdentityProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderId

`func (o *OrgFederationSettingsView) SetIdentityProviderId(v string)`

SetIdentityProviderId sets IdentityProviderId field to given value.

### HasIdentityProviderId

`func (o *OrgFederationSettingsView) HasIdentityProviderId() bool`

HasIdentityProviderId returns a boolean if a field has been set.

### GetIdentityProviderStatus

`func (o *OrgFederationSettingsView) GetIdentityProviderStatus() string`

GetIdentityProviderStatus returns the IdentityProviderStatus field if non-nil, zero value otherwise.

### GetIdentityProviderStatusOk

`func (o *OrgFederationSettingsView) GetIdentityProviderStatusOk() (*string, bool)`

GetIdentityProviderStatusOk returns a tuple with the IdentityProviderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderStatus

`func (o *OrgFederationSettingsView) SetIdentityProviderStatus(v string)`

SetIdentityProviderStatus sets IdentityProviderStatus field to given value.

### HasIdentityProviderStatus

`func (o *OrgFederationSettingsView) HasIdentityProviderStatus() bool`

HasIdentityProviderStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


