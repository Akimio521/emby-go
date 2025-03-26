# \GamesServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGamesSystemsummaries**](GamesServiceAPI.md#GetGamesSystemsummaries) | **Get** /Games/SystemSummaries | Finds games similar to a given game.



## GetGamesSystemsummaries

> []GameSystemSummary GetGamesSystemsummaries(ctx).UserId(userId).Execute()

Finds games similar to a given game.



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
	userId := "userId_example" // string | Optional. Filter by user id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesServiceAPI.GetGamesSystemsummaries(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesServiceAPI.GetGamesSystemsummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGamesSystemsummaries`: []GameSystemSummary
	fmt.Fprintf(os.Stdout, "Response from `GamesServiceAPI.GetGamesSystemsummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGamesSystemsummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional. Filter by user id | 

### Return type

[**[]GameSystemSummary**](GameSystemSummary.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

