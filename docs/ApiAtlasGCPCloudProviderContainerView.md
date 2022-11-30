# ApiAtlasGCPCloudProviderContainerView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtlasCidrBlock** | **string** | IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project. MongoDB Cloud assigns all of the project&#39;s clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.  These CIDR blocks must fall within the ranges reserved per RFC 1918. GCP further limits the block to a lower bound of the &#x60;/18&#x60; range.  To modify the CIDR block, the target project cannot have:  - Any M10 or greater clusters - Any other VPC peering connections   You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.   **Example:** A project in an Google Cloud (GCP) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of &#x60;/24&#x60; equals 27 three-node replica sets. | 
**GcpProjectId** | Pointer to **string** | Unique string that identifies the GCP project in which MongoDB Cloud clusters in this network peering container exist. The response returns **null** if no clusters exist in this network peering container. | [optional] [readonly] 
**NetworkName** | Pointer to **string** | Human-readable label that identifies the network in which MongoDB Cloud clusters in this network peering container exist. MongoDB Cloud returns **null** if no clusters exist in this network peering container. | [optional] [readonly] 
**Regions** | **[]string** | List of GCP regions to which you want to deploy this MongoDB Cloud network peering container.  In this MongoDB Cloud project, you can deploy clusters only to the GCP regions in this list. To deploy MongoDB Cloud clusters to other GCP regions, create additional projects. | 

## Methods

### NewApiAtlasGCPCloudProviderContainerView

`func NewApiAtlasGCPCloudProviderContainerView(atlasCidrBlock string, regions []string, ) *ApiAtlasGCPCloudProviderContainerView`

NewApiAtlasGCPCloudProviderContainerView instantiates a new ApiAtlasGCPCloudProviderContainerView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasGCPCloudProviderContainerViewWithDefaults

`func NewApiAtlasGCPCloudProviderContainerViewWithDefaults() *ApiAtlasGCPCloudProviderContainerView`

NewApiAtlasGCPCloudProviderContainerViewWithDefaults instantiates a new ApiAtlasGCPCloudProviderContainerView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlasCidrBlock

`func (o *ApiAtlasGCPCloudProviderContainerView) GetAtlasCidrBlock() string`

GetAtlasCidrBlock returns the AtlasCidrBlock field if non-nil, zero value otherwise.

### GetAtlasCidrBlockOk

`func (o *ApiAtlasGCPCloudProviderContainerView) GetAtlasCidrBlockOk() (*string, bool)`

GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasCidrBlock

`func (o *ApiAtlasGCPCloudProviderContainerView) SetAtlasCidrBlock(v string)`

SetAtlasCidrBlock sets AtlasCidrBlock field to given value.


### GetGcpProjectId

`func (o *ApiAtlasGCPCloudProviderContainerView) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *ApiAtlasGCPCloudProviderContainerView) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *ApiAtlasGCPCloudProviderContainerView) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *ApiAtlasGCPCloudProviderContainerView) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.

### GetNetworkName

`func (o *ApiAtlasGCPCloudProviderContainerView) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *ApiAtlasGCPCloudProviderContainerView) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *ApiAtlasGCPCloudProviderContainerView) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *ApiAtlasGCPCloudProviderContainerView) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetRegions

`func (o *ApiAtlasGCPCloudProviderContainerView) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ApiAtlasGCPCloudProviderContainerView) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ApiAtlasGCPCloudProviderContainerView) SetRegions(v []string)`

SetRegions sets Regions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


