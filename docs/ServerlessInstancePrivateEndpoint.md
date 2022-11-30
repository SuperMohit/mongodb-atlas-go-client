# ServerlessInstancePrivateEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | [**[]ServerlessInstancePrivateEndpointEndpoint**](ServerlessInstancePrivateEndpointEndpoint.md) | List that contains the private endpoints through which you connect to MongoDB Cloud when you use **connectionStrings.privateEndpoint[n].srvConnectionString**. | [readonly] 
**SrvConnectionString** | **string** | Private endpoint-aware connection string that uses the &#x60;mongodb+srv://&#x60; protocol to connect to MongoDB Cloud through a private endpoint. The &#x60;mongodb+srv&#x60; protocol tells the driver to look up the seed list of hosts in the Domain Name System (DNS). | [readonly] 
**Type** | **string** | MongoDB process type to which your application connects.  | [readonly] 

## Methods

### NewServerlessInstancePrivateEndpoint

`func NewServerlessInstancePrivateEndpoint(endpoints []ServerlessInstancePrivateEndpointEndpoint, srvConnectionString string, type_ string, ) *ServerlessInstancePrivateEndpoint`

NewServerlessInstancePrivateEndpoint instantiates a new ServerlessInstancePrivateEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessInstancePrivateEndpointWithDefaults

`func NewServerlessInstancePrivateEndpointWithDefaults() *ServerlessInstancePrivateEndpoint`

NewServerlessInstancePrivateEndpointWithDefaults instantiates a new ServerlessInstancePrivateEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *ServerlessInstancePrivateEndpoint) GetEndpoints() []ServerlessInstancePrivateEndpointEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ServerlessInstancePrivateEndpoint) GetEndpointsOk() (*[]ServerlessInstancePrivateEndpointEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ServerlessInstancePrivateEndpoint) SetEndpoints(v []ServerlessInstancePrivateEndpointEndpoint)`

SetEndpoints sets Endpoints field to given value.


### GetSrvConnectionString

`func (o *ServerlessInstancePrivateEndpoint) GetSrvConnectionString() string`

GetSrvConnectionString returns the SrvConnectionString field if non-nil, zero value otherwise.

### GetSrvConnectionStringOk

`func (o *ServerlessInstancePrivateEndpoint) GetSrvConnectionStringOk() (*string, bool)`

GetSrvConnectionStringOk returns a tuple with the SrvConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrvConnectionString

`func (o *ServerlessInstancePrivateEndpoint) SetSrvConnectionString(v string)`

SetSrvConnectionString sets SrvConnectionString field to given value.


### GetType

`func (o *ServerlessInstancePrivateEndpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerlessInstancePrivateEndpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerlessInstancePrivateEndpoint) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


