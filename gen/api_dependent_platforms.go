/*
CyberArkIAG

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 18a45ad8-77e8-4ecc-873e-787c6de10a60
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DependentPlatformsApiService DependentPlatformsApi service
type DependentPlatformsApiService service

type ApiDeleteDependentPlatformRequest struct {
	ctx _context.Context
	ApiService *DependentPlatformsApiService
	authorization *string
	contentType *string
	platformName string
}

// Session Authorization Token
func (r ApiDeleteDependentPlatformRequest) Authorization(authorization string) ApiDeleteDependentPlatformRequest {
	r.authorization = &authorization
	return r
}
func (r ApiDeleteDependentPlatformRequest) ContentType(contentType string) ApiDeleteDependentPlatformRequest {
	r.contentType = &contentType
	return r
}

func (r ApiDeleteDependentPlatformRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteDependentPlatformExecute(r)
}

/*
DeleteDependentPlatform Delete Dependent Platform

This method allows Vault Admins to delete a dependent platform.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param platformName
 @return ApiDeleteDependentPlatformRequest
*/
func (a *DependentPlatformsApiService) DeleteDependentPlatform(ctx _context.Context, platformName string) ApiDeleteDependentPlatformRequest {
	return ApiDeleteDependentPlatformRequest{
		ApiService: a,
		ctx: ctx,
		platformName: platformName,
	}
}

// Execute executes the request
func (a *DependentPlatformsApiService) DeleteDependentPlatformExecute(r ApiDeleteDependentPlatformRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DependentPlatformsApiService.DeleteDependentPlatform")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/passwordvault/api/platforms/dependents/{PlatformName}"
	localVarPath = strings.Replace(localVarPath, "{"+"PlatformName"+"}", _neturl.PathEscape(parameterToString(r.platformName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.authorization == nil {
		return nil, reportError("authorization is required and must be specified")
	}
	if r.contentType == nil {
		return nil, reportError("contentType is required and must be specified")
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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDuplicateDependentPlatformsRequest struct {
	ctx _context.Context
	ApiService *DependentPlatformsApiService
	authorization *string
	contentType *string
	platformName string
	body *string
}

// Session Authorization Token
func (r ApiDuplicateDependentPlatformsRequest) Authorization(authorization string) ApiDuplicateDependentPlatformsRequest {
	r.authorization = &authorization
	return r
}
func (r ApiDuplicateDependentPlatformsRequest) ContentType(contentType string) ApiDuplicateDependentPlatformsRequest {
	r.contentType = &contentType
	return r
}
// This method allows Vault Admins to duplicate dependent platforms.
func (r ApiDuplicateDependentPlatformsRequest) Body(body string) ApiDuplicateDependentPlatformsRequest {
	r.body = &body
	return r
}

func (r ApiDuplicateDependentPlatformsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DuplicateDependentPlatformsExecute(r)
}

/*
DuplicateDependentPlatforms Duplicate Dependent Platforms

This method allows Vault Admins to duplicate dependent platforms.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param platformName
 @return ApiDuplicateDependentPlatformsRequest
*/
func (a *DependentPlatformsApiService) DuplicateDependentPlatforms(ctx _context.Context, platformName string) ApiDuplicateDependentPlatformsRequest {
	return ApiDuplicateDependentPlatformsRequest{
		ApiService: a,
		ctx: ctx,
		platformName: platformName,
	}
}

// Execute executes the request
func (a *DependentPlatformsApiService) DuplicateDependentPlatformsExecute(r ApiDuplicateDependentPlatformsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DependentPlatformsApiService.DuplicateDependentPlatforms")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/passwordvault/api/platforms/dependent/{PlatformName}/duplicate"
	localVarPath = strings.Replace(localVarPath, "{"+"PlatformName"+"}", _neturl.PathEscape(parameterToString(r.platformName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.authorization == nil {
		return nil, reportError("authorization is required and must be specified")
	}
	if r.contentType == nil {
		return nil, reportError("contentType is required and must be specified")
	}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetDependentPlatformsRequest struct {
	ctx _context.Context
	ApiService *DependentPlatformsApiService
	search *float32
	authorization *string
	contentType *string
}

// Platform Name
func (r ApiGetDependentPlatformsRequest) Search(search float32) ApiGetDependentPlatformsRequest {
	r.search = &search
	return r
}
// Session Authorization Token
func (r ApiGetDependentPlatformsRequest) Authorization(authorization string) ApiGetDependentPlatformsRequest {
	r.authorization = &authorization
	return r
}
func (r ApiGetDependentPlatformsRequest) ContentType(contentType string) ApiGetDependentPlatformsRequest {
	r.contentType = &contentType
	return r
}

func (r ApiGetDependentPlatformsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetDependentPlatformsExecute(r)
}

/*
GetDependentPlatforms Get Dependent Platforms

This method allows Vault Admins to retrieve basic information about all existing dependent platforms.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDependentPlatformsRequest
*/
func (a *DependentPlatformsApiService) GetDependentPlatforms(ctx _context.Context) ApiGetDependentPlatformsRequest {
	return ApiGetDependentPlatformsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DependentPlatformsApiService) GetDependentPlatformsExecute(r ApiGetDependentPlatformsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DependentPlatformsApiService.GetDependentPlatforms")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/passwordvault/api/platforms/dependents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.search == nil {
		return nil, reportError("search is required and must be specified")
	}
	if r.authorization == nil {
		return nil, reportError("authorization is required and must be specified")
	}
	if r.contentType == nil {
		return nil, reportError("contentType is required and must be specified")
	}

	localVarQueryParams.Add("search", parameterToString(*r.search, ""))
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
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
