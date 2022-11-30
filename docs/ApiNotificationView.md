# ApiNotificationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | Pointer to **string** | Slack API token or Bot token that MongoDB Cloud needs to send alert notifications via Slack. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;SLACK\&quot;&#x60;. If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes the token. | [optional] 
**ChannelName** | Pointer to **string** | Name of the Slack channel to which MongoDB Cloud sends alert notifications. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;SLACK\&quot;&#x60;. | [optional] 
**DatadogApiKey** | Pointer to **string** | Datadog API Key that MongoDB Cloud needs to send alert notifications to Datadog. You can find this API key in the Datadog dashboard. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;DATADOG\&quot;&#x60;. | [optional] 
**DatadogRegion** | Pointer to **string** | Datadog region that indicates which API Uniform Resource Locator (URL) to use. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;DATADOG\&quot;&#x60;. | [optional] [default to "US"]
**DelayMin** | Pointer to **int32** | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. | [optional] 
**EmailAddress** | Pointer to **string** | Email address to which MongoDB Cloud sends alert notifications. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;EMAIL\&quot;&#x60;. You don’t need to set this value to send emails to individual or groups of MongoDB Cloud users including:  - specific MongoDB Cloud users (&#x60;\&quot;notifications.typeName\&quot; : \&quot;USER\&quot;&#x60;) - MongoDB Cloud users with specific project roles (&#x60;\&quot;notifications.typeName\&quot; : \&quot;GROUP\&quot;&#x60;) - MongoDB Cloud users with specific organization roles (&#x60;\&quot;notifications.typeName\&quot; : \&quot;ORG\&quot;&#x60;) - MongoDB Cloud teams (&#x60;\&quot;notifications.typeName\&quot; : \&quot;TEAM\&quot;&#x60;)  To send emails to one MongoDB Cloud user or grouping of users, set the **notifications.emailEnabled** parameter. | [optional] 
**EmailEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should send email notifications. The resource requires this parameter when one of the following values have been set:  - &#x60;\&quot;notifications.typeName\&quot; : \&quot;ORG\&quot;&#x60; - &#x60;\&quot;notifications.typeName\&quot; : \&quot;GROUP\&quot;&#x60; - &#x60;\&quot;notifications.typeName\&quot; : \&quot;USER\&quot;&#x60; | [optional] 
**IntervalMin** | Pointer to **int32** | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don&#39;t return this element. Configure and manage the notification interval within each of those services. | [optional] 
**MicrosoftTeamsWebhookUrl** | Pointer to **string** | Microsoft Teams Webhook Uniform Resource Locator (URL) that MongoDB Cloud needs to send this notification via Microsoft Teams. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;MICROSOFT_TEAMS\&quot;&#x60;. If the URL later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it. | [optional] 
**MobileNumber** | Pointer to **string** | Mobile phone number to which MongoDB Cloud sends alert notifications. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;SMS\&quot;&#x60;. | [optional] 
**NotificationToken** | Pointer to **string** | HipChat API token that MongoDB Cloud needs to send alert notifications to HipChat. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;HIP_CHAT\&quot;&#x60;\&quot;. If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes it. | [optional] 
**OpsGenieApiKey** | Pointer to **string** | API Key that MongoDB Cloud needs to send this notification via Opsgenie. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;OPS_GENIE\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it. | [optional] 
**OpsGenieRegion** | Pointer to **string** | Opsgenie region that indicates which API Uniform Resource Locator (URL) to use. | [optional] [default to "US"]
**Roles** | Pointer to **[]string** | List that contains the one or more organization or project roles that receive the configured alert. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;GROUP\&quot;&#x60; or &#x60;\&quot;notifications.typeName\&quot; : \&quot;ORG\&quot;&#x60;. If you include this parameter, MongoDB Cloud sends alerts only to users assigned the roles you specify in the array. If you omit this parameter, MongoDB Cloud sends alerts to users assigned any role. | [optional] 
**RoomName** | Pointer to **string** | HipChat API room name to which MongoDB Cloud sends alert notifications. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;HIP_CHAT\&quot;&#x60;\&quot;. | [optional] 
**ServiceKey** | Pointer to **string** | PagerDuty service key that MongoDB Cloud needs to send notifications via PagerDuty. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;PAGER_DUTY\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it. | [optional] 
**Severity** | Pointer to **string** | Degree of seriousness given to this notification. | [optional] 
**SmsEnabled** | Pointer to **bool** | Flag that indicates whether MongoDB Cloud should send text message notifications. The resource requires this parameter when one of the following values have been set:  - &#x60;\&quot;notifications.typeName\&quot; : \&quot;ORG\&quot;&#x60; - &#x60;\&quot;notifications.typeName\&quot; : \&quot;GROUP\&quot;&#x60; - &#x60;\&quot;notifications.typeName\&quot; : \&quot;USER\&quot;&#x60; | [optional] 
**TeamId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies one MongoDB Cloud team. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;TEAM\&quot;&#x60;. | [optional] 
**TeamName** | Pointer to **string** | Name of the MongoDB Cloud team that receives this notification. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;TEAM\&quot;&#x60;. | [optional] 
**TypeName** | Pointer to **string** | Human-readable label that displays the alert notification type. | [optional] 
**Username** | Pointer to **string** | MongoDB Cloud username of the person to whom MongoDB Cloud sends notifications. Specify only MongoDB Cloud users who belong to the project that owns the alert configuration. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;USER\&quot;&#x60;. | [optional] 
**VictorOpsApiKey** | Pointer to **string** | API key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;VICTOR_OPS\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it. | [optional] 
**VictorOpsRoutingKey** | Pointer to **string** | Routing key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when &#x60;\&quot;notifications.typeName\&quot; : \&quot;VICTOR_OPS\&quot;&#x60;. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it. | [optional] 
**WebhookSecret** | Pointer to **string** | An optional field for your webhook secret. | [optional] 
**WebhookUrl** | Pointer to **string** | Your webhook URL. | [optional] 

