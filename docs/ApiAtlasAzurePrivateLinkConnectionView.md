# ApiAtlasAzurePrivateLinkConnectionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Private Endpoint Service. | [optional] [readonly] 
**PrivateEndpoints** | Pointer to **[]string** | List of private endpoints assigned to this Azure Private Link Service. | [optional] [readonly] 
**PrivateLinkServiceName** | Pointer to **string** | Unique string that identifies the Azure Private Link Service that MongoDB Cloud manages. | [optional] [readonly] 
**PrivateLinkServiceResourceId** | Pointer to **string** | Root-relative path that identifies of the Azure Private Link Service that MongoDB Cloud manages. Use this value to create a private endpoint connection to an Azure VNet. | [optional] [readonly] 
**RegionName** | Pointer to **string** | Cloud provider region that manages this Private Endpoint Service. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the Private Endpoint Service connection when MongoDB Cloud received this request. | [optional] [readonly] 

## Methods

### NewApiAtlasAzurePrivateLinkConnectionView

`func NewApiAtlasAzurePrivateLinkConnectionView() *ApiAtlasAzurePrivateLinkConnectionView`

NewApiAtlasAzurePrivateLinkConnectionView instantiates a new ApiAtlasAzurePrivateLinkConnectionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasAzurePrivateLinkConnectionViewWithDefaults

`func NewApiAtlasAzurePrivateLinkConnectionViewWithDefaults() *ApiAtlasAzurePrivateLinkConnectionView`

NewApiAtlasAzurePrivateLinkConnectionViewWithDefaults instantiates a new ApiAtlasAzurePrivateLinkConnectionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiAtlasAzurePrivateLinkConnectionView) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiAtlasAzurePrivateLinkConnectionView) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasAzurePrivateLinkConnectionView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasAzurePrivateLinkConnectionView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrivateEndpoints

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetPrivateEndpoints() []string`

GetPrivateEndpoints returns the PrivateEndpoints field if non-nil, zero value otherwise.

### GetPrivateEndpointsOk

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetPrivateEndpointsOk() (*[]string, bool)`

GetPrivateEndpointsOk returns a tuple with the PrivateEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpoints

`func (o *ApiAtlasAzurePrivateLinkConnectionView) SetPrivateEndpoints(v []string)`

SetPrivateEndpoints sets PrivateEndpoints field to given value.

### HasPrivateEndpoints

`func (o *ApiAtlasAzurePrivateLinkConnectionView) HasPrivateEndpoints() bool`

HasPrivateEndpoints returns a boolean if a field has been set.

### GetPrivateLinkServiceName

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetPrivateLinkServiceName() string`

GetPrivateLinkServiceName returns the PrivateLinkServiceName field if non-nil, zero value otherwise.

### GetPrivateLinkServiceNameOk

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetPrivateLinkServiceNameOk() (*string, bool)`

GetPrivateLinkServiceNameOk returns a tuple with the PrivateLinkServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceName

`func (o *ApiAtlasAzurePrivateLinkConnectionView) SetPrivateLinkServiceName(v string)`

SetPrivateLinkServiceName sets PrivateLinkServiceName field to given value.

### HasPrivateLinkServiceName

`func (o *ApiAtlasAzurePrivateLinkConnectionView) HasPrivateLinkServiceName() bool`

HasPrivateLinkServiceName returns a boolean if a field has been set.

### GetPrivateLinkServiceResourceId

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *ApiAtlasAzurePrivateLinkConnectionView) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.

### HasPrivateLinkServiceResourceId

`func (o *ApiAtlasAzurePrivateLinkConnectionView) HasPrivateLinkServiceResourceId() bool`

HasPrivateLinkServiceResourceId returns a boolean if a field has been set.

### GetRegionName

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiAtlasAzurePrivateLinkConnectionView) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiAtlasAzurePrivateLinkConnectionView) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetStatus

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAtlasAzurePrivateLinkConnectionView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAtlasAzurePrivateLinkConnectionView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAtlasAzurePrivateLinkConnectionView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


