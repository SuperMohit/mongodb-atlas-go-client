# ApiAtlasClusterDescriptionV15View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupEnabled** | Pointer to **bool** | Flag that indicates whether the cluster can perform backups. If set to &#x60;true&#x60;, the cluster can perform backups. You must set this value to &#x60;true&#x60; for NVMe clusters. Backup uses [Cloud Backups](https://docs.atlas.mongodb.com/backup/cloud-backup/overview/) for dedicated clusters and [Shared Cluster Backups](https://docs.atlas.mongodb.com/backup/shared-tier/overview/) for tenant clusters. If set to &#x60;false&#x60;, the cluster doesn&#39;t use backups. | [optional] [default to false]
**BiConnector** | Pointer to [**ApiAtlasBiConnectorView**](ApiAtlasBiConnectorView.md) |  | [optional] 
**ClusterType** | Pointer to **string** | Configuration of nodes that comprise the cluster. | [optional] 
**ConnectionStrings** | Pointer to [**ApiAtlasClusterDescriptionConnectionStringsView**](ApiAtlasClusterDescriptionConnectionStringsView.md) |  | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this cluster. This parameter expresses its value in ISO 8601 format in UTC. | [optional] [readonly] 
**DiskSizeGB** | Pointer to **float64** | Storage capacity that the host&#39;s root volume possesses expressed in gigabytes. Increase this number to add capacity. MongoDB Cloud requires this parameter if you set **replicationSpecs**. If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value. Storage charge calculations depend on whether you choose the default value or a custom value.  The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier. | [optional] 
**EncryptionAtRestProvider** | Pointer to **string** | Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster. To enable customer key management for encryption at rest, the cluster **replicationSpecs[n].regionConfigs[m].{type}Specs.instanceSize** setting must be &#x60;M10&#x60; or higher and &#x60;\&quot;backupEnabled\&quot; : false&#x60; or omitted entirely. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Global Cluster. If you include existing zones in the request, you must specify this parameter. If you add a new zone to an existing Global Cluster, you may specify this parameter. The request deletes any existing zones in a Global Cluster that you exclude from the request. | [optional] [readonly] 
**Labels** | Pointer to [**[]ApiAtlasNDSLabelView**](ApiAtlasNDSLabelView.md) | Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. The MongoDB Cloud console doesn&#39;t display your labels. | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**MongoDBMajorVersion** | Pointer to **string** | Major MongoDB version of the cluster. MongoDB Cloud deploys the cluster with the latest stable release of the specified version. | [optional] [default to "5.0"]
**MongoDBVersion** | Pointer to **string** | Version of MongoDB that the cluster runs. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the advanced cluster. | [optional] 
**Paused** | Pointer to **bool** | Flag that indicates whether the cluster is paused. | [optional] 
**PitEnabled** | Pointer to **bool** | Flag that indicates whether the cluster uses continuous cloud backups. | [optional] 
**ReplicationSpecs** | Pointer to [**[]ApiAtlasReplicationSpecsV15Manual**](ApiAtlasReplicationSpecsV15Manual.md) | List of settings that configure your cluster regions. For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global sharded clusters and replica sets, this array has one object representing where your clusters nodes deploy. | [optional] 
**RootCertType** | Pointer to **string** | Root Certificate Authority that MongoDB Cloud cluster uses. MongoDB Cloud supports Internet Security Research Group. | [optional] [default to "ISRGROOTX1"]
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of this cluster. | [optional] [readonly] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]
**VersionReleaseSystem** | Pointer to **string** | Method by which the cluster maintains the MongoDB versions. If value is &#x60;CONTINUOUS&#x60;, you must not specify **mongoDBMajorVersion**. | [optional] [default to "LTS"]

## Methods

### NewApiAtlasClusterDescriptionV15View

`func NewApiAtlasClusterDescriptionV15View(links []Link, ) *ApiAtlasClusterDescriptionV15View`

NewApiAtlasClusterDescriptionV15View instantiates a new ApiAtlasClusterDescriptionV15View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasClusterDescriptionV15ViewWithDefaults

`func NewApiAtlasClusterDescriptionV15ViewWithDefaults() *ApiAtlasClusterDescriptionV15View`

NewApiAtlasClusterDescriptionV15ViewWithDefaults instantiates a new ApiAtlasClusterDescriptionV15View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupEnabled

