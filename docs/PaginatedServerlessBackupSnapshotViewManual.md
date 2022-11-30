# PaginatedServerlessBackupSnapshotViewManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.\&quot; | [readonly] 
**Results** | [**[]ApiAtlasServerlessBackupSnapshotViewManual**](ApiAtlasServerlessBackupSnapshotViewManual.md) | List of returned documents that MongoDB Cloud provides when completing this request. | [readonly] 
**TotalCount** | **float32** | Number of documents returned in this response. | [readonly] 

## Methods

### NewPaginatedServerlessBackupSnapshotViewManual

`func NewPaginatedServerlessBackupSnapshotViewManual(links []Link, results []ApiAtlasServerlessBackupSnapshotViewManual, totalCount float32, ) *PaginatedServerlessBackupSnapshotViewManual`

NewPaginatedServerlessBackupSnapshotViewManual instantiates a new PaginatedServerlessBackupSnapshotViewManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedServerlessBackupSnapshotViewManualWithDefaults

`func NewPaginatedServerlessBackupSnapshotViewManualWithDefaults() *PaginatedServerlessBackupSnapshotViewManual`

NewPaginatedServerlessBackupSnapshotViewManualWithDefaults instantiates a new PaginatedServerlessBackupSnapshotViewManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedServerlessBackupSnapshotViewManual) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedServerlessBackupSnapshotViewManual) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedServerlessBackupSnapshotViewManual) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetResults

`func (o *PaginatedServerlessBackupSnapshotViewManual) GetResults() []ApiAtlasServerlessBackupSnapshotViewManual`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedServerlessBackupSnapshotViewManual) GetResultsOk() (*[]ApiAtlasServerlessBackupSnapshotViewManual, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedServerlessBackupSnapshotViewManual) SetResults(v []ApiAtlasServerlessBackupSnapshotViewManual)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *PaginatedServerlessBackupSnapshotViewManual) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedServerlessBackupSnapshotViewManual) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedServerlessBackupSnapshotViewManual) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


