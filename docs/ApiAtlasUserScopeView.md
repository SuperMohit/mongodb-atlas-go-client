# ApiAtlasUserScopeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable label that identifies the cluster or MongoDB Atlas Data Lake that this database user can access. | 
**Type** | **string** | Category of resource that this database user can access. | 

## Methods

### NewApiAtlasUserScopeView

`func NewApiAtlasUserScopeView(name string, type_ string, ) *ApiAtlasUserScopeView`

NewApiAtlasUserScopeView instantiates a new ApiAtlasUserScopeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasUserScopeViewWithDefaults

`func NewApiAtlasUserScopeViewWithDefaults() *ApiAtlasUserScopeView`

NewApiAtlasUserScopeViewWithDefaults instantiates a new ApiAtlasUserScopeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiAtlasUserScopeView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasUserScopeView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasUserScopeView) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ApiAtlasUserScopeView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAtlasUserScopeView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAtlasUserScopeView) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


