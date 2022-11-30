# ApiDatadogViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | Key that allows MongoDB Cloud to access your Datadog account. | [optional] 
**Region** | Pointer to **string** | Two-letter code that indicates which regional URL MongoDB uses to access the Datadog API. | [optional] [default to "US"]

## Methods

### NewApiDatadogViewAllOf

`func NewApiDatadogViewAllOf() *ApiDatadogViewAllOf`

NewApiDatadogViewAllOf instantiates a new ApiDatadogViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDatadogViewAllOfWithDefaults

`func NewApiDatadogViewAllOfWithDefaults() *ApiDatadogViewAllOf`

NewApiDatadogViewAllOfWithDefaults instantiates a new ApiDatadogViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *ApiDatadogViewAllOf) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiDatadogViewAllOf) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiDatadogViewAllOf) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ApiDatadogViewAllOf) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetRegion

`func (o *ApiDatadogViewAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ApiDatadogViewAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ApiDatadogViewAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ApiDatadogViewAllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


