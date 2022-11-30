# ApiAtlasLegacyClusterDescriptionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**ApiAtlasAutoScalingView**](ApiAtlasAutoScalingView.md) |  | [optional] 
**BackupEnabled** | Pointer to **bool** | Flag that indicates whether the cluster can perform backups. If set to &#x60;true&#x60;, the cluster can perform backups. You must set this value to &#x60;true&#x60; for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to &#x60;false&#x60;, the cluster doesn&#39;t use MongoDB Cloud backups. | [optional] 
**BiConnector** | Pointer to [**ApiAtlasBiConnectorView**](ApiAtlasBiConnectorView.md) |  | [optional] 
**ClusterType** | Pointer to **string** | Configuration of nodes that comprise the cluster. | [optional] 
**ConnectionStrings** | Pointer to [**ApiAtlasClusterDescriptionConnectionStringsView**](ApiAtlasClusterDescriptionConnectionStringsView.md) |  | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this cluster. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**DiskSizeGB** | Pointer to **float32** | Storage capacity that the host&#39;s root volume possesses expressed in gigabytes. Increase this number to add capacity. MongoDB Cloud requires this parameter if you set **replicationSpecs**. If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value. Storage charge calculations depend on whether you choose the default value or a custom value.  The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier. | [optional] 
**EncryptionAtRestProvider** | Pointer to **string** | Cloud service provider that manages your customer keys to provide an additional layer of Encryption at Rest for the cluster. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster. | [optional] [readonly] 
**Labels** | Pointer to [**[]ApiAtlasNDSLabelView**](ApiAtlasNDSLabelView.md) | Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. The MongoDB Cloud console doesn&#39;t display your labels. | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**MongoDBMajorVersion** | Pointer to **string** | Major MongoDB version of the cluster. MongoDB Cloud deploys the cluster with the latest stable release of the specified version. | [optional] [default to "4.4"]
**MongoDBVersion** | Pointer to **string** | Version of MongoDB that the cluster runs. | [optional] [readonly] 
**MongoURI** | Pointer to **string** | Base connection string that you can use to connect to the cluster. MongoDB Cloud displays the string only after the cluster starts, not while it builds the cluster. | [optional] [readonly] 
**MongoURIUpdated** | Pointer to **time.Time** | Date and time when someone last updated the connection string. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**MongoURIWithOptions** | Pointer to **string** | Connection string that you can use to connect to the cluster including the &#x60;replicaSet&#x60;, &#x60;ssl&#x60;, and &#x60;authSource&#x60; query parameters with values appropriate for the cluster. You may need to add MongoDB database users. The response returns this parameter once the cluster can receive requests, not while it builds the cluster. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the cluster. | [optional] 
**NumShards** | Pointer to **int32** | Number of shards up to 50 to deploy for a sharded cluster. The resource returns &#x60;1&#x60; to indicate a replica set and values of &#x60;2&#x60; and higher to indicate a sharded cluster. The returned value equals the number of shards in the cluster. | [optional] [default to 1]
**Paused** | Pointer to **bool** | Flag that indicates whether the cluster is paused. | [optional] 
**PitEnabled** | Pointer to **bool** | Flag that indicates whether the cluster uses continuous cloud backups. | [optional] 
**ProviderBackupEnabled** | Pointer to **bool** | Flag that indicates whether the M10 or higher cluster can perform Cloud Backups. If set to &#x60;true&#x60;, the cluster can perform backups. If this and **backupEnabled** are set to &#x60;false&#x60;, the cluster doesn&#39;t use MongoDB Cloud backups. | [optional] 
**ProviderSettings** | Pointer to [**ApiAtlasProviderSettingsViewManual**](ApiAtlasProviderSettingsViewManual.md) |  | [optional] 
**ReplicationFactor** | Pointer to **int32** | Number of members that belong to the replica set. Each member retains a copy of your databases, providing high availability and data redundancy. Use **replicationSpecs** instead. | [optional] [default to 3]
**ReplicationSpec** | Pointer to [**map[string]ApiAtlasRegionSpecView**](ApiAtlasRegionSpecView.md) | Physical location where MongoDB Cloud provisions cluster nodes. | [optional] 
**ReplicationSpecs** | Pointer to [**[]ApiAtlasLegacyReplicationSpecViewManual**](ApiAtlasLegacyReplicationSpecViewManual.md) | List of settings that configure your cluster regions.  - For Global Clusters, each object in the array represents one zone where MongoDB Cloud deploys your clusters nodes. - For non-Global sharded clusters and replica sets, the single object represents where MongoDB Cloud deploys your clusters nodes. | [optional] 
**RootCertType** | Pointer to **string** | Root Certificate Authority that MongoDB Atlas clusters uses. MongoDB Cloud supports Internet Security Research Group. | [optional] [default to "ISRGROOTX1"]
**SrvAddress** | Pointer to **string** | Connection string that you can use to connect to the cluster. The &#x60;+srv&#x60; modifier forces the connection to use Transport Layer Security (TLS). The &#x60;mongoURI&#x60; parameter lists additional options. | [optional] [readonly] 
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of the cluster. | [optional] [readonly] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]
**VersionReleaseSystem** | Pointer to **string** | Method by which the cluster maintains the MongoDB versions. If value is &#x60;CONTINUOUS&#x60;, you must not specify **mongoDBMajorVersion**. | [optional] [default to "LTS"]

