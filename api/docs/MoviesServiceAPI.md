# \MoviesServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMoviesRecommendations**](MoviesServiceAPI.md#GetMoviesRecommendations) | **Get** /Movies/Recommendations | Gets movie recommendations



## GetMoviesRecommendations

> []RecommendationDto GetMoviesRecommendations(ctx).CategoryLimit(categoryLimit).ItemLimit(itemLimit).UserId(userId).ParentId(parentId).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Execute()

Gets movie recommendations



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
	categoryLimit := int32(56) // int32 | The max number of categories to return (optional)
	itemLimit := int32(56) // int32 | The max number of items to return per category (optional)
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoviesServiceAPI.GetMoviesRecommendations(context.Background()).CategoryLimit(categoryLimit).ItemLimit(itemLimit).UserId(userId).ParentId(parentId).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoviesServiceAPI.GetMoviesRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMoviesRecommendations`: []RecommendationDto
	fmt.Fprintf(os.Stdout, "Response from `MoviesServiceAPI.GetMoviesRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMoviesRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryLimit** | **int32** | The max number of categories to return | 
 **itemLimit** | **int32** | The max number of items to return per category | 
 **userId** | **string** | Optional. Filter by user id, and attach user data | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 

### Return type

[**[]RecommendationDto**](RecommendationDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

