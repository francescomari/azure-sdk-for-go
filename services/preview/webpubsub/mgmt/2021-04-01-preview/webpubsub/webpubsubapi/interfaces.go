package webpubsubapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/webpubsub/mgmt/2021-04-01-preview/webpubsub"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result webpubsub.OperationListPage, err error)
	ListComplete(ctx context.Context) (result webpubsub.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*webpubsub.OperationsClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, parameters webpubsub.NameAvailabilityParameters) (result webpubsub.NameAvailability, err error)
	CreateOrUpdate(ctx context.Context, parameters webpubsub.ResourceType, resourceGroupName string, resourceName string) (result webpubsub.CreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result webpubsub.DeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result webpubsub.ResourceType, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result webpubsub.ResourceListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result webpubsub.ResourceListIterator, err error)
	ListBySubscription(ctx context.Context) (result webpubsub.ResourceListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result webpubsub.ResourceListIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, resourceName string) (result webpubsub.Keys, err error)
	RegenerateKey(ctx context.Context, parameters webpubsub.RegenerateKeyParameters, resourceGroupName string, resourceName string) (result webpubsub.RegenerateKeyFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, resourceName string) (result webpubsub.RestartFuture, err error)
	Update(ctx context.Context, parameters webpubsub.ResourceType, resourceGroupName string, resourceName string) (result webpubsub.UpdateFuture, err error)
}

var _ ClientAPI = (*webpubsub.Client)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, location string) (result webpubsub.SignalRServiceUsageListPage, err error)
	ListComplete(ctx context.Context, location string) (result webpubsub.SignalRServiceUsageListIterator, err error)
}

var _ UsagesClientAPI = (*webpubsub.UsagesClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Delete(ctx context.Context, privateEndpointConnectionName string, resourceGroupName string, resourceName string) (result webpubsub.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, privateEndpointConnectionName string, resourceGroupName string, resourceName string) (result webpubsub.PrivateEndpointConnection, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result webpubsub.PrivateEndpointConnectionListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result webpubsub.PrivateEndpointConnectionListIterator, err error)
	Update(ctx context.Context, privateEndpointConnectionName string, parameters webpubsub.PrivateEndpointConnection, resourceGroupName string, resourceName string) (result webpubsub.PrivateEndpointConnection, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*webpubsub.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, resourceName string) (result webpubsub.PrivateLinkResourceListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result webpubsub.PrivateLinkResourceListIterator, err error)
}

var _ PrivateLinkResourcesClientAPI = (*webpubsub.PrivateLinkResourcesClient)(nil)

// SharedPrivateLinkResourcesClientAPI contains the set of methods on the SharedPrivateLinkResourcesClient type.
type SharedPrivateLinkResourcesClientAPI interface {
	CreateOrUpdate(ctx context.Context, sharedPrivateLinkResourceName string, parameters webpubsub.SharedPrivateLinkResource, resourceGroupName string, resourceName string) (result webpubsub.SharedPrivateLinkResourcesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string) (result webpubsub.SharedPrivateLinkResourcesDeleteFuture, err error)
	Get(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string) (result webpubsub.SharedPrivateLinkResource, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result webpubsub.SharedPrivateLinkResourceListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result webpubsub.SharedPrivateLinkResourceListIterator, err error)
}

var _ SharedPrivateLinkResourcesClientAPI = (*webpubsub.SharedPrivateLinkResourcesClient)(nil)
