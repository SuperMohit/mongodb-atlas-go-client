# ApiNewRelicView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique 40-hexadecimal digit string that identifies your New Relic account. | 
**LicenseKey** | **string** | Unique 40-hexadecimal digit string that identifies your New Relic license. | 
**ReadToken** | **string** | Query key used to access your New Relic account. | 
**WriteToken** | **string** | Insert key associated with your New Relic account. | 

## Methods

### NewApiNewRelicView

`func NewApiNewRelicView(accountId string, licenseKey string, readToken string, writeToken string, ) *ApiNewRelicView`

NewApiNewRelicView instantiates a new ApiNewRelicView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNewRelicViewWithDefaults

`func NewApiNewRelicViewWithDefaults() *ApiNewRelicView`

NewApiNewRelicViewWithDefaults instantiates a new ApiNewRelicView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ApiNewRelicView) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ApiNewRelicView) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ApiNewRelicView) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetLicenseKey

`func (o *ApiNewRelicView) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *ApiNewRelicView) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *ApiNewRelicView) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.


### GetReadToken

`func (o *ApiNewRelicView) GetReadToken() string`

GetReadToken returns the ReadToken field if non-nil, zero value otherwise.

### GetReadTokenOk

`func (o *ApiNewRelicView) GetReadTokenOk() (*string, bool)`

GetReadTokenOk returns a tuple with the ReadToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadToken

`func (o *ApiNewRelicView) SetReadToken(v string)`

SetReadToken sets ReadToken field to given value.


### GetWriteToken

`func (o *ApiNewRelicView) GetWriteToken() string`

GetWriteToken returns the WriteToken field if non-nil, zero value otherwise.

### GetWriteTokenOk

`func (o *ApiNewRelicView) GetWriteTokenOk() (*string, bool)`

GetWriteTokenOk returns a tuple with the WriteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteToken

`func (o *ApiNewRelicView) SetWriteToken(v string)`

SetWriteToken sets WriteToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


