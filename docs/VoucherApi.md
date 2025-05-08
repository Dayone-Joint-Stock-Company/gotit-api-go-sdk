# \VoucherAPI

All URIs are relative to *https://api-biz-stg.gotit.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVoucherLinkE**](VoucherAPI.md#CreateVoucherLinkE) | **Post** /api/v4.0/vouchers/e | Create voucher link e
[**CreateVoucherLinkG**](VoucherAPI.md#CreateVoucherLinkG) | **Post** /api/v4.0/vouchers/g | Create voucher link g
[**CreateVoucherLinkV**](VoucherAPI.md#CreateVoucherLinkV) | **Post** /api/v4.0/vouchers/v | Create voucher link v



## CreateVoucherLinkE

> VOUCHERERESPONSE CreateVoucherLinkE(ctx).XGIAuthorization(xGIAuthorization).REQUESTCREATEVOUCHERLINKE(rEQUESTCREATEVOUCHERLINKE).Execute()

Create voucher link e



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
	rEQUESTCREATEVOUCHERLINKE := *openapiclient.NewREQUESTCREATEVOUCHERLINKE("Got It Promotion - Jul 2023", "1000007_123abc", int64(100000), "2024-11-15", int64(1), "1gEGcntd") // REQUESTCREATEVOUCHERLINKE |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.CreateVoucherLinkE(context.Background()).XGIAuthorization(xGIAuthorization).REQUESTCREATEVOUCHERLINKE(rEQUESTCREATEVOUCHERLINKE).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.CreateVoucherLinkE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVoucherLinkE`: VOUCHERERESPONSE
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.CreateVoucherLinkE`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoucherLinkERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 
 **rEQUESTCREATEVOUCHERLINKE** | [**REQUESTCREATEVOUCHERLINKE**](REQUESTCREATEVOUCHERLINKE.md) |  | 

### Return type

[**VOUCHERERESPONSE**](VOUCHERERESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVoucherLinkG

> VOUCHERGRESPONSE CreateVoucherLinkG(ctx).XGIAuthorization(xGIAuthorization).REQUESTCREATEVOUCHERLINKG(rEQUESTCREATEVOUCHERLINKG).Execute()

Create voucher link g



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
	rEQUESTCREATEVOUCHERLINKG := *openapiclient.NewREQUESTCREATEVOUCHERLINKG([]openapiclient.PRODUCTDEFAULTLINKG{*openapiclient.NewPRODUCTDEFAULTLINKG(int64(1541), int64(2991), int64(1))}, "Got It Promotion - Jul 2023", "2024-11-15", "1000007_123abc") // REQUESTCREATEVOUCHERLINKG |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.CreateVoucherLinkG(context.Background()).XGIAuthorization(xGIAuthorization).REQUESTCREATEVOUCHERLINKG(rEQUESTCREATEVOUCHERLINKG).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.CreateVoucherLinkG``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVoucherLinkG`: VOUCHERGRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.CreateVoucherLinkG`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoucherLinkGRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 
 **rEQUESTCREATEVOUCHERLINKG** | [**REQUESTCREATEVOUCHERLINKG**](REQUESTCREATEVOUCHERLINKG.md) |  | 

### Return type

[**VOUCHERGRESPONSE**](VOUCHERGRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVoucherLinkV

> VOUCHERVRESPONSE CreateVoucherLinkV(ctx).XGIAuthorization(xGIAuthorization).REQUESTCREATEVOUCHERLINKV(rEQUESTCREATEVOUCHERLINKV).Execute()

Create voucher link v



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
	rEQUESTCREATEVOUCHERLINKV := *openapiclient.NewREQUESTCREATEVOUCHERLINKV(int64(1541), int64(2991), int64(1), "Got It Promotion - Jul 2023", "2024-11-15", "1000007_123abc") // REQUESTCREATEVOUCHERLINKV |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherAPI.CreateVoucherLinkV(context.Background()).XGIAuthorization(xGIAuthorization).REQUESTCREATEVOUCHERLINKV(rEQUESTCREATEVOUCHERLINKV).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherAPI.CreateVoucherLinkV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVoucherLinkV`: VOUCHERVRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `VoucherAPI.CreateVoucherLinkV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoucherLinkVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 
 **rEQUESTCREATEVOUCHERLINKV** | [**REQUESTCREATEVOUCHERLINKV**](REQUESTCREATEVOUCHERLINKV.md) |  | 

### Return type

[**VOUCHERVRESPONSE**](VOUCHERVRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

