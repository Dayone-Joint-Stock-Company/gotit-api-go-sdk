# PRODUCTSALLDETAIL

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
**StoreList** | Pointer to [**[]STOREPRODUCTSCHEMA**](STOREPRODUCTSCHEMA.md) |  | [optional] 
**TotalStore** | Pointer to **NullableInt64** | Total store | [optional] 
**StorePagination** | Pointer to [**PAGINGSCHEMA**](PAGINGSCHEMA.md) |  | [optional] 

## Methods

### NewPRODUCTSALLDETAIL

`func NewPRODUCTSALLDETAIL() *PRODUCTSALLDETAIL`

NewPRODUCTSALLDETAIL instantiates a new PRODUCTSALLDETAIL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPRODUCTSALLDETAILWithDefaults

`func NewPRODUCTSALLDETAILWithDefaults() *PRODUCTSALLDETAIL`

NewPRODUCTSALLDETAILWithDefaults instantiates a new PRODUCTSALLDETAIL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *PRODUCTSALLDETAIL) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PRODUCTSALLDETAIL) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PRODUCTSALLDETAIL) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PRODUCTSALLDETAIL) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *PRODUCTSALLDETAIL) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *PRODUCTSALLDETAIL) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetProductNm

`func (o *PRODUCTSALLDETAIL) GetProductNm() string`

GetProductNm returns the ProductNm field if non-nil, zero value otherwise.

### GetProductNmOk

`func (o *PRODUCTSALLDETAIL) GetProductNmOk() (*string, bool)`

GetProductNmOk returns a tuple with the ProductNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductNm

`func (o *PRODUCTSALLDETAIL) SetProductNm(v string)`

SetProductNm sets ProductNm field to given value.

### HasProductNm

`func (o *PRODUCTSALLDETAIL) HasProductNm() bool`

HasProductNm returns a boolean if a field has been set.

### SetProductNmNil

`func (o *PRODUCTSALLDETAIL) SetProductNmNil(b bool)`

 SetProductNmNil sets the value for ProductNm to be an explicit nil

### UnsetProductNm
`func (o *PRODUCTSALLDETAIL) UnsetProductNm()`

UnsetProductNm ensures that no value is present for ProductNm, not even an explicit nil
### GetProductImg

`func (o *PRODUCTSALLDETAIL) GetProductImg() string`

GetProductImg returns the ProductImg field if non-nil, zero value otherwise.

### GetProductImgOk

`func (o *PRODUCTSALLDETAIL) GetProductImgOk() (*string, bool)`

GetProductImgOk returns a tuple with the ProductImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImg

`func (o *PRODUCTSALLDETAIL) SetProductImg(v string)`

SetProductImg sets ProductImg field to given value.

### HasProductImg

`func (o *PRODUCTSALLDETAIL) HasProductImg() bool`

HasProductImg returns a boolean if a field has been set.

### SetProductImgNil

`func (o *PRODUCTSALLDETAIL) SetProductImgNil(b bool)`

 SetProductImgNil sets the value for ProductImg to be an explicit nil

### UnsetProductImg
`func (o *PRODUCTSALLDETAIL) UnsetProductImg()`

UnsetProductImg ensures that no value is present for ProductImg, not even an explicit nil
### GetProductDesc

`func (o *PRODUCTSALLDETAIL) GetProductDesc() string`

GetProductDesc returns the ProductDesc field if non-nil, zero value otherwise.

### GetProductDescOk

`func (o *PRODUCTSALLDETAIL) GetProductDescOk() (*string, bool)`

GetProductDescOk returns a tuple with the ProductDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDesc

`func (o *PRODUCTSALLDETAIL) SetProductDesc(v string)`

SetProductDesc sets ProductDesc field to given value.

### HasProductDesc

`func (o *PRODUCTSALLDETAIL) HasProductDesc() bool`

HasProductDesc returns a boolean if a field has been set.

### SetProductDescNil

`func (o *PRODUCTSALLDETAIL) SetProductDescNil(b bool)`

 SetProductDescNil sets the value for ProductDesc to be an explicit nil

