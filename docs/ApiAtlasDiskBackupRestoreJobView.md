# ApiAtlasDiskBackupRestoreJobView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | Pointer to **bool** | Flag that indicates whether someone canceled this restore job. | [optional] [readonly] 
**Components** | Pointer to [**[]ApiAtlasDiskBackupBaseRestoreMemberView**](ApiAtlasDiskBackupBaseRestoreMemberView.md) | Information on the restore job for each replica set in the sharded cluster. | [optional] [readonly] 
**DeliveryType** | **string** | Human-readable label that categorizes the restore job to create. | 
**DeliveryUrl** | Pointer to **[]string** | One or more Uniform Resource Locators (URLs) that point to the compressed snapshot files for manual download. MongoDB Cloud returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;download\&quot;&#x60;. | [optional] [readonly] 
**DesiredTimestamp** | Pointer to [**ApiBSONTimestampView**](ApiBSONTimestampView.md) |  | [optional] 
**Expired** | Pointer to **bool** | Flag that indicates whether the restore job expired. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Date and time when the restore job expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Failed** | Pointer to **bool** | Flag that indicates whether the restore job failed. | [optional] [readonly] 
**FinishedAt** | Pointer to **time.Time** | Date and time when the restore job completed. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the restore job. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**OplogInc** | Pointer to **int32** | Oplog operation number from which you want to restore this snapshot. This number represents the second part of an Oplog timestamp. The resource returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;pointInTime\&quot;&#x60; and **oplogTs** exceeds &#x60;0&#x60;. | [optional] 
**OplogTs** | Pointer to **int32** | Date and time from which you want to restore this snapshot. This parameter expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. This number represents the first part of an Oplog timestamp. The resource returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;pointInTime\&quot;&#x60; and **oplogTs** exceeds &#x60;0&#x60;. | [optional] 
**PointInTimeUTCSeconds** | Pointer to **int32** | Date and time from which MongoDB Cloud restored this snapshot. This parameter expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch. The resource returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;pointInTime\&quot;&#x60; and **pointInTimeUTCSeconds** exceeds &#x60;0&#x60;. | [optional] 
**SnapshotId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the snapshot. | [optional] 
**TargetClusterName** | **string** | Human-readable label that identifies the target cluster to which the restore job restores the snapshot. The resource returns this parameter when &#x60;\&quot;deliveryType\&quot;:&#x60; &#x60;\&quot;automated\&quot;&#x60;. | 
**TargetGroupId** | **string** | Unique 24-hexadecimal digit string that identifies the target project for the specified **targetClusterName**. | 
**Timestamp** | Pointer to **time.Time** | Date and time when MongoDB Cloud took the snapshot associated with **snapshotId**. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewApiAtlasDiskBackupRestoreJobView

`func NewApiAtlasDiskBackupRestoreJobView(deliveryType string, links []Link, targetClusterName string, targetGroupId string, ) *ApiAtlasDiskBackupRestoreJobView`

NewApiAtlasDiskBackupRestoreJobView instantiates a new ApiAtlasDiskBackupRestoreJobView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDiskBackupRestoreJobViewWithDefaults

`func NewApiAtlasDiskBackupRestoreJobViewWithDefaults() *ApiAtlasDiskBackupRestoreJobView`

NewApiAtlasDiskBackupRestoreJobViewWithDefaults instantiates a new ApiAtlasDiskBackupRestoreJobView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *ApiAtlasDiskBackupRestoreJobView) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *ApiAtlasDiskBackupRestoreJobView) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.

### HasCancelled

`func (o *ApiAtlasDiskBackupRestoreJobView) HasCancelled() bool`

HasCancelled returns a boolean if a field has been set.

### GetComponents

`func (o *ApiAtlasDiskBackupRestoreJobView) GetComponents() []ApiAtlasDiskBackupBaseRestoreMemberView`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetComponentsOk() (*[]ApiAtlasDiskBackupBaseRestoreMemberView, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ApiAtlasDiskBackupRestoreJobView) SetComponents(v []ApiAtlasDiskBackupBaseRestoreMemberView)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ApiAtlasDiskBackupRestoreJobView) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetDeliveryType

`func (o *ApiAtlasDiskBackupRestoreJobView) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *ApiAtlasDiskBackupRestoreJobView) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.


### GetDeliveryUrl

`func (o *ApiAtlasDiskBackupRestoreJobView) GetDeliveryUrl() []string`

GetDeliveryUrl returns the DeliveryUrl field if non-nil, zero value otherwise.

### GetDeliveryUrlOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetDeliveryUrlOk() (*[]string, bool)`

GetDeliveryUrlOk returns a tuple with the DeliveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryUrl

`func (o *ApiAtlasDiskBackupRestoreJobView) SetDeliveryUrl(v []string)`

SetDeliveryUrl sets DeliveryUrl field to given value.

### HasDeliveryUrl

`func (o *ApiAtlasDiskBackupRestoreJobView) HasDeliveryUrl() bool`

HasDeliveryUrl returns a boolean if a field has been set.

### GetDesiredTimestamp

