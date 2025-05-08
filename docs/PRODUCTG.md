# PRODUCTG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductImage** | Pointer to **NullableString** |  | [optional] 
**Link** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Serial** | Pointer to **NullableString** |  | [optional] 
**ProductId** | Pointer to **NullableInt64** |  | [optional] 
**PriceId** | Pointer to **NullableInt64** |  | [optional] 
**Value** | Pointer to **NullableInt64** |  | [optional] 
**ExpiredDate** | Pointer to **NullableString** |  | [optional] 
**Partner** | Pointer to [**VENDORSCHEMA**](VENDORSCHEMA.md) |  | [optional] 

## Methods

### NewPRODUCTG

`func NewPRODUCTG() *PRODUCTG`

NewPRODUCTG instantiates a new PRODUCTG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPRODUCTGWithDefaults

`func NewPRODUCTGWithDefaults() *PRODUCTG`

NewPRODUCTGWithDefaults instantiates a new PRODUCTG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductImage

`func (o *PRODUCTG) GetProductImage() string`

GetProductImage returns the ProductImage field if non-nil, zero value otherwise.

### GetProductImageOk

`func (o *PRODUCTG) GetProductImageOk() (*string, bool)`

GetProductImageOk returns a tuple with the ProductImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImage

`func (o *PRODUCTG) SetProductImage(v string)`

SetProductImage sets ProductImage field to given value.

### HasProductImage

`func (o *PRODUCTG) HasProductImage() bool`

HasProductImage returns a boolean if a field has been set.

### SetProductImageNil

`func (o *PRODUCTG) SetProductImageNil(b bool)`

 SetProductImageNil sets the value for ProductImage to be an explicit nil

### UnsetProductImage
`func (o *PRODUCTG) UnsetProductImage()`

UnsetProductImage ensures that no value is present for ProductImage, not even an explicit nil
### GetLink

`func (o *PRODUCTG) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *PRODUCTG) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *PRODUCTG) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *PRODUCTG) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *PRODUCTG) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *PRODUCTG) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetCode

`func (o *PRODUCTG) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PRODUCTG) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PRODUCTG) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PRODUCTG) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *PRODUCTG) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PRODUCTG) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetSerial

`func (o *PRODUCTG) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PRODUCTG) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PRODUCTG) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PRODUCTG) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *PRODUCTG) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *PRODUCTG) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetProductId

`func (o *PRODUCTG) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PRODUCTG) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PRODUCTG) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PRODUCTG) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *PRODUCTG) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *PRODUCTG) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetPriceId

`func (o *PRODUCTG) GetPriceId() int64`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *PRODUCTG) GetPriceIdOk() (*int64, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *PRODUCTG) SetPriceId(v int64)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *PRODUCTG) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### SetPriceIdNil

`func (o *PRODUCTG) SetPriceIdNil(b bool)`

 SetPriceIdNil sets the value for PriceId to be an explicit nil

### UnsetPriceId
`func (o *PRODUCTG) UnsetPriceId()`

UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
### GetValue

`func (o *PRODUCTG) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PRODUCTG) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PRODUCTG) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *PRODUCTG) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PRODUCTG) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PRODUCTG) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetExpiredDate

`func (o *PRODUCTG) GetExpiredDate() string`

GetExpiredDate returns the ExpiredDate field if non-nil, zero value otherwise.

### GetExpiredDateOk

`func (o *PRODUCTG) GetExpiredDateOk() (*string, bool)`

GetExpiredDateOk returns a tuple with the ExpiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredDate

`func (o *PRODUCTG) SetExpiredDate(v string)`

SetExpiredDate sets ExpiredDate field to given value.

### HasExpiredDate

`func (o *PRODUCTG) HasExpiredDate() bool`

HasExpiredDate returns a boolean if a field has been set.

### SetExpiredDateNil

`func (o *PRODUCTG) SetExpiredDateNil(b bool)`

 SetExpiredDateNil sets the value for ExpiredDate to be an explicit nil

### UnsetExpiredDate
`func (o *PRODUCTG) UnsetExpiredDate()`

UnsetExpiredDate ensures that no value is present for ExpiredDate, not even an explicit nil
### GetPartner

`func (o *PRODUCTG) GetPartner() VENDORSCHEMA`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *PRODUCTG) GetPartnerOk() (*VENDORSCHEMA, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *PRODUCTG) SetPartner(v VENDORSCHEMA)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *PRODUCTG) HasPartner() bool`

HasPartner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


