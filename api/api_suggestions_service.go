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


type SuggestionsServiceAPI interface {

	/*
	GetUsersByUseridSuggestions Gets items based on a query.

	No authentication required

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId
	@return ApiGetUsersByUseridSuggestionsRequest
	*/
	GetUsersByUseridSuggestions(ctx context.Context, userId string) ApiGetUsersByUseridSuggestionsRequest

	// GetUsersByUseridSuggestionsExecute executes the request
	//  @return QueryResultBaseItemDto
	GetUsersByUseridSuggestionsExecute(r ApiGetUsersByUseridSuggestionsRequest) (*QueryResultBaseItemDto, *http.Response, error)
}

// SuggestionsServiceAPIService SuggestionsServiceAPI service
type SuggestionsServiceAPIService service

type ApiGetUsersByUseridSuggestionsRequest struct {
	ctx context.Context
	ApiService SuggestionsServiceAPI
	userId string
}

func (r ApiGetUsersByUseridSuggestionsRequest) Execute() (*QueryResultBaseItemDto, *http.Response, error) {
	return r.ApiService.GetUsersByUseridSuggestionsExecute(r)
}

/*
GetUsersByUseridSuggestions Gets items based on a query.

No authentication required

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId
 @return ApiGetUsersByUseridSuggestionsRequest
*/
func (a *SuggestionsServiceAPIService) GetUsersByUseridSuggestions(ctx context.Context, userId string) ApiGetUsersByUseridSuggestionsRequest {
	return ApiGetUsersByUseridSuggestionsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return QueryResultBaseItemDto
func (a *SuggestionsServiceAPIService) GetUsersByUseridSuggestionsExecute(r ApiGetUsersByUseridSuggestionsRequest) (*QueryResultBaseItemDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueryResultBaseItemDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SuggestionsServiceAPIService.GetUsersByUseridSuggestions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Users/{UserId}/Suggestions"
	localVarPath = strings.Replace(localVarPath, "{"+"UserId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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