## Methods

### NewApiNotificationView

`func NewApiNotificationView() *ApiNotificationView`

NewApiNotificationView instantiates a new ApiNotificationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNotificationViewWithDefaults

`func NewApiNotificationViewWithDefaults() *ApiNotificationView`

NewApiNotificationViewWithDefaults instantiates a new ApiNotificationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiToken

`func (o *ApiNotificationView) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *ApiNotificationView) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *ApiNotificationView) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *ApiNotificationView) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetChannelName

`func (o *ApiNotificationView) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ApiNotificationView) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ApiNotificationView) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ApiNotificationView) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetDatadogApiKey

`func (o *ApiNotificationView) GetDatadogApiKey() string`

GetDatadogApiKey returns the DatadogApiKey field if non-nil, zero value otherwise.

### GetDatadogApiKeyOk

`func (o *ApiNotificationView) GetDatadogApiKeyOk() (*string, bool)`

GetDatadogApiKeyOk returns a tuple with the DatadogApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatadogApiKey

`func (o *ApiNotificationView) SetDatadogApiKey(v string)`

SetDatadogApiKey sets DatadogApiKey field to given value.

### HasDatadogApiKey

`func (o *ApiNotificationView) HasDatadogApiKey() bool`

HasDatadogApiKey returns a boolean if a field has been set.

### GetDatadogRegion

`func (o *ApiNotificationView) GetDatadogRegion() string`

GetDatadogRegion returns the DatadogRegion field if non-nil, zero value otherwise.

### GetDatadogRegionOk

`func (o *ApiNotificationView) GetDatadogRegionOk() (*string, bool)`

GetDatadogRegionOk returns a tuple with the DatadogRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatadogRegion

`func (o *ApiNotificationView) SetDatadogRegion(v string)`

SetDatadogRegion sets DatadogRegion field to given value.

### HasDatadogRegion

`func (o *ApiNotificationView) HasDatadogRegion() bool`

HasDatadogRegion returns a boolean if a field has been set.

### GetDelayMin

`func (o *ApiNotificationView) GetDelayMin() int32`

GetDelayMin returns the DelayMin field if non-nil, zero value otherwise.

### GetDelayMinOk

`func (o *ApiNotificationView) GetDelayMinOk() (*int32, bool)`

GetDelayMinOk returns a tuple with the DelayMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMin

`func (o *ApiNotificationView) SetDelayMin(v int32)`

SetDelayMin sets DelayMin field to given value.

### HasDelayMin

