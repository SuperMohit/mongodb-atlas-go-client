# ApiDatabaseUserX509ViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Human-readable label that represents the user that authenticates to MongoDB. This must be formatted as a [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name.  | [optional] 
**X509Type** | Pointer to **string** | X.509 method that MongoDB Cloud uses to authenticate the database user. - For MongoDB Cloud-managed X.509, specify &#x60;MANAGED&#x60;. - For self-managed X.509, specify &#x60;CUSTOMER&#x60;. Users created with the &#x60;CUSTOMER&#x60; method require a Common Name (CN) in the **username** parameter. You must create externally authenticated users on the &#x60;$external&#x60; database. | [optional] [default to "NONE"]

## Methods

### NewApiDatabaseUserX509ViewAllOf

`func NewApiDatabaseUserX509ViewAllOf() *ApiDatabaseUserX509ViewAllOf`

NewApiDatabaseUserX509ViewAllOf instantiates a new ApiDatabaseUserX509ViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDatabaseUserX509ViewAllOfWithDefaults

`func NewApiDatabaseUserX509ViewAllOfWithDefaults() *ApiDatabaseUserX509ViewAllOf`

NewApiDatabaseUserX509ViewAllOfWithDefaults instantiates a new ApiDatabaseUserX509ViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ApiDatabaseUserX509ViewAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiDatabaseUserX509ViewAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiDatabaseUserX509ViewAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiDatabaseUserX509ViewAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetX509Type

`func (o *ApiDatabaseUserX509ViewAllOf) GetX509Type() string`

GetX509Type returns the X509Type field if non-nil, zero value otherwise.

### GetX509TypeOk

`func (o *ApiDatabaseUserX509ViewAllOf) GetX509TypeOk() (*string, bool)`

GetX509TypeOk returns a tuple with the X509Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Type

`func (o *ApiDatabaseUserX509ViewAllOf) SetX509Type(v string)`

SetX509Type sets X509Type field to given value.

### HasX509Type

`func (o *ApiDatabaseUserX509ViewAllOf) HasX509Type() bool`

HasX509Type returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


