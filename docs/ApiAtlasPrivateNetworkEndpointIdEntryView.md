# ApiAtlasPrivateNetworkEndpointIdEntryView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Human-readable string to associate with this private endpoint. | [optional] 
**EndpointId** | **string** | Unique 22-character alphanumeric string that identifies the private endpoint. | 
**Provider** | Pointer to **string** | Human-readable label that identifies the cloud service provider. Atlas Data Lake supports Amazon Web Services only. | [optional] [default to "AWS"]
**Type** | Pointer to **string** | Human-readable label that identifies the resource type associated with this private endpoint. | [optional] [default to "DATA_LAKE"]

## Methods

### NewApiAtlasPrivateNetworkEndpointIdEntryView

`func NewApiAtlasPrivateNetworkEndpointIdEntryView(endpointId string, ) *ApiAtlasPrivateNetworkEndpointIdEntryView`

NewApiAtlasPrivateNetworkEndpointIdEntryView instantiates a new ApiAtlasPrivateNetworkEndpointIdEntryView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasPrivateNetworkEndpointIdEntryViewWithDefaults

`func NewApiAtlasPrivateNetworkEndpointIdEntryViewWithDefaults() *ApiAtlasPrivateNetworkEndpointIdEntryView`

NewApiAtlasPrivateNetworkEndpointIdEntryViewWithDefaults instantiates a new ApiAtlasPrivateNetworkEndpointIdEntryView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEndpointId

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.


### GetProvider

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetType

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiAtlasPrivateNetworkEndpointIdEntryView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


