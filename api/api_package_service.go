/*
Emby Server API

Explore the Emby Server API

API version: 4.1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type PackageServiceAPI interface {

	/*
	DeletePackagesInstallingById Cancels a package installation

	Requires authentication as administrator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Installation Id
	@return ApiDeletePackagesInstallingByIdRequest
	*/
	DeletePackagesInstallingById(ctx context.Context, id string) ApiDeletePackagesInstallingByIdRequest

	// DeletePackagesInstallingByIdExecute executes the request
	DeletePackagesInstallingByIdExecute(r ApiDeletePackagesInstallingByIdRequest) (*http.Response, error)

	/*
	GetPackages Gets available packages

	Requires authentication as user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPackagesRequest
	*/
	GetPackages(ctx context.Context) ApiGetPackagesRequest

	// GetPackagesExecute executes the request
	//  @return []UpdatesPackageInfo
	GetPackagesExecute(r ApiGetPackagesRequest) ([]UpdatesPackageInfo, *http.Response, error)

	/*
	GetPackagesByName Gets a package, by name or assembly guid

	Requires authentication as user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The name of the package
	@return ApiGetPackagesByNameRequest
	*/
	GetPackagesByName(ctx context.Context, name string) ApiGetPackagesByNameRequest

	// GetPackagesByNameExecute executes the request
	//  @return UpdatesPackageInfo
	GetPackagesByNameExecute(r ApiGetPackagesByNameRequest) (*UpdatesPackageInfo, *http.Response, error)

	/*
	GetPackagesUpdates Gets available package updates for currently installed packages

	Requires authentication as administrator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPackagesUpdatesRequest
	*/
	GetPackagesUpdates(ctx context.Context) ApiGetPackagesUpdatesRequest

	// GetPackagesUpdatesExecute executes the request
	//  @return []UpdatesPackageVersionInfo
	GetPackagesUpdatesExecute(r ApiGetPackagesUpdatesRequest) ([]UpdatesPackageVersionInfo, *http.Response, error)

	/*
	PostPackagesInstalledByName Installs a package

	Requires authentication as administrator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Package name
	@return ApiPostPackagesInstalledByNameRequest
	*/
	PostPackagesInstalledByName(ctx context.Context, name string) ApiPostPackagesInstalledByNameRequest

	// PostPackagesInstalledByNameExecute executes the request
	PostPackagesInstalledByNameExecute(r ApiPostPackagesInstalledByNameRequest) (*http.Response, error)
}

// PackageServiceAPIService PackageServiceAPI service
type PackageServiceAPIService service

type ApiDeletePackagesInstallingByIdRequest struct {
	ctx context.Context
	ApiService PackageServiceAPI
	id string
}

func (r ApiDeletePackagesInstallingByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePackagesInstallingByIdExecute(r)
}

/*
DeletePackagesInstallingById Cancels a package installation

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Installation Id
 @return ApiDeletePackagesInstallingByIdRequest
*/
func (a *PackageServiceAPIService) DeletePackagesInstallingById(ctx context.Context, id string) ApiDeletePackagesInstallingByIdRequest {
	return ApiDeletePackagesInstallingByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *PackageServiceAPIService) DeletePackagesInstallingByIdExecute(r ApiDeletePackagesInstallingByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageServiceAPIService.DeletePackagesInstallingById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Packages/Installing/{Id}"
	localVarPath = strings.Replace(localVarPath, "{"+"Id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetPackagesRequest struct {
	ctx context.Context
	ApiService PackageServiceAPI
	packageType *string
	targetSystems *string
	isPremium *bool
	isAdult *bool
}

// Optional package type filter (System/UserInstalled)
func (r ApiGetPackagesRequest) PackageType(packageType string) ApiGetPackagesRequest {
	r.packageType = &packageType
	return r
}

// Optional. Filter by target system type. Allows multiple, comma delimited.
func (r ApiGetPackagesRequest) TargetSystems(targetSystems string) ApiGetPackagesRequest {
	r.targetSystems = &targetSystems
	return r
}

// Optional. Filter by premium status
func (r ApiGetPackagesRequest) IsPremium(isPremium bool) ApiGetPackagesRequest {
	r.isPremium = &isPremium
	return r
}

// Optional. Filter by package that contain adult content.
func (r ApiGetPackagesRequest) IsAdult(isAdult bool) ApiGetPackagesRequest {
	r.isAdult = &isAdult
	return r
}

func (r ApiGetPackagesRequest) Execute() ([]UpdatesPackageInfo, *http.Response, error) {
	return r.ApiService.GetPackagesExecute(r)
}

/*
GetPackages Gets available packages

Requires authentication as user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPackagesRequest
*/
func (a *PackageServiceAPIService) GetPackages(ctx context.Context) ApiGetPackagesRequest {
	return ApiGetPackagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []UpdatesPackageInfo
func (a *PackageServiceAPIService) GetPackagesExecute(r ApiGetPackagesRequest) ([]UpdatesPackageInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []UpdatesPackageInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageServiceAPIService.GetPackages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Packages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.packageType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "PackageType", r.packageType, "form", "")
	}
	if r.targetSystems != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "TargetSystems", r.targetSystems, "form", "")
	}
	if r.isPremium != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsPremium", r.isPremium, "form", "")
	}
	if r.isAdult != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsAdult", r.isAdult, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPackagesByNameRequest struct {
	ctx context.Context
	ApiService PackageServiceAPI
	name string
	assemblyGuid *string
}

