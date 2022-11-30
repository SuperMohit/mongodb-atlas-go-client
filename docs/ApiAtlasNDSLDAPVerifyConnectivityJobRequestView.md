# ApiAtlasNDSLDAPVerifyConnectivityJobRequestView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project associated with this Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Request** | Pointer to [**ApiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView**](ApiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView.md) |  | [optional] 
**RequestId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this request to verify an Lightweight Directory Access Protocol (LDAP) configuration. | [optional] [readonly] 
**Status** | Pointer to **string** | Human-readable string that indicates the status of the Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration. | [optional] [readonly] 
**Validations** | Pointer to [**[]ApiAtlasNDSLDAPVerifyConnectivityJobRequestValidationView**](ApiAtlasNDSLDAPVerifyConnectivityJobRequestValidationView.md) | List that contains the validation messages related to the verification of the provided Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration details. The list contains a document for each test that MongoDB Cloud runs. MongoDB Cloud stops running tests after the first failure. | [optional] [readonly] 

## Methods

### NewApiAtlasNDSLDAPVerifyConnectivityJobRequestView

`func NewApiAtlasNDSLDAPVerifyConnectivityJobRequestView(links []Link, ) *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView`

NewApiAtlasNDSLDAPVerifyConnectivityJobRequestView instantiates a new ApiAtlasNDSLDAPVerifyConnectivityJobRequestView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasNDSLDAPVerifyConnectivityJobRequestViewWithDefaults

`func NewApiAtlasNDSLDAPVerifyConnectivityJobRequestViewWithDefaults() *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView`

NewApiAtlasNDSLDAPVerifyConnectivityJobRequestViewWithDefaults instantiates a new ApiAtlasNDSLDAPVerifyConnectivityJobRequestView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetRequest

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetRequest() ApiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetRequestOk() (*ApiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) SetRequest(v ApiAtlasNDSLDAPVerifyConnectivityJobRequestParamsView)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetRequestId

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValidations

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetValidations() []ApiAtlasNDSLDAPVerifyConnectivityJobRequestValidationView`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) GetValidationsOk() (*[]ApiAtlasNDSLDAPVerifyConnectivityJobRequestValidationView, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) SetValidations(v []ApiAtlasNDSLDAPVerifyConnectivityJobRequestValidationView)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *ApiAtlasNDSLDAPVerifyConnectivityJobRequestView) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


