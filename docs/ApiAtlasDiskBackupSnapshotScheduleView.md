# ApiAtlasDiskBackupSnapshotScheduleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoExportEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud automatically exports cloud backup snapshots to the AWS bucket. | [optional] 
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster with the snapshot you want to return. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster with the snapshot you want to return. | [optional] [readonly] 
**CopySettings** | Pointer to [**[]ApiAtlasDiskBackupCopySettingView**](ApiAtlasDiskBackupCopySettingView.md) | List that contains a document for each copy setting item in the desired backup policy. | [optional] 
**DeleteCopiedBackups** | Pointer to [**[]ApiDeleteCopiedBackupsView**](ApiDeleteCopiedBackupsView.md) | List that contains a document for each deleted copy setting whose backup copies you want to delete. | [optional] 
**Export** | Pointer to [**AutoExportPolicyView**](AutoExportPolicyView.md) |  | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**NextSnapshot** | Pointer to **time.Time** | Date and time when MongoDB Cloud takes the next snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Policies** | [**[]ApiPolicyView**](ApiPolicyView.md) | Rules set for this backup schedule. | 
**ReferenceHourOfDay** | Pointer to **int32** | Hour of day in Coordinated Universal Time (UTC) that represents when MongoDB Cloud takes the snapshot. | [optional] 
**ReferenceMinuteOfHour** | Pointer to **int32** | Minute of the **referenceHourOfDay** that represents when MongoDB Cloud takes the snapshot. | [optional] 
**RestoreWindowDays** | Pointer to **int32** | Number of previous days that you can restore back to with Continuous Cloud Backup accuracy. You must specify a positive, non-zero integer. This parameter applies to continuous cloud backups only. | [optional] 
**UpdateSnapshots** | Pointer to **bool** | Flag that indicates whether to apply the retention changes in the updated backup policy to snapshots that MongoDB Cloud took previously. | [optional] 
**UseOrgAndGroupNamesInExportPrefix** | Pointer to **bool** | Flag that indicates whether to use organization and project names instead of organization and project UUIDs in the path to the metadata files that MongoDB Cloud uploads to your AWS bucket. | [optional] 

## Methods

### NewApiAtlasDiskBackupSnapshotScheduleView

`func NewApiAtlasDiskBackupSnapshotScheduleView(links []Link, policies []ApiPolicyView, ) *ApiAtlasDiskBackupSnapshotScheduleView`

NewApiAtlasDiskBackupSnapshotScheduleView instantiates a new ApiAtlasDiskBackupSnapshotScheduleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDiskBackupSnapshotScheduleViewWithDefaults

`func NewApiAtlasDiskBackupSnapshotScheduleViewWithDefaults() *ApiAtlasDiskBackupSnapshotScheduleView`

NewApiAtlasDiskBackupSnapshotScheduleViewWithDefaults instantiates a new ApiAtlasDiskBackupSnapshotScheduleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoExportEnabled

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetAutoExportEnabled() bool`

GetAutoExportEnabled returns the AutoExportEnabled field if non-nil, zero value otherwise.

### GetAutoExportEnabledOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetAutoExportEnabledOk() (*bool, bool)`

GetAutoExportEnabledOk returns a tuple with the AutoExportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExportEnabled

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetAutoExportEnabled(v bool)`

SetAutoExportEnabled sets AutoExportEnabled field to given value.

### HasAutoExportEnabled

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasAutoExportEnabled() bool`

HasAutoExportEnabled returns a boolean if a field has been set.

### GetClusterId

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterName

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCopySettings

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetCopySettings() []ApiAtlasDiskBackupCopySettingView`

GetCopySettings returns the CopySettings field if non-nil, zero value otherwise.

### GetCopySettingsOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetCopySettingsOk() (*[]ApiAtlasDiskBackupCopySettingView, bool)`

GetCopySettingsOk returns a tuple with the CopySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySettings

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetCopySettings(v []ApiAtlasDiskBackupCopySettingView)`

SetCopySettings sets CopySettings field to given value.

### HasCopySettings

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasCopySettings() bool`

HasCopySettings returns a boolean if a field has been set.

### GetDeleteCopiedBackups

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetDeleteCopiedBackups() []ApiDeleteCopiedBackupsView`

GetDeleteCopiedBackups returns the DeleteCopiedBackups field if non-nil, zero value otherwise.

### GetDeleteCopiedBackupsOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetDeleteCopiedBackupsOk() (*[]ApiDeleteCopiedBackupsView, bool)`

GetDeleteCopiedBackupsOk returns a tuple with the DeleteCopiedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteCopiedBackups

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetDeleteCopiedBackups(v []ApiDeleteCopiedBackupsView)`

SetDeleteCopiedBackups sets DeleteCopiedBackups field to given value.

### HasDeleteCopiedBackups

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasDeleteCopiedBackups() bool`

