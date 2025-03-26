# \VideoServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemsFile**](VideoServiceAPI.md#GetItemsFile) | **Get** /Items/File | Gets the original file of an item
[**GetVideosByIdByContainer**](VideoServiceAPI.md#GetVideosByIdByContainer) | **Get** /Videos/{Id}/stream.{Container} | Gets a video stream
[**GetVideosByIdStream**](VideoServiceAPI.md#GetVideosByIdStream) | **Get** /Videos/{Id}/stream | Gets a video stream
[**HeadVideosByIdByContainer**](VideoServiceAPI.md#HeadVideosByIdByContainer) | **Head** /Videos/{Id}/stream.{Container} | Gets a video stream
[**HeadVideosByIdStream**](VideoServiceAPI.md#HeadVideosByIdStream) | **Head** /Videos/{Id}/stream | Gets a video stream



## GetItemsFile

> GetItemsFile(ctx).Execute()

Gets the original file of an item



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoServiceAPI.GetItemsFile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoServiceAPI.GetItemsFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsFileRequest struct via the builder pattern


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


## GetVideosByIdByContainer

> GetVideosByIdByContainer(ctx, id, container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()

Gets a video stream



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
	container := "container_example" // string | Container
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. Options: aac, mp3, vorbis, wma. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100 (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	audioChannels := int32(56) // int32 | Optional. Specify a specific number of audio channels to encode to, e.g. 2 (optional)
	maxAudioChannels := int32(56) // int32 | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 (optional)
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false (optional)
	profile := "profile_example" // string | Optional. Specify a specific h264 profile, e.g. main, baseline, high. (optional)
	level := "level_example" // string | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. (optional)
	framerate := float32(3.4) // float32 | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	maxFramerate := float32(3.4) // float32 | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	copyTimestamps := true // bool | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms (optional)
	width := int32(56) // int32 | Optional. The fixed horizontal resolution of the encoded video. (optional)
	height := int32(56) // int32 | Optional. The fixed vertical resolution of the encoded video. (optional)
	maxWidth := int32(56) // int32 | Optional. The maximum horizontal resolution of the encoded video. (optional)
	maxHeight := int32(56) // int32 | Optional. The maximum vertical resolution of the encoded video. (optional)
	videoBitRate := int32(56) // int32 | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. (optional)
	subtitleStreamIndex := int32(56) // int32 | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. (optional)
	subtitleMethod := "subtitleMethod_example" // string | Optional. Specify the subtitle delivery method. (optional)
	maxRefFrames := int32(56) // int32 | Optional. (optional)
	maxVideoBitDepth := int32(56) // int32 | Optional. (optional)
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. Options: h264, mpeg4, theora, vpx, wmv. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoServiceAPI.GetVideosByIdByContainer(context.Background(), id, container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoServiceAPI.GetVideosByIdByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**container** | **string** | Container | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosByIdByContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100 | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **int32** | Optional. Specify a specific number of audio channels to encode to, e.g. 2 | 
 **maxAudioChannels** | **int32** | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false | 
 **profile** | **string** | Optional. Specify a specific h264 profile, e.g. main, baseline, high. | 
 **level** | **string** | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. | 
 **framerate** | **float32** | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **float32** | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **bool** | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms | 
 **width** | **int32** | Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **int32** | Optional. The fixed vertical resolution of the encoded video. | 
 **maxWidth** | **int32** | Optional. The maximum horizontal resolution of the encoded video. | 
 **maxHeight** | **int32** | Optional. The maximum vertical resolution of the encoded video. | 
 **videoBitRate** | **int32** | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **int32** | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | **string** | Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **int32** | Optional. | 
 **maxVideoBitDepth** | **int32** | Optional. | 
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h264, mpeg4, theora, vpx, wmv. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 

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


## GetVideosByIdStream

> GetVideosByIdStream(ctx, id).Container(container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()

Gets a video stream



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
	container := "container_example" // string | Container
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. Options: aac, mp3, vorbis, wma. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100 (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	audioChannels := int32(56) // int32 | Optional. Specify a specific number of audio channels to encode to, e.g. 2 (optional)
	maxAudioChannels := int32(56) // int32 | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 (optional)
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false (optional)
	profile := "profile_example" // string | Optional. Specify a specific h264 profile, e.g. main, baseline, high. (optional)
	level := "level_example" // string | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. (optional)
	framerate := float32(3.4) // float32 | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	maxFramerate := float32(3.4) // float32 | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	copyTimestamps := true // bool | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms (optional)
	width := int32(56) // int32 | Optional. The fixed horizontal resolution of the encoded video. (optional)
	height := int32(56) // int32 | Optional. The fixed vertical resolution of the encoded video. (optional)
	maxWidth := int32(56) // int32 | Optional. The maximum horizontal resolution of the encoded video. (optional)
	maxHeight := int32(56) // int32 | Optional. The maximum vertical resolution of the encoded video. (optional)
	videoBitRate := int32(56) // int32 | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. (optional)
	subtitleStreamIndex := int32(56) // int32 | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. (optional)
	subtitleMethod := "subtitleMethod_example" // string | Optional. Specify the subtitle delivery method. (optional)
	maxRefFrames := int32(56) // int32 | Optional. (optional)
	maxVideoBitDepth := int32(56) // int32 | Optional. (optional)
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. Options: h264, mpeg4, theora, vpx, wmv. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoServiceAPI.GetVideosByIdStream(context.Background(), id).Container(container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoServiceAPI.GetVideosByIdStream``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetVideosByIdStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **string** | Container | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100 | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **int32** | Optional. Specify a specific number of audio channels to encode to, e.g. 2 | 
 **maxAudioChannels** | **int32** | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false | 
 **profile** | **string** | Optional. Specify a specific h264 profile, e.g. main, baseline, high. | 
 **level** | **string** | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. | 
 **framerate** | **float32** | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **float32** | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **bool** | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms | 
 **width** | **int32** | Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **int32** | Optional. The fixed vertical resolution of the encoded video. | 
 **maxWidth** | **int32** | Optional. The maximum horizontal resolution of the encoded video. | 
 **maxHeight** | **int32** | Optional. The maximum vertical resolution of the encoded video. | 
 **videoBitRate** | **int32** | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **int32** | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | **string** | Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **int32** | Optional. | 
 **maxVideoBitDepth** | **int32** | Optional. | 
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h264, mpeg4, theora, vpx, wmv. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 

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


## HeadVideosByIdByContainer

> HeadVideosByIdByContainer(ctx, id, container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()

Gets a video stream



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
	container := "container_example" // string | Container
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. Options: aac, mp3, vorbis, wma. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100 (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	audioChannels := int32(56) // int32 | Optional. Specify a specific number of audio channels to encode to, e.g. 2 (optional)
	maxAudioChannels := int32(56) // int32 | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 (optional)
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false (optional)
	profile := "profile_example" // string | Optional. Specify a specific h264 profile, e.g. main, baseline, high. (optional)
	level := "level_example" // string | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. (optional)
	framerate := float32(3.4) // float32 | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	maxFramerate := float32(3.4) // float32 | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	copyTimestamps := true // bool | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms (optional)
	width := int32(56) // int32 | Optional. The fixed horizontal resolution of the encoded video. (optional)
	height := int32(56) // int32 | Optional. The fixed vertical resolution of the encoded video. (optional)
	maxWidth := int32(56) // int32 | Optional. The maximum horizontal resolution of the encoded video. (optional)
	maxHeight := int32(56) // int32 | Optional. The maximum vertical resolution of the encoded video. (optional)
	videoBitRate := int32(56) // int32 | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. (optional)
	subtitleStreamIndex := int32(56) // int32 | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. (optional)
	subtitleMethod := "subtitleMethod_example" // string | Optional. Specify the subtitle delivery method. (optional)
	maxRefFrames := int32(56) // int32 | Optional. (optional)
	maxVideoBitDepth := int32(56) // int32 | Optional. (optional)
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. Options: h264, mpeg4, theora, vpx, wmv. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoServiceAPI.HeadVideosByIdByContainer(context.Background(), id, container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoServiceAPI.HeadVideosByIdByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**container** | **string** | Container | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadVideosByIdByContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100 | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **int32** | Optional. Specify a specific number of audio channels to encode to, e.g. 2 | 
 **maxAudioChannels** | **int32** | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false | 
 **profile** | **string** | Optional. Specify a specific h264 profile, e.g. main, baseline, high. | 
 **level** | **string** | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. | 
 **framerate** | **float32** | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **float32** | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **bool** | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms | 
 **width** | **int32** | Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **int32** | Optional. The fixed vertical resolution of the encoded video. | 
 **maxWidth** | **int32** | Optional. The maximum horizontal resolution of the encoded video. | 
 **maxHeight** | **int32** | Optional. The maximum vertical resolution of the encoded video. | 
 **videoBitRate** | **int32** | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **int32** | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | **string** | Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **int32** | Optional. | 
 **maxVideoBitDepth** | **int32** | Optional. | 
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h264, mpeg4, theora, vpx, wmv. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 

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


## HeadVideosByIdStream

> HeadVideosByIdStream(ctx, id).Container(container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()

Gets a video stream



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
	container := "container_example" // string | Container
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. Options: aac, mp3, vorbis, wma. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100 (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	audioChannels := int32(56) // int32 | Optional. Specify a specific number of audio channels to encode to, e.g. 2 (optional)
	maxAudioChannels := int32(56) // int32 | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 (optional)
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false (optional)
	profile := "profile_example" // string | Optional. Specify a specific h264 profile, e.g. main, baseline, high. (optional)
	level := "level_example" // string | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. (optional)
	framerate := float32(3.4) // float32 | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	maxFramerate := float32(3.4) // float32 | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	copyTimestamps := true // bool | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms (optional)
	width := int32(56) // int32 | Optional. The fixed horizontal resolution of the encoded video. (optional)
	height := int32(56) // int32 | Optional. The fixed vertical resolution of the encoded video. (optional)
	maxWidth := int32(56) // int32 | Optional. The maximum horizontal resolution of the encoded video. (optional)
	maxHeight := int32(56) // int32 | Optional. The maximum vertical resolution of the encoded video. (optional)
	videoBitRate := int32(56) // int32 | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. (optional)
	subtitleStreamIndex := int32(56) // int32 | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. (optional)
	subtitleMethod := "subtitleMethod_example" // string | Optional. Specify the subtitle delivery method. (optional)
	maxRefFrames := int32(56) // int32 | Optional. (optional)
	maxVideoBitDepth := int32(56) // int32 | Optional. (optional)
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. Options: h264, mpeg4, theora, vpx, wmv. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoServiceAPI.HeadVideosByIdStream(context.Background(), id).Container(container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoServiceAPI.HeadVideosByIdStream``: %v\n", err)
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

Other parameters are passed through a pointer to a apiHeadVideosByIdStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **string** | Container | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100 | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **int32** | Optional. Specify a specific number of audio channels to encode to, e.g. 2 | 
 **maxAudioChannels** | **int32** | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false | 
 **profile** | **string** | Optional. Specify a specific h264 profile, e.g. main, baseline, high. | 
 **level** | **string** | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. | 
 **framerate** | **float32** | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **float32** | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **bool** | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms | 
 **width** | **int32** | Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **int32** | Optional. The fixed vertical resolution of the encoded video. | 
 **maxWidth** | **int32** | Optional. The maximum horizontal resolution of the encoded video. | 
 **maxHeight** | **int32** | Optional. The maximum vertical resolution of the encoded video. | 
 **videoBitRate** | **int32** | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **int32** | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | **string** | Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **int32** | Optional. | 
 **maxVideoBitDepth** | **int32** | Optional. | 
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h264, mpeg4, theora, vpx, wmv. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 

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

