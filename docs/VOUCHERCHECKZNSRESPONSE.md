# VOUCHERCHECKZNSRESPONSE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | HTTP response status | [optional] 
**StatusCode** | Pointer to **int64** | HTTP response status codes | [optional] 
**Error** | Pointer to **string** | Error code | [optional] 
**Message** | Pointer to **string** | Message Error | [optional] 
**Data** | Pointer to [**VOUCHERCHECKZNSRESPONSEData**](VOUCHERCHECKZNSRESPONSEData.md) |  | [optional] 

## Methods

### NewVOUCHERCHECKZNSRESPONSE

`func NewVOUCHERCHECKZNSRESPONSE() *VOUCHERCHECKZNSRESPONSE`

NewVOUCHERCHECKZNSRESPONSE instantiates a new VOUCHERCHECKZNSRESPONSE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHERCHECKZNSRESPONSEWithDefaults

`func NewVOUCHERCHECKZNSRESPONSEWithDefaults() *VOUCHERCHECKZNSRESPONSE`

NewVOUCHERCHECKZNSRESPONSEWithDefaults instantiates a new VOUCHERCHECKZNSRESPONSE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VOUCHERCHECKZNSRESPONSE) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VOUCHERCHECKZNSRESPONSE) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VOUCHERCHECKZNSRESPONSE) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VOUCHERCHECKZNSRESPONSE) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *VOUCHERCHECKZNSRESPONSE) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *VOUCHERCHECKZNSRESPONSE) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *VOUCHERCHECKZNSRESPONSE) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *VOUCHERCHECKZNSRESPONSE) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *VOUCHERCHECKZNSRESPONSE) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VOUCHERCHECKZNSRESPONSE) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VOUCHERCHECKZNSRESPONSE) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *VOUCHERCHECKZNSRESPONSE) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *VOUCHERCHECKZNSRESPONSE) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VOUCHERCHECKZNSRESPONSE) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VOUCHERCHECKZNSRESPONSE) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VOUCHERCHECKZNSRESPONSE) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *VOUCHERCHECKZNSRESPONSE) GetData() VOUCHERCHECKZNSRESPONSEData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VOUCHERCHECKZNSRESPONSE) GetDataOk() (*VOUCHERCHECKZNSRESPONSEData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VOUCHERCHECKZNSRESPONSE) SetData(v VOUCHERCHECKZNSRESPONSEData)`

SetData sets Data field to given value.

### HasData

`func (o *VOUCHERCHECKZNSRESPONSE) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


