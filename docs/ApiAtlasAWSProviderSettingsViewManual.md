# ApiAtlasAWSProviderSettingsViewManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**ApiAtlasAWSProviderSettingsViewManualAllOfAutoScaling**](ApiAtlasAWSProviderSettingsViewManualAllOfAutoScaling.md) |  | [optional] 
**DiskIOPS** | Pointer to **int32** | Maximum Disk Input/Output Operations per Second (IOPS) that the database host can perform. | [optional] 
**EncryptEBSVolume** | Pointer to **bool** | Flag that indicates whether the Amazon Elastic Block Store (EBS) encryption feature encrypts the host&#39;s root volume for both data at rest within the volume and for data moving between the volume and the cluster. Clusters always have this setting enabled. | [optional] [default to true]
**InstanceSizeName** | Pointer to **string** | Cluster tier, with a default storage and memory capacity, that applies to all the data-bearing hosts in your cluster. | [optional] 
**RegionName** | Pointer to **string** | Human-readable label that identifies the geographic location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. For multi-region clusters, see **replicationSpec.region**. | [optional] 
**VolumeType** | Pointer to **string** | Disk Input/Output Operations per Second (IOPS) setting for Amazon Web Services (AWS) storage that you configure only for abbr title&#x3D;\&quot;Amazon Web Services\&quot;&gt;AWS&lt;/abbr&gt;. Specify whether Disk Input/Output Operations per Second (IOPS) must not exceed the default Input/Output Operations per Second (IOPS) rate for the selected volume size (&#x60;STANDARD&#x60;), or must fall within the allowable Input/Output Operations per Second (IOPS) range for the selected volume size (&#x60;PROVISIONED&#x60;). | [optional] 

## Methods

### NewApiAtlasAWSProviderSettingsViewManual

`func NewApiAtlasAWSProviderSettingsViewManual() *ApiAtlasAWSProviderSettingsViewManual`

NewApiAtlasAWSProviderSettingsViewManual instantiates a new ApiAtlasAWSProviderSettingsViewManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasAWSProviderSettingsViewManualWithDefaults

`func NewApiAtlasAWSProviderSettingsViewManualWithDefaults() *ApiAtlasAWSProviderSettingsViewManual`

NewApiAtlasAWSProviderSettingsViewManualWithDefaults instantiates a new ApiAtlasAWSProviderSettingsViewManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetAutoScaling() ApiAtlasAWSProviderSettingsViewManualAllOfAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetAutoScalingOk() (*ApiAtlasAWSProviderSettingsViewManualAllOfAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *ApiAtlasAWSProviderSettingsViewManual) SetAutoScaling(v ApiAtlasAWSProviderSettingsViewManualAllOfAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *ApiAtlasAWSProviderSettingsViewManual) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetDiskIOPS

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetDiskIOPS() int32`

GetDiskIOPS returns the DiskIOPS field if non-nil, zero value otherwise.

### GetDiskIOPSOk

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetDiskIOPSOk() (*int32, bool)`

GetDiskIOPSOk returns a tuple with the DiskIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIOPS

`func (o *ApiAtlasAWSProviderSettingsViewManual) SetDiskIOPS(v int32)`

SetDiskIOPS sets DiskIOPS field to given value.

### HasDiskIOPS

`func (o *ApiAtlasAWSProviderSettingsViewManual) HasDiskIOPS() bool`

HasDiskIOPS returns a boolean if a field has been set.

### GetEncryptEBSVolume

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetEncryptEBSVolume() bool`

GetEncryptEBSVolume returns the EncryptEBSVolume field if non-nil, zero value otherwise.

### GetEncryptEBSVolumeOk

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetEncryptEBSVolumeOk() (*bool, bool)`

GetEncryptEBSVolumeOk returns a tuple with the EncryptEBSVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptEBSVolume

`func (o *ApiAtlasAWSProviderSettingsViewManual) SetEncryptEBSVolume(v bool)`

SetEncryptEBSVolume sets EncryptEBSVolume field to given value.

### HasEncryptEBSVolume

`func (o *ApiAtlasAWSProviderSettingsViewManual) HasEncryptEBSVolume() bool`

HasEncryptEBSVolume returns a boolean if a field has been set.

### GetInstanceSizeName

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetInstanceSizeName() string`

GetInstanceSizeName returns the InstanceSizeName field if non-nil, zero value otherwise.

### GetInstanceSizeNameOk

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetInstanceSizeNameOk() (*string, bool)`

GetInstanceSizeNameOk returns a tuple with the InstanceSizeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizeName

`func (o *ApiAtlasAWSProviderSettingsViewManual) SetInstanceSizeName(v string)`

SetInstanceSizeName sets InstanceSizeName field to given value.

### HasInstanceSizeName

`func (o *ApiAtlasAWSProviderSettingsViewManual) HasInstanceSizeName() bool`

HasInstanceSizeName returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasAWSProviderSettingsViewManual) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasAWSProviderSettingsViewManual) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetVolumeType

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ApiAtlasAWSProviderSettingsViewManual) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ApiAtlasAWSProviderSettingsViewManual) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ApiAtlasAWSProviderSettingsViewManual) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


