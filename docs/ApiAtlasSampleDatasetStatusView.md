# ApiAtlasSampleDatasetStatusView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies this sample dataset. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster into which you loaded the sample dataset. | [optional] [readonly] 
**CompleteDate** | Pointer to **time.Time** | Date and time when the sample dataset load job completed. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**CreateDate** | Pointer to **time.Time** | Date and time when you started the sample dataset load job. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Details of the error returned when MongoDB Cloud loads the sample dataset. This endpoint returns null if state has a value other than FAILED. | [optional] [readonly] 
**State** | Pointer to **string** | Status of the sample dataset load job. | [optional] [readonly] 

## Methods

### NewApiAtlasSampleDatasetStatusView

`func NewApiAtlasSampleDatasetStatusView() *ApiAtlasSampleDatasetStatusView`

NewApiAtlasSampleDatasetStatusView instantiates a new ApiAtlasSampleDatasetStatusView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasSampleDatasetStatusViewWithDefaults

`func NewApiAtlasSampleDatasetStatusViewWithDefaults() *ApiAtlasSampleDatasetStatusView`

NewApiAtlasSampleDatasetStatusViewWithDefaults instantiates a new ApiAtlasSampleDatasetStatusView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiAtlasSampleDatasetStatusView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasSampleDatasetStatusView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasSampleDatasetStatusView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasSampleDatasetStatusView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClusterName

`func (o *ApiAtlasSampleDatasetStatusView) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ApiAtlasSampleDatasetStatusView) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ApiAtlasSampleDatasetStatusView) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ApiAtlasSampleDatasetStatusView) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCompleteDate

`func (o *ApiAtlasSampleDatasetStatusView) GetCompleteDate() time.Time`

GetCompleteDate returns the CompleteDate field if non-nil, zero value otherwise.

### GetCompleteDateOk

`func (o *ApiAtlasSampleDatasetStatusView) GetCompleteDateOk() (*time.Time, bool)`

GetCompleteDateOk returns a tuple with the CompleteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteDate

`func (o *ApiAtlasSampleDatasetStatusView) SetCompleteDate(v time.Time)`

SetCompleteDate sets CompleteDate field to given value.

### HasCompleteDate

`func (o *ApiAtlasSampleDatasetStatusView) HasCompleteDate() bool`

HasCompleteDate returns a boolean if a field has been set.

### GetCreateDate

`func (o *ApiAtlasSampleDatasetStatusView) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ApiAtlasSampleDatasetStatusView) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ApiAtlasSampleDatasetStatusView) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ApiAtlasSampleDatasetStatusView) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ApiAtlasSampleDatasetStatusView) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiAtlasSampleDatasetStatusView) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiAtlasSampleDatasetStatusView) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiAtlasSampleDatasetStatusView) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetState

`func (o *ApiAtlasSampleDatasetStatusView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiAtlasSampleDatasetStatusView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiAtlasSampleDatasetStatusView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiAtlasSampleDatasetStatusView) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


