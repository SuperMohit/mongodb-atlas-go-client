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

// ApiAtlasClusterRegionConfigsV15Manual - struct for ApiAtlasClusterRegionConfigsV15Manual
type ApiAtlasClusterRegionConfigsV15Manual struct {
	ApiAtlasClusterRegionConfigsV15AWSManual *ApiAtlasClusterRegionConfigsV15AWSManual
	ApiAtlasClusterRegionConfigsV15AzureManual *ApiAtlasClusterRegionConfigsV15AzureManual
	ApiAtlasClusterRegionConfigsV15GCPManual *ApiAtlasClusterRegionConfigsV15GCPManual
	ApiAtlasClusterRegionConfigsV15TenantManual *ApiAtlasClusterRegionConfigsV15TenantManual
}

// ApiAtlasClusterRegionConfigsV15AWSManualAsApiAtlasClusterRegionConfigsV15Manual is a convenience function that returns ApiAtlasClusterRegionConfigsV15AWSManual wrapped in ApiAtlasClusterRegionConfigsV15Manual
func ApiAtlasClusterRegionConfigsV15AWSManualAsApiAtlasClusterRegionConfigsV15Manual(v *ApiAtlasClusterRegionConfigsV15AWSManual) ApiAtlasClusterRegionConfigsV15Manual {
	return ApiAtlasClusterRegionConfigsV15Manual{
		ApiAtlasClusterRegionConfigsV15AWSManual: v,
	}
}

// ApiAtlasClusterRegionConfigsV15AzureManualAsApiAtlasClusterRegionConfigsV15Manual is a convenience function that returns ApiAtlasClusterRegionConfigsV15AzureManual wrapped in ApiAtlasClusterRegionConfigsV15Manual
func ApiAtlasClusterRegionConfigsV15AzureManualAsApiAtlasClusterRegionConfigsV15Manual(v *ApiAtlasClusterRegionConfigsV15AzureManual) ApiAtlasClusterRegionConfigsV15Manual {
	return ApiAtlasClusterRegionConfigsV15Manual{
		ApiAtlasClusterRegionConfigsV15AzureManual: v,
	}
}

// ApiAtlasClusterRegionConfigsV15GCPManualAsApiAtlasClusterRegionConfigsV15Manual is a convenience function that returns ApiAtlasClusterRegionConfigsV15GCPManual wrapped in ApiAtlasClusterRegionConfigsV15Manual
func ApiAtlasClusterRegionConfigsV15GCPManualAsApiAtlasClusterRegionConfigsV15Manual(v *ApiAtlasClusterRegionConfigsV15GCPManual) ApiAtlasClusterRegionConfigsV15Manual {
	return ApiAtlasClusterRegionConfigsV15Manual{
		ApiAtlasClusterRegionConfigsV15GCPManual: v,
	}
}

