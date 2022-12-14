# ApiAtlasGroupView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterCount** | **int64** | Quantity of MongoDB Cloud clusters deployed in this project. | [readonly] 
**Created** | **time.Time** | Date and time when MongoDB Cloud created this project. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [readonly] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud project. | [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Name** | **string** | Human-readable label that identifies the project included in the MongoDB Cloud organization. | 
**OrgId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud organization to which the project belongs. | 
**WithDefaultAlertsSettings** | Pointer to **bool** | Flag that indicates whether to create the project with default alert settings. | [optional] 

## Methods

### NewApiAtlasGroupView

`func NewApiAtlasGroupView(clusterCount int64, created time.Time, id string, links []Link, name string, orgId string, ) *ApiAtlasGroupView`

NewApiAtlasGroupView instantiates a new ApiAtlasGroupView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasGroupViewWithDefaults

`func NewApiAtlasGroupViewWithDefaults() *ApiAtlasGroupView`

NewApiAtlasGroupViewWithDefaults instantiates a new ApiAtlasGroupView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterCount

`func (o *ApiAtlasGroupView) GetClusterCount() int64`

GetClusterCount returns the ClusterCount field if non-nil, zero value otherwise.

### GetClusterCountOk

`func (o *ApiAtlasGroupView) GetClusterCountOk() (*int64, bool)`

GetClusterCountOk returns a tuple with the ClusterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCount

`func (o *ApiAtlasGroupView) SetClusterCount(v int64)`

SetClusterCount sets ClusterCount field to given value.


### GetCreated

`func (o *ApiAtlasGroupView) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiAtlasGroupView) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiAtlasGroupView) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *ApiAtlasGroupView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasGroupView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasGroupView) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *ApiAtlasGroupView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasGroupView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasGroupView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetName

`func (o *ApiAtlasGroupView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasGroupView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasGroupView) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *ApiAtlasGroupView) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ApiAtlasGroupView) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ApiAtlasGroupView) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetWithDefaultAlertsSettings

`func (o *ApiAtlasGroupView) GetWithDefaultAlertsSettings() bool`

GetWithDefaultAlertsSettings returns the WithDefaultAlertsSettings field if non-nil, zero value otherwise.

### GetWithDefaultAlertsSettingsOk

`func (o *ApiAtlasGroupView) GetWithDefaultAlertsSettingsOk() (*bool, bool)`

GetWithDefaultAlertsSettingsOk returns a tuple with the WithDefaultAlertsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithDefaultAlertsSettings

`func (o *ApiAtlasGroupView) SetWithDefaultAlertsSettings(v bool)`

SetWithDefaultAlertsSettings sets WithDefaultAlertsSettings field to given value.

### HasWithDefaultAlertsSettings

`func (o *ApiAtlasGroupView) HasWithDefaultAlertsSettings() bool`

HasWithDefaultAlertsSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


