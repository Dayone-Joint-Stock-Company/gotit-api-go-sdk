# PRODUCTSVOUCHERCHECK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **NullableInt64** | Product Id | [optional] 
**ProductNm** | Pointer to **NullableString** | Product Name | [optional] 
**ProductImg** | Pointer to **NullableString** | Link product image | [optional] 
**BrandId** | Pointer to **NullableInt64** | Brand id | [optional] 
**BrandNm** | Pointer to **NullableString** | Brand name | [optional] 
**BrandServiceGuide** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to [**PRODUCTPRICESCHEMA**](PRODUCTPRICESCHEMA.md) |  | [optional] 
**ProductDesc** | Pointer to **NullableString** |  | [optional] 
**Terms** | Pointer to **NullableString** |  | [optional] 
**StoreList** | Pointer to [**[]STOREPRODUCTSCHEMA**](STOREPRODUCTSCHEMA.md) |  | [optional] 
**TotalStore** | Pointer to **NullableInt64** |  | [optional] 
**StorePagination** | Pointer to [**STOREPAGINGSCHEMA**](STOREPAGINGSCHEMA.md) |  | [optional] 

## Methods

### NewPRODUCTSVOUCHERCHECK

`func NewPRODUCTSVOUCHERCHECK() *PRODUCTSVOUCHERCHECK`

NewPRODUCTSVOUCHERCHECK instantiates a new PRODUCTSVOUCHERCHECK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPRODUCTSVOUCHERCHECKWithDefaults

`func NewPRODUCTSVOUCHERCHECKWithDefaults() *PRODUCTSVOUCHERCHECK`

NewPRODUCTSVOUCHERCHECKWithDefaults instantiates a new PRODUCTSVOUCHERCHECK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *PRODUCTSVOUCHERCHECK) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PRODUCTSVOUCHERCHECK) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PRODUCTSVOUCHERCHECK) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PRODUCTSVOUCHERCHECK) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *PRODUCTSVOUCHERCHECK) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *PRODUCTSVOUCHERCHECK) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetProductNm

`func (o *PRODUCTSVOUCHERCHECK) GetProductNm() string`

GetProductNm returns the ProductNm field if non-nil, zero value otherwise.

### GetProductNmOk

`func (o *PRODUCTSVOUCHERCHECK) GetProductNmOk() (*string, bool)`

GetProductNmOk returns a tuple with the ProductNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductNm

`func (o *PRODUCTSVOUCHERCHECK) SetProductNm(v string)`

SetProductNm sets ProductNm field to given value.

### HasProductNm

`func (o *PRODUCTSVOUCHERCHECK) HasProductNm() bool`

HasProductNm returns a boolean if a field has been set.

### SetProductNmNil

`func (o *PRODUCTSVOUCHERCHECK) SetProductNmNil(b bool)`

 SetProductNmNil sets the value for ProductNm to be an explicit nil

### UnsetProductNm
`func (o *PRODUCTSVOUCHERCHECK) UnsetProductNm()`

UnsetProductNm ensures that no value is present for ProductNm, not even an explicit nil
### GetProductImg

`func (o *PRODUCTSVOUCHERCHECK) GetProductImg() string`

GetProductImg returns the ProductImg field if non-nil, zero value otherwise.

### GetProductImgOk

`func (o *PRODUCTSVOUCHERCHECK) GetProductImgOk() (*string, bool)`

GetProductImgOk returns a tuple with the ProductImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImg

`func (o *PRODUCTSVOUCHERCHECK) SetProductImg(v string)`

SetProductImg sets ProductImg field to given value.

### HasProductImg

`func (o *PRODUCTSVOUCHERCHECK) HasProductImg() bool`

HasProductImg returns a boolean if a field has been set.

### SetProductImgNil

`func (o *PRODUCTSVOUCHERCHECK) SetProductImgNil(b bool)`

 SetProductImgNil sets the value for ProductImg to be an explicit nil

