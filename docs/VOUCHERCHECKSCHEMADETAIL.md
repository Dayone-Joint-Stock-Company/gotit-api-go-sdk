# VOUCHERCHECKSCHEMADETAIL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionRefId** | Pointer to **NullableString** | TransactionRefId receive from client request | [optional] 
**VoucherCode** | Pointer to **NullableString** | Voucher code. Contains 10 numbers, unique and not duplicated with any other voucher. For some brands that do not use Got It code but use integerernal code (Big C, Mobile World), the system returns the integerernal code. | [optional] 
**VoucherLink** | Pointer to **NullableString** | Link of voucher | [optional] 
**VoucherCoverLink** | Pointer to **NullableString** | Cover link of voucher. | [optional] 
**VoucherSerial** | Pointer to **NullableString** | Is a unique code to identify voucher link v and is not valid for use. For example: V1562342ET2 | [optional] 
**GroupLink** | Pointer to **NullableString** | https://e-stg.gotit.vn/gLsZaFRN | [optional] 
**GroupVoucherSerial** | Pointer to **NullableString** | E2V2RML6F52 | [optional] 
**GroupCoverLink** | Pointer to **NullableString** | https://e-stg.gotit.vn/gLsZaFRN | [optional] 
**VoucherLinkCode** | Pointer to **NullableString** | Last 8 characters of voucher link | [optional] 
**VoucherImageLink** | Pointer to **NullableString** | Link of voucher image. To display vouchers as images | [optional] 
**ExpiryDate** | Pointer to **NullableString** | Voucher expiration date | [optional] 
**StateCode** | Pointer to **NullableInt64** | State Code | [optional] 
**StateText** | Pointer to **NullableString** | State Text | [optional] 
**UsedStore** | Pointer to [**[]STORESSCHEMA**](STORESSCHEMA.md) |  | [optional] 
**UsedTime** | Pointer to **NullableString** | Used Time | [optional] 
**UsedBrand** | Pointer to **NullableString** | State Text | [optional] 
**Product** | Pointer to [**PRODUCTV**](PRODUCTV.md) |  | [optional] 

## Methods

### NewVOUCHERCHECKSCHEMADETAIL

`func NewVOUCHERCHECKSCHEMADETAIL() *VOUCHERCHECKSCHEMADETAIL`

NewVOUCHERCHECKSCHEMADETAIL instantiates a new VOUCHERCHECKSCHEMADETAIL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHERCHECKSCHEMADETAILWithDefaults

`func NewVOUCHERCHECKSCHEMADETAILWithDefaults() *VOUCHERCHECKSCHEMADETAIL`

NewVOUCHERCHECKSCHEMADETAILWithDefaults instantiates a new VOUCHERCHECKSCHEMADETAIL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionRefId

`func (o *VOUCHERCHECKSCHEMADETAIL) GetTransactionRefId() string`

GetTransactionRefId returns the TransactionRefId field if non-nil, zero value otherwise.

### GetTransactionRefIdOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetTransactionRefIdOk() (*string, bool)`

GetTransactionRefIdOk returns a tuple with the TransactionRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRefId

`func (o *VOUCHERCHECKSCHEMADETAIL) SetTransactionRefId(v string)`

SetTransactionRefId sets TransactionRefId field to given value.

### HasTransactionRefId

`func (o *VOUCHERCHECKSCHEMADETAIL) HasTransactionRefId() bool`

HasTransactionRefId returns a boolean if a field has been set.

### SetTransactionRefIdNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetTransactionRefIdNil(b bool)`

 SetTransactionRefIdNil sets the value for TransactionRefId to be an explicit nil

### UnsetTransactionRefId
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetTransactionRefId()`

UnsetTransactionRefId ensures that no value is present for TransactionRefId, not even an explicit nil
### GetVoucherCode

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherCode() string`

GetVoucherCode returns the VoucherCode field if non-nil, zero value otherwise.

### GetVoucherCodeOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherCodeOk() (*string, bool)`

GetVoucherCodeOk returns a tuple with the VoucherCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherCode

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherCode(v string)`

SetVoucherCode sets VoucherCode field to given value.

### HasVoucherCode

`func (o *VOUCHERCHECKSCHEMADETAIL) HasVoucherCode() bool`

HasVoucherCode returns a boolean if a field has been set.

### SetVoucherCodeNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherCodeNil(b bool)`

 SetVoucherCodeNil sets the value for VoucherCode to be an explicit nil

