# ApiNamespaceObjView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** | Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. | [optional] [readonly] 
**Type** | Pointer to **string** | Human-readable label that identifies the type of namespace. | [optional] [readonly] [default to "COLLECTION"]

## Methods

### NewApiNamespaceObjView

`func NewApiNamespaceObjView() *ApiNamespaceObjView`

NewApiNamespaceObjView instantiates a new ApiNamespaceObjView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNamespaceObjViewWithDefaults

`func NewApiNamespaceObjViewWithDefaults() *ApiNamespaceObjView`

NewApiNamespaceObjViewWithDefaults instantiates a new ApiNamespaceObjView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ApiNamespaceObjView) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiNamespaceObjView) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiNamespaceObjView) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ApiNamespaceObjView) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetType

`func (o *ApiNamespaceObjView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiNamespaceObjView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiNamespaceObjView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiNamespaceObjView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


