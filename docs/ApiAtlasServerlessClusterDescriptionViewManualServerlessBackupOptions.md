# ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerlessContinuousBackupEnabled** | Pointer to **bool** | Flag that indicates whether the serverless instances uses **Serverless Continuous Backup**. If this parameter is &#x60;false&#x60;, the serverless instance uses **Basic Backup**. | Option | Description | |---|---| | Serverless Continuous Backup | Atlas takes incremental [snapshots](https://www.mongodb.com/docs/atlas/backup/cloud-backup/overview/#std-label-serverless-snapshots) of the data in your serverless instance every six hours and lets you restore the data from a selected point in time within the last 72 hours. Atlas also takes daily snapshots and retains these daily snapshots for 35 days. To learn more, see [Serverless Instance Costs](https://www.mongodb.com/docs/atlas/billing/serverless-instance-costs/#std-label-serverless-instance-costs). | | Basic Backup | Atlas takes incremental [snapshots](https://www.mongodb.com/docs/atlas/backup/cloud-backup/overview/#std-label-serverless-snapshots) of the data in your serverless instance every six hours and retains only the two most recent snapshots. You can use this option for free. | | [optional] 

## Methods

### NewApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions

`func NewApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions() *ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions`

NewApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions instantiates a new ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptionsWithDefaults

`func NewApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptionsWithDefaults() *ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions`

NewApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptionsWithDefaults instantiates a new ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerlessContinuousBackupEnabled

`func (o *ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions) GetServerlessContinuousBackupEnabled() bool`

GetServerlessContinuousBackupEnabled returns the ServerlessContinuousBackupEnabled field if non-nil, zero value otherwise.

### GetServerlessContinuousBackupEnabledOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions) GetServerlessContinuousBackupEnabledOk() (*bool, bool)`

GetServerlessContinuousBackupEnabledOk returns a tuple with the ServerlessContinuousBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerlessContinuousBackupEnabled

`func (o *ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions) SetServerlessContinuousBackupEnabled(v bool)`

SetServerlessContinuousBackupEnabled sets ServerlessContinuousBackupEnabled field to given value.

### HasServerlessContinuousBackupEnabled

`func (o *ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions) HasServerlessContinuousBackupEnabled() bool`

HasServerlessContinuousBackupEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


