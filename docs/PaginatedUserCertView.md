# PaginatedUserCertView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Results** | [**[]ApiAtlasUserCertView**](ApiAtlasUserCertView.md) | List of returned documents that MongoDB Cloud providers when completing this request. | [readonly] 
**TotalCount** | **int32** | Number of documents returned in this response. | [readonly] 

## Methods

### NewPaginatedUserCertView

`func NewPaginatedUserCertView(links []Link, results []ApiAtlasUserCertView, totalCount int32, ) *PaginatedUserCertView`

NewPaginatedUserCertView instantiates a new PaginatedUserCertView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedUserCertViewWithDefaults

`func NewPaginatedUserCertViewWithDefaults() *PaginatedUserCertView`

NewPaginatedUserCertViewWithDefaults instantiates a new PaginatedUserCertView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedUserCertView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedUserCertView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedUserCertView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetResults

`func (o *PaginatedUserCertView) GetResults() []ApiAtlasUserCertView`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedUserCertView) GetResultsOk() (*[]ApiAtlasUserCertView, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedUserCertView) SetResults(v []ApiAtlasUserCertView)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *PaginatedUserCertView) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedUserCertView) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedUserCertView) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


