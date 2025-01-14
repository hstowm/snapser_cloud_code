/*
Roi_dev

Your custom SDK

API version: Roi_dev: v9 SDK
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package snapser

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// StatisticsServiceApiService StatisticsServiceApi service
type StatisticsServiceApiService service

type ApiBatchUpdateUserStatisticsRequest struct {
	ctx        context.Context
	ApiService *StatisticsServiceApiService
	userId     string
	token      *string
	body       *BatchUpdateUserStatisticsRequest
}

// User session token
func (r ApiBatchUpdateUserStatisticsRequest) Token(token string) ApiBatchUpdateUserStatisticsRequest {
	r.token = &token
	return r
}

func (r ApiBatchUpdateUserStatisticsRequest) Body(body BatchUpdateUserStatisticsRequest) ApiBatchUpdateUserStatisticsRequest {
	r.body = &body
	return r
}

func (r ApiBatchUpdateUserStatisticsRequest) Execute() (*StatisticsBatchUpdateUserStatisticsResponse, *http.Response, error) {
	return r.ApiService.BatchUpdateUserStatisticsExecute(r)
}

/*
BatchUpdateUserStatistics User Statistics Batch Update API

Updates user statistics in bulk

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId User ID of the user for whom the stats are being updated
	@return ApiBatchUpdateUserStatisticsRequest
*/
func (a *StatisticsServiceApiService) BatchUpdateUserStatistics(ctx context.Context, userId string) ApiBatchUpdateUserStatisticsRequest {
	return ApiBatchUpdateUserStatisticsRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return	StatisticsBatchUpdateUserStatisticsResponse
func (a *StatisticsServiceApiService) BatchUpdateUserStatisticsExecute(r ApiBatchUpdateUserStatisticsRequest) (*StatisticsBatchUpdateUserStatisticsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StatisticsBatchUpdateUserStatisticsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsServiceApiService.BatchUpdateUserStatistics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/statistics/user-stats/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
	}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarHeaderParams["Token"] = parameterToString(*r.token, "")
	// body params
	localVarPostBody = r.body
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
		var v ErrorsApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiGetBulkUserStatisticsRequest struct {
	ctx        context.Context
	ApiService *StatisticsServiceApiService
	userId     string
	keys       []string
	token      *string
}

// User session token
func (r ApiGetBulkUserStatisticsRequest) Token(token string) ApiGetBulkUserStatisticsRequest {
	r.token = &token
	return r
}

func (r ApiGetBulkUserStatisticsRequest) Execute() (*StatisticsGetUserStatisticsResponse, *http.Response, error) {
	return r.ApiService.GetBulkUserStatisticsExecute(r)
}

/*
GetBulkUserStatistics User Statistics

Fetches multiple user statistic by user id and stat key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId User ID of the user
	@param keys Name of the user statistic
	@return ApiGetBulkUserStatisticsRequest
*/
func (a *StatisticsServiceApiService) GetBulkUserStatistics(ctx context.Context, userId string, keys []string) ApiGetBulkUserStatisticsRequest {
	return ApiGetBulkUserStatisticsRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		keys:       keys,
	}
}

// Execute executes the request
//
//	@return	StatisticsGetUserStatisticsResponse
func (a *StatisticsServiceApiService) GetBulkUserStatisticsExecute(r ApiGetBulkUserStatisticsRequest) (*StatisticsGetUserStatisticsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StatisticsGetUserStatisticsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsServiceApiService.GetBulkUserStatistics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/statistics/user-stats/{user_id}/{keys}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keys"+"}", url.PathEscape(parameterToString(r.keys, "csv")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if len(r.keys) < 1 {
		return localVarReturnValue, nil, reportError("keys must have at least 1 elements")
	}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
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
	localVarHeaderParams["Token"] = parameterToString(*r.token, "")
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
		var v ErrorsApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiGetUserStatisticRequest struct {
	ctx        context.Context
	ApiService *StatisticsServiceApiService
	userId     string
	key        string
	token      *string
}

// User session token
func (r ApiGetUserStatisticRequest) Token(token string) ApiGetUserStatisticRequest {
	r.token = &token
	return r
}

func (r ApiGetUserStatisticRequest) Execute() (*StatisticsUserStatistic, *http.Response, error) {
	return r.ApiService.GetUserStatisticExecute(r)
}

/*
GetUserStatistic User Statistic

Fetches a specific user statistic by user id and stat key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId User ID of the user
	@param key Name of the user statistic
	@return ApiGetUserStatisticRequest
*/
func (a *StatisticsServiceApiService) GetUserStatistic(ctx context.Context, userId string, key string) ApiGetUserStatisticRequest {
	return ApiGetUserStatisticRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		key:        key,
	}
}

