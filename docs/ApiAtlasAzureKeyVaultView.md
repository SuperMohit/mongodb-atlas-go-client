# ApiAtlasAzureKeyVaultView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureEnvironment** | Pointer to **string** | Azure environment in which your account credentials reside. | [optional] 
**ClientID** | Pointer to **string** | Unique 36-hexadecimal character string that identifies an Azure application associated with your Azure Active Directory tenant. | [optional] 
**Enabled** | Pointer to **bool** | Flag that indicates whether someone enabled encryption at rest for the specified  project. To disable encryption at rest using customer key management and remove the configuration details, pass only this parameter with a value of &#x60;false&#x60;. | [optional] 
**KeyIdentifier** | Pointer to **string** | Web address with a unique key that identifies for your Azure Key Vault. | [optional] 
**KeyVaultName** | Pointer to **string** | Unique string that identifies the Azure Key Vault that contains your key. | [optional] 
**ResourceGroupName** | Pointer to **string** | Name of the Azure resource group that contains your Azure Key Vault. | [optional] 
**Secret** | Pointer to **string** | Private data that you need secured and that belongs to the specified Azure Key Vault (AKV) tenant (**azureKeyVault.tenantID**). This data can include any type of sensitive data such as passwords, database connection strings, API keys, and the like. AKV stores this information as encrypted binary data. | [optional] 
**SubscriptionID** | Pointer to **string** | Unique 36-hexadecimal character string that identifies your Azure subscription. | [optional] 
**TenantID** | Pointer to **string** | Unique 36-hexadecimal character string that identifies the Azure Active Directory tenant within your Azure subscription. | [optional] 
**Valid** | Pointer to **bool** | Flag that indicates whether the Azure encryption key can encrypt and decrypt data. | [optional] [readonly] 

## Methods

### NewApiAtlasAzureKeyVaultView

`func NewApiAtlasAzureKeyVaultView() *ApiAtlasAzureKeyVaultView`

NewApiAtlasAzureKeyVaultView instantiates a new ApiAtlasAzureKeyVaultView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasAzureKeyVaultViewWithDefaults

`func NewApiAtlasAzureKeyVaultViewWithDefaults() *ApiAtlasAzureKeyVaultView`

NewApiAtlasAzureKeyVaultViewWithDefaults instantiates a new ApiAtlasAzureKeyVaultView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureEnvironment

`func (o *ApiAtlasAzureKeyVaultView) GetAzureEnvironment() string`

GetAzureEnvironment returns the AzureEnvironment field if non-nil, zero value otherwise.

### GetAzureEnvironmentOk

`func (o *ApiAtlasAzureKeyVaultView) GetAzureEnvironmentOk() (*string, bool)`

GetAzureEnvironmentOk returns a tuple with the AzureEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureEnvironment

`func (o *ApiAtlasAzureKeyVaultView) SetAzureEnvironment(v string)`

SetAzureEnvironment sets AzureEnvironment field to given value.

### HasAzureEnvironment

`func (o *ApiAtlasAzureKeyVaultView) HasAzureEnvironment() bool`

HasAzureEnvironment returns a boolean if a field has been set.

### GetClientID

`func (o *ApiAtlasAzureKeyVaultView) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *ApiAtlasAzureKeyVaultView) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *ApiAtlasAzureKeyVaultView) SetClientID(v string)`

SetClientID sets ClientID field to given value.

### HasClientID

`func (o *ApiAtlasAzureKeyVaultView) HasClientID() bool`

HasClientID returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiAtlasAzureKeyVaultView) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiAtlasAzureKeyVaultView) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiAtlasAzureKeyVaultView) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiAtlasAzureKeyVaultView) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetKeyIdentifier

`func (o *ApiAtlasAzureKeyVaultView) GetKeyIdentifier() string`

GetKeyIdentifier returns the KeyIdentifier field if non-nil, zero value otherwise.

### GetKeyIdentifierOk

`func (o *ApiAtlasAzureKeyVaultView) GetKeyIdentifierOk() (*string, bool)`

GetKeyIdentifierOk returns a tuple with the KeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIdentifier

`func (o *ApiAtlasAzureKeyVaultView) SetKeyIdentifier(v string)`

SetKeyIdentifier sets KeyIdentifier field to given value.

### HasKeyIdentifier

`func (o *ApiAtlasAzureKeyVaultView) HasKeyIdentifier() bool`

HasKeyIdentifier returns a boolean if a field has been set.

### GetKeyVaultName

`func (o *ApiAtlasAzureKeyVaultView) GetKeyVaultName() string`

GetKeyVaultName returns the KeyVaultName field if non-nil, zero value otherwise.

### GetKeyVaultNameOk

`func (o *ApiAtlasAzureKeyVaultView) GetKeyVaultNameOk() (*string, bool)`

GetKeyVaultNameOk returns a tuple with the KeyVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultName

`func (o *ApiAtlasAzureKeyVaultView) SetKeyVaultName(v string)`

SetKeyVaultName sets KeyVaultName field to given value.

### HasKeyVaultName

`func (o *ApiAtlasAzureKeyVaultView) HasKeyVaultName() bool`

HasKeyVaultName returns a boolean if a field has been set.

### GetResourceGroupName

`func (o *ApiAtlasAzureKeyVaultView) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *ApiAtlasAzureKeyVaultView) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *ApiAtlasAzureKeyVaultView) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *ApiAtlasAzureKeyVaultView) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### GetSecret

`func (o *ApiAtlasAzureKeyVaultView) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ApiAtlasAzureKeyVaultView) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ApiAtlasAzureKeyVaultView) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ApiAtlasAzureKeyVaultView) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSubscriptionID

`func (o *ApiAtlasAzureKeyVaultView) GetSubscriptionID() string`

GetSubscriptionID returns the SubscriptionID field if non-nil, zero value otherwise.

### GetSubscriptionIDOk

`func (o *ApiAtlasAzureKeyVaultView) GetSubscriptionIDOk() (*string, bool)`

GetSubscriptionIDOk returns a tuple with the SubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionID

`func (o *ApiAtlasAzureKeyVaultView) SetSubscriptionID(v string)`

SetSubscriptionID sets SubscriptionID field to given value.

### HasSubscriptionID

`func (o *ApiAtlasAzureKeyVaultView) HasSubscriptionID() bool`

HasSubscriptionID returns a boolean if a field has been set.

### GetTenantID

`func (o *ApiAtlasAzureKeyVaultView) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *ApiAtlasAzureKeyVaultView) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *ApiAtlasAzureKeyVaultView) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.

### HasTenantID

`func (o *ApiAtlasAzureKeyVaultView) HasTenantID() bool`

HasTenantID returns a boolean if a field has been set.

### GetValid

`func (o *ApiAtlasAzureKeyVaultView) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ApiAtlasAzureKeyVaultView) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ApiAtlasAzureKeyVaultView) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *ApiAtlasAzureKeyVaultView) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


