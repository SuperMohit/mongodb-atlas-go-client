# ApiAlertConfigView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created the alert configuration. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Flag that indicates whether someone enabled this alert configuration for the specified project. | [optional] [default to false]
**EventTypeName** | **string** | Event type that triggers an alert. | 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns this alert configuration. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this alert configuration. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Matchers** | Pointer to [**[]ApiMatcherView**](ApiMatcherView.md) | List of rules that determine whether MongoDB Cloud checks an object for the alert configuration. You can filter using the matchers array if the **eventTypeName** specifies an event for a host, replica set, or sharded cluster. | [optional] 
**MetricThreshold** | Pointer to [**ApiMetricThresholdView**](ApiMetricThresholdView.md) |  | [optional] 
**Notifications** | Pointer to [**[]ApiNotificationView**](ApiNotificationView.md) | List that contains the targets that MongoDB Cloud sends notifications. | [optional] 
**Threshold** | Pointer to [**ApiIntegerThresholdView**](ApiIntegerThresholdView.md) |  | [optional] 
**TypeName** | Pointer to **string** | Human-readable label that displays the alert type. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | Date and time when someone last updated this alert configuration. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewApiAlertConfigView

`func NewApiAlertConfigView(eventTypeName string, links []Link, ) *ApiAlertConfigView`

NewApiAlertConfigView instantiates a new ApiAlertConfigView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAlertConfigViewWithDefaults

`func NewApiAlertConfigViewWithDefaults() *ApiAlertConfigView`

NewApiAlertConfigViewWithDefaults instantiates a new ApiAlertConfigView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ApiAlertConfigView) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiAlertConfigView) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiAlertConfigView) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiAlertConfigView) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiAlertConfigView) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiAlertConfigView) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiAlertConfigView) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiAlertConfigView) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypeName

`func (o *ApiAlertConfigView) GetEventTypeName() string`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *ApiAlertConfigView) GetEventTypeNameOk() (*string, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *ApiAlertConfigView) SetEventTypeName(v string)`

SetEventTypeName sets EventTypeName field to given value.


### GetGroupId

`func (o *ApiAlertConfigView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAlertConfigView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAlertConfigView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAlertConfigView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ApiAlertConfigView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAlertConfigView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAlertConfigView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAlertConfigView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAlertConfigView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAlertConfigView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAlertConfigView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetMatchers

`func (o *ApiAlertConfigView) GetMatchers() []ApiMatcherView`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *ApiAlertConfigView) GetMatchersOk() (*[]ApiMatcherView, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *ApiAlertConfigView) SetMatchers(v []ApiMatcherView)`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *ApiAlertConfigView) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.

### GetMetricThreshold

`func (o *ApiAlertConfigView) GetMetricThreshold() ApiMetricThresholdView`

GetMetricThreshold returns the MetricThreshold field if non-nil, zero value otherwise.

### GetMetricThresholdOk

`func (o *ApiAlertConfigView) GetMetricThresholdOk() (*ApiMetricThresholdView, bool)`

GetMetricThresholdOk returns a tuple with the MetricThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricThreshold

`func (o *ApiAlertConfigView) SetMetricThreshold(v ApiMetricThresholdView)`

SetMetricThreshold sets MetricThreshold field to given value.

### HasMetricThreshold

`func (o *ApiAlertConfigView) HasMetricThreshold() bool`

HasMetricThreshold returns a boolean if a field has been set.

### GetNotifications

`func (o *ApiAlertConfigView) GetNotifications() []ApiNotificationView`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ApiAlertConfigView) GetNotificationsOk() (*[]ApiNotificationView, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ApiAlertConfigView) SetNotifications(v []ApiNotificationView)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ApiAlertConfigView) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetThreshold

`func (o *ApiAlertConfigView) GetThreshold() ApiIntegerThresholdView`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ApiAlertConfigView) GetThresholdOk() (*ApiIntegerThresholdView, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ApiAlertConfigView) SetThreshold(v ApiIntegerThresholdView)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ApiAlertConfigView) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTypeName

`func (o *ApiAlertConfigView) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *ApiAlertConfigView) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *ApiAlertConfigView) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *ApiAlertConfigView) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetUpdated

`func (o *ApiAlertConfigView) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ApiAlertConfigView) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ApiAlertConfigView) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ApiAlertConfigView) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


