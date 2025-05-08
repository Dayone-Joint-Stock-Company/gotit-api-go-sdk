# VOUCHERSENDZNSRESPONSE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | HTTP response status | [optional] 
**StatusCode** | Pointer to **int64** | HTTP response status codes | [optional] 
**Error** | Pointer to **string** | Error code | [optional] 
**Message** | Pointer to **string** | Message Error | [optional] 
**Data** | Pointer to [**VOUCHERSENDZNSRESPONSEData**](VOUCHERSENDZNSRESPONSEData.md) |  | [optional] 

## Methods

### NewVOUCHERSENDZNSRESPONSE

`func NewVOUCHERSENDZNSRESPONSE() *VOUCHERSENDZNSRESPONSE`

NewVOUCHERSENDZNSRESPONSE instantiates a new VOUCHERSENDZNSRESPONSE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHERSENDZNSRESPONSEWithDefaults

`func NewVOUCHERSENDZNSRESPONSEWithDefaults() *VOUCHERSENDZNSRESPONSE`

NewVOUCHERSENDZNSRESPONSEWithDefaults instantiates a new VOUCHERSENDZNSRESPONSE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VOUCHERSENDZNSRESPONSE) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VOUCHERSENDZNSRESPONSE) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VOUCHERSENDZNSRESPONSE) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VOUCHERSENDZNSRESPONSE) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *VOUCHERSENDZNSRESPONSE) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *VOUCHERSENDZNSRESPONSE) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *VOUCHERSENDZNSRESPONSE) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *VOUCHERSENDZNSRESPONSE) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *VOUCHERSENDZNSRESPONSE) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VOUCHERSENDZNSRESPONSE) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VOUCHERSENDZNSRESPONSE) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *VOUCHERSENDZNSRESPONSE) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *VOUCHERSENDZNSRESPONSE) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VOUCHERSENDZNSRESPONSE) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VOUCHERSENDZNSRESPONSE) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VOUCHERSENDZNSRESPONSE) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *VOUCHERSENDZNSRESPONSE) GetData() VOUCHERSENDZNSRESPONSEData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VOUCHERSENDZNSRESPONSE) GetDataOk() (*VOUCHERSENDZNSRESPONSEData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VOUCHERSENDZNSRESPONSE) SetData(v VOUCHERSENDZNSRESPONSEData)`

SetData sets Data field to given value.

### HasData

`func (o *VOUCHERSENDZNSRESPONSE) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


