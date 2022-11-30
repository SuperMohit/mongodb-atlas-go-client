# ApiDatadogView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | Key that allows MongoDB Cloud to access your Datadog account. | 
**Region** | Pointer to **string** | Two-letter code that indicates which regional URL MongoDB uses to access the Datadog API. | [optional] [default to "US"]

## Methods

### NewApiDatadogView

`func NewApiDatadogView(apiKey string, ) *ApiDatadogView`

NewApiDatadogView instantiates a new ApiDatadogView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDatadogViewWithDefaults

`func NewApiDatadogViewWithDefaults() *ApiDatadogView`

NewApiDatadogViewWithDefaults instantiates a new ApiDatadogView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *ApiDatadogView) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiDatadogView) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiDatadogView) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetRegion

`func (o *ApiDatadogView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ApiDatadogView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ApiDatadogView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ApiDatadogView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


