/**
 * (C) Copyright IBM Corp. 2024.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.93.0-c40121e6-20240729-182103
 */

// Package watsonxdatav2 : Operations and models for the WatsonxDataV2 service
package watsonxdatav2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/watsonxdata-go-sdk/common"
)

// WatsonxDataV2 : This is the Public API for IBM watsonx.data
//
// API Version: 2.0.0
type WatsonxDataV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://region.lakehouse.cloud.ibm.com/lakehouse/api/v2"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "watsonx_data"

// WatsonxDataV2Options : Service options
type WatsonxDataV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewWatsonxDataV2UsingExternalConfig : constructs an instance of WatsonxDataV2 with passed in options and external configuration.
func NewWatsonxDataV2UsingExternalConfig(options *WatsonxDataV2Options) (watsonxData *WatsonxDataV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	watsonxData, err = NewWatsonxDataV2(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = watsonxData.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = watsonxData.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewWatsonxDataV2 : constructs an instance of WatsonxDataV2 with passed in options.
func NewWatsonxDataV2(options *WatsonxDataV2Options) (service *WatsonxDataV2, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
			return
		}
	}

	service = &WatsonxDataV2{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "watsonxData" suitable for processing requests.
func (watsonxData *WatsonxDataV2) Clone() *WatsonxDataV2 {
	if core.IsNil(watsonxData) {
		return nil
	}
	clone := *watsonxData
	clone.Service = watsonxData.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (watsonxData *WatsonxDataV2) SetServiceURL(url string) error {
	err := watsonxData.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (watsonxData *WatsonxDataV2) GetServiceURL() string {
	return watsonxData.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (watsonxData *WatsonxDataV2) SetDefaultHeaders(headers http.Header) {
	watsonxData.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (watsonxData *WatsonxDataV2) SetEnableGzipCompression(enableGzip bool) {
	watsonxData.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (watsonxData *WatsonxDataV2) GetEnableGzipCompression() bool {
	return watsonxData.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (watsonxData *WatsonxDataV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	watsonxData.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (watsonxData *WatsonxDataV2) DisableRetries() {
	watsonxData.Service.DisableRetries()
}

// ListBucketRegistrations : Get bucket registrations
// Get list of registered buckets.
func (watsonxData *WatsonxDataV2) ListBucketRegistrations(listBucketRegistrationsOptions *ListBucketRegistrationsOptions) (result *BucketRegistrationCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListBucketRegistrationsWithContext(context.Background(), listBucketRegistrationsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListBucketRegistrationsWithContext is an alternate form of the ListBucketRegistrations method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListBucketRegistrationsWithContext(ctx context.Context, listBucketRegistrationsOptions *ListBucketRegistrationsOptions) (result *BucketRegistrationCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listBucketRegistrationsOptions, "listBucketRegistrationsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listBucketRegistrationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListBucketRegistrations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listBucketRegistrationsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listBucketRegistrationsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_bucket_registrations", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketRegistrationCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateBucketRegistration : Register bucket
// Register a new bucket.
func (watsonxData *WatsonxDataV2) CreateBucketRegistration(createBucketRegistrationOptions *CreateBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateBucketRegistrationWithContext(context.Background(), createBucketRegistrationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateBucketRegistrationWithContext is an alternate form of the CreateBucketRegistration method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateBucketRegistrationWithContext(ctx context.Context, createBucketRegistrationOptions *CreateBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createBucketRegistrationOptions, "createBucketRegistrationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createBucketRegistrationOptions, "createBucketRegistrationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createBucketRegistrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateBucketRegistration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createBucketRegistrationOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createBucketRegistrationOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createBucketRegistrationOptions.BucketDetails != nil {
		body["bucket_details"] = createBucketRegistrationOptions.BucketDetails
	}
	if createBucketRegistrationOptions.BucketType != nil {
		body["bucket_type"] = createBucketRegistrationOptions.BucketType
	}
	if createBucketRegistrationOptions.Description != nil {
		body["description"] = createBucketRegistrationOptions.Description
	}
	if createBucketRegistrationOptions.ManagedBy != nil {
		body["managed_by"] = createBucketRegistrationOptions.ManagedBy
	}
	if createBucketRegistrationOptions.AssociatedCatalog != nil {
		body["associated_catalog"] = createBucketRegistrationOptions.AssociatedCatalog
	}
	if createBucketRegistrationOptions.BucketDisplayName != nil {
		body["bucket_display_name"] = createBucketRegistrationOptions.BucketDisplayName
	}
	if createBucketRegistrationOptions.Region != nil {
		body["region"] = createBucketRegistrationOptions.Region
	}
	if createBucketRegistrationOptions.Tags != nil {
		body["tags"] = createBucketRegistrationOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_bucket_registration", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketRegistration)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetBucketRegistration : Get bucket
// Get a registered bucket.
func (watsonxData *WatsonxDataV2) GetBucketRegistration(getBucketRegistrationOptions *GetBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetBucketRegistrationWithContext(context.Background(), getBucketRegistrationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetBucketRegistrationWithContext is an alternate form of the GetBucketRegistration method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetBucketRegistrationWithContext(ctx context.Context, getBucketRegistrationOptions *GetBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBucketRegistrationOptions, "getBucketRegistrationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getBucketRegistrationOptions, "getBucketRegistrationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *getBucketRegistrationOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getBucketRegistrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetBucketRegistration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getBucketRegistrationOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getBucketRegistrationOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_bucket_registration", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketRegistration)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeregisterBucket : Deregister Bucket
// Deregister a bucket.
func (watsonxData *WatsonxDataV2) DeregisterBucket(deregisterBucketOptions *DeregisterBucketOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeregisterBucketWithContext(context.Background(), deregisterBucketOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeregisterBucketWithContext is an alternate form of the DeregisterBucket method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeregisterBucketWithContext(ctx context.Context, deregisterBucketOptions *DeregisterBucketOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deregisterBucketOptions, "deregisterBucketOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deregisterBucketOptions, "deregisterBucketOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *deregisterBucketOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deregisterBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeregisterBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deregisterBucketOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deregisterBucketOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "deregister_bucket", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateBucketRegistration : Update bucket
// Update bucket details & credentials.
func (watsonxData *WatsonxDataV2) UpdateBucketRegistration(updateBucketRegistrationOptions *UpdateBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.UpdateBucketRegistrationWithContext(context.Background(), updateBucketRegistrationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateBucketRegistrationWithContext is an alternate form of the UpdateBucketRegistration method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateBucketRegistrationWithContext(ctx context.Context, updateBucketRegistrationOptions *UpdateBucketRegistrationOptions) (result *BucketRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateBucketRegistrationOptions, "updateBucketRegistrationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateBucketRegistrationOptions, "updateBucketRegistrationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *updateBucketRegistrationOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateBucketRegistrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateBucketRegistration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updateBucketRegistrationOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateBucketRegistrationOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateBucketRegistrationOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_bucket_registration", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketRegistration)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateActivateBucket : Activate Bucket
// Activate a registered bucket.
func (watsonxData *WatsonxDataV2) CreateActivateBucket(createActivateBucketOptions *CreateActivateBucketOptions) (result *CreateActivateBucketCreatedBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateActivateBucketWithContext(context.Background(), createActivateBucketOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateActivateBucketWithContext is an alternate form of the CreateActivateBucket method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateActivateBucketWithContext(ctx context.Context, createActivateBucketOptions *CreateActivateBucketOptions) (result *CreateActivateBucketCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createActivateBucketOptions, "createActivateBucketOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createActivateBucketOptions, "createActivateBucketOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *createActivateBucketOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}/activate`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createActivateBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateActivateBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createActivateBucketOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createActivateBucketOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_activate_bucket", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateActivateBucketCreatedBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteDeactivateBucket : Deactivate Bucket
// Deactivate a bucket.
func (watsonxData *WatsonxDataV2) DeleteDeactivateBucket(deleteDeactivateBucketOptions *DeleteDeactivateBucketOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteDeactivateBucketWithContext(context.Background(), deleteDeactivateBucketOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDeactivateBucketWithContext is an alternate form of the DeleteDeactivateBucket method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteDeactivateBucketWithContext(ctx context.Context, deleteDeactivateBucketOptions *DeleteDeactivateBucketOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDeactivateBucketOptions, "deleteDeactivateBucketOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteDeactivateBucketOptions, "deleteDeactivateBucketOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *deleteDeactivateBucketOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}/deactivate`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteDeactivateBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteDeactivateBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteDeactivateBucketOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteDeactivateBucketOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_deactivate_bucket", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListBucketObjects : List bucket objects
// Fetch all objects from a given bucket.
func (watsonxData *WatsonxDataV2) ListBucketObjects(listBucketObjectsOptions *ListBucketObjectsOptions) (result *BucketRegistrationObjectCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListBucketObjectsWithContext(context.Background(), listBucketObjectsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListBucketObjectsWithContext is an alternate form of the ListBucketObjects method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListBucketObjectsWithContext(ctx context.Context, listBucketObjectsOptions *ListBucketObjectsOptions) (result *BucketRegistrationObjectCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listBucketObjectsOptions, "listBucketObjectsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listBucketObjectsOptions, "listBucketObjectsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"bucket_id": *listBucketObjectsOptions.BucketID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/bucket_registrations/{bucket_id}/objects`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listBucketObjectsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListBucketObjects")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listBucketObjectsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listBucketObjectsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_bucket_objects", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketRegistrationObjectCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListDatabaseRegistrations : Get databases
// Get list of databases.
func (watsonxData *WatsonxDataV2) ListDatabaseRegistrations(listDatabaseRegistrationsOptions *ListDatabaseRegistrationsOptions) (result *DatabaseRegistrationCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListDatabaseRegistrationsWithContext(context.Background(), listDatabaseRegistrationsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDatabaseRegistrationsWithContext is an alternate form of the ListDatabaseRegistrations method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListDatabaseRegistrationsWithContext(ctx context.Context, listDatabaseRegistrationsOptions *ListDatabaseRegistrationsOptions) (result *DatabaseRegistrationCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDatabaseRegistrationsOptions, "listDatabaseRegistrationsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/database_registrations`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listDatabaseRegistrationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListDatabaseRegistrations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listDatabaseRegistrationsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listDatabaseRegistrationsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_database_registrations", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDatabaseRegistrationCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateDatabaseRegistration : Add/Create database
// Add or create a new database.
func (watsonxData *WatsonxDataV2) CreateDatabaseRegistration(createDatabaseRegistrationOptions *CreateDatabaseRegistrationOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateDatabaseRegistrationWithContext(context.Background(), createDatabaseRegistrationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDatabaseRegistrationWithContext is an alternate form of the CreateDatabaseRegistration method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateDatabaseRegistrationWithContext(ctx context.Context, createDatabaseRegistrationOptions *CreateDatabaseRegistrationOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDatabaseRegistrationOptions, "createDatabaseRegistrationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createDatabaseRegistrationOptions, "createDatabaseRegistrationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/database_registrations`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createDatabaseRegistrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateDatabaseRegistration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createDatabaseRegistrationOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createDatabaseRegistrationOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createDatabaseRegistrationOptions.DatabaseDisplayName != nil {
		body["database_display_name"] = createDatabaseRegistrationOptions.DatabaseDisplayName
	}
	if createDatabaseRegistrationOptions.DatabaseType != nil {
		body["database_type"] = createDatabaseRegistrationOptions.DatabaseType
	}
	if createDatabaseRegistrationOptions.AssociatedCatalog != nil {
		body["associated_catalog"] = createDatabaseRegistrationOptions.AssociatedCatalog
	}
	if createDatabaseRegistrationOptions.CreatedOn != nil {
		body["created_on"] = createDatabaseRegistrationOptions.CreatedOn
	}
	if createDatabaseRegistrationOptions.DatabaseDetails != nil {
		body["database_details"] = createDatabaseRegistrationOptions.DatabaseDetails
	}
	if createDatabaseRegistrationOptions.DatabaseProperties != nil {
		body["database_properties"] = createDatabaseRegistrationOptions.DatabaseProperties
	}
	if createDatabaseRegistrationOptions.Description != nil {
		body["description"] = createDatabaseRegistrationOptions.Description
	}
	if createDatabaseRegistrationOptions.Tags != nil {
		body["tags"] = createDatabaseRegistrationOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_database_registration", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDatabaseRegistration)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDatabase : Get database
// Get a registered databases.
func (watsonxData *WatsonxDataV2) GetDatabase(getDatabaseOptions *GetDatabaseOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetDatabaseWithContext(context.Background(), getDatabaseOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDatabaseWithContext is an alternate form of the GetDatabase method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetDatabaseWithContext(ctx context.Context, getDatabaseOptions *GetDatabaseOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDatabaseOptions, "getDatabaseOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDatabaseOptions, "getDatabaseOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"database_id": *getDatabaseOptions.DatabaseID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/database_registrations/{database_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getDatabaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetDatabase")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDatabaseOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getDatabaseOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_database", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDatabaseRegistration)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteDatabaseCatalog : Delete database
// Delete a database.
func (watsonxData *WatsonxDataV2) DeleteDatabaseCatalog(deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteDatabaseCatalogWithContext(context.Background(), deleteDatabaseCatalogOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDatabaseCatalogWithContext is an alternate form of the DeleteDatabaseCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteDatabaseCatalogWithContext(ctx context.Context, deleteDatabaseCatalogOptions *DeleteDatabaseCatalogOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDatabaseCatalogOptions, "deleteDatabaseCatalogOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteDatabaseCatalogOptions, "deleteDatabaseCatalogOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"database_id": *deleteDatabaseCatalogOptions.DatabaseID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/database_registrations/{database_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteDatabaseCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteDatabaseCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteDatabaseCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteDatabaseCatalogOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_database_catalog", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateDatabase : Update database
// Update database details.
func (watsonxData *WatsonxDataV2) UpdateDatabase(updateDatabaseOptions *UpdateDatabaseOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.UpdateDatabaseWithContext(context.Background(), updateDatabaseOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateDatabaseWithContext is an alternate form of the UpdateDatabase method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateDatabaseWithContext(ctx context.Context, updateDatabaseOptions *UpdateDatabaseOptions) (result *DatabaseRegistration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDatabaseOptions, "updateDatabaseOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateDatabaseOptions, "updateDatabaseOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"database_id": *updateDatabaseOptions.DatabaseID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/database_registrations/{database_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateDatabaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateDatabase")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updateDatabaseOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateDatabaseOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateDatabaseOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_database", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDatabaseRegistration)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListOtherEngines : List other engines
// list all other engine details.
func (watsonxData *WatsonxDataV2) ListOtherEngines(listOtherEnginesOptions *ListOtherEnginesOptions) (result *OtherEngineCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListOtherEnginesWithContext(context.Background(), listOtherEnginesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListOtherEnginesWithContext is an alternate form of the ListOtherEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListOtherEnginesWithContext(ctx context.Context, listOtherEnginesOptions *ListOtherEnginesOptions) (result *OtherEngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listOtherEnginesOptions, "listOtherEnginesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/other_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listOtherEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListOtherEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listOtherEnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listOtherEnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_other_engines", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOtherEngineCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateOtherEngine : Create other engine
// Create a new engine.
func (watsonxData *WatsonxDataV2) CreateOtherEngine(createOtherEngineOptions *CreateOtherEngineOptions) (result *OtherEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateOtherEngineWithContext(context.Background(), createOtherEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateOtherEngineWithContext is an alternate form of the CreateOtherEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateOtherEngineWithContext(ctx context.Context, createOtherEngineOptions *CreateOtherEngineOptions) (result *OtherEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createOtherEngineOptions, "createOtherEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createOtherEngineOptions, "createOtherEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/other_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createOtherEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateOtherEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createOtherEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createOtherEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createOtherEngineOptions.EngineDetails != nil {
		body["engine_details"] = createOtherEngineOptions.EngineDetails
	}
	if createOtherEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createOtherEngineOptions.EngineDisplayName
	}
	if createOtherEngineOptions.Description != nil {
		body["description"] = createOtherEngineOptions.Description
	}
	if createOtherEngineOptions.Origin != nil {
		body["origin"] = createOtherEngineOptions.Origin
	}
	if createOtherEngineOptions.Tags != nil {
		body["tags"] = createOtherEngineOptions.Tags
	}
	if createOtherEngineOptions.Type != nil {
		body["type"] = createOtherEngineOptions.Type
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_other_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOtherEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteOtherEngine : Delete engine
// Delete an engine from lakehouse.
func (watsonxData *WatsonxDataV2) DeleteOtherEngine(deleteOtherEngineOptions *DeleteOtherEngineOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteOtherEngineWithContext(context.Background(), deleteOtherEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteOtherEngineWithContext is an alternate form of the DeleteOtherEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteOtherEngineWithContext(ctx context.Context, deleteOtherEngineOptions *DeleteOtherEngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteOtherEngineOptions, "deleteOtherEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteOtherEngineOptions, "deleteOtherEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteOtherEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/other_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteOtherEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteOtherEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteOtherEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteOtherEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_other_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListDb2Engines : Get list of db2 engines
// Get list of all db2 engines.
func (watsonxData *WatsonxDataV2) ListDb2Engines(listDb2EnginesOptions *ListDb2EnginesOptions) (result *Db2EngineCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListDb2EnginesWithContext(context.Background(), listDb2EnginesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDb2EnginesWithContext is an alternate form of the ListDb2Engines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListDb2EnginesWithContext(ctx context.Context, listDb2EnginesOptions *ListDb2EnginesOptions) (result *Db2EngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDb2EnginesOptions, "listDb2EnginesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/db2_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listDb2EnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListDb2Engines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listDb2EnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listDb2EnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_db2_engines", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDb2EngineCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateDb2Engine : Create db2 engine
// Create a new db2 engine.
func (watsonxData *WatsonxDataV2) CreateDb2Engine(createDb2EngineOptions *CreateDb2EngineOptions) (result *Db2Engine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateDb2EngineWithContext(context.Background(), createDb2EngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDb2EngineWithContext is an alternate form of the CreateDb2Engine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateDb2EngineWithContext(ctx context.Context, createDb2EngineOptions *CreateDb2EngineOptions) (result *Db2Engine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDb2EngineOptions, "createDb2EngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createDb2EngineOptions, "createDb2EngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/db2_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createDb2EngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateDb2Engine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createDb2EngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createDb2EngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createDb2EngineOptions.Origin != nil {
		body["origin"] = createDb2EngineOptions.Origin
	}
	if createDb2EngineOptions.Description != nil {
		body["description"] = createDb2EngineOptions.Description
	}
	if createDb2EngineOptions.EngineDetails != nil {
		body["engine_details"] = createDb2EngineOptions.EngineDetails
	}
	if createDb2EngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createDb2EngineOptions.EngineDisplayName
	}
	if createDb2EngineOptions.Tags != nil {
		body["tags"] = createDb2EngineOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_db2_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDb2Engine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteDb2Engine : Delete db2 engine
// Delete a db2 engine.
func (watsonxData *WatsonxDataV2) DeleteDb2Engine(deleteDb2EngineOptions *DeleteDb2EngineOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteDb2EngineWithContext(context.Background(), deleteDb2EngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDb2EngineWithContext is an alternate form of the DeleteDb2Engine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteDb2EngineWithContext(ctx context.Context, deleteDb2EngineOptions *DeleteDb2EngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDb2EngineOptions, "deleteDb2EngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteDb2EngineOptions, "deleteDb2EngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteDb2EngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/db2_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteDb2EngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteDb2Engine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteDb2EngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteDb2EngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_db2_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateDb2Engine : Update db2 engine
// Update details of db2 engine.
func (watsonxData *WatsonxDataV2) UpdateDb2Engine(updateDb2EngineOptions *UpdateDb2EngineOptions) (result *Db2Engine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.UpdateDb2EngineWithContext(context.Background(), updateDb2EngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateDb2EngineWithContext is an alternate form of the UpdateDb2Engine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateDb2EngineWithContext(ctx context.Context, updateDb2EngineOptions *UpdateDb2EngineOptions) (result *Db2Engine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDb2EngineOptions, "updateDb2EngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateDb2EngineOptions, "updateDb2EngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *updateDb2EngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/db2_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateDb2EngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateDb2Engine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updateDb2EngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateDb2EngineOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateDb2EngineOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_db2_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDb2Engine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListNetezzaEngines : Get list of netezza engines
// Get list of all netezza engines.
func (watsonxData *WatsonxDataV2) ListNetezzaEngines(listNetezzaEnginesOptions *ListNetezzaEnginesOptions) (result *NetezzaEngineCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListNetezzaEnginesWithContext(context.Background(), listNetezzaEnginesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListNetezzaEnginesWithContext is an alternate form of the ListNetezzaEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListNetezzaEnginesWithContext(ctx context.Context, listNetezzaEnginesOptions *ListNetezzaEnginesOptions) (result *NetezzaEngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listNetezzaEnginesOptions, "listNetezzaEnginesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/netezza_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listNetezzaEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListNetezzaEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listNetezzaEnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listNetezzaEnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_netezza_engines", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNetezzaEngineCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateNetezzaEngine : Create netezza engine
// Create a new netezza engine.
func (watsonxData *WatsonxDataV2) CreateNetezzaEngine(createNetezzaEngineOptions *CreateNetezzaEngineOptions) (result *NetezzaEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateNetezzaEngineWithContext(context.Background(), createNetezzaEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateNetezzaEngineWithContext is an alternate form of the CreateNetezzaEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateNetezzaEngineWithContext(ctx context.Context, createNetezzaEngineOptions *CreateNetezzaEngineOptions) (result *NetezzaEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createNetezzaEngineOptions, "createNetezzaEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createNetezzaEngineOptions, "createNetezzaEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/netezza_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createNetezzaEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateNetezzaEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createNetezzaEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createNetezzaEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createNetezzaEngineOptions.Origin != nil {
		body["origin"] = createNetezzaEngineOptions.Origin
	}
	if createNetezzaEngineOptions.Description != nil {
		body["description"] = createNetezzaEngineOptions.Description
	}
	if createNetezzaEngineOptions.EngineDetails != nil {
		body["engine_details"] = createNetezzaEngineOptions.EngineDetails
	}
	if createNetezzaEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createNetezzaEngineOptions.EngineDisplayName
	}
	if createNetezzaEngineOptions.Tags != nil {
		body["tags"] = createNetezzaEngineOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_netezza_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNetezzaEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteNetezzaEngine : Delete netezza engine
// Delete a netezza engine.
func (watsonxData *WatsonxDataV2) DeleteNetezzaEngine(deleteNetezzaEngineOptions *DeleteNetezzaEngineOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteNetezzaEngineWithContext(context.Background(), deleteNetezzaEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteNetezzaEngineWithContext is an alternate form of the DeleteNetezzaEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteNetezzaEngineWithContext(ctx context.Context, deleteNetezzaEngineOptions *DeleteNetezzaEngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteNetezzaEngineOptions, "deleteNetezzaEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteNetezzaEngineOptions, "deleteNetezzaEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteNetezzaEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/netezza_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteNetezzaEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteNetezzaEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteNetezzaEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteNetezzaEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_netezza_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateNetezzaEngine : Update netezza engine
// Update details of netezza engine.
func (watsonxData *WatsonxDataV2) UpdateNetezzaEngine(updateNetezzaEngineOptions *UpdateNetezzaEngineOptions) (result *NetezzaEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.UpdateNetezzaEngineWithContext(context.Background(), updateNetezzaEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateNetezzaEngineWithContext is an alternate form of the UpdateNetezzaEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateNetezzaEngineWithContext(ctx context.Context, updateNetezzaEngineOptions *UpdateNetezzaEngineOptions) (result *NetezzaEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateNetezzaEngineOptions, "updateNetezzaEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateNetezzaEngineOptions, "updateNetezzaEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *updateNetezzaEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/netezza_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateNetezzaEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateNetezzaEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updateNetezzaEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateNetezzaEngineOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateNetezzaEngineOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_netezza_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNetezzaEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListPrestissimoEngines : Get list of prestissimo engines
// Get list of all prestissimo engines.
func (watsonxData *WatsonxDataV2) ListPrestissimoEngines(listPrestissimoEnginesOptions *ListPrestissimoEnginesOptions) (result *PrestissimoEngineCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListPrestissimoEnginesWithContext(context.Background(), listPrestissimoEnginesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListPrestissimoEnginesWithContext is an alternate form of the ListPrestissimoEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListPrestissimoEnginesWithContext(ctx context.Context, listPrestissimoEnginesOptions *ListPrestissimoEnginesOptions) (result *PrestissimoEngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listPrestissimoEnginesOptions, "listPrestissimoEnginesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listPrestissimoEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListPrestissimoEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPrestissimoEnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listPrestissimoEnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_prestissimo_engines", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestissimoEngineCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreatePrestissimoEngine : Create prestissimo engine
// Create a new prestissimo engine.
func (watsonxData *WatsonxDataV2) CreatePrestissimoEngine(createPrestissimoEngineOptions *CreatePrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreatePrestissimoEngineWithContext(context.Background(), createPrestissimoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreatePrestissimoEngineWithContext is an alternate form of the CreatePrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreatePrestissimoEngineWithContext(ctx context.Context, createPrestissimoEngineOptions *CreatePrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPrestissimoEngineOptions, "createPrestissimoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createPrestissimoEngineOptions, "createPrestissimoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createPrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreatePrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createPrestissimoEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createPrestissimoEngineOptions.Origin != nil {
		body["origin"] = createPrestissimoEngineOptions.Origin
	}
	if createPrestissimoEngineOptions.AssociatedCatalogs != nil {
		body["associated_catalogs"] = createPrestissimoEngineOptions.AssociatedCatalogs
	}
	if createPrestissimoEngineOptions.Description != nil {
		body["description"] = createPrestissimoEngineOptions.Description
	}
	if createPrestissimoEngineOptions.EngineDetails != nil {
		body["engine_details"] = createPrestissimoEngineOptions.EngineDetails
	}
	if createPrestissimoEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createPrestissimoEngineOptions.EngineDisplayName
	}
	if createPrestissimoEngineOptions.Region != nil {
		body["region"] = createPrestissimoEngineOptions.Region
	}
	if createPrestissimoEngineOptions.Tags != nil {
		body["tags"] = createPrestissimoEngineOptions.Tags
	}
	if createPrestissimoEngineOptions.Version != nil {
		body["version"] = createPrestissimoEngineOptions.Version
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_prestissimo_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestissimoEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetPrestissimoEngine : Get prestissimo engine
// Get details of one prestissimo engine.
func (watsonxData *WatsonxDataV2) GetPrestissimoEngine(getPrestissimoEngineOptions *GetPrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetPrestissimoEngineWithContext(context.Background(), getPrestissimoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetPrestissimoEngineWithContext is an alternate form of the GetPrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetPrestissimoEngineWithContext(ctx context.Context, getPrestissimoEngineOptions *GetPrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPrestissimoEngineOptions, "getPrestissimoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getPrestissimoEngineOptions, "getPrestissimoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getPrestissimoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getPrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetPrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getPrestissimoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_prestissimo_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestissimoEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeletePrestissimoEngine : Delete prestissimo engine
// Delete a prestissimo engine.
func (watsonxData *WatsonxDataV2) DeletePrestissimoEngine(deletePrestissimoEngineOptions *DeletePrestissimoEngineOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeletePrestissimoEngineWithContext(context.Background(), deletePrestissimoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeletePrestissimoEngineWithContext is an alternate form of the DeletePrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeletePrestissimoEngineWithContext(ctx context.Context, deletePrestissimoEngineOptions *DeletePrestissimoEngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePrestissimoEngineOptions, "deletePrestissimoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deletePrestissimoEngineOptions, "deletePrestissimoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deletePrestissimoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deletePrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeletePrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deletePrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deletePrestissimoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_prestissimo_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdatePrestissimoEngine : Update prestissimo engine
// Update details of prestissimo engine.
func (watsonxData *WatsonxDataV2) UpdatePrestissimoEngine(updatePrestissimoEngineOptions *UpdatePrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.UpdatePrestissimoEngineWithContext(context.Background(), updatePrestissimoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdatePrestissimoEngineWithContext is an alternate form of the UpdatePrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdatePrestissimoEngineWithContext(ctx context.Context, updatePrestissimoEngineOptions *UpdatePrestissimoEngineOptions) (result *PrestissimoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePrestissimoEngineOptions, "updatePrestissimoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updatePrestissimoEngineOptions, "updatePrestissimoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *updatePrestissimoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updatePrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdatePrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updatePrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updatePrestissimoEngineOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updatePrestissimoEngineOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_prestissimo_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestissimoEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListPrestissimoEngineCatalogs : Get prestissimo engine catalogs
// Get list of all catalogs attached a prestissimo engine.
func (watsonxData *WatsonxDataV2) ListPrestissimoEngineCatalogs(listPrestissimoEngineCatalogsOptions *ListPrestissimoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListPrestissimoEngineCatalogsWithContext(context.Background(), listPrestissimoEngineCatalogsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListPrestissimoEngineCatalogsWithContext is an alternate form of the ListPrestissimoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListPrestissimoEngineCatalogsWithContext(ctx context.Context, listPrestissimoEngineCatalogsOptions *ListPrestissimoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPrestissimoEngineCatalogsOptions, "listPrestissimoEngineCatalogsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listPrestissimoEngineCatalogsOptions, "listPrestissimoEngineCatalogsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *listPrestissimoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listPrestissimoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListPrestissimoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPrestissimoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listPrestissimoEngineCatalogsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_prestissimo_engine_catalogs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// AddPrestissimoEngineCatalogs : Associate catalogs to a prestissimo engine
// Associate one or more catalogs to a prestissimo engine.
func (watsonxData *WatsonxDataV2) AddPrestissimoEngineCatalogs(addPrestissimoEngineCatalogsOptions *AddPrestissimoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.AddPrestissimoEngineCatalogsWithContext(context.Background(), addPrestissimoEngineCatalogsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// AddPrestissimoEngineCatalogsWithContext is an alternate form of the AddPrestissimoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) AddPrestissimoEngineCatalogsWithContext(ctx context.Context, addPrestissimoEngineCatalogsOptions *AddPrestissimoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addPrestissimoEngineCatalogsOptions, "addPrestissimoEngineCatalogsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(addPrestissimoEngineCatalogsOptions, "addPrestissimoEngineCatalogsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *addPrestissimoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range addPrestissimoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "AddPrestissimoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if addPrestissimoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*addPrestissimoEngineCatalogsOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if addPrestissimoEngineCatalogsOptions.CatalogNames != nil {
		body["catalog_names"] = addPrestissimoEngineCatalogsOptions.CatalogNames
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "add_prestissimo_engine_catalogs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeletePrestissimoEngineCatalogs : Disassociate catalogs from a prestissimo engine
// Disassociate one or more catalogs from a prestissimo engine.
func (watsonxData *WatsonxDataV2) DeletePrestissimoEngineCatalogs(deletePrestissimoEngineCatalogsOptions *DeletePrestissimoEngineCatalogsOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeletePrestissimoEngineCatalogsWithContext(context.Background(), deletePrestissimoEngineCatalogsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeletePrestissimoEngineCatalogsWithContext is an alternate form of the DeletePrestissimoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeletePrestissimoEngineCatalogsWithContext(ctx context.Context, deletePrestissimoEngineCatalogsOptions *DeletePrestissimoEngineCatalogsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePrestissimoEngineCatalogsOptions, "deletePrestissimoEngineCatalogsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deletePrestissimoEngineCatalogsOptions, "deletePrestissimoEngineCatalogsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deletePrestissimoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deletePrestissimoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeletePrestissimoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deletePrestissimoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deletePrestissimoEngineCatalogsOptions.AuthInstanceID))
	}

	builder.AddQuery("catalog_names", fmt.Sprint(*deletePrestissimoEngineCatalogsOptions.CatalogNames))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_prestissimo_engine_catalogs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetPrestissimoEngineCatalog : Get prestissimo engine catalog
// Get catalog attached to a prestissimo engine.
func (watsonxData *WatsonxDataV2) GetPrestissimoEngineCatalog(getPrestissimoEngineCatalogOptions *GetPrestissimoEngineCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetPrestissimoEngineCatalogWithContext(context.Background(), getPrestissimoEngineCatalogOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetPrestissimoEngineCatalogWithContext is an alternate form of the GetPrestissimoEngineCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetPrestissimoEngineCatalogWithContext(ctx context.Context, getPrestissimoEngineCatalogOptions *GetPrestissimoEngineCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPrestissimoEngineCatalogOptions, "getPrestissimoEngineCatalogOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getPrestissimoEngineCatalogOptions, "getPrestissimoEngineCatalogOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getPrestissimoEngineCatalogOptions.EngineID,
		"catalog_id": *getPrestissimoEngineCatalogOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/catalogs/{catalog_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getPrestissimoEngineCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetPrestissimoEngineCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPrestissimoEngineCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getPrestissimoEngineCatalogOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_prestissimo_engine_catalog", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalog)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// PausePrestissimoEngine : Pause prestissimo engine
// Pause a running prestissimo engine.
func (watsonxData *WatsonxDataV2) PausePrestissimoEngine(pausePrestissimoEngineOptions *PausePrestissimoEngineOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.PausePrestissimoEngineWithContext(context.Background(), pausePrestissimoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PausePrestissimoEngineWithContext is an alternate form of the PausePrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) PausePrestissimoEngineWithContext(ctx context.Context, pausePrestissimoEngineOptions *PausePrestissimoEngineOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(pausePrestissimoEngineOptions, "pausePrestissimoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(pausePrestissimoEngineOptions, "pausePrestissimoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *pausePrestissimoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/pause`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range pausePrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "PausePrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if pausePrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*pausePrestissimoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "pause_prestissimo_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// RunPrestissimoExplainStatement : Explain query
// Explain a query statement.
func (watsonxData *WatsonxDataV2) RunPrestissimoExplainStatement(runPrestissimoExplainStatementOptions *RunPrestissimoExplainStatementOptions) (result *ResultPrestissimoExplainStatement, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.RunPrestissimoExplainStatementWithContext(context.Background(), runPrestissimoExplainStatementOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// RunPrestissimoExplainStatementWithContext is an alternate form of the RunPrestissimoExplainStatement method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RunPrestissimoExplainStatementWithContext(ctx context.Context, runPrestissimoExplainStatementOptions *RunPrestissimoExplainStatementOptions) (result *ResultPrestissimoExplainStatement, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runPrestissimoExplainStatementOptions, "runPrestissimoExplainStatementOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(runPrestissimoExplainStatementOptions, "runPrestissimoExplainStatementOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *runPrestissimoExplainStatementOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/query_explain`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range runPrestissimoExplainStatementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RunPrestissimoExplainStatement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if runPrestissimoExplainStatementOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*runPrestissimoExplainStatementOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if runPrestissimoExplainStatementOptions.Statement != nil {
		body["statement"] = runPrestissimoExplainStatementOptions.Statement
	}
	if runPrestissimoExplainStatementOptions.Format != nil {
		body["format"] = runPrestissimoExplainStatementOptions.Format
	}
	if runPrestissimoExplainStatementOptions.Type != nil {
		body["type"] = runPrestissimoExplainStatementOptions.Type
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "run_prestissimo_explain_statement", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResultPrestissimoExplainStatement)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// RunPrestissimoExplainAnalyzeStatement : Explain analyze
// Return query metrics after query is complete.
func (watsonxData *WatsonxDataV2) RunPrestissimoExplainAnalyzeStatement(runPrestissimoExplainAnalyzeStatementOptions *RunPrestissimoExplainAnalyzeStatementOptions) (result *ResultRunPrestissimoExplainAnalyzeStatement, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.RunPrestissimoExplainAnalyzeStatementWithContext(context.Background(), runPrestissimoExplainAnalyzeStatementOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// RunPrestissimoExplainAnalyzeStatementWithContext is an alternate form of the RunPrestissimoExplainAnalyzeStatement method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RunPrestissimoExplainAnalyzeStatementWithContext(ctx context.Context, runPrestissimoExplainAnalyzeStatementOptions *RunPrestissimoExplainAnalyzeStatementOptions) (result *ResultRunPrestissimoExplainAnalyzeStatement, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runPrestissimoExplainAnalyzeStatementOptions, "runPrestissimoExplainAnalyzeStatementOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(runPrestissimoExplainAnalyzeStatementOptions, "runPrestissimoExplainAnalyzeStatementOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *runPrestissimoExplainAnalyzeStatementOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/query_explain_analyze`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range runPrestissimoExplainAnalyzeStatementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RunPrestissimoExplainAnalyzeStatement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if runPrestissimoExplainAnalyzeStatementOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*runPrestissimoExplainAnalyzeStatementOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if runPrestissimoExplainAnalyzeStatementOptions.Statement != nil {
		body["statement"] = runPrestissimoExplainAnalyzeStatementOptions.Statement
	}
	if runPrestissimoExplainAnalyzeStatementOptions.Verbose != nil {
		body["verbose"] = runPrestissimoExplainAnalyzeStatementOptions.Verbose
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "run_prestissimo_explain_analyze_statement", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResultRunPrestissimoExplainAnalyzeStatement)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// RestartPrestissimoEngine : Restart a prestissimo engine
// Restart an existing prestissimo engine.
func (watsonxData *WatsonxDataV2) RestartPrestissimoEngine(restartPrestissimoEngineOptions *RestartPrestissimoEngineOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.RestartPrestissimoEngineWithContext(context.Background(), restartPrestissimoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// RestartPrestissimoEngineWithContext is an alternate form of the RestartPrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RestartPrestissimoEngineWithContext(ctx context.Context, restartPrestissimoEngineOptions *RestartPrestissimoEngineOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(restartPrestissimoEngineOptions, "restartPrestissimoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(restartPrestissimoEngineOptions, "restartPrestissimoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *restartPrestissimoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/restart`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range restartPrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RestartPrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if restartPrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*restartPrestissimoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "restart_prestissimo_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ResumePrestissimoEngine : Resume prestissimo engine
// Resume a paused prestissimo engine.
func (watsonxData *WatsonxDataV2) ResumePrestissimoEngine(resumePrestissimoEngineOptions *ResumePrestissimoEngineOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ResumePrestissimoEngineWithContext(context.Background(), resumePrestissimoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ResumePrestissimoEngineWithContext is an alternate form of the ResumePrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ResumePrestissimoEngineWithContext(ctx context.Context, resumePrestissimoEngineOptions *ResumePrestissimoEngineOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resumePrestissimoEngineOptions, "resumePrestissimoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(resumePrestissimoEngineOptions, "resumePrestissimoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *resumePrestissimoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/resume`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range resumePrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ResumePrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if resumePrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*resumePrestissimoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "resume_prestissimo_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ScalePrestissimoEngine : Scale a prestissimo engine
// Scale an existing prestissimo engine.
func (watsonxData *WatsonxDataV2) ScalePrestissimoEngine(scalePrestissimoEngineOptions *ScalePrestissimoEngineOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ScalePrestissimoEngineWithContext(context.Background(), scalePrestissimoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ScalePrestissimoEngineWithContext is an alternate form of the ScalePrestissimoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ScalePrestissimoEngineWithContext(ctx context.Context, scalePrestissimoEngineOptions *ScalePrestissimoEngineOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(scalePrestissimoEngineOptions, "scalePrestissimoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(scalePrestissimoEngineOptions, "scalePrestissimoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *scalePrestissimoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/prestissimo_engines/{engine_id}/scale`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range scalePrestissimoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ScalePrestissimoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if scalePrestissimoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*scalePrestissimoEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if scalePrestissimoEngineOptions.Coordinator != nil {
		body["coordinator"] = scalePrestissimoEngineOptions.Coordinator
	}
	if scalePrestissimoEngineOptions.Worker != nil {
		body["worker"] = scalePrestissimoEngineOptions.Worker
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "scale_prestissimo_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListPrestoEngines : Get list of presto engines
// Get list of all presto engines.
func (watsonxData *WatsonxDataV2) ListPrestoEngines(listPrestoEnginesOptions *ListPrestoEnginesOptions) (result *PrestoEngineCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListPrestoEnginesWithContext(context.Background(), listPrestoEnginesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListPrestoEnginesWithContext is an alternate form of the ListPrestoEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListPrestoEnginesWithContext(ctx context.Context, listPrestoEnginesOptions *ListPrestoEnginesOptions) (result *PrestoEngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listPrestoEnginesOptions, "listPrestoEnginesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listPrestoEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListPrestoEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPrestoEnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listPrestoEnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_presto_engines", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestoEngineCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreatePrestoEngine : Create presto engine
// Create a new presto engine.
func (watsonxData *WatsonxDataV2) CreatePrestoEngine(createPrestoEngineOptions *CreatePrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreatePrestoEngineWithContext(context.Background(), createPrestoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreatePrestoEngineWithContext is an alternate form of the CreatePrestoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreatePrestoEngineWithContext(ctx context.Context, createPrestoEngineOptions *CreatePrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPrestoEngineOptions, "createPrestoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createPrestoEngineOptions, "createPrestoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createPrestoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreatePrestoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPrestoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createPrestoEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createPrestoEngineOptions.Origin != nil {
		body["origin"] = createPrestoEngineOptions.Origin
	}
	if createPrestoEngineOptions.AssociatedCatalogs != nil {
		body["associated_catalogs"] = createPrestoEngineOptions.AssociatedCatalogs
	}
	if createPrestoEngineOptions.Description != nil {
		body["description"] = createPrestoEngineOptions.Description
	}
	if createPrestoEngineOptions.EngineDetails != nil {
		body["engine_details"] = createPrestoEngineOptions.EngineDetails
	}
	if createPrestoEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createPrestoEngineOptions.EngineDisplayName
	}
	if createPrestoEngineOptions.Region != nil {
		body["region"] = createPrestoEngineOptions.Region
	}
	if createPrestoEngineOptions.Tags != nil {
		body["tags"] = createPrestoEngineOptions.Tags
	}
	if createPrestoEngineOptions.Version != nil {
		body["version"] = createPrestoEngineOptions.Version
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_presto_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestoEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetPrestoEngine : Get presto engine
// Get details of one presto engine.
func (watsonxData *WatsonxDataV2) GetPrestoEngine(getPrestoEngineOptions *GetPrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetPrestoEngineWithContext(context.Background(), getPrestoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetPrestoEngineWithContext is an alternate form of the GetPrestoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetPrestoEngineWithContext(ctx context.Context, getPrestoEngineOptions *GetPrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPrestoEngineOptions, "getPrestoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getPrestoEngineOptions, "getPrestoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getPrestoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getPrestoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetPrestoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPrestoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getPrestoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_presto_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestoEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteEngine : Delete presto engine
// Delete a presto engine.
func (watsonxData *WatsonxDataV2) DeleteEngine(deleteEngineOptions *DeleteEngineOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteEngineWithContext(context.Background(), deleteEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteEngineWithContext is an alternate form of the DeleteEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteEngineWithContext(ctx context.Context, deleteEngineOptions *DeleteEngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteEngineOptions, "deleteEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteEngineOptions, "deleteEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdatePrestoEngine : Update presto engine
// Update details of presto engine.
func (watsonxData *WatsonxDataV2) UpdatePrestoEngine(updatePrestoEngineOptions *UpdatePrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.UpdatePrestoEngineWithContext(context.Background(), updatePrestoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdatePrestoEngineWithContext is an alternate form of the UpdatePrestoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdatePrestoEngineWithContext(ctx context.Context, updatePrestoEngineOptions *UpdatePrestoEngineOptions) (result *PrestoEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePrestoEngineOptions, "updatePrestoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updatePrestoEngineOptions, "updatePrestoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *updatePrestoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updatePrestoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdatePrestoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updatePrestoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updatePrestoEngineOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updatePrestoEngineOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_presto_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrestoEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListPrestoEngineCatalogs : Get presto engine catalogs
// Get list of all catalogs attached to a presto engine.
func (watsonxData *WatsonxDataV2) ListPrestoEngineCatalogs(listPrestoEngineCatalogsOptions *ListPrestoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListPrestoEngineCatalogsWithContext(context.Background(), listPrestoEngineCatalogsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListPrestoEngineCatalogsWithContext is an alternate form of the ListPrestoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListPrestoEngineCatalogsWithContext(ctx context.Context, listPrestoEngineCatalogsOptions *ListPrestoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPrestoEngineCatalogsOptions, "listPrestoEngineCatalogsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listPrestoEngineCatalogsOptions, "listPrestoEngineCatalogsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *listPrestoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listPrestoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListPrestoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPrestoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listPrestoEngineCatalogsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_presto_engine_catalogs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// AddPrestoEngineCatalogs : Associate catalogs to presto engine
// Associate one or more catalogs to a presto engine.
func (watsonxData *WatsonxDataV2) AddPrestoEngineCatalogs(addPrestoEngineCatalogsOptions *AddPrestoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.AddPrestoEngineCatalogsWithContext(context.Background(), addPrestoEngineCatalogsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// AddPrestoEngineCatalogsWithContext is an alternate form of the AddPrestoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) AddPrestoEngineCatalogsWithContext(ctx context.Context, addPrestoEngineCatalogsOptions *AddPrestoEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addPrestoEngineCatalogsOptions, "addPrestoEngineCatalogsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(addPrestoEngineCatalogsOptions, "addPrestoEngineCatalogsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *addPrestoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range addPrestoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "AddPrestoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if addPrestoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*addPrestoEngineCatalogsOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if addPrestoEngineCatalogsOptions.CatalogNames != nil {
		body["catalog_names"] = addPrestoEngineCatalogsOptions.CatalogNames
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "add_presto_engine_catalogs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeletePrestoEngineCatalogs : Disassociate catalogs from a presto engine
// Disassociate one or more catalogs from a presto engine.
func (watsonxData *WatsonxDataV2) DeletePrestoEngineCatalogs(deletePrestoEngineCatalogsOptions *DeletePrestoEngineCatalogsOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeletePrestoEngineCatalogsWithContext(context.Background(), deletePrestoEngineCatalogsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeletePrestoEngineCatalogsWithContext is an alternate form of the DeletePrestoEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeletePrestoEngineCatalogsWithContext(ctx context.Context, deletePrestoEngineCatalogsOptions *DeletePrestoEngineCatalogsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePrestoEngineCatalogsOptions, "deletePrestoEngineCatalogsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deletePrestoEngineCatalogsOptions, "deletePrestoEngineCatalogsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deletePrestoEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deletePrestoEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeletePrestoEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deletePrestoEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deletePrestoEngineCatalogsOptions.AuthInstanceID))
	}

	builder.AddQuery("catalog_names", fmt.Sprint(*deletePrestoEngineCatalogsOptions.CatalogNames))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_presto_engine_catalogs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetPrestoEngineCatalog : Get presto engine catalog
// Get catalog attached to presto engine.
func (watsonxData *WatsonxDataV2) GetPrestoEngineCatalog(getPrestoEngineCatalogOptions *GetPrestoEngineCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetPrestoEngineCatalogWithContext(context.Background(), getPrestoEngineCatalogOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetPrestoEngineCatalogWithContext is an alternate form of the GetPrestoEngineCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetPrestoEngineCatalogWithContext(ctx context.Context, getPrestoEngineCatalogOptions *GetPrestoEngineCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPrestoEngineCatalogOptions, "getPrestoEngineCatalogOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getPrestoEngineCatalogOptions, "getPrestoEngineCatalogOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getPrestoEngineCatalogOptions.EngineID,
		"catalog_id": *getPrestoEngineCatalogOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/catalogs/{catalog_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getPrestoEngineCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetPrestoEngineCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPrestoEngineCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getPrestoEngineCatalogOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_presto_engine_catalog", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalog)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// PausePrestoEngine : Pause presto engine
// Pause a running presto engine.
func (watsonxData *WatsonxDataV2) PausePrestoEngine(pausePrestoEngineOptions *PausePrestoEngineOptions) (result *CreateEnginePauseCreatedBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.PausePrestoEngineWithContext(context.Background(), pausePrestoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PausePrestoEngineWithContext is an alternate form of the PausePrestoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) PausePrestoEngineWithContext(ctx context.Context, pausePrestoEngineOptions *PausePrestoEngineOptions) (result *CreateEnginePauseCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(pausePrestoEngineOptions, "pausePrestoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(pausePrestoEngineOptions, "pausePrestoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *pausePrestoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/pause`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range pausePrestoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "PausePrestoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if pausePrestoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*pausePrestoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "pause_presto_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateEnginePauseCreatedBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// RunExplainStatement : Explain presto query
// Explain a query statement.
func (watsonxData *WatsonxDataV2) RunExplainStatement(runExplainStatementOptions *RunExplainStatementOptions) (result *RunExplainStatementOKBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.RunExplainStatementWithContext(context.Background(), runExplainStatementOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// RunExplainStatementWithContext is an alternate form of the RunExplainStatement method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RunExplainStatementWithContext(ctx context.Context, runExplainStatementOptions *RunExplainStatementOptions) (result *RunExplainStatementOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runExplainStatementOptions, "runExplainStatementOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(runExplainStatementOptions, "runExplainStatementOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *runExplainStatementOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/query_explain`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range runExplainStatementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RunExplainStatement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if runExplainStatementOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*runExplainStatementOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if runExplainStatementOptions.Statement != nil {
		body["statement"] = runExplainStatementOptions.Statement
	}
	if runExplainStatementOptions.Format != nil {
		body["format"] = runExplainStatementOptions.Format
	}
	if runExplainStatementOptions.Type != nil {
		body["type"] = runExplainStatementOptions.Type
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "run_explain_statement", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRunExplainStatementOKBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// RunExplainAnalyzeStatement : Explain presto analyze
// Return query metrics after query is complete.
func (watsonxData *WatsonxDataV2) RunExplainAnalyzeStatement(runExplainAnalyzeStatementOptions *RunExplainAnalyzeStatementOptions) (result *RunExplainAnalyzeStatementOKBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.RunExplainAnalyzeStatementWithContext(context.Background(), runExplainAnalyzeStatementOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// RunExplainAnalyzeStatementWithContext is an alternate form of the RunExplainAnalyzeStatement method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RunExplainAnalyzeStatementWithContext(ctx context.Context, runExplainAnalyzeStatementOptions *RunExplainAnalyzeStatementOptions) (result *RunExplainAnalyzeStatementOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runExplainAnalyzeStatementOptions, "runExplainAnalyzeStatementOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(runExplainAnalyzeStatementOptions, "runExplainAnalyzeStatementOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *runExplainAnalyzeStatementOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/query_explain_analyze`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range runExplainAnalyzeStatementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RunExplainAnalyzeStatement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if runExplainAnalyzeStatementOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*runExplainAnalyzeStatementOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if runExplainAnalyzeStatementOptions.Statement != nil {
		body["statement"] = runExplainAnalyzeStatementOptions.Statement
	}
	if runExplainAnalyzeStatementOptions.Verbose != nil {
		body["verbose"] = runExplainAnalyzeStatementOptions.Verbose
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "run_explain_analyze_statement", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRunExplainAnalyzeStatementOKBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// RestartPrestoEngine : Restart a presto engine
// Restart an existing presto engine.
func (watsonxData *WatsonxDataV2) RestartPrestoEngine(restartPrestoEngineOptions *RestartPrestoEngineOptions) (result *CreateEngineRestartCreatedBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.RestartPrestoEngineWithContext(context.Background(), restartPrestoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// RestartPrestoEngineWithContext is an alternate form of the RestartPrestoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RestartPrestoEngineWithContext(ctx context.Context, restartPrestoEngineOptions *RestartPrestoEngineOptions) (result *CreateEngineRestartCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(restartPrestoEngineOptions, "restartPrestoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(restartPrestoEngineOptions, "restartPrestoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *restartPrestoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/restart`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range restartPrestoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RestartPrestoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if restartPrestoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*restartPrestoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "restart_presto_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateEngineRestartCreatedBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ResumePrestoEngine : Resume presto engine
// Resume a paused presto engine.
func (watsonxData *WatsonxDataV2) ResumePrestoEngine(resumePrestoEngineOptions *ResumePrestoEngineOptions) (result *CreateEngineResumeCreatedBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ResumePrestoEngineWithContext(context.Background(), resumePrestoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ResumePrestoEngineWithContext is an alternate form of the ResumePrestoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ResumePrestoEngineWithContext(ctx context.Context, resumePrestoEngineOptions *ResumePrestoEngineOptions) (result *CreateEngineResumeCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resumePrestoEngineOptions, "resumePrestoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(resumePrestoEngineOptions, "resumePrestoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *resumePrestoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/resume`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range resumePrestoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ResumePrestoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if resumePrestoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*resumePrestoEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "resume_presto_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateEngineResumeCreatedBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ScalePrestoEngine : Scale a presto engine
// Scale an existing presto engine.
func (watsonxData *WatsonxDataV2) ScalePrestoEngine(scalePrestoEngineOptions *ScalePrestoEngineOptions) (result *CreateEngineScaleCreatedBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ScalePrestoEngineWithContext(context.Background(), scalePrestoEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ScalePrestoEngineWithContext is an alternate form of the ScalePrestoEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ScalePrestoEngineWithContext(ctx context.Context, scalePrestoEngineOptions *ScalePrestoEngineOptions) (result *CreateEngineScaleCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(scalePrestoEngineOptions, "scalePrestoEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(scalePrestoEngineOptions, "scalePrestoEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *scalePrestoEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/presto_engines/{engine_id}/scale`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range scalePrestoEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ScalePrestoEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if scalePrestoEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*scalePrestoEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if scalePrestoEngineOptions.Coordinator != nil {
		body["coordinator"] = scalePrestoEngineOptions.Coordinator
	}
	if scalePrestoEngineOptions.Worker != nil {
		body["worker"] = scalePrestoEngineOptions.Worker
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "scale_presto_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateEngineScaleCreatedBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListSparkEngines : List all spark engines
// List all spark engines.
func (watsonxData *WatsonxDataV2) ListSparkEngines(listSparkEnginesOptions *ListSparkEnginesOptions) (result *SparkEngineCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListSparkEnginesWithContext(context.Background(), listSparkEnginesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListSparkEnginesWithContext is an alternate form of the ListSparkEngines method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListSparkEnginesWithContext(ctx context.Context, listSparkEnginesOptions *ListSparkEnginesOptions) (result *SparkEngineCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listSparkEnginesOptions, "listSparkEnginesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listSparkEnginesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListSparkEngines")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listSparkEnginesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listSparkEnginesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_spark_engines", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngineCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateSparkEngine : Create spark engine
// Create a new spark  engine.
func (watsonxData *WatsonxDataV2) CreateSparkEngine(createSparkEngineOptions *CreateSparkEngineOptions) (result *SparkEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateSparkEngineWithContext(context.Background(), createSparkEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateSparkEngineWithContext is an alternate form of the CreateSparkEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateSparkEngineWithContext(ctx context.Context, createSparkEngineOptions *CreateSparkEngineOptions) (result *SparkEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSparkEngineOptions, "createSparkEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createSparkEngineOptions, "createSparkEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createSparkEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateSparkEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createSparkEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createSparkEngineOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createSparkEngineOptions.Origin != nil {
		body["origin"] = createSparkEngineOptions.Origin
	}
	if createSparkEngineOptions.AssociatedCatalogs != nil {
		body["associated_catalogs"] = createSparkEngineOptions.AssociatedCatalogs
	}
	if createSparkEngineOptions.Description != nil {
		body["description"] = createSparkEngineOptions.Description
	}
	if createSparkEngineOptions.EngineDetails != nil {
		body["engine_details"] = createSparkEngineOptions.EngineDetails
	}
	if createSparkEngineOptions.EngineDisplayName != nil {
		body["engine_display_name"] = createSparkEngineOptions.EngineDisplayName
	}
	if createSparkEngineOptions.Status != nil {
		body["status"] = createSparkEngineOptions.Status
	}
	if createSparkEngineOptions.Tags != nil {
		body["tags"] = createSparkEngineOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_spark_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetSparkEngine : Get spark engine
// Get spark engine by ID.
func (watsonxData *WatsonxDataV2) GetSparkEngine(getSparkEngineOptions *GetSparkEngineOptions) (result *SparkEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetSparkEngineWithContext(context.Background(), getSparkEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetSparkEngineWithContext is an alternate form of the GetSparkEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetSparkEngineWithContext(ctx context.Context, getSparkEngineOptions *GetSparkEngineOptions) (result *SparkEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSparkEngineOptions, "getSparkEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getSparkEngineOptions, "getSparkEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getSparkEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getSparkEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetSparkEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSparkEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getSparkEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_spark_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteSparkEngine : Delete spark engine
// Delete a spark engine.
func (watsonxData *WatsonxDataV2) DeleteSparkEngine(deleteSparkEngineOptions *DeleteSparkEngineOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteSparkEngineWithContext(context.Background(), deleteSparkEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteSparkEngineWithContext is an alternate form of the DeleteSparkEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteSparkEngineWithContext(ctx context.Context, deleteSparkEngineOptions *DeleteSparkEngineOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSparkEngineOptions, "deleteSparkEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteSparkEngineOptions, "deleteSparkEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteSparkEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteSparkEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteSparkEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteSparkEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteSparkEngineOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_spark_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateSparkEngine : Update spark engine
// Update details of spark engine.
func (watsonxData *WatsonxDataV2) UpdateSparkEngine(updateSparkEngineOptions *UpdateSparkEngineOptions) (result *SparkEngine, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.UpdateSparkEngineWithContext(context.Background(), updateSparkEngineOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateSparkEngineWithContext is an alternate form of the UpdateSparkEngine method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateSparkEngineWithContext(ctx context.Context, updateSparkEngineOptions *UpdateSparkEngineOptions) (result *SparkEngine, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSparkEngineOptions, "updateSparkEngineOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateSparkEngineOptions, "updateSparkEngineOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *updateSparkEngineOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateSparkEngineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateSparkEngine")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updateSparkEngineOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateSparkEngineOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateSparkEngineOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_spark_engine", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngine)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListSparkEngineApplications : List all applications in a spark engine
// List all applications in a spark engine.
func (watsonxData *WatsonxDataV2) ListSparkEngineApplications(listSparkEngineApplicationsOptions *ListSparkEngineApplicationsOptions) (result *SparkEngineApplicationStatusCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListSparkEngineApplicationsWithContext(context.Background(), listSparkEngineApplicationsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListSparkEngineApplicationsWithContext is an alternate form of the ListSparkEngineApplications method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListSparkEngineApplicationsWithContext(ctx context.Context, listSparkEngineApplicationsOptions *ListSparkEngineApplicationsOptions) (result *SparkEngineApplicationStatusCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listSparkEngineApplicationsOptions, "listSparkEngineApplicationsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listSparkEngineApplicationsOptions, "listSparkEngineApplicationsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *listSparkEngineApplicationsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/applications`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listSparkEngineApplicationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListSparkEngineApplications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listSparkEngineApplicationsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listSparkEngineApplicationsOptions.AuthInstanceID))
	}

	if listSparkEngineApplicationsOptions.State != nil {
		builder.AddQuery("state", strings.Join(listSparkEngineApplicationsOptions.State, ","))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_spark_engine_applications", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngineApplicationStatusCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateSparkEngineApplication : Submit engine applications
// Submit engine applications.
func (watsonxData *WatsonxDataV2) CreateSparkEngineApplication(createSparkEngineApplicationOptions *CreateSparkEngineApplicationOptions) (result *SparkEngineApplicationStatus, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateSparkEngineApplicationWithContext(context.Background(), createSparkEngineApplicationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateSparkEngineApplicationWithContext is an alternate form of the CreateSparkEngineApplication method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateSparkEngineApplicationWithContext(ctx context.Context, createSparkEngineApplicationOptions *CreateSparkEngineApplicationOptions) (result *SparkEngineApplicationStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSparkEngineApplicationOptions, "createSparkEngineApplicationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createSparkEngineApplicationOptions, "createSparkEngineApplicationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createSparkEngineApplicationOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/applications`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createSparkEngineApplicationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateSparkEngineApplication")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createSparkEngineApplicationOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createSparkEngineApplicationOptions.AuthInstanceID))
	}

	if createSparkEngineApplicationOptions.State != nil {
		builder.AddQuery("state", strings.Join(createSparkEngineApplicationOptions.State, ","))
	}

	body := make(map[string]interface{})
	if createSparkEngineApplicationOptions.ApplicationDetails != nil {
		body["application_details"] = createSparkEngineApplicationOptions.ApplicationDetails
	}
	if createSparkEngineApplicationOptions.JobEndpoint != nil {
		body["job_endpoint"] = createSparkEngineApplicationOptions.JobEndpoint
	}
	if createSparkEngineApplicationOptions.ServiceInstanceID != nil {
		body["service_instance_id"] = createSparkEngineApplicationOptions.ServiceInstanceID
	}
	if createSparkEngineApplicationOptions.Type != nil {
		body["type"] = createSparkEngineApplicationOptions.Type
	}
	if createSparkEngineApplicationOptions.Volumes != nil {
		body["volumes"] = createSparkEngineApplicationOptions.Volumes
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_spark_engine_application", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngineApplicationStatus)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteSparkEngineApplications : Stop Spark Applications
// Stop a running spark application.
func (watsonxData *WatsonxDataV2) DeleteSparkEngineApplications(deleteSparkEngineApplicationsOptions *DeleteSparkEngineApplicationsOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteSparkEngineApplicationsWithContext(context.Background(), deleteSparkEngineApplicationsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteSparkEngineApplicationsWithContext is an alternate form of the DeleteSparkEngineApplications method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteSparkEngineApplicationsWithContext(ctx context.Context, deleteSparkEngineApplicationsOptions *DeleteSparkEngineApplicationsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSparkEngineApplicationsOptions, "deleteSparkEngineApplicationsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteSparkEngineApplicationsOptions, "deleteSparkEngineApplicationsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteSparkEngineApplicationsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/applications`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteSparkEngineApplicationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteSparkEngineApplications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteSparkEngineApplicationsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteSparkEngineApplicationsOptions.AuthInstanceID))
	}

	builder.AddQuery("application_id", fmt.Sprint(*deleteSparkEngineApplicationsOptions.ApplicationID))
	if deleteSparkEngineApplicationsOptions.State != nil {
		builder.AddQuery("state", strings.Join(deleteSparkEngineApplicationsOptions.State, ","))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_spark_engine_applications", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetSparkEngineApplicationStatus : Get spark application
// Get status of spark application.
func (watsonxData *WatsonxDataV2) GetSparkEngineApplicationStatus(getSparkEngineApplicationStatusOptions *GetSparkEngineApplicationStatusOptions) (result *SparkEngineApplicationStatus, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetSparkEngineApplicationStatusWithContext(context.Background(), getSparkEngineApplicationStatusOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetSparkEngineApplicationStatusWithContext is an alternate form of the GetSparkEngineApplicationStatus method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetSparkEngineApplicationStatusWithContext(ctx context.Context, getSparkEngineApplicationStatusOptions *GetSparkEngineApplicationStatusOptions) (result *SparkEngineApplicationStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSparkEngineApplicationStatusOptions, "getSparkEngineApplicationStatusOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getSparkEngineApplicationStatusOptions, "getSparkEngineApplicationStatusOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getSparkEngineApplicationStatusOptions.EngineID,
		"application_id": *getSparkEngineApplicationStatusOptions.ApplicationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/applications/{application_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getSparkEngineApplicationStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetSparkEngineApplicationStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSparkEngineApplicationStatusOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getSparkEngineApplicationStatusOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_spark_engine_application_status", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkEngineApplicationStatus)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListSparkEngineCatalogs : Get spark engine catalogs
// Get list of all catalogs attached to a spark engine.
func (watsonxData *WatsonxDataV2) ListSparkEngineCatalogs(listSparkEngineCatalogsOptions *ListSparkEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListSparkEngineCatalogsWithContext(context.Background(), listSparkEngineCatalogsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListSparkEngineCatalogsWithContext is an alternate form of the ListSparkEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListSparkEngineCatalogsWithContext(ctx context.Context, listSparkEngineCatalogsOptions *ListSparkEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listSparkEngineCatalogsOptions, "listSparkEngineCatalogsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listSparkEngineCatalogsOptions, "listSparkEngineCatalogsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *listSparkEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listSparkEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListSparkEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listSparkEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listSparkEngineCatalogsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_spark_engine_catalogs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// AddSparkEngineCatalogs : Associate catalogs to spark engine
// Associate one or more catalogs to a spark engine.
func (watsonxData *WatsonxDataV2) AddSparkEngineCatalogs(addSparkEngineCatalogsOptions *AddSparkEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.AddSparkEngineCatalogsWithContext(context.Background(), addSparkEngineCatalogsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// AddSparkEngineCatalogsWithContext is an alternate form of the AddSparkEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) AddSparkEngineCatalogsWithContext(ctx context.Context, addSparkEngineCatalogsOptions *AddSparkEngineCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addSparkEngineCatalogsOptions, "addSparkEngineCatalogsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(addSparkEngineCatalogsOptions, "addSparkEngineCatalogsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *addSparkEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range addSparkEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "AddSparkEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if addSparkEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*addSparkEngineCatalogsOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if addSparkEngineCatalogsOptions.CatalogNames != nil {
		body["catalog_names"] = addSparkEngineCatalogsOptions.CatalogNames
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "add_spark_engine_catalogs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteSparkEngineCatalogs : Disassociate catalogs from a spark engine
// Disassociate one or more catalogs from a spark engine.
func (watsonxData *WatsonxDataV2) DeleteSparkEngineCatalogs(deleteSparkEngineCatalogsOptions *DeleteSparkEngineCatalogsOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteSparkEngineCatalogsWithContext(context.Background(), deleteSparkEngineCatalogsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteSparkEngineCatalogsWithContext is an alternate form of the DeleteSparkEngineCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteSparkEngineCatalogsWithContext(ctx context.Context, deleteSparkEngineCatalogsOptions *DeleteSparkEngineCatalogsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSparkEngineCatalogsOptions, "deleteSparkEngineCatalogsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteSparkEngineCatalogsOptions, "deleteSparkEngineCatalogsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteSparkEngineCatalogsOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/catalogs`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteSparkEngineCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteSparkEngineCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteSparkEngineCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteSparkEngineCatalogsOptions.AuthInstanceID))
	}

	builder.AddQuery("catalog_names", fmt.Sprint(*deleteSparkEngineCatalogsOptions.CatalogNames))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_spark_engine_catalogs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetSparkEngineCatalog : Get spark engine catalog
// Get catalog attached to spark engine.
func (watsonxData *WatsonxDataV2) GetSparkEngineCatalog(getSparkEngineCatalogOptions *GetSparkEngineCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetSparkEngineCatalogWithContext(context.Background(), getSparkEngineCatalogOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetSparkEngineCatalogWithContext is an alternate form of the GetSparkEngineCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetSparkEngineCatalogWithContext(ctx context.Context, getSparkEngineCatalogOptions *GetSparkEngineCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSparkEngineCatalogOptions, "getSparkEngineCatalogOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getSparkEngineCatalogOptions, "getSparkEngineCatalogOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getSparkEngineCatalogOptions.EngineID,
		"catalog_id": *getSparkEngineCatalogOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/catalogs/{catalog_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getSparkEngineCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetSparkEngineCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSparkEngineCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getSparkEngineCatalogOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_spark_engine_catalog", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalog)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetSparkEngineHistoryServer : Get spark history server
// Get spark history server.
func (watsonxData *WatsonxDataV2) GetSparkEngineHistoryServer(getSparkEngineHistoryServerOptions *GetSparkEngineHistoryServerOptions) (result *SparkHistoryServer, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetSparkEngineHistoryServerWithContext(context.Background(), getSparkEngineHistoryServerOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetSparkEngineHistoryServerWithContext is an alternate form of the GetSparkEngineHistoryServer method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetSparkEngineHistoryServerWithContext(ctx context.Context, getSparkEngineHistoryServerOptions *GetSparkEngineHistoryServerOptions) (result *SparkHistoryServer, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSparkEngineHistoryServerOptions, "getSparkEngineHistoryServerOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getSparkEngineHistoryServerOptions, "getSparkEngineHistoryServerOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *getSparkEngineHistoryServerOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/history_server`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getSparkEngineHistoryServerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetSparkEngineHistoryServer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSparkEngineHistoryServerOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getSparkEngineHistoryServerOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_spark_engine_history_server", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkHistoryServer)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// StartSparkEngineHistoryServer : Start spark history server
// Start spark history server.
func (watsonxData *WatsonxDataV2) StartSparkEngineHistoryServer(startSparkEngineHistoryServerOptions *StartSparkEngineHistoryServerOptions) (result *SparkHistoryServer, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.StartSparkEngineHistoryServerWithContext(context.Background(), startSparkEngineHistoryServerOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// StartSparkEngineHistoryServerWithContext is an alternate form of the StartSparkEngineHistoryServer method which supports a Context parameter
func (watsonxData *WatsonxDataV2) StartSparkEngineHistoryServerWithContext(ctx context.Context, startSparkEngineHistoryServerOptions *StartSparkEngineHistoryServerOptions) (result *SparkHistoryServer, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(startSparkEngineHistoryServerOptions, "startSparkEngineHistoryServerOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(startSparkEngineHistoryServerOptions, "startSparkEngineHistoryServerOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *startSparkEngineHistoryServerOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/history_server`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range startSparkEngineHistoryServerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "StartSparkEngineHistoryServer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if startSparkEngineHistoryServerOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*startSparkEngineHistoryServerOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if startSparkEngineHistoryServerOptions.Cores != nil {
		body["cores"] = startSparkEngineHistoryServerOptions.Cores
	}
	if startSparkEngineHistoryServerOptions.Memory != nil {
		body["memory"] = startSparkEngineHistoryServerOptions.Memory
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "start_spark_engine_history_server", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSparkHistoryServer)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteSparkEngineHistoryServer : Stop spark history server
// Stop spark history server.
func (watsonxData *WatsonxDataV2) DeleteSparkEngineHistoryServer(deleteSparkEngineHistoryServerOptions *DeleteSparkEngineHistoryServerOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteSparkEngineHistoryServerWithContext(context.Background(), deleteSparkEngineHistoryServerOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteSparkEngineHistoryServerWithContext is an alternate form of the DeleteSparkEngineHistoryServer method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteSparkEngineHistoryServerWithContext(ctx context.Context, deleteSparkEngineHistoryServerOptions *DeleteSparkEngineHistoryServerOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSparkEngineHistoryServerOptions, "deleteSparkEngineHistoryServerOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteSparkEngineHistoryServerOptions, "deleteSparkEngineHistoryServerOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *deleteSparkEngineHistoryServerOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/history_server`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteSparkEngineHistoryServerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteSparkEngineHistoryServer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteSparkEngineHistoryServerOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteSparkEngineHistoryServerOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_spark_engine_history_server", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// CreateSparkEnginePause : Pause engine
// Pause engine.
func (watsonxData *WatsonxDataV2) CreateSparkEnginePause(createSparkEnginePauseOptions *CreateSparkEnginePauseOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateSparkEnginePauseWithContext(context.Background(), createSparkEnginePauseOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateSparkEnginePauseWithContext is an alternate form of the CreateSparkEnginePause method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateSparkEnginePauseWithContext(ctx context.Context, createSparkEnginePauseOptions *CreateSparkEnginePauseOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSparkEnginePauseOptions, "createSparkEnginePauseOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createSparkEnginePauseOptions, "createSparkEnginePauseOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createSparkEnginePauseOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/pause`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createSparkEnginePauseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateSparkEnginePause")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createSparkEnginePauseOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createSparkEnginePauseOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_spark_engine_pause", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateSparkEngineResume : Resume engine
// Resume engine.
func (watsonxData *WatsonxDataV2) CreateSparkEngineResume(createSparkEngineResumeOptions *CreateSparkEngineResumeOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateSparkEngineResumeWithContext(context.Background(), createSparkEngineResumeOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateSparkEngineResumeWithContext is an alternate form of the CreateSparkEngineResume method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateSparkEngineResumeWithContext(ctx context.Context, createSparkEngineResumeOptions *CreateSparkEngineResumeOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSparkEngineResumeOptions, "createSparkEngineResumeOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createSparkEngineResumeOptions, "createSparkEngineResumeOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createSparkEngineResumeOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/resume`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createSparkEngineResumeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateSparkEngineResume")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if createSparkEngineResumeOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createSparkEngineResumeOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_spark_engine_resume", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateSparkEngineScale : Scale Spark engine
// Scale Saprk engine.
func (watsonxData *WatsonxDataV2) CreateSparkEngineScale(createSparkEngineScaleOptions *CreateSparkEngineScaleOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateSparkEngineScaleWithContext(context.Background(), createSparkEngineScaleOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateSparkEngineScaleWithContext is an alternate form of the CreateSparkEngineScale method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateSparkEngineScaleWithContext(ctx context.Context, createSparkEngineScaleOptions *CreateSparkEngineScaleOptions) (result *SuccessResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSparkEngineScaleOptions, "createSparkEngineScaleOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createSparkEngineScaleOptions, "createSparkEngineScaleOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"engine_id": *createSparkEngineScaleOptions.EngineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_engines/{engine_id}/scale`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createSparkEngineScaleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateSparkEngineScale")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createSparkEngineScaleOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createSparkEngineScaleOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createSparkEngineScaleOptions.NumberOfNodes != nil {
		body["number_of_nodes"] = createSparkEngineScaleOptions.NumberOfNodes
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_spark_engine_scale", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListSparkVersions : List spark version
// List spark version.
func (watsonxData *WatsonxDataV2) ListSparkVersions(listSparkVersionsOptions *ListSparkVersionsOptions) (result *ListSparkVersionsOKBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListSparkVersionsWithContext(context.Background(), listSparkVersionsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListSparkVersionsWithContext is an alternate form of the ListSparkVersions method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListSparkVersionsWithContext(ctx context.Context, listSparkVersionsOptions *ListSparkVersionsOptions) (result *ListSparkVersionsOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listSparkVersionsOptions, "listSparkVersionsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/spark_versions`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listSparkVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListSparkVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listSparkVersionsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listSparkVersionsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_spark_versions", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListSparkVersionsOKBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListCatalogs : List all registered catalogs
// List all registered catalogs.
func (watsonxData *WatsonxDataV2) ListCatalogs(listCatalogsOptions *ListCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListCatalogsWithContext(context.Background(), listCatalogsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListCatalogsWithContext is an alternate form of the ListCatalogs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListCatalogsWithContext(ctx context.Context, listCatalogsOptions *ListCatalogsOptions) (result *CatalogCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCatalogsOptions, "listCatalogsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listCatalogsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listCatalogsOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_catalogs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetCatalog : Get catalog properties by catalog_id
// Get catalog properties of a catalog identified by catalog_id.
func (watsonxData *WatsonxDataV2) GetCatalog(getCatalogOptions *GetCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetCatalogWithContext(context.Background(), getCatalogOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetCatalogWithContext is an alternate form of the GetCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetCatalogWithContext(ctx context.Context, getCatalogOptions *GetCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCatalogOptions, "getCatalogOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getCatalogOptions, "getCatalogOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *getCatalogOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getCatalogOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_catalog", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalog)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListSchemas : List all schemas
// List all schemas in catalog.
func (watsonxData *WatsonxDataV2) ListSchemas(listSchemasOptions *ListSchemasOptions) (result *ListSchemasOKBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListSchemasWithContext(context.Background(), listSchemasOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListSchemasWithContext is an alternate form of the ListSchemas method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListSchemasWithContext(ctx context.Context, listSchemasOptions *ListSchemasOptions) (result *ListSchemasOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listSchemasOptions, "listSchemasOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listSchemasOptions, "listSchemasOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *listSchemasOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listSchemasOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListSchemas")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listSchemasOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listSchemasOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*listSchemasOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_schemas", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListSchemasOKBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateSchema : Create schema
// Create a new schema.
func (watsonxData *WatsonxDataV2) CreateSchema(createSchemaOptions *CreateSchemaOptions) (result *CreateSchemaCreatedBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateSchemaWithContext(context.Background(), createSchemaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateSchemaWithContext is an alternate form of the CreateSchema method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateSchemaWithContext(ctx context.Context, createSchemaOptions *CreateSchemaOptions) (result *CreateSchemaCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSchemaOptions, "createSchemaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createSchemaOptions, "createSchemaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *createSchemaOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createSchemaOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createSchemaOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*createSchemaOptions.EngineID))

	body := make(map[string]interface{})
	if createSchemaOptions.CustomPath != nil {
		body["custom_path"] = createSchemaOptions.CustomPath
	}
	if createSchemaOptions.SchemaName != nil {
		body["schema_name"] = createSchemaOptions.SchemaName
	}
	if createSchemaOptions.BucketName != nil {
		body["bucket_name"] = createSchemaOptions.BucketName
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_schema", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateSchemaCreatedBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteSchema : Delete schema
// Delete a schema.
func (watsonxData *WatsonxDataV2) DeleteSchema(deleteSchemaOptions *DeleteSchemaOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteSchemaWithContext(context.Background(), deleteSchemaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteSchemaWithContext is an alternate form of the DeleteSchema method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteSchemaWithContext(ctx context.Context, deleteSchemaOptions *DeleteSchemaOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSchemaOptions, "deleteSchemaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteSchemaOptions, "deleteSchemaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *deleteSchemaOptions.CatalogID,
		"schema_id": *deleteSchemaOptions.SchemaID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteSchemaOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteSchemaOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*deleteSchemaOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_schema", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListTables : List all tables
// List all tables in a schema in a catalog for a given engine.
func (watsonxData *WatsonxDataV2) ListTables(listTablesOptions *ListTablesOptions) (result *TableCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListTablesWithContext(context.Background(), listTablesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListTablesWithContext is an alternate form of the ListTables method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListTablesWithContext(ctx context.Context, listTablesOptions *ListTablesOptions) (result *TableCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTablesOptions, "listTablesOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listTablesOptions, "listTablesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *listTablesOptions.CatalogID,
		"schema_id": *listTablesOptions.SchemaID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listTablesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListTables")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listTablesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listTablesOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*listTablesOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_tables", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTableCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetTable : Get table details
// Get details of a given table in a catalog and schema.
func (watsonxData *WatsonxDataV2) GetTable(getTableOptions *GetTableOptions) (result *Table, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetTableWithContext(context.Background(), getTableOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetTableWithContext is an alternate form of the GetTable method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetTableWithContext(ctx context.Context, getTableOptions *GetTableOptions) (result *Table, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTableOptions, "getTableOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getTableOptions, "getTableOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *getTableOptions.CatalogID,
		"schema_id": *getTableOptions.SchemaID,
		"table_id": *getTableOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getTableOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetTable")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getTableOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getTableOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*getTableOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_table", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTable)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteTable : Delete table
// Delete table for a given schema and catalog.
func (watsonxData *WatsonxDataV2) DeleteTable(deleteTableOptions *DeleteTableOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteTableWithContext(context.Background(), deleteTableOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteTableWithContext is an alternate form of the DeleteTable method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteTableWithContext(ctx context.Context, deleteTableOptions *DeleteTableOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTableOptions, "deleteTableOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteTableOptions, "deleteTableOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *deleteTableOptions.CatalogID,
		"schema_id": *deleteTableOptions.SchemaID,
		"table_id": *deleteTableOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteTableOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteTable")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteTableOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteTableOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*deleteTableOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_table", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// RenameTable : Rename table
// Rename table.
func (watsonxData *WatsonxDataV2) RenameTable(renameTableOptions *RenameTableOptions) (result *Table, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.RenameTableWithContext(context.Background(), renameTableOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// RenameTableWithContext is an alternate form of the RenameTable method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RenameTableWithContext(ctx context.Context, renameTableOptions *RenameTableOptions) (result *Table, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(renameTableOptions, "renameTableOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(renameTableOptions, "renameTableOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *renameTableOptions.CatalogID,
		"schema_id": *renameTableOptions.SchemaID,
		"table_id": *renameTableOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range renameTableOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RenameTable")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if renameTableOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*renameTableOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*renameTableOptions.EngineID))

	_, err = builder.SetBodyContentJSON(renameTableOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "rename_table", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTable)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListColumns : List all columns of a table
// List all columns of a table in a given a schema for a given catalog.
func (watsonxData *WatsonxDataV2) ListColumns(listColumnsOptions *ListColumnsOptions) (result *ColumnCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListColumnsWithContext(context.Background(), listColumnsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListColumnsWithContext is an alternate form of the ListColumns method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListColumnsWithContext(ctx context.Context, listColumnsOptions *ListColumnsOptions) (result *ColumnCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listColumnsOptions, "listColumnsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listColumnsOptions, "listColumnsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *listColumnsOptions.CatalogID,
		"schema_id": *listColumnsOptions.SchemaID,
		"table_id": *listColumnsOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/columns`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listColumnsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListColumns")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listColumnsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listColumnsOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*listColumnsOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_columns", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalColumnCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateColumns : Add column(s)
// Add one or multiple columns to a table in a schema for a given catalog.
func (watsonxData *WatsonxDataV2) CreateColumns(createColumnsOptions *CreateColumnsOptions) (result *ColumnCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateColumnsWithContext(context.Background(), createColumnsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateColumnsWithContext is an alternate form of the CreateColumns method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateColumnsWithContext(ctx context.Context, createColumnsOptions *CreateColumnsOptions) (result *ColumnCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createColumnsOptions, "createColumnsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createColumnsOptions, "createColumnsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *createColumnsOptions.CatalogID,
		"schema_id": *createColumnsOptions.SchemaID,
		"table_id": *createColumnsOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/columns`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createColumnsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateColumns")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createColumnsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createColumnsOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*createColumnsOptions.EngineID))

	body := make(map[string]interface{})
	if createColumnsOptions.Columns != nil {
		body["columns"] = createColumnsOptions.Columns
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_columns", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalColumnCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteColumn : Delete column
// Delete column in a table for a given schema and catalog.
func (watsonxData *WatsonxDataV2) DeleteColumn(deleteColumnOptions *DeleteColumnOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteColumnWithContext(context.Background(), deleteColumnOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteColumnWithContext is an alternate form of the DeleteColumn method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteColumnWithContext(ctx context.Context, deleteColumnOptions *DeleteColumnOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteColumnOptions, "deleteColumnOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteColumnOptions, "deleteColumnOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *deleteColumnOptions.CatalogID,
		"schema_id": *deleteColumnOptions.SchemaID,
		"table_id": *deleteColumnOptions.TableID,
		"column_id": *deleteColumnOptions.ColumnID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/columns/{column_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteColumnOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteColumn")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteColumnOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteColumnOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*deleteColumnOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_column", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateColumn : Alter column
// Update the given column - rename column.
func (watsonxData *WatsonxDataV2) UpdateColumn(updateColumnOptions *UpdateColumnOptions) (result *Column, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.UpdateColumnWithContext(context.Background(), updateColumnOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateColumnWithContext is an alternate form of the UpdateColumn method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateColumnWithContext(ctx context.Context, updateColumnOptions *UpdateColumnOptions) (result *Column, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateColumnOptions, "updateColumnOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateColumnOptions, "updateColumnOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *updateColumnOptions.CatalogID,
		"schema_id": *updateColumnOptions.SchemaID,
		"table_id": *updateColumnOptions.TableID,
		"column_id": *updateColumnOptions.ColumnID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/columns/{column_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateColumnOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateColumn")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updateColumnOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateColumnOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*updateColumnOptions.EngineID))

	_, err = builder.SetBodyContentJSON(updateColumnOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_column", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalColumn)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListTableSnapshots : Get table snapshots
// List all table snapshots.
func (watsonxData *WatsonxDataV2) ListTableSnapshots(listTableSnapshotsOptions *ListTableSnapshotsOptions) (result *TableSnapshotCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListTableSnapshotsWithContext(context.Background(), listTableSnapshotsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListTableSnapshotsWithContext is an alternate form of the ListTableSnapshots method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListTableSnapshotsWithContext(ctx context.Context, listTableSnapshotsOptions *ListTableSnapshotsOptions) (result *TableSnapshotCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTableSnapshotsOptions, "listTableSnapshotsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listTableSnapshotsOptions, "listTableSnapshotsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *listTableSnapshotsOptions.CatalogID,
		"schema_id": *listTableSnapshotsOptions.SchemaID,
		"table_id": *listTableSnapshotsOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/snapshots`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listTableSnapshotsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListTableSnapshots")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listTableSnapshotsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listTableSnapshotsOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*listTableSnapshotsOptions.EngineID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_table_snapshots", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTableSnapshotCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// RollbackTable : Rollback table to snapshot
// Rollback table to a snapshot.
func (watsonxData *WatsonxDataV2) RollbackTable(rollbackTableOptions *RollbackTableOptions) (result *ReplaceSnapshotCreatedBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.RollbackTableWithContext(context.Background(), rollbackTableOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// RollbackTableWithContext is an alternate form of the RollbackTable method which supports a Context parameter
func (watsonxData *WatsonxDataV2) RollbackTableWithContext(ctx context.Context, rollbackTableOptions *RollbackTableOptions) (result *ReplaceSnapshotCreatedBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(rollbackTableOptions, "rollbackTableOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(rollbackTableOptions, "rollbackTableOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *rollbackTableOptions.CatalogID,
		"schema_id": *rollbackTableOptions.SchemaID,
		"table_id": *rollbackTableOptions.TableID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/schemas/{schema_id}/tables/{table_id}/rollback`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range rollbackTableOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "RollbackTable")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if rollbackTableOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*rollbackTableOptions.AuthInstanceID))
	}

	builder.AddQuery("engine_id", fmt.Sprint(*rollbackTableOptions.EngineID))

	body := make(map[string]interface{})
	if rollbackTableOptions.SnapshotID != nil {
		body["snapshot_id"] = rollbackTableOptions.SnapshotID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "rollback_table", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReplaceSnapshotCreatedBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateSyncCatalog : External Iceberg table registration
// Synchronize the external Iceberg table registration for a catalog identified by catalog_id.
func (watsonxData *WatsonxDataV2) UpdateSyncCatalog(updateSyncCatalogOptions *UpdateSyncCatalogOptions) (result *UpdateSyncCatalogOKBody, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.UpdateSyncCatalogWithContext(context.Background(), updateSyncCatalogOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateSyncCatalogWithContext is an alternate form of the UpdateSyncCatalog method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateSyncCatalogWithContext(ctx context.Context, updateSyncCatalogOptions *UpdateSyncCatalogOptions) (result *UpdateSyncCatalogOKBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSyncCatalogOptions, "updateSyncCatalogOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateSyncCatalogOptions, "updateSyncCatalogOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"catalog_id": *updateSyncCatalogOptions.CatalogID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/catalogs/{catalog_id}/sync`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateSyncCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateSyncCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updateSyncCatalogOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateSyncCatalogOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateSyncCatalogOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_sync_catalog", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateSyncCatalogOKBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListMilvusServices : Get list of milvus services
// Get list milvus services.
func (watsonxData *WatsonxDataV2) ListMilvusServices(listMilvusServicesOptions *ListMilvusServicesOptions) (result *MilvusServiceCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListMilvusServicesWithContext(context.Background(), listMilvusServicesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListMilvusServicesWithContext is an alternate form of the ListMilvusServices method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListMilvusServicesWithContext(ctx context.Context, listMilvusServicesOptions *ListMilvusServicesOptions) (result *MilvusServiceCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listMilvusServicesOptions, "listMilvusServicesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/milvus_services`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listMilvusServicesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListMilvusServices")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listMilvusServicesOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listMilvusServicesOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_milvus_services", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMilvusServiceCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateMilvusService : Create milvus service
// Create milvus service.
func (watsonxData *WatsonxDataV2) CreateMilvusService(createMilvusServiceOptions *CreateMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.CreateMilvusServiceWithContext(context.Background(), createMilvusServiceOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateMilvusServiceWithContext is an alternate form of the CreateMilvusService method which supports a Context parameter
func (watsonxData *WatsonxDataV2) CreateMilvusServiceWithContext(ctx context.Context, createMilvusServiceOptions *CreateMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createMilvusServiceOptions, "createMilvusServiceOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createMilvusServiceOptions, "createMilvusServiceOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/milvus_services`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createMilvusServiceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "CreateMilvusService")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createMilvusServiceOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*createMilvusServiceOptions.AuthInstanceID))
	}

	body := make(map[string]interface{})
	if createMilvusServiceOptions.Origin != nil {
		body["origin"] = createMilvusServiceOptions.Origin
	}
	if createMilvusServiceOptions.Description != nil {
		body["description"] = createMilvusServiceOptions.Description
	}
	if createMilvusServiceOptions.ServiceDisplayName != nil {
		body["service_display_name"] = createMilvusServiceOptions.ServiceDisplayName
	}
	if createMilvusServiceOptions.Tags != nil {
		body["tags"] = createMilvusServiceOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_milvus_service", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMilvusService)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetMilvusService : Get milvus service
// Get milvus service.
func (watsonxData *WatsonxDataV2) GetMilvusService(getMilvusServiceOptions *GetMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.GetMilvusServiceWithContext(context.Background(), getMilvusServiceOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetMilvusServiceWithContext is an alternate form of the GetMilvusService method which supports a Context parameter
func (watsonxData *WatsonxDataV2) GetMilvusServiceWithContext(ctx context.Context, getMilvusServiceOptions *GetMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getMilvusServiceOptions, "getMilvusServiceOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getMilvusServiceOptions, "getMilvusServiceOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"service_id": *getMilvusServiceOptions.ServiceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/milvus_services/{service_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getMilvusServiceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "GetMilvusService")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getMilvusServiceOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*getMilvusServiceOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_milvus_service", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMilvusService)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteMilvusService : Delete milvus service
// Delete milvus service.
func (watsonxData *WatsonxDataV2) DeleteMilvusService(deleteMilvusServiceOptions *DeleteMilvusServiceOptions) (response *core.DetailedResponse, err error) {
	response, err = watsonxData.DeleteMilvusServiceWithContext(context.Background(), deleteMilvusServiceOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteMilvusServiceWithContext is an alternate form of the DeleteMilvusService method which supports a Context parameter
func (watsonxData *WatsonxDataV2) DeleteMilvusServiceWithContext(ctx context.Context, deleteMilvusServiceOptions *DeleteMilvusServiceOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteMilvusServiceOptions, "deleteMilvusServiceOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteMilvusServiceOptions, "deleteMilvusServiceOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"service_id": *deleteMilvusServiceOptions.ServiceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/milvus_services/{service_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteMilvusServiceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "DeleteMilvusService")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteMilvusServiceOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*deleteMilvusServiceOptions.AuthInstanceID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = watsonxData.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_milvus_service", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateMilvusService : Update milvus service
// Update details of milvus service.
func (watsonxData *WatsonxDataV2) UpdateMilvusService(updateMilvusServiceOptions *UpdateMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.UpdateMilvusServiceWithContext(context.Background(), updateMilvusServiceOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateMilvusServiceWithContext is an alternate form of the UpdateMilvusService method which supports a Context parameter
func (watsonxData *WatsonxDataV2) UpdateMilvusServiceWithContext(ctx context.Context, updateMilvusServiceOptions *UpdateMilvusServiceOptions) (result *MilvusService, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateMilvusServiceOptions, "updateMilvusServiceOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateMilvusServiceOptions, "updateMilvusServiceOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"service_id": *updateMilvusServiceOptions.ServiceID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/milvus_services/{service_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateMilvusServiceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "UpdateMilvusService")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updateMilvusServiceOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*updateMilvusServiceOptions.AuthInstanceID))
	}

	_, err = builder.SetBodyContentJSON(updateMilvusServiceOptions.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_milvus_service", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMilvusService)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListIngestionJobs : Get ingestion jobs
// Get list of ingestion jobs.
func (watsonxData *WatsonxDataV2) ListIngestionJobs(listIngestionJobsOptions *ListIngestionJobsOptions) (result *IngestionJobCollection, response *core.DetailedResponse, err error) {
	result, response, err = watsonxData.ListIngestionJobsWithContext(context.Background(), listIngestionJobsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListIngestionJobsWithContext is an alternate form of the ListIngestionJobs method which supports a Context parameter
func (watsonxData *WatsonxDataV2) ListIngestionJobsWithContext(ctx context.Context, listIngestionJobsOptions *ListIngestionJobsOptions) (result *IngestionJobCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listIngestionJobsOptions, "listIngestionJobsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listIngestionJobsOptions, "listIngestionJobsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonxData.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonxData.Service.Options.URL, `/ingestion_jobs`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listIngestionJobsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watsonx_data", "V2", "ListIngestionJobs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listIngestionJobsOptions.AuthInstanceID != nil {
		builder.AddHeader("AuthInstanceId", fmt.Sprint(*listIngestionJobsOptions.AuthInstanceID))
	}

	if listIngestionJobsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listIngestionJobsOptions.Start))
	}
	if listIngestionJobsOptions.JobsPerPage != nil {
		builder.AddQuery("jobs_per_page", fmt.Sprint(*listIngestionJobsOptions.JobsPerPage))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonxData.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_ingestion_jobs", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIngestionJobCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "2.0.0")
}

// AddPrestissimoEngineCatalogsOptions : The AddPrestissimoEngineCatalogs options.
type AddPrestissimoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// catalog names.
	CatalogNames *string `json:"catalog_names,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewAddPrestissimoEngineCatalogsOptions : Instantiate AddPrestissimoEngineCatalogsOptions
func (*WatsonxDataV2) NewAddPrestissimoEngineCatalogsOptions(engineID string) *AddPrestissimoEngineCatalogsOptions {
	return &AddPrestissimoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *AddPrestissimoEngineCatalogsOptions) SetEngineID(engineID string) *AddPrestissimoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogNames : Allow user to set CatalogNames
func (_options *AddPrestissimoEngineCatalogsOptions) SetCatalogNames(catalogNames string) *AddPrestissimoEngineCatalogsOptions {
	_options.CatalogNames = core.StringPtr(catalogNames)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *AddPrestissimoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *AddPrestissimoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddPrestissimoEngineCatalogsOptions) SetHeaders(param map[string]string) *AddPrestissimoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// AddPrestoEngineCatalogsOptions : The AddPrestoEngineCatalogs options.
type AddPrestoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// catalog names.
	CatalogNames *string `json:"catalog_names,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewAddPrestoEngineCatalogsOptions : Instantiate AddPrestoEngineCatalogsOptions
func (*WatsonxDataV2) NewAddPrestoEngineCatalogsOptions(engineID string) *AddPrestoEngineCatalogsOptions {
	return &AddPrestoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *AddPrestoEngineCatalogsOptions) SetEngineID(engineID string) *AddPrestoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogNames : Allow user to set CatalogNames
func (_options *AddPrestoEngineCatalogsOptions) SetCatalogNames(catalogNames string) *AddPrestoEngineCatalogsOptions {
	_options.CatalogNames = core.StringPtr(catalogNames)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *AddPrestoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *AddPrestoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddPrestoEngineCatalogsOptions) SetHeaders(param map[string]string) *AddPrestoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// AddSparkEngineCatalogsOptions : The AddSparkEngineCatalogs options.
type AddSparkEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// catalog names.
	CatalogNames *string `json:"catalog_names,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewAddSparkEngineCatalogsOptions : Instantiate AddSparkEngineCatalogsOptions
func (*WatsonxDataV2) NewAddSparkEngineCatalogsOptions(engineID string) *AddSparkEngineCatalogsOptions {
	return &AddSparkEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *AddSparkEngineCatalogsOptions) SetEngineID(engineID string) *AddSparkEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogNames : Allow user to set CatalogNames
func (_options *AddSparkEngineCatalogsOptions) SetCatalogNames(catalogNames string) *AddSparkEngineCatalogsOptions {
	_options.CatalogNames = core.StringPtr(catalogNames)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *AddSparkEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *AddSparkEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddSparkEngineCatalogsOptions) SetHeaders(param map[string]string) *AddSparkEngineCatalogsOptions {
	options.Headers = param
	return options
}

// BucketCatalog : bucket catalog.
type BucketCatalog struct {
	// catalog name.
	CatalogName *string `json:"catalog_name,omitempty"`

	// catalog tags.
	CatalogTags []string `json:"catalog_tags,omitempty"`

	// catalog type.
	CatalogType *string `json:"catalog_type,omitempty"`
}

// UnmarshalBucketCatalog unmarshals an instance of BucketCatalog from the specified map of raw messages.
func UnmarshalBucketCatalog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketCatalog)
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_tags", &obj.CatalogTags)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_type", &obj.CatalogType)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketDetails : bucket details.
type BucketDetails struct {
	// Access key ID, encrypted during bucket registration.
	AccessKey *string `json:"access_key,omitempty"`

	// actual bucket name.
	BucketName *string `json:"bucket_name" validate:"required"`

	// Cos endpoint.
	Endpoint *string `json:"endpoint,omitempty"`

	// Secret access key, encrypted during bucket registration.
	SecretKey *string `json:"secret_key,omitempty"`
}

// NewBucketDetails : Instantiate BucketDetails (Generic Model Constructor)
func (*WatsonxDataV2) NewBucketDetails(bucketName string) (_model *BucketDetails, err error) {
	_model = &BucketDetails{
		BucketName: core.StringPtr(bucketName),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalBucketDetails unmarshals an instance of BucketDetails from the specified map of raw messages.
func UnmarshalBucketDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketDetails)
	err = core.UnmarshalPrimitive(m, "access_key", &obj.AccessKey)
	if err != nil {
		err = core.SDKErrorf(err, "", "access_key-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_name", &obj.BucketName)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "secret_key", &obj.SecretKey)
	if err != nil {
		err = core.SDKErrorf(err, "", "secret_key-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the BucketDetails
func (bucketDetails *BucketDetails) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(bucketDetails.AccessKey) {
		_patch["access_key"] = bucketDetails.AccessKey
	}
	if !core.IsNil(bucketDetails.BucketName) {
		_patch["bucket_name"] = bucketDetails.BucketName
	}
	if !core.IsNil(bucketDetails.Endpoint) {
		_patch["endpoint"] = bucketDetails.Endpoint
	}
	if !core.IsNil(bucketDetails.SecretKey) {
		_patch["secret_key"] = bucketDetails.SecretKey
	}

	return
}

// BucketRegistration : Bucket.
type BucketRegistration struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// bucket catalog.
	AssociatedCatalog *BucketCatalog `json:"associated_catalog" validate:"required"`

	// bucket details.
	BucketDetails *BucketDetails `json:"bucket_details,omitempty"`

	// bucket display name.
	BucketDisplayName *string `json:"bucket_display_name,omitempty"`

	// bucket ID auto generated during bucket registration.
	BucketID *string `json:"bucket_id,omitempty"`

	// bucket type.
	BucketType *string `json:"bucket_type" validate:"required"`

	// Username who created the bucket.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Creation date.
	CreatedOn *string `json:"created_on" validate:"required"`

	// bucket description.
	Description *string `json:"description" validate:"required"`

	// managed by.
	ManagedBy *string `json:"managed_by" validate:"required"`

	// Region where the bucket is located.
	Region *string `json:"region,omitempty"`

	// mark bucket active or inactive.
	State *string `json:"state" validate:"required"`

	// tags.
	Tags []string `json:"tags,omitempty"`
}

// Constants associated with the BucketRegistration.BucketType property.
// bucket type.
const (
	BucketRegistration_BucketType_AmazonS3 = "amazon_s3"
	BucketRegistration_BucketType_AwsS3 = "aws_s3"
	BucketRegistration_BucketType_IbmCeph = "ibm_ceph"
	BucketRegistration_BucketType_IbmCos = "ibm_cos"
	BucketRegistration_BucketType_Minio = "minio"
)

// Constants associated with the BucketRegistration.ManagedBy property.
// managed by.
const (
	BucketRegistration_ManagedBy_Customer = "customer"
	BucketRegistration_ManagedBy_Ibm = "ibm"
)

// Constants associated with the BucketRegistration.State property.
// mark bucket active or inactive.
const (
	BucketRegistration_State_Active = "active"
	BucketRegistration_State_Inactive = "inactive"
)

// UnmarshalBucketRegistration unmarshals an instance of BucketRegistration from the specified map of raw messages.
func UnmarshalBucketRegistration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketRegistration)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		err = core.SDKErrorf(err, "", "actions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "associated_catalog", &obj.AssociatedCatalog, UnmarshalBucketCatalog)
	if err != nil {
		err = core.SDKErrorf(err, "", "associated_catalog-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "bucket_details", &obj.BucketDetails, UnmarshalBucketDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_display_name", &obj.BucketDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_id", &obj.BucketID)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_type", &obj.BucketType)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_on-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "managed_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		err = core.SDKErrorf(err, "", "region-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketRegistrationCollection : List bucket registrations.
type BucketRegistrationCollection struct {
	// Buckets.
	BucketRegistrations []BucketRegistration `json:"bucket_registrations,omitempty"`
}

// UnmarshalBucketRegistrationCollection unmarshals an instance of BucketRegistrationCollection from the specified map of raw messages.
func UnmarshalBucketRegistrationCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketRegistrationCollection)
	err = core.UnmarshalModel(m, "bucket_registrations", &obj.BucketRegistrations, UnmarshalBucketRegistration)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_registrations-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketRegistrationObjectCollection : List bucket objects.
type BucketRegistrationObjectCollection struct {
	// bucket object.
	Objects []string `json:"objects,omitempty"`
}

// UnmarshalBucketRegistrationObjectCollection unmarshals an instance of BucketRegistrationObjectCollection from the specified map of raw messages.
func UnmarshalBucketRegistrationObjectCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketRegistrationObjectCollection)
	err = core.UnmarshalPrimitive(m, "objects", &obj.Objects)
	if err != nil {
		err = core.SDKErrorf(err, "", "objects-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketRegistrationPatch : Update bucket parameters.
type BucketRegistrationPatch struct {
	// bucket details.
	BucketDetails *BucketDetails `json:"bucket_details,omitempty"`

	// bucket display name.
	BucketDisplayName *string `json:"bucket_display_name,omitempty"`

	// modified description.
	Description *string `json:"description,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`
}

// UnmarshalBucketRegistrationPatch unmarshals an instance of BucketRegistrationPatch from the specified map of raw messages.
func UnmarshalBucketRegistrationPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketRegistrationPatch)
	err = core.UnmarshalModel(m, "bucket_details", &obj.BucketDetails, UnmarshalBucketDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_display_name", &obj.BucketDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the BucketRegistrationPatch
func (bucketRegistrationPatch *BucketRegistrationPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(bucketRegistrationPatch.BucketDetails) {
		_patch["bucket_details"] = bucketRegistrationPatch.BucketDetails.asPatch()
	}
	if !core.IsNil(bucketRegistrationPatch.BucketDisplayName) {
		_patch["bucket_display_name"] = bucketRegistrationPatch.BucketDisplayName
	}
	if !core.IsNil(bucketRegistrationPatch.Description) {
		_patch["description"] = bucketRegistrationPatch.Description
	}
	if !core.IsNil(bucketRegistrationPatch.Tags) {
		_patch["tags"] = bucketRegistrationPatch.Tags
	}

	return
}

// Catalog : Define the catalog details.
type Catalog struct {
	// list of allowed actions.
	Actions []string `json:"actions,omitempty"`

	// Associated buckets items.
	AssociatedBuckets []string `json:"associated_buckets,omitempty"`

	// Associated databases items.
	AssociatedDatabases []string `json:"associated_databases,omitempty"`

	// Associated engines items.
	AssociatedEngines []string `json:"associated_engines,omitempty"`

	// Name for the catalog.
	CatalogName *string `json:"catalog_name,omitempty"`

	// Table type.
	CatalogType *string `json:"catalog_type,omitempty"`

	// Created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// Description.
	Description *string `json:"description,omitempty"`

	// IBM thrift uri hostname.
	Hostname *string `json:"hostname,omitempty"`

	// Last sync time.
	LastSyncAt *string `json:"last_sync_at,omitempty"`

	// Managed by.
	ManagedBy *string `json:"managed_by,omitempty"`

	// Catalog name.
	Metastore *string `json:"metastore,omitempty"`

	// IBM thrift uri port.
	Port *string `json:"port,omitempty"`

	// Catalog status.
	Status *string `json:"status,omitempty"`

	// Sync description.
	SyncDescription *string `json:"sync_description,omitempty"`

	// Tables not sync because data is corrupted.
	SyncException []string `json:"sync_exception,omitempty"`

	// Sync status.
	SyncStatus *string `json:"sync_status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Customer thrift uri.
	ThriftURI *string `json:"thrift_uri,omitempty"`
}

// Constants associated with the Catalog.ManagedBy property.
// Managed by.
const (
	Catalog_ManagedBy_Customer = "customer"
	Catalog_ManagedBy_Ibm = "ibm"
)

// UnmarshalCatalog unmarshals an instance of Catalog from the specified map of raw messages.
func UnmarshalCatalog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Catalog)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		err = core.SDKErrorf(err, "", "actions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_buckets", &obj.AssociatedBuckets)
	if err != nil {
		err = core.SDKErrorf(err, "", "associated_buckets-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_databases", &obj.AssociatedDatabases)
	if err != nil {
		err = core.SDKErrorf(err, "", "associated_databases-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_engines", &obj.AssociatedEngines)
	if err != nil {
		err = core.SDKErrorf(err, "", "associated_engines-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_type", &obj.CatalogType)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_on-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		err = core.SDKErrorf(err, "", "hostname-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_sync_at", &obj.LastSyncAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_sync_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "managed_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "metastore", &obj.Metastore)
	if err != nil {
		err = core.SDKErrorf(err, "", "metastore-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "sync_description", &obj.SyncDescription)
	if err != nil {
		err = core.SDKErrorf(err, "", "sync_description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "sync_exception", &obj.SyncException)
	if err != nil {
		err = core.SDKErrorf(err, "", "sync_exception-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "sync_status", &obj.SyncStatus)
	if err != nil {
		err = core.SDKErrorf(err, "", "sync_status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "thrift_uri", &obj.ThriftURI)
	if err != nil {
		err = core.SDKErrorf(err, "", "thrift_uri-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogCollection : GetCatalogs OK.
type CatalogCollection struct {
	// Catalogs.
	Catalogs []Catalog `json:"catalogs,omitempty"`
}

// UnmarshalCatalogCollection unmarshals an instance of CatalogCollection from the specified map of raw messages.
func UnmarshalCatalogCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogCollection)
	err = core.UnmarshalModel(m, "catalogs", &obj.Catalogs, UnmarshalCatalog)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalogs-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Column : Column.
type Column struct {
	// Column name.
	ColumnName *string `json:"column_name,omitempty"`

	// Comment.
	Comment *string `json:"comment,omitempty"`

	// Extra.
	Extra *string `json:"extra,omitempty"`

	// length.
	Length *string `json:"length,omitempty"`

	// scale.
	Scale *string `json:"scale,omitempty"`

	// Data type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalColumn unmarshals an instance of Column from the specified map of raw messages.
func UnmarshalColumn(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Column)
	err = core.UnmarshalPrimitive(m, "column_name", &obj.ColumnName)
	if err != nil {
		err = core.SDKErrorf(err, "", "column_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		err = core.SDKErrorf(err, "", "comment-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "extra", &obj.Extra)
	if err != nil {
		err = core.SDKErrorf(err, "", "extra-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "length", &obj.Length)
	if err != nil {
		err = core.SDKErrorf(err, "", "length-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "scale", &obj.Scale)
	if err != nil {
		err = core.SDKErrorf(err, "", "scale-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ColumnCollection : list of columns in a table.
type ColumnCollection struct {
	// List of the columns present in the table.
	Columns []Column `json:"columns,omitempty"`
}

// UnmarshalColumnCollection unmarshals an instance of ColumnCollection from the specified map of raw messages.
func UnmarshalColumnCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ColumnCollection)
	err = core.UnmarshalModel(m, "columns", &obj.Columns, UnmarshalColumn)
	if err != nil {
		err = core.SDKErrorf(err, "", "columns-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ColumnPatch : list of columns to be added to a table.
type ColumnPatch struct {
	// Column name.
	ColumnName *string `json:"column_name,omitempty"`
}

// UnmarshalColumnPatch unmarshals an instance of ColumnPatch from the specified map of raw messages.
func UnmarshalColumnPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ColumnPatch)
	err = core.UnmarshalPrimitive(m, "column_name", &obj.ColumnName)
	if err != nil {
		err = core.SDKErrorf(err, "", "column_name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the ColumnPatch
func (columnPatch *ColumnPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(columnPatch.ColumnName) {
		_patch["column_name"] = columnPatch.ColumnName
	}

	return
}

// CreateActivateBucketCreatedBody : Activate bucket.
type CreateActivateBucketCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateActivateBucketCreatedBody unmarshals an instance of CreateActivateBucketCreatedBody from the specified map of raw messages.
func UnmarshalCreateActivateBucketCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateActivateBucketCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateActivateBucketOptions : The CreateActivateBucket options.
type CreateActivateBucketOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateActivateBucketOptions : Instantiate CreateActivateBucketOptions
func (*WatsonxDataV2) NewCreateActivateBucketOptions(bucketID string) *CreateActivateBucketOptions {
	return &CreateActivateBucketOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *CreateActivateBucketOptions) SetBucketID(bucketID string) *CreateActivateBucketOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateActivateBucketOptions) SetAuthInstanceID(authInstanceID string) *CreateActivateBucketOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateActivateBucketOptions) SetHeaders(param map[string]string) *CreateActivateBucketOptions {
	options.Headers = param
	return options
}

// CreateBucketRegistrationOptions : The CreateBucketRegistration options.
type CreateBucketRegistrationOptions struct {
	// bucket details.
	BucketDetails *BucketDetails `json:"bucket_details" validate:"required"`

	// bucket type.
	BucketType *string `json:"bucket_type" validate:"required"`

	// bucket description.
	Description *string `json:"description" validate:"required"`

	// managed by.
	ManagedBy *string `json:"managed_by" validate:"required"`

	// bucket catalog.
	AssociatedCatalog *BucketCatalog `json:"associated_catalog,omitempty"`

	// bucket display name.
	BucketDisplayName *string `json:"bucket_display_name,omitempty"`

	// region where the bucket is located.
	Region *string `json:"region,omitempty"`

	// tags.
	Tags []string `json:"tags,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateBucketRegistrationOptions.BucketType property.
// bucket type.
const (
	CreateBucketRegistrationOptions_BucketType_AwsS3 = "aws_s3"
	CreateBucketRegistrationOptions_BucketType_IbmCeph = "ibm_ceph"
	CreateBucketRegistrationOptions_BucketType_IbmCos = "ibm_cos"
	CreateBucketRegistrationOptions_BucketType_Minio = "minio"
)

// Constants associated with the CreateBucketRegistrationOptions.ManagedBy property.
// managed by.
const (
	CreateBucketRegistrationOptions_ManagedBy_Customer = "customer"
	CreateBucketRegistrationOptions_ManagedBy_Ibm = "ibm"
)

// NewCreateBucketRegistrationOptions : Instantiate CreateBucketRegistrationOptions
func (*WatsonxDataV2) NewCreateBucketRegistrationOptions(bucketDetails *BucketDetails, bucketType string, description string, managedBy string) *CreateBucketRegistrationOptions {
	return &CreateBucketRegistrationOptions{
		BucketDetails: bucketDetails,
		BucketType: core.StringPtr(bucketType),
		Description: core.StringPtr(description),
		ManagedBy: core.StringPtr(managedBy),
	}
}

// SetBucketDetails : Allow user to set BucketDetails
func (_options *CreateBucketRegistrationOptions) SetBucketDetails(bucketDetails *BucketDetails) *CreateBucketRegistrationOptions {
	_options.BucketDetails = bucketDetails
	return _options
}

// SetBucketType : Allow user to set BucketType
func (_options *CreateBucketRegistrationOptions) SetBucketType(bucketType string) *CreateBucketRegistrationOptions {
	_options.BucketType = core.StringPtr(bucketType)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateBucketRegistrationOptions) SetDescription(description string) *CreateBucketRegistrationOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetManagedBy : Allow user to set ManagedBy
func (_options *CreateBucketRegistrationOptions) SetManagedBy(managedBy string) *CreateBucketRegistrationOptions {
	_options.ManagedBy = core.StringPtr(managedBy)
	return _options
}

// SetAssociatedCatalog : Allow user to set AssociatedCatalog
func (_options *CreateBucketRegistrationOptions) SetAssociatedCatalog(associatedCatalog *BucketCatalog) *CreateBucketRegistrationOptions {
	_options.AssociatedCatalog = associatedCatalog
	return _options
}

// SetBucketDisplayName : Allow user to set BucketDisplayName
func (_options *CreateBucketRegistrationOptions) SetBucketDisplayName(bucketDisplayName string) *CreateBucketRegistrationOptions {
	_options.BucketDisplayName = core.StringPtr(bucketDisplayName)
	return _options
}

// SetRegion : Allow user to set Region
func (_options *CreateBucketRegistrationOptions) SetRegion(region string) *CreateBucketRegistrationOptions {
	_options.Region = core.StringPtr(region)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateBucketRegistrationOptions) SetTags(tags []string) *CreateBucketRegistrationOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateBucketRegistrationOptions) SetAuthInstanceID(authInstanceID string) *CreateBucketRegistrationOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateBucketRegistrationOptions) SetHeaders(param map[string]string) *CreateBucketRegistrationOptions {
	options.Headers = param
	return options
}

// CreateColumnsOptions : The CreateColumns options.
type CreateColumnsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded schema name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// List of the tables present in the schema.
	Columns []Column `json:"columns,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateColumnsOptions : Instantiate CreateColumnsOptions
func (*WatsonxDataV2) NewCreateColumnsOptions(engineID string, catalogID string, schemaID string, tableID string) *CreateColumnsOptions {
	return &CreateColumnsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateColumnsOptions) SetEngineID(engineID string) *CreateColumnsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *CreateColumnsOptions) SetCatalogID(catalogID string) *CreateColumnsOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *CreateColumnsOptions) SetSchemaID(schemaID string) *CreateColumnsOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *CreateColumnsOptions) SetTableID(tableID string) *CreateColumnsOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetColumns : Allow user to set Columns
func (_options *CreateColumnsOptions) SetColumns(columns []Column) *CreateColumnsOptions {
	_options.Columns = columns
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateColumnsOptions) SetAuthInstanceID(authInstanceID string) *CreateColumnsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateColumnsOptions) SetHeaders(param map[string]string) *CreateColumnsOptions {
	options.Headers = param
	return options
}

// CreateDatabaseRegistrationOptions : The CreateDatabaseRegistration options.
type CreateDatabaseRegistrationOptions struct {
	// Database display name.
	DatabaseDisplayName *string `json:"database_display_name" validate:"required"`

	// Connector type.
	DatabaseType *string `json:"database_type" validate:"required"`

	// database catalog.
	AssociatedCatalog *DatabaseCatalog `json:"associated_catalog,omitempty"`

	// Created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// database details.
	DatabaseDetails *DatabaseDetails `json:"database_details,omitempty"`

	// This will hold all the properties for a custom database.
	DatabaseProperties []DatabaseRegistrationPrototypeDatabasePropertiesItems `json:"database_properties,omitempty"`

	// Database description.
	Description *string `json:"description,omitempty"`

	// tags.
	Tags []string `json:"tags,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateDatabaseRegistrationOptions : Instantiate CreateDatabaseRegistrationOptions
func (*WatsonxDataV2) NewCreateDatabaseRegistrationOptions(databaseDisplayName string, databaseType string) *CreateDatabaseRegistrationOptions {
	return &CreateDatabaseRegistrationOptions{
		DatabaseDisplayName: core.StringPtr(databaseDisplayName),
		DatabaseType: core.StringPtr(databaseType),
	}
}

// SetDatabaseDisplayName : Allow user to set DatabaseDisplayName
func (_options *CreateDatabaseRegistrationOptions) SetDatabaseDisplayName(databaseDisplayName string) *CreateDatabaseRegistrationOptions {
	_options.DatabaseDisplayName = core.StringPtr(databaseDisplayName)
	return _options
}

// SetDatabaseType : Allow user to set DatabaseType
func (_options *CreateDatabaseRegistrationOptions) SetDatabaseType(databaseType string) *CreateDatabaseRegistrationOptions {
	_options.DatabaseType = core.StringPtr(databaseType)
	return _options
}

// SetAssociatedCatalog : Allow user to set AssociatedCatalog
func (_options *CreateDatabaseRegistrationOptions) SetAssociatedCatalog(associatedCatalog *DatabaseCatalog) *CreateDatabaseRegistrationOptions {
	_options.AssociatedCatalog = associatedCatalog
	return _options
}

// SetCreatedOn : Allow user to set CreatedOn
func (_options *CreateDatabaseRegistrationOptions) SetCreatedOn(createdOn string) *CreateDatabaseRegistrationOptions {
	_options.CreatedOn = core.StringPtr(createdOn)
	return _options
}

// SetDatabaseDetails : Allow user to set DatabaseDetails
func (_options *CreateDatabaseRegistrationOptions) SetDatabaseDetails(databaseDetails *DatabaseDetails) *CreateDatabaseRegistrationOptions {
	_options.DatabaseDetails = databaseDetails
	return _options
}

// SetDatabaseProperties : Allow user to set DatabaseProperties
func (_options *CreateDatabaseRegistrationOptions) SetDatabaseProperties(databaseProperties []DatabaseRegistrationPrototypeDatabasePropertiesItems) *CreateDatabaseRegistrationOptions {
	_options.DatabaseProperties = databaseProperties
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDatabaseRegistrationOptions) SetDescription(description string) *CreateDatabaseRegistrationOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateDatabaseRegistrationOptions) SetTags(tags []string) *CreateDatabaseRegistrationOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateDatabaseRegistrationOptions) SetAuthInstanceID(authInstanceID string) *CreateDatabaseRegistrationOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDatabaseRegistrationOptions) SetHeaders(param map[string]string) *CreateDatabaseRegistrationOptions {
	options.Headers = param
	return options
}

// CreateDb2EngineOptions : The CreateDb2Engine options.
type CreateDb2EngineOptions struct {
	// Origin - created or registered.
	Origin *string `json:"origin" validate:"required"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *Db2EngineDetailsBody `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateDb2EngineOptions.Origin property.
// Origin - created or registered.
const (
	CreateDb2EngineOptions_Origin_Discover = "discover"
	CreateDb2EngineOptions_Origin_External = "external"
	CreateDb2EngineOptions_Origin_Native = "native"
)

// NewCreateDb2EngineOptions : Instantiate CreateDb2EngineOptions
func (*WatsonxDataV2) NewCreateDb2EngineOptions(origin string) *CreateDb2EngineOptions {
	return &CreateDb2EngineOptions{
		Origin: core.StringPtr(origin),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreateDb2EngineOptions) SetOrigin(origin string) *CreateDb2EngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDb2EngineOptions) SetDescription(description string) *CreateDb2EngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreateDb2EngineOptions) SetEngineDetails(engineDetails *Db2EngineDetailsBody) *CreateDb2EngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreateDb2EngineOptions) SetEngineDisplayName(engineDisplayName string) *CreateDb2EngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateDb2EngineOptions) SetTags(tags []string) *CreateDb2EngineOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateDb2EngineOptions) SetAuthInstanceID(authInstanceID string) *CreateDb2EngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDb2EngineOptions) SetHeaders(param map[string]string) *CreateDb2EngineOptions {
	options.Headers = param
	return options
}

// CreateEnginePauseCreatedBody : Pause.
type CreateEnginePauseCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateEnginePauseCreatedBody unmarshals an instance of CreateEnginePauseCreatedBody from the specified map of raw messages.
func UnmarshalCreateEnginePauseCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEnginePauseCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateEngineRestartCreatedBody : restart engine.
type CreateEngineRestartCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateEngineRestartCreatedBody unmarshals an instance of CreateEngineRestartCreatedBody from the specified map of raw messages.
func UnmarshalCreateEngineRestartCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEngineRestartCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateEngineResumeCreatedBody : resume.
type CreateEngineResumeCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateEngineResumeCreatedBody unmarshals an instance of CreateEngineResumeCreatedBody from the specified map of raw messages.
func UnmarshalCreateEngineResumeCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEngineResumeCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateEngineScaleCreatedBody : scale engine.
type CreateEngineScaleCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateEngineScaleCreatedBody unmarshals an instance of CreateEngineScaleCreatedBody from the specified map of raw messages.
func UnmarshalCreateEngineScaleCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEngineScaleCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateMilvusServiceOptions : The CreateMilvusService options.
type CreateMilvusServiceOptions struct {
	// Origin - place holder.
	Origin *string `json:"origin" validate:"required"`

	// Service description.
	Description *string `json:"description,omitempty"`

	// Service display name.
	ServiceDisplayName *string `json:"service_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateMilvusServiceOptions : Instantiate CreateMilvusServiceOptions
func (*WatsonxDataV2) NewCreateMilvusServiceOptions(origin string) *CreateMilvusServiceOptions {
	return &CreateMilvusServiceOptions{
		Origin: core.StringPtr(origin),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreateMilvusServiceOptions) SetOrigin(origin string) *CreateMilvusServiceOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateMilvusServiceOptions) SetDescription(description string) *CreateMilvusServiceOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetServiceDisplayName : Allow user to set ServiceDisplayName
func (_options *CreateMilvusServiceOptions) SetServiceDisplayName(serviceDisplayName string) *CreateMilvusServiceOptions {
	_options.ServiceDisplayName = core.StringPtr(serviceDisplayName)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateMilvusServiceOptions) SetTags(tags []string) *CreateMilvusServiceOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateMilvusServiceOptions) SetAuthInstanceID(authInstanceID string) *CreateMilvusServiceOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateMilvusServiceOptions) SetHeaders(param map[string]string) *CreateMilvusServiceOptions {
	options.Headers = param
	return options
}

// CreateNetezzaEngineOptions : The CreateNetezzaEngine options.
type CreateNetezzaEngineOptions struct {
	// Origin - created or registered.
	Origin *string `json:"origin" validate:"required"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *NetezzaEngineDetailsBody `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateNetezzaEngineOptions.Origin property.
// Origin - created or registered.
const (
	CreateNetezzaEngineOptions_Origin_Discover = "discover"
	CreateNetezzaEngineOptions_Origin_External = "external"
	CreateNetezzaEngineOptions_Origin_Native = "native"
)

// NewCreateNetezzaEngineOptions : Instantiate CreateNetezzaEngineOptions
func (*WatsonxDataV2) NewCreateNetezzaEngineOptions(origin string) *CreateNetezzaEngineOptions {
	return &CreateNetezzaEngineOptions{
		Origin: core.StringPtr(origin),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreateNetezzaEngineOptions) SetOrigin(origin string) *CreateNetezzaEngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateNetezzaEngineOptions) SetDescription(description string) *CreateNetezzaEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreateNetezzaEngineOptions) SetEngineDetails(engineDetails *NetezzaEngineDetailsBody) *CreateNetezzaEngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreateNetezzaEngineOptions) SetEngineDisplayName(engineDisplayName string) *CreateNetezzaEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateNetezzaEngineOptions) SetTags(tags []string) *CreateNetezzaEngineOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateNetezzaEngineOptions) SetAuthInstanceID(authInstanceID string) *CreateNetezzaEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateNetezzaEngineOptions) SetHeaders(param map[string]string) *CreateNetezzaEngineOptions {
	options.Headers = param
	return options
}

// CreateOtherEngineOptions : The CreateOtherEngine options.
type CreateOtherEngineOptions struct {
	// External engine details.
	EngineDetails *OtherEngineDetailsBody `json:"engine_details" validate:"required"`

	// engine display name.
	EngineDisplayName *string `json:"engine_display_name" validate:"required"`

	// engine description.
	Description *string `json:"description,omitempty"`

	// Origin - created or registered.
	Origin *string `json:"origin,omitempty"`

	// other engine tags.
	Tags []string `json:"tags,omitempty"`

	// Engine type.
	Type *string `json:"type,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateOtherEngineOptions.Origin property.
// Origin - created or registered.
const (
	CreateOtherEngineOptions_Origin_Discover = "discover"
	CreateOtherEngineOptions_Origin_External = "external"
	CreateOtherEngineOptions_Origin_Native = "native"
)

// NewCreateOtherEngineOptions : Instantiate CreateOtherEngineOptions
func (*WatsonxDataV2) NewCreateOtherEngineOptions(engineDetails *OtherEngineDetailsBody, engineDisplayName string) *CreateOtherEngineOptions {
	return &CreateOtherEngineOptions{
		EngineDetails: engineDetails,
		EngineDisplayName: core.StringPtr(engineDisplayName),
	}
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreateOtherEngineOptions) SetEngineDetails(engineDetails *OtherEngineDetailsBody) *CreateOtherEngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreateOtherEngineOptions) SetEngineDisplayName(engineDisplayName string) *CreateOtherEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateOtherEngineOptions) SetDescription(description string) *CreateOtherEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetOrigin : Allow user to set Origin
func (_options *CreateOtherEngineOptions) SetOrigin(origin string) *CreateOtherEngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateOtherEngineOptions) SetTags(tags []string) *CreateOtherEngineOptions {
	_options.Tags = tags
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateOtherEngineOptions) SetType(typeVar string) *CreateOtherEngineOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateOtherEngineOptions) SetAuthInstanceID(authInstanceID string) *CreateOtherEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateOtherEngineOptions) SetHeaders(param map[string]string) *CreateOtherEngineOptions {
	options.Headers = param
	return options
}

// CreatePrestissimoEngineOptions : The CreatePrestissimoEngine options.
type CreatePrestissimoEngineOptions struct {
	// Origin - created or registered.
	Origin *string `json:"origin" validate:"required"`

	// Associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *PrestissimoEngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Region (cloud).
	Region *string `json:"region,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Version like 0.278 for prestissimo or else.
	Version *string `json:"version,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreatePrestissimoEngineOptions.Origin property.
// Origin - created or registered.
const (
	CreatePrestissimoEngineOptions_Origin_Discover = "discover"
	CreatePrestissimoEngineOptions_Origin_External = "external"
	CreatePrestissimoEngineOptions_Origin_Native = "native"
)

// NewCreatePrestissimoEngineOptions : Instantiate CreatePrestissimoEngineOptions
func (*WatsonxDataV2) NewCreatePrestissimoEngineOptions(origin string) *CreatePrestissimoEngineOptions {
	return &CreatePrestissimoEngineOptions{
		Origin: core.StringPtr(origin),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreatePrestissimoEngineOptions) SetOrigin(origin string) *CreatePrestissimoEngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetAssociatedCatalogs : Allow user to set AssociatedCatalogs
func (_options *CreatePrestissimoEngineOptions) SetAssociatedCatalogs(associatedCatalogs []string) *CreatePrestissimoEngineOptions {
	_options.AssociatedCatalogs = associatedCatalogs
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreatePrestissimoEngineOptions) SetDescription(description string) *CreatePrestissimoEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreatePrestissimoEngineOptions) SetEngineDetails(engineDetails *PrestissimoEngineDetails) *CreatePrestissimoEngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreatePrestissimoEngineOptions) SetEngineDisplayName(engineDisplayName string) *CreatePrestissimoEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetRegion : Allow user to set Region
func (_options *CreatePrestissimoEngineOptions) SetRegion(region string) *CreatePrestissimoEngineOptions {
	_options.Region = core.StringPtr(region)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreatePrestissimoEngineOptions) SetTags(tags []string) *CreatePrestissimoEngineOptions {
	_options.Tags = tags
	return _options
}

// SetVersion : Allow user to set Version
func (_options *CreatePrestissimoEngineOptions) SetVersion(version string) *CreatePrestissimoEngineOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreatePrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *CreatePrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePrestissimoEngineOptions) SetHeaders(param map[string]string) *CreatePrestissimoEngineOptions {
	options.Headers = param
	return options
}

// CreatePrestoEngineOptions : The CreatePrestoEngine options.
type CreatePrestoEngineOptions struct {
	// Origin - created or registered.
	Origin *string `json:"origin" validate:"required"`

	// Associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// Node details.
	EngineDetails *EngineDetailsBody `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Region (cloud).
	Region *string `json:"region,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Version like 0.278 for presto or else.
	Version *string `json:"version,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreatePrestoEngineOptions.Origin property.
// Origin - created or registered.
const (
	CreatePrestoEngineOptions_Origin_Discover = "discover"
	CreatePrestoEngineOptions_Origin_External = "external"
	CreatePrestoEngineOptions_Origin_Native = "native"
)

// NewCreatePrestoEngineOptions : Instantiate CreatePrestoEngineOptions
func (*WatsonxDataV2) NewCreatePrestoEngineOptions(origin string) *CreatePrestoEngineOptions {
	return &CreatePrestoEngineOptions{
		Origin: core.StringPtr(origin),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreatePrestoEngineOptions) SetOrigin(origin string) *CreatePrestoEngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetAssociatedCatalogs : Allow user to set AssociatedCatalogs
func (_options *CreatePrestoEngineOptions) SetAssociatedCatalogs(associatedCatalogs []string) *CreatePrestoEngineOptions {
	_options.AssociatedCatalogs = associatedCatalogs
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreatePrestoEngineOptions) SetDescription(description string) *CreatePrestoEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreatePrestoEngineOptions) SetEngineDetails(engineDetails *EngineDetailsBody) *CreatePrestoEngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreatePrestoEngineOptions) SetEngineDisplayName(engineDisplayName string) *CreatePrestoEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetRegion : Allow user to set Region
func (_options *CreatePrestoEngineOptions) SetRegion(region string) *CreatePrestoEngineOptions {
	_options.Region = core.StringPtr(region)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreatePrestoEngineOptions) SetTags(tags []string) *CreatePrestoEngineOptions {
	_options.Tags = tags
	return _options
}

// SetVersion : Allow user to set Version
func (_options *CreatePrestoEngineOptions) SetVersion(version string) *CreatePrestoEngineOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreatePrestoEngineOptions) SetAuthInstanceID(authInstanceID string) *CreatePrestoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePrestoEngineOptions) SetHeaders(param map[string]string) *CreatePrestoEngineOptions {
	options.Headers = param
	return options
}

// CreateSchemaCreatedBody : success response.
type CreateSchemaCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalCreateSchemaCreatedBody unmarshals an instance of CreateSchemaCreatedBody from the specified map of raw messages.
func UnmarshalCreateSchemaCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateSchemaCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateSchemaOptions : The CreateSchema options.
type CreateSchemaOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog name.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Path associated with bucket.
	CustomPath *string `json:"custom_path" validate:"required"`

	// Schema name.
	SchemaName *string `json:"schema_name" validate:"required"`

	// Bucket associated to metastore where schema will be added.
	BucketName *string `json:"bucket_name,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateSchemaOptions : Instantiate CreateSchemaOptions
func (*WatsonxDataV2) NewCreateSchemaOptions(engineID string, catalogID string, customPath string, schemaName string) *CreateSchemaOptions {
	return &CreateSchemaOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		CustomPath: core.StringPtr(customPath),
		SchemaName: core.StringPtr(schemaName),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateSchemaOptions) SetEngineID(engineID string) *CreateSchemaOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *CreateSchemaOptions) SetCatalogID(catalogID string) *CreateSchemaOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetCustomPath : Allow user to set CustomPath
func (_options *CreateSchemaOptions) SetCustomPath(customPath string) *CreateSchemaOptions {
	_options.CustomPath = core.StringPtr(customPath)
	return _options
}

// SetSchemaName : Allow user to set SchemaName
func (_options *CreateSchemaOptions) SetSchemaName(schemaName string) *CreateSchemaOptions {
	_options.SchemaName = core.StringPtr(schemaName)
	return _options
}

// SetBucketName : Allow user to set BucketName
func (_options *CreateSchemaOptions) SetBucketName(bucketName string) *CreateSchemaOptions {
	_options.BucketName = core.StringPtr(bucketName)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateSchemaOptions) SetAuthInstanceID(authInstanceID string) *CreateSchemaOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSchemaOptions) SetHeaders(param map[string]string) *CreateSchemaOptions {
	options.Headers = param
	return options
}

// CreateSparkEngineApplicationOptions : The CreateSparkEngineApplication options.
type CreateSparkEngineApplicationOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Application details.
	ApplicationDetails *SparkApplicationDetails `json:"application_details" validate:"required"`

	// Job endpoint.
	JobEndpoint *string `json:"job_endpoint,omitempty"`

	// Service Instance ID for POST.
	ServiceInstanceID *string `json:"service_instance_id,omitempty"`

	// Engine Type.
	Type *string `json:"type,omitempty"`

	// Spark application volumes to mount.
	Volumes []SparkVolumeDetails `json:"volumes,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// state.
	State []string `json:"state,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateSparkEngineApplicationOptions.Type property.
// Engine Type.
const (
	CreateSparkEngineApplicationOptions_Type_Emr = "emr"
	CreateSparkEngineApplicationOptions_Type_Iae = "iae"
)

// NewCreateSparkEngineApplicationOptions : Instantiate CreateSparkEngineApplicationOptions
func (*WatsonxDataV2) NewCreateSparkEngineApplicationOptions(engineID string, applicationDetails *SparkApplicationDetails) *CreateSparkEngineApplicationOptions {
	return &CreateSparkEngineApplicationOptions{
		EngineID: core.StringPtr(engineID),
		ApplicationDetails: applicationDetails,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateSparkEngineApplicationOptions) SetEngineID(engineID string) *CreateSparkEngineApplicationOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetApplicationDetails : Allow user to set ApplicationDetails
func (_options *CreateSparkEngineApplicationOptions) SetApplicationDetails(applicationDetails *SparkApplicationDetails) *CreateSparkEngineApplicationOptions {
	_options.ApplicationDetails = applicationDetails
	return _options
}

// SetJobEndpoint : Allow user to set JobEndpoint
func (_options *CreateSparkEngineApplicationOptions) SetJobEndpoint(jobEndpoint string) *CreateSparkEngineApplicationOptions {
	_options.JobEndpoint = core.StringPtr(jobEndpoint)
	return _options
}

// SetServiceInstanceID : Allow user to set ServiceInstanceID
func (_options *CreateSparkEngineApplicationOptions) SetServiceInstanceID(serviceInstanceID string) *CreateSparkEngineApplicationOptions {
	_options.ServiceInstanceID = core.StringPtr(serviceInstanceID)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateSparkEngineApplicationOptions) SetType(typeVar string) *CreateSparkEngineApplicationOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetVolumes : Allow user to set Volumes
func (_options *CreateSparkEngineApplicationOptions) SetVolumes(volumes []SparkVolumeDetails) *CreateSparkEngineApplicationOptions {
	_options.Volumes = volumes
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateSparkEngineApplicationOptions) SetAuthInstanceID(authInstanceID string) *CreateSparkEngineApplicationOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetState : Allow user to set State
func (_options *CreateSparkEngineApplicationOptions) SetState(state []string) *CreateSparkEngineApplicationOptions {
	_options.State = state
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSparkEngineApplicationOptions) SetHeaders(param map[string]string) *CreateSparkEngineApplicationOptions {
	options.Headers = param
	return options
}

// CreateSparkEngineOptions : The CreateSparkEngine options.
type CreateSparkEngineOptions struct {
	// Origin - created or registered.
	Origin *string `json:"origin" validate:"required"`

	// Associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// Node details.
	EngineDetails *SparkEngineDetailsPrototype `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateSparkEngineOptions.Origin property.
// Origin - created or registered.
const (
	CreateSparkEngineOptions_Origin_Discover = "discover"
	CreateSparkEngineOptions_Origin_External = "external"
	CreateSparkEngineOptions_Origin_Native = "native"
)

// NewCreateSparkEngineOptions : Instantiate CreateSparkEngineOptions
func (*WatsonxDataV2) NewCreateSparkEngineOptions(origin string) *CreateSparkEngineOptions {
	return &CreateSparkEngineOptions{
		Origin: core.StringPtr(origin),
	}
}

// SetOrigin : Allow user to set Origin
func (_options *CreateSparkEngineOptions) SetOrigin(origin string) *CreateSparkEngineOptions {
	_options.Origin = core.StringPtr(origin)
	return _options
}

// SetAssociatedCatalogs : Allow user to set AssociatedCatalogs
func (_options *CreateSparkEngineOptions) SetAssociatedCatalogs(associatedCatalogs []string) *CreateSparkEngineOptions {
	_options.AssociatedCatalogs = associatedCatalogs
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateSparkEngineOptions) SetDescription(description string) *CreateSparkEngineOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEngineDetails : Allow user to set EngineDetails
func (_options *CreateSparkEngineOptions) SetEngineDetails(engineDetails *SparkEngineDetailsPrototype) *CreateSparkEngineOptions {
	_options.EngineDetails = engineDetails
	return _options
}

// SetEngineDisplayName : Allow user to set EngineDisplayName
func (_options *CreateSparkEngineOptions) SetEngineDisplayName(engineDisplayName string) *CreateSparkEngineOptions {
	_options.EngineDisplayName = core.StringPtr(engineDisplayName)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *CreateSparkEngineOptions) SetStatus(status string) *CreateSparkEngineOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateSparkEngineOptions) SetTags(tags []string) *CreateSparkEngineOptions {
	_options.Tags = tags
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateSparkEngineOptions) SetAuthInstanceID(authInstanceID string) *CreateSparkEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSparkEngineOptions) SetHeaders(param map[string]string) *CreateSparkEngineOptions {
	options.Headers = param
	return options
}

// CreateSparkEnginePauseOptions : The CreateSparkEnginePause options.
type CreateSparkEnginePauseOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateSparkEnginePauseOptions : Instantiate CreateSparkEnginePauseOptions
func (*WatsonxDataV2) NewCreateSparkEnginePauseOptions(engineID string) *CreateSparkEnginePauseOptions {
	return &CreateSparkEnginePauseOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateSparkEnginePauseOptions) SetEngineID(engineID string) *CreateSparkEnginePauseOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateSparkEnginePauseOptions) SetAuthInstanceID(authInstanceID string) *CreateSparkEnginePauseOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSparkEnginePauseOptions) SetHeaders(param map[string]string) *CreateSparkEnginePauseOptions {
	options.Headers = param
	return options
}

// CreateSparkEngineResumeOptions : The CreateSparkEngineResume options.
type CreateSparkEngineResumeOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateSparkEngineResumeOptions : Instantiate CreateSparkEngineResumeOptions
func (*WatsonxDataV2) NewCreateSparkEngineResumeOptions(engineID string) *CreateSparkEngineResumeOptions {
	return &CreateSparkEngineResumeOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateSparkEngineResumeOptions) SetEngineID(engineID string) *CreateSparkEngineResumeOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateSparkEngineResumeOptions) SetAuthInstanceID(authInstanceID string) *CreateSparkEngineResumeOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSparkEngineResumeOptions) SetHeaders(param map[string]string) *CreateSparkEngineResumeOptions {
	options.Headers = param
	return options
}

// CreateSparkEngineScaleOptions : The CreateSparkEngineScale options.
type CreateSparkEngineScaleOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Node count.
	NumberOfNodes *int64 `json:"number_of_nodes,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateSparkEngineScaleOptions : Instantiate CreateSparkEngineScaleOptions
func (*WatsonxDataV2) NewCreateSparkEngineScaleOptions(engineID string) *CreateSparkEngineScaleOptions {
	return &CreateSparkEngineScaleOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *CreateSparkEngineScaleOptions) SetEngineID(engineID string) *CreateSparkEngineScaleOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetNumberOfNodes : Allow user to set NumberOfNodes
func (_options *CreateSparkEngineScaleOptions) SetNumberOfNodes(numberOfNodes int64) *CreateSparkEngineScaleOptions {
	_options.NumberOfNodes = core.Int64Ptr(numberOfNodes)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *CreateSparkEngineScaleOptions) SetAuthInstanceID(authInstanceID string) *CreateSparkEngineScaleOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSparkEngineScaleOptions) SetHeaders(param map[string]string) *CreateSparkEngineScaleOptions {
	options.Headers = param
	return options
}

// DatabaseCatalog : database catalog.
type DatabaseCatalog struct {
	// catalog name.
	CatalogName *string `json:"catalog_name,omitempty"`

	// catalog tags.
	CatalogTags []string `json:"catalog_tags,omitempty"`

	// catalog type.
	CatalogType *string `json:"catalog_type,omitempty"`
}

// UnmarshalDatabaseCatalog unmarshals an instance of DatabaseCatalog from the specified map of raw messages.
func UnmarshalDatabaseCatalog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseCatalog)
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_tags", &obj.CatalogTags)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_type", &obj.CatalogType)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatabaseDetails : database details.
type DatabaseDetails struct {
	// contents of a pem/crt file.
	Certificate *string `json:"certificate,omitempty"`

	// extension of the certificate file.
	CertificateExtension *string `json:"certificate_extension,omitempty"`

	// Database name.
	DatabaseName *string `json:"database_name,omitempty"`

	// Host name.
	Hostname *string `json:"hostname" validate:"required"`

	// Hostname in certificate.
	HostnameInCertificate *string `json:"hostname_in_certificate,omitempty"`

	// String of hostname:port.
	Hosts *string `json:"hosts,omitempty"`

	// Psssword.
	Password *string `json:"password,omitempty"`

	// Port.
	Port *int64 `json:"port" validate:"required"`

	// SASL Mode.
	Sasl *bool `json:"sasl,omitempty"`

	// SSL Mode.
	Ssl *bool `json:"ssl,omitempty"`

	// Only for Kafka - Add kafka tables.
	Tables *string `json:"tables,omitempty"`

	// Username.
	Username *string `json:"username,omitempty"`

	// Verify certificate.
	ValidateServerCertificate *bool `json:"validate_server_certificate,omitempty"`
}

// NewDatabaseDetails : Instantiate DatabaseDetails (Generic Model Constructor)
func (*WatsonxDataV2) NewDatabaseDetails(hostname string, port int64) (_model *DatabaseDetails, err error) {
	_model = &DatabaseDetails{
		Hostname: core.StringPtr(hostname),
		Port: core.Int64Ptr(port),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalDatabaseDetails unmarshals an instance of DatabaseDetails from the specified map of raw messages.
func UnmarshalDatabaseDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseDetails)
	err = core.UnmarshalPrimitive(m, "certificate", &obj.Certificate)
	if err != nil {
		err = core.SDKErrorf(err, "", "certificate-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_extension", &obj.CertificateExtension)
	if err != nil {
		err = core.SDKErrorf(err, "", "certificate_extension-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "database_name", &obj.DatabaseName)
	if err != nil {
		err = core.SDKErrorf(err, "", "database_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		err = core.SDKErrorf(err, "", "hostname-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname_in_certificate", &obj.HostnameInCertificate)
	if err != nil {
		err = core.SDKErrorf(err, "", "hostname_in_certificate-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "hosts", &obj.Hosts)
	if err != nil {
		err = core.SDKErrorf(err, "", "hosts-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		err = core.SDKErrorf(err, "", "password-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "sasl", &obj.Sasl)
	if err != nil {
		err = core.SDKErrorf(err, "", "sasl-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		err = core.SDKErrorf(err, "", "ssl-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tables", &obj.Tables)
	if err != nil {
		err = core.SDKErrorf(err, "", "tables-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		err = core.SDKErrorf(err, "", "username-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "validate_server_certificate", &obj.ValidateServerCertificate)
	if err != nil {
		err = core.SDKErrorf(err, "", "validate_server_certificate-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatabaseRegistration : database registration object.
type DatabaseRegistration struct {
	// actions.
	Actions []string `json:"actions,omitempty"`

	// database catalog.
	AssociatedCatalog *DatabaseCatalog `json:"associated_catalog,omitempty"`

	// Catalog name.
	CatalogName *string `json:"catalog_name,omitempty"`

	// Created by.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// database details.
	DatabaseDetails *DatabaseDetails `json:"database_details" validate:"required"`

	// Database display name.
	DatabaseDisplayName *string `json:"database_display_name" validate:"required"`

	// Database ID.
	DatabaseID *string `json:"database_id,omitempty"`

	// This will hold all the properties for a custom database.
	DatabaseProperties []DatabaseRegistrationDatabasePropertiesItems `json:"database_properties,omitempty"`

	// Connector type.
	DatabaseType *string `json:"database_type" validate:"required"`

	// Database description.
	Description *string `json:"description,omitempty"`

	// tags.
	Tags []string `json:"tags,omitempty"`
}

// UnmarshalDatabaseRegistration unmarshals an instance of DatabaseRegistration from the specified map of raw messages.
func UnmarshalDatabaseRegistration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseRegistration)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		err = core.SDKErrorf(err, "", "actions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "associated_catalog", &obj.AssociatedCatalog, UnmarshalDatabaseCatalog)
	if err != nil {
		err = core.SDKErrorf(err, "", "associated_catalog-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_on-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "database_details", &obj.DatabaseDetails, UnmarshalDatabaseDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "database_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "database_display_name", &obj.DatabaseDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "database_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "database_id", &obj.DatabaseID)
	if err != nil {
		err = core.SDKErrorf(err, "", "database_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "database_properties", &obj.DatabaseProperties, UnmarshalDatabaseRegistrationDatabasePropertiesItems)
	if err != nil {
		err = core.SDKErrorf(err, "", "database_properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "database_type", &obj.DatabaseType)
	if err != nil {
		err = core.SDKErrorf(err, "", "database_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatabaseRegistrationCollection : list database registrations.
type DatabaseRegistrationCollection struct {
	// Database body.
	DatabaseRegistrations []DatabaseRegistration `json:"database_registrations,omitempty"`
}

// UnmarshalDatabaseRegistrationCollection unmarshals an instance of DatabaseRegistrationCollection from the specified map of raw messages.
func UnmarshalDatabaseRegistrationCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseRegistrationCollection)
	err = core.UnmarshalModel(m, "database_registrations", &obj.DatabaseRegistrations, UnmarshalDatabaseRegistration)
	if err != nil {
		err = core.SDKErrorf(err, "", "database_registrations-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatabaseRegistrationDatabasePropertiesItems : Key value object.
type DatabaseRegistrationDatabasePropertiesItems struct {
	// Wether the value is to be encrypted before storing.
	Encrypt *bool `json:"encrypt" validate:"required"`

	// Key of the database property.
	Key *string `json:"key" validate:"required"`

	// Value of the database property.
	Value *string `json:"value" validate:"required"`
}

// UnmarshalDatabaseRegistrationDatabasePropertiesItems unmarshals an instance of DatabaseRegistrationDatabasePropertiesItems from the specified map of raw messages.
func UnmarshalDatabaseRegistrationDatabasePropertiesItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseRegistrationDatabasePropertiesItems)
	err = core.UnmarshalPrimitive(m, "encrypt", &obj.Encrypt)
	if err != nil {
		err = core.SDKErrorf(err, "", "encrypt-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		err = core.SDKErrorf(err, "", "key-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatabaseRegistrationPatch : update db body.
type DatabaseRegistrationPatch struct {
	// New database details.
	DatabaseDetails *DatabaseRegistrationPatchDatabaseDetails `json:"database_details,omitempty"`

	// New database display name.
	DatabaseDisplayName *string `json:"database_display_name,omitempty"`

	// New database description.
	Description *string `json:"description,omitempty"`

	// New tags.
	Tags []string `json:"tags,omitempty"`
}

// UnmarshalDatabaseRegistrationPatch unmarshals an instance of DatabaseRegistrationPatch from the specified map of raw messages.
func UnmarshalDatabaseRegistrationPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseRegistrationPatch)
	err = core.UnmarshalModel(m, "database_details", &obj.DatabaseDetails, UnmarshalDatabaseRegistrationPatchDatabaseDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "database_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "database_display_name", &obj.DatabaseDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "database_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the DatabaseRegistrationPatch
func (databaseRegistrationPatch *DatabaseRegistrationPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(databaseRegistrationPatch.DatabaseDetails) {
		_patch["database_details"] = databaseRegistrationPatch.DatabaseDetails.asPatch()
	}
	if !core.IsNil(databaseRegistrationPatch.DatabaseDisplayName) {
		_patch["database_display_name"] = databaseRegistrationPatch.DatabaseDisplayName
	}
	if !core.IsNil(databaseRegistrationPatch.Description) {
		_patch["description"] = databaseRegistrationPatch.Description
	}
	if !core.IsNil(databaseRegistrationPatch.Tags) {
		_patch["tags"] = databaseRegistrationPatch.Tags
	}

	return
}

// DatabaseRegistrationPatchDatabaseDetails : New database details.
type DatabaseRegistrationPatchDatabaseDetails struct {
	// New password.
	Password *string `json:"password,omitempty"`

	// New username.
	Username *string `json:"username,omitempty"`
}

// UnmarshalDatabaseRegistrationPatchDatabaseDetails unmarshals an instance of DatabaseRegistrationPatchDatabaseDetails from the specified map of raw messages.
func UnmarshalDatabaseRegistrationPatchDatabaseDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseRegistrationPatchDatabaseDetails)
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		err = core.SDKErrorf(err, "", "password-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		err = core.SDKErrorf(err, "", "username-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the DatabaseRegistrationPatchDatabaseDetails
func (databaseRegistrationPatchDatabaseDetails *DatabaseRegistrationPatchDatabaseDetails) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(databaseRegistrationPatchDatabaseDetails.Password) {
		_patch["password"] = databaseRegistrationPatchDatabaseDetails.Password
	}
	if !core.IsNil(databaseRegistrationPatchDatabaseDetails.Username) {
		_patch["username"] = databaseRegistrationPatchDatabaseDetails.Username
	}

	return
}

// DatabaseRegistrationPrototypeDatabasePropertiesItems : Key value object.
type DatabaseRegistrationPrototypeDatabasePropertiesItems struct {
	// Wether the value is to be encrypted before storing.
	Encrypt *bool `json:"encrypt" validate:"required"`

	// Key of the database property.
	Key *string `json:"key" validate:"required"`

	// Value of the database property.
	Value *string `json:"value" validate:"required"`
}

// NewDatabaseRegistrationPrototypeDatabasePropertiesItems : Instantiate DatabaseRegistrationPrototypeDatabasePropertiesItems (Generic Model Constructor)
func (*WatsonxDataV2) NewDatabaseRegistrationPrototypeDatabasePropertiesItems(encrypt bool, key string, value string) (_model *DatabaseRegistrationPrototypeDatabasePropertiesItems, err error) {
	_model = &DatabaseRegistrationPrototypeDatabasePropertiesItems{
		Encrypt: core.BoolPtr(encrypt),
		Key: core.StringPtr(key),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalDatabaseRegistrationPrototypeDatabasePropertiesItems unmarshals an instance of DatabaseRegistrationPrototypeDatabasePropertiesItems from the specified map of raw messages.
func UnmarshalDatabaseRegistrationPrototypeDatabasePropertiesItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseRegistrationPrototypeDatabasePropertiesItems)
	err = core.UnmarshalPrimitive(m, "encrypt", &obj.Encrypt)
	if err != nil {
		err = core.SDKErrorf(err, "", "encrypt-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		err = core.SDKErrorf(err, "", "key-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Db2Engine : Db2 engine details.
type Db2Engine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// watsonx.data build version.
	BuildVersion *string `json:"build_version,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *Db2EngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Engine host name. In case of OCP based clusters, this is internal hostname.
	HostName *string `json:"host_name,omitempty"`

	// Origin - place holder.
	Origin *string `json:"origin,omitempty"`

	// Engine port.
	Port *int64 `json:"port,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Engine type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalDb2Engine unmarshals an instance of Db2Engine from the specified map of raw messages.
func UnmarshalDb2Engine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Db2Engine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		err = core.SDKErrorf(err, "", "actions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "build_version", &obj.BuildVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "build_version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_on-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalDb2EngineDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		err = core.SDKErrorf(err, "", "host_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		err = core.SDKErrorf(err, "", "origin-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Db2EngineCollection : list db2 engines.
type Db2EngineCollection struct {
	// list db2 engines.
	Db2Engines []Db2Engine `json:"db2_engines,omitempty"`
}

// UnmarshalDb2EngineCollection unmarshals an instance of Db2EngineCollection from the specified map of raw messages.
func UnmarshalDb2EngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Db2EngineCollection)
	err = core.UnmarshalModel(m, "db2_engines", &obj.Db2Engines, UnmarshalDb2Engine)
	if err != nil {
		err = core.SDKErrorf(err, "", "db2_engines-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Db2EngineDetails : External engine details.
type Db2EngineDetails struct {
	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// Metastore host.
	MetastoreHost *string `json:"metastore_host,omitempty"`
}

// UnmarshalDb2EngineDetails unmarshals an instance of Db2EngineDetails from the specified map of raw messages.
func UnmarshalDb2EngineDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Db2EngineDetails)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		err = core.SDKErrorf(err, "", "connection_string-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "metastore_host", &obj.MetastoreHost)
	if err != nil {
		err = core.SDKErrorf(err, "", "metastore_host-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Db2EngineDetailsBody : External engine details.
type Db2EngineDetailsBody struct {
	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`
}

// UnmarshalDb2EngineDetailsBody unmarshals an instance of Db2EngineDetailsBody from the specified map of raw messages.
func UnmarshalDb2EngineDetailsBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Db2EngineDetailsBody)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		err = core.SDKErrorf(err, "", "connection_string-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Db2EnginePatch : UpdateEngine body.
type Db2EnginePatch struct {
	// Modified description.
	Description *string `json:"description,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`
}

// UnmarshalDb2EnginePatch unmarshals an instance of Db2EnginePatch from the specified map of raw messages.
func UnmarshalDb2EnginePatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Db2EnginePatch)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the Db2EnginePatch
func (db2EnginePatch *Db2EnginePatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(db2EnginePatch.Description) {
		_patch["description"] = db2EnginePatch.Description
	}
	if !core.IsNil(db2EnginePatch.EngineDisplayName) {
		_patch["engine_display_name"] = db2EnginePatch.EngineDisplayName
	}
	if !core.IsNil(db2EnginePatch.Tags) {
		_patch["tags"] = db2EnginePatch.Tags
	}

	return
}

// DeleteColumnOptions : The DeleteColumn options.
type DeleteColumnOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded schema name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// URL encoded schema name.
	ColumnID *string `json:"column_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteColumnOptions : Instantiate DeleteColumnOptions
func (*WatsonxDataV2) NewDeleteColumnOptions(engineID string, catalogID string, schemaID string, tableID string, columnID string) *DeleteColumnOptions {
	return &DeleteColumnOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
		ColumnID: core.StringPtr(columnID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteColumnOptions) SetEngineID(engineID string) *DeleteColumnOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *DeleteColumnOptions) SetCatalogID(catalogID string) *DeleteColumnOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *DeleteColumnOptions) SetSchemaID(schemaID string) *DeleteColumnOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *DeleteColumnOptions) SetTableID(tableID string) *DeleteColumnOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetColumnID : Allow user to set ColumnID
func (_options *DeleteColumnOptions) SetColumnID(columnID string) *DeleteColumnOptions {
	_options.ColumnID = core.StringPtr(columnID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteColumnOptions) SetAuthInstanceID(authInstanceID string) *DeleteColumnOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteColumnOptions) SetHeaders(param map[string]string) *DeleteColumnOptions {
	options.Headers = param
	return options
}

// DeleteDatabaseCatalogOptions : The DeleteDatabaseCatalog options.
type DeleteDatabaseCatalogOptions struct {
	// database id.
	DatabaseID *string `json:"database_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDatabaseCatalogOptions : Instantiate DeleteDatabaseCatalogOptions
func (*WatsonxDataV2) NewDeleteDatabaseCatalogOptions(databaseID string) *DeleteDatabaseCatalogOptions {
	return &DeleteDatabaseCatalogOptions{
		DatabaseID: core.StringPtr(databaseID),
	}
}

// SetDatabaseID : Allow user to set DatabaseID
func (_options *DeleteDatabaseCatalogOptions) SetDatabaseID(databaseID string) *DeleteDatabaseCatalogOptions {
	_options.DatabaseID = core.StringPtr(databaseID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteDatabaseCatalogOptions) SetAuthInstanceID(authInstanceID string) *DeleteDatabaseCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDatabaseCatalogOptions) SetHeaders(param map[string]string) *DeleteDatabaseCatalogOptions {
	options.Headers = param
	return options
}

// DeleteDb2EngineOptions : The DeleteDb2Engine options.
type DeleteDb2EngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDb2EngineOptions : Instantiate DeleteDb2EngineOptions
func (*WatsonxDataV2) NewDeleteDb2EngineOptions(engineID string) *DeleteDb2EngineOptions {
	return &DeleteDb2EngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteDb2EngineOptions) SetEngineID(engineID string) *DeleteDb2EngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteDb2EngineOptions) SetAuthInstanceID(authInstanceID string) *DeleteDb2EngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDb2EngineOptions) SetHeaders(param map[string]string) *DeleteDb2EngineOptions {
	options.Headers = param
	return options
}

// DeleteDeactivateBucketOptions : The DeleteDeactivateBucket options.
type DeleteDeactivateBucketOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDeactivateBucketOptions : Instantiate DeleteDeactivateBucketOptions
func (*WatsonxDataV2) NewDeleteDeactivateBucketOptions(bucketID string) *DeleteDeactivateBucketOptions {
	return &DeleteDeactivateBucketOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *DeleteDeactivateBucketOptions) SetBucketID(bucketID string) *DeleteDeactivateBucketOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteDeactivateBucketOptions) SetAuthInstanceID(authInstanceID string) *DeleteDeactivateBucketOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDeactivateBucketOptions) SetHeaders(param map[string]string) *DeleteDeactivateBucketOptions {
	options.Headers = param
	return options
}

// DeleteEngineOptions : The DeleteEngine options.
type DeleteEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteEngineOptions : Instantiate DeleteEngineOptions
func (*WatsonxDataV2) NewDeleteEngineOptions(engineID string) *DeleteEngineOptions {
	return &DeleteEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteEngineOptions) SetEngineID(engineID string) *DeleteEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteEngineOptions) SetAuthInstanceID(authInstanceID string) *DeleteEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteEngineOptions) SetHeaders(param map[string]string) *DeleteEngineOptions {
	options.Headers = param
	return options
}

// DeleteMilvusServiceOptions : The DeleteMilvusService options.
type DeleteMilvusServiceOptions struct {
	// service id.
	ServiceID *string `json:"service_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteMilvusServiceOptions : Instantiate DeleteMilvusServiceOptions
func (*WatsonxDataV2) NewDeleteMilvusServiceOptions(serviceID string) *DeleteMilvusServiceOptions {
	return &DeleteMilvusServiceOptions{
		ServiceID: core.StringPtr(serviceID),
	}
}

// SetServiceID : Allow user to set ServiceID
func (_options *DeleteMilvusServiceOptions) SetServiceID(serviceID string) *DeleteMilvusServiceOptions {
	_options.ServiceID = core.StringPtr(serviceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteMilvusServiceOptions) SetAuthInstanceID(authInstanceID string) *DeleteMilvusServiceOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteMilvusServiceOptions) SetHeaders(param map[string]string) *DeleteMilvusServiceOptions {
	options.Headers = param
	return options
}

// DeleteNetezzaEngineOptions : The DeleteNetezzaEngine options.
type DeleteNetezzaEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteNetezzaEngineOptions : Instantiate DeleteNetezzaEngineOptions
func (*WatsonxDataV2) NewDeleteNetezzaEngineOptions(engineID string) *DeleteNetezzaEngineOptions {
	return &DeleteNetezzaEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteNetezzaEngineOptions) SetEngineID(engineID string) *DeleteNetezzaEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteNetezzaEngineOptions) SetAuthInstanceID(authInstanceID string) *DeleteNetezzaEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteNetezzaEngineOptions) SetHeaders(param map[string]string) *DeleteNetezzaEngineOptions {
	options.Headers = param
	return options
}

// DeleteOtherEngineOptions : The DeleteOtherEngine options.
type DeleteOtherEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteOtherEngineOptions : Instantiate DeleteOtherEngineOptions
func (*WatsonxDataV2) NewDeleteOtherEngineOptions(engineID string) *DeleteOtherEngineOptions {
	return &DeleteOtherEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteOtherEngineOptions) SetEngineID(engineID string) *DeleteOtherEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteOtherEngineOptions) SetAuthInstanceID(authInstanceID string) *DeleteOtherEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteOtherEngineOptions) SetHeaders(param map[string]string) *DeleteOtherEngineOptions {
	options.Headers = param
	return options
}

// DeletePrestissimoEngineCatalogsOptions : The DeletePrestissimoEngineCatalogs options.
type DeletePrestissimoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Catalog id(s) to be stopped, comma separated.
	CatalogNames *string `json:"catalog_names" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeletePrestissimoEngineCatalogsOptions : Instantiate DeletePrestissimoEngineCatalogsOptions
func (*WatsonxDataV2) NewDeletePrestissimoEngineCatalogsOptions(engineID string, catalogNames string) *DeletePrestissimoEngineCatalogsOptions {
	return &DeletePrestissimoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogNames: core.StringPtr(catalogNames),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeletePrestissimoEngineCatalogsOptions) SetEngineID(engineID string) *DeletePrestissimoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogNames : Allow user to set CatalogNames
func (_options *DeletePrestissimoEngineCatalogsOptions) SetCatalogNames(catalogNames string) *DeletePrestissimoEngineCatalogsOptions {
	_options.CatalogNames = core.StringPtr(catalogNames)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeletePrestissimoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *DeletePrestissimoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePrestissimoEngineCatalogsOptions) SetHeaders(param map[string]string) *DeletePrestissimoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// DeletePrestissimoEngineOptions : The DeletePrestissimoEngine options.
type DeletePrestissimoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeletePrestissimoEngineOptions : Instantiate DeletePrestissimoEngineOptions
func (*WatsonxDataV2) NewDeletePrestissimoEngineOptions(engineID string) *DeletePrestissimoEngineOptions {
	return &DeletePrestissimoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeletePrestissimoEngineOptions) SetEngineID(engineID string) *DeletePrestissimoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeletePrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *DeletePrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePrestissimoEngineOptions) SetHeaders(param map[string]string) *DeletePrestissimoEngineOptions {
	options.Headers = param
	return options
}

// DeletePrestoEngineCatalogsOptions : The DeletePrestoEngineCatalogs options.
type DeletePrestoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Catalog id(s) to be stopped, comma separated.
	CatalogNames *string `json:"catalog_names" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeletePrestoEngineCatalogsOptions : Instantiate DeletePrestoEngineCatalogsOptions
func (*WatsonxDataV2) NewDeletePrestoEngineCatalogsOptions(engineID string, catalogNames string) *DeletePrestoEngineCatalogsOptions {
	return &DeletePrestoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogNames: core.StringPtr(catalogNames),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeletePrestoEngineCatalogsOptions) SetEngineID(engineID string) *DeletePrestoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogNames : Allow user to set CatalogNames
func (_options *DeletePrestoEngineCatalogsOptions) SetCatalogNames(catalogNames string) *DeletePrestoEngineCatalogsOptions {
	_options.CatalogNames = core.StringPtr(catalogNames)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeletePrestoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *DeletePrestoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePrestoEngineCatalogsOptions) SetHeaders(param map[string]string) *DeletePrestoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// DeleteSchemaOptions : The DeleteSchema options.
type DeleteSchemaOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog name.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded Schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteSchemaOptions : Instantiate DeleteSchemaOptions
func (*WatsonxDataV2) NewDeleteSchemaOptions(engineID string, catalogID string, schemaID string) *DeleteSchemaOptions {
	return &DeleteSchemaOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteSchemaOptions) SetEngineID(engineID string) *DeleteSchemaOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *DeleteSchemaOptions) SetCatalogID(catalogID string) *DeleteSchemaOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *DeleteSchemaOptions) SetSchemaID(schemaID string) *DeleteSchemaOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteSchemaOptions) SetAuthInstanceID(authInstanceID string) *DeleteSchemaOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSchemaOptions) SetHeaders(param map[string]string) *DeleteSchemaOptions {
	options.Headers = param
	return options
}

// DeleteSparkEngineApplicationsOptions : The DeleteSparkEngineApplications options.
type DeleteSparkEngineApplicationsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Application id(s) to be stopped, comma separated.
	ApplicationID *string `json:"application_id" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// state.
	State []string `json:"state,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteSparkEngineApplicationsOptions : Instantiate DeleteSparkEngineApplicationsOptions
func (*WatsonxDataV2) NewDeleteSparkEngineApplicationsOptions(engineID string, applicationID string) *DeleteSparkEngineApplicationsOptions {
	return &DeleteSparkEngineApplicationsOptions{
		EngineID: core.StringPtr(engineID),
		ApplicationID: core.StringPtr(applicationID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteSparkEngineApplicationsOptions) SetEngineID(engineID string) *DeleteSparkEngineApplicationsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetApplicationID : Allow user to set ApplicationID
func (_options *DeleteSparkEngineApplicationsOptions) SetApplicationID(applicationID string) *DeleteSparkEngineApplicationsOptions {
	_options.ApplicationID = core.StringPtr(applicationID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteSparkEngineApplicationsOptions) SetAuthInstanceID(authInstanceID string) *DeleteSparkEngineApplicationsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetState : Allow user to set State
func (_options *DeleteSparkEngineApplicationsOptions) SetState(state []string) *DeleteSparkEngineApplicationsOptions {
	_options.State = state
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSparkEngineApplicationsOptions) SetHeaders(param map[string]string) *DeleteSparkEngineApplicationsOptions {
	options.Headers = param
	return options
}

// DeleteSparkEngineCatalogsOptions : The DeleteSparkEngineCatalogs options.
type DeleteSparkEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Catalog id(s) to be stopped, comma separated.
	CatalogNames *string `json:"catalog_names" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteSparkEngineCatalogsOptions : Instantiate DeleteSparkEngineCatalogsOptions
func (*WatsonxDataV2) NewDeleteSparkEngineCatalogsOptions(engineID string, catalogNames string) *DeleteSparkEngineCatalogsOptions {
	return &DeleteSparkEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogNames: core.StringPtr(catalogNames),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteSparkEngineCatalogsOptions) SetEngineID(engineID string) *DeleteSparkEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogNames : Allow user to set CatalogNames
func (_options *DeleteSparkEngineCatalogsOptions) SetCatalogNames(catalogNames string) *DeleteSparkEngineCatalogsOptions {
	_options.CatalogNames = core.StringPtr(catalogNames)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteSparkEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *DeleteSparkEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSparkEngineCatalogsOptions) SetHeaders(param map[string]string) *DeleteSparkEngineCatalogsOptions {
	options.Headers = param
	return options
}

// DeleteSparkEngineHistoryServerOptions : The DeleteSparkEngineHistoryServer options.
type DeleteSparkEngineHistoryServerOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteSparkEngineHistoryServerOptions : Instantiate DeleteSparkEngineHistoryServerOptions
func (*WatsonxDataV2) NewDeleteSparkEngineHistoryServerOptions(engineID string) *DeleteSparkEngineHistoryServerOptions {
	return &DeleteSparkEngineHistoryServerOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteSparkEngineHistoryServerOptions) SetEngineID(engineID string) *DeleteSparkEngineHistoryServerOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteSparkEngineHistoryServerOptions) SetAuthInstanceID(authInstanceID string) *DeleteSparkEngineHistoryServerOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSparkEngineHistoryServerOptions) SetHeaders(param map[string]string) *DeleteSparkEngineHistoryServerOptions {
	options.Headers = param
	return options
}

// DeleteSparkEngineOptions : The DeleteSparkEngine options.
type DeleteSparkEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteSparkEngineOptions : Instantiate DeleteSparkEngineOptions
func (*WatsonxDataV2) NewDeleteSparkEngineOptions(engineID string) *DeleteSparkEngineOptions {
	return &DeleteSparkEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteSparkEngineOptions) SetEngineID(engineID string) *DeleteSparkEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteSparkEngineOptions) SetAuthInstanceID(authInstanceID string) *DeleteSparkEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSparkEngineOptions) SetHeaders(param map[string]string) *DeleteSparkEngineOptions {
	options.Headers = param
	return options
}

// DeleteTableOptions : The DeleteTable options.
type DeleteTableOptions struct {
	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded table name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteTableOptions : Instantiate DeleteTableOptions
func (*WatsonxDataV2) NewDeleteTableOptions(catalogID string, schemaID string, tableID string, engineID string) *DeleteTableOptions {
	return &DeleteTableOptions{
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
		EngineID: core.StringPtr(engineID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *DeleteTableOptions) SetCatalogID(catalogID string) *DeleteTableOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *DeleteTableOptions) SetSchemaID(schemaID string) *DeleteTableOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *DeleteTableOptions) SetTableID(tableID string) *DeleteTableOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *DeleteTableOptions) SetEngineID(engineID string) *DeleteTableOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeleteTableOptions) SetAuthInstanceID(authInstanceID string) *DeleteTableOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTableOptions) SetHeaders(param map[string]string) *DeleteTableOptions {
	options.Headers = param
	return options
}

// DeregisterBucketOptions : The DeregisterBucket options.
type DeregisterBucketOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeregisterBucketOptions : Instantiate DeregisterBucketOptions
func (*WatsonxDataV2) NewDeregisterBucketOptions(bucketID string) *DeregisterBucketOptions {
	return &DeregisterBucketOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *DeregisterBucketOptions) SetBucketID(bucketID string) *DeregisterBucketOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *DeregisterBucketOptions) SetAuthInstanceID(authInstanceID string) *DeregisterBucketOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeregisterBucketOptions) SetHeaders(param map[string]string) *DeregisterBucketOptions {
	options.Headers = param
	return options
}

// DisplayNameInfoResponse : DisplayNameInfoResponse.
type DisplayNameInfoResponse struct {
	// Display name.
	DisplayName *string `json:"display_name" validate:"required"`
}

// UnmarshalDisplayNameInfoResponse unmarshals an instance of DisplayNameInfoResponse from the specified map of raw messages.
func UnmarshalDisplayNameInfoResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DisplayNameInfoResponse)
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "display_name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Driver : Driver struct
type Driver struct {
	// Connection type.
	ConnectionType *string `json:"connection_type,omitempty"`

	// Driver name.
	DriverID *string `json:"driver_id,omitempty"`

	// Driver name.
	DriverName *string `json:"driver_name,omitempty"`

	// Driver version.
	DriverVersion *string `json:"driver_version,omitempty"`
}

// UnmarshalDriver unmarshals an instance of Driver from the specified map of raw messages.
func UnmarshalDriver(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Driver)
	err = core.UnmarshalPrimitive(m, "connection_type", &obj.ConnectionType)
	if err != nil {
		err = core.SDKErrorf(err, "", "connection_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "driver_id", &obj.DriverID)
	if err != nil {
		err = core.SDKErrorf(err, "", "driver_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "driver_name", &obj.DriverName)
	if err != nil {
		err = core.SDKErrorf(err, "", "driver_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "driver_version", &obj.DriverVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "driver_version-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EngineDetailsBody : Node details.
type EngineDetailsBody struct {
	// api key to work with the saas IAE instance.
	ApiKey *string `json:"api_key,omitempty"`

	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// coordinator/worker property settings.
	Coordinator *NodeDescriptionBody `json:"coordinator,omitempty"`

	// Instance to access the instance.
	InstanceID *string `json:"instance_id,omitempty"`

	// How is the spark instance managed.
	ManagedBy *string `json:"managed_by,omitempty"`

	// Size config.
	SizeConfig *string `json:"size_config,omitempty"`

	// coordinator/worker property settings.
	Worker *NodeDescriptionBody `json:"worker,omitempty"`
}

// Constants associated with the EngineDetailsBody.SizeConfig property.
// Size config.
const (
	EngineDetailsBody_SizeConfig_CacheOptimized = "cache_optimized"
	EngineDetailsBody_SizeConfig_ComputeOptimized = "compute_optimized"
	EngineDetailsBody_SizeConfig_Custom = "custom"
	EngineDetailsBody_SizeConfig_Large = "large"
	EngineDetailsBody_SizeConfig_Medium = "medium"
	EngineDetailsBody_SizeConfig_Small = "small"
	EngineDetailsBody_SizeConfig_Starter = "starter"
)

// UnmarshalEngineDetailsBody unmarshals an instance of EngineDetailsBody from the specified map of raw messages.
func UnmarshalEngineDetailsBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EngineDetailsBody)
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		err = core.SDKErrorf(err, "", "api_key-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		err = core.SDKErrorf(err, "", "connection_string-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "coordinator-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		err = core.SDKErrorf(err, "", "instance_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "managed_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "size_config", &obj.SizeConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "size_config-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "worker-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EnginePropertiesOaiGen1Configuration : Configuration settings.
type EnginePropertiesOaiGen1Configuration struct {
	// coordinator/worker property settings.
	Coordinator *NodeDescriptionBody `json:"coordinator,omitempty"`

	// coordinator/worker property settings.
	Worker *NodeDescriptionBody `json:"worker,omitempty"`
}

// UnmarshalEnginePropertiesOaiGen1Configuration unmarshals an instance of EnginePropertiesOaiGen1Configuration from the specified map of raw messages.
func UnmarshalEnginePropertiesOaiGen1Configuration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnginePropertiesOaiGen1Configuration)
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "coordinator-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "worker-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the EnginePropertiesOaiGen1Configuration
func (enginePropertiesOaiGen1Configuration *EnginePropertiesOaiGen1Configuration) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(enginePropertiesOaiGen1Configuration.Coordinator) {
		_patch["coordinator"] = enginePropertiesOaiGen1Configuration.Coordinator.asPatch()
	}
	if !core.IsNil(enginePropertiesOaiGen1Configuration.Worker) {
		_patch["worker"] = enginePropertiesOaiGen1Configuration.Worker.asPatch()
	}

	return
}

// EnginePropertiesOaiGen1Jvm : JVM settings.
type EnginePropertiesOaiGen1Jvm struct {
	// coordinator/worker property settings.
	Coordinator *NodeDescriptionBody `json:"coordinator,omitempty"`

	// coordinator/worker property settings.
	Worker *NodeDescriptionBody `json:"worker,omitempty"`
}

// UnmarshalEnginePropertiesOaiGen1Jvm unmarshals an instance of EnginePropertiesOaiGen1Jvm from the specified map of raw messages.
func UnmarshalEnginePropertiesOaiGen1Jvm(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnginePropertiesOaiGen1Jvm)
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "coordinator-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "worker-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the EnginePropertiesOaiGen1Jvm
func (enginePropertiesOaiGen1Jvm *EnginePropertiesOaiGen1Jvm) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(enginePropertiesOaiGen1Jvm.Coordinator) {
		_patch["coordinator"] = enginePropertiesOaiGen1Jvm.Coordinator.asPatch()
	}
	if !core.IsNil(enginePropertiesOaiGen1Jvm.Worker) {
		_patch["worker"] = enginePropertiesOaiGen1Jvm.Worker.asPatch()
	}

	return
}

// EnginePropertiesOaiGenConfiguration : Configuration settings for the engine properties.
type EnginePropertiesOaiGenConfiguration struct {
	// Node details.
	Coordinator *PrestissimoNodeDescriptionBody `json:"coordinator,omitempty"`

	// Node details.
	Worker *PrestissimoNodeDescriptionBody `json:"worker,omitempty"`
}

// UnmarshalEnginePropertiesOaiGenConfiguration unmarshals an instance of EnginePropertiesOaiGenConfiguration from the specified map of raw messages.
func UnmarshalEnginePropertiesOaiGenConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnginePropertiesOaiGenConfiguration)
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalPrestissimoNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "coordinator-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalPrestissimoNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "worker-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the EnginePropertiesOaiGenConfiguration
func (enginePropertiesOaiGenConfiguration *EnginePropertiesOaiGenConfiguration) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(enginePropertiesOaiGenConfiguration.Coordinator) {
		_patch["coordinator"] = enginePropertiesOaiGenConfiguration.Coordinator.asPatch()
	}
	if !core.IsNil(enginePropertiesOaiGenConfiguration.Worker) {
		_patch["worker"] = enginePropertiesOaiGenConfiguration.Worker.asPatch()
	}

	return
}

// GetBucketRegistrationOptions : The GetBucketRegistration options.
type GetBucketRegistrationOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetBucketRegistrationOptions : Instantiate GetBucketRegistrationOptions
func (*WatsonxDataV2) NewGetBucketRegistrationOptions(bucketID string) *GetBucketRegistrationOptions {
	return &GetBucketRegistrationOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *GetBucketRegistrationOptions) SetBucketID(bucketID string) *GetBucketRegistrationOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetBucketRegistrationOptions) SetAuthInstanceID(authInstanceID string) *GetBucketRegistrationOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBucketRegistrationOptions) SetHeaders(param map[string]string) *GetBucketRegistrationOptions {
	options.Headers = param
	return options
}

// GetCatalogOptions : The GetCatalog options.
type GetCatalogOptions struct {
	// catalog ID.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetCatalogOptions : Instantiate GetCatalogOptions
func (*WatsonxDataV2) NewGetCatalogOptions(catalogID string) *GetCatalogOptions {
	return &GetCatalogOptions{
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *GetCatalogOptions) SetCatalogID(catalogID string) *GetCatalogOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetCatalogOptions) SetAuthInstanceID(authInstanceID string) *GetCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCatalogOptions) SetHeaders(param map[string]string) *GetCatalogOptions {
	options.Headers = param
	return options
}

// GetDatabaseOptions : The GetDatabase options.
type GetDatabaseOptions struct {
	// database id.
	DatabaseID *string `json:"database_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDatabaseOptions : Instantiate GetDatabaseOptions
func (*WatsonxDataV2) NewGetDatabaseOptions(databaseID string) *GetDatabaseOptions {
	return &GetDatabaseOptions{
		DatabaseID: core.StringPtr(databaseID),
	}
}

// SetDatabaseID : Allow user to set DatabaseID
func (_options *GetDatabaseOptions) SetDatabaseID(databaseID string) *GetDatabaseOptions {
	_options.DatabaseID = core.StringPtr(databaseID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetDatabaseOptions) SetAuthInstanceID(authInstanceID string) *GetDatabaseOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDatabaseOptions) SetHeaders(param map[string]string) *GetDatabaseOptions {
	options.Headers = param
	return options
}

// GetMilvusServiceOptions : The GetMilvusService options.
type GetMilvusServiceOptions struct {
	// service id.
	ServiceID *string `json:"service_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetMilvusServiceOptions : Instantiate GetMilvusServiceOptions
func (*WatsonxDataV2) NewGetMilvusServiceOptions(serviceID string) *GetMilvusServiceOptions {
	return &GetMilvusServiceOptions{
		ServiceID: core.StringPtr(serviceID),
	}
}

// SetServiceID : Allow user to set ServiceID
func (_options *GetMilvusServiceOptions) SetServiceID(serviceID string) *GetMilvusServiceOptions {
	_options.ServiceID = core.StringPtr(serviceID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetMilvusServiceOptions) SetAuthInstanceID(authInstanceID string) *GetMilvusServiceOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetMilvusServiceOptions) SetHeaders(param map[string]string) *GetMilvusServiceOptions {
	options.Headers = param
	return options
}

// GetPrestissimoEngineCatalogOptions : The GetPrestissimoEngineCatalog options.
type GetPrestissimoEngineCatalogOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetPrestissimoEngineCatalogOptions : Instantiate GetPrestissimoEngineCatalogOptions
func (*WatsonxDataV2) NewGetPrestissimoEngineCatalogOptions(engineID string, catalogID string) *GetPrestissimoEngineCatalogOptions {
	return &GetPrestissimoEngineCatalogOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetPrestissimoEngineCatalogOptions) SetEngineID(engineID string) *GetPrestissimoEngineCatalogOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *GetPrestissimoEngineCatalogOptions) SetCatalogID(catalogID string) *GetPrestissimoEngineCatalogOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetPrestissimoEngineCatalogOptions) SetAuthInstanceID(authInstanceID string) *GetPrestissimoEngineCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPrestissimoEngineCatalogOptions) SetHeaders(param map[string]string) *GetPrestissimoEngineCatalogOptions {
	options.Headers = param
	return options
}

// GetPrestissimoEngineOptions : The GetPrestissimoEngine options.
type GetPrestissimoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetPrestissimoEngineOptions : Instantiate GetPrestissimoEngineOptions
func (*WatsonxDataV2) NewGetPrestissimoEngineOptions(engineID string) *GetPrestissimoEngineOptions {
	return &GetPrestissimoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetPrestissimoEngineOptions) SetEngineID(engineID string) *GetPrestissimoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetPrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *GetPrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPrestissimoEngineOptions) SetHeaders(param map[string]string) *GetPrestissimoEngineOptions {
	options.Headers = param
	return options
}

// GetPrestoEngineCatalogOptions : The GetPrestoEngineCatalog options.
type GetPrestoEngineCatalogOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetPrestoEngineCatalogOptions : Instantiate GetPrestoEngineCatalogOptions
func (*WatsonxDataV2) NewGetPrestoEngineCatalogOptions(engineID string, catalogID string) *GetPrestoEngineCatalogOptions {
	return &GetPrestoEngineCatalogOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetPrestoEngineCatalogOptions) SetEngineID(engineID string) *GetPrestoEngineCatalogOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *GetPrestoEngineCatalogOptions) SetCatalogID(catalogID string) *GetPrestoEngineCatalogOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetPrestoEngineCatalogOptions) SetAuthInstanceID(authInstanceID string) *GetPrestoEngineCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPrestoEngineCatalogOptions) SetHeaders(param map[string]string) *GetPrestoEngineCatalogOptions {
	options.Headers = param
	return options
}

// GetPrestoEngineOptions : The GetPrestoEngine options.
type GetPrestoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetPrestoEngineOptions : Instantiate GetPrestoEngineOptions
func (*WatsonxDataV2) NewGetPrestoEngineOptions(engineID string) *GetPrestoEngineOptions {
	return &GetPrestoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetPrestoEngineOptions) SetEngineID(engineID string) *GetPrestoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetPrestoEngineOptions) SetAuthInstanceID(authInstanceID string) *GetPrestoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPrestoEngineOptions) SetHeaders(param map[string]string) *GetPrestoEngineOptions {
	options.Headers = param
	return options
}

// GetSparkEngineApplicationStatusOptions : The GetSparkEngineApplicationStatus options.
type GetSparkEngineApplicationStatusOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Application id.
	ApplicationID *string `json:"application_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetSparkEngineApplicationStatusOptions : Instantiate GetSparkEngineApplicationStatusOptions
func (*WatsonxDataV2) NewGetSparkEngineApplicationStatusOptions(engineID string, applicationID string) *GetSparkEngineApplicationStatusOptions {
	return &GetSparkEngineApplicationStatusOptions{
		EngineID: core.StringPtr(engineID),
		ApplicationID: core.StringPtr(applicationID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetSparkEngineApplicationStatusOptions) SetEngineID(engineID string) *GetSparkEngineApplicationStatusOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetApplicationID : Allow user to set ApplicationID
func (_options *GetSparkEngineApplicationStatusOptions) SetApplicationID(applicationID string) *GetSparkEngineApplicationStatusOptions {
	_options.ApplicationID = core.StringPtr(applicationID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetSparkEngineApplicationStatusOptions) SetAuthInstanceID(authInstanceID string) *GetSparkEngineApplicationStatusOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSparkEngineApplicationStatusOptions) SetHeaders(param map[string]string) *GetSparkEngineApplicationStatusOptions {
	options.Headers = param
	return options
}

// GetSparkEngineCatalogOptions : The GetSparkEngineCatalog options.
type GetSparkEngineCatalogOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetSparkEngineCatalogOptions : Instantiate GetSparkEngineCatalogOptions
func (*WatsonxDataV2) NewGetSparkEngineCatalogOptions(engineID string, catalogID string) *GetSparkEngineCatalogOptions {
	return &GetSparkEngineCatalogOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetSparkEngineCatalogOptions) SetEngineID(engineID string) *GetSparkEngineCatalogOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *GetSparkEngineCatalogOptions) SetCatalogID(catalogID string) *GetSparkEngineCatalogOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetSparkEngineCatalogOptions) SetAuthInstanceID(authInstanceID string) *GetSparkEngineCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSparkEngineCatalogOptions) SetHeaders(param map[string]string) *GetSparkEngineCatalogOptions {
	options.Headers = param
	return options
}

// GetSparkEngineHistoryServerOptions : The GetSparkEngineHistoryServer options.
type GetSparkEngineHistoryServerOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetSparkEngineHistoryServerOptions : Instantiate GetSparkEngineHistoryServerOptions
func (*WatsonxDataV2) NewGetSparkEngineHistoryServerOptions(engineID string) *GetSparkEngineHistoryServerOptions {
	return &GetSparkEngineHistoryServerOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetSparkEngineHistoryServerOptions) SetEngineID(engineID string) *GetSparkEngineHistoryServerOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetSparkEngineHistoryServerOptions) SetAuthInstanceID(authInstanceID string) *GetSparkEngineHistoryServerOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSparkEngineHistoryServerOptions) SetHeaders(param map[string]string) *GetSparkEngineHistoryServerOptions {
	options.Headers = param
	return options
}

// GetSparkEngineOptions : The GetSparkEngine options.
type GetSparkEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetSparkEngineOptions : Instantiate GetSparkEngineOptions
func (*WatsonxDataV2) NewGetSparkEngineOptions(engineID string) *GetSparkEngineOptions {
	return &GetSparkEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *GetSparkEngineOptions) SetEngineID(engineID string) *GetSparkEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetSparkEngineOptions) SetAuthInstanceID(authInstanceID string) *GetSparkEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSparkEngineOptions) SetHeaders(param map[string]string) *GetSparkEngineOptions {
	options.Headers = param
	return options
}

// GetTableOptions : The GetTable options.
type GetTableOptions struct {
	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded table name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetTableOptions : Instantiate GetTableOptions
func (*WatsonxDataV2) NewGetTableOptions(catalogID string, schemaID string, tableID string, engineID string) *GetTableOptions {
	return &GetTableOptions{
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
		EngineID: core.StringPtr(engineID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *GetTableOptions) SetCatalogID(catalogID string) *GetTableOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *GetTableOptions) SetSchemaID(schemaID string) *GetTableOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *GetTableOptions) SetTableID(tableID string) *GetTableOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *GetTableOptions) SetEngineID(engineID string) *GetTableOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *GetTableOptions) SetAuthInstanceID(authInstanceID string) *GetTableOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTableOptions) SetHeaders(param map[string]string) *GetTableOptions {
	options.Headers = param
	return options
}

// IngestionJob : Ingestion job.
type IngestionJob struct {
	// Create new target table (if True); Insert into pre-existing target table (if False).
	CreateIfNotExist *bool `json:"create_if_not_exist,omitempty"`

	// Ingestion CSV properties.
	CsvProperty *IngestionJobCsvProperty `json:"csv_property,omitempty"`

	// Error messages of failed ingestion job.
	Details *string `json:"details,omitempty"`

	// Unix timestamp of ingestion job completing.
	EndTimestamp *string `json:"end_timestamp,omitempty"`

	// ID of the spark engine to be used for ingestion.
	EngineID *string `json:"engine_id,omitempty"`

	// Name of the spark engine to be used for ingestion.
	EngineName *string `json:"engine_name,omitempty"`

	// Ingestion engine configuration.
	ExecuteConfig *IngestionJobExecuteConfig `json:"execute_config,omitempty"`

	// Instance ID of the lakehouse where ingestion job is executed.
	InstanceID *string `json:"instance_id,omitempty"`

	// Job ID of the ingestion job.
	JobID *string `json:"job_id,omitempty"`

	// partition by expression of the target table.
	PartitionBy *string `json:"partition_by,omitempty"`

	// Schema definition of the source table.
	Schema *string `json:"schema,omitempty"`

	// Source data location of the ingestion job.
	SourceDataFiles *string `json:"source_data_files,omitempty"`

	// Source file types (parquet or csv).
	SourceFileType *string `json:"source_file_type,omitempty"`

	// Unix timestamp of ingestion job starting.
	StartTimestamp *string `json:"start_timestamp,omitempty"`

	// Current state of ingestion job.
	Status *string `json:"status,omitempty"`

	// Target table name in format catalog.schema.table.
	TargetTable *string `json:"target_table,omitempty"`

	// Ingestion job user.
	Username *string `json:"username,omitempty"`

	// Validate CSV header if the target table exist.
	ValidateCsvHeader *bool `json:"validate_csv_header,omitempty"`
}

// Constants associated with the IngestionJob.SourceFileType property.
// Source file types (parquet or csv).
const (
	IngestionJob_SourceFileType_Csv = "csv"
	IngestionJob_SourceFileType_Parquet = "parquet"
)

// UnmarshalIngestionJob unmarshals an instance of IngestionJob from the specified map of raw messages.
func UnmarshalIngestionJob(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IngestionJob)
	err = core.UnmarshalPrimitive(m, "create_if_not_exist", &obj.CreateIfNotExist)
	if err != nil {
		err = core.SDKErrorf(err, "", "create_if_not_exist-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "csv_property", &obj.CsvProperty, UnmarshalIngestionJobCsvProperty)
	if err != nil {
		err = core.SDKErrorf(err, "", "csv_property-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "details", &obj.Details)
	if err != nil {
		err = core.SDKErrorf(err, "", "details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "end_timestamp", &obj.EndTimestamp)
	if err != nil {
		err = core.SDKErrorf(err, "", "end_timestamp-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_name", &obj.EngineName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "execute_config", &obj.ExecuteConfig, UnmarshalIngestionJobExecuteConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "execute_config-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		err = core.SDKErrorf(err, "", "instance_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "job_id", &obj.JobID)
	if err != nil {
		err = core.SDKErrorf(err, "", "job_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "partition_by", &obj.PartitionBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "partition_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "schema", &obj.Schema)
	if err != nil {
		err = core.SDKErrorf(err, "", "schema-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "source_data_files", &obj.SourceDataFiles)
	if err != nil {
		err = core.SDKErrorf(err, "", "source_data_files-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "source_file_type", &obj.SourceFileType)
	if err != nil {
		err = core.SDKErrorf(err, "", "source_file_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "start_timestamp", &obj.StartTimestamp)
	if err != nil {
		err = core.SDKErrorf(err, "", "start_timestamp-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "target_table", &obj.TargetTable)
	if err != nil {
		err = core.SDKErrorf(err, "", "target_table-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		err = core.SDKErrorf(err, "", "username-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "validate_csv_header", &obj.ValidateCsvHeader)
	if err != nil {
		err = core.SDKErrorf(err, "", "validate_csv_header-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IngestionJobCollection : List ingestion jobs.
type IngestionJobCollection struct {
	// Ingestion jobs.
	IngestionJobs []IngestionJob `json:"ingestion_jobs,omitempty"`

	// A page in a pagination collection.
	First *IngestionJobCollectionPage `json:"first,omitempty"`

	// A page in a pagination collection.
	Next *IngestionJobCollectionPage `json:"next,omitempty"`
}

// UnmarshalIngestionJobCollection unmarshals an instance of IngestionJobCollection from the specified map of raw messages.
func UnmarshalIngestionJobCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IngestionJobCollection)
	err = core.UnmarshalModel(m, "ingestion_jobs", &obj.IngestionJobs, UnmarshalIngestionJob)
	if err != nil {
		err = core.SDKErrorf(err, "", "ingestion_jobs-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalIngestionJobCollectionPage)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalIngestionJobCollectionPage)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *IngestionJobCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil {
		err = core.SDKErrorf(err, "", "read-query-param-error", common.GetComponentInfo())
		return nil, err
	} else if start == nil {
		return nil, nil
	}
	return start, nil
}

// IngestionJobCollectionPage : A page in a pagination collection.
type IngestionJobCollectionPage struct {
	// Link to the a page in the collection.
	Href *string `json:"href,omitempty"`
}

// UnmarshalIngestionJobCollectionPage unmarshals an instance of IngestionJobCollectionPage from the specified map of raw messages.
func UnmarshalIngestionJobCollectionPage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IngestionJobCollectionPage)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IngestionJobCsvProperty : Ingestion CSV properties.
type IngestionJobCsvProperty struct {
	// Encoding used in CSV file.
	Encoding *string `json:"encoding,omitempty"`

	// Escape character of CSV file.
	EscapeCharacter *string `json:"escape_character,omitempty"`

	// Field delimiter of CSV file.
	FieldDelimiter *string `json:"field_delimiter,omitempty"`

	// Identify if header exists in CSV file.
	Header *bool `json:"header,omitempty"`

	// Line delimiter of CSV file.
	LineDelimiter *string `json:"line_delimiter,omitempty"`
}

// UnmarshalIngestionJobCsvProperty unmarshals an instance of IngestionJobCsvProperty from the specified map of raw messages.
func UnmarshalIngestionJobCsvProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IngestionJobCsvProperty)
	err = core.UnmarshalPrimitive(m, "encoding", &obj.Encoding)
	if err != nil {
		err = core.SDKErrorf(err, "", "encoding-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "escape_character", &obj.EscapeCharacter)
	if err != nil {
		err = core.SDKErrorf(err, "", "escape_character-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "field_delimiter", &obj.FieldDelimiter)
	if err != nil {
		err = core.SDKErrorf(err, "", "field_delimiter-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "header", &obj.Header)
	if err != nil {
		err = core.SDKErrorf(err, "", "header-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "line_delimiter", &obj.LineDelimiter)
	if err != nil {
		err = core.SDKErrorf(err, "", "line_delimiter-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IngestionJobExecuteConfig : Ingestion engine configuration.
type IngestionJobExecuteConfig struct {
	// Driver core(s) configuration for Spark engine.
	DriverCores *int64 `json:"driver_cores,omitempty"`

	// Driver memory configuration (in GB) for Spark engine.
	DriverMemory *string `json:"driver_memory,omitempty"`

	// Executor core(s) configuration for Spark engine.
	ExecutorCores *int64 `json:"executor_cores,omitempty"`

	// Executor memory configuration (in GB) for Spark engine.
	ExecutorMemory *string `json:"executor_memory,omitempty"`

	// Number of executors to assign for Spark engine.
	NumExecutors *int64 `json:"num_executors,omitempty"`
}

// UnmarshalIngestionJobExecuteConfig unmarshals an instance of IngestionJobExecuteConfig from the specified map of raw messages.
func UnmarshalIngestionJobExecuteConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IngestionJobExecuteConfig)
	err = core.UnmarshalPrimitive(m, "driver_cores", &obj.DriverCores)
	if err != nil {
		err = core.SDKErrorf(err, "", "driver_cores-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "driver_memory", &obj.DriverMemory)
	if err != nil {
		err = core.SDKErrorf(err, "", "driver_memory-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "executor_cores", &obj.ExecutorCores)
	if err != nil {
		err = core.SDKErrorf(err, "", "executor_cores-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "executor_memory", &obj.ExecutorMemory)
	if err != nil {
		err = core.SDKErrorf(err, "", "executor_memory-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "num_executors", &obj.NumExecutors)
	if err != nil {
		err = core.SDKErrorf(err, "", "num_executors-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListBucketObjectsOptions : The ListBucketObjects options.
type ListBucketObjectsOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListBucketObjectsOptions : Instantiate ListBucketObjectsOptions
func (*WatsonxDataV2) NewListBucketObjectsOptions(bucketID string) *ListBucketObjectsOptions {
	return &ListBucketObjectsOptions{
		BucketID: core.StringPtr(bucketID),
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *ListBucketObjectsOptions) SetBucketID(bucketID string) *ListBucketObjectsOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListBucketObjectsOptions) SetAuthInstanceID(authInstanceID string) *ListBucketObjectsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListBucketObjectsOptions) SetHeaders(param map[string]string) *ListBucketObjectsOptions {
	options.Headers = param
	return options
}

// ListBucketRegistrationsOptions : The ListBucketRegistrations options.
type ListBucketRegistrationsOptions struct {
	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListBucketRegistrationsOptions : Instantiate ListBucketRegistrationsOptions
func (*WatsonxDataV2) NewListBucketRegistrationsOptions() *ListBucketRegistrationsOptions {
	return &ListBucketRegistrationsOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListBucketRegistrationsOptions) SetAuthInstanceID(authInstanceID string) *ListBucketRegistrationsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListBucketRegistrationsOptions) SetHeaders(param map[string]string) *ListBucketRegistrationsOptions {
	options.Headers = param
	return options
}

// ListCatalogsOptions : The ListCatalogs options.
type ListCatalogsOptions struct {
	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListCatalogsOptions : Instantiate ListCatalogsOptions
func (*WatsonxDataV2) NewListCatalogsOptions() *ListCatalogsOptions {
	return &ListCatalogsOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListCatalogsOptions) SetAuthInstanceID(authInstanceID string) *ListCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListCatalogsOptions) SetHeaders(param map[string]string) *ListCatalogsOptions {
	options.Headers = param
	return options
}

// ListColumnsOptions : The ListColumns options.
type ListColumnsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded schema name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListColumnsOptions : Instantiate ListColumnsOptions
func (*WatsonxDataV2) NewListColumnsOptions(engineID string, catalogID string, schemaID string, tableID string) *ListColumnsOptions {
	return &ListColumnsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListColumnsOptions) SetEngineID(engineID string) *ListColumnsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ListColumnsOptions) SetCatalogID(catalogID string) *ListColumnsOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *ListColumnsOptions) SetSchemaID(schemaID string) *ListColumnsOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *ListColumnsOptions) SetTableID(tableID string) *ListColumnsOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListColumnsOptions) SetAuthInstanceID(authInstanceID string) *ListColumnsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListColumnsOptions) SetHeaders(param map[string]string) *ListColumnsOptions {
	options.Headers = param
	return options
}

// ListDatabaseRegistrationsOptions : The ListDatabaseRegistrations options.
type ListDatabaseRegistrationsOptions struct {
	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListDatabaseRegistrationsOptions : Instantiate ListDatabaseRegistrationsOptions
func (*WatsonxDataV2) NewListDatabaseRegistrationsOptions() *ListDatabaseRegistrationsOptions {
	return &ListDatabaseRegistrationsOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListDatabaseRegistrationsOptions) SetAuthInstanceID(authInstanceID string) *ListDatabaseRegistrationsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDatabaseRegistrationsOptions) SetHeaders(param map[string]string) *ListDatabaseRegistrationsOptions {
	options.Headers = param
	return options
}

// ListDb2EnginesOptions : The ListDb2Engines options.
type ListDb2EnginesOptions struct {
	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListDb2EnginesOptions : Instantiate ListDb2EnginesOptions
func (*WatsonxDataV2) NewListDb2EnginesOptions() *ListDb2EnginesOptions {
	return &ListDb2EnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListDb2EnginesOptions) SetAuthInstanceID(authInstanceID string) *ListDb2EnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDb2EnginesOptions) SetHeaders(param map[string]string) *ListDb2EnginesOptions {
	options.Headers = param
	return options
}

// ListIngestionJobsOptions : The ListIngestionJobs options.
type ListIngestionJobsOptions struct {
	// watsonx.data instance ID.
	AuthInstanceID *string `json:"AuthInstanceId" validate:"required"`

	// Page number of requested ingestion jobs.
	Start *string `json:"start,omitempty"`

	// Number of requested ingestion jobs.
	JobsPerPage *int64 `json:"jobs_per_page,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListIngestionJobsOptions : Instantiate ListIngestionJobsOptions
func (*WatsonxDataV2) NewListIngestionJobsOptions(authInstanceID string) *ListIngestionJobsOptions {
	return &ListIngestionJobsOptions{
		AuthInstanceID: core.StringPtr(authInstanceID),
	}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListIngestionJobsOptions) SetAuthInstanceID(authInstanceID string) *ListIngestionJobsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListIngestionJobsOptions) SetStart(start string) *ListIngestionJobsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetJobsPerPage : Allow user to set JobsPerPage
func (_options *ListIngestionJobsOptions) SetJobsPerPage(jobsPerPage int64) *ListIngestionJobsOptions {
	_options.JobsPerPage = core.Int64Ptr(jobsPerPage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListIngestionJobsOptions) SetHeaders(param map[string]string) *ListIngestionJobsOptions {
	options.Headers = param
	return options
}

// ListMilvusServicesOptions : The ListMilvusServices options.
type ListMilvusServicesOptions struct {
	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListMilvusServicesOptions : Instantiate ListMilvusServicesOptions
func (*WatsonxDataV2) NewListMilvusServicesOptions() *ListMilvusServicesOptions {
	return &ListMilvusServicesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListMilvusServicesOptions) SetAuthInstanceID(authInstanceID string) *ListMilvusServicesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListMilvusServicesOptions) SetHeaders(param map[string]string) *ListMilvusServicesOptions {
	options.Headers = param
	return options
}

// ListNetezzaEnginesOptions : The ListNetezzaEngines options.
type ListNetezzaEnginesOptions struct {
	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListNetezzaEnginesOptions : Instantiate ListNetezzaEnginesOptions
func (*WatsonxDataV2) NewListNetezzaEnginesOptions() *ListNetezzaEnginesOptions {
	return &ListNetezzaEnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListNetezzaEnginesOptions) SetAuthInstanceID(authInstanceID string) *ListNetezzaEnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListNetezzaEnginesOptions) SetHeaders(param map[string]string) *ListNetezzaEnginesOptions {
	options.Headers = param
	return options
}

// ListOtherEnginesOptions : The ListOtherEngines options.
type ListOtherEnginesOptions struct {
	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListOtherEnginesOptions : Instantiate ListOtherEnginesOptions
func (*WatsonxDataV2) NewListOtherEnginesOptions() *ListOtherEnginesOptions {
	return &ListOtherEnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListOtherEnginesOptions) SetAuthInstanceID(authInstanceID string) *ListOtherEnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListOtherEnginesOptions) SetHeaders(param map[string]string) *ListOtherEnginesOptions {
	options.Headers = param
	return options
}

// ListPrestissimoEngineCatalogsOptions : The ListPrestissimoEngineCatalogs options.
type ListPrestissimoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListPrestissimoEngineCatalogsOptions : Instantiate ListPrestissimoEngineCatalogsOptions
func (*WatsonxDataV2) NewListPrestissimoEngineCatalogsOptions(engineID string) *ListPrestissimoEngineCatalogsOptions {
	return &ListPrestissimoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListPrestissimoEngineCatalogsOptions) SetEngineID(engineID string) *ListPrestissimoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListPrestissimoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *ListPrestissimoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPrestissimoEngineCatalogsOptions) SetHeaders(param map[string]string) *ListPrestissimoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// ListPrestissimoEnginesOptions : The ListPrestissimoEngines options.
type ListPrestissimoEnginesOptions struct {
	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListPrestissimoEnginesOptions : Instantiate ListPrestissimoEnginesOptions
func (*WatsonxDataV2) NewListPrestissimoEnginesOptions() *ListPrestissimoEnginesOptions {
	return &ListPrestissimoEnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListPrestissimoEnginesOptions) SetAuthInstanceID(authInstanceID string) *ListPrestissimoEnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPrestissimoEnginesOptions) SetHeaders(param map[string]string) *ListPrestissimoEnginesOptions {
	options.Headers = param
	return options
}

// ListPrestoEngineCatalogsOptions : The ListPrestoEngineCatalogs options.
type ListPrestoEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListPrestoEngineCatalogsOptions : Instantiate ListPrestoEngineCatalogsOptions
func (*WatsonxDataV2) NewListPrestoEngineCatalogsOptions(engineID string) *ListPrestoEngineCatalogsOptions {
	return &ListPrestoEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListPrestoEngineCatalogsOptions) SetEngineID(engineID string) *ListPrestoEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListPrestoEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *ListPrestoEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPrestoEngineCatalogsOptions) SetHeaders(param map[string]string) *ListPrestoEngineCatalogsOptions {
	options.Headers = param
	return options
}

// ListPrestoEnginesOptions : The ListPrestoEngines options.
type ListPrestoEnginesOptions struct {
	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListPrestoEnginesOptions : Instantiate ListPrestoEnginesOptions
func (*WatsonxDataV2) NewListPrestoEnginesOptions() *ListPrestoEnginesOptions {
	return &ListPrestoEnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListPrestoEnginesOptions) SetAuthInstanceID(authInstanceID string) *ListPrestoEnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPrestoEnginesOptions) SetHeaders(param map[string]string) *ListPrestoEnginesOptions {
	options.Headers = param
	return options
}

// ListSchemasOKBody : GetSchemas OK.
type ListSchemasOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// Schemas.
	Schemas []string `json:"schemas" validate:"required"`
}

// UnmarshalListSchemasOKBody unmarshals an instance of ListSchemasOKBody from the specified map of raw messages.
func UnmarshalListSchemasOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListSchemasOKBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "schemas", &obj.Schemas)
	if err != nil {
		err = core.SDKErrorf(err, "", "schemas-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListSchemasOptions : The ListSchemas options.
type ListSchemasOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog name.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListSchemasOptions : Instantiate ListSchemasOptions
func (*WatsonxDataV2) NewListSchemasOptions(engineID string, catalogID string) *ListSchemasOptions {
	return &ListSchemasOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListSchemasOptions) SetEngineID(engineID string) *ListSchemasOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ListSchemasOptions) SetCatalogID(catalogID string) *ListSchemasOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListSchemasOptions) SetAuthInstanceID(authInstanceID string) *ListSchemasOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSchemasOptions) SetHeaders(param map[string]string) *ListSchemasOptions {
	options.Headers = param
	return options
}

// ListSparkEngineApplicationsOptions : The ListSparkEngineApplications options.
type ListSparkEngineApplicationsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// state.
	State []string `json:"state,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListSparkEngineApplicationsOptions : Instantiate ListSparkEngineApplicationsOptions
func (*WatsonxDataV2) NewListSparkEngineApplicationsOptions(engineID string) *ListSparkEngineApplicationsOptions {
	return &ListSparkEngineApplicationsOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListSparkEngineApplicationsOptions) SetEngineID(engineID string) *ListSparkEngineApplicationsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListSparkEngineApplicationsOptions) SetAuthInstanceID(authInstanceID string) *ListSparkEngineApplicationsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetState : Allow user to set State
func (_options *ListSparkEngineApplicationsOptions) SetState(state []string) *ListSparkEngineApplicationsOptions {
	_options.State = state
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSparkEngineApplicationsOptions) SetHeaders(param map[string]string) *ListSparkEngineApplicationsOptions {
	options.Headers = param
	return options
}

// ListSparkEngineCatalogsOptions : The ListSparkEngineCatalogs options.
type ListSparkEngineCatalogsOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListSparkEngineCatalogsOptions : Instantiate ListSparkEngineCatalogsOptions
func (*WatsonxDataV2) NewListSparkEngineCatalogsOptions(engineID string) *ListSparkEngineCatalogsOptions {
	return &ListSparkEngineCatalogsOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListSparkEngineCatalogsOptions) SetEngineID(engineID string) *ListSparkEngineCatalogsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListSparkEngineCatalogsOptions) SetAuthInstanceID(authInstanceID string) *ListSparkEngineCatalogsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSparkEngineCatalogsOptions) SetHeaders(param map[string]string) *ListSparkEngineCatalogsOptions {
	options.Headers = param
	return options
}

// ListSparkEnginesOptions : The ListSparkEngines options.
type ListSparkEnginesOptions struct {
	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListSparkEnginesOptions : Instantiate ListSparkEnginesOptions
func (*WatsonxDataV2) NewListSparkEnginesOptions() *ListSparkEnginesOptions {
	return &ListSparkEnginesOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListSparkEnginesOptions) SetAuthInstanceID(authInstanceID string) *ListSparkEnginesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSparkEnginesOptions) SetHeaders(param map[string]string) *ListSparkEnginesOptions {
	options.Headers = param
	return options
}

// ListSparkVersionsOKBody : List spark version.
type ListSparkVersionsOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// Spark versions list.
	SparkVersions []DisplayNameInfoResponse `json:"spark_versions" validate:"required"`
}

// UnmarshalListSparkVersionsOKBody unmarshals an instance of ListSparkVersionsOKBody from the specified map of raw messages.
func UnmarshalListSparkVersionsOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListSparkVersionsOKBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "spark_versions", &obj.SparkVersions, UnmarshalDisplayNameInfoResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_versions-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListSparkVersionsOptions : The ListSparkVersions options.
type ListSparkVersionsOptions struct {
	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListSparkVersionsOptions : Instantiate ListSparkVersionsOptions
func (*WatsonxDataV2) NewListSparkVersionsOptions() *ListSparkVersionsOptions {
	return &ListSparkVersionsOptions{}
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListSparkVersionsOptions) SetAuthInstanceID(authInstanceID string) *ListSparkVersionsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSparkVersionsOptions) SetHeaders(param map[string]string) *ListSparkVersionsOptions {
	options.Headers = param
	return options
}

// ListTableSnapshotsOptions : The ListTableSnapshots options.
type ListTableSnapshotsOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog ID.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Schema ID.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// Table ID.
	TableID *string `json:"table_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListTableSnapshotsOptions : Instantiate ListTableSnapshotsOptions
func (*WatsonxDataV2) NewListTableSnapshotsOptions(engineID string, catalogID string, schemaID string, tableID string) *ListTableSnapshotsOptions {
	return &ListTableSnapshotsOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ListTableSnapshotsOptions) SetEngineID(engineID string) *ListTableSnapshotsOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ListTableSnapshotsOptions) SetCatalogID(catalogID string) *ListTableSnapshotsOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *ListTableSnapshotsOptions) SetSchemaID(schemaID string) *ListTableSnapshotsOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *ListTableSnapshotsOptions) SetTableID(tableID string) *ListTableSnapshotsOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListTableSnapshotsOptions) SetAuthInstanceID(authInstanceID string) *ListTableSnapshotsOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTableSnapshotsOptions) SetHeaders(param map[string]string) *ListTableSnapshotsOptions {
	options.Headers = param
	return options
}

// ListTablesOptions : The ListTables options.
type ListTablesOptions struct {
	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListTablesOptions : Instantiate ListTablesOptions
func (*WatsonxDataV2) NewListTablesOptions(catalogID string, schemaID string, engineID string) *ListTablesOptions {
	return &ListTablesOptions{
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		EngineID: core.StringPtr(engineID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ListTablesOptions) SetCatalogID(catalogID string) *ListTablesOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *ListTablesOptions) SetSchemaID(schemaID string) *ListTablesOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *ListTablesOptions) SetEngineID(engineID string) *ListTablesOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ListTablesOptions) SetAuthInstanceID(authInstanceID string) *ListTablesOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTablesOptions) SetHeaders(param map[string]string) *ListTablesOptions {
	options.Headers = param
	return options
}

// MilvusService : milvus service details.
type MilvusService struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Service description.
	Description *string `json:"description,omitempty"`

	// milvus grpc_host.
	GrpcHost *string `json:"grpc_host,omitempty"`

	// milvus port.
	GrpcPort *int64 `json:"grpc_port,omitempty"`

	// milvus display name.
	HostName *string `json:"host_name,omitempty"`

	// milvus https_host.
	HttpsHost *string `json:"https_host,omitempty"`

	// milvus port.
	HttpsPort *int64 `json:"https_port,omitempty"`

	// Origin - place holder.
	Origin *string `json:"origin,omitempty"`

	// Service display name.
	ServiceDisplayName *string `json:"service_display_name,omitempty"`

	// Service programmatic name.
	ServiceID *string `json:"service_id,omitempty"`

	// milvus status.
	Status *string `json:"status,omitempty"`

	// milvus status code.
	StatusCode *int64 `json:"status_code" validate:"required"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// service type.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the MilvusService.Status property.
// milvus status.
const (
	MilvusService_Status_Pending = "pending"
	MilvusService_Status_Running = "running"
	MilvusService_Status_Stopped = "stopped"
)

// UnmarshalMilvusService unmarshals an instance of MilvusService from the specified map of raw messages.
func UnmarshalMilvusService(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MilvusService)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		err = core.SDKErrorf(err, "", "actions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_on-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "grpc_host", &obj.GrpcHost)
	if err != nil {
		err = core.SDKErrorf(err, "", "grpc_host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "grpc_port", &obj.GrpcPort)
	if err != nil {
		err = core.SDKErrorf(err, "", "grpc_port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		err = core.SDKErrorf(err, "", "host_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "https_host", &obj.HttpsHost)
	if err != nil {
		err = core.SDKErrorf(err, "", "https_host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "https_port", &obj.HttpsPort)
	if err != nil {
		err = core.SDKErrorf(err, "", "https_port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		err = core.SDKErrorf(err, "", "origin-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "service_display_name", &obj.ServiceDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "service_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "service_id", &obj.ServiceID)
	if err != nil {
		err = core.SDKErrorf(err, "", "service_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "status_code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MilvusServiceCollection : List milvus services.
type MilvusServiceCollection struct {
	// milvus service body.
	MilvusServices []MilvusService `json:"milvus_services,omitempty"`
}

// UnmarshalMilvusServiceCollection unmarshals an instance of MilvusServiceCollection from the specified map of raw messages.
func UnmarshalMilvusServiceCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MilvusServiceCollection)
	err = core.UnmarshalModel(m, "milvus_services", &obj.MilvusServices, UnmarshalMilvusService)
	if err != nil {
		err = core.SDKErrorf(err, "", "milvus_services-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MilvusServicePatch : UpdateService body.
type MilvusServicePatch struct {
	// Modified description.
	Description *string `json:"description,omitempty"`

	// Service display name.
	ServiceDisplayName *string `json:"service_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`
}

// UnmarshalMilvusServicePatch unmarshals an instance of MilvusServicePatch from the specified map of raw messages.
func UnmarshalMilvusServicePatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MilvusServicePatch)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "service_display_name", &obj.ServiceDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "service_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the MilvusServicePatch
func (milvusServicePatch *MilvusServicePatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(milvusServicePatch.Description) {
		_patch["description"] = milvusServicePatch.Description
	}
	if !core.IsNil(milvusServicePatch.ServiceDisplayName) {
		_patch["service_display_name"] = milvusServicePatch.ServiceDisplayName
	}
	if !core.IsNil(milvusServicePatch.Tags) {
		_patch["tags"] = milvusServicePatch.Tags
	}

	return
}

// NetezzaEngine : Netezza engine details.
type NetezzaEngine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// watsonx.data build version.
	BuildVersion *string `json:"build_version,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *NetezzaEngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Engine host name. In case of OCP based clusters, this is internal hostname.
	HostName *string `json:"host_name,omitempty"`

	// Origin - place holder.
	Origin *string `json:"origin,omitempty"`

	// Engine port.
	Port *int64 `json:"port,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Engine type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalNetezzaEngine unmarshals an instance of NetezzaEngine from the specified map of raw messages.
func UnmarshalNetezzaEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NetezzaEngine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		err = core.SDKErrorf(err, "", "actions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "build_version", &obj.BuildVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "build_version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_on-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalNetezzaEngineDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		err = core.SDKErrorf(err, "", "host_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		err = core.SDKErrorf(err, "", "origin-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NetezzaEngineCollection : list Netezza engines.
type NetezzaEngineCollection struct {
	// list Netezza engines.
	NetezzaEngines []NetezzaEngine `json:"netezza_engines,omitempty"`
}

// UnmarshalNetezzaEngineCollection unmarshals an instance of NetezzaEngineCollection from the specified map of raw messages.
func UnmarshalNetezzaEngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NetezzaEngineCollection)
	err = core.UnmarshalModel(m, "netezza_engines", &obj.NetezzaEngines, UnmarshalNetezzaEngine)
	if err != nil {
		err = core.SDKErrorf(err, "", "netezza_engines-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NetezzaEngineDetails : External engine details.
type NetezzaEngineDetails struct {
	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// Metastore host.
	MetastoreHost *string `json:"metastore_host,omitempty"`
}

// UnmarshalNetezzaEngineDetails unmarshals an instance of NetezzaEngineDetails from the specified map of raw messages.
func UnmarshalNetezzaEngineDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NetezzaEngineDetails)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		err = core.SDKErrorf(err, "", "connection_string-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "metastore_host", &obj.MetastoreHost)
	if err != nil {
		err = core.SDKErrorf(err, "", "metastore_host-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NetezzaEngineDetailsBody : External engine details.
type NetezzaEngineDetailsBody struct {
	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`
}

// UnmarshalNetezzaEngineDetailsBody unmarshals an instance of NetezzaEngineDetailsBody from the specified map of raw messages.
func UnmarshalNetezzaEngineDetailsBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NetezzaEngineDetailsBody)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		err = core.SDKErrorf(err, "", "connection_string-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NetezzaEnginePatch : UpdateEngine body.
type NetezzaEnginePatch struct {
	// Modified description.
	Description *string `json:"description,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`
}

// UnmarshalNetezzaEnginePatch unmarshals an instance of NetezzaEnginePatch from the specified map of raw messages.
func UnmarshalNetezzaEnginePatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NetezzaEnginePatch)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the NetezzaEnginePatch
func (netezzaEnginePatch *NetezzaEnginePatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(netezzaEnginePatch.Description) {
		_patch["description"] = netezzaEnginePatch.Description
	}
	if !core.IsNil(netezzaEnginePatch.EngineDisplayName) {
		_patch["engine_display_name"] = netezzaEnginePatch.EngineDisplayName
	}
	if !core.IsNil(netezzaEnginePatch.Tags) {
		_patch["tags"] = netezzaEnginePatch.Tags
	}

	return
}

// NodeDescription : NodeDescription.
type NodeDescription struct {
	// Node type.
	NodeType *string `json:"node_type,omitempty"`

	// Quantity.
	Quantity *int64 `json:"quantity,omitempty"`
}

// UnmarshalNodeDescription unmarshals an instance of NodeDescription from the specified map of raw messages.
func UnmarshalNodeDescription(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NodeDescription)
	err = core.UnmarshalPrimitive(m, "node_type", &obj.NodeType)
	if err != nil {
		err = core.SDKErrorf(err, "", "node_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity", &obj.Quantity)
	if err != nil {
		err = core.SDKErrorf(err, "", "quantity-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NodeDescriptionBody : coordinator/worker property settings.
type NodeDescriptionBody struct {
	// Node Type, r5, m, i..
	NodeType *string `json:"node_type,omitempty"`

	// Number of nodes.
	Quantity *int64 `json:"quantity,omitempty"`
}

// UnmarshalNodeDescriptionBody unmarshals an instance of NodeDescriptionBody from the specified map of raw messages.
func UnmarshalNodeDescriptionBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NodeDescriptionBody)
	err = core.UnmarshalPrimitive(m, "node_type", &obj.NodeType)
	if err != nil {
		err = core.SDKErrorf(err, "", "node_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity", &obj.Quantity)
	if err != nil {
		err = core.SDKErrorf(err, "", "quantity-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the NodeDescriptionBody
func (nodeDescriptionBody *NodeDescriptionBody) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(nodeDescriptionBody.NodeType) {
		_patch["node_type"] = nodeDescriptionBody.NodeType
	}
	if !core.IsNil(nodeDescriptionBody.Quantity) {
		_patch["quantity"] = nodeDescriptionBody.Quantity
	}

	return
}

// OtherEngine : external engine details.
type OtherEngine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *OtherEngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Engine host name. In case of OCP based clusters, this is internal hostname.
	HostName *string `json:"host_name,omitempty"`

	// origin.
	Origin *string `json:"origin,omitempty"`

	// Engine port.
	Port *int64 `json:"port,omitempty"`

	// engine status.
	Status *string `json:"status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Type like presto, netezza, external,..
	Type *string `json:"type,omitempty"`
}

// UnmarshalOtherEngine unmarshals an instance of OtherEngine from the specified map of raw messages.
func UnmarshalOtherEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OtherEngine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		err = core.SDKErrorf(err, "", "actions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_on-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalOtherEngineDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		err = core.SDKErrorf(err, "", "host_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		err = core.SDKErrorf(err, "", "origin-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OtherEngineCollection : list other engines.
type OtherEngineCollection struct {
	// list other engines.
	OtherEngines []OtherEngine `json:"other_engines,omitempty"`
}

// UnmarshalOtherEngineCollection unmarshals an instance of OtherEngineCollection from the specified map of raw messages.
func UnmarshalOtherEngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OtherEngineCollection)
	err = core.UnmarshalModel(m, "other_engines", &obj.OtherEngines, UnmarshalOtherEngine)
	if err != nil {
		err = core.SDKErrorf(err, "", "other_engines-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OtherEngineDetails : External engine details.
type OtherEngineDetails struct {
	// external engine connection string.
	ConnectionString *string `json:"connection_string" validate:"required"`

	// Actual engine type.
	EngineType *string `json:"engine_type" validate:"required"`

	// metastore host - not required while registering an engine.
	MetastoreHost *string `json:"metastore_host,omitempty"`
}

// UnmarshalOtherEngineDetails unmarshals an instance of OtherEngineDetails from the specified map of raw messages.
func UnmarshalOtherEngineDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OtherEngineDetails)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		err = core.SDKErrorf(err, "", "connection_string-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_type", &obj.EngineType)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "metastore_host", &obj.MetastoreHost)
	if err != nil {
		err = core.SDKErrorf(err, "", "metastore_host-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OtherEngineDetailsBody : External engine details.
type OtherEngineDetailsBody struct {
	// External engine connection string.
	ConnectionString *string `json:"connection_string" validate:"required"`

	// Actual engine type.
	EngineType *string `json:"engine_type" validate:"required"`
}

// NewOtherEngineDetailsBody : Instantiate OtherEngineDetailsBody (Generic Model Constructor)
func (*WatsonxDataV2) NewOtherEngineDetailsBody(connectionString string, engineType string) (_model *OtherEngineDetailsBody, err error) {
	_model = &OtherEngineDetailsBody{
		ConnectionString: core.StringPtr(connectionString),
		EngineType: core.StringPtr(engineType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalOtherEngineDetailsBody unmarshals an instance of OtherEngineDetailsBody from the specified map of raw messages.
func UnmarshalOtherEngineDetailsBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OtherEngineDetailsBody)
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		err = core.SDKErrorf(err, "", "connection_string-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_type", &obj.EngineType)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PausePrestissimoEngineOptions : The PausePrestissimoEngine options.
type PausePrestissimoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewPausePrestissimoEngineOptions : Instantiate PausePrestissimoEngineOptions
func (*WatsonxDataV2) NewPausePrestissimoEngineOptions(engineID string) *PausePrestissimoEngineOptions {
	return &PausePrestissimoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *PausePrestissimoEngineOptions) SetEngineID(engineID string) *PausePrestissimoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *PausePrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *PausePrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PausePrestissimoEngineOptions) SetHeaders(param map[string]string) *PausePrestissimoEngineOptions {
	options.Headers = param
	return options
}

// PausePrestoEngineOptions : The PausePrestoEngine options.
type PausePrestoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewPausePrestoEngineOptions : Instantiate PausePrestoEngineOptions
func (*WatsonxDataV2) NewPausePrestoEngineOptions(engineID string) *PausePrestoEngineOptions {
	return &PausePrestoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *PausePrestoEngineOptions) SetEngineID(engineID string) *PausePrestoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *PausePrestoEngineOptions) SetAuthInstanceID(authInstanceID string) *PausePrestoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PausePrestoEngineOptions) SetHeaders(param map[string]string) *PausePrestoEngineOptions {
	options.Headers = param
	return options
}

// PrestissimoEndpoints : Endpoints.
type PrestissimoEndpoints struct {
	// Application API.
	ApplicationsApi *string `json:"applications_api,omitempty"`

	// History server endpoint.
	HistoryServerEndpoint *string `json:"history_server_endpoint,omitempty"`

	// Spark kernel endpoint.
	SparkAccessEndpoint *string `json:"spark_access_endpoint,omitempty"`

	// Spark jobs V4 endpoint.
	SparkJobsV4Endpoint *string `json:"spark_jobs_v4_endpoint,omitempty"`

	// Spark kernel endpoint.
	SparkKernelEndpoint *string `json:"spark_kernel_endpoint,omitempty"`

	// View history server.
	ViewHistoryServer *string `json:"view_history_server,omitempty"`

	// Wxd application endpoint.
	WxdApplicationEndpoint *string `json:"wxd_application_endpoint,omitempty"`
}

// UnmarshalPrestissimoEndpoints unmarshals an instance of PrestissimoEndpoints from the specified map of raw messages.
func UnmarshalPrestissimoEndpoints(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEndpoints)
	err = core.UnmarshalPrimitive(m, "applications_api", &obj.ApplicationsApi)
	if err != nil {
		err = core.SDKErrorf(err, "", "applications_api-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "history_server_endpoint", &obj.HistoryServerEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "history_server_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_access_endpoint", &obj.SparkAccessEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_access_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_jobs_v4_endpoint", &obj.SparkJobsV4Endpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_jobs_v4_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_kernel_endpoint", &obj.SparkKernelEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_kernel_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "view_history_server", &obj.ViewHistoryServer)
	if err != nil {
		err = core.SDKErrorf(err, "", "view_history_server-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "wxd_application_endpoint", &obj.WxdApplicationEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "wxd_application_endpoint-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrestissimoEngine : EngineDetail.
type PrestissimoEngine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// Associated catalog.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// watsonx.data build version.
	BuildVersion *string `json:"build_version,omitempty"`

	// Node details.
	Coordinator *PrestissimoNodeDescriptionBody `json:"coordinator,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *PrestissimoEngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Engine properties.
	EngineProperties *PrestissimoEngineEngineProperties `json:"engine_properties,omitempty"`

	// Triggers engine restart if value is force.
	EngineRestart *string `json:"engine_restart,omitempty"`

	// Applicable only for OCP based clusters.  This is typically  servicename+route.
	ExternalHostName *string `json:"external_host_name" validate:"required"`

	// Group ID.
	GroupID *string `json:"group_id,omitempty"`

	// Engine host name. In case of OCP based clusters, this is internal hostname.
	HostName *string `json:"host_name,omitempty"`

	// Origin - place holder.
	Origin *string `json:"origin,omitempty"`

	// Engine port.
	Port *int64 `json:"port,omitempty"`

	// Region - place holder.
	Region *string `json:"region,omitempty"`

	// RemoveEngine properties.
	RemoveEngineProperties *RemoveEngineProperties `json:"remove_engine_properties,omitempty"`

	// Size config.
	SizeConfig *string `json:"size_config,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Engine status code.
	StatusCode *int64 `json:"status_code" validate:"required"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Engine type.
	Type *string `json:"type,omitempty"`

	// Version of the engine.
	Version *string `json:"version,omitempty"`

	// Node details.
	Worker *PrestissimoNodeDescriptionBody `json:"worker,omitempty"`
}

// Constants associated with the PrestissimoEngine.EngineRestart property.
// Triggers engine restart if value is force.
const (
	PrestissimoEngine_EngineRestart_False = "false"
	PrestissimoEngine_EngineRestart_Force = "force"
)

// Constants associated with the PrestissimoEngine.Origin property.
// Origin - place holder.
const (
	PrestissimoEngine_Origin_Discover = "discover"
	PrestissimoEngine_Origin_External = "external"
	PrestissimoEngine_Origin_Native = "native"
)

// Constants associated with the PrestissimoEngine.Status property.
// Engine status.
const (
	PrestissimoEngine_Status_Pending = "pending"
	PrestissimoEngine_Status_Running = "running"
	PrestissimoEngine_Status_Stopped = "stopped"
)

// UnmarshalPrestissimoEngine unmarshals an instance of PrestissimoEngine from the specified map of raw messages.
func UnmarshalPrestissimoEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEngine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		err = core.SDKErrorf(err, "", "actions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_catalogs", &obj.AssociatedCatalogs)
	if err != nil {
		err = core.SDKErrorf(err, "", "associated_catalogs-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "build_version", &obj.BuildVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "build_version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalPrestissimoNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "coordinator-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_on-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalPrestissimoEngineDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "engine_properties", &obj.EngineProperties, UnmarshalPrestissimoEngineEngineProperties)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_restart", &obj.EngineRestart)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_restart-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "external_host_name", &obj.ExternalHostName)
	if err != nil {
		err = core.SDKErrorf(err, "", "external_host_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "group_id", &obj.GroupID)
	if err != nil {
		err = core.SDKErrorf(err, "", "group_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		err = core.SDKErrorf(err, "", "host_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		err = core.SDKErrorf(err, "", "origin-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		err = core.SDKErrorf(err, "", "region-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "remove_engine_properties", &obj.RemoveEngineProperties, UnmarshalRemoveEngineProperties)
	if err != nil {
		err = core.SDKErrorf(err, "", "remove_engine_properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "size_config", &obj.SizeConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "size_config-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "status_code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalPrestissimoNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "worker-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrestissimoEngineCollection : list Prestissimo Engines.
type PrestissimoEngineCollection struct {
	// list prestissimo engines.
	PrestissimoEngines []PrestissimoEngine `json:"prestissimo_engines,omitempty"`
}

// UnmarshalPrestissimoEngineCollection unmarshals an instance of PrestissimoEngineCollection from the specified map of raw messages.
func UnmarshalPrestissimoEngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEngineCollection)
	err = core.UnmarshalModel(m, "prestissimo_engines", &obj.PrestissimoEngines, UnmarshalPrestissimoEngine)
	if err != nil {
		err = core.SDKErrorf(err, "", "prestissimo_engines-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrestissimoEngineDetails : External engine details.
type PrestissimoEngineDetails struct {
	// api key to work with the saas IAE instance.
	ApiKey *string `json:"api_key,omitempty"`

	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// Node details.
	Coordinator *PrestissimoNodeDescriptionBody `json:"coordinator,omitempty"`

	// Endpoints.
	Endpoints *PrestissimoEndpoints `json:"endpoints,omitempty"`

	// Instance to access the instance.
	InstanceID *string `json:"instance_id,omitempty"`

	// How is the spark instance managed.
	ManagedBy *string `json:"managed_by,omitempty"`

	// Metastore host.
	MetastoreHost *string `json:"metastore_host,omitempty"`

	// Size config.
	SizeConfig *string `json:"size_config,omitempty"`

	// Node details.
	Worker *PrestissimoNodeDescriptionBody `json:"worker,omitempty"`
}

// Constants associated with the PrestissimoEngineDetails.SizeConfig property.
// Size config.
const (
	PrestissimoEngineDetails_SizeConfig_CacheOptimized = "cache_optimized"
	PrestissimoEngineDetails_SizeConfig_ComputeOptimized = "compute_optimized"
	PrestissimoEngineDetails_SizeConfig_Custom = "custom"
	PrestissimoEngineDetails_SizeConfig_Large = "large"
	PrestissimoEngineDetails_SizeConfig_Medium = "medium"
	PrestissimoEngineDetails_SizeConfig_Small = "small"
	PrestissimoEngineDetails_SizeConfig_Starter = "starter"
)

// UnmarshalPrestissimoEngineDetails unmarshals an instance of PrestissimoEngineDetails from the specified map of raw messages.
func UnmarshalPrestissimoEngineDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEngineDetails)
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		err = core.SDKErrorf(err, "", "api_key-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		err = core.SDKErrorf(err, "", "connection_string-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalPrestissimoNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "coordinator-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "endpoints", &obj.Endpoints, UnmarshalPrestissimoEndpoints)
	if err != nil {
		err = core.SDKErrorf(err, "", "endpoints-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		err = core.SDKErrorf(err, "", "instance_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "managed_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "metastore_host", &obj.MetastoreHost)
	if err != nil {
		err = core.SDKErrorf(err, "", "metastore_host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "size_config", &obj.SizeConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "size_config-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalPrestissimoNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "worker-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrestissimoEngineEngineProperties : Engine properties.
type PrestissimoEngineEngineProperties struct {
	// catalog properties.
	Catalog *PrestissimoEnginePropertiesCatalog `json:"catalog,omitempty"`

	// Configuration settings for the engine properties.
	Configuration *EnginePropertiesOaiGenConfiguration `json:"configuration,omitempty"`

	// velox properties.
	Velox *PrestissimoEnginePropertiesVelox `json:"velox,omitempty"`

	// JVM settings.
	Jvm *PrestissimoEnginePropertiesOaiGen1Jvm `json:"jvm,omitempty"`
}

// UnmarshalPrestissimoEngineEngineProperties unmarshals an instance of PrestissimoEngineEngineProperties from the specified map of raw messages.
func UnmarshalPrestissimoEngineEngineProperties(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEngineEngineProperties)
	err = core.UnmarshalModel(m, "catalog", &obj.Catalog, UnmarshalPrestissimoEnginePropertiesCatalog)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "configuration", &obj.Configuration, UnmarshalEnginePropertiesOaiGenConfiguration)
	if err != nil {
		err = core.SDKErrorf(err, "", "configuration-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "velox", &obj.Velox, UnmarshalPrestissimoEnginePropertiesVelox)
	if err != nil {
		err = core.SDKErrorf(err, "", "velox-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "jvm", &obj.Jvm, UnmarshalPrestissimoEnginePropertiesOaiGen1Jvm)
	if err != nil {
		err = core.SDKErrorf(err, "", "jvm-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the PrestissimoEngineEngineProperties
func (prestissimoEngineEngineProperties *PrestissimoEngineEngineProperties) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(prestissimoEngineEngineProperties.Catalog) {
		_patch["catalog"] = prestissimoEngineEngineProperties.Catalog.asPatch()
	}
	if !core.IsNil(prestissimoEngineEngineProperties.Configuration) {
		_patch["configuration"] = prestissimoEngineEngineProperties.Configuration.asPatch()
	}
	if !core.IsNil(prestissimoEngineEngineProperties.Velox) {
		_patch["velox"] = prestissimoEngineEngineProperties.Velox.asPatch()
	}
	if !core.IsNil(prestissimoEngineEngineProperties.Jvm) {
		_patch["jvm"] = prestissimoEngineEngineProperties.Jvm.asPatch()
	}

	return
}

// PrestissimoEnginePatch : Update prestissimo engine body.
type PrestissimoEnginePatch struct {
	// Modified description.
	Description *string `json:"description,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine properties.
	EngineProperties *PrestissimoEngineEngineProperties `json:"engine_properties,omitempty"`

	// Triggers engine restart if value is force.
	EngineRestart *string `json:"engine_restart,omitempty"`

	// RemoveEngine properties.
	RemoveEngineProperties *RemoveEngineProperties `json:"remove_engine_properties,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`
}

// Constants associated with the PrestissimoEnginePatch.EngineRestart property.
// Triggers engine restart if value is force.
const (
	PrestissimoEnginePatch_EngineRestart_False = "false"
	PrestissimoEnginePatch_EngineRestart_Force = "force"
)

// UnmarshalPrestissimoEnginePatch unmarshals an instance of PrestissimoEnginePatch from the specified map of raw messages.
func UnmarshalPrestissimoEnginePatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEnginePatch)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "engine_properties", &obj.EngineProperties, UnmarshalPrestissimoEngineEngineProperties)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_restart", &obj.EngineRestart)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_restart-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "remove_engine_properties", &obj.RemoveEngineProperties, UnmarshalRemoveEngineProperties)
	if err != nil {
		err = core.SDKErrorf(err, "", "remove_engine_properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the PrestissimoEnginePatch
func (prestissimoEnginePatch *PrestissimoEnginePatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(prestissimoEnginePatch.Description) {
		_patch["description"] = prestissimoEnginePatch.Description
	}
	if !core.IsNil(prestissimoEnginePatch.EngineDisplayName) {
		_patch["engine_display_name"] = prestissimoEnginePatch.EngineDisplayName
	}
	if !core.IsNil(prestissimoEnginePatch.EngineProperties) {
		_patch["engine_properties"] = prestissimoEnginePatch.EngineProperties.asPatch()
	}
	if !core.IsNil(prestissimoEnginePatch.EngineRestart) {
		_patch["engine_restart"] = prestissimoEnginePatch.EngineRestart
	}
	if !core.IsNil(prestissimoEnginePatch.RemoveEngineProperties) {
		_patch["remove_engine_properties"] = prestissimoEnginePatch.RemoveEngineProperties.asPatch()
	}
	if !core.IsNil(prestissimoEnginePatch.Tags) {
		_patch["tags"] = prestissimoEnginePatch.Tags
	}

	return
}

// PrestissimoEnginePropertiesCatalog : catalog properties.
type PrestissimoEnginePropertiesCatalog struct {
	// catalog name.
	CatalogName []string `json:"catalog_name,omitempty"`
}

// UnmarshalPrestissimoEnginePropertiesCatalog unmarshals an instance of PrestissimoEnginePropertiesCatalog from the specified map of raw messages.
func UnmarshalPrestissimoEnginePropertiesCatalog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEnginePropertiesCatalog)
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the PrestissimoEnginePropertiesCatalog
func (prestissimoEnginePropertiesCatalog *PrestissimoEnginePropertiesCatalog) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(prestissimoEnginePropertiesCatalog.CatalogName) {
		_patch["catalog_name"] = prestissimoEnginePropertiesCatalog.CatalogName
	}

	return
}

// PrestissimoEnginePropertiesOaiGen1Jvm : JVM settings.
type PrestissimoEnginePropertiesOaiGen1Jvm struct {
	// coordinator/worker property settings.
	Coordinator *NodeDescriptionBody `json:"coordinator,omitempty"`
}

// UnmarshalPrestissimoEnginePropertiesOaiGen1Jvm unmarshals an instance of PrestissimoEnginePropertiesOaiGen1Jvm from the specified map of raw messages.
func UnmarshalPrestissimoEnginePropertiesOaiGen1Jvm(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEnginePropertiesOaiGen1Jvm)
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalNodeDescriptionBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "coordinator-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the PrestissimoEnginePropertiesOaiGen1Jvm
func (prestissimoEnginePropertiesOaiGen1Jvm *PrestissimoEnginePropertiesOaiGen1Jvm) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(prestissimoEnginePropertiesOaiGen1Jvm.Coordinator) {
		_patch["coordinator"] = prestissimoEnginePropertiesOaiGen1Jvm.Coordinator.asPatch()
	}

	return
}

// PrestissimoEnginePropertiesVelox : velox properties.
type PrestissimoEnginePropertiesVelox struct {
	// velox property.
	VeloxProperty []string `json:"velox_property,omitempty"`
}

// UnmarshalPrestissimoEnginePropertiesVelox unmarshals an instance of PrestissimoEnginePropertiesVelox from the specified map of raw messages.
func UnmarshalPrestissimoEnginePropertiesVelox(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoEnginePropertiesVelox)
	err = core.UnmarshalPrimitive(m, "velox_property", &obj.VeloxProperty)
	if err != nil {
		err = core.SDKErrorf(err, "", "velox_property-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the PrestissimoEnginePropertiesVelox
func (prestissimoEnginePropertiesVelox *PrestissimoEnginePropertiesVelox) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(prestissimoEnginePropertiesVelox.VeloxProperty) {
		_patch["velox_property"] = prestissimoEnginePropertiesVelox.VeloxProperty
	}

	return
}

// PrestissimoNodeDescriptionBody : Node details.
type PrestissimoNodeDescriptionBody struct {
	// Node Type, r5, m, i..
	NodeType *string `json:"node_type,omitempty"`

	// Number of nodes.
	Quantity *int64 `json:"quantity,omitempty"`
}

// UnmarshalPrestissimoNodeDescriptionBody unmarshals an instance of PrestissimoNodeDescriptionBody from the specified map of raw messages.
func UnmarshalPrestissimoNodeDescriptionBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestissimoNodeDescriptionBody)
	err = core.UnmarshalPrimitive(m, "node_type", &obj.NodeType)
	if err != nil {
		err = core.SDKErrorf(err, "", "node_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity", &obj.Quantity)
	if err != nil {
		err = core.SDKErrorf(err, "", "quantity-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the PrestissimoNodeDescriptionBody
func (prestissimoNodeDescriptionBody *PrestissimoNodeDescriptionBody) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(prestissimoNodeDescriptionBody.NodeType) {
		_patch["node_type"] = prestissimoNodeDescriptionBody.NodeType
	}
	if !core.IsNil(prestissimoNodeDescriptionBody.Quantity) {
		_patch["quantity"] = prestissimoNodeDescriptionBody.Quantity
	}

	return
}

// PrestoEngine : EngineDetail.
type PrestoEngine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// Associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// watsonx.data build version.
	BuildVersion *string `json:"build_version,omitempty"`

	// NodeDescription.
	Coordinator *NodeDescription `json:"coordinator,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// Driver details.
	Drivers []Driver `json:"drivers,omitempty"`

	// Node details.
	EngineDetails *EngineDetailsBody `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Engine properties.
	EngineProperties *PrestoEngineEngineProperties `json:"engine_properties,omitempty"`

	// Triggers engine restart if value is force.
	EngineRestart *string `json:"engine_restart,omitempty"`

	// Applicable only for OCP based clusters.  This is typically  servicename+route.
	ExternalHostName *string `json:"external_host_name" validate:"required"`

	// Group ID.
	GroupID *string `json:"group_id,omitempty"`

	// Engine host name. In case of OCP based clusters, this is internal hostname.
	HostName *string `json:"host_name,omitempty"`

	// Origin - created or registered.
	Origin *string `json:"origin,omitempty"`

	// Engine port.
	Port *int64 `json:"port,omitempty"`

	// Region (cloud).
	Region *string `json:"region,omitempty"`

	// RemoveEngine properties.
	RemoveEngineProperties *PrestoEnginePatchRemoveEngineProperties `json:"remove_engine_properties,omitempty"`

	// Size config.
	SizeConfig *string `json:"size_config,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Engine status code.
	StatusCode *int64 `json:"status_code" validate:"required"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Engine type presto.
	Type *string `json:"type,omitempty"`

	// Version of the engine.
	Version *string `json:"version,omitempty"`

	// NodeDescription.
	Worker *NodeDescription `json:"worker,omitempty"`
}

// Constants associated with the PrestoEngine.EngineRestart property.
// Triggers engine restart if value is force.
const (
	PrestoEngine_EngineRestart_False = "false"
	PrestoEngine_EngineRestart_Force = "force"
)

// Constants associated with the PrestoEngine.Origin property.
// Origin - created or registered.
const (
	PrestoEngine_Origin_Discover = "discover"
	PrestoEngine_Origin_External = "external"
	PrestoEngine_Origin_Native = "native"
)

// Constants associated with the PrestoEngine.Status property.
// Engine status.
const (
	PrestoEngine_Status_Pending = "pending"
	PrestoEngine_Status_Running = "running"
	PrestoEngine_Status_Stopped = "stopped"
)

// UnmarshalPrestoEngine unmarshals an instance of PrestoEngine from the specified map of raw messages.
func UnmarshalPrestoEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestoEngine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		err = core.SDKErrorf(err, "", "actions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_catalogs", &obj.AssociatedCatalogs)
	if err != nil {
		err = core.SDKErrorf(err, "", "associated_catalogs-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "build_version", &obj.BuildVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "build_version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "coordinator", &obj.Coordinator, UnmarshalNodeDescription)
	if err != nil {
		err = core.SDKErrorf(err, "", "coordinator-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_on-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "drivers", &obj.Drivers, UnmarshalDriver)
	if err != nil {
		err = core.SDKErrorf(err, "", "drivers-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalEngineDetailsBody)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "engine_properties", &obj.EngineProperties, UnmarshalPrestoEngineEngineProperties)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_restart", &obj.EngineRestart)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_restart-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "external_host_name", &obj.ExternalHostName)
	if err != nil {
		err = core.SDKErrorf(err, "", "external_host_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "group_id", &obj.GroupID)
	if err != nil {
		err = core.SDKErrorf(err, "", "group_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host_name", &obj.HostName)
	if err != nil {
		err = core.SDKErrorf(err, "", "host_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		err = core.SDKErrorf(err, "", "origin-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		err = core.SDKErrorf(err, "", "region-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "remove_engine_properties", &obj.RemoveEngineProperties, UnmarshalPrestoEnginePatchRemoveEngineProperties)
	if err != nil {
		err = core.SDKErrorf(err, "", "remove_engine_properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "size_config", &obj.SizeConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "size_config-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "status_code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "worker", &obj.Worker, UnmarshalNodeDescription)
	if err != nil {
		err = core.SDKErrorf(err, "", "worker-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrestoEngineCollection : List Presto engines.
type PrestoEngineCollection struct {
	// Presto engine.
	PrestoEngines []PrestoEngine `json:"presto_engines,omitempty"`
}

// UnmarshalPrestoEngineCollection unmarshals an instance of PrestoEngineCollection from the specified map of raw messages.
func UnmarshalPrestoEngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestoEngineCollection)
	err = core.UnmarshalModel(m, "presto_engines", &obj.PrestoEngines, UnmarshalPrestoEngine)
	if err != nil {
		err = core.SDKErrorf(err, "", "presto_engines-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrestoEngineEngineProperties : Engine properties.
type PrestoEngineEngineProperties struct {
	// Catalog configuration settings.
	Catalog *PrestoEnginePropertiesCatalog `json:"catalog,omitempty"`

	// Configuration settings.
	Configuration *EnginePropertiesOaiGen1Configuration `json:"configuration,omitempty"`

	// Global session is to accomodate all the custom properties that can be applicable for both coordinator and worker.
	Global *PrestoEnginePropertiesGlobal `json:"global,omitempty"`

	// JVM settings.
	Jvm *EnginePropertiesOaiGen1Jvm `json:"jvm,omitempty"`
}

// UnmarshalPrestoEngineEngineProperties unmarshals an instance of PrestoEngineEngineProperties from the specified map of raw messages.
func UnmarshalPrestoEngineEngineProperties(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestoEngineEngineProperties)
	err = core.UnmarshalModel(m, "catalog", &obj.Catalog, UnmarshalPrestoEnginePropertiesCatalog)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "configuration", &obj.Configuration, UnmarshalEnginePropertiesOaiGen1Configuration)
	if err != nil {
		err = core.SDKErrorf(err, "", "configuration-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "global", &obj.Global, UnmarshalPrestoEnginePropertiesGlobal)
	if err != nil {
		err = core.SDKErrorf(err, "", "global-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "jvm", &obj.Jvm, UnmarshalEnginePropertiesOaiGen1Jvm)
	if err != nil {
		err = core.SDKErrorf(err, "", "jvm-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the PrestoEngineEngineProperties
func (prestoEngineEngineProperties *PrestoEngineEngineProperties) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(prestoEngineEngineProperties.Catalog) {
		_patch["catalog"] = prestoEngineEngineProperties.Catalog.asPatch()
	}
	if !core.IsNil(prestoEngineEngineProperties.Configuration) {
		_patch["configuration"] = prestoEngineEngineProperties.Configuration.asPatch()
	}
	if !core.IsNil(prestoEngineEngineProperties.Global) {
		_patch["global"] = prestoEngineEngineProperties.Global.asPatch()
	}
	if !core.IsNil(prestoEngineEngineProperties.Jvm) {
		_patch["jvm"] = prestoEngineEngineProperties.Jvm.asPatch()
	}

	return
}

// PrestoEnginePatch : UpdateEngine body.
type PrestoEnginePatch struct {
	// Modified description.
	Description *string `json:"description,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine properties.
	EngineProperties *PrestoEngineEngineProperties `json:"engine_properties,omitempty"`

	// Triggers engine restart if value is force.
	EngineRestart *string `json:"engine_restart,omitempty"`

	// RemoveEngine properties.
	RemoveEngineProperties *PrestoEnginePatchRemoveEngineProperties `json:"remove_engine_properties,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`
}

// Constants associated with the PrestoEnginePatch.EngineRestart property.
// Triggers engine restart if value is force.
const (
	PrestoEnginePatch_EngineRestart_False = "false"
	PrestoEnginePatch_EngineRestart_Force = "force"
)

// UnmarshalPrestoEnginePatch unmarshals an instance of PrestoEnginePatch from the specified map of raw messages.
func UnmarshalPrestoEnginePatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestoEnginePatch)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "engine_properties", &obj.EngineProperties, UnmarshalPrestoEngineEngineProperties)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_restart", &obj.EngineRestart)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_restart-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "remove_engine_properties", &obj.RemoveEngineProperties, UnmarshalPrestoEnginePatchRemoveEngineProperties)
	if err != nil {
		err = core.SDKErrorf(err, "", "remove_engine_properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the PrestoEnginePatch
func (prestoEnginePatch *PrestoEnginePatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(prestoEnginePatch.Description) {
		_patch["description"] = prestoEnginePatch.Description
	}
	if !core.IsNil(prestoEnginePatch.EngineDisplayName) {
		_patch["engine_display_name"] = prestoEnginePatch.EngineDisplayName
	}
	if !core.IsNil(prestoEnginePatch.EngineProperties) {
		_patch["engine_properties"] = prestoEnginePatch.EngineProperties.asPatch()
	}
	if !core.IsNil(prestoEnginePatch.EngineRestart) {
		_patch["engine_restart"] = prestoEnginePatch.EngineRestart
	}
	if !core.IsNil(prestoEnginePatch.RemoveEngineProperties) {
		_patch["remove_engine_properties"] = prestoEnginePatch.RemoveEngineProperties.asPatch()
	}
	if !core.IsNil(prestoEnginePatch.Tags) {
		_patch["tags"] = prestoEnginePatch.Tags
	}

	return
}

// PrestoEnginePatchRemoveEngineProperties : RemoveEngine properties.
type PrestoEnginePatchRemoveEngineProperties struct {
	// Configuration settings for removing engine properties.
	Configuration *RemoveEnginePropertiesOaiGenConfiguration `json:"configuration,omitempty"`

	// JVM properties.
	Jvm *RemoveEnginePropertiesOaiGenJvm `json:"jvm,omitempty"`

	// Catalog configuration settings.
	Catalog *PrestoEnginePropertiesCatalog `json:"catalog,omitempty"`
}

// UnmarshalPrestoEnginePatchRemoveEngineProperties unmarshals an instance of PrestoEnginePatchRemoveEngineProperties from the specified map of raw messages.
func UnmarshalPrestoEnginePatchRemoveEngineProperties(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestoEnginePatchRemoveEngineProperties)
	err = core.UnmarshalModel(m, "configuration", &obj.Configuration, UnmarshalRemoveEnginePropertiesOaiGenConfiguration)
	if err != nil {
		err = core.SDKErrorf(err, "", "configuration-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "jvm", &obj.Jvm, UnmarshalRemoveEnginePropertiesOaiGenJvm)
	if err != nil {
		err = core.SDKErrorf(err, "", "jvm-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "catalog", &obj.Catalog, UnmarshalPrestoEnginePropertiesCatalog)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the PrestoEnginePatchRemoveEngineProperties
func (prestoEnginePatchRemoveEngineProperties *PrestoEnginePatchRemoveEngineProperties) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(prestoEnginePatchRemoveEngineProperties.Configuration) {
		_patch["configuration"] = prestoEnginePatchRemoveEngineProperties.Configuration.asPatch()
	}
	if !core.IsNil(prestoEnginePatchRemoveEngineProperties.Jvm) {
		_patch["jvm"] = prestoEnginePatchRemoveEngineProperties.Jvm.asPatch()
	}
	if !core.IsNil(prestoEnginePatchRemoveEngineProperties.Catalog) {
		_patch["catalog"] = prestoEnginePatchRemoveEngineProperties.Catalog.asPatch()
	}

	return
}

// PrestoEnginePropertiesCatalog : Catalog configuration settings.
type PrestoEnginePropertiesCatalog struct {
	// Name of the catalog.
	CatalogName *string `json:"catalog_name,omitempty"`
}

// UnmarshalPrestoEnginePropertiesCatalog unmarshals an instance of PrestoEnginePropertiesCatalog from the specified map of raw messages.
func UnmarshalPrestoEnginePropertiesCatalog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestoEnginePropertiesCatalog)
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog_name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the PrestoEnginePropertiesCatalog
func (prestoEnginePropertiesCatalog *PrestoEnginePropertiesCatalog) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(prestoEnginePropertiesCatalog.CatalogName) {
		_patch["catalog_name"] = prestoEnginePropertiesCatalog.CatalogName
	}

	return
}

// PrestoEnginePropertiesGlobal : Global session is to accomodate all the custom properties that can be applicable for both coordinator and worker.
type PrestoEnginePropertiesGlobal struct {
	// Global property settings.
	GlobalProperty *string `json:"global_property,omitempty"`
}

// UnmarshalPrestoEnginePropertiesGlobal unmarshals an instance of PrestoEnginePropertiesGlobal from the specified map of raw messages.
func UnmarshalPrestoEnginePropertiesGlobal(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrestoEnginePropertiesGlobal)
	err = core.UnmarshalPrimitive(m, "global_property", &obj.GlobalProperty)
	if err != nil {
		err = core.SDKErrorf(err, "", "global_property-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the PrestoEnginePropertiesGlobal
func (prestoEnginePropertiesGlobal *PrestoEnginePropertiesGlobal) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(prestoEnginePropertiesGlobal.GlobalProperty) {
		_patch["global_property"] = prestoEnginePropertiesGlobal.GlobalProperty
	}

	return
}

// RemoveEngineProperties : RemoveEngine properties.
type RemoveEngineProperties struct {
	// catalog properties.
	Catalog *PrestissimoEnginePropertiesCatalog `json:"catalog,omitempty"`

	// remove engine properties configuration.
	Configuration *RemoveEnginePropertiesConfiguration `json:"configuration,omitempty"`

	// remove engine properties configuration.
	Jvm *RemoveEnginePropertiesConfiguration `json:"jvm,omitempty"`

	// velox description.
	Velox []string `json:"velox,omitempty"`
}

// UnmarshalRemoveEngineProperties unmarshals an instance of RemoveEngineProperties from the specified map of raw messages.
func UnmarshalRemoveEngineProperties(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoveEngineProperties)
	err = core.UnmarshalModel(m, "catalog", &obj.Catalog, UnmarshalPrestissimoEnginePropertiesCatalog)
	if err != nil {
		err = core.SDKErrorf(err, "", "catalog-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "configuration", &obj.Configuration, UnmarshalRemoveEnginePropertiesConfiguration)
	if err != nil {
		err = core.SDKErrorf(err, "", "configuration-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "jvm", &obj.Jvm, UnmarshalRemoveEnginePropertiesConfiguration)
	if err != nil {
		err = core.SDKErrorf(err, "", "jvm-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "velox", &obj.Velox)
	if err != nil {
		err = core.SDKErrorf(err, "", "velox-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the RemoveEngineProperties
func (removeEngineProperties *RemoveEngineProperties) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(removeEngineProperties.Catalog) {
		_patch["catalog"] = removeEngineProperties.Catalog.asPatch()
	}
	if !core.IsNil(removeEngineProperties.Configuration) {
		_patch["configuration"] = removeEngineProperties.Configuration.asPatch()
	}
	if !core.IsNil(removeEngineProperties.Jvm) {
		_patch["jvm"] = removeEngineProperties.Jvm.asPatch()
	}
	if !core.IsNil(removeEngineProperties.Velox) {
		_patch["velox"] = removeEngineProperties.Velox
	}

	return
}

// RemoveEnginePropertiesConfiguration : remove engine properties configuration.
type RemoveEnginePropertiesConfiguration struct {
	// description for coordinator property.
	Coordinator []string `json:"coordinator,omitempty"`

	// description for worker property.
	Worker []string `json:"worker,omitempty"`
}

// UnmarshalRemoveEnginePropertiesConfiguration unmarshals an instance of RemoveEnginePropertiesConfiguration from the specified map of raw messages.
func UnmarshalRemoveEnginePropertiesConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoveEnginePropertiesConfiguration)
	err = core.UnmarshalPrimitive(m, "coordinator", &obj.Coordinator)
	if err != nil {
		err = core.SDKErrorf(err, "", "coordinator-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "worker", &obj.Worker)
	if err != nil {
		err = core.SDKErrorf(err, "", "worker-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the RemoveEnginePropertiesConfiguration
func (removeEnginePropertiesConfiguration *RemoveEnginePropertiesConfiguration) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(removeEnginePropertiesConfiguration.Coordinator) {
		_patch["coordinator"] = removeEnginePropertiesConfiguration.Coordinator
	}
	if !core.IsNil(removeEnginePropertiesConfiguration.Worker) {
		_patch["worker"] = removeEnginePropertiesConfiguration.Worker
	}

	return
}

// RemoveEnginePropertiesOaiGenConfiguration : Configuration settings for removing engine properties.
type RemoveEnginePropertiesOaiGenConfiguration struct {
	// List of coordinator properties.
	Coordinator []string `json:"coordinator,omitempty"`

	// List of worker properties.
	Worker []string `json:"worker,omitempty"`
}

// UnmarshalRemoveEnginePropertiesOaiGenConfiguration unmarshals an instance of RemoveEnginePropertiesOaiGenConfiguration from the specified map of raw messages.
func UnmarshalRemoveEnginePropertiesOaiGenConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoveEnginePropertiesOaiGenConfiguration)
	err = core.UnmarshalPrimitive(m, "coordinator", &obj.Coordinator)
	if err != nil {
		err = core.SDKErrorf(err, "", "coordinator-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "worker", &obj.Worker)
	if err != nil {
		err = core.SDKErrorf(err, "", "worker-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the RemoveEnginePropertiesOaiGenConfiguration
func (removeEnginePropertiesOaiGenConfiguration *RemoveEnginePropertiesOaiGenConfiguration) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(removeEnginePropertiesOaiGenConfiguration.Coordinator) {
		_patch["coordinator"] = removeEnginePropertiesOaiGenConfiguration.Coordinator
	}
	if !core.IsNil(removeEnginePropertiesOaiGenConfiguration.Worker) {
		_patch["worker"] = removeEnginePropertiesOaiGenConfiguration.Worker
	}

	return
}

// RemoveEnginePropertiesOaiGenJvm : JVM properties.
type RemoveEnginePropertiesOaiGenJvm struct {
	// List of coordinator properties.
	Coordinator []string `json:"coordinator,omitempty"`

	// List of worker properties.
	Worker []string `json:"worker,omitempty"`
}

// UnmarshalRemoveEnginePropertiesOaiGenJvm unmarshals an instance of RemoveEnginePropertiesOaiGenJvm from the specified map of raw messages.
func UnmarshalRemoveEnginePropertiesOaiGenJvm(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoveEnginePropertiesOaiGenJvm)
	err = core.UnmarshalPrimitive(m, "coordinator", &obj.Coordinator)
	if err != nil {
		err = core.SDKErrorf(err, "", "coordinator-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "worker", &obj.Worker)
	if err != nil {
		err = core.SDKErrorf(err, "", "worker-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the RemoveEnginePropertiesOaiGenJvm
func (removeEnginePropertiesOaiGenJvm *RemoveEnginePropertiesOaiGenJvm) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(removeEnginePropertiesOaiGenJvm.Coordinator) {
		_patch["coordinator"] = removeEnginePropertiesOaiGenJvm.Coordinator
	}
	if !core.IsNil(removeEnginePropertiesOaiGenJvm.Worker) {
		_patch["worker"] = removeEnginePropertiesOaiGenJvm.Worker
	}

	return
}

// RenameTableOptions : The RenameTable options.
type RenameTableOptions struct {
	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded table name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// Request body.
	Body map[string]interface{} `json:"body" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewRenameTableOptions : Instantiate RenameTableOptions
func (*WatsonxDataV2) NewRenameTableOptions(catalogID string, schemaID string, tableID string, engineID string, body map[string]interface{}) *RenameTableOptions {
	return &RenameTableOptions{
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *RenameTableOptions) SetCatalogID(catalogID string) *RenameTableOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *RenameTableOptions) SetSchemaID(schemaID string) *RenameTableOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *RenameTableOptions) SetTableID(tableID string) *RenameTableOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetEngineID : Allow user to set EngineID
func (_options *RenameTableOptions) SetEngineID(engineID string) *RenameTableOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *RenameTableOptions) SetBody(body map[string]interface{}) *RenameTableOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RenameTableOptions) SetAuthInstanceID(authInstanceID string) *RenameTableOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RenameTableOptions) SetHeaders(param map[string]string) *RenameTableOptions {
	options.Headers = param
	return options
}

// ReplaceSnapshotCreatedBody : success response.
type ReplaceSnapshotCreatedBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalReplaceSnapshotCreatedBody unmarshals an instance of ReplaceSnapshotCreatedBody from the specified map of raw messages.
func UnmarshalReplaceSnapshotCreatedBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReplaceSnapshotCreatedBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RestartPrestissimoEngineOptions : The RestartPrestissimoEngine options.
type RestartPrestissimoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewRestartPrestissimoEngineOptions : Instantiate RestartPrestissimoEngineOptions
func (*WatsonxDataV2) NewRestartPrestissimoEngineOptions(engineID string) *RestartPrestissimoEngineOptions {
	return &RestartPrestissimoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RestartPrestissimoEngineOptions) SetEngineID(engineID string) *RestartPrestissimoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RestartPrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *RestartPrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RestartPrestissimoEngineOptions) SetHeaders(param map[string]string) *RestartPrestissimoEngineOptions {
	options.Headers = param
	return options
}

// RestartPrestoEngineOptions : The RestartPrestoEngine options.
type RestartPrestoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewRestartPrestoEngineOptions : Instantiate RestartPrestoEngineOptions
func (*WatsonxDataV2) NewRestartPrestoEngineOptions(engineID string) *RestartPrestoEngineOptions {
	return &RestartPrestoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RestartPrestoEngineOptions) SetEngineID(engineID string) *RestartPrestoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RestartPrestoEngineOptions) SetAuthInstanceID(authInstanceID string) *RestartPrestoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RestartPrestoEngineOptions) SetHeaders(param map[string]string) *RestartPrestoEngineOptions {
	options.Headers = param
	return options
}

// ResultPrestissimoExplainStatement : ExplainStatement OK.
type ResultPrestissimoExplainStatement struct {
	// Result.
	Result *string `json:"result,omitempty"`
}

// UnmarshalResultPrestissimoExplainStatement unmarshals an instance of ResultPrestissimoExplainStatement from the specified map of raw messages.
func UnmarshalResultPrestissimoExplainStatement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResultPrestissimoExplainStatement)
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResultRunPrestissimoExplainAnalyzeStatement : explainAnalyzeStatement OK.
type ResultRunPrestissimoExplainAnalyzeStatement struct {
	// explainAnalyzeStatement result.
	Result *string `json:"result,omitempty"`
}

// UnmarshalResultRunPrestissimoExplainAnalyzeStatement unmarshals an instance of ResultRunPrestissimoExplainAnalyzeStatement from the specified map of raw messages.
func UnmarshalResultRunPrestissimoExplainAnalyzeStatement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResultRunPrestissimoExplainAnalyzeStatement)
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResumePrestissimoEngineOptions : The ResumePrestissimoEngine options.
type ResumePrestissimoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewResumePrestissimoEngineOptions : Instantiate ResumePrestissimoEngineOptions
func (*WatsonxDataV2) NewResumePrestissimoEngineOptions(engineID string) *ResumePrestissimoEngineOptions {
	return &ResumePrestissimoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ResumePrestissimoEngineOptions) SetEngineID(engineID string) *ResumePrestissimoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ResumePrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *ResumePrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ResumePrestissimoEngineOptions) SetHeaders(param map[string]string) *ResumePrestissimoEngineOptions {
	options.Headers = param
	return options
}

// ResumePrestoEngineOptions : The ResumePrestoEngine options.
type ResumePrestoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewResumePrestoEngineOptions : Instantiate ResumePrestoEngineOptions
func (*WatsonxDataV2) NewResumePrestoEngineOptions(engineID string) *ResumePrestoEngineOptions {
	return &ResumePrestoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ResumePrestoEngineOptions) SetEngineID(engineID string) *ResumePrestoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ResumePrestoEngineOptions) SetAuthInstanceID(authInstanceID string) *ResumePrestoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ResumePrestoEngineOptions) SetHeaders(param map[string]string) *ResumePrestoEngineOptions {
	options.Headers = param
	return options
}

// RollbackTableOptions : The RollbackTable options.
type RollbackTableOptions struct {
	// Engine name.
	EngineID *string `json:"engine_id" validate:"required"`

	// Catalog ID.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Schema ID.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// Table ID.
	TableID *string `json:"table_id" validate:"required,ne="`

	// Snapshot Id.
	SnapshotID *string `json:"snapshot_id,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewRollbackTableOptions : Instantiate RollbackTableOptions
func (*WatsonxDataV2) NewRollbackTableOptions(engineID string, catalogID string, schemaID string, tableID string) *RollbackTableOptions {
	return &RollbackTableOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RollbackTableOptions) SetEngineID(engineID string) *RollbackTableOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *RollbackTableOptions) SetCatalogID(catalogID string) *RollbackTableOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *RollbackTableOptions) SetSchemaID(schemaID string) *RollbackTableOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *RollbackTableOptions) SetTableID(tableID string) *RollbackTableOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetSnapshotID : Allow user to set SnapshotID
func (_options *RollbackTableOptions) SetSnapshotID(snapshotID string) *RollbackTableOptions {
	_options.SnapshotID = core.StringPtr(snapshotID)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RollbackTableOptions) SetAuthInstanceID(authInstanceID string) *RollbackTableOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RollbackTableOptions) SetHeaders(param map[string]string) *RollbackTableOptions {
	options.Headers = param
	return options
}

// RunExplainAnalyzeStatementOKBody : explainAnalyzeStatement OK.
type RunExplainAnalyzeStatementOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// explainAnalyzeStatement result.
	Result *string `json:"result" validate:"required"`
}

// UnmarshalRunExplainAnalyzeStatementOKBody unmarshals an instance of RunExplainAnalyzeStatementOKBody from the specified map of raw messages.
func UnmarshalRunExplainAnalyzeStatementOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RunExplainAnalyzeStatementOKBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RunExplainAnalyzeStatementOptions : The RunExplainAnalyzeStatement options.
type RunExplainAnalyzeStatementOptions struct {
	// Engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Presto query to show explain analyze.
	Statement *string `json:"statement" validate:"required"`

	// Verbose.
	Verbose *bool `json:"verbose,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewRunExplainAnalyzeStatementOptions : Instantiate RunExplainAnalyzeStatementOptions
func (*WatsonxDataV2) NewRunExplainAnalyzeStatementOptions(engineID string, statement string) *RunExplainAnalyzeStatementOptions {
	return &RunExplainAnalyzeStatementOptions{
		EngineID: core.StringPtr(engineID),
		Statement: core.StringPtr(statement),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RunExplainAnalyzeStatementOptions) SetEngineID(engineID string) *RunExplainAnalyzeStatementOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetStatement : Allow user to set Statement
func (_options *RunExplainAnalyzeStatementOptions) SetStatement(statement string) *RunExplainAnalyzeStatementOptions {
	_options.Statement = core.StringPtr(statement)
	return _options
}

// SetVerbose : Allow user to set Verbose
func (_options *RunExplainAnalyzeStatementOptions) SetVerbose(verbose bool) *RunExplainAnalyzeStatementOptions {
	_options.Verbose = core.BoolPtr(verbose)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RunExplainAnalyzeStatementOptions) SetAuthInstanceID(authInstanceID string) *RunExplainAnalyzeStatementOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunExplainAnalyzeStatementOptions) SetHeaders(param map[string]string) *RunExplainAnalyzeStatementOptions {
	options.Headers = param
	return options
}

// RunExplainStatementOKBody : ExplainStatement OK.
type RunExplainStatementOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response" validate:"required"`

	// Result.
	Result *string `json:"result" validate:"required"`
}

// UnmarshalRunExplainStatementOKBody unmarshals an instance of RunExplainStatementOKBody from the specified map of raw messages.
func UnmarshalRunExplainStatementOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RunExplainStatementOKBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
	if err != nil {
		err = core.SDKErrorf(err, "", "result-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RunExplainStatementOptions : The RunExplainStatement options.
type RunExplainStatementOptions struct {
	// Engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Presto query to determine explain plan.
	Statement *string `json:"statement" validate:"required"`

	// Format.
	Format *string `json:"format,omitempty"`

	// Type.
	Type *string `json:"type,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the RunExplainStatementOptions.Format property.
// Format.
const (
	RunExplainStatementOptions_Format_Graphviz = "graphviz"
	RunExplainStatementOptions_Format_JSON = "json"
	RunExplainStatementOptions_Format_Text = "text"
)

// Constants associated with the RunExplainStatementOptions.Type property.
// Type.
const (
	RunExplainStatementOptions_Type_Distributed = "distributed"
	RunExplainStatementOptions_Type_Io = "io"
	RunExplainStatementOptions_Type_Logical = "logical"
	RunExplainStatementOptions_Type_Validate = "validate"
)

// NewRunExplainStatementOptions : Instantiate RunExplainStatementOptions
func (*WatsonxDataV2) NewRunExplainStatementOptions(engineID string, statement string) *RunExplainStatementOptions {
	return &RunExplainStatementOptions{
		EngineID: core.StringPtr(engineID),
		Statement: core.StringPtr(statement),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RunExplainStatementOptions) SetEngineID(engineID string) *RunExplainStatementOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetStatement : Allow user to set Statement
func (_options *RunExplainStatementOptions) SetStatement(statement string) *RunExplainStatementOptions {
	_options.Statement = core.StringPtr(statement)
	return _options
}

// SetFormat : Allow user to set Format
func (_options *RunExplainStatementOptions) SetFormat(format string) *RunExplainStatementOptions {
	_options.Format = core.StringPtr(format)
	return _options
}

// SetType : Allow user to set Type
func (_options *RunExplainStatementOptions) SetType(typeVar string) *RunExplainStatementOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RunExplainStatementOptions) SetAuthInstanceID(authInstanceID string) *RunExplainStatementOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunExplainStatementOptions) SetHeaders(param map[string]string) *RunExplainStatementOptions {
	options.Headers = param
	return options
}

// RunPrestissimoExplainAnalyzeStatementOptions : The RunPrestissimoExplainAnalyzeStatement options.
type RunPrestissimoExplainAnalyzeStatementOptions struct {
	// Engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Presto query to show explain analyze.
	Statement *string `json:"statement" validate:"required"`

	// Verbose.
	Verbose *bool `json:"verbose,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewRunPrestissimoExplainAnalyzeStatementOptions : Instantiate RunPrestissimoExplainAnalyzeStatementOptions
func (*WatsonxDataV2) NewRunPrestissimoExplainAnalyzeStatementOptions(engineID string, statement string) *RunPrestissimoExplainAnalyzeStatementOptions {
	return &RunPrestissimoExplainAnalyzeStatementOptions{
		EngineID: core.StringPtr(engineID),
		Statement: core.StringPtr(statement),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RunPrestissimoExplainAnalyzeStatementOptions) SetEngineID(engineID string) *RunPrestissimoExplainAnalyzeStatementOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetStatement : Allow user to set Statement
func (_options *RunPrestissimoExplainAnalyzeStatementOptions) SetStatement(statement string) *RunPrestissimoExplainAnalyzeStatementOptions {
	_options.Statement = core.StringPtr(statement)
	return _options
}

// SetVerbose : Allow user to set Verbose
func (_options *RunPrestissimoExplainAnalyzeStatementOptions) SetVerbose(verbose bool) *RunPrestissimoExplainAnalyzeStatementOptions {
	_options.Verbose = core.BoolPtr(verbose)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RunPrestissimoExplainAnalyzeStatementOptions) SetAuthInstanceID(authInstanceID string) *RunPrestissimoExplainAnalyzeStatementOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunPrestissimoExplainAnalyzeStatementOptions) SetHeaders(param map[string]string) *RunPrestissimoExplainAnalyzeStatementOptions {
	options.Headers = param
	return options
}

// RunPrestissimoExplainStatementOptions : The RunPrestissimoExplainStatement options.
type RunPrestissimoExplainStatementOptions struct {
	// Engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Presto query to determine explain plan.
	Statement *string `json:"statement" validate:"required"`

	// Format.
	Format *string `json:"format,omitempty"`

	// Type.
	Type *string `json:"type,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the RunPrestissimoExplainStatementOptions.Format property.
// Format.
const (
	RunPrestissimoExplainStatementOptions_Format_Graphviz = "graphviz"
	RunPrestissimoExplainStatementOptions_Format_JSON = "json"
	RunPrestissimoExplainStatementOptions_Format_Text = "text"
)

// Constants associated with the RunPrestissimoExplainStatementOptions.Type property.
// Type.
const (
	RunPrestissimoExplainStatementOptions_Type_Distributed = "distributed"
	RunPrestissimoExplainStatementOptions_Type_Io = "io"
	RunPrestissimoExplainStatementOptions_Type_Logical = "logical"
	RunPrestissimoExplainStatementOptions_Type_Validate = "validate"
)

// NewRunPrestissimoExplainStatementOptions : Instantiate RunPrestissimoExplainStatementOptions
func (*WatsonxDataV2) NewRunPrestissimoExplainStatementOptions(engineID string, statement string) *RunPrestissimoExplainStatementOptions {
	return &RunPrestissimoExplainStatementOptions{
		EngineID: core.StringPtr(engineID),
		Statement: core.StringPtr(statement),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *RunPrestissimoExplainStatementOptions) SetEngineID(engineID string) *RunPrestissimoExplainStatementOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetStatement : Allow user to set Statement
func (_options *RunPrestissimoExplainStatementOptions) SetStatement(statement string) *RunPrestissimoExplainStatementOptions {
	_options.Statement = core.StringPtr(statement)
	return _options
}

// SetFormat : Allow user to set Format
func (_options *RunPrestissimoExplainStatementOptions) SetFormat(format string) *RunPrestissimoExplainStatementOptions {
	_options.Format = core.StringPtr(format)
	return _options
}

// SetType : Allow user to set Type
func (_options *RunPrestissimoExplainStatementOptions) SetType(typeVar string) *RunPrestissimoExplainStatementOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *RunPrestissimoExplainStatementOptions) SetAuthInstanceID(authInstanceID string) *RunPrestissimoExplainStatementOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunPrestissimoExplainStatementOptions) SetHeaders(param map[string]string) *RunPrestissimoExplainStatementOptions {
	options.Headers = param
	return options
}

// ScalePrestissimoEngineOptions : The ScalePrestissimoEngine options.
type ScalePrestissimoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Node details.
	Coordinator *PrestissimoNodeDescriptionBody `json:"coordinator,omitempty"`

	// Node details.
	Worker *PrestissimoNodeDescriptionBody `json:"worker,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewScalePrestissimoEngineOptions : Instantiate ScalePrestissimoEngineOptions
func (*WatsonxDataV2) NewScalePrestissimoEngineOptions(engineID string) *ScalePrestissimoEngineOptions {
	return &ScalePrestissimoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ScalePrestissimoEngineOptions) SetEngineID(engineID string) *ScalePrestissimoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCoordinator : Allow user to set Coordinator
func (_options *ScalePrestissimoEngineOptions) SetCoordinator(coordinator *PrestissimoNodeDescriptionBody) *ScalePrestissimoEngineOptions {
	_options.Coordinator = coordinator
	return _options
}

// SetWorker : Allow user to set Worker
func (_options *ScalePrestissimoEngineOptions) SetWorker(worker *PrestissimoNodeDescriptionBody) *ScalePrestissimoEngineOptions {
	_options.Worker = worker
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ScalePrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *ScalePrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ScalePrestissimoEngineOptions) SetHeaders(param map[string]string) *ScalePrestissimoEngineOptions {
	options.Headers = param
	return options
}

// ScalePrestoEngineOptions : The ScalePrestoEngine options.
type ScalePrestoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// NodeDescription.
	Coordinator *NodeDescription `json:"coordinator,omitempty"`

	// NodeDescription.
	Worker *NodeDescription `json:"worker,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewScalePrestoEngineOptions : Instantiate ScalePrestoEngineOptions
func (*WatsonxDataV2) NewScalePrestoEngineOptions(engineID string) *ScalePrestoEngineOptions {
	return &ScalePrestoEngineOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *ScalePrestoEngineOptions) SetEngineID(engineID string) *ScalePrestoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCoordinator : Allow user to set Coordinator
func (_options *ScalePrestoEngineOptions) SetCoordinator(coordinator *NodeDescription) *ScalePrestoEngineOptions {
	_options.Coordinator = coordinator
	return _options
}

// SetWorker : Allow user to set Worker
func (_options *ScalePrestoEngineOptions) SetWorker(worker *NodeDescription) *ScalePrestoEngineOptions {
	_options.Worker = worker
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *ScalePrestoEngineOptions) SetAuthInstanceID(authInstanceID string) *ScalePrestoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ScalePrestoEngineOptions) SetHeaders(param map[string]string) *ScalePrestoEngineOptions {
	options.Headers = param
	return options
}

// SparkApplicationConfig : Spark applications details configuration.
type SparkApplicationConfig struct {
	// spark_sample_config_properpty.
	SparkSampleConfigProperpty *string `json:"spark_sample_config_properpty,omitempty"`
}

// UnmarshalSparkApplicationConfig unmarshals an instance of SparkApplicationConfig from the specified map of raw messages.
func UnmarshalSparkApplicationConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkApplicationConfig)
	err = core.UnmarshalPrimitive(m, "spark_sample_config_properpty", &obj.SparkSampleConfigProperpty)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_sample_config_properpty-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkApplicationDetails : Application details.
type SparkApplicationDetails struct {
	// Application.
	Application *string `json:"application,omitempty"`

	// List of arguments.
	Arguments []string `json:"arguments,omitempty"`

	// Class.
	Class *string `json:"class,omitempty"`

	// Spark applications details configuration.
	Conf *SparkApplicationConfig `json:"conf,omitempty"`

	// Spark applications details env samples.
	Env *SparkApplicationEnv `json:"env,omitempty"`

	// Files.
	Files *string `json:"files,omitempty"`

	// Jars.
	Jars *string `json:"jars,omitempty"`

	// Display name of the spark application.
	Name *string `json:"name,omitempty"`

	// Packages.
	Packages *string `json:"packages,omitempty"`

	// Repositories.
	Repositories *string `json:"repositories,omitempty"`

	// Spark Version.
	SparkVersion *string `json:"spark_version,omitempty"`
}

// UnmarshalSparkApplicationDetails unmarshals an instance of SparkApplicationDetails from the specified map of raw messages.
func UnmarshalSparkApplicationDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkApplicationDetails)
	err = core.UnmarshalPrimitive(m, "application", &obj.Application)
	if err != nil {
		err = core.SDKErrorf(err, "", "application-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "arguments", &obj.Arguments)
	if err != nil {
		err = core.SDKErrorf(err, "", "arguments-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "class", &obj.Class)
	if err != nil {
		err = core.SDKErrorf(err, "", "class-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "conf", &obj.Conf, UnmarshalSparkApplicationConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "conf-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "env", &obj.Env, UnmarshalSparkApplicationEnv)
	if err != nil {
		err = core.SDKErrorf(err, "", "env-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "files", &obj.Files)
	if err != nil {
		err = core.SDKErrorf(err, "", "files-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "jars", &obj.Jars)
	if err != nil {
		err = core.SDKErrorf(err, "", "jars-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "packages", &obj.Packages)
	if err != nil {
		err = core.SDKErrorf(err, "", "packages-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "repositories", &obj.Repositories)
	if err != nil {
		err = core.SDKErrorf(err, "", "repositories-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_version", &obj.SparkVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_version-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkApplicationEnv : Spark applications details env samples.
type SparkApplicationEnv struct {
	// sample.
	SampleEnvKey *string `json:"sample_env_key,omitempty"`
}

// UnmarshalSparkApplicationEnv unmarshals an instance of SparkApplicationEnv from the specified map of raw messages.
func UnmarshalSparkApplicationEnv(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkApplicationEnv)
	err = core.UnmarshalPrimitive(m, "sample_env_key", &obj.SampleEnvKey)
	if err != nil {
		err = core.SDKErrorf(err, "", "sample_env_key-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkDefaultConfig : Spark Default Config details.
type SparkDefaultConfig struct {
	// config1.
	Config1 *string `json:"config1,omitempty"`

	// config2.
	Config2 *string `json:"config2,omitempty"`
}

// UnmarshalSparkDefaultConfig unmarshals an instance of SparkDefaultConfig from the specified map of raw messages.
func UnmarshalSparkDefaultConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkDefaultConfig)
	err = core.UnmarshalPrimitive(m, "config1", &obj.Config1)
	if err != nil {
		err = core.SDKErrorf(err, "", "config1-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "config2", &obj.Config2)
	if err != nil {
		err = core.SDKErrorf(err, "", "config2-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEndpoints : Application Endpoints.
type SparkEndpoints struct {
	// Application API.
	ApplicationsApi *string `json:"applications_api,omitempty"`

	// History server endpoint.
	HistoryServerEndpoint *string `json:"history_server_endpoint,omitempty"`

	// Spark kernel endpoint.
	SparkAccessEndpoint *string `json:"spark_access_endpoint,omitempty"`

	// Spark jobs V4 endpoint.
	SparkJobsV4Endpoint *string `json:"spark_jobs_v4_endpoint,omitempty"`

	// Spark kernel endpoint.
	SparkKernelEndpoint *string `json:"spark_kernel_endpoint,omitempty"`

	// View history server.
	ViewHistoryServer *string `json:"view_history_server,omitempty"`

	// Wxd application endpoint.
	WxdApplicationEndpoint *string `json:"wxd_application_endpoint,omitempty"`

	// Wxd engine endpoint.
	WxdEngineEndpoint *string `json:"wxd_engine_endpoint,omitempty"`

	// Wxd history_server endpoint.
	WxdHistoryServerEndpoint *string `json:"wxd_history_server_endpoint,omitempty"`

	// Wxd history_server endpoint.
	WxdHistoryServerUiEndpoint *string `json:"wxd_history_server_ui_endpoint,omitempty"`
}

// UnmarshalSparkEndpoints unmarshals an instance of SparkEndpoints from the specified map of raw messages.
func UnmarshalSparkEndpoints(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEndpoints)
	err = core.UnmarshalPrimitive(m, "applications_api", &obj.ApplicationsApi)
	if err != nil {
		err = core.SDKErrorf(err, "", "applications_api-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "history_server_endpoint", &obj.HistoryServerEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "history_server_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_access_endpoint", &obj.SparkAccessEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_access_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_jobs_v4_endpoint", &obj.SparkJobsV4Endpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_jobs_v4_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_kernel_endpoint", &obj.SparkKernelEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_kernel_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "view_history_server", &obj.ViewHistoryServer)
	if err != nil {
		err = core.SDKErrorf(err, "", "view_history_server-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "wxd_application_endpoint", &obj.WxdApplicationEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "wxd_application_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "wxd_engine_endpoint", &obj.WxdEngineEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "wxd_engine_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "wxd_history_server_endpoint", &obj.WxdHistoryServerEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "wxd_history_server_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "wxd_history_server_ui_endpoint", &obj.WxdHistoryServerUiEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "wxd_history_server_ui_endpoint-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngine : EngineDetail.
type SparkEngine struct {
	// Actions.
	Actions []string `json:"actions,omitempty"`

	// Associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`

	// watsonx.data build version.
	BuildVersion *string `json:"build_version,omitempty"`

	// Created user name.
	CreatedBy *string `json:"created_by,omitempty"`

	// Created time in epoch format.
	CreatedOn *int64 `json:"created_on,omitempty"`

	// Engine description.
	Description *string `json:"description,omitempty"`

	// External engine details.
	EngineDetails *SparkEngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Engine programmatic name.
	EngineID *string `json:"engine_id,omitempty"`

	// Origin - created or registered.
	Origin *string `json:"origin,omitempty"`

	// Engine status.
	Status *string `json:"status,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`

	// Type like spark, netezza,..
	Type *string `json:"type,omitempty"`
}

// Constants associated with the SparkEngine.Origin property.
// Origin - created or registered.
const (
	SparkEngine_Origin_Discover = "discover"
	SparkEngine_Origin_External = "external"
	SparkEngine_Origin_Native = "native"
)

// Constants associated with the SparkEngine.Type property.
// Type like spark, netezza,..
const (
	SparkEngine_Type_Spark = "spark"
)

// UnmarshalSparkEngine unmarshals an instance of SparkEngine from the specified map of raw messages.
func UnmarshalSparkEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngine)
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		err = core.SDKErrorf(err, "", "actions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_catalogs", &obj.AssociatedCatalogs)
	if err != nil {
		err = core.SDKErrorf(err, "", "associated_catalogs-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "build_version", &obj.BuildVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "build_version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_on-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalSparkEngineDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "origin", &obj.Origin)
	if err != nil {
		err = core.SDKErrorf(err, "", "origin-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineApplicationStatus : Engine Application Status.
type SparkEngineApplicationStatus struct {
	// Application details.
	ApplicationDetails *SparkApplicationDetails `json:"application_details,omitempty"`

	// Application ID.
	ApplicationID *string `json:"application_id,omitempty"`

	// Auto Termination Time.
	AutoTerminationTime *string `json:"auto_termination_time,omitempty"`

	// Creation time.
	CreationTime *string `json:"creation_time,omitempty"`

	// Deployment mode.
	DeployMode *string `json:"deploy_mode,omitempty"`

	// End Time.
	EndTime *string `json:"end_time,omitempty"`

	// Failed time.
	FailedTime *string `json:"failed_time,omitempty"`

	// Finish time.
	FinishTime *string `json:"finish_time,omitempty"`

	// Application ID.
	ID *string `json:"id,omitempty"`

	// Job endpoint.
	JobEndpoint *string `json:"job_endpoint,omitempty"`

	// Return code.
	ReturnCode *string `json:"return_code,omitempty"`

	// application run time.
	Runtime *SparkEngineApplicationStatusRuntime `json:"runtime,omitempty"`

	// Service Instance ID for POST.
	ServiceInstanceID *string `json:"service_instance_id,omitempty"`

	// Spark application ID.
	SparkApplicationID *string `json:"spark_application_id,omitempty"`

	// Spark application name.
	SparkApplicationName *string `json:"spark_application_name,omitempty"`

	// Spark Version.
	SparkVersion *string `json:"spark_version,omitempty"`

	// Start time.
	StartTime *string `json:"start_time,omitempty"`

	// Application state.
	State *string `json:"state,omitempty"`

	// Application state details.
	StateDetails []SparkEngineApplicationStatusStateDetailsItems `json:"state_details,omitempty"`

	// Application submission time.
	SubmissionTime *string `json:"submission_time,omitempty"`

	// Template ID.
	TemplateID *string `json:"template_id,omitempty"`

	// Engine Type.
	Type *string `json:"type,omitempty"`

	// Spark application volumes to mount.
	Volumes []SparkVolumeDetails `json:"volumes,omitempty"`

	// Wxd history_server endpoint.
	WxdApplicationUiEndpoint *string `json:"wxd_application_ui_endpoint,omitempty"`
}

// Constants associated with the SparkEngineApplicationStatus.Type property.
// Engine Type.
const (
	SparkEngineApplicationStatus_Type_Emr = "emr"
	SparkEngineApplicationStatus_Type_Iae = "iae"
)

// UnmarshalSparkEngineApplicationStatus unmarshals an instance of SparkEngineApplicationStatus from the specified map of raw messages.
func UnmarshalSparkEngineApplicationStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineApplicationStatus)
	err = core.UnmarshalModel(m, "application_details", &obj.ApplicationDetails, UnmarshalSparkApplicationDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "application_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
	if err != nil {
		err = core.SDKErrorf(err, "", "application_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_termination_time", &obj.AutoTerminationTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_termination_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "creation_time", &obj.CreationTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "creation_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "deploy_mode", &obj.DeployMode)
	if err != nil {
		err = core.SDKErrorf(err, "", "deploy_mode-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "end_time", &obj.EndTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "end_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "failed_time", &obj.FailedTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "failed_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "finish_time", &obj.FinishTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "finish_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "job_endpoint", &obj.JobEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "job_endpoint-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "return_code", &obj.ReturnCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "return_code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "runtime", &obj.Runtime, UnmarshalSparkEngineApplicationStatusRuntime)
	if err != nil {
		err = core.SDKErrorf(err, "", "runtime-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "service_instance_id", &obj.ServiceInstanceID)
	if err != nil {
		err = core.SDKErrorf(err, "", "service_instance_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_application_id", &obj.SparkApplicationID)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_application_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_application_name", &obj.SparkApplicationName)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_application_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "spark_version", &obj.SparkVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "start_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "state_details", &obj.StateDetails, UnmarshalSparkEngineApplicationStatusStateDetailsItems)
	if err != nil {
		err = core.SDKErrorf(err, "", "state_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "submission_time", &obj.SubmissionTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "submission_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "template_id", &obj.TemplateID)
	if err != nil {
		err = core.SDKErrorf(err, "", "template_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "volumes", &obj.Volumes, UnmarshalSparkVolumeDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "volumes-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "wxd_application_ui_endpoint", &obj.WxdApplicationUiEndpoint)
	if err != nil {
		err = core.SDKErrorf(err, "", "wxd_application_ui_endpoint-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineApplicationStatusCollection : Engine Application Detail.
type SparkEngineApplicationStatusCollection struct {
	// Application body.
	Applications []SparkEngineApplicationStatus `json:"applications,omitempty"`
}

// UnmarshalSparkEngineApplicationStatusCollection unmarshals an instance of SparkEngineApplicationStatusCollection from the specified map of raw messages.
func UnmarshalSparkEngineApplicationStatusCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineApplicationStatusCollection)
	err = core.UnmarshalModel(m, "applications", &obj.Applications, UnmarshalSparkEngineApplicationStatus)
	if err != nil {
		err = core.SDKErrorf(err, "", "applications-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineApplicationStatusRuntime : application run time.
type SparkEngineApplicationStatusRuntime struct {
	// Spark Version.
	SparkVersion *string `json:"spark_version,omitempty"`
}

// UnmarshalSparkEngineApplicationStatusRuntime unmarshals an instance of SparkEngineApplicationStatusRuntime from the specified map of raw messages.
func UnmarshalSparkEngineApplicationStatusRuntime(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineApplicationStatusRuntime)
	err = core.UnmarshalPrimitive(m, "spark_version", &obj.SparkVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_version-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineApplicationStatusStateDetailsItems : State details.
type SparkEngineApplicationStatusStateDetailsItems struct {
	// State details code.
	Code *string `json:"code,omitempty"`

	// State details message.
	Message *string `json:"message,omitempty"`

	// State details type.
	Type *string `json:"type,omitempty"`
}

// UnmarshalSparkEngineApplicationStatusStateDetailsItems unmarshals an instance of SparkEngineApplicationStatusStateDetailsItems from the specified map of raw messages.
func UnmarshalSparkEngineApplicationStatusStateDetailsItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineApplicationStatusStateDetailsItems)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		err = core.SDKErrorf(err, "", "code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineCollection : List spark engines.
type SparkEngineCollection struct {
	// List spark engines.
	SparkEngines []SparkEngine `json:"spark_engines,omitempty"`
}

// UnmarshalSparkEngineCollection unmarshals an instance of SparkEngineCollection from the specified map of raw messages.
func UnmarshalSparkEngineCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineCollection)
	err = core.UnmarshalModel(m, "spark_engines", &obj.SparkEngines, UnmarshalSparkEngine)
	if err != nil {
		err = core.SDKErrorf(err, "", "spark_engines-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineDetails : External engine details.
type SparkEngineDetails struct {
	// api key to work with the saas IAE instance.
	ApiKey *string `json:"api_key,omitempty"`

	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// Spark Default Config details.
	DefaultConfig *SparkDefaultConfig `json:"default_config,omitempty"`

	// The default spark version for the native engine.
	DefaultVersion *string `json:"default_version,omitempty"`

	// Application Endpoints.
	Endpoints *SparkEndpoints `json:"endpoints,omitempty"`

	// Default bucket for spark.
	EngineHomeBucketDisplayName *string `json:"engine_home_bucket_display_name,omitempty"`

	// Default bucket for spark.
	EngineHomeBucketName *string `json:"engine_home_bucket_name,omitempty"`

	// Path for spark.
	EngineHomePath *string `json:"engine_home_path,omitempty"`

	// Default volume for spark.
	EngineHomeVolume *string `json:"engine_home_volume,omitempty"`

	// Default volume for spark.
	EngineHomeVolumeID *string `json:"engine_home_volume_id,omitempty"`

	// Name of the volume.
	EngineHomeVolumeName *string `json:"engine_home_volume_name,omitempty"`

	// Storage class of the volume.
	EngineHomeVolumeStorageClass *string `json:"engine_home_volume_storage_class,omitempty"`

	// Storage size of the volume.
	EngineHomeVolumeStorageSize *string `json:"engine_home_volume_storage_size,omitempty"`

	// Instance to access the instance.
	InstanceID *string `json:"instance_id,omitempty"`

	// How is the spark instance managed.
	ManagedBy *string `json:"managed_by,omitempty"`

	// Spark instance scale configuration.
	ScaleConfig *SparkScaleConfig `json:"scale_config,omitempty"`
}

// UnmarshalSparkEngineDetails unmarshals an instance of SparkEngineDetails from the specified map of raw messages.
func UnmarshalSparkEngineDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineDetails)
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		err = core.SDKErrorf(err, "", "api_key-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		err = core.SDKErrorf(err, "", "connection_string-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "default_config", &obj.DefaultConfig, UnmarshalSparkDefaultConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "default_config-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "default_version", &obj.DefaultVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "default_version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "endpoints", &obj.Endpoints, UnmarshalSparkEndpoints)
	if err != nil {
		err = core.SDKErrorf(err, "", "endpoints-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_bucket_display_name", &obj.EngineHomeBucketDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_bucket_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_bucket_name", &obj.EngineHomeBucketName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_bucket_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_path", &obj.EngineHomePath)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_path-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_volume", &obj.EngineHomeVolume)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_volume-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_volume_id", &obj.EngineHomeVolumeID)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_volume_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_volume_name", &obj.EngineHomeVolumeName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_volume_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_volume_storage_class", &obj.EngineHomeVolumeStorageClass)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_volume_storage_class-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_volume_storage_size", &obj.EngineHomeVolumeStorageSize)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_volume_storage_size-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		err = core.SDKErrorf(err, "", "instance_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "managed_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "scale_config", &obj.ScaleConfig, UnmarshalSparkScaleConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "scale_config-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkEngineDetailsPrototype : Node details.
type SparkEngineDetailsPrototype struct {
	// api key to work with the saas IAE instance.
	ApiKey *string `json:"api_key,omitempty"`

	// External engine connection string.
	ConnectionString *string `json:"connection_string,omitempty"`

	// Spark Default Config details.
	DefaultConfig *SparkDefaultConfig `json:"default_config,omitempty"`

	// The default spark version for the native engine.
	DefaultVersion *string `json:"default_version,omitempty"`

	// Default bucket name for spark.
	EngineHomeBucketDisplayName *string `json:"engine_home_bucket_display_name,omitempty"`

	// Default bucket for spark.
	EngineHomeBucketName *string `json:"engine_home_bucket_name,omitempty"`

	// Path for spark.
	EngineHomePath *string `json:"engine_home_path,omitempty"`

	// Default volume for spark.
	EngineHomeVolumeID *string `json:"engine_home_volume_id,omitempty"`

	// Name of the volume.
	EngineHomeVolumeName *string `json:"engine_home_volume_name,omitempty"`

	// Storage class of the volume.
	EngineHomeVolumeStorageClass *string `json:"engine_home_volume_storage_class,omitempty"`

	// Storage size of the volume.
	EngineHomeVolumeStorageSize *string `json:"engine_home_volume_storage_size,omitempty"`

	// Instance to access the instance.
	InstanceID *string `json:"instance_id,omitempty"`

	// How is the spark instance managed.
	ManagedBy *string `json:"managed_by,omitempty"`

	// Spark instance scale configuration.
	ScaleConfig *SparkScaleConfig `json:"scale_config,omitempty"`
}

// UnmarshalSparkEngineDetailsPrototype unmarshals an instance of SparkEngineDetailsPrototype from the specified map of raw messages.
func UnmarshalSparkEngineDetailsPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkEngineDetailsPrototype)
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		err = core.SDKErrorf(err, "", "api_key-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		err = core.SDKErrorf(err, "", "connection_string-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "default_config", &obj.DefaultConfig, UnmarshalSparkDefaultConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "default_config-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "default_version", &obj.DefaultVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "default_version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_bucket_display_name", &obj.EngineHomeBucketDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_bucket_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_bucket_name", &obj.EngineHomeBucketName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_bucket_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_path", &obj.EngineHomePath)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_path-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_volume_id", &obj.EngineHomeVolumeID)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_volume_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_volume_name", &obj.EngineHomeVolumeName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_volume_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_volume_storage_class", &obj.EngineHomeVolumeStorageClass)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_volume_storage_class-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_home_volume_storage_size", &obj.EngineHomeVolumeStorageSize)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_home_volume_storage_size-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		err = core.SDKErrorf(err, "", "instance_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "managed_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "scale_config", &obj.ScaleConfig, UnmarshalSparkScaleConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "scale_config-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkHistoryServer : Native spark history server.
type SparkHistoryServer struct {
	// History server start time.
	AutoTerminationTime *string `json:"auto_termination_time,omitempty"`

	// History server cores.
	Cores *string `json:"cores,omitempty"`

	// History server memory.
	Memory *string `json:"memory,omitempty"`

	// History server start time.
	StartTime *string `json:"start_time,omitempty"`

	// History server state.
	State *string `json:"state,omitempty"`
}

// UnmarshalSparkHistoryServer unmarshals an instance of SparkHistoryServer from the specified map of raw messages.
func UnmarshalSparkHistoryServer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkHistoryServer)
	err = core.UnmarshalPrimitive(m, "auto_termination_time", &obj.AutoTerminationTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_termination_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "cores", &obj.Cores)
	if err != nil {
		err = core.SDKErrorf(err, "", "cores-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "memory", &obj.Memory)
	if err != nil {
		err = core.SDKErrorf(err, "", "memory-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "start_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkScaleConfig : Spark instance scale configuration.
type SparkScaleConfig struct {
	// Enable/disable autoscaling.
	AutoScaleEnabled *bool `json:"auto_scale_enabled,omitempty"`

	// Current node count.
	CurrentNumberOfNodes *int64 `json:"current_number_of_nodes,omitempty"`

	// Maximum node count.
	MaximumNumberOfNodes *int64 `json:"maximum_number_of_nodes,omitempty"`

	// Minimum node count.
	MinimumNumberOfNodes *int64 `json:"minimum_number_of_nodes,omitempty"`

	// Spark instance node type.
	NodeType *string `json:"node_type,omitempty"`

	// Node count.
	NumberOfNodes *int64 `json:"number_of_nodes,omitempty"`
}

// UnmarshalSparkScaleConfig unmarshals an instance of SparkScaleConfig from the specified map of raw messages.
func UnmarshalSparkScaleConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkScaleConfig)
	err = core.UnmarshalPrimitive(m, "auto_scale_enabled", &obj.AutoScaleEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scale_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "current_number_of_nodes", &obj.CurrentNumberOfNodes)
	if err != nil {
		err = core.SDKErrorf(err, "", "current_number_of_nodes-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_number_of_nodes", &obj.MaximumNumberOfNodes)
	if err != nil {
		err = core.SDKErrorf(err, "", "maximum_number_of_nodes-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_number_of_nodes", &obj.MinimumNumberOfNodes)
	if err != nil {
		err = core.SDKErrorf(err, "", "minimum_number_of_nodes-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "node_type", &obj.NodeType)
	if err != nil {
		err = core.SDKErrorf(err, "", "node_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "number_of_nodes", &obj.NumberOfNodes)
	if err != nil {
		err = core.SDKErrorf(err, "", "number_of_nodes-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkVolumeDetails : Spark application volume.
type SparkVolumeDetails struct {
	// Path in the spark cluster for the mounted volume.
	MountPath *string `json:"mount_path,omitempty"`

	// volume name.
	Name *string `json:"name,omitempty"`

	// Read only flag.
	ReadOnly *bool `json:"read_only,omitempty"`

	// Path in the volume to be mounted.
	SourceSubPath *string `json:"source_sub_path,omitempty"`
}

// UnmarshalSparkVolumeDetails unmarshals an instance of SparkVolumeDetails from the specified map of raw messages.
func UnmarshalSparkVolumeDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkVolumeDetails)
	err = core.UnmarshalPrimitive(m, "mount_path", &obj.MountPath)
	if err != nil {
		err = core.SDKErrorf(err, "", "mount_path-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "read_only", &obj.ReadOnly)
	if err != nil {
		err = core.SDKErrorf(err, "", "read_only-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "source_sub_path", &obj.SourceSubPath)
	if err != nil {
		err = core.SDKErrorf(err, "", "source_sub_path-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StartSparkEngineHistoryServerOptions : The StartSparkEngineHistoryServer options.
type StartSparkEngineHistoryServerOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// CPU count.
	Cores *string `json:"cores,omitempty"`

	// Memory in GiB.
	Memory *string `json:"memory,omitempty"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewStartSparkEngineHistoryServerOptions : Instantiate StartSparkEngineHistoryServerOptions
func (*WatsonxDataV2) NewStartSparkEngineHistoryServerOptions(engineID string) *StartSparkEngineHistoryServerOptions {
	return &StartSparkEngineHistoryServerOptions{
		EngineID: core.StringPtr(engineID),
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *StartSparkEngineHistoryServerOptions) SetEngineID(engineID string) *StartSparkEngineHistoryServerOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCores : Allow user to set Cores
func (_options *StartSparkEngineHistoryServerOptions) SetCores(cores string) *StartSparkEngineHistoryServerOptions {
	_options.Cores = core.StringPtr(cores)
	return _options
}

// SetMemory : Allow user to set Memory
func (_options *StartSparkEngineHistoryServerOptions) SetMemory(memory string) *StartSparkEngineHistoryServerOptions {
	_options.Memory = core.StringPtr(memory)
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *StartSparkEngineHistoryServerOptions) SetAuthInstanceID(authInstanceID string) *StartSparkEngineHistoryServerOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *StartSparkEngineHistoryServerOptions) SetHeaders(param map[string]string) *StartSparkEngineHistoryServerOptions {
	options.Headers = param
	return options
}

// SuccessResponse : Response of success.
type SuccessResponse struct {
	// Message.
	Message *string `json:"message,omitempty"`

	// Message code.
	MessageCode *string `json:"message_code,omitempty"`
}

// UnmarshalSuccessResponse unmarshals an instance of SuccessResponse from the specified map of raw messages.
func UnmarshalSuccessResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "message_code", &obj.MessageCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "message_code-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyncCatalogs : catalogs definition.
type SyncCatalogs struct {
	// Auto add new table.
	AutoAddNewTables *bool `json:"auto_add_new_tables" validate:"required"`

	// Sync iceberg metadata.
	SyncIcebergMd *bool `json:"sync_iceberg_md" validate:"required"`
}

// NewSyncCatalogs : Instantiate SyncCatalogs (Generic Model Constructor)
func (*WatsonxDataV2) NewSyncCatalogs(autoAddNewTables bool, syncIcebergMd bool) (_model *SyncCatalogs, err error) {
	_model = &SyncCatalogs{
		AutoAddNewTables: core.BoolPtr(autoAddNewTables),
		SyncIcebergMd: core.BoolPtr(syncIcebergMd),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalSyncCatalogs unmarshals an instance of SyncCatalogs from the specified map of raw messages.
func UnmarshalSyncCatalogs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyncCatalogs)
	err = core.UnmarshalPrimitive(m, "auto_add_new_tables", &obj.AutoAddNewTables)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_add_new_tables-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "sync_iceberg_md", &obj.SyncIcebergMd)
	if err != nil {
		err = core.SDKErrorf(err, "", "sync_iceberg_md-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the SyncCatalogs
func (syncCatalogs *SyncCatalogs) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(syncCatalogs.AutoAddNewTables) {
		_patch["auto_add_new_tables"] = syncCatalogs.AutoAddNewTables
	}
	if !core.IsNil(syncCatalogs.SyncIcebergMd) {
		_patch["sync_iceberg_md"] = syncCatalogs.SyncIcebergMd
	}

	return
}

// Table : GetColumns OK.
type Table struct {
	// Columns.
	Columns []Column `json:"columns,omitempty"`

	// Table name.
	TableName *string `json:"table_name,omitempty"`
}

// UnmarshalTable unmarshals an instance of Table from the specified map of raw messages.
func UnmarshalTable(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Table)
	err = core.UnmarshalModel(m, "columns", &obj.Columns, UnmarshalColumn)
	if err != nil {
		err = core.SDKErrorf(err, "", "columns-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "table_name", &obj.TableName)
	if err != nil {
		err = core.SDKErrorf(err, "", "table_name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableCollection : tables list.
type TableCollection struct {
	// List of the tables present in the schema.
	Tables []string `json:"tables,omitempty"`
}

// UnmarshalTableCollection unmarshals an instance of TableCollection from the specified map of raw messages.
func UnmarshalTableCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableCollection)
	err = core.UnmarshalPrimitive(m, "tables", &obj.Tables)
	if err != nil {
		err = core.SDKErrorf(err, "", "tables-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TablePatch : UpdateTable body.
type TablePatch struct {
	// New table name.
	TableName *string `json:"table_name,omitempty"`
}

// UnmarshalTablePatch unmarshals an instance of TablePatch from the specified map of raw messages.
func UnmarshalTablePatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TablePatch)
	err = core.UnmarshalPrimitive(m, "table_name", &obj.TableName)
	if err != nil {
		err = core.SDKErrorf(err, "", "table_name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the TablePatch
func (tablePatch *TablePatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(tablePatch.TableName) {
		_patch["table_name"] = tablePatch.TableName
	}

	return
}

// TableSnapshot : TableSnapshot.
type TableSnapshot struct {
	// Committed at.
	CommittedAt *string `json:"committed_at,omitempty"`

	// Operation.
	Operation *string `json:"operation,omitempty"`

	// Snapshot id.
	SnapshotID *string `json:"snapshot_id,omitempty"`

	// Summary.
	Summary *string `json:"summary,omitempty"`
}

// UnmarshalTableSnapshot unmarshals an instance of TableSnapshot from the specified map of raw messages.
func UnmarshalTableSnapshot(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableSnapshot)
	err = core.UnmarshalPrimitive(m, "committed_at", &obj.CommittedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "committed_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "operation", &obj.Operation)
	if err != nil {
		err = core.SDKErrorf(err, "", "operation-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "snapshot_id", &obj.SnapshotID)
	if err != nil {
		err = core.SDKErrorf(err, "", "snapshot_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "summary", &obj.Summary)
	if err != nil {
		err = core.SDKErrorf(err, "", "summary-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TableSnapshotCollection : TableSnapshot OK.
type TableSnapshotCollection struct {
	// Snapshots.
	Snapshots []TableSnapshot `json:"snapshots,omitempty"`
}

// UnmarshalTableSnapshotCollection unmarshals an instance of TableSnapshotCollection from the specified map of raw messages.
func UnmarshalTableSnapshotCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TableSnapshotCollection)
	err = core.UnmarshalModel(m, "snapshots", &obj.Snapshots, UnmarshalTableSnapshot)
	if err != nil {
		err = core.SDKErrorf(err, "", "snapshots-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateBucketRegistrationOptions : The UpdateBucketRegistration options.
type UpdateBucketRegistrationOptions struct {
	// bucket id.
	BucketID *string `json:"bucket_id" validate:"required,ne="`

	// Request body.
	Body map[string]interface{} `json:"body" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateBucketRegistrationOptions : Instantiate UpdateBucketRegistrationOptions
func (*WatsonxDataV2) NewUpdateBucketRegistrationOptions(bucketID string, body map[string]interface{}) *UpdateBucketRegistrationOptions {
	return &UpdateBucketRegistrationOptions{
		BucketID: core.StringPtr(bucketID),
		Body: body,
	}
}

// SetBucketID : Allow user to set BucketID
func (_options *UpdateBucketRegistrationOptions) SetBucketID(bucketID string) *UpdateBucketRegistrationOptions {
	_options.BucketID = core.StringPtr(bucketID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateBucketRegistrationOptions) SetBody(body map[string]interface{}) *UpdateBucketRegistrationOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateBucketRegistrationOptions) SetAuthInstanceID(authInstanceID string) *UpdateBucketRegistrationOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateBucketRegistrationOptions) SetHeaders(param map[string]string) *UpdateBucketRegistrationOptions {
	options.Headers = param
	return options
}

// UpdateColumnOptions : The UpdateColumn options.
type UpdateColumnOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required"`

	// catalog id.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// URL encoded schema name.
	SchemaID *string `json:"schema_id" validate:"required,ne="`

	// URL encoded schema name.
	TableID *string `json:"table_id" validate:"required,ne="`

	// URL encoded schema name.
	ColumnID *string `json:"column_id" validate:"required,ne="`

	// Request body.
	Body map[string]interface{} `json:"body" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateColumnOptions : Instantiate UpdateColumnOptions
func (*WatsonxDataV2) NewUpdateColumnOptions(engineID string, catalogID string, schemaID string, tableID string, columnID string, body map[string]interface{}) *UpdateColumnOptions {
	return &UpdateColumnOptions{
		EngineID: core.StringPtr(engineID),
		CatalogID: core.StringPtr(catalogID),
		SchemaID: core.StringPtr(schemaID),
		TableID: core.StringPtr(tableID),
		ColumnID: core.StringPtr(columnID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateColumnOptions) SetEngineID(engineID string) *UpdateColumnOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *UpdateColumnOptions) SetCatalogID(catalogID string) *UpdateColumnOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetSchemaID : Allow user to set SchemaID
func (_options *UpdateColumnOptions) SetSchemaID(schemaID string) *UpdateColumnOptions {
	_options.SchemaID = core.StringPtr(schemaID)
	return _options
}

// SetTableID : Allow user to set TableID
func (_options *UpdateColumnOptions) SetTableID(tableID string) *UpdateColumnOptions {
	_options.TableID = core.StringPtr(tableID)
	return _options
}

// SetColumnID : Allow user to set ColumnID
func (_options *UpdateColumnOptions) SetColumnID(columnID string) *UpdateColumnOptions {
	_options.ColumnID = core.StringPtr(columnID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateColumnOptions) SetBody(body map[string]interface{}) *UpdateColumnOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateColumnOptions) SetAuthInstanceID(authInstanceID string) *UpdateColumnOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateColumnOptions) SetHeaders(param map[string]string) *UpdateColumnOptions {
	options.Headers = param
	return options
}

// UpdateDatabaseOptions : The UpdateDatabase options.
type UpdateDatabaseOptions struct {
	// database id.
	DatabaseID *string `json:"database_id" validate:"required,ne="`

	// Request body.
	Body map[string]interface{} `json:"body" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateDatabaseOptions : Instantiate UpdateDatabaseOptions
func (*WatsonxDataV2) NewUpdateDatabaseOptions(databaseID string, body map[string]interface{}) *UpdateDatabaseOptions {
	return &UpdateDatabaseOptions{
		DatabaseID: core.StringPtr(databaseID),
		Body: body,
	}
}

// SetDatabaseID : Allow user to set DatabaseID
func (_options *UpdateDatabaseOptions) SetDatabaseID(databaseID string) *UpdateDatabaseOptions {
	_options.DatabaseID = core.StringPtr(databaseID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateDatabaseOptions) SetBody(body map[string]interface{}) *UpdateDatabaseOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateDatabaseOptions) SetAuthInstanceID(authInstanceID string) *UpdateDatabaseOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDatabaseOptions) SetHeaders(param map[string]string) *UpdateDatabaseOptions {
	options.Headers = param
	return options
}

// UpdateDb2EngineOptions : The UpdateDb2Engine options.
type UpdateDb2EngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Update Engine Body.
	Body map[string]interface{} `json:"body" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateDb2EngineOptions : Instantiate UpdateDb2EngineOptions
func (*WatsonxDataV2) NewUpdateDb2EngineOptions(engineID string, body map[string]interface{}) *UpdateDb2EngineOptions {
	return &UpdateDb2EngineOptions{
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateDb2EngineOptions) SetEngineID(engineID string) *UpdateDb2EngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateDb2EngineOptions) SetBody(body map[string]interface{}) *UpdateDb2EngineOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateDb2EngineOptions) SetAuthInstanceID(authInstanceID string) *UpdateDb2EngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDb2EngineOptions) SetHeaders(param map[string]string) *UpdateDb2EngineOptions {
	options.Headers = param
	return options
}

// UpdateMilvusServiceOptions : The UpdateMilvusService options.
type UpdateMilvusServiceOptions struct {
	// service id.
	ServiceID *string `json:"service_id" validate:"required,ne="`

	// Update milvus service.
	Body map[string]interface{} `json:"body" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateMilvusServiceOptions : Instantiate UpdateMilvusServiceOptions
func (*WatsonxDataV2) NewUpdateMilvusServiceOptions(serviceID string, body map[string]interface{}) *UpdateMilvusServiceOptions {
	return &UpdateMilvusServiceOptions{
		ServiceID: core.StringPtr(serviceID),
		Body: body,
	}
}

// SetServiceID : Allow user to set ServiceID
func (_options *UpdateMilvusServiceOptions) SetServiceID(serviceID string) *UpdateMilvusServiceOptions {
	_options.ServiceID = core.StringPtr(serviceID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateMilvusServiceOptions) SetBody(body map[string]interface{}) *UpdateMilvusServiceOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateMilvusServiceOptions) SetAuthInstanceID(authInstanceID string) *UpdateMilvusServiceOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateMilvusServiceOptions) SetHeaders(param map[string]string) *UpdateMilvusServiceOptions {
	options.Headers = param
	return options
}

// UpdateNetezzaEngineOptions : The UpdateNetezzaEngine options.
type UpdateNetezzaEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Update Engine Body.
	Body map[string]interface{} `json:"body" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateNetezzaEngineOptions : Instantiate UpdateNetezzaEngineOptions
func (*WatsonxDataV2) NewUpdateNetezzaEngineOptions(engineID string, body map[string]interface{}) *UpdateNetezzaEngineOptions {
	return &UpdateNetezzaEngineOptions{
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateNetezzaEngineOptions) SetEngineID(engineID string) *UpdateNetezzaEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateNetezzaEngineOptions) SetBody(body map[string]interface{}) *UpdateNetezzaEngineOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateNetezzaEngineOptions) SetAuthInstanceID(authInstanceID string) *UpdateNetezzaEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateNetezzaEngineOptions) SetHeaders(param map[string]string) *UpdateNetezzaEngineOptions {
	options.Headers = param
	return options
}

// UpdatePrestissimoEngineOptions : The UpdatePrestissimoEngine options.
type UpdatePrestissimoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Update prestissimo engine body.
	Body map[string]interface{} `json:"body" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdatePrestissimoEngineOptions : Instantiate UpdatePrestissimoEngineOptions
func (*WatsonxDataV2) NewUpdatePrestissimoEngineOptions(engineID string, body map[string]interface{}) *UpdatePrestissimoEngineOptions {
	return &UpdatePrestissimoEngineOptions{
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdatePrestissimoEngineOptions) SetEngineID(engineID string) *UpdatePrestissimoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdatePrestissimoEngineOptions) SetBody(body map[string]interface{}) *UpdatePrestissimoEngineOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdatePrestissimoEngineOptions) SetAuthInstanceID(authInstanceID string) *UpdatePrestissimoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePrestissimoEngineOptions) SetHeaders(param map[string]string) *UpdatePrestissimoEngineOptions {
	options.Headers = param
	return options
}

// UpdatePrestoEngineOptions : The UpdatePrestoEngine options.
type UpdatePrestoEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Update Engine Body.
	Body map[string]interface{} `json:"body" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdatePrestoEngineOptions : Instantiate UpdatePrestoEngineOptions
func (*WatsonxDataV2) NewUpdatePrestoEngineOptions(engineID string, body map[string]interface{}) *UpdatePrestoEngineOptions {
	return &UpdatePrestoEngineOptions{
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdatePrestoEngineOptions) SetEngineID(engineID string) *UpdatePrestoEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdatePrestoEngineOptions) SetBody(body map[string]interface{}) *UpdatePrestoEngineOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdatePrestoEngineOptions) SetAuthInstanceID(authInstanceID string) *UpdatePrestoEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePrestoEngineOptions) SetHeaders(param map[string]string) *UpdatePrestoEngineOptions {
	options.Headers = param
	return options
}

// UpdateSparkEngineBody : UpdateEngine body.
type UpdateSparkEngineBody struct {
	// Modified description.
	Description *string `json:"description,omitempty"`

	// Engine details.
	EngineDetails *UpdateSparkEngineBodyEngineDetails `json:"engine_details,omitempty"`

	// Engine display name.
	EngineDisplayName *string `json:"engine_display_name,omitempty"`

	// Tags.
	Tags []string `json:"tags,omitempty"`
}

// UnmarshalUpdateSparkEngineBody unmarshals an instance of UpdateSparkEngineBody from the specified map of raw messages.
func UnmarshalUpdateSparkEngineBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateSparkEngineBody)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalUpdateSparkEngineBodyEngineDetails)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_details-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_display_name", &obj.EngineDisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the UpdateSparkEngineBody
func (updateSparkEngineBody *UpdateSparkEngineBody) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(updateSparkEngineBody.Description) {
		_patch["description"] = updateSparkEngineBody.Description
	}
	if !core.IsNil(updateSparkEngineBody.EngineDetails) {
		_patch["engine_details"] = updateSparkEngineBody.EngineDetails.asPatch()
	}
	if !core.IsNil(updateSparkEngineBody.EngineDisplayName) {
		_patch["engine_display_name"] = updateSparkEngineBody.EngineDisplayName
	}
	if !core.IsNil(updateSparkEngineBody.Tags) {
		_patch["tags"] = updateSparkEngineBody.Tags
	}

	return
}

// UpdateSparkEngineBodyEngineDetails : Engine details.
type UpdateSparkEngineBodyEngineDetails struct {
	// Dynamic dict.
	DefaultConfig map[string]string `json:"default_config,omitempty"`

	// The default spark version for the native engine.
	DefaultVersion *string `json:"default_version,omitempty"`
}

// UnmarshalUpdateSparkEngineBodyEngineDetails unmarshals an instance of UpdateSparkEngineBodyEngineDetails from the specified map of raw messages.
func UnmarshalUpdateSparkEngineBodyEngineDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateSparkEngineBodyEngineDetails)
	err = core.UnmarshalPrimitive(m, "default_config", &obj.DefaultConfig)
	if err != nil {
		err = core.SDKErrorf(err, "", "default_config-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "default_version", &obj.DefaultVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "default_version-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the UpdateSparkEngineBodyEngineDetails
func (updateSparkEngineBodyEngineDetails *UpdateSparkEngineBodyEngineDetails) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(updateSparkEngineBodyEngineDetails.DefaultConfig) {
		_patch["default_config"] = updateSparkEngineBodyEngineDetails.DefaultConfig
	}
	if !core.IsNil(updateSparkEngineBodyEngineDetails.DefaultVersion) {
		_patch["default_version"] = updateSparkEngineBodyEngineDetails.DefaultVersion
	}

	return
}

// UpdateSparkEngineOptions : The UpdateSparkEngine options.
type UpdateSparkEngineOptions struct {
	// engine id.
	EngineID *string `json:"engine_id" validate:"required,ne="`

	// Update Engine Body.
	Body map[string]interface{} `json:"body" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateSparkEngineOptions : Instantiate UpdateSparkEngineOptions
func (*WatsonxDataV2) NewUpdateSparkEngineOptions(engineID string, body map[string]interface{}) *UpdateSparkEngineOptions {
	return &UpdateSparkEngineOptions{
		EngineID: core.StringPtr(engineID),
		Body: body,
	}
}

// SetEngineID : Allow user to set EngineID
func (_options *UpdateSparkEngineOptions) SetEngineID(engineID string) *UpdateSparkEngineOptions {
	_options.EngineID = core.StringPtr(engineID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateSparkEngineOptions) SetBody(body map[string]interface{}) *UpdateSparkEngineOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateSparkEngineOptions) SetAuthInstanceID(authInstanceID string) *UpdateSparkEngineOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSparkEngineOptions) SetHeaders(param map[string]string) *UpdateSparkEngineOptions {
	options.Headers = param
	return options
}

// UpdateSyncCatalogOKBody : success response.
type UpdateSyncCatalogOKBody struct {
	// Response of success.
	Response *SuccessResponse `json:"response,omitempty"`
}

// UnmarshalUpdateSyncCatalogOKBody unmarshals an instance of UpdateSyncCatalogOKBody from the specified map of raw messages.
func UnmarshalUpdateSyncCatalogOKBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateSyncCatalogOKBody)
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalSuccessResponse)
	if err != nil {
		err = core.SDKErrorf(err, "", "response-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateSyncCatalogOptions : The UpdateSyncCatalog options.
type UpdateSyncCatalogOptions struct {
	// catalog ID.
	CatalogID *string `json:"catalog_id" validate:"required,ne="`

	// Request body.
	Body map[string]interface{} `json:"body" validate:"required"`

	// CRN.
	AuthInstanceID *string `json:"AuthInstanceId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateSyncCatalogOptions : Instantiate UpdateSyncCatalogOptions
func (*WatsonxDataV2) NewUpdateSyncCatalogOptions(catalogID string, body map[string]interface{}) *UpdateSyncCatalogOptions {
	return &UpdateSyncCatalogOptions{
		CatalogID: core.StringPtr(catalogID),
		Body: body,
	}
}

// SetCatalogID : Allow user to set CatalogID
func (_options *UpdateSyncCatalogOptions) SetCatalogID(catalogID string) *UpdateSyncCatalogOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateSyncCatalogOptions) SetBody(body map[string]interface{}) *UpdateSyncCatalogOptions {
	_options.Body = body
	return _options
}

// SetAuthInstanceID : Allow user to set AuthInstanceID
func (_options *UpdateSyncCatalogOptions) SetAuthInstanceID(authInstanceID string) *UpdateSyncCatalogOptions {
	_options.AuthInstanceID = core.StringPtr(authInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSyncCatalogOptions) SetHeaders(param map[string]string) *UpdateSyncCatalogOptions {
	options.Headers = param
	return options
}

//
// IngestionJobsPager can be used to simplify the use of the "ListIngestionJobs" method.
//
type IngestionJobsPager struct {
	hasNext bool
	options *ListIngestionJobsOptions
	client  *WatsonxDataV2
	pageContext struct {
		next *string
	}
}

// NewIngestionJobsPager returns a new IngestionJobsPager instance.
func (watsonxData *WatsonxDataV2) NewIngestionJobsPager(options *ListIngestionJobsOptions) (pager *IngestionJobsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = core.SDKErrorf(nil, "the 'options.Start' field should not be set", "no-query-setting", common.GetComponentInfo())
		return
	}

	var optionsCopy ListIngestionJobsOptions = *options
	pager = &IngestionJobsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonxData,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *IngestionJobsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *IngestionJobsPager) GetNextWithContext(ctx context.Context) (page []IngestionJob, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListIngestionJobsWithContext(ctx, pager.options)
	if err != nil {
		err = core.RepurposeSDKProblem(err, "error-getting-next-page")
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			errMsg := fmt.Sprintf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			err = core.SDKErrorf(err, errMsg, "get-query-error", common.GetComponentInfo())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.IngestionJobs

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *IngestionJobsPager) GetAllWithContext(ctx context.Context) (allItems []IngestionJob, err error) {
	for pager.HasNext() {
		var nextPage []IngestionJob
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			err = core.RepurposeSDKProblem(err, "error-getting-next-page")
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *IngestionJobsPager) GetNext() (page []IngestionJob, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *IngestionJobsPager) GetAll() (allItems []IngestionJob, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}
