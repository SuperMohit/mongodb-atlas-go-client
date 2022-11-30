# ApiOrganizationInvitationRequestView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to **[]string** | One or more organization or project level roles to assign to the MongoDB Cloud user. | [optional] 
**TeamIds** | Pointer to **[]string** | List of teams to which you want to invite the desired MongoDB Cloud user. | [optional] 
**Username** | Pointer to **string** | Email address that belongs to the desired MongoDB Cloud user. | [optional] 

## Methods

### NewApiOrganizationInvitationRequestView

`func NewApiOrganizationInvitationRequestView() *ApiOrganizationInvitationRequestView`

NewApiOrganizationInvitationRequestView instantiates a new ApiOrganizationInvitationRequestView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOrganizationInvitationRequestViewWithDefaults

`func NewApiOrganizationInvitationRequestViewWithDefaults() *ApiOrganizationInvitationRequestView`

NewApiOrganizationInvitationRequestViewWithDefaults instantiates a new ApiOrganizationInvitationRequestView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *ApiOrganizationInvitationRequestView) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiOrganizationInvitationRequestView) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiOrganizationInvitationRequestView) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiOrganizationInvitationRequestView) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTeamIds

`func (o *ApiOrganizationInvitationRequestView) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *ApiOrganizationInvitationRequestView) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *ApiOrganizationInvitationRequestView) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *ApiOrganizationInvitationRequestView) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetUsername

`func (o *ApiOrganizationInvitationRequestView) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiOrganizationInvitationRequestView) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiOrganizationInvitationRequestView) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiOrganizationInvitationRequestView) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


