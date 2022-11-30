# ApiAtlasClusterOutageSimulationOutageFilterView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | The cloud provider of the region that undergoes the outage simulation. | [optional] 
**RegionName** | Pointer to **string** | The name of the region to undergo an outage simulation. | [optional] 
**Type** | Pointer to **string** | The type of cluster outage to simulate.  | Type       | Description | |------------|-------------| | &#x60;REGION&#x60;   | Simulates a cluster outage for a region.| | [optional] 

## Methods

### NewApiAtlasClusterOutageSimulationOutageFilterView

`func NewApiAtlasClusterOutageSimulationOutageFilterView() *ApiAtlasClusterOutageSimulationOutageFilterView`

NewApiAtlasClusterOutageSimulationOutageFilterView instantiates a new ApiAtlasClusterOutageSimulationOutageFilterView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasClusterOutageSimulationOutageFilterViewWithDefaults

`func NewApiAtlasClusterOutageSimulationOutageFilterViewWithDefaults() *ApiAtlasClusterOutageSimulationOutageFilterView`

NewApiAtlasClusterOutageSimulationOutageFilterViewWithDefaults instantiates a new ApiAtlasClusterOutageSimulationOutageFilterView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetType

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiAtlasClusterOutageSimulationOutageFilterView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


