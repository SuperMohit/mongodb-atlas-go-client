# Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** | Label that identifies the destination cluster. | 
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the destination project. | 
**HostnameSchemaType** | **string** | The network type to use between the migration host and the target cluster. | [default to "PUBLIC"]

## Methods

### NewDestination

`func NewDestination(clusterName string, groupId string, hostnameSchemaType string, ) *Destination`

NewDestination instantiates a new Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationWithDefaults

`func NewDestinationWithDefaults() *Destination`

NewDestinationWithDefaults instantiates a new Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *Destination) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Destination) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Destination) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetGroupId

`func (o *Destination) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Destination) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Destination) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetHostnameSchemaType

`func (o *Destination) GetHostnameSchemaType() string`

GetHostnameSchemaType returns the HostnameSchemaType field if non-nil, zero value otherwise.

### GetHostnameSchemaTypeOk

`func (o *Destination) GetHostnameSchemaTypeOk() (*string, bool)`

GetHostnameSchemaTypeOk returns a tuple with the HostnameSchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameSchemaType

`func (o *Destination) SetHostnameSchemaType(v string)`

SetHostnameSchemaType sets HostnameSchemaType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


