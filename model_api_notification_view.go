/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApiNotificationView One or more targets for MongoDB Cloud to send information when an event triggers an alert condition.
type ApiNotificationView struct {
	// Slack API token or Bot token that MongoDB Cloud needs to send alert notifications via Slack. The resource requires this parameter when `\"notifications.typeName\" : \"SLACK\"`. If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes the token.
	ApiToken *string `json:"apiToken,omitempty"`
	// Name of the Slack channel to which MongoDB Cloud sends alert notifications. The resource requires this parameter when `\"notifications.typeName\" : \"SLACK\"`.
	ChannelName *string `json:"channelName,omitempty"`
	// Datadog API Key that MongoDB Cloud needs to send alert notifications to Datadog. You can find this API key in the Datadog dashboard. The resource requires this parameter when `\"notifications.typeName\" : \"DATADOG\"`.
	DatadogApiKey *string `json:"datadogApiKey,omitempty"`
	// Datadog region that indicates which API Uniform Resource Locator (URL) to use. The resource requires this parameter when `\"notifications.typeName\" : \"DATADOG\"`.
	DatadogRegion *string `json:"datadogRegion,omitempty"`
	// Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.
	DelayMin *int32 `json:"delayMin,omitempty"`
	// Email address to which MongoDB Cloud sends alert notifications. The resource requires this parameter when `\"notifications.typeName\" : \"EMAIL\"`. You don’t need to set this value to send emails to individual or groups of MongoDB Cloud users including:  - specific MongoDB Cloud users (`\"notifications.typeName\" : \"USER\"`) - MongoDB Cloud users with specific project roles (`\"notifications.typeName\" : \"GROUP\"`) - MongoDB Cloud users with specific organization roles (`\"notifications.typeName\" : \"ORG\"`) - MongoDB Cloud teams (`\"notifications.typeName\" : \"TEAM\"`)  To send emails to one MongoDB Cloud user or grouping of users, set the **notifications.emailEnabled** parameter.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Flag that indicates whether MongoDB Cloud should send email notifications. The resource requires this parameter when one of the following values have been set:  - `\"notifications.typeName\" : \"ORG\"` - `\"notifications.typeName\" : \"GROUP\"` - `\"notifications.typeName\" : \"USER\"`
	EmailEnabled *bool `json:"emailEnabled,omitempty"`
	// Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.  PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.
	IntervalMin *int32 `json:"intervalMin,omitempty"`
	// Microsoft Teams Webhook Uniform Resource Locator (URL) that MongoDB Cloud needs to send this notification via Microsoft Teams. The resource requires this parameter when `\"notifications.typeName\" : \"MICROSOFT_TEAMS\"`. If the URL later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	MicrosoftTeamsWebhookUrl *string `json:"microsoftTeamsWebhookUrl,omitempty"`
	// Mobile phone number to which MongoDB Cloud sends alert notifications. The resource requires this parameter when `\"notifications.typeName\" : \"SMS\"`.
	MobileNumber *string `json:"mobileNumber,omitempty"`
	// HipChat API token that MongoDB Cloud needs to send alert notifications to HipChat. The resource requires this parameter when `\"notifications.typeName\" : \"HIP_CHAT\"`\". If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes it.
	NotificationToken *string `json:"notificationToken,omitempty"`
	// API Key that MongoDB Cloud needs to send this notification via Opsgenie. The resource requires this parameter when `\"notifications.typeName\" : \"OPS_GENIE\"`. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	OpsGenieApiKey *string `json:"opsGenieApiKey,omitempty"`
	// Opsgenie region that indicates which API Uniform Resource Locator (URL) to use.
	OpsGenieRegion *string `json:"opsGenieRegion,omitempty"`
	// List that contains the one or more organization or project roles that receive the configured alert. The resource requires this parameter when `\"notifications.typeName\" : \"GROUP\"` or `\"notifications.typeName\" : \"ORG\"`. If you include this parameter, MongoDB Cloud sends alerts only to users assigned the roles you specify in the array. If you omit this parameter, MongoDB Cloud sends alerts to users assigned any role.
	Roles []string `json:"roles,omitempty"`
	// HipChat API room name to which MongoDB Cloud sends alert notifications. The resource requires this parameter when `\"notifications.typeName\" : \"HIP_CHAT\"`\".
	RoomName *string `json:"roomName,omitempty"`
	// PagerDuty service key that MongoDB Cloud needs to send notifications via PagerDuty. The resource requires this parameter when `\"notifications.typeName\" : \"PAGER_DUTY\"`. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	ServiceKey *string `json:"serviceKey,omitempty"`
	// Degree of seriousness given to this notification.
	Severity *string `json:"severity,omitempty"`
	// Flag that indicates whether MongoDB Cloud should send text message notifications. The resource requires this parameter when one of the following values have been set:  - `\"notifications.typeName\" : \"ORG\"` - `\"notifications.typeName\" : \"GROUP\"` - `\"notifications.typeName\" : \"USER\"`
	SmsEnabled *bool `json:"smsEnabled,omitempty"`
	// Unique 24-hexadecimal digit string that identifies one MongoDB Cloud team. The resource requires this parameter when `\"notifications.typeName\" : \"TEAM\"`.
	TeamId *string `json:"teamId,omitempty"`
	// Name of the MongoDB Cloud team that receives this notification. The resource requires this parameter when `\"notifications.typeName\" : \"TEAM\"`.
	TeamName *string `json:"teamName,omitempty"`
	// Human-readable label that displays the alert notification type.
	TypeName *string `json:"typeName,omitempty"`
	// MongoDB Cloud username of the person to whom MongoDB Cloud sends notifications. Specify only MongoDB Cloud users who belong to the project that owns the alert configuration. The resource requires this parameter when `\"notifications.typeName\" : \"USER\"`.
	Username *string `json:"username,omitempty"`
	// API key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when `\"notifications.typeName\" : \"VICTOR_OPS\"`. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	VictorOpsApiKey *string `json:"victorOpsApiKey,omitempty"`
	// Routing key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. The resource requires this parameter when `\"notifications.typeName\" : \"VICTOR_OPS\"`. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.
	VictorOpsRoutingKey *string `json:"victorOpsRoutingKey,omitempty"`
	// An optional field for your webhook secret.
	WebhookSecret *string `json:"webhookSecret,omitempty"`
	// Your webhook URL.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}

// NewApiNotificationView instantiates a new ApiNotificationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNotificationView() *ApiNotificationView {
	this := ApiNotificationView{}
	var datadogRegion string = "US"
	this.DatadogRegion = &datadogRegion
	var opsGenieRegion string = "US"
	this.OpsGenieRegion = &opsGenieRegion
	return &this
}

