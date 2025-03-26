# \FilterServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemsFilters**](FilterServiceAPI.md#GetItemsFilters) | **Get** /Items/Filters | Gets branding configuration
[**GetItemsFilters2**](FilterServiceAPI.md#GetItemsFilters2) | **Get** /Items/Filters2 | Gets branding configuration



## GetItemsFilters

> QueryFiltersLegacy GetItemsFilters(ctx).UserId(userId).ParentId(parentId).IncludeItemTypes(includeItemTypes).MediaTypes(mediaTypes).Execute()

Gets branding configuration



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
	userId := "userId_example" // string | User Id (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterServiceAPI.GetItemsFilters(context.Background()).UserId(userId).ParentId(parentId).IncludeItemTypes(includeItemTypes).MediaTypes(mediaTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterServiceAPI.GetItemsFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsFilters`: QueryFiltersLegacy
	fmt.Fprintf(os.Stdout, "Response from `FilterServiceAPI.GetItemsFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User Id | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 

### Return type

[**QueryFiltersLegacy**](QueryFiltersLegacy.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsFilters2

> QueryFilters GetItemsFilters2(ctx).UserId(userId).ParentId(parentId).IncludeItemTypes(includeItemTypes).MediaTypes(mediaTypes).Execute()

Gets branding configuration



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
	userId := "userId_example" // string | User Id (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterServiceAPI.GetItemsFilters2(context.Background()).UserId(userId).ParentId(parentId).IncludeItemTypes(includeItemTypes).MediaTypes(mediaTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterServiceAPI.GetItemsFilters2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsFilters2`: QueryFilters
	fmt.Fprintf(os.Stdout, "Response from `FilterServiceAPI.GetItemsFilters2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsFilters2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User Id | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 

### Return type

[**QueryFilters**](QueryFilters.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

