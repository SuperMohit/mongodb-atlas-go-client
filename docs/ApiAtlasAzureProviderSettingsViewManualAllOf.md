# ApiAtlasAzureProviderSettingsViewManualAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**ApiAtlasAzureProviderSettingsViewManualAllOfAutoScaling**](ApiAtlasAzureProviderSettingsViewManualAllOfAutoScaling.md) |  | [optional] 
**DiskTypeName** | Pointer to **string** | Disk type that corresponds to the host&#39;s root volume for Azure instances. If omitted, the default disk type for the selected **providerSettings.instanceSizeName** applies. | [optional] 
**InstanceSizeName** | Pointer to **string** | Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. | [optional] 
**RegionName** | Pointer to **string** | Human-readable label that identifies the geographic location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. For multi-region clusters, see **replicationSpec.{region}**. | [optional] 

## Methods

### NewApiAtlasAzureProviderSettingsViewManualAllOf

`func NewApiAtlasAzureProviderSettingsViewManualAllOf() *ApiAtlasAzureProviderSettingsViewManualAllOf`

NewApiAtlasAzureProviderSettingsViewManualAllOf instantiates a new ApiAtlasAzureProviderSettingsViewManualAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasAzureProviderSettingsViewManualAllOfWithDefaults

`func NewApiAtlasAzureProviderSettingsViewManualAllOfWithDefaults() *ApiAtlasAzureProviderSettingsViewManualAllOf`

NewApiAtlasAzureProviderSettingsViewManualAllOfWithDefaults instantiates a new ApiAtlasAzureProviderSettingsViewManualAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) GetAutoScaling() ApiAtlasAzureProviderSettingsViewManualAllOfAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) GetAutoScalingOk() (*ApiAtlasAzureProviderSettingsViewManualAllOfAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) SetAutoScaling(v ApiAtlasAzureProviderSettingsViewManualAllOfAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetDiskTypeName

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) GetDiskTypeName() string`

GetDiskTypeName returns the DiskTypeName field if non-nil, zero value otherwise.

### GetDiskTypeNameOk

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) GetDiskTypeNameOk() (*string, bool)`

GetDiskTypeNameOk returns a tuple with the DiskTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTypeName

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) SetDiskTypeName(v string)`

SetDiskTypeName sets DiskTypeName field to given value.

### HasDiskTypeName

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) HasDiskTypeName() bool`

HasDiskTypeName returns a boolean if a field has been set.

### GetInstanceSizeName

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) GetInstanceSizeName() string`

GetInstanceSizeName returns the InstanceSizeName field if non-nil, zero value otherwise.

### GetInstanceSizeNameOk

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) GetInstanceSizeNameOk() (*string, bool)`

GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizeName

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) SetInstanceSizeName(v string)`

SetInstanceSizeName sets InstanceSizeName field to given value.

### HasInstanceSizeName

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) HasInstanceSizeName() bool`

HasInstanceSizeName returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasAzureProviderSettingsViewManualAllOf) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


