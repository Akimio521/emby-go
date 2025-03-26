# \PlaystateServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsersByUseridPlayeditemsById**](PlaystateServiceAPI.md#DeleteUsersByUseridPlayeditemsById) | **Delete** /Users/{UserId}/PlayedItems/{Id} | Marks an item as unplayed
[**DeleteUsersByUseridPlayingitemsById**](PlaystateServiceAPI.md#DeleteUsersByUseridPlayingitemsById) | **Delete** /Users/{UserId}/PlayingItems/{Id} | Reports that a user has stopped playing an item
[**PostSessionsPlaying**](PlaystateServiceAPI.md#PostSessionsPlaying) | **Post** /Sessions/Playing | Reports playback has started within a session
[**PostSessionsPlayingPing**](PlaystateServiceAPI.md#PostSessionsPlayingPing) | **Post** /Sessions/Playing/Ping | Pings a playback session
[**PostSessionsPlayingProgress**](PlaystateServiceAPI.md#PostSessionsPlayingProgress) | **Post** /Sessions/Playing/Progress | Reports playback progress within a session
[**PostSessionsPlayingStopped**](PlaystateServiceAPI.md#PostSessionsPlayingStopped) | **Post** /Sessions/Playing/Stopped | Reports playback has stopped within a session
[**PostUsersByUseridPlayeditemsById**](PlaystateServiceAPI.md#PostUsersByUseridPlayeditemsById) | **Post** /Users/{UserId}/PlayedItems/{Id} | Marks an item as played
[**PostUsersByUseridPlayingitemsById**](PlaystateServiceAPI.md#PostUsersByUseridPlayingitemsById) | **Post** /Users/{UserId}/PlayingItems/{Id} | Reports that a user has begun playing an item
[**PostUsersByUseridPlayingitemsByIdProgress**](PlaystateServiceAPI.md#PostUsersByUseridPlayingitemsByIdProgress) | **Post** /Users/{UserId}/PlayingItems/{Id}/Progress | Reports a user&#39;s playback progress



## DeleteUsersByUseridPlayeditemsById

> UserItemDataDto DeleteUsersByUseridPlayeditemsById(ctx, userId, id).Execute()

Marks an item as unplayed



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
	resp, r, err := apiClient.PlaystateServiceAPI.DeleteUsersByUseridPlayeditemsById(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.DeleteUsersByUseridPlayeditemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUsersByUseridPlayeditemsById`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `PlaystateServiceAPI.DeleteUsersByUseridPlayeditemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersByUseridPlayeditemsByIdRequest struct via the builder pattern


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


## DeleteUsersByUseridPlayingitemsById

> DeleteUsersByUseridPlayingitemsById(ctx, userId, id).MediaSourceId(mediaSourceId).NextMediaType(nextMediaType).PositionTicks(positionTicks).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()

Reports that a user has stopped playing an item



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
	mediaSourceId := "mediaSourceId_example" // string | The id of the MediaSource
	nextMediaType := "nextMediaType_example" // string | The next media type that will play
	positionTicks := int64(789) // int64 | Optional. The position, in ticks, where playback stopped. 1 tick = 10000 ms (optional)
	liveStreamId := "liveStreamId_example" // string |  (optional)
	playSessionId := "playSessionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.DeleteUsersByUseridPlayingitemsById(context.Background(), userId, id).MediaSourceId(mediaSourceId).NextMediaType(nextMediaType).PositionTicks(positionTicks).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.DeleteUsersByUseridPlayingitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersByUseridPlayingitemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **string** | The id of the MediaSource | 
 **nextMediaType** | **string** | The next media type that will play | 
 **positionTicks** | **int64** | Optional. The position, in ticks, where playback stopped. 1 tick &#x3D; 10000 ms | 
 **liveStreamId** | **string** |  | 
 **playSessionId** | **string** |  | 

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


## PostSessionsPlaying

> PostSessionsPlaying(ctx).PlaybackStartInfo(playbackStartInfo).Execute()

Reports playback has started within a session



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
	playbackStartInfo := *openapiclient.NewPlaybackStartInfo() // PlaybackStartInfo | PlaybackStartInfo: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostSessionsPlaying(context.Background()).PlaybackStartInfo(playbackStartInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostSessionsPlaying``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsPlayingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackStartInfo** | [**PlaybackStartInfo**](PlaybackStartInfo.md) | PlaybackStartInfo:  | 

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


## PostSessionsPlayingPing

> PostSessionsPlayingPing(ctx).PlaySessionId(playSessionId).Execute()

Pings a playback session



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
	playSessionId := "playSessionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostSessionsPlayingPing(context.Background()).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostSessionsPlayingPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsPlayingPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playSessionId** | **string** |  | 

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


## PostSessionsPlayingProgress

> PostSessionsPlayingProgress(ctx).PlaybackProgressInfo(playbackProgressInfo).Execute()

Reports playback progress within a session



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
	playbackProgressInfo := *openapiclient.NewPlaybackProgressInfo() // PlaybackProgressInfo | PlaybackProgressInfo: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostSessionsPlayingProgress(context.Background()).PlaybackProgressInfo(playbackProgressInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostSessionsPlayingProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsPlayingProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackProgressInfo** | [**PlaybackProgressInfo**](PlaybackProgressInfo.md) | PlaybackProgressInfo:  | 

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


## PostSessionsPlayingStopped

> PostSessionsPlayingStopped(ctx).PlaybackStopInfo(playbackStopInfo).Execute()

Reports playback has stopped within a session



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
	playbackStopInfo := *openapiclient.NewPlaybackStopInfo() // PlaybackStopInfo | PlaybackStopInfo: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostSessionsPlayingStopped(context.Background()).PlaybackStopInfo(playbackStopInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostSessionsPlayingStopped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsPlayingStoppedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackStopInfo** | [**PlaybackStopInfo**](PlaybackStopInfo.md) | PlaybackStopInfo:  | 

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


## PostUsersByUseridPlayeditemsById

> UserItemDataDto PostUsersByUseridPlayeditemsById(ctx, userId, id).DatePlayed(datePlayed).Execute()

Marks an item as played



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
	datePlayed := "datePlayed_example" // string | The date the item was played (if any). Format = yyyyMMddHHmmss (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaystateServiceAPI.PostUsersByUseridPlayeditemsById(context.Background(), userId, id).DatePlayed(datePlayed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostUsersByUseridPlayeditemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersByUseridPlayeditemsById`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `PlaystateServiceAPI.PostUsersByUseridPlayeditemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridPlayeditemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datePlayed** | **string** | The date the item was played (if any). Format &#x3D; yyyyMMddHHmmss | 

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


## PostUsersByUseridPlayingitemsById

> PostUsersByUseridPlayingitemsById(ctx, userId, id).MediaSourceId(mediaSourceId).CanSeek(canSeek).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()

Reports that a user has begun playing an item



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
	mediaSourceId := "mediaSourceId_example" // string | The id of the MediaSource
	canSeek := true // bool | Indicates if the client can seek (optional)
	audioStreamIndex := int32(56) // int32 |  (optional)
	subtitleStreamIndex := int32(56) // int32 |  (optional)
	playMethod := "playMethod_example" // string |  (optional)
	liveStreamId := "liveStreamId_example" // string |  (optional)
	playSessionId := "playSessionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostUsersByUseridPlayingitemsById(context.Background(), userId, id).MediaSourceId(mediaSourceId).CanSeek(canSeek).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostUsersByUseridPlayingitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridPlayingitemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **string** | The id of the MediaSource | 
 **canSeek** | **bool** | Indicates if the client can seek | 
 **audioStreamIndex** | **int32** |  | 
 **subtitleStreamIndex** | **int32** |  | 
 **playMethod** | **string** |  | 
 **liveStreamId** | **string** |  | 
 **playSessionId** | **string** |  | 

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


## PostUsersByUseridPlayingitemsByIdProgress

> PostUsersByUseridPlayingitemsByIdProgress(ctx, userId, id).MediaSourceId(mediaSourceId).PositionTicks(positionTicks).IsPaused(isPaused).IsMuted(isMuted).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).VolumeLevel(volumeLevel).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).RepeatMode(repeatMode).Execute()

Reports a user's playback progress



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
	mediaSourceId := "mediaSourceId_example" // string | The id of the MediaSource
	positionTicks := int64(789) // int64 | Optional. The current position, in ticks. 1 tick = 10000 ms (optional)
	isPaused := true // bool | Indicates if the player is paused. (optional)
	isMuted := true // bool | Indicates if the player is muted. (optional)
	audioStreamIndex := int32(56) // int32 |  (optional)
	subtitleStreamIndex := int32(56) // int32 |  (optional)
	volumeLevel := int32(56) // int32 | Scale of 0-100 (optional)
	playMethod := "playMethod_example" // string |  (optional)
	liveStreamId := "liveStreamId_example" // string |  (optional)
	playSessionId := "playSessionId_example" // string |  (optional)
	repeatMode := "repeatMode_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateServiceAPI.PostUsersByUseridPlayingitemsByIdProgress(context.Background(), userId, id).MediaSourceId(mediaSourceId).PositionTicks(positionTicks).IsPaused(isPaused).IsMuted(isMuted).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).VolumeLevel(volumeLevel).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).RepeatMode(repeatMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateServiceAPI.PostUsersByUseridPlayingitemsByIdProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridPlayingitemsByIdProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **string** | The id of the MediaSource | 
 **positionTicks** | **int64** | Optional. The current position, in ticks. 1 tick &#x3D; 10000 ms | 
 **isPaused** | **bool** | Indicates if the player is paused. | 
 **isMuted** | **bool** | Indicates if the player is muted. | 
 **audioStreamIndex** | **int32** |  | 
 **subtitleStreamIndex** | **int32** |  | 
 **volumeLevel** | **int32** | Scale of 0-100 | 
 **playMethod** | **string** |  | 
 **liveStreamId** | **string** |  | 
 **playSessionId** | **string** |  | 
 **repeatMode** | **string** |  | 

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

