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

// ApiAtlasGCPEndpointServiceView Group of Private Endpoint Service settings.
type ApiAtlasGCPEndpointServiceView struct {
	// List of Google Cloud network endpoint groups that corresponds to the Private Service Connect endpoint service.
	EndpointGroupNames []string `json:"endpointGroupNames,omitempty"`
	// Error message returned when requesting private connection resource. The resource returns `null` if the request succeeded.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Private Endpoint Service.
	Id *string `json:"id,omitempty"`
	// Cloud provider region that manages this Private Endpoint Service.
	RegionName *string `json:"regionName,omitempty"`
	// List of Uniform Resource Locators (URLs) that identifies endpoints that MongoDB Cloud can use to access one Google Cloud Service across a Google Cloud Virtual Private Connection (VPC) network.
	ServiceAttachmentNames []string `json:"serviceAttachmentNames,omitempty"`
	// State of the Private Endpoint Service connection when MongoDB Cloud received this request.
	Status *string `json:"status,omitempty"`
}

// NewApiAtlasGCPEndpointServiceView instantiates a new ApiAtlasGCPEndpointServiceView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasGCPEndpointServiceView() *ApiAtlasGCPEndpointServiceView {
	this := ApiAtlasGCPEndpointServiceView{}
	return &this
}

// NewApiAtlasGCPEndpointServiceViewWithDefaults instantiates a new ApiAtlasGCPEndpointServiceView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasGCPEndpointServiceViewWithDefaults() *ApiAtlasGCPEndpointServiceView {
	this := ApiAtlasGCPEndpointServiceView{}
	return &this
}

// GetEndpointGroupNames returns the EndpointGroupNames field value if set, zero value otherwise.
func (o *ApiAtlasGCPEndpointServiceView) GetEndpointGroupNames() []string {
	if o == nil || isNil(o.EndpointGroupNames) {
		var ret []string
		return ret
	}
	return o.EndpointGroupNames
}

// GetEndpointGroupNamesOk returns a tuple with the EndpointGroupNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGCPEndpointServiceView) GetEndpointGroupNamesOk() ([]string, bool) {
	if o == nil || isNil(o.EndpointGroupNames) {
    return nil, false
	}
	return o.EndpointGroupNames, true
}

// HasEndpointGroupNames returns a boolean if a field has been set.
func (o *ApiAtlasGCPEndpointServiceView) HasEndpointGroupNames() bool {
	if o != nil && !isNil(o.EndpointGroupNames) {
		return true
	}

	return false
}

// SetEndpointGroupNames gets a reference to the given []string and assigns it to the EndpointGroupNames field.
func (o *ApiAtlasGCPEndpointServiceView) SetEndpointGroupNames(v []string) {
	o.EndpointGroupNames = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ApiAtlasGCPEndpointServiceView) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGCPEndpointServiceView) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
    return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ApiAtlasGCPEndpointServiceView) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ApiAtlasGCPEndpointServiceView) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAtlasGCPEndpointServiceView) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGCPEndpointServiceView) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAtlasGCPEndpointServiceView) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAtlasGCPEndpointServiceView) SetId(v string) {
	o.Id = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *ApiAtlasGCPEndpointServiceView) GetRegionName() string {
	if o == nil || isNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGCPEndpointServiceView) GetRegionNameOk() (*string, bool) {
	if o == nil || isNil(o.RegionName) {
    return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ApiAtlasGCPEndpointServiceView) HasRegionName() bool {
	if o != nil && !isNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ApiAtlasGCPEndpointServiceView) SetRegionName(v string) {
	o.RegionName = &v
}

// GetServiceAttachmentNames returns the ServiceAttachmentNames field value if set, zero value otherwise.
func (o *ApiAtlasGCPEndpointServiceView) GetServiceAttachmentNames() []string {
	if o == nil || isNil(o.ServiceAttachmentNames) {
		var ret []string
		return ret
	}
	return o.ServiceAttachmentNames
}

// GetServiceAttachmentNamesOk returns a tuple with the ServiceAttachmentNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGCPEndpointServiceView) GetServiceAttachmentNamesOk() ([]string, bool) {
	if o == nil || isNil(o.ServiceAttachmentNames) {
    return nil, false
	}
	return o.ServiceAttachmentNames, true
}

// HasServiceAttachmentNames returns a boolean if a field has been set.
func (o *ApiAtlasGCPEndpointServiceView) HasServiceAttachmentNames() bool {
	if o != nil && !isNil(o.ServiceAttachmentNames) {
		return true
	}

	return false
}

// SetServiceAttachmentNames gets a reference to the given []string and assigns it to the ServiceAttachmentNames field.
func (o *ApiAtlasGCPEndpointServiceView) SetServiceAttachmentNames(v []string) {
	o.ServiceAttachmentNames = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiAtlasGCPEndpointServiceView) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGCPEndpointServiceView) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiAtlasGCPEndpointServiceView) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiAtlasGCPEndpointServiceView) SetStatus(v string) {
	o.Status = &v
}

func (o ApiAtlasGCPEndpointServiceView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EndpointGroupNames) {
		toSerialize["endpointGroupNames"] = o.EndpointGroupNames
	}
	if !isNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !isNil(o.ServiceAttachmentNames) {
		toSerialize["serviceAttachmentNames"] = o.ServiceAttachmentNames
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasGCPEndpointServiceView struct {
	value *ApiAtlasGCPEndpointServiceView
	isSet bool
}

func (v NullableApiAtlasGCPEndpointServiceView) Get() *ApiAtlasGCPEndpointServiceView {
	return v.value
}

func (v *NullableApiAtlasGCPEndpointServiceView) Set(val *ApiAtlasGCPEndpointServiceView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasGCPEndpointServiceView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasGCPEndpointServiceView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasGCPEndpointServiceView(val *ApiAtlasGCPEndpointServiceView) *NullableApiAtlasGCPEndpointServiceView {
	return &NullableApiAtlasGCPEndpointServiceView{value: val, isSet: true}
}

func (v NullableApiAtlasGCPEndpointServiceView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasGCPEndpointServiceView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


