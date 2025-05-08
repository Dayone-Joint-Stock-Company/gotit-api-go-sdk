# PRODUCTDETAIL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **int64** | Product Id | [optional] 
**ProductNm** | Pointer to **string** | Product Name | [optional] 
**ProductImg** | Pointer to **string** | Link product image | [optional] 
**ProductSubImg** | Pointer to **[]interface{}** | Array of link image | [optional] 
**BrandId** | Pointer to **int64** | Brand id | [optional] 
**BrandNm** | Pointer to **string** | Brand name | [optional] 
**ProductType** | Pointer to **string** | c (cash) hoặc i (item) | [optional] 
**BrandNameSlug** | Pointer to **string** | Slug of brand | [optional] 
**BrandPhone** | Pointer to **string** | Phone of brand | [optional] 
**BrandAddress** | Pointer to **string** | Address of brand | [optional] 
**BrandDesc** | Pointer to **string** | Description of brand | [optional] 
**BrandServiceGuide** | Pointer to **string** | T&amp;C of brand | [optional] 
**ServiceGuide** | Pointer to **string** | T&amp;C of product | [optional] 
**BrandLogo** | Pointer to **string** | Link to brand logo image | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Prices** | Pointer to [**[]PRODUCTPRICESCHEMA**](PRODUCTPRICESCHEMA.md) |  | [optional] 
**NameSlug** | Pointer to **string** | Slug of product | [optional] 
**ProductDesc** | Pointer to **string** | Product Description | [optional] 
**ProductShortDesc** | Pointer to **string** | Product Short Description | [optional] 
**Terms** | Pointer to **string** | Terms of use | [optional] 
**CategoryId** | Pointer to **int64** | Category Id | [optional] 
**CategoryNm** | Pointer to **string** | Category Name | [optional] 
**CategoryImg** | Pointer to **string** | Category Image | [optional] 
**BrandRedeem** | Pointer to [**[]BRANDREEDEMSCHEMA**](BRANDREEDEMSCHEMA.md) |  | [optional] 
**StoreList** | Pointer to [**[]STORESSCHEMA**](STORESSCHEMA.md) |  | [optional] 

## Methods

### NewPRODUCTDETAIL

`func NewPRODUCTDETAIL() *PRODUCTDETAIL`

NewPRODUCTDETAIL instantiates a new PRODUCTDETAIL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPRODUCTDETAILWithDefaults

`func NewPRODUCTDETAILWithDefaults() *PRODUCTDETAIL`

NewPRODUCTDETAILWithDefaults instantiates a new PRODUCTDETAIL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *PRODUCTDETAIL) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PRODUCTDETAIL) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PRODUCTDETAIL) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PRODUCTDETAIL) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductNm

`func (o *PRODUCTDETAIL) GetProductNm() string`

GetProductNm returns the ProductNm field if non-nil, zero value otherwise.

### GetProductNmOk

`func (o *PRODUCTDETAIL) GetProductNmOk() (*string, bool)`

GetProductNmOk returns a tuple with the ProductNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductNm

`func (o *PRODUCTDETAIL) SetProductNm(v string)`

SetProductNm sets ProductNm field to given value.

### HasProductNm

`func (o *PRODUCTDETAIL) HasProductNm() bool`

HasProductNm returns a boolean if a field has been set.

### GetProductImg

`func (o *PRODUCTDETAIL) GetProductImg() string`

GetProductImg returns the ProductImg field if non-nil, zero value otherwise.

### GetProductImgOk

`func (o *PRODUCTDETAIL) GetProductImgOk() (*string, bool)`

GetProductImgOk returns a tuple with the ProductImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImg

`func (o *PRODUCTDETAIL) SetProductImg(v string)`

SetProductImg sets ProductImg field to given value.

### HasProductImg

`func (o *PRODUCTDETAIL) HasProductImg() bool`

HasProductImg returns a boolean if a field has been set.

### GetProductSubImg

`func (o *PRODUCTDETAIL) GetProductSubImg() []interface{}`

