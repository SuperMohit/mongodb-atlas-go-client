# ApiEventView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertConfigId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the alert configuration associated with the **alertId**. | [optional] [readonly] 
**AlertId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the alert associated with the event. | [optional] [readonly] 
**ApiKeyId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the API Key that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **userId** parameter. | [optional] [readonly] 
**Collection** | Pointer to **string** | Human-readable label of the collection on which the event occurred. The resource returns this parameter when the **eventTypeName** includes &#x60;DATA_EXPLORER&#x60;. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when this event occurred. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**CurrentValue** | Pointer to [**ApiMetricValueView**](ApiMetricValueView.md) |  | [optional] 
**Database** | Pointer to **string** | Human-readable label of the database on which this incident occurred. The resource returns this parameter when &#x60;\&quot;eventTypeName\&quot; : \&quot;DATA_EXPLORER\&quot;&#x60; or &#x60;\&quot;eventTypeName\&quot; : \&quot;DATA_EXPLORER_CRUD\&quot;&#x60;. | [optional] [readonly] 
**EndpointId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the endpoint associated with this event. | [optional] [readonly] 
**EventTypeName** | Pointer to **string** | Category of incident recorded at this moment in time. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the event. | [optional] [readonly] 
**InvoiceId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies of the invoice associated with the event. | [optional] [readonly] 
**IsGlobalAdmin** | Pointer to **bool** | Flag that indicates whether a MongoDB employee triggered the specified event. | [optional] [readonly] [default to false]
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**MetricName** | Pointer to **string** | Human-readable label of the metric associated with the **alertId**. | [optional] [readonly] 
**OpType** | Pointer to **string** | Action that the database attempted to execute when the event triggered. The response returns this parameter when &#x60;eventTypeName\&quot; : \&quot;DATA_EXPLORER\&quot;&#x60;. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies for the organization to which these events apply. | [optional] [readonly] 
**PaymentId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the invoice payment associated with this event. | [optional] [readonly] 
**Port** | Pointer to **int32** | IANA port on which the MongoDB process listens for requests. | [optional] [readonly] 
**ProviderEndpointId** | Pointer to **string** | Unique identification string that the cloud provider uses to identify the private endpoint. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | Public part of the API key that triggered the event. If this resource returns this parameter, it doesn&#39;t return the **username** parameter. | [optional] [readonly] 
**Raw** | Pointer to **string** | Additional meta information captured about this event. The response returns this parameter as a JSON object when the query parameter &#x60;includeRaw&#x3D;true&#x60;. The values in the raw document may change. Don&#39;t rely on raw values for formal monitoring. | [optional] [readonly] 
**RemoteAddress** | Pointer to **string** | IPv4 or IPv6 address from which the user triggered this event. | [optional] [readonly] 
**ReplicaSetName** | Pointer to **string** | Human-readable label of the replica set associated with the event. | [optional] [readonly] 
**ShardName** | Pointer to **string** | Human-readable label of the shard associated with the event. | [optional] [readonly] 
**TargetPublicKey** | Pointer to **string** | Public part of the API key that this event targets. | [optional] [readonly] 
**TargetUsername** | Pointer to **string** | Email address for the console user that this event targets. The resource returns this parameter when &#x60;\&quot;eventTypeName\&quot; : \&quot;USER\&quot;&#x60;. | [optional] [readonly] 
**TeamId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization team associated with this event. | [optional] [readonly] 
**UserId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn&#39;t return the **apiKeyId** parameter. | [optional] [readonly] 
**Username** | Pointer to **string** | Email address for the user who triggered this event. If this resource returns this parameter, it doesn&#39;t return the **publicApiKey** parameter. | [optional] [readonly] 
**WhitelistEntry** | Pointer to **string** | Entry in the list of source host addresses that the API key accepts and this event targets. | [optional] [readonly] 

## Methods

### NewApiEventView

`func NewApiEventView(links []Link, ) *ApiEventView`

NewApiEventView instantiates a new ApiEventView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEventViewWithDefaults

`func NewApiEventViewWithDefaults() *ApiEventView`

NewApiEventViewWithDefaults instantiates a new ApiEventView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertConfigId

