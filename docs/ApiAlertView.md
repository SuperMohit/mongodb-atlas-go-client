# ApiAlertView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedUntil** | **time.Time** | Date and time until which this alert has been acknowledged. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert.  - To acknowledge this alert forever, set the parameter value to 100 years in the future.  - To unacknowledge a previously acknowledged alert, set the parameter value to a date in the past. | 
**AcknowledgementComment** | Pointer to **string** | Comment that a MongoDB Cloud user submitted when acknowledging the alert. | [optional] 
**AcknowledgingUsername** | Pointer to **string** | MongoDB Cloud username of the person who acknowledged the alert. The response returns this parameter if a MongoDB Cloud user previously acknowledged this alert. | [optional] [readonly] 
**AlertConfigId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster to which this alert applies. This resource returns this parameter for alerts of events impacting backups, replica sets, or sharded clusters. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this alert. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**CurrentValue** | Pointer to [**ApiMetricValueView**](ApiMetricValueView.md) |  | [optional] 
**EventTypeName** | Pointer to **string** | Incident that triggered this alert. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns this alert. | [optional] [readonly] 
**HostnameAndPort** | Pointer to **string** | Hostname and port of the host to which this alert applies. The resource returns this parameter for alerts of events impacting hosts or replica sets. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this alert. | [optional] [readonly] 
**LastNotified** | Pointer to **time.Time** | Date and time that any notifications were last sent for this alert. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter if MongoDB Cloud has sent notifications for this alert. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**MetricName** | Pointer to **string** | Human-readable label that identifies the metric against which MongoDB Cloud checks the alert. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Name of the replica set to which this alert applies. The response returns this parameter for alerts of events impacting backups, hosts, or replica sets. | [optional] [readonly] 
**Resolved** | Pointer to **time.Time** | Date and time that this alert changed to &#x60;\&quot;status\&quot; : \&quot;CLOSED\&quot;&#x60;. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter once &#x60;\&quot;status\&quot; : \&quot;CLOSED\&quot;&#x60;. | [optional] [readonly] 
**Status** | Pointer to **string** | State of this alert at the time you requested its details. | [optional] [readonly] 
**TypeName** | Pointer to **string** | Category in which MongoDB Cloud classifies this alert. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | Date and time when someone last updated this alert. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewApiAlertView

`func NewApiAlertView(acknowledgedUntil time.Time, links []Link, ) *ApiAlertView`

NewApiAlertView instantiates a new ApiAlertView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAlertViewWithDefaults

`func NewApiAlertViewWithDefaults() *ApiAlertView`

NewApiAlertViewWithDefaults instantiates a new ApiAlertView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgedUntil

`func (o *ApiAlertView) GetAcknowledgedUntil() time.Time`

GetAcknowledgedUntil returns the AcknowledgedUntil field if non-nil, zero value otherwise.

### GetAcknowledgedUntilOk

`func (o *ApiAlertView) GetAcknowledgedUntilOk() (*time.Time, bool)`

GetAcknowledgedUntilOk returns a tuple with the AcknowledgedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedUntil

`func (o *ApiAlertView) SetAcknowledgedUntil(v time.Time)`

SetAcknowledgedUntil sets AcknowledgedUntil field to given value.


### GetAcknowledgementComment

`func (o *ApiAlertView) GetAcknowledgementComment() string`

GetAcknowledgementComment returns the AcknowledgementComment field if non-nil, zero value otherwise.

### GetAcknowledgementCommentOk

`func (o *ApiAlertView) GetAcknowledgementCommentOk() (*string, bool)`

GetAcknowledgementCommentOk returns a tuple with the AcknowledgementComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgementComment

`func (o *ApiAlertView) SetAcknowledgementComment(v string)`

SetAcknowledgementComment sets AcknowledgementComment field to given value.

### HasAcknowledgementComment

`func (o *ApiAlertView) HasAcknowledgementComment() bool`

HasAcknowledgementComment returns a boolean if a field has been set.

### GetAcknowledgingUsername

`func (o *ApiAlertView) GetAcknowledgingUsername() string`

GetAcknowledgingUsername returns the AcknowledgingUsername field if non-nil, zero value otherwise.

### GetAcknowledgingUsernameOk

