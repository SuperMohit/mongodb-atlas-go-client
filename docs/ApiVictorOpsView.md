# ApiVictorOpsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | Key that allows MongoDB Cloud to access your Splunk On-Call account. | 
**RoutingKey** | Pointer to **string** | Routing key associated with your Splunk On-Call account. | [optional] 

## Methods

### NewApiVictorOpsView

`func NewApiVictorOpsView(apiKey string, ) *ApiVictorOpsView`

NewApiVictorOpsView instantiates a new ApiVictorOpsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiVictorOpsViewWithDefaults

`func NewApiVictorOpsViewWithDefaults() *ApiVictorOpsView`

NewApiVictorOpsViewWithDefaults instantiates a new ApiVictorOpsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *ApiVictorOpsView) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiVictorOpsView) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiVictorOpsView) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetRoutingKey

`func (o *ApiVictorOpsView) GetRoutingKey() string`

GetRoutingKey returns the RoutingKey field if non-nil, zero value otherwise.

### GetRoutingKeyOk

`func (o *ApiVictorOpsView) GetRoutingKeyOk() (*string, bool)`

GetRoutingKeyOk returns a tuple with the RoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKey

`func (o *ApiVictorOpsView) SetRoutingKey(v string)`

SetRoutingKey sets RoutingKey field to given value.

### HasRoutingKey

`func (o *ApiVictorOpsView) HasRoutingKey() bool`

HasRoutingKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


