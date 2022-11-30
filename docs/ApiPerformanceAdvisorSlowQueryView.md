# ApiPerformanceAdvisorSlowQueryView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | Pointer to **string** | Text of the MongoDB log related to this slow query. | [optional] [readonly] 
**Namespace** | Pointer to **string** | Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. | [optional] [readonly] 

## Methods

### NewApiPerformanceAdvisorSlowQueryView

`func NewApiPerformanceAdvisorSlowQueryView() *ApiPerformanceAdvisorSlowQueryView`

NewApiPerformanceAdvisorSlowQueryView instantiates a new ApiPerformanceAdvisorSlowQueryView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPerformanceAdvisorSlowQueryViewWithDefaults

`func NewApiPerformanceAdvisorSlowQueryViewWithDefaults() *ApiPerformanceAdvisorSlowQueryView`

NewApiPerformanceAdvisorSlowQueryViewWithDefaults instantiates a new ApiPerformanceAdvisorSlowQueryView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *ApiPerformanceAdvisorSlowQueryView) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ApiPerformanceAdvisorSlowQueryView) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ApiPerformanceAdvisorSlowQueryView) SetLine(v string)`

SetLine sets Line field to given value.

### HasLine

`func (o *ApiPerformanceAdvisorSlowQueryView) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetNamespace

`func (o *ApiPerformanceAdvisorSlowQueryView) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiPerformanceAdvisorSlowQueryView) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiPerformanceAdvisorSlowQueryView) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ApiPerformanceAdvisorSlowQueryView) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


