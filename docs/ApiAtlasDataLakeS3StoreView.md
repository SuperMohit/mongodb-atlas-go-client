# ApiAtlasDataLakeS3StoreView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalStorageClasses** | Pointer to **[]string** | Collection of AWS S3 [storage classes](https://aws.amazon.com/s3/storage-classes/). Atlas Data Lake includes the files in these storage classes in the query results. | [optional] 
**Bucket** | Pointer to **string** | Human-readable label that identifies the AWS S3 bucket. This label must exactly match the name of an S3 bucket that the data lake can access with the configured AWS Identity and Access Management (IAM) credentials. | [optional] 
**Delimiter** | Pointer to **string** | The delimiter that separates **databases.[n].collections.[n].dataSources.[n].path** segments in the data store. MongoDB Cloud uses the delimiter to efficiently traverse S3 buckets with a hierarchical directory structure. You can specify any character supported by the S3 object keys as the delimiter. For example, you can specify an underscore (_) or a plus sign (+) or multiple characters, such as double underscores (__) as the delimiter. If omitted, defaults to &#x60;/&#x60;. | [optional] 
**IncludeTags** | Pointer to **bool** | Flag that indicates whether to use S3 tags on the files in the given path as additional partition attributes. If set to &#x60;true&#x60;, data lake adds the S3 tags as additional partition attributes and adds new top-level BSON elements associating each tag to each document. | [optional] [default to false]
**Prefix** | Pointer to **string** | Prefix that MongoDB Cloud applies when searching for files in the S3 bucket. The data store prepends the value of prefix to the **databases.[n].collections.[n].dataSources.[n].path** to create the full path for files to ingest. If omitted, MongoDB Cloud searches all files from the root of the S3 bucket. | [optional] 
**Public** | Pointer to **bool** | Flag that indicates whether the bucket is public. If set to &#x60;true&#x60;, MongoDB Cloud doesn&#39;t use the configured AWS Identity and Access Management (IAM) role to access the S3 bucket. If set to &#x60;false&#x60;, the configured AWS IAM role must include permissions to access the S3 bucket. | [optional] [default to false]
**Region** | Pointer to **string** |  Physical location where MongoDB Cloud deploys your AWS-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts. | [optional] 

## Methods

### NewApiAtlasDataLakeS3StoreView

`func NewApiAtlasDataLakeS3StoreView() *ApiAtlasDataLakeS3StoreView`

NewApiAtlasDataLakeS3StoreView instantiates a new ApiAtlasDataLakeS3StoreView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDataLakeS3StoreViewWithDefaults

`func NewApiAtlasDataLakeS3StoreViewWithDefaults() *ApiAtlasDataLakeS3StoreView`

NewApiAtlasDataLakeS3StoreViewWithDefaults instantiates a new ApiAtlasDataLakeS3StoreView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalStorageClasses

`func (o *ApiAtlasDataLakeS3StoreView) GetAdditionalStorageClasses() []string`

GetAdditionalStorageClasses returns the AdditionalStorageClasses field if non-nil, zero value otherwise.

### GetAdditionalStorageClassesOk

`func (o *ApiAtlasDataLakeS3StoreView) GetAdditionalStorageClassesOk() (*[]string, bool)`

GetAdditionalStorageClassesOk returns a tuple with the AdditionalStorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalStorageClasses

`func (o *ApiAtlasDataLakeS3StoreView) SetAdditionalStorageClasses(v []string)`

SetAdditionalStorageClasses sets AdditionalStorageClasses field to given value.

### HasAdditionalStorageClasses

`func (o *ApiAtlasDataLakeS3StoreView) HasAdditionalStorageClasses() bool`

HasAdditionalStorageClasses returns a boolean if a field has been set.

### GetBucket

`func (o *ApiAtlasDataLakeS3StoreView) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ApiAtlasDataLakeS3StoreView) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ApiAtlasDataLakeS3StoreView) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ApiAtlasDataLakeS3StoreView) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetDelimiter

`func (o *ApiAtlasDataLakeS3StoreView) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *ApiAtlasDataLakeS3StoreView) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *ApiAtlasDataLakeS3StoreView) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *ApiAtlasDataLakeS3StoreView) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetIncludeTags

`func (o *ApiAtlasDataLakeS3StoreView) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *ApiAtlasDataLakeS3StoreView) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *ApiAtlasDataLakeS3StoreView) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *ApiAtlasDataLakeS3StoreView) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetPrefix

`func (o *ApiAtlasDataLakeS3StoreView) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ApiAtlasDataLakeS3StoreView) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ApiAtlasDataLakeS3StoreView) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ApiAtlasDataLakeS3StoreView) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPublic

`func (o *ApiAtlasDataLakeS3StoreView) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ApiAtlasDataLakeS3StoreView) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ApiAtlasDataLakeS3StoreView) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *ApiAtlasDataLakeS3StoreView) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRegion

`func (o *ApiAtlasDataLakeS3StoreView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ApiAtlasDataLakeS3StoreView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ApiAtlasDataLakeS3StoreView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ApiAtlasDataLakeS3StoreView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


