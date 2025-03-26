# \EnvironmentServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnvironmentDefaultdirectorybrowser**](EnvironmentServiceAPI.md#GetEnvironmentDefaultdirectorybrowser) | **Get** /Environment/DefaultDirectoryBrowser | Gets the parent path of a given path
[**GetEnvironmentDirectorycontents**](EnvironmentServiceAPI.md#GetEnvironmentDirectorycontents) | **Get** /Environment/DirectoryContents | Gets the contents of a given directory in the file system
[**GetEnvironmentDrives**](EnvironmentServiceAPI.md#GetEnvironmentDrives) | **Get** /Environment/Drives | Gets available drives from the server&#39;s file system
[**GetEnvironmentNetworkdevices**](EnvironmentServiceAPI.md#GetEnvironmentNetworkdevices) | **Get** /Environment/NetworkDevices | Gets a list of devices on the network
[**GetEnvironmentNetworkshares**](EnvironmentServiceAPI.md#GetEnvironmentNetworkshares) | **Get** /Environment/NetworkShares | Gets shares from a network device
[**GetEnvironmentParentpath**](EnvironmentServiceAPI.md#GetEnvironmentParentpath) | **Get** /Environment/ParentPath | Gets the parent path of a given path
[**PostEnvironmentValidatepath**](EnvironmentServiceAPI.md#PostEnvironmentValidatepath) | **Post** /Environment/ValidatePath | Gets the contents of a given directory in the file system



## GetEnvironmentDefaultdirectorybrowser

> DefaultDirectoryBrowserInfo GetEnvironmentDefaultdirectorybrowser(ctx).Execute()

Gets the parent path of a given path



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
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentDefaultdirectorybrowser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentDefaultdirectorybrowser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentDefaultdirectorybrowser`: DefaultDirectoryBrowserInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentDefaultdirectorybrowser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentDefaultdirectorybrowserRequest struct via the builder pattern


### Return type

[**DefaultDirectoryBrowserInfo**](DefaultDirectoryBrowserInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentDirectorycontents

> []IOFileSystemEntryInfo GetEnvironmentDirectorycontents(ctx).Path(path).IncludeFiles(includeFiles).IncludeDirectories(includeDirectories).Execute()

Gets the contents of a given directory in the file system



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
	path := "path_example" // string | 
	includeFiles := true // bool | An optional filter to include or exclude files from the results. true/false (optional)
	includeDirectories := true // bool | An optional filter to include or exclude folders from the results. true/false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentDirectorycontents(context.Background()).Path(path).IncludeFiles(includeFiles).IncludeDirectories(includeDirectories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentDirectorycontents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentDirectorycontents`: []IOFileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentDirectorycontents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentDirectorycontentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **includeFiles** | **bool** | An optional filter to include or exclude files from the results. true/false | 
 **includeDirectories** | **bool** | An optional filter to include or exclude folders from the results. true/false | 

### Return type

[**[]IOFileSystemEntryInfo**](IOFileSystemEntryInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentDrives

> []IOFileSystemEntryInfo GetEnvironmentDrives(ctx).Execute()

Gets available drives from the server's file system



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
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentDrives(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentDrives`: []IOFileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentDrives`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentDrivesRequest struct via the builder pattern


### Return type

[**[]IOFileSystemEntryInfo**](IOFileSystemEntryInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentNetworkdevices

> []IOFileSystemEntryInfo GetEnvironmentNetworkdevices(ctx).Execute()

Gets a list of devices on the network



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
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentNetworkdevices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentNetworkdevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentNetworkdevices`: []IOFileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentNetworkdevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentNetworkdevicesRequest struct via the builder pattern


### Return type

[**[]IOFileSystemEntryInfo**](IOFileSystemEntryInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentNetworkshares

> []IOFileSystemEntryInfo GetEnvironmentNetworkshares(ctx).Path(path).Execute()

Gets shares from a network device



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentNetworkshares(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentNetworkshares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentNetworkshares`: []IOFileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentNetworkshares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentNetworksharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

### Return type

[**[]IOFileSystemEntryInfo**](IOFileSystemEntryInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentParentpath

> string GetEnvironmentParentpath(ctx).Path(path).Execute()

Gets the parent path of a given path



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentParentpath(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentParentpath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentParentpath`: string
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentParentpath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentParentpathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

### Return type

**string**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEnvironmentValidatepath

> PostEnvironmentValidatepath(ctx).Path(path).ValidatePath(validatePath).Execute()

Gets the contents of a given directory in the file system



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
	path := "path_example" // string | 
	validatePath := *openapiclient.NewValidatePath() // ValidatePath | ValidatePath

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentServiceAPI.PostEnvironmentValidatepath(context.Background()).Path(path).ValidatePath(validatePath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.PostEnvironmentValidatepath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEnvironmentValidatepathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **validatePath** | [**ValidatePath**](ValidatePath.md) | ValidatePath | 

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

