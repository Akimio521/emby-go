# \UserLibraryServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsersByUseridFavoriteitemsById**](UserLibraryServiceAPI.md#DeleteUsersByUseridFavoriteitemsById) | **Delete** /Users/{UserId}/FavoriteItems/{Id} | Unmarks an item as a favorite
[**DeleteUsersByUseridItemsByIdRating**](UserLibraryServiceAPI.md#DeleteUsersByUseridItemsByIdRating) | **Delete** /Users/{UserId}/Items/{Id}/Rating | Deletes a user&#39;s saved personal rating for an item
[**GetLivetvProgramsById**](UserLibraryServiceAPI.md#GetLivetvProgramsById) | **Get** /LiveTv/Programs/{Id} | Gets a live tv program
[**GetUsersByUseridItemsById**](UserLibraryServiceAPI.md#GetUsersByUseridItemsById) | **Get** /Users/{UserId}/Items/{Id} | Gets an item from a user&#39;s library
[**GetUsersByUseridItemsByIdIntros**](UserLibraryServiceAPI.md#GetUsersByUseridItemsByIdIntros) | **Get** /Users/{UserId}/Items/{Id}/Intros | Gets intros to play before the main media item plays
[**GetUsersByUseridItemsByIdLocaltrailers**](UserLibraryServiceAPI.md#GetUsersByUseridItemsByIdLocaltrailers) | **Get** /Users/{UserId}/Items/{Id}/LocalTrailers | Gets local trailers for an item
[**GetUsersByUseridItemsByIdSpecialfeatures**](UserLibraryServiceAPI.md#GetUsersByUseridItemsByIdSpecialfeatures) | **Get** /Users/{UserId}/Items/{Id}/SpecialFeatures | Gets special features for an item
[**GetUsersByUseridItemsLatest**](UserLibraryServiceAPI.md#GetUsersByUseridItemsLatest) | **Get** /Users/{UserId}/Items/Latest | Gets latest media
[**GetUsersByUseridItemsRoot**](UserLibraryServiceAPI.md#GetUsersByUseridItemsRoot) | **Get** /Users/{UserId}/Items/Root | Gets the root folder from a user&#39;s library
[**PostUsersByUseridFavoriteitemsById**](UserLibraryServiceAPI.md#PostUsersByUseridFavoriteitemsById) | **Post** /Users/{UserId}/FavoriteItems/{Id} | Marks an item as a favorite
[**PostUsersByUseridItemsByIdRating**](UserLibraryServiceAPI.md#PostUsersByUseridItemsByIdRating) | **Post** /Users/{UserId}/Items/{Id}/Rating | Updates a user&#39;s rating for an item



## DeleteUsersByUseridFavoriteitemsById

> UserItemDataDto DeleteUsersByUseridFavoriteitemsById(ctx, userId, id).Execute()

Unmarks an item as a favorite



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
	userId := "userId_example" // string | User Id
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryServiceAPI.DeleteUsersByUseridFavoriteitemsById(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryServiceAPI.DeleteUsersByUseridFavoriteitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUsersByUseridFavoriteitemsById`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryServiceAPI.DeleteUsersByUseridFavoriteitemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersByUseridFavoriteitemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsersByUseridItemsByIdRating

> UserItemDataDto DeleteUsersByUseridItemsByIdRating(ctx, userId, id).Execute()

Deletes a user's saved personal rating for an item



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
	userId := "userId_example" // string | User Id
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryServiceAPI.DeleteUsersByUseridItemsByIdRating(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryServiceAPI.DeleteUsersByUseridItemsByIdRating``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUsersByUseridItemsByIdRating`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryServiceAPI.DeleteUsersByUseridItemsByIdRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersByUseridItemsByIdRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvProgramsById

> BaseItemDto GetLivetvProgramsById(ctx, id).Execute()

Gets a live tv program



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
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryServiceAPI.GetLivetvProgramsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryServiceAPI.GetLivetvProgramsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvProgramsById`: BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryServiceAPI.GetLivetvProgramsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvProgramsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersByUseridItemsById

> BaseItemDto GetUsersByUseridItemsById(ctx, userId, id).Execute()

Gets an item from a user's library



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
	userId := "userId_example" // string | User Id
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryServiceAPI.GetUsersByUseridItemsById(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryServiceAPI.GetUsersByUseridItemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersByUseridItemsById`: BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryServiceAPI.GetUsersByUseridItemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByUseridItemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersByUseridItemsByIdIntros

> QueryResultBaseItemDto GetUsersByUseridItemsByIdIntros(ctx, userId, id).Execute()

Gets intros to play before the main media item plays



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
	userId := "userId_example" // string | User Id
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryServiceAPI.GetUsersByUseridItemsByIdIntros(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryServiceAPI.GetUsersByUseridItemsByIdIntros``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersByUseridItemsByIdIntros`: QueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryServiceAPI.GetUsersByUseridItemsByIdIntros`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByUseridItemsByIdIntrosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**QueryResultBaseItemDto**](QueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersByUseridItemsByIdLocaltrailers

> []BaseItemDto GetUsersByUseridItemsByIdLocaltrailers(ctx, userId, id).Execute()

Gets local trailers for an item



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
	userId := "userId_example" // string | User Id
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryServiceAPI.GetUsersByUseridItemsByIdLocaltrailers(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryServiceAPI.GetUsersByUseridItemsByIdLocaltrailers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersByUseridItemsByIdLocaltrailers`: []BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryServiceAPI.GetUsersByUseridItemsByIdLocaltrailers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByUseridItemsByIdLocaltrailersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersByUseridItemsByIdSpecialfeatures

> []BaseItemDto GetUsersByUseridItemsByIdSpecialfeatures(ctx, userId, id).Execute()

Gets special features for an item



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
	userId := "userId_example" // string | User Id
	id := "id_example" // string | Movie Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryServiceAPI.GetUsersByUseridItemsByIdSpecialfeatures(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryServiceAPI.GetUsersByUseridItemsByIdSpecialfeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersByUseridItemsByIdSpecialfeatures`: []BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryServiceAPI.GetUsersByUseridItemsByIdSpecialfeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Movie Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByUseridItemsByIdSpecialfeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersByUseridItemsLatest

> []BaseItemDto GetUsersByUseridItemsLatest(ctx, userId).Limit(limit).ParentId(parentId).Fields(fields).IncludeItemTypes(includeItemTypes).IsFolder(isFolder).IsPlayed(isPlayed).GroupItems(groupItems).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()

Gets latest media



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
	userId := "userId_example" // string | User Id
	limit := int32(56) // int32 | Limit (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, SortName, Studios, Taglines (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	isFolder := true // bool | Filter by items that are folders, or not. (optional)
	isPlayed := true // bool | Filter by items that are played, or not. (optional)
	groupItems := true // bool | Whether or not to group items into a parent container. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional, include user data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryServiceAPI.GetUsersByUseridItemsLatest(context.Background(), userId).Limit(limit).ParentId(parentId).Fields(fields).IncludeItemTypes(includeItemTypes).IsFolder(isFolder).IsPlayed(isPlayed).GroupItems(groupItems).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryServiceAPI.GetUsersByUseridItemsLatest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersByUseridItemsLatest`: []BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryServiceAPI.GetUsersByUseridItemsLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByUseridItemsLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, SortName, Studios, Taglines | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **isFolder** | **bool** | Filter by items that are folders, or not. | 
 **isPlayed** | **bool** | Filter by items that are played, or not. | 
 **groupItems** | **bool** | Whether or not to group items into a parent container. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional, include user data | 

### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersByUseridItemsRoot

> BaseItemDto GetUsersByUseridItemsRoot(ctx, userId).Execute()

Gets the root folder from a user's library



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
	userId := "userId_example" // string | User Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryServiceAPI.GetUsersByUseridItemsRoot(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryServiceAPI.GetUsersByUseridItemsRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersByUseridItemsRoot`: BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryServiceAPI.GetUsersByUseridItemsRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByUseridItemsRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersByUseridFavoriteitemsById

> UserItemDataDto PostUsersByUseridFavoriteitemsById(ctx, userId, id).Execute()

Marks an item as a favorite



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
	userId := "userId_example" // string | User Id
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryServiceAPI.PostUsersByUseridFavoriteitemsById(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryServiceAPI.PostUsersByUseridFavoriteitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersByUseridFavoriteitemsById`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryServiceAPI.PostUsersByUseridFavoriteitemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridFavoriteitemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersByUseridItemsByIdRating

> UserItemDataDto PostUsersByUseridItemsByIdRating(ctx, userId, id).Likes(likes).Execute()

Updates a user's rating for an item



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
	userId := "userId_example" // string | User Id
	id := "id_example" // string | Item Id
	likes := true // bool | Whether the user likes the item or not. true/false

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryServiceAPI.PostUsersByUseridItemsByIdRating(context.Background(), userId, id).Likes(likes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryServiceAPI.PostUsersByUseridItemsByIdRating``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersByUseridItemsByIdRating`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryServiceAPI.PostUsersByUseridItemsByIdRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridItemsByIdRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **likes** | **bool** | Whether the user likes the item or not. true/false | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

