# \VoucherSendMethodAPI

All URIs are relative to *https://api-biz-stg.gotit.vn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckStatusZNS**](VoucherSendMethodAPI.md#CheckStatusZNS) | **Post** /api/v4.0/vouchers/send/zns/check | Check status zns
[**SendVoucherByEmail**](VoucherSendMethodAPI.md#SendVoucherByEmail) | **Post** /api/v4.0/vouchers/send/email | Send voucher by mail
[**SendVoucherBySMS**](VoucherSendMethodAPI.md#SendVoucherBySMS) | **Post** /api/v4.0/vouchers/send/sms | Send voucher by sms
[**SendVoucherByZNS**](VoucherSendMethodAPI.md#SendVoucherByZNS) | **Post** /api/v4.0/vouchers/send/zns | Send voucher by zns



## CheckStatusZNS

> VOUCHERCHECKZNSRESPONSE CheckStatusZNS(ctx).XGIAuthorization(xGIAuthorization).REQUESTCHECKSTATUSZNS(rEQUESTCHECKSTATUSZNS).Execute()

Check status zns



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
	rEQUESTCHECKSTATUSZNS := *openapiclient.NewREQUESTCHECKSTATUSZNS() // REQUESTCHECKSTATUSZNS |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherSendMethodAPI.CheckStatusZNS(context.Background()).XGIAuthorization(xGIAuthorization).REQUESTCHECKSTATUSZNS(rEQUESTCHECKSTATUSZNS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherSendMethodAPI.CheckStatusZNS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckStatusZNS`: VOUCHERCHECKZNSRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `VoucherSendMethodAPI.CheckStatusZNS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckStatusZNSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 
 **rEQUESTCHECKSTATUSZNS** | [**REQUESTCHECKSTATUSZNS**](REQUESTCHECKSTATUSZNS.md) |  | 

### Return type

[**VOUCHERCHECKZNSRESPONSE**](VOUCHERCHECKZNSRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendVoucherByEmail

> VOUCHERSENDEMAILRESPONSE SendVoucherByEmail(ctx).XGIAuthorization(xGIAuthorization).REQUESTSENDVOUCHERBYEMAIL(rEQUESTSENDVOUCHERBYEMAIL).Execute()

Send voucher by mail



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
	rEQUESTSENDVOUCHERBYEMAIL := *openapiclient.NewREQUESTSENDVOUCHERBYEMAIL("1gEGcntd", "quang.huynh@gotit.vn") // REQUESTSENDVOUCHERBYEMAIL |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherSendMethodAPI.SendVoucherByEmail(context.Background()).XGIAuthorization(xGIAuthorization).REQUESTSENDVOUCHERBYEMAIL(rEQUESTSENDVOUCHERBYEMAIL).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherSendMethodAPI.SendVoucherByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendVoucherByEmail`: VOUCHERSENDEMAILRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `VoucherSendMethodAPI.SendVoucherByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendVoucherByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 
 **rEQUESTSENDVOUCHERBYEMAIL** | [**REQUESTSENDVOUCHERBYEMAIL**](REQUESTSENDVOUCHERBYEMAIL.md) |  | 

### Return type

[**VOUCHERSENDEMAILRESPONSE**](VOUCHERSENDEMAILRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendVoucherBySMS

> VOUCHERSENDSMSRESPONSE SendVoucherBySMS(ctx).XGIAuthorization(xGIAuthorization).REQUESTSENDVOUCHERBYSMS(rEQUESTSENDVOUCHERBYSMS).Execute()

Send voucher by sms



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
	rEQUESTSENDVOUCHERBYSMS := *openapiclient.NewREQUESTSENDVOUCHERBYSMS("1gEGcntd", "0394162063") // REQUESTSENDVOUCHERBYSMS | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherSendMethodAPI.SendVoucherBySMS(context.Background()).XGIAuthorization(xGIAuthorization).REQUESTSENDVOUCHERBYSMS(rEQUESTSENDVOUCHERBYSMS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherSendMethodAPI.SendVoucherBySMS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendVoucherBySMS`: VOUCHERSENDSMSRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `VoucherSendMethodAPI.SendVoucherBySMS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendVoucherBySMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 
 **rEQUESTSENDVOUCHERBYSMS** | [**REQUESTSENDVOUCHERBYSMS**](REQUESTSENDVOUCHERBYSMS.md) |  | 

### Return type

[**VOUCHERSENDSMSRESPONSE**](VOUCHERSENDSMSRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendVoucherByZNS

> VOUCHERSENDZNSRESPONSE SendVoucherByZNS(ctx).XGIAuthorization(xGIAuthorization).REQUESTSENDVOUCHERBYZNS(rEQUESTSENDVOUCHERBYZNS).Execute()

Send voucher by zns



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
	rEQUESTSENDVOUCHERBYZNS := *openapiclient.NewREQUESTSENDVOUCHERBYZNS("0394162063", []string{"1234567890123456"}, "transactionID") // REQUESTSENDVOUCHERBYZNS |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoucherSendMethodAPI.SendVoucherByZNS(context.Background()).XGIAuthorization(xGIAuthorization).REQUESTSENDVOUCHERBYZNS(rEQUESTSENDVOUCHERBYZNS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoucherSendMethodAPI.SendVoucherByZNS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendVoucherByZNS`: VOUCHERSENDZNSRESPONSE
	fmt.Fprintf(os.Stdout, "Response from `VoucherSendMethodAPI.SendVoucherByZNS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendVoucherByZNSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGIAuthorization** | **string** | Authorization | 
 **rEQUESTSENDVOUCHERBYZNS** | [**REQUESTSENDVOUCHERBYZNS**](REQUESTSENDVOUCHERBYZNS.md) |  | 

### Return type

[**VOUCHERSENDZNSRESPONSE**](VOUCHERSENDZNSRESPONSE.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

