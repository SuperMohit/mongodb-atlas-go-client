# ApiIndexRequestView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collation** | Pointer to [**Collation**](Collation.md) |  | [optional] 
**Collection** | **string** | Human-readable label of the collection for which MongoDB Cloud creates an index. | 
**Db** | **string** | Human-readable label of the database that holds the collection on which MongoDB Cloud creates an index. | 
**Keys** | **[]map[string]string** | List that contains one or more objects that describe the parameters that you want to index. | 
**Options** | Pointer to [**IndexOptions**](IndexOptions.md) |  | [optional] 

## Methods

### NewApiIndexRequestView

`func NewApiIndexRequestView(collection string, db string, keys []map[string]string, ) *ApiIndexRequestView`

NewApiIndexRequestView instantiates a new ApiIndexRequestView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiIndexRequestViewWithDefaults

`func NewApiIndexRequestViewWithDefaults() *ApiIndexRequestView`

NewApiIndexRequestViewWithDefaults instantiates a new ApiIndexRequestView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollation

`func (o *ApiIndexRequestView) GetCollation() Collation`

GetCollation returns the Collation field if non-nil, zero value otherwise.

### GetCollationOk

`func (o *ApiIndexRequestView) GetCollationOk() (*Collation, bool)`

GetCollationOk returns a tuple with the Collation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollation

`func (o *ApiIndexRequestView) SetCollation(v Collation)`

SetCollation sets Collation field to given value.

### HasCollation

`func (o *ApiIndexRequestView) HasCollation() bool`

HasCollation returns a boolean if a field has been set.

### GetCollection

`func (o *ApiIndexRequestView) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ApiIndexRequestView) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ApiIndexRequestView) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetDb

`func (o *ApiIndexRequestView) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *ApiIndexRequestView) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *ApiIndexRequestView) SetDb(v string)`

SetDb sets Db field to given value.


### GetKeys

`func (o *ApiIndexRequestView) GetKeys() []map[string]string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ApiIndexRequestView) GetKeysOk() (*[]map[string]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ApiIndexRequestView) SetKeys(v []map[string]string)`

SetKeys sets Keys field to given value.


### GetOptions

`func (o *ApiIndexRequestView) GetOptions() IndexOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApiIndexRequestView) GetOptionsOk() (*IndexOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApiIndexRequestView) SetOptions(v IndexOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApiIndexRequestView) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


