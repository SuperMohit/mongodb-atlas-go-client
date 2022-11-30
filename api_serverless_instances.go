/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ServerlessInstancesApiService ServerlessInstancesApi service
type ServerlessInstancesApiService service

type ApiCreateOneServerlessInstanceInOneProjectRequest struct {
	ctx context.Context
	ApiService *ServerlessInstancesApiService
	groupId string
	apiAtlasServerlessClusterDescriptionViewManual *ApiAtlasServerlessClusterDescriptionViewManual
	envelope *bool
	pretty *bool
}

// Create One Serverless Instance in One Project.
func (r ApiCreateOneServerlessInstanceInOneProjectRequest) ApiAtlasServerlessClusterDescriptionViewManual(apiAtlasServerlessClusterDescriptionViewManual ApiAtlasServerlessClusterDescriptionViewManual) ApiCreateOneServerlessInstanceInOneProjectRequest {
	r.apiAtlasServerlessClusterDescriptionViewManual = &apiAtlasServerlessClusterDescriptionViewManual
	return r
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiCreateOneServerlessInstanceInOneProjectRequest) Envelope(envelope bool) ApiCreateOneServerlessInstanceInOneProjectRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiCreateOneServerlessInstanceInOneProjectRequest) Pretty(pretty bool) ApiCreateOneServerlessInstanceInOneProjectRequest {
	r.pretty = &pretty
	return r
}

func (r ApiCreateOneServerlessInstanceInOneProjectRequest) Execute() (*ApiAtlasServerlessClusterDescriptionViewManual, *http.Response, error) {
	return r.ApiService.CreateOneServerlessInstanceInOneProjectExecute(r)
}

/*
CreateOneServerlessInstanceInOneProject Create One Serverless Instance in One Project

Creates one serverless instance in the specified project. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @return ApiCreateOneServerlessInstanceInOneProjectRequest
*/
func (a *ServerlessInstancesApiService) CreateOneServerlessInstanceInOneProject(ctx context.Context, groupId string) ApiCreateOneServerlessInstanceInOneProjectRequest {
	return ApiCreateOneServerlessInstanceInOneProjectRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return ApiAtlasServerlessClusterDescriptionViewManual
func (a *ServerlessInstancesApiService) CreateOneServerlessInstanceInOneProjectExecute(r ApiCreateOneServerlessInstanceInOneProjectRequest) (*ApiAtlasServerlessClusterDescriptionViewManual, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiAtlasServerlessClusterDescriptionViewManual
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessInstancesApiService.CreateOneServerlessInstanceInOneProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/serverless"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if r.apiAtlasServerlessClusterDescriptionViewManual == nil {
		return localVarReturnValue, nil, reportError("apiAtlasServerlessClusterDescriptionViewManual is required and must be specified")
	}

	if r.envelope != nil {
		localVarQueryParams.Add("envelope", parameterToString(*r.envelope, ""))
	}
	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
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
	localVarPostBody = r.apiAtlasServerlessClusterDescriptionViewManual
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type ApiRemoveOneServerlessInstanceFromOneProjectRequest struct {
	ctx context.Context
	ApiService *ServerlessInstancesApiService
	groupId string
	name string
	envelope *bool
	pretty *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiRemoveOneServerlessInstanceFromOneProjectRequest) Envelope(envelope bool) ApiRemoveOneServerlessInstanceFromOneProjectRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiRemoveOneServerlessInstanceFromOneProjectRequest) Pretty(pretty bool) ApiRemoveOneServerlessInstanceFromOneProjectRequest {
	r.pretty = &pretty
	return r
}

func (r ApiRemoveOneServerlessInstanceFromOneProjectRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveOneServerlessInstanceFromOneProjectExecute(r)
}

/*
RemoveOneServerlessInstanceFromOneProject Remove One Serverless Instance from One Project

Removes one serverless instance from the specified project. The serverless instance must have termination protection disabled in order to be deleted. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @param name Human-readable label that identifies the serverless instance.
 @return ApiRemoveOneServerlessInstanceFromOneProjectRequest
*/
func (a *ServerlessInstancesApiService) RemoveOneServerlessInstanceFromOneProject(ctx context.Context, groupId string, name string) ApiRemoveOneServerlessInstanceFromOneProjectRequest {
	return ApiRemoveOneServerlessInstanceFromOneProjectRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		name: name,
	}
}

