# ApiAtlasAWSInterfaceEndpointView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStatus** | Pointer to **string** | State of the Amazon Web Service PrivateLink connection when MongoDB Cloud received this request. | [optional] [readonly] 
**DeleteRequested** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud received a request to remove the specified private endpoint from the private endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**InterfaceEndpointId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the interface endpoint. | [optional] [readonly] 

## Methods

### NewApiAtlasAWSInterfaceEndpointView

`func NewApiAtlasAWSInterfaceEndpointView() *ApiAtlasAWSInterfaceEndpointView`

NewApiAtlasAWSInterfaceEndpointView instantiates a new ApiAtlasAWSInterfaceEndpointView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasAWSInterfaceEndpointViewWithDefaults

`func NewApiAtlasAWSInterfaceEndpointViewWithDefaults() *ApiAtlasAWSInterfaceEndpointView`

NewApiAtlasAWSInterfaceEndpointViewWithDefaults instantiates a new ApiAtlasAWSInterfaceEndpointView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatus

`func (o *ApiAtlasAWSInterfaceEndpointView) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *ApiAtlasAWSInterfaceEndpointView) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *ApiAtlasAWSInterfaceEndpointView) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *ApiAtlasAWSInterfaceEndpointView) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDeleteRequested

`func (o *ApiAtlasAWSInterfaceEndpointView) GetDeleteRequested() bool`

GetDeleteRequested returns the DeleteRequested field if non-nil, zero value otherwise.

### GetDeleteRequestedOk

`func (o *ApiAtlasAWSInterfaceEndpointView) GetDeleteRequestedOk() (*bool, bool)`

GetDeleteRequestedOk returns a tuple with the DeleteRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRequested

`func (o *ApiAtlasAWSInterfaceEndpointView) SetDeleteRequested(v bool)`

SetDeleteRequested sets DeleteRequested field to given value.

### HasDeleteRequested

`func (o *ApiAtlasAWSInterfaceEndpointView) HasDeleteRequested() bool`

HasDeleteRequested returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ApiAtlasAWSInterfaceEndpointView) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiAtlasAWSInterfaceEndpointView) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiAtlasAWSInterfaceEndpointView) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiAtlasAWSInterfaceEndpointView) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetInterfaceEndpointId

`func (o *ApiAtlasAWSInterfaceEndpointView) GetInterfaceEndpointId() string`

GetInterfaceEndpointId returns the InterfaceEndpointId field if non-nil, zero value otherwise.

### GetInterfaceEndpointIdOk

`func (o *ApiAtlasAWSInterfaceEndpointView) GetInterfaceEndpointIdOk() (*string, bool)`

GetInterfaceEndpointIdOk returns a tuple with the InterfaceEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceEndpointId

`func (o *ApiAtlasAWSInterfaceEndpointView) SetInterfaceEndpointId(v string)`

SetInterfaceEndpointId sets InterfaceEndpointId field to given value.

### HasInterfaceEndpointId

`func (o *ApiAtlasAWSInterfaceEndpointView) HasInterfaceEndpointId() bool`

HasInterfaceEndpointId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


