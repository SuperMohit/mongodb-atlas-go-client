# ApiMatcherView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | Name of the parameter in the target object that MongoDB Cloud checks. The parameter must match all rules for MongoDB Cloud to check for alert configurations. | [optional] 
**Operator** | Pointer to **string** | Comparison operator to apply when checking the current metric value against **matcher[n].value**. | [optional] 
**Value** | Pointer to **string** | Value to match or exceed using the specified **matchers.operator**. | [optional] 

## Methods

### NewApiMatcherView

`func NewApiMatcherView() *ApiMatcherView`

NewApiMatcherView instantiates a new ApiMatcherView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMatcherViewWithDefaults

`func NewApiMatcherViewWithDefaults() *ApiMatcherView`

NewApiMatcherViewWithDefaults instantiates a new ApiMatcherView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *ApiMatcherView) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ApiMatcherView) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ApiMatcherView) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ApiMatcherView) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetOperator

`func (o *ApiMatcherView) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ApiMatcherView) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ApiMatcherView) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ApiMatcherView) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *ApiMatcherView) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiMatcherView) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiMatcherView) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApiMatcherView) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


