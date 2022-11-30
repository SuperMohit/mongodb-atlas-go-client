# ApiTeamView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-hexadecimal digit string that identifies this team. | [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Name** | **string** | Human-readable label that identifies the team. | 
**Usernames** | **[]string** | List that contains the MongoDB Cloud users in this team. | 

## Methods

### NewApiTeamView

`func NewApiTeamView(id string, links []Link, name string, usernames []string, ) *ApiTeamView`

NewApiTeamView instantiates a new ApiTeamView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTeamViewWithDefaults

`func NewApiTeamViewWithDefaults() *ApiTeamView`

NewApiTeamViewWithDefaults instantiates a new ApiTeamView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiTeamView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiTeamView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiTeamView) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *ApiTeamView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiTeamView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiTeamView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetName

`func (o *ApiTeamView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiTeamView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiTeamView) SetName(v string)`

SetName sets Name field to given value.


### GetUsernames

`func (o *ApiTeamView) GetUsernames() []string`

GetUsernames returns the Usernames field if non-nil, zero value otherwise.

### GetUsernamesOk

`func (o *ApiTeamView) GetUsernamesOk() (*[]string, bool)`

GetUsernamesOk returns a tuple with the Usernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernames

`func (o *ApiTeamView) SetUsernames(v []string)`

SetUsernames sets Usernames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


