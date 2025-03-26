# \ServerApiEndpointsAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostNotificationSMTPTestByUserid**](ServerApiEndpointsAPI.md#PostNotificationSMTPTestByUserid) | **Post** /Notification/SMTP/Test/{UserID} | Tests SMTP



## PostNotificationSMTPTestByUserid

> PostNotificationSMTPTestByUserid(ctx, userID).Execute()

Tests SMTP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Akimio521/emby-go/api"
)

func main() {
	userID := "userID_example" // string | User Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerApiEndpointsAPI.PostNotificationSMTPTestByUserid(context.Background(), userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerApiEndpointsAPI.PostNotificationSMTPTestByUserid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNotificationSMTPTestByUseridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

