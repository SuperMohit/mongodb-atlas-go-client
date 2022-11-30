# ApiAtlasClusterDescriptionProcessArgsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultReadConcern** | Pointer to **string** | [Default level of acknowledgment requested from MongoDB for read operations](https://docs.mongodb.com/manual/reference/read-concern/) set for this cluster.  MongoDB 4.4 clusters default to &#x60;available&#x60;. MongoDB 5.0 and later clusters default to &#x60;local&#x60;. | [optional] [default to "available"]
**DefaultWriteConcern** | Pointer to **string** | [Default level of acknowledgment requested from MongoDB for write operations](https://docs.mongodb.com/manual/reference/write-concern/) set for this cluster.  MongoDB 4.4 clusters default to &#x60;1&#x60;. MongoDB 5.0 and later clusters default to &#x60;majority&#x60;. | [optional] [default to "1"]
**FailIndexKeyTooLong** | Pointer to **bool** | Flag that indicates whether you can insert or update documents where all indexed entries don&#39;t exceed 1024 bytes. If you set this to false, [mongod](https://docs.mongodb.com/upcoming/reference/program/mongod/#mongodb-binary-bin.mongod) writes documents that exceed this limit but doesn&#39;t index them. | [optional] [default to true]
**JavascriptEnabled** | Pointer to **bool** | Flag that indicates whether the cluster allows execution of operations that perform server-side executions of JavaScript. | [optional] [default to true]
**MinimumEnabledTlsProtocol** | Pointer to **string** | Minimum Transport Layer Security (TLS) version that the cluster accepts for incoming connections. Clusters using TLS 1.0 or 1.1 should consider setting TLS 1.2 as the minimum TLS protocol version. | [optional] 
**NoTableScan** | Pointer to **bool** | Flag that indicates whether the cluster disables executing any query that requires a collection scan to return results. | [optional] [default to false]
**OplogMinRetentionHours** | Pointer to **NullableFloat64** | Minimum retention window for cluster&#39;s oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates. | [optional] 
**OplogSizeMB** | Pointer to **NullableInt32** | Storage limit of cluster&#39;s oplog expressed in megabytes. A value of null indicates that the cluster uses the default oplog size that MongoDB Cloud calculates. | [optional] 
**SampleRefreshIntervalBIConnector** | Pointer to **int32** | Interval in seconds at which the mongosqld process re-samples data to create its relational schema. | [optional] [default to 0]
**SampleSizeBIConnector** | Pointer to **int32** | Number of documents per database to sample when gathering schema information. | [optional] [default to 1000]

## Methods

### NewApiAtlasClusterDescriptionProcessArgsView

`func NewApiAtlasClusterDescriptionProcessArgsView() *ApiAtlasClusterDescriptionProcessArgsView`

NewApiAtlasClusterDescriptionProcessArgsView instantiates a new ApiAtlasClusterDescriptionProcessArgsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasClusterDescriptionProcessArgsViewWithDefaults

`func NewApiAtlasClusterDescriptionProcessArgsViewWithDefaults() *ApiAtlasClusterDescriptionProcessArgsView`

NewApiAtlasClusterDescriptionProcessArgsViewWithDefaults instantiates a new ApiAtlasClusterDescriptionProcessArgsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultReadConcern

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetDefaultReadConcern() string`

GetDefaultReadConcern returns the DefaultReadConcern field if non-nil, zero value otherwise.

### GetDefaultReadConcernOk

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetDefaultReadConcernOk() (*string, bool)`

GetDefaultReadConcernOk returns a tuple with the DefaultReadConcern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReadConcern

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetDefaultReadConcern(v string)`

SetDefaultReadConcern sets DefaultReadConcern field to given value.

### HasDefaultReadConcern

`func (o *ApiAtlasClusterDescriptionProcessArgsView) HasDefaultReadConcern() bool`

HasDefaultReadConcern returns a boolean if a field has been set.

### GetDefaultWriteConcern

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetDefaultWriteConcern() string`

GetDefaultWriteConcern returns the DefaultWriteConcern field if non-nil, zero value otherwise.

### GetDefaultWriteConcernOk

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetDefaultWriteConcernOk() (*string, bool)`

GetDefaultWriteConcernOk returns a tuple with the DefaultWriteConcern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWriteConcern

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetDefaultWriteConcern(v string)`

SetDefaultWriteConcern sets DefaultWriteConcern field to given value.

### HasDefaultWriteConcern

`func (o *ApiAtlasClusterDescriptionProcessArgsView) HasDefaultWriteConcern() bool`

HasDefaultWriteConcern returns a boolean if a field has been set.

### GetFailIndexKeyTooLong

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetFailIndexKeyTooLong() bool`

GetFailIndexKeyTooLong returns the FailIndexKeyTooLong field if non-nil, zero value otherwise.

### GetFailIndexKeyTooLongOk

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetFailIndexKeyTooLongOk() (*bool, bool)`

GetFailIndexKeyTooLongOk returns a tuple with the FailIndexKeyTooLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailIndexKeyTooLong

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetFailIndexKeyTooLong(v bool)`

SetFailIndexKeyTooLong sets FailIndexKeyTooLong field to given value.

### HasFailIndexKeyTooLong

`func (o *ApiAtlasClusterDescriptionProcessArgsView) HasFailIndexKeyTooLong() bool`

HasFailIndexKeyTooLong returns a boolean if a field has been set.

### GetJavascriptEnabled

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetJavascriptEnabled() bool`

GetJavascriptEnabled returns the JavascriptEnabled field if non-nil, zero value otherwise.

### GetJavascriptEnabledOk

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetJavascriptEnabledOk() (*bool, bool)`

GetJavascriptEnabledOk returns a tuple with the JavascriptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascriptEnabled

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetJavascriptEnabled(v bool)`

SetJavascriptEnabled sets JavascriptEnabled field to given value.

### HasJavascriptEnabled

`func (o *ApiAtlasClusterDescriptionProcessArgsView) HasJavascriptEnabled() bool`

HasJavascriptEnabled returns a boolean if a field has been set.

### GetMinimumEnabledTlsProtocol

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetMinimumEnabledTlsProtocol() string`

GetMinimumEnabledTlsProtocol returns the MinimumEnabledTlsProtocol field if non-nil, zero value otherwise.

### GetMinimumEnabledTlsProtocolOk

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetMinimumEnabledTlsProtocolOk() (*string, bool)`

GetMinimumEnabledTlsProtocolOk returns a tuple with the MinimumEnabledTlsProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumEnabledTlsProtocol

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetMinimumEnabledTlsProtocol(v string)`

SetMinimumEnabledTlsProtocol sets MinimumEnabledTlsProtocol field to given value.

### HasMinimumEnabledTlsProtocol

`func (o *ApiAtlasClusterDescriptionProcessArgsView) HasMinimumEnabledTlsProtocol() bool`

HasMinimumEnabledTlsProtocol returns a boolean if a field has been set.

### GetNoTableScan

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetNoTableScan() bool`

GetNoTableScan returns the NoTableScan field if non-nil, zero value otherwise.

### GetNoTableScanOk

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetNoTableScanOk() (*bool, bool)`

GetNoTableScanOk returns a tuple with the NoTableScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoTableScan

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetNoTableScan(v bool)`

SetNoTableScan sets NoTableScan field to given value.

### HasNoTableScan

`func (o *ApiAtlasClusterDescriptionProcessArgsView) HasNoTableScan() bool`

HasNoTableScan returns a boolean if a field has been set.

### GetOplogMinRetentionHours

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetOplogMinRetentionHours() float64`

GetOplogMinRetentionHours returns the OplogMinRetentionHours field if non-nil, zero value otherwise.

### GetOplogMinRetentionHoursOk

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetOplogMinRetentionHoursOk() (*float64, bool)`

GetOplogMinRetentionHoursOk returns a tuple with the OplogMinRetentionHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogMinRetentionHours

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetOplogMinRetentionHours(v float64)`

SetOplogMinRetentionHours sets OplogMinRetentionHours field to given value.

### HasOplogMinRetentionHours

`func (o *ApiAtlasClusterDescriptionProcessArgsView) HasOplogMinRetentionHours() bool`

HasOplogMinRetentionHours returns a boolean if a field has been set.

### SetOplogMinRetentionHoursNil

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetOplogMinRetentionHoursNil(b bool)`

 SetOplogMinRetentionHoursNil sets the value for OplogMinRetentionHours to be an explicit nil

### UnsetOplogMinRetentionHours
`func (o *ApiAtlasClusterDescriptionProcessArgsView) UnsetOplogMinRetentionHours()`

UnsetOplogMinRetentionHours ensures that no value is present for OplogMinRetentionHours, not even an explicit nil
### GetOplogSizeMB

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetOplogSizeMB() int32`

GetOplogSizeMB returns the OplogSizeMB field if non-nil, zero value otherwise.

### GetOplogSizeMBOk

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetOplogSizeMBOk() (*int32, bool)`

GetOplogSizeMBOk returns a tuple with the OplogSizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogSizeMB

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetOplogSizeMB(v int32)`

SetOplogSizeMB sets OplogSizeMB field to given value.

### HasOplogSizeMB

`func (o *ApiAtlasClusterDescriptionProcessArgsView) HasOplogSizeMB() bool`

HasOplogSizeMB returns a boolean if a field has been set.

### SetOplogSizeMBNil

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetOplogSizeMBNil(b bool)`

 SetOplogSizeMBNil sets the value for OplogSizeMB to be an explicit nil

### UnsetOplogSizeMB
`func (o *ApiAtlasClusterDescriptionProcessArgsView) UnsetOplogSizeMB()`

UnsetOplogSizeMB ensures that no value is present for OplogSizeMB, not even an explicit nil
### GetSampleRefreshIntervalBIConnector

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetSampleRefreshIntervalBIConnector() int32`

GetSampleRefreshIntervalBIConnector returns the SampleRefreshIntervalBIConnector field if non-nil, zero value otherwise.

### GetSampleRefreshIntervalBIConnectorOk

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetSampleRefreshIntervalBIConnectorOk() (*int32, bool)`

GetSampleRefreshIntervalBIConnectorOk returns a tuple with the SampleRefreshIntervalBIConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRefreshIntervalBIConnector

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetSampleRefreshIntervalBIConnector(v int32)`

SetSampleRefreshIntervalBIConnector sets SampleRefreshIntervalBIConnector field to given value.

### HasSampleRefreshIntervalBIConnector

`func (o *ApiAtlasClusterDescriptionProcessArgsView) HasSampleRefreshIntervalBIConnector() bool`

HasSampleRefreshIntervalBIConnector returns a boolean if a field has been set.

### GetSampleSizeBIConnector

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetSampleSizeBIConnector() int32`

GetSampleSizeBIConnector returns the SampleSizeBIConnector field if non-nil, zero value otherwise.

### GetSampleSizeBIConnectorOk

`func (o *ApiAtlasClusterDescriptionProcessArgsView) GetSampleSizeBIConnectorOk() (*int32, bool)`

GetSampleSizeBIConnectorOk returns a tuple with the SampleSizeBIConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSizeBIConnector

`func (o *ApiAtlasClusterDescriptionProcessArgsView) SetSampleSizeBIConnector(v int32)`

SetSampleSizeBIConnector sets SampleSizeBIConnector field to given value.

### HasSampleSizeBIConnector

`func (o *ApiAtlasClusterDescriptionProcessArgsView) HasSampleSizeBIConnector() bool`

HasSampleSizeBIConnector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


