# ApiAtlasAutoScalingView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**ApiAtlasComputeAutoScalingView**](ApiAtlasComputeAutoScalingView.md) |  | [optional] 
**DiskGBEnabled** | Pointer to **bool** | Flag that indicates whether someone enabled disk auto-scaling for this cluster. | [optional] 

## Methods

### NewApiAtlasAutoScalingView

`func NewApiAtlasAutoScalingView() *ApiAtlasAutoScalingView`

NewApiAtlasAutoScalingView instantiates a new ApiAtlasAutoScalingView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasAutoScalingViewWithDefaults

`func NewApiAtlasAutoScalingViewWithDefaults() *ApiAtlasAutoScalingView`

NewApiAtlasAutoScalingViewWithDefaults instantiates a new ApiAtlasAutoScalingView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *ApiAtlasAutoScalingView) GetCompute() ApiAtlasComputeAutoScalingView`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *ApiAtlasAutoScalingView) GetComputeOk() (*ApiAtlasComputeAutoScalingView, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *ApiAtlasAutoScalingView) SetCompute(v ApiAtlasComputeAutoScalingView)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *ApiAtlasAutoScalingView) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetDiskGBEnabled

`func (o *ApiAtlasAutoScalingView) GetDiskGBEnabled() bool`

GetDiskGBEnabled returns the DiskGBEnabled field if non-nil, zero value otherwise.

### GetDiskGBEnabledOk

`func (o *ApiAtlasAutoScalingView) GetDiskGBEnabledOk() (*bool, bool)`

GetDiskGBEnabledOk returns a tuple with the DiskGBEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGBEnabled

`func (o *ApiAtlasAutoScalingView) SetDiskGBEnabled(v bool)`

SetDiskGBEnabled sets DiskGBEnabled field to given value.

### HasDiskGBEnabled

`func (o *ApiAtlasAutoScalingView) HasDiskGBEnabled() bool`

HasDiskGBEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