`func (o *ApiAtlasClusterDescriptionV15View) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *ApiAtlasClusterDescriptionV15View) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *ApiAtlasClusterDescriptionV15View) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *ApiAtlasClusterDescriptionV15View) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.

### GetBiConnector

`func (o *ApiAtlasClusterDescriptionV15View) GetBiConnector() ApiAtlasBiConnectorView`

GetBiConnector returns the BiConnector field if non-nil, zero value otherwise.

### GetBiConnectorOk

`func (o *ApiAtlasClusterDescriptionV15View) GetBiConnectorOk() (*ApiAtlasBiConnectorView, bool)`

GetBiConnectorOk returns a tuple with the BiConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiConnector

`func (o *ApiAtlasClusterDescriptionV15View) SetBiConnector(v ApiAtlasBiConnectorView)`

SetBiConnector sets BiConnector field to given value.

### HasBiConnector

`func (o *ApiAtlasClusterDescriptionV15View) HasBiConnector() bool`

HasBiConnector returns a boolean if a field has been set.

### GetClusterType

`func (o *ApiAtlasClusterDescriptionV15View) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *ApiAtlasClusterDescriptionV15View) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *ApiAtlasClusterDescriptionV15View) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *ApiAtlasClusterDescriptionV15View) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetConnectionStrings

`func (o *ApiAtlasClusterDescriptionV15View) GetConnectionStrings() ApiAtlasClusterDescriptionConnectionStringsView`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *ApiAtlasClusterDescriptionV15View) GetConnectionStringsOk() (*ApiAtlasClusterDescriptionConnectionStringsView, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *ApiAtlasClusterDescriptionV15View) SetConnectionStrings(v ApiAtlasClusterDescriptionConnectionStringsView)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *ApiAtlasClusterDescriptionV15View) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.

### GetCreateDate

`func (o *ApiAtlasClusterDescriptionV15View) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ApiAtlasClusterDescriptionV15View) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ApiAtlasClusterDescriptionV15View) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ApiAtlasClusterDescriptionV15View) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *ApiAtlasClusterDescriptionV15View) GetDiskSizeGB() float64`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *ApiAtlasClusterDescriptionV15View) GetDiskSizeGBOk() (*float64, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *ApiAtlasClusterDescriptionV15View) SetDiskSizeGB(v float64)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *ApiAtlasClusterDescriptionV15View) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.

### GetEncryptionAtRestProvider

`func (o *ApiAtlasClusterDescriptionV15View) GetEncryptionAtRestProvider() string`

GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field if non-nil, zero value otherwise.

### GetEncryptionAtRestProviderOk

`func (o *ApiAtlasClusterDescriptionV15View) GetEncryptionAtRestProviderOk() (*string, bool)`

GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestProvider

`func (o *ApiAtlasClusterDescriptionV15View) SetEncryptionAtRestProvider(v string)`

SetEncryptionAtRestProvider sets EncryptionAtRestProvider field to given value.

### HasEncryptionAtRestProvider

`func (o *ApiAtlasClusterDescriptionV15View) HasEncryptionAtRestProvider() bool`

HasEncryptionAtRestProvider returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiAtlasClusterDescriptionV15View) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasClusterDescriptionV15View) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasClusterDescriptionV15View) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAtlasClusterDescriptionV15View) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasClusterDescriptionV15View) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasClusterDescriptionV15View) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasClusterDescriptionV15View) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasClusterDescriptionV15View) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *ApiAtlasClusterDescriptionV15View) GetLabels() []ApiAtlasNDSLabelView`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiAtlasClusterDescriptionV15View) GetLabelsOk() (*[]ApiAtlasNDSLabelView, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiAtlasClusterDescriptionV15View) SetLabels(v []ApiAtlasNDSLabelView)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiAtlasClusterDescriptionV15View) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasClusterDescriptionV15View) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasClusterDescriptionV15View) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasClusterDescriptionV15View) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetMongoDBMajorVersion

`func (o *ApiAtlasClusterDescriptionV15View) GetMongoDBMajorVersion() string`

GetMongoDBMajorVersion returns the MongoDBMajorVersion field if non-nil, zero value otherwise.

### GetMongoDBMajorVersionOk

`func (o *ApiAtlasClusterDescriptionV15View) GetMongoDBMajorVersionOk() (*string, bool)`

GetMongoDBMajorVersionOk returns a tuple with the MongoDBMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBMajorVersion

`func (o *ApiAtlasClusterDescriptionV15View) SetMongoDBMajorVersion(v string)`

SetMongoDBMajorVersion sets MongoDBMajorVersion field to given value.

### HasMongoDBMajorVersion

`func (o *ApiAtlasClusterDescriptionV15View) HasMongoDBMajorVersion() bool`

HasMongoDBMajorVersion returns a boolean if a field has been set.

### GetMongoDBVersion

`func (o *ApiAtlasClusterDescriptionV15View) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *ApiAtlasClusterDescriptionV15View) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *ApiAtlasClusterDescriptionV15View) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *ApiAtlasClusterDescriptionV15View) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.

