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


type EncodingInfoServiceAPI interface {

	/*
	GetEncodingCodecconfigurationDefaults Gets default codec configurations

	Requires authentication as user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetEncodingCodecconfigurationDefaultsRequest
	*/
	GetEncodingCodecconfigurationDefaults(ctx context.Context) ApiGetEncodingCodecconfigurationDefaultsRequest

	// GetEncodingCodecconfigurationDefaultsExecute executes the request
	//  @return []ConfigurationCodecConfiguration
	GetEncodingCodecconfigurationDefaultsExecute(r ApiGetEncodingCodecconfigurationDefaultsRequest) ([]ConfigurationCodecConfiguration, *http.Response, error)

	/*
	GetEncodingCodecinformationVideo Gets details about available video encoders and decoders

	Requires authentication as user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetEncodingCodecinformationVideoRequest
	*/
	GetEncodingCodecinformationVideo(ctx context.Context) ApiGetEncodingCodecinformationVideoRequest

	// GetEncodingCodecinformationVideoExecute executes the request
	//  @return []MediaEncodingCodecsVideoCodecsVideoCodecBase
	GetEncodingCodecinformationVideoExecute(r ApiGetEncodingCodecinformationVideoRequest) ([]MediaEncodingCodecsVideoCodecsVideoCodecBase, *http.Response, error)
}

// EncodingInfoServiceAPIService EncodingInfoServiceAPI service
type EncodingInfoServiceAPIService service

type ApiGetEncodingCodecconfigurationDefaultsRequest struct {
	ctx context.Context
	ApiService EncodingInfoServiceAPI
}

func (r ApiGetEncodingCodecconfigurationDefaultsRequest) Execute() ([]ConfigurationCodecConfiguration, *http.Response, error) {
	return r.ApiService.GetEncodingCodecconfigurationDefaultsExecute(r)
}

/*
GetEncodingCodecconfigurationDefaults Gets default codec configurations

Requires authentication as user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEncodingCodecconfigurationDefaultsRequest
*/
func (a *EncodingInfoServiceAPIService) GetEncodingCodecconfigurationDefaults(ctx context.Context) ApiGetEncodingCodecconfigurationDefaultsRequest {
	return ApiGetEncodingCodecconfigurationDefaultsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ConfigurationCodecConfiguration
func (a *EncodingInfoServiceAPIService) GetEncodingCodecconfigurationDefaultsExecute(r ApiGetEncodingCodecconfigurationDefaultsRequest) ([]ConfigurationCodecConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ConfigurationCodecConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EncodingInfoServiceAPIService.GetEncodingCodecconfigurationDefaults")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Encoding/CodecConfiguration/Defaults"

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

type ApiGetEncodingCodecinformationVideoRequest struct {
	ctx context.Context
	ApiService EncodingInfoServiceAPI
}

func (r ApiGetEncodingCodecinformationVideoRequest) Execute() ([]MediaEncodingCodecsVideoCodecsVideoCodecBase, *http.Response, error) {
	return r.ApiService.GetEncodingCodecinformationVideoExecute(r)
}

/*
GetEncodingCodecinformationVideo Gets details about available video encoders and decoders

Requires authentication as user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEncodingCodecinformationVideoRequest
*/
func (a *EncodingInfoServiceAPIService) GetEncodingCodecinformationVideo(ctx context.Context) ApiGetEncodingCodecinformationVideoRequest {
	return ApiGetEncodingCodecinformationVideoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MediaEncodingCodecsVideoCodecsVideoCodecBase
func (a *EncodingInfoServiceAPIService) GetEncodingCodecinformationVideoExecute(r ApiGetEncodingCodecinformationVideoRequest) ([]MediaEncodingCodecsVideoCodecsVideoCodecBase, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MediaEncodingCodecsVideoCodecsVideoCodecBase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EncodingInfoServiceAPIService.GetEncodingCodecinformationVideo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Encoding/CodecInformation/Video"

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
