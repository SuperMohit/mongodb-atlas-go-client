# ApiAtlasServerlessBackupRestoreJobViewManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | Pointer to **bool** | Flag that indicates whether someone canceled this restore job. | [optional] [readonly] 
**DeliveryType** | **string** | Human-readable label that categorizes the restore job to create. | 
**DeliveryUrl** | Pointer to **[]string** | One or more Uniform Resource Locators (URLs) that point to the compressed snapshot files for manual download. MongoDB Cloud returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;download\&quot;&#x60;.\&quot; | [optional] [readonly] 
**DesiredTimestamp** | Pointer to [**ApiBSONTimestampView**](ApiBSONTimestampView.md) |  | [optional] 
**Expired** | Pointer to **bool** | Flag that indicates whether the restore job expired. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Date and time when the restore job expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Failed** | Pointer to **bool** | Flag that indicates whether the restore job failed. | [optional] [readonly] 
**FinishedAt** | Pointer to **time.Time** | Date and time when the restore job completed. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the restore job. | [optional] [readonly] 
**SnapshotId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the snapshot. | [optional] 
**TargetClusterName** | **string** | Human-readable label that identifies the target cluster to which the restore job restores the snapshot. The resource returns this parameter when &#x60;\&quot;deliveryType\&quot;:&#x60; &#x60;\&quot;automated\&quot;&#x60;. | 
**TargetGroupId** | **string** | Unique 24-hexadecimal digit string that identifies the target project for the specified **targetClusterName**. | 
**Timestamp** | Pointer to **time.Time** | Date and time when MongoDB Cloud took the snapshot associated with **snapshotId**. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewApiAtlasServerlessBackupRestoreJobViewManual

`func NewApiAtlasServerlessBackupRestoreJobViewManual(deliveryType string, targetClusterName string, targetGroupId string, ) *ApiAtlasServerlessBackupRestoreJobViewManual`

NewApiAtlasServerlessBackupRestoreJobViewManual instantiates a new ApiAtlasServerlessBackupRestoreJobViewManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasServerlessBackupRestoreJobViewManualWithDefaults

`func NewApiAtlasServerlessBackupRestoreJobViewManualWithDefaults() *ApiAtlasServerlessBackupRestoreJobViewManual`

NewApiAtlasServerlessBackupRestoreJobViewManualWithDefaults instantiates a new ApiAtlasServerlessBackupRestoreJobViewManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.

### HasCancelled

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasCancelled() bool`

HasCancelled returns a boolean if a field has been set.

### GetDeliveryType

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.


### GetDeliveryUrl

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDeliveryUrl() []string`

GetDeliveryUrl returns the DeliveryUrl field if non-nil, zero value otherwise.

### GetDeliveryUrlOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDeliveryUrlOk() (*[]string, bool)`

GetDeliveryUrlOk returns a tuple with the DeliveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryUrl

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetDeliveryUrl(v []string)`

SetDeliveryUrl sets DeliveryUrl field to given value.

### HasDeliveryUrl

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasDeliveryUrl() bool`

HasDeliveryUrl returns a boolean if a field has been set.

### GetDesiredTimestamp

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDesiredTimestamp() ApiBSONTimestampView`

GetDesiredTimestamp returns the DesiredTimestamp field if non-nil, zero value otherwise.

### GetDesiredTimestampOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetDesiredTimestampOk() (*ApiBSONTimestampView, bool)`

GetDesiredTimestampOk returns a tuple with the DesiredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTimestamp

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetDesiredTimestamp(v ApiBSONTimestampView)`

SetDesiredTimestamp sets DesiredTimestamp field to given value.

### HasDesiredTimestamp

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasDesiredTimestamp() bool`

HasDesiredTimestamp returns a boolean if a field has been set.

### GetExpired

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFailed

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetFinishedAt

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSnapshotId

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetTargetClusterName

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTargetClusterName() string`

GetTargetClusterName returns the TargetClusterName field if non-nil, zero value otherwise.

### GetTargetClusterNameOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTargetClusterNameOk() (*string, bool)`

GetTargetClusterNameOk returns a tuple with the TargetClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterName

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetTargetClusterName(v string)`

SetTargetClusterName sets TargetClusterName field to given value.


### GetTargetGroupId

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.


### GetTimestamp

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiAtlasServerlessBackupRestoreJobViewManual) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


