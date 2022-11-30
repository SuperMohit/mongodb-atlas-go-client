# ApiAtlasClusterRegionConfigsV15TenantManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsSpecs** | Pointer to [**ApiAtlasClusterHardwareSpecsV15TenantManual**](ApiAtlasClusterHardwareSpecsV15TenantManual.md) |  | [optional] 
**BackingProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisioned the tenant cluster. | [optional] 
**ElectableSpecs** | Pointer to [**ApiAtlasClusterHardwareSpecsV15TenantManual**](ApiAtlasClusterHardwareSpecsV15TenantManual.md) |  | [optional] 
**Priority** | Pointer to **int32** | Precedence that MongoDB Cloud gives this region when a primary election occurs. If your **regionConfigs** has only **readOnlySpecs**, **analyticsSpecs**, or both, set this value to &#x60;0&#x60;. If you have multiple **regionConfigs** objects, set priorities in descending order starting from &#x60;7&#x60;.  Example: If you have three regions, their priorities would be &#x60;7&#x60;, &#x60;6&#x60;, and &#x60;5&#x60; respectively. If you added two more regions for supporting electable nodes, the priorities of those regions would be &#x60;4&#x60; and &#x60;3&#x60; respectively.  | [optional] 
**ProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisions the MongoDB nodes.  | [optional] 
**ReadOnlySpecs** | Pointer to [**ApiAtlasClusterHardwareSpecsV15TenantManual**](ApiAtlasClusterHardwareSpecsV15TenantManual.md) |  | [optional] 
**RegionName** | Pointer to [**ApiAtlasClusterTenantRegionsV15Manual**](ApiAtlasClusterTenantRegionsV15Manual.md) |  | [optional] 

## Methods

### NewApiAtlasClusterRegionConfigsV15TenantManual

`func NewApiAtlasClusterRegionConfigsV15TenantManual() *ApiAtlasClusterRegionConfigsV15TenantManual`

NewApiAtlasClusterRegionConfigsV15TenantManual instantiates a new ApiAtlasClusterRegionConfigsV15TenantManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasClusterRegionConfigsV15TenantManualWithDefaults

`func NewApiAtlasClusterRegionConfigsV15TenantManualWithDefaults() *ApiAtlasClusterRegionConfigsV15TenantManual`

NewApiAtlasClusterRegionConfigsV15TenantManualWithDefaults instantiates a new ApiAtlasClusterRegionConfigsV15TenantManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsSpecs

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetAnalyticsSpecs() ApiAtlasClusterHardwareSpecsV15TenantManual`

GetAnalyticsSpecs returns the AnalyticsSpecs field if non-nil, zero value otherwise.

### GetAnalyticsSpecsOk

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetAnalyticsSpecsOk() (*ApiAtlasClusterHardwareSpecsV15TenantManual, bool)`

GetAnalyticsSpecsOk returns a tuple with the AnalyticsSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsSpecs

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) SetAnalyticsSpecs(v ApiAtlasClusterHardwareSpecsV15TenantManual)`

SetAnalyticsSpecs sets AnalyticsSpecs field to given value.

### HasAnalyticsSpecs

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) HasAnalyticsSpecs() bool`

HasAnalyticsSpecs returns a boolean if a field has been set.

### GetBackingProviderName

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetBackingProviderName() string`

GetBackingProviderName returns the BackingProviderName field if non-nil, zero value otherwise.

### GetBackingProviderNameOk

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetBackingProviderNameOk() (*string, bool)`

GetBackingProviderNameOk returns a tuple with the BackingProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingProviderName

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) SetBackingProviderName(v string)`

SetBackingProviderName sets BackingProviderName field to given value.

### HasBackingProviderName

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) HasBackingProviderName() bool`

HasBackingProviderName returns a boolean if a field has been set.

### GetElectableSpecs

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetElectableSpecs() ApiAtlasClusterHardwareSpecsV15TenantManual`

GetElectableSpecs returns the ElectableSpecs field if non-nil, zero value otherwise.

### GetElectableSpecsOk

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetElectableSpecsOk() (*ApiAtlasClusterHardwareSpecsV15TenantManual, bool)`

GetElectableSpecsOk returns a tuple with the ElectableSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectableSpecs

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) SetElectableSpecs(v ApiAtlasClusterHardwareSpecsV15TenantManual)`

SetElectableSpecs sets ElectableSpecs field to given value.

### HasElectableSpecs

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) HasElectableSpecs() bool`

HasElectableSpecs returns a boolean if a field has been set.

### GetPriority

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProviderName

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetReadOnlySpecs

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetReadOnlySpecs() ApiAtlasClusterHardwareSpecsV15TenantManual`

GetReadOnlySpecs returns the ReadOnlySpecs field if non-nil, zero value otherwise.

### GetReadOnlySpecsOk

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetReadOnlySpecsOk() (*ApiAtlasClusterHardwareSpecsV15TenantManual, bool)`

GetReadOnlySpecsOk returns a tuple with the ReadOnlySpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlySpecs

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) SetReadOnlySpecs(v ApiAtlasClusterHardwareSpecsV15TenantManual)`

SetReadOnlySpecs sets ReadOnlySpecs field to given value.

### HasReadOnlySpecs

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) HasReadOnlySpecs() bool`

HasReadOnlySpecs returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetRegionName() ApiAtlasClusterTenantRegionsV15Manual`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) GetRegionNameOk() (*ApiAtlasClusterTenantRegionsV15Manual, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) SetRegionName(v ApiAtlasClusterTenantRegionsV15Manual)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasClusterRegionConfigsV15TenantManual) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


