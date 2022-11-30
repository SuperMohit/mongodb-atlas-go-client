# ApiAtlasDiskBackupCopySettingView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider that stores the snapshot copy. | [optional] 
**Frequencies** | Pointer to **[]string** | List that describes which types of snapshots to copy. | [optional] 
**RegionName** | Pointer to **string** | Target region to copy snapshots belonging to replicationSpecId to. Please supply the &#39;Atlas Region&#39; which can be found under https://www.mongodb.com/docs/atlas/reference/cloud-providers/ &#39;regions&#39; link | [optional] 
**ReplicationSpecId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Replication Spec Id, do a GET request to Return One Cluster in One Project and consult the replicationSpecs array https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/returnOneCluster | [optional] 
**ShouldCopyOplogs** | Pointer to **bool** | Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores. | [optional] 

## Methods

### NewApiAtlasDiskBackupCopySettingView

`func NewApiAtlasDiskBackupCopySettingView() *ApiAtlasDiskBackupCopySettingView`

NewApiAtlasDiskBackupCopySettingView instantiates a new ApiAtlasDiskBackupCopySettingView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDiskBackupCopySettingViewWithDefaults

`func NewApiAtlasDiskBackupCopySettingViewWithDefaults() *ApiAtlasDiskBackupCopySettingView`

NewApiAtlasDiskBackupCopySettingViewWithDefaults instantiates a new ApiAtlasDiskBackupCopySettingView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ApiAtlasDiskBackupCopySettingView) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ApiAtlasDiskBackupCopySettingView) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ApiAtlasDiskBackupCopySettingView) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ApiAtlasDiskBackupCopySettingView) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetFrequencies

`func (o *ApiAtlasDiskBackupCopySettingView) GetFrequencies() []string`

GetFrequencies returns the Frequencies field if non-nil, zero value otherwise.

### GetFrequenciesOk

`func (o *ApiAtlasDiskBackupCopySettingView) GetFrequenciesOk() (*[]string, bool)`

GetFrequenciesOk returns a tuple with the Frequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencies

`func (o *ApiAtlasDiskBackupCopySettingView) SetFrequencies(v []string)`

SetFrequencies sets Frequencies field to given value.

### HasFrequencies

`func (o *ApiAtlasDiskBackupCopySettingView) HasFrequencies() bool`

HasFrequencies returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasDiskBackupCopySettingView) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasDiskBackupCopySettingView) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasDiskBackupCopySettingView) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasDiskBackupCopySettingView) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetReplicationSpecId

`func (o *ApiAtlasDiskBackupCopySettingView) GetReplicationSpecId() string`

GetReplicationSpecId returns the ReplicationSpecId field if non-nil, zero value otherwise.

### GetReplicationSpecIdOk

`func (o *ApiAtlasDiskBackupCopySettingView) GetReplicationSpecIdOk() (*string, bool)`

GetReplicationSpecIdOk returns a tuple with the ReplicationSpecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpecId

`func (o *ApiAtlasDiskBackupCopySettingView) SetReplicationSpecId(v string)`

SetReplicationSpecId sets ReplicationSpecId field to given value.

### HasReplicationSpecId

`func (o *ApiAtlasDiskBackupCopySettingView) HasReplicationSpecId() bool`

HasReplicationSpecId returns a boolean if a field has been set.

### GetShouldCopyOplogs

`func (o *ApiAtlasDiskBackupCopySettingView) GetShouldCopyOplogs() bool`

GetShouldCopyOplogs returns the ShouldCopyOplogs field if non-nil, zero value otherwise.

### GetShouldCopyOplogsOk

`func (o *ApiAtlasDiskBackupCopySettingView) GetShouldCopyOplogsOk() (*bool, bool)`

GetShouldCopyOplogsOk returns a tuple with the ShouldCopyOplogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCopyOplogs

`func (o *ApiAtlasDiskBackupCopySettingView) SetShouldCopyOplogs(v bool)`

SetShouldCopyOplogs sets ShouldCopyOplogs field to given value.

### HasShouldCopyOplogs

`func (o *ApiAtlasDiskBackupCopySettingView) HasShouldCopyOplogs() bool`

HasShouldCopyOplogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