### UnsetProductImg
`func (o *PRODUCTSVOUCHERCHECK) UnsetProductImg()`

UnsetProductImg ensures that no value is present for ProductImg, not even an explicit nil
### GetBrandId

`func (o *PRODUCTSVOUCHERCHECK) GetBrandId() int64`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *PRODUCTSVOUCHERCHECK) GetBrandIdOk() (*int64, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *PRODUCTSVOUCHERCHECK) SetBrandId(v int64)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *PRODUCTSVOUCHERCHECK) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### SetBrandIdNil

`func (o *PRODUCTSVOUCHERCHECK) SetBrandIdNil(b bool)`

 SetBrandIdNil sets the value for BrandId to be an explicit nil

### UnsetBrandId
`func (o *PRODUCTSVOUCHERCHECK) UnsetBrandId()`

UnsetBrandId ensures that no value is present for BrandId, not even an explicit nil
### GetBrandNm

`func (o *PRODUCTSVOUCHERCHECK) GetBrandNm() string`

GetBrandNm returns the BrandNm field if non-nil, zero value otherwise.

### GetBrandNmOk

`func (o *PRODUCTSVOUCHERCHECK) GetBrandNmOk() (*string, bool)`

GetBrandNmOk returns a tuple with the BrandNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNm

`func (o *PRODUCTSVOUCHERCHECK) SetBrandNm(v string)`

SetBrandNm sets BrandNm field to given value.

### HasBrandNm

`func (o *PRODUCTSVOUCHERCHECK) HasBrandNm() bool`

HasBrandNm returns a boolean if a field has been set.

### SetBrandNmNil

`func (o *PRODUCTSVOUCHERCHECK) SetBrandNmNil(b bool)`

 SetBrandNmNil sets the value for BrandNm to be an explicit nil

### UnsetBrandNm
`func (o *PRODUCTSVOUCHERCHECK) UnsetBrandNm()`

UnsetBrandNm ensures that no value is present for BrandNm, not even an explicit nil
### GetBrandServiceGuide

`func (o *PRODUCTSVOUCHERCHECK) GetBrandServiceGuide() string`

GetBrandServiceGuide returns the BrandServiceGuide field if non-nil, zero value otherwise.

### GetBrandServiceGuideOk

`func (o *PRODUCTSVOUCHERCHECK) GetBrandServiceGuideOk() (*string, bool)`

GetBrandServiceGuideOk returns a tuple with the BrandServiceGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandServiceGuide

`func (o *PRODUCTSVOUCHERCHECK) SetBrandServiceGuide(v string)`

SetBrandServiceGuide sets BrandServiceGuide field to given value.

### HasBrandServiceGuide

`func (o *PRODUCTSVOUCHERCHECK) HasBrandServiceGuide() bool`

HasBrandServiceGuide returns a boolean if a field has been set.

### SetBrandServiceGuideNil

`func (o *PRODUCTSVOUCHERCHECK) SetBrandServiceGuideNil(b bool)`

 SetBrandServiceGuideNil sets the value for BrandServiceGuide to be an explicit nil

### UnsetBrandServiceGuide
`func (o *PRODUCTSVOUCHERCHECK) UnsetBrandServiceGuide()`

UnsetBrandServiceGuide ensures that no value is present for BrandServiceGuide, not even an explicit nil
### GetPrice

`func (o *PRODUCTSVOUCHERCHECK) GetPrice() PRODUCTPRICESCHEMA`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PRODUCTSVOUCHERCHECK) GetPriceOk() (*PRODUCTPRICESCHEMA, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PRODUCTSVOUCHERCHECK) SetPrice(v PRODUCTPRICESCHEMA)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PRODUCTSVOUCHERCHECK) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProductDesc

`func (o *PRODUCTSVOUCHERCHECK) GetProductDesc() string`

GetProductDesc returns the ProductDesc field if non-nil, zero value otherwise.

### GetProductDescOk

`func (o *PRODUCTSVOUCHERCHECK) GetProductDescOk() (*string, bool)`

