/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ApiLineItemView One service included in this invoice.
type ApiLineItemView struct {
	// Human-readable label that identifies the cluster that incurred the charge.
	ClusterName *string `json:"clusterName,omitempty"`
	// Date and time when MongoDB Cloud created this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Sum by which MongoDB discounted this line item. MongoDB Cloud expresses this value in cents (100ths of one US Dollar). The resource returns this parameter when a discount applies.
	DiscountCents *int64 `json:"discountCents,omitempty"`
	// Date and time when when MongoDB Cloud finished charging for this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	EndDate *time.Time `json:"endDate,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project associated to this line item.
	GroupId *string `json:"groupId,omitempty"`
	// Human-readable label that identifies the project.
	GroupName *string `json:"groupName,omitempty"`
	// Comment that applies to this line item.
	Note *string `json:"note,omitempty"`
	// Percentage by which MongoDB discounted this line item. The resource returns this parameter when a discount applies.
	PercentDiscount *float32 `json:"percentDiscount,omitempty"`
	// Number of units included for the line item. These can be expressions of storage (GB), time (hours), or other units.
	Quantity *float64 `json:"quantity,omitempty"`
	// Human-readable description of the service that this line item provided. This Stock Keeping Unit (SKU) could be the instance type, a support charge, advanced security, or another service.
	Sku *string `json:"sku,omitempty"`
	// Date and time when MongoDB Cloud began charging for this line item. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Human-readable label that identifies the Atlas App Service associated with this line item.
	StitchAppName *string `json:"stitchAppName,omitempty"`
	// Lower bound for usage amount range in current SKU tier.   **NOTE**: **lineItems[n].tierLowerBound** appears only if your **lineItems[n].sku** is tiered.
	TierLowerBound *float64 `json:"tierLowerBound,omitempty"`
	// Upper bound for usage amount range in current SKU tier.   **NOTE**: **lineItems[n].tierUpperBound** appears only if your **lineItems[n].sku** is tiered.
	TierUpperBound *float64 `json:"tierUpperBound,omitempty"`
	// Sum of the cost set for this line item. MongoDB Cloud expresses this value in cents (100ths of one US Dollar) and calculates this value as **unitPriceDollars** × **quantity** × 100.
	TotalPriceCents *int64 `json:"totalPriceCents,omitempty"`
	// Element used to express what **quantity** this line item measures. This value can be elements of time, storage capacity, and the like.
	Unit *string `json:"unit,omitempty"`
	// Value per **unit** for this line item expressed in US Dollars.
	UnitPriceDollars *float64 `json:"unitPriceDollars,omitempty"`
}

// NewApiLineItemView instantiates a new ApiLineItemView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLineItemView() *ApiLineItemView {
	this := ApiLineItemView{}
	return &this
}

// NewApiLineItemViewWithDefaults instantiates a new ApiLineItemView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLineItemViewWithDefaults() *ApiLineItemView {
	this := ApiLineItemView{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *ApiLineItemView) GetClusterName() string {
	if o == nil || isNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetClusterNameOk() (*string, bool) {
	if o == nil || isNil(o.ClusterName) {
    return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ApiLineItemView) HasClusterName() bool {
	if o != nil && !isNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ApiLineItemView) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ApiLineItemView) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
    return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ApiLineItemView) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ApiLineItemView) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDiscountCents returns the DiscountCents field value if set, zero value otherwise.
func (o *ApiLineItemView) GetDiscountCents() int64 {
	if o == nil || isNil(o.DiscountCents) {
		var ret int64
		return ret
	}
	return *o.DiscountCents
}

// GetDiscountCentsOk returns a tuple with the DiscountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetDiscountCentsOk() (*int64, bool) {
	if o == nil || isNil(o.DiscountCents) {
    return nil, false
	}
	return o.DiscountCents, true
}

// HasDiscountCents returns a boolean if a field has been set.
func (o *ApiLineItemView) HasDiscountCents() bool {
	if o != nil && !isNil(o.DiscountCents) {
		return true
	}

	return false
}

