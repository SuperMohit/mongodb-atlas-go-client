# ServerlessInstancePrivateEndpointEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | **string** | Unique provider identifier of the private endpoint.  | [readonly] 
**ProviderName** | **string** | Cloud provider where the private endpoint is deployed.  | [readonly] 
**Region** | **string** | Region where the private endpoint is deployed.  | [readonly] 

## Methods

### NewServerlessInstancePrivateEndpointEndpoint

`func NewServerlessInstancePrivateEndpointEndpoint(endpointId string, providerName string, region string, ) *ServerlessInstancePrivateEndpointEndpoint`

NewServerlessInstancePrivateEndpointEndpoint instantiates a new ServerlessInstancePrivateEndpointEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessInstancePrivateEndpointEndpointWithDefaults

`func NewServerlessInstancePrivateEndpointEndpointWithDefaults() *ServerlessInstancePrivateEndpointEndpoint`

NewServerlessInstancePrivateEndpointEndpointWithDefaults instantiates a new ServerlessInstancePrivateEndpointEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *ServerlessInstancePrivateEndpointEndpoint) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ServerlessInstancePrivateEndpointEndpoint) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ServerlessInstancePrivateEndpointEndpoint) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.


### GetProviderName

`func (o *ServerlessInstancePrivateEndpointEndpoint) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ServerlessInstancePrivateEndpointEndpoint) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ServerlessInstancePrivateEndpointEndpoint) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetRegion

`func (o *ServerlessInstancePrivateEndpointEndpoint) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ServerlessInstancePrivateEndpointEndpoint) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ServerlessInstancePrivateEndpointEndpoint) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


