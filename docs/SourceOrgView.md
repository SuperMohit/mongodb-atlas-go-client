# SourceOrgView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The ISO-8601-formatted timestamp when Cloud Manager or Ops Manager created the link. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Error message pertaining to the organization sync. Returns null if there are no errors. | [optional] [readonly] 
**LastSyncedAt** | Pointer to **time.Time** | The ISO-8601-formatted timestamp when Cloud Manager or Ops Manager last synced with Atlas. | [optional] [readonly] 
**SourceOrgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains the data you want to migrate. | [readonly] 
**TargetOrg** | Pointer to [**TargetOrgView**](TargetOrgView.md) |  | [optional] 

## Methods

### NewSourceOrgView

`func NewSourceOrgView(sourceOrgId string, ) *SourceOrgView`

NewSourceOrgView instantiates a new SourceOrgView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceOrgViewWithDefaults

`func NewSourceOrgViewWithDefaults() *SourceOrgView`

NewSourceOrgViewWithDefaults instantiates a new SourceOrgView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SourceOrgView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SourceOrgView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SourceOrgView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SourceOrgView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetErrorMessage

`func (o *SourceOrgView) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SourceOrgView) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SourceOrgView) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SourceOrgView) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *SourceOrgView) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *SourceOrgView) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *SourceOrgView) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *SourceOrgView) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetSourceOrgId

`func (o *SourceOrgView) GetSourceOrgId() string`

GetSourceOrgId returns the SourceOrgId field if non-nil, zero value otherwise.

### GetSourceOrgIdOk

`func (o *SourceOrgView) GetSourceOrgIdOk() (*string, bool)`

GetSourceOrgIdOk returns a tuple with the SourceOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrgId

`func (o *SourceOrgView) SetSourceOrgId(v string)`

SetSourceOrgId sets SourceOrgId field to given value.


### GetTargetOrg

`func (o *SourceOrgView) GetTargetOrg() TargetOrgView`

GetTargetOrg returns the TargetOrg field if non-nil, zero value otherwise.

### GetTargetOrgOk

`func (o *SourceOrgView) GetTargetOrgOk() (*TargetOrgView, bool)`

GetTargetOrgOk returns a tuple with the TargetOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOrg

`func (o *SourceOrgView) SetTargetOrg(v TargetOrgView)`

SetTargetOrg sets TargetOrg field to given value.

### HasTargetOrg

`func (o *SourceOrgView) HasTargetOrg() bool`

HasTargetOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


