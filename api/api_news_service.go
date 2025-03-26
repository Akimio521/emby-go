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


type NewsServiceAPI interface {

	/*
	GetNewsProduct Gets the latest product news.

	No authentication required

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetNewsProductRequest
	*/
	GetNewsProduct(ctx context.Context) ApiGetNewsProductRequest

	// GetNewsProductExecute executes the request
	//  @return QueryResultNewsNewsItem
	GetNewsProductExecute(r ApiGetNewsProductRequest) (*QueryResultNewsNewsItem, *http.Response, error)
}

// NewsServiceAPIService NewsServiceAPI service
type NewsServiceAPIService service

type ApiGetNewsProductRequest struct {
	ctx context.Context
	ApiService NewsServiceAPI
	startIndex *int32
	limit *int32
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetNewsProductRequest) StartIndex(startIndex int32) ApiGetNewsProductRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return
func (r ApiGetNewsProductRequest) Limit(limit int32) ApiGetNewsProductRequest {
	r.limit = &limit
	return r
}

func (r ApiGetNewsProductRequest) Execute() (*QueryResultNewsNewsItem, *http.Response, error) {
	return r.ApiService.GetNewsProductExecute(r)
}

/*
GetNewsProduct Gets the latest product news.

No authentication required

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNewsProductRequest
*/
func (a *NewsServiceAPIService) GetNewsProduct(ctx context.Context) ApiGetNewsProductRequest {
	return ApiGetNewsProductRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QueryResultNewsNewsItem
func (a *NewsServiceAPIService) GetNewsProductExecute(r ApiGetNewsProductRequest) (*QueryResultNewsNewsItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueryResultNewsNewsItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NewsServiceAPIService.GetNewsProduct")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/News/Product"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
