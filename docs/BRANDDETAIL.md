# BRANDDETAIL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandId** | Pointer to **NullableInt64** | Brand id | [optional] 
**BrandNm** | Pointer to **NullableString** | Brand name | [optional] 
**BrandAddr** | Pointer to **NullableString** | Address of brand | [optional] 
**BrandLogo** | Pointer to **NullableString** | Link to brand logo image | [optional] 
**Slug** | Pointer to **string** | Brand name used for URL | [optional] 
**ShortDesc** | Pointer to **NullableString** | Brand short description | [optional] 
**Description** | Pointer to **NullableString** | Brand description | [optional] 
**CategoryID** | Pointer to **[]interface{}** | Array containing a list of Categories codes to which the Brand belongs | [optional] 
**Order** | Pointer to **NullableInt64** | The serial number displays the brand | [optional] 
**ServiceGuide** | Pointer to **NullableString** | Describe the brand&#39;s terms of reference (T&amp;C). HTML format | [optional] 
**Stores** | Pointer to [**[]STORESSCHEMA**](STORESSCHEMA.md) |  | [optional] 
**UsageMethods** | Pointer to [**[]USAGEMETHODSCHEMA**](USAGEMETHODSCHEMA.md) | Information on the brand&#39;s applicable channels | [optional] 

## Methods

### NewBRANDDETAIL

`func NewBRANDDETAIL() *BRANDDETAIL`

NewBRANDDETAIL instantiates a new BRANDDETAIL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBRANDDETAILWithDefaults

`func NewBRANDDETAILWithDefaults() *BRANDDETAIL`

NewBRANDDETAILWithDefaults instantiates a new BRANDDETAIL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandId

`func (o *BRANDDETAIL) GetBrandId() int64`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *BRANDDETAIL) GetBrandIdOk() (*int64, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *BRANDDETAIL) SetBrandId(v int64)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *BRANDDETAIL) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### SetBrandIdNil

`func (o *BRANDDETAIL) SetBrandIdNil(b bool)`

 SetBrandIdNil sets the value for BrandId to be an explicit nil

### UnsetBrandId
`func (o *BRANDDETAIL) UnsetBrandId()`

UnsetBrandId ensures that no value is present for BrandId, not even an explicit nil
### GetBrandNm

`func (o *BRANDDETAIL) GetBrandNm() string`

GetBrandNm returns the BrandNm field if non-nil, zero value otherwise.

### GetBrandNmOk

`func (o *BRANDDETAIL) GetBrandNmOk() (*string, bool)`

GetBrandNmOk returns a tuple with the BrandNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNm

`func (o *BRANDDETAIL) SetBrandNm(v string)`

SetBrandNm sets BrandNm field to given value.

### HasBrandNm

`func (o *BRANDDETAIL) HasBrandNm() bool`

HasBrandNm returns a boolean if a field has been set.

### SetBrandNmNil

`func (o *BRANDDETAIL) SetBrandNmNil(b bool)`

 SetBrandNmNil sets the value for BrandNm to be an explicit nil

### UnsetBrandNm
`func (o *BRANDDETAIL) UnsetBrandNm()`

UnsetBrandNm ensures that no value is present for BrandNm, not even an explicit nil
### GetBrandAddr

`func (o *BRANDDETAIL) GetBrandAddr() string`

GetBrandAddr returns the BrandAddr field if non-nil, zero value otherwise.

### GetBrandAddrOk

`func (o *BRANDDETAIL) GetBrandAddrOk() (*string, bool)`

GetBrandAddrOk returns a tuple with the BrandAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandAddr

`func (o *BRANDDETAIL) SetBrandAddr(v string)`

SetBrandAddr sets BrandAddr field to given value.

### HasBrandAddr

`func (o *BRANDDETAIL) HasBrandAddr() bool`

HasBrandAddr returns a boolean if a field has been set.

### SetBrandAddrNil

`func (o *BRANDDETAIL) SetBrandAddrNil(b bool)`

 SetBrandAddrNil sets the value for BrandAddr to be an explicit nil

### UnsetBrandAddr
`func (o *BRANDDETAIL) UnsetBrandAddr()`

UnsetBrandAddr ensures that no value is present for BrandAddr, not even an explicit nil
### GetBrandLogo

`func (o *BRANDDETAIL) GetBrandLogo() string`

GetBrandLogo returns the BrandLogo field if non-nil, zero value otherwise.

### GetBrandLogoOk

`func (o *BRANDDETAIL) GetBrandLogoOk() (*string, bool)`

GetBrandLogoOk returns a tuple with the BrandLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandLogo

`func (o *BRANDDETAIL) SetBrandLogo(v string)`

SetBrandLogo sets BrandLogo field to given value.

### HasBrandLogo

`func (o *BRANDDETAIL) HasBrandLogo() bool`

HasBrandLogo returns a boolean if a field has been set.

### SetBrandLogoNil

