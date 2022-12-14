# CreateOnePrivateEndpointForOneProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique string that identifies the private endpoint&#39;s network interface that someone added to this private endpoint service. | [optional] 
**PrivateEndpointIPAddress** | Pointer to **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | [optional] 
**EndpointGroupName** | Pointer to **string** | Human-readable label that identifies a set of endpoints. | [optional] 
**Endpoints** | Pointer to [**[]ApiAtlasCreateGCPForwardingRuleRequestView**](ApiAtlasCreateGCPForwardingRuleRequestView.md) | List of individual private endpoints that comprise this endpoint group. | [optional] 
**GcpProjectId** | Pointer to **string** | Unique string that identifies the Google Cloud project in which you created the endpoints. | [optional] 

## Methods

### NewCreateOnePrivateEndpointForOneProviderRequest

`func NewCreateOnePrivateEndpointForOneProviderRequest() *CreateOnePrivateEndpointForOneProviderRequest`

NewCreateOnePrivateEndpointForOneProviderRequest instantiates a new CreateOnePrivateEndpointForOneProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOnePrivateEndpointForOneProviderRequestWithDefaults

`func NewCreateOnePrivateEndpointForOneProviderRequestWithDefaults() *CreateOnePrivateEndpointForOneProviderRequest`

NewCreateOnePrivateEndpointForOneProviderRequestWithDefaults instantiates a new CreateOnePrivateEndpointForOneProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOnePrivateEndpointForOneProviderRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOnePrivateEndpointForOneProviderRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOnePrivateEndpointForOneProviderRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOnePrivateEndpointForOneProviderRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrivateEndpointIPAddress

`func (o *CreateOnePrivateEndpointForOneProviderRequest) GetPrivateEndpointIPAddress() string`

GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIPAddressOk

`func (o *CreateOnePrivateEndpointForOneProviderRequest) GetPrivateEndpointIPAddressOk() (*string, bool)`

GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIPAddress

`func (o *CreateOnePrivateEndpointForOneProviderRequest) SetPrivateEndpointIPAddress(v string)`

SetPrivateEndpointIPAddress sets PrivateEndpointIPAddress field to given value.

### HasPrivateEndpointIPAddress

`func (o *CreateOnePrivateEndpointForOneProviderRequest) HasPrivateEndpointIPAddress() bool`

HasPrivateEndpointIPAddress returns a boolean if a field has been set.

### GetEndpointGroupName

`func (o *CreateOnePrivateEndpointForOneProviderRequest) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *CreateOnePrivateEndpointForOneProviderRequest) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *CreateOnePrivateEndpointForOneProviderRequest) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *CreateOnePrivateEndpointForOneProviderRequest) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetEndpoints

`func (o *CreateOnePrivateEndpointForOneProviderRequest) GetEndpoints() []ApiAtlasCreateGCPForwardingRuleRequestView`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *CreateOnePrivateEndpointForOneProviderRequest) GetEndpointsOk() (*[]ApiAtlasCreateGCPForwardingRuleRequestView, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *CreateOnePrivateEndpointForOneProviderRequest) SetEndpoints(v []ApiAtlasCreateGCPForwardingRuleRequestView)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *CreateOnePrivateEndpointForOneProviderRequest) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetGcpProjectId

`func (o *CreateOnePrivateEndpointForOneProviderRequest) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *CreateOnePrivateEndpointForOneProviderRequest) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *CreateOnePrivateEndpointForOneProviderRequest) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *CreateOnePrivateEndpointForOneProviderRequest) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


