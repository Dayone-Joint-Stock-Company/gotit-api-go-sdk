# PRODUCTSDETAIL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **NullableInt64** | Product Id | [optional] 
**ProductNm** | Pointer to **NullableString** | Product Name | [optional] 
**ProductImg** | Pointer to **NullableString** | Link product image | [optional] 
**ProductDesc** | Pointer to **NullableString** | Product Description | [optional] 
**ProductShortDesc** | Pointer to **NullableString** | Product Short Description | [optional] 
**Terms** | Pointer to **NullableString** | Terms of use | [optional] 
**BrandId** | Pointer to **NullableInt64** | Brand id | [optional] 
**BrandNm** | Pointer to **NullableString** | Brand name | [optional] 
**BrandLogo** | Pointer to **NullableString** | Link to brand logo image | [optional] 
**BrandServiceGuide** | Pointer to **NullableString** | T&amp;C of brand | [optional] 
**CategoryId** | Pointer to **NullableInt64** | Category Id | [optional] 
**CategoryNm** | Pointer to **NullableString** | Category Name | [optional] 
**CategoryImg** | Pointer to **NullableString** | Category Image | [optional] 
**ProductType** | Pointer to **NullableString** | c (cash) hoặc i (item) | [optional] 
**Prices** | Pointer to [**[]PRODUCTPRICESCHEMA**](PRODUCTPRICESCHEMA.md) |  | [optional] 
**Link** | Pointer to **NullableString** |  | [optional] 
**StoreList** | Pointer to [**[]STORESSCHEMA**](STORESSCHEMA.md) |  | [optional] 

## Methods

### NewPRODUCTSDETAIL

`func NewPRODUCTSDETAIL() *PRODUCTSDETAIL`

NewPRODUCTSDETAIL instantiates a new PRODUCTSDETAIL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPRODUCTSDETAILWithDefaults

`func NewPRODUCTSDETAILWithDefaults() *PRODUCTSDETAIL`

NewPRODUCTSDETAILWithDefaults instantiates a new PRODUCTSDETAIL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *PRODUCTSDETAIL) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PRODUCTSDETAIL) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PRODUCTSDETAIL) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PRODUCTSDETAIL) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *PRODUCTSDETAIL) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *PRODUCTSDETAIL) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetProductNm

`func (o *PRODUCTSDETAIL) GetProductNm() string`

GetProductNm returns the ProductNm field if non-nil, zero value otherwise.

### GetProductNmOk

`func (o *PRODUCTSDETAIL) GetProductNmOk() (*string, bool)`

GetProductNmOk returns a tuple with the ProductNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductNm

`func (o *PRODUCTSDETAIL) SetProductNm(v string)`

SetProductNm sets ProductNm field to given value.

### HasProductNm

`func (o *PRODUCTSDETAIL) HasProductNm() bool`

HasProductNm returns a boolean if a field has been set.

### SetProductNmNil

`func (o *PRODUCTSDETAIL) SetProductNmNil(b bool)`

 SetProductNmNil sets the value for ProductNm to be an explicit nil

### UnsetProductNm
`func (o *PRODUCTSDETAIL) UnsetProductNm()`

UnsetProductNm ensures that no value is present for ProductNm, not even an explicit nil
### GetProductImg

`func (o *PRODUCTSDETAIL) GetProductImg() string`

GetProductImg returns the ProductImg field if non-nil, zero value otherwise.

### GetProductImgOk

`func (o *PRODUCTSDETAIL) GetProductImgOk() (*string, bool)`

GetProductImgOk returns a tuple with the ProductImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImg

`func (o *PRODUCTSDETAIL) SetProductImg(v string)`

SetProductImg sets ProductImg field to given value.

### HasProductImg

`func (o *PRODUCTSDETAIL) HasProductImg() bool`

HasProductImg returns a boolean if a field has been set.

### SetProductImgNil

`func (o *PRODUCTSDETAIL) SetProductImgNil(b bool)`

 SetProductImgNil sets the value for ProductImg to be an explicit nil

### UnsetProductImg
`func (o *PRODUCTSDETAIL) UnsetProductImg()`

UnsetProductImg ensures that no value is present for ProductImg, not even an explicit nil
### GetProductDesc

`func (o *PRODUCTSDETAIL) GetProductDesc() string`

GetProductDesc returns the ProductDesc field if non-nil, zero value otherwise.

### GetProductDescOk

`func (o *PRODUCTSDETAIL) GetProductDescOk() (*string, bool)`

GetProductDescOk returns a tuple with the ProductDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDesc

`func (o *PRODUCTSDETAIL) SetProductDesc(v string)`

SetProductDesc sets ProductDesc field to given value.

### HasProductDesc

`func (o *PRODUCTSDETAIL) HasProductDesc() bool`

HasProductDesc returns a boolean if a field has been set.

### SetProductDescNil

`func (o *PRODUCTSDETAIL) SetProductDescNil(b bool)`

 SetProductDescNil sets the value for ProductDesc to be an explicit nil

