package web

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ManagedHostingEnvironmentsClient is the use these APIs to manage Azure
// Websites resources through the Azure Resource Manager. All task operations
// conform to the HTTP/1.1 protocol specification and each operation returns
// an x-ms-request-id header that can be used to obtain information about the
// request. You must make sure that requests made to these resources are
// secure. For more information, see <a
// href="https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx">Authenticating
// Azure Resource Manager requests.</a>
type ManagedHostingEnvironmentsClient struct {
	ManagementClient
}

// NewManagedHostingEnvironmentsClient creates an instance of the
// ManagedHostingEnvironmentsClient client.
func NewManagedHostingEnvironmentsClient(subscriptionID string) ManagedHostingEnvironmentsClient {
	return NewManagedHostingEnvironmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedHostingEnvironmentsClientWithBaseURI creates an instance of the
// ManagedHostingEnvironmentsClient client.
func NewManagedHostingEnvironmentsClientWithBaseURI(baseURI string, subscriptionID string) ManagedHostingEnvironmentsClient {
	return ManagedHostingEnvironmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdateManagedHostingEnvironment sends the create or update managed
// hosting environment request. This method may poll for completion. Polling
// can be canceled by passing the cancel channel argument. The channel will
// be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment managedHostingEnvironmentEnvelope is properties of managed
// hosting environment
func (client ManagedHostingEnvironmentsClient) CreateOrUpdateManagedHostingEnvironment(resourceGroupName string, name string, managedHostingEnvironmentEnvelope HostingEnvironment, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.CreateOrUpdateManagedHostingEnvironmentPreparer(resourceGroupName, name, managedHostingEnvironmentEnvelope, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "CreateOrUpdateManagedHostingEnvironment", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateManagedHostingEnvironmentSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "CreateOrUpdateManagedHostingEnvironment", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateManagedHostingEnvironmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "CreateOrUpdateManagedHostingEnvironment", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateManagedHostingEnvironmentPreparer prepares the CreateOrUpdateManagedHostingEnvironment request.
func (client ManagedHostingEnvironmentsClient) CreateOrUpdateManagedHostingEnvironmentPreparer(resourceGroupName string, name string, managedHostingEnvironmentEnvelope HostingEnvironment, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}", pathParameters),
		autorest.WithJSON(managedHostingEnvironmentEnvelope),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateManagedHostingEnvironmentSender sends the CreateOrUpdateManagedHostingEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) CreateOrUpdateManagedHostingEnvironmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateManagedHostingEnvironmentResponder handles the response to the CreateOrUpdateManagedHostingEnvironment request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) CreateOrUpdateManagedHostingEnvironmentResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusBadRequest, http.StatusNotFound, http.StatusConflict),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteManagedHostingEnvironment sends the delete managed hosting
// environment request. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used
// to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment forceDelete is delete even if the managed hosting environment
// contains resources
func (client ManagedHostingEnvironmentsClient) DeleteManagedHostingEnvironment(resourceGroupName string, name string, forceDelete *bool, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.DeleteManagedHostingEnvironmentPreparer(resourceGroupName, name, forceDelete, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "DeleteManagedHostingEnvironment", nil, "Failure preparing request")
	}

	resp, err := client.DeleteManagedHostingEnvironmentSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "DeleteManagedHostingEnvironment", resp, "Failure sending request")
	}

	result, err = client.DeleteManagedHostingEnvironmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "DeleteManagedHostingEnvironment", resp, "Failure responding to request")
	}

	return
}

// DeleteManagedHostingEnvironmentPreparer prepares the DeleteManagedHostingEnvironment request.
func (client ManagedHostingEnvironmentsClient) DeleteManagedHostingEnvironmentPreparer(resourceGroupName string, name string, forceDelete *bool, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if forceDelete != nil {
		queryParameters["forceDelete"] = autorest.Encode("query", *forceDelete)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteManagedHostingEnvironmentSender sends the DeleteManagedHostingEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) DeleteManagedHostingEnvironmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteManagedHostingEnvironmentResponder handles the response to the DeleteManagedHostingEnvironment request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) DeleteManagedHostingEnvironmentResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusBadRequest, http.StatusNotFound, http.StatusConflict),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetManagedHostingEnvironment sends the get managed hosting environment
