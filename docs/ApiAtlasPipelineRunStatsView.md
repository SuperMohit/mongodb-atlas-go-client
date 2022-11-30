# ApiAtlasPipelineRunStatsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesExported** | Pointer to **int64** | Total data size in bytes exported for this pipeline run. | [optional] [readonly] 
**NumDocs** | Pointer to **int64** | Number of docs ingested for a this pipeline run. | [optional] [readonly] 

## Methods

### NewApiAtlasPipelineRunStatsView

`func NewApiAtlasPipelineRunStatsView() *ApiAtlasPipelineRunStatsView`

NewApiAtlasPipelineRunStatsView instantiates a new ApiAtlasPipelineRunStatsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasPipelineRunStatsViewWithDefaults

`func NewApiAtlasPipelineRunStatsViewWithDefaults() *ApiAtlasPipelineRunStatsView`

NewApiAtlasPipelineRunStatsViewWithDefaults instantiates a new ApiAtlasPipelineRunStatsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesExported

`func (o *ApiAtlasPipelineRunStatsView) GetBytesExported() int64`

GetBytesExported returns the BytesExported field if non-nil, zero value otherwise.

### GetBytesExportedOk

`func (o *ApiAtlasPipelineRunStatsView) GetBytesExportedOk() (*int64, bool)`

GetBytesExportedOk returns a tuple with the BytesExported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesExported

`func (o *ApiAtlasPipelineRunStatsView) SetBytesExported(v int64)`

SetBytesExported sets BytesExported field to given value.

### HasBytesExported

`func (o *ApiAtlasPipelineRunStatsView) HasBytesExported() bool`

HasBytesExported returns a boolean if a field has been set.

### GetNumDocs

`func (o *ApiAtlasPipelineRunStatsView) GetNumDocs() int64`

GetNumDocs returns the NumDocs field if non-nil, zero value otherwise.

### GetNumDocsOk

`func (o *ApiAtlasPipelineRunStatsView) GetNumDocsOk() (*int64, bool)`

GetNumDocsOk returns a tuple with the NumDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDocs

`func (o *ApiAtlasPipelineRunStatsView) SetNumDocs(v int64)`

SetNumDocs sets NumDocs field to given value.

### HasNumDocs

`func (o *ApiAtlasPipelineRunStatsView) HasNumDocs() bool`

HasNumDocs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


