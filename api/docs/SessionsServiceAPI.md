# \SessionsServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuthKeysByKey**](SessionsServiceAPI.md#DeleteAuthKeysByKey) | **Delete** /Auth/Keys/{Key} | 
[**DeleteSessionsByIdUsersByUserid**](SessionsServiceAPI.md#DeleteSessionsByIdUsersByUserid) | **Delete** /Sessions/{Id}/Users/{UserId} | Removes an additional user from a session
[**GetAuthKeys**](SessionsServiceAPI.md#GetAuthKeys) | **Get** /Auth/Keys | 
[**GetAuthProviders**](SessionsServiceAPI.md#GetAuthProviders) | **Get** /Auth/Providers | 
[**GetSessions**](SessionsServiceAPI.md#GetSessions) | **Get** /Sessions | Gets a list of sessions
[**PostAuthKeys**](SessionsServiceAPI.md#PostAuthKeys) | **Post** /Auth/Keys | 
[**PostSessionsByIdCommand**](SessionsServiceAPI.md#PostSessionsByIdCommand) | **Post** /Sessions/{Id}/Command | Issues a system command to a client
[**PostSessionsByIdCommandByCommand**](SessionsServiceAPI.md#PostSessionsByIdCommandByCommand) | **Post** /Sessions/{Id}/Command/{Command} | Issues a system command to a client
[**PostSessionsByIdMessage**](SessionsServiceAPI.md#PostSessionsByIdMessage) | **Post** /Sessions/{Id}/Message | Issues a command to a client to display a message to the user
[**PostSessionsByIdPlaying**](SessionsServiceAPI.md#PostSessionsByIdPlaying) | **Post** /Sessions/{Id}/Playing | Instructs a session to play an item
[**PostSessionsByIdPlayingByCommand**](SessionsServiceAPI.md#PostSessionsByIdPlayingByCommand) | **Post** /Sessions/{Id}/Playing/{Command} | Issues a playstate command to a client
[**PostSessionsByIdSystemByCommand**](SessionsServiceAPI.md#PostSessionsByIdSystemByCommand) | **Post** /Sessions/{Id}/System/{Command} | Issues a system command to a client
[**PostSessionsByIdUsersByUserid**](SessionsServiceAPI.md#PostSessionsByIdUsersByUserid) | **Post** /Sessions/{Id}/Users/{UserId} | Adds an additional user to a session
[**PostSessionsByIdViewing**](SessionsServiceAPI.md#PostSessionsByIdViewing) | **Post** /Sessions/{Id}/Viewing | Instructs a session to browse to an item or view
[**PostSessionsCapabilities**](SessionsServiceAPI.md#PostSessionsCapabilities) | **Post** /Sessions/Capabilities | Updates capabilities for a device
[**PostSessionsCapabilitiesFull**](SessionsServiceAPI.md#PostSessionsCapabilitiesFull) | **Post** /Sessions/Capabilities/Full | Updates capabilities for a device
[**PostSessionsLogout**](SessionsServiceAPI.md#PostSessionsLogout) | **Post** /Sessions/Logout | Reports that a session has ended



## DeleteAuthKeysByKey

> DeleteAuthKeysByKey(ctx, key).Execute()





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
	key := "key_example" // string | Auth Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.DeleteAuthKeysByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.DeleteAuthKeysByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Auth Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthKeysByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteSessionsByIdUsersByUserid

> DeleteSessionsByIdUsersByUserid(ctx, id, userId).Execute()

Removes an additional user from a session



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
	id := "id_example" // string | Session Id
	userId := "userId_example" // string | UserId Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.DeleteSessionsByIdUsersByUserid(context.Background(), id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.DeleteSessionsByIdUsersByUserid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Session Id | 
**userId** | **string** | UserId Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionsByIdUsersByUseridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetAuthKeys

> GetAuthKeys(ctx).Execute()





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
	r, err := apiClient.SessionsServiceAPI.GetAuthKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.GetAuthKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthKeysRequest struct via the builder pattern


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


## GetAuthProviders

> []NameIdPair GetAuthProviders(ctx).Execute()





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
	resp, r, err := apiClient.SessionsServiceAPI.GetAuthProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.GetAuthProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthProviders`: []NameIdPair
	fmt.Fprintf(os.Stdout, "Response from `SessionsServiceAPI.GetAuthProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthProvidersRequest struct via the builder pattern


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


## GetSessions

> []SessionSessionInfo GetSessions(ctx).ControllableByUserId(controllableByUserId).DeviceId(deviceId).Execute()

Gets a list of sessions



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
	controllableByUserId := "controllableByUserId_example" // string | Optional. Filter by sessions that a given user is allowed to remote control. (optional)
	deviceId := "deviceId_example" // string | Optional. Filter by device id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsServiceAPI.GetSessions(context.Background()).ControllableByUserId(controllableByUserId).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.GetSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessions`: []SessionSessionInfo
	fmt.Fprintf(os.Stdout, "Response from `SessionsServiceAPI.GetSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controllableByUserId** | **string** | Optional. Filter by sessions that a given user is allowed to remote control. | 
 **deviceId** | **string** | Optional. Filter by device id. | 

### Return type

[**[]SessionSessionInfo**](SessionSessionInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthKeys

> PostAuthKeys(ctx).App(app).Execute()





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
	app := "app_example" // string | App

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.PostAuthKeys(context.Background()).App(app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostAuthKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **string** | App | 

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


## PostSessionsByIdCommand

> PostSessionsByIdCommand(ctx, id).GeneralCommand(generalCommand).Execute()

Issues a system command to a client



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
	id := "id_example" // string | Session Id
	generalCommand := *openapiclient.NewGeneralCommand() // GeneralCommand | GeneralCommand: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.PostSessionsByIdCommand(context.Background(), id).GeneralCommand(generalCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostSessionsByIdCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Session Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsByIdCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generalCommand** | [**GeneralCommand**](GeneralCommand.md) | GeneralCommand:  | 

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


## PostSessionsByIdCommandByCommand

> PostSessionsByIdCommandByCommand(ctx, id, command).Execute()

Issues a system command to a client



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
	id := "id_example" // string | Session Id
	command := "command_example" // string | The command to send.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.PostSessionsByIdCommandByCommand(context.Background(), id, command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostSessionsByIdCommandByCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Session Id | 
**command** | **string** | The command to send. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsByIdCommandByCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## PostSessionsByIdMessage

> PostSessionsByIdMessage(ctx, id).Text(text).Header(header).TimeoutMs(timeoutMs).Execute()

Issues a command to a client to display a message to the user



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
	id := "id_example" // string | Session Id
	text := "text_example" // string | The message text.
	header := "header_example" // string | The message header.
	timeoutMs := int64(789) // int64 | The message timeout. If omitted the user will have to confirm viewing the message. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.PostSessionsByIdMessage(context.Background(), id).Text(text).Header(header).TimeoutMs(timeoutMs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostSessionsByIdMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Session Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsByIdMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **text** | **string** | The message text. | 
 **header** | **string** | The message header. | 
 **timeoutMs** | **int64** | The message timeout. If omitted the user will have to confirm viewing the message. | 

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


## PostSessionsByIdPlaying

> PostSessionsByIdPlaying(ctx, id).ItemIds(itemIds).PlayCommand(playCommand).PlayRequest(playRequest).StartPositionTicks(startPositionTicks).Execute()

Instructs a session to play an item



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
	id := "id_example" // string | Session Id
	itemIds := []int64{int64(123)} // []int64 | The ids of the items to play, comma delimited
	playCommand := "playCommand_example" // string | The type of play command to issue (PlayNow, PlayNext, PlayLast). Clients who have not yet implemented play next and play last may play now.
	playRequest := *openapiclient.NewPlayRequest() // PlayRequest | PlayRequest: 
	startPositionTicks := int64(789) // int64 | The starting position of the first item. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.PostSessionsByIdPlaying(context.Background(), id).ItemIds(itemIds).PlayCommand(playCommand).PlayRequest(playRequest).StartPositionTicks(startPositionTicks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostSessionsByIdPlaying``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Session Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsByIdPlayingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemIds** | **[]int64** | The ids of the items to play, comma delimited | 
 **playCommand** | **string** | The type of play command to issue (PlayNow, PlayNext, PlayLast). Clients who have not yet implemented play next and play last may play now. | 
 **playRequest** | [**PlayRequest**](PlayRequest.md) | PlayRequest:  | 
 **startPositionTicks** | **int64** | The starting position of the first item. | 

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


## PostSessionsByIdPlayingByCommand

> PostSessionsByIdPlayingByCommand(ctx, id, command).PlaystateRequest(playstateRequest).Execute()

Issues a playstate command to a client



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
	id := "id_example" // string | Session Id
	command := "command_example" // string | 
	playstateRequest := *openapiclient.NewPlaystateRequest() // PlaystateRequest | PlaystateRequest: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.PostSessionsByIdPlayingByCommand(context.Background(), id, command).PlaystateRequest(playstateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostSessionsByIdPlayingByCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Session Id | 
**command** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsByIdPlayingByCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **playstateRequest** | [**PlaystateRequest**](PlaystateRequest.md) | PlaystateRequest:  | 

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


## PostSessionsByIdSystemByCommand

> PostSessionsByIdSystemByCommand(ctx, id, command).Execute()

Issues a system command to a client



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
	id := "id_example" // string | Session Id
	command := "command_example" // string | The command to send.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.PostSessionsByIdSystemByCommand(context.Background(), id, command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostSessionsByIdSystemByCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Session Id | 
**command** | **string** | The command to send. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsByIdSystemByCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## PostSessionsByIdUsersByUserid

> PostSessionsByIdUsersByUserid(ctx, id, userId).Execute()

Adds an additional user to a session



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
	id := "id_example" // string | Session Id
	userId := "userId_example" // string | UserId Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.PostSessionsByIdUsersByUserid(context.Background(), id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostSessionsByIdUsersByUserid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Session Id | 
**userId** | **string** | UserId Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsByIdUsersByUseridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## PostSessionsByIdViewing

> PostSessionsByIdViewing(ctx, id).ItemType(itemType).ItemId(itemId).ItemName(itemName).Execute()

Instructs a session to browse to an item or view



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
	id := "id_example" // string | Session Id
	itemType := "itemType_example" // string | The type of item to browse to.
	itemId := "itemId_example" // string | The Id of the item.
	itemName := "itemName_example" // string | The name of the item.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.PostSessionsByIdViewing(context.Background(), id).ItemType(itemType).ItemId(itemId).ItemName(itemName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostSessionsByIdViewing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Session Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsByIdViewingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemType** | **string** | The type of item to browse to. | 
 **itemId** | **string** | The Id of the item. | 
 **itemName** | **string** | The name of the item. | 

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


## PostSessionsCapabilities

> PostSessionsCapabilities(ctx).Id(id).PlayableMediaTypes(playableMediaTypes).SupportedCommands(supportedCommands).SupportsMediaControl(supportsMediaControl).SupportsSync(supportsSync).SupportsPersistentIdentifier(supportsPersistentIdentifier).Execute()

Updates capabilities for a device



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
	id := "id_example" // string | Session Id
	playableMediaTypes := "playableMediaTypes_example" // string | A list of playable media types, comma delimited. Audio, Video, Book, Game, Photo. (optional)
	supportedCommands := "supportedCommands_example" // string | A list of supported remote control commands, comma delimited (optional)
	supportsMediaControl := true // bool | Determines whether media can be played remotely. (optional)
	supportsSync := true // bool | Determines whether sync is supported. (optional)
	supportsPersistentIdentifier := true // bool | Determines whether the device supports a unique identifier. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.PostSessionsCapabilities(context.Background()).Id(id).PlayableMediaTypes(playableMediaTypes).SupportedCommands(supportedCommands).SupportsMediaControl(supportsMediaControl).SupportsSync(supportsSync).SupportsPersistentIdentifier(supportsPersistentIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostSessionsCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Session Id | 
 **playableMediaTypes** | **string** | A list of playable media types, comma delimited. Audio, Video, Book, Game, Photo. | 
 **supportedCommands** | **string** | A list of supported remote control commands, comma delimited | 
 **supportsMediaControl** | **bool** | Determines whether media can be played remotely. | 
 **supportsSync** | **bool** | Determines whether sync is supported. | 
 **supportsPersistentIdentifier** | **bool** | Determines whether the device supports a unique identifier. | 

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


## PostSessionsCapabilitiesFull

> PostSessionsCapabilitiesFull(ctx).Id(id).ClientCapabilities(clientCapabilities).Execute()

Updates capabilities for a device



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
	id := "id_example" // string | Session Id
	clientCapabilities := *openapiclient.NewClientCapabilities() // ClientCapabilities | ClientCapabilities: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsServiceAPI.PostSessionsCapabilitiesFull(context.Background()).Id(id).ClientCapabilities(clientCapabilities).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostSessionsCapabilitiesFull``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsCapabilitiesFullRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Session Id | 
 **clientCapabilities** | [**ClientCapabilities**](ClientCapabilities.md) | ClientCapabilities:  | 

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


## PostSessionsLogout

> PostSessionsLogout(ctx).Execute()

Reports that a session has ended



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
	r, err := apiClient.SessionsServiceAPI.PostSessionsLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsServiceAPI.PostSessionsLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionsLogoutRequest struct via the builder pattern


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

