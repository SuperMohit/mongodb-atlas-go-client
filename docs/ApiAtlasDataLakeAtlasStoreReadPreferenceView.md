# ApiAtlasDataLakeAtlasStoreReadPreferenceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxStalenessSeconds** | Pointer to **int32** | Maximum replication lag, or **staleness**, for reads from secondaries. | [optional] 
**Mode** | Pointer to **string** | [Read preference mode](https://docs.mongodb.com/manual/core/read-preference/#read-preference-modes) that specifies to which replica set member to route the read requests. | [optional] 
**TagSets** | Pointer to [**[][]ApiAtlasDataLakeAtlasStoreReadPreferenceTagView**]([]ApiAtlasDataLakeAtlasStoreReadPreferenceTagView.md) | List that contains [tag sets](https://docs.mongodb.com/manual/core/read-preference-tags/) or tag specification documents. If specified, Atlas Data Lake routes read requests to replica set member or members that are associated with the specified tags. | [optional] 

## Methods

### NewApiAtlasDataLakeAtlasStoreReadPreferenceView

`func NewApiAtlasDataLakeAtlasStoreReadPreferenceView() *ApiAtlasDataLakeAtlasStoreReadPreferenceView`

NewApiAtlasDataLakeAtlasStoreReadPreferenceView instantiates a new ApiAtlasDataLakeAtlasStoreReadPreferenceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDataLakeAtlasStoreReadPreferenceViewWithDefaults

`func NewApiAtlasDataLakeAtlasStoreReadPreferenceViewWithDefaults() *ApiAtlasDataLakeAtlasStoreReadPreferenceView`

NewApiAtlasDataLakeAtlasStoreReadPreferenceViewWithDefaults instantiates a new ApiAtlasDataLakeAtlasStoreReadPreferenceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxStalenessSeconds

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) GetMaxStalenessSeconds() int32`

GetMaxStalenessSeconds returns the MaxStalenessSeconds field if non-nil, zero value otherwise.

### GetMaxStalenessSecondsOk

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) GetMaxStalenessSecondsOk() (*int32, bool)`

GetMaxStalenessSecondsOk returns a tuple with the MaxStalenessSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStalenessSeconds

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) SetMaxStalenessSeconds(v int32)`

SetMaxStalenessSeconds sets MaxStalenessSeconds field to given value.

### HasMaxStalenessSeconds

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) HasMaxStalenessSeconds() bool`

HasMaxStalenessSeconds returns a boolean if a field has been set.

### GetMode

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTagSets

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) GetTagSets() [][]ApiAtlasDataLakeAtlasStoreReadPreferenceTagView`

GetTagSets returns the TagSets field if non-nil, zero value otherwise.

### GetTagSetsOk

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) GetTagSetsOk() (*[][]ApiAtlasDataLakeAtlasStoreReadPreferenceTagView, bool)`

GetTagSetsOk returns a tuple with the TagSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagSets

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) SetTagSets(v [][]ApiAtlasDataLakeAtlasStoreReadPreferenceTagView)`

SetTagSets sets TagSets field to given value.

### HasTagSets

`func (o *ApiAtlasDataLakeAtlasStoreReadPreferenceView) HasTagSets() bool`

HasTagSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


