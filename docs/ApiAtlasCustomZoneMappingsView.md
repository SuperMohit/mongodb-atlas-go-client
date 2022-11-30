# ApiAtlasCustomZoneMappingsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomZoneMappings** | [**[]ApiAtlasZoneMappingView**](ApiAtlasZoneMappingView.md) | List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to the human-readable label for the desired custom zone. MongoDB Cloud maps the ISO 3166-1a2 code to the nearest geographical zone by default. Include this parameter to override the default mappings.  This parameter returns an empty object if no custom zones exist. | 

## Methods

### NewApiAtlasCustomZoneMappingsView

`func NewApiAtlasCustomZoneMappingsView(customZoneMappings []ApiAtlasZoneMappingView, ) *ApiAtlasCustomZoneMappingsView`

NewApiAtlasCustomZoneMappingsView instantiates a new ApiAtlasCustomZoneMappingsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCustomZoneMappingsViewWithDefaults

`func NewApiAtlasCustomZoneMappingsViewWithDefaults() *ApiAtlasCustomZoneMappingsView`

NewApiAtlasCustomZoneMappingsViewWithDefaults instantiates a new ApiAtlasCustomZoneMappingsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomZoneMappings

`func (o *ApiAtlasCustomZoneMappingsView) GetCustomZoneMappings() []ApiAtlasZoneMappingView`

GetCustomZoneMappings returns the CustomZoneMappings field if non-nil, zero value otherwise.

### GetCustomZoneMappingsOk

`func (o *ApiAtlasCustomZoneMappingsView) GetCustomZoneMappingsOk() (*[]ApiAtlasZoneMappingView, bool)`

GetCustomZoneMappingsOk returns a tuple with the CustomZoneMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomZoneMappings

`func (o *ApiAtlasCustomZoneMappingsView) SetCustomZoneMappings(v []ApiAtlasZoneMappingView)`

SetCustomZoneMappings sets CustomZoneMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


