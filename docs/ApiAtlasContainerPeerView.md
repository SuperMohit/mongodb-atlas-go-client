# ApiAtlasContainerPeerView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection. | 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the network peering connection. | [optional] [readonly] 

## Methods

### NewApiAtlasContainerPeerView

`func NewApiAtlasContainerPeerView(containerId string, ) *ApiAtlasContainerPeerView`

NewApiAtlasContainerPeerView instantiates a new ApiAtlasContainerPeerView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasContainerPeerViewWithDefaults

`func NewApiAtlasContainerPeerViewWithDefaults() *ApiAtlasContainerPeerView`

NewApiAtlasContainerPeerViewWithDefaults instantiates a new ApiAtlasContainerPeerView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *ApiAtlasContainerPeerView) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ApiAtlasContainerPeerView) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ApiAtlasContainerPeerView) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetId

`func (o *ApiAtlasContainerPeerView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasContainerPeerView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasContainerPeerView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasContainerPeerView) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