HasDeleteCopiedBackups returns a boolean if a field has been set.

### GetExport

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetExport() AutoExportPolicyView`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetExportOk() (*AutoExportPolicyView, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetExport(v AutoExportPolicyView)`

SetExport sets Export field to given value.

### HasExport

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetNextSnapshot

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetNextSnapshot() time.Time`

GetNextSnapshot returns the NextSnapshot field if non-nil, zero value otherwise.

### GetNextSnapshotOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetNextSnapshotOk() (*time.Time, bool)`

GetNextSnapshotOk returns a tuple with the NextSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSnapshot

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetNextSnapshot(v time.Time)`

SetNextSnapshot sets NextSnapshot field to given value.

### HasNextSnapshot

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasNextSnapshot() bool`

HasNextSnapshot returns a boolean if a field has been set.

### GetPolicies

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetPolicies() []ApiPolicyView`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetPoliciesOk() (*[]ApiPolicyView, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetPolicies(v []ApiPolicyView)`

SetPolicies sets Policies field to given value.


### GetReferenceHourOfDay

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetReferenceHourOfDay() int32`

GetReferenceHourOfDay returns the ReferenceHourOfDay field if non-nil, zero value otherwise.

### GetReferenceHourOfDayOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetReferenceHourOfDayOk() (*int32, bool)`

GetReferenceHourOfDayOk returns a tuple with the ReferenceHourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceHourOfDay

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetReferenceHourOfDay(v int32)`

SetReferenceHourOfDay sets ReferenceHourOfDay field to given value.

### HasReferenceHourOfDay

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasReferenceHourOfDay() bool`

HasReferenceHourOfDay returns a boolean if a field has been set.

### GetReferenceMinuteOfHour

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetReferenceMinuteOfHour() int32`

GetReferenceMinuteOfHour returns the ReferenceMinuteOfHour field if non-nil, zero value otherwise.

### GetReferenceMinuteOfHourOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetReferenceMinuteOfHourOk() (*int32, bool)`

GetReferenceMinuteOfHourOk returns a tuple with the ReferenceMinuteOfHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceMinuteOfHour

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetReferenceMinuteOfHour(v int32)`

SetReferenceMinuteOfHour sets ReferenceMinuteOfHour field to given value.

### HasReferenceMinuteOfHour

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasReferenceMinuteOfHour() bool`

HasReferenceMinuteOfHour returns a boolean if a field has been set.

### GetRestoreWindowDays

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetRestoreWindowDays() int32`

GetRestoreWindowDays returns the RestoreWindowDays field if non-nil, zero value otherwise.

### GetRestoreWindowDaysOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetRestoreWindowDaysOk() (*int32, bool)`

GetRestoreWindowDaysOk returns a tuple with the RestoreWindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreWindowDays

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetRestoreWindowDays(v int32)`

SetRestoreWindowDays sets RestoreWindowDays field to given value.

### HasRestoreWindowDays

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasRestoreWindowDays() bool`

HasRestoreWindowDays returns a boolean if a field has been set.

### GetUpdateSnapshots

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetUpdateSnapshots() bool`

GetUpdateSnapshots returns the UpdateSnapshots field if non-nil, zero value otherwise.

### GetUpdateSnapshotsOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetUpdateSnapshotsOk() (*bool, bool)`

GetUpdateSnapshotsOk returns a tuple with the UpdateSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSnapshots

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetUpdateSnapshots(v bool)`

SetUpdateSnapshots sets UpdateSnapshots field to given value.

### HasUpdateSnapshots

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasUpdateSnapshots() bool`

HasUpdateSnapshots returns a boolean if a field has been set.

### GetUseOrgAndGroupNamesInExportPrefix

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetUseOrgAndGroupNamesInExportPrefix() bool`

GetUseOrgAndGroupNamesInExportPrefix returns the UseOrgAndGroupNamesInExportPrefix field if non-nil, zero value otherwise.

### GetUseOrgAndGroupNamesInExportPrefixOk

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) GetUseOrgAndGroupNamesInExportPrefixOk() (*bool, bool)`

GetUseOrgAndGroupNamesInExportPrefixOk returns a tuple with the UseOrgAndGroupNamesInExportPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOrgAndGroupNamesInExportPrefix

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) SetUseOrgAndGroupNamesInExportPrefix(v bool)`

SetUseOrgAndGroupNamesInExportPrefix sets UseOrgAndGroupNamesInExportPrefix field to given value.

### HasUseOrgAndGroupNamesInExportPrefix

`func (o *ApiAtlasDiskBackupSnapshotScheduleView) HasUseOrgAndGroupNamesInExportPrefix() bool`

HasUseOrgAndGroupNamesInExportPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


