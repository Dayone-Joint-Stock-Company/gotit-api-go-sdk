# VOUCHERV

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderName** | Pointer to **NullableString** | Order name from client request | [optional] 
**Vouchers** | Pointer to [**[]VOUCHERVSCHEMA**](VOUCHERVSCHEMA.md) |  | [optional] 

## Methods

### NewVOUCHERV

`func NewVOUCHERV() *VOUCHERV`

NewVOUCHERV instantiates a new VOUCHERV object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHERVWithDefaults

`func NewVOUCHERVWithDefaults() *VOUCHERV`

NewVOUCHERVWithDefaults instantiates a new VOUCHERV object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderName

`func (o *VOUCHERV) GetOrderName() string`

GetOrderName returns the OrderName field if non-nil, zero value otherwise.

### GetOrderNameOk

`func (o *VOUCHERV) GetOrderNameOk() (*string, bool)`

GetOrderNameOk returns a tuple with the OrderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderName

`func (o *VOUCHERV) SetOrderName(v string)`

SetOrderName sets OrderName field to given value.

### HasOrderName

`func (o *VOUCHERV) HasOrderName() bool`

HasOrderName returns a boolean if a field has been set.

### SetOrderNameNil

`func (o *VOUCHERV) SetOrderNameNil(b bool)`

 SetOrderNameNil sets the value for OrderName to be an explicit nil

### UnsetOrderName
`func (o *VOUCHERV) UnsetOrderName()`

UnsetOrderName ensures that no value is present for OrderName, not even an explicit nil
### GetVouchers

`func (o *VOUCHERV) GetVouchers() []VOUCHERVSCHEMA`

GetVouchers returns the Vouchers field if non-nil, zero value otherwise.

### GetVouchersOk

`func (o *VOUCHERV) GetVouchersOk() (*[]VOUCHERVSCHEMA, bool)`

GetVouchersOk returns a tuple with the Vouchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVouchers

`func (o *VOUCHERV) SetVouchers(v []VOUCHERVSCHEMA)`

SetVouchers sets Vouchers field to given value.

### HasVouchers

`func (o *VOUCHERV) HasVouchers() bool`

HasVouchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


