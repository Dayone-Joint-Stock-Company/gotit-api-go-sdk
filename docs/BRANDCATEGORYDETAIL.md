# BRANDCATEGORYDETAIL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandId** | Pointer to **NullableInt64** | Brand id | [optional] 
**BrandNm** | Pointer to **NullableString** | Brand name | [optional] 
**BrandLogo** | Pointer to **NullableString** | Link to brand logo image | [optional] 
**ShortDesc** | Pointer to **NullableString** | Brand short description | [optional] 
**Description** | Pointer to **NullableString** | Brand description | [optional] 
**Order** | Pointer to **NullableInt64** |  | [optional] 
**ServiceGuide** | Pointer to **NullableString** | Describe the brand&#39;s terms of reference (T&amp;C). HTML format | [optional] 
**UsageMethods** | Pointer to [**[]USAGEMETHODSCHEMA**](USAGEMETHODSCHEMA.md) | Information on the brand&#39;s applicable channels | [optional] 

## Methods

### NewBRANDCATEGORYDETAIL

`func NewBRANDCATEGORYDETAIL() *BRANDCATEGORYDETAIL`

NewBRANDCATEGORYDETAIL instantiates a new BRANDCATEGORYDETAIL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBRANDCATEGORYDETAILWithDefaults

`func NewBRANDCATEGORYDETAILWithDefaults() *BRANDCATEGORYDETAIL`

NewBRANDCATEGORYDETAILWithDefaults instantiates a new BRANDCATEGORYDETAIL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandId

`func (o *BRANDCATEGORYDETAIL) GetBrandId() int64`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *BRANDCATEGORYDETAIL) GetBrandIdOk() (*int64, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *BRANDCATEGORYDETAIL) SetBrandId(v int64)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *BRANDCATEGORYDETAIL) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### SetBrandIdNil

`func (o *BRANDCATEGORYDETAIL) SetBrandIdNil(b bool)`

 SetBrandIdNil sets the value for BrandId to be an explicit nil

### UnsetBrandId
`func (o *BRANDCATEGORYDETAIL) UnsetBrandId()`

UnsetBrandId ensures that no value is present for BrandId, not even an explicit nil
### GetBrandNm

`func (o *BRANDCATEGORYDETAIL) GetBrandNm() string`

GetBrandNm returns the BrandNm field if non-nil, zero value otherwise.

### GetBrandNmOk

`func (o *BRANDCATEGORYDETAIL) GetBrandNmOk() (*string, bool)`

GetBrandNmOk returns a tuple with the BrandNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNm

`func (o *BRANDCATEGORYDETAIL) SetBrandNm(v string)`

SetBrandNm sets BrandNm field to given value.

### HasBrandNm

`func (o *BRANDCATEGORYDETAIL) HasBrandNm() bool`

HasBrandNm returns a boolean if a field has been set.

### SetBrandNmNil

`func (o *BRANDCATEGORYDETAIL) SetBrandNmNil(b bool)`

 SetBrandNmNil sets the value for BrandNm to be an explicit nil

### UnsetBrandNm
`func (o *BRANDCATEGORYDETAIL) UnsetBrandNm()`

UnsetBrandNm ensures that no value is present for BrandNm, not even an explicit nil
### GetBrandLogo

`func (o *BRANDCATEGORYDETAIL) GetBrandLogo() string`

GetBrandLogo returns the BrandLogo field if non-nil, zero value otherwise.

### GetBrandLogoOk

`func (o *BRANDCATEGORYDETAIL) GetBrandLogoOk() (*string, bool)`

GetBrandLogoOk returns a tuple with the BrandLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandLogo

`func (o *BRANDCATEGORYDETAIL) SetBrandLogo(v string)`

SetBrandLogo sets BrandLogo field to given value.

### HasBrandLogo

`func (o *BRANDCATEGORYDETAIL) HasBrandLogo() bool`

HasBrandLogo returns a boolean if a field has been set.

