# ApiAtlasDataLakeStoreView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human-readable label that identifies the data store. The **databases.[n].collections.[n].dataSources.[n].storeName** field references this values as part of the mapping configuration. To use MongoDB Cloud as a data store, the data lake requires a serverless instance or an &#x60;M10&#x60; or higher cluster. | [optional] 
**Provider** | **string** |  | 

## Methods

### NewApiAtlasDataLakeStoreView

`func NewApiAtlasDataLakeStoreView(provider string, ) *ApiAtlasDataLakeStoreView`

NewApiAtlasDataLakeStoreView instantiates a new ApiAtlasDataLakeStoreView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDataLakeStoreViewWithDefaults

`func NewApiAtlasDataLakeStoreViewWithDefaults() *ApiAtlasDataLakeStoreView`

NewApiAtlasDataLakeStoreViewWithDefaults instantiates a new ApiAtlasDataLakeStoreView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiAtlasDataLakeStoreView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasDataLakeStoreView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasDataLakeStoreView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAtlasDataLakeStoreView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *ApiAtlasDataLakeStoreView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApiAtlasDataLakeStoreView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApiAtlasDataLakeStoreView) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


