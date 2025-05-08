# \CategoriesAPI

All URIs are relative to *https://api-biz-stg.gotit.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCategoryByProduct**](CategoriesAPI.md#GetCategoryByProduct) | **Get** /api/v4.0/categories/categoriesByProducts | Get category by product
[**GetListOfCategories**](CategoriesAPI.md#GetListOfCategories) | **Get** /api/v4.0/categories | Get lists category



## GetCategoryByProduct

> CATEGORIESRESPONSE GetCategoryByProduct(ctx).XGIAuthorization(xGIAuthorization).Execute()

Get category by product



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
	resp, r, err := apiClient.CategoriesAPI.GetCategoryByProduct(context.Background()).XGIAuthorization(xGIAuthorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.GetCategoryByProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoryByProduct`: CATEGORIESRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.GetCategoryByProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryByProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 

### Return type

[**CATEGORIESRESPONSE**](CATEGORIESRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfCategories

> CATEGORIESRESPONSE GetListOfCategories(ctx).XGIAuthorization(xGIAuthorization).Execute()

Get lists category



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
	resp, r, err := apiClient.CategoriesAPI.GetListOfCategories(context.Background()).XGIAuthorization(xGIAuthorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.GetListOfCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListOfCategories`: CATEGORIESRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.GetListOfCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 

### Return type

[**CATEGORIESRESPONSE**](CATEGORIESRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

