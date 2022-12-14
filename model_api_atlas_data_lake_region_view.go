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

// ApiAtlasDataLakeRegionView Atlas Data Lake Regions
type ApiAtlasDataLakeRegionView string

// List of ApiAtlasDataLakeRegionView
const (
	DUBLIN_IRL ApiAtlasDataLakeRegionView = "DUBLIN_IRL"
	FRANKFURT_DEU ApiAtlasDataLakeRegionView = "FRANKFURT_DEU"
	LONDON_GBR ApiAtlasDataLakeRegionView = "LONDON_GBR"
	MUMBAI_IND ApiAtlasDataLakeRegionView = "MUMBAI_IND"
	OREGON_USA ApiAtlasDataLakeRegionView = "OREGON_USA"
	SYDNEY_AUS ApiAtlasDataLakeRegionView = "SYDNEY_AUS"
	VIRGINIA_USA ApiAtlasDataLakeRegionView = "VIRGINIA_USA"
	SAOPAULO_BRA ApiAtlasDataLakeRegionView = "SAOPAULO_BRA"
)

// All allowed values of ApiAtlasDataLakeRegionView enum
var AllowedApiAtlasDataLakeRegionViewEnumValues = []ApiAtlasDataLakeRegionView{
	"DUBLIN_IRL",
	"FRANKFURT_DEU",
	"LONDON_GBR",
	"MUMBAI_IND",
	"OREGON_USA",
	"SYDNEY_AUS",
	"VIRGINIA_USA",
	"SAOPAULO_BRA",
}

func (v *ApiAtlasDataLakeRegionView) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiAtlasDataLakeRegionView(value)
	for _, existing := range AllowedApiAtlasDataLakeRegionViewEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiAtlasDataLakeRegionView", value)
}

// NewApiAtlasDataLakeRegionViewFromValue returns a pointer to a valid ApiAtlasDataLakeRegionView
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiAtlasDataLakeRegionViewFromValue(v string) (*ApiAtlasDataLakeRegionView, error) {
	ev := ApiAtlasDataLakeRegionView(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiAtlasDataLakeRegionView: valid values are %v", v, AllowedApiAtlasDataLakeRegionViewEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiAtlasDataLakeRegionView) IsValid() bool {
	for _, existing := range AllowedApiAtlasDataLakeRegionViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApiAtlasDataLakeRegionView value
func (v ApiAtlasDataLakeRegionView) Ptr() *ApiAtlasDataLakeRegionView {
	return &v
}

type NullableApiAtlasDataLakeRegionView struct {
	value *ApiAtlasDataLakeRegionView
	isSet bool
}

func (v NullableApiAtlasDataLakeRegionView) Get() *ApiAtlasDataLakeRegionView {
	return v.value
}

func (v *NullableApiAtlasDataLakeRegionView) Set(val *ApiAtlasDataLakeRegionView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasDataLakeRegionView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasDataLakeRegionView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasDataLakeRegionView(val *ApiAtlasDataLakeRegionView) *NullableApiAtlasDataLakeRegionView {
	return &NullableApiAtlasDataLakeRegionView{value: val, isSet: true}
}

func (v NullableApiAtlasDataLakeRegionView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasDataLakeRegionView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

