# VOUCHERE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionRefId** | Pointer to **NullableString** | TransactionRefId receive from client request | [optional] 
**Vouchers** | Pointer to [**[]VOUCHERESCHEMA**](VOUCHERESCHEMA.md) |  | [optional] 

## Methods

### NewVOUCHERE

`func NewVOUCHERE() *VOUCHERE`

NewVOUCHERE instantiates a new VOUCHERE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHEREWithDefaults

`func NewVOUCHEREWithDefaults() *VOUCHERE`

NewVOUCHEREWithDefaults instantiates a new VOUCHERE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionRefId

`func (o *VOUCHERE) GetTransactionRefId() string`

GetTransactionRefId returns the TransactionRefId field if non-nil, zero value otherwise.

### GetTransactionRefIdOk

`func (o *VOUCHERE) GetTransactionRefIdOk() (*string, bool)`

GetTransactionRefIdOk returns a tuple with the TransactionRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRefId

`func (o *VOUCHERE) SetTransactionRefId(v string)`

SetTransactionRefId sets TransactionRefId field to given value.

### HasTransactionRefId

`func (o *VOUCHERE) HasTransactionRefId() bool`

HasTransactionRefId returns a boolean if a field has been set.

### SetTransactionRefIdNil

`func (o *VOUCHERE) SetTransactionRefIdNil(b bool)`

 SetTransactionRefIdNil sets the value for TransactionRefId to be an explicit nil

### UnsetTransactionRefId
`func (o *VOUCHERE) UnsetTransactionRefId()`

UnsetTransactionRefId ensures that no value is present for TransactionRefId, not even an explicit nil
### GetVouchers

`func (o *VOUCHERE) GetVouchers() []VOUCHERESCHEMA`

GetVouchers returns the Vouchers field if non-nil, zero value otherwise.

### GetVouchersOk

`func (o *VOUCHERE) GetVouchersOk() (*[]VOUCHERESCHEMA, bool)`

GetVouchersOk returns a tuple with the Vouchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVouchers

`func (o *VOUCHERE) SetVouchers(v []VOUCHERESCHEMA)`

SetVouchers sets Vouchers field to given value.

### HasVouchers

`func (o *VOUCHERE) HasVouchers() bool`

HasVouchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


