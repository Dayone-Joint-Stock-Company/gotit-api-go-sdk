# STORESSCHEMA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreId** | Pointer to **NullableInt64** | Store Id | [optional] 
**StoreNm** | Pointer to **NullableString** | Store Name | [optional] 
**StoreAddr** | Pointer to **NullableString** | Store Address | [optional] 
**StoreEm** | Pointer to **NullableString** | Store Email | [optional] 
**StorePh** | Pointer to **NullableString** | Store Phone | [optional] 
**Lat** | Pointer to **NullableFloat32** | Lat location on Google maps | [optional] 
**Long** | Pointer to **NullableFloat32** | Long location on Google maps | [optional] 
**BrandNm** | Pointer to **NullableString** | Brand Name | [optional] 
**DistrictId** | Pointer to **NullableInt64** | District code (Got It identifier) | [optional] 
**DistrictNm** | Pointer to **NullableString** | District name | [optional] 
**CityId** | Pointer to **NullableInt64** | City code (Got It identifier) | [optional] 
**CityNm** | Pointer to **NullableString** | City name | [optional] 

## Methods

### NewSTORESSCHEMA

`func NewSTORESSCHEMA() *STORESSCHEMA`

NewSTORESSCHEMA instantiates a new STORESSCHEMA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSTORESSCHEMAWithDefaults

`func NewSTORESSCHEMAWithDefaults() *STORESSCHEMA`

NewSTORESSCHEMAWithDefaults instantiates a new STORESSCHEMA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreId

`func (o *STORESSCHEMA) GetStoreId() int64`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *STORESSCHEMA) GetStoreIdOk() (*int64, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *STORESSCHEMA) SetStoreId(v int64)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *STORESSCHEMA) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### SetStoreIdNil

`func (o *STORESSCHEMA) SetStoreIdNil(b bool)`

 SetStoreIdNil sets the value for StoreId to be an explicit nil

### UnsetStoreId
`func (o *STORESSCHEMA) UnsetStoreId()`

UnsetStoreId ensures that no value is present for StoreId, not even an explicit nil
### GetStoreNm

`func (o *STORESSCHEMA) GetStoreNm() string`

GetStoreNm returns the StoreNm field if non-nil, zero value otherwise.

### GetStoreNmOk

`func (o *STORESSCHEMA) GetStoreNmOk() (*string, bool)`

GetStoreNmOk returns a tuple with the StoreNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreNm

`func (o *STORESSCHEMA) SetStoreNm(v string)`

SetStoreNm sets StoreNm field to given value.

### HasStoreNm

`func (o *STORESSCHEMA) HasStoreNm() bool`

HasStoreNm returns a boolean if a field has been set.

### SetStoreNmNil

`func (o *STORESSCHEMA) SetStoreNmNil(b bool)`

 SetStoreNmNil sets the value for StoreNm to be an explicit nil

### UnsetStoreNm
`func (o *STORESSCHEMA) UnsetStoreNm()`

UnsetStoreNm ensures that no value is present for StoreNm, not even an explicit nil
### GetStoreAddr

`func (o *STORESSCHEMA) GetStoreAddr() string`

GetStoreAddr returns the StoreAddr field if non-nil, zero value otherwise.

### GetStoreAddrOk

`func (o *STORESSCHEMA) GetStoreAddrOk() (*string, bool)`

GetStoreAddrOk returns a tuple with the StoreAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAddr

`func (o *STORESSCHEMA) SetStoreAddr(v string)`

SetStoreAddr sets StoreAddr field to given value.

### HasStoreAddr

`func (o *STORESSCHEMA) HasStoreAddr() bool`

HasStoreAddr returns a boolean if a field has been set.

### SetStoreAddrNil

`func (o *STORESSCHEMA) SetStoreAddrNil(b bool)`

 SetStoreAddrNil sets the value for StoreAddr to be an explicit nil

### UnsetStoreAddr
`func (o *STORESSCHEMA) UnsetStoreAddr()`

UnsetStoreAddr ensures that no value is present for StoreAddr, not even an explicit nil
### GetStoreEm

`func (o *STORESSCHEMA) GetStoreEm() string`

GetStoreEm returns the StoreEm field if non-nil, zero value otherwise.

### GetStoreEmOk

`func (o *STORESSCHEMA) GetStoreEmOk() (*string, bool)`

GetStoreEmOk returns a tuple with the StoreEm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreEm

`func (o *STORESSCHEMA) SetStoreEm(v string)`

SetStoreEm sets StoreEm field to given value.

### HasStoreEm

`func (o *STORESSCHEMA) HasStoreEm() bool`

HasStoreEm returns a boolean if a field has been set.

### SetStoreEmNil

`func (o *STORESSCHEMA) SetStoreEmNil(b bool)`

 SetStoreEmNil sets the value for StoreEm to be an explicit nil

### UnsetStoreEm
`func (o *STORESSCHEMA) UnsetStoreEm()`

UnsetStoreEm ensures that no value is present for StoreEm, not even an explicit nil
### GetStorePh

`func (o *STORESSCHEMA) GetStorePh() string`

GetStorePh returns the StorePh field if non-nil, zero value otherwise.

### GetStorePhOk

`func (o *STORESSCHEMA) GetStorePhOk() (*string, bool)`

GetStorePhOk returns a tuple with the StorePh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePh

`func (o *STORESSCHEMA) SetStorePh(v string)`

SetStorePh sets StorePh field to given value.

### HasStorePh

`func (o *STORESSCHEMA) HasStorePh() bool`

