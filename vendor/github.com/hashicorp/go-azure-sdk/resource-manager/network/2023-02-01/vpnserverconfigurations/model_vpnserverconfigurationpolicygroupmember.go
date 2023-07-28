package vpnserverconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnServerConfigurationPolicyGroupMember struct {
	AttributeType  *VpnPolicyMemberAttributeType `json:"attributeType,omitempty"`
	AttributeValue *string                       `json:"attributeValue,omitempty"`
	Name           *string                       `json:"name,omitempty"`
}