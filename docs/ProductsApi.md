# \ProductsAPI

All URIs are relative to *https://api-biz-stg.gotit.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListOfProducts**](ProductsAPI.md#GetListOfProducts) | **Get** /api/v4.0/products | Get all products master data
[**GetProductDetail**](ProductsAPI.md#GetProductDetail) | **Get** /api/v4.0/products/{id} | Get product detail data
[**GetStoresOfProduct**](ProductsAPI.md#GetStoresOfProduct) | **Get** /api/v4.0/products/{id}/stores | Get stores of this product



## GetListOfProducts

> PRODUCTSRESPONSE GetListOfProducts(ctx).XGIAuthorization(xGIAuthorization).Page(page).PageSize(pageSize).MinPrice(minPrice).MaxPrice(maxPrice).IsExcludeStoreListInfo(isExcludeStoreListInfo).StoreListPage(storeListPage).StoreListPageSize(storeListPageSize).Execute()

Get all products master data



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
	page := int64(1) // int64 | Page
	pageSize := int64(200) // int64 | Page Size
	minPrice := int64(1000) // int64 | Min price (optional)
	maxPrice := int64(10000000) // int64 | Max price (optional)
	isExcludeStoreListInfo := true // bool | Stores (optional)
	storeListPage := int64(1) // int64 | Store Page (optional)
	storeListPageSize := int64(200) // int64 | Store Page Size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.GetListOfProducts(context.Background()).XGIAuthorization(xGIAuthorization).Page(page).PageSize(pageSize).MinPrice(minPrice).MaxPrice(maxPrice).IsExcludeStoreListInfo(isExcludeStoreListInfo).StoreListPage(storeListPage).StoreListPageSize(storeListPageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetListOfProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListOfProducts`: PRODUCTSRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetListOfProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 
 **page** | **int64** | Page | 
 **pageSize** | **int64** | Page Size | 
 **minPrice** | **int64** | Min price | 
 **maxPrice** | **int64** | Max price | 
 **isExcludeStoreListInfo** | **bool** | Stores | 
 **storeListPage** | **int64** | Store Page | 
 **storeListPageSize** | **int64** | Store Page Size | 

### Return type

[**PRODUCTSRESPONSE**](PRODUCTSRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductDetail

> PRODUCTDETAILRESPONSE GetProductDetail(ctx, id).XGIAuthorization(xGIAuthorization).IsExcludeStoreListInfo(isExcludeStoreListInfo).StoreListPage(storeListPage).StoreListPageSize(storeListPageSize).Execute()

Get product detail data



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
	id := int64(1541) // int64 | Product ID
	isExcludeStoreListInfo := true // bool | Exclude store list information (optional)
	storeListPage := int64(1) // int64 | Store Page (optional)
	storeListPageSize := int64(200) // int64 | Store Page Size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.GetProductDetail(context.Background(), id).XGIAuthorization(xGIAuthorization).IsExcludeStoreListInfo(isExcludeStoreListInfo).StoreListPage(storeListPage).StoreListPageSize(storeListPageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetProductDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductDetail`: PRODUCTDETAILRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetProductDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 

 **isExcludeStoreListInfo** | **bool** | Exclude store list information | 
 **storeListPage** | **int64** | Store Page | 
 **storeListPageSize** | **int64** | Store Page Size | 

### Return type

[**PRODUCTDETAILRESPONSE**](PRODUCTDETAILRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoresOfProduct

> STORESRESPONSE GetStoresOfProduct(ctx, id).XGIAuthorization(xGIAuthorization).Execute()

Get stores of this product



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
	id := int64(1541) // int64 | Product ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.GetStoresOfProduct(context.Background(), id).XGIAuthorization(xGIAuthorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetStoresOfProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStoresOfProduct`: STORESRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetStoresOfProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoresOfProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 


### Return type

[**STORESRESPONSE**](STORESRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

