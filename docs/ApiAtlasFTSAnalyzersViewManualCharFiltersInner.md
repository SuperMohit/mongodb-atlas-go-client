# ApiAtlasFTSAnalyzersViewManualCharFiltersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoredTags** | Pointer to **[]string** | The HTML tags that you want to exclude from filtering. | [optional] 
**Type** | **string** | Human-readable label that identifies this character filter type. | 
**Mappings** | [**CharFiltermappingMappings**](CharFiltermappingMappings.md) |  | 

## Methods

### NewApiAtlasFTSAnalyzersViewManualCharFiltersInner

`func NewApiAtlasFTSAnalyzersViewManualCharFiltersInner(type_ string, mappings CharFiltermappingMappings, ) *ApiAtlasFTSAnalyzersViewManualCharFiltersInner`

NewApiAtlasFTSAnalyzersViewManualCharFiltersInner instantiates a new ApiAtlasFTSAnalyzersViewManualCharFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasFTSAnalyzersViewManualCharFiltersInnerWithDefaults

`func NewApiAtlasFTSAnalyzersViewManualCharFiltersInnerWithDefaults() *ApiAtlasFTSAnalyzersViewManualCharFiltersInner`

NewApiAtlasFTSAnalyzersViewManualCharFiltersInnerWithDefaults instantiates a new ApiAtlasFTSAnalyzersViewManualCharFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoredTags

`func (o *ApiAtlasFTSAnalyzersViewManualCharFiltersInner) GetIgnoredTags() []string`

GetIgnoredTags returns the IgnoredTags field if non-nil, zero value otherwise.

### GetIgnoredTagsOk

`func (o *ApiAtlasFTSAnalyzersViewManualCharFiltersInner) GetIgnoredTagsOk() (*[]string, bool)`

GetIgnoredTagsOk returns a tuple with the IgnoredTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredTags

`func (o *ApiAtlasFTSAnalyzersViewManualCharFiltersInner) SetIgnoredTags(v []string)`

SetIgnoredTags sets IgnoredTags field to given value.

### HasIgnoredTags

`func (o *ApiAtlasFTSAnalyzersViewManualCharFiltersInner) HasIgnoredTags() bool`

HasIgnoredTags returns a boolean if a field has been set.

### GetType

`func (o *ApiAtlasFTSAnalyzersViewManualCharFiltersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAtlasFTSAnalyzersViewManualCharFiltersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAtlasFTSAnalyzersViewManualCharFiltersInner) SetType(v string)`

SetType sets Type field to given value.


### GetMappings

`func (o *ApiAtlasFTSAnalyzersViewManualCharFiltersInner) GetMappings() CharFiltermappingMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *ApiAtlasFTSAnalyzersViewManualCharFiltersInner) GetMappingsOk() (*CharFiltermappingMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *ApiAtlasFTSAnalyzersViewManualCharFiltersInner) SetMappings(v CharFiltermappingMappings)`

SetMappings sets Mappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


