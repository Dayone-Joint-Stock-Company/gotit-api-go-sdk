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

// checks if the VOUCHERSENDZNSRESPONSEData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VOUCHERSENDZNSRESPONSEData{}

// VOUCHERSENDZNSRESPONSEData Response data containing the transactionId
type VOUCHERSENDZNSRESPONSEData struct {
	// Unique identifier for the transaction
	TransactionId *string `json:"transactionId,omitempty"`
}

// NewVOUCHERSENDZNSRESPONSEData instantiates a new VOUCHERSENDZNSRESPONSEData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVOUCHERSENDZNSRESPONSEData() *VOUCHERSENDZNSRESPONSEData {
	this := VOUCHERSENDZNSRESPONSEData{}
	return &this
}

// NewVOUCHERSENDZNSRESPONSEDataWithDefaults instantiates a new VOUCHERSENDZNSRESPONSEData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVOUCHERSENDZNSRESPONSEDataWithDefaults() *VOUCHERSENDZNSRESPONSEData {
	this := VOUCHERSENDZNSRESPONSEData{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *VOUCHERSENDZNSRESPONSEData) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VOUCHERSENDZNSRESPONSEData) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *VOUCHERSENDZNSRESPONSEData) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *VOUCHERSENDZNSRESPONSEData) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o VOUCHERSENDZNSRESPONSEData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VOUCHERSENDZNSRESPONSEData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	return toSerialize, nil
}

type NullableVOUCHERSENDZNSRESPONSEData struct {
	value *VOUCHERSENDZNSRESPONSEData
	isSet bool
}

func (v NullableVOUCHERSENDZNSRESPONSEData) Get() *VOUCHERSENDZNSRESPONSEData {
	return v.value
}

func (v *NullableVOUCHERSENDZNSRESPONSEData) Set(val *VOUCHERSENDZNSRESPONSEData) {
	v.value = val
	v.isSet = true
}

func (v NullableVOUCHERSENDZNSRESPONSEData) IsSet() bool {
	return v.isSet
}

func (v *NullableVOUCHERSENDZNSRESPONSEData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVOUCHERSENDZNSRESPONSEData(val *VOUCHERSENDZNSRESPONSEData) *NullableVOUCHERSENDZNSRESPONSEData {
	return &NullableVOUCHERSENDZNSRESPONSEData{value: val, isSet: true}
}

func (v NullableVOUCHERSENDZNSRESPONSEData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVOUCHERSENDZNSRESPONSEData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


