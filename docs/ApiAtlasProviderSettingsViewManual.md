# ApiAtlasProviderSettingsViewManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Cloud service provider that applies to the provisioned the hosts. The &#x60;TENANT&#x60; value corresponds to an &#x60;M2&#x60; or &#x60;M5&#x60; multi-tenant cluster. For multi-tenant deployments, set **providerSettings.providerName** to &#x60;TENANT&#x60; and specify the cloud service provider in **providerSettings.backingProviderName**. | 

## Methods

### NewApiAtlasProviderSettingsViewManual

`func NewApiAtlasProviderSettingsViewManual(providerName string, ) *ApiAtlasProviderSettingsViewManual`

NewApiAtlasProviderSettingsViewManual instantiates a new ApiAtlasProviderSettingsViewManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasProviderSettingsViewManualWithDefaults

`func NewApiAtlasProviderSettingsViewManualWithDefaults() *ApiAtlasProviderSettingsViewManual`

NewApiAtlasProviderSettingsViewManualWithDefaults instantiates a new ApiAtlasProviderSettingsViewManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *ApiAtlasProviderSettingsViewManual) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ApiAtlasProviderSettingsViewManual) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ApiAtlasProviderSettingsViewManual) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


