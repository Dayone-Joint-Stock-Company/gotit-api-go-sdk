# \BrandsAPI

All URIs are relative to *https://api-biz-stg.gotit.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrandsByProducts**](BrandsAPI.md#BrandsByProducts) | **Get** /api/v4.0/brands/brandsByProducts | Get brand by product
[**GetDetailOfBrand**](BrandsAPI.md#GetDetailOfBrand) | **Get** /api/v4.0/brands/{id} | Get brand detail
[**GetListOfBrands**](BrandsAPI.md#GetListOfBrands) | **Get** /api/v4.0/brands | Get list of brands



## BrandsByProducts

> BRANDDETAILRESPONSE BrandsByProducts(ctx).XGIAuthorization(xGIAuthorization).Execute()

Get brand by product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xGIAuthorization := "API key GotIt provided" // string | Authorization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.BrandsByProducts(context.Background()).XGIAuthorization(xGIAuthorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.BrandsByProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrandsByProducts`: BRANDDETAILRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.BrandsByProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrandsByProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 

### Return type

[**BRANDDETAILRESPONSE**](BRANDDETAILRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetailOfBrand

> BRANDDETAILRESPONSE GetDetailOfBrand(ctx, id).XGIAuthorization(xGIAuthorization).Execute()

Get brand detail



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xGIAuthorization := "API key GotIt provided" // string | Authorization
	id := int64(46) // int64 | Brand Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.GetDetailOfBrand(context.Background(), id).XGIAuthorization(xGIAuthorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.GetDetailOfBrand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetailOfBrand`: BRANDDETAILRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.GetDetailOfBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Brand Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailOfBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 


### Return type

[**BRANDDETAILRESPONSE**](BRANDDETAILRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfBrands

> BRANDSRESPONSE GetListOfBrands(ctx).XGIAuthorization(xGIAuthorization).Execute()

Get list of brands



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xGIAuthorization := "API key GotIt provided" // string | Authorization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.GetListOfBrands(context.Background()).XGIAuthorization(xGIAuthorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.GetListOfBrands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListOfBrands`: BRANDSRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.GetListOfBrands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfBrandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 

### Return type

[**BRANDSRESPONSE**](BRANDSRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

