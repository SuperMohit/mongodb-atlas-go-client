# ApiAtlasContainerPeerViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection. | 
**ProviderName** | **string** | Cloud service provider that determines the needed settings to configure the network connection for a virtual private connection. | 
**AccepterRegionName** | **string** | Amazon Web Services (AWS) region where the Virtual Peering Connection (VPC) that you peered with the MongoDB Cloud VPC resides. The resource returns &#x60;null&#x60; if your VPC and the MongoDB Cloud VPC reside in the same region. | 
**AwsAccountId** | **string** | Unique twelve-digit string that identifies the Amazon Web Services (AWS) account that owns the VPC that you peered with the MongoDB Cloud VPC. | 
**ConnectionId** | Pointer to **string** | Unique string that identifies the peering connection on AWS. | [optional] [readonly] 
**ErrorStateName** | Pointer to **string** | Type of error that can be returned when requesting an Amazon Web Services (AWS) peering connection. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the network peering connection. | [optional] [readonly] 
**RouteTableCidrBlock** | **string** | Internet Protocol (IP) addresses expressed in Classless Inter-Domain Routing (CIDR) notation of the VPC&#39;s subnet that you want to peer with the MongoDB Cloud VPC. | 
**StatusName** | Pointer to **string** | State of the network peering connection at the time you made the request. | [optional] [readonly] 
**VpcId** | **string** | Unique string that identifies the VPC on Amazon Web Services (AWS) that you want to peer with the MongoDB Cloud VPC. | 
**AzureDirectoryId** | **string** | Unique string that identifies the Azure AD directory in which the VNet peered with the MongoDB Cloud VNet resides. | 
**AzureSubscriptionId** | **string** | Unique string that identifies the Azure subscription in which the VNet you peered with the MongoDB Cloud VNet resides. | 
**ErrorState** | Pointer to **string** | Error message returned when a requested Azure network peering resource returns &#x60;\&quot;status\&quot; : \&quot;FAILED\&quot;&#x60;. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**ResourceGroupName** | **string** | Human-readable label that identifies the resource group in which the VNet to peer with the MongoDB Cloud VNet resides. | 
**Status** | Pointer to **string** | State of the network peering connection at the time you made the request. | [optional] [readonly] 
**VnetName** | **string** | Human-readable label that identifies the VNet that you want to peer with the MongoDB Cloud VNet. | 
**ErrorMessage** | Pointer to **string** | Details of the error returned when requesting a GCP network peering resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**GcpProjectId** | **string** | Human-readable label that identifies the GCP project that contains the network that you want to peer with the MongoDB Cloud VPC. | 
**NetworkName** | **string** | Human-readable label that identifies the network to peer with the MongoDB Cloud VPC. | 

## Methods

### NewApiAtlasContainerPeerViewRequest

`func NewApiAtlasContainerPeerViewRequest(containerId string, providerName string, accepterRegionName string, awsAccountId string, routeTableCidrBlock string, vpcId string, azureDirectoryId string, azureSubscriptionId string, resourceGroupName string, vnetName string, gcpProjectId string, networkName string, ) *ApiAtlasContainerPeerViewRequest`

NewApiAtlasContainerPeerViewRequest instantiates a new ApiAtlasContainerPeerViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasContainerPeerViewRequestWithDefaults

`func NewApiAtlasContainerPeerViewRequestWithDefaults() *ApiAtlasContainerPeerViewRequest`

NewApiAtlasContainerPeerViewRequestWithDefaults instantiates a new ApiAtlasContainerPeerViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *ApiAtlasContainerPeerViewRequest) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ApiAtlasContainerPeerViewRequest) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ApiAtlasContainerPeerViewRequest) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetProviderName

`func (o *ApiAtlasContainerPeerViewRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ApiAtlasContainerPeerViewRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ApiAtlasContainerPeerViewRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetAccepterRegionName

`func (o *ApiAtlasContainerPeerViewRequest) GetAccepterRegionName() string`

GetAccepterRegionName returns the AccepterRegionName field if non-nil, zero value otherwise.

### GetAccepterRegionNameOk

`func (o *ApiAtlasContainerPeerViewRequest) GetAccepterRegionNameOk() (*string, bool)`

GetAccepterRegionNameOk returns a tuple with the AccepterRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterRegionName

`func (o *ApiAtlasContainerPeerViewRequest) SetAccepterRegionName(v string)`

SetAccepterRegionName sets AccepterRegionName field to given value.


### GetAwsAccountId

`func (o *ApiAtlasContainerPeerViewRequest) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *ApiAtlasContainerPeerViewRequest) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *ApiAtlasContainerPeerViewRequest) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.


### GetConnectionId

`func (o *ApiAtlasContainerPeerViewRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ApiAtlasContainerPeerViewRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ApiAtlasContainerPeerViewRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ApiAtlasContainerPeerViewRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetErrorStateName

`func (o *ApiAtlasContainerPeerViewRequest) GetErrorStateName() string`

GetErrorStateName returns the ErrorStateName field if non-nil, zero value otherwise.

### GetErrorStateNameOk

`func (o *ApiAtlasContainerPeerViewRequest) GetErrorStateNameOk() (*string, bool)`

GetErrorStateNameOk returns a tuple with the ErrorStateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorStateName

`func (o *ApiAtlasContainerPeerViewRequest) SetErrorStateName(v string)`

SetErrorStateName sets ErrorStateName field to given value.

### HasErrorStateName

`func (o *ApiAtlasContainerPeerViewRequest) HasErrorStateName() bool`

HasErrorStateName returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasContainerPeerViewRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasContainerPeerViewRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasContainerPeerViewRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasContainerPeerViewRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRouteTableCidrBlock

`func (o *ApiAtlasContainerPeerViewRequest) GetRouteTableCidrBlock() string`

GetRouteTableCidrBlock returns the RouteTableCidrBlock field if non-nil, zero value otherwise.

### GetRouteTableCidrBlockOk

`func (o *ApiAtlasContainerPeerViewRequest) GetRouteTableCidrBlockOk() (*string, bool)`

GetRouteTableCidrBlockOk returns a tuple with the RouteTableCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableCidrBlock

`func (o *ApiAtlasContainerPeerViewRequest) SetRouteTableCidrBlock(v string)`

SetRouteTableCidrBlock sets RouteTableCidrBlock field to given value.


### GetStatusName

`func (o *ApiAtlasContainerPeerViewRequest) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *ApiAtlasContainerPeerViewRequest) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *ApiAtlasContainerPeerViewRequest) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *ApiAtlasContainerPeerViewRequest) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetVpcId

`func (o *ApiAtlasContainerPeerViewRequest) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *ApiAtlasContainerPeerViewRequest) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *ApiAtlasContainerPeerViewRequest) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetAzureDirectoryId

`func (o *ApiAtlasContainerPeerViewRequest) GetAzureDirectoryId() string`

GetAzureDirectoryId returns the AzureDirectoryId field if non-nil, zero value otherwise.

### GetAzureDirectoryIdOk

`func (o *ApiAtlasContainerPeerViewRequest) GetAzureDirectoryIdOk() (*string, bool)`

GetAzureDirectoryIdOk returns a tuple with the AzureDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDirectoryId

`func (o *ApiAtlasContainerPeerViewRequest) SetAzureDirectoryId(v string)`

SetAzureDirectoryId sets AzureDirectoryId field to given value.


### GetAzureSubscriptionId

`func (o *ApiAtlasContainerPeerViewRequest) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *ApiAtlasContainerPeerViewRequest) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *ApiAtlasContainerPeerViewRequest) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### GetErrorState

`func (o *ApiAtlasContainerPeerViewRequest) GetErrorState() string`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *ApiAtlasContainerPeerViewRequest) GetErrorStateOk() (*string, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *ApiAtlasContainerPeerViewRequest) SetErrorState(v string)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *ApiAtlasContainerPeerViewRequest) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetResourceGroupName

`func (o *ApiAtlasContainerPeerViewRequest) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *ApiAtlasContainerPeerViewRequest) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *ApiAtlasContainerPeerViewRequest) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetStatus

`func (o *ApiAtlasContainerPeerViewRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAtlasContainerPeerViewRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAtlasContainerPeerViewRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAtlasContainerPeerViewRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVnetName

`func (o *ApiAtlasContainerPeerViewRequest) GetVnetName() string`

GetVnetName returns the VnetName field if non-nil, zero value otherwise.

### GetVnetNameOk

`func (o *ApiAtlasContainerPeerViewRequest) GetVnetNameOk() (*string, bool)`

GetVnetNameOk returns a tuple with the VnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetName

`func (o *ApiAtlasContainerPeerViewRequest) SetVnetName(v string)`

SetVnetName sets VnetName field to given value.


### GetErrorMessage

`func (o *ApiAtlasContainerPeerViewRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiAtlasContainerPeerViewRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiAtlasContainerPeerViewRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiAtlasContainerPeerViewRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetGcpProjectId

`func (o *ApiAtlasContainerPeerViewRequest) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *ApiAtlasContainerPeerViewRequest) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *ApiAtlasContainerPeerViewRequest) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.


### GetNetworkName

`func (o *ApiAtlasContainerPeerViewRequest) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *ApiAtlasContainerPeerViewRequest) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *ApiAtlasContainerPeerViewRequest) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


