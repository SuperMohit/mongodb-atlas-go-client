# ApiAtlasGCPPeerVpcRequestView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection. | 
**ProviderName** | **string** | Cloud service provider that determines the needed settings to configure the network connection for a virtual private connection. | 
**ErrorMessage** | Pointer to **string** | Details of the error returned when requesting a GCP network peering resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**GcpProjectId** | **string** | Human-readable label that identifies the GCP project that contains the network that you want to peer with the MongoDB Cloud VPC. | 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the network peering connection. | [optional] [readonly] 
**NetworkName** | **string** | Human-readable label that identifies the network to peer with the MongoDB Cloud VPC. | 
**Status** | Pointer to **string** | State of the network peering connection at the time you made the request. | [optional] [readonly] 

## Methods

### NewApiAtlasGCPPeerVpcRequestView

`func NewApiAtlasGCPPeerVpcRequestView(containerId string, providerName string, gcpProjectId string, networkName string, ) *ApiAtlasGCPPeerVpcRequestView`

NewApiAtlasGCPPeerVpcRequestView instantiates a new ApiAtlasGCPPeerVpcRequestView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasGCPPeerVpcRequestViewWithDefaults

`func NewApiAtlasGCPPeerVpcRequestViewWithDefaults() *ApiAtlasGCPPeerVpcRequestView`

NewApiAtlasGCPPeerVpcRequestViewWithDefaults instantiates a new ApiAtlasGCPPeerVpcRequestView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *ApiAtlasGCPPeerVpcRequestView) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ApiAtlasGCPPeerVpcRequestView) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ApiAtlasGCPPeerVpcRequestView) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetProviderName

`func (o *ApiAtlasGCPPeerVpcRequestView) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ApiAtlasGCPPeerVpcRequestView) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ApiAtlasGCPPeerVpcRequestView) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetErrorMessage

`func (o *ApiAtlasGCPPeerVpcRequestView) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiAtlasGCPPeerVpcRequestView) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiAtlasGCPPeerVpcRequestView) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiAtlasGCPPeerVpcRequestView) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetGcpProjectId

`func (o *ApiAtlasGCPPeerVpcRequestView) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *ApiAtlasGCPPeerVpcRequestView) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *ApiAtlasGCPPeerVpcRequestView) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.


### GetId

`func (o *ApiAtlasGCPPeerVpcRequestView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasGCPPeerVpcRequestView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasGCPPeerVpcRequestView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasGCPPeerVpcRequestView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkName

`func (o *ApiAtlasGCPPeerVpcRequestView) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *ApiAtlasGCPPeerVpcRequestView) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *ApiAtlasGCPPeerVpcRequestView) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetStatus

`func (o *ApiAtlasGCPPeerVpcRequestView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAtlasGCPPeerVpcRequestView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAtlasGCPPeerVpcRequestView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAtlasGCPPeerVpcRequestView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


