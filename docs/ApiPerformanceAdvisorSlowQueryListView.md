# ApiPerformanceAdvisorSlowQueryListView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SlowQueries** | Pointer to [**[]ApiPerformanceAdvisorSlowQueryView**](ApiPerformanceAdvisorSlowQueryView.md) | List of operations that the Performance Advisor detected that took longer to execute than a specified threshold. | [optional] [readonly] 

## Methods

### NewApiPerformanceAdvisorSlowQueryListView

`func NewApiPerformanceAdvisorSlowQueryListView() *ApiPerformanceAdvisorSlowQueryListView`

NewApiPerformanceAdvisorSlowQueryListView instantiates a new ApiPerformanceAdvisorSlowQueryListView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPerformanceAdvisorSlowQueryListViewWithDefaults

`func NewApiPerformanceAdvisorSlowQueryListViewWithDefaults() *ApiPerformanceAdvisorSlowQueryListView`

NewApiPerformanceAdvisorSlowQueryListViewWithDefaults instantiates a new ApiPerformanceAdvisorSlowQueryListView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlowQueries

`func (o *ApiPerformanceAdvisorSlowQueryListView) GetSlowQueries() []ApiPerformanceAdvisorSlowQueryView`

GetSlowQueries returns the SlowQueries field if non-nil, zero value otherwise.

### GetSlowQueriesOk

`func (o *ApiPerformanceAdvisorSlowQueryListView) GetSlowQueriesOk() (*[]ApiPerformanceAdvisorSlowQueryView, bool)`

GetSlowQueriesOk returns a tuple with the SlowQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowQueries

`func (o *ApiPerformanceAdvisorSlowQueryListView) SetSlowQueries(v []ApiPerformanceAdvisorSlowQueryView)`

SetSlowQueries sets SlowQueries field to given value.

### HasSlowQueries

`func (o *ApiPerformanceAdvisorSlowQueryListView) HasSlowQueries() bool`

HasSlowQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


