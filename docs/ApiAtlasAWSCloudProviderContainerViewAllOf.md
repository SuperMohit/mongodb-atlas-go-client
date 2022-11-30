# ApiAtlasAWSCloudProviderContainerViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtlasCidrBlock** | Pointer to **string** | IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project. MongoDB Cloud assigns all of the project&#39;s clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.  These CIDR blocks must fall within the ranges reserved per RFC 1918. AWS and Azure further limit the block to between the &#x60;/24&#x60; and  &#x60;/21&#x60; ranges.  To modify the CIDR block, the target project cannot have:  - Any M10 or greater clusters - Any other VPC peering connections   You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.   **Example:** A project in an Amazon Web Services (AWS) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of &#x60;/24&#x60; equals 27 three-node replica sets. | [optional] 
**RegionName** | Pointer to **string** | Geographic area that Amazon Web Services (AWS) defines to which MongoDB Cloud deployed this network peering container. | [optional] 
**VpcId** | Pointer to **string** | Unique string that identifies the MongoDB Cloud VPC on AWS. | [optional] [readonly] 

## Methods

### NewApiAtlasAWSCloudProviderContainerViewAllOf

`func NewApiAtlasAWSCloudProviderContainerViewAllOf() *ApiAtlasAWSCloudProviderContainerViewAllOf`

NewApiAtlasAWSCloudProviderContainerViewAllOf instantiates a new ApiAtlasAWSCloudProviderContainerViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasAWSCloudProviderContainerViewAllOfWithDefaults

`func NewApiAtlasAWSCloudProviderContainerViewAllOfWithDefaults() *ApiAtlasAWSCloudProviderContainerViewAllOf`

NewApiAtlasAWSCloudProviderContainerViewAllOfWithDefaults instantiates a new ApiAtlasAWSCloudProviderContainerViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlasCidrBlock

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) GetAtlasCidrBlock() string`

GetAtlasCidrBlock returns the AtlasCidrBlock field if non-nil, zero value otherwise.

### GetAtlasCidrBlockOk

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) GetAtlasCidrBlockOk() (*string, bool)`

GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasCidrBlock

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) SetAtlasCidrBlock(v string)`

SetAtlasCidrBlock sets AtlasCidrBlock field to given value.

### HasAtlasCidrBlock

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) HasAtlasCidrBlock() bool`

HasAtlasCidrBlock returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetVpcId

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *ApiAtlasAWSCloudProviderContainerViewAllOf) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


