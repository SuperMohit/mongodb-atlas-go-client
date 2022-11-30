# ApiAtlasIngestionPipelineView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Data Lake Pipeline. | [optional] [readonly] 
**CreatedDate** | Pointer to **time.Time** | Timestamp that indicates when the Data Lake Pipeline was created. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the group. | [optional] [readonly] 
**LastUpdatedDate** | Pointer to **time.Time** | Timestamp that indicates the last time that the Data Lake Pipeline was updated. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of this Data Lake Pipeline. | [optional] 
**Sink** | Pointer to [**ApiAtlasIngestionSinkView**](ApiAtlasIngestionSinkView.md) |  | [optional] 
**Source** | Pointer to [**ApiAtlasIngestionSourceView**](ApiAtlasIngestionSourceView.md) |  | [optional] 
**State** | Pointer to **string** | State of this Data Lake Pipeline. | [optional] [readonly] 
**Transformations** | Pointer to [**[]ApiAtlasFieldTransformationView**](ApiAtlasFieldTransformationView.md) | Fields to be excluded for this Data Lake Pipeline. | [optional] 

## Methods

### NewApiAtlasIngestionPipelineView

`func NewApiAtlasIngestionPipelineView() *ApiAtlasIngestionPipelineView`

NewApiAtlasIngestionPipelineView instantiates a new ApiAtlasIngestionPipelineView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasIngestionPipelineViewWithDefaults

`func NewApiAtlasIngestionPipelineViewWithDefaults() *ApiAtlasIngestionPipelineView`

NewApiAtlasIngestionPipelineViewWithDefaults instantiates a new ApiAtlasIngestionPipelineView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAtlasIngestionPipelineView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasIngestionPipelineView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasIngestionPipelineView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasIngestionPipelineView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ApiAtlasIngestionPipelineView) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiAtlasIngestionPipelineView) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiAtlasIngestionPipelineView) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiAtlasIngestionPipelineView) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiAtlasIngestionPipelineView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasIngestionPipelineView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasIngestionPipelineView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAtlasIngestionPipelineView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *ApiAtlasIngestionPipelineView) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *ApiAtlasIngestionPipelineView) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *ApiAtlasIngestionPipelineView) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *ApiAtlasIngestionPipelineView) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetName

`func (o *ApiAtlasIngestionPipelineView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasIngestionPipelineView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasIngestionPipelineView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAtlasIngestionPipelineView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSink

`func (o *ApiAtlasIngestionPipelineView) GetSink() ApiAtlasIngestionSinkView`

GetSink returns the Sink field if non-nil, zero value otherwise.

### GetSinkOk

`func (o *ApiAtlasIngestionPipelineView) GetSinkOk() (*ApiAtlasIngestionSinkView, bool)`

GetSinkOk returns a tuple with the Sink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSink

`func (o *ApiAtlasIngestionPipelineView) SetSink(v ApiAtlasIngestionSinkView)`

SetSink sets Sink field to given value.

### HasSink

`func (o *ApiAtlasIngestionPipelineView) HasSink() bool`

HasSink returns a boolean if a field has been set.

### GetSource

`func (o *ApiAtlasIngestionPipelineView) GetSource() ApiAtlasIngestionSourceView`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiAtlasIngestionPipelineView) GetSourceOk() (*ApiAtlasIngestionSourceView, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiAtlasIngestionPipelineView) SetSource(v ApiAtlasIngestionSourceView)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiAtlasIngestionPipelineView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *ApiAtlasIngestionPipelineView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiAtlasIngestionPipelineView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiAtlasIngestionPipelineView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiAtlasIngestionPipelineView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTransformations

`func (o *ApiAtlasIngestionPipelineView) GetTransformations() []ApiAtlasFieldTransformationView`

GetTransformations returns the Transformations field if non-nil, zero value otherwise.

### GetTransformationsOk

`func (o *ApiAtlasIngestionPipelineView) GetTransformationsOk() (*[]ApiAtlasFieldTransformationView, bool)`

GetTransformationsOk returns a tuple with the Transformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformations

`func (o *ApiAtlasIngestionPipelineView) SetTransformations(v []ApiAtlasFieldTransformationView)`

SetTransformations sets Transformations field to given value.

### HasTransformations

`func (o *ApiAtlasIngestionPipelineView) HasTransformations() bool`

HasTransformations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