### SetBrandLogoNil

`func (o *BRANDCATEGORYDETAIL) SetBrandLogoNil(b bool)`

 SetBrandLogoNil sets the value for BrandLogo to be an explicit nil

### UnsetBrandLogo
`func (o *BRANDCATEGORYDETAIL) UnsetBrandLogo()`

UnsetBrandLogo ensures that no value is present for BrandLogo, not even an explicit nil
### GetShortDesc

`func (o *BRANDCATEGORYDETAIL) GetShortDesc() string`

GetShortDesc returns the ShortDesc field if non-nil, zero value otherwise.

### GetShortDescOk

`func (o *BRANDCATEGORYDETAIL) GetShortDescOk() (*string, bool)`

GetShortDescOk returns a tuple with the ShortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDesc

`func (o *BRANDCATEGORYDETAIL) SetShortDesc(v string)`

SetShortDesc sets ShortDesc field to given value.

### HasShortDesc

`func (o *BRANDCATEGORYDETAIL) HasShortDesc() bool`

HasShortDesc returns a boolean if a field has been set.

### SetShortDescNil

`func (o *BRANDCATEGORYDETAIL) SetShortDescNil(b bool)`

 SetShortDescNil sets the value for ShortDesc to be an explicit nil

### UnsetShortDesc
`func (o *BRANDCATEGORYDETAIL) UnsetShortDesc()`

UnsetShortDesc ensures that no value is present for ShortDesc, not even an explicit nil
### GetDescription

`func (o *BRANDCATEGORYDETAIL) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BRANDCATEGORYDETAIL) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BRANDCATEGORYDETAIL) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BRANDCATEGORYDETAIL) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BRANDCATEGORYDETAIL) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BRANDCATEGORYDETAIL) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *BRANDCATEGORYDETAIL) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *BRANDCATEGORYDETAIL) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *BRANDCATEGORYDETAIL) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *BRANDCATEGORYDETAIL) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *BRANDCATEGORYDETAIL) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *BRANDCATEGORYDETAIL) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetServiceGuide

`func (o *BRANDCATEGORYDETAIL) GetServiceGuide() string`

GetServiceGuide returns the ServiceGuide field if non-nil, zero value otherwise.

### GetServiceGuideOk

`func (o *BRANDCATEGORYDETAIL) GetServiceGuideOk() (*string, bool)`

GetServiceGuideOk returns a tuple with the ServiceGuide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGuide

`func (o *BRANDCATEGORYDETAIL) SetServiceGuide(v string)`

SetServiceGuide sets ServiceGuide field to given value.

### HasServiceGuide

`func (o *BRANDCATEGORYDETAIL) HasServiceGuide() bool`

HasServiceGuide returns a boolean if a field has been set.

### SetServiceGuideNil

`func (o *BRANDCATEGORYDETAIL) SetServiceGuideNil(b bool)`

 SetServiceGuideNil sets the value for ServiceGuide to be an explicit nil

### UnsetServiceGuide
`func (o *BRANDCATEGORYDETAIL) UnsetServiceGuide()`

UnsetServiceGuide ensures that no value is present for ServiceGuide, not even an explicit nil
### GetUsageMethods

`func (o *BRANDCATEGORYDETAIL) GetUsageMethods() []USAGEMETHODSCHEMA`

GetUsageMethods returns the UsageMethods field if non-nil, zero value otherwise.

### GetUsageMethodsOk

`func (o *BRANDCATEGORYDETAIL) GetUsageMethodsOk() (*[]USAGEMETHODSCHEMA, bool)`

GetUsageMethodsOk returns a tuple with the UsageMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMethods

`func (o *BRANDCATEGORYDETAIL) SetUsageMethods(v []USAGEMETHODSCHEMA)`

SetUsageMethods sets UsageMethods field to given value.

### HasUsageMethods

`func (o *BRANDCATEGORYDETAIL) HasUsageMethods() bool`

HasUsageMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


