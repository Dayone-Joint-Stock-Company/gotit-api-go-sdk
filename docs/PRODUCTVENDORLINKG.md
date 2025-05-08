# PRODUCTVENDORLINKG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **int64** | Product Id. You can see the Products section for the product id. Note: productId type auto convert merchant code does not support creating cover links. | 
**ProductPriceId** | **int64** | Product price id. Each product will have 1 or more denomination codes corresponding to the value or converted product size. | 
**Quantity** | **int64** | Number of vouchers to be issued | 

## Methods

### NewPRODUCTVENDORLINKG

`func NewPRODUCTVENDORLINKG(productId int64, productPriceId int64, quantity int64, ) *PRODUCTVENDORLINKG`

NewPRODUCTVENDORLINKG instantiates a new PRODUCTVENDORLINKG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPRODUCTVENDORLINKGWithDefaults

`func NewPRODUCTVENDORLINKGWithDefaults() *PRODUCTVENDORLINKG`

NewPRODUCTVENDORLINKGWithDefaults instantiates a new PRODUCTVENDORLINKG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *PRODUCTVENDORLINKG) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PRODUCTVENDORLINKG) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PRODUCTVENDORLINKG) SetProductId(v int64)`

SetProductId sets ProductId field to given value.


### GetProductPriceId

`func (o *PRODUCTVENDORLINKG) GetProductPriceId() int64`

GetProductPriceId returns the ProductPriceId field if non-nil, zero value otherwise.

### GetProductPriceIdOk

`func (o *PRODUCTVENDORLINKG) GetProductPriceIdOk() (*int64, bool)`

GetProductPriceIdOk returns a tuple with the ProductPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPriceId

`func (o *PRODUCTVENDORLINKG) SetProductPriceId(v int64)`

SetProductPriceId sets ProductPriceId field to given value.


### GetQuantity

`func (o *PRODUCTVENDORLINKG) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PRODUCTVENDORLINKG) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PRODUCTVENDORLINKG) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


