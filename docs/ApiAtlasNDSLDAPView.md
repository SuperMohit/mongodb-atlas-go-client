# ApiAtlasNDSLDAPView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationEnabled** | Pointer to **bool** | Flag that indicates whether users can authenticate using an Lightweight Directory Access Protocol (LDAP) host. | [optional] 
**AuthorizationEnabled** | Pointer to **bool** | Flag that indicates whether users can authorize access to MongoDB Cloud resources using an Lightweight Directory Access Protocol (LDAP) host. | [optional] 
**AuthzQueryTemplate** | Pointer to **string** | Lightweight Directory Access Protocol (LDAP) query template that MongoDB Cloud runs to obtain the LDAP groups associated with the authenticated user. MongoDB Cloud uses this parameter only for user authorization. Use the &#x60;{USER}&#x60; placeholder in the Uniform Resource Locator (URL) to substitute the authenticated username. The query relates to the host specified with the hostname. Format this query according to [RFC 4515](https://tools.ietf.org/search/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516). | [optional] [default to "{USER}?memberOf?base"]
**BindPassword** | Pointer to **string** | Password that MongoDB Cloud uses to authenticate the **bindUsername**. | [optional] 
**BindUsername** | Pointer to **string** | Full Distinguished Name (DN) of the Lightweight Directory Access Protocol (LDAP) user that MongoDB Cloud uses to connect to the LDAP host. LDAP distinguished names must be formatted according to RFC 2253. | [optional] 
**CaCertificate** | Pointer to **string** | Certificate Authority (CA) certificate that MongoDB Cloud uses to verify the identity of the Lightweight Directory Access Protocol (LDAP) host. MongoDB Cloud allows self-signed certificates. To delete an assigned value, pass an empty string: &#x60;\&quot;caCertificate\&quot;: \&quot;\&quot;&#x60; | [optional] 
**Hostname** | Pointer to **string** | Human-readable label that identifies the hostname or Internet Protocol (IP) address of the Lightweight Directory Access Protocol (LDAP) host. This host must have access to the internet or have a Virtual Private Cloud (VPC) peering connection to your cluster. | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Port** | Pointer to **int32** | Port to which the Lightweight Directory Access Protocol (LDAP) host listens for client connections. | [optional] [default to 636]
**UserToDNMapping** | Pointer to [**[]ApiAtlasNDSUserToDNMappingView**](ApiAtlasNDSUserToDNMappingView.md) | User-to-Distinguished Name (DN) map that MongoDB Cloud uses to transform a Lightweight Directory Access Protocol (LDAP) username into an LDAP DN. | [optional] 

## Methods

### NewApiAtlasNDSLDAPView

`func NewApiAtlasNDSLDAPView(links []Link, ) *ApiAtlasNDSLDAPView`

NewApiAtlasNDSLDAPView instantiates a new ApiAtlasNDSLDAPView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasNDSLDAPViewWithDefaults

`func NewApiAtlasNDSLDAPViewWithDefaults() *ApiAtlasNDSLDAPView`

NewApiAtlasNDSLDAPViewWithDefaults instantiates a new ApiAtlasNDSLDAPView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationEnabled

`func (o *ApiAtlasNDSLDAPView) GetAuthenticationEnabled() bool`

GetAuthenticationEnabled returns the AuthenticationEnabled field if non-nil, zero value otherwise.

### GetAuthenticationEnabledOk

`func (o *ApiAtlasNDSLDAPView) GetAuthenticationEnabledOk() (*bool, bool)`

GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEnabled

`func (o *ApiAtlasNDSLDAPView) SetAuthenticationEnabled(v bool)`

SetAuthenticationEnabled sets AuthenticationEnabled field to given value.

### HasAuthenticationEnabled

`func (o *ApiAtlasNDSLDAPView) HasAuthenticationEnabled() bool`

HasAuthenticationEnabled returns a boolean if a field has been set.

### GetAuthorizationEnabled

`func (o *ApiAtlasNDSLDAPView) GetAuthorizationEnabled() bool`

GetAuthorizationEnabled returns the AuthorizationEnabled field if non-nil, zero value otherwise.

