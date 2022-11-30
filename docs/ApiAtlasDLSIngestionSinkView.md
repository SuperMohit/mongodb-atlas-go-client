# ApiAtlasDLSIngestionSinkView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataProvider** | Pointer to **string** | Target cloud provider for this Data Lake Pipeline. | [optional] 
**MetadataRegion** | Pointer to **string** | Target cloud provider region for this Data Lake Pipeline. | [optional] 
**PartitionFields** | Pointer to [**[]ApiAtlasPartitionFieldView**](ApiAtlasPartitionFieldView.md) | Ordered fields used to physically organize data in the destination. | [optional] 

## Methods

### NewApiAtlasDLSIngestionSinkView

`func NewApiAtlasDLSIngestionSinkView() *ApiAtlasDLSIngestionSinkView`

NewApiAtlasDLSIngestionSinkView instantiates a new ApiAtlasDLSIngestionSinkView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDLSIngestionSinkViewWithDefaults

`func NewApiAtlasDLSIngestionSinkViewWithDefaults() *ApiAtlasDLSIngestionSinkView`

NewApiAtlasDLSIngestionSinkViewWithDefaults instantiates a new ApiAtlasDLSIngestionSinkView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataProvider

`func (o *ApiAtlasDLSIngestionSinkView) GetMetadataProvider() string`

GetMetadataProvider returns the MetadataProvider field if non-nil, zero value otherwise.

### GetMetadataProviderOk

`func (o *ApiAtlasDLSIngestionSinkView) GetMetadataProviderOk() (*string, bool)`

GetMetadataProviderOk returns a tuple with the MetadataProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProvider

`func (o *ApiAtlasDLSIngestionSinkView) SetMetadataProvider(v string)`

SetMetadataProvider sets MetadataProvider field to given value.

### HasMetadataProvider

`func (o *ApiAtlasDLSIngestionSinkView) HasMetadataProvider() bool`

HasMetadataProvider returns a boolean if a field has been set.

### GetMetadataRegion

`func (o *ApiAtlasDLSIngestionSinkView) GetMetadataRegion() string`

GetMetadataRegion returns the MetadataRegion field if non-nil, zero value otherwise.

### GetMetadataRegionOk

`func (o *ApiAtlasDLSIngestionSinkView) GetMetadataRegionOk() (*string, bool)`

GetMetadataRegionOk returns a tuple with the MetadataRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataRegion

`func (o *ApiAtlasDLSIngestionSinkView) SetMetadataRegion(v string)`

SetMetadataRegion sets MetadataRegion field to given value.

### HasMetadataRegion

`func (o *ApiAtlasDLSIngestionSinkView) HasMetadataRegion() bool`

HasMetadataRegion returns a boolean if a field has been set.

### GetPartitionFields

`func (o *ApiAtlasDLSIngestionSinkView) GetPartitionFields() []ApiAtlasPartitionFieldView`

GetPartitionFields returns the PartitionFields field if non-nil, zero value otherwise.

### GetPartitionFieldsOk

`func (o *ApiAtlasDLSIngestionSinkView) GetPartitionFieldsOk() (*[]ApiAtlasPartitionFieldView, bool)`

GetPartitionFieldsOk returns a tuple with the PartitionFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionFields

`func (o *ApiAtlasDLSIngestionSinkView) SetPartitionFields(v []ApiAtlasPartitionFieldView)`

SetPartitionFields sets PartitionFields field to given value.

### HasPartitionFields

`func (o *ApiAtlasDLSIngestionSinkView) HasPartitionFields() bool`

HasPartitionFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


