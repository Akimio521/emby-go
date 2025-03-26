# \EncodingInfoServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEncodingCodecconfigurationDefaults**](EncodingInfoServiceAPI.md#GetEncodingCodecconfigurationDefaults) | **Get** /Encoding/CodecConfiguration/Defaults | Gets default codec configurations
[**GetEncodingCodecinformationVideo**](EncodingInfoServiceAPI.md#GetEncodingCodecinformationVideo) | **Get** /Encoding/CodecInformation/Video | Gets details about available video encoders and decoders



## GetEncodingCodecconfigurationDefaults

> []ConfigurationCodecConfiguration GetEncodingCodecconfigurationDefaults(ctx).Execute()

Gets default codec configurations



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
	resp, r, err := apiClient.EncodingInfoServiceAPI.GetEncodingCodecconfigurationDefaults(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncodingInfoServiceAPI.GetEncodingCodecconfigurationDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingCodecconfigurationDefaults`: []ConfigurationCodecConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EncodingInfoServiceAPI.GetEncodingCodecconfigurationDefaults`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingCodecconfigurationDefaultsRequest struct via the builder pattern


### Return type

[**[]ConfigurationCodecConfiguration**](ConfigurationCodecConfiguration.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncodingCodecinformationVideo

> []MediaEncodingCodecsVideoCodecsVideoCodecBase GetEncodingCodecinformationVideo(ctx).Execute()

Gets details about available video encoders and decoders



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
	resp, r, err := apiClient.EncodingInfoServiceAPI.GetEncodingCodecinformationVideo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncodingInfoServiceAPI.GetEncodingCodecinformationVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingCodecinformationVideo`: []MediaEncodingCodecsVideoCodecsVideoCodecBase
	fmt.Fprintf(os.Stdout, "Response from `EncodingInfoServiceAPI.GetEncodingCodecinformationVideo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingCodecinformationVideoRequest struct via the builder pattern


### Return type

[**[]MediaEncodingCodecsVideoCodecsVideoCodecBase**](MediaEncodingCodecsVideoCodecsVideoCodecBase.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

