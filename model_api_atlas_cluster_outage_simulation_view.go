/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ApiAtlasClusterOutageSimulationView struct for ApiAtlasClusterOutageSimulationView
type ApiAtlasClusterOutageSimulationView struct {
	// Human-readable label that identifies the cluster that undergoes outage simulation.
	ClusterName *string `json:"clusterName,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project that contains the cluster to undergo outage simulation.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal character string that identifies the outage simulation.
	Id *string `json:"id,omitempty"`
	// List of settings that specify the type of cluster outage simulation.
	OutageFilters []ApiAtlasClusterOutageSimulationOutageFilterView `json:"outageFilters,omitempty"`
	// Date and time when MongoDB Cloud started the regional outage simulation.
	StartRequestDate *time.Time `json:"startRequestDate,omitempty"`
	// Phase of the outage simulation.  | State       | Indication | |-------------|------------| | `START_REQUESTED`    | User has requested cluster outage simulation.| | `STARTING`           | MongoDB Cloud is starting cluster outage simulation.| | `SIMULATING`         | MongoDB Cloud is simulating cluster outage.| | `RECOVERY_REQUESTED` | User has requested recovery from the simulated outage.| | `RECOVERING`         | MongoDB Cloud is recovering the cluster from the simulated outage.| | `COMPLETE`           | MongoDB Cloud has completed the cluster outage simulation.|
	State *string `json:"state,omitempty"`
}

// NewApiAtlasClusterOutageSimulationView instantiates a new ApiAtlasClusterOutageSimulationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasClusterOutageSimulationView() *ApiAtlasClusterOutageSimulationView {
	this := ApiAtlasClusterOutageSimulationView{}
	return &this
}

// NewApiAtlasClusterOutageSimulationViewWithDefaults instantiates a new ApiAtlasClusterOutageSimulationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasClusterOutageSimulationViewWithDefaults() *ApiAtlasClusterOutageSimulationView {
	this := ApiAtlasClusterOutageSimulationView{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *ApiAtlasClusterOutageSimulationView) GetClusterName() string {
	if o == nil || isNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterOutageSimulationView) GetClusterNameOk() (*string, bool) {
	if o == nil || isNil(o.ClusterName) {
    return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ApiAtlasClusterOutageSimulationView) HasClusterName() bool {
	if o != nil && !isNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ApiAtlasClusterOutageSimulationView) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApiAtlasClusterOutageSimulationView) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterOutageSimulationView) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiAtlasClusterOutageSimulationView) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiAtlasClusterOutageSimulationView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAtlasClusterOutageSimulationView) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterOutageSimulationView) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAtlasClusterOutageSimulationView) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAtlasClusterOutageSimulationView) SetId(v string) {
	o.Id = &v
}

// GetOutageFilters returns the OutageFilters field value if set, zero value otherwise.
func (o *ApiAtlasClusterOutageSimulationView) GetOutageFilters() []ApiAtlasClusterOutageSimulationOutageFilterView {
	if o == nil || isNil(o.OutageFilters) {
		var ret []ApiAtlasClusterOutageSimulationOutageFilterView
		return ret
	}
	return o.OutageFilters
}

// GetOutageFiltersOk returns a tuple with the OutageFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterOutageSimulationView) GetOutageFiltersOk() ([]ApiAtlasClusterOutageSimulationOutageFilterView, bool) {
	if o == nil || isNil(o.OutageFilters) {
    return nil, false
	}
	return o.OutageFilters, true
}

// HasOutageFilters returns a boolean if a field has been set.
func (o *ApiAtlasClusterOutageSimulationView) HasOutageFilters() bool {
	if o != nil && !isNil(o.OutageFilters) {
		return true
	}

	return false
}

// SetOutageFilters gets a reference to the given []ApiAtlasClusterOutageSimulationOutageFilterView and assigns it to the OutageFilters field.
func (o *ApiAtlasClusterOutageSimulationView) SetOutageFilters(v []ApiAtlasClusterOutageSimulationOutageFilterView) {
	o.OutageFilters = v
}

// GetStartRequestDate returns the StartRequestDate field value if set, zero value otherwise.
func (o *ApiAtlasClusterOutageSimulationView) GetStartRequestDate() time.Time {
	if o == nil || isNil(o.StartRequestDate) {
		var ret time.Time
		return ret
	}
	return *o.StartRequestDate
}

// GetStartRequestDateOk returns a tuple with the StartRequestDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterOutageSimulationView) GetStartRequestDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartRequestDate) {
    return nil, false
	}
	return o.StartRequestDate, true
}

// HasStartRequestDate returns a boolean if a field has been set.
func (o *ApiAtlasClusterOutageSimulationView) HasStartRequestDate() bool {
	if o != nil && !isNil(o.StartRequestDate) {
		return true
	}

	return false
}

// SetStartRequestDate gets a reference to the given time.Time and assigns it to the StartRequestDate field.
func (o *ApiAtlasClusterOutageSimulationView) SetStartRequestDate(v time.Time) {
	o.StartRequestDate = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ApiAtlasClusterOutageSimulationView) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasClusterOutageSimulationView) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApiAtlasClusterOutageSimulationView) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ApiAtlasClusterOutageSimulationView) SetState(v string) {
	o.State = &v
}

func (o ApiAtlasClusterOutageSimulationView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClusterName) {
		toSerialize["clusterName"] = o.ClusterName
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OutageFilters) {
		toSerialize["outageFilters"] = o.OutageFilters
	}
	if !isNil(o.StartRequestDate) {
		toSerialize["startRequestDate"] = o.StartRequestDate
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableApiAtlasClusterOutageSimulationView struct {
	value *ApiAtlasClusterOutageSimulationView
	isSet bool
}

func (v NullableApiAtlasClusterOutageSimulationView) Get() *ApiAtlasClusterOutageSimulationView {
	return v.value
}

func (v *NullableApiAtlasClusterOutageSimulationView) Set(val *ApiAtlasClusterOutageSimulationView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAtlasClusterOutageSimulationView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAtlasClusterOutageSimulationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAtlasClusterOutageSimulationView(val *ApiAtlasClusterOutageSimulationView) *NullableApiAtlasClusterOutageSimulationView {
	return &NullableApiAtlasClusterOutageSimulationView{value: val, isSet: true}
}

func (v NullableApiAtlasClusterOutageSimulationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAtlasClusterOutageSimulationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

