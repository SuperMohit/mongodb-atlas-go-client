# ApiAtlasReplSpecsRegionSpecViewManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsNodes** | Pointer to **float32** | Number of analytics nodes to deploy in the specified region. Analytics nodes handle dedicated queries for reporting such as those from MongoDB Connector for Business Intelligence on MongoDB Cloud. You can&#39;t write to an analytics node. Elections can&#39;t make an analytic node the primary node. | [optional] 
**ElectableNodes** | Pointer to **float32** | Number of electable nodes to deploy in the specified region. Electable nodes can become the primary and can facilitate local reads. | [optional] 
**Priority** | Pointer to **float32** | Number that indicates the election priority of the region. - To identify the preferred region of the cluster, set this parameter to &#x60;7&#x60;.   The primary node runs in the **Preferred Region**. - To identify a read-only region, set this parameter to &#x60;0&#x60;. | [optional] 
**ReadOnlyNodes** | Pointer to **float32** | Number of read-only nodes to deploy in the specified region. Read-only nodes can never become the primary member, but can facilitate local reads. | [optional] 

## Methods

### NewApiAtlasReplSpecsRegionSpecViewManual

`func NewApiAtlasReplSpecsRegionSpecViewManual() *ApiAtlasReplSpecsRegionSpecViewManual`

NewApiAtlasReplSpecsRegionSpecViewManual instantiates a new ApiAtlasReplSpecsRegionSpecViewManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasReplSpecsRegionSpecViewManualWithDefaults

`func NewApiAtlasReplSpecsRegionSpecViewManualWithDefaults() *ApiAtlasReplSpecsRegionSpecViewManual`

NewApiAtlasReplSpecsRegionSpecViewManualWithDefaults instantiates a new ApiAtlasReplSpecsRegionSpecViewManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsNodes

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) GetAnalyticsNodes() float32`

GetAnalyticsNodes returns the AnalyticsNodes field if non-nil, zero value otherwise.

### GetAnalyticsNodesOk

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) GetAnalyticsNodesOk() (*float32, bool)`

GetAnalyticsNodesOk returns a tuple with the AnalyticsNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsNodes

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) SetAnalyticsNodes(v float32)`

SetAnalyticsNodes sets AnalyticsNodes field to given value.

### HasAnalyticsNodes

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) HasAnalyticsNodes() bool`

HasAnalyticsNodes returns a boolean if a field has been set.

### GetElectableNodes

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) GetElectableNodes() float32`

GetElectableNodes returns the ElectableNodes field if non-nil, zero value otherwise.

### GetElectableNodesOk

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) GetElectableNodesOk() (*float32, bool)`

GetElectableNodesOk returns a tuple with the ElectableNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectableNodes

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) SetElectableNodes(v float32)`

SetElectableNodes sets ElectableNodes field to given value.

### HasElectableNodes

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) HasElectableNodes() bool`

HasElectableNodes returns a boolean if a field has been set.

### GetPriority

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReadOnlyNodes

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) GetReadOnlyNodes() float32`

GetReadOnlyNodes returns the ReadOnlyNodes field if non-nil, zero value otherwise.

### GetReadOnlyNodesOk

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) GetReadOnlyNodesOk() (*float32, bool)`

GetReadOnlyNodesOk returns a tuple with the ReadOnlyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyNodes

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) SetReadOnlyNodes(v float32)`

SetReadOnlyNodes sets ReadOnlyNodes field to given value.

### HasReadOnlyNodes

`func (o *ApiAtlasReplSpecsRegionSpecViewManual) HasReadOnlyNodes() bool`

HasReadOnlyNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