GetProductSubImg returns the ProductSubImg field if non-nil, zero value otherwise.

### GetProductSubImgOk

`func (o *PRODUCTDETAIL) GetProductSubImgOk() (*[]interface{}, bool)`

GetProductSubImgOk returns a tuple with the ProductSubImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubImg

`func (o *PRODUCTDETAIL) SetProductSubImg(v []interface{})`

SetProductSubImg sets ProductSubImg field to given value.

### HasProductSubImg

`func (o *PRODUCTDETAIL) HasProductSubImg() bool`

HasProductSubImg returns a boolean if a field has been set.

### GetBrandId

`func (o *PRODUCTDETAIL) GetBrandId() int64`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *PRODUCTDETAIL) GetBrandIdOk() (*int64, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *PRODUCTDETAIL) SetBrandId(v int64)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *PRODUCTDETAIL) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetBrandNm

`func (o *PRODUCTDETAIL) GetBrandNm() string`

GetBrandNm returns the BrandNm field if non-nil, zero value otherwise.

### GetBrandNmOk

`func (o *PRODUCTDETAIL) GetBrandNmOk() (*string, bool)`

GetBrandNmOk returns a tuple with the BrandNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNm

`func (o *PRODUCTDETAIL) SetBrandNm(v string)`

SetBrandNm sets BrandNm field to given value.

### HasBrandNm

`func (o *PRODUCTDETAIL) HasBrandNm() bool`

HasBrandNm returns a boolean if a field has been set.

### GetProductType

