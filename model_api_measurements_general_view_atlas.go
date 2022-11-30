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

// ApiMeasurementsGeneralViewAtlas struct for ApiMeasurementsGeneralViewAtlas
type ApiMeasurementsGeneralViewAtlas struct {
	// Human-readable label that identifies the database that the specified MongoDB process serves.
	DatabaseName *string `json:"databaseName,omitempty"`
	// Date and time that specifies when to stop retrieving measurements. If you set **end**, you must set **start**. You can't set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	End *time.Time `json:"end,omitempty"`
	// Duration that specifies the interval between measurement data points. The parameter expresses its value in ISO 8601 timestamp format in UTC. If you set this parameter, you must set either **period** or **start** and **end**.
	Granularity *string `json:"granularity,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project. The project contains MongoDB processes that you want to return. The MongoDB process can be either the `mongod` or `mongos`.
	GroupId *string `json:"groupId,omitempty"`
	// Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	HostId *string `json:"hostId,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []LinkAtlas `json:"links"`
	// List that contains measurements and their data points.
	Measurements []ApiMeasurementViewAtlas `json:"measurements,omitempty"`
	// Human-readable label of the disk or partition to which the measurements apply.
	PartitionName *string `json:"partitionName,omitempty"`
	// Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	ProcessId *string `json:"processId,omitempty"`
	// Date and time that specifies when to start retrieving measurements. If you set **start**, you must set **end**. You can't set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Start *time.Time `json:"start,omitempty"`
}

// NewApiMeasurementsGeneralViewAtlas instantiates a new ApiMeasurementsGeneralViewAtlas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMeasurementsGeneralViewAtlas(links []LinkAtlas) *ApiMeasurementsGeneralViewAtlas {
	this := ApiMeasurementsGeneralViewAtlas{}
	this.Links = links
	return &this
}

