# ApiLineItemView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster that incurred the charge. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**DiscountCents** | Pointer to **int64** | Sum by which MongoDB discounted this line item. MongoDB Cloud expresses this value in cents (100ths of one US Dollar). The resource returns this parameter when a discount applies. | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | Date and time when when MongoDB Cloud finished charging for this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project associated to this line item. | [optional] [readonly] 
**GroupName** | Pointer to **string** | Human-readable label that identifies the project. | [optional] 
**Note** | Pointer to **string** | Comment that applies to this line item. | [optional] [readonly] 
**PercentDiscount** | Pointer to **float32** | Percentage by which MongoDB discounted this line item. The resource returns this parameter when a discount applies. | [optional] [readonly] 
**Quantity** | Pointer to **float64** | Number of units included for the line item. These can be expressions of storage (GB), time (hours), or other units. | [optional] [readonly] 
**Sku** | Pointer to **string** | Human-readable description of the service that this line item provided. This Stock Keeping Unit (SKU) could be the instance type, a support charge, advanced security, or another service. | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud began charging for this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**StitchAppName** | Pointer to **string** | Human-readable label that identifies the Atlas App Service associated with this line item. | [optional] [readonly] 
**TierLowerBound** | Pointer to **float64** | Lower bound for usage amount range in current SKU tier.   **NOTE**: **lineItems[n].tierLowerBound** appears only if your **lineItems[n].sku** is tiered. | [optional] [readonly] 
**TierUpperBound** | Pointer to **float64** | Upper bound for usage amount range in current SKU tier.   **NOTE**: **lineItems[n].tierUpperBound** appears only if your **lineItems[n].sku** is tiered. | [optional] [readonly] 
**TotalPriceCents** | Pointer to **int64** | Sum of the cost set for this line item. MongoDB Cloud expresses this value in cents (100ths of one US Dollar) and calculates this value as **unitPriceDollars** × **quantity** × 100. | [optional] [readonly] 
**Unit** | Pointer to **string** | Element used to express what **quantity** this line item measures. This value can be elements of time, storage capacity, and the like. | [optional] [readonly] 
**UnitPriceDollars** | Pointer to **float64** | Value per **unit** for this line item expressed in US Dollars. | [optional] [readonly] 

## Methods

### NewApiLineItemView

`func NewApiLineItemView() *ApiLineItemView`

NewApiLineItemView instantiates a new ApiLineItemView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLineItemViewWithDefaults

`func NewApiLineItemViewWithDefaults() *ApiLineItemView`

NewApiLineItemViewWithDefaults instantiates a new ApiLineItemView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ApiLineItemView) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ApiLineItemView) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ApiLineItemView) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ApiLineItemView) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCreated

`func (o *ApiLineItemView) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiLineItemView) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiLineItemView) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiLineItemView) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDiscountCents

`func (o *ApiLineItemView) GetDiscountCents() int64`

GetDiscountCents returns the DiscountCents field if non-nil, zero value otherwise.

### GetDiscountCentsOk

`func (o *ApiLineItemView) GetDiscountCentsOk() (*int64, bool)`

GetDiscountCentsOk returns a tuple with the DiscountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCents

`func (o *ApiLineItemView) SetDiscountCents(v int64)`

SetDiscountCents sets DiscountCents field to given value.

### HasDiscountCents

`func (o *ApiLineItemView) HasDiscountCents() bool`

HasDiscountCents returns a boolean if a field has been set.

### GetEndDate

`func (o *ApiLineItemView) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ApiLineItemView) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ApiLineItemView) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ApiLineItemView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiLineItemView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiLineItemView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiLineItemView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiLineItemView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *ApiLineItemView) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ApiLineItemView) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ApiLineItemView) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ApiLineItemView) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetNote

`func (o *ApiLineItemView) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ApiLineItemView) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ApiLineItemView) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ApiLineItemView) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPercentDiscount

`func (o *ApiLineItemView) GetPercentDiscount() float32`

GetPercentDiscount returns the PercentDiscount field if non-nil, zero value otherwise.

### GetPercentDiscountOk

`func (o *ApiLineItemView) GetPercentDiscountOk() (*float32, bool)`

GetPercentDiscountOk returns a tuple with the PercentDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentDiscount

`func (o *ApiLineItemView) SetPercentDiscount(v float32)`

SetPercentDiscount sets PercentDiscount field to given value.

### HasPercentDiscount

`func (o *ApiLineItemView) HasPercentDiscount() bool`

HasPercentDiscount returns a boolean if a field has been set.

### GetQuantity

`func (o *ApiLineItemView) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ApiLineItemView) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ApiLineItemView) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ApiLineItemView) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSku

