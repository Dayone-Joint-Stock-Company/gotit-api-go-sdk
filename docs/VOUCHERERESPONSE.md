# VOUCHERERESPONSE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | HTTP response status | [optional] 
**StatusCode** | Pointer to **int64** | HTTP response status codes | [optional] 
**Error** | Pointer to **string** | Error code | [optional] 
**Message** | Pointer to **string** | Message Error | [optional] 
**Data** | Pointer to [**[]VOUCHERE**](VOUCHERE.md) |  | [optional] 

## Methods

### NewVOUCHERERESPONSE

`func NewVOUCHERERESPONSE() *VOUCHERERESPONSE`

NewVOUCHERERESPONSE instantiates a new VOUCHERERESPONSE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHERERESPONSEWithDefaults

`func NewVOUCHERERESPONSEWithDefaults() *VOUCHERERESPONSE`

NewVOUCHERERESPONSEWithDefaults instantiates a new VOUCHERERESPONSE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VOUCHERERESPONSE) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VOUCHERERESPONSE) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VOUCHERERESPONSE) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VOUCHERERESPONSE) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *VOUCHERERESPONSE) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *VOUCHERERESPONSE) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *VOUCHERERESPONSE) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *VOUCHERERESPONSE) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *VOUCHERERESPONSE) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VOUCHERERESPONSE) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VOUCHERERESPONSE) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *VOUCHERERESPONSE) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *VOUCHERERESPONSE) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VOUCHERERESPONSE) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VOUCHERERESPONSE) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VOUCHERERESPONSE) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *VOUCHERERESPONSE) GetData() []VOUCHERE`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VOUCHERERESPONSE) GetDataOk() (*[]VOUCHERE, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VOUCHERERESPONSE) SetData(v []VOUCHERE)`

SetData sets Data field to given value.

### HasData

`func (o *VOUCHERERESPONSE) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


