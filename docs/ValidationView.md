# ValidationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **map[string]interface{}** | List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.  This parameter returns an empty object if no custom zones exist. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Reason why the validation job failed. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project to validate. | [optional] [readonly] 
**SourceGroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the source project. | [optional] 
**Status** | Pointer to **string** | State of the specified validation job returned at the time of the request. | [optional] [readonly] 

## Methods

### NewValidationView

`func NewValidationView() *ValidationView`

NewValidationView instantiates a new ValidationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationViewWithDefaults

`func NewValidationViewWithDefaults() *ValidationView`

NewValidationViewWithDefaults instantiates a new ValidationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ValidationView) GetId() map[string]interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ValidationView) GetIdOk() (*map[string]interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ValidationView) SetId(v map[string]interface{})`

SetId sets Id field to given value.

### HasId

`func (o *ValidationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ValidationView) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ValidationView) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ValidationView) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ValidationView) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetGroupId

`func (o *ValidationView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ValidationView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ValidationView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ValidationView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSourceGroupId

`func (o *ValidationView) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *ValidationView) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *ValidationView) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.

### HasSourceGroupId

`func (o *ValidationView) HasSourceGroupId() bool`

HasSourceGroupId returns a boolean if a field has been set.

### GetStatus

`func (o *ValidationView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidationView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidationView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ValidationView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


