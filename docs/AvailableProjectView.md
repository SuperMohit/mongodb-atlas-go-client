# AvailableProjectView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | [**[]AvailableDeploymentView**](AvailableDeploymentView.md) | List of clusters that can be migrated to MongoDB Cloud. | 
**MigrationHosts** | **[]string** |  | 
**Name** | **string** | Human-readable label that identifies this project. | [readonly] 
**ProjectId** | **string** | Unique 24-hexadecimal digit string that identifies the project to be migrated. | [readonly] 

## Methods

### NewAvailableProjectView

`func NewAvailableProjectView(deployments []AvailableDeploymentView, migrationHosts []string, name string, projectId string, ) *AvailableProjectView`

NewAvailableProjectView instantiates a new AvailableProjectView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableProjectViewWithDefaults

`func NewAvailableProjectViewWithDefaults() *AvailableProjectView`

NewAvailableProjectViewWithDefaults instantiates a new AvailableProjectView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *AvailableProjectView) GetDeployments() []AvailableDeploymentView`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *AvailableProjectView) GetDeploymentsOk() (*[]AvailableDeploymentView, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *AvailableProjectView) SetDeployments(v []AvailableDeploymentView)`

SetDeployments sets Deployments field to given value.


### GetMigrationHosts

`func (o *AvailableProjectView) GetMigrationHosts() []string`

GetMigrationHosts returns the MigrationHosts field if non-nil, zero value otherwise.

### GetMigrationHostsOk

`func (o *AvailableProjectView) GetMigrationHostsOk() (*[]string, bool)`

GetMigrationHostsOk returns a tuple with the MigrationHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationHosts

`func (o *AvailableProjectView) SetMigrationHosts(v []string)`

SetMigrationHosts sets MigrationHosts field to given value.


### GetName

`func (o *AvailableProjectView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailableProjectView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailableProjectView) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *AvailableProjectView) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AvailableProjectView) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AvailableProjectView) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


