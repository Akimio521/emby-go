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
)


type ChannelServiceAPI interface {

	/*
	GetChannels Gets available channels

	Requires authentication as user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetChannelsRequest
	*/
	GetChannels(ctx context.Context) ApiGetChannelsRequest

	// GetChannelsExecute executes the request
	//  @return QueryResultBaseItemDto
	GetChannelsExecute(r ApiGetChannelsRequest) (*QueryResultBaseItemDto, *http.Response, error)

	/*
	GetChannelsByIdItems Gets channel items

	Requires authentication as user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetChannelsByIdItemsRequest
	*/
	GetChannelsByIdItems(ctx context.Context) ApiGetChannelsByIdItemsRequest

	// GetChannelsByIdItemsExecute executes the request
	GetChannelsByIdItemsExecute(r ApiGetChannelsByIdItemsRequest) (*http.Response, error)

	/*
	GetChannelsItemsLatest Gets channel items

	Requires authentication as user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetChannelsItemsLatestRequest
	*/
	GetChannelsItemsLatest(ctx context.Context) ApiGetChannelsItemsLatestRequest

	// GetChannelsItemsLatestExecute executes the request
	GetChannelsItemsLatestExecute(r ApiGetChannelsItemsLatestRequest) (*http.Response, error)
}

// ChannelServiceAPIService ChannelServiceAPI service
type ChannelServiceAPIService service

type ApiGetChannelsRequest struct {
	ctx context.Context
	ApiService ChannelServiceAPI
	userId *string
	startIndex *int32
	limit *int32
}

// User Id
func (r ApiGetChannelsRequest) UserId(userId string) ApiGetChannelsRequest {
	r.userId = &userId
	return r
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetChannelsRequest) StartIndex(startIndex int32) ApiGetChannelsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return
func (r ApiGetChannelsRequest) Limit(limit int32) ApiGetChannelsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetChannelsRequest) Execute() (*QueryResultBaseItemDto, *http.Response, error) {
	return r.ApiService.GetChannelsExecute(r)
}

/*
GetChannels Gets available channels

Requires authentication as user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetChannelsRequest
*/
func (a *ChannelServiceAPIService) GetChannels(ctx context.Context) ApiGetChannelsRequest {
	return ApiGetChannelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QueryResultBaseItemDto
func (a *ChannelServiceAPIService) GetChannelsExecute(r ApiGetChannelsRequest) (*QueryResultBaseItemDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueryResultBaseItemDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelServiceAPIService.GetChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Channels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "UserId", r.userId, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "StartIndex", r.startIndex, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Limit", r.limit, "form", "")
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

type ApiGetChannelsByIdItemsRequest struct {
	ctx context.Context
	ApiService ChannelServiceAPI
}

func (r ApiGetChannelsByIdItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetChannelsByIdItemsExecute(r)
}

/*
GetChannelsByIdItems Gets channel items

Requires authentication as user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetChannelsByIdItemsRequest
*/
func (a *ChannelServiceAPIService) GetChannelsByIdItems(ctx context.Context) ApiGetChannelsByIdItemsRequest {
	return ApiGetChannelsByIdItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ChannelServiceAPIService) GetChannelsByIdItemsExecute(r ApiGetChannelsByIdItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelServiceAPIService.GetChannelsByIdItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Channels/{Id}/Items"

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

type ApiGetChannelsItemsLatestRequest struct {
	ctx context.Context
	ApiService ChannelServiceAPI
}

func (r ApiGetChannelsItemsLatestRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetChannelsItemsLatestExecute(r)
}

/*
GetChannelsItemsLatest Gets channel items

Requires authentication as user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetChannelsItemsLatestRequest
*/
func (a *ChannelServiceAPIService) GetChannelsItemsLatest(ctx context.Context) ApiGetChannelsItemsLatestRequest {
	return ApiGetChannelsItemsLatestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ChannelServiceAPIService) GetChannelsItemsLatestExecute(r ApiGetChannelsItemsLatestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelServiceAPIService.GetChannelsItemsLatest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Channels/Items/Latest"

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
