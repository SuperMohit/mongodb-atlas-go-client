# ApiAtlasClusterV15ComputeAutoScalingGCPManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Flag that indicates whether someone enabled instance size auto-scaling.  -  Set to &#x60;true&#x60; to enable instance size auto-scaling. If enabled, you must specify a value for **replicationSpecs[n].regionConfigs[m].autoScaling.compute.maxInstanceSize**. -  Set to &#x60;false&#x60; to disable instance size automatic scaling.  | [optional] 
**MaxInstanceSize** | Pointer to **string** | Maximum instance size to which your cluster can automatically scale. MongoDB Cloud requires this parameter if &#x60;\&quot;replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled\&quot; : true&#x60;. | [optional] 
**MinInstanceSize** | Pointer to **string** | Minimum instance size to which your cluster can automatically scale. MongoDB Cloud requires this parameter if &#x60;\&quot;replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled\&quot; : true&#x60;. | [optional] 
**ScaleDownEnabled** | Pointer to **bool** | Flag that indicates whether the instance size may scale down. MongoDB Cloud requires this parameter if &#x60;\&quot;replicationSpecs[n].regionConfigs[m].autoScaling.compute.enabled\&quot; : true&#x60;. If you enable this option, specify a value for **replicationSpecs[n].regionConfigs[m].autoScaling.compute.minInstanceSize**. | [optional] 

## Methods

### NewApiAtlasClusterV15ComputeAutoScalingGCPManual

`func NewApiAtlasClusterV15ComputeAutoScalingGCPManual() *ApiAtlasClusterV15ComputeAutoScalingGCPManual`

NewApiAtlasClusterV15ComputeAutoScalingGCPManual instantiates a new ApiAtlasClusterV15ComputeAutoScalingGCPManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasClusterV15ComputeAutoScalingGCPManualWithDefaults

`func NewApiAtlasClusterV15ComputeAutoScalingGCPManualWithDefaults() *ApiAtlasClusterV15ComputeAutoScalingGCPManual`

NewApiAtlasClusterV15ComputeAutoScalingGCPManualWithDefaults instantiates a new ApiAtlasClusterV15ComputeAutoScalingGCPManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxInstanceSize

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) GetMaxInstanceSize() string`

GetMaxInstanceSize returns the MaxInstanceSize field if non-nil, zero value otherwise.

### GetMaxInstanceSizeOk

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) GetMaxInstanceSizeOk() (*string, bool)`

GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstanceSize

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) SetMaxInstanceSize(v string)`

SetMaxInstanceSize sets MaxInstanceSize field to given value.

### HasMaxInstanceSize

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) HasMaxInstanceSize() bool`

HasMaxInstanceSize returns a boolean if a field has been set.

### GetMinInstanceSize

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) GetMinInstanceSize() string`

GetMinInstanceSize returns the MinInstanceSize field if non-nil, zero value otherwise.

### GetMinInstanceSizeOk

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) GetMinInstanceSizeOk() (*string, bool)`

GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInstanceSize

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) SetMinInstanceSize(v string)`

SetMinInstanceSize sets MinInstanceSize field to given value.

### HasMinInstanceSize

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) HasMinInstanceSize() bool`

HasMinInstanceSize returns a boolean if a field has been set.

### GetScaleDownEnabled

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) GetScaleDownEnabled() bool`

GetScaleDownEnabled returns the ScaleDownEnabled field if non-nil, zero value otherwise.

### GetScaleDownEnabledOk

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) GetScaleDownEnabledOk() (*bool, bool)`

GetScaleDownEnabledOk returns a tuple with the ScaleDownEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleDownEnabled

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) SetScaleDownEnabled(v bool)`

SetScaleDownEnabled sets ScaleDownEnabled field to given value.

### HasScaleDownEnabled

`func (o *ApiAtlasClusterV15ComputeAutoScalingGCPManual) HasScaleDownEnabled() bool`

HasScaleDownEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


