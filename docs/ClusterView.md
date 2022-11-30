# ClusterView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertCount** | Pointer to **int32** | Whole number that indicates the quantity of alerts open on the cluster. | [optional] [readonly] 
**AuthEnabled** | Pointer to **bool** | Flag that indicates whether authentication is required to access the nodes in this cluster. | [optional] [readonly] 
**Availability** | Pointer to **string** | Term that expresses how many nodes of the cluster can be accessed when MongoDB Cloud receives this request. This parameter returns &#x60;available&#x60; when all nodes are accessible, &#x60;warning&#x60; only when some nodes in the cluster can be accessed, &#x60;unavailable&#x60; when the cluster can&#39;t be accessed, or &#x60;dead&#x60; when the cluster has been deactivated. | [optional] [readonly] 
**BackupEnabled** | Pointer to **bool** | Flag that indicates whether the cluster can perform backups. If set to &#x60;true&#x60;, the cluster can perform backups. You must set this value to &#x60;true&#x60; for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to &#x60;false&#x60;, the cluster doesn&#39;t use MongoDB Cloud backups. | [optional] [readonly] 
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the cluster. | [optional] [readonly] 
**DataSizeBytes** | Pointer to **int64** | Total size of the data stored on each node in the cluster. The resource expresses this value in bytes. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the cluster. | [optional] [readonly] 
**NodeCount** | Pointer to **int32** | Whole number that indicates the quantity of nodes that comprise the cluster. | [optional] [readonly] 
**SslEnabled** | Pointer to **bool** | Flag that indicates whether TLS authentication is required to access the nodes in this cluster. | [optional] [readonly] 
**Type** | Pointer to **string** | Human-readable label that indicates the cluster type. | [optional] [readonly] 
**Versions** | Pointer to **[]string** | List that contains the versions of MongoDB that each node in the cluster runs. | [optional] [readonly] 

## Methods

### NewClusterView

`func NewClusterView() *ClusterView`

NewClusterView instantiates a new ClusterView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterViewWithDefaults

`func NewClusterViewWithDefaults() *ClusterView`

NewClusterViewWithDefaults instantiates a new ClusterView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertCount

`func (o *ClusterView) GetAlertCount() int32`

GetAlertCount returns the AlertCount field if non-nil, zero value otherwise.

### GetAlertCountOk

`func (o *ClusterView) GetAlertCountOk() (*int32, bool)`

GetAlertCountOk returns a tuple with the AlertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCount

`func (o *ClusterView) SetAlertCount(v int32)`

SetAlertCount sets AlertCount field to given value.

### HasAlertCount

`func (o *ClusterView) HasAlertCount() bool`

HasAlertCount returns a boolean if a field has been set.

### GetAuthEnabled

`func (o *ClusterView) GetAuthEnabled() bool`

GetAuthEnabled returns the AuthEnabled field if non-nil, zero value otherwise.

### GetAuthEnabledOk

`func (o *ClusterView) GetAuthEnabledOk() (*bool, bool)`

GetAuthEnabledOk returns a tuple with the AuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEnabled

`func (o *ClusterView) SetAuthEnabled(v bool)`

SetAuthEnabled sets AuthEnabled field to given value.

### HasAuthEnabled

`func (o *ClusterView) HasAuthEnabled() bool`

HasAuthEnabled returns a boolean if a field has been set.

### GetAvailability

`func (o *ClusterView) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ClusterView) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ClusterView) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ClusterView) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetBackupEnabled

`func (o *ClusterView) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *ClusterView) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *ClusterView) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *ClusterView) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.

### GetClusterId

`func (o *ClusterView) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterView) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterView) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClusterView) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetDataSizeBytes

`func (o *ClusterView) GetDataSizeBytes() int64`

GetDataSizeBytes returns the DataSizeBytes field if non-nil, zero value otherwise.

### GetDataSizeBytesOk

`func (o *ClusterView) GetDataSizeBytesOk() (*int64, bool)`

GetDataSizeBytesOk returns a tuple with the DataSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSizeBytes

`func (o *ClusterView) SetDataSizeBytes(v int64)`

SetDataSizeBytes sets DataSizeBytes field to given value.

### HasDataSizeBytes

`func (o *ClusterView) HasDataSizeBytes() bool`

HasDataSizeBytes returns a boolean if a field has been set.

### GetName

`func (o *ClusterView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterView) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterView) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterView) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterView) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetSslEnabled

`func (o *ClusterView) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *ClusterView) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *ClusterView) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *ClusterView) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### GetType

`func (o *ClusterView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersions

`func (o *ClusterView) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ClusterView) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ClusterView) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ClusterView) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


