# \UniversalAudioServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAudioByIdByContainer**](UniversalAudioServiceAPI.md#GetAudioByIdByContainer) | **Get** /Audio/{Id}/universal.{Container} | Gets an audio stream
[**GetAudioByIdUniversal**](UniversalAudioServiceAPI.md#GetAudioByIdUniversal) | **Get** /Audio/{Id}/universal | Gets an audio stream
[**HeadAudioByIdByContainer**](UniversalAudioServiceAPI.md#HeadAudioByIdByContainer) | **Head** /Audio/{Id}/universal.{Container} | Gets an audio stream
[**HeadAudioByIdUniversal**](UniversalAudioServiceAPI.md#HeadAudioByIdUniversal) | **Head** /Audio/{Id}/universal | Gets an audio stream



## GetAudioByIdByContainer

> GetAudioByIdByContainer(ctx, id, container).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()

Gets an audio stream



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
	container := "container_example" // string | 
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniversalAudioServiceAPI.GetAudioByIdByContainer(context.Background(), id, container).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversalAudioServiceAPI.GetAudioByIdByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**container** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioByIdByContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms | 

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


## GetAudioByIdUniversal

> GetAudioByIdUniversal(ctx, id).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()

Gets an audio stream



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
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniversalAudioServiceAPI.GetAudioByIdUniversal(context.Background(), id).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversalAudioServiceAPI.GetAudioByIdUniversal``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAudioByIdUniversalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms | 

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


## HeadAudioByIdByContainer

> HeadAudioByIdByContainer(ctx, id, container).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()

Gets an audio stream



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
	container := "container_example" // string | 
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniversalAudioServiceAPI.HeadAudioByIdByContainer(context.Background(), id, container).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversalAudioServiceAPI.HeadAudioByIdByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**container** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadAudioByIdByContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms | 

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


## HeadAudioByIdUniversal

> HeadAudioByIdUniversal(ctx, id).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()

Gets an audio stream



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
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniversalAudioServiceAPI.HeadAudioByIdUniversal(context.Background(), id).DeviceId(deviceId).StartTimeTicks(startTimeTicks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversalAudioServiceAPI.HeadAudioByIdUniversal``: %v\n", err)
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

Other parameters are passed through a pointer to a apiHeadAudioByIdUniversalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms | 

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

