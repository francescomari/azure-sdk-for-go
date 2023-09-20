//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/communityGalleryExamples/CommunityGalleryImageVersion_Get.json
func ExampleCommunityGalleryImageVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommunityGalleryImageVersionsClient().Get(ctx, "myLocation", "publicGalleryName", "myGalleryImageName", "myGalleryImageVersionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommunityGalleryImageVersion = armcompute.CommunityGalleryImageVersion{
	// 	Name: to.Ptr("myGalleryImageVersionName"),
	// 	Type: to.Ptr("Microsoft.Compute/Locations/CommunityGalleryImageVersion"),
	// 	Identifier: &armcompute.CommunityGalleryIdentifier{
	// 		UniqueID: to.Ptr("/CommunityGalleries/publicGalleryName/Images/myGalleryImageName/Versions/myGalleryImageVersionName"),
	// 	},
	// 	Location: to.Ptr("myLocation"),
	// 	Properties: &armcompute.CommunityGalleryImageVersionProperties{
	// 		EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T09:12:28Z"); return t}()),
	// 		ExcludeFromLatest: to.Ptr(false),
	// 		PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-20T09:12:28Z"); return t}()),
	// 		StorageProfile: &armcompute.SharedGalleryImageVersionStorageProfile{
	// 			OSDiskImage: &armcompute.SharedGalleryOSDiskImage{
	// 				DiskSizeGB: to.Ptr[int32](29),
	// 				HostCaching: to.Ptr(armcompute.SharedGalleryHostCachingNone),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/communityGalleryExamples/CommunityGalleryImageVersion_List.json
func ExampleCommunityGalleryImageVersionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommunityGalleryImageVersionsClient().NewListPager("myLocation", "publicGalleryName", "myGalleryImageName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.CommunityGalleryImageVersionList = armcompute.CommunityGalleryImageVersionList{
		// 	Value: []*armcompute.CommunityGalleryImageVersion{
		// 		{
		// 			Name: to.Ptr("myGalleryImageVersionName"),
		// 			Identifier: &armcompute.CommunityGalleryIdentifier{
		// 				UniqueID: to.Ptr("/CommunityGalleries/publicGalleryName/Images/myGalleryImageName/Versions/myGalleryImageVersionName"),
		// 			},
		// 			Location: to.Ptr("myLocation"),
		// 			Properties: &armcompute.CommunityGalleryImageVersionProperties{
		// 				EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T09:12:28Z"); return t}()),
		// 				ExcludeFromLatest: to.Ptr(false),
		// 				PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-20T09:12:28Z"); return t}()),
		// 				StorageProfile: &armcompute.SharedGalleryImageVersionStorageProfile{
		// 					OSDiskImage: &armcompute.SharedGalleryOSDiskImage{
		// 						DiskSizeGB: to.Ptr[int32](29),
		// 						HostCaching: to.Ptr(armcompute.SharedGalleryHostCachingNone),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
