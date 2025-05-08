# Go API client for gotit_api_go_sdk

SDK Technical document for GotIt APIs


## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import gotit_api_go_sdk "https://github.com/Dayone-Joint-Stock-Company/gotit-api-go-sdk.git"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `gotit_api_go_sdk.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), gotit_api_go_sdk.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `gotit_api_go_sdk.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), gotit_api_go_sdk.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `gotit_api_go_sdk.ContextOperationServerIndices` and `gotit_api_go_sdk.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), gotit_api_go_sdk.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), gotit_api_go_sdk.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api-biz-stg.gotit.vn*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BrandsAPI* | [**BrandsByProducts**](docs/BrandsAPI.md#brandsbyproducts) | **Get** /api/v4.0/brands/brandsByProducts | Get brand by product
*BrandsAPI* | [**GetDetailOfBrand**](docs/BrandsAPI.md#getdetailofbrand) | **Get** /api/v4.0/brands/{id} | Get brand detail
*BrandsAPI* | [**GetListOfBrands**](docs/BrandsAPI.md#getlistofbrands) | **Get** /api/v4.0/brands | Get list of brands
*CategoriesAPI* | [**GetCategoryByProduct**](docs/CategoriesAPI.md#getcategorybyproduct) | **Get** /api/v4.0/categories/categoriesByProducts | Get category by product
*CategoriesAPI* | [**GetListOfCategories**](docs/CategoriesAPI.md#getlistofcategories) | **Get** /api/v4.0/categories | Get lists category
*ProductsAPI* | [**GetListOfProducts**](docs/ProductsAPI.md#getlistofproducts) | **Get** /api/v4.0/products | Get all products master data
*ProductsAPI* | [**GetProductDetail**](docs/ProductsAPI.md#getproductdetail) | **Get** /api/v4.0/products/{id} | Get product detail data
*ProductsAPI* | [**GetStoresOfProduct**](docs/ProductsAPI.md#getstoresofproduct) | **Get** /api/v4.0/products/{id}/stores | Get stores of this product
*VoucherAPI* | [**CreateVoucherLinkE**](docs/VoucherAPI.md#createvoucherlinke) | **Post** /api/v4.0/vouchers/e | Create voucher link e
*VoucherAPI* | [**CreateVoucherLinkG**](docs/VoucherAPI.md#createvoucherlinkg) | **Post** /api/v4.0/vouchers/g | Create voucher link g
*VoucherAPI* | [**CreateVoucherLinkV**](docs/VoucherAPI.md#createvoucherlinkv) | **Post** /api/v4.0/vouchers/v | Create voucher link v
*VoucherSendMethodAPI* | [**CheckStatusZNS**](docs/VoucherSendMethodAPI.md#checkstatuszns) | **Post** /api/v4.0/vouchers/send/zns/check | Check status zns
*VoucherSendMethodAPI* | [**SendVoucherByEmail**](docs/VoucherSendMethodAPI.md#sendvoucherbyemail) | **Post** /api/v4.0/vouchers/send/email | Send voucher by mail
*VoucherSendMethodAPI* | [**SendVoucherBySMS**](docs/VoucherSendMethodAPI.md#sendvoucherbysms) | **Post** /api/v4.0/vouchers/send/sms | Send voucher by sms
*VoucherSendMethodAPI* | [**SendVoucherByZNS**](docs/VoucherSendMethodAPI.md#sendvoucherbyzns) | **Post** /api/v4.0/vouchers/send/zns | Send voucher by zns
*VoucherStatusAPI* | [**CheckVoucher**](docs/VoucherStatusAPI.md#checkvoucher) | **Get** /api/v4.0/vouchers/multiple/status/{refId} | Check voucher status


## Documentation For Models

 - [BRANDCATEGORYDETAIL](docs/BRANDCATEGORYDETAIL.md)
 - [BRANDCATEGORYDETAILRESPONSE](docs/BRANDCATEGORYDETAILRESPONSE.md)
 - [BRANDDETAIL](docs/BRANDDETAIL.md)
 - [BRANDDETAILRESPONSE](docs/BRANDDETAILRESPONSE.md)
 - [BRANDREEDEMSCHEMA](docs/BRANDREEDEMSCHEMA.md)
 - [BRANDSDETAIL](docs/BRANDSDETAIL.md)
 - [BRANDSRESPONSE](docs/BRANDSRESPONSE.md)
 - [CATEGORIESDETAIL](docs/CATEGORIESDETAIL.md)
 - [CATEGORIESRESPONSE](docs/CATEGORIESRESPONSE.md)
 - [GROUPVOUCHERSCHEMA](docs/GROUPVOUCHERSCHEMA.md)
 - [HTTPBADREQUEST](docs/HTTPBADREQUEST.md)
 - [HTTPINTERNALSERVERERROR](docs/HTTPINTERNALSERVERERROR.md)
 - [HTTPNOTFOUND](docs/HTTPNOTFOUND.md)
 - [HTTPUNAUTHORIZED](docs/HTTPUNAUTHORIZED.md)
 - [PAGINGSCHEMA](docs/PAGINGSCHEMA.md)
 - [PRODUCTDEFAULTLINKG](docs/PRODUCTDEFAULTLINKG.md)
 - [PRODUCTDETAIL](docs/PRODUCTDETAIL.md)
 - [PRODUCTDETAILRESPONSE](docs/PRODUCTDETAILRESPONSE.md)
 - [PRODUCTG](docs/PRODUCTG.md)
 - [PRODUCTPRICESCHEMA](docs/PRODUCTPRICESCHEMA.md)
 - [PRODUCTSALLDETAIL](docs/PRODUCTSALLDETAIL.md)
 - [PRODUCTSDEFAULTLINKG](docs/PRODUCTSDEFAULTLINKG.md)
 - [PRODUCTSDETAIL](docs/PRODUCTSDETAIL.md)
 - [PRODUCTSRESPONSE](docs/PRODUCTSRESPONSE.md)
 - [PRODUCTSRESPONSEDataInner](docs/PRODUCTSRESPONSEDataInner.md)
 - [PRODUCTSVOUCHERCHECK](docs/PRODUCTSVOUCHERCHECK.md)
 - [PRODUCTV](docs/PRODUCTV.md)
 - [PRODUCTVENDORLINKG](docs/PRODUCTVENDORLINKG.md)
 - [REQUESTCHECKSTATUSZNS](docs/REQUESTCHECKSTATUSZNS.md)
 - [REQUESTCREATEVOUCHERLINKE](docs/REQUESTCREATEVOUCHERLINKE.md)
 - [REQUESTCREATEVOUCHERLINKG](docs/REQUESTCREATEVOUCHERLINKG.md)
 - [REQUESTCREATEVOUCHERLINKV](docs/REQUESTCREATEVOUCHERLINKV.md)
 - [REQUESTSENDVOUCHERBYEMAIL](docs/REQUESTSENDVOUCHERBYEMAIL.md)
 - [REQUESTSENDVOUCHERBYSMS](docs/REQUESTSENDVOUCHERBYSMS.md)
 - [REQUESTSENDVOUCHERBYZNS](docs/REQUESTSENDVOUCHERBYZNS.md)
 - [STOREPAGINGSCHEMA](docs/STOREPAGINGSCHEMA.md)
 - [STOREPRODUCTSCHEMA](docs/STOREPRODUCTSCHEMA.md)
 - [STORESRESPONSE](docs/STORESRESPONSE.md)
 - [STORESRESPONSEDataInner](docs/STORESRESPONSEDataInner.md)
 - [STORESSCHEMA](docs/STORESSCHEMA.md)
 - [USAGEMETHODSCHEMA](docs/USAGEMETHODSCHEMA.md)
 - [VENDORSCHEMA](docs/VENDORSCHEMA.md)
 - [VOUCHERCHECKRESPONSE](docs/VOUCHERCHECKRESPONSE.md)
 - [VOUCHERCHECKSCHEMA](docs/VOUCHERCHECKSCHEMA.md)
 - [VOUCHERCHECKSCHEMADETAIL](docs/VOUCHERCHECKSCHEMADETAIL.md)
 - [VOUCHERCHECKZNSRESPONSE](docs/VOUCHERCHECKZNSRESPONSE.md)
 - [VOUCHERCHECKZNSRESPONSEData](docs/VOUCHERCHECKZNSRESPONSEData.md)
 - [VOUCHERE](docs/VOUCHERE.md)
 - [VOUCHERERESPONSE](docs/VOUCHERERESPONSE.md)
 - [VOUCHERESCHEMA](docs/VOUCHERESCHEMA.md)
 - [VOUCHERG](docs/VOUCHERG.md)
 - [VOUCHERGRESPONSE](docs/VOUCHERGRESPONSE.md)
 - [VOUCHERGSCHEMA](docs/VOUCHERGSCHEMA.md)
 - [VOUCHERSENDEMAILRESPONSE](docs/VOUCHERSENDEMAILRESPONSE.md)
 - [VOUCHERSENDEMAILSCHEMA](docs/VOUCHERSENDEMAILSCHEMA.md)
 - [VOUCHERSENDSMSRESPONSE](docs/VOUCHERSENDSMSRESPONSE.md)
 - [VOUCHERSENDSMSSCHEMA](docs/VOUCHERSENDSMSSCHEMA.md)
 - [VOUCHERSENDZNSRESPONSE](docs/VOUCHERSENDZNSRESPONSE.md)
 - [VOUCHERSENDZNSRESPONSEData](docs/VOUCHERSENDZNSRESPONSEData.md)
 - [VOUCHERV](docs/VOUCHERV.md)
 - [VOUCHERVRESPONSE](docs/VOUCHERVRESPONSE.md)
 - [VOUCHERVSCHEMA](docs/VOUCHERVSCHEMA.md)
 - [VOUCHERVSCHEMAProduct](docs/VOUCHERVSCHEMAProduct.md)



## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

quang.huynh@gotit.vn

## Running Examples with Docker

The `example` directory contains a sample application that demonstrates how to use the SDK. It includes endpoints for brands, categories, products, and vouchers.

### Prerequisites

- Docker
- Docker Compose
- Go 1.21 or later

### Environment Variables

Create a `.env` file in the `example` directory with the following variables:

```env
X_GI_AUTHORIZATION=your_api_key_here
PRIVATE_KEY_STR=your_private_key_here
```

### Installation and Running

1. Navigate to the example directory:
```bash
cd example
```

2. Build and run the application using Docker Compose:
```bash
docker-compose up --build
```

The application will be available at `http://localhost:5004`

### Available Endpoints

- Brands:
  - GET `/brands` - Get list of brands
  - GET `/brands/:id` - Get brand details
  - GET `/brands/brand_product` - Get brand products

- Categories:
  - GET `/categories` - Get list of categories
  - GET `/categories/category_product` - Get category products

- Products:
  - GET `/products` - Get list of products
  - GET `/products/:id` - Get product details
  - GET `/products/:id/stores` - Get product stores

- Vouchers:
  - GET `/vouchers/e` - Generate E vouchers
  - GET `/vouchers/v` - Generate V vouchers
  - GET `/vouchers/g` - Generate G vouchers
  - GET `/vouchers/send_sms` - Send voucher via SMS
  - GET `/vouchers/send_email` - Send voucher via email
  - GET `/vouchers/send_zns` - Send voucher via ZNS
  - GET `/vouchers/check_zns` - Check ZNS status
  - GET `/vouchers/check_status_voucher` - Check voucher status

### Example Usage

```bash
# Get list of brands
curl http://localhost:5004/brands

# Get product details
curl http://localhost:5004/products/1541

# Generate E voucher
curl http://localhost:5004/vouchers/e
```

Note: Make sure to set the required environment variables before running the application. The private key is only required for generating E vouchers.

