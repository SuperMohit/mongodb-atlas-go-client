# ApiAtlasClusterRegionConfigsV15AWSManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsAutoScaling** | Pointer to [**ApiAtlasClusterV15AutoScalingAWSManual**](ApiAtlasClusterV15AutoScalingAWSManual.md) |  | [optional] 
**AnalyticsSpecs** | Pointer to [**ApiAtlasClusterHardwareSpecsV15AWSManual**](ApiAtlasClusterHardwareSpecsV15AWSManual.md) |  | [optional] 
**AutoScaling** | Pointer to [**ApiAtlasClusterV15AutoScalingAWSManual**](ApiAtlasClusterV15AutoScalingAWSManual.md) |  | [optional] 
**ElectableSpecs** | Pointer to [**ApiAtlasClusterHardwareSpecsV15AWSManual**](ApiAtlasClusterHardwareSpecsV15AWSManual.md) |  | [optional] 
**Priority** | Pointer to **int32** | Precedence that MongoDB Cloud gives this region when a primary election occurs. If your **regionConfigs** has only **readOnlySpecs**, **analyticsSpecs**, or both, set this value to &#x60;0&#x60;. If you have multiple **regionConfigs** objects, set priorities in descending order starting from &#x60;7&#x60;.  Example: If you have three regions, their priorities would be &#x60;7&#x60;, &#x60;6&#x60;, and &#x60;5&#x60; respectively. If you added two more regions for supporting electable nodes, the priorities of those regions would be &#x60;4&#x60; and &#x60;3&#x60; respectively.  | [optional] 
**ProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisions the hosts.  | [optional] 
**ReadOnlySpecs** | Pointer to [**ApiAtlasClusterHardwareSpecsV15AWSManual**](ApiAtlasClusterHardwareSpecsV15AWSManual.md) |  | [optional] 
**RegionName** | Pointer to **string** | Physical location where MongoDB Cloud deploys your AWS-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts. | [optional] 

## Methods

### NewApiAtlasClusterRegionConfigsV15AWSManual

`func NewApiAtlasClusterRegionConfigsV15AWSManual() *ApiAtlasClusterRegionConfigsV15AWSManual`

NewApiAtlasClusterRegionConfigsV15AWSManual instantiates a new ApiAtlasClusterRegionConfigsV15AWSManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasClusterRegionConfigsV15AWSManualWithDefaults

`func NewApiAtlasClusterRegionConfigsV15AWSManualWithDefaults() *ApiAtlasClusterRegionConfigsV15AWSManual`

NewApiAtlasClusterRegionConfigsV15AWSManualWithDefaults instantiates a new ApiAtlasClusterRegionConfigsV15AWSManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetAnalyticsAutoScaling() ApiAtlasClusterV15AutoScalingAWSManual`

GetAnalyticsAutoScaling returns the AnalyticsAutoScaling field if non-nil, zero value otherwise.

### GetAnalyticsAutoScalingOk

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetAnalyticsAutoScalingOk() (*ApiAtlasClusterV15AutoScalingAWSManual, bool)`

GetAnalyticsAutoScalingOk returns a tuple with the AnalyticsAutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) SetAnalyticsAutoScaling(v ApiAtlasClusterV15AutoScalingAWSManual)`

SetAnalyticsAutoScaling sets AnalyticsAutoScaling field to given value.

### HasAnalyticsAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) HasAnalyticsAutoScaling() bool`

HasAnalyticsAutoScaling returns a boolean if a field has been set.

### GetAnalyticsSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetAnalyticsSpecs() ApiAtlasClusterHardwareSpecsV15AWSManual`

GetAnalyticsSpecs returns the AnalyticsSpecs field if non-nil, zero value otherwise.

### GetAnalyticsSpecsOk

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetAnalyticsSpecsOk() (*ApiAtlasClusterHardwareSpecsV15AWSManual, bool)`

GetAnalyticsSpecsOk returns a tuple with the AnalyticsSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) SetAnalyticsSpecs(v ApiAtlasClusterHardwareSpecsV15AWSManual)`

SetAnalyticsSpecs sets AnalyticsSpecs field to given value.

### HasAnalyticsSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) HasAnalyticsSpecs() bool`

HasAnalyticsSpecs returns a boolean if a field has been set.

### GetAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetAutoScaling() ApiAtlasClusterV15AutoScalingAWSManual`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetAutoScalingOk() (*ApiAtlasClusterV15AutoScalingAWSManual, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) SetAutoScaling(v ApiAtlasClusterV15AutoScalingAWSManual)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetElectableSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetElectableSpecs() ApiAtlasClusterHardwareSpecsV15AWSManual`

GetElectableSpecs returns the ElectableSpecs field if non-nil, zero value otherwise.

### GetElectableSpecsOk

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetElectableSpecsOk() (*ApiAtlasClusterHardwareSpecsV15AWSManual, bool)`

GetElectableSpecsOk returns a tuple with the ElectableSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectableSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) SetElectableSpecs(v ApiAtlasClusterHardwareSpecsV15AWSManual)`

SetElectableSpecs sets ElectableSpecs field to given value.

### HasElectableSpecs

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) HasElectableSpecs() bool`

HasElectableSpecs returns a boolean if a field has been set.

### GetPriority

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProviderName

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetReadOnlySpecs

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetReadOnlySpecs() ApiAtlasClusterHardwareSpecsV15AWSManual`

GetReadOnlySpecs returns the ReadOnlySpecs field if non-nil, zero value otherwise.

### GetReadOnlySpecsOk

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetReadOnlySpecsOk() (*ApiAtlasClusterHardwareSpecsV15AWSManual, bool)`

GetReadOnlySpecsOk returns a tuple with the ReadOnlySpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlySpecs

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) SetReadOnlySpecs(v ApiAtlasClusterHardwareSpecsV15AWSManual)`

SetReadOnlySpecs sets ReadOnlySpecs field to given value.

### HasReadOnlySpecs

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) HasReadOnlySpecs() bool`

HasReadOnlySpecs returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasClusterRegionConfigsV15AWSManual) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


