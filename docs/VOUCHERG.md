# VOUCHERG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderName** | Pointer to **NullableString** | Order name from client request | [optional] 
**Vouchers** | Pointer to [**[]VOUCHERGSCHEMA**](VOUCHERGSCHEMA.md) |  | [optional] 

## Methods

### NewVOUCHERG

`func NewVOUCHERG() *VOUCHERG`

NewVOUCHERG instantiates a new VOUCHERG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHERGWithDefaults

`func NewVOUCHERGWithDefaults() *VOUCHERG`

NewVOUCHERGWithDefaults instantiates a new VOUCHERG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderName

`func (o *VOUCHERG) GetOrderName() string`

GetOrderName returns the OrderName field if non-nil, zero value otherwise.

### GetOrderNameOk

`func (o *VOUCHERG) GetOrderNameOk() (*string, bool)`

GetOrderNameOk returns a tuple with the OrderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderName

`func (o *VOUCHERG) SetOrderName(v string)`

SetOrderName sets OrderName field to given value.

### HasOrderName

`func (o *VOUCHERG) HasOrderName() bool`

HasOrderName returns a boolean if a field has been set.

### SetOrderNameNil

`func (o *VOUCHERG) SetOrderNameNil(b bool)`

 SetOrderNameNil sets the value for OrderName to be an explicit nil

### UnsetOrderName
`func (o *VOUCHERG) UnsetOrderName()`

UnsetOrderName ensures that no value is present for OrderName, not even an explicit nil
### GetVouchers

`func (o *VOUCHERG) GetVouchers() []VOUCHERGSCHEMA`

GetVouchers returns the Vouchers field if non-nil, zero value otherwise.

### GetVouchersOk

`func (o *VOUCHERG) GetVouchersOk() (*[]VOUCHERGSCHEMA, bool)`

GetVouchersOk returns a tuple with the Vouchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVouchers

`func (o *VOUCHERG) SetVouchers(v []VOUCHERGSCHEMA)`

SetVouchers sets Vouchers field to given value.

### HasVouchers

`func (o *VOUCHERG) HasVouchers() bool`

HasVouchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


