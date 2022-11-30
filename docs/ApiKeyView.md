# ApiKeyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessList** | [**[]AccessListItemView**](AccessListItemView.md) | List of network addresses granted access to this API using this API key. | 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key. | [readonly] 
**PublicKey** | **string** | Public API key value set for the specified organization API key.  | [readonly] 
**Roles** | [**[]ApiRoleAssignmentView**](ApiRoleAssignmentView.md) | List that contains roles that the API key needs to have. All roles you provide must be valid for the specified project or organization. Each request must include a minimum of one valid role. The resource returns all project and organization roles assigned to the Cloud user. | [readonly] 

## Methods

### NewApiKeyView

`func NewApiKeyView(accessList []AccessListItemView, id string, publicKey string, roles []ApiRoleAssignmentView, ) *ApiKeyView`

NewApiKeyView instantiates a new ApiKeyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyViewWithDefaults

`func NewApiKeyViewWithDefaults() *ApiKeyView`

NewApiKeyViewWithDefaults instantiates a new ApiKeyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessList

`func (o *ApiKeyView) GetAccessList() []AccessListItemView`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *ApiKeyView) GetAccessListOk() (*[]AccessListItemView, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *ApiKeyView) SetAccessList(v []AccessListItemView)`

SetAccessList sets AccessList field to given value.


### GetId

`func (o *ApiKeyView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyView) SetId(v string)`

SetId sets Id field to given value.


### GetPublicKey

`func (o *ApiKeyView) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ApiKeyView) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ApiKeyView) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetRoles

`func (o *ApiKeyView) GetRoles() []ApiRoleAssignmentView`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiKeyView) GetRolesOk() (*[]ApiRoleAssignmentView, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiKeyView) SetRoles(v []ApiRoleAssignmentView)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


