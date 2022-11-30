# ApiAtlasServerlessClusterDescriptionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStrings** | Pointer to [**ApiAtlasClusterDescriptionConnectionStringsView**](ApiAtlasClusterDescriptionConnectionStringsView.md) |  | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this cluster. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**MongoDBVersion** | Pointer to **string** | Version of MongoDB that the cluster runs. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the cluster. | [optional] 
**ProviderSettings** | Pointer to [**ApiAtlasProviderSettingsViewManual**](ApiAtlasProviderSettingsViewManual.md) |  | [optional] 
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of the cluster. | [optional] [readonly] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the cluster. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the cluster. If set to &#x60;false&#x60;, MongoDB Cloud will delete the cluster. | [optional] [default to false]

## Methods

### NewApiAtlasServerlessClusterDescriptionView

`func NewApiAtlasServerlessClusterDescriptionView(links []Link, ) *ApiAtlasServerlessClusterDescriptionView`

NewApiAtlasServerlessClusterDescriptionView instantiates a new ApiAtlasServerlessClusterDescriptionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasServerlessClusterDescriptionViewWithDefaults

`func NewApiAtlasServerlessClusterDescriptionViewWithDefaults() *ApiAtlasServerlessClusterDescriptionView`

NewApiAtlasServerlessClusterDescriptionViewWithDefaults instantiates a new ApiAtlasServerlessClusterDescriptionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStrings

`func (o *ApiAtlasServerlessClusterDescriptionView) GetConnectionStrings() ApiAtlasClusterDescriptionConnectionStringsView`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *ApiAtlasServerlessClusterDescriptionView) GetConnectionStringsOk() (*ApiAtlasClusterDescriptionConnectionStringsView, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *ApiAtlasServerlessClusterDescriptionView) SetConnectionStrings(v ApiAtlasClusterDescriptionConnectionStringsView)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *ApiAtlasServerlessClusterDescriptionView) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.

### GetCreateDate

`func (o *ApiAtlasServerlessClusterDescriptionView) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ApiAtlasServerlessClusterDescriptionView) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ApiAtlasServerlessClusterDescriptionView) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ApiAtlasServerlessClusterDescriptionView) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiAtlasServerlessClusterDescriptionView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasServerlessClusterDescriptionView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasServerlessClusterDescriptionView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAtlasServerlessClusterDescriptionView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasServerlessClusterDescriptionView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasServerlessClusterDescriptionView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasServerlessClusterDescriptionView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasServerlessClusterDescriptionView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasServerlessClusterDescriptionView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasServerlessClusterDescriptionView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasServerlessClusterDescriptionView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetMongoDBVersion

`func (o *ApiAtlasServerlessClusterDescriptionView) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *ApiAtlasServerlessClusterDescriptionView) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *ApiAtlasServerlessClusterDescriptionView) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *ApiAtlasServerlessClusterDescriptionView) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.

### GetName

`func (o *ApiAtlasServerlessClusterDescriptionView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasServerlessClusterDescriptionView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasServerlessClusterDescriptionView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAtlasServerlessClusterDescriptionView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviderSettings

`func (o *ApiAtlasServerlessClusterDescriptionView) GetProviderSettings() ApiAtlasProviderSettingsViewManual`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ApiAtlasServerlessClusterDescriptionView) GetProviderSettingsOk() (*ApiAtlasProviderSettingsViewManual, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ApiAtlasServerlessClusterDescriptionView) SetProviderSettings(v ApiAtlasProviderSettingsViewManual)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *ApiAtlasServerlessClusterDescriptionView) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetStateName

`func (o *ApiAtlasServerlessClusterDescriptionView) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *ApiAtlasServerlessClusterDescriptionView) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *ApiAtlasServerlessClusterDescriptionView) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *ApiAtlasServerlessClusterDescriptionView) HasStateName() bool`

HasStateName returns a boolean if a field has been set.

### GetTerminationProtectionEnabled

`func (o *ApiAtlasServerlessClusterDescriptionView) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *ApiAtlasServerlessClusterDescriptionView) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *ApiAtlasServerlessClusterDescriptionView) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *ApiAtlasServerlessClusterDescriptionView) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


