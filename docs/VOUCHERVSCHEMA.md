# VOUCHERVSCHEMA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionRefId** | Pointer to **NullableString** | TransactionRefId receive from client request | [optional] 
**VoucherCode** | Pointer to **NullableString** | Voucher code. Contains 10 numbers, unique and not duplicated with any other voucher. For some brands that do not use Got It code but use integerernal code (Big C, Mobile World), the system returns the integerernal code. | [optional] 
**VoucherLink** | Pointer to **NullableString** | Link of voucher | [optional] 
**VoucherLinkCode** | Pointer to **NullableString** | Last 8 characters of voucher link | [optional] 
**VoucherImageLink** | Pointer to **NullableString** | Link of voucher image. To display vouchers as images | [optional] 
**VoucherCoverLink** | Pointer to **NullableString** | Cover link of voucher. | [optional] 
**VoucherCoverLinkCode** | Pointer to **NullableString** | Last 8 characters of voucher cover link | [optional] 
**VoucherSerial** | Pointer to **NullableString** | Is a unique code to identify voucher link v and is not valid for use. For example: V1562342ET2 | [optional] 
**ExpiryDate** | Pointer to **NullableString** | Voucher expiration date | [optional] 
**Product** | Pointer to [**NullableVOUCHERVSCHEMAProduct**](VOUCHERVSCHEMAProduct.md) |  | [optional] 

## Methods

### NewVOUCHERVSCHEMA

`func NewVOUCHERVSCHEMA() *VOUCHERVSCHEMA`

NewVOUCHERVSCHEMA instantiates a new VOUCHERVSCHEMA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHERVSCHEMAWithDefaults

`func NewVOUCHERVSCHEMAWithDefaults() *VOUCHERVSCHEMA`

NewVOUCHERVSCHEMAWithDefaults instantiates a new VOUCHERVSCHEMA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionRefId

`func (o *VOUCHERVSCHEMA) GetTransactionRefId() string`

GetTransactionRefId returns the TransactionRefId field if non-nil, zero value otherwise.

### GetTransactionRefIdOk

`func (o *VOUCHERVSCHEMA) GetTransactionRefIdOk() (*string, bool)`

GetTransactionRefIdOk returns a tuple with the TransactionRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRefId

`func (o *VOUCHERVSCHEMA) SetTransactionRefId(v string)`

SetTransactionRefId sets TransactionRefId field to given value.

### HasTransactionRefId

`func (o *VOUCHERVSCHEMA) HasTransactionRefId() bool`

HasTransactionRefId returns a boolean if a field has been set.

### SetTransactionRefIdNil

`func (o *VOUCHERVSCHEMA) SetTransactionRefIdNil(b bool)`

 SetTransactionRefIdNil sets the value for TransactionRefId to be an explicit nil

### UnsetTransactionRefId
`func (o *VOUCHERVSCHEMA) UnsetTransactionRefId()`

UnsetTransactionRefId ensures that no value is present for TransactionRefId, not even an explicit nil
### GetVoucherCode

`func (o *VOUCHERVSCHEMA) GetVoucherCode() string`

GetVoucherCode returns the VoucherCode field if non-nil, zero value otherwise.

### GetVoucherCodeOk

`func (o *VOUCHERVSCHEMA) GetVoucherCodeOk() (*string, bool)`

GetVoucherCodeOk returns a tuple with the VoucherCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherCode

`func (o *VOUCHERVSCHEMA) SetVoucherCode(v string)`

SetVoucherCode sets VoucherCode field to given value.

### HasVoucherCode

`func (o *VOUCHERVSCHEMA) HasVoucherCode() bool`

HasVoucherCode returns a boolean if a field has been set.

### SetVoucherCodeNil

`func (o *VOUCHERVSCHEMA) SetVoucherCodeNil(b bool)`

 SetVoucherCodeNil sets the value for VoucherCode to be an explicit nil

### UnsetVoucherCode
`func (o *VOUCHERVSCHEMA) UnsetVoucherCode()`

UnsetVoucherCode ensures that no value is present for VoucherCode, not even an explicit nil
### GetVoucherLink

