# ApiAtlasCustomDBRoleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]ApiAtlasDBActionView**](ApiAtlasDBActionView.md) | List of the individual privilege actions that the role grants. | [optional] 
**InheritedRoles** | Pointer to [**[]ApiAtlasInheritedRoleView**](ApiAtlasInheritedRoleView.md) | List of the built-in roles that this custom role inherits. | [optional] 
**RoleName** | **string** | Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project. | 

## Methods

### NewApiAtlasCustomDBRoleView

`func NewApiAtlasCustomDBRoleView(roleName string, ) *ApiAtlasCustomDBRoleView`

NewApiAtlasCustomDBRoleView instantiates a new ApiAtlasCustomDBRoleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCustomDBRoleViewWithDefaults

`func NewApiAtlasCustomDBRoleViewWithDefaults() *ApiAtlasCustomDBRoleView`

NewApiAtlasCustomDBRoleViewWithDefaults instantiates a new ApiAtlasCustomDBRoleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ApiAtlasCustomDBRoleView) GetActions() []ApiAtlasDBActionView`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiAtlasCustomDBRoleView) GetActionsOk() (*[]ApiAtlasDBActionView, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiAtlasCustomDBRoleView) SetActions(v []ApiAtlasDBActionView)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ApiAtlasCustomDBRoleView) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetInheritedRoles

`func (o *ApiAtlasCustomDBRoleView) GetInheritedRoles() []ApiAtlasInheritedRoleView`

GetInheritedRoles returns the InheritedRoles field if non-nil, zero value otherwise.

### GetInheritedRolesOk

`func (o *ApiAtlasCustomDBRoleView) GetInheritedRolesOk() (*[]ApiAtlasInheritedRoleView, bool)`

GetInheritedRolesOk returns a tuple with the InheritedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedRoles

`func (o *ApiAtlasCustomDBRoleView) SetInheritedRoles(v []ApiAtlasInheritedRoleView)`

SetInheritedRoles sets InheritedRoles field to given value.

### HasInheritedRoles

`func (o *ApiAtlasCustomDBRoleView) HasInheritedRoles() bool`

HasInheritedRoles returns a boolean if a field has been set.

### GetRoleName

`func (o *ApiAtlasCustomDBRoleView) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *ApiAtlasCustomDBRoleView) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *ApiAtlasCustomDBRoleView) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