`func (o *PRODUCTDETAIL) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PRODUCTDETAIL) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PRODUCTDETAIL) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *PRODUCTDETAIL) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetBrandNameSlug

`func (o *PRODUCTDETAIL) GetBrandNameSlug() string`

GetBrandNameSlug returns the BrandNameSlug field if non-nil, zero value otherwise.

### GetBrandNameSlugOk

`func (o *PRODUCTDETAIL) GetBrandNameSlugOk() (*string, bool)`

GetBrandNameSlugOk returns a tuple with the BrandNameSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNameSlug

`func (o *PRODUCTDETAIL) SetBrandNameSlug(v string)`

SetBrandNameSlug sets BrandNameSlug field to given value.

### HasBrandNameSlug

`func (o *PRODUCTDETAIL) HasBrandNameSlug() bool`

HasBrandNameSlug returns a boolean if a field has been set.

### GetBrandPhone

`func (o *PRODUCTDETAIL) GetBrandPhone() string`

GetBrandPhone returns the BrandPhone field if non-nil, zero value otherwise.

### GetBrandPhoneOk

`func (o *PRODUCTDETAIL) GetBrandPhoneOk() (*string, bool)`

GetBrandPhoneOk returns a tuple with the BrandPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandPhone

`func (o *PRODUCTDETAIL) SetBrandPhone(v string)`

SetBrandPhone sets BrandPhone field to given value.

### HasBrandPhone

`func (o *PRODUCTDETAIL) HasBrandPhone() bool`

HasBrandPhone returns a boolean if a field has been set.

### GetBrandAddress

`func (o *PRODUCTDETAIL) GetBrandAddress() string`

GetBrandAddress returns the BrandAddress field if non-nil, zero value otherwise.

### GetBrandAddressOk

`func (o *PRODUCTDETAIL) GetBrandAddressOk() (*string, bool)`

GetBrandAddressOk returns a tuple with the BrandAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandAddress

`func (o *PRODUCTDETAIL) SetBrandAddress(v string)`

SetBrandAddress sets BrandAddress field to given value.

### HasBrandAddress

`func (o *PRODUCTDETAIL) HasBrandAddress() bool`

HasBrandAddress returns a boolean if a field has been set.

### GetBrandDesc

`func (o *PRODUCTDETAIL) GetBrandDesc() string`

GetBrandDesc returns the BrandDesc field if non-nil, zero value otherwise.

### GetBrandDescOk

`func (o *PRODUCTDETAIL) GetBrandDescOk() (*string, bool)`

GetBrandDescOk returns a tuple with the BrandDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandDesc

`func (o *PRODUCTDETAIL) SetBrandDesc(v string)`

SetBrandDesc sets BrandDesc field to given value.

### HasBrandDesc

`func (o *PRODUCTDETAIL) HasBrandDesc() bool`

HasBrandDesc returns a boolean if a field has been set.

### GetBrandServiceGuide

`func (o *PRODUCTDETAIL) GetBrandServiceGuide() string`

GetBrandServiceGuide returns the BrandServiceGuide field if non-nil, zero value otherwise.

### GetBrandServiceGuideOk

`func (o *PRODUCTDETAIL) GetBrandServiceGuideOk() (*string, bool)`

GetBrandServiceGuideOk returns a tuple with the BrandServiceGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandServiceGuide

`func (o *PRODUCTDETAIL) SetBrandServiceGuide(v string)`

SetBrandServiceGuide sets BrandServiceGuide field to given value.

### HasBrandServiceGuide

`func (o *PRODUCTDETAIL) HasBrandServiceGuide() bool`

HasBrandServiceGuide returns a boolean if a field has been set.

### GetServiceGuide

`func (o *PRODUCTDETAIL) GetServiceGuide() string`

GetServiceGuide returns the ServiceGuide field if non-nil, zero value otherwise.

### GetServiceGuideOk

`func (o *PRODUCTDETAIL) GetServiceGuideOk() (*string, bool)`

GetServiceGuideOk returns a tuple with the ServiceGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGuide

`func (o *PRODUCTDETAIL) SetServiceGuide(v string)`

SetServiceGuide sets ServiceGuide field to given value.

### HasServiceGuide

`func (o *PRODUCTDETAIL) HasServiceGuide() bool`

HasServiceGuide returns a boolean if a field has been set.

### GetBrandLogo

`func (o *PRODUCTDETAIL) GetBrandLogo() string`

GetBrandLogo returns the BrandLogo field if non-nil, zero value otherwise.

### GetBrandLogoOk

`func (o *PRODUCTDETAIL) GetBrandLogoOk() (*string, bool)`

GetBrandLogoOk returns a tuple with the BrandLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandLogo

`func (o *PRODUCTDETAIL) SetBrandLogo(v string)`

SetBrandLogo sets BrandLogo field to given value.

### HasBrandLogo

`func (o *PRODUCTDETAIL) HasBrandLogo() bool`

HasBrandLogo returns a boolean if a field has been set.

### GetLink

`func (o *PRODUCTDETAIL) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *PRODUCTDETAIL) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *PRODUCTDETAIL) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *PRODUCTDETAIL) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPrices

`func (o *PRODUCTDETAIL) GetPrices() []PRODUCTPRICESCHEMA`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *PRODUCTDETAIL) GetPricesOk() (*[]PRODUCTPRICESCHEMA, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *PRODUCTDETAIL) SetPrices(v []PRODUCTPRICESCHEMA)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *PRODUCTDETAIL) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetNameSlug

`func (o *PRODUCTDETAIL) GetNameSlug() string`

GetNameSlug returns the NameSlug field if non-nil, zero value otherwise.

### GetNameSlugOk

`func (o *PRODUCTDETAIL) GetNameSlugOk() (*string, bool)`

GetNameSlugOk returns a tuple with the NameSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSlug

`func (o *PRODUCTDETAIL) SetNameSlug(v string)`

SetNameSlug sets NameSlug field to given value.

### HasNameSlug

`func (o *PRODUCTDETAIL) HasNameSlug() bool`

HasNameSlug returns a boolean if a field has been set.

### GetProductDesc

`func (o *PRODUCTDETAIL) GetProductDesc() string`

GetProductDesc returns the ProductDesc field if non-nil, zero value otherwise.

### GetProductDescOk

`func (o *PRODUCTDETAIL) GetProductDescOk() (*string, bool)`

GetProductDescOk returns a tuple with the ProductDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDesc

`func (o *PRODUCTDETAIL) SetProductDesc(v string)`

