# VOUCHERGSCHEMA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **NullableInt64** | Order Id | [optional] 
**OrderName** | Pointer to **NullableString** | Order name from client request | [optional] 
**GroupVouchers** | Pointer to [**GROUPVOUCHERSCHEMA**](GROUPVOUCHERSCHEMA.md) |  | [optional] 
**Vouchers** | Pointer to [**[]PRODUCTG**](PRODUCTG.md) |  | [optional] 

## Methods

### NewVOUCHERGSCHEMA

`func NewVOUCHERGSCHEMA() *VOUCHERGSCHEMA`

NewVOUCHERGSCHEMA instantiates a new VOUCHERGSCHEMA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHERGSCHEMAWithDefaults

`func NewVOUCHERGSCHEMAWithDefaults() *VOUCHERGSCHEMA`

NewVOUCHERGSCHEMAWithDefaults instantiates a new VOUCHERGSCHEMA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *VOUCHERGSCHEMA) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *VOUCHERGSCHEMA) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *VOUCHERGSCHEMA) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *VOUCHERGSCHEMA) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *VOUCHERGSCHEMA) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *VOUCHERGSCHEMA) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetOrderName

`func (o *VOUCHERGSCHEMA) GetOrderName() string`

GetOrderName returns the OrderName field if non-nil, zero value otherwise.

### GetOrderNameOk

`func (o *VOUCHERGSCHEMA) GetOrderNameOk() (*string, bool)`

GetOrderNameOk returns a tuple with the OrderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderName

`func (o *VOUCHERGSCHEMA) SetOrderName(v string)`

SetOrderName sets OrderName field to given value.

### HasOrderName

`func (o *VOUCHERGSCHEMA) HasOrderName() bool`

HasOrderName returns a boolean if a field has been set.

### SetOrderNameNil

`func (o *VOUCHERGSCHEMA) SetOrderNameNil(b bool)`

 SetOrderNameNil sets the value for OrderName to be an explicit nil

### UnsetOrderName
`func (o *VOUCHERGSCHEMA) UnsetOrderName()`

UnsetOrderName ensures that no value is present for OrderName, not even an explicit nil
### GetGroupVouchers

`func (o *VOUCHERGSCHEMA) GetGroupVouchers() GROUPVOUCHERSCHEMA`

GetGroupVouchers returns the GroupVouchers field if non-nil, zero value otherwise.

### GetGroupVouchersOk

`func (o *VOUCHERGSCHEMA) GetGroupVouchersOk() (*GROUPVOUCHERSCHEMA, bool)`

GetGroupVouchersOk returns a tuple with the GroupVouchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupVouchers

`func (o *VOUCHERGSCHEMA) SetGroupVouchers(v GROUPVOUCHERSCHEMA)`

SetGroupVouchers sets GroupVouchers field to given value.

### HasGroupVouchers

`func (o *VOUCHERGSCHEMA) HasGroupVouchers() bool`

HasGroupVouchers returns a boolean if a field has been set.

### GetVouchers

`func (o *VOUCHERGSCHEMA) GetVouchers() []PRODUCTG`

GetVouchers returns the Vouchers field if non-nil, zero value otherwise.

### GetVouchersOk

`func (o *VOUCHERGSCHEMA) GetVouchersOk() (*[]PRODUCTG, bool)`

GetVouchersOk returns a tuple with the Vouchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVouchers

`func (o *VOUCHERGSCHEMA) SetVouchers(v []PRODUCTG)`

SetVouchers sets Vouchers field to given value.

### HasVouchers

`func (o *VOUCHERGSCHEMA) HasVouchers() bool`

HasVouchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


