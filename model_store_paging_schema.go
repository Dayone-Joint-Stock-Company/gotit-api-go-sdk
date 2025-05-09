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

// checks if the STOREPAGINGSCHEMA type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &STOREPAGINGSCHEMA{}

// STOREPAGINGSCHEMA struct for STOREPAGINGSCHEMA
type STOREPAGINGSCHEMA struct {
	// Number of page
	Page NullableInt64 `json:"page,omitempty"`
	// Size of page
	PageSize NullableInt64 `json:"pageSize,omitempty"`
	// Total page
	TotalPage NullableInt64 `json:"totalPage,omitempty"`
}

// NewSTOREPAGINGSCHEMA instantiates a new STOREPAGINGSCHEMA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSTOREPAGINGSCHEMA() *STOREPAGINGSCHEMA {
	this := STOREPAGINGSCHEMA{}
	return &this
}

// NewSTOREPAGINGSCHEMAWithDefaults instantiates a new STOREPAGINGSCHEMA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSTOREPAGINGSCHEMAWithDefaults() *STOREPAGINGSCHEMA {
	this := STOREPAGINGSCHEMA{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *STOREPAGINGSCHEMA) GetPage() int64 {
	if o == nil || IsNil(o.Page.Get()) {
		var ret int64
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *STOREPAGINGSCHEMA) GetPageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field has been set.
func (o *STOREPAGINGSCHEMA) HasPage() bool {
	if o != nil && o.Page.IsSet() {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableInt64 and assigns it to the Page field.
func (o *STOREPAGINGSCHEMA) SetPage(v int64) {
	o.Page.Set(&v)
}
// SetPageNil sets the value for Page to be an explicit nil
func (o *STOREPAGINGSCHEMA) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *STOREPAGINGSCHEMA) UnsetPage() {
	o.Page.Unset()
}

// GetPageSize returns the PageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *STOREPAGINGSCHEMA) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize.Get()) {
		var ret int64
		return ret
	}
	return *o.PageSize.Get()
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *STOREPAGINGSCHEMA) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageSize.Get(), o.PageSize.IsSet()
}

// HasPageSize returns a boolean if a field has been set.
func (o *STOREPAGINGSCHEMA) HasPageSize() bool {
	if o != nil && o.PageSize.IsSet() {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given NullableInt64 and assigns it to the PageSize field.
func (o *STOREPAGINGSCHEMA) SetPageSize(v int64) {
	o.PageSize.Set(&v)
}
// SetPageSizeNil sets the value for PageSize to be an explicit nil
func (o *STOREPAGINGSCHEMA) SetPageSizeNil() {
	o.PageSize.Set(nil)
}

// UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
func (o *STOREPAGINGSCHEMA) UnsetPageSize() {
	o.PageSize.Unset()
}

// GetTotalPage returns the TotalPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *STOREPAGINGSCHEMA) GetTotalPage() int64 {
	if o == nil || IsNil(o.TotalPage.Get()) {
		var ret int64
		return ret
	}
	return *o.TotalPage.Get()
}

// GetTotalPageOk returns a tuple with the TotalPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *STOREPAGINGSCHEMA) GetTotalPageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalPage.Get(), o.TotalPage.IsSet()
}

// HasTotalPage returns a boolean if a field has been set.
func (o *STOREPAGINGSCHEMA) HasTotalPage() bool {
	if o != nil && o.TotalPage.IsSet() {
		return true
	}

	return false
}

// SetTotalPage gets a reference to the given NullableInt64 and assigns it to the TotalPage field.
func (o *STOREPAGINGSCHEMA) SetTotalPage(v int64) {
	o.TotalPage.Set(&v)
}
// SetTotalPageNil sets the value for TotalPage to be an explicit nil
func (o *STOREPAGINGSCHEMA) SetTotalPageNil() {
	o.TotalPage.Set(nil)
}

// UnsetTotalPage ensures that no value is present for TotalPage, not even an explicit nil
func (o *STOREPAGINGSCHEMA) UnsetTotalPage() {
	o.TotalPage.Unset()
}

func (o STOREPAGINGSCHEMA) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o STOREPAGINGSCHEMA) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	if o.PageSize.IsSet() {
		toSerialize["pageSize"] = o.PageSize.Get()
	}
	if o.TotalPage.IsSet() {
		toSerialize["totalPage"] = o.TotalPage.Get()
	}
	return toSerialize, nil
}

type NullableSTOREPAGINGSCHEMA struct {
	value *STOREPAGINGSCHEMA
	isSet bool
}

func (v NullableSTOREPAGINGSCHEMA) Get() *STOREPAGINGSCHEMA {
	return v.value
}

func (v *NullableSTOREPAGINGSCHEMA) Set(val *STOREPAGINGSCHEMA) {
	v.value = val
	v.isSet = true
}

func (v NullableSTOREPAGINGSCHEMA) IsSet() bool {
	return v.isSet
}

func (v *NullableSTOREPAGINGSCHEMA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSTOREPAGINGSCHEMA(val *STOREPAGINGSCHEMA) *NullableSTOREPAGINGSCHEMA {
	return &NullableSTOREPAGINGSCHEMA{value: val, isSet: true}
}

func (v NullableSTOREPAGINGSCHEMA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSTOREPAGINGSCHEMA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