// NewApiNotificationViewWithDefaults instantiates a new ApiNotificationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNotificationViewWithDefaults() *ApiNotificationView {
	this := ApiNotificationView{}
	var datadogRegion string = "US"
	this.DatadogRegion = &datadogRegion
	var opsGenieRegion string = "US"
	this.OpsGenieRegion = &opsGenieRegion
	return &this
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise.
func (o *ApiNotificationView) GetApiToken() string {
	if o == nil || isNil(o.ApiToken) {
		var ret string
		return ret
	}
	return *o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetApiTokenOk() (*string, bool) {
	if o == nil || isNil(o.ApiToken) {
    return nil, false
	}
	return o.ApiToken, true
}

// HasApiToken returns a boolean if a field has been set.
func (o *ApiNotificationView) HasApiToken() bool {
	if o != nil && !isNil(o.ApiToken) {
		return true
	}

	return false
}

// SetApiToken gets a reference to the given string and assigns it to the ApiToken field.
func (o *ApiNotificationView) SetApiToken(v string) {
	o.ApiToken = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *ApiNotificationView) GetChannelName() string {
	if o == nil || isNil(o.ChannelName) {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetChannelNameOk() (*string, bool) {
	if o == nil || isNil(o.ChannelName) {
    return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *ApiNotificationView) HasChannelName() bool {
	if o != nil && !isNil(o.ChannelName) {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *ApiNotificationView) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetDatadogApiKey returns the DatadogApiKey field value if set, zero value otherwise.
func (o *ApiNotificationView) GetDatadogApiKey() string {
	if o == nil || isNil(o.DatadogApiKey) {
		var ret string
		return ret
	}
	return *o.DatadogApiKey
}

// GetDatadogApiKeyOk returns a tuple with the DatadogApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetDatadogApiKeyOk() (*string, bool) {
	if o == nil || isNil(o.DatadogApiKey) {
    return nil, false
	}
	return o.DatadogApiKey, true
}

// HasDatadogApiKey returns a boolean if a field has been set.
func (o *ApiNotificationView) HasDatadogApiKey() bool {
	if o != nil && !isNil(o.DatadogApiKey) {
		return true
	}

	return false
}

// SetDatadogApiKey gets a reference to the given string and assigns it to the DatadogApiKey field.
func (o *ApiNotificationView) SetDatadogApiKey(v string) {
	o.DatadogApiKey = &v
}

// GetDatadogRegion returns the DatadogRegion field value if set, zero value otherwise.
func (o *ApiNotificationView) GetDatadogRegion() string {
	if o == nil || isNil(o.DatadogRegion) {
		var ret string
		return ret
	}
	return *o.DatadogRegion
}

// GetDatadogRegionOk returns a tuple with the DatadogRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetDatadogRegionOk() (*string, bool) {
	if o == nil || isNil(o.DatadogRegion) {
    return nil, false
	}
	return o.DatadogRegion, true
}

// HasDatadogRegion returns a boolean if a field has been set.
func (o *ApiNotificationView) HasDatadogRegion() bool {
	if o != nil && !isNil(o.DatadogRegion) {
		return true
	}

	return false
}

// SetDatadogRegion gets a reference to the given string and assigns it to the DatadogRegion field.
func (o *ApiNotificationView) SetDatadogRegion(v string) {
	o.DatadogRegion = &v
}

// GetDelayMin returns the DelayMin field value if set, zero value otherwise.
func (o *ApiNotificationView) GetDelayMin() int32 {
	if o == nil || isNil(o.DelayMin) {
		var ret int32
		return ret
	}
	return *o.DelayMin
}

// GetDelayMinOk returns a tuple with the DelayMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetDelayMinOk() (*int32, bool) {
	if o == nil || isNil(o.DelayMin) {
    return nil, false
	}
	return o.DelayMin, true
}

// HasDelayMin returns a boolean if a field has been set.
func (o *ApiNotificationView) HasDelayMin() bool {
	if o != nil && !isNil(o.DelayMin) {
		return true
	}

	return false
}

// SetDelayMin gets a reference to the given int32 and assigns it to the DelayMin field.
func (o *ApiNotificationView) SetDelayMin(v int32) {
	o.DelayMin = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ApiNotificationView) GetEmailAddress() string {
	if o == nil || isNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetEmailAddressOk() (*string, bool) {
	if o == nil || isNil(o.EmailAddress) {
    return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *ApiNotificationView) HasEmailAddress() bool {
	if o != nil && !isNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ApiNotificationView) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetEmailEnabled returns the EmailEnabled field value if set, zero value otherwise.
func (o *ApiNotificationView) GetEmailEnabled() bool {
	if o == nil || isNil(o.EmailEnabled) {
		var ret bool
		return ret
	}
	return *o.EmailEnabled
}

// GetEmailEnabledOk returns a tuple with the EmailEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetEmailEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.EmailEnabled) {
    return nil, false
	}
	return o.EmailEnabled, true
}

// HasEmailEnabled returns a boolean if a field has been set.
func (o *ApiNotificationView) HasEmailEnabled() bool {
	if o != nil && !isNil(o.EmailEnabled) {
		return true
	}

	return false
}

// SetEmailEnabled gets a reference to the given bool and assigns it to the EmailEnabled field.
func (o *ApiNotificationView) SetEmailEnabled(v bool) {
	o.EmailEnabled = &v
}

// GetIntervalMin returns the IntervalMin field value if set, zero value otherwise.
func (o *ApiNotificationView) GetIntervalMin() int32 {
	if o == nil || isNil(o.IntervalMin) {
		var ret int32
		return ret
	}
	return *o.IntervalMin
}

// GetIntervalMinOk returns a tuple with the IntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetIntervalMinOk() (*int32, bool) {
	if o == nil || isNil(o.IntervalMin) {
    return nil, false
	}
	return o.IntervalMin, true
}

// HasIntervalMin returns a boolean if a field has been set.
func (o *ApiNotificationView) HasIntervalMin() bool {
	if o != nil && !isNil(o.IntervalMin) {
		return true
	}

	return false
}

// SetIntervalMin gets a reference to the given int32 and assigns it to the IntervalMin field.
func (o *ApiNotificationView) SetIntervalMin(v int32) {
	o.IntervalMin = &v
}

// GetMicrosoftTeamsWebhookUrl returns the MicrosoftTeamsWebhookUrl field value if set, zero value otherwise.
func (o *ApiNotificationView) GetMicrosoftTeamsWebhookUrl() string {
	if o == nil || isNil(o.MicrosoftTeamsWebhookUrl) {
		var ret string
		return ret
	}
	return *o.MicrosoftTeamsWebhookUrl
}

// GetMicrosoftTeamsWebhookUrlOk returns a tuple with the MicrosoftTeamsWebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetMicrosoftTeamsWebhookUrlOk() (*string, bool) {
	if o == nil || isNil(o.MicrosoftTeamsWebhookUrl) {
    return nil, false
	}
	return o.MicrosoftTeamsWebhookUrl, true
}

