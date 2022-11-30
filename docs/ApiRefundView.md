# ApiRefundView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int64** | Sum of the funds returned to the specified organization expressed in cents (100th of US Dollar). | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this refund. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**PaymentId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the payment that the organization had made. | [optional] [readonly] 
**Reason** | Pointer to **string** | Justification that MongoDB accepted to return funds to the organization. | [optional] [readonly] 

## Methods

### NewApiRefundView

`func NewApiRefundView() *ApiRefundView`

NewApiRefundView instantiates a new ApiRefundView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRefundViewWithDefaults

`func NewApiRefundViewWithDefaults() *ApiRefundView`

NewApiRefundViewWithDefaults instantiates a new ApiRefundView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *ApiRefundView) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *ApiRefundView) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *ApiRefundView) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *ApiRefundView) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCreated

`func (o *ApiRefundView) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiRefundView) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiRefundView) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiRefundView) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetPaymentId

`func (o *ApiRefundView) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *ApiRefundView) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *ApiRefundView) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *ApiRefundView) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetReason

`func (o *ApiRefundView) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiRefundView) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiRefundView) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ApiRefundView) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


