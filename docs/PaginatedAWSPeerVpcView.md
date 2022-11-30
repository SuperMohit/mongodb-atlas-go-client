# PaginatedAWSPeerVpcView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Results** | [**[]ApiAtlasAWSPeerVpcView**](ApiAtlasAWSPeerVpcView.md) | List of returned documents that MongoDB Cloud providers when completing this request. | [readonly] 
**TotalCount** | **int32** | Number of documents returned in this response. | [readonly] 

## Methods

### NewPaginatedAWSPeerVpcView

`func NewPaginatedAWSPeerVpcView(links []Link, results []ApiAtlasAWSPeerVpcView, totalCount int32, ) *PaginatedAWSPeerVpcView`

NewPaginatedAWSPeerVpcView instantiates a new PaginatedAWSPeerVpcView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAWSPeerVpcViewWithDefaults

`func NewPaginatedAWSPeerVpcViewWithDefaults() *PaginatedAWSPeerVpcView`

NewPaginatedAWSPeerVpcViewWithDefaults instantiates a new PaginatedAWSPeerVpcView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedAWSPeerVpcView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedAWSPeerVpcView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedAWSPeerVpcView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetResults

`func (o *PaginatedAWSPeerVpcView) GetResults() []ApiAtlasAWSPeerVpcView`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedAWSPeerVpcView) GetResultsOk() (*[]ApiAtlasAWSPeerVpcView, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedAWSPeerVpcView) SetResults(v []ApiAtlasAWSPeerVpcView)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *PaginatedAWSPeerVpcView) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedAWSPeerVpcView) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedAWSPeerVpcView) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