### GetName

`func (o *ApiAtlasClusterDescriptionV15View) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasClusterDescriptionV15View) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasClusterDescriptionV15View) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAtlasClusterDescriptionV15View) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaused

`func (o *ApiAtlasClusterDescriptionV15View) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ApiAtlasClusterDescriptionV15View) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ApiAtlasClusterDescriptionV15View) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ApiAtlasClusterDescriptionV15View) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPitEnabled

`func (o *ApiAtlasClusterDescriptionV15View) GetPitEnabled() bool`

GetPitEnabled returns the PitEnabled field if non-nil, zero value otherwise.

### GetPitEnabledOk

`func (o *ApiAtlasClusterDescriptionV15View) GetPitEnabledOk() (*bool, bool)`

GetPitEnabledOk returns a tuple with the PitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitEnabled

`func (o *ApiAtlasClusterDescriptionV15View) SetPitEnabled(v bool)`

SetPitEnabled sets PitEnabled field to given value.

### HasPitEnabled

`func (o *ApiAtlasClusterDescriptionV15View) HasPitEnabled() bool`

HasPitEnabled returns a boolean if a field has been set.

### GetReplicationSpecs

`func (o *ApiAtlasClusterDescriptionV15View) GetReplicationSpecs() []ApiAtlasReplicationSpecsV15Manual`

GetReplicationSpecs returns the ReplicationSpecs field if non-nil, zero value otherwise.

### GetReplicationSpecsOk

`func (o *ApiAtlasClusterDescriptionV15View) GetReplicationSpecsOk() (*[]ApiAtlasReplicationSpecsV15Manual, bool)`

GetReplicationSpecsOk returns a tuple with the ReplicationSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSpecs

`func (o *ApiAtlasClusterDescriptionV15View) SetReplicationSpecs(v []ApiAtlasReplicationSpecsV15Manual)`

SetReplicationSpecs sets ReplicationSpecs field to given value.

### HasReplicationSpecs

`func (o *ApiAtlasClusterDescriptionV15View) HasReplicationSpecs() bool`

HasReplicationSpecs returns a boolean if a field has been set.

### GetRootCertType

`func (o *ApiAtlasClusterDescriptionV15View) GetRootCertType() string`

GetRootCertType returns the RootCertType field if non-nil, zero value otherwise.

### GetRootCertTypeOk

`func (o *ApiAtlasClusterDescriptionV15View) GetRootCertTypeOk() (*string, bool)`

GetRootCertTypeOk returns a tuple with the RootCertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCertType

`func (o *ApiAtlasClusterDescriptionV15View) SetRootCertType(v string)`

SetRootCertType sets RootCertType field to given value.

### HasRootCertType

`func (o *ApiAtlasClusterDescriptionV15View) HasRootCertType() bool`

HasRootCertType returns a boolean if a field has been set.

### GetStateName

`func (o *ApiAtlasClusterDescriptionV15View) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *ApiAtlasClusterDescriptionV15View) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *ApiAtlasClusterDescriptionV15View) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *ApiAtlasClusterDescriptionV15View) HasStateName() bool`

HasStateName returns a boolean if a field has been set.

### GetTerminationProtectionEnabled

`func (o *ApiAtlasClusterDescriptionV15View) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *ApiAtlasClusterDescriptionV15View) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *ApiAtlasClusterDescriptionV15View) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *ApiAtlasClusterDescriptionV15View) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.

### GetVersionReleaseSystem

`func (o *ApiAtlasClusterDescriptionV15View) GetVersionReleaseSystem() string`

GetVersionReleaseSystem returns the VersionReleaseSystem field if non-nil, zero value otherwise.

### GetVersionReleaseSystemOk

`func (o *ApiAtlasClusterDescriptionV15View) GetVersionReleaseSystemOk() (*string, bool)`

GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionReleaseSystem

`func (o *ApiAtlasClusterDescriptionV15View) SetVersionReleaseSystem(v string)`

SetVersionReleaseSystem sets VersionReleaseSystem field to given value.

### HasVersionReleaseSystem

`func (o *ApiAtlasClusterDescriptionV15View) HasVersionReleaseSystem() bool`

HasVersionReleaseSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