`func (o *VOUCHERVSCHEMA) GetVoucherLink() string`

GetVoucherLink returns the VoucherLink field if non-nil, zero value otherwise.

### GetVoucherLinkOk

`func (o *VOUCHERVSCHEMA) GetVoucherLinkOk() (*string, bool)`

GetVoucherLinkOk returns a tuple with the VoucherLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherLink

`func (o *VOUCHERVSCHEMA) SetVoucherLink(v string)`

SetVoucherLink sets VoucherLink field to given value.

### HasVoucherLink

`func (o *VOUCHERVSCHEMA) HasVoucherLink() bool`

HasVoucherLink returns a boolean if a field has been set.

### SetVoucherLinkNil

`func (o *VOUCHERVSCHEMA) SetVoucherLinkNil(b bool)`

 SetVoucherLinkNil sets the value for VoucherLink to be an explicit nil

### UnsetVoucherLink
`func (o *VOUCHERVSCHEMA) UnsetVoucherLink()`

UnsetVoucherLink ensures that no value is present for VoucherLink, not even an explicit nil
### GetVoucherLinkCode

`func (o *VOUCHERVSCHEMA) GetVoucherLinkCode() string`

GetVoucherLinkCode returns the VoucherLinkCode field if non-nil, zero value otherwise.

### GetVoucherLinkCodeOk

`func (o *VOUCHERVSCHEMA) GetVoucherLinkCodeOk() (*string, bool)`

GetVoucherLinkCodeOk returns a tuple with the VoucherLinkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherLinkCode

`func (o *VOUCHERVSCHEMA) SetVoucherLinkCode(v string)`

SetVoucherLinkCode sets VoucherLinkCode field to given value.

### HasVoucherLinkCode

`func (o *VOUCHERVSCHEMA) HasVoucherLinkCode() bool`

HasVoucherLinkCode returns a boolean if a field has been set.

### SetVoucherLinkCodeNil

`func (o *VOUCHERVSCHEMA) SetVoucherLinkCodeNil(b bool)`

 SetVoucherLinkCodeNil sets the value for VoucherLinkCode to be an explicit nil

### UnsetVoucherLinkCode
`func (o *VOUCHERVSCHEMA) UnsetVoucherLinkCode()`

UnsetVoucherLinkCode ensures that no value is present for VoucherLinkCode, not even an explicit nil
### GetVoucherImageLink

`func (o *VOUCHERVSCHEMA) GetVoucherImageLink() string`

GetVoucherImageLink returns the VoucherImageLink field if non-nil, zero value otherwise.

### GetVoucherImageLinkOk

`func (o *VOUCHERVSCHEMA) GetVoucherImageLinkOk() (*string, bool)`

GetVoucherImageLinkOk returns a tuple with the VoucherImageLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherImageLink

`func (o *VOUCHERVSCHEMA) SetVoucherImageLink(v string)`

SetVoucherImageLink sets VoucherImageLink field to given value.

### HasVoucherImageLink

`func (o *VOUCHERVSCHEMA) HasVoucherImageLink() bool`

HasVoucherImageLink returns a boolean if a field has been set.

### SetVoucherImageLinkNil

`func (o *VOUCHERVSCHEMA) SetVoucherImageLinkNil(b bool)`

 SetVoucherImageLinkNil sets the value for VoucherImageLink to be an explicit nil

### UnsetVoucherImageLink
`func (o *VOUCHERVSCHEMA) UnsetVoucherImageLink()`

UnsetVoucherImageLink ensures that no value is present for VoucherImageLink, not even an explicit nil
### GetVoucherCoverLink

`func (o *VOUCHERVSCHEMA) GetVoucherCoverLink() string`

GetVoucherCoverLink returns the VoucherCoverLink field if non-nil, zero value otherwise.

### GetVoucherCoverLinkOk

`func (o *VOUCHERVSCHEMA) GetVoucherCoverLinkOk() (*string, bool)`

GetVoucherCoverLinkOk returns a tuple with the VoucherCoverLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherCoverLink

