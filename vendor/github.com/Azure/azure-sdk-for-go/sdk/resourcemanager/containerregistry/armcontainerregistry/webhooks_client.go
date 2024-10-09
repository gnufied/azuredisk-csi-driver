//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WebhooksClient contains the methods for the Webhooks group.
// Don't use this type directly, use NewWebhooksClient() instead.
type WebhooksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWebhooksClient creates a new instance of WebhooksClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWebhooksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebhooksClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WebhooksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a webhook for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - webhookName - The name of the webhook.
//   - webhookCreateParameters - The parameters for creating a webhook.
//   - options - WebhooksClientBeginCreateOptions contains the optional parameters for the WebhooksClient.BeginCreate method.
func (client *WebhooksClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookCreateParameters WebhookCreateParameters, options *WebhooksClientBeginCreateOptions) (*runtime.Poller[WebhooksClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, webhookName, webhookCreateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WebhooksClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WebhooksClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates a webhook for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *WebhooksClient) create(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookCreateParameters WebhookCreateParameters, options *WebhooksClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "WebhooksClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, webhookName, webhookCreateParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *WebhooksClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookCreateParameters WebhookCreateParameters, options *WebhooksClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if webhookName == "" {
		return nil, errors.New("parameter webhookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webhookName}", url.PathEscape(webhookName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, webhookCreateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a webhook from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - webhookName - The name of the webhook.
//   - options - WebhooksClientBeginDeleteOptions contains the optional parameters for the WebhooksClient.BeginDelete method.
func (client *WebhooksClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, webhookName string, options *WebhooksClientBeginDeleteOptions) (*runtime.Poller[WebhooksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, webhookName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WebhooksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WebhooksClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a webhook from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *WebhooksClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, webhookName string, options *WebhooksClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "WebhooksClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, webhookName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WebhooksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, webhookName string, options *WebhooksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if webhookName == "" {
		return nil, errors.New("parameter webhookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webhookName}", url.PathEscape(webhookName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the properties of the specified webhook.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - webhookName - The name of the webhook.
//   - options - WebhooksClientGetOptions contains the optional parameters for the WebhooksClient.Get method.
func (client *WebhooksClient) Get(ctx context.Context, resourceGroupName string, registryName string, webhookName string, options *WebhooksClientGetOptions) (WebhooksClientGetResponse, error) {
	var err error
	const operationName = "WebhooksClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, webhookName, options)
	if err != nil {
		return WebhooksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebhooksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebhooksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WebhooksClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, webhookName string, options *WebhooksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if webhookName == "" {
		return nil, errors.New("parameter webhookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webhookName}", url.PathEscape(webhookName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebhooksClient) getHandleResponse(resp *http.Response) (WebhooksClientGetResponse, error) {
	result := WebhooksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Webhook); err != nil {
		return WebhooksClientGetResponse{}, err
	}
	return result, nil
}

// GetCallbackConfig - Gets the configuration of service URI and custom headers for the webhook.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - webhookName - The name of the webhook.
//   - options - WebhooksClientGetCallbackConfigOptions contains the optional parameters for the WebhooksClient.GetCallbackConfig
//     method.
func (client *WebhooksClient) GetCallbackConfig(ctx context.Context, resourceGroupName string, registryName string, webhookName string, options *WebhooksClientGetCallbackConfigOptions) (WebhooksClientGetCallbackConfigResponse, error) {
	var err error
	const operationName = "WebhooksClient.GetCallbackConfig"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCallbackConfigCreateRequest(ctx, resourceGroupName, registryName, webhookName, options)
	if err != nil {
		return WebhooksClientGetCallbackConfigResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebhooksClientGetCallbackConfigResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebhooksClientGetCallbackConfigResponse{}, err
	}
	resp, err := client.getCallbackConfigHandleResponse(httpResp)
	return resp, err
}

// getCallbackConfigCreateRequest creates the GetCallbackConfig request.
func (client *WebhooksClient) getCallbackConfigCreateRequest(ctx context.Context, resourceGroupName string, registryName string, webhookName string, options *WebhooksClientGetCallbackConfigOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}/getCallbackConfig"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if webhookName == "" {
		return nil, errors.New("parameter webhookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webhookName}", url.PathEscape(webhookName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getCallbackConfigHandleResponse handles the GetCallbackConfig response.
func (client *WebhooksClient) getCallbackConfigHandleResponse(resp *http.Response) (WebhooksClientGetCallbackConfigResponse, error) {
	result := WebhooksClientGetCallbackConfigResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CallbackConfig); err != nil {
		return WebhooksClientGetCallbackConfigResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the webhooks for the specified container registry.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - options - WebhooksClientListOptions contains the optional parameters for the WebhooksClient.NewListPager method.
func (client *WebhooksClient) NewListPager(resourceGroupName string, registryName string, options *WebhooksClientListOptions) *runtime.Pager[WebhooksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebhooksClientListResponse]{
		More: func(page WebhooksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebhooksClientListResponse) (WebhooksClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WebhooksClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			}, nil)
			if err != nil {
				return WebhooksClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *WebhooksClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *WebhooksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WebhooksClient) listHandleResponse(resp *http.Response) (WebhooksClientListResponse, error) {
	result := WebhooksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebhookListResult); err != nil {
		return WebhooksClientListResponse{}, err
	}
	return result, nil
}

// NewListEventsPager - Lists recent events for the specified webhook.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - webhookName - The name of the webhook.
//   - options - WebhooksClientListEventsOptions contains the optional parameters for the WebhooksClient.NewListEventsPager method.
func (client *WebhooksClient) NewListEventsPager(resourceGroupName string, registryName string, webhookName string, options *WebhooksClientListEventsOptions) *runtime.Pager[WebhooksClientListEventsResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebhooksClientListEventsResponse]{
		More: func(page WebhooksClientListEventsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebhooksClientListEventsResponse) (WebhooksClientListEventsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WebhooksClient.NewListEventsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listEventsCreateRequest(ctx, resourceGroupName, registryName, webhookName, options)
			}, nil)
			if err != nil {
				return WebhooksClientListEventsResponse{}, err
			}
			return client.listEventsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listEventsCreateRequest creates the ListEvents request.
func (client *WebhooksClient) listEventsCreateRequest(ctx context.Context, resourceGroupName string, registryName string, webhookName string, options *WebhooksClientListEventsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}/listEvents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if webhookName == "" {
		return nil, errors.New("parameter webhookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webhookName}", url.PathEscape(webhookName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listEventsHandleResponse handles the ListEvents response.
func (client *WebhooksClient) listEventsHandleResponse(resp *http.Response) (WebhooksClientListEventsResponse, error) {
	result := WebhooksClientListEventsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventListResult); err != nil {
		return WebhooksClientListEventsResponse{}, err
	}
	return result, nil
}

// Ping - Triggers a ping event to be sent to the webhook.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - webhookName - The name of the webhook.
//   - options - WebhooksClientPingOptions contains the optional parameters for the WebhooksClient.Ping method.
func (client *WebhooksClient) Ping(ctx context.Context, resourceGroupName string, registryName string, webhookName string, options *WebhooksClientPingOptions) (WebhooksClientPingResponse, error) {
	var err error
	const operationName = "WebhooksClient.Ping"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.pingCreateRequest(ctx, resourceGroupName, registryName, webhookName, options)
	if err != nil {
		return WebhooksClientPingResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebhooksClientPingResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebhooksClientPingResponse{}, err
	}
	resp, err := client.pingHandleResponse(httpResp)
	return resp, err
}

// pingCreateRequest creates the Ping request.
func (client *WebhooksClient) pingCreateRequest(ctx context.Context, resourceGroupName string, registryName string, webhookName string, options *WebhooksClientPingOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}/ping"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if webhookName == "" {
		return nil, errors.New("parameter webhookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webhookName}", url.PathEscape(webhookName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// pingHandleResponse handles the Ping response.
func (client *WebhooksClient) pingHandleResponse(resp *http.Response) (WebhooksClientPingResponse, error) {
	result := WebhooksClientPingResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventInfo); err != nil {
		return WebhooksClientPingResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a webhook with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - webhookName - The name of the webhook.
//   - webhookUpdateParameters - The parameters for updating a webhook.
//   - options - WebhooksClientBeginUpdateOptions contains the optional parameters for the WebhooksClient.BeginUpdate method.
func (client *WebhooksClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookUpdateParameters WebhookUpdateParameters, options *WebhooksClientBeginUpdateOptions) (*runtime.Poller[WebhooksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, registryName, webhookName, webhookUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WebhooksClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WebhooksClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates a webhook with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *WebhooksClient) update(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookUpdateParameters WebhookUpdateParameters, options *WebhooksClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "WebhooksClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, webhookName, webhookUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *WebhooksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookUpdateParameters WebhookUpdateParameters, options *WebhooksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if webhookName == "" {
		return nil, errors.New("parameter webhookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webhookName}", url.PathEscape(webhookName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, webhookUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
