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

// ApiRefundView One payment that MongoDB returned to the organization for this invoice.
type ApiRefundView struct {
	// Sum of the funds returned to the specified organization expressed in cents (100th of US Dollar).
	AmountCents *int64 `json:"amountCents,omitempty"`
	// Date and time when MongoDB Cloud created this refund. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the payment that the organization had made.
	PaymentId *string `json:"paymentId,omitempty"`
	// Justification that MongoDB accepted to return funds to the organization.
	Reason *string `json:"reason,omitempty"`
}

// NewApiRefundView instantiates a new ApiRefundView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRefundView() *ApiRefundView {
	this := ApiRefundView{}
	return &this
}

// NewApiRefundViewWithDefaults instantiates a new ApiRefundView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRefundViewWithDefaults() *ApiRefundView {
	this := ApiRefundView{}
	return &this
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise.
func (o *ApiRefundView) GetAmountCents() int64 {
	if o == nil || isNil(o.AmountCents) {
		var ret int64
		return ret
	}
	return *o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRefundView) GetAmountCentsOk() (*int64, bool) {
	if o == nil || isNil(o.AmountCents) {
    return nil, false
	}
	return o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *ApiRefundView) HasAmountCents() bool {
	if o != nil && !isNil(o.AmountCents) {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given int64 and assigns it to the AmountCents field.
func (o *ApiRefundView) SetAmountCents(v int64) {
	o.AmountCents = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ApiRefundView) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRefundView) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
    return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ApiRefundView) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ApiRefundView) SetCreated(v time.Time) {
	o.Created = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *ApiRefundView) GetPaymentId() string {
	if o == nil || isNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRefundView) GetPaymentIdOk() (*string, bool) {
	if o == nil || isNil(o.PaymentId) {
    return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *ApiRefundView) HasPaymentId() bool {
	if o != nil && !isNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *ApiRefundView) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ApiRefundView) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRefundView) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ApiRefundView) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ApiRefundView) SetReason(v string) {
	o.Reason = &v
}

func (o ApiRefundView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AmountCents) {
		toSerialize["amountCents"] = o.AmountCents
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableApiRefundView struct {
	value *ApiRefundView
	isSet bool
}

func (v NullableApiRefundView) Get() *ApiRefundView {
	return v.value
}

func (v *NullableApiRefundView) Set(val *ApiRefundView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRefundView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRefundView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRefundView(val *ApiRefundView) *NullableApiRefundView {
	return &NullableApiRefundView{value: val, isSet: true}
}

func (v NullableApiRefundView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRefundView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