`func (o *ApiEventView) GetAlertConfigId() string`

GetAlertConfigId returns the AlertConfigId field if non-nil, zero value otherwise.

### GetAlertConfigIdOk

`func (o *ApiEventView) GetAlertConfigIdOk() (*string, bool)`

GetAlertConfigIdOk returns a tuple with the AlertConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigId

`func (o *ApiEventView) SetAlertConfigId(v string)`

SetAlertConfigId sets AlertConfigId field to given value.

### HasAlertConfigId

`func (o *ApiEventView) HasAlertConfigId() bool`

HasAlertConfigId returns a boolean if a field has been set.

### GetAlertId

`func (o *ApiEventView) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *ApiEventView) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *ApiEventView) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *ApiEventView) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetApiKeyId

`func (o *ApiEventView) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *ApiEventView) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *ApiEventView) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *ApiEventView) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### GetCollection

`func (o *ApiEventView) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ApiEventView) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ApiEventView) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *ApiEventView) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetCreated

`func (o *ApiEventView) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiEventView) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiEventView) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiEventView) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCurrentValue

`func (o *ApiEventView) GetCurrentValue() ApiMetricValueView`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *ApiEventView) GetCurrentValueOk() (*ApiMetricValueView, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *ApiEventView) SetCurrentValue(v ApiMetricValueView)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *ApiEventView) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetDatabase

`func (o *ApiEventView) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ApiEventView) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ApiEventView) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *ApiEventView) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetEndpointId

`func (o *ApiEventView) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ApiEventView) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ApiEventView) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *ApiEventView) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEventTypeName

`func (o *ApiEventView) GetEventTypeName() string`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *ApiEventView) GetEventTypeNameOk() (*string, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *ApiEventView) SetEventTypeName(v string)`

SetEventTypeName sets EventTypeName field to given value.

### HasEventTypeName

`func (o *ApiEventView) HasEventTypeName() bool`

HasEventTypeName returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiEventView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiEventView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiEventView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiEventView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ApiEventView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiEventView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiEventView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiEventView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *ApiEventView) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ApiEventView) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ApiEventView) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *ApiEventView) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetIsGlobalAdmin

`func (o *ApiEventView) GetIsGlobalAdmin() bool`

GetIsGlobalAdmin returns the IsGlobalAdmin field if non-nil, zero value otherwise.

### GetIsGlobalAdminOk

`func (o *ApiEventView) GetIsGlobalAdminOk() (*bool, bool)`

GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalAdmin

`func (o *ApiEventView) SetIsGlobalAdmin(v bool)`

SetIsGlobalAdmin sets IsGlobalAdmin field to given value.

### HasIsGlobalAdmin

`func (o *ApiEventView) HasIsGlobalAdmin() bool`

HasIsGlobalAdmin returns a boolean if a field has been set.

### GetLinks

`func (o *ApiEventView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiEventView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiEventView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetMetricName

`func (o *ApiEventView) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ApiEventView) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ApiEventView) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *ApiEventView) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetOpType

`func (o *ApiEventView) GetOpType() string`

GetOpType returns the OpType field if non-nil, zero value otherwise.

### GetOpTypeOk

`func (o *ApiEventView) GetOpTypeOk() (*string, bool)`

GetOpTypeOk returns a tuple with the OpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpType

`func (o *ApiEventView) SetOpType(v string)`

SetOpType sets OpType field to given value.

### HasOpType

`func (o *ApiEventView) HasOpType() bool`

HasOpType returns a boolean if a field has been set.

### GetOrgId

`func (o *ApiEventView) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ApiEventView) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ApiEventView) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ApiEventView) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPaymentId

`func (o *ApiEventView) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *ApiEventView) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *ApiEventView) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *ApiEventView) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPort

`func (o *ApiEventView) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApiEventView) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApiEventView) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApiEventView) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProviderEndpointId

`func (o *ApiEventView) GetProviderEndpointId() string`

GetProviderEndpointId returns the ProviderEndpointId field if non-nil, zero value otherwise.

### GetProviderEndpointIdOk

`func (o *ApiEventView) GetProviderEndpointIdOk() (*string, bool)`