### UnsetProductDesc
`func (o *PRODUCTSALLDETAIL) UnsetProductDesc()`

UnsetProductDesc ensures that no value is present for ProductDesc, not even an explicit nil
### GetProductShortDesc

`func (o *PRODUCTSALLDETAIL) GetProductShortDesc() string`

GetProductShortDesc returns the ProductShortDesc field if non-nil, zero value otherwise.

### GetProductShortDescOk

`func (o *PRODUCTSALLDETAIL) GetProductShortDescOk() (*string, bool)`

GetProductShortDescOk returns a tuple with the ProductShortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductShortDesc

`func (o *PRODUCTSALLDETAIL) SetProductShortDesc(v string)`

SetProductShortDesc sets ProductShortDesc field to given value.

### HasProductShortDesc

`func (o *PRODUCTSALLDETAIL) HasProductShortDesc() bool`

HasProductShortDesc returns a boolean if a field has been set.

### SetProductShortDescNil

`func (o *PRODUCTSALLDETAIL) SetProductShortDescNil(b bool)`

 SetProductShortDescNil sets the value for ProductShortDesc to be an explicit nil

### UnsetProductShortDesc
`func (o *PRODUCTSALLDETAIL) UnsetProductShortDesc()`

UnsetProductShortDesc ensures that no value is present for ProductShortDesc, not even an explicit nil
### GetTerms

`func (o *PRODUCTSALLDETAIL) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *PRODUCTSALLDETAIL) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *PRODUCTSALLDETAIL) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *PRODUCTSALLDETAIL) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### SetTermsNil

`func (o *PRODUCTSALLDETAIL) SetTermsNil(b bool)`

 SetTermsNil sets the value for Terms to be an explicit nil

### UnsetTerms
`func (o *PRODUCTSALLDETAIL) UnsetTerms()`

UnsetTerms ensures that no value is present for Terms, not even an explicit nil
### GetBrandId

`func (o *PRODUCTSALLDETAIL) GetBrandId() int64`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *PRODUCTSALLDETAIL) GetBrandIdOk() (*int64, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *PRODUCTSALLDETAIL) SetBrandId(v int64)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *PRODUCTSALLDETAIL) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### SetBrandIdNil

`func (o *PRODUCTSALLDETAIL) SetBrandIdNil(b bool)`

 SetBrandIdNil sets the value for BrandId to be an explicit nil

### UnsetBrandId
`func (o *PRODUCTSALLDETAIL) UnsetBrandId()`

UnsetBrandId ensures that no value is present for BrandId, not even an explicit nil
### GetBrandNm

`func (o *PRODUCTSALLDETAIL) GetBrandNm() string`

GetBrandNm returns the BrandNm field if non-nil, zero value otherwise.

### GetBrandNmOk

`func (o *PRODUCTSALLDETAIL) GetBrandNmOk() (*string, bool)`

GetBrandNmOk returns a tuple with the BrandNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNm

`func (o *PRODUCTSALLDETAIL) SetBrandNm(v string)`

SetBrandNm sets BrandNm field to given value.

### HasBrandNm

`func (o *PRODUCTSALLDETAIL) HasBrandNm() bool`

HasBrandNm returns a boolean if a field has been set.

### SetBrandNmNil

`func (o *PRODUCTSALLDETAIL) SetBrandNmNil(b bool)`

 SetBrandNmNil sets the value for BrandNm to be an explicit nil

### UnsetBrandNm
`func (o *PRODUCTSALLDETAIL) UnsetBrandNm()`

UnsetBrandNm ensures that no value is present for BrandNm, not even an explicit nil
### GetBrandLogo

`func (o *PRODUCTSALLDETAIL) GetBrandLogo() string`

GetBrandLogo returns the BrandLogo field if non-nil, zero value otherwise.

### GetBrandLogoOk

`func (o *PRODUCTSALLDETAIL) GetBrandLogoOk() (*string, bool)`

GetBrandLogoOk returns a tuple with the BrandLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandLogo

`func (o *PRODUCTSALLDETAIL) SetBrandLogo(v string)`

SetBrandLogo sets BrandLogo field to given value.

### HasBrandLogo

`func (o *PRODUCTSALLDETAIL) HasBrandLogo() bool`

HasBrandLogo returns a boolean if a field has been set.

### SetBrandLogoNil

`func (o *PRODUCTSALLDETAIL) SetBrandLogoNil(b bool)`

 SetBrandLogoNil sets the value for BrandLogo to be an explicit nil

### UnsetBrandLogo
`func (o *PRODUCTSALLDETAIL) UnsetBrandLogo()`

UnsetBrandLogo ensures that no value is present for BrandLogo, not even an explicit nil
### GetBrandServiceGuide

`func (o *PRODUCTSALLDETAIL) GetBrandServiceGuide() string`

GetBrandServiceGuide returns the BrandServiceGuide field if non-nil, zero value otherwise.

### GetBrandServiceGuideOk

`func (o *PRODUCTSALLDETAIL) GetBrandServiceGuideOk() (*string, bool)`

GetBrandServiceGuideOk returns a tuple with the BrandServiceGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandServiceGuide

`func (o *PRODUCTSALLDETAIL) SetBrandServiceGuide(v string)`

SetBrandServiceGuide sets BrandServiceGuide field to given value.

### HasBrandServiceGuide

`func (o *PRODUCTSALLDETAIL) HasBrandServiceGuide() bool`

HasBrandServiceGuide returns a boolean if a field has been set.

### SetBrandServiceGuideNil

`func (o *PRODUCTSALLDETAIL) SetBrandServiceGuideNil(b bool)`

 SetBrandServiceGuideNil sets the value for BrandServiceGuide to be an explicit nil

### UnsetBrandServiceGuide
`func (o *PRODUCTSALLDETAIL) UnsetBrandServiceGuide()`

UnsetBrandServiceGuide ensures that no value is present for BrandServiceGuide, not even an explicit nil
### GetCategoryId

`func (o *PRODUCTSALLDETAIL) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *PRODUCTSALLDETAIL) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *PRODUCTSALLDETAIL) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *PRODUCTSALLDETAIL) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *PRODUCTSALLDETAIL) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *PRODUCTSALLDETAIL) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetCategoryNm

`func (o *PRODUCTSALLDETAIL) GetCategoryNm() string`

GetCategoryNm returns the CategoryNm field if non-nil, zero value otherwise.

### GetCategoryNmOk

`func (o *PRODUCTSALLDETAIL) GetCategoryNmOk() (*string, bool)`

GetCategoryNmOk returns a tuple with the CategoryNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryNm

`func (o *PRODUCTSALLDETAIL) SetCategoryNm(v string)`

SetCategoryNm sets CategoryNm field to given value.

### HasCategoryNm

`func (o *PRODUCTSALLDETAIL) HasCategoryNm() bool`

HasCategoryNm returns a boolean if a field has been set.

### SetCategoryNmNil

`func (o *PRODUCTSALLDETAIL) SetCategoryNmNil(b bool)`

 SetCategoryNmNil sets the value for CategoryNm to be an explicit nil

### UnsetCategoryNm
`func (o *PRODUCTSALLDETAIL) UnsetCategoryNm()`

UnsetCategoryNm ensures that no value is present for CategoryNm, not even an explicit nil
### GetCategoryImg

`func (o *PRODUCTSALLDETAIL) GetCategoryImg() string`

GetCategoryImg returns the CategoryImg field if non-nil, zero value otherwise.

### GetCategoryImgOk

`func (o *PRODUCTSALLDETAIL) GetCategoryImgOk() (*string, bool)`

GetCategoryImgOk returns a tuple with the CategoryImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryImg

`func (o *PRODUCTSALLDETAIL) SetCategoryImg(v string)`

SetCategoryImg sets CategoryImg field to given value.

### HasCategoryImg

`func (o *PRODUCTSALLDETAIL) HasCategoryImg() bool`

HasCategoryImg returns a boolean if a field has been set.

### SetCategoryImgNil