`func (o *ApiNotificationView) HasDelayMin() bool`

HasDelayMin returns a boolean if a field has been set.

### GetEmailAddress

`func (o *ApiNotificationView) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ApiNotificationView) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ApiNotificationView) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ApiNotificationView) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailEnabled

`func (o *ApiNotificationView) GetEmailEnabled() bool`

GetEmailEnabled returns the EmailEnabled field if non-nil, zero value otherwise.

### GetEmailEnabledOk

`func (o *ApiNotificationView) GetEmailEnabledOk() (*bool, bool)`

GetEmailEnabledOk returns a tuple with the EmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailEnabled

`func (o *ApiNotificationView) SetEmailEnabled(v bool)`

SetEmailEnabled sets EmailEnabled field to given value.

### HasEmailEnabled

`func (o *ApiNotificationView) HasEmailEnabled() bool`

HasEmailEnabled returns a boolean if a field has been set.

### GetIntervalMin

`func (o *ApiNotificationView) GetIntervalMin() int32`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *ApiNotificationView) GetIntervalMinOk() (*int32, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *ApiNotificationView) SetIntervalMin(v int32)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *ApiNotificationView) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetMicrosoftTeamsWebhookUrl

`func (o *ApiNotificationView) GetMicrosoftTeamsWebhookUrl() string`

GetMicrosoftTeamsWebhookUrl returns the MicrosoftTeamsWebhookUrl field if non-nil, zero value otherwise.

### GetMicrosoftTeamsWebhookUrlOk

`func (o *ApiNotificationView) GetMicrosoftTeamsWebhookUrlOk() (*string, bool)`

GetMicrosoftTeamsWebhookUrlOk returns a tuple with the MicrosoftTeamsWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftTeamsWebhookUrl

`func (o *ApiNotificationView) SetMicrosoftTeamsWebhookUrl(v string)`

SetMicrosoftTeamsWebhookUrl sets MicrosoftTeamsWebhookUrl field to given value.

### HasMicrosoftTeamsWebhookUrl

`func (o *ApiNotificationView) HasMicrosoftTeamsWebhookUrl() bool`

HasMicrosoftTeamsWebhookUrl returns a boolean if a field has been set.

### GetMobileNumber

`func (o *ApiNotificationView) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *ApiNotificationView) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *ApiNotificationView) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *ApiNotificationView) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetNotificationToken

`func (o *ApiNotificationView) GetNotificationToken() string`

GetNotificationToken returns the NotificationToken field if non-nil, zero value otherwise.

### GetNotificationTokenOk

`func (o *ApiNotificationView) GetNotificationTokenOk() (*string, bool)`

GetNotificationTokenOk returns a tuple with the NotificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationToken

`func (o *ApiNotificationView) SetNotificationToken(v string)`

SetNotificationToken sets NotificationToken field to given value.

### HasNotificationToken

`func (o *ApiNotificationView) HasNotificationToken() bool`

HasNotificationToken returns a boolean if a field has been set.

### GetOpsGenieApiKey

`func (o *ApiNotificationView) GetOpsGenieApiKey() string`

GetOpsGenieApiKey returns the OpsGenieApiKey field if non-nil, zero value otherwise.

### GetOpsGenieApiKeyOk

`func (o *ApiNotificationView) GetOpsGenieApiKeyOk() (*string, bool)`

GetOpsGenieApiKeyOk returns a tuple with the OpsGenieApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsGenieApiKey

`func (o *ApiNotificationView) SetOpsGenieApiKey(v string)`

SetOpsGenieApiKey sets OpsGenieApiKey field to given value.

### HasOpsGenieApiKey

`func (o *ApiNotificationView) HasOpsGenieApiKey() bool`

HasOpsGenieApiKey returns a boolean if a field has been set.

### GetOpsGenieRegion

`func (o *ApiNotificationView) GetOpsGenieRegion() string`

GetOpsGenieRegion returns the OpsGenieRegion field if non-nil, zero value otherwise.

### GetOpsGenieRegionOk

`func (o *ApiNotificationView) GetOpsGenieRegionOk() (*string, bool)`

GetOpsGenieRegionOk returns a tuple with the OpsGenieRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsGenieRegion

`func (o *ApiNotificationView) SetOpsGenieRegion(v string)`

