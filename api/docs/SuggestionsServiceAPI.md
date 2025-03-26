# \SuggestionsServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsersByUseridSuggestions**](SuggestionsServiceAPI.md#GetUsersByUseridSuggestions) | **Get** /Users/{UserId}/Suggestions | Gets items based on a query.



## GetUsersByUseridSuggestions

> QueryResultBaseItemDto GetUsersByUseridSuggestions(ctx, userId).Execute()

Gets items based on a query.



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuggestionsServiceAPI.GetUsersByUseridSuggestions(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsServiceAPI.GetUsersByUseridSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersByUseridSuggestions`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `SuggestionsServiceAPI.GetUsersByUseridSuggestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByUseridSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

