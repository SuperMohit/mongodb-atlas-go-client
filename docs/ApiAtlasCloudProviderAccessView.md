# ApiAtlasCloudProviderAccessView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsIamRoles** | Pointer to [**[]ApiAtlasCloudProviderAccessAWSIAMRoleView**](ApiAtlasCloudProviderAccessAWSIAMRoleView.md) | List that contains the Amazon Web Services (AWS) IAM roles registered and authorized with MongoDB Cloud. | [optional] 

## Methods

### NewApiAtlasCloudProviderAccessView

`func NewApiAtlasCloudProviderAccessView() *ApiAtlasCloudProviderAccessView`

NewApiAtlasCloudProviderAccessView instantiates a new ApiAtlasCloudProviderAccessView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCloudProviderAccessViewWithDefaults

`func NewApiAtlasCloudProviderAccessViewWithDefaults() *ApiAtlasCloudProviderAccessView`

NewApiAtlasCloudProviderAccessViewWithDefaults instantiates a new ApiAtlasCloudProviderAccessView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsIamRoles

`func (o *ApiAtlasCloudProviderAccessView) GetAwsIamRoles() []ApiAtlasCloudProviderAccessAWSIAMRoleView`

GetAwsIamRoles returns the AwsIamRoles field if non-nil, zero value otherwise.

### GetAwsIamRolesOk

`func (o *ApiAtlasCloudProviderAccessView) GetAwsIamRolesOk() (*[]ApiAtlasCloudProviderAccessAWSIAMRoleView, bool)`

GetAwsIamRolesOk returns a tuple with the AwsIamRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIamRoles

`func (o *ApiAtlasCloudProviderAccessView) SetAwsIamRoles(v []ApiAtlasCloudProviderAccessAWSIAMRoleView)`

SetAwsIamRoles sets AwsIamRoles field to given value.

### HasAwsIamRoles

`func (o *ApiAtlasCloudProviderAccessView) HasAwsIamRoles() bool`

HasAwsIamRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


