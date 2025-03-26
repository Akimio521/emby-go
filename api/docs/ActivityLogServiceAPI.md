# \ActivityLogServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemActivitylogEntries**](ActivityLogServiceAPI.md#GetSystemActivitylogEntries) | **Get** /System/ActivityLog/Entries | Gets activity log entries



## GetSystemActivitylogEntries

> QueryResultActivityLogEntry GetSystemActivitylogEntries(ctx).StartIndex(startIndex).Limit(limit).MinDate(minDate).Execute()

Gets activity log entries



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
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	minDate := "minDate_example" // string | Optional. The minimum date. Format = ISO (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityLogServiceAPI.GetSystemActivitylogEntries(context.Background()).StartIndex(startIndex).Limit(limit).MinDate(minDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityLogServiceAPI.GetSystemActivitylogEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemActivitylogEntries`: QueryResultActivityLogEntry
	fmt.Fprintf(os.Stdout, "Response from `ActivityLogServiceAPI.GetSystemActivitylogEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemActivitylogEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **minDate** | **string** | Optional. The minimum date. Format &#x3D; ISO | 

### Return type

[**QueryResultActivityLogEntry**](QueryResultActivityLogEntry.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