`func (o *ApiAlertView) GetAcknowledgingUsernameOk() (*string, bool)`

GetAcknowledgingUsernameOk returns a tuple with the AcknowledgingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgingUsername

`func (o *ApiAlertView) SetAcknowledgingUsername(v string)`

SetAcknowledgingUsername sets AcknowledgingUsername field to given value.

### HasAcknowledgingUsername

`func (o *ApiAlertView) HasAcknowledgingUsername() bool`

HasAcknowledgingUsername returns a boolean if a field has been set.

### GetAlertConfigId

`func (o *ApiAlertView) GetAlertConfigId() string`

GetAlertConfigId returns the AlertConfigId field if non-nil, zero value otherwise.

### GetAlertConfigIdOk

`func (o *ApiAlertView) GetAlertConfigIdOk() (*string, bool)`

GetAlertConfigIdOk returns a tuple with the AlertConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigId

`func (o *ApiAlertView) SetAlertConfigId(v string)`

SetAlertConfigId sets AlertConfigId field to given value.

### HasAlertConfigId

`func (o *ApiAlertView) HasAlertConfigId() bool`

HasAlertConfigId returns a boolean if a field has been set.

### GetClusterName

`func (o *ApiAlertView) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ApiAlertView) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ApiAlertView) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ApiAlertView) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCreated

`func (o *ApiAlertView) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiAlertView) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiAlertView) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiAlertView) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCurrentValue

`func (o *ApiAlertView) GetCurrentValue() ApiMetricValueView`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *ApiAlertView) GetCurrentValueOk() (*ApiMetricValueView, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *ApiAlertView) SetCurrentValue(v ApiMetricValueView)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *ApiAlertView) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetEventTypeName

`func (o *ApiAlertView) GetEventTypeName() string`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *ApiAlertView) GetEventTypeNameOk() (*string, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *ApiAlertView) SetEventTypeName(v string)`

SetEventTypeName sets EventTypeName field to given value.

### HasEventTypeName

`func (o *ApiAlertView) HasEventTypeName() bool`

HasEventTypeName returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiAlertView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAlertView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAlertView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAlertView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHostnameAndPort

`func (o *ApiAlertView) GetHostnameAndPort() string`

GetHostnameAndPort returns the HostnameAndPort field if non-nil, zero value otherwise.

### GetHostnameAndPortOk

`func (o *ApiAlertView) GetHostnameAndPortOk() (*string, bool)`

GetHostnameAndPortOk returns a tuple with the HostnameAndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameAndPort

`func (o *ApiAlertView) SetHostnameAndPort(v string)`

SetHostnameAndPort sets HostnameAndPort field to given value.

### HasHostnameAndPort

`func (o *ApiAlertView) HasHostnameAndPort() bool`

HasHostnameAndPort returns a boolean if a field has been set.

### GetId

`func (o *ApiAlertView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAlertView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAlertView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAlertView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastNotified

`func (o *ApiAlertView) GetLastNotified() time.Time`

GetLastNotified returns the LastNotified field if non-nil, zero value otherwise.

### GetLastNotifiedOk

`func (o *ApiAlertView) GetLastNotifiedOk() (*time.Time, bool)`

GetLastNotifiedOk returns a tuple with the LastNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotified

`func (o *ApiAlertView) SetLastNotified(v time.Time)`

SetLastNotified sets LastNotified field to given value.

### HasLastNotified

`func (o *ApiAlertView) HasLastNotified() bool`

HasLastNotified returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAlertView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAlertView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAlertView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetMetricName

`func (o *ApiAlertView) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ApiAlertView) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ApiAlertView) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *ApiAlertView) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetReplicaSetName

`func (o *ApiAlertView) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *ApiAlertView) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *ApiAlertView) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *ApiAlertView) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.

### GetResolved

`func (o *ApiAlertView) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *ApiAlertView) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *ApiAlertView) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *ApiAlertView) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetStatus

`func (o *ApiAlertView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAlertView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAlertView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAlertView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTypeName

`func (o *ApiAlertView) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *ApiAlertView) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *ApiAlertView) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *ApiAlertView) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetUpdated

`func (o *ApiAlertView) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ApiAlertView) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ApiAlertView) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ApiAlertView) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


