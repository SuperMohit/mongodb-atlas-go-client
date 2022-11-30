# ApiOrganizationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization. | [optional] [readonly] 
**IsDeleted** | Pointer to **bool** | Flag that indicates whether this organization has been deleted. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Name** | **string** | Human-readable label that identifies the organization. | 

## Methods

### NewApiOrganizationView

`func NewApiOrganizationView(links []Link, name string, ) *ApiOrganizationView`

NewApiOrganizationView instantiates a new ApiOrganizationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOrganizationViewWithDefaults

`func NewApiOrganizationViewWithDefaults() *ApiOrganizationView`

NewApiOrganizationViewWithDefaults instantiates a new ApiOrganizationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiOrganizationView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiOrganizationView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiOrganizationView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiOrganizationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *ApiOrganizationView) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ApiOrganizationView) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ApiOrganizationView) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ApiOrganizationView) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLinks

`func (o *ApiOrganizationView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiOrganizationView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiOrganizationView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetName

`func (o *ApiOrganizationView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiOrganizationView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiOrganizationView) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


