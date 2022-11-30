# ApiAtlasDiskBackupExportJobRequestView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomData** | Pointer to [**[]ApiAtlasLabelView**](ApiAtlasLabelView.md) | Collection of key-value pairs that represent custom data to add to the metadata file that MongoDB Cloud uploads to the bucket when the export job finishes. | [optional] 
**ExportBucketId** | **string** | Unique 24-hexadecimal character string that identifies the AWS bucket to which MongoDB Cloud exports the Cloud Backup snapshot. | 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**SnapshotId** | **string** | Unique 24-hexadecimal character string that identifies the Cloud Backup snasphot to export. | 

## Methods

### NewApiAtlasDiskBackupExportJobRequestView

`func NewApiAtlasDiskBackupExportJobRequestView(exportBucketId string, links []Link, snapshotId string, ) *ApiAtlasDiskBackupExportJobRequestView`

NewApiAtlasDiskBackupExportJobRequestView instantiates a new ApiAtlasDiskBackupExportJobRequestView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDiskBackupExportJobRequestViewWithDefaults

`func NewApiAtlasDiskBackupExportJobRequestViewWithDefaults() *ApiAtlasDiskBackupExportJobRequestView`

NewApiAtlasDiskBackupExportJobRequestViewWithDefaults instantiates a new ApiAtlasDiskBackupExportJobRequestView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomData

`func (o *ApiAtlasDiskBackupExportJobRequestView) GetCustomData() []ApiAtlasLabelView`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *ApiAtlasDiskBackupExportJobRequestView) GetCustomDataOk() (*[]ApiAtlasLabelView, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *ApiAtlasDiskBackupExportJobRequestView) SetCustomData(v []ApiAtlasLabelView)`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *ApiAtlasDiskBackupExportJobRequestView) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetExportBucketId

`func (o *ApiAtlasDiskBackupExportJobRequestView) GetExportBucketId() string`

GetExportBucketId returns the ExportBucketId field if non-nil, zero value otherwise.

### GetExportBucketIdOk

`func (o *ApiAtlasDiskBackupExportJobRequestView) GetExportBucketIdOk() (*string, bool)`

GetExportBucketIdOk returns a tuple with the ExportBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportBucketId

`func (o *ApiAtlasDiskBackupExportJobRequestView) SetExportBucketId(v string)`

SetExportBucketId sets ExportBucketId field to given value.


### GetLinks

`func (o *ApiAtlasDiskBackupExportJobRequestView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasDiskBackupExportJobRequestView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasDiskBackupExportJobRequestView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetSnapshotId

`func (o *ApiAtlasDiskBackupExportJobRequestView) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ApiAtlasDiskBackupExportJobRequestView) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ApiAtlasDiskBackupExportJobRequestView) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