`func (o *ApiLineItemView) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ApiLineItemView) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ApiLineItemView) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ApiLineItemView) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetStartDate

`func (o *ApiLineItemView) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ApiLineItemView) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ApiLineItemView) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ApiLineItemView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStitchAppName

`func (o *ApiLineItemView) GetStitchAppName() string`

GetStitchAppName returns the StitchAppName field if non-nil, zero value otherwise.

### GetStitchAppNameOk

`func (o *ApiLineItemView) GetStitchAppNameOk() (*string, bool)`

GetStitchAppNameOk returns a tuple with the StitchAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStitchAppName

`func (o *ApiLineItemView) SetStitchAppName(v string)`

SetStitchAppName sets StitchAppName field to given value.

### HasStitchAppName

`func (o *ApiLineItemView) HasStitchAppName() bool`

HasStitchAppName returns a boolean if a field has been set.

### GetTierLowerBound

`func (o *ApiLineItemView) GetTierLowerBound() float64`

GetTierLowerBound returns the TierLowerBound field if non-nil, zero value otherwise.

### GetTierLowerBoundOk

`func (o *ApiLineItemView) GetTierLowerBoundOk() (*float64, bool)`

GetTierLowerBoundOk returns a tuple with the TierLowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierLowerBound

`func (o *ApiLineItemView) SetTierLowerBound(v float64)`

SetTierLowerBound sets TierLowerBound field to given value.

### HasTierLowerBound

`func (o *ApiLineItemView) HasTierLowerBound() bool`

HasTierLowerBound returns a boolean if a field has been set.

### GetTierUpperBound

`func (o *ApiLineItemView) GetTierUpperBound() float64`

GetTierUpperBound returns the TierUpperBound field if non-nil, zero value otherwise.

### GetTierUpperBoundOk

`func (o *ApiLineItemView) GetTierUpperBoundOk() (*float64, bool)`

GetTierUpperBoundOk returns a tuple with the TierUpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierUpperBound

`func (o *ApiLineItemView) SetTierUpperBound(v float64)`

SetTierUpperBound sets TierUpperBound field to given value.

### HasTierUpperBound

`func (o *ApiLineItemView) HasTierUpperBound() bool`

HasTierUpperBound returns a boolean if a field has been set.

### GetTotalPriceCents

`func (o *ApiLineItemView) GetTotalPriceCents() int64`

GetTotalPriceCents returns the TotalPriceCents field if non-nil, zero value otherwise.

### GetTotalPriceCentsOk

`func (o *ApiLineItemView) GetTotalPriceCentsOk() (*int64, bool)`

GetTotalPriceCentsOk returns a tuple with the TotalPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceCents

`func (o *ApiLineItemView) SetTotalPriceCents(v int64)`

SetTotalPriceCents sets TotalPriceCents field to given value.

### HasTotalPriceCents

`func (o *ApiLineItemView) HasTotalPriceCents() bool`

HasTotalPriceCents returns a boolean if a field has been set.

### GetUnit

`func (o *ApiLineItemView) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ApiLineItemView) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ApiLineItemView) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ApiLineItemView) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitPriceDollars

`func (o *ApiLineItemView) GetUnitPriceDollars() float64`

GetUnitPriceDollars returns the UnitPriceDollars field if non-nil, zero value otherwise.

### GetUnitPriceDollarsOk

`func (o *ApiLineItemView) GetUnitPriceDollarsOk() (*float64, bool)`

GetUnitPriceDollarsOk returns a tuple with the UnitPriceDollars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceDollars

`func (o *ApiLineItemView) SetUnitPriceDollars(v float64)`

SetUnitPriceDollars sets UnitPriceDollars field to given value.

### HasUnitPriceDollars

`func (o *ApiLineItemView) HasUnitPriceDollars() bool`

HasUnitPriceDollars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


