# \DeviceServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDevices**](DeviceServiceAPI.md#DeleteDevices) | **Delete** /Devices | Deletes a device
[**GetDevices**](DeviceServiceAPI.md#GetDevices) | **Get** /Devices | Gets all devices
[**GetDevicesCamerauploads**](DeviceServiceAPI.md#GetDevicesCamerauploads) | **Get** /Devices/CameraUploads | Gets camera upload history for a device
[**GetDevicesInfo**](DeviceServiceAPI.md#GetDevicesInfo) | **Get** /Devices/Info | Gets info for a device
[**GetDevicesOptions**](DeviceServiceAPI.md#GetDevicesOptions) | **Get** /Devices/Options | Gets options for a device
[**PostDevicesCamerauploads**](DeviceServiceAPI.md#PostDevicesCamerauploads) | **Post** /Devices/CameraUploads | Uploads content
[**PostDevicesOptions**](DeviceServiceAPI.md#PostDevicesOptions) | **Post** /Devices/Options | Updates device options



## DeleteDevices

> DeleteDevices(ctx).Id(id).Execute()

Deletes a device



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
	id := "id_example" // string | Device Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceServiceAPI.DeleteDevices(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.DeleteDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id | 

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


## GetDevices

> QueryResultDevicesDeviceInfo GetDevices(ctx).Execute()

Gets all devices



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
	resp, r, err := apiClient.DeviceServiceAPI.GetDevices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.GetDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevices`: QueryResultDevicesDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `DeviceServiceAPI.GetDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


### Return type

[**QueryResultDevicesDeviceInfo**](QueryResultDevicesDeviceInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesCamerauploads

> DevicesContentUploadHistory GetDevicesCamerauploads(ctx).DeviceId(deviceId).Execute()

Gets camera upload history for a device



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
	deviceId := "deviceId_example" // string | Device Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceServiceAPI.GetDevicesCamerauploads(context.Background()).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.GetDevicesCamerauploads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesCamerauploads`: DevicesContentUploadHistory
	fmt.Fprintf(os.Stdout, "Response from `DeviceServiceAPI.GetDevicesCamerauploads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesCamerauploadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 

### Return type

[**DevicesContentUploadHistory**](DevicesContentUploadHistory.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesInfo

> DevicesDeviceInfo GetDevicesInfo(ctx).Id(id).Execute()

Gets info for a device



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
	id := "id_example" // string | Device Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceServiceAPI.GetDevicesInfo(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.GetDevicesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesInfo`: DevicesDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `DeviceServiceAPI.GetDevicesInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id | 

### Return type

[**DevicesDeviceInfo**](DevicesDeviceInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesOptions

> DevicesDeviceOptions GetDevicesOptions(ctx).Id(id).Execute()

Gets options for a device



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
	id := "id_example" // string | Device Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceServiceAPI.GetDevicesOptions(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.GetDevicesOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesOptions`: DevicesDeviceOptions
	fmt.Fprintf(os.Stdout, "Response from `DeviceServiceAPI.GetDevicesOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id | 

### Return type

[**DevicesDeviceOptions**](DevicesDeviceOptions.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDevicesCamerauploads

> PostDevicesCamerauploads(ctx).DeviceId(deviceId).Album(album).Name(name).Id(id).Body(body).Execute()

Uploads content



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
	deviceId := "deviceId_example" // string | Device Id
	album := "album_example" // string | Album
	name := "name_example" // string | Name
	id := "id_example" // string | Id
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceServiceAPI.PostDevicesCamerauploads(context.Background()).DeviceId(deviceId).Album(album).Name(name).Id(id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.PostDevicesCamerauploads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDevicesCamerauploadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **album** | **string** | Album | 
 **name** | **string** | Name | 
 **id** | **string** | Id | 
 **body** | ***os.File** | Binary stream | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDevicesOptions

> PostDevicesOptions(ctx).DevicesDeviceOptions(devicesDeviceOptions).Execute()

Updates device options



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
	devicesDeviceOptions := *openapiclient.NewDevicesDeviceOptions() // DevicesDeviceOptions | DeviceOptions: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceServiceAPI.PostDevicesOptions(context.Background()).DevicesDeviceOptions(devicesDeviceOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceServiceAPI.PostDevicesOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDevicesOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **devicesDeviceOptions** | [**DevicesDeviceOptions**](DevicesDeviceOptions.md) | DeviceOptions:  | 

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