// Execute executes the request
//
//	@return	StatisticsUserStatistic
func (a *StatisticsServiceApiService) GetUserStatisticExecute(r ApiGetUserStatisticRequest) (*StatisticsUserStatistic, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StatisticsUserStatistic
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsServiceApiService.GetUserStatistic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/statistics/user-stats/{user_id}/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.PathEscape(parameterToString(r.key, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
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
	localVarHeaderParams["Token"] = parameterToString(*r.token, "")
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
		var v ErrorsApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiGetUserStatisticsRequest struct {
	ctx        context.Context
	ApiService *StatisticsServiceApiService
	userId     string
	token      *string
}

// User session token
func (r ApiGetUserStatisticsRequest) Token(token string) ApiGetUserStatisticsRequest {
	r.token = &token
	return r
}

func (r ApiGetUserStatisticsRequest) Execute() (*StatisticsGetUserStatisticsResponse, *http.Response, error) {
	return r.ApiService.GetUserStatisticsExecute(r)
}

/*
GetUserStatistics User Statistics

Fetches all user-statistics for user or stat

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId User ID of the user
	@return ApiGetUserStatisticsRequest
*/
func (a *StatisticsServiceApiService) GetUserStatistics(ctx context.Context, userId string) ApiGetUserStatisticsRequest {
	return ApiGetUserStatisticsRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return	StatisticsGetUserStatisticsResponse
func (a *StatisticsServiceApiService) GetUserStatisticsExecute(r ApiGetUserStatisticsRequest) (*StatisticsGetUserStatisticsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StatisticsGetUserStatisticsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsServiceApiService.GetUserStatistics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/statistics/settings/user-stats/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
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
	localVarHeaderParams["Token"] = parameterToString(*r.token, "")
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
		var v ErrorsApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiIncrementUserStatisticRequest struct {
	ctx        context.Context
	ApiService *StatisticsServiceApiService
	userId     string
	key        string
	token      *string
	body       *IncrementUserStatisticRequest
}

// User session token
func (r ApiIncrementUserStatisticRequest) Token(token string) ApiIncrementUserStatisticRequest {
	r.token = &token
	return r
}

func (r ApiIncrementUserStatisticRequest) Body(body IncrementUserStatisticRequest) ApiIncrementUserStatisticRequest {
	r.body = &body
	return r
}

func (r ApiIncrementUserStatisticRequest) Execute() (*StatisticsUserStatistic, *http.Response, error) {
	return r.ApiService.IncrementUserStatisticExecute(r)
}

/*
IncrementUserStatistic User Statistic

Increments a user-statistic

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId User ID of the user who's stats are being requested
	@param key Name of the user statistic
	@return ApiIncrementUserStatisticRequest
*/
func (a *StatisticsServiceApiService) IncrementUserStatistic(ctx context.Context, userId string, key string) ApiIncrementUserStatisticRequest {
	return ApiIncrementUserStatisticRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		key:        key,
	}
}

// Execute executes the request
//
//	@return	StatisticsUserStatistic
func (a *StatisticsServiceApiService) IncrementUserStatisticExecute(r ApiIncrementUserStatisticRequest) (*StatisticsUserStatistic, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StatisticsUserStatistic
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsServiceApiService.IncrementUserStatistic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/statistics/user-stats/{user_id}/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.PathEscape(parameterToString(r.key, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
	}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarHeaderParams["Token"] = parameterToString(*r.token, "")
	// body params
	localVarPostBody = r.body
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
		var v ErrorsApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiIsUserInSegmentRequest struct {
	ctx        context.Context
	ApiService *StatisticsServiceApiService
	name       string
	userId     string
	token      *string
}

// User session token
func (r ApiIsUserInSegmentRequest) Token(token string) ApiIsUserInSegmentRequest {
	r.token = &token
	return r
}

func (r ApiIsUserInSegmentRequest) Execute() (*StatisticsIsUserInSegmentResponse, *http.Response, error) {
	return r.ApiService.IsUserInSegmentExecute(r)
}

/*
IsUserInSegment User Segment Lookup

Returns true/false if user is in segment or not

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of the segment
	@param userId User ID of the user
	@return ApiIsUserInSegmentRequest
*/
func (a *StatisticsServiceApiService) IsUserInSegment(ctx context.Context, name string, userId string) ApiIsUserInSegmentRequest {
	return ApiIsUserInSegmentRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
		userId:     userId,
	}
}

// Execute executes the request
//
//	@return	StatisticsIsUserInSegmentResponse
func (a *StatisticsServiceApiService) IsUserInSegmentExecute(r ApiIsUserInSegmentRequest) (*StatisticsIsUserInSegmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StatisticsIsUserInSegmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsServiceApiService.IsUserInSegment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/statistics/segments/{name}/users/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
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
	localVarHeaderParams["Token"] = parameterToString(*r.token, "")
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
		var v ErrorsApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiSetUserStatisticRequest struct {
	ctx        context.Context
	ApiService *StatisticsServiceApiService
	userId     string
	key        string
	token      *string
	body       *SetUserStatisticRequest
}

// User session token
func (r ApiSetUserStatisticRequest) Token(token string) ApiSetUserStatisticRequest {
	r.token = &token
	return r
}

func (r ApiSetUserStatisticRequest) Body(body SetUserStatisticRequest) ApiSetUserStatisticRequest {
	r.body = &body
	return r
}

func (r ApiSetUserStatisticRequest) Execute() (*StatisticsUserStatistic, *http.Response, error) {
	return r.ApiService.SetUserStatisticExecute(r)
}

/*
SetUserStatistic User Statistic

Sets a user statistic to exact value provided

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId User ID of the user who's stats are being requested
	@param key Name of the user statistic
	@return ApiSetUserStatisticRequest
*/
func (a *StatisticsServiceApiService) SetUserStatistic(ctx context.Context, userId string, key string) ApiSetUserStatisticRequest {
	return ApiSetUserStatisticRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		key:        key,
	}
}

// Execute executes the request
//
//	@return	StatisticsUserStatistic
func (a *StatisticsServiceApiService) SetUserStatisticExecute(r ApiSetUserStatisticRequest) (*StatisticsUserStatistic, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StatisticsUserStatistic
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatisticsServiceApiService.SetUserStatistic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/statistics/user-stats/{user_id}/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.PathEscape(parameterToString(r.key, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
	}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarHeaderParams["Token"] = parameterToString(*r.token, "")
	// body params
	localVarPostBody = r.body
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
		var v ErrorsApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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
