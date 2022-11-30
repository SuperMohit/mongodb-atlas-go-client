# ApiRestoreJobDeliveryView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthHeader** | Pointer to **string** | Header name to use when downloading the restore, used with &#x60;\&quot;delivery.methodName\&quot; : \&quot;HTTP\&quot;&#x60;. | [optional] [readonly] 
**AuthValue** | Pointer to **string** | Header value to use when downloading the restore, used with &#x60;\&quot;delivery.methodName\&quot; : \&quot;HTTP\&quot;&#x60;. | [optional] [readonly] 
**ExpirationHours** | Pointer to **int32** | Number of hours after the restore job completes that indicates when the Uniform Resource Locator (URL) for the snapshot download file expires. The resource returns this parameter when &#x60;\&quot;delivery.methodName\&quot; : \&quot;HTTP\&quot;&#x60;. | [optional] 
**Expires** | Pointer to **time.Time** | Date and time when the Uniform Resource Locator (URL) for the snapshot download file expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter when &#x60;\&quot;delivery.methodName\&quot; : \&quot;HTTP\&quot;&#x60;. | [optional] [readonly] 
**MaxDownloads** | Pointer to **int32** | Positive integer that indicates how many times you can use the Uniform Resource Locator (URL) for the snapshot download file. The resource returns this parameter when &#x60;\&quot;delivery.methodName\&quot; : \&quot;HTTP\&quot;&#x60;. | [optional] 
**MethodName** | **string** | Human-readable label that identifies the means for delivering the data. If you set &#x60;\&quot;delivery.methodName\&quot; : \&quot;AUTOMATED_RESTORE\&quot;&#x60;, you must also set: **delivery.targetGroupId** and **delivery.targetClusterName** or **delivery.targetClusterId**. The response returns &#x60;\&quot;delivery.methodName\&quot; : \&quot;HTTP\&quot;&#x60; as an automated restore uses HyperText Transport Protocol (HTTP) to deliver the restore job to the target host. | 
**StatusName** | Pointer to **string** | State of the downloadable snapshot file when MongoDB Cloud received this request. | [optional] [readonly] 
**TargetClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the target cluster. Use the **clusterId** returned in the response body of the **Get All Snapshots** and **Get a Snapshot** endpoints. This parameter applies when &#x60;\&quot;delivery.methodName\&quot; : \&quot;AUTOMATED_RESTORE\&quot;&#x60;.   If the target cluster doesn&#39;t have backup enabled, two resources return parameters with empty values:  - **Get All Snapshots** endpoint returns an empty results array without **clusterId** elements - **Get a Snapshot** endpoint doesn&#39;t return a **clusterId** parameter.  To return a response with the **clusterId** parameter, either use the **delivery.targetClusterName** parameter or enable backup on the target cluster. | [optional] 
**TargetClusterName** | Pointer to **string** | Human-readable label that identifies the target cluster. Use the **clusterName** returned in the response body of the **Get All Snapshots** and **Get a Snapshot** endpoints.  This parameter applies when &#x60;\&quot;delivery.methodName\&quot; : \&quot;AUTOMATED_RESTORE\&quot;&#x60;.  If the target cluster doesn&#39;t have backup enabled, two resources return parameters with empty values:  - **Get All Snapshots** endpoint returns an empty results array without **clusterId** elements - **Get a Snapshot** endpoint doesn&#39;t return a **clusterId** parameter.  To return a response with the **clusterId** parameter, either use the **delivery.targetClusterName** parameter or enable backup on the target cluster. | [optional] 
**TargetGroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that contains the destination cluster for the restore job. The resource returns this parameter when &#x60;\&quot;delivery.methodName\&quot; : \&quot;AUTOMATED_RESTORE\&quot;&#x60;. | [optional] 
**Url** | Pointer to **string** | Uniform Resource Locator (URL) from which you can download the restored snapshot data. Url includes the verification key. The resource returns this parameter when &#x60;\&quot;delivery.methodName\&quot; : \&quot;HTTP\&quot;&#x60;. | [optional] [readonly] 
**UrlV2** | Pointer to **string** | Uniform Resource Locator (URL) from which you can download the restored snapshot data. This should be preferred over **url**. The verification key must be sent as an HTTP header. The resource returns this parameter when &#x60;\&quot;delivery.methodName\&quot; : \&quot;HTTP\&quot;&#x60;. | [optional] [readonly] 

## Methods

### NewApiRestoreJobDeliveryView

`func NewApiRestoreJobDeliveryView(methodName string, ) *ApiRestoreJobDeliveryView`

NewApiRestoreJobDeliveryView instantiates a new ApiRestoreJobDeliveryView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRestoreJobDeliveryViewWithDefaults

`func NewApiRestoreJobDeliveryViewWithDefaults() *ApiRestoreJobDeliveryView`

NewApiRestoreJobDeliveryViewWithDefaults instantiates a new ApiRestoreJobDeliveryView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthHeader

`func (o *ApiRestoreJobDeliveryView) GetAuthHeader() string`

GetAuthHeader returns the AuthHeader field if non-nil, zero value otherwise.

### GetAuthHeaderOk

`func (o *ApiRestoreJobDeliveryView) GetAuthHeaderOk() (*string, bool)`

GetAuthHeaderOk returns a tuple with the AuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeader

`func (o *ApiRestoreJobDeliveryView) SetAuthHeader(v string)`

