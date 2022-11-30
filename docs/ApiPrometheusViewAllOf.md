# ApiPrometheusViewAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Flag that indicates whether someone has activated the Prometheus integration. | [optional] 
**ListenAddress** | Pointer to **string** | Combination of IPv4 address and Internet Assigned Numbers Authority (IANA) port or the IANA port alone to which Prometheus binds to ingest MongoDB metrics. | [optional] [default to ":9216"]
**Password** | Pointer to **string** |  | [optional] 
**Scheme** | Pointer to **string** | Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud. | [optional] 
**ServiceDiscovery** | Pointer to **string** | Desired method to discover the Prometheus service. | [optional] 
**TlsPemPath** | Pointer to **string** | Root-relative path to the Transport Layer Security (TLS) Privacy Enhanced Mail (PEM) key and certificate file on the host. | [optional] 
**Username** | Pointer to **string** | Human-readable label that identifies your Prometheus incoming webhook. | [optional] 

## Methods

### NewApiPrometheusViewAllOf

`func NewApiPrometheusViewAllOf() *ApiPrometheusViewAllOf`

NewApiPrometheusViewAllOf instantiates a new ApiPrometheusViewAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrometheusViewAllOfWithDefaults

`func NewApiPrometheusViewAllOfWithDefaults() *ApiPrometheusViewAllOf`

NewApiPrometheusViewAllOfWithDefaults instantiates a new ApiPrometheusViewAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApiPrometheusViewAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiPrometheusViewAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiPrometheusViewAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiPrometheusViewAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetListenAddress

`func (o *ApiPrometheusViewAllOf) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *ApiPrometheusViewAllOf) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *ApiPrometheusViewAllOf) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *ApiPrometheusViewAllOf) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetPassword

`func (o *ApiPrometheusViewAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiPrometheusViewAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiPrometheusViewAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiPrometheusViewAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetScheme

`func (o *ApiPrometheusViewAllOf) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ApiPrometheusViewAllOf) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ApiPrometheusViewAllOf) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ApiPrometheusViewAllOf) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetServiceDiscovery

`func (o *ApiPrometheusViewAllOf) GetServiceDiscovery() string`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *ApiPrometheusViewAllOf) GetServiceDiscoveryOk() (*string, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *ApiPrometheusViewAllOf) SetServiceDiscovery(v string)`

SetServiceDiscovery sets ServiceDiscovery field to given value.

### HasServiceDiscovery

`func (o *ApiPrometheusViewAllOf) HasServiceDiscovery() bool`

HasServiceDiscovery returns a boolean if a field has been set.

### GetTlsPemPath

`func (o *ApiPrometheusViewAllOf) GetTlsPemPath() string`

GetTlsPemPath returns the TlsPemPath field if non-nil, zero value otherwise.

### GetTlsPemPathOk

`func (o *ApiPrometheusViewAllOf) GetTlsPemPathOk() (*string, bool)`

GetTlsPemPathOk returns a tuple with the TlsPemPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsPemPath

`func (o *ApiPrometheusViewAllOf) SetTlsPemPath(v string)`

SetTlsPemPath sets TlsPemPath field to given value.

### HasTlsPemPath

`func (o *ApiPrometheusViewAllOf) HasTlsPemPath() bool`

HasTlsPemPath returns a boolean if a field has been set.

### GetUsername

`func (o *ApiPrometheusViewAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiPrometheusViewAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiPrometheusViewAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiPrometheusViewAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


