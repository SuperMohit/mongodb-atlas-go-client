# ApiAtlasDataLakeDataProcessRegionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Name of the cloud service that hosts the data lake&#39;s data stores. | 
**Region** | [**ApiAtlasDataLakeRegionView**](ApiAtlasDataLakeRegionView.md) |  | 

## Methods

### NewApiAtlasDataLakeDataProcessRegionView

`func NewApiAtlasDataLakeDataProcessRegionView(cloudProvider string, region ApiAtlasDataLakeRegionView, ) *ApiAtlasDataLakeDataProcessRegionView`

NewApiAtlasDataLakeDataProcessRegionView instantiates a new ApiAtlasDataLakeDataProcessRegionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDataLakeDataProcessRegionViewWithDefaults

`func NewApiAtlasDataLakeDataProcessRegionViewWithDefaults() *ApiAtlasDataLakeDataProcessRegionView`

NewApiAtlasDataLakeDataProcessRegionViewWithDefaults instantiates a new ApiAtlasDataLakeDataProcessRegionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ApiAtlasDataLakeDataProcessRegionView) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ApiAtlasDataLakeDataProcessRegionView) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ApiAtlasDataLakeDataProcessRegionView) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegion

`func (o *ApiAtlasDataLakeDataProcessRegionView) GetRegion() ApiAtlasDataLakeRegionView`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ApiAtlasDataLakeDataProcessRegionView) GetRegionOk() (*ApiAtlasDataLakeRegionView, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ApiAtlasDataLakeDataProcessRegionView) SetRegion(v ApiAtlasDataLakeRegionView)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


