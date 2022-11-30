# ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtlasAWSAccountArn** | Pointer to **string** | Amazon Resource Name that identifies the Amazon Web Services (AWS) user account that MongoDB Cloud uses when it assumes the Identity and Access Management (IAM) role. | [optional] [readonly] 
**AtlasAssumedRoleExternalId** | Pointer to **string** | Unique external ID that MongoDB Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account. | [optional] [readonly] 
**AuthorizedDate** | Pointer to **time.Time** | Date and time when someone authorized this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**CreatedDate** | Pointer to **time.Time** | Date and time when someone created this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FeatureUsages** | Pointer to [**[]ApiAtlasCloudProviderAccessFeatureUsageView**](ApiAtlasCloudProviderAccessFeatureUsageView.md) | List that contains application features associated with this Amazon Web Services (AWS) Identity and Access Management (IAM) role. | [optional] [readonly] 
**IamAssumedRoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that MongoDB Cloud assumes when it accesses resources in your AWS account. | [optional] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the role. | [optional] [readonly] 

## Methods

### NewApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf

`func NewApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf() *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf`

NewApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf instantiates a new ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCloudProviderAccessAWSIAMRoleViewAllOfWithDefaults

`func NewApiAtlasCloudProviderAccessAWSIAMRoleViewAllOfWithDefaults() *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf`

NewApiAtlasCloudProviderAccessAWSIAMRoleViewAllOfWithDefaults instantiates a new ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtlasAWSAccountArn

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAtlasAWSAccountArn() string`

GetAtlasAWSAccountArn returns the AtlasAWSAccountArn field if non-nil, zero value otherwise.

### GetAtlasAWSAccountArnOk

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAtlasAWSAccountArnOk() (*string, bool)`

GetAtlasAWSAccountArnOk returns a tuple with the AtlasAWSAccountArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAWSAccountArn

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetAtlasAWSAccountArn(v string)`

SetAtlasAWSAccountArn sets AtlasAWSAccountArn field to given value.

### HasAtlasAWSAccountArn

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasAtlasAWSAccountArn() bool`

HasAtlasAWSAccountArn returns a boolean if a field has been set.

### GetAtlasAssumedRoleExternalId

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAtlasAssumedRoleExternalId() string`

GetAtlasAssumedRoleExternalId returns the AtlasAssumedRoleExternalId field if non-nil, zero value otherwise.

### GetAtlasAssumedRoleExternalIdOk

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAtlasAssumedRoleExternalIdOk() (*string, bool)`

GetAtlasAssumedRoleExternalIdOk returns a tuple with the AtlasAssumedRoleExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasAssumedRoleExternalId

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetAtlasAssumedRoleExternalId(v string)`

SetAtlasAssumedRoleExternalId sets AtlasAssumedRoleExternalId field to given value.

### HasAtlasAssumedRoleExternalId

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasAtlasAssumedRoleExternalId() bool`

HasAtlasAssumedRoleExternalId returns a boolean if a field has been set.

### GetAuthorizedDate

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAuthorizedDate() time.Time`

GetAuthorizedDate returns the AuthorizedDate field if non-nil, zero value otherwise.

### GetAuthorizedDateOk

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetAuthorizedDateOk() (*time.Time, bool)`

GetAuthorizedDateOk returns a tuple with the AuthorizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedDate

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetAuthorizedDate(v time.Time)`

SetAuthorizedDate sets AuthorizedDate field to given value.

### HasAuthorizedDate

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasAuthorizedDate() bool`

HasAuthorizedDate returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetFeatureUsages

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetFeatureUsages() []ApiAtlasCloudProviderAccessFeatureUsageView`

GetFeatureUsages returns the FeatureUsages field if non-nil, zero value otherwise.

### GetFeatureUsagesOk

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetFeatureUsagesOk() (*[]ApiAtlasCloudProviderAccessFeatureUsageView, bool)`

GetFeatureUsagesOk returns a tuple with the FeatureUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureUsages

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetFeatureUsages(v []ApiAtlasCloudProviderAccessFeatureUsageView)`

SetFeatureUsages sets FeatureUsages field to given value.

### HasFeatureUsages

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasFeatureUsages() bool`

HasFeatureUsages returns a boolean if a field has been set.

### GetIamAssumedRoleArn

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetIamAssumedRoleArn() string`

GetIamAssumedRoleArn returns the IamAssumedRoleArn field if non-nil, zero value otherwise.

### GetIamAssumedRoleArnOk

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetIamAssumedRoleArnOk() (*string, bool)`

GetIamAssumedRoleArnOk returns a tuple with the IamAssumedRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamAssumedRoleArn

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetIamAssumedRoleArn(v string)`

SetIamAssumedRoleArn sets IamAssumedRoleArn field to given value.

### HasIamAssumedRoleArn

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasIamAssumedRoleArn() bool`

HasIamAssumedRoleArn returns a boolean if a field has been set.

### GetRoleId

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ApiAtlasCloudProviderAccessAWSIAMRoleViewAllOf) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


