# CustomCriteriaView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | MongoDB find query that selects documents to archive. The specified query follows the syntax of the &#x60;db.collection.find(query)&#x60; command. This query can&#39;t use the empty document (&#x60;{}&#x60;) to return all documents. Set this parameter when **\&quot;criteria.type\&quot; : \&quot;CUSTOM\&quot;**. | 

## Methods

### NewCustomCriteriaView

`func NewCustomCriteriaView(query string, ) *CustomCriteriaView`

NewCustomCriteriaView instantiates a new CustomCriteriaView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCriteriaViewWithDefaults

`func NewCustomCriteriaViewWithDefaults() *CustomCriteriaView`

NewCustomCriteriaViewWithDefaults instantiates a new CustomCriteriaView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CustomCriteriaView) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CustomCriteriaView) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CustomCriteriaView) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


