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


type DlnaServiceAPI interface {

	/*
	DeleteDlnaProfilesById Deletes a profile

	Requires authentication as administrator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Profile Id
	@return ApiDeleteDlnaProfilesByIdRequest
	*/
	DeleteDlnaProfilesById(ctx context.Context, id string) ApiDeleteDlnaProfilesByIdRequest

	// DeleteDlnaProfilesByIdExecute executes the request
	DeleteDlnaProfilesByIdExecute(r ApiDeleteDlnaProfilesByIdRequest) (*http.Response, error)

	/*
	GetDlnaProfileinfos Gets a list of profiles

	Requires authentication as administrator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDlnaProfileinfosRequest
	*/
	GetDlnaProfileinfos(ctx context.Context) ApiGetDlnaProfileinfosRequest

	// GetDlnaProfileinfosExecute executes the request
	//  @return []DlnaDeviceProfileInfo
	GetDlnaProfileinfosExecute(r ApiGetDlnaProfileinfosRequest) ([]DlnaDeviceProfileInfo, *http.Response, error)

	/*
	GetDlnaProfilesById Gets a single profile

	Requires authentication as administrator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Profile Id
	@return ApiGetDlnaProfilesByIdRequest
	*/
	GetDlnaProfilesById(ctx context.Context, id string) ApiGetDlnaProfilesByIdRequest

	// GetDlnaProfilesByIdExecute executes the request
	//  @return DlnaDeviceProfile
	GetDlnaProfilesByIdExecute(r ApiGetDlnaProfilesByIdRequest) (*DlnaDeviceProfile, *http.Response, error)

	/*
	GetDlnaProfilesDefault Gets the default profile

	Requires authentication as administrator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDlnaProfilesDefaultRequest
	*/
	GetDlnaProfilesDefault(ctx context.Context) ApiGetDlnaProfilesDefaultRequest

	// GetDlnaProfilesDefaultExecute executes the request
	//  @return DlnaDeviceProfile
	GetDlnaProfilesDefaultExecute(r ApiGetDlnaProfilesDefaultRequest) (*DlnaDeviceProfile, *http.Response, error)

	/*
	PostDlnaProfiles Creates a profile

	Requires authentication as administrator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostDlnaProfilesRequest
	*/
	PostDlnaProfiles(ctx context.Context) ApiPostDlnaProfilesRequest

	// PostDlnaProfilesExecute executes the request
	PostDlnaProfilesExecute(r ApiPostDlnaProfilesRequest) (*http.Response, error)

	/*
	PostDlnaProfilesById Updates a profile

	Requires authentication as administrator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiPostDlnaProfilesByIdRequest
	*/
	PostDlnaProfilesById(ctx context.Context, id string) ApiPostDlnaProfilesByIdRequest

	// PostDlnaProfilesByIdExecute executes the request
	PostDlnaProfilesByIdExecute(r ApiPostDlnaProfilesByIdRequest) (*http.Response, error)
}

// DlnaServiceAPIService DlnaServiceAPI service
type DlnaServiceAPIService service

type ApiDeleteDlnaProfilesByIdRequest struct {
	ctx context.Context
	ApiService DlnaServiceAPI
	id string
}

func (r ApiDeleteDlnaProfilesByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDlnaProfilesByIdExecute(r)
}

