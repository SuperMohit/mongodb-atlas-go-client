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
	"time"
	"reflect"
)


// EventsApiService EventsApi service
type EventsApiService service

type ApiReturnAllEventsFromOneOrganizationRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	orgId string
	envelope *bool
	includeCount *bool
	itemsPerPage *int32
	pageNum *int32
	pretty *bool
	eventType *string
	includeRaw *bool
	maxDate *time.Time
	minDate *time.Time
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiReturnAllEventsFromOneOrganizationRequest) Envelope(envelope bool) ApiReturnAllEventsFromOneOrganizationRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ApiReturnAllEventsFromOneOrganizationRequest) IncludeCount(includeCount bool) ApiReturnAllEventsFromOneOrganizationRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ApiReturnAllEventsFromOneOrganizationRequest) ItemsPerPage(itemsPerPage int32) ApiReturnAllEventsFromOneOrganizationRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ApiReturnAllEventsFromOneOrganizationRequest) PageNum(pageNum int32) ApiReturnAllEventsFromOneOrganizationRequest {
	r.pageNum = &pageNum
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiReturnAllEventsFromOneOrganizationRequest) Pretty(pretty bool) ApiReturnAllEventsFromOneOrganizationRequest {
	r.pretty = &pretty
	return r
}

// Category of incident recorded at this moment in time.
func (r ApiReturnAllEventsFromOneOrganizationRequest) EventType(eventType string) ApiReturnAllEventsFromOneOrganizationRequest {
	r.eventType = &eventType
	return r
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r ApiReturnAllEventsFromOneOrganizationRequest) IncludeRaw(includeRaw bool) ApiReturnAllEventsFromOneOrganizationRequest {
	r.includeRaw = &includeRaw
	return r
}

// Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC.
func (r ApiReturnAllEventsFromOneOrganizationRequest) MaxDate(maxDate time.Time) ApiReturnAllEventsFromOneOrganizationRequest {
	r.maxDate = &maxDate
	return r
}

// Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC.
func (r ApiReturnAllEventsFromOneOrganizationRequest) MinDate(minDate time.Time) ApiReturnAllEventsFromOneOrganizationRequest {
	r.minDate = &minDate
	return r
}

func (r ApiReturnAllEventsFromOneOrganizationRequest) Execute() ([]ApiEventView, *http.Response, error) {
	return r.ApiService.ReturnAllEventsFromOneOrganizationExecute(r)
}

/*
ReturnAllEventsFromOneOrganization Return All Events from One Organization

Returns all events for the specified organization. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-orgs-get-all/).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
 @return ApiReturnAllEventsFromOneOrganizationRequest
*/
func (a *EventsApiService) ReturnAllEventsFromOneOrganization(ctx context.Context, orgId string) ApiReturnAllEventsFromOneOrganizationRequest {
	return ApiReturnAllEventsFromOneOrganizationRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return []ApiEventView
func (a *EventsApiService) ReturnAllEventsFromOneOrganizationExecute(r ApiReturnAllEventsFromOneOrganizationRequest) ([]ApiEventView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiEventView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ReturnAllEventsFromOneOrganization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/orgs/{orgId}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterToString(r.orgId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
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
	if r.eventType != nil {
		localVarQueryParams.Add("eventType", parameterToString(*r.eventType, ""))
	}
	if r.includeRaw != nil {
		localVarQueryParams.Add("includeRaw", parameterToString(*r.includeRaw, ""))
	}
	if r.maxDate != nil {
		localVarQueryParams.Add("maxDate", parameterToString(*r.maxDate, ""))
	}
	if r.minDate != nil {
		localVarQueryParams.Add("minDate", parameterToString(*r.minDate, ""))
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

type ApiReturnAllEventsFromOneProjectRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	groupId string
	envelope *bool
	includeCount *bool
	itemsPerPage *int32
	pageNum *int32
	pretty *bool
	clusterNames *[]string
	eventType *string
	includeRaw *bool
	maxDate *time.Time
	minDate *time.Time
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiReturnAllEventsFromOneProjectRequest) Envelope(envelope bool) ApiReturnAllEventsFromOneProjectRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ApiReturnAllEventsFromOneProjectRequest) IncludeCount(includeCount bool) ApiReturnAllEventsFromOneProjectRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ApiReturnAllEventsFromOneProjectRequest) ItemsPerPage(itemsPerPage int32) ApiReturnAllEventsFromOneProjectRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ApiReturnAllEventsFromOneProjectRequest) PageNum(pageNum int32) ApiReturnAllEventsFromOneProjectRequest {
	r.pageNum = &pageNum
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiReturnAllEventsFromOneProjectRequest) Pretty(pretty bool) ApiReturnAllEventsFromOneProjectRequest {
	r.pretty = &pretty
	return r
}