## Methods

### NewApiAtlasLegacyClusterDescriptionView

`func NewApiAtlasLegacyClusterDescriptionView(links []Link, ) *ApiAtlasLegacyClusterDescriptionView`

NewApiAtlasLegacyClusterDescriptionView instantiates a new ApiAtlasLegacyClusterDescriptionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasLegacyClusterDescriptionViewWithDefaults

`func NewApiAtlasLegacyClusterDescriptionViewWithDefaults() *ApiAtlasLegacyClusterDescriptionView`

NewApiAtlasLegacyClusterDescriptionViewWithDefaults instantiates a new ApiAtlasLegacyClusterDescriptionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *ApiAtlasLegacyClusterDescriptionView) GetAutoScaling() ApiAtlasAutoScalingView`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetAutoScalingOk() (*ApiAtlasAutoScalingView, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *ApiAtlasLegacyClusterDescriptionView) SetAutoScaling(v ApiAtlasAutoScalingView)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *ApiAtlasLegacyClusterDescriptionView) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetBackupEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.

### GetBiConnector

`func (o *ApiAtlasLegacyClusterDescriptionView) GetBiConnector() ApiAtlasBiConnectorView`

GetBiConnector returns the BiConnector field if non-nil, zero value otherwise.

### GetBiConnectorOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetBiConnectorOk() (*ApiAtlasBiConnectorView, bool)`

GetBiConnectorOk returns a tuple with the BiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiConnector

`func (o *ApiAtlasLegacyClusterDescriptionView) SetBiConnector(v ApiAtlasBiConnectorView)`

SetBiConnector sets BiConnector field to given value.

### HasBiConnector

`func (o *ApiAtlasLegacyClusterDescriptionView) HasBiConnector() bool`

HasBiConnector returns a boolean if a field has been set.

### GetClusterType

`func (o *ApiAtlasLegacyClusterDescriptionView) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *ApiAtlasLegacyClusterDescriptionView) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *ApiAtlasLegacyClusterDescriptionView) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetConnectionStrings

`func (o *ApiAtlasLegacyClusterDescriptionView) GetConnectionStrings() ApiAtlasClusterDescriptionConnectionStringsView`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetConnectionStringsOk() (*ApiAtlasClusterDescriptionConnectionStringsView, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *ApiAtlasLegacyClusterDescriptionView) SetConnectionStrings(v ApiAtlasClusterDescriptionConnectionStringsView)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *ApiAtlasLegacyClusterDescriptionView) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.

### GetCreateDate

`func (o *ApiAtlasLegacyClusterDescriptionView) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ApiAtlasLegacyClusterDescriptionView) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ApiAtlasLegacyClusterDescriptionView) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *ApiAtlasLegacyClusterDescriptionView) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *ApiAtlasLegacyClusterDescriptionView) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *ApiAtlasLegacyClusterDescriptionView) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.