// The guid of the associated assembly
func (r ApiGetPackagesByNameRequest) AssemblyGuid(assemblyGuid string) ApiGetPackagesByNameRequest {
	r.assemblyGuid = &assemblyGuid
	return r
}

func (r ApiGetPackagesByNameRequest) Execute() (*UpdatesPackageInfo, *http.Response, error) {
	return r.ApiService.GetPackagesByNameExecute(r)
}

/*
GetPackagesByName Gets a package, by name or assembly guid

Requires authentication as user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name The name of the package
 @return ApiGetPackagesByNameRequest
*/
func (a *PackageServiceAPIService) GetPackagesByName(ctx context.Context, name string) ApiGetPackagesByNameRequest {
	return ApiGetPackagesByNameRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return UpdatesPackageInfo
func (a *PackageServiceAPIService) GetPackagesByNameExecute(r ApiGetPackagesByNameRequest) (*UpdatesPackageInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdatesPackageInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageServiceAPIService.GetPackagesByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Packages/{Name}"
	localVarPath = strings.Replace(localVarPath, "{"+"Name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.assemblyGuid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "AssemblyGuid", r.assemblyGuid, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPackagesUpdatesRequest struct {
	ctx context.Context
	ApiService PackageServiceAPI
	packageType *string
}

// Package type filter (System/UserInstalled)
func (r ApiGetPackagesUpdatesRequest) PackageType(packageType string) ApiGetPackagesUpdatesRequest {
	r.packageType = &packageType
	return r
}

func (r ApiGetPackagesUpdatesRequest) Execute() ([]UpdatesPackageVersionInfo, *http.Response, error) {
	return r.ApiService.GetPackagesUpdatesExecute(r)
}

/*
GetPackagesUpdates Gets available package updates for currently installed packages

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPackagesUpdatesRequest
*/
func (a *PackageServiceAPIService) GetPackagesUpdates(ctx context.Context) ApiGetPackagesUpdatesRequest {
	return ApiGetPackagesUpdatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []UpdatesPackageVersionInfo
func (a *PackageServiceAPIService) GetPackagesUpdatesExecute(r ApiGetPackagesUpdatesRequest) ([]UpdatesPackageVersionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []UpdatesPackageVersionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageServiceAPIService.GetPackagesUpdates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Packages/Updates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.packageType == nil {
		return localVarReturnValue, nil, reportError("packageType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "PackageType", r.packageType, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostPackagesInstalledByNameRequest struct {
	ctx context.Context
	ApiService PackageServiceAPI
	name string
	assemblyGuid *string
	version *string
	updateClass *string
}

// Guid of the associated assembly
func (r ApiPostPackagesInstalledByNameRequest) AssemblyGuid(assemblyGuid string) ApiPostPackagesInstalledByNameRequest {
	r.assemblyGuid = &assemblyGuid
	return r
}

// Optional version. Defaults to latest version.
func (r ApiPostPackagesInstalledByNameRequest) Version(version string) ApiPostPackagesInstalledByNameRequest {
	r.version = &version
	return r
}

// Optional update class (Dev, Beta, Release). Defaults to Release.
func (r ApiPostPackagesInstalledByNameRequest) UpdateClass(updateClass string) ApiPostPackagesInstalledByNameRequest {
	r.updateClass = &updateClass
	return r
}

func (r ApiPostPackagesInstalledByNameRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostPackagesInstalledByNameExecute(r)
}

/*
PostPackagesInstalledByName Installs a package

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Package name
 @return ApiPostPackagesInstalledByNameRequest
*/
func (a *PackageServiceAPIService) PostPackagesInstalledByName(ctx context.Context, name string) ApiPostPackagesInstalledByNameRequest {
	return ApiPostPackagesInstalledByNameRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
func (a *PackageServiceAPIService) PostPackagesInstalledByNameExecute(r ApiPostPackagesInstalledByNameRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageServiceAPIService.PostPackagesInstalledByName")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Packages/Installed/{Name}"
	localVarPath = strings.Replace(localVarPath, "{"+"Name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.assemblyGuid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "AssemblyGuid", r.assemblyGuid, "form", "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Version", r.version, "form", "")
	}
	if r.updateClass != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "UpdateClass", r.updateClass, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
