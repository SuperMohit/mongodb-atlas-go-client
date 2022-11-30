# ApiAtlasReplicationSpecsV15Manual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Multi-Cloud Cluster. If you include existing zones in the request, you must specify this parameter. If you add a new zone to an existing Multi-Cloud Cluster, you may specify this parameter. The request deletes any existing zones in the Multi-Cloud Cluster that you exclude from the request. | [optional] 
**NumShards** | Pointer to **int32** | Positive integer that specifies the number of shards to deploy in each specified zone. If you set this value to &#x60;1&#x60; and &#x60;\&quot;clusterType\&quot; : \&quot;SHARDED\&quot;&#x60;, MongoDB Cloud deploys a single-shard sharded cluster. Don&#39;t create a sharded cluster with a single shard for production environments. Single-shard sharded clusters don&#39;t provide the same benefits as multi-shard configurations. | [optional] 
**RegionConfigs** | Pointer to [**[]ApiAtlasClusterRegionConfigsV15Manual**](ApiAtlasClusterRegionConfigsV15Manual.md) | Hardware specifications for nodes set for a given region. Each **regionConfigs** object describes the region&#39;s priority in elections and the number and type of MongoDB nodes that MongoDB Cloud deploys to the region. Each **regionConfigs** object must have either an **analyticsSpecs** object, **electableSpecs** object, or **readOnlySpecs** object. Tenant clusters only require **electableSpecs. Dedicated** clusters can specify any of these specifications, but must have at least one **electableSpecs** object within a **replicationSpec**. Every hardware specification must use the same **instanceSize**.  **Example:**  If you set &#x60;\&quot;replicationSpecs[n].regionConfigs[m].analyticsSpecs.instanceSize\&quot; : \&quot;M30\&quot;&#x60;, set &#x60;\&quot;replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize\&quot; : &#x60;\&quot;M30\&quot;&#x60; if you have electable nodes and &#x60;\&quot;replicationSpecs[n].regionConfigs[m].readOnlySpecs.instanceSize\&quot; : &#x60;\&quot;M30\&quot;&#x60; if you have read-only nodes.\&quot;, | [optional] 
**ZoneName** | Pointer to **string** | Human-readable label that identifies the zone in a Global Cluster. Provide this value only if &#x60;\&quot;clusterType\&quot; : \&quot;GEOSHARDED\&quot;&#x60;. | [optional] 

## Methods

### NewApiAtlasReplicationSpecsV15Manual

`func NewApiAtlasReplicationSpecsV15Manual() *ApiAtlasReplicationSpecsV15Manual`

NewApiAtlasReplicationSpecsV15Manual instantiates a new ApiAtlasReplicationSpecsV15Manual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasReplicationSpecsV15ManualWithDefaults

`func NewApiAtlasReplicationSpecsV15ManualWithDefaults() *ApiAtlasReplicationSpecsV15Manual`

NewApiAtlasReplicationSpecsV15ManualWithDefaults instantiates a new ApiAtlasReplicationSpecsV15Manual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAtlasReplicationSpecsV15Manual) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasReplicationSpecsV15Manual) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasReplicationSpecsV15Manual) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasReplicationSpecsV15Manual) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumShards

`func (o *ApiAtlasReplicationSpecsV15Manual) GetNumShards() int32`

GetNumShards returns the NumShards field if non-nil, zero value otherwise.

### GetNumShardsOk

`func (o *ApiAtlasReplicationSpecsV15Manual) GetNumShardsOk() (*int32, bool)`

GetNumShardsOk returns a tuple with the NumShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumShards

`func (o *ApiAtlasReplicationSpecsV15Manual) SetNumShards(v int32)`

SetNumShards sets NumShards field to given value.

### HasNumShards

`func (o *ApiAtlasReplicationSpecsV15Manual) HasNumShards() bool`

HasNumShards returns a boolean if a field has been set.

### GetRegionConfigs

`func (o *ApiAtlasReplicationSpecsV15Manual) GetRegionConfigs() []ApiAtlasClusterRegionConfigsV15Manual`

GetRegionConfigs returns the RegionConfigs field if non-nil, zero value otherwise.

### GetRegionConfigsOk

`func (o *ApiAtlasReplicationSpecsV15Manual) GetRegionConfigsOk() (*[]ApiAtlasClusterRegionConfigsV15Manual, bool)`

GetRegionConfigsOk returns a tuple with the RegionConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionConfigs

`func (o *ApiAtlasReplicationSpecsV15Manual) SetRegionConfigs(v []ApiAtlasClusterRegionConfigsV15Manual)`

SetRegionConfigs sets RegionConfigs field to given value.

### HasRegionConfigs

`func (o *ApiAtlasReplicationSpecsV15Manual) HasRegionConfigs() bool`

HasRegionConfigs returns a boolean if a field has been set.

### GetZoneName

`func (o *ApiAtlasReplicationSpecsV15Manual) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ApiAtlasReplicationSpecsV15Manual) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ApiAtlasReplicationSpecsV15Manual) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ApiAtlasReplicationSpecsV15Manual) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


