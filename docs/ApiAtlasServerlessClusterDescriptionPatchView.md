# ApiAtlasServerlessClusterDescriptionPatchView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]

## Methods

### NewApiAtlasServerlessClusterDescriptionPatchView

`func NewApiAtlasServerlessClusterDescriptionPatchView() *ApiAtlasServerlessClusterDescriptionPatchView`

NewApiAtlasServerlessClusterDescriptionPatchView instantiates a new ApiAtlasServerlessClusterDescriptionPatchView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasServerlessClusterDescriptionPatchViewWithDefaults

`func NewApiAtlasServerlessClusterDescriptionPatchViewWithDefaults() *ApiAtlasServerlessClusterDescriptionPatchView`

NewApiAtlasServerlessClusterDescriptionPatchViewWithDefaults instantiates a new ApiAtlasServerlessClusterDescriptionPatchView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerminationProtectionEnabled

`func (o *ApiAtlasServerlessClusterDescriptionPatchView) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *ApiAtlasServerlessClusterDescriptionPatchView) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *ApiAtlasServerlessClusterDescriptionPatchView) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *ApiAtlasServerlessClusterDescriptionPatchView) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


