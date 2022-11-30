# ApiAtlasTenantRestoreView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Human-readable label that identifies the source cluster. | [optional] [readonly] 
**DeliveryType** | Pointer to **string** | Means by which this resource returns the snapshot to the requesting MongoDB Cloud user. | [optional] [readonly] 
**ExpirationDate** | Pointer to **time.Time** | Date and time when the download link no longer works. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the restore job. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**ProjectId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project from which the restore job originated. | [optional] [readonly] 
**RestoreFinishedDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud completed writing this snapshot. MongoDB Cloud changes the status of the restore job to &#x60;CLOSED&#x60;. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**RestoreScheduledDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud will restore this snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**SnapshotFinishedDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud completed writing this snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**SnapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the snapshot to restore. | 
**SnapshotUrl** | Pointer to **string** | Internet address from which you can download the compressed snapshot files. The resource returns this parameter when  &#x60;\&quot;deliveryType\&quot; : \&quot;DOWNLOAD\&quot;&#x60;. | [optional] [readonly] 
**Status** | Pointer to **string** | Phase of the restore workflow for this job at the time this resource made this request. | [optional] [readonly] 
**TargetDeploymentItemName** | **string** | Human-readable label that identifies the cluster on the target project to which you want to restore the snapshot. You can restore the snapshot to a cluster tier *M2* or greater. | 
**TargetProjectId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that contains the cluster to which you want to restore the snapshot. | [optional] 

## Methods

### NewApiAtlasTenantRestoreView

`func NewApiAtlasTenantRestoreView(links []Link, snapshotId string, targetDeploymentItemName string, ) *ApiAtlasTenantRestoreView`

NewApiAtlasTenantRestoreView instantiates a new ApiAtlasTenantRestoreView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasTenantRestoreViewWithDefaults

`func NewApiAtlasTenantRestoreViewWithDefaults() *ApiAtlasTenantRestoreView`

NewApiAtlasTenantRestoreViewWithDefaults instantiates a new ApiAtlasTenantRestoreView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ApiAtlasTenantRestoreView) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ApiAtlasTenantRestoreView) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ApiAtlasTenantRestoreView) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ApiAtlasTenantRestoreView) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetDeliveryType

`func (o *ApiAtlasTenantRestoreView) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *ApiAtlasTenantRestoreView) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *ApiAtlasTenantRestoreView) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *ApiAtlasTenantRestoreView) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ApiAtlasTenantRestoreView) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ApiAtlasTenantRestoreView) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ApiAtlasTenantRestoreView) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ApiAtlasTenantRestoreView) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasTenantRestoreView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasTenantRestoreView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasTenantRestoreView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasTenantRestoreView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasTenantRestoreView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasTenantRestoreView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasTenantRestoreView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetProjectId

`func (o *ApiAtlasTenantRestoreView) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ApiAtlasTenantRestoreView) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ApiAtlasTenantRestoreView) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ApiAtlasTenantRestoreView) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRestoreFinishedDate

`func (o *ApiAtlasTenantRestoreView) GetRestoreFinishedDate() time.Time`

GetRestoreFinishedDate returns the RestoreFinishedDate field if non-nil, zero value otherwise.

### GetRestoreFinishedDateOk

`func (o *ApiAtlasTenantRestoreView) GetRestoreFinishedDateOk() (*time.Time, bool)`

GetRestoreFinishedDateOk returns a tuple with the RestoreFinishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreFinishedDate

`func (o *ApiAtlasTenantRestoreView) SetRestoreFinishedDate(v time.Time)`

SetRestoreFinishedDate sets RestoreFinishedDate field to given value.

### HasRestoreFinishedDate

`func (o *ApiAtlasTenantRestoreView) HasRestoreFinishedDate() bool`

HasRestoreFinishedDate returns a boolean if a field has been set.

### GetRestoreScheduledDate

`func (o *ApiAtlasTenantRestoreView) GetRestoreScheduledDate() time.Time`

GetRestoreScheduledDate returns the RestoreScheduledDate field if non-nil, zero value otherwise.

