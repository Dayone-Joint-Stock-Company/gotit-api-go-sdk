# VOUCHERSENDSMSRESPONSE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | HTTP response status | [optional] 
**StatusCode** | Pointer to **int64** | HTTP response status codes | [optional] 
**Error** | Pointer to **string** | Error code | [optional] 
**Message** | Pointer to **string** | Message Error | [optional] 
**Data** | Pointer to [**[]VOUCHERSENDSMSSCHEMA**](VOUCHERSENDSMSSCHEMA.md) |  | [optional] 

## Methods

### NewVOUCHERSENDSMSRESPONSE

`func NewVOUCHERSENDSMSRESPONSE() *VOUCHERSENDSMSRESPONSE`

NewVOUCHERSENDSMSRESPONSE instantiates a new VOUCHERSENDSMSRESPONSE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHERSENDSMSRESPONSEWithDefaults

`func NewVOUCHERSENDSMSRESPONSEWithDefaults() *VOUCHERSENDSMSRESPONSE`

NewVOUCHERSENDSMSRESPONSEWithDefaults instantiates a new VOUCHERSENDSMSRESPONSE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VOUCHERSENDSMSRESPONSE) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VOUCHERSENDSMSRESPONSE) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VOUCHERSENDSMSRESPONSE) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VOUCHERSENDSMSRESPONSE) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *VOUCHERSENDSMSRESPONSE) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *VOUCHERSENDSMSRESPONSE) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *VOUCHERSENDSMSRESPONSE) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *VOUCHERSENDSMSRESPONSE) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *VOUCHERSENDSMSRESPONSE) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VOUCHERSENDSMSRESPONSE) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VOUCHERSENDSMSRESPONSE) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *VOUCHERSENDSMSRESPONSE) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *VOUCHERSENDSMSRESPONSE) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VOUCHERSENDSMSRESPONSE) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VOUCHERSENDSMSRESPONSE) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VOUCHERSENDSMSRESPONSE) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *VOUCHERSENDSMSRESPONSE) GetData() []VOUCHERSENDSMSSCHEMA`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VOUCHERSENDSMSRESPONSE) GetDataOk() (*[]VOUCHERSENDSMSSCHEMA, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VOUCHERSENDSMSRESPONSE) SetData(v []VOUCHERSENDSMSSCHEMA)`

SetData sets Data field to given value.

### HasData

`func (o *VOUCHERSENDSMSRESPONSE) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


