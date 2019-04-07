// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: dab57ee609fffdc578f48519d5cdc980efd8cc00

package compute

import original "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-03-30/compute"

type VirtualMachineScaleSetsClient = original.VirtualMachineScaleSetsClient
type SnapshotsClient = original.SnapshotsClient
type UsageClient = original.UsageClient
type VirtualMachineExtensionImagesClient = original.VirtualMachineExtensionImagesClient
type VirtualMachineImagesClient = original.VirtualMachineImagesClient
type VirtualMachineSizesClient = original.VirtualMachineSizesClient
type AvailabilitySetsClient = original.AvailabilitySetsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type VirtualMachineExtensionsClient = original.VirtualMachineExtensionsClient
type VirtualMachineScaleSetExtensionsClient = original.VirtualMachineScaleSetExtensionsClient
type VirtualMachineScaleSetVMsClient = original.VirtualMachineScaleSetVMsClient
type DisksClient = original.DisksClient
type ResourceSkusClient = original.ResourceSkusClient
type VirtualMachineRunCommandsClient = original.VirtualMachineRunCommandsClient
type VirtualMachinesClient = original.VirtualMachinesClient
type ImagesClient = original.ImagesClient
type AccessLevel = original.AccessLevel

const (
	None AccessLevel = original.None
	Read AccessLevel = original.Read
)

type CachingTypes = original.CachingTypes

const (
	CachingTypesNone      CachingTypes = original.CachingTypesNone
	CachingTypesReadOnly  CachingTypes = original.CachingTypesReadOnly
	CachingTypesReadWrite CachingTypes = original.CachingTypesReadWrite
)

type ComponentNames = original.ComponentNames

const (
	MicrosoftWindowsShellSetup ComponentNames = original.MicrosoftWindowsShellSetup
)

type DiskCreateOption = original.DiskCreateOption

const (
	Attach    DiskCreateOption = original.Attach
	Copy      DiskCreateOption = original.Copy
	Empty     DiskCreateOption = original.Empty
	FromImage DiskCreateOption = original.FromImage
	Import    DiskCreateOption = original.Import
)

type DiskCreateOptionTypes = original.DiskCreateOptionTypes

const (
	DiskCreateOptionTypesAttach    DiskCreateOptionTypes = original.DiskCreateOptionTypesAttach
	DiskCreateOptionTypesEmpty     DiskCreateOptionTypes = original.DiskCreateOptionTypesEmpty
	DiskCreateOptionTypesFromImage DiskCreateOptionTypes = original.DiskCreateOptionTypesFromImage
)

type InstanceViewTypes = original.InstanceViewTypes

const (
	InstanceView InstanceViewTypes = original.InstanceView
)

type IPVersion = original.IPVersion

const (
	IPv4 IPVersion = original.IPv4
	IPv6 IPVersion = original.IPv6
)

type MaintenanceOperationResultCodeTypes = original.MaintenanceOperationResultCodeTypes

const (
	MaintenanceOperationResultCodeTypesMaintenanceAborted   MaintenanceOperationResultCodeTypes = original.MaintenanceOperationResultCodeTypesMaintenanceAborted
	MaintenanceOperationResultCodeTypesMaintenanceCompleted MaintenanceOperationResultCodeTypes = original.MaintenanceOperationResultCodeTypesMaintenanceCompleted
	MaintenanceOperationResultCodeTypesNone                 MaintenanceOperationResultCodeTypes = original.MaintenanceOperationResultCodeTypesNone
	MaintenanceOperationResultCodeTypesRetryLater           MaintenanceOperationResultCodeTypes = original.MaintenanceOperationResultCodeTypesRetryLater
)

type OperatingSystemStateTypes = original.OperatingSystemStateTypes

const (
	Generalized OperatingSystemStateTypes = original.Generalized
	Specialized OperatingSystemStateTypes = original.Specialized
)

type OperatingSystemTypes = original.OperatingSystemTypes

