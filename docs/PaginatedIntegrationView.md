# PaginatedIntegrationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Results** | [**[]ApiIntegrationViewManual**](ApiIntegrationViewManual.md) | List of returned documents that MongoDB Cloud provides when completing this request. | [readonly] 
**TotalCount** | Pointer to **float32** | Number of documents returned in this response. | [optional] [readonly] 

## Methods

### NewPaginatedIntegrationView

`func NewPaginatedIntegrationView(links []Link, results []ApiIntegrationViewManual, ) *PaginatedIntegrationView`

NewPaginatedIntegrationView instantiates a new PaginatedIntegrationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedIntegrationViewWithDefaults

`func NewPaginatedIntegrationViewWithDefaults() *PaginatedIntegrationView`

NewPaginatedIntegrationViewWithDefaults instantiates a new PaginatedIntegrationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedIntegrationView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedIntegrationView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedIntegrationView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetResults

`func (o *PaginatedIntegrationView) GetResults() []ApiIntegrationViewManual`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedIntegrationView) GetResultsOk() (*[]ApiIntegrationViewManual, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedIntegrationView) SetResults(v []ApiIntegrationViewManual)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *PaginatedIntegrationView) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedIntegrationView) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedIntegrationView) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedIntegrationView) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


