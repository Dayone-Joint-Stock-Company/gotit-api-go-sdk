/*
Sample API

Technical document APIs for API Version 4.

API version: 4.0.0
Contact: quang.huynh@gotit.vn
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gotit_api_go_sdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the REQUESTSENDVOUCHERBYEMAIL type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &REQUESTSENDVOUCHERBYEMAIL{}

// REQUESTSENDVOUCHERBYEMAIL struct for REQUESTSENDVOUCHERBYEMAIL
type REQUESTSENDVOUCHERBYEMAIL struct {
	VoucherLinkCode string `json:"voucherLinkCode"`
	Email string `json:"email"`
	ReceiverNm *string `json:"receiverNm,omitempty"`
	SenderNm *string `json:"senderNm,omitempty"`
	Message *string `json:"message,omitempty"`
}

type _REQUESTSENDVOUCHERBYEMAIL REQUESTSENDVOUCHERBYEMAIL

// NewREQUESTSENDVOUCHERBYEMAIL instantiates a new REQUESTSENDVOUCHERBYEMAIL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewREQUESTSENDVOUCHERBYEMAIL(voucherLinkCode string, email string) *REQUESTSENDVOUCHERBYEMAIL {
	this := REQUESTSENDVOUCHERBYEMAIL{}
	this.VoucherLinkCode = voucherLinkCode
	this.Email = email
	return &this
}

// NewREQUESTSENDVOUCHERBYEMAILWithDefaults instantiates a new REQUESTSENDVOUCHERBYEMAIL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewREQUESTSENDVOUCHERBYEMAILWithDefaults() *REQUESTSENDVOUCHERBYEMAIL {
	this := REQUESTSENDVOUCHERBYEMAIL{}
	return &this
}

// GetVoucherLinkCode returns the VoucherLinkCode field value
func (o *REQUESTSENDVOUCHERBYEMAIL) GetVoucherLinkCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VoucherLinkCode
}

// GetVoucherLinkCodeOk returns a tuple with the VoucherLinkCode field value
// and a boolean to check if the value has been set.
func (o *REQUESTSENDVOUCHERBYEMAIL) GetVoucherLinkCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VoucherLinkCode, true
}

// SetVoucherLinkCode sets field value
func (o *REQUESTSENDVOUCHERBYEMAIL) SetVoucherLinkCode(v string) {
	o.VoucherLinkCode = v
}

// GetEmail returns the Email field value
func (o *REQUESTSENDVOUCHERBYEMAIL) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *REQUESTSENDVOUCHERBYEMAIL) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *REQUESTSENDVOUCHERBYEMAIL) SetEmail(v string) {
	o.Email = v
}

// GetReceiverNm returns the ReceiverNm field value if set, zero value otherwise.
func (o *REQUESTSENDVOUCHERBYEMAIL) GetReceiverNm() string {
	if o == nil || IsNil(o.ReceiverNm) {
		var ret string
		return ret
	}
	return *o.ReceiverNm
}

// GetReceiverNmOk returns a tuple with the ReceiverNm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *REQUESTSENDVOUCHERBYEMAIL) GetReceiverNmOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiverNm) {
		return nil, false
	}
	return o.ReceiverNm, true
}

// HasReceiverNm returns a boolean if a field has been set.
func (o *REQUESTSENDVOUCHERBYEMAIL) HasReceiverNm() bool {
	if o != nil && !IsNil(o.ReceiverNm) {
		return true
	}

	return false
}

// SetReceiverNm gets a reference to the given string and assigns it to the ReceiverNm field.
func (o *REQUESTSENDVOUCHERBYEMAIL) SetReceiverNm(v string) {
	o.ReceiverNm = &v
}

// GetSenderNm returns the SenderNm field value if set, zero value otherwise.
func (o *REQUESTSENDVOUCHERBYEMAIL) GetSenderNm() string {
	if o == nil || IsNil(o.SenderNm) {
		var ret string
		return ret
	}
	return *o.SenderNm
}

// GetSenderNmOk returns a tuple with the SenderNm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *REQUESTSENDVOUCHERBYEMAIL) GetSenderNmOk() (*string, bool) {
	if o == nil || IsNil(o.SenderNm) {
		return nil, false
	}
	return o.SenderNm, true
}

// HasSenderNm returns a boolean if a field has been set.
func (o *REQUESTSENDVOUCHERBYEMAIL) HasSenderNm() bool {
	if o != nil && !IsNil(o.SenderNm) {
		return true
	}

	return false
}

// SetSenderNm gets a reference to the given string and assigns it to the SenderNm field.
func (o *REQUESTSENDVOUCHERBYEMAIL) SetSenderNm(v string) {
	o.SenderNm = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *REQUESTSENDVOUCHERBYEMAIL) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *REQUESTSENDVOUCHERBYEMAIL) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *REQUESTSENDVOUCHERBYEMAIL) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *REQUESTSENDVOUCHERBYEMAIL) SetMessage(v string) {
	o.Message = &v
}

func (o REQUESTSENDVOUCHERBYEMAIL) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o REQUESTSENDVOUCHERBYEMAIL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["voucherLinkCode"] = o.VoucherLinkCode
	toSerialize["email"] = o.Email
	if !IsNil(o.ReceiverNm) {
		toSerialize["receiverNm"] = o.ReceiverNm
	}
	if !IsNil(o.SenderNm) {
		toSerialize["senderNm"] = o.SenderNm
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

func (o *REQUESTSENDVOUCHERBYEMAIL) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"voucherLinkCode",
		"email",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varREQUESTSENDVOUCHERBYEMAIL := _REQUESTSENDVOUCHERBYEMAIL{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varREQUESTSENDVOUCHERBYEMAIL)

	if err != nil {
		return err
	}

	*o = REQUESTSENDVOUCHERBYEMAIL(varREQUESTSENDVOUCHERBYEMAIL)

	return err
}

type NullableREQUESTSENDVOUCHERBYEMAIL struct {
	value *REQUESTSENDVOUCHERBYEMAIL
	isSet bool
}

func (v NullableREQUESTSENDVOUCHERBYEMAIL) Get() *REQUESTSENDVOUCHERBYEMAIL {
	return v.value
}

func (v *NullableREQUESTSENDVOUCHERBYEMAIL) Set(val *REQUESTSENDVOUCHERBYEMAIL) {
	v.value = val
	v.isSet = true
}

func (v NullableREQUESTSENDVOUCHERBYEMAIL) IsSet() bool {
	return v.isSet
}

func (v *NullableREQUESTSENDVOUCHERBYEMAIL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableREQUESTSENDVOUCHERBYEMAIL(val *REQUESTSENDVOUCHERBYEMAIL) *NullableREQUESTSENDVOUCHERBYEMAIL {
	return &NullableREQUESTSENDVOUCHERBYEMAIL{value: val, isSet: true}
}

func (v NullableREQUESTSENDVOUCHERBYEMAIL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableREQUESTSENDVOUCHERBYEMAIL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


