# ApiWebhookViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | Parameter returned if someone configure this webhook with a secret. | [optional] 
**Url** | Pointer to **string** | Endpoint web address to which MongoDB Cloud sends notifications. | [optional] 

## Methods

### NewApiWebhookViewAllOf

`func NewApiWebhookViewAllOf() *ApiWebhookViewAllOf`

NewApiWebhookViewAllOf instantiates a new ApiWebhookViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWebhookViewAllOfWithDefaults

`func NewApiWebhookViewAllOfWithDefaults() *ApiWebhookViewAllOf`

NewApiWebhookViewAllOfWithDefaults instantiates a new ApiWebhookViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *ApiWebhookViewAllOf) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ApiWebhookViewAllOf) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ApiWebhookViewAllOf) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ApiWebhookViewAllOf) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUrl

`func (o *ApiWebhookViewAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiWebhookViewAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiWebhookViewAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApiWebhookViewAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


