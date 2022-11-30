# ApiDatabaseUserX509View

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
**Username** | **string** | Human-readable label that represents the user that authenticates to MongoDB. This must be formatted as a [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name.  | 
**X509Type** | **string** | X.509 method that MongoDB Cloud uses to authenticate the database user. - For MongoDB Cloud-managed X.509, specify &#x60;MANAGED&#x60;. - For self-managed X.509, specify &#x60;CUSTOMER&#x60;. Users created with the &#x60;CUSTOMER&#x60; method require a Common Name (CN) in the **username** parameter. You must create externally authenticated users on the &#x60;$external&#x60; database. | [default to "NONE"]

## Methods

### NewApiDatabaseUserX509View

`func NewApiDatabaseUserX509View(databaseName string, groupId string, roles []ApiAtlasRoleView, username string, x509Type string, ) *ApiDatabaseUserX509View`

NewApiDatabaseUserX509View instantiates a new ApiDatabaseUserX509View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDatabaseUserX509ViewWithDefaults

`func NewApiDatabaseUserX509ViewWithDefaults() *ApiDatabaseUserX509View`

NewApiDatabaseUserX509ViewWithDefaults instantiates a new ApiDatabaseUserX509View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *ApiDatabaseUserX509View) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ApiDatabaseUserX509View) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ApiDatabaseUserX509View) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDeleteAfterDate

`func (o *ApiDatabaseUserX509View) GetDeleteAfterDate() string`

GetDeleteAfterDate returns the DeleteAfterDate field if non-nil, zero value otherwise.

### GetDeleteAfterDateOk

`func (o *ApiDatabaseUserX509View) GetDeleteAfterDateOk() (*string, bool)`

GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfterDate

`func (o *ApiDatabaseUserX509View) SetDeleteAfterDate(v string)`

SetDeleteAfterDate sets DeleteAfterDate field to given value.

### HasDeleteAfterDate

`func (o *ApiDatabaseUserX509View) HasDeleteAfterDate() bool`

HasDeleteAfterDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiDatabaseUserX509View) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiDatabaseUserX509View) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiDatabaseUserX509View) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetLabels

`func (o *ApiDatabaseUserX509View) GetLabels() []ApiAtlasNDSLabelView`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiDatabaseUserX509View) GetLabelsOk() (*[]ApiAtlasNDSLabelView, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiDatabaseUserX509View) SetLabels(v []ApiAtlasNDSLabelView)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiDatabaseUserX509View) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLinks

`func (o *ApiDatabaseUserX509View) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiDatabaseUserX509View) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiDatabaseUserX509View) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiDatabaseUserX509View) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRoles

`func (o *ApiDatabaseUserX509View) GetRoles() []ApiAtlasRoleView`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiDatabaseUserX509View) GetRolesOk() (*[]ApiAtlasRoleView, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiDatabaseUserX509View) SetRoles(v []ApiAtlasRoleView)`

SetRoles sets Roles field to given value.


### GetScopes

`func (o *ApiDatabaseUserX509View) GetScopes() []ApiAtlasUserScopeView`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiDatabaseUserX509View) GetScopesOk() (*[]ApiAtlasUserScopeView, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiDatabaseUserX509View) SetScopes(v []ApiAtlasUserScopeView)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApiDatabaseUserX509View) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetUsername

`func (o *ApiDatabaseUserX509View) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiDatabaseUserX509View) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiDatabaseUserX509View) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetX509Type

`func (o *ApiDatabaseUserX509View) GetX509Type() string`

GetX509Type returns the X509Type field if non-nil, zero value otherwise.

### GetX509TypeOk

`func (o *ApiDatabaseUserX509View) GetX509TypeOk() (*string, bool)`

GetX509TypeOk returns a tuple with the X509Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Type

`func (o *ApiDatabaseUserX509View) SetX509Type(v string)`

SetX509Type sets X509Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