### UnsetProductDesc
`func (o *PRODUCTSDETAIL) UnsetProductDesc()`

UnsetProductDesc ensures that no value is present for ProductDesc, not even an explicit nil
### GetProductShortDesc

`func (o *PRODUCTSDETAIL) GetProductShortDesc() string`

GetProductShortDesc returns the ProductShortDesc field if non-nil, zero value otherwise.

### GetProductShortDescOk

`func (o *PRODUCTSDETAIL) GetProductShortDescOk() (*string, bool)`

GetProductShortDescOk returns a tuple with the ProductShortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductShortDesc

`func (o *PRODUCTSDETAIL) SetProductShortDesc(v string)`

SetProductShortDesc sets ProductShortDesc field to given value.

### HasProductShortDesc

`func (o *PRODUCTSDETAIL) HasProductShortDesc() bool`

HasProductShortDesc returns a boolean if a field has been set.

### SetProductShortDescNil

`func (o *PRODUCTSDETAIL) SetProductShortDescNil(b bool)`

 SetProductShortDescNil sets the value for ProductShortDesc to be an explicit nil

### UnsetProductShortDesc
`func (o *PRODUCTSDETAIL) UnsetProductShortDesc()`

UnsetProductShortDesc ensures that no value is present for ProductShortDesc, not even an explicit nil
### GetTerms

`func (o *PRODUCTSDETAIL) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *PRODUCTSDETAIL) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *PRODUCTSDETAIL) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *PRODUCTSDETAIL) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### SetTermsNil

`func (o *PRODUCTSDETAIL) SetTermsNil(b bool)`

 SetTermsNil sets the value for Terms to be an explicit nil

### UnsetTerms
`func (o *PRODUCTSDETAIL) UnsetTerms()`

UnsetTerms ensures that no value is present for Terms, not even an explicit nil
### GetBrandId

`func (o *PRODUCTSDETAIL) GetBrandId() int64`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *PRODUCTSDETAIL) GetBrandIdOk() (*int64, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *PRODUCTSDETAIL) SetBrandId(v int64)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *PRODUCTSDETAIL) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### SetBrandIdNil

`func (o *PRODUCTSDETAIL) SetBrandIdNil(b bool)`

 SetBrandIdNil sets the value for BrandId to be an explicit nil

### UnsetBrandId
`func (o *PRODUCTSDETAIL) UnsetBrandId()`

UnsetBrandId ensures that no value is present for BrandId, not even an explicit nil
### GetBrandNm

`func (o *PRODUCTSDETAIL) GetBrandNm() string`

GetBrandNm returns the BrandNm field if non-nil, zero value otherwise.

### GetBrandNmOk

`func (o *PRODUCTSDETAIL) GetBrandNmOk() (*string, bool)`

GetBrandNmOk returns a tuple with the BrandNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNm

`func (o *PRODUCTSDETAIL) SetBrandNm(v string)`

SetBrandNm sets BrandNm field to given value.

### HasBrandNm

`func (o *PRODUCTSDETAIL) HasBrandNm() bool`

HasBrandNm returns a boolean if a field has been set.

### SetBrandNmNil

`func (o *PRODUCTSDETAIL) SetBrandNmNil(b bool)`

 SetBrandNmNil sets the value for BrandNm to be an explicit nil

### UnsetBrandNm
`func (o *PRODUCTSDETAIL) UnsetBrandNm()`

UnsetBrandNm ensures that no value is present for BrandNm, not even an explicit nil
### GetBrandLogo

`func (o *PRODUCTSDETAIL) GetBrandLogo() string`

GetBrandLogo returns the BrandLogo field if non-nil, zero value otherwise.

### GetBrandLogoOk

`func (o *PRODUCTSDETAIL) GetBrandLogoOk() (*string, bool)`

GetBrandLogoOk returns a tuple with the BrandLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandLogo

`func (o *PRODUCTSDETAIL) SetBrandLogo(v string)`

SetBrandLogo sets BrandLogo field to given value.

### HasBrandLogo

`func (o *PRODUCTSDETAIL) HasBrandLogo() bool`

HasBrandLogo returns a boolean if a field has been set.

### SetBrandLogoNil

`func (o *PRODUCTSDETAIL) SetBrandLogoNil(b bool)`

 SetBrandLogoNil sets the value for BrandLogo to be an explicit nil

### UnsetBrandLogo
`func (o *PRODUCTSDETAIL) UnsetBrandLogo()`

UnsetBrandLogo ensures that no value is present for BrandLogo, not even an explicit nil
### GetBrandServiceGuide

`func (o *PRODUCTSDETAIL) GetBrandServiceGuide() string`

GetBrandServiceGuide returns the BrandServiceGuide field if non-nil, zero value otherwise.

### GetBrandServiceGuideOk

`func (o *PRODUCTSDETAIL) GetBrandServiceGuideOk() (*string, bool)`

GetBrandServiceGuideOk returns a tuple with the BrandServiceGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandServiceGuide

`func (o *PRODUCTSDETAIL) SetBrandServiceGuide(v string)`

SetBrandServiceGuide sets BrandServiceGuide field to given value.

### HasBrandServiceGuide

`func (o *PRODUCTSDETAIL) HasBrandServiceGuide() bool`

HasBrandServiceGuide returns a boolean if a field has been set.

### SetBrandServiceGuideNil

`func (o *PRODUCTSDETAIL) SetBrandServiceGuideNil(b bool)`

 SetBrandServiceGuideNil sets the value for BrandServiceGuide to be an explicit nil

### UnsetBrandServiceGuide
`func (o *PRODUCTSDETAIL) UnsetBrandServiceGuide()`

UnsetBrandServiceGuide ensures that no value is present for BrandServiceGuide, not even an explicit nil
### GetCategoryId

`func (o *PRODUCTSDETAIL) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *PRODUCTSDETAIL) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *PRODUCTSDETAIL) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *PRODUCTSDETAIL) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *PRODUCTSDETAIL) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *PRODUCTSDETAIL) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetCategoryNm

`func (o *PRODUCTSDETAIL) GetCategoryNm() string`

GetCategoryNm returns the CategoryNm field if non-nil, zero value otherwise.

### GetCategoryNmOk

`func (o *PRODUCTSDETAIL) GetCategoryNmOk() (*string, bool)`

GetCategoryNmOk returns a tuple with the CategoryNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryNm

`func (o *PRODUCTSDETAIL) SetCategoryNm(v string)`

SetCategoryNm sets CategoryNm field to given value.

### HasCategoryNm

`func (o *PRODUCTSDETAIL) HasCategoryNm() bool`

HasCategoryNm returns a boolean if a field has been set.

### SetCategoryNmNil

`func (o *PRODUCTSDETAIL) SetCategoryNmNil(b bool)`

 SetCategoryNmNil sets the value for CategoryNm to be an explicit nil

### UnsetCategoryNm
`func (o *PRODUCTSDETAIL) UnsetCategoryNm()`

UnsetCategoryNm ensures that no value is present for CategoryNm, not even an explicit nil
### GetCategoryImg

`func (o *PRODUCTSDETAIL) GetCategoryImg() string`

GetCategoryImg returns the CategoryImg field if non-nil, zero value otherwise.

### GetCategoryImgOk

`func (o *PRODUCTSDETAIL) GetCategoryImgOk() (*string, bool)`

GetCategoryImgOk returns a tuple with the CategoryImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryImg

`func (o *PRODUCTSDETAIL) SetCategoryImg(v string)`

SetCategoryImg sets CategoryImg field to given value.

### HasCategoryImg

`func (o *PRODUCTSDETAIL) HasCategoryImg() bool`

HasCategoryImg returns a boolean if a field has been set.

### SetCategoryImgNil

`func (o *PRODUCTSDETAIL) SetCategoryImgNil(b bool)`

 SetCategoryImgNil sets the value for CategoryImg to be an explicit nil

### UnsetCategoryImg
`func (o *PRODUCTSDETAIL) UnsetCategoryImg()`

UnsetCategoryImg ensures that no value is present for CategoryImg, not even an explicit nil
### GetProductType

`func (o *PRODUCTSDETAIL) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PRODUCTSDETAIL) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PRODUCTSDETAIL) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *PRODUCTSDETAIL) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### SetProductTypeNil

`func (o *PRODUCTSDETAIL) SetProductTypeNil(b bool)`

 SetProductTypeNil sets the value for ProductType to be an explicit nil

### UnsetProductType
`func (o *PRODUCTSDETAIL) UnsetProductType()`

UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
### GetPrices

`func (o *PRODUCTSDETAIL) GetPrices() []PRODUCTPRICESCHEMA`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *PRODUCTSDETAIL) GetPricesOk() (*[]PRODUCTPRICESCHEMA, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *PRODUCTSDETAIL) SetPrices(v []PRODUCTPRICESCHEMA)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *PRODUCTSDETAIL) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetLink

`func (o *PRODUCTSDETAIL) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *PRODUCTSDETAIL) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *PRODUCTSDETAIL) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *PRODUCTSDETAIL) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *PRODUCTSDETAIL) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *PRODUCTSDETAIL) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetStoreList

`func (o *PRODUCTSDETAIL) GetStoreList() []STORESSCHEMA`

GetStoreList returns the StoreList field if non-nil, zero value otherwise.

### GetStoreListOk

`func (o *PRODUCTSDETAIL) GetStoreListOk() (*[]STORESSCHEMA, bool)`

GetStoreListOk returns a tuple with the StoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreList

`func (o *PRODUCTSDETAIL) SetStoreList(v []STORESSCHEMA)`

SetStoreList sets StoreList field to given value.

### HasStoreList

`func (o *PRODUCTSDETAIL) HasStoreList() bool`

HasStoreList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