// ApiAtlasClusterRegionConfigsV15TenantManualAsApiAtlasClusterRegionConfigsV15Manual is a convenience function that returns ApiAtlasClusterRegionConfigsV15TenantManual wrapped in ApiAtlasClusterRegionConfigsV15Manual
func ApiAtlasClusterRegionConfigsV15TenantManualAsApiAtlasClusterRegionConfigsV15Manual(v *ApiAtlasClusterRegionConfigsV15TenantManual) ApiAtlasClusterRegionConfigsV15Manual {
	return ApiAtlasClusterRegionConfigsV15Manual{
		ApiAtlasClusterRegionConfigsV15TenantManual: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApiAtlasClusterRegionConfigsV15Manual) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApiAtlasClusterRegionConfigsV15AWSManual
	err = newStrictDecoder(data).Decode(&dst.ApiAtlasClusterRegionConfigsV15AWSManual)
	if err == nil {
		jsonApiAtlasClusterRegionConfigsV15AWSManual, _ := json.Marshal(dst.ApiAtlasClusterRegionConfigsV15AWSManual)
		if string(jsonApiAtlasClusterRegionConfigsV15AWSManual) == "{}" { // empty struct
			dst.ApiAtlasClusterRegionConfigsV15AWSManual = nil
		} else {
			match++
		}
	} else {
		dst.ApiAtlasClusterRegionConfigsV15AWSManual = nil
	}

	// try to unmarshal data into ApiAtlasClusterRegionConfigsV15AzureManual
	err = newStrictDecoder(data).Decode(&dst.ApiAtlasClusterRegionConfigsV15AzureManual)
	if err == nil {
		jsonApiAtlasClusterRegionConfigsV15AzureManual, _ := json.Marshal(dst.ApiAtlasClusterRegionConfigsV15AzureManual)
		if string(jsonApiAtlasClusterRegionConfigsV15AzureManual) == "{}" { // empty struct
			dst.ApiAtlasClusterRegionConfigsV15AzureManual = nil
		} else {
			match++
		}
	} else {
		dst.ApiAtlasClusterRegionConfigsV15AzureManual = nil
	}

	// try to unmarshal data into ApiAtlasClusterRegionConfigsV15GCPManual
	err = newStrictDecoder(data).Decode(&dst.ApiAtlasClusterRegionConfigsV15GCPManual)
	if err == nil {
		jsonApiAtlasClusterRegionConfigsV15GCPManual, _ := json.Marshal(dst.ApiAtlasClusterRegionConfigsV15GCPManual)
		if string(jsonApiAtlasClusterRegionConfigsV15GCPManual) == "{}" { // empty struct
			dst.ApiAtlasClusterRegionConfigsV15GCPManual = nil
		} else {
			match++
		}
	} else {
		dst.ApiAtlasClusterRegionConfigsV15GCPManual = nil
	}

	// try to unmarshal data into ApiAtlasClusterRegionConfigsV15TenantManual
	err = newStrictDecoder(data).Decode(&dst.ApiAtlasClusterRegionConfigsV15TenantManual)
	if err == nil {
		jsonApiAtlasClusterRegionConfigsV15TenantManual, _ := json.Marshal(dst.ApiAtlasClusterRegionConfigsV15TenantManual)
		if string(jsonApiAtlasClusterRegionConfigsV15TenantManual) == "{}" { // empty struct
			dst.ApiAtlasClusterRegionConfigsV15TenantManual = nil
		} else {
			match++
		}
	} else {
		dst.ApiAtlasClusterRegionConfigsV15TenantManual = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApiAtlasClusterRegionConfigsV15AWSManual = nil
		dst.ApiAtlasClusterRegionConfigsV15AzureManual = nil
		dst.ApiAtlasClusterRegionConfigsV15GCPManual = nil
		dst.ApiAtlasClusterRegionConfigsV15TenantManual = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApiAtlasClusterRegionConfigsV15Manual)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApiAtlasClusterRegionConfigsV15Manual)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApiAtlasClusterRegionConfigsV15Manual) MarshalJSON() ([]byte, error) {
	if src.ApiAtlasClusterRegionConfigsV15AWSManual != nil {
		return json.Marshal(&src.ApiAtlasClusterRegionConfigsV15AWSManual)
	}

	if src.ApiAtlasClusterRegionConfigsV15AzureManual != nil {
		return json.Marshal(&src.ApiAtlasClusterRegionConfigsV15AzureManual)
	}

	if src.ApiAtlasClusterRegionConfigsV15GCPManual != nil {
		return json.Marshal(&src.ApiAtlasClusterRegionConfigsV15GCPManual)
	}

	if src.ApiAtlasClusterRegionConfigsV15TenantManual != nil {
		return json.Marshal(&src.ApiAtlasClusterRegionConfigsV15TenantManual)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApiAtlasClusterRegionConfigsV15Manual) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApiAtlasClusterRegionConfigsV15AWSManual != nil {
		return obj.ApiAtlasClusterRegionConfigsV15AWSManual
	}

	if obj.ApiAtlasClusterRegionConfigsV15AzureManual != nil {
		return obj.ApiAtlasClusterRegionConfigsV15AzureManual
	}

	if obj.ApiAtlasClusterRegionConfigsV15GCPManual != nil {
		return obj.ApiAtlasClusterRegionConfigsV15GCPManual
	}

	if obj.ApiAtlasClusterRegionConfigsV15TenantManual != nil {
		return obj.ApiAtlasClusterRegionConfigsV15TenantManual
	}

	// all schemas are nil
	return nil
}

type NullableApiAtlasClusterRegionConfigsV15Manual struct {
	value *ApiAtlasClusterRegionConfigsV15Manual
	isSet bool
}

func (v NullableApiAtlasClusterRegionConfigsV15Manual) Get() *ApiAtlasClusterRegionConfigsV15Manual {
	return v.value
}

func (v *NullableApiAtlasClusterRegionConfigsV15Manual) Set(val *ApiAtlasClusterRegionConfigsV15Manual) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasClusterRegionConfigsV15Manual) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasClusterRegionConfigsV15Manual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasClusterRegionConfigsV15Manual(val *ApiAtlasClusterRegionConfigsV15Manual) *NullableApiAtlasClusterRegionConfigsV15Manual {
	return &NullableApiAtlasClusterRegionConfigsV15Manual{value: val, isSet: true}
}

func (v NullableApiAtlasClusterRegionConfigsV15Manual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasClusterRegionConfigsV15Manual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

