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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/restorePointExamples/RestorePoint_Copy_BetweenRegions.json
func ExampleRestorePointsClient_BeginCreate_copyARestorePointToADifferentRegion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRestorePointsClient().BeginCreate(ctx, "myResourceGroup", "rpcName", "rpName", armcompute.RestorePoint{
		Properties: &armcompute.RestorePointProperties{
			SourceRestorePoint: &armcompute.APIEntityReference{
				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/sourceRpcName/restorePoints/sourceRpName"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/restorePointExamples/RestorePoint_Create.json
func ExampleRestorePointsClient_BeginCreate_createARestorePoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRestorePointsClient().BeginCreate(ctx, "myResourceGroup", "rpcName", "rpName", armcompute.RestorePoint{
		Properties: &armcompute.RestorePointProperties{
			ExcludeDisks: []*armcompute.APIEntityReference{
				{
					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/restorePointExamples/RestorePoint_Delete_MaximumSet_Gen.json
func ExampleRestorePointsClient_BeginDelete_restorePointDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRestorePointsClient().BeginDelete(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaa", "a", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/restorePointExamples/RestorePoint_Delete_MinimumSet_Gen.json
func ExampleRestorePointsClient_BeginDelete_restorePointDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRestorePointsClient().BeginDelete(ctx, "rgcompute", "aaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/restorePointExamples/RestorePoint_Get.json
func ExampleRestorePointsClient_Get_getARestorePoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointsClient().Get(ctx, "myResourceGroup", "rpcName", "rpName", &armcompute.RestorePointsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePoint = armcompute.RestorePoint{
	// 	Name: to.Ptr("rpName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/rpName"),
	// 	Properties: &armcompute.RestorePointProperties{
	// 		ConsistencyMode: to.Ptr(armcompute.ConsistencyModeTypesApplicationConsistent),
	// 		ExcludeDisks: []*armcompute.APIEntityReference{
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm8768_disk2_fe6ffde4f69b491ca33fb984d5bcd89f"),
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SourceMetadata: &armcompute.RestorePointSourceMetadata{
	// 			DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 				BootDiagnostics: &armcompute.BootDiagnostics{
	// 					Enabled: to.Ptr(true),
	// 				},
	// 			},
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardB1S),
	// 			},
	// 			Location: to.Ptr("westus"),
	// 			OSProfile: &armcompute.OSProfile{
	// 				AdminUsername: to.Ptr("admin"),
	// 				AllowExtensionOperations: to.Ptr(true),
	// 				ComputerName: to.Ptr("computerName"),
	// 				RequireGuestProvisionSignal: to.Ptr(true),
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 					EnableAutomaticUpdates: to.Ptr(true),
	// 					ProvisionVMAgent: to.Ptr(true),
	// 				},
	// 			},
	// 			StorageProfile: &armcompute.RestorePointSourceVMStorageProfile{
	// 				DataDisks: []*armcompute.RestorePointSourceVMDataDisk{
	// 					{
	// 						Name: to.Ptr("testingexcludedisk_DataDisk_1"),
	// 						Caching: to.Ptr(armcompute.CachingTypesNone),
	// 						DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/userdata/providers/Microsoft.Compute/restorePointCollections/mynewrpc/restorePoints/restorepointtwo/diskRestorePoints/testingexcludedisk_DataDisk_1_68785190-1acb-4d5e-a8ae-705b45f3dca5"),
	// 						},
	// 						Lun: to.Ptr[int32](1),
	// 						ManagedDisk: &armcompute.ManagedDiskParameters{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/userdata/providers/Microsoft.Compute/disks/testingexcludedisk_DataDisk_1"),
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 						},
	// 				}},
	// 				OSDisk: &armcompute.RestorePointSourceVMOSDisk{
	// 					Name: to.Ptr("testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/rpName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
	// 					},
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypeWindows),
	// 				},
	// 			},
	// 			VMID: to.Ptr("76d6541e-80bd-4dc1-932b-3cae4cfb80e7"),
	// 		},
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-27T20:35:05.8401519+00:00"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/restorePointExamples/RestorePoint_Get_WithInstanceView.json
func ExampleRestorePointsClient_Get_getRestorePointWithInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointsClient().Get(ctx, "myResourceGroup", "rpcName", "rpName", &armcompute.RestorePointsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePoint = armcompute.RestorePoint{
	// 	Name: to.Ptr("rpName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/rpName"),
	// 	Properties: &armcompute.RestorePointProperties{
	// 		ConsistencyMode: to.Ptr(armcompute.ConsistencyModeTypesApplicationConsistent),
	// 		ExcludeDisks: []*armcompute.APIEntityReference{
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm8768_disk2_fe6ffde4f69b491ca33fb984d5bcd89f"),
	// 		}},
	// 		InstanceView: &armcompute.RestorePointInstanceView{
	// 			DiskRestorePoints: []*armcompute.DiskRestorePointInstanceView{
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/rpName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
	// 					ReplicationStatus: &armcompute.DiskRestorePointReplicationStatus{
	// 						CompletionPercent: to.Ptr[int32](100),
	// 						Status: &armcompute.InstanceViewStatus{
	// 							Code: to.Ptr("ReplicationState/succeeded"),
	// 							DisplayStatus: to.Ptr("Succeeded"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/userdata/providers/Microsoft.Compute/restorePointCollections/mynewrpc/restorePoints/restorepointtwo/diskRestorePoints/testingexcludedisk_DataDisk_1_68785190-1acb-4d5e-a8ae-705b45f3dca5"),
	// 					ReplicationStatus: &armcompute.DiskRestorePointReplicationStatus{
	// 						CompletionPercent: to.Ptr[int32](100),
	// 						Status: &armcompute.InstanceViewStatus{
	// 							Code: to.Ptr("ReplicationState/succeeded"),
	// 							DisplayStatus: to.Ptr("Succeeded"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						},
	// 					},
	// 			}},
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("ReplicationState/succeeded"),
	// 					DisplayStatus: to.Ptr("Succeeded"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 			}},
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SourceMetadata: &armcompute.RestorePointSourceMetadata{
	// 			DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 				BootDiagnostics: &armcompute.BootDiagnostics{
	// 					Enabled: to.Ptr(true),
	// 				},
	// 			},
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardB1S),
	// 			},
	// 			Location: to.Ptr("westus"),
	// 			OSProfile: &armcompute.OSProfile{
	// 				AdminUsername: to.Ptr("admin"),
	// 				AllowExtensionOperations: to.Ptr(true),
	// 				ComputerName: to.Ptr("computerName"),
	// 				RequireGuestProvisionSignal: to.Ptr(true),
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 					EnableAutomaticUpdates: to.Ptr(true),
	// 					ProvisionVMAgent: to.Ptr(true),
	// 				},
	// 			},
	// 			StorageProfile: &armcompute.RestorePointSourceVMStorageProfile{
	// 				DataDisks: []*armcompute.RestorePointSourceVMDataDisk{
	// 					{
	// 						Name: to.Ptr("testingexcludedisk_DataDisk_1"),
	// 						Caching: to.Ptr(armcompute.CachingTypesNone),
	// 						DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/userdata/providers/Microsoft.Compute/restorePointCollections/mynewrpc/restorePoints/restorepointtwo/diskRestorePoints/testingexcludedisk_DataDisk_1_68785190-1acb-4d5e-a8ae-705b45f3dca5"),
	// 						},
	// 						Lun: to.Ptr[int32](1),
	// 						ManagedDisk: &armcompute.ManagedDiskParameters{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/userdata/providers/Microsoft.Compute/disks/testingexcludedisk_DataDisk_1"),
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 						},
	// 				}},
	// 				OSDisk: &armcompute.RestorePointSourceVMOSDisk{
	// 					Name: to.Ptr("testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/rpName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
	// 					},
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypeWindows),
	// 				},
	// 			},
	// 			VMID: to.Ptr("76d6541e-80bd-4dc1-932b-3cae4cfb80e7"),
	// 		},
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-27T20:35:05.8401519+00:00"); return t}()),
	// 	},
	// }
}