// NewApiMeasurementsGeneralViewAtlasWithDefaults instantiates a new ApiMeasurementsGeneralViewAtlas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMeasurementsGeneralViewAtlasWithDefaults() *ApiMeasurementsGeneralViewAtlas {
	this := ApiMeasurementsGeneralViewAtlas{}
	return &this
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *ApiMeasurementsGeneralViewAtlas) GetDatabaseName() string {
	if o == nil || isNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsGeneralViewAtlas) GetDatabaseNameOk() (*string, bool) {
	if o == nil || isNil(o.DatabaseName) {
    return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *ApiMeasurementsGeneralViewAtlas) HasDatabaseName() bool {
	if o != nil && !isNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *ApiMeasurementsGeneralViewAtlas) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ApiMeasurementsGeneralViewAtlas) GetEnd() time.Time {
	if o == nil || isNil(o.End) {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsGeneralViewAtlas) GetEndOk() (*time.Time, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ApiMeasurementsGeneralViewAtlas) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *ApiMeasurementsGeneralViewAtlas) SetEnd(v time.Time) {
	o.End = &v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *ApiMeasurementsGeneralViewAtlas) GetGranularity() string {
	if o == nil || isNil(o.Granularity) {
		var ret string
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsGeneralViewAtlas) GetGranularityOk() (*string, bool) {
	if o == nil || isNil(o.Granularity) {
    return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *ApiMeasurementsGeneralViewAtlas) HasGranularity() bool {
	if o != nil && !isNil(o.Granularity) {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given string and assigns it to the Granularity field.
func (o *ApiMeasurementsGeneralViewAtlas) SetGranularity(v string) {
	o.Granularity = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApiMeasurementsGeneralViewAtlas) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsGeneralViewAtlas) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiMeasurementsGeneralViewAtlas) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiMeasurementsGeneralViewAtlas) SetGroupId(v string) {
	o.GroupId = &v
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *ApiMeasurementsGeneralViewAtlas) GetHostId() string {
	if o == nil || isNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsGeneralViewAtlas) GetHostIdOk() (*string, bool) {
	if o == nil || isNil(o.HostId) {
    return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *ApiMeasurementsGeneralViewAtlas) HasHostId() bool {
	if o != nil && !isNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *ApiMeasurementsGeneralViewAtlas) SetHostId(v string) {
	o.HostId = &v
}

// GetLinks returns the Links field value
func (o *ApiMeasurementsGeneralViewAtlas) GetLinks() []LinkAtlas {
	if o == nil {
		var ret []LinkAtlas
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsGeneralViewAtlas) GetLinksOk() ([]LinkAtlas, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ApiMeasurementsGeneralViewAtlas) SetLinks(v []LinkAtlas) {
	o.Links = v
}

// GetMeasurements returns the Measurements field value if set, zero value otherwise.
func (o *ApiMeasurementsGeneralViewAtlas) GetMeasurements() []ApiMeasurementViewAtlas {
	if o == nil || isNil(o.Measurements) {
		var ret []ApiMeasurementViewAtlas
		return ret
	}
	return o.Measurements
}

// GetMeasurementsOk returns a tuple with the Measurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsGeneralViewAtlas) GetMeasurementsOk() ([]ApiMeasurementViewAtlas, bool) {
	if o == nil || isNil(o.Measurements) {
    return nil, false
	}
	return o.Measurements, true
}

// HasMeasurements returns a boolean if a field has been set.
func (o *ApiMeasurementsGeneralViewAtlas) HasMeasurements() bool {
	if o != nil && !isNil(o.Measurements) {
		return true
	}

	return false
}

// SetMeasurements gets a reference to the given []ApiMeasurementViewAtlas and assigns it to the Measurements field.
func (o *ApiMeasurementsGeneralViewAtlas) SetMeasurements(v []ApiMeasurementViewAtlas) {
	o.Measurements = v
}

// GetPartitionName returns the PartitionName field value if set, zero value otherwise.
func (o *ApiMeasurementsGeneralViewAtlas) GetPartitionName() string {
	if o == nil || isNil(o.PartitionName) {
		var ret string
		return ret
	}
	return *o.PartitionName
}

// GetPartitionNameOk returns a tuple with the PartitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsGeneralViewAtlas) GetPartitionNameOk() (*string, bool) {
	if o == nil || isNil(o.PartitionName) {
    return nil, false
	}
	return o.PartitionName, true
}

// HasPartitionName returns a boolean if a field has been set.
func (o *ApiMeasurementsGeneralViewAtlas) HasPartitionName() bool {
	if o != nil && !isNil(o.PartitionName) {
		return true
	}

	return false
}

// SetPartitionName gets a reference to the given string and assigns it to the PartitionName field.
func (o *ApiMeasurementsGeneralViewAtlas) SetPartitionName(v string) {
	o.PartitionName = &v
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *ApiMeasurementsGeneralViewAtlas) GetProcessId() string {
	if o == nil || isNil(o.ProcessId) {
		var ret string
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsGeneralViewAtlas) GetProcessIdOk() (*string, bool) {
	if o == nil || isNil(o.ProcessId) {
    return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *ApiMeasurementsGeneralViewAtlas) HasProcessId() bool {
	if o != nil && !isNil(o.ProcessId) {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given string and assigns it to the ProcessId field.
func (o *ApiMeasurementsGeneralViewAtlas) SetProcessId(v string) {
	o.ProcessId = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ApiMeasurementsGeneralViewAtlas) GetStart() time.Time {
	if o == nil || isNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMeasurementsGeneralViewAtlas) GetStartOk() (*time.Time, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ApiMeasurementsGeneralViewAtlas) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *ApiMeasurementsGeneralViewAtlas) SetStart(v time.Time) {
	o.Start = &v
}

func (o ApiMeasurementsGeneralViewAtlas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DatabaseName) {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Granularity) {
		toSerialize["granularity"] = o.Granularity
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.HostId) {
		toSerialize["hostId"] = o.HostId
	}
	if true {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.Measurements) {
		toSerialize["measurements"] = o.Measurements
	}
	if !isNil(o.PartitionName) {
		toSerialize["partitionName"] = o.PartitionName
	}
	if !isNil(o.ProcessId) {
		toSerialize["processId"] = o.ProcessId
	}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	return json.Marshal(toSerialize)
}

type NullableApiMeasurementsGeneralViewAtlas struct {
	value *ApiMeasurementsGeneralViewAtlas
	isSet bool
}

func (v NullableApiMeasurementsGeneralViewAtlas) Get() *ApiMeasurementsGeneralViewAtlas {
	return v.value
}

func (v *NullableApiMeasurementsGeneralViewAtlas) Set(val *ApiMeasurementsGeneralViewAtlas) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMeasurementsGeneralViewAtlas) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMeasurementsGeneralViewAtlas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMeasurementsGeneralViewAtlas(val *ApiMeasurementsGeneralViewAtlas) *NullableApiMeasurementsGeneralViewAtlas {
	return &NullableApiMeasurementsGeneralViewAtlas{value: val, isSet: true}
}

func (v NullableApiMeasurementsGeneralViewAtlas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMeasurementsGeneralViewAtlas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


