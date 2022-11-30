# ApiAtlasDiskBackupExportJobView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]ApiAtlasDiskBackupBaseRestoreMemberView**](ApiAtlasDiskBackupBaseRestoreMemberView.md) | Information on the export job for each replica set in the sharded cluster. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Date and time when someone created this export job. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**CustomData** | Pointer to [**[]ApiAtlasLabelView**](ApiAtlasLabelView.md) | Collection of key-value pairs that represent custom data for the metadata file that MongoDB Cloud uploads to the bucket when the export job finishes. | [optional] 
**DeliveryUrl** | Pointer to **[]string** | One or more Uniform Resource Locators (URLs) that point to the compressed snapshot files for manual download. MongoDB Cloud returns this parameter when &#x60;\&quot;deliveryType\&quot; : \&quot;download\&quot;&#x60;. | [optional] [readonly] 
**ExportBucketId** | **string** | Unique 24-hexadecimal character string that identifies the AWS bucket to which MongoDB Cloud exports the Cloud Backup snapshot. | [readonly] 
**ExportStatus** | Pointer to [**ApiExportStatusView**](ApiExportStatusView.md) |  | [optional] 
**FinishedAt** | Pointer to **time.Time** | Date and time when this export job completed. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the restore job. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Prefix** | Pointer to **string** | Full path on the cloud provider bucket to the folder where the snapshot is exported. | [optional] [readonly] 
**SnapshotId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the snapshot. | [optional] 
**State** | Pointer to **string** | State of the export job. | [optional] [readonly] 

## Methods

### NewApiAtlasDiskBackupExportJobView

`func NewApiAtlasDiskBackupExportJobView(exportBucketId string, links []Link, ) *ApiAtlasDiskBackupExportJobView`

NewApiAtlasDiskBackupExportJobView instantiates a new ApiAtlasDiskBackupExportJobView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDiskBackupExportJobViewWithDefaults

`func NewApiAtlasDiskBackupExportJobViewWithDefaults() *ApiAtlasDiskBackupExportJobView`

NewApiAtlasDiskBackupExportJobViewWithDefaults instantiates a new ApiAtlasDiskBackupExportJobView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *ApiAtlasDiskBackupExportJobView) GetComponents() []ApiAtlasDiskBackupBaseRestoreMemberView`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ApiAtlasDiskBackupExportJobView) GetComponentsOk() (*[]ApiAtlasDiskBackupBaseRestoreMemberView, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ApiAtlasDiskBackupExportJobView) SetComponents(v []ApiAtlasDiskBackupBaseRestoreMemberView)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ApiAtlasDiskBackupExportJobView) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiAtlasDiskBackupExportJobView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiAtlasDiskBackupExportJobView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiAtlasDiskBackupExportJobView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiAtlasDiskBackupExportJobView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomData

`func (o *ApiAtlasDiskBackupExportJobView) GetCustomData() []ApiAtlasLabelView`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *ApiAtlasDiskBackupExportJobView) GetCustomDataOk() (*[]ApiAtlasLabelView, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *ApiAtlasDiskBackupExportJobView) SetCustomData(v []ApiAtlasLabelView)`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *ApiAtlasDiskBackupExportJobView) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetDeliveryUrl

`func (o *ApiAtlasDiskBackupExportJobView) GetDeliveryUrl() []string`

GetDeliveryUrl returns the DeliveryUrl field if non-nil, zero value otherwise.

### GetDeliveryUrlOk

`func (o *ApiAtlasDiskBackupExportJobView) GetDeliveryUrlOk() (*[]string, bool)`

GetDeliveryUrlOk returns a tuple with the DeliveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryUrl

`func (o *ApiAtlasDiskBackupExportJobView) SetDeliveryUrl(v []string)`

SetDeliveryUrl sets DeliveryUrl field to given value.

### HasDeliveryUrl

`func (o *ApiAtlasDiskBackupExportJobView) HasDeliveryUrl() bool`

HasDeliveryUrl returns a boolean if a field has been set.

### GetExportBucketId

`func (o *ApiAtlasDiskBackupExportJobView) GetExportBucketId() string`

GetExportBucketId returns the ExportBucketId field if non-nil, zero value otherwise.

### GetExportBucketIdOk

`func (o *ApiAtlasDiskBackupExportJobView) GetExportBucketIdOk() (*string, bool)`

GetExportBucketIdOk returns a tuple with the ExportBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportBucketId

`func (o *ApiAtlasDiskBackupExportJobView) SetExportBucketId(v string)`

SetExportBucketId sets ExportBucketId field to given value.


### GetExportStatus

`func (o *ApiAtlasDiskBackupExportJobView) GetExportStatus() ApiExportStatusView`

GetExportStatus returns the ExportStatus field if non-nil, zero value otherwise.

### GetExportStatusOk

`func (o *ApiAtlasDiskBackupExportJobView) GetExportStatusOk() (*ApiExportStatusView, bool)`

GetExportStatusOk returns a tuple with the ExportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportStatus

`func (o *ApiAtlasDiskBackupExportJobView) SetExportStatus(v ApiExportStatusView)`

SetExportStatus sets ExportStatus field to given value.

### HasExportStatus

`func (o *ApiAtlasDiskBackupExportJobView) HasExportStatus() bool`

HasExportStatus returns a boolean if a field has been set.

### GetFinishedAt

`func (o *ApiAtlasDiskBackupExportJobView) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ApiAtlasDiskBackupExportJobView) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ApiAtlasDiskBackupExportJobView) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ApiAtlasDiskBackupExportJobView) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasDiskBackupExportJobView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasDiskBackupExportJobView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasDiskBackupExportJobView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasDiskBackupExportJobView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasDiskBackupExportJobView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasDiskBackupExportJobView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasDiskBackupExportJobView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetPrefix

`func (o *ApiAtlasDiskBackupExportJobView) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ApiAtlasDiskBackupExportJobView) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ApiAtlasDiskBackupExportJobView) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ApiAtlasDiskBackupExportJobView) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSnapshotId

`func (o *ApiAtlasDiskBackupExportJobView) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ApiAtlasDiskBackupExportJobView) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ApiAtlasDiskBackupExportJobView) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ApiAtlasDiskBackupExportJobView) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetState

`func (o *ApiAtlasDiskBackupExportJobView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiAtlasDiskBackupExportJobView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiAtlasDiskBackupExportJobView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiAtlasDiskBackupExportJobView) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


