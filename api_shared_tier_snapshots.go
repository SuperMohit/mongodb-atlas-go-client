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


// SharedTierSnapshotsApiService SharedTierSnapshotsApi service
type SharedTierSnapshotsApiService service

type ApiDownloadOneM2OrM5ClusterSnapshotRequest struct {
	ctx context.Context
	ApiService *SharedTierSnapshotsApiService
	clusterName string
	groupId string
	apiAtlasTenantRestoreView *ApiAtlasTenantRestoreView
	envelope *bool
	pretty *bool
}

// Snapshot to be downloaded.
func (r ApiDownloadOneM2OrM5ClusterSnapshotRequest) ApiAtlasTenantRestoreView(apiAtlasTenantRestoreView ApiAtlasTenantRestoreView) ApiDownloadOneM2OrM5ClusterSnapshotRequest {
	r.apiAtlasTenantRestoreView = &apiAtlasTenantRestoreView
	return r
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiDownloadOneM2OrM5ClusterSnapshotRequest) Envelope(envelope bool) ApiDownloadOneM2OrM5ClusterSnapshotRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiDownloadOneM2OrM5ClusterSnapshotRequest) Pretty(pretty bool) ApiDownloadOneM2OrM5ClusterSnapshotRequest {
	r.pretty = &pretty
	return r
}

func (r ApiDownloadOneM2OrM5ClusterSnapshotRequest) Execute() (*ApiAtlasTenantRestoreView, *http.Response, error) {
	return r.ApiService.DownloadOneM2OrM5ClusterSnapshotExecute(r)
}

/*
DownloadOneM2OrM5ClusterSnapshot Download One M2 or M5 Cluster Snapshot

Requests one snapshot for the specified shared cluster. This resource returns a `snapshotURL` that you can use to download the snapshot. This `snapshotURL` remains active for four hours after you make the request. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterName Human-readable label that identifies the cluster.
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @return ApiDownloadOneM2OrM5ClusterSnapshotRequest
*/
func (a *SharedTierSnapshotsApiService) DownloadOneM2OrM5ClusterSnapshot(ctx context.Context, clusterName string, groupId string) ApiDownloadOneM2OrM5ClusterSnapshotRequest {
	return ApiDownloadOneM2OrM5ClusterSnapshotRequest{
		ApiService: a,
		ctx: ctx,
		clusterName: clusterName,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return ApiAtlasTenantRestoreView
func (a *SharedTierSnapshotsApiService) DownloadOneM2OrM5ClusterSnapshotExecute(r ApiDownloadOneM2OrM5ClusterSnapshotRequest) (*ApiAtlasTenantRestoreView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiAtlasTenantRestoreView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierSnapshotsApiService.DownloadOneM2OrM5ClusterSnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/tenant/download"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if r.apiAtlasTenantRestoreView == nil {
		return localVarReturnValue, nil, reportError("apiAtlasTenantRestoreView is required and must be specified")
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
	localVarPostBody = r.apiAtlasTenantRestoreView
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
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 403 {
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

type ApiReturnAllSnapshotsForOneM2OrM5ClusterRequest struct {
	ctx context.Context
	ApiService *SharedTierSnapshotsApiService
	groupId string
	clusterName string
	envelope *bool
	pretty *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiReturnAllSnapshotsForOneM2OrM5ClusterRequest) Envelope(envelope bool) ApiReturnAllSnapshotsForOneM2OrM5ClusterRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiReturnAllSnapshotsForOneM2OrM5ClusterRequest) Pretty(pretty bool) ApiReturnAllSnapshotsForOneM2OrM5ClusterRequest {
	r.pretty = &pretty
	return r
}

func (r ApiReturnAllSnapshotsForOneM2OrM5ClusterRequest) Execute() (*PaginatedTenantSnapshotView, *http.Response, error) {
	return r.ApiService.ReturnAllSnapshotsForOneM2OrM5ClusterExecute(r)
}

/*
ReturnAllSnapshotsForOneM2OrM5Cluster Return All Snapshots for One M2 or M5 Cluster

Returns details for all snapshots for the specified shared cluster. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @param clusterName Human-readable label that identifies the cluster.
 @return ApiReturnAllSnapshotsForOneM2OrM5ClusterRequest
*/
func (a *SharedTierSnapshotsApiService) ReturnAllSnapshotsForOneM2OrM5Cluster(ctx context.Context, groupId string, clusterName string) ApiReturnAllSnapshotsForOneM2OrM5ClusterRequest {
	return ApiReturnAllSnapshotsForOneM2OrM5ClusterRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//  @return PaginatedTenantSnapshotView
func (a *SharedTierSnapshotsApiService) ReturnAllSnapshotsForOneM2OrM5ClusterExecute(r ApiReturnAllSnapshotsForOneM2OrM5ClusterRequest) (*PaginatedTenantSnapshotView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedTenantSnapshotView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierSnapshotsApiService.ReturnAllSnapshotsForOneM2OrM5Cluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots"
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

type ApiReturnOneSnapshotForOneM2OrM5ClusterRequest struct {
	ctx context.Context
	ApiService *SharedTierSnapshotsApiService
	groupId string
	clusterName string
	snapshotId string
	envelope *bool
	pretty *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiReturnOneSnapshotForOneM2OrM5ClusterRequest) Envelope(envelope bool) ApiReturnOneSnapshotForOneM2OrM5ClusterRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiReturnOneSnapshotForOneM2OrM5ClusterRequest) Pretty(pretty bool) ApiReturnOneSnapshotForOneM2OrM5ClusterRequest {
	r.pretty = &pretty
	return r
}

func (r ApiReturnOneSnapshotForOneM2OrM5ClusterRequest) Execute() (*ApiAtlasTenantSnapshotView, *http.Response, error) {
	return r.ApiService.ReturnOneSnapshotForOneM2OrM5ClusterExecute(r)
}

/*
ReturnOneSnapshotForOneM2OrM5Cluster Return One Snapshot for One M2 or M5 Cluster

Returns details for one snapshot for the specified shared cluster. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @param clusterName Human-readable label that identifies the cluster.
 @param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
 @return ApiReturnOneSnapshotForOneM2OrM5ClusterRequest
*/
func (a *SharedTierSnapshotsApiService) ReturnOneSnapshotForOneM2OrM5Cluster(ctx context.Context, groupId string, clusterName string, snapshotId string) ApiReturnOneSnapshotForOneM2OrM5ClusterRequest {
	return ApiReturnOneSnapshotForOneM2OrM5ClusterRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		clusterName: clusterName,
		snapshotId: snapshotId,
	}
}

// Execute executes the request
//  @return ApiAtlasTenantSnapshotView
func (a *SharedTierSnapshotsApiService) ReturnOneSnapshotForOneM2OrM5ClusterExecute(r ApiReturnOneSnapshotForOneM2OrM5ClusterRequest) (*ApiAtlasTenantSnapshotView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiAtlasTenantSnapshotView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierSnapshotsApiService.ReturnOneSnapshotForOneM2OrM5Cluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterToString(r.snapshotId, "")), -1)

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
	if strlen(r.snapshotId) < 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have at least 24 elements")
	}
	if strlen(r.snapshotId) > 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have less than 24 elements")
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
