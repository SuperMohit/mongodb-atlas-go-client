# ApiAtlasIngestionPipelineRunView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run. | [optional] [readonly] 
**BackupFrequencyType** | Pointer to **string** | Backup schedule interval of the Data Lake Pipeline. | [optional] [readonly] 
**CreatedDate** | Pointer to **time.Time** | Timestamp that indicates when the pipeline run was created. | [optional] [readonly] 
**DatasetName** | Pointer to **string** | Human-readable label that identifies the dataset that Atlas generates during this pipeline run. You can use this dataset as a &#x60;dataSource&#x60; in a Federated Database collection. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**LastUpdatedDate** | Pointer to **time.Time** | Timestamp that indicates the last time that the pipeline run was updated. | [optional] [readonly] 
**Phase** | Pointer to **string** | Processing phase of the Data Lake Pipeline. | [optional] [readonly] 
**PipelineId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies a Data Lake Pipeline. | [optional] [readonly] 
**SnapshotId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the snapshot of a cluster. | [optional] [readonly] 
**State** | Pointer to **string** | State of the pipeline run. | [optional] [readonly] 
**Stats** | Pointer to [**ApiAtlasPipelineRunStatsView**](ApiAtlasPipelineRunStatsView.md) |  | [optional] 

## Methods

### NewApiAtlasIngestionPipelineRunView

`func NewApiAtlasIngestionPipelineRunView() *ApiAtlasIngestionPipelineRunView`

NewApiAtlasIngestionPipelineRunView instantiates a new ApiAtlasIngestionPipelineRunView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasIngestionPipelineRunViewWithDefaults

`func NewApiAtlasIngestionPipelineRunViewWithDefaults() *ApiAtlasIngestionPipelineRunView`

NewApiAtlasIngestionPipelineRunViewWithDefaults instantiates a new ApiAtlasIngestionPipelineRunView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAtlasIngestionPipelineRunView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasIngestionPipelineRunView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasIngestionPipelineRunView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasIngestionPipelineRunView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBackupFrequencyType

`func (o *ApiAtlasIngestionPipelineRunView) GetBackupFrequencyType() string`

GetBackupFrequencyType returns the BackupFrequencyType field if non-nil, zero value otherwise.

### GetBackupFrequencyTypeOk

`func (o *ApiAtlasIngestionPipelineRunView) GetBackupFrequencyTypeOk() (*string, bool)`

GetBackupFrequencyTypeOk returns a tuple with the BackupFrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFrequencyType

`func (o *ApiAtlasIngestionPipelineRunView) SetBackupFrequencyType(v string)`

SetBackupFrequencyType sets BackupFrequencyType field to given value.

### HasBackupFrequencyType

`func (o *ApiAtlasIngestionPipelineRunView) HasBackupFrequencyType() bool`

HasBackupFrequencyType returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ApiAtlasIngestionPipelineRunView) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiAtlasIngestionPipelineRunView) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiAtlasIngestionPipelineRunView) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiAtlasIngestionPipelineRunView) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDatasetName

`func (o *ApiAtlasIngestionPipelineRunView) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *ApiAtlasIngestionPipelineRunView) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *ApiAtlasIngestionPipelineRunView) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *ApiAtlasIngestionPipelineRunView) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiAtlasIngestionPipelineRunView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasIngestionPipelineRunView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasIngestionPipelineRunView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAtlasIngestionPipelineRunView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *ApiAtlasIngestionPipelineRunView) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *ApiAtlasIngestionPipelineRunView) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *ApiAtlasIngestionPipelineRunView) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *ApiAtlasIngestionPipelineRunView) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetPhase

`func (o *ApiAtlasIngestionPipelineRunView) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ApiAtlasIngestionPipelineRunView) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ApiAtlasIngestionPipelineRunView) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ApiAtlasIngestionPipelineRunView) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPipelineId

`func (o *ApiAtlasIngestionPipelineRunView) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *ApiAtlasIngestionPipelineRunView) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *ApiAtlasIngestionPipelineRunView) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *ApiAtlasIngestionPipelineRunView) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetSnapshotId

`func (o *ApiAtlasIngestionPipelineRunView) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ApiAtlasIngestionPipelineRunView) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ApiAtlasIngestionPipelineRunView) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ApiAtlasIngestionPipelineRunView) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetState

`func (o *ApiAtlasIngestionPipelineRunView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiAtlasIngestionPipelineRunView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiAtlasIngestionPipelineRunView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiAtlasIngestionPipelineRunView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStats

`func (o *ApiAtlasIngestionPipelineRunView) GetStats() ApiAtlasPipelineRunStatsView`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ApiAtlasIngestionPipelineRunView) GetStatsOk() (*ApiAtlasPipelineRunStatsView, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ApiAtlasIngestionPipelineRunView) SetStats(v ApiAtlasPipelineRunStatsView)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ApiAtlasIngestionPipelineRunView) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


