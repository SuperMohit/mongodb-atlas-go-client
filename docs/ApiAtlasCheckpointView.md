# ApiAtlasCheckpointView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster that contains the checkpoint. | [optional] [readonly] 
**Completed** | Pointer to **time.Time** | Date and time when the checkpoint completed and the balancer restarted. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns the checkpoints. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies checkpoint. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Parts** | Pointer to [**[]ApiCheckpointPartView**](ApiCheckpointPartView.md) | Metadata that describes the complete snapshot.  - For a replica set, this array contains a single document. - For a sharded cluster, this array contains one document for each shard plus one document for the config host. | [optional] [readonly] 
**Restorable** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud can use the checkpoint for a restore. | [optional] [readonly] 
**Started** | Pointer to **time.Time** | Date and time when the balancer stopped and began the checkpoint. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** | Date and time to which the checkpoint restores. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewApiAtlasCheckpointView

`func NewApiAtlasCheckpointView(links []Link, ) *ApiAtlasCheckpointView`

NewApiAtlasCheckpointView instantiates a new ApiAtlasCheckpointView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCheckpointViewWithDefaults

`func NewApiAtlasCheckpointViewWithDefaults() *ApiAtlasCheckpointView`

NewApiAtlasCheckpointViewWithDefaults instantiates a new ApiAtlasCheckpointView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ApiAtlasCheckpointView) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ApiAtlasCheckpointView) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ApiAtlasCheckpointView) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ApiAtlasCheckpointView) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCompleted

`func (o *ApiAtlasCheckpointView) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *ApiAtlasCheckpointView) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *ApiAtlasCheckpointView) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *ApiAtlasCheckpointView) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiAtlasCheckpointView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasCheckpointView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasCheckpointView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAtlasCheckpointView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasCheckpointView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasCheckpointView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasCheckpointView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasCheckpointView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasCheckpointView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasCheckpointView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasCheckpointView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetParts

`func (o *ApiAtlasCheckpointView) GetParts() []ApiCheckpointPartView`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *ApiAtlasCheckpointView) GetPartsOk() (*[]ApiCheckpointPartView, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *ApiAtlasCheckpointView) SetParts(v []ApiCheckpointPartView)`

SetParts sets Parts field to given value.

### HasParts

`func (o *ApiAtlasCheckpointView) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetRestorable

`func (o *ApiAtlasCheckpointView) GetRestorable() bool`

GetRestorable returns the Restorable field if non-nil, zero value otherwise.

### GetRestorableOk

`func (o *ApiAtlasCheckpointView) GetRestorableOk() (*bool, bool)`

GetRestorableOk returns a tuple with the Restorable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorable

`func (o *ApiAtlasCheckpointView) SetRestorable(v bool)`

SetRestorable sets Restorable field to given value.

### HasRestorable

`func (o *ApiAtlasCheckpointView) HasRestorable() bool`

HasRestorable returns a boolean if a field has been set.

### GetStarted

`func (o *ApiAtlasCheckpointView) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *ApiAtlasCheckpointView) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *ApiAtlasCheckpointView) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *ApiAtlasCheckpointView) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetTimestamp

`func (o *ApiAtlasCheckpointView) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiAtlasCheckpointView) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiAtlasCheckpointView) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiAtlasCheckpointView) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


