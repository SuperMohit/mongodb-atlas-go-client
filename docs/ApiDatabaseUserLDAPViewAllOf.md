# ApiDatabaseUserLDAPViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapAuthType** | Pointer to **string** | Part of the Lightweight Directory Access Protocol (LDAP) record that the database uses to authenticate this database user on the LDAP host. | [optional] [default to "NONE"]
**Username** | Pointer to **string** | Human-readable label that represents the user that authenticates to MongoDB. This must be formatted as a [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name.  | [optional] 

## Methods

### NewApiDatabaseUserLDAPViewAllOf

`func NewApiDatabaseUserLDAPViewAllOf() *ApiDatabaseUserLDAPViewAllOf`

NewApiDatabaseUserLDAPViewAllOf instantiates a new ApiDatabaseUserLDAPViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDatabaseUserLDAPViewAllOfWithDefaults

`func NewApiDatabaseUserLDAPViewAllOfWithDefaults() *ApiDatabaseUserLDAPViewAllOf`

NewApiDatabaseUserLDAPViewAllOfWithDefaults instantiates a new ApiDatabaseUserLDAPViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdapAuthType

`func (o *ApiDatabaseUserLDAPViewAllOf) GetLdapAuthType() string`

GetLdapAuthType returns the LdapAuthType field if non-nil, zero value otherwise.

### GetLdapAuthTypeOk

`func (o *ApiDatabaseUserLDAPViewAllOf) GetLdapAuthTypeOk() (*string, bool)`

GetLdapAuthTypeOk returns a tuple with the LdapAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAuthType

`func (o *ApiDatabaseUserLDAPViewAllOf) SetLdapAuthType(v string)`

SetLdapAuthType sets LdapAuthType field to given value.

### HasLdapAuthType

`func (o *ApiDatabaseUserLDAPViewAllOf) HasLdapAuthType() bool`

HasLdapAuthType returns a boolean if a field has been set.

### GetUsername

`func (o *ApiDatabaseUserLDAPViewAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiDatabaseUserLDAPViewAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiDatabaseUserLDAPViewAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiDatabaseUserLDAPViewAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


