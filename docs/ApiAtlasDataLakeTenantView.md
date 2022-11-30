# ApiAtlasDataLakeTenantView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderConfig** | Pointer to [**ApiAtlasDataLakeCloudProviderConfigView**](ApiAtlasDataLakeCloudProviderConfigView.md) |  | [optional] 
**DataProcessRegion** | Pointer to [**ApiAtlasDataLakeDataProcessRegionView**](ApiAtlasDataLakeDataProcessRegionView.md) |  | [optional] 
**Name** | Pointer to **string** | Human-readable label that identifies the data lake. | [optional] [readonly] 
**Storage** | Pointer to [**ApiAtlasDataLakeStorageView**](ApiAtlasDataLakeStorageView.md) |  | [optional] 

## Methods

### NewApiAtlasDataLakeTenantView

`func NewApiAtlasDataLakeTenantView() *ApiAtlasDataLakeTenantView`

NewApiAtlasDataLakeTenantView instantiates a new ApiAtlasDataLakeTenantView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDataLakeTenantViewWithDefaults

`func NewApiAtlasDataLakeTenantViewWithDefaults() *ApiAtlasDataLakeTenantView`

NewApiAtlasDataLakeTenantViewWithDefaults instantiates a new ApiAtlasDataLakeTenantView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderConfig

`func (o *ApiAtlasDataLakeTenantView) GetCloudProviderConfig() ApiAtlasDataLakeCloudProviderConfigView`

GetCloudProviderConfig returns the CloudProviderConfig field if non-nil, zero value otherwise.

### GetCloudProviderConfigOk

`func (o *ApiAtlasDataLakeTenantView) GetCloudProviderConfigOk() (*ApiAtlasDataLakeCloudProviderConfigView, bool)`

GetCloudProviderConfigOk returns a tuple with the CloudProviderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderConfig

`func (o *ApiAtlasDataLakeTenantView) SetCloudProviderConfig(v ApiAtlasDataLakeCloudProviderConfigView)`

SetCloudProviderConfig sets CloudProviderConfig field to given value.

### HasCloudProviderConfig

`func (o *ApiAtlasDataLakeTenantView) HasCloudProviderConfig() bool`

HasCloudProviderConfig returns a boolean if a field has been set.

### GetDataProcessRegion

`func (o *ApiAtlasDataLakeTenantView) GetDataProcessRegion() ApiAtlasDataLakeDataProcessRegionView`

GetDataProcessRegion returns the DataProcessRegion field if non-nil, zero value otherwise.

### GetDataProcessRegionOk

`func (o *ApiAtlasDataLakeTenantView) GetDataProcessRegionOk() (*ApiAtlasDataLakeDataProcessRegionView, bool)`

GetDataProcessRegionOk returns a tuple with the DataProcessRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProcessRegion

`func (o *ApiAtlasDataLakeTenantView) SetDataProcessRegion(v ApiAtlasDataLakeDataProcessRegionView)`

SetDataProcessRegion sets DataProcessRegion field to given value.

### HasDataProcessRegion

`func (o *ApiAtlasDataLakeTenantView) HasDataProcessRegion() bool`

HasDataProcessRegion returns a boolean if a field has been set.

### GetName

`func (o *ApiAtlasDataLakeTenantView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasDataLakeTenantView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasDataLakeTenantView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAtlasDataLakeTenantView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorage

`func (o *ApiAtlasDataLakeTenantView) GetStorage() ApiAtlasDataLakeStorageView`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ApiAtlasDataLakeTenantView) GetStorageOk() (*ApiAtlasDataLakeStorageView, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ApiAtlasDataLakeTenantView) SetStorage(v ApiAtlasDataLakeStorageView)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ApiAtlasDataLakeTenantView) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


