# ApiSlackView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | **string** | Key that allows MongoDB Cloud to access your Slack account. | 
**ChannelName** | **NullableString** | Name of the Slack channel to which MongoDB Cloud sends alert notifications. | 
**TeamName** | Pointer to **string** | Human-readable label that identifies your Slack team. Set this parameter when you configure a legacy Slack integration. | [optional] 

## Methods

### NewApiSlackView

`func NewApiSlackView(apiToken string, channelName NullableString, ) *ApiSlackView`

NewApiSlackView instantiates a new ApiSlackView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSlackViewWithDefaults

`func NewApiSlackViewWithDefaults() *ApiSlackView`

NewApiSlackViewWithDefaults instantiates a new ApiSlackView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiToken

`func (o *ApiSlackView) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *ApiSlackView) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *ApiSlackView) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetChannelName

`func (o *ApiSlackView) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ApiSlackView) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ApiSlackView) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.


### SetChannelNameNil

`func (o *ApiSlackView) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *ApiSlackView) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetTeamName

`func (o *ApiSlackView) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ApiSlackView) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ApiSlackView) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *ApiSlackView) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


