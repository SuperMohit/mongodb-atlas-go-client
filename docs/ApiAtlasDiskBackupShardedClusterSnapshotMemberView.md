# ApiAtlasDiskBackupShardedClusterSnapshotMemberView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Human-readable label that identifies the cloud provider that stores this snapshot. The resource returns this parameter when &#x60;\&quot;type\&quot;: \&quot;replicaSet\&quot;.&#x60; | [readonly] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies the snapshot. | [readonly] 
**ReplicaSetName** | **string** | Human-readable label that identifies the shard or config host from which MongoDB Cloud took this snapshot. | [readonly] 

## Methods

### NewApiAtlasDiskBackupShardedClusterSnapshotMemberView

`func NewApiAtlasDiskBackupShardedClusterSnapshotMemberView(cloudProvider string, id string, replicaSetName string, ) *ApiAtlasDiskBackupShardedClusterSnapshotMemberView`

NewApiAtlasDiskBackupShardedClusterSnapshotMemberView instantiates a new ApiAtlasDiskBackupShardedClusterSnapshotMemberView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDiskBackupShardedClusterSnapshotMemberViewWithDefaults

`func NewApiAtlasDiskBackupShardedClusterSnapshotMemberViewWithDefaults() *ApiAtlasDiskBackupShardedClusterSnapshotMemberView`

NewApiAtlasDiskBackupShardedClusterSnapshotMemberViewWithDefaults instantiates a new ApiAtlasDiskBackupShardedClusterSnapshotMemberView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetId

`func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) SetId(v string)`

SetId sets Id field to given value.


### GetReplicaSetName

`func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *ApiAtlasDiskBackupShardedClusterSnapshotMemberView) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


