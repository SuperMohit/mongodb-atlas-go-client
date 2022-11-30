# ApiNamespacesView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespaces** | Pointer to [**[]ApiNamespaceObjView**](ApiNamespaceObjView.md) | List that contains each combination of database, collection, and type on the specified host. | [optional] [readonly] 

## Methods

### NewApiNamespacesView

`func NewApiNamespacesView() *ApiNamespacesView`

NewApiNamespacesView instantiates a new ApiNamespacesView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNamespacesViewWithDefaults

`func NewApiNamespacesViewWithDefaults() *ApiNamespacesView`

NewApiNamespacesViewWithDefaults instantiates a new ApiNamespacesView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaces

`func (o *ApiNamespacesView) GetNamespaces() []ApiNamespaceObjView`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *ApiNamespacesView) GetNamespacesOk() (*[]ApiNamespaceObjView, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *ApiNamespacesView) SetNamespaces(v []ApiNamespaceObjView)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *ApiNamespacesView) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