`func (o *ApiAtlasDiskBackupRestoreJobView) GetDesiredTimestamp() ApiBSONTimestampView`

GetDesiredTimestamp returns the DesiredTimestamp field if non-nil, zero value otherwise.

### GetDesiredTimestampOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetDesiredTimestampOk() (*ApiBSONTimestampView, bool)`

GetDesiredTimestampOk returns a tuple with the DesiredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTimestamp

`func (o *ApiAtlasDiskBackupRestoreJobView) SetDesiredTimestamp(v ApiBSONTimestampView)`

SetDesiredTimestamp sets DesiredTimestamp field to given value.

### HasDesiredTimestamp

`func (o *ApiAtlasDiskBackupRestoreJobView) HasDesiredTimestamp() bool`

HasDesiredTimestamp returns a boolean if a field has been set.

### GetExpired

`func (o *ApiAtlasDiskBackupRestoreJobView) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ApiAtlasDiskBackupRestoreJobView) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *ApiAtlasDiskBackupRestoreJobView) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiAtlasDiskBackupRestoreJobView) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiAtlasDiskBackupRestoreJobView) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiAtlasDiskBackupRestoreJobView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFailed

`func (o *ApiAtlasDiskBackupRestoreJobView) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ApiAtlasDiskBackupRestoreJobView) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ApiAtlasDiskBackupRestoreJobView) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetFinishedAt

`func (o *ApiAtlasDiskBackupRestoreJobView) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ApiAtlasDiskBackupRestoreJobView) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ApiAtlasDiskBackupRestoreJobView) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasDiskBackupRestoreJobView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasDiskBackupRestoreJobView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasDiskBackupRestoreJobView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasDiskBackupRestoreJobView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasDiskBackupRestoreJobView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetOplogInc

`func (o *ApiAtlasDiskBackupRestoreJobView) GetOplogInc() int32`

GetOplogInc returns the OplogInc field if non-nil, zero value otherwise.

### GetOplogIncOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetOplogIncOk() (*int32, bool)`

GetOplogIncOk returns a tuple with the OplogInc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogInc

`func (o *ApiAtlasDiskBackupRestoreJobView) SetOplogInc(v int32)`

SetOplogInc sets OplogInc field to given value.

### HasOplogInc

`func (o *ApiAtlasDiskBackupRestoreJobView) HasOplogInc() bool`

HasOplogInc returns a boolean if a field has been set.

### GetOplogTs

`func (o *ApiAtlasDiskBackupRestoreJobView) GetOplogTs() int32`

GetOplogTs returns the OplogTs field if non-nil, zero value otherwise.

### GetOplogTsOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetOplogTsOk() (*int32, bool)`

GetOplogTsOk returns a tuple with the OplogTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogTs

`func (o *ApiAtlasDiskBackupRestoreJobView) SetOplogTs(v int32)`

SetOplogTs sets OplogTs field to given value.

### HasOplogTs

`func (o *ApiAtlasDiskBackupRestoreJobView) HasOplogTs() bool`

HasOplogTs returns a boolean if a field has been set.

### GetPointInTimeUTCSeconds

`func (o *ApiAtlasDiskBackupRestoreJobView) GetPointInTimeUTCSeconds() int32`

GetPointInTimeUTCSeconds returns the PointInTimeUTCSeconds field if non-nil, zero value otherwise.

### GetPointInTimeUTCSecondsOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetPointInTimeUTCSecondsOk() (*int32, bool)`

GetPointInTimeUTCSecondsOk returns a tuple with the PointInTimeUTCSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeUTCSeconds

`func (o *ApiAtlasDiskBackupRestoreJobView) SetPointInTimeUTCSeconds(v int32)`

SetPointInTimeUTCSeconds sets PointInTimeUTCSeconds field to given value.

### HasPointInTimeUTCSeconds

`func (o *ApiAtlasDiskBackupRestoreJobView) HasPointInTimeUTCSeconds() bool`

HasPointInTimeUTCSeconds returns a boolean if a field has been set.

### GetSnapshotId

`func (o *ApiAtlasDiskBackupRestoreJobView) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ApiAtlasDiskBackupRestoreJobView) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ApiAtlasDiskBackupRestoreJobView) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetTargetClusterName

`func (o *ApiAtlasDiskBackupRestoreJobView) GetTargetClusterName() string`

GetTargetClusterName returns the TargetClusterName field if non-nil, zero value otherwise.

### GetTargetClusterNameOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetTargetClusterNameOk() (*string, bool)`

GetTargetClusterNameOk returns a tuple with the TargetClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterName

`func (o *ApiAtlasDiskBackupRestoreJobView) SetTargetClusterName(v string)`

SetTargetClusterName sets TargetClusterName field to given value.


### GetTargetGroupId

`func (o *ApiAtlasDiskBackupRestoreJobView) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *ApiAtlasDiskBackupRestoreJobView) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.


### GetTimestamp

`func (o *ApiAtlasDiskBackupRestoreJobView) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiAtlasDiskBackupRestoreJobView) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiAtlasDiskBackupRestoreJobView) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiAtlasDiskBackupRestoreJobView) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


