# \DlnaServerServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDlnaByUuidConnectionmanagerConnectionmanager**](DlnaServerServiceAPI.md#GetDlnaByUuidConnectionmanagerConnectionmanager) | **Get** /Dlna/{UuId}/connectionmanager/connectionmanager | Gets dlna connection manager xml
[**GetDlnaByUuidConnectionmanagerConnectionmanagerXml**](DlnaServerServiceAPI.md#GetDlnaByUuidConnectionmanagerConnectionmanagerXml) | **Get** /Dlna/{UuId}/connectionmanager/connectionmanager.xml | Gets dlna connection manager xml
[**GetDlnaByUuidContentdirectoryContentdirectory**](DlnaServerServiceAPI.md#GetDlnaByUuidContentdirectoryContentdirectory) | **Get** /Dlna/{UuId}/contentdirectory/contentdirectory | Gets dlna content directory xml
[**GetDlnaByUuidContentdirectoryContentdirectoryXml**](DlnaServerServiceAPI.md#GetDlnaByUuidContentdirectoryContentdirectoryXml) | **Get** /Dlna/{UuId}/contentdirectory/contentdirectory.xml | Gets dlna content directory xml
[**GetDlnaByUuidDescription**](DlnaServerServiceAPI.md#GetDlnaByUuidDescription) | **Get** /Dlna/{UuId}/description | Gets dlna server info
[**GetDlnaByUuidDescriptionXml**](DlnaServerServiceAPI.md#GetDlnaByUuidDescriptionXml) | **Get** /Dlna/{UuId}/description.xml | Gets dlna server info
[**GetDlnaByUuidIconsByFilename**](DlnaServerServiceAPI.md#GetDlnaByUuidIconsByFilename) | **Get** /Dlna/{UuId}/icons/{Filename} | Gets a server icon
[**GetDlnaIconsByFilename**](DlnaServerServiceAPI.md#GetDlnaIconsByFilename) | **Get** /Dlna/icons/{Filename} | Gets a server icon
[**PostDlnaByUuidConnectionmanagerControl**](DlnaServerServiceAPI.md#PostDlnaByUuidConnectionmanagerControl) | **Post** /Dlna/{UuId}/connectionmanager/control | Processes a control request
[**PostDlnaByUuidContentdirectoryControl**](DlnaServerServiceAPI.md#PostDlnaByUuidContentdirectoryControl) | **Post** /Dlna/{UuId}/contentdirectory/control | Processes a control request



## GetDlnaByUuidConnectionmanagerConnectionmanager

> GetDlnaByUuidConnectionmanagerConnectionmanager(ctx, uuId).Execute()

Gets dlna connection manager xml



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
	uuId := "uuId_example" // string | Server UuId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidConnectionmanagerConnectionmanager(context.Background(), uuId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerServiceAPI.GetDlnaByUuidConnectionmanagerConnectionmanager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuId** | **string** | Server UuId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlnaByUuidConnectionmanagerConnectionmanagerRequest struct via the builder pattern


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


## GetDlnaByUuidConnectionmanagerConnectionmanagerXml

> GetDlnaByUuidConnectionmanagerConnectionmanagerXml(ctx, uuId).Execute()

Gets dlna connection manager xml



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
	uuId := "uuId_example" // string | Server UuId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidConnectionmanagerConnectionmanagerXml(context.Background(), uuId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerServiceAPI.GetDlnaByUuidConnectionmanagerConnectionmanagerXml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuId** | **string** | Server UuId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlnaByUuidConnectionmanagerConnectionmanagerXmlRequest struct via the builder pattern


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


## GetDlnaByUuidContentdirectoryContentdirectory

> GetDlnaByUuidContentdirectoryContentdirectory(ctx, uuId).Execute()

Gets dlna content directory xml



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
	uuId := "uuId_example" // string | Server UuId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidContentdirectoryContentdirectory(context.Background(), uuId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerServiceAPI.GetDlnaByUuidContentdirectoryContentdirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuId** | **string** | Server UuId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlnaByUuidContentdirectoryContentdirectoryRequest struct via the builder pattern


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


## GetDlnaByUuidContentdirectoryContentdirectoryXml

> GetDlnaByUuidContentdirectoryContentdirectoryXml(ctx, uuId).Execute()

Gets dlna content directory xml



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
	uuId := "uuId_example" // string | Server UuId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidContentdirectoryContentdirectoryXml(context.Background(), uuId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerServiceAPI.GetDlnaByUuidContentdirectoryContentdirectoryXml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuId** | **string** | Server UuId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlnaByUuidContentdirectoryContentdirectoryXmlRequest struct via the builder pattern


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


## GetDlnaByUuidDescription

> GetDlnaByUuidDescription(ctx, uuId).Execute()

Gets dlna server info



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
	uuId := "uuId_example" // string | Server UuId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidDescription(context.Background(), uuId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerServiceAPI.GetDlnaByUuidDescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuId** | **string** | Server UuId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlnaByUuidDescriptionRequest struct via the builder pattern


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


## GetDlnaByUuidDescriptionXml

> GetDlnaByUuidDescriptionXml(ctx, uuId).Execute()

Gets dlna server info



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
	uuId := "uuId_example" // string | Server UuId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidDescriptionXml(context.Background(), uuId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerServiceAPI.GetDlnaByUuidDescriptionXml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuId** | **string** | Server UuId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlnaByUuidDescriptionXmlRequest struct via the builder pattern


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


## GetDlnaByUuidIconsByFilename

> GetDlnaByUuidIconsByFilename(ctx, uuId, filename).Execute()

Gets a server icon



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
	uuId := "uuId_example" // string | Server UuId
	filename := "filename_example" // string | The icon filename

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidIconsByFilename(context.Background(), uuId, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerServiceAPI.GetDlnaByUuidIconsByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuId** | **string** | Server UuId | 
**filename** | **string** | The icon filename | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlnaByUuidIconsByFilenameRequest struct via the builder pattern


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


## GetDlnaIconsByFilename

> GetDlnaIconsByFilename(ctx, filename).UuId(uuId).Execute()

Gets a server icon



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
	filename := "filename_example" // string | The icon filename
	uuId := "uuId_example" // string | Server UuId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServerServiceAPI.GetDlnaIconsByFilename(context.Background(), filename).UuId(uuId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerServiceAPI.GetDlnaIconsByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | The icon filename | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlnaIconsByFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uuId** | **string** | Server UuId | 

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


## PostDlnaByUuidConnectionmanagerControl

> PostDlnaByUuidConnectionmanagerControl(ctx, uuId).Body(body).Execute()

Processes a control request



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
	uuId := "uuId_example" // string | Server UuId
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServerServiceAPI.PostDlnaByUuidConnectionmanagerControl(context.Background(), uuId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerServiceAPI.PostDlnaByUuidConnectionmanagerControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuId** | **string** | Server UuId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDlnaByUuidConnectionmanagerControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | Binary stream | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDlnaByUuidContentdirectoryControl

> PostDlnaByUuidContentdirectoryControl(ctx, uuId).Body(body).Execute()

Processes a control request



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
	uuId := "uuId_example" // string | Server UuId
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaServerServiceAPI.PostDlnaByUuidContentdirectoryControl(context.Background(), uuId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerServiceAPI.PostDlnaByUuidContentdirectoryControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuId** | **string** | Server UuId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDlnaByUuidContentdirectoryControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | Binary stream | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

