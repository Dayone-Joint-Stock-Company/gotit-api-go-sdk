# USAGEMETHODSCHEMA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Type of method used &#39;AT_STORE&#39;: at the store &#39;HOTLINE&#39;: via hot line number &#39;WEBSITE&#39;: via website &#39;FORM&#39;: via order form &#39;BUY_IN_WEBVIEW&#39;: via web &#39;BUY_IN_APP&#39;: via App | [optional] 
**Title** | Pointer to **NullableString** | Method title | [optional] 
**Description** | Pointer to **NullableString** | Description of the method | [optional] 
**Order** | Pointer to **NullableInt64** | Number | [optional] 
**Link** | Pointer to **NullableString** | Link to the usage method | [optional] 
**Phone1** | Pointer to **NullableString** | Phone number 1 | [optional] 
**Phone2** | Pointer to **NullableString** | Phone number 2 | [optional] 

## Methods

### NewUSAGEMETHODSCHEMA

`func NewUSAGEMETHODSCHEMA() *USAGEMETHODSCHEMA`

NewUSAGEMETHODSCHEMA instantiates a new USAGEMETHODSCHEMA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUSAGEMETHODSCHEMAWithDefaults

`func NewUSAGEMETHODSCHEMAWithDefaults() *USAGEMETHODSCHEMA`

NewUSAGEMETHODSCHEMAWithDefaults instantiates a new USAGEMETHODSCHEMA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *USAGEMETHODSCHEMA) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *USAGEMETHODSCHEMA) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *USAGEMETHODSCHEMA) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *USAGEMETHODSCHEMA) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *USAGEMETHODSCHEMA) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *USAGEMETHODSCHEMA) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTitle

`func (o *USAGEMETHODSCHEMA) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *USAGEMETHODSCHEMA) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *USAGEMETHODSCHEMA) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *USAGEMETHODSCHEMA) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *USAGEMETHODSCHEMA) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *USAGEMETHODSCHEMA) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *USAGEMETHODSCHEMA) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *USAGEMETHODSCHEMA) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *USAGEMETHODSCHEMA) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *USAGEMETHODSCHEMA) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *USAGEMETHODSCHEMA) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *USAGEMETHODSCHEMA) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *USAGEMETHODSCHEMA) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *USAGEMETHODSCHEMA) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *USAGEMETHODSCHEMA) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *USAGEMETHODSCHEMA) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *USAGEMETHODSCHEMA) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *USAGEMETHODSCHEMA) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetLink

`func (o *USAGEMETHODSCHEMA) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *USAGEMETHODSCHEMA) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *USAGEMETHODSCHEMA) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *USAGEMETHODSCHEMA) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *USAGEMETHODSCHEMA) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *USAGEMETHODSCHEMA) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetPhone1

`func (o *USAGEMETHODSCHEMA) GetPhone1() string`

GetPhone1 returns the Phone1 field if non-nil, zero value otherwise.

### GetPhone1Ok

`func (o *USAGEMETHODSCHEMA) GetPhone1Ok() (*string, bool)`

GetPhone1Ok returns a tuple with the Phone1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone1

`func (o *USAGEMETHODSCHEMA) SetPhone1(v string)`

SetPhone1 sets Phone1 field to given value.

### HasPhone1

`func (o *USAGEMETHODSCHEMA) HasPhone1() bool`

HasPhone1 returns a boolean if a field has been set.

### SetPhone1Nil

`func (o *USAGEMETHODSCHEMA) SetPhone1Nil(b bool)`

 SetPhone1Nil sets the value for Phone1 to be an explicit nil

### UnsetPhone1
`func (o *USAGEMETHODSCHEMA) UnsetPhone1()`

UnsetPhone1 ensures that no value is present for Phone1, not even an explicit nil
### GetPhone2

`func (o *USAGEMETHODSCHEMA) GetPhone2() string`

GetPhone2 returns the Phone2 field if non-nil, zero value otherwise.

### GetPhone2Ok

`func (o *USAGEMETHODSCHEMA) GetPhone2Ok() (*string, bool)`

GetPhone2Ok returns a tuple with the Phone2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone2

`func (o *USAGEMETHODSCHEMA) SetPhone2(v string)`

SetPhone2 sets Phone2 field to given value.

### HasPhone2

`func (o *USAGEMETHODSCHEMA) HasPhone2() bool`

HasPhone2 returns a boolean if a field has been set.

### SetPhone2Nil

`func (o *USAGEMETHODSCHEMA) SetPhone2Nil(b bool)`

 SetPhone2Nil sets the value for Phone2 to be an explicit nil

### UnsetPhone2
`func (o *USAGEMETHODSCHEMA) UnsetPhone2()`

UnsetPhone2 ensures that no value is present for Phone2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


