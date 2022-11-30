# LiveMigrationRequestView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the migration request. | [optional] [readonly] 
**Destination** | [**Destination**](Destination.md) |  | 
**DropEnabled** | **bool** | Flag that indicates whether the migration process drops all collections from the destination cluster before the migration starts. | 
**MigrationHosts** | **[]string** | List of migration hosts used for this migration. | 
**Source** | [**Source**](Source.md) |  | 

## Methods

### NewLiveMigrationRequestView

`func NewLiveMigrationRequestView(destination Destination, dropEnabled bool, migrationHosts []string, source Source, ) *LiveMigrationRequestView`

NewLiveMigrationRequestView instantiates a new LiveMigrationRequestView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveMigrationRequestViewWithDefaults

`func NewLiveMigrationRequestViewWithDefaults() *LiveMigrationRequestView`

NewLiveMigrationRequestViewWithDefaults instantiates a new LiveMigrationRequestView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LiveMigrationRequestView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveMigrationRequestView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveMigrationRequestView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveMigrationRequestView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDestination

`func (o *LiveMigrationRequestView) GetDestination() Destination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *LiveMigrationRequestView) GetDestinationOk() (*Destination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *LiveMigrationRequestView) SetDestination(v Destination)`

SetDestination sets Destination field to given value.


### GetDropEnabled

`func (o *LiveMigrationRequestView) GetDropEnabled() bool`

GetDropEnabled returns the DropEnabled field if non-nil, zero value otherwise.

### GetDropEnabledOk

`func (o *LiveMigrationRequestView) GetDropEnabledOk() (*bool, bool)`

GetDropEnabledOk returns a tuple with the DropEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropEnabled

`func (o *LiveMigrationRequestView) SetDropEnabled(v bool)`

SetDropEnabled sets DropEnabled field to given value.


### GetMigrationHosts

`func (o *LiveMigrationRequestView) GetMigrationHosts() []string`

GetMigrationHosts returns the MigrationHosts field if non-nil, zero value otherwise.

### GetMigrationHostsOk

`func (o *LiveMigrationRequestView) GetMigrationHostsOk() (*[]string, bool)`

GetMigrationHostsOk returns a tuple with the MigrationHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationHosts

`func (o *LiveMigrationRequestView) SetMigrationHosts(v []string)`

SetMigrationHosts sets MigrationHosts field to given value.


### GetSource

`func (o *LiveMigrationRequestView) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LiveMigrationRequestView) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LiveMigrationRequestView) SetSource(v Source)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