### UnsetVoucherCode
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetVoucherCode()`

UnsetVoucherCode ensures that no value is present for VoucherCode, not even an explicit nil
### GetVoucherLink

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherLink() string`

GetVoucherLink returns the VoucherLink field if non-nil, zero value otherwise.

### GetVoucherLinkOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherLinkOk() (*string, bool)`

GetVoucherLinkOk returns a tuple with the VoucherLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherLink

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherLink(v string)`

SetVoucherLink sets VoucherLink field to given value.

### HasVoucherLink

`func (o *VOUCHERCHECKSCHEMADETAIL) HasVoucherLink() bool`

HasVoucherLink returns a boolean if a field has been set.

### SetVoucherLinkNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherLinkNil(b bool)`

 SetVoucherLinkNil sets the value for VoucherLink to be an explicit nil

### UnsetVoucherLink
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetVoucherLink()`

UnsetVoucherLink ensures that no value is present for VoucherLink, not even an explicit nil
### GetVoucherCoverLink

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherCoverLink() string`

GetVoucherCoverLink returns the VoucherCoverLink field if non-nil, zero value otherwise.

### GetVoucherCoverLinkOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherCoverLinkOk() (*string, bool)`

GetVoucherCoverLinkOk returns a tuple with the VoucherCoverLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherCoverLink

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherCoverLink(v string)`

SetVoucherCoverLink sets VoucherCoverLink field to given value.

### HasVoucherCoverLink

`func (o *VOUCHERCHECKSCHEMADETAIL) HasVoucherCoverLink() bool`

HasVoucherCoverLink returns a boolean if a field has been set.

### SetVoucherCoverLinkNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherCoverLinkNil(b bool)`

 SetVoucherCoverLinkNil sets the value for VoucherCoverLink to be an explicit nil

### UnsetVoucherCoverLink
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetVoucherCoverLink()`

UnsetVoucherCoverLink ensures that no value is present for VoucherCoverLink, not even an explicit nil
### GetVoucherSerial

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherSerial() string`

GetVoucherSerial returns the VoucherSerial field if non-nil, zero value otherwise.

### GetVoucherSerialOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherSerialOk() (*string, bool)`

GetVoucherSerialOk returns a tuple with the VoucherSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherSerial

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherSerial(v string)`

SetVoucherSerial sets VoucherSerial field to given value.

### HasVoucherSerial

`func (o *VOUCHERCHECKSCHEMADETAIL) HasVoucherSerial() bool`

HasVoucherSerial returns a boolean if a field has been set.

### SetVoucherSerialNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherSerialNil(b bool)`

 SetVoucherSerialNil sets the value for VoucherSerial to be an explicit nil

### UnsetVoucherSerial
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetVoucherSerial()`

UnsetVoucherSerial ensures that no value is present for VoucherSerial, not even an explicit nil
### GetGroupLink

`func (o *VOUCHERCHECKSCHEMADETAIL) GetGroupLink() string`

GetGroupLink returns the GroupLink field if non-nil, zero value otherwise.

### GetGroupLinkOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetGroupLinkOk() (*string, bool)`

GetGroupLinkOk returns a tuple with the GroupLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLink

`func (o *VOUCHERCHECKSCHEMADETAIL) SetGroupLink(v string)`

SetGroupLink sets GroupLink field to given value.

### HasGroupLink

`func (o *VOUCHERCHECKSCHEMADETAIL) HasGroupLink() bool`

HasGroupLink returns a boolean if a field has been set.

### SetGroupLinkNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetGroupLinkNil(b bool)`

 SetGroupLinkNil sets the value for GroupLink to be an explicit nil

### UnsetGroupLink
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetGroupLink()`

UnsetGroupLink ensures that no value is present for GroupLink, not even an explicit nil
### GetGroupVoucherSerial

`func (o *VOUCHERCHECKSCHEMADETAIL) GetGroupVoucherSerial() string`

GetGroupVoucherSerial returns the GroupVoucherSerial field if non-nil, zero value otherwise.

### GetGroupVoucherSerialOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetGroupVoucherSerialOk() (*string, bool)`

GetGroupVoucherSerialOk returns a tuple with the GroupVoucherSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupVoucherSerial

`func (o *VOUCHERCHECKSCHEMADETAIL) SetGroupVoucherSerial(v string)`

