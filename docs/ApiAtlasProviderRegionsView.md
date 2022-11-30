# ApiAtlasProviderRegionsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSizes** | Pointer to [**[]ApiInstanceSizeView**](ApiInstanceSizeView.md) | List of instances sizes that this cloud provider supports. | [optional] [readonly] 
**Provider** | Pointer to **string** | Human-readable label that identifies the Cloud provider. | [optional] 

## Methods

### NewApiAtlasProviderRegionsView

`func NewApiAtlasProviderRegionsView() *ApiAtlasProviderRegionsView`

NewApiAtlasProviderRegionsView instantiates a new ApiAtlasProviderRegionsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasProviderRegionsViewWithDefaults

`func NewApiAtlasProviderRegionsViewWithDefaults() *ApiAtlasProviderRegionsView`

NewApiAtlasProviderRegionsViewWithDefaults instantiates a new ApiAtlasProviderRegionsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSizes

`func (o *ApiAtlasProviderRegionsView) GetInstanceSizes() []ApiInstanceSizeView`

GetInstanceSizes returns the InstanceSizes field if non-nil, zero value otherwise.

### GetInstanceSizesOk

`func (o *ApiAtlasProviderRegionsView) GetInstanceSizesOk() (*[]ApiInstanceSizeView, bool)`

GetInstanceSizesOk returns a tuple with the InstanceSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizes

`func (o *ApiAtlasProviderRegionsView) SetInstanceSizes(v []ApiInstanceSizeView)`

SetInstanceSizes sets InstanceSizes field to given value.

### HasInstanceSizes

`func (o *ApiAtlasProviderRegionsView) HasInstanceSizes() bool`

HasInstanceSizes returns a boolean if a field has been set.

### GetProvider

`func (o *ApiAtlasProviderRegionsView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApiAtlasProviderRegionsView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApiAtlasProviderRegionsView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ApiAtlasProviderRegionsView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


