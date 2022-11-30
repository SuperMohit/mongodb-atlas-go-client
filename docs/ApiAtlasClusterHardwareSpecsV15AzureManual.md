# ApiAtlasClusterHardwareSpecsV15AzureManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSize** | Pointer to **string** | Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size. If you deploy a Global Cluster, you must choose a instance size of &#x60;M30&#x60; or greater. | [optional] 
**NodeCount** | Pointer to **int32** | Number of read-only nodes for MongoDB Cloud deploys to the region. Read-only nodes can never become the primary, but can enable local reads. | [optional] 

## Methods

### NewApiAtlasClusterHardwareSpecsV15AzureManual

`func NewApiAtlasClusterHardwareSpecsV15AzureManual() *ApiAtlasClusterHardwareSpecsV15AzureManual`

NewApiAtlasClusterHardwareSpecsV15AzureManual instantiates a new ApiAtlasClusterHardwareSpecsV15AzureManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasClusterHardwareSpecsV15AzureManualWithDefaults

`func NewApiAtlasClusterHardwareSpecsV15AzureManualWithDefaults() *ApiAtlasClusterHardwareSpecsV15AzureManual`

NewApiAtlasClusterHardwareSpecsV15AzureManualWithDefaults instantiates a new ApiAtlasClusterHardwareSpecsV15AzureManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSize

`func (o *ApiAtlasClusterHardwareSpecsV15AzureManual) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *ApiAtlasClusterHardwareSpecsV15AzureManual) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *ApiAtlasClusterHardwareSpecsV15AzureManual) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *ApiAtlasClusterHardwareSpecsV15AzureManual) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.

### GetNodeCount

`func (o *ApiAtlasClusterHardwareSpecsV15AzureManual) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ApiAtlasClusterHardwareSpecsV15AzureManual) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ApiAtlasClusterHardwareSpecsV15AzureManual) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ApiAtlasClusterHardwareSpecsV15AzureManual) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


