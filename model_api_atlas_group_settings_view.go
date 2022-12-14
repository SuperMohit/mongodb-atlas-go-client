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

// ApiAtlasGroupSettingsView Collection of settings that configures the project.
type ApiAtlasGroupSettingsView struct {
	// Flag that indicates whether to collect database-specific metrics  for the specified project.
	IsCollectDatabaseSpecificsStatisticsEnabled *bool `json:"isCollectDatabaseSpecificsStatisticsEnabled,omitempty"`
	// Flag that indicates whether to enable the Data Explorer for the specified project.
	IsDataExplorerEnabled *bool `json:"isDataExplorerEnabled,omitempty"`
	// Flag that indicates whether to enable extended storage sizes  for the specified project.
	IsExtendedStorageSizesEnabled *bool `json:"isExtendedStorageSizesEnabled,omitempty"`
	// Flag that indicates whether to enable the Performance Advisor and Profiler  for the specified project.
	IsPerformanceAdvisorEnabled *bool `json:"isPerformanceAdvisorEnabled,omitempty"`
	// Flag that indicates whether to enable the Real Time Performance Panel for the specified project.
	IsRealtimePerformancePanelEnabled *bool `json:"isRealtimePerformancePanelEnabled,omitempty"`
	// Flag that indicates whether to enable the Schema Advisor for the specified project.
	IsSchemaAdvisorEnabled *bool `json:"isSchemaAdvisorEnabled,omitempty"`
}

// NewApiAtlasGroupSettingsView instantiates a new ApiAtlasGroupSettingsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasGroupSettingsView() *ApiAtlasGroupSettingsView {
	this := ApiAtlasGroupSettingsView{}
	return &this
}

// NewApiAtlasGroupSettingsViewWithDefaults instantiates a new ApiAtlasGroupSettingsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasGroupSettingsViewWithDefaults() *ApiAtlasGroupSettingsView {
	this := ApiAtlasGroupSettingsView{}
	return &this
}

// GetIsCollectDatabaseSpecificsStatisticsEnabled returns the IsCollectDatabaseSpecificsStatisticsEnabled field value if set, zero value otherwise.
func (o *ApiAtlasGroupSettingsView) GetIsCollectDatabaseSpecificsStatisticsEnabled() bool {
	if o == nil || isNil(o.IsCollectDatabaseSpecificsStatisticsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsCollectDatabaseSpecificsStatisticsEnabled
}

// GetIsCollectDatabaseSpecificsStatisticsEnabledOk returns a tuple with the IsCollectDatabaseSpecificsStatisticsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupSettingsView) GetIsCollectDatabaseSpecificsStatisticsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsCollectDatabaseSpecificsStatisticsEnabled) {
    return nil, false
	}
	return o.IsCollectDatabaseSpecificsStatisticsEnabled, true
}

// HasIsCollectDatabaseSpecificsStatisticsEnabled returns a boolean if a field has been set.
func (o *ApiAtlasGroupSettingsView) HasIsCollectDatabaseSpecificsStatisticsEnabled() bool {
	if o != nil && !isNil(o.IsCollectDatabaseSpecificsStatisticsEnabled) {
		return true
	}

	return false
}

// SetIsCollectDatabaseSpecificsStatisticsEnabled gets a reference to the given bool and assigns it to the IsCollectDatabaseSpecificsStatisticsEnabled field.
func (o *ApiAtlasGroupSettingsView) SetIsCollectDatabaseSpecificsStatisticsEnabled(v bool) {
	o.IsCollectDatabaseSpecificsStatisticsEnabled = &v
}

// GetIsDataExplorerEnabled returns the IsDataExplorerEnabled field value if set, zero value otherwise.
func (o *ApiAtlasGroupSettingsView) GetIsDataExplorerEnabled() bool {
	if o == nil || isNil(o.IsDataExplorerEnabled) {
		var ret bool
		return ret
	}
	return *o.IsDataExplorerEnabled
}

