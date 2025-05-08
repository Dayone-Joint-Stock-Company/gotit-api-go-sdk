# BRANDSDETAIL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandId** | Pointer to **NullableInt64** | Brand id | [optional] 
**BrandNm** | Pointer to **NullableString** | Brand name | [optional] 
**BrandLogo** | Pointer to **NullableString** | Link to brand logo image | [optional] 
**Slug** | Pointer to **NullableString** | Brand name used for URL | [optional] 
**ShortDesc** | Pointer to **interface{}** | Brand short description | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CategoryID** | Pointer to **[]interface{}** | Array containing a list of Categories codes to which the Brand belongs | [optional] 
**Order** | Pointer to **NullableInt64** | The serial number displays the brand | [optional] 
**ServiceGuide** | Pointer to **interface{}** | Describe the brand&#39;s terms of reference (T&amp;C). HTML format | [optional] 
**Stores** | Pointer to [**[]STORESSCHEMA**](STORESSCHEMA.md) |  | [optional] 
**UsageMethods** | Pointer to [**[]USAGEMETHODSCHEMA**](USAGEMETHODSCHEMA.md) | Information on the brand&#39;s applicable channels | [optional] 

## Methods

### NewBRANDSDETAIL

`func NewBRANDSDETAIL() *BRANDSDETAIL`

NewBRANDSDETAIL instantiates a new BRANDSDETAIL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBRANDSDETAILWithDefaults

`func NewBRANDSDETAILWithDefaults() *BRANDSDETAIL`

NewBRANDSDETAILWithDefaults instantiates a new BRANDSDETAIL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandId

`func (o *BRANDSDETAIL) GetBrandId() int64`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *BRANDSDETAIL) GetBrandIdOk() (*int64, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *BRANDSDETAIL) SetBrandId(v int64)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *BRANDSDETAIL) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### SetBrandIdNil

`func (o *BRANDSDETAIL) SetBrandIdNil(b bool)`

 SetBrandIdNil sets the value for BrandId to be an explicit nil

### UnsetBrandId
`func (o *BRANDSDETAIL) UnsetBrandId()`

UnsetBrandId ensures that no value is present for BrandId, not even an explicit nil
### GetBrandNm

`func (o *BRANDSDETAIL) GetBrandNm() string`

GetBrandNm returns the BrandNm field if non-nil, zero value otherwise.

### GetBrandNmOk

`func (o *BRANDSDETAIL) GetBrandNmOk() (*string, bool)`

GetBrandNmOk returns a tuple with the BrandNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNm

`func (o *BRANDSDETAIL) SetBrandNm(v string)`

SetBrandNm sets BrandNm field to given value.

### HasBrandNm

`func (o *BRANDSDETAIL) HasBrandNm() bool`

HasBrandNm returns a boolean if a field has been set.

### SetBrandNmNil

`func (o *BRANDSDETAIL) SetBrandNmNil(b bool)`

 SetBrandNmNil sets the value for BrandNm to be an explicit nil

### UnsetBrandNm
`func (o *BRANDSDETAIL) UnsetBrandNm()`

UnsetBrandNm ensures that no value is present for BrandNm, not even an explicit nil
### GetBrandLogo

`func (o *BRANDSDETAIL) GetBrandLogo() string`

GetBrandLogo returns the BrandLogo field if non-nil, zero value otherwise.

### GetBrandLogoOk

`func (o *BRANDSDETAIL) GetBrandLogoOk() (*string, bool)`

GetBrandLogoOk returns a tuple with the BrandLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandLogo

`func (o *BRANDSDETAIL) SetBrandLogo(v string)`

SetBrandLogo sets BrandLogo field to given value.

### HasBrandLogo

`func (o *BRANDSDETAIL) HasBrandLogo() bool`

HasBrandLogo returns a boolean if a field has been set.

### SetBrandLogoNil

`func (o *BRANDSDETAIL) SetBrandLogoNil(b bool)`

 SetBrandLogoNil sets the value for BrandLogo to be an explicit nil

### UnsetBrandLogo
`func (o *BRANDSDETAIL) UnsetBrandLogo()`

UnsetBrandLogo ensures that no value is present for BrandLogo, not even an explicit nil
### GetSlug

