# CriteriaView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Means by which MongoDB Cloud selects data to archive. Data can be chosen using the age of the data or a MongoDB query. **DATE** selects documents to archive based on a date. **CUSTOM** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn&#39;t support **CUSTOM** when &#x60;\&quot;collectionType\&quot;: \&quot;TIMESERIES\&quot;&#x60;. | [optional] 

## Methods

### NewCriteriaView

`func NewCriteriaView() *CriteriaView`

NewCriteriaView instantiates a new CriteriaView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriteriaViewWithDefaults

`func NewCriteriaViewWithDefaults() *CriteriaView`

NewCriteriaViewWithDefaults instantiates a new CriteriaView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CriteriaView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CriteriaView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CriteriaView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CriteriaView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


