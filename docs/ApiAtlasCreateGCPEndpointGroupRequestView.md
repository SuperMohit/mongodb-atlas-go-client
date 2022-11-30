# ApiAtlasCreateGCPEndpointGroupRequestView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointGroupName** | Pointer to **string** | Human-readable label that identifies a set of endpoints. | [optional] 
**Endpoints** | Pointer to [**[]ApiAtlasCreateGCPForwardingRuleRequestView**](ApiAtlasCreateGCPForwardingRuleRequestView.md) | List of individual private endpoints that comprise this endpoint group. | [optional] 
**GcpProjectId** | Pointer to **string** | Unique string that identifies the Google Cloud project in which you created the endpoints. | [optional] 

## Methods

### NewApiAtlasCreateGCPEndpointGroupRequestView

`func NewApiAtlasCreateGCPEndpointGroupRequestView() *ApiAtlasCreateGCPEndpointGroupRequestView`

NewApiAtlasCreateGCPEndpointGroupRequestView instantiates a new ApiAtlasCreateGCPEndpointGroupRequestView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCreateGCPEndpointGroupRequestViewWithDefaults

`func NewApiAtlasCreateGCPEndpointGroupRequestViewWithDefaults() *ApiAtlasCreateGCPEndpointGroupRequestView`

NewApiAtlasCreateGCPEndpointGroupRequestViewWithDefaults instantiates a new ApiAtlasCreateGCPEndpointGroupRequestView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointGroupName

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetEndpoints

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) GetEndpoints() []ApiAtlasCreateGCPForwardingRuleRequestView`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) GetEndpointsOk() (*[]ApiAtlasCreateGCPForwardingRuleRequestView, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) SetEndpoints(v []ApiAtlasCreateGCPForwardingRuleRequestView)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetGcpProjectId

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *ApiAtlasCreateGCPEndpointGroupRequestView) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


