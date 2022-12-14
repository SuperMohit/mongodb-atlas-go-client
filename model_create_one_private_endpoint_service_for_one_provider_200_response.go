/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CreateOnePrivateEndpointServiceForOneProvider200Response - struct for CreateOnePrivateEndpointServiceForOneProvider200Response
type CreateOnePrivateEndpointServiceForOneProvider200Response struct {
	ApiAtlasAWSPrivateLinkConnectionView *ApiAtlasAWSPrivateLinkConnectionView
	ApiAtlasAzurePrivateLinkConnectionView *ApiAtlasAzurePrivateLinkConnectionView
	ApiAtlasGCPEndpointServiceView *ApiAtlasGCPEndpointServiceView
}

// ApiAtlasAWSPrivateLinkConnectionViewAsCreateOnePrivateEndpointServiceForOneProvider200Response is a convenience function that returns ApiAtlasAWSPrivateLinkConnectionView wrapped in CreateOnePrivateEndpointServiceForOneProvider200Response
func ApiAtlasAWSPrivateLinkConnectionViewAsCreateOnePrivateEndpointServiceForOneProvider200Response(v *ApiAtlasAWSPrivateLinkConnectionView) CreateOnePrivateEndpointServiceForOneProvider200Response {
	return CreateOnePrivateEndpointServiceForOneProvider200Response{
		ApiAtlasAWSPrivateLinkConnectionView: v,
	}
}

// ApiAtlasAzurePrivateLinkConnectionViewAsCreateOnePrivateEndpointServiceForOneProvider200Response is a convenience function that returns ApiAtlasAzurePrivateLinkConnectionView wrapped in CreateOnePrivateEndpointServiceForOneProvider200Response
func ApiAtlasAzurePrivateLinkConnectionViewAsCreateOnePrivateEndpointServiceForOneProvider200Response(v *ApiAtlasAzurePrivateLinkConnectionView) CreateOnePrivateEndpointServiceForOneProvider200Response {
	return CreateOnePrivateEndpointServiceForOneProvider200Response{
		ApiAtlasAzurePrivateLinkConnectionView: v,
	}
}

// ApiAtlasGCPEndpointServiceViewAsCreateOnePrivateEndpointServiceForOneProvider200Response is a convenience function that returns ApiAtlasGCPEndpointServiceView wrapped in CreateOnePrivateEndpointServiceForOneProvider200Response
func ApiAtlasGCPEndpointServiceViewAsCreateOnePrivateEndpointServiceForOneProvider200Response(v *ApiAtlasGCPEndpointServiceView) CreateOnePrivateEndpointServiceForOneProvider200Response {
	return CreateOnePrivateEndpointServiceForOneProvider200Response{
		ApiAtlasGCPEndpointServiceView: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateOnePrivateEndpointServiceForOneProvider200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApiAtlasAWSPrivateLinkConnectionView
	err = newStrictDecoder(data).Decode(&dst.ApiAtlasAWSPrivateLinkConnectionView)
	if err == nil {
		jsonApiAtlasAWSPrivateLinkConnectionView, _ := json.Marshal(dst.ApiAtlasAWSPrivateLinkConnectionView)
		if string(jsonApiAtlasAWSPrivateLinkConnectionView) == "{}" { // empty struct
			dst.ApiAtlasAWSPrivateLinkConnectionView = nil
		} else {
			match++
		}
	} else {
		dst.ApiAtlasAWSPrivateLinkConnectionView = nil
	}

	// try to unmarshal data into ApiAtlasAzurePrivateLinkConnectionView
	err = newStrictDecoder(data).Decode(&dst.ApiAtlasAzurePrivateLinkConnectionView)
	if err == nil {
		jsonApiAtlasAzurePrivateLinkConnectionView, _ := json.Marshal(dst.ApiAtlasAzurePrivateLinkConnectionView)
		if string(jsonApiAtlasAzurePrivateLinkConnectionView) == "{}" { // empty struct
			dst.ApiAtlasAzurePrivateLinkConnectionView = nil
		} else {
			match++
		}
	} else {
		dst.ApiAtlasAzurePrivateLinkConnectionView = nil
	}

	// try to unmarshal data into ApiAtlasGCPEndpointServiceView
	err = newStrictDecoder(data).Decode(&dst.ApiAtlasGCPEndpointServiceView)
	if err == nil {
		jsonApiAtlasGCPEndpointServiceView, _ := json.Marshal(dst.ApiAtlasGCPEndpointServiceView)
		if string(jsonApiAtlasGCPEndpointServiceView) == "{}" { // empty struct
			dst.ApiAtlasGCPEndpointServiceView = nil
		} else {
			match++
		}
	} else {
		dst.ApiAtlasGCPEndpointServiceView = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApiAtlasAWSPrivateLinkConnectionView = nil
		dst.ApiAtlasAzurePrivateLinkConnectionView = nil
		dst.ApiAtlasGCPEndpointServiceView = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateOnePrivateEndpointServiceForOneProvider200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateOnePrivateEndpointServiceForOneProvider200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateOnePrivateEndpointServiceForOneProvider200Response) MarshalJSON() ([]byte, error) {
	if src.ApiAtlasAWSPrivateLinkConnectionView != nil {
		return json.Marshal(&src.ApiAtlasAWSPrivateLinkConnectionView)
	}

	if src.ApiAtlasAzurePrivateLinkConnectionView != nil {
		return json.Marshal(&src.ApiAtlasAzurePrivateLinkConnectionView)
	}

	if src.ApiAtlasGCPEndpointServiceView != nil {
		return json.Marshal(&src.ApiAtlasGCPEndpointServiceView)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateOnePrivateEndpointServiceForOneProvider200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApiAtlasAWSPrivateLinkConnectionView != nil {
		return obj.ApiAtlasAWSPrivateLinkConnectionView
	}

	if obj.ApiAtlasAzurePrivateLinkConnectionView != nil {
		return obj.ApiAtlasAzurePrivateLinkConnectionView
	}

	if obj.ApiAtlasGCPEndpointServiceView != nil {
		return obj.ApiAtlasGCPEndpointServiceView
	}

	// all schemas are nil
	return nil
}

type NullableCreateOnePrivateEndpointServiceForOneProvider200Response struct {
	value *CreateOnePrivateEndpointServiceForOneProvider200Response
	isSet bool
}

func (v NullableCreateOnePrivateEndpointServiceForOneProvider200Response) Get() *CreateOnePrivateEndpointServiceForOneProvider200Response {
	return v.value
}

func (v *NullableCreateOnePrivateEndpointServiceForOneProvider200Response) Set(val *CreateOnePrivateEndpointServiceForOneProvider200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOnePrivateEndpointServiceForOneProvider200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOnePrivateEndpointServiceForOneProvider200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOnePrivateEndpointServiceForOneProvider200Response(val *CreateOnePrivateEndpointServiceForOneProvider200Response) *NullableCreateOnePrivateEndpointServiceForOneProvider200Response {
	return &NullableCreateOnePrivateEndpointServiceForOneProvider200Response{value: val, isSet: true}
}

func (v NullableCreateOnePrivateEndpointServiceForOneProvider200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOnePrivateEndpointServiceForOneProvider200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