SetGroupVoucherSerial sets GroupVoucherSerial field to given value.

### HasGroupVoucherSerial

`func (o *VOUCHERCHECKSCHEMADETAIL) HasGroupVoucherSerial() bool`

HasGroupVoucherSerial returns a boolean if a field has been set.

### SetGroupVoucherSerialNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetGroupVoucherSerialNil(b bool)`

 SetGroupVoucherSerialNil sets the value for GroupVoucherSerial to be an explicit nil

### UnsetGroupVoucherSerial
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetGroupVoucherSerial()`

UnsetGroupVoucherSerial ensures that no value is present for GroupVoucherSerial, not even an explicit nil
### GetGroupCoverLink

`func (o *VOUCHERCHECKSCHEMADETAIL) GetGroupCoverLink() string`

GetGroupCoverLink returns the GroupCoverLink field if non-nil, zero value otherwise.

### GetGroupCoverLinkOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetGroupCoverLinkOk() (*string, bool)`

GetGroupCoverLinkOk returns a tuple with the GroupCoverLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCoverLink

`func (o *VOUCHERCHECKSCHEMADETAIL) SetGroupCoverLink(v string)`

SetGroupCoverLink sets GroupCoverLink field to given value.

### HasGroupCoverLink

`func (o *VOUCHERCHECKSCHEMADETAIL) HasGroupCoverLink() bool`

HasGroupCoverLink returns a boolean if a field has been set.

### SetGroupCoverLinkNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetGroupCoverLinkNil(b bool)`

 SetGroupCoverLinkNil sets the value for GroupCoverLink to be an explicit nil

### UnsetGroupCoverLink
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetGroupCoverLink()`

UnsetGroupCoverLink ensures that no value is present for GroupCoverLink, not even an explicit nil
### GetVoucherLinkCode

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherLinkCode() string`

GetVoucherLinkCode returns the VoucherLinkCode field if non-nil, zero value otherwise.

### GetVoucherLinkCodeOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherLinkCodeOk() (*string, bool)`

GetVoucherLinkCodeOk returns a tuple with the VoucherLinkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherLinkCode

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherLinkCode(v string)`

SetVoucherLinkCode sets VoucherLinkCode field to given value.

### HasVoucherLinkCode

`func (o *VOUCHERCHECKSCHEMADETAIL) HasVoucherLinkCode() bool`

HasVoucherLinkCode returns a boolean if a field has been set.

### SetVoucherLinkCodeNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherLinkCodeNil(b bool)`

 SetVoucherLinkCodeNil sets the value for VoucherLinkCode to be an explicit nil

### UnsetVoucherLinkCode
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetVoucherLinkCode()`

UnsetVoucherLinkCode ensures that no value is present for VoucherLinkCode, not even an explicit nil
### GetVoucherImageLink

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherImageLink() string`

GetVoucherImageLink returns the VoucherImageLink field if non-nil, zero value otherwise.

### GetVoucherImageLinkOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetVoucherImageLinkOk() (*string, bool)`

GetVoucherImageLinkOk returns a tuple with the VoucherImageLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherImageLink

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherImageLink(v string)`

SetVoucherImageLink sets VoucherImageLink field to given value.

### HasVoucherImageLink

`func (o *VOUCHERCHECKSCHEMADETAIL) HasVoucherImageLink() bool`

HasVoucherImageLink returns a boolean if a field has been set.

### SetVoucherImageLinkNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetVoucherImageLinkNil(b bool)`

 SetVoucherImageLinkNil sets the value for VoucherImageLink to be an explicit nil

### UnsetVoucherImageLink
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetVoucherImageLink()`

UnsetVoucherImageLink ensures that no value is present for VoucherImageLink, not even an explicit nil
### GetExpiryDate

`func (o *VOUCHERCHECKSCHEMADETAIL) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *VOUCHERCHECKSCHEMADETAIL) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *VOUCHERCHECKSCHEMADETAIL) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetStateCode

`func (o *VOUCHERCHECKSCHEMADETAIL) GetStateCode() int64`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetStateCodeOk() (*int64, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *VOUCHERCHECKSCHEMADETAIL) SetStateCode(v int64)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *VOUCHERCHECKSCHEMADETAIL) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### SetStateCodeNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetStateCodeNil(b bool)`

 SetStateCodeNil sets the value for StateCode to be an explicit nil

