# ApiAtlasDatabaseUserViewManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | **string** | MongoDB database against which the MongoDB database user authenticates. MongoDB database users must provide both a username and authentication database to log into MongoDB. | [default to "admin"]
**DeleteAfterDate** | Pointer to **string** | Date and time when MongoDB Cloud deletes the user. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request. | [optional] 
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the project. | 
**Labels** | Pointer to [**[]ApiAtlasNDSLabelView**](ApiAtlasNDSLabelView.md) | List that contains the key-value pairs for tagging and categorizing the MongoDB database user. The labels that you define do not appear in the console. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Roles** | [**[]ApiAtlasRoleView**](ApiAtlasRoleView.md) | List that provides the pairings of one role with one applicable database. | 
**Scopes** | Pointer to [**[]ApiAtlasUserScopeView**](ApiAtlasUserScopeView.md) | List that contains clusters and MongoDB Atlas Data Lakes that this database user can access. If omitted, MongoDB Cloud grants the database user access to all the clusters and MongoDB Atlas Data Lakes in the project. | [optional] 
**Password** | **string** | Alphanumeric string that authenticates this database user against the database specified in &#x60;databaseName&#x60;. To authenticate with SCRAM-SHA, you must specify this parameter. | 
**Username** | **string** | Human-readable label that represents the user that authenticates to MongoDB.  | 
**LdapAuthType** | **string** | Part of the Lightweight Directory Access Protocol (LDAP) record that the database uses to authenticate this database user on the LDAP host. | [default to "NONE"]
**X509Type** | **string** | X.509 method that MongoDB Cloud uses to authenticate the database user. - For MongoDB Cloud-managed X.509, specify &#x60;MANAGED&#x60;. - For self-managed X.509, specify &#x60;CUSTOMER&#x60;. Users created with the &#x60;CUSTOMER&#x60; method require a Common Name (CN) in the **username** parameter. You must create externally authenticated users on the &#x60;$external&#x60; database. | [default to "NONE"]
**AwsIAMType** | **string** | Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user&#39;s role. | [default to "NONE"]

## Methods

### NewApiAtlasDatabaseUserViewManual

`func NewApiAtlasDatabaseUserViewManual(databaseName string, groupId string, roles []ApiAtlasRoleView, password string, username string, ldapAuthType string, x509Type string, awsIAMType string, ) *ApiAtlasDatabaseUserViewManual`

NewApiAtlasDatabaseUserViewManual instantiates a new ApiAtlasDatabaseUserViewManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDatabaseUserViewManualWithDefaults

`func NewApiAtlasDatabaseUserViewManualWithDefaults() *ApiAtlasDatabaseUserViewManual`

NewApiAtlasDatabaseUserViewManualWithDefaults instantiates a new ApiAtlasDatabaseUserViewManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *ApiAtlasDatabaseUserViewManual) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ApiAtlasDatabaseUserViewManual) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ApiAtlasDatabaseUserViewManual) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDeleteAfterDate

`func (o *ApiAtlasDatabaseUserViewManual) GetDeleteAfterDate() string`

GetDeleteAfterDate returns the DeleteAfterDate field if non-nil, zero value otherwise.

### GetDeleteAfterDateOk

`func (o *ApiAtlasDatabaseUserViewManual) GetDeleteAfterDateOk() (*string, bool)`

GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfterDate

`func (o *ApiAtlasDatabaseUserViewManual) SetDeleteAfterDate(v string)`

SetDeleteAfterDate sets DeleteAfterDate field to given value.

### HasDeleteAfterDate

`func (o *ApiAtlasDatabaseUserViewManual) HasDeleteAfterDate() bool`

HasDeleteAfterDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiAtlasDatabaseUserViewManual) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasDatabaseUserViewManual) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasDatabaseUserViewManual) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetLabels

`func (o *ApiAtlasDatabaseUserViewManual) GetLabels() []ApiAtlasNDSLabelView`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiAtlasDatabaseUserViewManual) GetLabelsOk() (*[]ApiAtlasNDSLabelView, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiAtlasDatabaseUserViewManual) SetLabels(v []ApiAtlasNDSLabelView)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiAtlasDatabaseUserViewManual) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasDatabaseUserViewManual) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasDatabaseUserViewManual) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasDatabaseUserViewManual) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiAtlasDatabaseUserViewManual) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRoles

`func (o *ApiAtlasDatabaseUserViewManual) GetRoles() []ApiAtlasRoleView`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiAtlasDatabaseUserViewManual) GetRolesOk() (*[]ApiAtlasRoleView, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiAtlasDatabaseUserViewManual) SetRoles(v []ApiAtlasRoleView)`

SetRoles sets Roles field to given value.


### GetScopes

`func (o *ApiAtlasDatabaseUserViewManual) GetScopes() []ApiAtlasUserScopeView`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiAtlasDatabaseUserViewManual) GetScopesOk() (*[]ApiAtlasUserScopeView, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiAtlasDatabaseUserViewManual) SetScopes(v []ApiAtlasUserScopeView)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApiAtlasDatabaseUserViewManual) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPassword

`func (o *ApiAtlasDatabaseUserViewManual) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiAtlasDatabaseUserViewManual) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiAtlasDatabaseUserViewManual) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *ApiAtlasDatabaseUserViewManual) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiAtlasDatabaseUserViewManual) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiAtlasDatabaseUserViewManual) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetLdapAuthType

`func (o *ApiAtlasDatabaseUserViewManual) GetLdapAuthType() string`

GetLdapAuthType returns the LdapAuthType field if non-nil, zero value otherwise.

### GetLdapAuthTypeOk

`func (o *ApiAtlasDatabaseUserViewManual) GetLdapAuthTypeOk() (*string, bool)`

GetLdapAuthTypeOk returns a tuple with the LdapAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAuthType

`func (o *ApiAtlasDatabaseUserViewManual) SetLdapAuthType(v string)`

SetLdapAuthType sets LdapAuthType field to given value.


### GetX509Type

`func (o *ApiAtlasDatabaseUserViewManual) GetX509Type() string`

GetX509Type returns the X509Type field if non-nil, zero value otherwise.

### GetX509TypeOk

`func (o *ApiAtlasDatabaseUserViewManual) GetX509TypeOk() (*string, bool)`

GetX509TypeOk returns a tuple with the X509Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Type

`func (o *ApiAtlasDatabaseUserViewManual) SetX509Type(v string)`

SetX509Type sets X509Type field to given value.


### GetAwsIAMType

`func (o *ApiAtlasDatabaseUserViewManual) GetAwsIAMType() string`

GetAwsIAMType returns the AwsIAMType field if non-nil, zero value otherwise.

### GetAwsIAMTypeOk

`func (o *ApiAtlasDatabaseUserViewManual) GetAwsIAMTypeOk() (*string, bool)`

GetAwsIAMTypeOk returns a tuple with the AwsIAMType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIAMType

`func (o *ApiAtlasDatabaseUserViewManual) SetAwsIAMType(v string)`

SetAwsIAMType sets AwsIAMType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


