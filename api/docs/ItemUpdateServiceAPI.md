# \ItemUpdateServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemsByItemidMetadataeditor**](ItemUpdateServiceAPI.md#GetItemsByItemidMetadataeditor) | **Get** /Items/{ItemId}/MetadataEditor | Gets metadata editor info for an item
[**PostItemsByItemid**](ItemUpdateServiceAPI.md#PostItemsByItemid) | **Post** /Items/{ItemId} | Updates an item



## GetItemsByItemidMetadataeditor

> MetadataEditorInfo GetItemsByItemidMetadataeditor(ctx, itemId).Execute()

Gets metadata editor info for an item



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
	itemId := "itemId_example" // string | The id of the item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemUpdateServiceAPI.GetItemsByItemidMetadataeditor(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemUpdateServiceAPI.GetItemsByItemidMetadataeditor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByItemidMetadataeditor`: MetadataEditorInfo
	fmt.Fprintf(os.Stdout, "Response from `ItemUpdateServiceAPI.GetItemsByItemidMetadataeditor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The id of the item | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByItemidMetadataeditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataEditorInfo**](MetadataEditorInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsByItemid

> PostItemsByItemid(ctx, itemId).BaseItemDto(baseItemDto).Execute()

Updates an item



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
	itemId := "itemId_example" // string | The id of the item
	baseItemDto := *openapiclient.NewBaseItemDto() // BaseItemDto | BaseItemDto: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemUpdateServiceAPI.PostItemsByItemid(context.Background(), itemId).BaseItemDto(baseItemDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemUpdateServiceAPI.PostItemsByItemid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The id of the item | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsByItemidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseItemDto** | [**BaseItemDto**](BaseItemDto.md) | BaseItemDto:  | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

