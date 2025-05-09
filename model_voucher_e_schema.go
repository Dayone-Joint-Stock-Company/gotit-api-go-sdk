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
)

// checks if the VOUCHERESCHEMA type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VOUCHERESCHEMA{}

// VOUCHERESCHEMA struct for VOUCHERESCHEMA
type VOUCHERESCHEMA struct {
	// Voucher of link e
	VoucherLink NullableString `json:"voucher_link,omitempty"`
	// Voucher cover of link e
	VoucherCoverLink NullableString `json:"voucher_cover_link,omitempty"`
	// Voucher serial of link e
	VoucherSerial NullableString `json:"voucher_serial,omitempty"`
	// Voucher value of link e
	Value NullableString `json:"value,omitempty"`
	// Expiry date of voucher link Format: 'YYYY-MM-DD'
	ExpiredDate NullableString `json:"expired_date,omitempty"`
}

// NewVOUCHERESCHEMA instantiates a new VOUCHERESCHEMA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVOUCHERESCHEMA() *VOUCHERESCHEMA {
	this := VOUCHERESCHEMA{}
	return &this
}

// NewVOUCHERESCHEMAWithDefaults instantiates a new VOUCHERESCHEMA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVOUCHERESCHEMAWithDefaults() *VOUCHERESCHEMA {
	this := VOUCHERESCHEMA{}
	return &this
}

// GetVoucherLink returns the VoucherLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VOUCHERESCHEMA) GetVoucherLink() string {
	if o == nil || IsNil(o.VoucherLink.Get()) {
		var ret string
		return ret
	}
	return *o.VoucherLink.Get()
}

// GetVoucherLinkOk returns a tuple with the VoucherLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VOUCHERESCHEMA) GetVoucherLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VoucherLink.Get(), o.VoucherLink.IsSet()
}

// HasVoucherLink returns a boolean if a field has been set.
func (o *VOUCHERESCHEMA) HasVoucherLink() bool {
	if o != nil && o.VoucherLink.IsSet() {
		return true
	}

	return false
}

// SetVoucherLink gets a reference to the given NullableString and assigns it to the VoucherLink field.
func (o *VOUCHERESCHEMA) SetVoucherLink(v string) {
	o.VoucherLink.Set(&v)
}
// SetVoucherLinkNil sets the value for VoucherLink to be an explicit nil
func (o *VOUCHERESCHEMA) SetVoucherLinkNil() {
	o.VoucherLink.Set(nil)
}

// UnsetVoucherLink ensures that no value is present for VoucherLink, not even an explicit nil
func (o *VOUCHERESCHEMA) UnsetVoucherLink() {
	o.VoucherLink.Unset()
}

// GetVoucherCoverLink returns the VoucherCoverLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VOUCHERESCHEMA) GetVoucherCoverLink() string {
	if o == nil || IsNil(o.VoucherCoverLink.Get()) {
		var ret string
		return ret
	}
	return *o.VoucherCoverLink.Get()
}

// GetVoucherCoverLinkOk returns a tuple with the VoucherCoverLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VOUCHERESCHEMA) GetVoucherCoverLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VoucherCoverLink.Get(), o.VoucherCoverLink.IsSet()
}

// HasVoucherCoverLink returns a boolean if a field has been set.
func (o *VOUCHERESCHEMA) HasVoucherCoverLink() bool {
	if o != nil && o.VoucherCoverLink.IsSet() {
		return true
	}

	return false
}

// SetVoucherCoverLink gets a reference to the given NullableString and assigns it to the VoucherCoverLink field.
func (o *VOUCHERESCHEMA) SetVoucherCoverLink(v string) {
	o.VoucherCoverLink.Set(&v)
}
// SetVoucherCoverLinkNil sets the value for VoucherCoverLink to be an explicit nil
func (o *VOUCHERESCHEMA) SetVoucherCoverLinkNil() {
	o.VoucherCoverLink.Set(nil)
}

// UnsetVoucherCoverLink ensures that no value is present for VoucherCoverLink, not even an explicit nil
func (o *VOUCHERESCHEMA) UnsetVoucherCoverLink() {
	o.VoucherCoverLink.Unset()
}