// SetDiscountCents gets a reference to the given int64 and assigns it to the DiscountCents field.
func (o *ApiLineItemView) SetDiscountCents(v int64) {
	o.DiscountCents = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ApiLineItemView) GetEndDate() time.Time {
	if o == nil || isNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetEndDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndDate) {
    return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ApiLineItemView) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ApiLineItemView) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApiLineItemView) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiLineItemView) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiLineItemView) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ApiLineItemView) GetGroupName() string {
	if o == nil || isNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetGroupNameOk() (*string, bool) {
	if o == nil || isNil(o.GroupName) {
    return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ApiLineItemView) HasGroupName() bool {
	if o != nil && !isNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ApiLineItemView) SetGroupName(v string) {
	o.GroupName = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *ApiLineItemView) GetNote() string {
	if o == nil || isNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetNoteOk() (*string, bool) {
	if o == nil || isNil(o.Note) {
    return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *ApiLineItemView) HasNote() bool {
	if o != nil && !isNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *ApiLineItemView) SetNote(v string) {
	o.Note = &v
}

// GetPercentDiscount returns the PercentDiscount field value if set, zero value otherwise.
func (o *ApiLineItemView) GetPercentDiscount() float32 {
	if o == nil || isNil(o.PercentDiscount) {
		var ret float32
		return ret
	}
	return *o.PercentDiscount
}

// GetPercentDiscountOk returns a tuple with the PercentDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetPercentDiscountOk() (*float32, bool) {
	if o == nil || isNil(o.PercentDiscount) {
    return nil, false
	}
	return o.PercentDiscount, true
}

// HasPercentDiscount returns a boolean if a field has been set.
func (o *ApiLineItemView) HasPercentDiscount() bool {
	if o != nil && !isNil(o.PercentDiscount) {
		return true
	}

	return false
}

// SetPercentDiscount gets a reference to the given float32 and assigns it to the PercentDiscount field.
func (o *ApiLineItemView) SetPercentDiscount(v float32) {
	o.PercentDiscount = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ApiLineItemView) GetQuantity() float64 {
	if o == nil || isNil(o.Quantity) {
		var ret float64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetQuantityOk() (*float64, bool) {
	if o == nil || isNil(o.Quantity) {
    return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ApiLineItemView) HasQuantity() bool {
	if o != nil && !isNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float64 and assigns it to the Quantity field.
func (o *ApiLineItemView) SetQuantity(v float64) {
	o.Quantity = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ApiLineItemView) GetSku() string {
	if o == nil || isNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetSkuOk() (*string, bool) {
	if o == nil || isNil(o.Sku) {
    return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ApiLineItemView) HasSku() bool {
	if o != nil && !isNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ApiLineItemView) SetSku(v string) {
	o.Sku = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ApiLineItemView) GetStartDate() time.Time {
	if o == nil || isNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetStartDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartDate) {
    return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ApiLineItemView) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ApiLineItemView) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetStitchAppName returns the StitchAppName field value if set, zero value otherwise.
func (o *ApiLineItemView) GetStitchAppName() string {
	if o == nil || isNil(o.StitchAppName) {
		var ret string
		return ret
	}
	return *o.StitchAppName
}

// GetStitchAppNameOk returns a tuple with the StitchAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetStitchAppNameOk() (*string, bool) {
	if o == nil || isNil(o.StitchAppName) {
    return nil, false
	}
	return o.StitchAppName, true
}

// HasStitchAppName returns a boolean if a field has been set.
func (o *ApiLineItemView) HasStitchAppName() bool {
	if o != nil && !isNil(o.StitchAppName) {
		return true
	}

	return false
}

// SetStitchAppName gets a reference to the given string and assigns it to the StitchAppName field.
func (o *ApiLineItemView) SetStitchAppName(v string) {
	o.StitchAppName = &v
}

// GetTierLowerBound returns the TierLowerBound field value if set, zero value otherwise.
func (o *ApiLineItemView) GetTierLowerBound() float64 {
	if o == nil || isNil(o.TierLowerBound) {
		var ret float64
		return ret
	}
	return *o.TierLowerBound
}

// GetTierLowerBoundOk returns a tuple with the TierLowerBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetTierLowerBoundOk() (*float64, bool) {
	if o == nil || isNil(o.TierLowerBound) {
    return nil, false
	}
	return o.TierLowerBound, true
}

// HasTierLowerBound returns a boolean if a field has been set.
func (o *ApiLineItemView) HasTierLowerBound() bool {
	if o != nil && !isNil(o.TierLowerBound) {
		return true
	}

	return false
}

// SetTierLowerBound gets a reference to the given float64 and assigns it to the TierLowerBound field.
func (o *ApiLineItemView) SetTierLowerBound(v float64) {
	o.TierLowerBound = &v
}

// GetTierUpperBound returns the TierUpperBound field value if set, zero value otherwise.
func (o *ApiLineItemView) GetTierUpperBound() float64 {
	if o == nil || isNil(o.TierUpperBound) {
		var ret float64
		return ret
	}
	return *o.TierUpperBound
}

// GetTierUpperBoundOk returns a tuple with the TierUpperBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetTierUpperBoundOk() (*float64, bool) {
	if o == nil || isNil(o.TierUpperBound) {
    return nil, false
	}
	return o.TierUpperBound, true
}

// HasTierUpperBound returns a boolean if a field has been set.
func (o *ApiLineItemView) HasTierUpperBound() bool {
	if o != nil && !isNil(o.TierUpperBound) {
		return true
	}

	return false
}

// SetTierUpperBound gets a reference to the given float64 and assigns it to the TierUpperBound field.
func (o *ApiLineItemView) SetTierUpperBound(v float64) {
	o.TierUpperBound = &v
}

// GetTotalPriceCents returns the TotalPriceCents field value if set, zero value otherwise.
func (o *ApiLineItemView) GetTotalPriceCents() int64 {
	if o == nil || isNil(o.TotalPriceCents) {
		var ret int64
		return ret
	}
	return *o.TotalPriceCents
}

// GetTotalPriceCentsOk returns a tuple with the TotalPriceCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetTotalPriceCentsOk() (*int64, bool) {
	if o == nil || isNil(o.TotalPriceCents) {
    return nil, false
	}
	return o.TotalPriceCents, true
}

// HasTotalPriceCents returns a boolean if a field has been set.
func (o *ApiLineItemView) HasTotalPriceCents() bool {
	if o != nil && !isNil(o.TotalPriceCents) {
		return true
	}

	return false
}

// SetTotalPriceCents gets a reference to the given int64 and assigns it to the TotalPriceCents field.
func (o *ApiLineItemView) SetTotalPriceCents(v int64) {
	o.TotalPriceCents = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ApiLineItemView) GetUnit() string {
	if o == nil || isNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetUnitOk() (*string, bool) {
	if o == nil || isNil(o.Unit) {
    return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ApiLineItemView) HasUnit() bool {
	if o != nil && !isNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *ApiLineItemView) SetUnit(v string) {
	o.Unit = &v
}

// GetUnitPriceDollars returns the UnitPriceDollars field value if set, zero value otherwise.
func (o *ApiLineItemView) GetUnitPriceDollars() float64 {
	if o == nil || isNil(o.UnitPriceDollars) {
		var ret float64
		return ret
	}
	return *o.UnitPriceDollars
}

// GetUnitPriceDollarsOk returns a tuple with the UnitPriceDollars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLineItemView) GetUnitPriceDollarsOk() (*float64, bool) {
	if o == nil || isNil(o.UnitPriceDollars) {
    return nil, false
	}
	return o.UnitPriceDollars, true
}

// HasUnitPriceDollars returns a boolean if a field has been set.
func (o *ApiLineItemView) HasUnitPriceDollars() bool {
	if o != nil && !isNil(o.UnitPriceDollars) {
		return true
	}

	return false
}

// SetUnitPriceDollars gets a reference to the given float64 and assigns it to the UnitPriceDollars field.
func (o *ApiLineItemView) SetUnitPriceDollars(v float64) {
	o.UnitPriceDollars = &v
}

func (o ApiLineItemView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClusterName) {
		toSerialize["clusterName"] = o.ClusterName
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.DiscountCents) {
		toSerialize["discountCents"] = o.DiscountCents
	}
	if !isNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !isNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !isNil(o.PercentDiscount) {
		toSerialize["percentDiscount"] = o.PercentDiscount
	}
	if !isNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !isNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !isNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !isNil(o.StitchAppName) {
		toSerialize["stitchAppName"] = o.StitchAppName
	}
	if !isNil(o.TierLowerBound) {
		toSerialize["tierLowerBound"] = o.TierLowerBound
	}
	if !isNil(o.TierUpperBound) {
		toSerialize["tierUpperBound"] = o.TierUpperBound
	}
	if !isNil(o.TotalPriceCents) {
		toSerialize["totalPriceCents"] = o.TotalPriceCents
	}
	if !isNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !isNil(o.UnitPriceDollars) {
		toSerialize["unitPriceDollars"] = o.UnitPriceDollars
	}
	return json.Marshal(toSerialize)
}

type NullableApiLineItemView struct {
	value *ApiLineItemView
	isSet bool
}

func (v NullableApiLineItemView) Get() *ApiLineItemView {
	return v.value
}

func (v *NullableApiLineItemView) Set(val *ApiLineItemView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLineItemView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLineItemView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLineItemView(val *ApiLineItemView) *NullableApiLineItemView {
	return &NullableApiLineItemView{value: val, isSet: true}
}

func (v NullableApiLineItemView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLineItemView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


