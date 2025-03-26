# \SubtitleServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVideosByIdSubtitlesByIndex**](SubtitleServiceAPI.md#DeleteVideosByIdSubtitlesByIndex) | **Delete** /Videos/{Id}/Subtitles/{Index} | Deletes an external subtitle file
[**GetItemsByIdRemotesearchSubtitlesByLanguage**](SubtitleServiceAPI.md#GetItemsByIdRemotesearchSubtitlesByLanguage) | **Get** /Items/{Id}/RemoteSearch/Subtitles/{Language} | 
[**GetProvidersSubtitlesSubtitlesById**](SubtitleServiceAPI.md#GetProvidersSubtitlesSubtitlesById) | **Get** /Providers/Subtitles/Subtitles/{Id} | 
[**GetVideosByIdByMediasourceidSubtitlesByIndexByFormat**](SubtitleServiceAPI.md#GetVideosByIdByMediasourceidSubtitlesByIndexByFormat) | **Get** /Videos/{Id}/{MediaSourceId}/Subtitles/{Index}/Stream.{Format} | Gets subtitles in a specified format.
[**GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksByFormat**](SubtitleServiceAPI.md#GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksByFormat) | **Get** /Videos/{Id}/{MediaSourceId}/Subtitles/{Index}/{StartPositionTicks}/Stream.{Format} | Gets subtitles in a specified format.
[**PostItemsByIdRemotesearchSubtitlesBySubtitleid**](SubtitleServiceAPI.md#PostItemsByIdRemotesearchSubtitlesBySubtitleid) | **Post** /Items/{Id}/RemoteSearch/Subtitles/{SubtitleId} | 



## DeleteVideosByIdSubtitlesByIndex

> DeleteVideosByIdSubtitlesByIndex(ctx, id, index).Execute()

Deletes an external subtitle file



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
	index := int32(56) // int32 | The subtitle stream index

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.DeleteVideosByIdSubtitlesByIndex(context.Background(), id, index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.DeleteVideosByIdSubtitlesByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**index** | **int32** | The subtitle stream index | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideosByIdSubtitlesByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdRemotesearchSubtitlesByLanguage

> []RemoteSubtitleInfo GetItemsByIdRemotesearchSubtitlesByLanguage(ctx, id, language).IsPerfectMatch(isPerfectMatch).IsForced(isForced).Execute()





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
	language := "language_example" // string | Language
	isPerfectMatch := true // bool | IsPerfectMatch (optional)
	isForced := true // bool | IsForced (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubtitleServiceAPI.GetItemsByIdRemotesearchSubtitlesByLanguage(context.Background(), id, language).IsPerfectMatch(isPerfectMatch).IsForced(isForced).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.GetItemsByIdRemotesearchSubtitlesByLanguage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdRemotesearchSubtitlesByLanguage`: []RemoteSubtitleInfo
	fmt.Fprintf(os.Stdout, "Response from `SubtitleServiceAPI.GetItemsByIdRemotesearchSubtitlesByLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**language** | **string** | Language | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdRemotesearchSubtitlesByLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isPerfectMatch** | **bool** | IsPerfectMatch | 
 **isForced** | **bool** | IsForced | 

### Return type

[**[]RemoteSubtitleInfo**](RemoteSubtitleInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvidersSubtitlesSubtitlesById

> GetProvidersSubtitlesSubtitlesById(ctx, id).Execute()





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
	r, err := apiClient.SubtitleServiceAPI.GetProvidersSubtitlesSubtitlesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.GetProvidersSubtitlesSubtitlesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvidersSubtitlesSubtitlesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideosByIdByMediasourceidSubtitlesByIndexByFormat

> GetVideosByIdByMediasourceidSubtitlesByIndexByFormat(ctx, id, mediaSourceId, index, format).StartPositionTicks(startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()

Gets subtitles in a specified format.



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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	index := int32(56) // int32 | The subtitle stream index
	format := "format_example" // string | Format
	startPositionTicks := int64(789) // int64 | StartPositionTicks (optional)
	endPositionTicks := int64(789) // int64 | EndPositionTicks (optional)
	copyTimestamps := true // bool | CopyTimestamps (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.GetVideosByIdByMediasourceidSubtitlesByIndexByFormat(context.Background(), id, mediaSourceId, index, format).StartPositionTicks(startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.GetVideosByIdByMediasourceidSubtitlesByIndexByFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**mediaSourceId** | **string** | MediaSourceId | 
**index** | **int32** | The subtitle stream index | 
**format** | **string** | Format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosByIdByMediasourceidSubtitlesByIndexByFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **startPositionTicks** | **int64** | StartPositionTicks | 
 **endPositionTicks** | **int64** | EndPositionTicks | 
 **copyTimestamps** | **bool** | CopyTimestamps | 

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


## GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksByFormat

> GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksByFormat(ctx, id, mediaSourceId, index, format, startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()

Gets subtitles in a specified format.



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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	index := int32(56) // int32 | The subtitle stream index
	format := "format_example" // string | Format
	startPositionTicks := int64(789) // int64 | StartPositionTicks
	endPositionTicks := int64(789) // int64 | EndPositionTicks (optional)
	copyTimestamps := true // bool | CopyTimestamps (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksByFormat(context.Background(), id, mediaSourceId, index, format, startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksByFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**mediaSourceId** | **string** | MediaSourceId | 
**index** | **int32** | The subtitle stream index | 
**format** | **string** | Format | 
**startPositionTicks** | **int64** | StartPositionTicks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksByFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **endPositionTicks** | **int64** | EndPositionTicks | 
 **copyTimestamps** | **bool** | CopyTimestamps | 

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


## PostItemsByIdRemotesearchSubtitlesBySubtitleid

> PostItemsByIdRemotesearchSubtitlesBySubtitleid(ctx, id, subtitleId).Execute()





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
	subtitleId := "subtitleId_example" // string | SubtitleId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.PostItemsByIdRemotesearchSubtitlesBySubtitleid(context.Background(), id, subtitleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.PostItemsByIdRemotesearchSubtitlesBySubtitleid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**subtitleId** | **string** | SubtitleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsByIdRemotesearchSubtitlesBySubtitleidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

