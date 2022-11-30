# ApiAtlasUserSecurityView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerX509** | Pointer to [**ApiAtlasCustomerX509View**](ApiAtlasCustomerX509View.md) |  | [optional] 
**Ldap** | Pointer to [**ApiAtlasNDSLDAPView**](ApiAtlasNDSLDAPView.md) |  | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 

## Methods

### NewApiAtlasUserSecurityView

`func NewApiAtlasUserSecurityView(links []Link, ) *ApiAtlasUserSecurityView`

NewApiAtlasUserSecurityView instantiates a new ApiAtlasUserSecurityView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasUserSecurityViewWithDefaults

`func NewApiAtlasUserSecurityViewWithDefaults() *ApiAtlasUserSecurityView`

NewApiAtlasUserSecurityViewWithDefaults instantiates a new ApiAtlasUserSecurityView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerX509

`func (o *ApiAtlasUserSecurityView) GetCustomerX509() ApiAtlasCustomerX509View`

GetCustomerX509 returns the CustomerX509 field if non-nil, zero value otherwise.

### GetCustomerX509Ok

`func (o *ApiAtlasUserSecurityView) GetCustomerX509Ok() (*ApiAtlasCustomerX509View, bool)`

GetCustomerX509Ok returns a tuple with the CustomerX509 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerX509

`func (o *ApiAtlasUserSecurityView) SetCustomerX509(v ApiAtlasCustomerX509View)`

SetCustomerX509 sets CustomerX509 field to given value.

### HasCustomerX509

`func (o *ApiAtlasUserSecurityView) HasCustomerX509() bool`

HasCustomerX509 returns a boolean if a field has been set.

### GetLdap

`func (o *ApiAtlasUserSecurityView) GetLdap() ApiAtlasNDSLDAPView`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *ApiAtlasUserSecurityView) GetLdapOk() (*ApiAtlasNDSLDAPView, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *ApiAtlasUserSecurityView) SetLdap(v ApiAtlasNDSLDAPView)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *ApiAtlasUserSecurityView) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasUserSecurityView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasUserSecurityView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasUserSecurityView) SetLinks(v []Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


