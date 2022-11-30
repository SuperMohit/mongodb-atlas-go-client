# ApiAtlasSnapshotScheduleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterCheckpointIntervalMin** | **int32** | Quantity of time expressed in minutes between successive cluster checkpoints. This parameter applies only to sharded clusters. This number determines the granularity of continuous cloud backups for sharded clusters. | 
**ClusterId** | **string** | Unique 24-hexadecimal digit string that identifies the cluster with the snapshot you want to return. | 
**DailySnapshotRetentionDays** | **int32** | Quantity of time to keep daily snapshots. MongoDB Cloud expresses this value in days. Set this value to &#x60;0&#x60; to disable daily snapshot retention. | 
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the project that contains the cluster. | [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**MonthlySnapshotRetentionMonths** | **int32** | Number of months that MongoDB Cloud must keep monthly snapshots. Set this value to &#x60;0&#x60; to disable monthly snapshot retention. | 
**PointInTimeWindowHours** | **int32** | Number of hours before the current time from which MongoDB Cloud can create a Continuous Cloud Backup snapshot. | 
**SnapshotIntervalHours** | **int32** | Number of hours that must elapse before taking another snapshot. | 
**SnapshotRetentionDays** | **int32** | Number of days that MongoDB Cloud must keep recent snapshots. | 
**WeeklySnapshotRetentionWeeks** | **int32** | Number of weeks that MongoDB Cloud must keep weekly snapshots. Set this value to &#x60;0&#x60; to disable weekly snapshot retention. | 

## Methods

### NewApiAtlasSnapshotScheduleView

`func NewApiAtlasSnapshotScheduleView(clusterCheckpointIntervalMin int32, clusterId string, dailySnapshotRetentionDays int32, groupId string, links []Link, monthlySnapshotRetentionMonths int32, pointInTimeWindowHours int32, snapshotIntervalHours int32, snapshotRetentionDays int32, weeklySnapshotRetentionWeeks int32, ) *ApiAtlasSnapshotScheduleView`

NewApiAtlasSnapshotScheduleView instantiates a new ApiAtlasSnapshotScheduleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasSnapshotScheduleViewWithDefaults

`func NewApiAtlasSnapshotScheduleViewWithDefaults() *ApiAtlasSnapshotScheduleView`

NewApiAtlasSnapshotScheduleViewWithDefaults instantiates a new ApiAtlasSnapshotScheduleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterCheckpointIntervalMin

`func (o *ApiAtlasSnapshotScheduleView) GetClusterCheckpointIntervalMin() int32`

GetClusterCheckpointIntervalMin returns the ClusterCheckpointIntervalMin field if non-nil, zero value otherwise.

### GetClusterCheckpointIntervalMinOk

`func (o *ApiAtlasSnapshotScheduleView) GetClusterCheckpointIntervalMinOk() (*int32, bool)`

GetClusterCheckpointIntervalMinOk returns a tuple with the ClusterCheckpointIntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCheckpointIntervalMin

`func (o *ApiAtlasSnapshotScheduleView) SetClusterCheckpointIntervalMin(v int32)`

SetClusterCheckpointIntervalMin sets ClusterCheckpointIntervalMin field to given value.


### GetClusterId

`func (o *ApiAtlasSnapshotScheduleView) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ApiAtlasSnapshotScheduleView) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ApiAtlasSnapshotScheduleView) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetDailySnapshotRetentionDays

`func (o *ApiAtlasSnapshotScheduleView) GetDailySnapshotRetentionDays() int32`

GetDailySnapshotRetentionDays returns the DailySnapshotRetentionDays field if non-nil, zero value otherwise.

### GetDailySnapshotRetentionDaysOk

`func (o *ApiAtlasSnapshotScheduleView) GetDailySnapshotRetentionDaysOk() (*int32, bool)`

GetDailySnapshotRetentionDaysOk returns a tuple with the DailySnapshotRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySnapshotRetentionDays

`func (o *ApiAtlasSnapshotScheduleView) SetDailySnapshotRetentionDays(v int32)`

SetDailySnapshotRetentionDays sets DailySnapshotRetentionDays field to given value.


### GetGroupId

`func (o *ApiAtlasSnapshotScheduleView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasSnapshotScheduleView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasSnapshotScheduleView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetLinks

`func (o *ApiAtlasSnapshotScheduleView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasSnapshotScheduleView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasSnapshotScheduleView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetMonthlySnapshotRetentionMonths

`func (o *ApiAtlasSnapshotScheduleView) GetMonthlySnapshotRetentionMonths() int32`

GetMonthlySnapshotRetentionMonths returns the MonthlySnapshotRetentionMonths field if non-nil, zero value otherwise.

### GetMonthlySnapshotRetentionMonthsOk

`func (o *ApiAtlasSnapshotScheduleView) GetMonthlySnapshotRetentionMonthsOk() (*int32, bool)`

GetMonthlySnapshotRetentionMonthsOk returns a tuple with the MonthlySnapshotRetentionMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlySnapshotRetentionMonths

`func (o *ApiAtlasSnapshotScheduleView) SetMonthlySnapshotRetentionMonths(v int32)`

SetMonthlySnapshotRetentionMonths sets MonthlySnapshotRetentionMonths field to given value.


### GetPointInTimeWindowHours

`func (o *ApiAtlasSnapshotScheduleView) GetPointInTimeWindowHours() int32`

GetPointInTimeWindowHours returns the PointInTimeWindowHours field if non-nil, zero value otherwise.

### GetPointInTimeWindowHoursOk

`func (o *ApiAtlasSnapshotScheduleView) GetPointInTimeWindowHoursOk() (*int32, bool)`

GetPointInTimeWindowHoursOk returns a tuple with the PointInTimeWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeWindowHours

`func (o *ApiAtlasSnapshotScheduleView) SetPointInTimeWindowHours(v int32)`

SetPointInTimeWindowHours sets PointInTimeWindowHours field to given value.


### GetSnapshotIntervalHours

`func (o *ApiAtlasSnapshotScheduleView) GetSnapshotIntervalHours() int32`

GetSnapshotIntervalHours returns the SnapshotIntervalHours field if non-nil, zero value otherwise.

### GetSnapshotIntervalHoursOk

`func (o *ApiAtlasSnapshotScheduleView) GetSnapshotIntervalHoursOk() (*int32, bool)`

GetSnapshotIntervalHoursOk returns a tuple with the SnapshotIntervalHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIntervalHours

`func (o *ApiAtlasSnapshotScheduleView) SetSnapshotIntervalHours(v int32)`

SetSnapshotIntervalHours sets SnapshotIntervalHours field to given value.


### GetSnapshotRetentionDays

`func (o *ApiAtlasSnapshotScheduleView) GetSnapshotRetentionDays() int32`

GetSnapshotRetentionDays returns the SnapshotRetentionDays field if non-nil, zero value otherwise.

### GetSnapshotRetentionDaysOk

`func (o *ApiAtlasSnapshotScheduleView) GetSnapshotRetentionDaysOk() (*int32, bool)`

GetSnapshotRetentionDaysOk returns a tuple with the SnapshotRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionDays

`func (o *ApiAtlasSnapshotScheduleView) SetSnapshotRetentionDays(v int32)`

SetSnapshotRetentionDays sets SnapshotRetentionDays field to given value.


### GetWeeklySnapshotRetentionWeeks

`func (o *ApiAtlasSnapshotScheduleView) GetWeeklySnapshotRetentionWeeks() int32`

GetWeeklySnapshotRetentionWeeks returns the WeeklySnapshotRetentionWeeks field if non-nil, zero value otherwise.

### GetWeeklySnapshotRetentionWeeksOk

`func (o *ApiAtlasSnapshotScheduleView) GetWeeklySnapshotRetentionWeeksOk() (*int32, bool)`

GetWeeklySnapshotRetentionWeeksOk returns a tuple with the WeeklySnapshotRetentionWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklySnapshotRetentionWeeks

`func (o *ApiAtlasSnapshotScheduleView) SetWeeklySnapshotRetentionWeeks(v int32)`

SetWeeklySnapshotRetentionWeeks sets WeeklySnapshotRetentionWeeks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


