# \PackageServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePackagesInstallingById**](PackageServiceAPI.md#DeletePackagesInstallingById) | **Delete** /Packages/Installing/{Id} | Cancels a package installation
[**GetPackages**](PackageServiceAPI.md#GetPackages) | **Get** /Packages | Gets available packages
[**GetPackagesByName**](PackageServiceAPI.md#GetPackagesByName) | **Get** /Packages/{Name} | Gets a package, by name or assembly guid
[**GetPackagesUpdates**](PackageServiceAPI.md#GetPackagesUpdates) | **Get** /Packages/Updates | Gets available package updates for currently installed packages
[**PostPackagesInstalledByName**](PackageServiceAPI.md#PostPackagesInstalledByName) | **Post** /Packages/Installed/{Name} | Installs a package



## DeletePackagesInstallingById

> DeletePackagesInstallingById(ctx, id).Execute()

Cancels a package installation



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
	id := "id_example" // string | Installation Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PackageServiceAPI.DeletePackagesInstallingById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageServiceAPI.DeletePackagesInstallingById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Installation Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePackagesInstallingByIdRequest struct via the builder pattern


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


## GetPackages

> []UpdatesPackageInfo GetPackages(ctx).PackageType(packageType).TargetSystems(targetSystems).IsPremium(isPremium).IsAdult(isAdult).Execute()

Gets available packages



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
	packageType := "packageType_example" // string | Optional package type filter (System/UserInstalled) (optional)
	targetSystems := "targetSystems_example" // string | Optional. Filter by target system type. Allows multiple, comma delimited. (optional)
	isPremium := true // bool | Optional. Filter by premium status (optional)
	isAdult := true // bool | Optional. Filter by package that contain adult content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageServiceAPI.GetPackages(context.Background()).PackageType(packageType).TargetSystems(targetSystems).IsPremium(isPremium).IsAdult(isAdult).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageServiceAPI.GetPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackages`: []UpdatesPackageInfo
	fmt.Fprintf(os.Stdout, "Response from `PackageServiceAPI.GetPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageType** | **string** | Optional package type filter (System/UserInstalled) | 
 **targetSystems** | **string** | Optional. Filter by target system type. Allows multiple, comma delimited. | 
 **isPremium** | **bool** | Optional. Filter by premium status | 
 **isAdult** | **bool** | Optional. Filter by package that contain adult content. | 

### Return type

[**[]UpdatesPackageInfo**](UpdatesPackageInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackagesByName

> UpdatesPackageInfo GetPackagesByName(ctx, name).AssemblyGuid(assemblyGuid).Execute()

Gets a package, by name or assembly guid



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
	name := "name_example" // string | The name of the package
	assemblyGuid := "assemblyGuid_example" // string | The guid of the associated assembly (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageServiceAPI.GetPackagesByName(context.Background(), name).AssemblyGuid(assemblyGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageServiceAPI.GetPackagesByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackagesByName`: UpdatesPackageInfo
	fmt.Fprintf(os.Stdout, "Response from `PackageServiceAPI.GetPackagesByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the package | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assemblyGuid** | **string** | The guid of the associated assembly | 

### Return type

[**UpdatesPackageInfo**](UpdatesPackageInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackagesUpdates

> []UpdatesPackageVersionInfo GetPackagesUpdates(ctx).PackageType(packageType).Execute()

Gets available package updates for currently installed packages



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
	packageType := "packageType_example" // string | Package type filter (System/UserInstalled)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageServiceAPI.GetPackagesUpdates(context.Background()).PackageType(packageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageServiceAPI.GetPackagesUpdates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackagesUpdates`: []UpdatesPackageVersionInfo
	fmt.Fprintf(os.Stdout, "Response from `PackageServiceAPI.GetPackagesUpdates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageType** | **string** | Package type filter (System/UserInstalled) | 

### Return type

[**[]UpdatesPackageVersionInfo**](UpdatesPackageVersionInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPackagesInstalledByName

> PostPackagesInstalledByName(ctx, name).AssemblyGuid(assemblyGuid).Version(version).UpdateClass(updateClass).Execute()

Installs a package



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
	name := "name_example" // string | Package name
	assemblyGuid := "assemblyGuid_example" // string | Guid of the associated assembly (optional)
	version := "version_example" // string | Optional version. Defaults to latest version. (optional)
	updateClass := "updateClass_example" // string | Optional update class (Dev, Beta, Release). Defaults to Release. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PackageServiceAPI.PostPackagesInstalledByName(context.Background(), name).AssemblyGuid(assemblyGuid).Version(version).UpdateClass(updateClass).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageServiceAPI.PostPackagesInstalledByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Package name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPackagesInstalledByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assemblyGuid** | **string** | Guid of the associated assembly | 
 **version** | **string** | Optional version. Defaults to latest version. | 
 **updateClass** | **string** | Optional update class (Dev, Beta, Release). Defaults to Release. | 

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

