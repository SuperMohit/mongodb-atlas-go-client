# ApiAtlasRoleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | **string** | Collection on which this role applies. | 
**DatabaseName** | **string** | Database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. | [default to "admin"]
**RoleName** | **string** | Human-readable label that identifies a group of privileges assigned to a database user. This value can either be a built-in role or a custom role. | 

## Methods

### NewApiAtlasRoleView

`func NewApiAtlasRoleView(collectionName string, databaseName string, roleName string, ) *ApiAtlasRoleView`

NewApiAtlasRoleView instantiates a new ApiAtlasRoleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasRoleViewWithDefaults

`func NewApiAtlasRoleViewWithDefaults() *ApiAtlasRoleView`

NewApiAtlasRoleViewWithDefaults instantiates a new ApiAtlasRoleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *ApiAtlasRoleView) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *ApiAtlasRoleView) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *ApiAtlasRoleView) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.


### GetDatabaseName

`func (o *ApiAtlasRoleView) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ApiAtlasRoleView) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ApiAtlasRoleView) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetRoleName

`func (o *ApiAtlasRoleView) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *ApiAtlasRoleView) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *ApiAtlasRoleView) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


