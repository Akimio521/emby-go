# \SyncServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSyncByTargetidItems**](SyncServiceAPI.md#DeleteSyncByTargetidItems) | **Delete** /Sync/{TargetId}/Items | Cancels items from a sync target
[**DeleteSyncJobitemsById**](SyncServiceAPI.md#DeleteSyncJobitemsById) | **Delete** /Sync/JobItems/{Id} | Cancels a sync job item
[**DeleteSyncJobsById**](SyncServiceAPI.md#DeleteSyncJobsById) | **Delete** /Sync/Jobs/{Id} | Cancels a sync job.
[**GetSyncItemsReady**](SyncServiceAPI.md#GetSyncItemsReady) | **Get** /Sync/Items/Ready | Gets ready to download sync items.
[**GetSyncJobitems**](SyncServiceAPI.md#GetSyncJobitems) | **Get** /Sync/JobItems | Gets sync job items.
[**GetSyncJobitemsByIdAdditionalfiles**](SyncServiceAPI.md#GetSyncJobitemsByIdAdditionalfiles) | **Get** /Sync/JobItems/{Id}/AdditionalFiles | Gets a sync job item file
[**GetSyncJobitemsByIdFile**](SyncServiceAPI.md#GetSyncJobitemsByIdFile) | **Get** /Sync/JobItems/{Id}/File | Gets a sync job item file
[**GetSyncJobs**](SyncServiceAPI.md#GetSyncJobs) | **Get** /Sync/Jobs | Gets sync jobs.
[**GetSyncJobsById**](SyncServiceAPI.md#GetSyncJobsById) | **Get** /Sync/Jobs/{Id} | Gets a sync job.
[**GetSyncOptions**](SyncServiceAPI.md#GetSyncOptions) | **Get** /Sync/Options | Gets a list of available sync targets.
[**GetSyncTargets**](SyncServiceAPI.md#GetSyncTargets) | **Get** /Sync/Targets | Gets a list of available sync targets.
[**PostSyncByItemidStatus**](SyncServiceAPI.md#PostSyncByItemidStatus) | **Post** /Sync/{ItemId}/Status | Gets sync status for an item.
[**PostSyncData**](SyncServiceAPI.md#PostSyncData) | **Post** /Sync/Data | Syncs data between device and server
[**PostSyncItemsCancel**](SyncServiceAPI.md#PostSyncItemsCancel) | **Post** /Sync/Items/Cancel | Cancels items from a sync target
[**PostSyncJobitemsByIdEnable**](SyncServiceAPI.md#PostSyncJobitemsByIdEnable) | **Post** /Sync/JobItems/{Id}/Enable | Enables a cancelled or queued sync job item
[**PostSyncJobitemsByIdMarkforremoval**](SyncServiceAPI.md#PostSyncJobitemsByIdMarkforremoval) | **Post** /Sync/JobItems/{Id}/MarkForRemoval | Marks a job item for removal
[**PostSyncJobitemsByIdTransferred**](SyncServiceAPI.md#PostSyncJobitemsByIdTransferred) | **Post** /Sync/JobItems/{Id}/Transferred | Reports that a sync job item has successfully been transferred.
[**PostSyncJobitemsByIdUnmarkforremoval**](SyncServiceAPI.md#PostSyncJobitemsByIdUnmarkforremoval) | **Post** /Sync/JobItems/{Id}/UnmarkForRemoval | Unmarks a job item for removal
[**PostSyncJobs**](SyncServiceAPI.md#PostSyncJobs) | **Post** /Sync/Jobs | Gets sync jobs.
[**PostSyncJobsById**](SyncServiceAPI.md#PostSyncJobsById) | **Post** /Sync/Jobs/{Id} | Updates a sync job.
[**PostSyncOfflineactions**](SyncServiceAPI.md#PostSyncOfflineactions) | **Post** /Sync/OfflineActions | Reports an action that occurred while offline.



## DeleteSyncByTargetidItems

> DeleteSyncByTargetidItems(ctx, targetId).Execute()

Cancels items from a sync target



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
	targetId := "targetId_example" // string | TargetId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.DeleteSyncByTargetidItems(context.Background(), targetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.DeleteSyncByTargetidItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetId** | **string** | TargetId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyncByTargetidItemsRequest struct via the builder pattern


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


## DeleteSyncJobitemsById

> DeleteSyncJobitemsById(ctx, id).Execute()

Cancels a sync job item



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.DeleteSyncJobitemsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.DeleteSyncJobitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyncJobitemsByIdRequest struct via the builder pattern


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


## DeleteSyncJobsById

> DeleteSyncJobsById(ctx, id).Execute()

Cancels a sync job.



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.DeleteSyncJobsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.DeleteSyncJobsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyncJobsByIdRequest struct via the builder pattern


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


## GetSyncItemsReady

> []SyncModelSyncedItem GetSyncItemsReady(ctx).TargetId(targetId).Execute()

Gets ready to download sync items.



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
	targetId := "targetId_example" // string | TargetId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.GetSyncItemsReady(context.Background()).TargetId(targetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncItemsReady``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncItemsReady`: []SyncModelSyncedItem
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncItemsReady`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncItemsReadyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetId** | **string** | TargetId | 

### Return type

[**[]SyncModelSyncedItem**](SyncModelSyncedItem.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncJobitems

> QueryResultSyncModelSyncJobItem GetSyncJobitems(ctx).Execute()

Gets sync job items.



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
	resp, r, err := apiClient.SyncServiceAPI.GetSyncJobitems(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncJobitems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncJobitems`: QueryResultSyncModelSyncJobItem
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncJobitems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncJobitemsRequest struct via the builder pattern


### Return type

[**QueryResultSyncModelSyncJobItem**](QueryResultSyncModelSyncJobItem.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncJobitemsByIdAdditionalfiles

> GetSyncJobitemsByIdAdditionalfiles(ctx, id).Name(name).Execute()

Gets a sync job item file



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
	id := "id_example" // string | Id
	name := "name_example" // string | Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.GetSyncJobitemsByIdAdditionalfiles(context.Background(), id).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncJobitemsByIdAdditionalfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncJobitemsByIdAdditionalfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name | 

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


## GetSyncJobitemsByIdFile

> GetSyncJobitemsByIdFile(ctx, id).Execute()

Gets a sync job item file



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.GetSyncJobitemsByIdFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncJobitemsByIdFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncJobitemsByIdFileRequest struct via the builder pattern


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


## GetSyncJobs

> QueryResultSyncSyncJob GetSyncJobs(ctx).Execute()

Gets sync jobs.



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
	resp, r, err := apiClient.SyncServiceAPI.GetSyncJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncJobs`: QueryResultSyncSyncJob
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncJobsRequest struct via the builder pattern


### Return type

[**QueryResultSyncSyncJob**](QueryResultSyncSyncJob.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncJobsById

> SyncSyncJob GetSyncJobsById(ctx, id).Execute()

Gets a sync job.



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.GetSyncJobsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncJobsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncJobsById`: SyncSyncJob
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncJobsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncJobsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncSyncJob**](SyncSyncJob.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncOptions

> SyncModelSyncDialogOptions GetSyncOptions(ctx).UserId(userId).ItemIds(itemIds).ParentId(parentId).TargetId(targetId).Category(category).Execute()

Gets a list of available sync targets.



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
	userId := "userId_example" // string | UserId
	itemIds := "itemIds_example" // string | ItemIds (optional)
	parentId := "parentId_example" // string | ParentId (optional)
	targetId := "targetId_example" // string | TargetId (optional)
	category := "category_example" // string | Category (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.GetSyncOptions(context.Background()).UserId(userId).ItemIds(itemIds).ParentId(parentId).TargetId(targetId).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncOptions`: SyncModelSyncDialogOptions
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | UserId | 
 **itemIds** | **string** | ItemIds | 
 **parentId** | **string** | ParentId | 
 **targetId** | **string** | TargetId | 
 **category** | **string** | Category | 

### Return type

[**SyncModelSyncDialogOptions**](SyncModelSyncDialogOptions.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncTargets

> []SyncSyncTarget GetSyncTargets(ctx).UserId(userId).Execute()

Gets a list of available sync targets.



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
	userId := "userId_example" // string | UserId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.GetSyncTargets(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.GetSyncTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncTargets`: []SyncSyncTarget
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.GetSyncTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | UserId | 

### Return type

[**[]SyncSyncTarget**](SyncSyncTarget.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSyncByItemidStatus

> PostSyncByItemidStatus(ctx, itemId).SyncModelSyncedItemProgress(syncModelSyncedItemProgress).Execute()

Gets sync status for an item.



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
	itemId := "itemId_example" // string | 
	syncModelSyncedItemProgress := *openapiclient.NewSyncModelSyncedItemProgress() // SyncModelSyncedItemProgress | SyncedItemProgress: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncByItemidStatus(context.Background(), itemId).SyncModelSyncedItemProgress(syncModelSyncedItemProgress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncByItemidStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncByItemidStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncModelSyncedItemProgress** | [**SyncModelSyncedItemProgress**](SyncModelSyncedItemProgress.md) | SyncedItemProgress:  | 

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


## PostSyncData

> SyncModelSyncDataResponse PostSyncData(ctx).SyncModelSyncDataRequest(syncModelSyncDataRequest).Execute()

Syncs data between device and server



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
	syncModelSyncDataRequest := *openapiclient.NewSyncModelSyncDataRequest() // SyncModelSyncDataRequest | SyncDataRequest: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.PostSyncData(context.Background()).SyncModelSyncDataRequest(syncModelSyncDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSyncData`: SyncModelSyncDataResponse
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.PostSyncData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncModelSyncDataRequest** | [**SyncModelSyncDataRequest**](SyncModelSyncDataRequest.md) | SyncDataRequest:  | 

### Return type

[**SyncModelSyncDataResponse**](SyncModelSyncDataResponse.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSyncItemsCancel

> PostSyncItemsCancel(ctx).ItemIds(itemIds).Execute()

Cancels items from a sync target



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
	itemIds := "itemIds_example" // string | ItemIds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncItemsCancel(context.Background()).ItemIds(itemIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncItemsCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncItemsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemIds** | **string** | ItemIds | 

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


## PostSyncJobitemsByIdEnable

> PostSyncJobitemsByIdEnable(ctx, id).Execute()

Enables a cancelled or queued sync job item



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobitemsByIdEnable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobitemsByIdEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobitemsByIdEnableRequest struct via the builder pattern


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


## PostSyncJobitemsByIdMarkforremoval

> PostSyncJobitemsByIdMarkforremoval(ctx, id).Execute()

Marks a job item for removal



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobitemsByIdMarkforremoval(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobitemsByIdMarkforremoval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobitemsByIdMarkforremovalRequest struct via the builder pattern


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


## PostSyncJobitemsByIdTransferred

> PostSyncJobitemsByIdTransferred(ctx, id).Execute()

Reports that a sync job item has successfully been transferred.



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobitemsByIdTransferred(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobitemsByIdTransferred``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobitemsByIdTransferredRequest struct via the builder pattern


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


## PostSyncJobitemsByIdUnmarkforremoval

> PostSyncJobitemsByIdUnmarkforremoval(ctx, id).Execute()

Unmarks a job item for removal



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
	id := "id_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobitemsByIdUnmarkforremoval(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobitemsByIdUnmarkforremoval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobitemsByIdUnmarkforremovalRequest struct via the builder pattern


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


## PostSyncJobs

> SyncModelSyncJobCreationResult PostSyncJobs(ctx).SyncModelSyncJobRequest(syncModelSyncJobRequest).Execute()

Gets sync jobs.



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
	syncModelSyncJobRequest := *openapiclient.NewSyncModelSyncJobRequest() // SyncModelSyncJobRequest | SyncJobRequest: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncServiceAPI.PostSyncJobs(context.Background()).SyncModelSyncJobRequest(syncModelSyncJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSyncJobs`: SyncModelSyncJobCreationResult
	fmt.Fprintf(os.Stdout, "Response from `SyncServiceAPI.PostSyncJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncModelSyncJobRequest** | [**SyncModelSyncJobRequest**](SyncModelSyncJobRequest.md) | SyncJobRequest:  | 

### Return type

[**SyncModelSyncJobCreationResult**](SyncModelSyncJobCreationResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSyncJobsById

> PostSyncJobsById(ctx, id).SyncSyncJob(syncSyncJob).Execute()

Updates a sync job.



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
	id := int64(789) // int64 | 
	syncSyncJob := *openapiclient.NewSyncSyncJob() // SyncSyncJob | SyncJob: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncJobsById(context.Background(), id).SyncSyncJob(syncSyncJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncJobsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncJobsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncSyncJob** | [**SyncSyncJob**](SyncSyncJob.md) | SyncJob:  | 

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


## PostSyncOfflineactions

> PostSyncOfflineactions(ctx).UsersUserAction(usersUserAction).Execute()

Reports an action that occurred while offline.



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
	usersUserAction := []openapiclient.UsersUserAction{*openapiclient.NewUsersUserAction()} // []UsersUserAction | List`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncServiceAPI.PostSyncOfflineactions(context.Background()).UsersUserAction(usersUserAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncServiceAPI.PostSyncOfflineactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSyncOfflineactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersUserAction** | [**[]UsersUserAction**](UsersUserAction.md) | List&#x60;1:  | 

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

