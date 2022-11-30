# ApiApiUserView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** | Purpose or explanation provided when someone created this organization API key. | [optional] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this organization API key assigned to this project. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**PrivateKey** | Pointer to **string** | Redacted private key returned for this organization API key. This key displays unredacted when first created. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | Public API key value set for the specified organization API key. | [optional] [readonly] 
**Roles** | Pointer to [**[]ApiRoleAssignmentView**](ApiRoleAssignmentView.md) | List that contains the roles that the API key needs to have. All roles you provide must be valid for the specified project or organization. Each request must include a minimum of one valid role. The resource returns all project and organization roles assigned to the API key. | [optional] 

## Methods

### NewApiApiUserView

`func NewApiApiUserView(links []Link, ) *ApiApiUserView`

NewApiApiUserView instantiates a new ApiApiUserView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiApiUserViewWithDefaults

`func NewApiApiUserViewWithDefaults() *ApiApiUserView`

NewApiApiUserViewWithDefaults instantiates a new ApiApiUserView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *ApiApiUserView) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ApiApiUserView) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ApiApiUserView) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *ApiApiUserView) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetId

`func (o *ApiApiUserView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiApiUserView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiApiUserView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiApiUserView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiApiUserView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiApiUserView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiApiUserView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetPrivateKey

`func (o *ApiApiUserView) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ApiApiUserView) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ApiApiUserView) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *ApiApiUserView) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *ApiApiUserView) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ApiApiUserView) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ApiApiUserView) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *ApiApiUserView) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRoles

`func (o *ApiApiUserView) GetRoles() []ApiRoleAssignmentView`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiApiUserView) GetRolesOk() (*[]ApiRoleAssignmentView, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiApiUserView) SetRoles(v []ApiRoleAssignmentView)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiApiUserView) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


