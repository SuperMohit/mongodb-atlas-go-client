# ApiAtlasCloudProviderContainerView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the network peering container. | [optional] [readonly] 
**ProviderName** | Pointer to **string** | Cloud service provider that serves the requested network peering containers. | [optional] 
**Provisioned** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud clusters exist in the specified network peering container. | [optional] [readonly] 

## Methods

### NewApiAtlasCloudProviderContainerView

`func NewApiAtlasCloudProviderContainerView() *ApiAtlasCloudProviderContainerView`

NewApiAtlasCloudProviderContainerView instantiates a new ApiAtlasCloudProviderContainerView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCloudProviderContainerViewWithDefaults

`func NewApiAtlasCloudProviderContainerViewWithDefaults() *ApiAtlasCloudProviderContainerView`

NewApiAtlasCloudProviderContainerViewWithDefaults instantiates a new ApiAtlasCloudProviderContainerView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAtlasCloudProviderContainerView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasCloudProviderContainerView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasCloudProviderContainerView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasCloudProviderContainerView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderName

`func (o *ApiAtlasCloudProviderContainerView) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ApiAtlasCloudProviderContainerView) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ApiAtlasCloudProviderContainerView) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ApiAtlasCloudProviderContainerView) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetProvisioned

`func (o *ApiAtlasCloudProviderContainerView) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *ApiAtlasCloudProviderContainerView) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *ApiAtlasCloudProviderContainerView) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *ApiAtlasCloudProviderContainerView) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


