# ApiExportStatusView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportedCollections** | Pointer to **int32** | Number of collections on the replica set that MongoDB Cloud exported. | [optional] [readonly] 
**TotalCollections** | Pointer to **int32** | Total number of collections on the replica set to export. | [optional] [readonly] 

## Methods

### NewApiExportStatusView

`func NewApiExportStatusView() *ApiExportStatusView`

NewApiExportStatusView instantiates a new ApiExportStatusView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExportStatusViewWithDefaults

`func NewApiExportStatusViewWithDefaults() *ApiExportStatusView`

NewApiExportStatusViewWithDefaults instantiates a new ApiExportStatusView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportedCollections

`func (o *ApiExportStatusView) GetExportedCollections() int32`

GetExportedCollections returns the ExportedCollections field if non-nil, zero value otherwise.

### GetExportedCollectionsOk

`func (o *ApiExportStatusView) GetExportedCollectionsOk() (*int32, bool)`

GetExportedCollectionsOk returns a tuple with the ExportedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedCollections

`func (o *ApiExportStatusView) SetExportedCollections(v int32)`

SetExportedCollections sets ExportedCollections field to given value.

### HasExportedCollections

`func (o *ApiExportStatusView) HasExportedCollections() bool`

HasExportedCollections returns a boolean if a field has been set.

### GetTotalCollections

`func (o *ApiExportStatusView) GetTotalCollections() int32`

GetTotalCollections returns the TotalCollections field if non-nil, zero value otherwise.

### GetTotalCollectionsOk

`func (o *ApiExportStatusView) GetTotalCollectionsOk() (*int32, bool)`

GetTotalCollectionsOk returns a tuple with the TotalCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCollections

`func (o *ApiExportStatusView) SetTotalCollections(v int32)`

SetTotalCollections sets TotalCollections field to given value.

### HasTotalCollections

`func (o *ApiExportStatusView) HasTotalCollections() bool`

HasTotalCollections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


