# ApiAtlasGCPProviderSettingsViewManualAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling**](ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling.md) |  | [optional] 
**InstanceSizeName** | Pointer to **string** | Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. | [optional] 
**RegionName** | Pointer to **string** | Human-readable label that identifies the geographic location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. For multi-region clusters, see **replicationSpec.{region}**. | [optional] 

## Methods

### NewApiAtlasGCPProviderSettingsViewManualAllOf

`func NewApiAtlasGCPProviderSettingsViewManualAllOf() *ApiAtlasGCPProviderSettingsViewManualAllOf`

NewApiAtlasGCPProviderSettingsViewManualAllOf instantiates a new ApiAtlasGCPProviderSettingsViewManualAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasGCPProviderSettingsViewManualAllOfWithDefaults

`func NewApiAtlasGCPProviderSettingsViewManualAllOfWithDefaults() *ApiAtlasGCPProviderSettingsViewManualAllOf`

NewApiAtlasGCPProviderSettingsViewManualAllOfWithDefaults instantiates a new ApiAtlasGCPProviderSettingsViewManualAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) GetAutoScaling() ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) GetAutoScalingOk() (*ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) SetAutoScaling(v ApiAtlasGCPProviderSettingsViewManualAllOfAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetInstanceSizeName

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) GetInstanceSizeName() string`

GetInstanceSizeName returns the InstanceSizeName field if non-nil, zero value otherwise.

### GetInstanceSizeNameOk

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) GetInstanceSizeNameOk() (*string, bool)`

GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizeName

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) SetInstanceSizeName(v string)`

SetInstanceSizeName sets InstanceSizeName field to given value.

### HasInstanceSizeName

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) HasInstanceSizeName() bool`

HasInstanceSizeName returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasGCPProviderSettingsViewManualAllOf) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


