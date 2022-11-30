# ApiAtlasDatabaseUserViewBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | Pointer to **string** | MongoDB database against which the MongoDB database user authenticates. MongoDB database users must provide both a username and authentication database to log into MongoDB. | [optional] [default to "admin"]
**DeleteAfterDate** | Pointer to **string** | Date and time when MongoDB Cloud deletes the user. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project. | [optional] 
**Labels** | Pointer to [**[]ApiAtlasNDSLabelView**](ApiAtlasNDSLabelView.md) | List that contains the key-value pairs for tagging and categorizing the MongoDB database user. The labels that you define do not appear in the console. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Roles** | Pointer to [**[]ApiAtlasRoleView**](ApiAtlasRoleView.md) | List that provides the pairings of one role with one applicable database. | [optional] 
**Scopes** | Pointer to [**[]ApiAtlasUserScopeView**](ApiAtlasUserScopeView.md) | List that contains clusters and MongoDB Atlas Data Lakes that this database user can access. If omitted, MongoDB Cloud grants the database user access to all the clusters and MongoDB Atlas Data Lakes in the project. | [optional] 

## Methods

### NewApiAtlasDatabaseUserViewBase

`func NewApiAtlasDatabaseUserViewBase() *ApiAtlasDatabaseUserViewBase`

NewApiAtlasDatabaseUserViewBase instantiates a new ApiAtlasDatabaseUserViewBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDatabaseUserViewBaseWithDefaults

`func NewApiAtlasDatabaseUserViewBaseWithDefaults() *ApiAtlasDatabaseUserViewBase`

NewApiAtlasDatabaseUserViewBaseWithDefaults instantiates a new ApiAtlasDatabaseUserViewBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *ApiAtlasDatabaseUserViewBase) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ApiAtlasDatabaseUserViewBase) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ApiAtlasDatabaseUserViewBase) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *ApiAtlasDatabaseUserViewBase) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetDeleteAfterDate

`func (o *ApiAtlasDatabaseUserViewBase) GetDeleteAfterDate() string`

GetDeleteAfterDate returns the DeleteAfterDate field if non-nil, zero value otherwise.

### GetDeleteAfterDateOk

`func (o *ApiAtlasDatabaseUserViewBase) GetDeleteAfterDateOk() (*string, bool)`

GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfterDate

`func (o *ApiAtlasDatabaseUserViewBase) SetDeleteAfterDate(v string)`

SetDeleteAfterDate sets DeleteAfterDate field to given value.

### HasDeleteAfterDate

`func (o *ApiAtlasDatabaseUserViewBase) HasDeleteAfterDate() bool`

HasDeleteAfterDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiAtlasDatabaseUserViewBase) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasDatabaseUserViewBase) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasDatabaseUserViewBase) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAtlasDatabaseUserViewBase) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLabels

`func (o *ApiAtlasDatabaseUserViewBase) GetLabels() []ApiAtlasNDSLabelView`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiAtlasDatabaseUserViewBase) GetLabelsOk() (*[]ApiAtlasNDSLabelView, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiAtlasDatabaseUserViewBase) SetLabels(v []ApiAtlasNDSLabelView)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiAtlasDatabaseUserViewBase) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasDatabaseUserViewBase) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasDatabaseUserViewBase) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasDatabaseUserViewBase) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiAtlasDatabaseUserViewBase) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRoles

`func (o *ApiAtlasDatabaseUserViewBase) GetRoles() []ApiAtlasRoleView`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiAtlasDatabaseUserViewBase) GetRolesOk() (*[]ApiAtlasRoleView, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiAtlasDatabaseUserViewBase) SetRoles(v []ApiAtlasRoleView)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiAtlasDatabaseUserViewBase) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetScopes

`func (o *ApiAtlasDatabaseUserViewBase) GetScopes() []ApiAtlasUserScopeView`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiAtlasDatabaseUserViewBase) GetScopesOk() (*[]ApiAtlasUserScopeView, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiAtlasDatabaseUserViewBase) SetScopes(v []ApiAtlasUserScopeView)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApiAtlasDatabaseUserViewBase) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


