# ApiInvoiceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountBilledCents** | Pointer to **int64** | Sum of services that the specified organization consumed in the period covered in this invoice. This parameter expresses its value in cents (100ths of one US Dollar) and calculates its value as **subtotalCents** + **salesTaxCents** - **startingBalanceCents**. | [optional] [readonly] 
**AmountPaidCents** | Pointer to **int64** | Sum that the specified organization paid toward this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when MongoDB Cloud created this invoice. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**CreditsCents** | Pointer to **int64** | Sum that MongoDB credited the specified organization toward this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud finished the billing period that this invoice covers. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project associated to this invoice. This identifying string doesn&#39;t appear on all invoices. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | [optional] [readonly] 
**LineItems** | Pointer to [**[]ApiLineItemView**](ApiLineItemView.md) | List that contains individual services included in this invoice. | [optional] [readonly] 
**LinkedInvoices** | Pointer to [**[]ApiInvoiceView**](ApiInvoiceView.md) | List that contains the invoices for organizations linked to the paying organization. | [optional] [readonly] 
**Links** | [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization charged for services consumed from MongoDB Cloud. | [optional] [readonly] 
**Payments** | Pointer to [**[]ApiPaymentView**](ApiPaymentView.md) | List that contains funds transferred to MongoDB to cover the specified service noted in this invoice. | [optional] [readonly] 
**Refunds** | Pointer to [**[]ApiRefundView**](ApiRefundView.md) | List that contains payments that MongoDB returned to the organization for this invoice. | [optional] [readonly] 
**SalesTaxCents** | Pointer to **int64** | Sum of sales tax applied to this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud began the billing period that this invoice covers. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**StartingBalanceCents** | Pointer to **int64** | Sum that the specified organization owed to MongoDB when MongoDB issued this invoice. This parameter expresses its value in US Dollars. | [optional] [readonly] 
**StatusName** | Pointer to **string** | Phase of payment processing in which this invoice exists when you made this request. Accepted phases include:  | Phase Value | Reason | |---|---| | CLOSED | MongoDB finalized all charges in the billing cycle but has yet to charge the customer. | | FAILED | MongoDB attempted to charge the provided credit card but charge for that amount failed. | | FORGIVEN | Customer initiated payment which MongoDB later forgave. | | FREE | All charges totalled zero so the customer won&#39;t be charged. | | INVOICED | MongoDB handled these charges using elastic invoicing. | | PAID | MongoDB succeeded in charging the provided credit card. | | PENDING | Invoice includes charges for the current billing cycle. | | PREPAID | Customer has a pre-paid plan so they won&#39;t be charged. |  | [optional] 
**SubtotalCents** | Pointer to **int64** | Sum of all positive invoice line items contained in this invoice. This parameter expresses its value in cents (100ths of one US Dollar). | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | Date and time when MongoDB Cloud last updated the value of this payment. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 

## Methods

### NewApiInvoiceView

`func NewApiInvoiceView(links []Link, ) *ApiInvoiceView`

NewApiInvoiceView instantiates a new ApiInvoiceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInvoiceViewWithDefaults

`func NewApiInvoiceViewWithDefaults() *ApiInvoiceView`

NewApiInvoiceViewWithDefaults instantiates a new ApiInvoiceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountBilledCents

`func (o *ApiInvoiceView) GetAmountBilledCents() int64`

GetAmountBilledCents returns the AmountBilledCents field if non-nil, zero value otherwise.

### GetAmountBilledCentsOk

`func (o *ApiInvoiceView) GetAmountBilledCentsOk() (*int64, bool)`

GetAmountBilledCentsOk returns a tuple with the AmountBilledCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBilledCents

`func (o *ApiInvoiceView) SetAmountBilledCents(v int64)`

SetAmountBilledCents sets AmountBilledCents field to given value.

### HasAmountBilledCents

`func (o *ApiInvoiceView) HasAmountBilledCents() bool`

HasAmountBilledCents returns a boolean if a field has been set.

### GetAmountPaidCents

`func (o *ApiInvoiceView) GetAmountPaidCents() int64`

GetAmountPaidCents returns the AmountPaidCents field if non-nil, zero value otherwise.

### GetAmountPaidCentsOk

`func (o *ApiInvoiceView) GetAmountPaidCentsOk() (*int64, bool)`

GetAmountPaidCentsOk returns a tuple with the AmountPaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaidCents

`func (o *ApiInvoiceView) SetAmountPaidCents(v int64)`

SetAmountPaidCents sets AmountPaidCents field to given value.

### HasAmountPaidCents

`func (o *ApiInvoiceView) HasAmountPaidCents() bool`

HasAmountPaidCents returns a boolean if a field has been set.

### GetCreated

`func (o *ApiInvoiceView) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiInvoiceView) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiInvoiceView) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiInvoiceView) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreditsCents

`func (o *ApiInvoiceView) GetCreditsCents() int64`

GetCreditsCents returns the CreditsCents field if non-nil, zero value otherwise.

### GetCreditsCentsOk

`func (o *ApiInvoiceView) GetCreditsCentsOk() (*int64, bool)`

GetCreditsCentsOk returns a tuple with the CreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCents

`func (o *ApiInvoiceView) SetCreditsCents(v int64)`

SetCreditsCents sets CreditsCents field to given value.

### HasCreditsCents

`func (o *ApiInvoiceView) HasCreditsCents() bool`

HasCreditsCents returns a boolean if a field has been set.

### GetEndDate

`func (o *ApiInvoiceView) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ApiInvoiceView) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ApiInvoiceView) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ApiInvoiceView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiInvoiceView) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiInvoiceView) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiInvoiceView) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ApiInvoiceView) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *ApiInvoiceView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiInvoiceView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiInvoiceView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiInvoiceView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLineItems

