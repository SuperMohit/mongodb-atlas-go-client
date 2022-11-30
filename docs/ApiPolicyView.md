# ApiPolicyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this backup policy. | [optional] 
**PolicyItems** | [**[]ApiPolicyItemView**](ApiPolicyItemView.md) | List that contains the specifications for one policy. | 

## Methods

### NewApiPolicyView

`func NewApiPolicyView(policyItems []ApiPolicyItemView, ) *ApiPolicyView`

NewApiPolicyView instantiates a new ApiPolicyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyViewWithDefaults

`func NewApiPolicyViewWithDefaults() *ApiPolicyView`

NewApiPolicyViewWithDefaults instantiates a new ApiPolicyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiPolicyView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiPolicyView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiPolicyView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiPolicyView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicyItems

`func (o *ApiPolicyView) GetPolicyItems() []ApiPolicyItemView`

GetPolicyItems returns the PolicyItems field if non-nil, zero value otherwise.

### GetPolicyItemsOk

`func (o *ApiPolicyView) GetPolicyItemsOk() (*[]ApiPolicyItemView, bool)`

GetPolicyItemsOk returns a tuple with the PolicyItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyItems

`func (o *ApiPolicyView) SetPolicyItems(v []ApiPolicyItemView)`

SetPolicyItems sets PolicyItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


