# HTTPBADREQUEST

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewHTTPBADREQUEST

`func NewHTTPBADREQUEST() *HTTPBADREQUEST`

NewHTTPBADREQUEST instantiates a new HTTPBADREQUEST object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPBADREQUESTWithDefaults

`func NewHTTPBADREQUESTWithDefaults() *HTTPBADREQUEST`

NewHTTPBADREQUESTWithDefaults instantiates a new HTTPBADREQUEST object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HTTPBADREQUEST) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HTTPBADREQUEST) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HTTPBADREQUEST) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HTTPBADREQUEST) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *HTTPBADREQUEST) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *HTTPBADREQUEST) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *HTTPBADREQUEST) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *HTTPBADREQUEST) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *HTTPBADREQUEST) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HTTPBADREQUEST) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HTTPBADREQUEST) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *HTTPBADREQUEST) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *HTTPBADREQUEST) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HTTPBADREQUEST) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HTTPBADREQUEST) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HTTPBADREQUEST) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *HTTPBADREQUEST) GetData() []interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HTTPBADREQUEST) GetDataOk() (*[]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HTTPBADREQUEST) SetData(v []interface{})`

SetData sets Data field to given value.

### HasData

`func (o *HTTPBADREQUEST) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


