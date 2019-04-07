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

package notificationhubs

import original "github.com/Azure/azure-sdk-for-go/services/notificationhubs/mgmt/2017-04-01/notificationhubs"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type GroupClient = original.GroupClient
type HubsClient = original.HubsClient
type AccessRights = original.AccessRights

const (
	Listen AccessRights = original.Listen
	Manage AccessRights = original.Manage
	Send   AccessRights = original.Send
)

type NamespaceType = original.NamespaceType

const (
	Messaging       NamespaceType = original.Messaging
	NotificationHub NamespaceType = original.NotificationHub
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Free     SkuName = original.Free
	Standard SkuName = original.Standard
)

type AdmCredential = original.AdmCredential
type AdmCredentialProperties = original.AdmCredentialProperties
type ApnsCredential = original.ApnsCredential
type ApnsCredentialProperties = original.ApnsCredentialProperties
type BaiduCredential = original.BaiduCredential
type BaiduCredentialProperties = original.BaiduCredentialProperties
type CheckAvailabilityParameters = original.CheckAvailabilityParameters
type CheckAvailabilityResult = original.CheckAvailabilityResult
type CheckNameAvailabilityRequestParameters = original.CheckNameAvailabilityRequestParameters
type CheckNameAvailabilityResponse = original.CheckNameAvailabilityResponse
type CreateOrUpdateParameters = original.CreateOrUpdateParameters
type GcmCredential = original.GcmCredential
type GcmCredentialProperties = original.GcmCredentialProperties
type ListResult = original.ListResult
type MpnsCredential = original.MpnsCredential
type MpnsCredentialProperties = original.MpnsCredentialProperties
type NamespaceCreateOrUpdateParameters = original.NamespaceCreateOrUpdateParameters
type NamespaceListResult = original.NamespaceListResult
type NamespacePatchParameters = original.NamespacePatchParameters
type NamespaceProperties = original.NamespaceProperties
type NamespaceResource = original.NamespaceResource
type PnsCredentialsProperties = original.PnsCredentialsProperties
type PnsCredentialsResource = original.PnsCredentialsResource
type PolicykeyResource = original.PolicykeyResource
type Properties = original.Properties
type Resource = original.Resource
type ResourceListKeys = original.ResourceListKeys
type ResourceType = original.ResourceType
type SharedAccessAuthorizationRuleCreateOrUpdateParameters = original.SharedAccessAuthorizationRuleCreateOrUpdateParameters
type SharedAccessAuthorizationRuleListResult = original.SharedAccessAuthorizationRuleListResult
type SharedAccessAuthorizationRuleProperties = original.SharedAccessAuthorizationRuleProperties
type SharedAccessAuthorizationRuleResource = original.SharedAccessAuthorizationRuleResource
type Sku = original.Sku
type SubResource = original.SubResource
type WnsCredential = original.WnsCredential
type WnsCredentialProperties = original.WnsCredentialProperties
type NameClient = original.NameClient
type NamespacesClient = original.NamespacesClient

func NewNameClient(subscriptionID string) NameClient {
	return original.NewNameClient(subscriptionID)
}
func NewNameClientWithBaseURI(baseURI string, subscriptionID string) NameClient {
	return original.NewNameClientWithBaseURI(baseURI, subscriptionID)
}
func NewNamespacesClient(subscriptionID string) NamespacesClient {
	return original.NewNamespacesClient(subscriptionID)
}
func NewNamespacesClientWithBaseURI(baseURI string, subscriptionID string) NamespacesClient {
	return original.NewNamespacesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewGroupClient(subscriptionID string) GroupClient {
	return original.NewGroupClient(subscriptionID)
}
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI, subscriptionID)
}
func NewHubsClient(subscriptionID string) HubsClient {
	return original.NewHubsClient(subscriptionID)
}
func NewHubsClientWithBaseURI(baseURI string, subscriptionID string) HubsClient {
	return original.NewHubsClientWithBaseURI(baseURI, subscriptionID)
}
