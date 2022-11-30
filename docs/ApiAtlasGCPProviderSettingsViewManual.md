# ApiAtlasGCPProviderSettingsViewManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling**](ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling.md) |  | [optional] 
**InstanceSizeName** | Pointer to **string** | Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. | [optional] 
**RegionName** | Pointer to **string** | Human-readable label that identifies the geographic location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. For multi-region clusters, see **replicationSpec.{region}**. | [optional] 

## Methods

### NewApiAtlasGCPProviderSettingsViewManual

`func NewApiAtlasGCPProviderSettingsViewManual() *ApiAtlasGCPProviderSettingsViewManual`

NewApiAtlasGCPProviderSettingsViewManual instantiates a new ApiAtlasGCPProviderSettingsViewManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasGCPProviderSettingsViewManualWithDefaults

`func NewApiAtlasGCPProviderSettingsViewManualWithDefaults() *ApiAtlasGCPProviderSettingsViewManual`

NewApiAtlasGCPProviderSettingsViewManualWithDefaults instantiates a new ApiAtlasGCPProviderSettingsViewManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *ApiAtlasGCPProviderSettingsViewManual) GetAutoScaling() ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *ApiAtlasGCPProviderSettingsViewManual) GetAutoScalingOk() (*ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *ApiAtlasGCPProviderSettingsViewManual) SetAutoScaling(v ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *ApiAtlasGCPProviderSettingsViewManual) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetInstanceSizeName

`func (o *ApiAtlasGCPProviderSettingsViewManual) GetInstanceSizeName() string`

GetInstanceSizeName returns the InstanceSizeName field if non-nil, zero value otherwise.

### GetInstanceSizeNameOk

`func (o *ApiAtlasGCPProviderSettingsViewManual) GetInstanceSizeNameOk() (*string, bool)`

GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizeName

`func (o *ApiAtlasGCPProviderSettingsViewManual) SetInstanceSizeName(v string)`

SetInstanceSizeName sets InstanceSizeName field to given value.

### HasInstanceSizeName

`func (o *ApiAtlasGCPProviderSettingsViewManual) HasInstanceSizeName() bool`

HasInstanceSizeName returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasGCPProviderSettingsViewManual) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasGCPProviderSettingsViewManual) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasGCPProviderSettingsViewManual) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasGCPProviderSettingsViewManual) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


