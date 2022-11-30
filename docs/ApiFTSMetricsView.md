# ApiFTSMetricsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the project. | [readonly] 
**HardwareMetrics** | [**[]ApiFTSMetricView**](ApiFTSMetricView.md) | List that contains all host compute, memory, and storage utilization dedicated to Atlas Search when MongoDB Atlas received this request. | [readonly] 
**IndexMetrics** | [**[]ApiFTSMetricView**](ApiFTSMetricView.md) | List that contains all performance and utilization measurements that Atlas Search index performed by the time MongoDB Atlas received this request. | [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**ProcessId** | **string** | Hostname and port that identifies the process. | [readonly] 
**StatusMetrics** | [**[]ApiFTSMetricView**](ApiFTSMetricView.md) | List that contains all available Atlas Search status metrics when MongoDB Atlas received this request. | [readonly] 

## Methods

### NewApiFTSMetricsView

`func NewApiFTSMetricsView(groupId string, hardwareMetrics []ApiFTSMetricView, indexMetrics []ApiFTSMetricView, links []Link, processId string, statusMetrics []ApiFTSMetricView, ) *ApiFTSMetricsView`

NewApiFTSMetricsView instantiates a new ApiFTSMetricsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFTSMetricsViewWithDefaults

`func NewApiFTSMetricsViewWithDefaults() *ApiFTSMetricsView`

NewApiFTSMetricsViewWithDefaults instantiates a new ApiFTSMetricsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ApiFTSMetricsView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiFTSMetricsView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiFTSMetricsView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetHardwareMetrics

`func (o *ApiFTSMetricsView) GetHardwareMetrics() []ApiFTSMetricView`

GetHardwareMetrics returns the HardwareMetrics field if non-nil, zero value otherwise.

### GetHardwareMetricsOk

`func (o *ApiFTSMetricsView) GetHardwareMetricsOk() (*[]ApiFTSMetricView, bool)`

GetHardwareMetricsOk returns a tuple with the HardwareMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareMetrics

`func (o *ApiFTSMetricsView) SetHardwareMetrics(v []ApiFTSMetricView)`

SetHardwareMetrics sets HardwareMetrics field to given value.


### GetIndexMetrics

`func (o *ApiFTSMetricsView) GetIndexMetrics() []ApiFTSMetricView`

GetIndexMetrics returns the IndexMetrics field if non-nil, zero value otherwise.

### GetIndexMetricsOk

`func (o *ApiFTSMetricsView) GetIndexMetricsOk() (*[]ApiFTSMetricView, bool)`

GetIndexMetricsOk returns a tuple with the IndexMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexMetrics

`func (o *ApiFTSMetricsView) SetIndexMetrics(v []ApiFTSMetricView)`

SetIndexMetrics sets IndexMetrics field to given value.


### GetLinks

`func (o *ApiFTSMetricsView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiFTSMetricsView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiFTSMetricsView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetProcessId

`func (o *ApiFTSMetricsView) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ApiFTSMetricsView) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ApiFTSMetricsView) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetStatusMetrics

`func (o *ApiFTSMetricsView) GetStatusMetrics() []ApiFTSMetricView`

GetStatusMetrics returns the StatusMetrics field if non-nil, zero value otherwise.

### GetStatusMetricsOk

`func (o *ApiFTSMetricsView) GetStatusMetricsOk() (*[]ApiFTSMetricView, bool)`

GetStatusMetricsOk returns a tuple with the StatusMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMetrics

`func (o *ApiFTSMetricsView) SetStatusMetrics(v []ApiFTSMetricView)`

SetStatusMetrics sets StatusMetrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