SetProductDesc sets ProductDesc field to given value.

### HasProductDesc

`func (o *PRODUCTDETAIL) HasProductDesc() bool`

HasProductDesc returns a boolean if a field has been set.

### GetProductShortDesc

`func (o *PRODUCTDETAIL) GetProductShortDesc() string`

GetProductShortDesc returns the ProductShortDesc field if non-nil, zero value otherwise.

### GetProductShortDescOk

`func (o *PRODUCTDETAIL) GetProductShortDescOk() (*string, bool)`

GetProductShortDescOk returns a tuple with the ProductShortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductShortDesc

`func (o *PRODUCTDETAIL) SetProductShortDesc(v string)`

SetProductShortDesc sets ProductShortDesc field to given value.

### HasProductShortDesc

`func (o *PRODUCTDETAIL) HasProductShortDesc() bool`

HasProductShortDesc returns a boolean if a field has been set.

### GetTerms

`func (o *PRODUCTDETAIL) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *PRODUCTDETAIL) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *PRODUCTDETAIL) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *PRODUCTDETAIL) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetCategoryId

`func (o *PRODUCTDETAIL) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *PRODUCTDETAIL) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *PRODUCTDETAIL) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *PRODUCTDETAIL) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCategoryNm

`func (o *PRODUCTDETAIL) GetCategoryNm() string`

GetCategoryNm returns the CategoryNm field if non-nil, zero value otherwise.

### GetCategoryNmOk

`func (o *PRODUCTDETAIL) GetCategoryNmOk() (*string, bool)`

GetCategoryNmOk returns a tuple with the CategoryNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryNm

`func (o *PRODUCTDETAIL) SetCategoryNm(v string)`

SetCategoryNm sets CategoryNm field to given value.

### HasCategoryNm

`func (o *PRODUCTDETAIL) HasCategoryNm() bool`

HasCategoryNm returns a boolean if a field has been set.

### GetCategoryImg

`func (o *PRODUCTDETAIL) GetCategoryImg() string`

GetCategoryImg returns the CategoryImg field if non-nil, zero value otherwise.

### GetCategoryImgOk

`func (o *PRODUCTDETAIL) GetCategoryImgOk() (*string, bool)`

GetCategoryImgOk returns a tuple with the CategoryImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryImg

`func (o *PRODUCTDETAIL) SetCategoryImg(v string)`

SetCategoryImg sets CategoryImg field to given value.

### HasCategoryImg

`func (o *PRODUCTDETAIL) HasCategoryImg() bool`

HasCategoryImg returns a boolean if a field has been set.

### GetBrandRedeem

`func (o *PRODUCTDETAIL) GetBrandRedeem() []BRANDREEDEMSCHEMA`

GetBrandRedeem returns the BrandRedeem field if non-nil, zero value otherwise.

### GetBrandRedeemOk

`func (o *PRODUCTDETAIL) GetBrandRedeemOk() (*[]BRANDREEDEMSCHEMA, bool)`

GetBrandRedeemOk returns a tuple with the BrandRedeem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandRedeem

`func (o *PRODUCTDETAIL) SetBrandRedeem(v []BRANDREEDEMSCHEMA)`

SetBrandRedeem sets BrandRedeem field to given value.

### HasBrandRedeem

`func (o *PRODUCTDETAIL) HasBrandRedeem() bool`

HasBrandRedeem returns a boolean if a field has been set.

### GetStoreList

`func (o *PRODUCTDETAIL) GetStoreList() []STORESSCHEMA`

GetStoreList returns the StoreList field if non-nil, zero value otherwise.

### GetStoreListOk

`func (o *PRODUCTDETAIL) GetStoreListOk() (*[]STORESSCHEMA, bool)`

GetStoreListOk returns a tuple with the StoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreList

`func (o *PRODUCTDETAIL) SetStoreList(v []STORESSCHEMA)`

SetStoreList sets StoreList field to given value.

### HasStoreList

`func (o *PRODUCTDETAIL) HasStoreList() bool`

HasStoreList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


