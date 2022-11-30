# ApiSystemStatusView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | [**ApiKeyView**](ApiKeyView.md) |  | 
**AppName** | **string** | Human-readable label that identifies the service from which you requested this response. | 
**Build** | **string** | Unique 40-hexadecimal digit hash that identifies the latest git commit merged for this application. | [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**Throttling** | **bool** | Flag that indicates whether someone enabled throttling on this service. | 

## Methods

### NewApiSystemStatusView

`func NewApiSystemStatusView(apiKey ApiKeyView, appName string, build string, links []Link, throttling bool, ) *ApiSystemStatusView`

NewApiSystemStatusView instantiates a new ApiSystemStatusView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSystemStatusViewWithDefaults

`func NewApiSystemStatusViewWithDefaults() *ApiSystemStatusView`

NewApiSystemStatusViewWithDefaults instantiates a new ApiSystemStatusView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *ApiSystemStatusView) GetApiKey() ApiKeyView`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiSystemStatusView) GetApiKeyOk() (*ApiKeyView, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiSystemStatusView) SetApiKey(v ApiKeyView)`

SetApiKey sets ApiKey field to given value.


### GetAppName

`func (o *ApiSystemStatusView) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ApiSystemStatusView) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ApiSystemStatusView) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetBuild

`func (o *ApiSystemStatusView) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ApiSystemStatusView) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ApiSystemStatusView) SetBuild(v string)`

SetBuild sets Build field to given value.


### GetLinks

`func (o *ApiSystemStatusView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiSystemStatusView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiSystemStatusView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetThrottling

`func (o *ApiSystemStatusView) GetThrottling() bool`

GetThrottling returns the Throttling field if non-nil, zero value otherwise.

### GetThrottlingOk

`func (o *ApiSystemStatusView) GetThrottlingOk() (*bool, bool)`

GetThrottlingOk returns a tuple with the Throttling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottling

`func (o *ApiSystemStatusView) SetThrottling(v bool)`

SetThrottling sets Throttling field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


