# ApiDatabaseUserIAMViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsIAMType** | Pointer to **string** | Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user&#39;s role. | [optional] [default to "NONE"]
**Username** | Pointer to **string** | Human-readable label that represents the user that authenticates to MongoDB.  | [optional] 

## Methods

### NewApiDatabaseUserIAMViewAllOf

`func NewApiDatabaseUserIAMViewAllOf() *ApiDatabaseUserIAMViewAllOf`

NewApiDatabaseUserIAMViewAllOf instantiates a new ApiDatabaseUserIAMViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDatabaseUserIAMViewAllOfWithDefaults

`func NewApiDatabaseUserIAMViewAllOfWithDefaults() *ApiDatabaseUserIAMViewAllOf`

NewApiDatabaseUserIAMViewAllOfWithDefaults instantiates a new ApiDatabaseUserIAMViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsIAMType

`func (o *ApiDatabaseUserIAMViewAllOf) GetAwsIAMType() string`

GetAwsIAMType returns the AwsIAMType field if non-nil, zero value otherwise.

### GetAwsIAMTypeOk

`func (o *ApiDatabaseUserIAMViewAllOf) GetAwsIAMTypeOk() (*string, bool)`

GetAwsIAMTypeOk returns a tuple with the AwsIAMType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIAMType

`func (o *ApiDatabaseUserIAMViewAllOf) SetAwsIAMType(v string)`

SetAwsIAMType sets AwsIAMType field to given value.

### HasAwsIAMType

`func (o *ApiDatabaseUserIAMViewAllOf) HasAwsIAMType() bool`

HasAwsIAMType returns a boolean if a field has been set.

### GetUsername

`func (o *ApiDatabaseUserIAMViewAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiDatabaseUserIAMViewAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiDatabaseUserIAMViewAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiDatabaseUserIAMViewAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


