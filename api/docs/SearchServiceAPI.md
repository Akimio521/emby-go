# \SearchServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearchHints**](SearchServiceAPI.md#GetSearchHints) | **Get** /Search/Hints | Gets search hints based on a search term



## GetSearchHints

> SearchSearchHintResult GetSearchHints(ctx).SearchTerm(searchTerm).StartIndex(startIndex).Limit(limit).UserId(userId).IncludePeople(includePeople).IncludeMedia(includeMedia).IncludeGenres(includeGenres).IncludeStudios(includeStudios).IncludeArtists(includeArtists).IncludeItemTypes(includeItemTypes).ExcludeItemTypes(excludeItemTypes).MediaTypes(mediaTypes).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).Execute()

Gets search hints based on a search term



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
	searchTerm := "searchTerm_example" // string | The search term to filter on
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	userId := "userId_example" // string | Optional. Supply a user id to search within a user's library or omit to search all. (optional)
	includePeople := true // bool |  (optional)
	includeMedia := true // bool |  (optional)
	includeGenres := true // bool |  (optional)
	includeStudios := true // bool |  (optional)
	includeArtists := true // bool |  (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	mediaTypes := "mediaTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for movies. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchServiceAPI.GetSearchHints(context.Background()).SearchTerm(searchTerm).StartIndex(startIndex).Limit(limit).UserId(userId).IncludePeople(includePeople).IncludeMedia(includeMedia).IncludeGenres(includeGenres).IncludeStudios(includeStudios).IncludeArtists(includeArtists).IncludeItemTypes(includeItemTypes).ExcludeItemTypes(excludeItemTypes).MediaTypes(mediaTypes).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchServiceAPI.GetSearchHints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchHints`: SearchSearchHintResult
	fmt.Fprintf(os.Stdout, "Response from `SearchServiceAPI.GetSearchHints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchHintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** | The search term to filter on | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **userId** | **string** | Optional. Supply a user id to search within a user&#39;s library or omit to search all. | 
 **includePeople** | **bool** |  | 
 **includeMedia** | **bool** |  | 
 **includeGenres** | **bool** |  | 
 **includeStudios** | **bool** |  | 
 **includeArtists** | **bool** |  | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **mediaTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for movies. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 

### Return type

[**SearchSearchHintResult**](SearchSearchHintResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

