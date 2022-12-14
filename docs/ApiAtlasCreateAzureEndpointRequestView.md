# ApiAtlasCreateAzureEndpointRequestView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique string that identifies the private endpoint&#39;s network interface that someone added to this private endpoint service. | [optional] 
**PrivateEndpointIPAddress** | Pointer to **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | [optional] 

## Methods

### NewApiAtlasCreateAzureEndpointRequestView

`func NewApiAtlasCreateAzureEndpointRequestView() *ApiAtlasCreateAzureEndpointRequestView`

NewApiAtlasCreateAzureEndpointRequestView instantiates a new ApiAtlasCreateAzureEndpointRequestView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCreateAzureEndpointRequestViewWithDefaults

`func NewApiAtlasCreateAzureEndpointRequestViewWithDefaults() *ApiAtlasCreateAzureEndpointRequestView`

NewApiAtlasCreateAzureEndpointRequestViewWithDefaults instantiates a new ApiAtlasCreateAzureEndpointRequestView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAtlasCreateAzureEndpointRequestView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasCreateAzureEndpointRequestView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasCreateAzureEndpointRequestView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasCreateAzureEndpointRequestView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrivateEndpointIPAddress

`func (o *ApiAtlasCreateAzureEndpointRequestView) GetPrivateEndpointIPAddress() string`

GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIPAddressOk

`func (o *ApiAtlasCreateAzureEndpointRequestView) GetPrivateEndpointIPAddressOk() (*string, bool)`

GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIPAddress

`func (o *ApiAtlasCreateAzureEndpointRequestView) SetPrivateEndpointIPAddress(v string)`

SetPrivateEndpointIPAddress sets PrivateEndpointIPAddress field to given value.

### HasPrivateEndpointIPAddress

`func (o *ApiAtlasCreateAzureEndpointRequestView) HasPrivateEndpointIPAddress() bool`

HasPrivateEndpointIPAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


