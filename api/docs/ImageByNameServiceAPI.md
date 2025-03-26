# \ImageByNameServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImagesGeneral**](ImageByNameServiceAPI.md#GetImagesGeneral) | **Get** /Images/General | Gets all general images by name
[**GetImagesGeneralByNameByType**](ImageByNameServiceAPI.md#GetImagesGeneralByNameByType) | **Get** /Images/General/{Name}/{Type} | Gets a general image by name
[**GetImagesMediainfo**](ImageByNameServiceAPI.md#GetImagesMediainfo) | **Get** /Images/MediaInfo | Gets all media info image by name
[**GetImagesMediainfoByThemeByName**](ImageByNameServiceAPI.md#GetImagesMediainfoByThemeByName) | **Get** /Images/MediaInfo/{Theme}/{Name} | Gets a media info image by name
[**GetImagesRatings**](ImageByNameServiceAPI.md#GetImagesRatings) | **Get** /Images/Ratings | Gets all rating images by name
[**GetImagesRatingsByThemeByName**](ImageByNameServiceAPI.md#GetImagesRatingsByThemeByName) | **Get** /Images/Ratings/{Theme}/{Name} | Gets a rating image by name



## GetImagesGeneral

> []ImageByNameInfo GetImagesGeneral(ctx).Execute()

Gets all general images by name



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
	resp, r, err := apiClient.ImageByNameServiceAPI.GetImagesGeneral(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameServiceAPI.GetImagesGeneral``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImagesGeneral`: []ImageByNameInfo
	fmt.Fprintf(os.Stdout, "Response from `ImageByNameServiceAPI.GetImagesGeneral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesGeneralRequest struct via the builder pattern


### Return type

[**[]ImageByNameInfo**](ImageByNameInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImagesGeneralByNameByType

> GetImagesGeneralByNameByType(ctx, name, type_).Execute()

Gets a general image by name



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
	name := "name_example" // string | The name of the image
	type_ := "type__example" // string | Image Type (primary, backdrop, logo, etc).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageByNameServiceAPI.GetImagesGeneralByNameByType(context.Background(), name, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameServiceAPI.GetImagesGeneralByNameByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the image | 
**type_** | **string** | Image Type (primary, backdrop, logo, etc). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesGeneralByNameByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetImagesMediainfo

> []ImageByNameInfo GetImagesMediainfo(ctx).Execute()

Gets all media info image by name



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
	resp, r, err := apiClient.ImageByNameServiceAPI.GetImagesMediainfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameServiceAPI.GetImagesMediainfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImagesMediainfo`: []ImageByNameInfo
	fmt.Fprintf(os.Stdout, "Response from `ImageByNameServiceAPI.GetImagesMediainfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesMediainfoRequest struct via the builder pattern


### Return type

[**[]ImageByNameInfo**](ImageByNameInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImagesMediainfoByThemeByName

> GetImagesMediainfoByThemeByName(ctx, name, theme).Execute()

Gets a media info image by name



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
	name := "name_example" // string | The name of the image
	theme := "theme_example" // string | The theme to get the image from

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageByNameServiceAPI.GetImagesMediainfoByThemeByName(context.Background(), name, theme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameServiceAPI.GetImagesMediainfoByThemeByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the image | 
**theme** | **string** | The theme to get the image from | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesMediainfoByThemeByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetImagesRatings

> []ImageByNameInfo GetImagesRatings(ctx).Execute()

Gets all rating images by name



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
	resp, r, err := apiClient.ImageByNameServiceAPI.GetImagesRatings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameServiceAPI.GetImagesRatings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImagesRatings`: []ImageByNameInfo
	fmt.Fprintf(os.Stdout, "Response from `ImageByNameServiceAPI.GetImagesRatings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesRatingsRequest struct via the builder pattern


### Return type

[**[]ImageByNameInfo**](ImageByNameInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImagesRatingsByThemeByName

> GetImagesRatingsByThemeByName(ctx, name, theme).Execute()

Gets a rating image by name



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
	name := "name_example" // string | The name of the image
	theme := "theme_example" // string | The theme to get the image from

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageByNameServiceAPI.GetImagesRatingsByThemeByName(context.Background(), name, theme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameServiceAPI.GetImagesRatingsByThemeByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the image | 
**theme** | **string** | The theme to get the image from | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesRatingsByThemeByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

