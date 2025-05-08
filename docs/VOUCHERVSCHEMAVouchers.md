# VOUCHERVSCHEMAVouchers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **NullableInt64** | Product Id | [optional] 
**ProductNm** | Pointer to **NullableString** | Product Name | [optional] 
**ProductImg** | Pointer to **NullableString** | Link product image | [optional] 
**BrandId** | Pointer to **NullableInt64** | Brand id | [optional] 
**BrandNm** | Pointer to **NullableString** | Brand name | [optional] 
**BrandServiceGuide** | Pointer to **NullableString** |  | [optional] 
**Link** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to [**PRODUCTPRICESCHEMA**](PRODUCTPRICESCHEMA.md) |  | [optional] 
**ProductDesc** | Pointer to **NullableString** |  | [optional] 
**Terms** | Pointer to **NullableString** |  | [optional] 
**ProductType** | Pointer to **NullableString** |  | [optional] 
**StoreList** | Pointer to [**[]STOREPRODUCTSCHEMA**](STOREPRODUCTSCHEMA.md) |  | [optional] 

## Methods

### NewVOUCHERVSCHEMAVouchers

`func NewVOUCHERVSCHEMAVouchers() *VOUCHERVSCHEMAVouchers`

NewVOUCHERVSCHEMAVouchers instantiates a new VOUCHERVSCHEMAVouchers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVOUCHERVSCHEMAVouchersWithDefaults

`func NewVOUCHERVSCHEMAVouchersWithDefaults() *VOUCHERVSCHEMAVouchers`

NewVOUCHERVSCHEMAVouchersWithDefaults instantiates a new VOUCHERVSCHEMAVouchers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *VOUCHERVSCHEMAVouchers) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *VOUCHERVSCHEMAVouchers) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *VOUCHERVSCHEMAVouchers) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *VOUCHERVSCHEMAVouchers) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *VOUCHERVSCHEMAVouchers) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *VOUCHERVSCHEMAVouchers) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetProductNm

`func (o *VOUCHERVSCHEMAVouchers) GetProductNm() string`

GetProductNm returns the ProductNm field if non-nil, zero value otherwise.

### GetProductNmOk

`func (o *VOUCHERVSCHEMAVouchers) GetProductNmOk() (*string, bool)`

GetProductNmOk returns a tuple with the ProductNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductNm

`func (o *VOUCHERVSCHEMAVouchers) SetProductNm(v string)`

SetProductNm sets ProductNm field to given value.

### HasProductNm

`func (o *VOUCHERVSCHEMAVouchers) HasProductNm() bool`

HasProductNm returns a boolean if a field has been set.

### SetProductNmNil

`func (o *VOUCHERVSCHEMAVouchers) SetProductNmNil(b bool)`

 SetProductNmNil sets the value for ProductNm to be an explicit nil

### UnsetProductNm
`func (o *VOUCHERVSCHEMAVouchers) UnsetProductNm()`

UnsetProductNm ensures that no value is present for ProductNm, not even an explicit nil
### GetProductImg

`func (o *VOUCHERVSCHEMAVouchers) GetProductImg() string`

GetProductImg returns the ProductImg field if non-nil, zero value otherwise.

### GetProductImgOk

`func (o *VOUCHERVSCHEMAVouchers) GetProductImgOk() (*string, bool)`

GetProductImgOk returns a tuple with the ProductImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImg

`func (o *VOUCHERVSCHEMAVouchers) SetProductImg(v string)`

SetProductImg sets ProductImg field to given value.

### HasProductImg

`func (o *VOUCHERVSCHEMAVouchers) HasProductImg() bool`

HasProductImg returns a boolean if a field has been set.

### SetProductImgNil

`func (o *VOUCHERVSCHEMAVouchers) SetProductImgNil(b bool)`

 SetProductImgNil sets the value for ProductImg to be an explicit nil

### UnsetProductImg
`func (o *VOUCHERVSCHEMAVouchers) UnsetProductImg()`

UnsetProductImg ensures that no value is present for ProductImg, not even an explicit nil
### GetBrandId

`func (o *VOUCHERVSCHEMAVouchers) GetBrandId() int64`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *VOUCHERVSCHEMAVouchers) GetBrandIdOk() (*int64, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *VOUCHERVSCHEMAVouchers) SetBrandId(v int64)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *VOUCHERVSCHEMAVouchers) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### SetBrandIdNil

`func (o *VOUCHERVSCHEMAVouchers) SetBrandIdNil(b bool)`

 SetBrandIdNil sets the value for BrandId to be an explicit nil

### UnsetBrandId
`func (o *VOUCHERVSCHEMAVouchers) UnsetBrandId()`

UnsetBrandId ensures that no value is present for BrandId, not even an explicit nil
### GetBrandNm

`func (o *VOUCHERVSCHEMAVouchers) GetBrandNm() string`

GetBrandNm returns the BrandNm field if non-nil, zero value otherwise.

### GetBrandNmOk

`func (o *VOUCHERVSCHEMAVouchers) GetBrandNmOk() (*string, bool)`

GetBrandNmOk returns a tuple with the BrandNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNm

`func (o *VOUCHERVSCHEMAVouchers) SetBrandNm(v string)`

SetBrandNm sets BrandNm field to given value.

### HasBrandNm

`func (o *VOUCHERVSCHEMAVouchers) HasBrandNm() bool`