`func (o *BRANDDETAIL) SetBrandLogoNil(b bool)`

 SetBrandLogoNil sets the value for BrandLogo to be an explicit nil

### UnsetBrandLogo
`func (o *BRANDDETAIL) UnsetBrandLogo()`

UnsetBrandLogo ensures that no value is present for BrandLogo, not even an explicit nil
### GetSlug

`func (o *BRANDDETAIL) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BRANDDETAIL) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BRANDDETAIL) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BRANDDETAIL) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetShortDesc

`func (o *BRANDDETAIL) GetShortDesc() string`

GetShortDesc returns the ShortDesc field if non-nil, zero value otherwise.

### GetShortDescOk

`func (o *BRANDDETAIL) GetShortDescOk() (*string, bool)`

GetShortDescOk returns a tuple with the ShortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDesc

`func (o *BRANDDETAIL) SetShortDesc(v string)`

SetShortDesc sets ShortDesc field to given value.

### HasShortDesc

`func (o *BRANDDETAIL) HasShortDesc() bool`

HasShortDesc returns a boolean if a field has been set.

### SetShortDescNil

`func (o *BRANDDETAIL) SetShortDescNil(b bool)`

 SetShortDescNil sets the value for ShortDesc to be an explicit nil

### UnsetShortDesc
`func (o *BRANDDETAIL) UnsetShortDesc()`

UnsetShortDesc ensures that no value is present for ShortDesc, not even an explicit nil
### GetDescription

`func (o *BRANDDETAIL) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BRANDDETAIL) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BRANDDETAIL) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BRANDDETAIL) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BRANDDETAIL) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BRANDDETAIL) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategoryID

`func (o *BRANDDETAIL) GetCategoryID() []interface{}`

GetCategoryID returns the CategoryID field if non-nil, zero value otherwise.

### GetCategoryIDOk

`func (o *BRANDDETAIL) GetCategoryIDOk() (*[]interface{}, bool)`

GetCategoryIDOk returns a tuple with the CategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryID

`func (o *BRANDDETAIL) SetCategoryID(v []interface{})`

SetCategoryID sets CategoryID field to given value.

### HasCategoryID

`func (o *BRANDDETAIL) HasCategoryID() bool`

HasCategoryID returns a boolean if a field has been set.

### SetCategoryIDNil

`func (o *BRANDDETAIL) SetCategoryIDNil(b bool)`

 SetCategoryIDNil sets the value for CategoryID to be an explicit nil

### UnsetCategoryID
`func (o *BRANDDETAIL) UnsetCategoryID()`

UnsetCategoryID ensures that no value is present for CategoryID, not even an explicit nil
### GetOrder

`func (o *BRANDDETAIL) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *BRANDDETAIL) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *BRANDDETAIL) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *BRANDDETAIL) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *BRANDDETAIL) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *BRANDDETAIL) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetServiceGuide

`func (o *BRANDDETAIL) GetServiceGuide() string`

GetServiceGuide returns the ServiceGuide field if non-nil, zero value otherwise.

### GetServiceGuideOk

`func (o *BRANDDETAIL) GetServiceGuideOk() (*string, bool)`

GetServiceGuideOk returns a tuple with the ServiceGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGuide

`func (o *BRANDDETAIL) SetServiceGuide(v string)`

SetServiceGuide sets ServiceGuide field to given value.

### HasServiceGuide

`func (o *BRANDDETAIL) HasServiceGuide() bool`

HasServiceGuide returns a boolean if a field has been set.

### SetServiceGuideNil

`func (o *BRANDDETAIL) SetServiceGuideNil(b bool)`

 SetServiceGuideNil sets the value for ServiceGuide to be an explicit nil

### UnsetServiceGuide
`func (o *BRANDDETAIL) UnsetServiceGuide()`

UnsetServiceGuide ensures that no value is present for ServiceGuide, not even an explicit nil
### GetStores

`func (o *BRANDDETAIL) GetStores() []STORESSCHEMA`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *BRANDDETAIL) GetStoresOk() (*[]STORESSCHEMA, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *BRANDDETAIL) SetStores(v []STORESSCHEMA)`

SetStores sets Stores field to given value.

### HasStores

`func (o *BRANDDETAIL) HasStores() bool`

HasStores returns a boolean if a field has been set.

### GetUsageMethods

`func (o *BRANDDETAIL) GetUsageMethods() []USAGEMETHODSCHEMA`

GetUsageMethods returns the UsageMethods field if non-nil, zero value otherwise.

### GetUsageMethodsOk

`func (o *BRANDDETAIL) GetUsageMethodsOk() (*[]USAGEMETHODSCHEMA, bool)`

GetUsageMethodsOk returns a tuple with the UsageMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMethods

`func (o *BRANDDETAIL) SetUsageMethods(v []USAGEMETHODSCHEMA)`

SetUsageMethods sets UsageMethods field to given value.

### HasUsageMethods

`func (o *BRANDDETAIL) HasUsageMethods() bool`

HasUsageMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


