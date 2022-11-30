# ApiSlackViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | Pointer to **string** | Key that allows MongoDB Cloud to access your Slack account. | [optional] 
**ChannelName** | Pointer to **NullableString** | Name of the Slack channel to which MongoDB Cloud sends alert notifications. | [optional] 
**TeamName** | Pointer to **string** | Human-readable label that identifies your Slack team. Set this parameter when you configure a legacy Slack integration. | [optional] 

## Methods

### NewApiSlackViewAllOf

`func NewApiSlackViewAllOf() *ApiSlackViewAllOf`

NewApiSlackViewAllOf instantiates a new ApiSlackViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSlackViewAllOfWithDefaults

`func NewApiSlackViewAllOfWithDefaults() *ApiSlackViewAllOf`

NewApiSlackViewAllOfWithDefaults instantiates a new ApiSlackViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiToken

`func (o *ApiSlackViewAllOf) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *ApiSlackViewAllOf) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *ApiSlackViewAllOf) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *ApiSlackViewAllOf) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetChannelName

`func (o *ApiSlackViewAllOf) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ApiSlackViewAllOf) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ApiSlackViewAllOf) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ApiSlackViewAllOf) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *ApiSlackViewAllOf) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *ApiSlackViewAllOf) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetTeamName

`func (o *ApiSlackViewAllOf) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ApiSlackViewAllOf) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ApiSlackViewAllOf) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *ApiSlackViewAllOf) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


