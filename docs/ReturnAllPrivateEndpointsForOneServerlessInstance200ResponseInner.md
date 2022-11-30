# ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the private endpoint. | [optional] [readonly] 
**CloudProviderEndpointId** | Pointer to **string** | Unique string that identifies the Azure private endpoint&#39;s network interface that someone added to this private endpoint service. | [optional] [readonly] 
**Comment** | Pointer to **string** | Human-readable comment associated with the private endpoint. | [optional] [readonly] 
**EndpointServiceName** | Pointer to **string** | Unique string that identifies the Azure private endpoint service. MongoDB Cloud returns null while it creates the endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Human-readable error message that indicates error condition associated with establishing the private endpoint connection. | [optional] [readonly] 
**Status** | Pointer to **string** | Human-readable label that indicates the current operating status of the private endpoint. | [optional] [readonly] 
**PrivateEndpointIpAddress** | Pointer to **string** | IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service. | [optional] [readonly] 
**PrivateLinkServiceResourceId** | Pointer to **string** | Root-relative path that identifies the Azure Private Link Service that MongoDB Cloud manages. MongoDB Cloud returns null while it creates the endpoint service. | [optional] [readonly] 

## Methods

### NewReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner

`func NewReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner() *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner`

NewReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner instantiates a new ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInnerWithDefaults

`func NewReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInnerWithDefaults() *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner`

NewReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInnerWithDefaults instantiates a new ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCloudProviderEndpointId

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetCloudProviderEndpointId() string`

GetCloudProviderEndpointId returns the CloudProviderEndpointId field if non-nil, zero value otherwise.

### GetCloudProviderEndpointIdOk

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetCloudProviderEndpointIdOk() (*string, bool)`

GetCloudProviderEndpointIdOk returns a tuple with the CloudProviderEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderEndpointId

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) SetCloudProviderEndpointId(v string)`

SetCloudProviderEndpointId sets CloudProviderEndpointId field to given value.

### HasCloudProviderEndpointId

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) HasCloudProviderEndpointId() bool`

HasCloudProviderEndpointId returns a boolean if a field has been set.

### GetComment

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEndpointServiceName

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetEndpointServiceName() string`

GetEndpointServiceName returns the EndpointServiceName field if non-nil, zero value otherwise.

### GetEndpointServiceNameOk

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetEndpointServiceNameOk() (*string, bool)`

GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointServiceName

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) SetEndpointServiceName(v string)`

SetEndpointServiceName sets EndpointServiceName field to given value.

### HasEndpointServiceName

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) HasEndpointServiceName() bool`

HasEndpointServiceName returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrivateEndpointIpAddress

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetPrivateEndpointIpAddress() string`

GetPrivateEndpointIpAddress returns the PrivateEndpointIpAddress field if non-nil, zero value otherwise.

### GetPrivateEndpointIpAddressOk

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetPrivateEndpointIpAddressOk() (*string, bool)`

GetPrivateEndpointIpAddressOk returns a tuple with the PrivateEndpointIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointIpAddress

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) SetPrivateEndpointIpAddress(v string)`

SetPrivateEndpointIpAddress sets PrivateEndpointIpAddress field to given value.

### HasPrivateEndpointIpAddress

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) HasPrivateEndpointIpAddress() bool`

HasPrivateEndpointIpAddress returns a boolean if a field has been set.

### GetPrivateLinkServiceResourceId

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.

### HasPrivateLinkServiceResourceId

`func (o *ReturnAllPrivateEndpointsForOneServerlessInstance200ResponseInner) HasPrivateLinkServiceResourceId() bool`

HasPrivateLinkServiceResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


