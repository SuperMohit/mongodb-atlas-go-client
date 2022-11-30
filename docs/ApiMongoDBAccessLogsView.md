# ApiMongoDBAccessLogsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthResult** | Pointer to **bool** | Flag that indicates whether the response should return successful authentication attempts only. | [optional] 
**AuthSource** | Pointer to **string** | Database against which someone attempted to authenticate. | [optional] [readonly] 
**FailureReason** | Pointer to **string** | Reason that the authentication failed. Null if authentication succeeded.  | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Hostname** | Pointer to **string** | Human-readable label that identifies the hostname of the target node that received the authentication attempt. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | Internet Protocol address that attempted to authenticate with the database. | [optional] [readonly] 
**LogLine** | Pointer to **string** | Text of the host log concerning the authentication attempt. | [optional] [readonly] 
**Timestamp** | Pointer to **string** | Date and time when someone made this authentication attempt. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**Username** | Pointer to **string** | Username used to authenticate against the database. | [optional] [readonly] 

## Methods

### NewApiMongoDBAccessLogsView

`func NewApiMongoDBAccessLogsView() *ApiMongoDBAccessLogsView`

NewApiMongoDBAccessLogsView instantiates a new ApiMongoDBAccessLogsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMongoDBAccessLogsViewWithDefaults

`func NewApiMongoDBAccessLogsViewWithDefaults() *ApiMongoDBAccessLogsView`

NewApiMongoDBAccessLogsViewWithDefaults instantiates a new ApiMongoDBAccessLogsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthResult

`func (o *ApiMongoDBAccessLogsView) GetAuthResult() bool`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *ApiMongoDBAccessLogsView) GetAuthResultOk() (*bool, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *ApiMongoDBAccessLogsView) SetAuthResult(v bool)`

SetAuthResult sets AuthResult field to given value.

### HasAuthResult

`func (o *ApiMongoDBAccessLogsView) HasAuthResult() bool`

HasAuthResult returns a boolean if a field has been set.

### GetAuthSource

`func (o *ApiMongoDBAccessLogsView) GetAuthSource() string`

GetAuthSource returns the AuthSource field if non-nil, zero value otherwise.

### GetAuthSourceOk

`func (o *ApiMongoDBAccessLogsView) GetAuthSourceOk() (*string, bool)`

GetAuthSourceOk returns a tuple with the AuthSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSource

`func (o *ApiMongoDBAccessLogsView) SetAuthSource(v string)`

SetAuthSource sets AuthSource field to given value.

### HasAuthSource

`func (o *ApiMongoDBAccessLogsView) HasAuthSource() bool`

HasAuthSource returns a boolean if a field has been set.

### GetFailureReason

`func (o *ApiMongoDBAccessLogsView) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ApiMongoDBAccessLogsView) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ApiMongoDBAccessLogsView) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ApiMongoDBAccessLogsView) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiMongoDBAccessLogsView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiMongoDBAccessLogsView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiMongoDBAccessLogsView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiMongoDBAccessLogsView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHostname

`func (o *ApiMongoDBAccessLogsView) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApiMongoDBAccessLogsView) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApiMongoDBAccessLogsView) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApiMongoDBAccessLogsView) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpAddress

`func (o *ApiMongoDBAccessLogsView) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ApiMongoDBAccessLogsView) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ApiMongoDBAccessLogsView) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ApiMongoDBAccessLogsView) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLogLine

`func (o *ApiMongoDBAccessLogsView) GetLogLine() string`

GetLogLine returns the LogLine field if non-nil, zero value otherwise.

### GetLogLineOk

`func (o *ApiMongoDBAccessLogsView) GetLogLineOk() (*string, bool)`

GetLogLineOk returns a tuple with the LogLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLine

`func (o *ApiMongoDBAccessLogsView) SetLogLine(v string)`

SetLogLine sets LogLine field to given value.

### HasLogLine

`func (o *ApiMongoDBAccessLogsView) HasLogLine() bool`

HasLogLine returns a boolean if a field has been set.

### GetTimestamp

`func (o *ApiMongoDBAccessLogsView) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiMongoDBAccessLogsView) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiMongoDBAccessLogsView) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiMongoDBAccessLogsView) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUsername

`func (o *ApiMongoDBAccessLogsView) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiMongoDBAccessLogsView) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiMongoDBAccessLogsView) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiMongoDBAccessLogsView) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


