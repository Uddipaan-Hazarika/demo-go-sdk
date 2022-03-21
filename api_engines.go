/*
Delphix API Gateway

Delphix API Gateway API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

// EnginesApiService EnginesApi service
type EnginesApiService service

type ApiConnectivityCheckRequest struct {
	ctx context.Context
	ApiService *EnginesApiService
	engineConnectivityCheckRequest *EngineConnectivityCheckRequest
}

// The api to check connectivity of engine and a remote host on given port.
func (r ApiConnectivityCheckRequest) EngineConnectivityCheckRequest(engineConnectivityCheckRequest EngineConnectivityCheckRequest) ApiConnectivityCheckRequest {
	r.engineConnectivityCheckRequest = &engineConnectivityCheckRequest
	return r
}

func (r ApiConnectivityCheckRequest) Execute() (*EngineConnectivityCheckResponse, *http.Response, error) {
	return r.ApiService.ConnectivityCheckExecute(r)
}

/*
ConnectivityCheck Checks connectivity between an engine and a remote host on a given port.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConnectivityCheckRequest
*/
func (a *EnginesApiService) ConnectivityCheck(ctx context.Context) ApiConnectivityCheckRequest {
	return ApiConnectivityCheckRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EngineConnectivityCheckResponse
func (a *EnginesApiService) ConnectivityCheckExecute(r ApiConnectivityCheckRequest) (*EngineConnectivityCheckResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EngineConnectivityCheckResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnginesApiService.ConnectivityCheck")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/engines/connnectivity/check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.engineConnectivityCheckRequest == nil {
		return localVarReturnValue, nil, reportError("engineConnectivityCheckRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.engineConnectivityCheckRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
