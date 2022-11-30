# ApiAtlasServerlessClusterDescriptionViewManual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStrings** | Pointer to [**ServerlessInstanceConnectionStrings**](ServerlessInstanceConnectionStrings.md) |  | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this serverless instance. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the serverless instance. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MongoDBVersion** | Pointer to **string** | Version of MongoDB that the serverless instance runs. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the serverless instance. | [optional] 
**ProviderSettings** | Pointer to [**ServerlessInstanceProviderSettings**](ServerlessInstanceProviderSettings.md) |  | [optional] 
**ServerlessBackupOptions** | Pointer to [**ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions**](ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions.md) |  | [optional] 
**StateName** | Pointer to **string** | Human-readable label that indicates the current operating condition of the serverless instance. | [optional] [readonly] 
**TerminationProtectionEnabled** | Pointer to **bool** | Flag that indicates whether termination protection is enabled on the serverless instance. If set to &#x60;true&#x60;, MongoDB Cloud won&#39;t delete the serverless instance. If set to &#x60;false&#x60;, MongoDB cloud will delete the serverless instance.\&quot; | [optional] [default to false]

## Methods

### NewApiAtlasServerlessClusterDescriptionViewManual

`func NewApiAtlasServerlessClusterDescriptionViewManual() *ApiAtlasServerlessClusterDescriptionViewManual`

NewApiAtlasServerlessClusterDescriptionViewManual instantiates a new ApiAtlasServerlessClusterDescriptionViewManual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasServerlessClusterDescriptionViewManualWithDefaults

`func NewApiAtlasServerlessClusterDescriptionViewManualWithDefaults() *ApiAtlasServerlessClusterDescriptionViewManual`

NewApiAtlasServerlessClusterDescriptionViewManualWithDefaults instantiates a new ApiAtlasServerlessClusterDescriptionViewManual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStrings

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetConnectionStrings() ServerlessInstanceConnectionStrings`

GetConnectionStrings returns the ConnectionStrings field if non-nil, zero value otherwise.

### GetConnectionStringsOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetConnectionStringsOk() (*ServerlessInstanceConnectionStrings, bool)`

GetConnectionStringsOk returns a tuple with the ConnectionStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStrings

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) SetConnectionStrings(v ServerlessInstanceConnectionStrings)`

SetConnectionStrings sets ConnectionStrings field to given value.

### HasConnectionStrings

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) HasConnectionStrings() bool`

HasConnectionStrings returns a boolean if a field has been set.

### GetCreateDate

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMongoDBVersion

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.

### GetName

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviderSettings

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetProviderSettings() ServerlessInstanceProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetProviderSettingsOk() (*ServerlessInstanceProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) SetProviderSettings(v ServerlessInstanceProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetServerlessBackupOptions

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetServerlessBackupOptions() ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions`

GetServerlessBackupOptions returns the ServerlessBackupOptions field if non-nil, zero value otherwise.

### GetServerlessBackupOptionsOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetServerlessBackupOptionsOk() (*ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions, bool)`

GetServerlessBackupOptionsOk returns a tuple with the ServerlessBackupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerlessBackupOptions

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) SetServerlessBackupOptions(v ApiAtlasServerlessClusterDescriptionViewManualServerlessBackupOptions)`

SetServerlessBackupOptions sets ServerlessBackupOptions field to given value.

### HasServerlessBackupOptions

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) HasServerlessBackupOptions() bool`

HasServerlessBackupOptions returns a boolean if a field has been set.

### GetStateName

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) HasStateName() bool`

HasStateName returns a boolean if a field has been set.

### GetTerminationProtectionEnabled

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetTerminationProtectionEnabled() bool`

GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field if non-nil, zero value otherwise.

### GetTerminationProtectionEnabledOk

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) GetTerminationProtectionEnabledOk() (*bool, bool)`

GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationProtectionEnabled

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) SetTerminationProtectionEnabled(v bool)`

SetTerminationProtectionEnabled sets TerminationProtectionEnabled field to given value.

### HasTerminationProtectionEnabled

`func (o *ApiAtlasServerlessClusterDescriptionViewManual) HasTerminationProtectionEnabled() bool`

HasTerminationProtectionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