// Human-readable label that identifies the cluster.
func (r ApiReturnAllEventsFromOneProjectRequest) ClusterNames(clusterNames []string) ApiReturnAllEventsFromOneProjectRequest {
	r.clusterNames = &clusterNames
	return r
}

// Category of incident recorded at this moment in time.
func (r ApiReturnAllEventsFromOneProjectRequest) EventType(eventType string) ApiReturnAllEventsFromOneProjectRequest {
	r.eventType = &eventType
	return r
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r ApiReturnAllEventsFromOneProjectRequest) IncludeRaw(includeRaw bool) ApiReturnAllEventsFromOneProjectRequest {
	r.includeRaw = &includeRaw
	return r
}

// Date and time from when MongoDB Cloud stops returning events. This parameter uses the ISO 8601 timestamp format in UTC.
func (r ApiReturnAllEventsFromOneProjectRequest) MaxDate(maxDate time.Time) ApiReturnAllEventsFromOneProjectRequest {
	r.maxDate = &maxDate
	return r
}

// Date and time from when MongoDB Cloud starts returning events. This parameter uses the ISO 8601 timestamp format in UTC.
func (r ApiReturnAllEventsFromOneProjectRequest) MinDate(minDate time.Time) ApiReturnAllEventsFromOneProjectRequest {
	r.minDate = &minDate
	return r
}

func (r ApiReturnAllEventsFromOneProjectRequest) Execute() (*PaginatedEventView, *http.Response, error) {
	return r.ApiService.ReturnAllEventsFromOneProjectExecute(r)
}

/*
ReturnAllEventsFromOneProject Return All Events from One Project

Returns one event for the specified project. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-projects-get-all/).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @return ApiReturnAllEventsFromOneProjectRequest
*/
func (a *EventsApiService) ReturnAllEventsFromOneProject(ctx context.Context, groupId string) ApiReturnAllEventsFromOneProjectRequest {
	return ApiReturnAllEventsFromOneProjectRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return PaginatedEventView
func (a *EventsApiService) ReturnAllEventsFromOneProjectExecute(r ApiReturnAllEventsFromOneProjectRequest) (*PaginatedEventView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedEventView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ReturnAllEventsFromOneProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/events"
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
	if r.clusterNames != nil {
		t := *r.clusterNames
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("clusterNames", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("clusterNames", parameterToString(t, "multi"))
		}
	}
	if r.eventType != nil {
		localVarQueryParams.Add("eventType", parameterToString(*r.eventType, ""))
	}
	if r.includeRaw != nil {
		localVarQueryParams.Add("includeRaw", parameterToString(*r.includeRaw, ""))
	}
	if r.maxDate != nil {
		localVarQueryParams.Add("maxDate", parameterToString(*r.maxDate, ""))
	}
	if r.minDate != nil {
		localVarQueryParams.Add("minDate", parameterToString(*r.minDate, ""))
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

type ApiReturnOneEventFromOneOrganizationRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	orgId string
	eventId string
	envelope *bool
	pretty *bool
	includeRaw *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiReturnOneEventFromOneOrganizationRequest) Envelope(envelope bool) ApiReturnOneEventFromOneOrganizationRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiReturnOneEventFromOneOrganizationRequest) Pretty(pretty bool) ApiReturnOneEventFromOneOrganizationRequest {
	r.pretty = &pretty
	return r
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r ApiReturnOneEventFromOneOrganizationRequest) IncludeRaw(includeRaw bool) ApiReturnOneEventFromOneOrganizationRequest {
	r.includeRaw = &includeRaw
	return r
}

func (r ApiReturnOneEventFromOneOrganizationRequest) Execute() (*ApiEventView, *http.Response, error) {
	return r.ApiService.ReturnOneEventFromOneOrganizationExecute(r)
}

/*
ReturnOneEventFromOneOrganization Return One Event from One Organization

Returns one event for the specified organization. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Organization Member role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-orgs-get-one/).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects.
 @param eventId Unique 24-hexadecimal digit string that identifies the event that you want to return.
 @return ApiReturnOneEventFromOneOrganizationRequest
*/
func (a *EventsApiService) ReturnOneEventFromOneOrganization(ctx context.Context, orgId string, eventId string) ApiReturnOneEventFromOneOrganizationRequest {
	return ApiReturnOneEventFromOneOrganizationRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		eventId: eventId,
	}
}

