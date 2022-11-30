# ManagedNamespaceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to **string** | Human-readable label of the collection to manage for this Global Cluster. | [optional] 
**CustomShardKey** | Pointer to **string** | Database parameter used to divide the *collection* into shards. Global clusters require a compound shard key. This compound shard key combines the location parameter and the user-selected custom key. | [optional] 
**Db** | Pointer to **string** | Human-readable label of the database to manage for this Global Cluster. | [optional] 
**IsCustomShardKeyHashed** | Pointer to **bool** | Flag that indicates whether someone hashed the custom shard key. If this parameter returns &#x60;false&#x60;, this cluster uses ranged sharding. | [optional] [default to false]
**IsShardKeyUnique** | Pointer to **bool** | Flag that indicates whether the underlying index enforces unique values. | [optional] [default to false]
**NumInitialChunks** | Pointer to **int64** | Minimum number of chunks to create initially when sharding an empty collection with a hashed shard key. | [optional] 
**PresplitHashedZones** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should create and distribute initial chunks for an empty or non-existing collection. MongoDB Cloud distributes data based on the defined zones and zone ranges for the collection. | [optional] [default to false]

## Methods

### NewManagedNamespaceView

`func NewManagedNamespaceView() *ManagedNamespaceView`

NewManagedNamespaceView instantiates a new ManagedNamespaceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedNamespaceViewWithDefaults

`func NewManagedNamespaceViewWithDefaults() *ManagedNamespaceView`

NewManagedNamespaceViewWithDefaults instantiates a new ManagedNamespaceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *ManagedNamespaceView) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ManagedNamespaceView) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ManagedNamespaceView) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *ManagedNamespaceView) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetCustomShardKey

`func (o *ManagedNamespaceView) GetCustomShardKey() string`

GetCustomShardKey returns the CustomShardKey field if non-nil, zero value otherwise.

### GetCustomShardKeyOk

`func (o *ManagedNamespaceView) GetCustomShardKeyOk() (*string, bool)`

GetCustomShardKeyOk returns a tuple with the CustomShardKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomShardKey

`func (o *ManagedNamespaceView) SetCustomShardKey(v string)`

SetCustomShardKey sets CustomShardKey field to given value.

### HasCustomShardKey

`func (o *ManagedNamespaceView) HasCustomShardKey() bool`

HasCustomShardKey returns a boolean if a field has been set.

### GetDb

`func (o *ManagedNamespaceView) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *ManagedNamespaceView) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *ManagedNamespaceView) SetDb(v string)`

SetDb sets Db field to given value.

### HasDb

`func (o *ManagedNamespaceView) HasDb() bool`

HasDb returns a boolean if a field has been set.

### GetIsCustomShardKeyHashed

`func (o *ManagedNamespaceView) GetIsCustomShardKeyHashed() bool`

GetIsCustomShardKeyHashed returns the IsCustomShardKeyHashed field if non-nil, zero value otherwise.

### GetIsCustomShardKeyHashedOk

`func (o *ManagedNamespaceView) GetIsCustomShardKeyHashedOk() (*bool, bool)`

GetIsCustomShardKeyHashedOk returns a tuple with the IsCustomShardKeyHashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomShardKeyHashed

`func (o *ManagedNamespaceView) SetIsCustomShardKeyHashed(v bool)`

SetIsCustomShardKeyHashed sets IsCustomShardKeyHashed field to given value.

### HasIsCustomShardKeyHashed

`func (o *ManagedNamespaceView) HasIsCustomShardKeyHashed() bool`

HasIsCustomShardKeyHashed returns a boolean if a field has been set.

### GetIsShardKeyUnique

`func (o *ManagedNamespaceView) GetIsShardKeyUnique() bool`

GetIsShardKeyUnique returns the IsShardKeyUnique field if non-nil, zero value otherwise.

### GetIsShardKeyUniqueOk

`func (o *ManagedNamespaceView) GetIsShardKeyUniqueOk() (*bool, bool)`

GetIsShardKeyUniqueOk returns a tuple with the IsShardKeyUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShardKeyUnique

`func (o *ManagedNamespaceView) SetIsShardKeyUnique(v bool)`

SetIsShardKeyUnique sets IsShardKeyUnique field to given value.

### HasIsShardKeyUnique

`func (o *ManagedNamespaceView) HasIsShardKeyUnique() bool`

HasIsShardKeyUnique returns a boolean if a field has been set.

### GetNumInitialChunks

`func (o *ManagedNamespaceView) GetNumInitialChunks() int64`

GetNumInitialChunks returns the NumInitialChunks field if non-nil, zero value otherwise.

### GetNumInitialChunksOk

`func (o *ManagedNamespaceView) GetNumInitialChunksOk() (*int64, bool)`

GetNumInitialChunksOk returns a tuple with the NumInitialChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInitialChunks

`func (o *ManagedNamespaceView) SetNumInitialChunks(v int64)`

SetNumInitialChunks sets NumInitialChunks field to given value.

### HasNumInitialChunks

`func (o *ManagedNamespaceView) HasNumInitialChunks() bool`

HasNumInitialChunks returns a boolean if a field has been set.

### GetPresplitHashedZones

`func (o *ManagedNamespaceView) GetPresplitHashedZones() bool`

GetPresplitHashedZones returns the PresplitHashedZones field if non-nil, zero value otherwise.

### GetPresplitHashedZonesOk

`func (o *ManagedNamespaceView) GetPresplitHashedZonesOk() (*bool, bool)`

GetPresplitHashedZonesOk returns a tuple with the PresplitHashedZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresplitHashedZones

`func (o *ManagedNamespaceView) SetPresplitHashedZones(v bool)`

SetPresplitHashedZones sets PresplitHashedZones field to given value.

### HasPresplitHashedZones

`func (o *ManagedNamespaceView) HasPresplitHashedZones() bool`

HasPresplitHashedZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