HasStorePh returns a boolean if a field has been set.

### SetStorePhNil

`func (o *STORESSCHEMA) SetStorePhNil(b bool)`

 SetStorePhNil sets the value for StorePh to be an explicit nil

### UnsetStorePh
`func (o *STORESSCHEMA) UnsetStorePh()`

UnsetStorePh ensures that no value is present for StorePh, not even an explicit nil
### GetLat

`func (o *STORESSCHEMA) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *STORESSCHEMA) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *STORESSCHEMA) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *STORESSCHEMA) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *STORESSCHEMA) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *STORESSCHEMA) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLong

`func (o *STORESSCHEMA) GetLong() float32`

GetLong returns the Long field if non-nil, zero value otherwise.

### GetLongOk

`func (o *STORESSCHEMA) GetLongOk() (*float32, bool)`

GetLongOk returns a tuple with the Long field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLong

`func (o *STORESSCHEMA) SetLong(v float32)`

SetLong sets Long field to given value.

### HasLong

`func (o *STORESSCHEMA) HasLong() bool`

HasLong returns a boolean if a field has been set.

### SetLongNil

`func (o *STORESSCHEMA) SetLongNil(b bool)`

 SetLongNil sets the value for Long to be an explicit nil

### UnsetLong
`func (o *STORESSCHEMA) UnsetLong()`

UnsetLong ensures that no value is present for Long, not even an explicit nil
### GetBrandNm

`func (o *STORESSCHEMA) GetBrandNm() string`

GetBrandNm returns the BrandNm field if non-nil, zero value otherwise.

### GetBrandNmOk

`func (o *STORESSCHEMA) GetBrandNmOk() (*string, bool)`

GetBrandNmOk returns a tuple with the BrandNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNm

`func (o *STORESSCHEMA) SetBrandNm(v string)`

SetBrandNm sets BrandNm field to given value.

### HasBrandNm

`func (o *STORESSCHEMA) HasBrandNm() bool`

HasBrandNm returns a boolean if a field has been set.

### SetBrandNmNil

`func (o *STORESSCHEMA) SetBrandNmNil(b bool)`

 SetBrandNmNil sets the value for BrandNm to be an explicit nil

### UnsetBrandNm
`func (o *STORESSCHEMA) UnsetBrandNm()`

UnsetBrandNm ensures that no value is present for BrandNm, not even an explicit nil
### GetDistrictId

`func (o *STORESSCHEMA) GetDistrictId() int64`

GetDistrictId returns the DistrictId field if non-nil, zero value otherwise.

### GetDistrictIdOk

`func (o *STORESSCHEMA) GetDistrictIdOk() (*int64, bool)`

GetDistrictIdOk returns a tuple with the DistrictId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictId

`func (o *STORESSCHEMA) SetDistrictId(v int64)`

SetDistrictId sets DistrictId field to given value.

### HasDistrictId

`func (o *STORESSCHEMA) HasDistrictId() bool`

HasDistrictId returns a boolean if a field has been set.

### SetDistrictIdNil

`func (o *STORESSCHEMA) SetDistrictIdNil(b bool)`

 SetDistrictIdNil sets the value for DistrictId to be an explicit nil

### UnsetDistrictId
`func (o *STORESSCHEMA) UnsetDistrictId()`

UnsetDistrictId ensures that no value is present for DistrictId, not even an explicit nil
### GetDistrictNm

`func (o *STORESSCHEMA) GetDistrictNm() string`

GetDistrictNm returns the DistrictNm field if non-nil, zero value otherwise.

### GetDistrictNmOk

`func (o *STORESSCHEMA) GetDistrictNmOk() (*string, bool)`

GetDistrictNmOk returns a tuple with the DistrictNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictNm

`func (o *STORESSCHEMA) SetDistrictNm(v string)`

SetDistrictNm sets DistrictNm field to given value.

### HasDistrictNm

`func (o *STORESSCHEMA) HasDistrictNm() bool`

HasDistrictNm returns a boolean if a field has been set.

### SetDistrictNmNil

`func (o *STORESSCHEMA) SetDistrictNmNil(b bool)`

 SetDistrictNmNil sets the value for DistrictNm to be an explicit nil

### UnsetDistrictNm
`func (o *STORESSCHEMA) UnsetDistrictNm()`

UnsetDistrictNm ensures that no value is present for DistrictNm, not even an explicit nil
### GetCityId

`func (o *STORESSCHEMA) GetCityId() int64`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *STORESSCHEMA) GetCityIdOk() (*int64, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *STORESSCHEMA) SetCityId(v int64)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *STORESSCHEMA) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *STORESSCHEMA) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *STORESSCHEMA) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetCityNm

`func (o *STORESSCHEMA) GetCityNm() string`

GetCityNm returns the CityNm field if non-nil, zero value otherwise.

### GetCityNmOk

`func (o *STORESSCHEMA) GetCityNmOk() (*string, bool)`

GetCityNmOk returns a tuple with the CityNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityNm

`func (o *STORESSCHEMA) SetCityNm(v string)`

SetCityNm sets CityNm field to given value.

### HasCityNm

`func (o *STORESSCHEMA) HasCityNm() bool`

HasCityNm returns a boolean if a field has been set.

### SetCityNmNil

`func (o *STORESSCHEMA) SetCityNmNil(b bool)`

 SetCityNmNil sets the value for CityNm to be an explicit nil

### UnsetCityNm
`func (o *STORESSCHEMA) UnsetCityNm()`

UnsetCityNm ensures that no value is present for CityNm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