// HasMicrosoftTeamsWebhookUrl returns a boolean if a field has been set.
func (o *ApiNotificationView) HasMicrosoftTeamsWebhookUrl() bool {
	if o != nil && !isNil(o.MicrosoftTeamsWebhookUrl) {
		return true
	}

	return false
}

// SetMicrosoftTeamsWebhookUrl gets a reference to the given string and assigns it to the MicrosoftTeamsWebhookUrl field.
func (o *ApiNotificationView) SetMicrosoftTeamsWebhookUrl(v string) {
	o.MicrosoftTeamsWebhookUrl = &v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *ApiNotificationView) GetMobileNumber() string {
	if o == nil || isNil(o.MobileNumber) {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetMobileNumberOk() (*string, bool) {
	if o == nil || isNil(o.MobileNumber) {
    return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *ApiNotificationView) HasMobileNumber() bool {
	if o != nil && !isNil(o.MobileNumber) {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *ApiNotificationView) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetNotificationToken returns the NotificationToken field value if set, zero value otherwise.
func (o *ApiNotificationView) GetNotificationToken() string {
	if o == nil || isNil(o.NotificationToken) {
		var ret string
		return ret
	}
	return *o.NotificationToken
}

// GetNotificationTokenOk returns a tuple with the NotificationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetNotificationTokenOk() (*string, bool) {
	if o == nil || isNil(o.NotificationToken) {
    return nil, false
	}
	return o.NotificationToken, true
}

// HasNotificationToken returns a boolean if a field has been set.
func (o *ApiNotificationView) HasNotificationToken() bool {
	if o != nil && !isNil(o.NotificationToken) {
		return true
	}

	return false
}

// SetNotificationToken gets a reference to the given string and assigns it to the NotificationToken field.
func (o *ApiNotificationView) SetNotificationToken(v string) {
	o.NotificationToken = &v
}

// GetOpsGenieApiKey returns the OpsGenieApiKey field value if set, zero value otherwise.
func (o *ApiNotificationView) GetOpsGenieApiKey() string {
	if o == nil || isNil(o.OpsGenieApiKey) {
		var ret string
		return ret
	}
	return *o.OpsGenieApiKey
}

// GetOpsGenieApiKeyOk returns a tuple with the OpsGenieApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetOpsGenieApiKeyOk() (*string, bool) {
	if o == nil || isNil(o.OpsGenieApiKey) {
    return nil, false
	}
	return o.OpsGenieApiKey, true
}

// HasOpsGenieApiKey returns a boolean if a field has been set.
func (o *ApiNotificationView) HasOpsGenieApiKey() bool {
	if o != nil && !isNil(o.OpsGenieApiKey) {
		return true
	}

	return false
}

// SetOpsGenieApiKey gets a reference to the given string and assigns it to the OpsGenieApiKey field.
func (o *ApiNotificationView) SetOpsGenieApiKey(v string) {
	o.OpsGenieApiKey = &v
}

// GetOpsGenieRegion returns the OpsGenieRegion field value if set, zero value otherwise.
func (o *ApiNotificationView) GetOpsGenieRegion() string {
	if o == nil || isNil(o.OpsGenieRegion) {
		var ret string
		return ret
	}
	return *o.OpsGenieRegion
}

// GetOpsGenieRegionOk returns a tuple with the OpsGenieRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetOpsGenieRegionOk() (*string, bool) {
	if o == nil || isNil(o.OpsGenieRegion) {
    return nil, false
	}
	return o.OpsGenieRegion, true
}

// HasOpsGenieRegion returns a boolean if a field has been set.
func (o *ApiNotificationView) HasOpsGenieRegion() bool {
	if o != nil && !isNil(o.OpsGenieRegion) {
		return true
	}

	return false
}

// SetOpsGenieRegion gets a reference to the given string and assigns it to the OpsGenieRegion field.
func (o *ApiNotificationView) SetOpsGenieRegion(v string) {
	o.OpsGenieRegion = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApiNotificationView) GetRoles() []string {
	if o == nil || isNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetRolesOk() ([]string, bool) {
	if o == nil || isNil(o.Roles) {
    return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApiNotificationView) HasRoles() bool {
	if o != nil && !isNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *ApiNotificationView) SetRoles(v []string) {
	o.Roles = v
}

// GetRoomName returns the RoomName field value if set, zero value otherwise.
func (o *ApiNotificationView) GetRoomName() string {
	if o == nil || isNil(o.RoomName) {
		var ret string
		return ret
	}
	return *o.RoomName
}

// GetRoomNameOk returns a tuple with the RoomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetRoomNameOk() (*string, bool) {
	if o == nil || isNil(o.RoomName) {
    return nil, false
	}
	return o.RoomName, true
}

// HasRoomName returns a boolean if a field has been set.
func (o *ApiNotificationView) HasRoomName() bool {
	if o != nil && !isNil(o.RoomName) {
		return true
	}

	return false
}

// SetRoomName gets a reference to the given string and assigns it to the RoomName field.
func (o *ApiNotificationView) SetRoomName(v string) {
	o.RoomName = &v
}

// GetServiceKey returns the ServiceKey field value if set, zero value otherwise.
func (o *ApiNotificationView) GetServiceKey() string {
	if o == nil || isNil(o.ServiceKey) {
		var ret string
		return ret
	}
	return *o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetServiceKeyOk() (*string, bool) {
	if o == nil || isNil(o.ServiceKey) {
    return nil, false
	}
	return o.ServiceKey, true
}

// HasServiceKey returns a boolean if a field has been set.
func (o *ApiNotificationView) HasServiceKey() bool {
	if o != nil && !isNil(o.ServiceKey) {
		return true
	}

	return false
}

// SetServiceKey gets a reference to the given string and assigns it to the ServiceKey field.
func (o *ApiNotificationView) SetServiceKey(v string) {
	o.ServiceKey = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ApiNotificationView) GetSeverity() string {
	if o == nil || isNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetSeverityOk() (*string, bool) {
	if o == nil || isNil(o.Severity) {
    return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ApiNotificationView) HasSeverity() bool {
	if o != nil && !isNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ApiNotificationView) SetSeverity(v string) {
	o.Severity = &v
}

// GetSmsEnabled returns the SmsEnabled field value if set, zero value otherwise.
func (o *ApiNotificationView) GetSmsEnabled() bool {
	if o == nil || isNil(o.SmsEnabled) {
		var ret bool
		return ret
	}
	return *o.SmsEnabled
}

// GetSmsEnabledOk returns a tuple with the SmsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetSmsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.SmsEnabled) {
    return nil, false
	}
	return o.SmsEnabled, true
}

// HasSmsEnabled returns a boolean if a field has been set.
func (o *ApiNotificationView) HasSmsEnabled() bool {
	if o != nil && !isNil(o.SmsEnabled) {
		return true
	}

	return false
}

// SetSmsEnabled gets a reference to the given bool and assigns it to the SmsEnabled field.
func (o *ApiNotificationView) SetSmsEnabled(v bool) {
	o.SmsEnabled = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *ApiNotificationView) GetTeamId() string {
	if o == nil || isNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetTeamIdOk() (*string, bool) {
	if o == nil || isNil(o.TeamId) {
    return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *ApiNotificationView) HasTeamId() bool {
	if o != nil && !isNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *ApiNotificationView) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *ApiNotificationView) GetTeamName() string {
	if o == nil || isNil(o.TeamName) {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetTeamNameOk() (*string, bool) {
	if o == nil || isNil(o.TeamName) {
    return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *ApiNotificationView) HasTeamName() bool {
	if o != nil && !isNil(o.TeamName) {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *ApiNotificationView) SetTeamName(v string) {
	o.TeamName = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *ApiNotificationView) GetTypeName() string {
	if o == nil || isNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetTypeNameOk() (*string, bool) {
	if o == nil || isNil(o.TypeName) {
    return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *ApiNotificationView) HasTypeName() bool {
	if o != nil && !isNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *ApiNotificationView) SetTypeName(v string) {
	o.TypeName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiNotificationView) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiNotificationView) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiNotificationView) SetUsername(v string) {
	o.Username = &v
}

// GetVictorOpsApiKey returns the VictorOpsApiKey field value if set, zero value otherwise.
func (o *ApiNotificationView) GetVictorOpsApiKey() string {
	if o == nil || isNil(o.VictorOpsApiKey) {
		var ret string
		return ret
	}
	return *o.VictorOpsApiKey
}

// GetVictorOpsApiKeyOk returns a tuple with the VictorOpsApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetVictorOpsApiKeyOk() (*string, bool) {
	if o == nil || isNil(o.VictorOpsApiKey) {
    return nil, false
	}
	return o.VictorOpsApiKey, true
}

// HasVictorOpsApiKey returns a boolean if a field has been set.
func (o *ApiNotificationView) HasVictorOpsApiKey() bool {
	if o != nil && !isNil(o.VictorOpsApiKey) {
		return true
	}

	return false
}

// SetVictorOpsApiKey gets a reference to the given string and assigns it to the VictorOpsApiKey field.
func (o *ApiNotificationView) SetVictorOpsApiKey(v string) {
	o.VictorOpsApiKey = &v
}

// GetVictorOpsRoutingKey returns the VictorOpsRoutingKey field value if set, zero value otherwise.
func (o *ApiNotificationView) GetVictorOpsRoutingKey() string {
	if o == nil || isNil(o.VictorOpsRoutingKey) {
		var ret string
		return ret
	}
	return *o.VictorOpsRoutingKey
}

// GetVictorOpsRoutingKeyOk returns a tuple with the VictorOpsRoutingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetVictorOpsRoutingKeyOk() (*string, bool) {
	if o == nil || isNil(o.VictorOpsRoutingKey) {
    return nil, false
	}
	return o.VictorOpsRoutingKey, true
}

// HasVictorOpsRoutingKey returns a boolean if a field has been set.
func (o *ApiNotificationView) HasVictorOpsRoutingKey() bool {
	if o != nil && !isNil(o.VictorOpsRoutingKey) {
		return true
	}

	return false
}

// SetVictorOpsRoutingKey gets a reference to the given string and assigns it to the VictorOpsRoutingKey field.
func (o *ApiNotificationView) SetVictorOpsRoutingKey(v string) {
	o.VictorOpsRoutingKey = &v
}

// GetWebhookSecret returns the WebhookSecret field value if set, zero value otherwise.
func (o *ApiNotificationView) GetWebhookSecret() string {
	if o == nil || isNil(o.WebhookSecret) {
		var ret string
		return ret
	}
	return *o.WebhookSecret
}

// GetWebhookSecretOk returns a tuple with the WebhookSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetWebhookSecretOk() (*string, bool) {
	if o == nil || isNil(o.WebhookSecret) {
    return nil, false
	}
	return o.WebhookSecret, true
}

// HasWebhookSecret returns a boolean if a field has been set.
func (o *ApiNotificationView) HasWebhookSecret() bool {
	if o != nil && !isNil(o.WebhookSecret) {
		return true
	}

	return false
}

// SetWebhookSecret gets a reference to the given string and assigns it to the WebhookSecret field.
func (o *ApiNotificationView) SetWebhookSecret(v string) {
	o.WebhookSecret = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *ApiNotificationView) GetWebhookUrl() string {
	if o == nil || isNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationView) GetWebhookUrlOk() (*string, bool) {
	if o == nil || isNil(o.WebhookUrl) {
    return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *ApiNotificationView) HasWebhookUrl() bool {
	if o != nil && !isNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *ApiNotificationView) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

func (o ApiNotificationView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiToken) {
		toSerialize["apiToken"] = o.ApiToken
	}
	if !isNil(o.ChannelName) {
		toSerialize["channelName"] = o.ChannelName
	}
	if !isNil(o.DatadogApiKey) {
		toSerialize["datadogApiKey"] = o.DatadogApiKey
	}
	if !isNil(o.DatadogRegion) {
		toSerialize["datadogRegion"] = o.DatadogRegion
	}
	if !isNil(o.DelayMin) {
		toSerialize["delayMin"] = o.DelayMin
	}
	if !isNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !isNil(o.EmailEnabled) {
		toSerialize["emailEnabled"] = o.EmailEnabled
	}
	if !isNil(o.IntervalMin) {
		toSerialize["intervalMin"] = o.IntervalMin
	}
	if !isNil(o.MicrosoftTeamsWebhookUrl) {
		toSerialize["microsoftTeamsWebhookUrl"] = o.MicrosoftTeamsWebhookUrl
	}
	if !isNil(o.MobileNumber) {
		toSerialize["mobileNumber"] = o.MobileNumber
	}
	if !isNil(o.NotificationToken) {
		toSerialize["notificationToken"] = o.NotificationToken
	}
	if !isNil(o.OpsGenieApiKey) {
		toSerialize["opsGenieApiKey"] = o.OpsGenieApiKey
	}
	if !isNil(o.OpsGenieRegion) {
		toSerialize["opsGenieRegion"] = o.OpsGenieRegion
	}
	if !isNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !isNil(o.RoomName) {
		toSerialize["roomName"] = o.RoomName
	}
	if !isNil(o.ServiceKey) {
		toSerialize["serviceKey"] = o.ServiceKey
	}
	if !isNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !isNil(o.SmsEnabled) {
		toSerialize["smsEnabled"] = o.SmsEnabled
	}
	if !isNil(o.TeamId) {
		toSerialize["teamId"] = o.TeamId
	}
	if !isNil(o.TeamName) {
		toSerialize["teamName"] = o.TeamName
	}
	if !isNil(o.TypeName) {
		toSerialize["typeName"] = o.TypeName
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.VictorOpsApiKey) {
		toSerialize["victorOpsApiKey"] = o.VictorOpsApiKey
	}
	if !isNil(o.VictorOpsRoutingKey) {
		toSerialize["victorOpsRoutingKey"] = o.VictorOpsRoutingKey
	}
	if !isNil(o.WebhookSecret) {
		toSerialize["webhookSecret"] = o.WebhookSecret
	}
	if !isNil(o.WebhookUrl) {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	return json.Marshal(toSerialize)
}

type NullableApiNotificationView struct {
	value *ApiNotificationView
	isSet bool
}

func (v NullableApiNotificationView) Get() *ApiNotificationView {
	return v.value
}

func (v *NullableApiNotificationView) Set(val *ApiNotificationView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNotificationView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNotificationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNotificationView(val *ApiNotificationView) *NullableApiNotificationView {
	return &NullableApiNotificationView{value: val, isSet: true}
}

func (v NullableApiNotificationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNotificationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


