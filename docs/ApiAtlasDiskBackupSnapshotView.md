# ApiAtlasDiskBackupSnapshotView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud took the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Description** | Pointer to **string** | Human-readable phrase or sentence that explains the purpose of the snapshot. The resource returns this parameter when &#x60;\&quot;status\&quot;: \&quot;onDemand\&quot;&#x60;. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud deletes the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FrequencyType** | Pointer to **string** | Human-readable label that identifies how often this snapshot triggers. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the snapshot. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**MasterKeyUUID** | Pointer to **string** | Unique string that identifies the Amazon Web Services (AWS) Key Management Service (KMS) Customer Master Key (CMK) used to encrypt the snapshot. The resource returns this value when &#x60;\&quot;encryptionEnabled\&quot; : true&#x60;. | [optional] [readonly] 
**MongodVersion** | Pointer to **string** | Version of the MongoDB host that this snapshot backs up. | [optional] [readonly] 
**PolicyItems** | Pointer to **[]string** | List that contains unique identifiers for the policy items. | [optional] [readonly] 
**SnapshotType** | Pointer to **string** | Human-readable label that identifies when this snapshot triggers. | [optional] [readonly] 
**Status** | Pointer to **string** | Human-readable label that indicates the stage of the backup process for this snapshot. | [optional] [readonly] 
**StorageSizeBytes** | Pointer to **int64** | Number of bytes taken to store the backup snapshot. | [optional] [readonly] 
**Type** | Pointer to **string** | Human-readable label that categorizes the cluster as a replica set or sharded cluster. | [optional] [readonly] 

## Methods

### NewApiAtlasDiskBackupSnapshotView

`func NewApiAtlasDiskBackupSnapshotView(links []Link, ) *ApiAtlasDiskBackupSnapshotView`

NewApiAtlasDiskBackupSnapshotView instantiates a new ApiAtlasDiskBackupSnapshotView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDiskBackupSnapshotViewWithDefaults

`func NewApiAtlasDiskBackupSnapshotViewWithDefaults() *ApiAtlasDiskBackupSnapshotView`

NewApiAtlasDiskBackupSnapshotViewWithDefaults instantiates a new ApiAtlasDiskBackupSnapshotView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ApiAtlasDiskBackupSnapshotView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiAtlasDiskBackupSnapshotView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiAtlasDiskBackupSnapshotView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ApiAtlasDiskBackupSnapshotView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiAtlasDiskBackupSnapshotView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiAtlasDiskBackupSnapshotView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiAtlasDiskBackupSnapshotView) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiAtlasDiskBackupSnapshotView) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiAtlasDiskBackupSnapshotView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFrequencyType

`func (o *ApiAtlasDiskBackupSnapshotView) GetFrequencyType() string`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetFrequencyTypeOk() (*string, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *ApiAtlasDiskBackupSnapshotView) SetFrequencyType(v string)`

SetFrequencyType sets FrequencyType field to given value.

### HasFrequencyType

`func (o *ApiAtlasDiskBackupSnapshotView) HasFrequencyType() bool`

HasFrequencyType returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasDiskBackupSnapshotView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasDiskBackupSnapshotView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasDiskBackupSnapshotView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasDiskBackupSnapshotView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasDiskBackupSnapshotView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetMasterKeyUUID

`func (o *ApiAtlasDiskBackupSnapshotView) GetMasterKeyUUID() string`

GetMasterKeyUUID returns the MasterKeyUUID field if non-nil, zero value otherwise.

### GetMasterKeyUUIDOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetMasterKeyUUIDOk() (*string, bool)`

GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyUUID

`func (o *ApiAtlasDiskBackupSnapshotView) SetMasterKeyUUID(v string)`

SetMasterKeyUUID sets MasterKeyUUID field to given value.

### HasMasterKeyUUID

`func (o *ApiAtlasDiskBackupSnapshotView) HasMasterKeyUUID() bool`

HasMasterKeyUUID returns a boolean if a field has been set.

### GetMongodVersion

`func (o *ApiAtlasDiskBackupSnapshotView) GetMongodVersion() string`

GetMongodVersion returns the MongodVersion field if non-nil, zero value otherwise.

### GetMongodVersionOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetMongodVersionOk() (*string, bool)`

GetMongodVersionOk returns a tuple with the MongodVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodVersion

`func (o *ApiAtlasDiskBackupSnapshotView) SetMongodVersion(v string)`

SetMongodVersion sets MongodVersion field to given value.

### HasMongodVersion

`func (o *ApiAtlasDiskBackupSnapshotView) HasMongodVersion() bool`

HasMongodVersion returns a boolean if a field has been set.

### GetPolicyItems

`func (o *ApiAtlasDiskBackupSnapshotView) GetPolicyItems() []string`

GetPolicyItems returns the PolicyItems field if non-nil, zero value otherwise.

### GetPolicyItemsOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetPolicyItemsOk() (*[]string, bool)`

GetPolicyItemsOk returns a tuple with the PolicyItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyItems

`func (o *ApiAtlasDiskBackupSnapshotView) SetPolicyItems(v []string)`

SetPolicyItems sets PolicyItems field to given value.

### HasPolicyItems

`func (o *ApiAtlasDiskBackupSnapshotView) HasPolicyItems() bool`

HasPolicyItems returns a boolean if a field has been set.

### GetSnapshotType

`func (o *ApiAtlasDiskBackupSnapshotView) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *ApiAtlasDiskBackupSnapshotView) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *ApiAtlasDiskBackupSnapshotView) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetStatus

`func (o *ApiAtlasDiskBackupSnapshotView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAtlasDiskBackupSnapshotView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAtlasDiskBackupSnapshotView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageSizeBytes

`func (o *ApiAtlasDiskBackupSnapshotView) GetStorageSizeBytes() int64`

GetStorageSizeBytes returns the StorageSizeBytes field if non-nil, zero value otherwise.

### GetStorageSizeBytesOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetStorageSizeBytesOk() (*int64, bool)`

GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeBytes

`func (o *ApiAtlasDiskBackupSnapshotView) SetStorageSizeBytes(v int64)`

SetStorageSizeBytes sets StorageSizeBytes field to given value.

### HasStorageSizeBytes

`func (o *ApiAtlasDiskBackupSnapshotView) HasStorageSizeBytes() bool`

HasStorageSizeBytes returns a boolean if a field has been set.

### GetType

`func (o *ApiAtlasDiskBackupSnapshotView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAtlasDiskBackupSnapshotView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAtlasDiskBackupSnapshotView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiAtlasDiskBackupSnapshotView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


