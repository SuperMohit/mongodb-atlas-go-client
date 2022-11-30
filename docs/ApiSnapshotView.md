# ApiSnapshotView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster with the snapshots you want to return. | [optional] [readonly] 
**Complete** | Pointer to **bool** | Flag that indicates whether the snapshot exists. This flag returns &#x60;false&#x60; while MongoDB Cloud creates the snapshot. | [optional] [readonly] 
**Created** | Pointer to [**ApiBSONTimestampView**](ApiBSONTimestampView.md) |  | [optional] 
**DoNotDelete** | Pointer to **bool** | Flag that indicates whether someone can delete this snapshot. You can&#39;t set &#x60;\&quot;doNotDelete\&quot; : true&#x60; and set a timestamp for **expires** in the same request. | [optional] 
**Expires** | Pointer to **time.Time** | Date and time when MongoDB Cloud deletes the snapshot. If &#x60;\&quot;doNotDelete\&quot; : true&#x60;, MongoDB Cloud removes any value set for this parameter. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns the snapshots. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the snapshot. | [optional] [readonly] 
**LastOplogAppliedTimestamp** | Pointer to [**ApiBSONTimestampView**](ApiBSONTimestampView.md) |  | [optional] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Parts** | Pointer to [**[]ApiSnapshotPartView**](ApiSnapshotPartView.md) | Metadata that describes the complete snapshot.  - For a replica set, this array contains a single document. - For a sharded cluster, this array contains one document for each shard plus one document for the config host. | [optional] [readonly] 

## Methods

### NewApiSnapshotView

`func NewApiSnapshotView(links []Link, ) *ApiSnapshotView`

NewApiSnapshotView instantiates a new ApiSnapshotView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSnapshotViewWithDefaults

`func NewApiSnapshotViewWithDefaults() *ApiSnapshotView`

NewApiSnapshotViewWithDefaults instantiates a new ApiSnapshotView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ApiSnapshotView) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ApiSnapshotView) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ApiSnapshotView) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ApiSnapshotView) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetComplete

`func (o *ApiSnapshotView) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *ApiSnapshotView) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *ApiSnapshotView) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *ApiSnapshotView) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetCreated

`func (o *ApiSnapshotView) GetCreated() ApiBSONTimestampView`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiSnapshotView) GetCreatedOk() (*ApiBSONTimestampView, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiSnapshotView) SetCreated(v ApiBSONTimestampView)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiSnapshotView) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDoNotDelete

`func (o *ApiSnapshotView) GetDoNotDelete() bool`

GetDoNotDelete returns the DoNotDelete field if non-nil, zero value otherwise.

### GetDoNotDeleteOk

`func (o *ApiSnapshotView) GetDoNotDeleteOk() (*bool, bool)`

GetDoNotDeleteOk returns a tuple with the DoNotDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDelete

`func (o *ApiSnapshotView) SetDoNotDelete(v bool)`

SetDoNotDelete sets DoNotDelete field to given value.

### HasDoNotDelete

`func (o *ApiSnapshotView) HasDoNotDelete() bool`

HasDoNotDelete returns a boolean if a field has been set.

### GetExpires

`func (o *ApiSnapshotView) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ApiSnapshotView) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ApiSnapshotView) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ApiSnapshotView) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiSnapshotView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiSnapshotView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiSnapshotView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiSnapshotView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ApiSnapshotView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiSnapshotView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiSnapshotView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiSnapshotView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastOplogAppliedTimestamp

`func (o *ApiSnapshotView) GetLastOplogAppliedTimestamp() ApiBSONTimestampView`

GetLastOplogAppliedTimestamp returns the LastOplogAppliedTimestamp field if non-nil, zero value otherwise.

### GetLastOplogAppliedTimestampOk

`func (o *ApiSnapshotView) GetLastOplogAppliedTimestampOk() (*ApiBSONTimestampView, bool)`

GetLastOplogAppliedTimestampOk returns a tuple with the LastOplogAppliedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOplogAppliedTimestamp

`func (o *ApiSnapshotView) SetLastOplogAppliedTimestamp(v ApiBSONTimestampView)`

SetLastOplogAppliedTimestamp sets LastOplogAppliedTimestamp field to given value.

### HasLastOplogAppliedTimestamp

`func (o *ApiSnapshotView) HasLastOplogAppliedTimestamp() bool`

HasLastOplogAppliedTimestamp returns a boolean if a field has been set.

### GetLinks

`func (o *ApiSnapshotView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiSnapshotView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiSnapshotView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetParts

`func (o *ApiSnapshotView) GetParts() []ApiSnapshotPartView`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *ApiSnapshotView) GetPartsOk() (*[]ApiSnapshotPartView, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *ApiSnapshotView) SetParts(v []ApiSnapshotPartView)`

SetParts sets Parts field to given value.

### HasParts

`func (o *ApiSnapshotView) HasParts() bool`

HasParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