// GetIsDataExplorerEnabledOk returns a tuple with the IsDataExplorerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupSettingsView) GetIsDataExplorerEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsDataExplorerEnabled) {
    return nil, false
	}
	return o.IsDataExplorerEnabled, true
}

// HasIsDataExplorerEnabled returns a boolean if a field has been set.
func (o *ApiAtlasGroupSettingsView) HasIsDataExplorerEnabled() bool {
	if o != nil && !isNil(o.IsDataExplorerEnabled) {
		return true
	}

	return false
}

// SetIsDataExplorerEnabled gets a reference to the given bool and assigns it to the IsDataExplorerEnabled field.
func (o *ApiAtlasGroupSettingsView) SetIsDataExplorerEnabled(v bool) {
	o.IsDataExplorerEnabled = &v
}

// GetIsExtendedStorageSizesEnabled returns the IsExtendedStorageSizesEnabled field value if set, zero value otherwise.
func (o *ApiAtlasGroupSettingsView) GetIsExtendedStorageSizesEnabled() bool {
	if o == nil || isNil(o.IsExtendedStorageSizesEnabled) {
		var ret bool
		return ret
	}
	return *o.IsExtendedStorageSizesEnabled
}

// GetIsExtendedStorageSizesEnabledOk returns a tuple with the IsExtendedStorageSizesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupSettingsView) GetIsExtendedStorageSizesEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsExtendedStorageSizesEnabled) {
    return nil, false
	}
	return o.IsExtendedStorageSizesEnabled, true
}

// HasIsExtendedStorageSizesEnabled returns a boolean if a field has been set.
func (o *ApiAtlasGroupSettingsView) HasIsExtendedStorageSizesEnabled() bool {
	if o != nil && !isNil(o.IsExtendedStorageSizesEnabled) {
		return true
	}

	return false
}

// SetIsExtendedStorageSizesEnabled gets a reference to the given bool and assigns it to the IsExtendedStorageSizesEnabled field.
func (o *ApiAtlasGroupSettingsView) SetIsExtendedStorageSizesEnabled(v bool) {
	o.IsExtendedStorageSizesEnabled = &v
}

// GetIsPerformanceAdvisorEnabled returns the IsPerformanceAdvisorEnabled field value if set, zero value otherwise.
func (o *ApiAtlasGroupSettingsView) GetIsPerformanceAdvisorEnabled() bool {
	if o == nil || isNil(o.IsPerformanceAdvisorEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPerformanceAdvisorEnabled
}

// GetIsPerformanceAdvisorEnabledOk returns a tuple with the IsPerformanceAdvisorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupSettingsView) GetIsPerformanceAdvisorEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsPerformanceAdvisorEnabled) {
    return nil, false
	}
	return o.IsPerformanceAdvisorEnabled, true
}

// HasIsPerformanceAdvisorEnabled returns a boolean if a field has been set.
func (o *ApiAtlasGroupSettingsView) HasIsPerformanceAdvisorEnabled() bool {
	if o != nil && !isNil(o.IsPerformanceAdvisorEnabled) {
		return true
	}

	return false
}

// SetIsPerformanceAdvisorEnabled gets a reference to the given bool and assigns it to the IsPerformanceAdvisorEnabled field.
func (o *ApiAtlasGroupSettingsView) SetIsPerformanceAdvisorEnabled(v bool) {
	o.IsPerformanceAdvisorEnabled = &v
}

// GetIsRealtimePerformancePanelEnabled returns the IsRealtimePerformancePanelEnabled field value if set, zero value otherwise.
func (o *ApiAtlasGroupSettingsView) GetIsRealtimePerformancePanelEnabled() bool {
	if o == nil || isNil(o.IsRealtimePerformancePanelEnabled) {
		var ret bool
		return ret
	}
	return *o.IsRealtimePerformancePanelEnabled
}