// Execute executes the request
func (a *ServerlessInstancesApiService) RemoveOneServerlessInstanceFromOneProjectExecute(r ApiRemoveOneServerlessInstanceFromOneProjectRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessInstancesApiService.RemoveOneServerlessInstanceFromOneProject")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/serverless/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.name) < 1 {
		return nil, reportError("name must have at least 1 elements")
	}
	if strlen(r.name) > 64 {
		return nil, reportError("name must have less than 64 elements")
	}

	if r.envelope != nil {
		localVarQueryParams.Add("envelope", parameterToString(*r.envelope, ""))
	}
	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiReturnAllServerlessInstancesFromOneProjectRequest struct {
	ctx context.Context
	ApiService *ServerlessInstancesApiService
	groupId string
	envelope *bool
	includeCount *bool
	itemsPerPage *int32
	pageNum *int32
	pretty *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiReturnAllServerlessInstancesFromOneProjectRequest) Envelope(envelope bool) ApiReturnAllServerlessInstancesFromOneProjectRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ApiReturnAllServerlessInstancesFromOneProjectRequest) IncludeCount(includeCount bool) ApiReturnAllServerlessInstancesFromOneProjectRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ApiReturnAllServerlessInstancesFromOneProjectRequest) ItemsPerPage(itemsPerPage int32) ApiReturnAllServerlessInstancesFromOneProjectRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ApiReturnAllServerlessInstancesFromOneProjectRequest) PageNum(pageNum int32) ApiReturnAllServerlessInstancesFromOneProjectRequest {
	r.pageNum = &pageNum
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiReturnAllServerlessInstancesFromOneProjectRequest) Pretty(pretty bool) ApiReturnAllServerlessInstancesFromOneProjectRequest {
	r.pretty = &pretty
	return r
}

func (r ApiReturnAllServerlessInstancesFromOneProjectRequest) Execute() (*PaginatedServerlessInstancesView, *http.Response, error) {
	return r.ApiService.ReturnAllServerlessInstancesFromOneProjectExecute(r)
}

/*
ReturnAllServerlessInstancesFromOneProject Return All Serverless Instances from One Project

Returns details for all serverless instances in the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @return ApiReturnAllServerlessInstancesFromOneProjectRequest
*/
func (a *ServerlessInstancesApiService) ReturnAllServerlessInstancesFromOneProject(ctx context.Context, groupId string) ApiReturnAllServerlessInstancesFromOneProjectRequest {
	return ApiReturnAllServerlessInstancesFromOneProjectRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return PaginatedServerlessInstancesView
func (a *ServerlessInstancesApiService) ReturnAllServerlessInstancesFromOneProjectExecute(r ApiReturnAllServerlessInstancesFromOneProjectRequest) (*PaginatedServerlessInstancesView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedServerlessInstancesView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessInstancesApiService.ReturnAllServerlessInstancesFromOneProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/serverless"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}

	if r.envelope != nil {
		localVarQueryParams.Add("envelope", parameterToString(*r.envelope, ""))
	}
	if r.includeCount != nil {
		localVarQueryParams.Add("includeCount", parameterToString(*r.includeCount, ""))
	}
	if r.itemsPerPage != nil {
		localVarQueryParams.Add("itemsPerPage", parameterToString(*r.itemsPerPage, ""))
	}
	if r.pageNum != nil {
		localVarQueryParams.Add("pageNum", parameterToString(*r.pageNum, ""))
	}
	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type ApiReturnOneServerlessInstanceFromOneProjectRequest struct {
	ctx context.Context
	ApiService *ServerlessInstancesApiService
	groupId string
	name string
	envelope *bool
	pretty *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiReturnOneServerlessInstanceFromOneProjectRequest) Envelope(envelope bool) ApiReturnOneServerlessInstanceFromOneProjectRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiReturnOneServerlessInstanceFromOneProjectRequest) Pretty(pretty bool) ApiReturnOneServerlessInstanceFromOneProjectRequest {
	r.pretty = &pretty
	return r
}

func (r ApiReturnOneServerlessInstanceFromOneProjectRequest) Execute() (*ApiAtlasServerlessClusterDescriptionViewManual, *http.Response, error) {
	return r.ApiService.ReturnOneServerlessInstanceFromOneProjectExecute(r)
}

/*
ReturnOneServerlessInstanceFromOneProject Return One Serverless Instance from One Project

Returns details for one serverless instance in the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @param name Human-readable label that identifies the serverless instance.
 @return ApiReturnOneServerlessInstanceFromOneProjectRequest
*/
func (a *ServerlessInstancesApiService) ReturnOneServerlessInstanceFromOneProject(ctx context.Context, groupId string, name string) ApiReturnOneServerlessInstanceFromOneProjectRequest {
	return ApiReturnOneServerlessInstanceFromOneProjectRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		name: name,
	}
}

