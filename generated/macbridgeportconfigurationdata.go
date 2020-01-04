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

// MacBridgePortConfigurationDataClassID is the 16-bit ID for the OMCI
// Managed entity MAC bridge port configuration data
const MacBridgePortConfigurationDataClassID ClassID = ClassID(47)

var macbridgeportconfigurationdataBME *ManagedEntityDefinition

// MacBridgePortConfigurationData (class ID #47)
//	This ME models a port on a MAC bridge. Instances of this ME are created and deleted by the OLT.
//
//	Relationships
//		An instance of this ME is linked to an instance of the MAC bridge service profile. Additional
//		bridge port control capabilities are provided by implicitly linked instances of some or all of:////		o	MAC bridge port filter table data;////		o	MAC bridge port filter pre-assign table;////		o	VLAN tagging filter data;////		o	Dot1 rate limiter.////		Real-time status of the bridge port is provided by implicitly linked instances of:////		o	MAC bridge port designation data;////		o	MAC bridge port bridge table data;////		o	Multicast subscriber monitor.////		Bridge port PM collection is provided by implicitly linked instances of:////		o	MAC bridge port PM history data;////		o	Ethernet frame PM history data upstream and downstream;////		o	Ethernet frame extended PM (preferred).
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Bridge Id Pointer
//			Bridge ID pointer: This attribute points to an instance of the MAC bridge service profile.
//			(R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Port Num
//			Port num:	This attribute is the bridge port number. It must be unique among all ports associated
//			with a particular MAC bridge service profile. (R,-W, setbycreate) (mandatory) (1-byte)
//
//		Tp Type
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Tp Pointer
//			NOTE 1 - When the TP type is very high-speed digital subscriber line (VDSL) or xDSL, the two
//			MSBs may be used to indicate a bearer channel.
//
//		Port Priority
//			Port priority:	This attribute denotes the priority of the port for use in (rapid) spanning tree
//			algorithms. The range is 0..255. (R,-W, setbycreate) (optional) (2-bytes)
//
//		Port Path Cost
//			Port path cost: This attribute specifies the contribution of the port to the path cost towards
//			the spanning tree root bridge. The range is 1..65535. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Port Spanning Tree Ind
//			Port spanning tree ind: The Boolean value true enables (R)STP LAN topology change detection at
//			this port. The value false disables topology change detection. (R,-W, setbycreate) (mandatory)
//			(1-byte)
//
//		Deprecated 1
//			Deprecated 1: This attribute is not used. If present, it should be ignored by both the ONU and
//			the OLT, except as necessary to comply with OMCI message definitions. (R,-W, setbycreate)
//			(optional) (1-byte)
//
//		Deprecated 2
//			Deprecated 2: This attribute is not used. If present, it should be ignored by both the ONU and
//			the OLT, except as necessary to comply with OMCI message definitions. (R,-W, setbycreate)
//			(1-byte) (optional)
//
//		Port Mac Address
//			Port MAC address: If the TP associated with this port has a MAC address, this attribute
//			specifies it. (R) (optional) (6-bytes)
//
//		Outbound Td Pointer
//			Outbound TD pointer: This attribute points to a traffic descriptor that limits the traffic rate
//			leaving the MAC bridge. (R,-W) (optional) (2-byte)
//
//		Inbound Td Pointer
//			Inbound TD pointer: This attribute points to a traffic descriptor that limits the traffic rate
//			entering the MAC bridge. (R,-W) (optional) (2-byte)
//
//		Mac Learning Depth
//			NOTE 2 - If this attribute is not zero, its value overrides the value set in the MAC learning
//			depth attribute of the MAC bridge service profile.
//
type MacBridgePortConfigurationData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	macbridgeportconfigurationdataBME = &ManagedEntityDefinition{
		Name:    "MacBridgePortConfigurationData",
		ClassID: 47,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xfff8,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 0),
			1:  Uint16Field("BridgeIdPointer", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 1),
			2:  ByteField("PortNum", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 2),
			3:  ByteField("TpType", EnumerationAttributeType, 1, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 3),
			4:  Uint16Field("TpPointer", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 4),
			5:  Uint16Field("PortPriority", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, 5),
			6:  Uint16Field("PortPathCost", UnsignedIntegerAttributeType, 1, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 6),
			7:  ByteField("PortSpanningTreeInd", EnumerationAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 7),
			8:  ByteField("Deprecated1", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, 8),
			9:  ByteField("Deprecated2", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, 9),
			10: MultiByteField("PortMacAddress", OctetsAttributeType, 6, toOctets("AAAAAAAA"), mapset.NewSetWith(Read), false, true, 10),
			11: Uint16Field("OutboundTdPointer", PointerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 11),
			12: Uint16Field("InboundTdPointer", PointerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 12),
			13: ByteField("MacLearningDepth", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, 13),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewMacBridgePortConfigurationData (class ID 47) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewMacBridgePortConfigurationData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*macbridgeportconfigurationdataBME, params...)
}
