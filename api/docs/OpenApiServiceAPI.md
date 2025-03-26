# \OpenApiServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOpenapi**](OpenApiServiceAPI.md#GetOpenapi) | **Get** /openapi | Gets the OpenAPI 3 specifications
[**GetOpenapiJson**](OpenApiServiceAPI.md#GetOpenapiJson) | **Get** /openapi.json | Gets OpenAPI 3 specifications
[**GetSwagger**](OpenApiServiceAPI.md#GetSwagger) | **Get** /swagger | Gets the swagger specifications
[**GetSwaggerJson**](OpenApiServiceAPI.md#GetSwaggerJson) | **Get** /swagger.json | Gets the swagger specifications



## GetOpenapi

> string GetOpenapi(ctx).Execute()

Gets the OpenAPI 3 specifications



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
	resp, r, err := apiClient.OpenApiServiceAPI.GetOpenapi(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenApiServiceAPI.GetOpenapi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenapi`: string
	fmt.Fprintf(os.Stdout, "Response from `OpenApiServiceAPI.GetOpenapi`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenapiRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenapiJson

> string GetOpenapiJson(ctx).Execute()

Gets OpenAPI 3 specifications



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
	resp, r, err := apiClient.OpenApiServiceAPI.GetOpenapiJson(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenApiServiceAPI.GetOpenapiJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenapiJson`: string
	fmt.Fprintf(os.Stdout, "Response from `OpenApiServiceAPI.GetOpenapiJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenapiJsonRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwagger

> string GetSwagger(ctx).Execute()

Gets the swagger specifications



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
	resp, r, err := apiClient.OpenApiServiceAPI.GetSwagger(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenApiServiceAPI.GetSwagger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwagger`: string
	fmt.Fprintf(os.Stdout, "Response from `OpenApiServiceAPI.GetSwagger`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwaggerRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwaggerJson

> string GetSwaggerJson(ctx).Execute()

Gets the swagger specifications



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
	resp, r, err := apiClient.OpenApiServiceAPI.GetSwaggerJson(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenApiServiceAPI.GetSwaggerJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSwaggerJson`: string
	fmt.Fprintf(os.Stdout, "Response from `OpenApiServiceAPI.GetSwaggerJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwaggerJsonRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

