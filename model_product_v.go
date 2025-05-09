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

// checks if the PRODUCTV type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PRODUCTV{}

// PRODUCTV struct for PRODUCTV
type PRODUCTV struct {
	// Product Id
	ProductId NullableInt64 `json:"productId,omitempty"`
	// Product Name
	ProductNm NullableString `json:"productNm,omitempty"`
	// Link product image
	ProductImg NullableString `json:"productImg,omitempty"`
	// Brand id
	BrandId NullableInt64 `json:"brandId,omitempty"`
	// Brand name
	BrandNm NullableString `json:"brandNm,omitempty"`
	BrandServiceGuide NullableString `json:"brandServiceGuide,omitempty"`
	Link NullableString `json:"link,omitempty"`
	Price *PRODUCTPRICESCHEMA `json:"price,omitempty"`
	ProductDesc NullableString `json:"productDesc,omitempty"`
	Terms NullableString `json:"terms,omitempty"`
	ProductType NullableString `json:"productType,omitempty"`
	StoreList []STOREPRODUCTSCHEMA `json:"storeList,omitempty"`
}

// NewPRODUCTV instantiates a new PRODUCTV object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPRODUCTV() *PRODUCTV {
	this := PRODUCTV{}
	return &this
}

// NewPRODUCTVWithDefaults instantiates a new PRODUCTV object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPRODUCTVWithDefaults() *PRODUCTV {
	this := PRODUCTV{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PRODUCTV) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId.Get()) {
		var ret int64
		return ret
	}
	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRODUCTV) GetProductIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// HasProductId returns a boolean if a field has been set.
func (o *PRODUCTV) HasProductId() bool {
	if o != nil && o.ProductId.IsSet() {
		return true
	}

	return false
}

// SetProductId gets a reference to the given NullableInt64 and assigns it to the ProductId field.
func (o *PRODUCTV) SetProductId(v int64) {
	o.ProductId.Set(&v)
}
// SetProductIdNil sets the value for ProductId to be an explicit nil
func (o *PRODUCTV) SetProductIdNil() {
	o.ProductId.Set(nil)
}

// UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
func (o *PRODUCTV) UnsetProductId() {
	o.ProductId.Unset()
}

// GetProductNm returns the ProductNm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PRODUCTV) GetProductNm() string {
	if o == nil || IsNil(o.ProductNm.Get()) {
		var ret string
		return ret
	}
	return *o.ProductNm.Get()
}

// GetProductNmOk returns a tuple with the ProductNm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRODUCTV) GetProductNmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductNm.Get(), o.ProductNm.IsSet()
}

// HasProductNm returns a boolean if a field has been set.
func (o *PRODUCTV) HasProductNm() bool {
	if o != nil && o.ProductNm.IsSet() {
		return true
	}

	return false
}

// SetProductNm gets a reference to the given NullableString and assigns it to the ProductNm field.
func (o *PRODUCTV) SetProductNm(v string) {
	o.ProductNm.Set(&v)
}
// SetProductNmNil sets the value for ProductNm to be an explicit nil
func (o *PRODUCTV) SetProductNmNil() {
	o.ProductNm.Set(nil)
}

// UnsetProductNm ensures that no value is present for ProductNm, not even an explicit nil
func (o *PRODUCTV) UnsetProductNm() {
	o.ProductNm.Unset()
}

// GetProductImg returns the ProductImg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PRODUCTV) GetProductImg() string {
	if o == nil || IsNil(o.ProductImg.Get()) {
		var ret string
		return ret
	}
	return *o.ProductImg.Get()
}

// GetProductImgOk returns a tuple with the ProductImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRODUCTV) GetProductImgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductImg.Get(), o.ProductImg.IsSet()
}

// HasProductImg returns a boolean if a field has been set.
func (o *PRODUCTV) HasProductImg() bool {
	if o != nil && o.ProductImg.IsSet() {
		return true
	}

	return false
}

// SetProductImg gets a reference to the given NullableString and assigns it to the ProductImg field.
func (o *PRODUCTV) SetProductImg(v string) {
	o.ProductImg.Set(&v)
}
// SetProductImgNil sets the value for ProductImg to be an explicit nil
func (o *PRODUCTV) SetProductImgNil() {
	o.ProductImg.Set(nil)
}

// UnsetProductImg ensures that no value is present for ProductImg, not even an explicit nil
func (o *PRODUCTV) UnsetProductImg() {
	o.ProductImg.Unset()
}

// GetBrandId returns the BrandId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PRODUCTV) GetBrandId() int64 {
	if o == nil || IsNil(o.BrandId.Get()) {
		var ret int64
		return ret
	}
	return *o.BrandId.Get()
}

// GetBrandIdOk returns a tuple with the BrandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRODUCTV) GetBrandIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrandId.Get(), o.BrandId.IsSet()
}

// HasBrandId returns a boolean if a field has been set.
func (o *PRODUCTV) HasBrandId() bool {
	if o != nil && o.BrandId.IsSet() {
		return true
	}

	return false
}