/*
DeleteDlnaProfilesById Deletes a profile

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Profile Id
 @return ApiDeleteDlnaProfilesByIdRequest
*/
func (a *DlnaServiceAPIService) DeleteDlnaProfilesById(ctx context.Context, id string) ApiDeleteDlnaProfilesByIdRequest {
	return ApiDeleteDlnaProfilesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *DlnaServiceAPIService) DeleteDlnaProfilesByIdExecute(r ApiDeleteDlnaProfilesByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DlnaServiceAPIService.DeleteDlnaProfilesById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Dlna/Profiles/{Id}"
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

type ApiGetDlnaProfileinfosRequest struct {
	ctx context.Context
	ApiService DlnaServiceAPI
}

func (r ApiGetDlnaProfileinfosRequest) Execute() ([]DlnaDeviceProfileInfo, *http.Response, error) {
	return r.ApiService.GetDlnaProfileinfosExecute(r)
}

/*
GetDlnaProfileinfos Gets a list of profiles

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDlnaProfileinfosRequest
*/
func (a *DlnaServiceAPIService) GetDlnaProfileinfos(ctx context.Context) ApiGetDlnaProfileinfosRequest {
	return ApiGetDlnaProfileinfosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []DlnaDeviceProfileInfo
func (a *DlnaServiceAPIService) GetDlnaProfileinfosExecute(r ApiGetDlnaProfileinfosRequest) ([]DlnaDeviceProfileInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DlnaDeviceProfileInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DlnaServiceAPIService.GetDlnaProfileinfos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Dlna/ProfileInfos"

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

type ApiGetDlnaProfilesByIdRequest struct {
	ctx context.Context
	ApiService DlnaServiceAPI
	id string
}

func (r ApiGetDlnaProfilesByIdRequest) Execute() (*DlnaDeviceProfile, *http.Response, error) {
	return r.ApiService.GetDlnaProfilesByIdExecute(r)
}

/*
GetDlnaProfilesById Gets a single profile

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Profile Id
 @return ApiGetDlnaProfilesByIdRequest
*/
func (a *DlnaServiceAPIService) GetDlnaProfilesById(ctx context.Context, id string) ApiGetDlnaProfilesByIdRequest {
	return ApiGetDlnaProfilesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return DlnaDeviceProfile
func (a *DlnaServiceAPIService) GetDlnaProfilesByIdExecute(r ApiGetDlnaProfilesByIdRequest) (*DlnaDeviceProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DlnaDeviceProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DlnaServiceAPIService.GetDlnaProfilesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Dlna/Profiles/{Id}"
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

type ApiGetDlnaProfilesDefaultRequest struct {
	ctx context.Context
	ApiService DlnaServiceAPI
}

func (r ApiGetDlnaProfilesDefaultRequest) Execute() (*DlnaDeviceProfile, *http.Response, error) {
	return r.ApiService.GetDlnaProfilesDefaultExecute(r)
}

/*
GetDlnaProfilesDefault Gets the default profile

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDlnaProfilesDefaultRequest
*/
func (a *DlnaServiceAPIService) GetDlnaProfilesDefault(ctx context.Context) ApiGetDlnaProfilesDefaultRequest {
	return ApiGetDlnaProfilesDefaultRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DlnaDeviceProfile
func (a *DlnaServiceAPIService) GetDlnaProfilesDefaultExecute(r ApiGetDlnaProfilesDefaultRequest) (*DlnaDeviceProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DlnaDeviceProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DlnaServiceAPIService.GetDlnaProfilesDefault")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Dlna/Profiles/Default"

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

type ApiPostDlnaProfilesRequest struct {
	ctx context.Context
	ApiService DlnaServiceAPI
	dlnaDeviceProfile *DlnaDeviceProfile
}

// DeviceProfile: 
func (r ApiPostDlnaProfilesRequest) DlnaDeviceProfile(dlnaDeviceProfile DlnaDeviceProfile) ApiPostDlnaProfilesRequest {
	r.dlnaDeviceProfile = &dlnaDeviceProfile
	return r
}

func (r ApiPostDlnaProfilesRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostDlnaProfilesExecute(r)
}

/*
PostDlnaProfiles Creates a profile

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostDlnaProfilesRequest
*/
func (a *DlnaServiceAPIService) PostDlnaProfiles(ctx context.Context) ApiPostDlnaProfilesRequest {
	return ApiPostDlnaProfilesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DlnaServiceAPIService) PostDlnaProfilesExecute(r ApiPostDlnaProfilesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DlnaServiceAPIService.PostDlnaProfiles")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Dlna/Profiles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dlnaDeviceProfile == nil {
		return nil, reportError("dlnaDeviceProfile is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

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
	// body params
	localVarPostBody = r.dlnaDeviceProfile
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

type ApiPostDlnaProfilesByIdRequest struct {
	ctx context.Context
	ApiService DlnaServiceAPI
	id string
	dlnaDeviceProfile *DlnaDeviceProfile
}

// DeviceProfile: 
func (r ApiPostDlnaProfilesByIdRequest) DlnaDeviceProfile(dlnaDeviceProfile DlnaDeviceProfile) ApiPostDlnaProfilesByIdRequest {
	r.dlnaDeviceProfile = &dlnaDeviceProfile
	return r
}

func (r ApiPostDlnaProfilesByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostDlnaProfilesByIdExecute(r)
}

/*
PostDlnaProfilesById Updates a profile

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiPostDlnaProfilesByIdRequest
*/
func (a *DlnaServiceAPIService) PostDlnaProfilesById(ctx context.Context, id string) ApiPostDlnaProfilesByIdRequest {
	return ApiPostDlnaProfilesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *DlnaServiceAPIService) PostDlnaProfilesByIdExecute(r ApiPostDlnaProfilesByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DlnaServiceAPIService.PostDlnaProfilesById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Dlna/Profiles/{Id}"
	localVarPath = strings.Replace(localVarPath, "{"+"Id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dlnaDeviceProfile == nil {
		return nil, reportError("dlnaDeviceProfile is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

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
	// body params
	localVarPostBody = r.dlnaDeviceProfile
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
