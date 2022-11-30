# ApiAtlasLegacyReplicationSpecViewManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Global Cluster. - If you include existing zones in the request, you must specify this parameter. - If you add a new zone to an existing Global Cluster, you may specify this   parameter. The request deletes any existing zones in a Global Cluster that you   exclude from the request. | [optional] 
**NumShards** | Pointer to **float32** | Positive integer that specifies the number of shards to deploy in each specified zone. If you set this value to &#x60;1&#x60; and &#x60;\&quot;clusterType\&quot; : \&quot;SHARDED\&quot;&#x60;, MongoDB Cloud deploys a single-shard sharded cluster. Don&#39;t create a sharded cluster with a single shard for production environments. Single-shard sharded clusters don&#39;t provide the same benefits as multi-shard configurations. | [optional] [default to 1]
**RegionsConfig** | Pointer to [**ApiAtlasReplSpecsRegionViewManual**](ApiAtlasReplSpecsRegionViewManual.md) |  | [optional] 
**ZoneName** | Pointer to **string** | Human-readable label that identifies the zone in a Global Cluster. Provide this value only if &#x60;\&quot;clusterType\&quot; : \&quot;GEOSHARDED\&quot;&#x60;. | [optional] 

## Methods

### NewApiAtlasLegacyReplicationSpecViewManual

`func NewApiAtlasLegacyReplicationSpecViewManual() *ApiAtlasLegacyReplicationSpecViewManual`

NewApiAtlasLegacyReplicationSpecViewManual instantiates a new ApiAtlasLegacyReplicationSpecViewManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasLegacyReplicationSpecViewManualWithDefaults

`func NewApiAtlasLegacyReplicationSpecViewManualWithDefaults() *ApiAtlasLegacyReplicationSpecViewManual`

NewApiAtlasLegacyReplicationSpecViewManualWithDefaults instantiates a new ApiAtlasLegacyReplicationSpecViewManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAtlasLegacyReplicationSpecViewManual) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasLegacyReplicationSpecViewManual) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasLegacyReplicationSpecViewManual) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasLegacyReplicationSpecViewManual) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumShards

`func (o *ApiAtlasLegacyReplicationSpecViewManual) GetNumShards() float32`

GetNumShards returns the NumShards field if non-nil, zero value otherwise.

### GetNumShardsOk

`func (o *ApiAtlasLegacyReplicationSpecViewManual) GetNumShardsOk() (*float32, bool)`

GetNumShardsOk returns a tuple with the NumShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumShards

`func (o *ApiAtlasLegacyReplicationSpecViewManual) SetNumShards(v float32)`

SetNumShards sets NumShards field to given value.

### HasNumShards

`func (o *ApiAtlasLegacyReplicationSpecViewManual) HasNumShards() bool`

HasNumShards returns a boolean if a field has been set.

### GetRegionsConfig

`func (o *ApiAtlasLegacyReplicationSpecViewManual) GetRegionsConfig() ApiAtlasReplSpecsRegionViewManual`

GetRegionsConfig returns the RegionsConfig field if non-nil, zero value otherwise.

### GetRegionsConfigOk

`func (o *ApiAtlasLegacyReplicationSpecViewManual) GetRegionsConfigOk() (*ApiAtlasReplSpecsRegionViewManual, bool)`

GetRegionsConfigOk returns a tuple with the RegionsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionsConfig

`func (o *ApiAtlasLegacyReplicationSpecViewManual) SetRegionsConfig(v ApiAtlasReplSpecsRegionViewManual)`

SetRegionsConfig sets RegionsConfig field to given value.

### HasRegionsConfig

`func (o *ApiAtlasLegacyReplicationSpecViewManual) HasRegionsConfig() bool`

HasRegionsConfig returns a boolean if a field has been set.

### GetZoneName

`func (o *ApiAtlasLegacyReplicationSpecViewManual) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ApiAtlasLegacyReplicationSpecViewManual) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ApiAtlasLegacyReplicationSpecViewManual) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ApiAtlasLegacyReplicationSpecViewManual) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