SetAuthHeader sets AuthHeader field to given value.

### HasAuthHeader

`func (o *ApiRestoreJobDeliveryView) HasAuthHeader() bool`

HasAuthHeader returns a boolean if a field has been set.

### GetAuthValue

`func (o *ApiRestoreJobDeliveryView) GetAuthValue() string`

GetAuthValue returns the AuthValue field if non-nil, zero value otherwise.

### GetAuthValueOk

`func (o *ApiRestoreJobDeliveryView) GetAuthValueOk() (*string, bool)`

GetAuthValueOk returns a tuple with the AuthValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthValue

`func (o *ApiRestoreJobDeliveryView) SetAuthValue(v string)`

SetAuthValue sets AuthValue field to given value.

### HasAuthValue

`func (o *ApiRestoreJobDeliveryView) HasAuthValue() bool`

HasAuthValue returns a boolean if a field has been set.

### GetExpirationHours

`func (o *ApiRestoreJobDeliveryView) GetExpirationHours() int32`

GetExpirationHours returns the ExpirationHours field if non-nil, zero value otherwise.

### GetExpirationHoursOk

`func (o *ApiRestoreJobDeliveryView) GetExpirationHoursOk() (*int32, bool)`

GetExpirationHoursOk returns a tuple with the ExpirationHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationHours

`func (o *ApiRestoreJobDeliveryView) SetExpirationHours(v int32)`

SetExpirationHours sets ExpirationHours field to given value.

### HasExpirationHours

`func (o *ApiRestoreJobDeliveryView) HasExpirationHours() bool`

HasExpirationHours returns a boolean if a field has been set.

### GetExpires

`func (o *ApiRestoreJobDeliveryView) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ApiRestoreJobDeliveryView) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ApiRestoreJobDeliveryView) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ApiRestoreJobDeliveryView) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetMaxDownloads

`func (o *ApiRestoreJobDeliveryView) GetMaxDownloads() int32`

GetMaxDownloads returns the MaxDownloads field if non-nil, zero value otherwise.

### GetMaxDownloadsOk

`func (o *ApiRestoreJobDeliveryView) GetMaxDownloadsOk() (*int32, bool)`

GetMaxDownloadsOk returns a tuple with the MaxDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDownloads

`func (o *ApiRestoreJobDeliveryView) SetMaxDownloads(v int32)`

SetMaxDownloads sets MaxDownloads field to given value.

### HasMaxDownloads

`func (o *ApiRestoreJobDeliveryView) HasMaxDownloads() bool`

HasMaxDownloads returns a boolean if a field has been set.

### GetMethodName

`func (o *ApiRestoreJobDeliveryView) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *ApiRestoreJobDeliveryView) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *ApiRestoreJobDeliveryView) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.


### GetStatusName

`func (o *ApiRestoreJobDeliveryView) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *ApiRestoreJobDeliveryView) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *ApiRestoreJobDeliveryView) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *ApiRestoreJobDeliveryView) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetTargetClusterId

`func (o *ApiRestoreJobDeliveryView) GetTargetClusterId() string`

GetTargetClusterId returns the TargetClusterId field if non-nil, zero value otherwise.

### GetTargetClusterIdOk

`func (o *ApiRestoreJobDeliveryView) GetTargetClusterIdOk() (*string, bool)`

GetTargetClusterIdOk returns a tuple with the TargetClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterId

`func (o *ApiRestoreJobDeliveryView) SetTargetClusterId(v string)`

SetTargetClusterId sets TargetClusterId field to given value.

### HasTargetClusterId

`func (o *ApiRestoreJobDeliveryView) HasTargetClusterId() bool`

HasTargetClusterId returns a boolean if a field has been set.

### GetTargetClusterName

`func (o *ApiRestoreJobDeliveryView) GetTargetClusterName() string`

GetTargetClusterName returns the TargetClusterName field if non-nil, zero value otherwise.

### GetTargetClusterNameOk

`func (o *ApiRestoreJobDeliveryView) GetTargetClusterNameOk() (*string, bool)`

GetTargetClusterNameOk returns a tuple with the TargetClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterName

`func (o *ApiRestoreJobDeliveryView) SetTargetClusterName(v string)`

SetTargetClusterName sets TargetClusterName field to given value.

### HasTargetClusterName

`func (o *ApiRestoreJobDeliveryView) HasTargetClusterName() bool`

HasTargetClusterName returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *ApiRestoreJobDeliveryView) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *ApiRestoreJobDeliveryView) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *ApiRestoreJobDeliveryView) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *ApiRestoreJobDeliveryView) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetUrl

`func (o *ApiRestoreJobDeliveryView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiRestoreJobDeliveryView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiRestoreJobDeliveryView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApiRestoreJobDeliveryView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlV2

`func (o *ApiRestoreJobDeliveryView) GetUrlV2() string`

GetUrlV2 returns the UrlV2 field if non-nil, zero value otherwise.

### GetUrlV2Ok

`func (o *ApiRestoreJobDeliveryView) GetUrlV2Ok() (*string, bool)`

GetUrlV2Ok returns a tuple with the UrlV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlV2

`func (o *ApiRestoreJobDeliveryView) SetUrlV2(v string)`

SetUrlV2 sets UrlV2 field to given value.

### HasUrlV2

`func (o *ApiRestoreJobDeliveryView) HasUrlV2() bool`

HasUrlV2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


