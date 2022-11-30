# ApiAtlasDataLakeDatabaseView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]ApiAtlasDataLakeDatabaseCollectionView**](ApiAtlasDataLakeDatabaseCollectionView.md) | Array of collections and data sources that map to a &#x60;&#x60;stores&#x60;&#x60; data store. | [optional] [readonly] 
**MaxWildcardCollections** | Pointer to **int32** | Maximum number of wildcard collections in the database. This only applies to S3 data sources. | [optional] [readonly] [default to 100]
**Name** | Pointer to **string** | Human-readable label that identifies the database to which the data lake maps data. | [optional] [readonly] 
**Views** | Pointer to [**[]ApiAtlasDataLakeViewView**](ApiAtlasDataLakeViewView.md) | Array of aggregation pipelines that apply to the collection. This only applies to S3 data sources. | [optional] [readonly] 

## Methods

### NewApiAtlasDataLakeDatabaseView

`func NewApiAtlasDataLakeDatabaseView() *ApiAtlasDataLakeDatabaseView`

NewApiAtlasDataLakeDatabaseView instantiates a new ApiAtlasDataLakeDatabaseView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDataLakeDatabaseViewWithDefaults

`func NewApiAtlasDataLakeDatabaseViewWithDefaults() *ApiAtlasDataLakeDatabaseView`

NewApiAtlasDataLakeDatabaseViewWithDefaults instantiates a new ApiAtlasDataLakeDatabaseView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *ApiAtlasDataLakeDatabaseView) GetCollections() []ApiAtlasDataLakeDatabaseCollectionView`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *ApiAtlasDataLakeDatabaseView) GetCollectionsOk() (*[]ApiAtlasDataLakeDatabaseCollectionView, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *ApiAtlasDataLakeDatabaseView) SetCollections(v []ApiAtlasDataLakeDatabaseCollectionView)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *ApiAtlasDataLakeDatabaseView) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetMaxWildcardCollections

`func (o *ApiAtlasDataLakeDatabaseView) GetMaxWildcardCollections() int32`

GetMaxWildcardCollections returns the MaxWildcardCollections field if non-nil, zero value otherwise.

### GetMaxWildcardCollectionsOk

`func (o *ApiAtlasDataLakeDatabaseView) GetMaxWildcardCollectionsOk() (*int32, bool)`

GetMaxWildcardCollectionsOk returns a tuple with the MaxWildcardCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWildcardCollections

`func (o *ApiAtlasDataLakeDatabaseView) SetMaxWildcardCollections(v int32)`

SetMaxWildcardCollections sets MaxWildcardCollections field to given value.

### HasMaxWildcardCollections

`func (o *ApiAtlasDataLakeDatabaseView) HasMaxWildcardCollections() bool`

HasMaxWildcardCollections returns a boolean if a field has been set.

### GetName

`func (o *ApiAtlasDataLakeDatabaseView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasDataLakeDatabaseView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasDataLakeDatabaseView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAtlasDataLakeDatabaseView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetViews

`func (o *ApiAtlasDataLakeDatabaseView) GetViews() []ApiAtlasDataLakeViewView`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *ApiAtlasDataLakeDatabaseView) GetViewsOk() (*[]ApiAtlasDataLakeViewView, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *ApiAtlasDataLakeDatabaseView) SetViews(v []ApiAtlasDataLakeViewView)`

SetViews sets Views field to given value.

### HasViews

`func (o *ApiAtlasDataLakeDatabaseView) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


