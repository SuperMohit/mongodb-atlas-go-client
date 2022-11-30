# ApiAtlasRegionSpecView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsNodes** | Pointer to **int32** | Number of analytics nodes in the region. Analytics nodes handle analytic data such as reporting queries from MongoDB Connector for Business Intelligence on MongoDB Cloud. Analytics nodes are read-only, and can never become the primary. Use **replicationSpecs[n].{region}.analyticsNodes** instead. | [optional] 
**ElectableNodes** | Pointer to **int32** | Number of electable nodes to deploy in the specified region. Electable nodes can become the primary and can facilitate local reads. Use **replicationSpecs[n].{region}.electableNodes** instead. | [optional] 
**Priority** | Pointer to **int32** | Number that indicates the election priority of the region. To identify the Preferred Region of the cluster, set this parameter to &#x60;7&#x60;. The primary node runs in the **Preferred Region**. To identify a read-only region, set this parameter to &#x60;0&#x60;. | [optional] 
**ReadOnlyNodes** | Pointer to **int32** | Number of read-only nodes in the region. Read-only nodes can never become the primary member, but can facilitate local reads. Use **replicationSpecs[n].{region}.readOnlyNodes** instead. | [optional] 

## Methods

### NewApiAtlasRegionSpecView

`func NewApiAtlasRegionSpecView() *ApiAtlasRegionSpecView`

NewApiAtlasRegionSpecView instantiates a new ApiAtlasRegionSpecView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasRegionSpecViewWithDefaults

`func NewApiAtlasRegionSpecViewWithDefaults() *ApiAtlasRegionSpecView`

NewApiAtlasRegionSpecViewWithDefaults instantiates a new ApiAtlasRegionSpecView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsNodes

`func (o *ApiAtlasRegionSpecView) GetAnalyticsNodes() int32`

GetAnalyticsNodes returns the AnalyticsNodes field if non-nil, zero value otherwise.

### GetAnalyticsNodesOk

`func (o *ApiAtlasRegionSpecView) GetAnalyticsNodesOk() (*int32, bool)`

GetAnalyticsNodesOk returns a tuple with the AnalyticsNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsNodes

`func (o *ApiAtlasRegionSpecView) SetAnalyticsNodes(v int32)`

SetAnalyticsNodes sets AnalyticsNodes field to given value.

### HasAnalyticsNodes

`func (o *ApiAtlasRegionSpecView) HasAnalyticsNodes() bool`

HasAnalyticsNodes returns a boolean if a field has been set.

### GetElectableNodes

`func (o *ApiAtlasRegionSpecView) GetElectableNodes() int32`

GetElectableNodes returns the ElectableNodes field if non-nil, zero value otherwise.

### GetElectableNodesOk

`func (o *ApiAtlasRegionSpecView) GetElectableNodesOk() (*int32, bool)`

GetElectableNodesOk returns a tuple with the ElectableNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectableNodes

`func (o *ApiAtlasRegionSpecView) SetElectableNodes(v int32)`

SetElectableNodes sets ElectableNodes field to given value.

### HasElectableNodes

`func (o *ApiAtlasRegionSpecView) HasElectableNodes() bool`

HasElectableNodes returns a boolean if a field has been set.

### GetPriority

`func (o *ApiAtlasRegionSpecView) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApiAtlasRegionSpecView) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApiAtlasRegionSpecView) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApiAtlasRegionSpecView) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReadOnlyNodes

`func (o *ApiAtlasRegionSpecView) GetReadOnlyNodes() int32`

GetReadOnlyNodes returns the ReadOnlyNodes field if non-nil, zero value otherwise.

### GetReadOnlyNodesOk

`func (o *ApiAtlasRegionSpecView) GetReadOnlyNodesOk() (*int32, bool)`

GetReadOnlyNodesOk returns a tuple with the ReadOnlyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyNodes

`func (o *ApiAtlasRegionSpecView) SetReadOnlyNodes(v int32)`

SetReadOnlyNodes sets ReadOnlyNodes field to given value.

### HasReadOnlyNodes

`func (o *ApiAtlasRegionSpecView) HasReadOnlyNodes() bool`

HasReadOnlyNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


