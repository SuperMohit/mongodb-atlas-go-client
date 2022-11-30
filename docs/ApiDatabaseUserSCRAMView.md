# ApiDatabaseUserSCRAMView

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
**Username** | **string** | Human-readable label that represents the user that authenticates to MongoDB. This must be formatted as a [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name.  | 

## Methods

### NewApiDatabaseUserSCRAMView

`func NewApiDatabaseUserSCRAMView(databaseName string, groupId string, roles []ApiAtlasRoleView, password string, username string, ) *ApiDatabaseUserSCRAMView`

NewApiDatabaseUserSCRAMView instantiates a new ApiDatabaseUserSCRAMView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDatabaseUserSCRAMViewWithDefaults

`func NewApiDatabaseUserSCRAMViewWithDefaults() *ApiDatabaseUserSCRAMView`

NewApiDatabaseUserSCRAMViewWithDefaults instantiates a new ApiDatabaseUserSCRAMView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *ApiDatabaseUserSCRAMView) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ApiDatabaseUserSCRAMView) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ApiDatabaseUserSCRAMView) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDeleteAfterDate

`func (o *ApiDatabaseUserSCRAMView) GetDeleteAfterDate() string`

GetDeleteAfterDate returns the DeleteAfterDate field if non-nil, zero value otherwise.

### GetDeleteAfterDateOk

`func (o *ApiDatabaseUserSCRAMView) GetDeleteAfterDateOk() (*string, bool)`

GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfterDate

`func (o *ApiDatabaseUserSCRAMView) SetDeleteAfterDate(v string)`

SetDeleteAfterDate sets DeleteAfterDate field to given value.

### HasDeleteAfterDate

`func (o *ApiDatabaseUserSCRAMView) HasDeleteAfterDate() bool`

HasDeleteAfterDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiDatabaseUserSCRAMView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiDatabaseUserSCRAMView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiDatabaseUserSCRAMView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetLabels

`func (o *ApiDatabaseUserSCRAMView) GetLabels() []ApiAtlasNDSLabelView`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiDatabaseUserSCRAMView) GetLabelsOk() (*[]ApiAtlasNDSLabelView, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiDatabaseUserSCRAMView) SetLabels(v []ApiAtlasNDSLabelView)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiDatabaseUserSCRAMView) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLinks

`func (o *ApiDatabaseUserSCRAMView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiDatabaseUserSCRAMView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiDatabaseUserSCRAMView) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiDatabaseUserSCRAMView) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRoles

`func (o *ApiDatabaseUserSCRAMView) GetRoles() []ApiAtlasRoleView`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiDatabaseUserSCRAMView) GetRolesOk() (*[]ApiAtlasRoleView, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiDatabaseUserSCRAMView) SetRoles(v []ApiAtlasRoleView)`

SetRoles sets Roles field to given value.


### GetScopes

`func (o *ApiDatabaseUserSCRAMView) GetScopes() []ApiAtlasUserScopeView`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiDatabaseUserSCRAMView) GetScopesOk() (*[]ApiAtlasUserScopeView, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiDatabaseUserSCRAMView) SetScopes(v []ApiAtlasUserScopeView)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApiDatabaseUserSCRAMView) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPassword

`func (o *ApiDatabaseUserSCRAMView) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiDatabaseUserSCRAMView) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiDatabaseUserSCRAMView) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *ApiDatabaseUserSCRAMView) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiDatabaseUserSCRAMView) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiDatabaseUserSCRAMView) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