### GetRestoreScheduledDateOk

`func (o *ApiAtlasTenantRestoreView) GetRestoreScheduledDateOk() (*time.Time, bool)`

GetRestoreScheduledDateOk returns a tuple with the RestoreScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreScheduledDate

`func (o *ApiAtlasTenantRestoreView) SetRestoreScheduledDate(v time.Time)`

SetRestoreScheduledDate sets RestoreScheduledDate field to given value.

### HasRestoreScheduledDate

`func (o *ApiAtlasTenantRestoreView) HasRestoreScheduledDate() bool`

HasRestoreScheduledDate returns a boolean if a field has been set.

### GetSnapshotFinishedDate

`func (o *ApiAtlasTenantRestoreView) GetSnapshotFinishedDate() time.Time`

GetSnapshotFinishedDate returns the SnapshotFinishedDate field if non-nil, zero value otherwise.

### GetSnapshotFinishedDateOk

`func (o *ApiAtlasTenantRestoreView) GetSnapshotFinishedDateOk() (*time.Time, bool)`

GetSnapshotFinishedDateOk returns a tuple with the SnapshotFinishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotFinishedDate

`func (o *ApiAtlasTenantRestoreView) SetSnapshotFinishedDate(v time.Time)`

SetSnapshotFinishedDate sets SnapshotFinishedDate field to given value.

### HasSnapshotFinishedDate

`func (o *ApiAtlasTenantRestoreView) HasSnapshotFinishedDate() bool`

HasSnapshotFinishedDate returns a boolean if a field has been set.

### GetSnapshotId

`func (o *ApiAtlasTenantRestoreView) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ApiAtlasTenantRestoreView) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ApiAtlasTenantRestoreView) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetSnapshotUrl

`func (o *ApiAtlasTenantRestoreView) GetSnapshotUrl() string`

GetSnapshotUrl returns the SnapshotUrl field if non-nil, zero value otherwise.

### GetSnapshotUrlOk

`func (o *ApiAtlasTenantRestoreView) GetSnapshotUrlOk() (*string, bool)`

GetSnapshotUrlOk returns a tuple with the SnapshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotUrl

`func (o *ApiAtlasTenantRestoreView) SetSnapshotUrl(v string)`

SetSnapshotUrl sets SnapshotUrl field to given value.

### HasSnapshotUrl

`func (o *ApiAtlasTenantRestoreView) HasSnapshotUrl() bool`

HasSnapshotUrl returns a boolean if a field has been set.

### GetStatus

`func (o *ApiAtlasTenantRestoreView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAtlasTenantRestoreView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAtlasTenantRestoreView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAtlasTenantRestoreView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetDeploymentItemName

`func (o *ApiAtlasTenantRestoreView) GetTargetDeploymentItemName() string`

GetTargetDeploymentItemName returns the TargetDeploymentItemName field if non-nil, zero value otherwise.

### GetTargetDeploymentItemNameOk

`func (o *ApiAtlasTenantRestoreView) GetTargetDeploymentItemNameOk() (*string, bool)`

GetTargetDeploymentItemNameOk returns a tuple with the TargetDeploymentItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDeploymentItemName

`func (o *ApiAtlasTenantRestoreView) SetTargetDeploymentItemName(v string)`

SetTargetDeploymentItemName sets TargetDeploymentItemName field to given value.


### GetTargetProjectId

`func (o *ApiAtlasTenantRestoreView) GetTargetProjectId() string`

GetTargetProjectId returns the TargetProjectId field if non-nil, zero value otherwise.

### GetTargetProjectIdOk

`func (o *ApiAtlasTenantRestoreView) GetTargetProjectIdOk() (*string, bool)`

GetTargetProjectIdOk returns a tuple with the TargetProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProjectId

`func (o *ApiAtlasTenantRestoreView) SetTargetProjectId(v string)`

SetTargetProjectId sets TargetProjectId field to given value.

### HasTargetProjectId

`func (o *ApiAtlasTenantRestoreView) HasTargetProjectId() bool`

HasTargetProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