`func (o *VOUCHERVSCHEMA) SetVoucherCoverLink(v string)`

SetVoucherCoverLink sets VoucherCoverLink field to given value.

### HasVoucherCoverLink

`func (o *VOUCHERVSCHEMA) HasVoucherCoverLink() bool`

HasVoucherCoverLink returns a boolean if a field has been set.

### SetVoucherCoverLinkNil

`func (o *VOUCHERVSCHEMA) SetVoucherCoverLinkNil(b bool)`

 SetVoucherCoverLinkNil sets the value for VoucherCoverLink to be an explicit nil

### UnsetVoucherCoverLink
`func (o *VOUCHERVSCHEMA) UnsetVoucherCoverLink()`

UnsetVoucherCoverLink ensures that no value is present for VoucherCoverLink, not even an explicit nil
### GetVoucherCoverLinkCode

`func (o *VOUCHERVSCHEMA) GetVoucherCoverLinkCode() string`

GetVoucherCoverLinkCode returns the VoucherCoverLinkCode field if non-nil, zero value otherwise.

### GetVoucherCoverLinkCodeOk

`func (o *VOUCHERVSCHEMA) GetVoucherCoverLinkCodeOk() (*string, bool)`

GetVoucherCoverLinkCodeOk returns a tuple with the VoucherCoverLinkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherCoverLinkCode

`func (o *VOUCHERVSCHEMA) SetVoucherCoverLinkCode(v string)`

SetVoucherCoverLinkCode sets VoucherCoverLinkCode field to given value.

### HasVoucherCoverLinkCode

`func (o *VOUCHERVSCHEMA) HasVoucherCoverLinkCode() bool`

HasVoucherCoverLinkCode returns a boolean if a field has been set.

### SetVoucherCoverLinkCodeNil

`func (o *VOUCHERVSCHEMA) SetVoucherCoverLinkCodeNil(b bool)`

 SetVoucherCoverLinkCodeNil sets the value for VoucherCoverLinkCode to be an explicit nil

### UnsetVoucherCoverLinkCode
`func (o *VOUCHERVSCHEMA) UnsetVoucherCoverLinkCode()`

UnsetVoucherCoverLinkCode ensures that no value is present for VoucherCoverLinkCode, not even an explicit nil
### GetVoucherSerial

`func (o *VOUCHERVSCHEMA) GetVoucherSerial() string`

GetVoucherSerial returns the VoucherSerial field if non-nil, zero value otherwise.

### GetVoucherSerialOk

`func (o *VOUCHERVSCHEMA) GetVoucherSerialOk() (*string, bool)`

GetVoucherSerialOk returns a tuple with the VoucherSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherSerial

`func (o *VOUCHERVSCHEMA) SetVoucherSerial(v string)`

SetVoucherSerial sets VoucherSerial field to given value.

### HasVoucherSerial

`func (o *VOUCHERVSCHEMA) HasVoucherSerial() bool`

HasVoucherSerial returns a boolean if a field has been set.

### SetVoucherSerialNil

`func (o *VOUCHERVSCHEMA) SetVoucherSerialNil(b bool)`

 SetVoucherSerialNil sets the value for VoucherSerial to be an explicit nil

### UnsetVoucherSerial
`func (o *VOUCHERVSCHEMA) UnsetVoucherSerial()`

UnsetVoucherSerial ensures that no value is present for VoucherSerial, not even an explicit nil
### GetExpiryDate

`func (o *VOUCHERVSCHEMA) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *VOUCHERVSCHEMA) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *VOUCHERVSCHEMA) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *VOUCHERVSCHEMA) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *VOUCHERVSCHEMA) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *VOUCHERVSCHEMA) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetProduct

`func (o *VOUCHERVSCHEMA) GetProduct() VOUCHERVSCHEMAProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *VOUCHERVSCHEMA) GetProductOk() (*VOUCHERVSCHEMAProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *VOUCHERVSCHEMA) SetProduct(v VOUCHERVSCHEMAProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *VOUCHERVSCHEMA) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *VOUCHERVSCHEMA) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *VOUCHERVSCHEMA) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


