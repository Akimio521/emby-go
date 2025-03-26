# \LocalizationServiceAPI

All URIs are relative to *https://home.ourflix.de:32865/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocalizationCountries**](LocalizationServiceAPI.md#GetLocalizationCountries) | **Get** /Localization/Countries | Gets known countries
[**GetLocalizationCultures**](LocalizationServiceAPI.md#GetLocalizationCultures) | **Get** /Localization/Cultures | Gets known cultures
[**GetLocalizationOptions**](LocalizationServiceAPI.md#GetLocalizationOptions) | **Get** /Localization/Options | Gets localization options
[**GetLocalizationParentalratings**](LocalizationServiceAPI.md#GetLocalizationParentalratings) | **Get** /Localization/ParentalRatings | Gets known parental ratings



## GetLocalizationCountries

> []GlobalizationCountryInfo GetLocalizationCountries(ctx).Execute()

Gets known countries



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
	resp, r, err := apiClient.LocalizationServiceAPI.GetLocalizationCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationServiceAPI.GetLocalizationCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationCountries`: []GlobalizationCountryInfo
	fmt.Fprintf(os.Stdout, "Response from `LocalizationServiceAPI.GetLocalizationCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationCountriesRequest struct via the builder pattern


### Return type

[**[]GlobalizationCountryInfo**](GlobalizationCountryInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalizationCultures

> []GlobalizationCultureDto GetLocalizationCultures(ctx).Execute()

Gets known cultures



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
	resp, r, err := apiClient.LocalizationServiceAPI.GetLocalizationCultures(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationServiceAPI.GetLocalizationCultures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationCultures`: []GlobalizationCultureDto
	fmt.Fprintf(os.Stdout, "Response from `LocalizationServiceAPI.GetLocalizationCultures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationCulturesRequest struct via the builder pattern


### Return type

[**[]GlobalizationCultureDto**](GlobalizationCultureDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalizationOptions

> []GlobalizationLocalizatonOption GetLocalizationOptions(ctx).Execute()

Gets localization options



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
	resp, r, err := apiClient.LocalizationServiceAPI.GetLocalizationOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationServiceAPI.GetLocalizationOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationOptions`: []GlobalizationLocalizatonOption
	fmt.Fprintf(os.Stdout, "Response from `LocalizationServiceAPI.GetLocalizationOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationOptionsRequest struct via the builder pattern


### Return type

[**[]GlobalizationLocalizatonOption**](GlobalizationLocalizatonOption.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalizationParentalratings

> []ParentalRating GetLocalizationParentalratings(ctx).Execute()

Gets known parental ratings



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
	resp, r, err := apiClient.LocalizationServiceAPI.GetLocalizationParentalratings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationServiceAPI.GetLocalizationParentalratings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationParentalratings`: []ParentalRating
	fmt.Fprintf(os.Stdout, "Response from `LocalizationServiceAPI.GetLocalizationParentalratings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationParentalratingsRequest struct via the builder pattern


### Return type

[**[]ParentalRating**](ParentalRating.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

