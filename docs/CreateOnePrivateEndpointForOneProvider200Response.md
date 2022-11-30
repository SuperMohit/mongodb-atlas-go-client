# CreateOnePrivateEndpointForOneProvider200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStatus** | Pointer to **string** | State of the Amazon Web Service PrivateLink connection when MongoDB Cloud received this request. | [optional] [readonly] 
**DeleteRequested** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud received a request to remove the specified private endpoint from the private endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**InterfaceEndpointId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the interface endpoint. | [optional] [readonly] 
**PrivateEndpointConnectionName** | Pointer to **string** | Human-readable label that MongoDB Cloud generates that identifies the private endpoint connection. | [optional] [readonly] 
**PrivateEndpointIPAddress** | Pointer to **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | [optional] 
**PrivateEndpointResourceId** | Pointer to **string** | Unique string that identifies the Azure private endpoint&#39;s network interface that someone added to this private endpoint service. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the Google Cloud network endpoint group when MongoDB Cloud received this request. | [optional] [readonly] 
**EndpointGroupName** | Pointer to **string** | Human-readable label that identifies a set of endpoints. | [optional] [readonly] 
**Endpoints** | Pointer to [**[]ApiAtlasGCPConsumerForwardingRuleView**](ApiAtlasGCPConsumerForwardingRuleView.md) | List of individual private endpoints that comprise this endpoint group. | [optional] [readonly] 

## Methods

### NewCreateOnePrivateEndpointForOneProvider200Response

`func NewCreateOnePrivateEndpointForOneProvider200Response() *CreateOnePrivateEndpointForOneProvider200Response`

NewCreateOnePrivateEndpointForOneProvider200Response instantiates a new CreateOnePrivateEndpointForOneProvider200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOnePrivateEndpointForOneProvider200ResponseWithDefaults

`func NewCreateOnePrivateEndpointForOneProvider200ResponseWithDefaults() *CreateOnePrivateEndpointForOneProvider200Response`

NewCreateOnePrivateEndpointForOneProvider200ResponseWithDefaults instantiates a new CreateOnePrivateEndpointForOneProvider200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatus

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *CreateOnePrivateEndpointForOneProvider200Response) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *CreateOnePrivateEndpointForOneProvider200Response) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDeleteRequested

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetDeleteRequested() bool`

GetDeleteRequested returns the DeleteRequested field if non-nil, zero value otherwise.

### GetDeleteRequestedOk

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetDeleteRequestedOk() (*bool, bool)`

GetDeleteRequestedOk returns a tuple with the DeleteRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRequested

`func (o *CreateOnePrivateEndpointForOneProvider200Response) SetDeleteRequested(v bool)`

SetDeleteRequested sets DeleteRequested field to given value.

### HasDeleteRequested

`func (o *CreateOnePrivateEndpointForOneProvider200Response) HasDeleteRequested() bool`

HasDeleteRequested returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CreateOnePrivateEndpointForOneProvider200Response) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CreateOnePrivateEndpointForOneProvider200Response) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetInterfaceEndpointId

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetInterfaceEndpointId() string`

GetInterfaceEndpointId returns the InterfaceEndpointId field if non-nil, zero value otherwise.

### GetInterfaceEndpointIdOk

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetInterfaceEndpointIdOk() (*string, bool)`

GetInterfaceEndpointIdOk returns a tuple with the InterfaceEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceEndpointId

`func (o *CreateOnePrivateEndpointForOneProvider200Response) SetInterfaceEndpointId(v string)`

SetInterfaceEndpointId sets InterfaceEndpointId field to given value.

### HasInterfaceEndpointId

`func (o *CreateOnePrivateEndpointForOneProvider200Response) HasInterfaceEndpointId() bool`

HasInterfaceEndpointId returns a boolean if a field has been set.

### GetPrivateEndpointConnectionName

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetPrivateEndpointConnectionName() string`

GetPrivateEndpointConnectionName returns the PrivateEndpointConnectionName field if non-nil, zero value otherwise.

### GetPrivateEndpointConnectionNameOk

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetPrivateEndpointConnectionNameOk() (*string, bool)`

GetPrivateEndpointConnectionNameOk returns a tuple with the PrivateEndpointConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointConnectionName

`func (o *CreateOnePrivateEndpointForOneProvider200Response) SetPrivateEndpointConnectionName(v string)`

SetPrivateEndpointConnectionName sets PrivateEndpointConnectionName field to given value.

### HasPrivateEndpointConnectionName

`func (o *CreateOnePrivateEndpointForOneProvider200Response) HasPrivateEndpointConnectionName() bool`

HasPrivateEndpointConnectionName returns a boolean if a field has been set.

### GetPrivateEndpointIPAddress

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetPrivateEndpointIPAddress() string`

GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIPAddressOk

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetPrivateEndpointIPAddressOk() (*string, bool)`

GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIPAddress

`func (o *CreateOnePrivateEndpointForOneProvider200Response) SetPrivateEndpointIPAddress(v string)`

SetPrivateEndpointIPAddress sets PrivateEndpointIPAddress field to given value.

### HasPrivateEndpointIPAddress

`func (o *CreateOnePrivateEndpointForOneProvider200Response) HasPrivateEndpointIPAddress() bool`

HasPrivateEndpointIPAddress returns a boolean if a field has been set.

### GetPrivateEndpointResourceId

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetPrivateEndpointResourceId() string`

GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field if non-nil, zero value otherwise.

### GetPrivateEndpointResourceIdOk

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetPrivateEndpointResourceIdOk() (*string, bool)`

GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointResourceId

`func (o *CreateOnePrivateEndpointForOneProvider200Response) SetPrivateEndpointResourceId(v string)`

SetPrivateEndpointResourceId sets PrivateEndpointResourceId field to given value.

### HasPrivateEndpointResourceId

`func (o *CreateOnePrivateEndpointForOneProvider200Response) HasPrivateEndpointResourceId() bool`

HasPrivateEndpointResourceId returns a boolean if a field has been set.

### GetStatus

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateOnePrivateEndpointForOneProvider200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateOnePrivateEndpointForOneProvider200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEndpointGroupName

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *CreateOnePrivateEndpointForOneProvider200Response) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *CreateOnePrivateEndpointForOneProvider200Response) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetEndpoints

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetEndpoints() []ApiAtlasGCPConsumerForwardingRuleView`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *CreateOnePrivateEndpointForOneProvider200Response) GetEndpointsOk() (*[]ApiAtlasGCPConsumerForwardingRuleView, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *CreateOnePrivateEndpointForOneProvider200Response) SetEndpoints(v []ApiAtlasGCPConsumerForwardingRuleView)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *CreateOnePrivateEndpointForOneProvider200Response) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


