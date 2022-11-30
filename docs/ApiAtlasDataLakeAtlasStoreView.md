# ApiAtlasDataLakeAtlasStoreView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Human-readable label of the MongoDB Cloud cluster on which the store is based. | [optional] 
**ProjectId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project. | [optional] [readonly] 
**ReadPreference** | Pointer to [**ApiAtlasDataLakeAtlasStoreReadPreferenceView**](ApiAtlasDataLakeAtlasStoreReadPreferenceView.md) |  | [optional] 

## Methods

### NewApiAtlasDataLakeAtlasStoreView

`func NewApiAtlasDataLakeAtlasStoreView() *ApiAtlasDataLakeAtlasStoreView`

NewApiAtlasDataLakeAtlasStoreView instantiates a new ApiAtlasDataLakeAtlasStoreView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDataLakeAtlasStoreViewWithDefaults

`func NewApiAtlasDataLakeAtlasStoreViewWithDefaults() *ApiAtlasDataLakeAtlasStoreView`

NewApiAtlasDataLakeAtlasStoreViewWithDefaults instantiates a new ApiAtlasDataLakeAtlasStoreView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ApiAtlasDataLakeAtlasStoreView) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ApiAtlasDataLakeAtlasStoreView) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ApiAtlasDataLakeAtlasStoreView) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ApiAtlasDataLakeAtlasStoreView) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetProjectId

`func (o *ApiAtlasDataLakeAtlasStoreView) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ApiAtlasDataLakeAtlasStoreView) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ApiAtlasDataLakeAtlasStoreView) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ApiAtlasDataLakeAtlasStoreView) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReadPreference

`func (o *ApiAtlasDataLakeAtlasStoreView) GetReadPreference() ApiAtlasDataLakeAtlasStoreReadPreferenceView`

GetReadPreference returns the ReadPreference field if non-nil, zero value otherwise.

### GetReadPreferenceOk

`func (o *ApiAtlasDataLakeAtlasStoreView) GetReadPreferenceOk() (*ApiAtlasDataLakeAtlasStoreReadPreferenceView, bool)`

GetReadPreferenceOk returns a tuple with the ReadPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPreference

`func (o *ApiAtlasDataLakeAtlasStoreView) SetReadPreference(v ApiAtlasDataLakeAtlasStoreReadPreferenceView)`

SetReadPreference sets ReadPreference field to given value.

### HasReadPreference

`func (o *ApiAtlasDataLakeAtlasStoreView) HasReadPreference() bool`

HasReadPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


