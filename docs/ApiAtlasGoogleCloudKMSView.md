# ApiAtlasGoogleCloudKMSView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Flag that indicates whether someone enabled encryption at rest for the specified  project. To disable encryption at rest using customer key management and remove the configuration details, pass only this parameter with a value of &#x60;false&#x60;. | [optional] 
**KeyVersionResourceID** | Pointer to **string** | Resource path that displays the key version resource ID for your Google Cloud KMS. | [optional] 
**ServiceAccountKey** | Pointer to **string** | JavaScript Object Notation (JSON) object that contains the Google Cloud Key Management Service (KMS). Format the JSON as a string and not as an object. | [optional] 
**Valid** | Pointer to **bool** | Flag that indicates whether the Google Cloud Key Management Service (KMS) encryption key can encrypt and decrypt data. | [optional] [readonly] 

## Methods

### NewApiAtlasGoogleCloudKMSView

`func NewApiAtlasGoogleCloudKMSView() *ApiAtlasGoogleCloudKMSView`

NewApiAtlasGoogleCloudKMSView instantiates a new ApiAtlasGoogleCloudKMSView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasGoogleCloudKMSViewWithDefaults

`func NewApiAtlasGoogleCloudKMSViewWithDefaults() *ApiAtlasGoogleCloudKMSView`

NewApiAtlasGoogleCloudKMSViewWithDefaults instantiates a new ApiAtlasGoogleCloudKMSView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApiAtlasGoogleCloudKMSView) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiAtlasGoogleCloudKMSView) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiAtlasGoogleCloudKMSView) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiAtlasGoogleCloudKMSView) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetKeyVersionResourceID

`func (o *ApiAtlasGoogleCloudKMSView) GetKeyVersionResourceID() string`

GetKeyVersionResourceID returns the KeyVersionResourceID field if non-nil, zero value otherwise.

### GetKeyVersionResourceIDOk

`func (o *ApiAtlasGoogleCloudKMSView) GetKeyVersionResourceIDOk() (*string, bool)`

GetKeyVersionResourceIDOk returns a tuple with the KeyVersionResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersionResourceID

`func (o *ApiAtlasGoogleCloudKMSView) SetKeyVersionResourceID(v string)`

SetKeyVersionResourceID sets KeyVersionResourceID field to given value.

### HasKeyVersionResourceID

`func (o *ApiAtlasGoogleCloudKMSView) HasKeyVersionResourceID() bool`

HasKeyVersionResourceID returns a boolean if a field has been set.

### GetServiceAccountKey

`func (o *ApiAtlasGoogleCloudKMSView) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *ApiAtlasGoogleCloudKMSView) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *ApiAtlasGoogleCloudKMSView) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.

### HasServiceAccountKey

`func (o *ApiAtlasGoogleCloudKMSView) HasServiceAccountKey() bool`

HasServiceAccountKey returns a boolean if a field has been set.

### GetValid

`func (o *ApiAtlasGoogleCloudKMSView) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ApiAtlasGoogleCloudKMSView) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ApiAtlasGoogleCloudKMSView) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *ApiAtlasGoogleCloudKMSView) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


