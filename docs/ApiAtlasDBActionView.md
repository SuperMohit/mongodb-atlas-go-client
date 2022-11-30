# ApiAtlasDBActionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Human-readable label that identifies the privilege action. | 
**Resources** | [**[]ApiAtlasDBResourceView**](ApiAtlasDBResourceView.md) | List of resources on which you grant the action. | 

## Methods

### NewApiAtlasDBActionView

`func NewApiAtlasDBActionView(action string, resources []ApiAtlasDBResourceView, ) *ApiAtlasDBActionView`

NewApiAtlasDBActionView instantiates a new ApiAtlasDBActionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasDBActionViewWithDefaults

`func NewApiAtlasDBActionViewWithDefaults() *ApiAtlasDBActionView`

NewApiAtlasDBActionViewWithDefaults instantiates a new ApiAtlasDBActionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ApiAtlasDBActionView) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ApiAtlasDBActionView) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ApiAtlasDBActionView) SetAction(v string)`

SetAction sets Action field to given value.


### GetResources

`func (o *ApiAtlasDBActionView) GetResources() []ApiAtlasDBResourceView`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ApiAtlasDBActionView) GetResourcesOk() (*[]ApiAtlasDBResourceView, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ApiAtlasDBActionView) SetResources(v []ApiAtlasDBResourceView)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


