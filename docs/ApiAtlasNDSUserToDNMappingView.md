# ApiAtlasNDSUserToDNMappingView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapQuery** | **string** | Lightweight Directory Access Protocol (LDAP) query template that inserts the LDAP name that the regular expression matches into an LDAP query Uniform Resource Identifier (URI). The formatting for the query must conform to [RFC 4515](https://datatracker.ietf.org/doc/html/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516). | 
**Match** | **string** | Regular expression that MongoDB Cloud uses to match against the provided Lightweight Directory Access Protocol (LDAP) username. Each parenthesis-enclosed section represents a regular expression capture group that the substitution or &#x60;ldapQuery&#x60; template uses. | 
**Substitution** | **string** | Lightweight Directory Access Protocol (LDAP) Distinguished Name (DN) template that converts the LDAP username that matches regular expression in the *match* parameter into an LDAP Distinguished Name (DN). | 

## Methods

### NewApiAtlasNDSUserToDNMappingView

`func NewApiAtlasNDSUserToDNMappingView(ldapQuery string, match string, substitution string, ) *ApiAtlasNDSUserToDNMappingView`

NewApiAtlasNDSUserToDNMappingView instantiates a new ApiAtlasNDSUserToDNMappingView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasNDSUserToDNMappingViewWithDefaults

`func NewApiAtlasNDSUserToDNMappingViewWithDefaults() *ApiAtlasNDSUserToDNMappingView`

NewApiAtlasNDSUserToDNMappingViewWithDefaults instantiates a new ApiAtlasNDSUserToDNMappingView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdapQuery

`func (o *ApiAtlasNDSUserToDNMappingView) GetLdapQuery() string`

GetLdapQuery returns the LdapQuery field if non-nil, zero value otherwise.

### GetLdapQueryOk

`func (o *ApiAtlasNDSUserToDNMappingView) GetLdapQueryOk() (*string, bool)`

GetLdapQueryOk returns a tuple with the LdapQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapQuery

`func (o *ApiAtlasNDSUserToDNMappingView) SetLdapQuery(v string)`

SetLdapQuery sets LdapQuery field to given value.


### GetMatch

`func (o *ApiAtlasNDSUserToDNMappingView) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *ApiAtlasNDSUserToDNMappingView) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *ApiAtlasNDSUserToDNMappingView) SetMatch(v string)`

SetMatch sets Match field to given value.


### GetSubstitution

`func (o *ApiAtlasNDSUserToDNMappingView) GetSubstitution() string`

GetSubstitution returns the Substitution field if non-nil, zero value otherwise.

### GetSubstitutionOk

`func (o *ApiAtlasNDSUserToDNMappingView) GetSubstitutionOk() (*string, bool)`

GetSubstitutionOk returns a tuple with the Substitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitution

`func (o *ApiAtlasNDSUserToDNMappingView) SetSubstitution(v string)`

SetSubstitution sets Substitution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