// Execute executes the request
//  @return ApiEventView
func (a *EventsApiService) ReturnOneEventFromOneOrganizationExecute(r ApiReturnOneEventFromOneOrganizationRequest) (*ApiEventView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiEventView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ReturnOneEventFromOneOrganization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/orgs/{orgId}/events/{eventId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(parameterToString(r.eventId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.orgId) < 24 {
		return localVarReturnValue, nil, reportError("orgId must have at least 24 elements")
	}
	if strlen(r.orgId) > 24 {
		return localVarReturnValue, nil, reportError("orgId must have less than 24 elements")
	}
	if strlen(r.eventId) < 24 {
		return localVarReturnValue, nil, reportError("eventId must have at least 24 elements")
	}
	if strlen(r.eventId) > 24 {
		return localVarReturnValue, nil, reportError("eventId must have less than 24 elements")
	}

	if r.envelope != nil {
		localVarQueryParams.Add("envelope", parameterToString(*r.envelope, ""))
	}
	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.includeRaw != nil {
		localVarQueryParams.Add("includeRaw", parameterToString(*r.includeRaw, ""))
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

type ApiReturnOneEventFromOneProjectRequest struct {
	ctx context.Context
	ApiService *EventsApiService
	groupId string
	eventId string
	envelope *bool
	pretty *bool
	includeRaw *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r ApiReturnOneEventFromOneProjectRequest) Envelope(envelope bool) ApiReturnOneEventFromOneProjectRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the prettyprint format.
func (r ApiReturnOneEventFromOneProjectRequest) Pretty(pretty bool) ApiReturnOneEventFromOneProjectRequest {
	r.pretty = &pretty
	return r
}

// Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.
func (r ApiReturnOneEventFromOneProjectRequest) IncludeRaw(includeRaw bool) ApiReturnOneEventFromOneProjectRequest {
	r.includeRaw = &includeRaw
	return r
}

func (r ApiReturnOneEventFromOneProjectRequest) Execute() (*ApiEventView, *http.Response, error) {
	return r.ApiService.ReturnOneEventFromOneProjectExecute(r)
}

/*
ReturnOneEventFromOneProject Return One Event from One Project

Returns one event for the specified project. Events identify significant database, billing, or security activities or status changes. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 This resource remains under revision and may change. Refer to the [legacy documentation for this resource](https://www.mongodb.com/docs/atlas/reference/api/events-projects-get-one/).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project.
 @param eventId Unique 24-hexadecimal digit string that identifies the event that you want to return.
 @return ApiReturnOneEventFromOneProjectRequest
*/
func (a *EventsApiService) ReturnOneEventFromOneProject(ctx context.Context, groupId string, eventId string) ApiReturnOneEventFromOneProjectRequest {
	return ApiReturnOneEventFromOneProjectRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		eventId: eventId,
	}
}

// Execute executes the request
//  @return ApiEventView
func (a *EventsApiService) ReturnOneEventFromOneProjectExecute(r ApiReturnOneEventFromOneProjectRequest) (*ApiEventView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiEventView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.ReturnOneEventFromOneProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v1.0/groups/{groupId}/events/{eventId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(parameterToString(r.eventId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.eventId) < 24 {
		return localVarReturnValue, nil, reportError("eventId must have at least 24 elements")
	}
	if strlen(r.eventId) > 24 {
		return localVarReturnValue, nil, reportError("eventId must have less than 24 elements")
	}

	if r.envelope != nil {
		localVarQueryParams.Add("envelope", parameterToString(*r.envelope, ""))
	}
	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.includeRaw != nil {
		localVarQueryParams.Add("includeRaw", parameterToString(*r.includeRaw, ""))
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