`func (o *ApiInvoiceView) GetLineItems() []ApiLineItemView`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *ApiInvoiceView) GetLineItemsOk() (*[]ApiLineItemView, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *ApiInvoiceView) SetLineItems(v []ApiLineItemView)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *ApiInvoiceView) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetLinkedInvoices

`func (o *ApiInvoiceView) GetLinkedInvoices() []ApiInvoiceView`

GetLinkedInvoices returns the LinkedInvoices field if non-nil, zero value otherwise.

### GetLinkedInvoicesOk

`func (o *ApiInvoiceView) GetLinkedInvoicesOk() (*[]ApiInvoiceView, bool)`

GetLinkedInvoicesOk returns a tuple with the LinkedInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInvoices

`func (o *ApiInvoiceView) SetLinkedInvoices(v []ApiInvoiceView)`

SetLinkedInvoices sets LinkedInvoices field to given value.

### HasLinkedInvoices

`func (o *ApiInvoiceView) HasLinkedInvoices() bool`

HasLinkedInvoices returns a boolean if a field has been set.

### GetLinks

`func (o *ApiInvoiceView) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiInvoiceView) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiInvoiceView) SetLinks(v []Link)`

SetLinks sets Links field to given value.


### GetOrgId

`func (o *ApiInvoiceView) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ApiInvoiceView) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ApiInvoiceView) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ApiInvoiceView) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPayments

`func (o *ApiInvoiceView) GetPayments() []ApiPaymentView`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *ApiInvoiceView) GetPaymentsOk() (*[]ApiPaymentView, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *ApiInvoiceView) SetPayments(v []ApiPaymentView)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *ApiInvoiceView) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetRefunds

`func (o *ApiInvoiceView) GetRefunds() []ApiRefundView`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *ApiInvoiceView) GetRefundsOk() (*[]ApiRefundView, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *ApiInvoiceView) SetRefunds(v []ApiRefundView)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *ApiInvoiceView) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetSalesTaxCents

`func (o *ApiInvoiceView) GetSalesTaxCents() int64`

GetSalesTaxCents returns the SalesTaxCents field if non-nil, zero value otherwise.

### GetSalesTaxCentsOk

`func (o *ApiInvoiceView) GetSalesTaxCentsOk() (*int64, bool)`

GetSalesTaxCentsOk returns a tuple with the SalesTaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCents

`func (o *ApiInvoiceView) SetSalesTaxCents(v int64)`

SetSalesTaxCents sets SalesTaxCents field to given value.

### HasSalesTaxCents

`func (o *ApiInvoiceView) HasSalesTaxCents() bool`

HasSalesTaxCents returns a boolean if a field has been set.

### GetStartDate

`func (o *ApiInvoiceView) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ApiInvoiceView) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ApiInvoiceView) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ApiInvoiceView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartingBalanceCents

`func (o *ApiInvoiceView) GetStartingBalanceCents() int64`

GetStartingBalanceCents returns the StartingBalanceCents field if non-nil, zero value otherwise.

### GetStartingBalanceCentsOk

`func (o *ApiInvoiceView) GetStartingBalanceCentsOk() (*int64, bool)`

GetStartingBalanceCentsOk returns a tuple with the StartingBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingBalanceCents

`func (o *ApiInvoiceView) SetStartingBalanceCents(v int64)`

SetStartingBalanceCents sets StartingBalanceCents field to given value.

### HasStartingBalanceCents

`func (o *ApiInvoiceView) HasStartingBalanceCents() bool`

HasStartingBalanceCents returns a boolean if a field has been set.

### GetStatusName

`func (o *ApiInvoiceView) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *ApiInvoiceView) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *ApiInvoiceView) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *ApiInvoiceView) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetSubtotalCents

`func (o *ApiInvoiceView) GetSubtotalCents() int64`

GetSubtotalCents returns the SubtotalCents field if non-nil, zero value otherwise.

### GetSubtotalCentsOk

`func (o *ApiInvoiceView) GetSubtotalCentsOk() (*int64, bool)`

GetSubtotalCentsOk returns a tuple with the SubtotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalCents

`func (o *ApiInvoiceView) SetSubtotalCents(v int64)`

SetSubtotalCents sets SubtotalCents field to given value.

### HasSubtotalCents

`func (o *ApiInvoiceView) HasSubtotalCents() bool`

HasSubtotalCents returns a boolean if a field has been set.

### GetUpdated

`func (o *ApiInvoiceView) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ApiInvoiceView) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ApiInvoiceView) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ApiInvoiceView) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