HasBrandNm returns a boolean if a field has been set.

### SetBrandNmNil

`func (o *VOUCHERVSCHEMAVouchers) SetBrandNmNil(b bool)`

 SetBrandNmNil sets the value for BrandNm to be an explicit nil

### UnsetBrandNm
`func (o *VOUCHERVSCHEMAVouchers) UnsetBrandNm()`

UnsetBrandNm ensures that no value is present for BrandNm, not even an explicit nil
### GetBrandServiceGuide

`func (o *VOUCHERVSCHEMAVouchers) GetBrandServiceGuide() string`

GetBrandServiceGuide returns the BrandServiceGuide field if non-nil, zero value otherwise.

### GetBrandServiceGuideOk

`func (o *VOUCHERVSCHEMAVouchers) GetBrandServiceGuideOk() (*string, bool)`

GetBrandServiceGuideOk returns a tuple with the BrandServiceGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandServiceGuide

`func (o *VOUCHERVSCHEMAVouchers) SetBrandServiceGuide(v string)`

SetBrandServiceGuide sets BrandServiceGuide field to given value.

### HasBrandServiceGuide

`func (o *VOUCHERVSCHEMAVouchers) HasBrandServiceGuide() bool`

HasBrandServiceGuide returns a boolean if a field has been set.

### SetBrandServiceGuideNil

`func (o *VOUCHERVSCHEMAVouchers) SetBrandServiceGuideNil(b bool)`

 SetBrandServiceGuideNil sets the value for BrandServiceGuide to be an explicit nil

### UnsetBrandServiceGuide
`func (o *VOUCHERVSCHEMAVouchers) UnsetBrandServiceGuide()`

UnsetBrandServiceGuide ensures that no value is present for BrandServiceGuide, not even an explicit nil
### GetLink

`func (o *VOUCHERVSCHEMAVouchers) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *VOUCHERVSCHEMAVouchers) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *VOUCHERVSCHEMAVouchers) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *VOUCHERVSCHEMAVouchers) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *VOUCHERVSCHEMAVouchers) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *VOUCHERVSCHEMAVouchers) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetPrice

`func (o *VOUCHERVSCHEMAVouchers) GetPrice() PRODUCTPRICESCHEMA`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *VOUCHERVSCHEMAVouchers) GetPriceOk() (*PRODUCTPRICESCHEMA, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *VOUCHERVSCHEMAVouchers) SetPrice(v PRODUCTPRICESCHEMA)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *VOUCHERVSCHEMAVouchers) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProductDesc

`func (o *VOUCHERVSCHEMAVouchers) GetProductDesc() string`

GetProductDesc returns the ProductDesc field if non-nil, zero value otherwise.

### GetProductDescOk

`func (o *VOUCHERVSCHEMAVouchers) GetProductDescOk() (*string, bool)`

GetProductDescOk returns a tuple with the ProductDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDesc

`func (o *VOUCHERVSCHEMAVouchers) SetProductDesc(v string)`

SetProductDesc sets ProductDesc field to given value.

### HasProductDesc

`func (o *VOUCHERVSCHEMAVouchers) HasProductDesc() bool`

HasProductDesc returns a boolean if a field has been set.

### SetProductDescNil

`func (o *VOUCHERVSCHEMAVouchers) SetProductDescNil(b bool)`

 SetProductDescNil sets the value for ProductDesc to be an explicit nil

### UnsetProductDesc
`func (o *VOUCHERVSCHEMAVouchers) UnsetProductDesc()`

UnsetProductDesc ensures that no value is present for ProductDesc, not even an explicit nil
### GetTerms

`func (o *VOUCHERVSCHEMAVouchers) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *VOUCHERVSCHEMAVouchers) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *VOUCHERVSCHEMAVouchers) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *VOUCHERVSCHEMAVouchers) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### SetTermsNil

`func (o *VOUCHERVSCHEMAVouchers) SetTermsNil(b bool)`

 SetTermsNil sets the value for Terms to be an explicit nil

### UnsetTerms
`func (o *VOUCHERVSCHEMAVouchers) UnsetTerms()`

UnsetTerms ensures that no value is present for Terms, not even an explicit nil
### GetProductType

`func (o *VOUCHERVSCHEMAVouchers) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *VOUCHERVSCHEMAVouchers) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *VOUCHERVSCHEMAVouchers) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *VOUCHERVSCHEMAVouchers) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### SetProductTypeNil

`func (o *VOUCHERVSCHEMAVouchers) SetProductTypeNil(b bool)`

 SetProductTypeNil sets the value for ProductType to be an explicit nil

### UnsetProductType
`func (o *VOUCHERVSCHEMAVouchers) UnsetProductType()`

UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
### GetStoreList

`func (o *VOUCHERVSCHEMAVouchers) GetStoreList() []STOREPRODUCTSCHEMA`

GetStoreList returns the StoreList field if non-nil, zero value otherwise.

### GetStoreListOk

`func (o *VOUCHERVSCHEMAVouchers) GetStoreListOk() (*[]STOREPRODUCTSCHEMA, bool)`

GetStoreListOk returns a tuple with the StoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreList

`func (o *VOUCHERVSCHEMAVouchers) SetStoreList(v []STOREPRODUCTSCHEMA)`

SetStoreList sets StoreList field to given value.

### HasStoreList

`func (o *VOUCHERVSCHEMAVouchers) HasStoreList() bool`

HasStoreList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


