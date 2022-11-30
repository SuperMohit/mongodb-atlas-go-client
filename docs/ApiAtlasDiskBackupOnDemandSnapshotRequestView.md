# ApiAtlasDiskBackupOnDemandSnapshotRequestView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Human-readable phrase or sentence that explains the purpose of the snapshot. The resource returns this parameter when &#x60;\&quot;status\&quot; : \&quot;onDemand\&quot;&#x60;. | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**RetentionInDays** | Pointer to **int32** | Number of days that MongoDB Cloud should retain the on-demand snapshot. Must be at least **1** | [optional] 

## Methods

### NewApiAtlasDiskBackupOnDemandSnapshotRequestView

`func NewApiAtlasDiskBackupOnDemandSnapshotRequestView(links []Link, ) *ApiAtlasDiskBackupOnDemandSnapshotRequestView`

NewApiAtlasDiskBackupOnDemandSnapshotRequestView instantiates a new ApiAtlasDiskBackupOnDemandSnapshotRequestView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDiskBackupOnDemandSnapshotRequestViewWithDefaults

`func NewApiAtlasDiskBackupOnDemandSnapshotRequestViewWithDefaults() *ApiAtlasDiskBackupOnDemandSnapshotRequestView`

NewApiAtlasDiskBackupOnDemandSnapshotRequestViewWithDefaults instantiates a new ApiAtlasDiskBackupOnDemandSnapshotRequestView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetRetentionInDays

`func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetRetentionInDays() int32`

GetRetentionInDays returns the RetentionInDays field if non-nil, zero value otherwise.

### GetRetentionInDaysOk

`func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) GetRetentionInDaysOk() (*int32, bool)`

GetRetentionInDaysOk returns a tuple with the RetentionInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionInDays

`func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) SetRetentionInDays(v int32)`

SetRetentionInDays sets RetentionInDays field to given value.

### HasRetentionInDays

`func (o *ApiAtlasDiskBackupOnDemandSnapshotRequestView) HasRetentionInDays() bool`

HasRetentionInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