const (
	Linux   OperatingSystemTypes = original.Linux
	Windows OperatingSystemTypes = original.Windows
)

type PassNames = original.PassNames

const (
	OobeSystem PassNames = original.OobeSystem
)

type ProtocolTypes = original.ProtocolTypes

const (
	HTTP  ProtocolTypes = original.HTTP
	HTTPS ProtocolTypes = original.HTTPS
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleType

const (
	ResourceSkuCapacityScaleTypeAutomatic ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleTypeAutomatic
	ResourceSkuCapacityScaleTypeManual    ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleTypeManual
	ResourceSkuCapacityScaleTypeNone      ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleTypeNone
)

type ResourceSkuRestrictionsReasonCode = original.ResourceSkuRestrictionsReasonCode

const (
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = original.NotAvailableForSubscription
	QuotaID                     ResourceSkuRestrictionsReasonCode = original.QuotaID
)

type ResourceSkuRestrictionsType = original.ResourceSkuRestrictionsType

const (
	Location ResourceSkuRestrictionsType = original.Location
)

type RollingUpgradeActionType = original.RollingUpgradeActionType

const (
	Cancel RollingUpgradeActionType = original.Cancel
	Start  RollingUpgradeActionType = original.Start
)

type RollingUpgradeStatusCode = original.RollingUpgradeStatusCode

const (
	Cancelled      RollingUpgradeStatusCode = original.Cancelled
	Completed      RollingUpgradeStatusCode = original.Completed
	Faulted        RollingUpgradeStatusCode = original.Faulted
	RollingForward RollingUpgradeStatusCode = original.RollingForward
)

type SettingNames = original.SettingNames

const (
	AutoLogon          SettingNames = original.AutoLogon
	FirstLogonCommands SettingNames = original.FirstLogonCommands
)

type StatusLevelTypes = original.StatusLevelTypes

const (
	Error   StatusLevelTypes = original.Error
	Info    StatusLevelTypes = original.Info
	Warning StatusLevelTypes = original.Warning
)

type StorageAccountTypes = original.StorageAccountTypes

const (
	PremiumLRS  StorageAccountTypes = original.PremiumLRS
	StandardLRS StorageAccountTypes = original.StandardLRS
)

type UpgradeMode = original.UpgradeMode

const (
	Automatic UpgradeMode = original.Automatic
	Manual    UpgradeMode = original.Manual
	Rolling   UpgradeMode = original.Rolling
)

type VirtualMachineScaleSetSkuScaleType = original.VirtualMachineScaleSetSkuScaleType

const (
	VirtualMachineScaleSetSkuScaleTypeAutomatic VirtualMachineScaleSetSkuScaleType = original.VirtualMachineScaleSetSkuScaleTypeAutomatic
	VirtualMachineScaleSetSkuScaleTypeNone      VirtualMachineScaleSetSkuScaleType = original.VirtualMachineScaleSetSkuScaleTypeNone
)

type VirtualMachineSizeTypes = original.VirtualMachineSizeTypes

const (
	BasicA0        VirtualMachineSizeTypes = original.BasicA0
	BasicA1        VirtualMachineSizeTypes = original.BasicA1
	BasicA2        VirtualMachineSizeTypes = original.BasicA2
	BasicA3        VirtualMachineSizeTypes = original.BasicA3
	BasicA4        VirtualMachineSizeTypes = original.BasicA4
	StandardA0     VirtualMachineSizeTypes = original.StandardA0
	StandardA1     VirtualMachineSizeTypes = original.StandardA1
	StandardA10    VirtualMachineSizeTypes = original.StandardA10
	StandardA11    VirtualMachineSizeTypes = original.StandardA11
	StandardA1V2   VirtualMachineSizeTypes = original.StandardA1V2
	StandardA2     VirtualMachineSizeTypes = original.StandardA2
	StandardA2mV2  VirtualMachineSizeTypes = original.StandardA2mV2
	StandardA2V2   VirtualMachineSizeTypes = original.StandardA2V2
	StandardA3     VirtualMachineSizeTypes = original.StandardA3
	StandardA4     VirtualMachineSizeTypes = original.StandardA4
	StandardA4mV2  VirtualMachineSizeTypes = original.StandardA4mV2
	StandardA4V2   VirtualMachineSizeTypes = original.StandardA4V2
	StandardA5     VirtualMachineSizeTypes = original.StandardA5
	StandardA6     VirtualMachineSizeTypes = original.StandardA6
	StandardA7     VirtualMachineSizeTypes = original.StandardA7
	StandardA8     VirtualMachineSizeTypes = original.StandardA8
	StandardA8mV2  VirtualMachineSizeTypes = original.StandardA8mV2
	StandardA8V2   VirtualMachineSizeTypes = original.StandardA8V2
	StandardA9     VirtualMachineSizeTypes = original.StandardA9
	StandardD1     VirtualMachineSizeTypes = original.StandardD1
	StandardD11    VirtualMachineSizeTypes = original.StandardD11
	StandardD11V2  VirtualMachineSizeTypes = original.StandardD11V2
	StandardD12    VirtualMachineSizeTypes = original.StandardD12
	StandardD12V2  VirtualMachineSizeTypes = original.StandardD12V2
	StandardD13    VirtualMachineSizeTypes = original.StandardD13
	StandardD13V2  VirtualMachineSizeTypes = original.StandardD13V2
	StandardD14    VirtualMachineSizeTypes = original.StandardD14
	StandardD14V2  VirtualMachineSizeTypes = original.StandardD14V2
	StandardD15V2  VirtualMachineSizeTypes = original.StandardD15V2
	StandardD1V2   VirtualMachineSizeTypes = original.StandardD1V2
	StandardD2     VirtualMachineSizeTypes = original.StandardD2
	StandardD2V2   VirtualMachineSizeTypes = original.StandardD2V2
	StandardD3     VirtualMachineSizeTypes = original.StandardD3
	StandardD3V2   VirtualMachineSizeTypes = original.StandardD3V2
	StandardD4     VirtualMachineSizeTypes = original.StandardD4
	StandardD4V2   VirtualMachineSizeTypes = original.StandardD4V2
	StandardD5V2   VirtualMachineSizeTypes = original.StandardD5V2
	StandardDS1    VirtualMachineSizeTypes = original.StandardDS1
	StandardDS11   VirtualMachineSizeTypes = original.StandardDS11
	StandardDS11V2 VirtualMachineSizeTypes = original.StandardDS11V2
	StandardDS12   VirtualMachineSizeTypes = original.StandardDS12
	StandardDS12V2 VirtualMachineSizeTypes = original.StandardDS12V2
	StandardDS13   VirtualMachineSizeTypes = original.StandardDS13
	StandardDS13V2 VirtualMachineSizeTypes = original.StandardDS13V2
	StandardDS14   VirtualMachineSizeTypes = original.StandardDS14
	StandardDS14V2 VirtualMachineSizeTypes = original.StandardDS14V2
	StandardDS15V2 VirtualMachineSizeTypes = original.StandardDS15V2
	StandardDS1V2  VirtualMachineSizeTypes = original.StandardDS1V2
	StandardDS2    VirtualMachineSizeTypes = original.StandardDS2
	StandardDS2V2  VirtualMachineSizeTypes = original.StandardDS2V2
	StandardDS3    VirtualMachineSizeTypes = original.StandardDS3
	StandardDS3V2  VirtualMachineSizeTypes = original.StandardDS3V2
	StandardDS4    VirtualMachineSizeTypes = original.StandardDS4
	StandardDS4V2  VirtualMachineSizeTypes = original.StandardDS4V2
	StandardDS5V2  VirtualMachineSizeTypes = original.StandardDS5V2
	StandardF1     VirtualMachineSizeTypes = original.StandardF1
	StandardF16    VirtualMachineSizeTypes = original.StandardF16
	StandardF16s   VirtualMachineSizeTypes = original.StandardF16s
	StandardF1s    VirtualMachineSizeTypes = original.StandardF1s
	StandardF2     VirtualMachineSizeTypes = original.StandardF2
	StandardF2s    VirtualMachineSizeTypes = original.StandardF2s
	StandardF4     VirtualMachineSizeTypes = original.StandardF4
	StandardF4s    VirtualMachineSizeTypes = original.StandardF4s
	StandardF8     VirtualMachineSizeTypes = original.StandardF8
	StandardF8s    VirtualMachineSizeTypes = original.StandardF8s
	StandardG1     VirtualMachineSizeTypes = original.StandardG1
	StandardG2     VirtualMachineSizeTypes = original.StandardG2
	StandardG3     VirtualMachineSizeTypes = original.StandardG3
	StandardG4     VirtualMachineSizeTypes = original.StandardG4
	StandardG5     VirtualMachineSizeTypes = original.StandardG5
	StandardGS1    VirtualMachineSizeTypes = original.StandardGS1
	StandardGS2    VirtualMachineSizeTypes = original.StandardGS2
	StandardGS3    VirtualMachineSizeTypes = original.StandardGS3
	StandardGS4    VirtualMachineSizeTypes = original.StandardGS4
	StandardGS5    VirtualMachineSizeTypes = original.StandardGS5
	StandardH16    VirtualMachineSizeTypes = original.StandardH16
	StandardH16m   VirtualMachineSizeTypes = original.StandardH16m
	StandardH16mr  VirtualMachineSizeTypes = original.StandardH16mr
	StandardH16r   VirtualMachineSizeTypes = original.StandardH16r
	StandardH8     VirtualMachineSizeTypes = original.StandardH8
	StandardH8m    VirtualMachineSizeTypes = original.StandardH8m
	StandardL16s   VirtualMachineSizeTypes = original.StandardL16s
	StandardL32s   VirtualMachineSizeTypes = original.StandardL32s
	StandardL4s    VirtualMachineSizeTypes = original.StandardL4s
	StandardL8s    VirtualMachineSizeTypes = original.StandardL8s
	StandardNC12   VirtualMachineSizeTypes = original.StandardNC12
	StandardNC24   VirtualMachineSizeTypes = original.StandardNC24
	StandardNC24r  VirtualMachineSizeTypes = original.StandardNC24r
	StandardNC6    VirtualMachineSizeTypes = original.StandardNC6
	StandardNV12   VirtualMachineSizeTypes = original.StandardNV12
	StandardNV24   VirtualMachineSizeTypes = original.StandardNV24
	StandardNV6    VirtualMachineSizeTypes = original.StandardNV6
)

type AccessURI = original.AccessURI
type AccessURIOutput = original.AccessURIOutput
type AccessURIRaw = original.AccessURIRaw
type AdditionalUnattendContent = original.AdditionalUnattendContent
type APIEntityReference = original.APIEntityReference
type APIError = original.APIError
type APIErrorBase = original.APIErrorBase
type AvailabilitySet = original.AvailabilitySet
type AvailabilitySetListResult = original.AvailabilitySetListResult
type AvailabilitySetProperties = original.AvailabilitySetProperties
type BootDiagnostics = original.BootDiagnostics
type BootDiagnosticsInstanceView = original.BootDiagnosticsInstanceView
type CreationData = original.CreationData
type DataDisk = original.DataDisk
type DataDiskImage = original.DataDiskImage
type DiagnosticsProfile = original.DiagnosticsProfile
type Disk = original.Disk
type DiskEncryptionSettings = original.DiskEncryptionSettings
type DiskInstanceView = original.DiskInstanceView
type DiskList = original.DiskList
type DiskProperties = original.DiskProperties
type DiskSku = original.DiskSku
type DiskUpdate = original.DiskUpdate
type DiskUpdateProperties = original.DiskUpdateProperties
type EncryptionSettings = original.EncryptionSettings
type GrantAccessData = original.GrantAccessData
type HardwareProfile = original.HardwareProfile
type Image = original.Image
type ImageDataDisk = original.ImageDataDisk
type ImageDiskReference = original.ImageDiskReference
type ImageListResult = original.ImageListResult
type ImageOSDisk = original.ImageOSDisk
type ImageProperties = original.ImageProperties
type ImageReference = original.ImageReference
type ImageStorageProfile = original.ImageStorageProfile
type InnerError = original.InnerError
type InstanceViewStatus = original.InstanceViewStatus
type KeyVaultAndKeyReference = original.KeyVaultAndKeyReference
type KeyVaultAndSecretReference = original.KeyVaultAndSecretReference
type KeyVaultKeyReference = original.KeyVaultKeyReference
type KeyVaultSecretReference = original.KeyVaultSecretReference
type LinuxConfiguration = original.LinuxConfiguration
type ListUsagesResult = original.ListUsagesResult
type ListVirtualMachineExtensionImage = original.ListVirtualMachineExtensionImage
type ListVirtualMachineImageResource = original.ListVirtualMachineImageResource
type LongRunningOperationProperties = original.LongRunningOperationProperties
type MaintenanceRedeployStatus = original.MaintenanceRedeployStatus
type ManagedDiskParameters = original.ManagedDiskParameters
type NetworkInterfaceReference = original.NetworkInterfaceReference
type NetworkInterfaceReferenceProperties = original.NetworkInterfaceReferenceProperties
type NetworkProfile = original.NetworkProfile
type OperationStatusResponse = original.OperationStatusResponse
type OSDisk = original.OSDisk
type OSDiskImage = original.OSDiskImage
type OSProfile = original.OSProfile
type Plan = original.Plan
type PurchasePlan = original.PurchasePlan
type Resource = original.Resource
type ResourceSku = original.ResourceSku
type ResourceSkuCapabilities = original.ResourceSkuCapabilities
type ResourceSkuCapacity = original.ResourceSkuCapacity
type ResourceSkuCosts = original.ResourceSkuCosts
type ResourceSkuRestrictions = original.ResourceSkuRestrictions
type ResourceSkusResult = original.ResourceSkusResult
type ResourceUpdate = original.ResourceUpdate
type RollingUpgradePolicy = original.RollingUpgradePolicy
type RollingUpgradeProgressInfo = original.RollingUpgradeProgressInfo
type RollingUpgradeRunningStatus = original.RollingUpgradeRunningStatus
type RollingUpgradeStatusInfo = original.RollingUpgradeStatusInfo
type RollingUpgradeStatusInfoProperties = original.RollingUpgradeStatusInfoProperties
type RunCommandDocument = original.RunCommandDocument
type RunCommandDocumentBase = original.RunCommandDocumentBase
type RunCommandInput = original.RunCommandInput
type RunCommandInputParameter = original.RunCommandInputParameter
type RunCommandListResult = original.RunCommandListResult
type RunCommandParameterDefinition = original.RunCommandParameterDefinition
type RunCommandResult = original.RunCommandResult
type RunCommandResultProperties = original.RunCommandResultProperties
type Sku = original.Sku
type Snapshot = original.Snapshot
type SnapshotList = original.SnapshotList
type SnapshotUpdate = original.SnapshotUpdate
type SourceVault = original.SourceVault
type SSHConfiguration = original.SSHConfiguration
type SSHPublicKey = original.SSHPublicKey
type StorageProfile = original.StorageProfile
type SubResource = original.SubResource
type SubResourceReadOnly = original.SubResourceReadOnly
type UpdateResource = original.UpdateResource
type UpgradePolicy = original.UpgradePolicy
type Usage = original.Usage
type UsageName = original.UsageName
type VaultCertificate = original.VaultCertificate
type VaultSecretGroup = original.VaultSecretGroup
type VirtualHardDisk = original.VirtualHardDisk
type VirtualMachine = original.VirtualMachine
type VirtualMachineAgentInstanceView = original.VirtualMachineAgentInstanceView
type VirtualMachineCaptureParameters = original.VirtualMachineCaptureParameters
type VirtualMachineCaptureResult = original.VirtualMachineCaptureResult
type VirtualMachineCaptureResultProperties = original.VirtualMachineCaptureResultProperties
type VirtualMachineExtension = original.VirtualMachineExtension
type VirtualMachineExtensionHandlerInstanceView = original.VirtualMachineExtensionHandlerInstanceView
type VirtualMachineExtensionImage = original.VirtualMachineExtensionImage
type VirtualMachineExtensionImageProperties = original.VirtualMachineExtensionImageProperties
type VirtualMachineExtensionInstanceView = original.VirtualMachineExtensionInstanceView
type VirtualMachineExtensionProperties = original.VirtualMachineExtensionProperties
type VirtualMachineHealthStatus = original.VirtualMachineHealthStatus
type VirtualMachineIdentity = original.VirtualMachineIdentity
type VirtualMachineImage = original.VirtualMachineImage
type VirtualMachineImageProperties = original.VirtualMachineImageProperties
type VirtualMachineImageResource = original.VirtualMachineImageResource
type VirtualMachineInstanceView = original.VirtualMachineInstanceView
type VirtualMachineListResult = original.VirtualMachineListResult
type VirtualMachineProperties = original.VirtualMachineProperties
type VirtualMachineScaleSet = original.VirtualMachineScaleSet
type VirtualMachineScaleSetDataDisk = original.VirtualMachineScaleSetDataDisk
type VirtualMachineScaleSetExtension = original.VirtualMachineScaleSetExtension
type VirtualMachineScaleSetExtensionListResult = original.VirtualMachineScaleSetExtensionListResult
type VirtualMachineScaleSetExtensionProfile = original.VirtualMachineScaleSetExtensionProfile
type VirtualMachineScaleSetExtensionProperties = original.VirtualMachineScaleSetExtensionProperties
type VirtualMachineScaleSetIdentity = original.VirtualMachineScaleSetIdentity
type VirtualMachineScaleSetInstanceView = original.VirtualMachineScaleSetInstanceView
type VirtualMachineScaleSetInstanceViewStatusesSummary = original.VirtualMachineScaleSetInstanceViewStatusesSummary
type VirtualMachineScaleSetIPConfiguration = original.VirtualMachineScaleSetIPConfiguration
type VirtualMachineScaleSetIPConfigurationProperties = original.VirtualMachineScaleSetIPConfigurationProperties
type VirtualMachineScaleSetListResult = original.VirtualMachineScaleSetListResult
type VirtualMachineScaleSetListSkusResult = original.VirtualMachineScaleSetListSkusResult
type VirtualMachineScaleSetListWithLinkResult = original.VirtualMachineScaleSetListWithLinkResult
type VirtualMachineScaleSetManagedDiskParameters = original.VirtualMachineScaleSetManagedDiskParameters
type VirtualMachineScaleSetNetworkConfiguration = original.VirtualMachineScaleSetNetworkConfiguration
type VirtualMachineScaleSetNetworkConfigurationDNSSettings = original.VirtualMachineScaleSetNetworkConfigurationDNSSettings
type VirtualMachineScaleSetNetworkConfigurationProperties = original.VirtualMachineScaleSetNetworkConfigurationProperties
type VirtualMachineScaleSetNetworkProfile = original.VirtualMachineScaleSetNetworkProfile
type VirtualMachineScaleSetOSDisk = original.VirtualMachineScaleSetOSDisk
type VirtualMachineScaleSetOSProfile = original.VirtualMachineScaleSetOSProfile
type VirtualMachineScaleSetProperties = original.VirtualMachineScaleSetProperties
type VirtualMachineScaleSetPublicIPAddressConfiguration = original.VirtualMachineScaleSetPublicIPAddressConfiguration
type VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings = original.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings
type VirtualMachineScaleSetPublicIPAddressConfigurationProperties = original.VirtualMachineScaleSetPublicIPAddressConfigurationProperties
type VirtualMachineScaleSetSku = original.VirtualMachineScaleSetSku
type VirtualMachineScaleSetSkuCapacity = original.VirtualMachineScaleSetSkuCapacity
type VirtualMachineScaleSetStorageProfile = original.VirtualMachineScaleSetStorageProfile
type VirtualMachineScaleSetUpdate = original.VirtualMachineScaleSetUpdate
type VirtualMachineScaleSetUpdateIPConfiguration = original.VirtualMachineScaleSetUpdateIPConfiguration
type VirtualMachineScaleSetUpdateIPConfigurationProperties = original.VirtualMachineScaleSetUpdateIPConfigurationProperties
type VirtualMachineScaleSetUpdateNetworkConfiguration = original.VirtualMachineScaleSetUpdateNetworkConfiguration
type VirtualMachineScaleSetUpdateNetworkConfigurationProperties = original.VirtualMachineScaleSetUpdateNetworkConfigurationProperties
type VirtualMachineScaleSetUpdateNetworkProfile = original.VirtualMachineScaleSetUpdateNetworkProfile
type VirtualMachineScaleSetUpdateOSDisk = original.VirtualMachineScaleSetUpdateOSDisk
type VirtualMachineScaleSetUpdateOSProfile = original.VirtualMachineScaleSetUpdateOSProfile
type VirtualMachineScaleSetUpdateProperties = original.VirtualMachineScaleSetUpdateProperties
type VirtualMachineScaleSetUpdatePublicIPAddressConfiguration = original.VirtualMachineScaleSetUpdatePublicIPAddressConfiguration
type VirtualMachineScaleSetUpdatePublicIPAddressConfigurationProperties = original.VirtualMachineScaleSetUpdatePublicIPAddressConfigurationProperties
type VirtualMachineScaleSetUpdateStorageProfile = original.VirtualMachineScaleSetUpdateStorageProfile
type VirtualMachineScaleSetUpdateVMProfile = original.VirtualMachineScaleSetUpdateVMProfile
type VirtualMachineScaleSetVM = original.VirtualMachineScaleSetVM
type VirtualMachineScaleSetVMExtensionsSummary = original.VirtualMachineScaleSetVMExtensionsSummary
type VirtualMachineScaleSetVMInstanceIDs = original.VirtualMachineScaleSetVMInstanceIDs
type VirtualMachineScaleSetVMInstanceRequiredIDs = original.VirtualMachineScaleSetVMInstanceRequiredIDs
type VirtualMachineScaleSetVMInstanceView = original.VirtualMachineScaleSetVMInstanceView
type VirtualMachineScaleSetVMListResult = original.VirtualMachineScaleSetVMListResult
type VirtualMachineScaleSetVMProfile = original.VirtualMachineScaleSetVMProfile
type VirtualMachineScaleSetVMProperties = original.VirtualMachineScaleSetVMProperties
type VirtualMachineSize = original.VirtualMachineSize
type VirtualMachineSizeListResult = original.VirtualMachineSizeListResult
type VirtualMachineStatusCodeCount = original.VirtualMachineStatusCodeCount
type WindowsConfiguration = original.WindowsConfiguration
type WinRMConfiguration = original.WinRMConfiguration
type WinRMListener = original.WinRMListener
type VirtualMachineScaleSetRollingUpgradesClient = original.VirtualMachineScaleSetRollingUpgradesClient

func NewVirtualMachineSizesClient(subscriptionID string) VirtualMachineSizesClient {
	return original.NewVirtualMachineSizesClient(subscriptionID)
}
func NewVirtualMachineSizesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineSizesClient {
	return original.NewVirtualMachineSizesClientWithBaseURI(baseURI, subscriptionID)
}
func NewAvailabilitySetsClient(subscriptionID string) AvailabilitySetsClient {
	return original.NewAvailabilitySetsClient(subscriptionID)
}
func NewAvailabilitySetsClientWithBaseURI(baseURI string, subscriptionID string) AvailabilitySetsClient {
	return original.NewAvailabilitySetsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
func NewVirtualMachineExtensionsClient(subscriptionID string) VirtualMachineExtensionsClient {
	return original.NewVirtualMachineExtensionsClient(subscriptionID)
}
func NewVirtualMachineExtensionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineExtensionsClient {
	return original.NewVirtualMachineExtensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineScaleSetExtensionsClient(subscriptionID string) VirtualMachineScaleSetExtensionsClient {
	return original.NewVirtualMachineScaleSetExtensionsClient(subscriptionID)
}
func NewVirtualMachineScaleSetExtensionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetExtensionsClient {
	return original.NewVirtualMachineScaleSetExtensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineScaleSetVMsClient(subscriptionID string) VirtualMachineScaleSetVMsClient {
	return original.NewVirtualMachineScaleSetVMsClient(subscriptionID)
}
func NewVirtualMachineScaleSetVMsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetVMsClient {
	return original.NewVirtualMachineScaleSetVMsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDisksClient(subscriptionID string) DisksClient {
	return original.NewDisksClient(subscriptionID)
}
func NewDisksClientWithBaseURI(baseURI string, subscriptionID string) DisksClient {
	return original.NewDisksClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceSkusClient(subscriptionID string) ResourceSkusClient {
	return original.NewResourceSkusClient(subscriptionID)
}
func NewResourceSkusClientWithBaseURI(baseURI string, subscriptionID string) ResourceSkusClient {
	return original.NewResourceSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineRunCommandsClient(subscriptionID string) VirtualMachineRunCommandsClient {
	return original.NewVirtualMachineRunCommandsClient(subscriptionID)
}
func NewVirtualMachineRunCommandsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineRunCommandsClient {
	return original.NewVirtualMachineRunCommandsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachinesClient(subscriptionID string) VirtualMachinesClient {
	return original.NewVirtualMachinesClient(subscriptionID)
}
func NewVirtualMachinesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachinesClient {
	return original.NewVirtualMachinesClientWithBaseURI(baseURI, subscriptionID)
}
func NewImagesClient(subscriptionID string) ImagesClient {
	return original.NewImagesClient(subscriptionID)
}
func NewImagesClientWithBaseURI(baseURI string, subscriptionID string) ImagesClient {
	return original.NewImagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineScaleSetRollingUpgradesClient(subscriptionID string) VirtualMachineScaleSetRollingUpgradesClient {
	return original.NewVirtualMachineScaleSetRollingUpgradesClient(subscriptionID)
}
func NewVirtualMachineScaleSetRollingUpgradesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetRollingUpgradesClient {
	return original.NewVirtualMachineScaleSetRollingUpgradesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineScaleSetsClient(subscriptionID string) VirtualMachineScaleSetsClient {
	return original.NewVirtualMachineScaleSetsClient(subscriptionID)
}
func NewVirtualMachineScaleSetsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineScaleSetsClient {
	return original.NewVirtualMachineScaleSetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSnapshotsClient(subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClient(subscriptionID)
}
func NewSnapshotsClientWithBaseURI(baseURI string, subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageClient(subscriptionID string) UsageClient {
	return original.NewUsageClient(subscriptionID)
}
func NewUsageClientWithBaseURI(baseURI string, subscriptionID string) UsageClient {
	return original.NewUsageClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineExtensionImagesClient(subscriptionID string) VirtualMachineExtensionImagesClient {
	return original.NewVirtualMachineExtensionImagesClient(subscriptionID)
}
func NewVirtualMachineExtensionImagesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineExtensionImagesClient {
	return original.NewVirtualMachineExtensionImagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineImagesClient(subscriptionID string) VirtualMachineImagesClient {
	return original.NewVirtualMachineImagesClient(subscriptionID)
}
func NewVirtualMachineImagesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineImagesClient {
	return original.NewVirtualMachineImagesClientWithBaseURI(baseURI, subscriptionID)
}
