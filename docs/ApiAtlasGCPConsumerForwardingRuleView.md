# ApiAtlasGCPConsumerForwardingRuleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointName** | Pointer to **string** | Human-readable label that identifies the Google Cloud consumer forwarding rule that you created. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | One Private Internet Protocol version 4 (IPv4) address to which this Google Cloud consumer forwarding rule resolves. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the MongoDB Cloud endpoint group when MongoDB Cloud received this request. | [optional] [readonly] 

## Methods

### NewApiAtlasGCPConsumerForwardingRuleView

`func NewApiAtlasGCPConsumerForwardingRuleView() *ApiAtlasGCPConsumerForwardingRuleView`

NewApiAtlasGCPConsumerForwardingRuleView instantiates a new ApiAtlasGCPConsumerForwardingRuleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasGCPConsumerForwardingRuleViewWithDefaults

`func NewApiAtlasGCPConsumerForwardingRuleViewWithDefaults() *ApiAtlasGCPConsumerForwardingRuleView`

NewApiAtlasGCPConsumerForwardingRuleViewWithDefaults instantiates a new ApiAtlasGCPConsumerForwardingRuleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointName

`func (o *ApiAtlasGCPConsumerForwardingRuleView) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *ApiAtlasGCPConsumerForwardingRuleView) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *ApiAtlasGCPConsumerForwardingRuleView) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *ApiAtlasGCPConsumerForwardingRuleView) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetIpAddress

`func (o *ApiAtlasGCPConsumerForwardingRuleView) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ApiAtlasGCPConsumerForwardingRuleView) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ApiAtlasGCPConsumerForwardingRuleView) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ApiAtlasGCPConsumerForwardingRuleView) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetStatus

`func (o *ApiAtlasGCPConsumerForwardingRuleView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAtlasGCPConsumerForwardingRuleView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAtlasGCPConsumerForwardingRuleView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAtlasGCPConsumerForwardingRuleView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


