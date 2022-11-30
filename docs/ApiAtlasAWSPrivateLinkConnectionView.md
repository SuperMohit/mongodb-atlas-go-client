# ApiAtlasAWSPrivateLinkConnectionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointServiceName** | Pointer to **string** | Unique string that identifies the Amazon Web Services (AWS) PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Private Endpoint Service. | [optional] [readonly] 
**InterfaceEndpoints** | Pointer to **[]string** | List of strings that identify private endpoint interfaces applied to the specified project. | [optional] [readonly] 
**RegionName** | Pointer to **string** | Cloud provider region that manages this Private Endpoint Service. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the Private Endpoint Service connection when MongoDB Cloud received this request. | [optional] [readonly] 

## Methods

### NewApiAtlasAWSPrivateLinkConnectionView

`func NewApiAtlasAWSPrivateLinkConnectionView() *ApiAtlasAWSPrivateLinkConnectionView`

NewApiAtlasAWSPrivateLinkConnectionView instantiates a new ApiAtlasAWSPrivateLinkConnectionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasAWSPrivateLinkConnectionViewWithDefaults

`func NewApiAtlasAWSPrivateLinkConnectionViewWithDefaults() *ApiAtlasAWSPrivateLinkConnectionView`

NewApiAtlasAWSPrivateLinkConnectionViewWithDefaults instantiates a new ApiAtlasAWSPrivateLinkConnectionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointServiceName

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetEndpointServiceName() string`

GetEndpointServiceName returns the EndpointServiceName field if non-nil, zero value otherwise.

### GetEndpointServiceNameOk

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetEndpointServiceNameOk() (*string, bool)`

GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointServiceName

`func (o *ApiAtlasAWSPrivateLinkConnectionView) SetEndpointServiceName(v string)`

SetEndpointServiceName sets EndpointServiceName field to given value.

### HasEndpointServiceName

`func (o *ApiAtlasAWSPrivateLinkConnectionView) HasEndpointServiceName() bool`

HasEndpointServiceName returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiAtlasAWSPrivateLinkConnectionView) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiAtlasAWSPrivateLinkConnectionView) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasAWSPrivateLinkConnectionView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasAWSPrivateLinkConnectionView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceEndpoints

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetInterfaceEndpoints() []string`

GetInterfaceEndpoints returns the InterfaceEndpoints field if non-nil, zero value otherwise.

### GetInterfaceEndpointsOk

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetInterfaceEndpointsOk() (*[]string, bool)`

GetInterfaceEndpointsOk returns a tuple with the InterfaceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceEndpoints

`func (o *ApiAtlasAWSPrivateLinkConnectionView) SetInterfaceEndpoints(v []string)`

SetInterfaceEndpoints sets InterfaceEndpoints field to given value.

### HasInterfaceEndpoints

`func (o *ApiAtlasAWSPrivateLinkConnectionView) HasInterfaceEndpoints() bool`

HasInterfaceEndpoints returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasAWSPrivateLinkConnectionView) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasAWSPrivateLinkConnectionView) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetStatus

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAtlasAWSPrivateLinkConnectionView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAtlasAWSPrivateLinkConnectionView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAtlasAWSPrivateLinkConnectionView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


