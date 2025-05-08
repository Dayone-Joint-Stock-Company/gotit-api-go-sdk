# \VoucherStatusAPI

All URIs are relative to *https://api-biz-stg.gotit.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckVoucher**](VoucherStatusAPI.md#CheckVoucher) | **Get** /api/v4.0/vouchers/multiple/status/{refId} | Check voucher status



## CheckVoucher

> VOUCHERCHECKRESPONSE CheckVoucher(ctx, refId).XGIAuthorization(xGIAuthorization).IsExcludeStoreListInfo(isExcludeStoreListInfo).StoreListPage(storeListPage).StoreListPageSize(storeListPageSize).Execute()

Check voucher status



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
	refId := "voucher_0711_102" // string | Transaction RefId
	isExcludeStoreListInfo := true // bool | Exclude store list information (optional)
	storeListPage := int64(1) // int64 | Store Page (optional)
	storeListPageSize := int64(200) // int64 | Store Page Size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherStatusAPI.CheckVoucher(context.Background(), refId).XGIAuthorization(xGIAuthorization).IsExcludeStoreListInfo(isExcludeStoreListInfo).StoreListPage(storeListPage).StoreListPageSize(storeListPageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherStatusAPI.CheckVoucher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckVoucher`: VOUCHERCHECKRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `VoucherStatusAPI.CheckVoucher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refId** | **string** | Transaction RefId | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckVoucherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 

 **isExcludeStoreListInfo** | **bool** | Exclude store list information | 
 **storeListPage** | **int64** | Store Page | 
 **storeListPageSize** | **int64** | Store Page Size | 

### Return type

[**VOUCHERCHECKRESPONSE**](VOUCHERCHECKRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

