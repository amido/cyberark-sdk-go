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
)

// Linger please
var (
	_ _context.Context
)

// SharedLogonAuthenticationApiService SharedLogonAuthenticationApi service
type SharedLogonAuthenticationApiService service

type ApiLogoff1Request struct {
	ctx _context.Context
	ApiService *SharedLogonAuthenticationApiService
	contentType *string
	authorization *string
}

func (r ApiLogoff1Request) ContentType(contentType string) ApiLogoff1Request {
	r.contentType = &contentType
	return r
}
func (r ApiLogoff1Request) Authorization(authorization string) ApiLogoff1Request {
	r.authorization = &authorization
	return r
}

func (r ApiLogoff1Request) Execute() (*_nethttp.Response, error) {
	return r.ApiService.Logoff1Execute(r)
}

/*
Logoff1 Logoff

This method logs off the shared user and removes the Vault session.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLogoff1Request
*/
func (a *SharedLogonAuthenticationApiService) Logoff1(ctx _context.Context) ApiLogoff1Request {
	return ApiLogoff1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SharedLogonAuthenticationApiService) Logoff1Execute(r ApiLogoff1Request) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedLogonAuthenticationApiService.Logoff1")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/PasswordVault/WebServices/auth/Shared/RestfulAuthenticationService.svc/Logoff"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.contentType == nil {
		return nil, reportError("contentType is required and must be specified")
	}
	if r.authorization == nil {
		return nil, reportError("authorization is required and must be specified")
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
	localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
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

type ApiLogon0Request struct {
	ctx _context.Context
	ApiService *SharedLogonAuthenticationApiService
	contentType *string
	body *string
}

func (r ApiLogon0Request) ContentType(contentType string) ApiLogon0Request {
	r.contentType = &contentType
	return r
}
// Shared authentication is based on a user credential file that is stored in the PVWA web server. During shared authentication, only the user defined in the credential file can log on to the PVWA, but multiple users can use the logon token.  This type of authentication **requires** the application using the REST services to manage the users as the Vault can&#39;t identify which specific user performs each action.  Multiple concurrent connections can be created using the same token, without affecting each other.  The shared user is defined in a user credential file, whose location is specified in the WSCredentialFile parameter, in the appsettings section of the PVWAweb.config file:  &#x60;&#x60;&#x60; &lt;add key&#x3D;\&quot;WSCredentialFile\&quot; value&#x3D;\&quot;C:\\CyberArk\\Password Vault Web Access\\CredFiles\\WSUser.ini\&quot;/&gt; &#x60;&#x60;&#x60;  * Make sure that this user can access the PVWA interface. * Make sure the user only has the permissions in the Vault that they require.  For information about securing communication when using the SDK, refer to the following:  * [Securing application-to-REST communication](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/SDK/Securing%20Communication.htm) * [Configuring client authentication via certificates](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/SDK/Configuring%20Client%20Authentication%20via%20Client%20Certificates.htm)  This method authenticates to the Vault with a shared webservices user and returns a token that will be used in subsequent web services calls.  This is supported for CyberArk authentication only, and not for third party authentication.
func (r ApiLogon0Request) Body(body string) ApiLogon0Request {
	r.body = &body
	return r
}

func (r ApiLogon0Request) Execute() (*_nethttp.Response, error) {
	return r.ApiService.Logon0Execute(r)
}

/*
Logon0 Logon

Shared authentication is based on a user credential file that is stored in the PVWA web server. During shared authentication, only the user defined in the credential file can log on to the PVWA, but multiple users can use the logon token.

This type of authentication **requires** the application using the REST services to manage the users as the Vault can't identify which specific user performs each action.

Multiple concurrent connections can be created using the same token, without affecting each other.

The shared user is defined in a user credential file, whose location is specified in the WSCredentialFile parameter, in the appsettings section of the PVWAweb.config file:

```
<add key="WSCredentialFile" value="C:\CyberArk\Password Vault Web Access\CredFiles\WSUser.ini"/>
```

* Make sure that this user can access the PVWA interface.
* Make sure the user only has the permissions in the Vault that they require.

For information about securing communication when using the SDK, refer to the following:

* [Securing application-to-REST communication](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/SDK/Securing%20Communication.htm)
* [Configuring client authentication via certificates](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/SDK/Configuring%20Client%20Authentication%20via%20Client%20Certificates.htm)

This method authenticates to the Vault with a shared webservices user and returns a token that will be used in subsequent web services calls.

This is supported for CyberArk authentication only, and not for third party authentication.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLogon0Request
*/
func (a *SharedLogonAuthenticationApiService) Logon0(ctx _context.Context) ApiLogon0Request {
	return ApiLogon0Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SharedLogonAuthenticationApiService) Logon0Execute(r ApiLogon0Request) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedLogonAuthenticationApiService.Logon0")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/PasswordVault/WebServices/auth/Shared/RestfulAuthenticationService.svc/Logon"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
