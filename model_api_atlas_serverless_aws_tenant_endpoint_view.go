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

// ApiAtlasServerlessAWSTenantEndpointView View for a serverless AWS tenant endpoint.
type ApiAtlasServerlessAWSTenantEndpointView struct {
	// Unique 24-hexadecimal digit string that identifies the private endpoint.
	Id *string `json:"_id,omitempty"`
	// Unique string that identifies the private endpoint's network interface.
	CloudProviderEndpointId *string `json:"cloudProviderEndpointId,omitempty"`
	// Human-readable comment associated with the private endpoint.
	Comment *string `json:"comment,omitempty"`
	// Unique string that identifies the Amazon Web Services (AWS) PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service.
	EndpointServiceName *string `json:"endpointServiceName,omitempty"`
	// Human-readable error message that indicates error condition associated with establishing the private endpoint connection.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Human-readable label that indicates the current operating status of the private endpoint.
	Status *string `json:"status,omitempty"`
}

// NewApiAtlasServerlessAWSTenantEndpointView instantiates a new ApiAtlasServerlessAWSTenantEndpointView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasServerlessAWSTenantEndpointView() *ApiAtlasServerlessAWSTenantEndpointView {
	this := ApiAtlasServerlessAWSTenantEndpointView{}
	return &this
}

// NewApiAtlasServerlessAWSTenantEndpointViewWithDefaults instantiates a new ApiAtlasServerlessAWSTenantEndpointView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasServerlessAWSTenantEndpointViewWithDefaults() *ApiAtlasServerlessAWSTenantEndpointView {
	this := ApiAtlasServerlessAWSTenantEndpointView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAtlasServerlessAWSTenantEndpointView) SetId(v string) {
	o.Id = &v
}

// GetCloudProviderEndpointId returns the CloudProviderEndpointId field value if set, zero value otherwise.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetCloudProviderEndpointId() string {
	if o == nil || isNil(o.CloudProviderEndpointId) {
		var ret string
		return ret
	}
	return *o.CloudProviderEndpointId
}

// GetCloudProviderEndpointIdOk returns a tuple with the CloudProviderEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetCloudProviderEndpointIdOk() (*string, bool) {
	if o == nil || isNil(o.CloudProviderEndpointId) {
    return nil, false
	}
	return o.CloudProviderEndpointId, true
}

// HasCloudProviderEndpointId returns a boolean if a field has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) HasCloudProviderEndpointId() bool {
	if o != nil && !isNil(o.CloudProviderEndpointId) {
		return true
	}

	return false
}

// SetCloudProviderEndpointId gets a reference to the given string and assigns it to the CloudProviderEndpointId field.
func (o *ApiAtlasServerlessAWSTenantEndpointView) SetCloudProviderEndpointId(v string) {
	o.CloudProviderEndpointId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApiAtlasServerlessAWSTenantEndpointView) SetComment(v string) {
	o.Comment = &v
}

// GetEndpointServiceName returns the EndpointServiceName field value if set, zero value otherwise.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetEndpointServiceName() string {
	if o == nil || isNil(o.EndpointServiceName) {
		var ret string
		return ret
	}
	return *o.EndpointServiceName
}

// GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetEndpointServiceNameOk() (*string, bool) {
	if o == nil || isNil(o.EndpointServiceName) {
    return nil, false
	}
	return o.EndpointServiceName, true
}

// HasEndpointServiceName returns a boolean if a field has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) HasEndpointServiceName() bool {
	if o != nil && !isNil(o.EndpointServiceName) {
		return true
	}

	return false
}

// SetEndpointServiceName gets a reference to the given string and assigns it to the EndpointServiceName field.
func (o *ApiAtlasServerlessAWSTenantEndpointView) SetEndpointServiceName(v string) {
	o.EndpointServiceName = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
    return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ApiAtlasServerlessAWSTenantEndpointView) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiAtlasServerlessAWSTenantEndpointView) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiAtlasServerlessAWSTenantEndpointView) SetStatus(v string) {
	o.Status = &v
}

func (o ApiAtlasServerlessAWSTenantEndpointView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !isNil(o.CloudProviderEndpointId) {
		toSerialize["cloudProviderEndpointId"] = o.CloudProviderEndpointId
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !isNil(o.EndpointServiceName) {
		toSerialize["endpointServiceName"] = o.EndpointServiceName
	}
	if !isNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasServerlessAWSTenantEndpointView struct {
	value *ApiAtlasServerlessAWSTenantEndpointView
	isSet bool
}

func (v NullableApiAtlasServerlessAWSTenantEndpointView) Get() *ApiAtlasServerlessAWSTenantEndpointView {
	return v.value
}

func (v *NullableApiAtlasServerlessAWSTenantEndpointView) Set(val *ApiAtlasServerlessAWSTenantEndpointView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasServerlessAWSTenantEndpointView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasServerlessAWSTenantEndpointView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasServerlessAWSTenantEndpointView(val *ApiAtlasServerlessAWSTenantEndpointView) *NullableApiAtlasServerlessAWSTenantEndpointView {
	return &NullableApiAtlasServerlessAWSTenantEndpointView{value: val, isSet: true}
}

func (v NullableApiAtlasServerlessAWSTenantEndpointView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasServerlessAWSTenantEndpointView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

