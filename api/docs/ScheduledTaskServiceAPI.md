# \ScheduledTaskServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteScheduledtasksRunningById**](ScheduledTaskServiceAPI.md#DeleteScheduledtasksRunningById) | **Delete** /ScheduledTasks/Running/{Id} | Stops a scheduled task
[**GetScheduledtasks**](ScheduledTaskServiceAPI.md#GetScheduledtasks) | **Get** /ScheduledTasks | Gets scheduled tasks
[**GetScheduledtasksById**](ScheduledTaskServiceAPI.md#GetScheduledtasksById) | **Get** /ScheduledTasks/{Id} | Gets a scheduled task, by Id
[**PostScheduledtasksByIdTriggers**](ScheduledTaskServiceAPI.md#PostScheduledtasksByIdTriggers) | **Post** /ScheduledTasks/{Id}/Triggers | Updates the triggers for a scheduled task
[**PostScheduledtasksRunningById**](ScheduledTaskServiceAPI.md#PostScheduledtasksRunningById) | **Post** /ScheduledTasks/Running/{Id} | Starts a scheduled task



## DeleteScheduledtasksRunningById

> DeleteScheduledtasksRunningById(ctx, id).Execute()

Stops a scheduled task



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScheduledTaskServiceAPI.DeleteScheduledtasksRunningById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledTaskServiceAPI.DeleteScheduledtasksRunningById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduledtasksRunningByIdRequest struct via the builder pattern


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


## GetScheduledtasks

> []TasksTaskInfo GetScheduledtasks(ctx).IsHidden(isHidden).IsEnabled(isEnabled).Execute()

Gets scheduled tasks



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
	isHidden := true // bool | Optional filter tasks that are hidden, or not. (optional)
	isEnabled := true // bool | Optional filter tasks that are enabled, or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduledTaskServiceAPI.GetScheduledtasks(context.Background()).IsHidden(isHidden).IsEnabled(isEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledTaskServiceAPI.GetScheduledtasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduledtasks`: []TasksTaskInfo
	fmt.Fprintf(os.Stdout, "Response from `ScheduledTaskServiceAPI.GetScheduledtasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduledtasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **bool** | Optional filter tasks that are hidden, or not. | 
 **isEnabled** | **bool** | Optional filter tasks that are enabled, or not. | 

### Return type

[**[]TasksTaskInfo**](TasksTaskInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduledtasksById

> TasksTaskInfo GetScheduledtasksById(ctx, id).Execute()

Gets a scheduled task, by Id



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduledTaskServiceAPI.GetScheduledtasksById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledTaskServiceAPI.GetScheduledtasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduledtasksById`: TasksTaskInfo
	fmt.Fprintf(os.Stdout, "Response from `ScheduledTaskServiceAPI.GetScheduledtasksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduledtasksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TasksTaskInfo**](TasksTaskInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostScheduledtasksByIdTriggers

> PostScheduledtasksByIdTriggers(ctx, id).TasksTaskTriggerInfo(tasksTaskTriggerInfo).Execute()

Updates the triggers for a scheduled task



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
	id := "id_example" // string | 
	tasksTaskTriggerInfo := []openapiclient.TasksTaskTriggerInfo{*openapiclient.NewTasksTaskTriggerInfo()} // []TasksTaskTriggerInfo | List`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScheduledTaskServiceAPI.PostScheduledtasksByIdTriggers(context.Background(), id).TasksTaskTriggerInfo(tasksTaskTriggerInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledTaskServiceAPI.PostScheduledtasksByIdTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostScheduledtasksByIdTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tasksTaskTriggerInfo** | [**[]TasksTaskTriggerInfo**](TasksTaskTriggerInfo.md) | List&#x60;1:  | 

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


## PostScheduledtasksRunningById

> PostScheduledtasksRunningById(ctx, id).Execute()

Starts a scheduled task



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScheduledTaskServiceAPI.PostScheduledtasksRunningById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledTaskServiceAPI.PostScheduledtasksRunningById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostScheduledtasksRunningByIdRequest struct via the builder pattern


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

