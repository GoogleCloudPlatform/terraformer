package compute

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DedicatedHostsClient is the compute Client
type DedicatedHostsClient struct {
	BaseClient
}

// NewDedicatedHostsClient creates an instance of the DedicatedHostsClient client.
func NewDedicatedHostsClient(subscriptionID string) DedicatedHostsClient {
	return NewDedicatedHostsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDedicatedHostsClientWithBaseURI creates an instance of the DedicatedHostsClient client.
func NewDedicatedHostsClientWithBaseURI(baseURI string, subscriptionID string) DedicatedHostsClient {
	return DedicatedHostsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a dedicated host .
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
// hostName - the name of the dedicated host .
// parameters - parameters supplied to the Create Dedicated Host.
func (client DedicatedHostsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost) (result DedicatedHostsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.DedicatedHostProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.DedicatedHostProperties.PlatformFaultDomain", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.DedicatedHostProperties.PlatformFaultDomain", Name: validation.InclusiveMaximum, Rule: int64(2), Chain: nil},
						{Target: "parameters.DedicatedHostProperties.PlatformFaultDomain", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
					}},
				}},
				{Target: "parameters.Sku", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("compute.DedicatedHostsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, hostGroupName, hostName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DedicatedHostsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) CreateOrUpdateSender(req *http.Request) (future DedicatedHostsCreateOrUpdateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) CreateOrUpdateResponder(resp *http.Response) (result DedicatedHost, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a dedicated host.
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
// hostName - the name of the dedicated host.
func (client DedicatedHostsClient) Delete(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string) (result DedicatedHostsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, hostGroupName, hostName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DedicatedHostsClient) DeletePreparer(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) DeleteSender(req *http.Request) (future DedicatedHostsDeleteFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves information about a dedicated host.
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
// hostName - the name of the dedicated host.
// expand - the expand expression to apply on the operation.
func (client DedicatedHostsClient) Get(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, expand InstanceViewTypes) (result DedicatedHost, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, hostGroupName, hostName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DedicatedHostsClient) GetPreparer(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, expand InstanceViewTypes) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(expand)) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) GetResponder(resp *http.Response) (result DedicatedHost, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHostGroup lists all of the dedicated hosts in the specified dedicated host group. Use the nextLink property in
// the response to get the next page of dedicated hosts.
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
func (client DedicatedHostsClient) ListByHostGroup(ctx context.Context, resourceGroupName string, hostGroupName string) (result DedicatedHostListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.ListByHostGroup")
		defer func() {
			sc := -1
			if result.dhlr.Response.Response != nil {
				sc = result.dhlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByHostGroupNextResults
	req, err := client.ListByHostGroupPreparer(ctx, resourceGroupName, hostGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "ListByHostGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHostGroupSender(req)
	if err != nil {
		result.dhlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "ListByHostGroup", resp, "Failure sending request")
		return
	}

	result.dhlr, err = client.ListByHostGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "ListByHostGroup", resp, "Failure responding to request")
	}

	return
}

// ListByHostGroupPreparer prepares the ListByHostGroup request.
func (client DedicatedHostsClient) ListByHostGroupPreparer(ctx context.Context, resourceGroupName string, hostGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHostGroupSender sends the ListByHostGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) ListByHostGroupSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByHostGroupResponder handles the response to the ListByHostGroup request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) ListByHostGroupResponder(resp *http.Response) (result DedicatedHostListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHostGroupNextResults retrieves the next set of results, if any.
func (client DedicatedHostsClient) listByHostGroupNextResults(ctx context.Context, lastResults DedicatedHostListResult) (result DedicatedHostListResult, err error) {
	req, err := lastResults.dedicatedHostListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "listByHostGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHostGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "listByHostGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHostGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "listByHostGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHostGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DedicatedHostsClient) ListByHostGroupComplete(ctx context.Context, resourceGroupName string, hostGroupName string) (result DedicatedHostListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.ListByHostGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHostGroup(ctx, resourceGroupName, hostGroupName)
	return
}

// Update update an dedicated host .
// Parameters:
// resourceGroupName - the name of the resource group.
// hostGroupName - the name of the dedicated host group.
// hostName - the name of the dedicated host .
// parameters - parameters supplied to the Update Dedicated Host operation.
func (client DedicatedHostsClient) Update(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate) (result DedicatedHostsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DedicatedHostsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, hostGroupName, hostName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DedicatedHostsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DedicatedHostsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostGroupName":     autorest.Encode("path", hostGroupName),
		"hostName":          autorest.Encode("path", hostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DedicatedHostsClient) UpdateSender(req *http.Request) (future DedicatedHostsUpdateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DedicatedHostsClient) UpdateResponder(resp *http.Response) (result DedicatedHost, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
