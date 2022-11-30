# ApiAtlasClusterStatusView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeStatus** | Pointer to **string** | State of cluster at the time of this request. | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 

## Methods

### NewApiAtlasClusterStatusView

`func NewApiAtlasClusterStatusView(links []Link, ) *ApiAtlasClusterStatusView`

NewApiAtlasClusterStatusView instantiates a new ApiAtlasClusterStatusView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasClusterStatusViewWithDefaults

`func NewApiAtlasClusterStatusViewWithDefaults() *ApiAtlasClusterStatusView`

NewApiAtlasClusterStatusViewWithDefaults instantiates a new ApiAtlasClusterStatusView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeStatus

`func (o *ApiAtlasClusterStatusView) GetChangeStatus() string`

GetChangeStatus returns the ChangeStatus field if non-nil, zero value otherwise.

### GetChangeStatusOk

`func (o *ApiAtlasClusterStatusView) GetChangeStatusOk() (*string, bool)`

GetChangeStatusOk returns a tuple with the ChangeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeStatus

`func (o *ApiAtlasClusterStatusView) SetChangeStatus(v string)`

SetChangeStatus sets ChangeStatus field to given value.

### HasChangeStatus

`func (o *ApiAtlasClusterStatusView) HasChangeStatus() bool`

HasChangeStatus returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasClusterStatusView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasClusterStatusView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasClusterStatusView) SetLinks(v []Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


