# PaginatedDatabaseUserView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.  | [readonly] 
**Results** | [**[]ApiAtlasDatabaseUserViewBase**](ApiAtlasDatabaseUserViewBase.md) | List of returned documents that MongoDB Cloud provides when completing this request.  | [readonly] 
**TotalCount** | Pointer to **float32** | Number of documents returned in this response. | [optional] [readonly] 

## Methods

### NewPaginatedDatabaseUserView

`func NewPaginatedDatabaseUserView(links []Link, results []ApiAtlasDatabaseUserViewBase, ) *PaginatedDatabaseUserView`

NewPaginatedDatabaseUserView instantiates a new PaginatedDatabaseUserView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedDatabaseUserViewWithDefaults

`func NewPaginatedDatabaseUserViewWithDefaults() *PaginatedDatabaseUserView`

NewPaginatedDatabaseUserViewWithDefaults instantiates a new PaginatedDatabaseUserView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedDatabaseUserView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedDatabaseUserView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedDatabaseUserView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetResults

`func (o *PaginatedDatabaseUserView) GetResults() []ApiAtlasDatabaseUserViewBase`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedDatabaseUserView) GetResultsOk() (*[]ApiAtlasDatabaseUserViewBase, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedDatabaseUserView) SetResults(v []ApiAtlasDatabaseUserViewBase)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *PaginatedDatabaseUserView) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedDatabaseUserView) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedDatabaseUserView) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedDatabaseUserView) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


