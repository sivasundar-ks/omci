/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */

package generated

import "github.com/deckarep/golang-set"

// VlanTaggingOperationConfigurationDataClassID is the 16-bit ID for the OMCI
// Managed entity VLAN tagging operation configuration data
const VlanTaggingOperationConfigurationDataClassID ClassID = ClassID(78)

var vlantaggingoperationconfigurationdataBME *ManagedEntityDefinition

// VlanTaggingOperationConfigurationData (class ID #78)
//	This ME organizes data associated with VLAN tagging. Instances of this ME are created and
//	deleted by the OLT.
//
//	NOTE 1 - The extended VLAN tagging operation configuration data of clause 9.3.13 is preferred
//	for new implementations.
//
//	Relationships
//		Zero or one instance of this ME may exist for an instance of any ME that can terminate or modify
//		an Ethernet stream.////		When this ME is associated with a UNI-side TP, it performs its upstream classification and
//		tagging operations before offering the upstream frame to other filtering, bridging or switching
//		functions. In the downstream direction, the defined inverse operation is the last operation
//		performed on the frame before offering it to the UNI-side termination.////		When this ME is associated with an ANI-side TP, it performs its upstream classification and
//		tagging operations as the last step before queueing for transmission to the OLT, after having
//		received the upstream frame from other filtering, bridging or switching functions. In the
//		downstream direction, the defined inverse operation is the first operation performed on the
//		frame before offering it to possible filter, bridge or switch functions.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. When the
//			optional association type attribute is 0 or undefined, this attribute's value is the same as the
//			ID of the ME with which this VLAN tagging operation configuration data instance is associated,
//			which may be either a PPTP Ethernet UNI or an IP host config data or an IPv6 host config data
//			ME. Otherwise, the value of the ME ID is unconstrained except by the need to be unique. (R, set-
//			by-create) (mandatory) (2 bytes)
//
//		Upstream Vlan Tagging Operation Mode
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Upstream Vlan Tag Tci Value
//			Upstream VLAN tag TCI value: This attribute specifies the TCI for upstream VLAN tagging. It is
//			used when the upstream VLAN tagging operation mode is 1 or 2. (R,-W, setbycreate) (mandatory)
//			(2-bytes)
//
//		Downstream Vlan Tagging Operation Mode
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Association Type
//			The associated ME instance is identified by the associated ME pointer. (R,-W, setbycreate)
//			(optional) (1-byte)
//
//		Associated Me Pointer
//			NOTE 3 - When the association type is xDSL, the two MSBs may be used to indicate a bearer
//			channel.
//
type VlanTaggingOperationConfigurationData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	vlantaggingoperationconfigurationdataBME = &ManagedEntityDefinition{
		Name:    "VlanTaggingOperationConfigurationData",
		ClassID: 78,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xf800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 0),
			1: ByteField("UpstreamVlanTaggingOperationMode", EnumerationAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 1),
			2: Uint16Field("UpstreamVlanTagTciValue", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 2),
			3: ByteField("DownstreamVlanTaggingOperationMode", EnumerationAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 3),
			4: ByteField("AssociationType", EnumerationAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, 4),
			5: Uint16Field("AssociatedMePointer", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, 5),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewVlanTaggingOperationConfigurationData (class ID 78) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewVlanTaggingOperationConfigurationData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*vlantaggingoperationconfigurationdataBME, params...)
}