### UnsetStateCode
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetStateCode()`

UnsetStateCode ensures that no value is present for StateCode, not even an explicit nil
### GetStateText

`func (o *VOUCHERCHECKSCHEMADETAIL) GetStateText() string`

GetStateText returns the StateText field if non-nil, zero value otherwise.

### GetStateTextOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetStateTextOk() (*string, bool)`

GetStateTextOk returns a tuple with the StateText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateText

`func (o *VOUCHERCHECKSCHEMADETAIL) SetStateText(v string)`

SetStateText sets StateText field to given value.

### HasStateText

`func (o *VOUCHERCHECKSCHEMADETAIL) HasStateText() bool`

HasStateText returns a boolean if a field has been set.

### SetStateTextNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetStateTextNil(b bool)`

 SetStateTextNil sets the value for StateText to be an explicit nil

### UnsetStateText
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetStateText()`

UnsetStateText ensures that no value is present for StateText, not even an explicit nil
### GetUsedStore

`func (o *VOUCHERCHECKSCHEMADETAIL) GetUsedStore() []STORESSCHEMA`

GetUsedStore returns the UsedStore field if non-nil, zero value otherwise.

### GetUsedStoreOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetUsedStoreOk() (*[]STORESSCHEMA, bool)`

GetUsedStoreOk returns a tuple with the UsedStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStore

`func (o *VOUCHERCHECKSCHEMADETAIL) SetUsedStore(v []STORESSCHEMA)`

SetUsedStore sets UsedStore field to given value.

### HasUsedStore

`func (o *VOUCHERCHECKSCHEMADETAIL) HasUsedStore() bool`

HasUsedStore returns a boolean if a field has been set.

### SetUsedStoreNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetUsedStoreNil(b bool)`

 SetUsedStoreNil sets the value for UsedStore to be an explicit nil

### UnsetUsedStore
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetUsedStore()`

UnsetUsedStore ensures that no value is present for UsedStore, not even an explicit nil
### GetUsedTime

`func (o *VOUCHERCHECKSCHEMADETAIL) GetUsedTime() string`

GetUsedTime returns the UsedTime field if non-nil, zero value otherwise.

### GetUsedTimeOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetUsedTimeOk() (*string, bool)`

GetUsedTimeOk returns a tuple with the UsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTime

`func (o *VOUCHERCHECKSCHEMADETAIL) SetUsedTime(v string)`

SetUsedTime sets UsedTime field to given value.

### HasUsedTime

`func (o *VOUCHERCHECKSCHEMADETAIL) HasUsedTime() bool`

HasUsedTime returns a boolean if a field has been set.

### SetUsedTimeNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetUsedTimeNil(b bool)`

 SetUsedTimeNil sets the value for UsedTime to be an explicit nil

### UnsetUsedTime
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetUsedTime()`

UnsetUsedTime ensures that no value is present for UsedTime, not even an explicit nil
### GetUsedBrand

`func (o *VOUCHERCHECKSCHEMADETAIL) GetUsedBrand() string`

GetUsedBrand returns the UsedBrand field if non-nil, zero value otherwise.

### GetUsedBrandOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetUsedBrandOk() (*string, bool)`

GetUsedBrandOk returns a tuple with the UsedBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBrand

`func (o *VOUCHERCHECKSCHEMADETAIL) SetUsedBrand(v string)`

SetUsedBrand sets UsedBrand field to given value.

### HasUsedBrand

`func (o *VOUCHERCHECKSCHEMADETAIL) HasUsedBrand() bool`

HasUsedBrand returns a boolean if a field has been set.

### SetUsedBrandNil

`func (o *VOUCHERCHECKSCHEMADETAIL) SetUsedBrandNil(b bool)`

 SetUsedBrandNil sets the value for UsedBrand to be an explicit nil

### UnsetUsedBrand
`func (o *VOUCHERCHECKSCHEMADETAIL) UnsetUsedBrand()`

UnsetUsedBrand ensures that no value is present for UsedBrand, not even an explicit nil
### GetProduct

`func (o *VOUCHERCHECKSCHEMADETAIL) GetProduct() PRODUCTV`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *VOUCHERCHECKSCHEMADETAIL) GetProductOk() (*PRODUCTV, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *VOUCHERCHECKSCHEMADETAIL) SetProduct(v PRODUCTV)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *VOUCHERCHECKSCHEMADETAIL) HasProduct() bool`

HasProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