// GetVoucherSerial returns the VoucherSerial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VOUCHERESCHEMA) GetVoucherSerial() string {
	if o == nil || IsNil(o.VoucherSerial.Get()) {
		var ret string
		return ret
	}
	return *o.VoucherSerial.Get()
}

// GetVoucherSerialOk returns a tuple with the VoucherSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VOUCHERESCHEMA) GetVoucherSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VoucherSerial.Get(), o.VoucherSerial.IsSet()
}

// HasVoucherSerial returns a boolean if a field has been set.
func (o *VOUCHERESCHEMA) HasVoucherSerial() bool {
	if o != nil && o.VoucherSerial.IsSet() {
		return true
	}

	return false
}

// SetVoucherSerial gets a reference to the given NullableString and assigns it to the VoucherSerial field.
func (o *VOUCHERESCHEMA) SetVoucherSerial(v string) {
	o.VoucherSerial.Set(&v)
}
// SetVoucherSerialNil sets the value for VoucherSerial to be an explicit nil
func (o *VOUCHERESCHEMA) SetVoucherSerialNil() {
	o.VoucherSerial.Set(nil)
}

// UnsetVoucherSerial ensures that no value is present for VoucherSerial, not even an explicit nil
func (o *VOUCHERESCHEMA) UnsetVoucherSerial() {
	o.VoucherSerial.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VOUCHERESCHEMA) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VOUCHERESCHEMA) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *VOUCHERESCHEMA) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *VOUCHERESCHEMA) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *VOUCHERESCHEMA) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *VOUCHERESCHEMA) UnsetValue() {
	o.Value.Unset()
}

// GetExpiredDate returns the ExpiredDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VOUCHERESCHEMA) GetExpiredDate() string {
	if o == nil || IsNil(o.ExpiredDate.Get()) {
		var ret string
		return ret
	}
	return *o.ExpiredDate.Get()
}

// GetExpiredDateOk returns a tuple with the ExpiredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VOUCHERESCHEMA) GetExpiredDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredDate.Get(), o.ExpiredDate.IsSet()
}

// HasExpiredDate returns a boolean if a field has been set.
func (o *VOUCHERESCHEMA) HasExpiredDate() bool {
	if o != nil && o.ExpiredDate.IsSet() {
		return true
	}

	return false
}

// SetExpiredDate gets a reference to the given NullableString and assigns it to the ExpiredDate field.
func (o *VOUCHERESCHEMA) SetExpiredDate(v string) {
	o.ExpiredDate.Set(&v)
}
// SetExpiredDateNil sets the value for ExpiredDate to be an explicit nil
func (o *VOUCHERESCHEMA) SetExpiredDateNil() {
	o.ExpiredDate.Set(nil)
}

// UnsetExpiredDate ensures that no value is present for ExpiredDate, not even an explicit nil
func (o *VOUCHERESCHEMA) UnsetExpiredDate() {
	o.ExpiredDate.Unset()
}

func (o VOUCHERESCHEMA) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VOUCHERESCHEMA) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.VoucherLink.IsSet() {
		toSerialize["voucher_link"] = o.VoucherLink.Get()
	}
	if o.VoucherCoverLink.IsSet() {
		toSerialize["voucher_cover_link"] = o.VoucherCoverLink.Get()
	}
	if o.VoucherSerial.IsSet() {
		toSerialize["voucher_serial"] = o.VoucherSerial.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.ExpiredDate.IsSet() {
		toSerialize["expired_date"] = o.ExpiredDate.Get()
	}
	return toSerialize, nil
}

type NullableVOUCHERESCHEMA struct {
	value *VOUCHERESCHEMA
	isSet bool
}

func (v NullableVOUCHERESCHEMA) Get() *VOUCHERESCHEMA {
	return v.value
}

func (v *NullableVOUCHERESCHEMA) Set(val *VOUCHERESCHEMA) {
	v.value = val
	v.isSet = true
}

func (v NullableVOUCHERESCHEMA) IsSet() bool {
	return v.isSet
}

func (v *NullableVOUCHERESCHEMA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVOUCHERESCHEMA(val *VOUCHERESCHEMA) *NullableVOUCHERESCHEMA {
	return &NullableVOUCHERESCHEMA{value: val, isSet: true}
}

func (v NullableVOUCHERESCHEMA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVOUCHERESCHEMA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


