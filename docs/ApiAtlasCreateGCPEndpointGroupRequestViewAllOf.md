# ApiAtlasCreateGCPEndpointGroupRequestViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointGroupName** | Pointer to **string** | Human-readable label that identifies a set of endpoints. | [optional] 
**Endpoints** | Pointer to [**[]ApiAtlasCreateGCPForwardingRuleRequestView**](ApiAtlasCreateGCPForwardingRuleRequestView.md) | List of individual private endpoints that comprise this endpoint group. | [optional] 
**GcpProjectId** | Pointer to **string** | Unique string that identifies the Google Cloud project in which you created the endpoints. | [optional] 

## Methods

### NewApiAtlasCreateGCPEndpointGroupRequestViewAllOf

`func NewApiAtlasCreateGCPEndpointGroupRequestViewAllOf() *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf`

NewApiAtlasCreateGCPEndpointGroupRequestViewAllOf instantiates a new ApiAtlasCreateGCPEndpointGroupRequestViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCreateGCPEndpointGroupRequestViewAllOfWithDefaults

`func NewApiAtlasCreateGCPEndpointGroupRequestViewAllOfWithDefaults() *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf`

NewApiAtlasCreateGCPEndpointGroupRequestViewAllOfWithDefaults instantiates a new ApiAtlasCreateGCPEndpointGroupRequestViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointGroupName

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetEndpoints

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) GetEndpoints() []ApiAtlasCreateGCPForwardingRuleRequestView`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) GetEndpointsOk() (*[]ApiAtlasCreateGCPForwardingRuleRequestView, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) SetEndpoints(v []ApiAtlasCreateGCPForwardingRuleRequestView)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetGcpProjectId

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *ApiAtlasCreateGCPEndpointGroupRequestViewAllOf) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


