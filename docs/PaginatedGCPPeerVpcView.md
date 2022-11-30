# PaginatedGCPPeerVpcView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Results** | [**[]ApiAtlasGCPPeerVpcView**](ApiAtlasGCPPeerVpcView.md) | List of returned documents that MongoDB Cloud providers when completing this request. | [readonly] 
**TotalCount** | **int32** | Number of documents returned in this response. | [readonly] 

## Methods

### NewPaginatedGCPPeerVpcView

`func NewPaginatedGCPPeerVpcView(links []Link, results []ApiAtlasGCPPeerVpcView, totalCount int32, ) *PaginatedGCPPeerVpcView`

NewPaginatedGCPPeerVpcView instantiates a new PaginatedGCPPeerVpcView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedGCPPeerVpcViewWithDefaults

`func NewPaginatedGCPPeerVpcViewWithDefaults() *PaginatedGCPPeerVpcView`

NewPaginatedGCPPeerVpcViewWithDefaults instantiates a new PaginatedGCPPeerVpcView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedGCPPeerVpcView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedGCPPeerVpcView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedGCPPeerVpcView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetResults

`func (o *PaginatedGCPPeerVpcView) GetResults() []ApiAtlasGCPPeerVpcView`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedGCPPeerVpcView) GetResultsOk() (*[]ApiAtlasGCPPeerVpcView, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedGCPPeerVpcView) SetResults(v []ApiAtlasGCPPeerVpcView)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *PaginatedGCPPeerVpcView) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedGCPPeerVpcView) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedGCPPeerVpcView) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


