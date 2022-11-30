# ServerlessInstanceProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackingProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisioned the serverless instance. | [optional] 
**ProviderName** | **string** | Human-readable label that identifies the cloud service provider. | [default to "SERVERLESS"]
**RegionName** | Pointer to **string** | Human-readable label that identifies the geographic location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/). | [optional] 

## Methods

### NewServerlessInstanceProviderSettings

`func NewServerlessInstanceProviderSettings(providerName string, ) *ServerlessInstanceProviderSettings`

NewServerlessInstanceProviderSettings instantiates a new ServerlessInstanceProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessInstanceProviderSettingsWithDefaults

`func NewServerlessInstanceProviderSettingsWithDefaults() *ServerlessInstanceProviderSettings`

NewServerlessInstanceProviderSettingsWithDefaults instantiates a new ServerlessInstanceProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackingProviderName

`func (o *ServerlessInstanceProviderSettings) GetBackingProviderName() string`

GetBackingProviderName returns the BackingProviderName field if non-nil, zero value otherwise.

### GetBackingProviderNameOk

`func (o *ServerlessInstanceProviderSettings) GetBackingProviderNameOk() (*string, bool)`

GetBackingProviderNameOk returns a tuple with the BackingProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingProviderName

`func (o *ServerlessInstanceProviderSettings) SetBackingProviderName(v string)`

SetBackingProviderName sets BackingProviderName field to given value.

### HasBackingProviderName

`func (o *ServerlessInstanceProviderSettings) HasBackingProviderName() bool`

HasBackingProviderName returns a boolean if a field has been set.

### GetProviderName

`func (o *ServerlessInstanceProviderSettings) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ServerlessInstanceProviderSettings) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ServerlessInstanceProviderSettings) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetRegionName

`func (o *ServerlessInstanceProviderSettings) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ServerlessInstanceProviderSettings) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ServerlessInstanceProviderSettings) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ServerlessInstanceProviderSettings) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