GetProductDescOk returns a tuple with the ProductDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDesc

`func (o *PRODUCTSVOUCHERCHECK) SetProductDesc(v string)`

SetProductDesc sets ProductDesc field to given value.

### HasProductDesc

`func (o *PRODUCTSVOUCHERCHECK) HasProductDesc() bool`

HasProductDesc returns a boolean if a field has been set.

### SetProductDescNil

`func (o *PRODUCTSVOUCHERCHECK) SetProductDescNil(b bool)`

 SetProductDescNil sets the value for ProductDesc to be an explicit nil

### UnsetProductDesc
`func (o *PRODUCTSVOUCHERCHECK) UnsetProductDesc()`

UnsetProductDesc ensures that no value is present for ProductDesc, not even an explicit nil
### GetTerms

`func (o *PRODUCTSVOUCHERCHECK) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *PRODUCTSVOUCHERCHECK) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *PRODUCTSVOUCHERCHECK) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *PRODUCTSVOUCHERCHECK) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### SetTermsNil

`func (o *PRODUCTSVOUCHERCHECK) SetTermsNil(b bool)`

 SetTermsNil sets the value for Terms to be an explicit nil

### UnsetTerms
`func (o *PRODUCTSVOUCHERCHECK) UnsetTerms()`

UnsetTerms ensures that no value is present for Terms, not even an explicit nil
### GetStoreList

`func (o *PRODUCTSVOUCHERCHECK) GetStoreList() []STOREPRODUCTSCHEMA`

GetStoreList returns the StoreList field if non-nil, zero value otherwise.

### GetStoreListOk

`func (o *PRODUCTSVOUCHERCHECK) GetStoreListOk() (*[]STOREPRODUCTSCHEMA, bool)`

GetStoreListOk returns a tuple with the StoreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreList

`func (o *PRODUCTSVOUCHERCHECK) SetStoreList(v []STOREPRODUCTSCHEMA)`

SetStoreList sets StoreList field to given value.

### HasStoreList

`func (o *PRODUCTSVOUCHERCHECK) HasStoreList() bool`

HasStoreList returns a boolean if a field has been set.

### GetTotalStore

`func (o *PRODUCTSVOUCHERCHECK) GetTotalStore() int64`

GetTotalStore returns the TotalStore field if non-nil, zero value otherwise.

### GetTotalStoreOk

`func (o *PRODUCTSVOUCHERCHECK) GetTotalStoreOk() (*int64, bool)`

GetTotalStoreOk returns a tuple with the TotalStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStore

`func (o *PRODUCTSVOUCHERCHECK) SetTotalStore(v int64)`

SetTotalStore sets TotalStore field to given value.

### HasTotalStore

`func (o *PRODUCTSVOUCHERCHECK) HasTotalStore() bool`

HasTotalStore returns a boolean if a field has been set.

### SetTotalStoreNil

`func (o *PRODUCTSVOUCHERCHECK) SetTotalStoreNil(b bool)`

 SetTotalStoreNil sets the value for TotalStore to be an explicit nil

### UnsetTotalStore
`func (o *PRODUCTSVOUCHERCHECK) UnsetTotalStore()`

UnsetTotalStore ensures that no value is present for TotalStore, not even an explicit nil
### GetStorePagination

`func (o *PRODUCTSVOUCHERCHECK) GetStorePagination() STOREPAGINGSCHEMA`

GetStorePagination returns the StorePagination field if non-nil, zero value otherwise.

### GetStorePaginationOk

`func (o *PRODUCTSVOUCHERCHECK) GetStorePaginationOk() (*STOREPAGINGSCHEMA, bool)`

GetStorePaginationOk returns a tuple with the StorePagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePagination

`func (o *PRODUCTSVOUCHERCHECK) SetStorePagination(v STOREPAGINGSCHEMA)`

SetStorePagination sets StorePagination field to given value.

### HasStorePagination

`func (o *PRODUCTSVOUCHERCHECK) HasStorePagination() bool`

HasStorePagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


