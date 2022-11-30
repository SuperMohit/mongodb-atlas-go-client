# ApiAtlasAzureCloudProviderContainerViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtlasCidrBlock** | Pointer to **string** | IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project. MongoDB Cloud assigns all of the project&#39;s clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.  These CIDR blocks must fall within the ranges reserved per RFC 1918. AWS and Azure further limit the block to between the &#x60;/24&#x60; and  &#x60;/21&#x60; ranges.  To modify the CIDR block, the target project cannot have:  - Any M10 or greater clusters - Any other VPC peering connections   You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.   **Example:** A project in an Amazon Web Services (AWS) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of &#x60;/24&#x60; equals 27 three-node replica sets. | [optional] 
**AzureSubscriptionId** | Pointer to **string** | Unique string that identifies the Azure subscription in which the MongoDB Cloud VNet resides. | [optional] [readonly] 
**Region** | Pointer to **string** | Azure region to which MongoDB Cloud deployed this network peering container. | [optional] 
**VnetName** | Pointer to **string** | Unique string that identifies the Azure VNet in which MongoDB Cloud clusters in this network peering container exist. The response returns **null** if no clusters exist in this network peering container. | [optional] [readonly] 

## Methods

### NewApiAtlasAzureCloudProviderContainerViewAllOf

`func NewApiAtlasAzureCloudProviderContainerViewAllOf() *ApiAtlasAzureCloudProviderContainerViewAllOf`

NewApiAtlasAzureCloudProviderContainerViewAllOf instantiates a new ApiAtlasAzureCloudProviderContainerViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasAzureCloudProviderContainerViewAllOfWithDefaults

`func NewApiAtlasAzureCloudProviderContainerViewAllOfWithDefaults() *ApiAtlasAzureCloudProviderContainerViewAllOf`

NewApiAtlasAzureCloudProviderContainerViewAllOfWithDefaults instantiates a new ApiAtlasAzureCloudProviderContainerViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlasCidrBlock

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) GetAtlasCidrBlock() string`

GetAtlasCidrBlock returns the AtlasCidrBlock field if non-nil, zero value otherwise.

### GetAtlasCidrBlockOk

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) GetAtlasCidrBlockOk() (*string, bool)`

GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasCidrBlock

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) SetAtlasCidrBlock(v string)`

SetAtlasCidrBlock sets AtlasCidrBlock field to given value.

### HasAtlasCidrBlock

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) HasAtlasCidrBlock() bool`

HasAtlasCidrBlock returns a boolean if a field has been set.

### GetAzureSubscriptionId

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.

### GetRegion

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVnetName

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) GetVnetName() string`

GetVnetName returns the VnetName field if non-nil, zero value otherwise.

### GetVnetNameOk

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) GetVnetNameOk() (*string, bool)`

GetVnetNameOk returns a tuple with the VnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetName

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) SetVnetName(v string)`

SetVnetName sets VnetName field to given value.

### HasVnetName

`func (o *ApiAtlasAzureCloudProviderContainerViewAllOf) HasVnetName() bool`

HasVnetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


