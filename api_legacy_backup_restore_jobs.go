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


// LegacyBackupRestoreJobsApiService LegacyBackupRestoreJobsApi service
type LegacyBackupRestoreJobsApiService service

type ApiCreateOneLegacyBackupRestoreJobRequest struct {
	ctx context.Context
	ApiService *LegacyBackupRestoreJobsApiService
	groupId string
	clusterName string
	apiRestoreJobView *ApiRestoreJobView
	envelope *bool
	pretty *bool
}

// Legacy backup to restore to one cluster in the specified project.
func (r ApiCreateOneLegacyBackupRestoreJobRequest) ApiRestoreJobView(apiRestoreJobView ApiRestoreJobView) ApiCreateOneLegacyBackupRestoreJobRequest {
	r.apiRestoreJobView = &apiRestoreJobView
	return r
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiCreateOneLegacyBackupRestoreJobRequest) Envelope(envelope bool) ApiCreateOneLegacyBackupRestoreJobRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiCreateOneLegacyBackupRestoreJobRequest) Pretty(pretty bool) ApiCreateOneLegacyBackupRestoreJobRequest {
	r.pretty = &pretty
	return r
}

func (r ApiCreateOneLegacyBackupRestoreJobRequest) Execute() (*PaginatedRestoreJobView, *http.Response, error) {
	return r.ApiService.CreateOneLegacyBackupRestoreJobExecute(r)
}

/*
CreateOneLegacyBackupRestoreJob Create One Legacy Backup Restore Job

Restores one legacy backup for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Atlas Admin role and an entry for the project access list.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
 @return ApiCreateOneLegacyBackupRestoreJobRequest

Deprecated
*/
func (a *LegacyBackupRestoreJobsApiService) CreateOneLegacyBackupRestoreJob(ctx context.Context, groupId string, clusterName string) ApiCreateOneLegacyBackupRestoreJobRequest {
	return ApiCreateOneLegacyBackupRestoreJobRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//  @return PaginatedRestoreJobView
// Deprecated
func (a *LegacyBackupRestoreJobsApiService) CreateOneLegacyBackupRestoreJobExecute(r ApiCreateOneLegacyBackupRestoreJobRequest) (*PaginatedRestoreJobView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRestoreJobView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupRestoreJobsApiService.CreateOneLegacyBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/restoreJobs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterToString(r.clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if r.apiRestoreJobView == nil {
		return localVarReturnValue, nil, reportError("apiRestoreJobView is required and must be specified")
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
	localVarPostBody = r.apiRestoreJobView
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

type ApiReturnAllLegacyBackupRestoreJobsRequest struct {
	ctx context.Context
	ApiService *LegacyBackupRestoreJobsApiService
	groupId string
	clusterName string
	envelope *bool
	includeCount *bool
	itemsPerPage *int32
	pageNum *int32
	pretty *bool
	batchId *string
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiReturnAllLegacyBackupRestoreJobsRequest) Envelope(envelope bool) ApiReturnAllLegacyBackupRestoreJobsRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ApiReturnAllLegacyBackupRestoreJobsRequest) IncludeCount(includeCount bool) ApiReturnAllLegacyBackupRestoreJobsRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ApiReturnAllLegacyBackupRestoreJobsRequest) ItemsPerPage(itemsPerPage int32) ApiReturnAllLegacyBackupRestoreJobsRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ApiReturnAllLegacyBackupRestoreJobsRequest) PageNum(pageNum int32) ApiReturnAllLegacyBackupRestoreJobsRequest {
	r.pageNum = &pageNum
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiReturnAllLegacyBackupRestoreJobsRequest) Pretty(pretty bool) ApiReturnAllLegacyBackupRestoreJobsRequest {
	r.pretty = &pretty
	return r
}

// Unique 24-hexadecimal digit string that identifies the batch of restore jobs to return. Timestamp in ISO 8601 date and time format in UTC when creating a restore job for a sharded cluster, Application creates a separate job for each shard, plus another for the config host. Each of these jobs comprise one batch. A restore job for a replica set can&#39;t be part of a batch.
func (r ApiReturnAllLegacyBackupRestoreJobsRequest) BatchId(batchId string) ApiReturnAllLegacyBackupRestoreJobsRequest {
	r.batchId = &batchId
	return r
}

func (r ApiReturnAllLegacyBackupRestoreJobsRequest) Execute() (*PaginatedRestoreJobView, *http.Response, error) {
	return r.ApiService.ReturnAllLegacyBackupRestoreJobsExecute(r)
}

/*
ReturnAllLegacyBackupRestoreJobs Return All Legacy Backup Restore Jobs

Returns all legacy backup restore jobs for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
 @return ApiReturnAllLegacyBackupRestoreJobsRequest

Deprecated
*/
func (a *LegacyBackupRestoreJobsApiService) ReturnAllLegacyBackupRestoreJobs(ctx context.Context, groupId string, clusterName string) ApiReturnAllLegacyBackupRestoreJobsRequest {
	return ApiReturnAllLegacyBackupRestoreJobsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//  @return PaginatedRestoreJobView
// Deprecated
func (a *LegacyBackupRestoreJobsApiService) ReturnAllLegacyBackupRestoreJobsExecute(r ApiReturnAllLegacyBackupRestoreJobsRequest) (*PaginatedRestoreJobView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRestoreJobView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupRestoreJobsApiService.ReturnAllLegacyBackupRestoreJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/restoreJobs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterToString(r.clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
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
	if r.batchId != nil {
		localVarQueryParams.Add("batchId", parameterToString(*r.batchId, ""))
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

type ApiReturnOneLegacyBackupRestoreJobRequest struct {
	ctx context.Context
	ApiService *LegacyBackupRestoreJobsApiService
	groupId string
	clusterName string
	jobId string
	envelope *bool
	pretty *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiReturnOneLegacyBackupRestoreJobRequest) Envelope(envelope bool) ApiReturnOneLegacyBackupRestoreJobRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiReturnOneLegacyBackupRestoreJobRequest) Pretty(pretty bool) ApiReturnOneLegacyBackupRestoreJobRequest {
	r.pretty = &pretty
	return r
}

func (r ApiReturnOneLegacyBackupRestoreJobRequest) Execute() (*ApiRestoreJobView, *http.Response, error) {
	return r.ApiService.ReturnOneLegacyBackupRestoreJobExecute(r)
}

/*
ReturnOneLegacyBackupRestoreJob Return One Legacy Backup Restore Job

Returns one legacy backup restore job for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
 @param jobId Unique 24-hexadecimal digit string that identifies the restore job.
 @return ApiReturnOneLegacyBackupRestoreJobRequest

Deprecated
*/
func (a *LegacyBackupRestoreJobsApiService) ReturnOneLegacyBackupRestoreJob(ctx context.Context, groupId string, clusterName string, jobId string) ApiReturnOneLegacyBackupRestoreJobRequest {
	return ApiReturnOneLegacyBackupRestoreJobRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		clusterName: clusterName,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return ApiRestoreJobView
// Deprecated
func (a *LegacyBackupRestoreJobsApiService) ReturnOneLegacyBackupRestoreJobExecute(r ApiReturnOneLegacyBackupRestoreJobRequest) (*ApiRestoreJobView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiRestoreJobView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupRestoreJobsApiService.ReturnOneLegacyBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/restoreJobs/{jobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.jobId) < 24 {
		return localVarReturnValue, nil, reportError("jobId must have at least 24 elements")
	}
	if strlen(r.jobId) > 24 {
		return localVarReturnValue, nil, reportError("jobId must have less than 24 elements")
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
