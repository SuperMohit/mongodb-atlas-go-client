# CreateOnePrivateEndpointServiceForOneProvider200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointServiceName** | Pointer to **string** | Unique string that identifies the Amazon Web Services (AWS) PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Private Endpoint Service. | [optional] [readonly] 
**InterfaceEndpoints** | Pointer to **[]string** | List of strings that identify private endpoint interfaces applied to the specified project. | [optional] [readonly] 
**RegionName** | Pointer to **string** | Cloud provider region that manages this Private Endpoint Service. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the Private Endpoint Service connection when MongoDB Cloud received this request. | [optional] [readonly] 
**PrivateEndpoints** | Pointer to **[]string** | List of private endpoints assigned to this Azure Private Link Service. | [optional] [readonly] 
**PrivateLinkServiceName** | Pointer to **string** | Unique string that identifies the Azure Private Link Service that MongoDB Cloud manages. | [optional] [readonly] 
**PrivateLinkServiceResourceId** | Pointer to **string** | Root-relative path that identifies of the Azure Private Link Service that MongoDB Cloud manages. Use this value to create a private endpoint connection to an Azure VNet. | [optional] [readonly] 
**EndpointGroupNames** | Pointer to **[]string** | List of Google Cloud network endpoint groups that corresponds to the Private Service Connect endpoint service. | [optional] 
**ServiceAttachmentNames** | Pointer to **[]string** | List of Uniform Resource Locators (URLs) that identifies endpoints that MongoDB Cloud can use to access one Google Cloud Service across a Google Cloud Virtual Private Connection (VPC) network. | [optional] 

## Methods

### NewCreateOnePrivateEndpointServiceForOneProvider200Response

`func NewCreateOnePrivateEndpointServiceForOneProvider200Response() *CreateOnePrivateEndpointServiceForOneProvider200Response`

NewCreateOnePrivateEndpointServiceForOneProvider200Response instantiates a new CreateOnePrivateEndpointServiceForOneProvider200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOnePrivateEndpointServiceForOneProvider200ResponseWithDefaults

`func NewCreateOnePrivateEndpointServiceForOneProvider200ResponseWithDefaults() *CreateOnePrivateEndpointServiceForOneProvider200Response`

NewCreateOnePrivateEndpointServiceForOneProvider200ResponseWithDefaults instantiates a new CreateOnePrivateEndpointServiceForOneProvider200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointServiceName

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetEndpointServiceName() string`

GetEndpointServiceName returns the EndpointServiceName field if non-nil, zero value otherwise.

### GetEndpointServiceNameOk

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetEndpointServiceNameOk() (*string, bool)`

GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointServiceName

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) SetEndpointServiceName(v string)`

SetEndpointServiceName sets EndpointServiceName field to given value.

### HasEndpointServiceName

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) HasEndpointServiceName() bool`

HasEndpointServiceName returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetId

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceEndpoints

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetInterfaceEndpoints() []string`

GetInterfaceEndpoints returns the InterfaceEndpoints field if non-nil, zero value otherwise.

### GetInterfaceEndpointsOk

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetInterfaceEndpointsOk() (*[]string, bool)`

GetInterfaceEndpointsOk returns a tuple with the InterfaceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceEndpoints

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) SetInterfaceEndpoints(v []string)`

SetInterfaceEndpoints sets InterfaceEndpoints field to given value.

### HasInterfaceEndpoints

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) HasInterfaceEndpoints() bool`

HasInterfaceEndpoints returns a boolean if a field has been set.

### GetRegionName

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetStatus

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrivateEndpoints

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetPrivateEndpoints() []string`

GetPrivateEndpoints returns the PrivateEndpoints field if non-nil, zero value otherwise.

### GetPrivateEndpointsOk

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetPrivateEndpointsOk() (*[]string, bool)`

GetPrivateEndpointsOk returns a tuple with the PrivateEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpoints

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) SetPrivateEndpoints(v []string)`

SetPrivateEndpoints sets PrivateEndpoints field to given value.

### HasPrivateEndpoints

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) HasPrivateEndpoints() bool`

HasPrivateEndpoints returns a boolean if a field has been set.

### GetPrivateLinkServiceName

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetPrivateLinkServiceName() string`

GetPrivateLinkServiceName returns the PrivateLinkServiceName field if non-nil, zero value otherwise.

### GetPrivateLinkServiceNameOk

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetPrivateLinkServiceNameOk() (*string, bool)`

GetPrivateLinkServiceNameOk returns a tuple with the PrivateLinkServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceName

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) SetPrivateLinkServiceName(v string)`

SetPrivateLinkServiceName sets PrivateLinkServiceName field to given value.

### HasPrivateLinkServiceName

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) HasPrivateLinkServiceName() bool`

HasPrivateLinkServiceName returns a boolean if a field has been set.

### GetPrivateLinkServiceResourceId

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.

### HasPrivateLinkServiceResourceId

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) HasPrivateLinkServiceResourceId() bool`

HasPrivateLinkServiceResourceId returns a boolean if a field has been set.

### GetEndpointGroupNames

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetEndpointGroupNames() []string`

GetEndpointGroupNames returns the EndpointGroupNames field if non-nil, zero value otherwise.

### GetEndpointGroupNamesOk

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetEndpointGroupNamesOk() (*[]string, bool)`

GetEndpointGroupNamesOk returns a tuple with the EndpointGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupNames

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) SetEndpointGroupNames(v []string)`

SetEndpointGroupNames sets EndpointGroupNames field to given value.

### HasEndpointGroupNames

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) HasEndpointGroupNames() bool`

HasEndpointGroupNames returns a boolean if a field has been set.

### GetServiceAttachmentNames

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetServiceAttachmentNames() []string`

GetServiceAttachmentNames returns the ServiceAttachmentNames field if non-nil, zero value otherwise.

### GetServiceAttachmentNamesOk

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) GetServiceAttachmentNamesOk() (*[]string, bool)`

GetServiceAttachmentNamesOk returns a tuple with the ServiceAttachmentNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttachmentNames

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) SetServiceAttachmentNames(v []string)`

SetServiceAttachmentNames sets ServiceAttachmentNames field to given value.

### HasServiceAttachmentNames

`func (o *CreateOnePrivateEndpointServiceForOneProvider200Response) HasServiceAttachmentNames() bool`

HasServiceAttachmentNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


