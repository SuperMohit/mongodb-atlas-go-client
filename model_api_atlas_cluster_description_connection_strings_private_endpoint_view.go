/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView List of private endpoint connection strings that you can use to connect to this cluster through a private endpoint. This parameter returns only if you deployed a private endpoint to all regions to which you deployed this clusters' nodes.
type ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView struct {
	// Private endpoint-aware connection string that uses the `mongodb://` protocol to connect to MongoDB Cloud through a private endpoint.
	ConnectionString *string `json:"connectionString,omitempty"`
	// List that contains the private endpoints through which you connect to MongoDB Cloud when you use **connectionStrings.privateEndpoint[n].connectionString** or **connectionStrings.privateEndpoint[n].srvConnectionString**.
	Endpoints []ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointEndpointView `json:"endpoints,omitempty"`
	// Private endpoint-aware connection string that uses the `mongodb+srv://` protocol to connect to MongoDB Cloud through a private endpoint. The `mongodb+srv` protocol tells the driver to look up the seed list of hosts in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don't need to append the seed list or change the Uniform Resource Identifier (URI) if the nodes change. Use this Uniform Resource Identifier (URI) format if your application supports it. If it doesn't, use connectionStrings.privateEndpoint[n].connectionString.
	SrvConnectionString *string `json:"srvConnectionString,omitempty"`
	// MongoDB process type to which your application connects. Use `MONGOD` for replica sets and `MONGOS` for sharded clusters.
	Type *string `json:"type,omitempty"`
}

// NewApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView instantiates a new ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView() *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView {
	this := ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView{}
	return &this
}

// NewApiAtlasClusterDescriptionConnectionStringsPrivateEndpointViewWithDefaults instantiates a new ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasClusterDescriptionConnectionStringsPrivateEndpointViewWithDefaults() *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView {
	this := ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView{}
	return &this
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) GetConnectionString() string {
	if o == nil || isNil(o.ConnectionString) {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) GetConnectionStringOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionString) {
    return nil, false
	}
	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) HasConnectionString() bool {
	if o != nil && !isNil(o.ConnectionString) {
		return true
	}

	return false
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) SetConnectionString(v string) {
	o.ConnectionString = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) GetEndpoints() []ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointEndpointView {
	if o == nil || isNil(o.Endpoints) {
		var ret []ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointEndpointView
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) GetEndpointsOk() ([]ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointEndpointView, bool) {
	if o == nil || isNil(o.Endpoints) {
    return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) HasEndpoints() bool {
	if o != nil && !isNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointEndpointView and assigns it to the Endpoints field.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) SetEndpoints(v []ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointEndpointView) {
	o.Endpoints = v
}

// GetSrvConnectionString returns the SrvConnectionString field value if set, zero value otherwise.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) GetSrvConnectionString() string {
	if o == nil || isNil(o.SrvConnectionString) {
		var ret string
		return ret
	}
	return *o.SrvConnectionString
}

// GetSrvConnectionStringOk returns a tuple with the SrvConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) GetSrvConnectionStringOk() (*string, bool) {
	if o == nil || isNil(o.SrvConnectionString) {
    return nil, false
	}
	return o.SrvConnectionString, true
}

// HasSrvConnectionString returns a boolean if a field has been set.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) HasSrvConnectionString() bool {
	if o != nil && !isNil(o.SrvConnectionString) {
		return true
	}

	return false
}

// SetSrvConnectionString gets a reference to the given string and assigns it to the SrvConnectionString field.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) SetSrvConnectionString(v string) {
	o.SrvConnectionString = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) SetType(v string) {
	o.Type = &v
}

func (o ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ConnectionString) {
		toSerialize["connectionString"] = o.ConnectionString
	}
	if !isNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	if !isNil(o.SrvConnectionString) {
		toSerialize["srvConnectionString"] = o.SrvConnectionString
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView struct {
	value *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView
	isSet bool
}

func (v NullableApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) Get() *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView {
	return v.value
}

func (v *NullableApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) Set(val *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView(val *ApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) *NullableApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView {
	return &NullableApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView{value: val, isSet: true}
}

func (v NullableApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasClusterDescriptionConnectionStringsPrivateEndpointView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

