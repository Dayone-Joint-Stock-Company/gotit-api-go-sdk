# STOREPRODUCTSCHEMA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreId** | Pointer to **NullableInt64** | Store Id | [optional] 
**StoreNm** | Pointer to **NullableString** | Store Name | [optional] 
**StoreAddr** | Pointer to **NullableString** | Store Address | [optional] 
**Lat** | Pointer to **NullableFloat32** | Lat location on Google maps | [optional] 
**Long** | Pointer to **NullableFloat32** | Long location on Google maps | [optional] 
**Phone** | Pointer to **NullableString** | Store Phone | [optional] 
**CityId** | Pointer to **NullableInt64** | City code (Got It identifier) | [optional] 
**City** | Pointer to **NullableString** | City name | [optional] 
**DistId** | Pointer to **NullableInt64** | District code (Got It identifier) | [optional] 
**District** | Pointer to **NullableString** | District name | [optional] 

## Methods

### NewSTOREPRODUCTSCHEMA

`func NewSTOREPRODUCTSCHEMA() *STOREPRODUCTSCHEMA`

NewSTOREPRODUCTSCHEMA instantiates a new STOREPRODUCTSCHEMA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSTOREPRODUCTSCHEMAWithDefaults

`func NewSTOREPRODUCTSCHEMAWithDefaults() *STOREPRODUCTSCHEMA`

NewSTOREPRODUCTSCHEMAWithDefaults instantiates a new STOREPRODUCTSCHEMA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreId

`func (o *STOREPRODUCTSCHEMA) GetStoreId() int64`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *STOREPRODUCTSCHEMA) GetStoreIdOk() (*int64, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *STOREPRODUCTSCHEMA) SetStoreId(v int64)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *STOREPRODUCTSCHEMA) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### SetStoreIdNil

`func (o *STOREPRODUCTSCHEMA) SetStoreIdNil(b bool)`

 SetStoreIdNil sets the value for StoreId to be an explicit nil

### UnsetStoreId
`func (o *STOREPRODUCTSCHEMA) UnsetStoreId()`

UnsetStoreId ensures that no value is present for StoreId, not even an explicit nil
### GetStoreNm

`func (o *STOREPRODUCTSCHEMA) GetStoreNm() string`

GetStoreNm returns the StoreNm field if non-nil, zero value otherwise.

### GetStoreNmOk

`func (o *STOREPRODUCTSCHEMA) GetStoreNmOk() (*string, bool)`

GetStoreNmOk returns a tuple with the StoreNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreNm

`func (o *STOREPRODUCTSCHEMA) SetStoreNm(v string)`

SetStoreNm sets StoreNm field to given value.

### HasStoreNm

`func (o *STOREPRODUCTSCHEMA) HasStoreNm() bool`

HasStoreNm returns a boolean if a field has been set.

### SetStoreNmNil

`func (o *STOREPRODUCTSCHEMA) SetStoreNmNil(b bool)`

 SetStoreNmNil sets the value for StoreNm to be an explicit nil

### UnsetStoreNm
`func (o *STOREPRODUCTSCHEMA) UnsetStoreNm()`

UnsetStoreNm ensures that no value is present for StoreNm, not even an explicit nil
### GetStoreAddr

`func (o *STOREPRODUCTSCHEMA) GetStoreAddr() string`

GetStoreAddr returns the StoreAddr field if non-nil, zero value otherwise.

### GetStoreAddrOk

`func (o *STOREPRODUCTSCHEMA) GetStoreAddrOk() (*string, bool)`

GetStoreAddrOk returns a tuple with the StoreAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAddr

`func (o *STOREPRODUCTSCHEMA) SetStoreAddr(v string)`

SetStoreAddr sets StoreAddr field to given value.

### HasStoreAddr

`func (o *STOREPRODUCTSCHEMA) HasStoreAddr() bool`

HasStoreAddr returns a boolean if a field has been set.

### SetStoreAddrNil

`func (o *STOREPRODUCTSCHEMA) SetStoreAddrNil(b bool)`

 SetStoreAddrNil sets the value for StoreAddr to be an explicit nil

### UnsetStoreAddr
`func (o *STOREPRODUCTSCHEMA) UnsetStoreAddr()`

UnsetStoreAddr ensures that no value is present for StoreAddr, not even an explicit nil
### GetLat

`func (o *STOREPRODUCTSCHEMA) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *STOREPRODUCTSCHEMA) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *STOREPRODUCTSCHEMA) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *STOREPRODUCTSCHEMA) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *STOREPRODUCTSCHEMA) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *STOREPRODUCTSCHEMA) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLong

`func (o *STOREPRODUCTSCHEMA) GetLong() float32`

GetLong returns the Long field if non-nil, zero value otherwise.

### GetLongOk

`func (o *STOREPRODUCTSCHEMA) GetLongOk() (*float32, bool)`

GetLongOk returns a tuple with the Long field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLong

`func (o *STOREPRODUCTSCHEMA) SetLong(v float32)`

SetLong sets Long field to given value.

### HasLong

`func (o *STOREPRODUCTSCHEMA) HasLong() bool`

HasLong returns a boolean if a field has been set.

### SetLongNil

`func (o *STOREPRODUCTSCHEMA) SetLongNil(b bool)`

 SetLongNil sets the value for Long to be an explicit nil

### UnsetLong
`func (o *STOREPRODUCTSCHEMA) UnsetLong()`

UnsetLong ensures that no value is present for Long, not even an explicit nil
### GetPhone

`func (o *STOREPRODUCTSCHEMA) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *STOREPRODUCTSCHEMA) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *STOREPRODUCTSCHEMA) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *STOREPRODUCTSCHEMA) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *STOREPRODUCTSCHEMA) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *STOREPRODUCTSCHEMA) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCityId

`func (o *STOREPRODUCTSCHEMA) GetCityId() int64`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *STOREPRODUCTSCHEMA) GetCityIdOk() (*int64, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *STOREPRODUCTSCHEMA) SetCityId(v int64)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *STOREPRODUCTSCHEMA) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *STOREPRODUCTSCHEMA) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *STOREPRODUCTSCHEMA) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetCity

`func (o *STOREPRODUCTSCHEMA) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *STOREPRODUCTSCHEMA) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *STOREPRODUCTSCHEMA) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *STOREPRODUCTSCHEMA) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *STOREPRODUCTSCHEMA) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *STOREPRODUCTSCHEMA) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetDistId

`func (o *STOREPRODUCTSCHEMA) GetDistId() int64`

GetDistId returns the DistId field if non-nil, zero value otherwise.

### GetDistIdOk

`func (o *STOREPRODUCTSCHEMA) GetDistIdOk() (*int64, bool)`

GetDistIdOk returns a tuple with the DistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistId

`func (o *STOREPRODUCTSCHEMA) SetDistId(v int64)`

SetDistId sets DistId field to given value.

### HasDistId

`func (o *STOREPRODUCTSCHEMA) HasDistId() bool`

HasDistId returns a boolean if a field has been set.

### SetDistIdNil

`func (o *STOREPRODUCTSCHEMA) SetDistIdNil(b bool)`

 SetDistIdNil sets the value for DistId to be an explicit nil

### UnsetDistId
`func (o *STOREPRODUCTSCHEMA) UnsetDistId()`

UnsetDistId ensures that no value is present for DistId, not even an explicit nil
### GetDistrict

`func (o *STOREPRODUCTSCHEMA) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *STOREPRODUCTSCHEMA) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *STOREPRODUCTSCHEMA) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *STOREPRODUCTSCHEMA) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *STOREPRODUCTSCHEMA) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *STOREPRODUCTSCHEMA) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