`func (o *BRANDSDETAIL) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BRANDSDETAIL) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BRANDSDETAIL) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BRANDSDETAIL) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *BRANDSDETAIL) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *BRANDSDETAIL) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetShortDesc

`func (o *BRANDSDETAIL) GetShortDesc() interface{}`

GetShortDesc returns the ShortDesc field if non-nil, zero value otherwise.

### GetShortDescOk

`func (o *BRANDSDETAIL) GetShortDescOk() (*interface{}, bool)`

GetShortDescOk returns a tuple with the ShortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDesc

`func (o *BRANDSDETAIL) SetShortDesc(v interface{})`

SetShortDesc sets ShortDesc field to given value.

### HasShortDesc

`func (o *BRANDSDETAIL) HasShortDesc() bool`

HasShortDesc returns a boolean if a field has been set.

### SetShortDescNil

`func (o *BRANDSDETAIL) SetShortDescNil(b bool)`

 SetShortDescNil sets the value for ShortDesc to be an explicit nil

### UnsetShortDesc
`func (o *BRANDSDETAIL) UnsetShortDesc()`

UnsetShortDesc ensures that no value is present for ShortDesc, not even an explicit nil
### GetDescription

`func (o *BRANDSDETAIL) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BRANDSDETAIL) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BRANDSDETAIL) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BRANDSDETAIL) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BRANDSDETAIL) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BRANDSDETAIL) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategoryID

`func (o *BRANDSDETAIL) GetCategoryID() []interface{}`

GetCategoryID returns the CategoryID field if non-nil, zero value otherwise.

### GetCategoryIDOk

`func (o *BRANDSDETAIL) GetCategoryIDOk() (*[]interface{}, bool)`

GetCategoryIDOk returns a tuple with the CategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryID

`func (o *BRANDSDETAIL) SetCategoryID(v []interface{})`

SetCategoryID sets CategoryID field to given value.

### HasCategoryID

`func (o *BRANDSDETAIL) HasCategoryID() bool`

HasCategoryID returns a boolean if a field has been set.

### SetCategoryIDNil

`func (o *BRANDSDETAIL) SetCategoryIDNil(b bool)`

 SetCategoryIDNil sets the value for CategoryID to be an explicit nil

### UnsetCategoryID
`func (o *BRANDSDETAIL) UnsetCategoryID()`

UnsetCategoryID ensures that no value is present for CategoryID, not even an explicit nil
### GetOrder

`func (o *BRANDSDETAIL) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *BRANDSDETAIL) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *BRANDSDETAIL) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *BRANDSDETAIL) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *BRANDSDETAIL) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *BRANDSDETAIL) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetServiceGuide

`func (o *BRANDSDETAIL) GetServiceGuide() interface{}`

GetServiceGuide returns the ServiceGuide field if non-nil, zero value otherwise.

### GetServiceGuideOk

`func (o *BRANDSDETAIL) GetServiceGuideOk() (*interface{}, bool)`

GetServiceGuideOk returns a tuple with the ServiceGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGuide

`func (o *BRANDSDETAIL) SetServiceGuide(v interface{})`

SetServiceGuide sets ServiceGuide field to given value.

### HasServiceGuide

`func (o *BRANDSDETAIL) HasServiceGuide() bool`

HasServiceGuide returns a boolean if a field has been set.

### SetServiceGuideNil

`func (o *BRANDSDETAIL) SetServiceGuideNil(b bool)`

 SetServiceGuideNil sets the value for ServiceGuide to be an explicit nil

### UnsetServiceGuide
`func (o *BRANDSDETAIL) UnsetServiceGuide()`

UnsetServiceGuide ensures that no value is present for ServiceGuide, not even an explicit nil
### GetStores

`func (o *BRANDSDETAIL) GetStores() []STORESSCHEMA`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *BRANDSDETAIL) GetStoresOk() (*[]STORESSCHEMA, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *BRANDSDETAIL) SetStores(v []STORESSCHEMA)`

SetStores sets Stores field to given value.

### HasStores

`func (o *BRANDSDETAIL) HasStores() bool`

HasStores returns a boolean if a field has been set.

### GetUsageMethods

`func (o *BRANDSDETAIL) GetUsageMethods() []USAGEMETHODSCHEMA`

GetUsageMethods returns the UsageMethods field if non-nil, zero value otherwise.

### GetUsageMethodsOk

`func (o *BRANDSDETAIL) GetUsageMethodsOk() (*[]USAGEMETHODSCHEMA, bool)`

GetUsageMethodsOk returns a tuple with the UsageMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMethods

`func (o *BRANDSDETAIL) SetUsageMethods(v []USAGEMETHODSCHEMA)`

SetUsageMethods sets UsageMethods field to given value.

### HasUsageMethods

`func (o *BRANDSDETAIL) HasUsageMethods() bool`

HasUsageMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


