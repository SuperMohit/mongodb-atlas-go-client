# ApiTeamRoleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**RoleNames** | Pointer to **[]string** | One or more organization- or project-level roles to assign to the MongoDB Cloud user. | [optional] 
**TeamId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the team. | [optional] 

## Methods

### NewApiTeamRoleView

`func NewApiTeamRoleView(links []Link, ) *ApiTeamRoleView`

NewApiTeamRoleView instantiates a new ApiTeamRoleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTeamRoleViewWithDefaults

`func NewApiTeamRoleViewWithDefaults() *ApiTeamRoleView`

NewApiTeamRoleViewWithDefaults instantiates a new ApiTeamRoleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApiTeamRoleView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiTeamRoleView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiTeamRoleView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetRoleNames

`func (o *ApiTeamRoleView) GetRoleNames() []string`

GetRoleNames returns the RoleNames field if non-nil, zero value otherwise.

### GetRoleNamesOk

`func (o *ApiTeamRoleView) GetRoleNamesOk() (*[]string, bool)`

GetRoleNamesOk returns a tuple with the RoleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNames

`func (o *ApiTeamRoleView) SetRoleNames(v []string)`

SetRoleNames sets RoleNames field to given value.

### HasRoleNames

`func (o *ApiTeamRoleView) HasRoleNames() bool`

HasRoleNames returns a boolean if a field has been set.

### GetTeamId

`func (o *ApiTeamRoleView) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ApiTeamRoleView) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ApiTeamRoleView) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *ApiTeamRoleView) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


