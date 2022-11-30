# CustomCriteriaViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | MongoDB find query that selects documents to archive. The specified query follows the syntax of the &#x60;db.collection.find(query)&#x60; command. This query can&#39;t use the empty document (&#x60;{}&#x60;) to return all documents. Set this parameter when **\&quot;criteria.type\&quot; : \&quot;CUSTOM\&quot;**. | [optional] 

## Methods

### NewCustomCriteriaViewAllOf

`func NewCustomCriteriaViewAllOf() *CustomCriteriaViewAllOf`

NewCustomCriteriaViewAllOf instantiates a new CustomCriteriaViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCriteriaViewAllOfWithDefaults

`func NewCustomCriteriaViewAllOfWithDefaults() *CustomCriteriaViewAllOf`

NewCustomCriteriaViewAllOfWithDefaults instantiates a new CustomCriteriaViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CustomCriteriaViewAllOf) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CustomCriteriaViewAllOf) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CustomCriteriaViewAllOf) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CustomCriteriaViewAllOf) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


