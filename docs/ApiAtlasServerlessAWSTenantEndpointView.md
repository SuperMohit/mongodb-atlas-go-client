# ApiAtlasServerlessAWSTenantEndpointView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the private endpoint. | [optional] [readonly] 
**CloudProviderEndpointId** | Pointer to **string** | Unique string that identifies the private endpoint&#39;s network interface. | [optional] [readonly] 
**Comment** | Pointer to **string** | Human-readable comment associated with the private endpoint. | [optional] [readonly] 
**EndpointServiceName** | Pointer to **string** | Unique string that identifies the Amazon Web Services (AWS) PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Human-readable error message that indicates error condition associated with establishing the private endpoint connection. | [optional] [readonly] 
**Status** | Pointer to **string** | Human-readable label that indicates the current operating status of the private endpoint. | [optional] [readonly] 

## Methods

### NewApiAtlasServerlessAWSTenantEndpointView

`func NewApiAtlasServerlessAWSTenantEndpointView() *ApiAtlasServerlessAWSTenantEndpointView`

NewApiAtlasServerlessAWSTenantEndpointView instantiates a new ApiAtlasServerlessAWSTenantEndpointView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasServerlessAWSTenantEndpointViewWithDefaults

`func NewApiAtlasServerlessAWSTenantEndpointViewWithDefaults() *ApiAtlasServerlessAWSTenantEndpointView`

NewApiAtlasServerlessAWSTenantEndpointViewWithDefaults instantiates a new ApiAtlasServerlessAWSTenantEndpointView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasServerlessAWSTenantEndpointView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasServerlessAWSTenantEndpointView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCloudProviderEndpointId

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetCloudProviderEndpointId() string`

GetCloudProviderEndpointId returns the CloudProviderEndpointId field if non-nil, zero value otherwise.

### GetCloudProviderEndpointIdOk

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetCloudProviderEndpointIdOk() (*string, bool)`

GetCloudProviderEndpointIdOk returns a tuple with the CloudProviderEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderEndpointId

`func (o *ApiAtlasServerlessAWSTenantEndpointView) SetCloudProviderEndpointId(v string)`

SetCloudProviderEndpointId sets CloudProviderEndpointId field to given value.

### HasCloudProviderEndpointId

`func (o *ApiAtlasServerlessAWSTenantEndpointView) HasCloudProviderEndpointId() bool`

HasCloudProviderEndpointId returns a boolean if a field has been set.

### GetComment

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiAtlasServerlessAWSTenantEndpointView) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiAtlasServerlessAWSTenantEndpointView) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEndpointServiceName

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetEndpointServiceName() string`

GetEndpointServiceName returns the EndpointServiceName field if non-nil, zero value otherwise.

### GetEndpointServiceNameOk

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetEndpointServiceNameOk() (*string, bool)`

GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointServiceName

`func (o *ApiAtlasServerlessAWSTenantEndpointView) SetEndpointServiceName(v string)`

SetEndpointServiceName sets EndpointServiceName field to given value.

### HasEndpointServiceName

`func (o *ApiAtlasServerlessAWSTenantEndpointView) HasEndpointServiceName() bool`

HasEndpointServiceName returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiAtlasServerlessAWSTenantEndpointView) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiAtlasServerlessAWSTenantEndpointView) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAtlasServerlessAWSTenantEndpointView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAtlasServerlessAWSTenantEndpointView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAtlasServerlessAWSTenantEndpointView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