// request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironment(resourceGroupName string, name string) (result ManagedHostingEnvironment, err error) {
	req, err := client.GetManagedHostingEnvironmentPreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironment", nil, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironment", resp, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironment", resp, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentPreparer prepares the GetManagedHostingEnvironment request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetManagedHostingEnvironmentSender sends the GetManagedHostingEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetManagedHostingEnvironmentResponder handles the response to the GetManagedHostingEnvironment request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentResponder(resp *http.Response) (result ManagedHostingEnvironment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironmentOperation sends the get managed hosting
// environment operation request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment operationID is operation identifier GUID
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentOperation(resourceGroupName string, name string, operationID string) (result SetObject, err error) {
	req, err := client.GetManagedHostingEnvironmentOperationPreparer(resourceGroupName, name, operationID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentOperation", nil, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentOperationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentOperation", resp, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentOperationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentOperation", resp, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentOperationPreparer prepares the GetManagedHostingEnvironmentOperation request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentOperationPreparer(resourceGroupName string, name string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}/operations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetManagedHostingEnvironmentOperationSender sends the GetManagedHostingEnvironmentOperation request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentOperationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetManagedHostingEnvironmentOperationResponder handles the response to the GetManagedHostingEnvironmentOperation request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentOperationResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironments sends the get managed hosting environments
// request.
//
// resourceGroupName is name of resource group
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironments(resourceGroupName string) (result HostingEnvironmentCollection, err error) {
	req, err := client.GetManagedHostingEnvironmentsPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironments", nil, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironments", resp, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironments", resp, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentsPreparer prepares the GetManagedHostingEnvironments request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentsPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetManagedHostingEnvironmentsSender sends the GetManagedHostingEnvironments request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetManagedHostingEnvironmentsResponder handles the response to the GetManagedHostingEnvironments request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentsResponder(resp *http.Response) (result HostingEnvironmentCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironmentsNextResults retrieves the next set of results, if any.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentsNextResults(lastResults HostingEnvironmentCollection) (result HostingEnvironmentCollection, err error) {
	req, err := lastResults.HostingEnvironmentCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironments", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetManagedHostingEnvironmentsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironments", resp, "Failure sending next results request")
	}

	result, err = client.GetManagedHostingEnvironmentsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironments", resp, "Failure responding to next results request")
	}

	return
}

// GetManagedHostingEnvironmentServerFarms sends the get managed hosting
// environment server farms request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentServerFarms(resourceGroupName string, name string) (result ServerFarmCollection, err error) {
	req, err := client.GetManagedHostingEnvironmentServerFarmsPreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentServerFarms", nil, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentServerFarmsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentServerFarms", resp, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentServerFarmsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentServerFarms", resp, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentServerFarmsPreparer prepares the GetManagedHostingEnvironmentServerFarms request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentServerFarmsPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}/serverfarms", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetManagedHostingEnvironmentServerFarmsSender sends the GetManagedHostingEnvironmentServerFarms request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentServerFarmsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetManagedHostingEnvironmentServerFarmsResponder handles the response to the GetManagedHostingEnvironmentServerFarms request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentServerFarmsResponder(resp *http.Response) (result ServerFarmCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironmentServerFarmsNextResults retrieves the next set of results, if any.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentServerFarmsNextResults(lastResults ServerFarmCollection) (result ServerFarmCollection, err error) {
	req, err := lastResults.ServerFarmCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentServerFarms", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetManagedHostingEnvironmentServerFarmsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentServerFarms", resp, "Failure sending next results request")
	}

	result, err = client.GetManagedHostingEnvironmentServerFarmsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentServerFarms", resp, "Failure responding to next results request")
	}

	return
}

