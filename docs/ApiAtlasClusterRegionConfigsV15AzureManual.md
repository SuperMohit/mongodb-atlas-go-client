# ApiAtlasClusterRegionConfigsV15AzureManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsAutoScaling** | Pointer to [**ApiAtlasClusterV15AutoScalingAzureManual**](ApiAtlasClusterV15AutoScalingAzureManual.md) |  | [optional] 
**AnalyticsSpecs** | Pointer to [**ApiAtlasClusterHardwareSpecsV15AzureManual**](ApiAtlasClusterHardwareSpecsV15AzureManual.md) |  | [optional] 
**AutoScaling** | Pointer to [**ApiAtlasClusterV15AutoScalingAzureManual**](ApiAtlasClusterV15AutoScalingAzureManual.md) |  | [optional] 
**ElectableSpecs** | Pointer to [**ApiAtlasClusterHardwareSpecsV15AzureManual**](ApiAtlasClusterHardwareSpecsV15AzureManual.md) |  | [optional] 
**Priority** | Pointer to **int32** | Precedence that MongoDB Cloud gives this region when a primary election occurs. If your **regionConfigs** has only **readOnlySpecs**, **analyticsSpecs**, or both, set this value to &#x60;0&#x60;. If you have multiple **regionConfigs** objects, set priorities in descending order starting from &#x60;7&#x60;.  Example: If you have three regions, their priorities would be &#x60;7&#x60;, &#x60;6&#x60;, and &#x60;5&#x60; respectively. If you added two more regions for supporting electable nodes, the priorities of those regions would be &#x60;4&#x60; and &#x60;3&#x60; respectively.  | [optional] 
**ProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisions the hosts.  | [optional] 
**ReadOnlySpecs** | Pointer to [**ApiAtlasClusterHardwareSpecsV15AzureManual**](ApiAtlasClusterHardwareSpecsV15AzureManual.md) |  | [optional] 
**RegionName** | Pointer to **string** | Physical location where MongoDB Cloud deploys your Azure-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts. | [optional] 

## Methods

### NewApiAtlasClusterRegionConfigsV15AzureManual

`func NewApiAtlasClusterRegionConfigsV15AzureManual() *ApiAtlasClusterRegionConfigsV15AzureManual`

NewApiAtlasClusterRegionConfigsV15AzureManual instantiates a new ApiAtlasClusterRegionConfigsV15AzureManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasClusterRegionConfigsV15AzureManualWithDefaults

`func NewApiAtlasClusterRegionConfigsV15AzureManualWithDefaults() *ApiAtlasClusterRegionConfigsV15AzureManual`

NewApiAtlasClusterRegionConfigsV15AzureManualWithDefaults instantiates a new ApiAtlasClusterRegionConfigsV15AzureManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetAnalyticsAutoScaling() ApiAtlasClusterV15AutoScalingAzureManual`

GetAnalyticsAutoScaling returns the AnalyticsAutoScaling field if non-nil, zero value otherwise.

### GetAnalyticsAutoScalingOk

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetAnalyticsAutoScalingOk() (*ApiAtlasClusterV15AutoScalingAzureManual, bool)`

GetAnalyticsAutoScalingOk returns a tuple with the AnalyticsAutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) SetAnalyticsAutoScaling(v ApiAtlasClusterV15AutoScalingAzureManual)`

SetAnalyticsAutoScaling sets AnalyticsAutoScaling field to given value.

### HasAnalyticsAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) HasAnalyticsAutoScaling() bool`

HasAnalyticsAutoScaling returns a boolean if a field has been set.

### GetAnalyticsSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetAnalyticsSpecs() ApiAtlasClusterHardwareSpecsV15AzureManual`

GetAnalyticsSpecs returns the AnalyticsSpecs field if non-nil, zero value otherwise.

### GetAnalyticsSpecsOk

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetAnalyticsSpecsOk() (*ApiAtlasClusterHardwareSpecsV15AzureManual, bool)`

GetAnalyticsSpecsOk returns a tuple with the AnalyticsSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) SetAnalyticsSpecs(v ApiAtlasClusterHardwareSpecsV15AzureManual)`

SetAnalyticsSpecs sets AnalyticsSpecs field to given value.

### HasAnalyticsSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) HasAnalyticsSpecs() bool`

HasAnalyticsSpecs returns a boolean if a field has been set.

### GetAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetAutoScaling() ApiAtlasClusterV15AutoScalingAzureManual`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetAutoScalingOk() (*ApiAtlasClusterV15AutoScalingAzureManual, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) SetAutoScaling(v ApiAtlasClusterV15AutoScalingAzureManual)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetElectableSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetElectableSpecs() ApiAtlasClusterHardwareSpecsV15AzureManual`

GetElectableSpecs returns the ElectableSpecs field if non-nil, zero value otherwise.

### GetElectableSpecsOk

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetElectableSpecsOk() (*ApiAtlasClusterHardwareSpecsV15AzureManual, bool)`

GetElectableSpecsOk returns a tuple with the ElectableSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectableSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) SetElectableSpecs(v ApiAtlasClusterHardwareSpecsV15AzureManual)`

SetElectableSpecs sets ElectableSpecs field to given value.

### HasElectableSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) HasElectableSpecs() bool`

HasElectableSpecs returns a boolean if a field has been set.

### GetPriority

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProviderName

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetReadOnlySpecs

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetReadOnlySpecs() ApiAtlasClusterHardwareSpecsV15AzureManual`

GetReadOnlySpecs returns the ReadOnlySpecs field if non-nil, zero value otherwise.

### GetReadOnlySpecsOk

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetReadOnlySpecsOk() (*ApiAtlasClusterHardwareSpecsV15AzureManual, bool)`

GetReadOnlySpecsOk returns a tuple with the ReadOnlySpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlySpecs

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) SetReadOnlySpecs(v ApiAtlasClusterHardwareSpecsV15AzureManual)`

SetReadOnlySpecs sets ReadOnlySpecs field to given value.

### HasReadOnlySpecs

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) HasReadOnlySpecs() bool`

HasReadOnlySpecs returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasClusterRegionConfigsV15AzureManual) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


