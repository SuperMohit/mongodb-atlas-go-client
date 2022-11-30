# ReturnOneNetworkPeeringConnectionInOneProject200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccepterRegionName** | **string** | Amazon Web Services (AWS) region where the Virtual Peering Connection (VPC) that you peered with the MongoDB Cloud VPC resides. The resource returns &#x60;null&#x60; if your VPC and the MongoDB Cloud VPC reside in the same region. | 
**AwsAccountId** | **string** | Unique twelve-digit string that identifies the Amazon Web Services (AWS) account that owns the VPC that you peered with the MongoDB Cloud VPC. | 
**ConnectionId** | Pointer to **string** | Unique string that identifies the peering connection on AWS. | [optional] [readonly] 
**ContainerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection. | 
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

### NewReturnOneNetworkPeeringConnectionInOneProject200Response

`func NewReturnOneNetworkPeeringConnectionInOneProject200Response(accepterRegionName string, awsAccountId string, containerId string, routeTableCidrBlock string, vpcId string, azureDirectoryId string, azureSubscriptionId string, resourceGroupName string, vnetName string, gcpProjectId string, networkName string, ) *ReturnOneNetworkPeeringConnectionInOneProject200Response`

NewReturnOneNetworkPeeringConnectionInOneProject200Response instantiates a new ReturnOneNetworkPeeringConnectionInOneProject200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnOneNetworkPeeringConnectionInOneProject200ResponseWithDefaults

`func NewReturnOneNetworkPeeringConnectionInOneProject200ResponseWithDefaults() *ReturnOneNetworkPeeringConnectionInOneProject200Response`

NewReturnOneNetworkPeeringConnectionInOneProject200ResponseWithDefaults instantiates a new ReturnOneNetworkPeeringConnectionInOneProject200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepterRegionName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetAccepterRegionName() string`

GetAccepterRegionName returns the AccepterRegionName field if non-nil, zero value otherwise.

### GetAccepterRegionNameOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetAccepterRegionNameOk() (*string, bool)`

GetAccepterRegionNameOk returns a tuple with the AccepterRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepterRegionName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetAccepterRegionName(v string)`

SetAccepterRegionName sets AccepterRegionName field to given value.


### GetAwsAccountId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.


### GetConnectionId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetContainerId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetErrorStateName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetErrorStateName() string`

GetErrorStateName returns the ErrorStateName field if non-nil, zero value otherwise.

### GetErrorStateNameOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetErrorStateNameOk() (*string, bool)`

GetErrorStateNameOk returns a tuple with the ErrorStateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorStateName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetErrorStateName(v string)`

SetErrorStateName sets ErrorStateName field to given value.

### HasErrorStateName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) HasErrorStateName() bool`

HasErrorStateName returns a boolean if a field has been set.

### GetId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRouteTableCidrBlock

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetRouteTableCidrBlock() string`

GetRouteTableCidrBlock returns the RouteTableCidrBlock field if non-nil, zero value otherwise.

### GetRouteTableCidrBlockOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetRouteTableCidrBlockOk() (*string, bool)`

GetRouteTableCidrBlockOk returns a tuple with the RouteTableCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableCidrBlock

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetRouteTableCidrBlock(v string)`

SetRouteTableCidrBlock sets RouteTableCidrBlock field to given value.


### GetStatusName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetVpcId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetAzureDirectoryId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetAzureDirectoryId() string`

GetAzureDirectoryId returns the AzureDirectoryId field if non-nil, zero value otherwise.

### GetAzureDirectoryIdOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetAzureDirectoryIdOk() (*string, bool)`

GetAzureDirectoryIdOk returns a tuple with the AzureDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDirectoryId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetAzureDirectoryId(v string)`

SetAzureDirectoryId sets AzureDirectoryId field to given value.


### GetAzureSubscriptionId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### GetErrorState

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetErrorState() string`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetErrorStateOk() (*string, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetErrorState(v string)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetResourceGroupName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetStatus

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVnetName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetVnetName() string`

GetVnetName returns the VnetName field if non-nil, zero value otherwise.

### GetVnetNameOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetVnetNameOk() (*string, bool)`

GetVnetNameOk returns a tuple with the VnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetVnetName(v string)`

SetVnetName sets VnetName field to given value.


### GetErrorMessage

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetGcpProjectId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.


### GetNetworkName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *ReturnOneNetworkPeeringConnectionInOneProject200Response) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


