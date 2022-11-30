# AutoExportPolicyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportBucketId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the AWS Bucket. | [optional] 
**FrequencyType** | Pointer to **string** | Human-readable label that indicates the rate at which the export policy item occurs. | [optional] 

## Methods

### NewAutoExportPolicyView

`func NewAutoExportPolicyView() *AutoExportPolicyView`

NewAutoExportPolicyView instantiates a new AutoExportPolicyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoExportPolicyViewWithDefaults

`func NewAutoExportPolicyViewWithDefaults() *AutoExportPolicyView`

NewAutoExportPolicyViewWithDefaults instantiates a new AutoExportPolicyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportBucketId

`func (o *AutoExportPolicyView) GetExportBucketId() string`

GetExportBucketId returns the ExportBucketId field if non-nil, zero value otherwise.

### GetExportBucketIdOk

`func (o *AutoExportPolicyView) GetExportBucketIdOk() (*string, bool)`

GetExportBucketIdOk returns a tuple with the ExportBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportBucketId

`func (o *AutoExportPolicyView) SetExportBucketId(v string)`

SetExportBucketId sets ExportBucketId field to given value.

### HasExportBucketId

`func (o *AutoExportPolicyView) HasExportBucketId() bool`

HasExportBucketId returns a boolean if a field has been set.

### GetFrequencyType

`func (o *AutoExportPolicyView) GetFrequencyType() string`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *AutoExportPolicyView) GetFrequencyTypeOk() (*string, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *AutoExportPolicyView) SetFrequencyType(v string)`

SetFrequencyType sets FrequencyType field to given value.

### HasFrequencyType

`func (o *AutoExportPolicyView) HasFrequencyType() bool`

HasFrequencyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