GetProviderEndpointIdOk returns a tuple with the ProviderEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderEndpointId

`func (o *ApiEventView) SetProviderEndpointId(v string)`

SetProviderEndpointId sets ProviderEndpointId field to given value.

### HasProviderEndpointId

`func (o *ApiEventView) HasProviderEndpointId() bool`

HasProviderEndpointId returns a boolean if a field has been set.

### GetPublicKey

`func (o *ApiEventView) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ApiEventView) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ApiEventView) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *ApiEventView) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRaw

`func (o *ApiEventView) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *ApiEventView) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *ApiEventView) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *ApiEventView) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *ApiEventView) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *ApiEventView) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *ApiEventView) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *ApiEventView) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetReplicaSetName

`func (o *ApiEventView) GetReplicaSetName() string`

GetReplicaSetName returns the ReplicaSetName field if non-nil, zero value otherwise.

### GetReplicaSetNameOk

`func (o *ApiEventView) GetReplicaSetNameOk() (*string, bool)`

GetReplicaSetNameOk returns a tuple with the ReplicaSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaSetName

`func (o *ApiEventView) SetReplicaSetName(v string)`

SetReplicaSetName sets ReplicaSetName field to given value.

### HasReplicaSetName

`func (o *ApiEventView) HasReplicaSetName() bool`

HasReplicaSetName returns a boolean if a field has been set.

### GetShardName

`func (o *ApiEventView) GetShardName() string`

GetShardName returns the ShardName field if non-nil, zero value otherwise.

### GetShardNameOk

`func (o *ApiEventView) GetShardNameOk() (*string, bool)`

GetShardNameOk returns a tuple with the ShardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardName

`func (o *ApiEventView) SetShardName(v string)`

SetShardName sets ShardName field to given value.

### HasShardName

`func (o *ApiEventView) HasShardName() bool`

HasShardName returns a boolean if a field has been set.

### GetTargetPublicKey

`func (o *ApiEventView) GetTargetPublicKey() string`

GetTargetPublicKey returns the TargetPublicKey field if non-nil, zero value otherwise.

### GetTargetPublicKeyOk

`func (o *ApiEventView) GetTargetPublicKeyOk() (*string, bool)`

GetTargetPublicKeyOk returns a tuple with the TargetPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPublicKey

`func (o *ApiEventView) SetTargetPublicKey(v string)`

SetTargetPublicKey sets TargetPublicKey field to given value.

### HasTargetPublicKey

`func (o *ApiEventView) HasTargetPublicKey() bool`

HasTargetPublicKey returns a boolean if a field has been set.

### GetTargetUsername

`func (o *ApiEventView) GetTargetUsername() string`

GetTargetUsername returns the TargetUsername field if non-nil, zero value otherwise.

### GetTargetUsernameOk

`func (o *ApiEventView) GetTargetUsernameOk() (*string, bool)`

GetTargetUsernameOk returns a tuple with the TargetUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUsername

`func (o *ApiEventView) SetTargetUsername(v string)`

SetTargetUsername sets TargetUsername field to given value.

### HasTargetUsername

`func (o *ApiEventView) HasTargetUsername() bool`

HasTargetUsername returns a boolean if a field has been set.

### GetTeamId

`func (o *ApiEventView) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ApiEventView) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ApiEventView) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *ApiEventView) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetUserId

`func (o *ApiEventView) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiEventView) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiEventView) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiEventView) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *ApiEventView) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiEventView) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiEventView) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiEventView) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWhitelistEntry

`func (o *ApiEventView) GetWhitelistEntry() string`

GetWhitelistEntry returns the WhitelistEntry field if non-nil, zero value otherwise.

### GetWhitelistEntryOk

`func (o *ApiEventView) GetWhitelistEntryOk() (*string, bool)`

GetWhitelistEntryOk returns a tuple with the WhitelistEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistEntry

`func (o *ApiEventView) SetWhitelistEntry(v string)`

SetWhitelistEntry sets WhitelistEntry field to given value.

### HasWhitelistEntry

`func (o *ApiEventView) HasWhitelistEntry() bool`

HasWhitelistEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


