# REQUESTSENDVOUCHERBYZNS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNo** | **string** |  | 
**ReceiverNm** | Pointer to **string** |  | [optional] 
**VoucherSerials** | **[]string** |  | 
**TransactionId** | **string** |  | 

## Methods

### NewREQUESTSENDVOUCHERBYZNS

`func NewREQUESTSENDVOUCHERBYZNS(phoneNo string, voucherSerials []string, transactionId string, ) *REQUESTSENDVOUCHERBYZNS`

NewREQUESTSENDVOUCHERBYZNS instantiates a new REQUESTSENDVOUCHERBYZNS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewREQUESTSENDVOUCHERBYZNSWithDefaults

`func NewREQUESTSENDVOUCHERBYZNSWithDefaults() *REQUESTSENDVOUCHERBYZNS`

NewREQUESTSENDVOUCHERBYZNSWithDefaults instantiates a new REQUESTSENDVOUCHERBYZNS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNo

`func (o *REQUESTSENDVOUCHERBYZNS) GetPhoneNo() string`

GetPhoneNo returns the PhoneNo field if non-nil, zero value otherwise.

### GetPhoneNoOk

`func (o *REQUESTSENDVOUCHERBYZNS) GetPhoneNoOk() (*string, bool)`

GetPhoneNoOk returns a tuple with the PhoneNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNo

`func (o *REQUESTSENDVOUCHERBYZNS) SetPhoneNo(v string)`

SetPhoneNo sets PhoneNo field to given value.


### GetReceiverNm

`func (o *REQUESTSENDVOUCHERBYZNS) GetReceiverNm() string`

GetReceiverNm returns the ReceiverNm field if non-nil, zero value otherwise.

### GetReceiverNmOk

`func (o *REQUESTSENDVOUCHERBYZNS) GetReceiverNmOk() (*string, bool)`

GetReceiverNmOk returns a tuple with the ReceiverNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverNm

`func (o *REQUESTSENDVOUCHERBYZNS) SetReceiverNm(v string)`

SetReceiverNm sets ReceiverNm field to given value.

### HasReceiverNm

`func (o *REQUESTSENDVOUCHERBYZNS) HasReceiverNm() bool`

HasReceiverNm returns a boolean if a field has been set.

### GetVoucherSerials

`func (o *REQUESTSENDVOUCHERBYZNS) GetVoucherSerials() []string`

GetVoucherSerials returns the VoucherSerials field if non-nil, zero value otherwise.

### GetVoucherSerialsOk

`func (o *REQUESTSENDVOUCHERBYZNS) GetVoucherSerialsOk() (*[]string, bool)`

GetVoucherSerialsOk returns a tuple with the VoucherSerials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherSerials

`func (o *REQUESTSENDVOUCHERBYZNS) SetVoucherSerials(v []string)`

SetVoucherSerials sets VoucherSerials field to given value.


### GetTransactionId

`func (o *REQUESTSENDVOUCHERBYZNS) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *REQUESTSENDVOUCHERBYZNS) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *REQUESTSENDVOUCHERBYZNS) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