// SetBrandId gets a reference to the given NullableInt64 and assigns it to the BrandId field.
func (o *PRODUCTV) SetBrandId(v int64) {
	o.BrandId.Set(&v)
}
// SetBrandIdNil sets the value for BrandId to be an explicit nil
func (o *PRODUCTV) SetBrandIdNil() {
	o.BrandId.Set(nil)
}

// UnsetBrandId ensures that no value is present for BrandId, not even an explicit nil
func (o *PRODUCTV) UnsetBrandId() {
	o.BrandId.Unset()
}

// GetBrandNm returns the BrandNm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PRODUCTV) GetBrandNm() string {
	if o == nil || IsNil(o.BrandNm.Get()) {
		var ret string
		return ret
	}
	return *o.BrandNm.Get()
}

// GetBrandNmOk returns a tuple with the BrandNm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRODUCTV) GetBrandNmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrandNm.Get(), o.BrandNm.IsSet()
}

// HasBrandNm returns a boolean if a field has been set.
func (o *PRODUCTV) HasBrandNm() bool {
	if o != nil && o.BrandNm.IsSet() {
		return true
	}

	return false
}

// SetBrandNm gets a reference to the given NullableString and assigns it to the BrandNm field.
func (o *PRODUCTV) SetBrandNm(v string) {
	o.BrandNm.Set(&v)
}
// SetBrandNmNil sets the value for BrandNm to be an explicit nil
func (o *PRODUCTV) SetBrandNmNil() {
	o.BrandNm.Set(nil)
}

// UnsetBrandNm ensures that no value is present for BrandNm, not even an explicit nil
func (o *PRODUCTV) UnsetBrandNm() {
	o.BrandNm.Unset()
}

// GetBrandServiceGuide returns the BrandServiceGuide field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PRODUCTV) GetBrandServiceGuide() string {
	if o == nil || IsNil(o.BrandServiceGuide.Get()) {
		var ret string
		return ret
	}
	return *o.BrandServiceGuide.Get()
}

// GetBrandServiceGuideOk returns a tuple with the BrandServiceGuide field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRODUCTV) GetBrandServiceGuideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrandServiceGuide.Get(), o.BrandServiceGuide.IsSet()
}

// HasBrandServiceGuide returns a boolean if a field has been set.
func (o *PRODUCTV) HasBrandServiceGuide() bool {
	if o != nil && o.BrandServiceGuide.IsSet() {
		return true
	}

	return false
}

// SetBrandServiceGuide gets a reference to the given NullableString and assigns it to the BrandServiceGuide field.
func (o *PRODUCTV) SetBrandServiceGuide(v string) {
	o.BrandServiceGuide.Set(&v)
}
// SetBrandServiceGuideNil sets the value for BrandServiceGuide to be an explicit nil
func (o *PRODUCTV) SetBrandServiceGuideNil() {
	o.BrandServiceGuide.Set(nil)
}

// UnsetBrandServiceGuide ensures that no value is present for BrandServiceGuide, not even an explicit nil
func (o *PRODUCTV) UnsetBrandServiceGuide() {
	o.BrandServiceGuide.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PRODUCTV) GetLink() string {
	if o == nil || IsNil(o.Link.Get()) {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRODUCTV) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *PRODUCTV) HasLink() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *PRODUCTV) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *PRODUCTV) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *PRODUCTV) UnsetLink() {
	o.Link.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PRODUCTV) GetPrice() PRODUCTPRICESCHEMA {
	if o == nil || IsNil(o.Price) {
		var ret PRODUCTPRICESCHEMA
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PRODUCTV) GetPriceOk() (*PRODUCTPRICESCHEMA, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PRODUCTV) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given PRODUCTPRICESCHEMA and assigns it to the Price field.
func (o *PRODUCTV) SetPrice(v PRODUCTPRICESCHEMA) {
	o.Price = &v
}

// GetProductDesc returns the ProductDesc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PRODUCTV) GetProductDesc() string {
	if o == nil || IsNil(o.ProductDesc.Get()) {
		var ret string
		return ret
	}
	return *o.ProductDesc.Get()
}

// GetProductDescOk returns a tuple with the ProductDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRODUCTV) GetProductDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductDesc.Get(), o.ProductDesc.IsSet()
}

// HasProductDesc returns a boolean if a field has been set.
func (o *PRODUCTV) HasProductDesc() bool {
	if o != nil && o.ProductDesc.IsSet() {
		return true
	}

	return false
}

// SetProductDesc gets a reference to the given NullableString and assigns it to the ProductDesc field.
func (o *PRODUCTV) SetProductDesc(v string) {
	o.ProductDesc.Set(&v)
}
// SetProductDescNil sets the value for ProductDesc to be an explicit nil
func (o *PRODUCTV) SetProductDescNil() {
	o.ProductDesc.Set(nil)
}