### GetEncryptionAtRestProvider

`func (o *ApiAtlasLegacyClusterDescriptionView) GetEncryptionAtRestProvider() string`

GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field if non-nil, zero value otherwise.

### GetEncryptionAtRestProviderOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetEncryptionAtRestProviderOk() (*string, bool)`

GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestProvider

`func (o *ApiAtlasLegacyClusterDescriptionView) SetEncryptionAtRestProvider(v string)`

SetEncryptionAtRestProvider sets EncryptionAtRestProvider field to given value.

### HasEncryptionAtRestProvider

`func (o *ApiAtlasLegacyClusterDescriptionView) HasEncryptionAtRestProvider() bool`

HasEncryptionAtRestProvider returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiAtlasLegacyClusterDescriptionView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasLegacyClusterDescriptionView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAtlasLegacyClusterDescriptionView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasLegacyClusterDescriptionView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasLegacyClusterDescriptionView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasLegacyClusterDescriptionView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *ApiAtlasLegacyClusterDescriptionView) GetLabels() []ApiAtlasNDSLabelView`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetLabelsOk() (*[]ApiAtlasNDSLabelView, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiAtlasLegacyClusterDescriptionView) SetLabels(v []ApiAtlasNDSLabelView)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiAtlasLegacyClusterDescriptionView) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasLegacyClusterDescriptionView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasLegacyClusterDescriptionView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetMongoDBMajorVersion

`func (o *ApiAtlasLegacyClusterDescriptionView) GetMongoDBMajorVersion() string`

GetMongoDBMajorVersion returns the MongoDBMajorVersion field if non-nil, zero value otherwise.

### GetMongoDBMajorVersionOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetMongoDBMajorVersionOk() (*string, bool)`

GetMongoDBMajorVersionOk returns a tuple with the MongoDBMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBMajorVersion

`func (o *ApiAtlasLegacyClusterDescriptionView) SetMongoDBMajorVersion(v string)`

SetMongoDBMajorVersion sets MongoDBMajorVersion field to given value.

### HasMongoDBMajorVersion

`func (o *ApiAtlasLegacyClusterDescriptionView) HasMongoDBMajorVersion() bool`

HasMongoDBMajorVersion returns a boolean if a field has been set.

### GetMongoDBVersion

`func (o *ApiAtlasLegacyClusterDescriptionView) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *ApiAtlasLegacyClusterDescriptionView) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *ApiAtlasLegacyClusterDescriptionView) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.

### GetMongoURI

`func (o *ApiAtlasLegacyClusterDescriptionView) GetMongoURI() string`

GetMongoURI returns the MongoURI field if non-nil, zero value otherwise.

### GetMongoURIOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetMongoURIOk() (*string, bool)`

GetMongoURIOk returns a tuple with the MongoURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURI

`func (o *ApiAtlasLegacyClusterDescriptionView) SetMongoURI(v string)`

SetMongoURI sets MongoURI field to given value.

### HasMongoURI

`func (o *ApiAtlasLegacyClusterDescriptionView) HasMongoURI() bool`

HasMongoURI returns a boolean if a field has been set.

### GetMongoURIUpdated

`func (o *ApiAtlasLegacyClusterDescriptionView) GetMongoURIUpdated() time.Time`

GetMongoURIUpdated returns the MongoURIUpdated field if non-nil, zero value otherwise.

### GetMongoURIUpdatedOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetMongoURIUpdatedOk() (*time.Time, bool)`

GetMongoURIUpdatedOk returns a tuple with the MongoURIUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURIUpdated

`func (o *ApiAtlasLegacyClusterDescriptionView) SetMongoURIUpdated(v time.Time)`

SetMongoURIUpdated sets MongoURIUpdated field to given value.

### HasMongoURIUpdated