// Execute executes the request
//  @return ApiAtlasServerlessClusterDescriptionViewManual
func (a *ServerlessInstancesApiService) ReturnOneServerlessInstanceFromOneProjectExecute(r ApiReturnOneServerlessInstanceFromOneProjectRequest) (*ApiAtlasServerlessClusterDescriptionViewManual, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiAtlasServerlessClusterDescriptionViewManual
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessInstancesApiService.ReturnOneServerlessInstanceFromOneProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/serverless/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.name) < 1 {
		return localVarReturnValue, nil, reportError("name must have at least 1 elements")
	}
	if strlen(r.name) > 64 {
		return localVarReturnValue, nil, reportError("name must have less than 64 elements")
	}

	if r.envelope != nil {
		localVarQueryParams.Add("envelope", parameterToString(*r.envelope, ""))
	}
	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type ApiUpdateOneServerlessInstanceInOneProjectRequest struct {
	ctx context.Context
	ApiService *ServerlessInstancesApiService
	groupId string
	name string
	apiAtlasServerlessClusterDescriptionPatchView *ApiAtlasServerlessClusterDescriptionPatchView
	envelope *bool
	pretty *bool
}

// Update One Serverless Instance in One Project.
func (r ApiUpdateOneServerlessInstanceInOneProjectRequest) ApiAtlasServerlessClusterDescriptionPatchView(apiAtlasServerlessClusterDescriptionPatchView ApiAtlasServerlessClusterDescriptionPatchView) ApiUpdateOneServerlessInstanceInOneProjectRequest {
	r.apiAtlasServerlessClusterDescriptionPatchView = &apiAtlasServerlessClusterDescriptionPatchView
	return r
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiUpdateOneServerlessInstanceInOneProjectRequest) Envelope(envelope bool) ApiUpdateOneServerlessInstanceInOneProjectRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiUpdateOneServerlessInstanceInOneProjectRequest) Pretty(pretty bool) ApiUpdateOneServerlessInstanceInOneProjectRequest {
	r.pretty = &pretty
	return r
}

func (r ApiUpdateOneServerlessInstanceInOneProjectRequest) Execute() (*ApiAtlasServerlessClusterDescriptionPatchView, *http.Response, error) {
	return r.ApiService.UpdateOneServerlessInstanceInOneProjectExecute(r)
}

/*
UpdateOneServerlessInstanceInOneProject Update One Serverless Instance in One Project

Updates one serverless instance in the specified project. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @param name Human-readable label that identifies the serverless instance.
 @return ApiUpdateOneServerlessInstanceInOneProjectRequest
*/
func (a *ServerlessInstancesApiService) UpdateOneServerlessInstanceInOneProject(ctx context.Context, groupId string, name string) ApiUpdateOneServerlessInstanceInOneProjectRequest {
	return ApiUpdateOneServerlessInstanceInOneProjectRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		name: name,
	}
}

// Execute executes the request
//  @return ApiAtlasServerlessClusterDescriptionPatchView
func (a *ServerlessInstancesApiService) UpdateOneServerlessInstanceInOneProjectExecute(r ApiUpdateOneServerlessInstanceInOneProjectRequest) (*ApiAtlasServerlessClusterDescriptionPatchView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiAtlasServerlessClusterDescriptionPatchView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessInstancesApiService.UpdateOneServerlessInstanceInOneProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/serverless/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.name) < 1 {
		return localVarReturnValue, nil, reportError("name must have at least 1 elements")
	}
	if strlen(r.name) > 64 {
		return localVarReturnValue, nil, reportError("name must have less than 64 elements")
	}
	if r.apiAtlasServerlessClusterDescriptionPatchView == nil {
		return localVarReturnValue, nil, reportError("apiAtlasServerlessClusterDescriptionPatchView is required and must be specified")
	}

	if r.envelope != nil {
		localVarQueryParams.Add("envelope", parameterToString(*r.envelope, ""))
	}
	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
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
	localVarPostBody = r.apiAtlasServerlessClusterDescriptionPatchView
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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
