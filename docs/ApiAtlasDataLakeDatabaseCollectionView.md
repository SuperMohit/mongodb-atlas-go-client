# ApiAtlasDataLakeDatabaseCollectionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSources** | Pointer to [**[]ApiAtlasDataLakeDatabaseDataSourceView**](ApiAtlasDataLakeDatabaseDataSourceView.md) | Array that contains the data stores that map to a collection for this data lake. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the collection to which MongoDB Cloud maps the data in the data stores. | [optional] [readonly] 

## Methods

### NewApiAtlasDataLakeDatabaseCollectionView

`func NewApiAtlasDataLakeDatabaseCollectionView() *ApiAtlasDataLakeDatabaseCollectionView`

NewApiAtlasDataLakeDatabaseCollectionView instantiates a new ApiAtlasDataLakeDatabaseCollectionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDataLakeDatabaseCollectionViewWithDefaults

`func NewApiAtlasDataLakeDatabaseCollectionViewWithDefaults() *ApiAtlasDataLakeDatabaseCollectionView`

NewApiAtlasDataLakeDatabaseCollectionViewWithDefaults instantiates a new ApiAtlasDataLakeDatabaseCollectionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSources

`func (o *ApiAtlasDataLakeDatabaseCollectionView) GetDataSources() []ApiAtlasDataLakeDatabaseDataSourceView`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *ApiAtlasDataLakeDatabaseCollectionView) GetDataSourcesOk() (*[]ApiAtlasDataLakeDatabaseDataSourceView, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *ApiAtlasDataLakeDatabaseCollectionView) SetDataSources(v []ApiAtlasDataLakeDatabaseDataSourceView)`

SetDataSources sets DataSources field to given value.

### HasDataSources

`func (o *ApiAtlasDataLakeDatabaseCollectionView) HasDataSources() bool`

HasDataSources returns a boolean if a field has been set.

### GetName

`func (o *ApiAtlasDataLakeDatabaseCollectionView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasDataLakeDatabaseCollectionView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasDataLakeDatabaseCollectionView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAtlasDataLakeDatabaseCollectionView) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


