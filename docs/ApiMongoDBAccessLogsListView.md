# ApiMongoDBAccessLogsListView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLogs** | Pointer to [**[]ApiMongoDBAccessLogsView**](ApiMongoDBAccessLogsView.md) | Authentication attempt, one per object, made against the cluster. | [optional] [readonly] 

## Methods

### NewApiMongoDBAccessLogsListView

`func NewApiMongoDBAccessLogsListView() *ApiMongoDBAccessLogsListView`

NewApiMongoDBAccessLogsListView instantiates a new ApiMongoDBAccessLogsListView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMongoDBAccessLogsListViewWithDefaults

`func NewApiMongoDBAccessLogsListViewWithDefaults() *ApiMongoDBAccessLogsListView`

NewApiMongoDBAccessLogsListViewWithDefaults instantiates a new ApiMongoDBAccessLogsListView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLogs

`func (o *ApiMongoDBAccessLogsListView) GetAccessLogs() []ApiMongoDBAccessLogsView`

GetAccessLogs returns the AccessLogs field if non-nil, zero value otherwise.

### GetAccessLogsOk

`func (o *ApiMongoDBAccessLogsListView) GetAccessLogsOk() (*[]ApiMongoDBAccessLogsView, bool)`

GetAccessLogsOk returns a tuple with the AccessLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLogs

`func (o *ApiMongoDBAccessLogsListView) SetAccessLogs(v []ApiMongoDBAccessLogsView)`

SetAccessLogs sets AccessLogs field to given value.

### HasAccessLogs

`func (o *ApiMongoDBAccessLogsListView) HasAccessLogs() bool`

HasAccessLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