### GetAuthorizationEnabledOk

`func (o *ApiAtlasNDSLDAPView) GetAuthorizationEnabledOk() (*bool, bool)`

GetAuthorizationEnabledOk returns a tuple with the AuthorizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEnabled

`func (o *ApiAtlasNDSLDAPView) SetAuthorizationEnabled(v bool)`

SetAuthorizationEnabled sets AuthorizationEnabled field to given value.

### HasAuthorizationEnabled

`func (o *ApiAtlasNDSLDAPView) HasAuthorizationEnabled() bool`

HasAuthorizationEnabled returns a boolean if a field has been set.

### GetAuthzQueryTemplate

`func (o *ApiAtlasNDSLDAPView) GetAuthzQueryTemplate() string`

GetAuthzQueryTemplate returns the AuthzQueryTemplate field if non-nil, zero value otherwise.

### GetAuthzQueryTemplateOk

`func (o *ApiAtlasNDSLDAPView) GetAuthzQueryTemplateOk() (*string, bool)`

GetAuthzQueryTemplateOk returns a tuple with the AuthzQueryTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthzQueryTemplate

`func (o *ApiAtlasNDSLDAPView) SetAuthzQueryTemplate(v string)`

SetAuthzQueryTemplate sets AuthzQueryTemplate field to given value.

### HasAuthzQueryTemplate

`func (o *ApiAtlasNDSLDAPView) HasAuthzQueryTemplate() bool`

HasAuthzQueryTemplate returns a boolean if a field has been set.

### GetBindPassword

`func (o *ApiAtlasNDSLDAPView) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *ApiAtlasNDSLDAPView) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *ApiAtlasNDSLDAPView) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.

### HasBindPassword

`func (o *ApiAtlasNDSLDAPView) HasBindPassword() bool`

HasBindPassword returns a boolean if a field has been set.

### GetBindUsername

`func (o *ApiAtlasNDSLDAPView) GetBindUsername() string`

GetBindUsername returns the BindUsername field if non-nil, zero value otherwise.

### GetBindUsernameOk

`func (o *ApiAtlasNDSLDAPView) GetBindUsernameOk() (*string, bool)`

GetBindUsernameOk returns a tuple with the BindUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindUsername

`func (o *ApiAtlasNDSLDAPView) SetBindUsername(v string)`

SetBindUsername sets BindUsername field to given value.

### HasBindUsername

`func (o *ApiAtlasNDSLDAPView) HasBindUsername() bool`

HasBindUsername returns a boolean if a field has been set.

### GetCaCertificate

`func (o *ApiAtlasNDSLDAPView) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *ApiAtlasNDSLDAPView) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *ApiAtlasNDSLDAPView) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *ApiAtlasNDSLDAPView) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetHostname

`func (o *ApiAtlasNDSLDAPView) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApiAtlasNDSLDAPView) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApiAtlasNDSLDAPView) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApiAtlasNDSLDAPView) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasNDSLDAPView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasNDSLDAPView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasNDSLDAPView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetPort

`func (o *ApiAtlasNDSLDAPView) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApiAtlasNDSLDAPView) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApiAtlasNDSLDAPView) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApiAtlasNDSLDAPView) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUserToDNMapping

`func (o *ApiAtlasNDSLDAPView) GetUserToDNMapping() []ApiAtlasNDSUserToDNMappingView`

GetUserToDNMapping returns the UserToDNMapping field if non-nil, zero value otherwise.

### GetUserToDNMappingOk

`func (o *ApiAtlasNDSLDAPView) GetUserToDNMappingOk() (*[]ApiAtlasNDSUserToDNMappingView, bool)`

GetUserToDNMappingOk returns a tuple with the UserToDNMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToDNMapping

`func (o *ApiAtlasNDSLDAPView) SetUserToDNMapping(v []ApiAtlasNDSUserToDNMappingView)`

SetUserToDNMapping sets UserToDNMapping field to given value.

### HasUserToDNMapping

`func (o *ApiAtlasNDSLDAPView) HasUserToDNMapping() bool`

HasUserToDNMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


