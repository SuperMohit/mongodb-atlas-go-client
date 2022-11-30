# PaginatedServerlessBackupRestoreJobViewManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Results** | [**[]ApiAtlasServerlessBackupRestoreJobViewManual**](ApiAtlasServerlessBackupRestoreJobViewManual.md) | List of returned documents that MongoDB Cloud provides when completing this request. | [readonly] 
**TotalCount** | **float32** | Number of documents returned in this response. | [readonly] 

## Methods

### NewPaginatedServerlessBackupRestoreJobViewManual

`func NewPaginatedServerlessBackupRestoreJobViewManual(links []Link, results []ApiAtlasServerlessBackupRestoreJobViewManual, totalCount float32, ) *PaginatedServerlessBackupRestoreJobViewManual`

NewPaginatedServerlessBackupRestoreJobViewManual instantiates a new PaginatedServerlessBackupRestoreJobViewManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedServerlessBackupRestoreJobViewManualWithDefaults

`func NewPaginatedServerlessBackupRestoreJobViewManualWithDefaults() *PaginatedServerlessBackupRestoreJobViewManual`

NewPaginatedServerlessBackupRestoreJobViewManualWithDefaults instantiates a new PaginatedServerlessBackupRestoreJobViewManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedServerlessBackupRestoreJobViewManual) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedServerlessBackupRestoreJobViewManual) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedServerlessBackupRestoreJobViewManual) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetResults

`func (o *PaginatedServerlessBackupRestoreJobViewManual) GetResults() []ApiAtlasServerlessBackupRestoreJobViewManual`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedServerlessBackupRestoreJobViewManual) GetResultsOk() (*[]ApiAtlasServerlessBackupRestoreJobViewManual, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedServerlessBackupRestoreJobViewManual) SetResults(v []ApiAtlasServerlessBackupRestoreJobViewManual)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *PaginatedServerlessBackupRestoreJobViewManual) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedServerlessBackupRestoreJobViewManual) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedServerlessBackupRestoreJobViewManual) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


