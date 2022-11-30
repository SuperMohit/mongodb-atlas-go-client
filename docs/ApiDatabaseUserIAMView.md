# ApiDatabaseUserIAMView

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
**AwsIAMType** | **string** | Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user&#39;s role. | [default to "NONE"]
**Username** | **string** | Human-readable label that represents the user that authenticates to MongoDB.  | 

## Methods

### NewApiDatabaseUserIAMView

`func NewApiDatabaseUserIAMView(databaseName string, groupId string, roles []ApiAtlasRoleView, awsIAMType string, username string, ) *ApiDatabaseUserIAMView`

NewApiDatabaseUserIAMView instantiates a new ApiDatabaseUserIAMView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDatabaseUserIAMViewWithDefaults

`func NewApiDatabaseUserIAMViewWithDefaults() *ApiDatabaseUserIAMView`

NewApiDatabaseUserIAMViewWithDefaults instantiates a new ApiDatabaseUserIAMView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *ApiDatabaseUserIAMView) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ApiDatabaseUserIAMView) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ApiDatabaseUserIAMView) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDeleteAfterDate

`func (o *ApiDatabaseUserIAMView) GetDeleteAfterDate() string`

GetDeleteAfterDate returns the DeleteAfterDate field if non-nil, zero value otherwise.

### GetDeleteAfterDateOk

`func (o *ApiDatabaseUserIAMView) GetDeleteAfterDateOk() (*string, bool)`

GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfterDate

`func (o *ApiDatabaseUserIAMView) SetDeleteAfterDate(v string)`

SetDeleteAfterDate sets DeleteAfterDate field to given value.

### HasDeleteAfterDate

`func (o *ApiDatabaseUserIAMView) HasDeleteAfterDate() bool`

HasDeleteAfterDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiDatabaseUserIAMView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiDatabaseUserIAMView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiDatabaseUserIAMView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetLabels

`func (o *ApiDatabaseUserIAMView) GetLabels() []ApiAtlasNDSLabelView`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiDatabaseUserIAMView) GetLabelsOk() (*[]ApiAtlasNDSLabelView, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiDatabaseUserIAMView) SetLabels(v []ApiAtlasNDSLabelView)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiDatabaseUserIAMView) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLinks

`func (o *ApiDatabaseUserIAMView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiDatabaseUserIAMView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiDatabaseUserIAMView) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiDatabaseUserIAMView) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRoles

`func (o *ApiDatabaseUserIAMView) GetRoles() []ApiAtlasRoleView`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiDatabaseUserIAMView) GetRolesOk() (*[]ApiAtlasRoleView, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiDatabaseUserIAMView) SetRoles(v []ApiAtlasRoleView)`

SetRoles sets Roles field to given value.


### GetScopes

`func (o *ApiDatabaseUserIAMView) GetScopes() []ApiAtlasUserScopeView`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiDatabaseUserIAMView) GetScopesOk() (*[]ApiAtlasUserScopeView, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiDatabaseUserIAMView) SetScopes(v []ApiAtlasUserScopeView)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ApiDatabaseUserIAMView) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetAwsIAMType

`func (o *ApiDatabaseUserIAMView) GetAwsIAMType() string`

GetAwsIAMType returns the AwsIAMType field if non-nil, zero value otherwise.

### GetAwsIAMTypeOk

`func (o *ApiDatabaseUserIAMView) GetAwsIAMTypeOk() (*string, bool)`

GetAwsIAMTypeOk returns a tuple with the AwsIAMType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIAMType

`func (o *ApiDatabaseUserIAMView) SetAwsIAMType(v string)`

SetAwsIAMType sets AwsIAMType field to given value.


### GetUsername

`func (o *ApiDatabaseUserIAMView) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiDatabaseUserIAMView) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiDatabaseUserIAMView) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