// GetIsRealtimePerformancePanelEnabledOk returns a tuple with the IsRealtimePerformancePanelEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupSettingsView) GetIsRealtimePerformancePanelEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsRealtimePerformancePanelEnabled) {
    return nil, false
	}
	return o.IsRealtimePerformancePanelEnabled, true
}

// HasIsRealtimePerformancePanelEnabled returns a boolean if a field has been set.
func (o *ApiAtlasGroupSettingsView) HasIsRealtimePerformancePanelEnabled() bool {
	if o != nil && !isNil(o.IsRealtimePerformancePanelEnabled) {
		return true
	}

	return false
}

// SetIsRealtimePerformancePanelEnabled gets a reference to the given bool and assigns it to the IsRealtimePerformancePanelEnabled field.
func (o *ApiAtlasGroupSettingsView) SetIsRealtimePerformancePanelEnabled(v bool) {
	o.IsRealtimePerformancePanelEnabled = &v
}

// GetIsSchemaAdvisorEnabled returns the IsSchemaAdvisorEnabled field value if set, zero value otherwise.
func (o *ApiAtlasGroupSettingsView) GetIsSchemaAdvisorEnabled() bool {
	if o == nil || isNil(o.IsSchemaAdvisorEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSchemaAdvisorEnabled
}

// GetIsSchemaAdvisorEnabledOk returns a tuple with the IsSchemaAdvisorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasGroupSettingsView) GetIsSchemaAdvisorEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsSchemaAdvisorEnabled) {
    return nil, false
	}
	return o.IsSchemaAdvisorEnabled, true
}

// HasIsSchemaAdvisorEnabled returns a boolean if a field has been set.
func (o *ApiAtlasGroupSettingsView) HasIsSchemaAdvisorEnabled() bool {
	if o != nil && !isNil(o.IsSchemaAdvisorEnabled) {
		return true
	}

	return false
}

// SetIsSchemaAdvisorEnabled gets a reference to the given bool and assigns it to the IsSchemaAdvisorEnabled field.
func (o *ApiAtlasGroupSettingsView) SetIsSchemaAdvisorEnabled(v bool) {
	o.IsSchemaAdvisorEnabled = &v
}

func (o ApiAtlasGroupSettingsView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsCollectDatabaseSpecificsStatisticsEnabled) {
		toSerialize["isCollectDatabaseSpecificsStatisticsEnabled"] = o.IsCollectDatabaseSpecificsStatisticsEnabled
	}
	if !isNil(o.IsDataExplorerEnabled) {
		toSerialize["isDataExplorerEnabled"] = o.IsDataExplorerEnabled
	}
	if !isNil(o.IsExtendedStorageSizesEnabled) {
		toSerialize["isExtendedStorageSizesEnabled"] = o.IsExtendedStorageSizesEnabled
	}
	if !isNil(o.IsPerformanceAdvisorEnabled) {
		toSerialize["isPerformanceAdvisorEnabled"] = o.IsPerformanceAdvisorEnabled
	}
	if !isNil(o.IsRealtimePerformancePanelEnabled) {
		toSerialize["isRealtimePerformancePanelEnabled"] = o.IsRealtimePerformancePanelEnabled
	}
	if !isNil(o.IsSchemaAdvisorEnabled) {
		toSerialize["isSchemaAdvisorEnabled"] = o.IsSchemaAdvisorEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasGroupSettingsView struct {
	value *ApiAtlasGroupSettingsView
	isSet bool
}

func (v NullableApiAtlasGroupSettingsView) Get() *ApiAtlasGroupSettingsView {
	return v.value
}

func (v *NullableApiAtlasGroupSettingsView) Set(val *ApiAtlasGroupSettingsView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasGroupSettingsView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasGroupSettingsView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasGroupSettingsView(val *ApiAtlasGroupSettingsView) *NullableApiAtlasGroupSettingsView {
	return &NullableApiAtlasGroupSettingsView{value: val, isSet: true}
}

func (v NullableApiAtlasGroupSettingsView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasGroupSettingsView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


