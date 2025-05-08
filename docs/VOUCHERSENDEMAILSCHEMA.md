# VOUCHERSENDEMAILSCHEMA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VoucherLinkCode** | Pointer to **NullableString** | Last 8 characters of voucher link | [optional] 
**Email** | Pointer to **NullableString** | Email of user receive voucher | [optional] 
**ReceiverNm** | Pointer to **NullableString** | Email of user receive voucher | [optional] 
**SenderNm** | Pointer to **NullableString** | Name of sender | [optional] 
**Message** | Pointer to **NullableString** | Message send | [optional] 

## Methods

### NewVOUCHERSENDEMAILSCHEMA

`func NewVOUCHERSENDEMAILSCHEMA() *VOUCHERSENDEMAILSCHEMA`

NewVOUCHERSENDEMAILSCHEMA instantiates a new VOUCHERSENDEMAILSCHEMA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHERSENDEMAILSCHEMAWithDefaults

`func NewVOUCHERSENDEMAILSCHEMAWithDefaults() *VOUCHERSENDEMAILSCHEMA`

NewVOUCHERSENDEMAILSCHEMAWithDefaults instantiates a new VOUCHERSENDEMAILSCHEMA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoucherLinkCode

`func (o *VOUCHERSENDEMAILSCHEMA) GetVoucherLinkCode() string`

GetVoucherLinkCode returns the VoucherLinkCode field if non-nil, zero value otherwise.

### GetVoucherLinkCodeOk

`func (o *VOUCHERSENDEMAILSCHEMA) GetVoucherLinkCodeOk() (*string, bool)`

GetVoucherLinkCodeOk returns a tuple with the VoucherLinkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherLinkCode

`func (o *VOUCHERSENDEMAILSCHEMA) SetVoucherLinkCode(v string)`

SetVoucherLinkCode sets VoucherLinkCode field to given value.

### HasVoucherLinkCode

`func (o *VOUCHERSENDEMAILSCHEMA) HasVoucherLinkCode() bool`

HasVoucherLinkCode returns a boolean if a field has been set.

### SetVoucherLinkCodeNil

`func (o *VOUCHERSENDEMAILSCHEMA) SetVoucherLinkCodeNil(b bool)`

 SetVoucherLinkCodeNil sets the value for VoucherLinkCode to be an explicit nil

### UnsetVoucherLinkCode
`func (o *VOUCHERSENDEMAILSCHEMA) UnsetVoucherLinkCode()`

UnsetVoucherLinkCode ensures that no value is present for VoucherLinkCode, not even an explicit nil
### GetEmail

`func (o *VOUCHERSENDEMAILSCHEMA) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VOUCHERSENDEMAILSCHEMA) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VOUCHERSENDEMAILSCHEMA) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *VOUCHERSENDEMAILSCHEMA) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *VOUCHERSENDEMAILSCHEMA) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *VOUCHERSENDEMAILSCHEMA) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetReceiverNm

`func (o *VOUCHERSENDEMAILSCHEMA) GetReceiverNm() string`

GetReceiverNm returns the ReceiverNm field if non-nil, zero value otherwise.

### GetReceiverNmOk

`func (o *VOUCHERSENDEMAILSCHEMA) GetReceiverNmOk() (*string, bool)`

GetReceiverNmOk returns a tuple with the ReceiverNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverNm

`func (o *VOUCHERSENDEMAILSCHEMA) SetReceiverNm(v string)`

SetReceiverNm sets ReceiverNm field to given value.

### HasReceiverNm

`func (o *VOUCHERSENDEMAILSCHEMA) HasReceiverNm() bool`

HasReceiverNm returns a boolean if a field has been set.

### SetReceiverNmNil

`func (o *VOUCHERSENDEMAILSCHEMA) SetReceiverNmNil(b bool)`

 SetReceiverNmNil sets the value for ReceiverNm to be an explicit nil

### UnsetReceiverNm
`func (o *VOUCHERSENDEMAILSCHEMA) UnsetReceiverNm()`

UnsetReceiverNm ensures that no value is present for ReceiverNm, not even an explicit nil
### GetSenderNm

`func (o *VOUCHERSENDEMAILSCHEMA) GetSenderNm() string`

GetSenderNm returns the SenderNm field if non-nil, zero value otherwise.

### GetSenderNmOk

`func (o *VOUCHERSENDEMAILSCHEMA) GetSenderNmOk() (*string, bool)`

GetSenderNmOk returns a tuple with the SenderNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderNm

`func (o *VOUCHERSENDEMAILSCHEMA) SetSenderNm(v string)`

SetSenderNm sets SenderNm field to given value.

### HasSenderNm

`func (o *VOUCHERSENDEMAILSCHEMA) HasSenderNm() bool`

HasSenderNm returns a boolean if a field has been set.

### SetSenderNmNil

`func (o *VOUCHERSENDEMAILSCHEMA) SetSenderNmNil(b bool)`

 SetSenderNmNil sets the value for SenderNm to be an explicit nil

### UnsetSenderNm
`func (o *VOUCHERSENDEMAILSCHEMA) UnsetSenderNm()`

UnsetSenderNm ensures that no value is present for SenderNm, not even an explicit nil
### GetMessage

`func (o *VOUCHERSENDEMAILSCHEMA) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VOUCHERSENDEMAILSCHEMA) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VOUCHERSENDEMAILSCHEMA) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VOUCHERSENDEMAILSCHEMA) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *VOUCHERSENDEMAILSCHEMA) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *VOUCHERSENDEMAILSCHEMA) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


