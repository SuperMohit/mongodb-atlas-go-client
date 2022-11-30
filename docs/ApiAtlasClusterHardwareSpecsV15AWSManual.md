# ApiAtlasClusterHardwareSpecsV15AWSManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskIOPS** | Pointer to **int32** | Target throughput desired for storage attached to your AWS-provisioned cluster. Only change this parameter if you:  - set &#x60;\&quot;replicationSpecs[n].regionConfigs[m].providerName\&quot; : \&quot;AWS\&quot;&#x60;. - set &#x60;\&quot;replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize\&quot; : \&quot;M30\&quot;&#x60; or greater not including &#x60;Mxx_NVME&#x60; tiers.  The maximum input/output operations per second (IOPS) depend on the selected **.instanceSize** and **.diskSizeGB**. This parameter defaults to the cluster tier&#39;s standard IOPS value. Changing this value impacts cluster cost. MongoDB Cloud enforces minimum ratios of storage capacity to system memory for given cluster tiers. This keeps cluster performance consistent with large datasets.  - Instance sizes &#x60;M10&#x60; to &#x60;M40&#x60; have a ratio of disk capacity to system memory of 60:1. - Instance sizes greater than &#x60;M40&#x60; have a ratio of 120:1. | [optional] 
**EbsVolumeType** | Pointer to **string** | Type of storage you want to attach to your AWS-provisioned cluster.  - &#x60;STANDARD&#x60; volume types can&#39;t exceed the default input/output operations per second (IOPS) rate for the selected volume size.  - &#x60;PROVISIONED&#x60; volume types must fall within the allowable IOPS range for the selected volume size.\&quot;, | [optional] [default to "PROVISIONED"]
**InstanceSize** | Pointer to **string** | Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size. If you deploy a Global Cluster, you must choose a instance size of &#x60;M30&#x60; or greater. | [optional] 
**NodeCount** | Pointer to **int32** | Number of read-only nodes for MongoDB Cloud deploys to the region. Read-only nodes can never become the primary, but can enable local reads. | [optional] 

## Methods

### NewApiAtlasClusterHardwareSpecsV15AWSManual

`func NewApiAtlasClusterHardwareSpecsV15AWSManual() *ApiAtlasClusterHardwareSpecsV15AWSManual`

NewApiAtlasClusterHardwareSpecsV15AWSManual instantiates a new ApiAtlasClusterHardwareSpecsV15AWSManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasClusterHardwareSpecsV15AWSManualWithDefaults

`func NewApiAtlasClusterHardwareSpecsV15AWSManualWithDefaults() *ApiAtlasClusterHardwareSpecsV15AWSManual`

NewApiAtlasClusterHardwareSpecsV15AWSManualWithDefaults instantiates a new ApiAtlasClusterHardwareSpecsV15AWSManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskIOPS

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) GetDiskIOPS() int32`

GetDiskIOPS returns the DiskIOPS field if non-nil, zero value otherwise.

### GetDiskIOPSOk

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) GetDiskIOPSOk() (*int32, bool)`

GetDiskIOPSOk returns a tuple with the DiskIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIOPS

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) SetDiskIOPS(v int32)`

SetDiskIOPS sets DiskIOPS field to given value.

### HasDiskIOPS

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) HasDiskIOPS() bool`

HasDiskIOPS returns a boolean if a field has been set.

### GetEbsVolumeType

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) GetEbsVolumeType() string`

GetEbsVolumeType returns the EbsVolumeType field if non-nil, zero value otherwise.

### GetEbsVolumeTypeOk

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) GetEbsVolumeTypeOk() (*string, bool)`

GetEbsVolumeTypeOk returns a tuple with the EbsVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsVolumeType

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) SetEbsVolumeType(v string)`

SetEbsVolumeType sets EbsVolumeType field to given value.

### HasEbsVolumeType

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) HasEbsVolumeType() bool`

HasEbsVolumeType returns a boolean if a field has been set.

### GetInstanceSize

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.

### GetNodeCount

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ApiAtlasClusterHardwareSpecsV15AWSManual) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