`func (o *PRODUCTSALLDETAIL) SetCategoryImgNil(b bool)`

 SetCategoryImgNil sets the value for CategoryImg to be an explicit nil

### UnsetCategoryImg
`func (o *PRODUCTSALLDETAIL) UnsetCategoryImg()`

UnsetCategoryImg ensures that no value is present for CategoryImg, not even an explicit nil
### GetProductType

`func (o *PRODUCTSALLDETAIL) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PRODUCTSALLDETAIL) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PRODUCTSALLDETAIL) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *PRODUCTSALLDETAIL) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### SetProductTypeNil

`func (o *PRODUCTSALLDETAIL) SetProductTypeNil(b bool)`

 SetProductTypeNil sets the value for ProductType to be an explicit nil

### UnsetProductType
`func (o *PRODUCTSALLDETAIL) UnsetProductType()`

UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
### GetPrices

`func (o *PRODUCTSALLDETAIL) GetPrices() []PRODUCTPRICESCHEMA`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *PRODUCTSALLDETAIL) GetPricesOk() (*[]PRODUCTPRICESCHEMA, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *PRODUCTSALLDETAIL) SetPrices(v []PRODUCTPRICESCHEMA)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *PRODUCTSALLDETAIL) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetLink

`func (o *PRODUCTSALLDETAIL) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *PRODUCTSALLDETAIL) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *PRODUCTSALLDETAIL) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *PRODUCTSALLDETAIL) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *PRODUCTSALLDETAIL) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *PRODUCTSALLDETAIL) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetStoreList

`func (o *PRODUCTSALLDETAIL) GetStoreList() []STOREPRODUCTSCHEMA`

GetStoreList returns the StoreList field if non-nil, zero value otherwise.

### GetStoreListOk

`func (o *PRODUCTSALLDETAIL) GetStoreListOk() (*[]STOREPRODUCTSCHEMA, bool)`

GetStoreListOk returns a tuple with the StoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreList

`func (o *PRODUCTSALLDETAIL) SetStoreList(v []STOREPRODUCTSCHEMA)`

SetStoreList sets StoreList field to given value.

### HasStoreList

`func (o *PRODUCTSALLDETAIL) HasStoreList() bool`

HasStoreList returns a boolean if a field has been set.

### GetTotalStore

`func (o *PRODUCTSALLDETAIL) GetTotalStore() int64`

GetTotalStore returns the TotalStore field if non-nil, zero value otherwise.

### GetTotalStoreOk

`func (o *PRODUCTSALLDETAIL) GetTotalStoreOk() (*int64, bool)`

GetTotalStoreOk returns a tuple with the TotalStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStore

`func (o *PRODUCTSALLDETAIL) SetTotalStore(v int64)`

SetTotalStore sets TotalStore field to given value.

### HasTotalStore

`func (o *PRODUCTSALLDETAIL) HasTotalStore() bool`

HasTotalStore returns a boolean if a field has been set.

### SetTotalStoreNil

`func (o *PRODUCTSALLDETAIL) SetTotalStoreNil(b bool)`

 SetTotalStoreNil sets the value for TotalStore to be an explicit nil

### UnsetTotalStore
`func (o *PRODUCTSALLDETAIL) UnsetTotalStore()`

UnsetTotalStore ensures that no value is present for TotalStore, not even an explicit nil
### GetStorePagination

`func (o *PRODUCTSALLDETAIL) GetStorePagination() PAGINGSCHEMA`

GetStorePagination returns the StorePagination field if non-nil, zero value otherwise.

### GetStorePaginationOk

`func (o *PRODUCTSALLDETAIL) GetStorePaginationOk() (*PAGINGSCHEMA, bool)`

GetStorePaginationOk returns a tuple with the StorePagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePagination

`func (o *PRODUCTSALLDETAIL) SetStorePagination(v PAGINGSCHEMA)`

SetStorePagination sets StorePagination field to given value.

### HasStorePagination

`func (o *PRODUCTSALLDETAIL) HasStorePagination() bool`

HasStorePagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


