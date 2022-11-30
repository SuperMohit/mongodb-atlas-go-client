# ApiAtlasDiskBackupSnapshotAWSExportBucketView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the Amazon Web Services (AWS) Simple Storage Service (S3) export bucket. | [optional] [readonly] 
**BucketName** | Pointer to **string** | Human-readable label that identifies the AWS bucket that the role is authorized to access. | [optional] 
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider that stores this snapshot. | [optional] 
**IamRoleId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the AWS IAM role that MongoDB Cloud uses to access the AWS S3 bucket. | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 

## Methods

### NewApiAtlasDiskBackupSnapshotAWSExportBucketView

`func NewApiAtlasDiskBackupSnapshotAWSExportBucketView(links []Link, ) *ApiAtlasDiskBackupSnapshotAWSExportBucketView`

NewApiAtlasDiskBackupSnapshotAWSExportBucketView instantiates a new ApiAtlasDiskBackupSnapshotAWSExportBucketView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDiskBackupSnapshotAWSExportBucketViewWithDefaults

`func NewApiAtlasDiskBackupSnapshotAWSExportBucketViewWithDefaults() *ApiAtlasDiskBackupSnapshotAWSExportBucketView`

NewApiAtlasDiskBackupSnapshotAWSExportBucketViewWithDefaults instantiates a new ApiAtlasDiskBackupSnapshotAWSExportBucketView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBucketName

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetCloudProvider

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetIamRoleId

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### HasIamRoleId

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) HasIamRoleId() bool`

HasIamRoleId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasDiskBackupSnapshotAWSExportBucketView) SetLinks(v []Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


