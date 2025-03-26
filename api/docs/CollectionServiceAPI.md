# \CollectionServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCollectionsByIdItems**](CollectionServiceAPI.md#DeleteCollectionsByIdItems) | **Delete** /Collections/{Id}/Items | Removes items from a collection
[**PostCollections**](CollectionServiceAPI.md#PostCollections) | **Post** /Collections | Creates a new collection
[**PostCollectionsByIdItems**](CollectionServiceAPI.md#PostCollectionsByIdItems) | **Post** /Collections/{Id}/Items | Adds items to a collection



## DeleteCollectionsByIdItems

> DeleteCollectionsByIdItems(ctx, id).Ids(ids).Execute()

Removes items from a collection



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
	ids := "ids_example" // string | Item id, comma delimited
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionServiceAPI.DeleteCollectionsByIdItems(context.Background(), id).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionServiceAPI.DeleteCollectionsByIdItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCollectionsByIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Item id, comma delimited | 


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


## PostCollections

> CollectionsCollectionCreationResult PostCollections(ctx).IsLocked(isLocked).Name(name).ParentId(parentId).Ids(ids).Execute()

Creates a new collection



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
	isLocked := true // bool | Whether or not to lock the new collection. (optional)
	name := "name_example" // string | The name of the new collection. (optional)
	parentId := "parentId_example" // string | Optional - create the collection within a specific folder (optional)
	ids := "ids_example" // string | Item Ids to add to the collection (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionServiceAPI.PostCollections(context.Background()).IsLocked(isLocked).Name(name).ParentId(parentId).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionServiceAPI.PostCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCollections`: CollectionsCollectionCreationResult
	fmt.Fprintf(os.Stdout, "Response from `CollectionServiceAPI.PostCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isLocked** | **bool** | Whether or not to lock the new collection. | 
 **name** | **string** | The name of the new collection. | 
 **parentId** | **string** | Optional - create the collection within a specific folder | 
 **ids** | **string** | Item Ids to add to the collection | 

### Return type

[**CollectionsCollectionCreationResult**](CollectionsCollectionCreationResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCollectionsByIdItems

> PostCollectionsByIdItems(ctx, id).Ids(ids).Execute()

Adds items to a collection



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
	ids := "ids_example" // string | Item id, comma delimited
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionServiceAPI.PostCollectionsByIdItems(context.Background(), id).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionServiceAPI.PostCollectionsByIdItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostCollectionsByIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Item id, comma delimited | 


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

