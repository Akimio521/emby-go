# \NotificationsServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationsByUserid**](NotificationsServiceAPI.md#GetNotificationsByUserid) | **Get** /Notifications/{UserId} | Gets notifications
[**GetNotificationsByUseridSummary**](NotificationsServiceAPI.md#GetNotificationsByUseridSummary) | **Get** /Notifications/{UserId}/Summary | Gets a notification summary for a user
[**GetNotificationsServices**](NotificationsServiceAPI.md#GetNotificationsServices) | **Get** /Notifications/Services | Gets notification types
[**GetNotificationsTypes**](NotificationsServiceAPI.md#GetNotificationsTypes) | **Get** /Notifications/Types | Gets notification types
[**PostNotificationsAdmin**](NotificationsServiceAPI.md#PostNotificationsAdmin) | **Post** /Notifications/Admin | Sends a notification to all admin users
[**PostNotificationsByUseridRead**](NotificationsServiceAPI.md#PostNotificationsByUseridRead) | **Post** /Notifications/{UserId}/Read | Marks notifications as read
[**PostNotificationsByUseridUnread**](NotificationsServiceAPI.md#PostNotificationsByUseridUnread) | **Post** /Notifications/{UserId}/Unread | Marks notifications as unread



## GetNotificationsByUserid

> EmbyNotificationsApiNotificationResult GetNotificationsByUserid(ctx, userId).IsRead(isRead).StartIndex(startIndex).Limit(limit).Execute()

Gets notifications



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
	isRead := true // bool | An optional filter by IsRead (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsServiceAPI.GetNotificationsByUserid(context.Background(), userId).IsRead(isRead).StartIndex(startIndex).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsServiceAPI.GetNotificationsByUserid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsByUserid`: EmbyNotificationsApiNotificationResult
	fmt.Fprintf(os.Stdout, "Response from `NotificationsServiceAPI.GetNotificationsByUserid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsByUseridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isRead** | **bool** | An optional filter by IsRead | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 

### Return type

[**EmbyNotificationsApiNotificationResult**](EmbyNotificationsApiNotificationResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsByUseridSummary

> EmbyNotificationsApiNotificationsSummary GetNotificationsByUseridSummary(ctx, userId).Execute()

Gets a notification summary for a user



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsServiceAPI.GetNotificationsByUseridSummary(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsServiceAPI.GetNotificationsByUseridSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsByUseridSummary`: EmbyNotificationsApiNotificationsSummary
	fmt.Fprintf(os.Stdout, "Response from `NotificationsServiceAPI.GetNotificationsByUseridSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsByUseridSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmbyNotificationsApiNotificationsSummary**](EmbyNotificationsApiNotificationsSummary.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsServices

> []NameIdPair GetNotificationsServices(ctx).Execute()

Gets notification types



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
	resp, r, err := apiClient.NotificationsServiceAPI.GetNotificationsServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsServiceAPI.GetNotificationsServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsServices`: []NameIdPair
	fmt.Fprintf(os.Stdout, "Response from `NotificationsServiceAPI.GetNotificationsServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsServicesRequest struct via the builder pattern


### Return type

[**[]NameIdPair**](NameIdPair.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsTypes

> []NotificationsNotificationTypeInfo GetNotificationsTypes(ctx).Execute()

Gets notification types



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
	resp, r, err := apiClient.NotificationsServiceAPI.GetNotificationsTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsServiceAPI.GetNotificationsTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsTypes`: []NotificationsNotificationTypeInfo
	fmt.Fprintf(os.Stdout, "Response from `NotificationsServiceAPI.GetNotificationsTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsTypesRequest struct via the builder pattern


### Return type

[**[]NotificationsNotificationTypeInfo**](NotificationsNotificationTypeInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotificationsAdmin

> PostNotificationsAdmin(ctx).Name(name).Description(description).ImageUrl(imageUrl).Url(url).Level(level).Execute()

Sends a notification to all admin users



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
	name := "name_example" // string | The notification's name
	description := "description_example" // string | The notification's description
	imageUrl := "imageUrl_example" // string | The notification's image url (optional)
	url := "url_example" // string | The notification's info url (optional)
	level := "level_example" // string | The notification level (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsServiceAPI.PostNotificationsAdmin(context.Background()).Name(name).Description(description).ImageUrl(imageUrl).Url(url).Level(level).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsServiceAPI.PostNotificationsAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNotificationsAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The notification&#39;s name | 
 **description** | **string** | The notification&#39;s description | 
 **imageUrl** | **string** | The notification&#39;s image url | 
 **url** | **string** | The notification&#39;s info url | 
 **level** | **string** | The notification level | 

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


## PostNotificationsByUseridRead

> PostNotificationsByUseridRead(ctx, userId).Ids(ids).Execute()

Marks notifications as read



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
	ids := "ids_example" // string | A list of notification ids, comma delimited

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsServiceAPI.PostNotificationsByUseridRead(context.Background(), userId).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsServiceAPI.PostNotificationsByUseridRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNotificationsByUseridReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **string** | A list of notification ids, comma delimited | 

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


## PostNotificationsByUseridUnread

> PostNotificationsByUseridUnread(ctx, userId).Ids(ids).Execute()

Marks notifications as unread



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
	ids := "ids_example" // string | A list of notification ids, comma delimited

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsServiceAPI.PostNotificationsByUseridUnread(context.Background(), userId).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsServiceAPI.PostNotificationsByUseridUnread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNotificationsByUseridUnreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **string** | A list of notification ids, comma delimited | 

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

