# ApiAtlasFTSMappingsViewManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dynamic** | Pointer to **bool** | Flag that indicates whether the index uses dynamic or static mappings. Required if **mappings.fields** is omitted. | [optional] [default to false]
**Fields** | Pointer to **map[string]map[string]interface{}** | One or more field specifications for the Atlas Search index. Required if **mappings.dynamic** is omitted or set to **false**. | [optional] 

## Methods

### NewApiAtlasFTSMappingsViewManual

`func NewApiAtlasFTSMappingsViewManual() *ApiAtlasFTSMappingsViewManual`

NewApiAtlasFTSMappingsViewManual instantiates a new ApiAtlasFTSMappingsViewManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasFTSMappingsViewManualWithDefaults

`func NewApiAtlasFTSMappingsViewManualWithDefaults() *ApiAtlasFTSMappingsViewManual`

NewApiAtlasFTSMappingsViewManualWithDefaults instantiates a new ApiAtlasFTSMappingsViewManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDynamic

`func (o *ApiAtlasFTSMappingsViewManual) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *ApiAtlasFTSMappingsViewManual) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *ApiAtlasFTSMappingsViewManual) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *ApiAtlasFTSMappingsViewManual) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetFields

`func (o *ApiAtlasFTSMappingsViewManual) GetFields() map[string]map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ApiAtlasFTSMappingsViewManual) GetFieldsOk() (*map[string]map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ApiAtlasFTSMappingsViewManual) SetFields(v map[string]map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *ApiAtlasFTSMappingsViewManual) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


