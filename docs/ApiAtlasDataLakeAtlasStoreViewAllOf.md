# ApiAtlasDataLakeAtlasStoreViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Human-readable label of the MongoDB Cloud cluster on which the store is based. | [optional] 
**ProjectId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project. | [optional] [readonly] 
**ReadPreference** | Pointer to [**ApiAtlasDataLakeAtlasStoreReadPreferenceView**](ApiAtlasDataLakeAtlasStoreReadPreferenceView.md) |  | [optional] 

## Methods

### NewApiAtlasDataLakeAtlasStoreViewAllOf

`func NewApiAtlasDataLakeAtlasStoreViewAllOf() *ApiAtlasDataLakeAtlasStoreViewAllOf`

NewApiAtlasDataLakeAtlasStoreViewAllOf instantiates a new ApiAtlasDataLakeAtlasStoreViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDataLakeAtlasStoreViewAllOfWithDefaults

`func NewApiAtlasDataLakeAtlasStoreViewAllOfWithDefaults() *ApiAtlasDataLakeAtlasStoreViewAllOf`

NewApiAtlasDataLakeAtlasStoreViewAllOfWithDefaults instantiates a new ApiAtlasDataLakeAtlasStoreViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetProjectId

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReadPreference

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) GetReadPreference() ApiAtlasDataLakeAtlasStoreReadPreferenceView`

GetReadPreference returns the ReadPreference field if non-nil, zero value otherwise.

### GetReadPreferenceOk

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) GetReadPreferenceOk() (*ApiAtlasDataLakeAtlasStoreReadPreferenceView, bool)`

GetReadPreferenceOk returns a tuple with the ReadPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPreference

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) SetReadPreference(v ApiAtlasDataLakeAtlasStoreReadPreferenceView)`

SetReadPreference sets ReadPreference field to given value.

### HasReadPreference

`func (o *ApiAtlasDataLakeAtlasStoreViewAllOf) HasReadPreference() bool`

HasReadPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