`func (o *ApiAtlasLegacyClusterDescriptionView) HasMongoURIUpdated() bool`

HasMongoURIUpdated returns a boolean if a field has been set.

### GetMongoURIWithOptions

`func (o *ApiAtlasLegacyClusterDescriptionView) GetMongoURIWithOptions() string`

GetMongoURIWithOptions returns the MongoURIWithOptions field if non-nil, zero value otherwise.

### GetMongoURIWithOptionsOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetMongoURIWithOptionsOk() (*string, bool)`

GetMongoURIWithOptionsOk returns a tuple with the MongoURIWithOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoURIWithOptions

`func (o *ApiAtlasLegacyClusterDescriptionView) SetMongoURIWithOptions(v string)`

SetMongoURIWithOptions sets MongoURIWithOptions field to given value.

### HasMongoURIWithOptions

`func (o *ApiAtlasLegacyClusterDescriptionView) HasMongoURIWithOptions() bool`

HasMongoURIWithOptions returns a boolean if a field has been set.

### GetName

`func (o *ApiAtlasLegacyClusterDescriptionView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasLegacyClusterDescriptionView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAtlasLegacyClusterDescriptionView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumShards

`func (o *ApiAtlasLegacyClusterDescriptionView) GetNumShards() int32`

GetNumShards returns the NumShards field if non-nil, zero value otherwise.

### GetNumShardsOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetNumShardsOk() (*int32, bool)`

GetNumShardsOk returns a tuple with the NumShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumShards

`func (o *ApiAtlasLegacyClusterDescriptionView) SetNumShards(v int32)`

SetNumShards sets NumShards field to given value.

### HasNumShards

`func (o *ApiAtlasLegacyClusterDescriptionView) HasNumShards() bool`

HasNumShards returns a boolean if a field has been set.

### GetPaused

`func (o *ApiAtlasLegacyClusterDescriptionView) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ApiAtlasLegacyClusterDescriptionView) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ApiAtlasLegacyClusterDescriptionView) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPitEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) GetPitEnabled() bool`

GetPitEnabled returns the PitEnabled field if non-nil, zero value otherwise.

### GetPitEnabledOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetPitEnabledOk() (*bool, bool)`

GetPitEnabledOk returns a tuple with the PitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) SetPitEnabled(v bool)`

SetPitEnabled sets PitEnabled field to given value.

### HasPitEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) HasPitEnabled() bool`

HasPitEnabled returns a boolean if a field has been set.

### GetProviderBackupEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) GetProviderBackupEnabled() bool`

GetProviderBackupEnabled returns the ProviderBackupEnabled field if non-nil, zero value otherwise.

### GetProviderBackupEnabledOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetProviderBackupEnabledOk() (*bool, bool)`

GetProviderBackupEnabledOk returns a tuple with the ProviderBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderBackupEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) SetProviderBackupEnabled(v bool)`

SetProviderBackupEnabled sets ProviderBackupEnabled field to given value.

### HasProviderBackupEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) HasProviderBackupEnabled() bool`

HasProviderBackupEnabled returns a boolean if a field has been set.

### GetProviderSettings

`func (o *ApiAtlasLegacyClusterDescriptionView) GetProviderSettings() ApiAtlasProviderSettingsViewManual`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetProviderSettingsOk() (*ApiAtlasProviderSettingsViewManual, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ApiAtlasLegacyClusterDescriptionView) SetProviderSettings(v ApiAtlasProviderSettingsViewManual)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *ApiAtlasLegacyClusterDescriptionView) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *ApiAtlasLegacyClusterDescriptionView) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *ApiAtlasLegacyClusterDescriptionView) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *ApiAtlasLegacyClusterDescriptionView) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetReplicationSpec

`func (o *ApiAtlasLegacyClusterDescriptionView) GetReplicationSpec() map[string]ApiAtlasRegionSpecView`

GetReplicationSpec returns the ReplicationSpec field if non-nil, zero value otherwise.

### GetReplicationSpecOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetReplicationSpecOk() (*map[string]ApiAtlasRegionSpecView, bool)`

GetReplicationSpecOk returns a tuple with the ReplicationSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpec

`func (o *ApiAtlasLegacyClusterDescriptionView) SetReplicationSpec(v map[string]ApiAtlasRegionSpecView)`

SetReplicationSpec sets ReplicationSpec field to given value.

### HasReplicationSpec

`func (o *ApiAtlasLegacyClusterDescriptionView) HasReplicationSpec() bool`

HasReplicationSpec returns a boolean if a field has been set.

### GetReplicationSpecs

`func (o *ApiAtlasLegacyClusterDescriptionView) GetReplicationSpecs() []ApiAtlasLegacyReplicationSpecViewManual`

GetReplicationSpecs returns the ReplicationSpecs field if non-nil, zero value otherwise.

### GetReplicationSpecsOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetReplicationSpecsOk() (*[]ApiAtlasLegacyReplicationSpecViewManual, bool)`

GetReplicationSpecsOk returns a tuple with the ReplicationSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpecs

`func (o *ApiAtlasLegacyClusterDescriptionView) SetReplicationSpecs(v []ApiAtlasLegacyReplicationSpecViewManual)`

SetReplicationSpecs sets ReplicationSpecs field to given value.

### HasReplicationSpecs

`func (o *ApiAtlasLegacyClusterDescriptionView) HasReplicationSpecs() bool`

HasReplicationSpecs returns a boolean if a field has been set.

### GetRootCertType

`func (o *ApiAtlasLegacyClusterDescriptionView) GetRootCertType() string`

GetRootCertType returns the RootCertType field if non-nil, zero value otherwise.

### GetRootCertTypeOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetRootCertTypeOk() (*string, bool)`

GetRootCertTypeOk returns a tuple with the RootCertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCertType

`func (o *ApiAtlasLegacyClusterDescriptionView) SetRootCertType(v string)`

SetRootCertType sets RootCertType field to given value.

### HasRootCertType

`func (o *ApiAtlasLegacyClusterDescriptionView) HasRootCertType() bool`

HasRootCertType returns a boolean if a field has been set.

### GetSrvAddress

`func (o *ApiAtlasLegacyClusterDescriptionView) GetSrvAddress() string`

GetSrvAddress returns the SrvAddress field if non-nil, zero value otherwise.

### GetSrvAddressOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetSrvAddressOk() (*string, bool)`

GetSrvAddressOk returns a tuple with the SrvAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrvAddress

`func (o *ApiAtlasLegacyClusterDescriptionView) SetSrvAddress(v string)`

SetSrvAddress sets SrvAddress field to given value.

### HasSrvAddress

`func (o *ApiAtlasLegacyClusterDescriptionView) HasSrvAddress() bool`

HasSrvAddress returns a boolean if a field has been set.

### GetStateName

`func (o *ApiAtlasLegacyClusterDescriptionView) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *ApiAtlasLegacyClusterDescriptionView) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *ApiAtlasLegacyClusterDescriptionView) HasStateName() bool`

HasStateName returns a boolean if a field has been set.

### GetTerminationProtectionEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *ApiAtlasLegacyClusterDescriptionView) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.

### GetVersionReleaseSystem

`func (o *ApiAtlasLegacyClusterDescriptionView) GetVersionReleaseSystem() string`

GetVersionReleaseSystem returns the VersionReleaseSystem field if non-nil, zero value otherwise.

### GetVersionReleaseSystemOk

`func (o *ApiAtlasLegacyClusterDescriptionView) GetVersionReleaseSystemOk() (*string, bool)`

GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionReleaseSystem

`func (o *ApiAtlasLegacyClusterDescriptionView) SetVersionReleaseSystem(v string)`

SetVersionReleaseSystem sets VersionReleaseSystem field to given value.

### HasVersionReleaseSystem

`func (o *ApiAtlasLegacyClusterDescriptionView) HasVersionReleaseSystem() bool`

HasVersionReleaseSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


