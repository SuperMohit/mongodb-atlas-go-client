# ApiMeasurementsIndexesView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionName** | Pointer to **string** | Human-readable label that identifies the collection. | [optional] [readonly] 
**DatabaseName** | Pointer to **string** | Human-readable label that identifies the database that the specified MongoDB process serves. | [optional] [readonly] 
**End** | Pointer to **time.Time** | Date and time that specifies when to stop retrieving measurements. If you set **end**, you must set **start**. You can&#39;t set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Granularity** | Pointer to **string** | Duration that specifies the interval between measurement data points. The parameter expresses its value in ISO 8601 timestamp format in UTC. If you set this parameter, you must set either **period** or **start** and **end**. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project. The project contains MongoDB processes that you want to return. The MongoDB process can be either the &#x60;mongod&#x60; or &#x60;mongos&#x60;. | [optional] [readonly] 
**IndexIds** | Pointer to **[]string** | List that contains the Atlas Search index identifiers. | [optional] [readonly] 
**IndexStatsMeasurements** | Pointer to [**[]ApiMeasurementView**](ApiMeasurementView.md) | List that contains the Atlas Search index stats measurements. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**ProcessId** | Pointer to **string** | Combination of hostname and Internet Assigned Numbers Authority (IANA) port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (&#x60;mongod&#x60; or &#x60;mongos&#x60;). The port must be the IANA port on which the MongoDB process listens for requests. | [optional] [readonly] 
**Start** | Pointer to **time.Time** | Date and time that specifies when to start retrieving measurements. If you set **start**, you must set **end**. You can&#39;t set this parameter and **period** in the same request. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewApiMeasurementsIndexesView

`func NewApiMeasurementsIndexesView(links []Link, ) *ApiMeasurementsIndexesView`

NewApiMeasurementsIndexesView instantiates a new ApiMeasurementsIndexesView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMeasurementsIndexesViewWithDefaults

`func NewApiMeasurementsIndexesViewWithDefaults() *ApiMeasurementsIndexesView`

NewApiMeasurementsIndexesViewWithDefaults instantiates a new ApiMeasurementsIndexesView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionName

`func (o *ApiMeasurementsIndexesView) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *ApiMeasurementsIndexesView) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *ApiMeasurementsIndexesView) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *ApiMeasurementsIndexesView) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.

### GetDatabaseName

`func (o *ApiMeasurementsIndexesView) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *ApiMeasurementsIndexesView) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *ApiMeasurementsIndexesView) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *ApiMeasurementsIndexesView) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetEnd

`func (o *ApiMeasurementsIndexesView) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ApiMeasurementsIndexesView) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ApiMeasurementsIndexesView) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ApiMeasurementsIndexesView) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetGranularity

`func (o *ApiMeasurementsIndexesView) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *ApiMeasurementsIndexesView) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *ApiMeasurementsIndexesView) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *ApiMeasurementsIndexesView) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiMeasurementsIndexesView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiMeasurementsIndexesView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiMeasurementsIndexesView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiMeasurementsIndexesView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIndexIds

`func (o *ApiMeasurementsIndexesView) GetIndexIds() []string`

GetIndexIds returns the IndexIds field if non-nil, zero value otherwise.

### GetIndexIdsOk

`func (o *ApiMeasurementsIndexesView) GetIndexIdsOk() (*[]string, bool)`

GetIndexIdsOk returns a tuple with the IndexIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexIds

`func (o *ApiMeasurementsIndexesView) SetIndexIds(v []string)`

SetIndexIds sets IndexIds field to given value.

### HasIndexIds

`func (o *ApiMeasurementsIndexesView) HasIndexIds() bool`

HasIndexIds returns a boolean if a field has been set.

### GetIndexStatsMeasurements

`func (o *ApiMeasurementsIndexesView) GetIndexStatsMeasurements() []ApiMeasurementView`

GetIndexStatsMeasurements returns the IndexStatsMeasurements field if non-nil, zero value otherwise.

### GetIndexStatsMeasurementsOk

`func (o *ApiMeasurementsIndexesView) GetIndexStatsMeasurementsOk() (*[]ApiMeasurementView, bool)`

GetIndexStatsMeasurementsOk returns a tuple with the IndexStatsMeasurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexStatsMeasurements

`func (o *ApiMeasurementsIndexesView) SetIndexStatsMeasurements(v []ApiMeasurementView)`

SetIndexStatsMeasurements sets IndexStatsMeasurements field to given value.

### HasIndexStatsMeasurements

`func (o *ApiMeasurementsIndexesView) HasIndexStatsMeasurements() bool`

HasIndexStatsMeasurements returns a boolean if a field has been set.

### GetLinks

`func (o *ApiMeasurementsIndexesView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiMeasurementsIndexesView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiMeasurementsIndexesView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetProcessId

`func (o *ApiMeasurementsIndexesView) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ApiMeasurementsIndexesView) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ApiMeasurementsIndexesView) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *ApiMeasurementsIndexesView) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetStart

`func (o *ApiMeasurementsIndexesView) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ApiMeasurementsIndexesView) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ApiMeasurementsIndexesView) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *ApiMeasurementsIndexesView) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


