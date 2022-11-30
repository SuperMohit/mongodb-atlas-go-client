# ApiPrometheusView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Flag that indicates whether someone has activated the Prometheus integration. | 
**ListenAddress** | Pointer to **string** | Combination of IPv4 address and Internet Assigned Numbers Authority (IANA) port or the IANA port alone to which Prometheus binds to ingest MongoDB metrics. | [optional] [default to ":9216"]
**Password** | Pointer to **string** |  | [optional] 
**Scheme** | **string** | Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud. | 
**ServiceDiscovery** | **string** | Desired method to discover the Prometheus service. | 
**TlsPemPath** | Pointer to **string** | Root-relative path to the Transport Layer Security (TLS) Privacy Enhanced Mail (PEM) key and certificate file on the host. | [optional] 
**Username** | **string** | Human-readable label that identifies your Prometheus incoming webhook. | 

## Methods

### NewApiPrometheusView

`func NewApiPrometheusView(enabled bool, scheme string, serviceDiscovery string, username string, ) *ApiPrometheusView`

NewApiPrometheusView instantiates a new ApiPrometheusView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrometheusViewWithDefaults

`func NewApiPrometheusViewWithDefaults() *ApiPrometheusView`

NewApiPrometheusViewWithDefaults instantiates a new ApiPrometheusView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApiPrometheusView) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiPrometheusView) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiPrometheusView) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetListenAddress

`func (o *ApiPrometheusView) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *ApiPrometheusView) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *ApiPrometheusView) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.

### HasListenAddress

`func (o *ApiPrometheusView) HasListenAddress() bool`

HasListenAddress returns a boolean if a field has been set.

### GetPassword

`func (o *ApiPrometheusView) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiPrometheusView) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiPrometheusView) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiPrometheusView) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetScheme

`func (o *ApiPrometheusView) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ApiPrometheusView) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ApiPrometheusView) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetServiceDiscovery

`func (o *ApiPrometheusView) GetServiceDiscovery() string`

GetServiceDiscovery returns the ServiceDiscovery field if non-nil, zero value otherwise.

### GetServiceDiscoveryOk

`func (o *ApiPrometheusView) GetServiceDiscoveryOk() (*string, bool)`

GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDiscovery

`func (o *ApiPrometheusView) SetServiceDiscovery(v string)`

SetServiceDiscovery sets ServiceDiscovery field to given value.


### GetTlsPemPath

`func (o *ApiPrometheusView) GetTlsPemPath() string`

GetTlsPemPath returns the TlsPemPath field if non-nil, zero value otherwise.

### GetTlsPemPathOk

`func (o *ApiPrometheusView) GetTlsPemPathOk() (*string, bool)`

GetTlsPemPathOk returns a tuple with the TlsPemPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsPemPath

`func (o *ApiPrometheusView) SetTlsPemPath(v string)`

SetTlsPemPath sets TlsPemPath field to given value.

### HasTlsPemPath

`func (o *ApiPrometheusView) HasTlsPemPath() bool`

HasTlsPemPath returns a boolean if a field has been set.

### GetUsername

`func (o *ApiPrometheusView) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiPrometheusView) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiPrometheusView) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