SetOpsGenieRegion sets OpsGenieRegion field to given value.

### HasOpsGenieRegion

`func (o *ApiNotificationView) HasOpsGenieRegion() bool`

HasOpsGenieRegion returns a boolean if a field has been set.

### GetRoles

`func (o *ApiNotificationView) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiNotificationView) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiNotificationView) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiNotificationView) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRoomName

`func (o *ApiNotificationView) GetRoomName() string`

GetRoomName returns the RoomName field if non-nil, zero value otherwise.

### GetRoomNameOk

`func (o *ApiNotificationView) GetRoomNameOk() (*string, bool)`

GetRoomNameOk returns a tuple with the RoomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomName

`func (o *ApiNotificationView) SetRoomName(v string)`

SetRoomName sets RoomName field to given value.

### HasRoomName

`func (o *ApiNotificationView) HasRoomName() bool`

HasRoomName returns a boolean if a field has been set.

### GetServiceKey

`func (o *ApiNotificationView) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *ApiNotificationView) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *ApiNotificationView) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *ApiNotificationView) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetSeverity

`func (o *ApiNotificationView) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApiNotificationView) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApiNotificationView) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApiNotificationView) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSmsEnabled

`func (o *ApiNotificationView) GetSmsEnabled() bool`

GetSmsEnabled returns the SmsEnabled field if non-nil, zero value otherwise.

### GetSmsEnabledOk

`func (o *ApiNotificationView) GetSmsEnabledOk() (*bool, bool)`

GetSmsEnabledOk returns a tuple with the SmsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsEnabled

`func (o *ApiNotificationView) SetSmsEnabled(v bool)`

SetSmsEnabled sets SmsEnabled field to given value.

### HasSmsEnabled

`func (o *ApiNotificationView) HasSmsEnabled() bool`

HasSmsEnabled returns a boolean if a field has been set.

### GetTeamId

`func (o *ApiNotificationView) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ApiNotificationView) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ApiNotificationView) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *ApiNotificationView) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTeamName

`func (o *ApiNotificationView) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ApiNotificationView) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ApiNotificationView) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *ApiNotificationView) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetTypeName

`func (o *ApiNotificationView) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *ApiNotificationView) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *ApiNotificationView) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *ApiNotificationView) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetUsername

`func (o *ApiNotificationView) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiNotificationView) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiNotificationView) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiNotificationView) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVictorOpsApiKey

`func (o *ApiNotificationView) GetVictorOpsApiKey() string`

GetVictorOpsApiKey returns the VictorOpsApiKey field if non-nil, zero value otherwise.

### GetVictorOpsApiKeyOk

`func (o *ApiNotificationView) GetVictorOpsApiKeyOk() (*string, bool)`

GetVictorOpsApiKeyOk returns a tuple with the VictorOpsApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictorOpsApiKey

`func (o *ApiNotificationView) SetVictorOpsApiKey(v string)`

SetVictorOpsApiKey sets VictorOpsApiKey field to given value.

### HasVictorOpsApiKey

`func (o *ApiNotificationView) HasVictorOpsApiKey() bool`

HasVictorOpsApiKey returns a boolean if a field has been set.

### GetVictorOpsRoutingKey

`func (o *ApiNotificationView) GetVictorOpsRoutingKey() string`

GetVictorOpsRoutingKey returns the VictorOpsRoutingKey field if non-nil, zero value otherwise.

### GetVictorOpsRoutingKeyOk

`func (o *ApiNotificationView) GetVictorOpsRoutingKeyOk() (*string, bool)`

GetVictorOpsRoutingKeyOk returns a tuple with the VictorOpsRoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVictorOpsRoutingKey

`func (o *ApiNotificationView) SetVictorOpsRoutingKey(v string)`

SetVictorOpsRoutingKey sets VictorOpsRoutingKey field to given value.

### HasVictorOpsRoutingKey

`func (o *ApiNotificationView) HasVictorOpsRoutingKey() bool`

HasVictorOpsRoutingKey returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *ApiNotificationView) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *ApiNotificationView) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *ApiNotificationView) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *ApiNotificationView) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *ApiNotificationView) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ApiNotificationView) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ApiNotificationView) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ApiNotificationView) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