// GetManagedHostingEnvironmentSites sends the get managed hosting environment
// sites request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment propertiesToInclude is comma separated list of site properties
// to include
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentSites(resourceGroupName string, name string, propertiesToInclude string) (result SiteCollection, err error) {
	req, err := client.GetManagedHostingEnvironmentSitesPreparer(resourceGroupName, name, propertiesToInclude)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentSites", nil, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentSitesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentSites", resp, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentSitesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentSites", resp, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentSitesPreparer prepares the GetManagedHostingEnvironmentSites request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentSitesPreparer(resourceGroupName string, name string, propertiesToInclude string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(propertiesToInclude) > 0 {
		queryParameters["propertiesToInclude"] = autorest.Encode("query", propertiesToInclude)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}/sites", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetManagedHostingEnvironmentSitesSender sends the GetManagedHostingEnvironmentSites request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentSitesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetManagedHostingEnvironmentSitesResponder handles the response to the GetManagedHostingEnvironmentSites request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentSitesResponder(resp *http.Response) (result SiteCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironmentSitesNextResults retrieves the next set of results, if any.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentSitesNextResults(lastResults SiteCollection) (result SiteCollection, err error) {
	req, err := lastResults.SiteCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentSites", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetManagedHostingEnvironmentSitesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentSites", resp, "Failure sending next results request")
	}

	result, err = client.GetManagedHostingEnvironmentSitesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentSites", resp, "Failure responding to next results request")
	}

	return
}

// GetManagedHostingEnvironmentVips sends the get managed hosting environment
// vips request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentVips(resourceGroupName string, name string) (result AddressResponse, err error) {
	req, err := client.GetManagedHostingEnvironmentVipsPreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentVips", nil, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentVipsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentVips", resp, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentVipsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentVips", resp, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentVipsPreparer prepares the GetManagedHostingEnvironmentVips request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentVipsPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}/capacities/virtualip", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetManagedHostingEnvironmentVipsSender sends the GetManagedHostingEnvironmentVips request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentVipsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetManagedHostingEnvironmentVipsResponder handles the response to the GetManagedHostingEnvironmentVips request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentVipsResponder(resp *http.Response) (result AddressResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironmentWebHostingPlans sends the get managed hosting
// environment web hosting plans request.
//
// resourceGroupName is name of resource group name is name of managed hosting
// environment
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentWebHostingPlans(resourceGroupName string, name string) (result ServerFarmCollection, err error) {
	req, err := client.GetManagedHostingEnvironmentWebHostingPlansPreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentWebHostingPlans", nil, "Failure preparing request")
	}

	resp, err := client.GetManagedHostingEnvironmentWebHostingPlansSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentWebHostingPlans", resp, "Failure sending request")
	}

	result, err = client.GetManagedHostingEnvironmentWebHostingPlansResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentWebHostingPlans", resp, "Failure responding to request")
	}

	return
}

// GetManagedHostingEnvironmentWebHostingPlansPreparer prepares the GetManagedHostingEnvironmentWebHostingPlans request.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentWebHostingPlansPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}/webhostingplans", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetManagedHostingEnvironmentWebHostingPlansSender sends the GetManagedHostingEnvironmentWebHostingPlans request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentWebHostingPlansSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetManagedHostingEnvironmentWebHostingPlansResponder handles the response to the GetManagedHostingEnvironmentWebHostingPlans request. The method always
// closes the http.Response Body.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentWebHostingPlansResponder(resp *http.Response) (result ServerFarmCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetManagedHostingEnvironmentWebHostingPlansNextResults retrieves the next set of results, if any.
func (client ManagedHostingEnvironmentsClient) GetManagedHostingEnvironmentWebHostingPlansNextResults(lastResults ServerFarmCollection) (result ServerFarmCollection, err error) {
	req, err := lastResults.ServerFarmCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentWebHostingPlans", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetManagedHostingEnvironmentWebHostingPlansSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentWebHostingPlans", resp, "Failure sending next results request")
	}

	result, err = client.GetManagedHostingEnvironmentWebHostingPlansResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ManagedHostingEnvironmentsClient", "GetManagedHostingEnvironmentWebHostingPlans", resp, "Failure responding to next results request")
	}

	return
}
