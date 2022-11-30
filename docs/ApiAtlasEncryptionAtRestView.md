# ApiAtlasEncryptionAtRestView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKms** | Pointer to [**ApiAtlasAWSKMSView**](ApiAtlasAWSKMSView.md) |  | [optional] 
**AzureKeyVault** | Pointer to [**ApiAtlasAzureKeyVaultView**](ApiAtlasAzureKeyVaultView.md) |  | [optional] 
**GoogleCloudKms** | Pointer to [**ApiAtlasGoogleCloudKMSView**](ApiAtlasGoogleCloudKMSView.md) |  | [optional] 

## Methods

### NewApiAtlasEncryptionAtRestView

`func NewApiAtlasEncryptionAtRestView() *ApiAtlasEncryptionAtRestView`

NewApiAtlasEncryptionAtRestView instantiates a new ApiAtlasEncryptionAtRestView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasEncryptionAtRestViewWithDefaults

`func NewApiAtlasEncryptionAtRestViewWithDefaults() *ApiAtlasEncryptionAtRestView`

NewApiAtlasEncryptionAtRestViewWithDefaults instantiates a new ApiAtlasEncryptionAtRestView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKms

`func (o *ApiAtlasEncryptionAtRestView) GetAwsKms() ApiAtlasAWSKMSView`

GetAwsKms returns the AwsKms field if non-nil, zero value otherwise.

### GetAwsKmsOk

`func (o *ApiAtlasEncryptionAtRestView) GetAwsKmsOk() (*ApiAtlasAWSKMSView, bool)`

GetAwsKmsOk returns a tuple with the AwsKms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKms

`func (o *ApiAtlasEncryptionAtRestView) SetAwsKms(v ApiAtlasAWSKMSView)`

SetAwsKms sets AwsKms field to given value.

### HasAwsKms

`func (o *ApiAtlasEncryptionAtRestView) HasAwsKms() bool`

HasAwsKms returns a boolean if a field has been set.

### GetAzureKeyVault

`func (o *ApiAtlasEncryptionAtRestView) GetAzureKeyVault() ApiAtlasAzureKeyVaultView`

GetAzureKeyVault returns the AzureKeyVault field if non-nil, zero value otherwise.

### GetAzureKeyVaultOk

`func (o *ApiAtlasEncryptionAtRestView) GetAzureKeyVaultOk() (*ApiAtlasAzureKeyVaultView, bool)`

GetAzureKeyVaultOk returns a tuple with the AzureKeyVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureKeyVault

`func (o *ApiAtlasEncryptionAtRestView) SetAzureKeyVault(v ApiAtlasAzureKeyVaultView)`

SetAzureKeyVault sets AzureKeyVault field to given value.

### HasAzureKeyVault

`func (o *ApiAtlasEncryptionAtRestView) HasAzureKeyVault() bool`

HasAzureKeyVault returns a boolean if a field has been set.

### GetGoogleCloudKms

`func (o *ApiAtlasEncryptionAtRestView) GetGoogleCloudKms() ApiAtlasGoogleCloudKMSView`

GetGoogleCloudKms returns the GoogleCloudKms field if non-nil, zero value otherwise.

### GetGoogleCloudKmsOk

`func (o *ApiAtlasEncryptionAtRestView) GetGoogleCloudKmsOk() (*ApiAtlasGoogleCloudKMSView, bool)`

GetGoogleCloudKmsOk returns a tuple with the GoogleCloudKms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCloudKms

`func (o *ApiAtlasEncryptionAtRestView) SetGoogleCloudKms(v ApiAtlasGoogleCloudKMSView)`

SetGoogleCloudKms sets GoogleCloudKms field to given value.

### HasGoogleCloudKms

`func (o *ApiAtlasEncryptionAtRestView) HasGoogleCloudKms() bool`

HasGoogleCloudKms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