// UnsetProductDesc ensures that no value is present for ProductDesc, not even an explicit nil
func (o *PRODUCTV) UnsetProductDesc() {
	o.ProductDesc.Unset()
}

// GetTerms returns the Terms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PRODUCTV) GetTerms() string {
	if o == nil || IsNil(o.Terms.Get()) {
		var ret string
		return ret
	}
	return *o.Terms.Get()
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRODUCTV) GetTermsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Terms.Get(), o.Terms.IsSet()
}

// HasTerms returns a boolean if a field has been set.
func (o *PRODUCTV) HasTerms() bool {
	if o != nil && o.Terms.IsSet() {
		return true
	}

	return false
}

// SetTerms gets a reference to the given NullableString and assigns it to the Terms field.
func (o *PRODUCTV) SetTerms(v string) {
	o.Terms.Set(&v)
}
// SetTermsNil sets the value for Terms to be an explicit nil
func (o *PRODUCTV) SetTermsNil() {
	o.Terms.Set(nil)
}

// UnsetTerms ensures that no value is present for Terms, not even an explicit nil
func (o *PRODUCTV) UnsetTerms() {
	o.Terms.Unset()
}

// GetProductType returns the ProductType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PRODUCTV) GetProductType() string {
	if o == nil || IsNil(o.ProductType.Get()) {
		var ret string
		return ret
	}
	return *o.ProductType.Get()
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PRODUCTV) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductType.Get(), o.ProductType.IsSet()
}

// HasProductType returns a boolean if a field has been set.
func (o *PRODUCTV) HasProductType() bool {
	if o != nil && o.ProductType.IsSet() {
		return true
	}

	return false
}

// SetProductType gets a reference to the given NullableString and assigns it to the ProductType field.
func (o *PRODUCTV) SetProductType(v string) {
	o.ProductType.Set(&v)
}
// SetProductTypeNil sets the value for ProductType to be an explicit nil
func (o *PRODUCTV) SetProductTypeNil() {
	o.ProductType.Set(nil)
}

// UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
func (o *PRODUCTV) UnsetProductType() {
	o.ProductType.Unset()
}

// GetStoreList returns the StoreList field value if set, zero value otherwise.
func (o *PRODUCTV) GetStoreList() []STOREPRODUCTSCHEMA {
	if o == nil || IsNil(o.StoreList) {
		var ret []STOREPRODUCTSCHEMA
		return ret
	}
	return o.StoreList
}

// GetStoreListOk returns a tuple with the StoreList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PRODUCTV) GetStoreListOk() ([]STOREPRODUCTSCHEMA, bool) {
	if o == nil || IsNil(o.StoreList) {
		return nil, false
	}
	return o.StoreList, true
}

// HasStoreList returns a boolean if a field has been set.
func (o *PRODUCTV) HasStoreList() bool {
	if o != nil && !IsNil(o.StoreList) {
		return true
	}

	return false
}

// SetStoreList gets a reference to the given []STOREPRODUCTSCHEMA and assigns it to the StoreList field.
func (o *PRODUCTV) SetStoreList(v []STOREPRODUCTSCHEMA) {
	o.StoreList = v
}

func (o PRODUCTV) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PRODUCTV) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductId.IsSet() {
		toSerialize["productId"] = o.ProductId.Get()
	}
	if o.ProductNm.IsSet() {
		toSerialize["productNm"] = o.ProductNm.Get()
	}
	if o.ProductImg.IsSet() {
		toSerialize["productImg"] = o.ProductImg.Get()
	}
	if o.BrandId.IsSet() {
		toSerialize["brandId"] = o.BrandId.Get()
	}
	if o.BrandNm.IsSet() {
		toSerialize["brandNm"] = o.BrandNm.Get()
	}
	if o.BrandServiceGuide.IsSet() {
		toSerialize["brandServiceGuide"] = o.BrandServiceGuide.Get()
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if o.ProductDesc.IsSet() {
		toSerialize["productDesc"] = o.ProductDesc.Get()
	}
	if o.Terms.IsSet() {
		toSerialize["terms"] = o.Terms.Get()
	}
	if o.ProductType.IsSet() {
		toSerialize["productType"] = o.ProductType.Get()
	}
	if !IsNil(o.StoreList) {
		toSerialize["storeList"] = o.StoreList
	}
	return toSerialize, nil
}

type NullablePRODUCTV struct {
	value *PRODUCTV
	isSet bool
}

func (v NullablePRODUCTV) Get() *PRODUCTV {
	return v.value
}

func (v *NullablePRODUCTV) Set(val *PRODUCTV) {
	v.value = val
	v.isSet = true
}

func (v NullablePRODUCTV) IsSet() bool {
	return v.isSet
}

func (v *NullablePRODUCTV) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePRODUCTV(val *PRODUCTV) *NullablePRODUCTV {
	return &NullablePRODUCTV{value: val, isSet: true}
}

func (v NullablePRODUCTV) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePRODUCTV) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


