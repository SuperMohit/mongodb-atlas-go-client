# ApiAtlasAzurePeerNetworkRequestView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection. | 
**ProviderName** | **string** | Cloud service provider that determines the needed settings to configure the network connection for a virtual private connection. | 
**AzureDirectoryId** | **string** | Unique string that identifies the Azure AD directory in which the VNet peered with the MongoDB Cloud VNet resides. | 
**AzureSubscriptionId** | **string** | Unique string that identifies the Azure subscription in which the VNet you peered with the MongoDB Cloud VNet resides. | 
**ErrorState** | Pointer to **string** | Error message returned when a requested Azure network peering resource returns &#x60;\&quot;status\&quot; : \&quot;FAILED\&quot;&#x60;. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the network peering connection. | [optional] [readonly] 
**ResourceGroupName** | **string** | Human-readable label that identifies the resource group in which the VNet to peer with the MongoDB Cloud VNet resides. | 
**Status** | Pointer to **string** | State of the network peering connection at the time you made the request. | [optional] [readonly] 
**VnetName** | **string** | Human-readable label that identifies the VNet that you want to peer with the MongoDB Cloud VNet. | 

## Methods

### NewApiAtlasAzurePeerNetworkRequestView

`func NewApiAtlasAzurePeerNetworkRequestView(containerId string, providerName string, azureDirectoryId string, azureSubscriptionId string, resourceGroupName string, vnetName string, ) *ApiAtlasAzurePeerNetworkRequestView`

NewApiAtlasAzurePeerNetworkRequestView instantiates a new ApiAtlasAzurePeerNetworkRequestView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasAzurePeerNetworkRequestViewWithDefaults

`func NewApiAtlasAzurePeerNetworkRequestViewWithDefaults() *ApiAtlasAzurePeerNetworkRequestView`

NewApiAtlasAzurePeerNetworkRequestViewWithDefaults instantiates a new ApiAtlasAzurePeerNetworkRequestView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ApiAtlasAzurePeerNetworkRequestView) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetProviderName

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ApiAtlasAzurePeerNetworkRequestView) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetAzureDirectoryId

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetAzureDirectoryId() string`

GetAzureDirectoryId returns the AzureDirectoryId field if non-nil, zero value otherwise.

### GetAzureDirectoryIdOk

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetAzureDirectoryIdOk() (*string, bool)`

GetAzureDirectoryIdOk returns a tuple with the AzureDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDirectoryId

`func (o *ApiAtlasAzurePeerNetworkRequestView) SetAzureDirectoryId(v string)`

SetAzureDirectoryId sets AzureDirectoryId field to given value.


### GetAzureSubscriptionId

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *ApiAtlasAzurePeerNetworkRequestView) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### GetErrorState

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetErrorState() string`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetErrorStateOk() (*string, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *ApiAtlasAzurePeerNetworkRequestView) SetErrorState(v string)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *ApiAtlasAzurePeerNetworkRequestView) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasAzurePeerNetworkRequestView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasAzurePeerNetworkRequestView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceGroupName

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *ApiAtlasAzurePeerNetworkRequestView) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetStatus

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAtlasAzurePeerNetworkRequestView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAtlasAzurePeerNetworkRequestView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVnetName

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetVnetName() string`

GetVnetName returns the VnetName field if non-nil, zero value otherwise.

### GetVnetNameOk

`func (o *ApiAtlasAzurePeerNetworkRequestView) GetVnetNameOk() (*string, bool)`

GetVnetNameOk returns a tuple with the VnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetName

`func (o *ApiAtlasAzurePeerNetworkRequestView) SetVnetName(v string)`

SetVnetName sets VnetName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


