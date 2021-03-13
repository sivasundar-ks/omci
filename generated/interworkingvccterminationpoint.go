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

// InterworkingVccTerminationPointClassID is the 16-bit ID for the OMCI
// Managed entity Interworking VCC termination point
const InterworkingVccTerminationPointClassID ClassID = ClassID(14)

var interworkingvccterminationpointBME *ManagedEntityDefinition

// InterworkingVccTerminationPoint (class ID #14)
//	An instance of this ME represents a point in the ONU where the IW of a service or underlying
//	physical infrastructure (e.g., ADSL) to an ATM layer takes place. At this point, ATM cells are
//	generated from a bit stream (e.g., Ethernet) or a bit stream is reconstructed from ATM cells.
//
//	Instances of this ME are created and deleted by the OLT.
//
//	Relationships
//		One instance of this ME exists for each occurrence of transformation of a data stream into ATM
//		cells and vice versa.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. (R,-setbycreate)
//			(mandatory) (2-bytes)
//
//		Vci Value
//			VCI value:	This attribute identifies the VCI value associated with this IW VCC TP. (R,-W,
//			setbycreate) (mandatory) (2-bytes)
//
//		Vp Network Ctp Connectivity Pointer
//			VP network CTP connectivity pointer: This attribute points to the VP network CTP associated with
//			this IW VCC TP. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Deprecated 1
//			Deprecated 1: Not used; should be set to 0. (R,-W, setbycreate) (mandatory) (1-byte)
//
//		Deprecated 2
//			Deprecated 2: Not used; should be set to 0. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Aal5 Profile Pointer
//			AAL5 profile pointer: This attribute points to an instance of the AAL5 profile. (R,-W,
//			setbycreate) (mandatory) (2-bytes)
//
//		Deprecated 3
//			Deprecated 3: Not used; should be set to 0. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Aal Loopback Configuration
//			The default value of this attribute is 0. (R,-W) (mandatory) (1-byte)
//
//		Pptp Counter
//			PPTP counter: This value is the number of instances of PPTP MEs associated with this instance of
//			the IW VCC TP. (R) (optional) (1-byte)
//
//		Operational State
//			Operational state: This attribute indicates whether the ME is capable of performing its
//			function. Valid values are enabled (0) and disabled (1). (R) (optional) (1-byte)
//
type InterworkingVccTerminationPoint struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	interworkingvccterminationpointBME = &ManagedEntityDefinition{
		Name:    "InterworkingVccTerminationPoint",
		ClassID: 14,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xff80,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1: Uint16Field("VciValue", UnsignedIntegerAttributeType, 0x8000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 1),
			2: Uint16Field("VpNetworkCtpConnectivityPointer", UnsignedIntegerAttributeType, 0x4000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 2),
			3: ByteField("Deprecated1", UnsignedIntegerAttributeType, 0x2000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, 3),
			4: Uint16Field("Deprecated2", UnsignedIntegerAttributeType, 0x1000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, 4),
			5: Uint16Field("Aal5ProfilePointer", UnsignedIntegerAttributeType, 0x0800, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 5),
			6: Uint16Field("Deprecated3", UnsignedIntegerAttributeType, 0x0400, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, true, 6),
			7: ByteField("AalLoopbackConfiguration", UnsignedIntegerAttributeType, 0x0200, 0, mapset.NewSetWith(Read, Write), false, false, false, 7),
			8: ByteField("PptpCounter", UnsignedIntegerAttributeType, 0x0100, 0, mapset.NewSetWith(Read), false, true, false, 8),
			9: ByteField("OperationalState", UnsignedIntegerAttributeType, 0x0080, 0, mapset.NewSetWith(Read), true, true, false, 9),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
		Alarms: AlarmMap{
			0: "End-to-end VC AIS layer management indication receiving (LMIR)",
			1: "End-to-end VC RDI LMIR",
			2: "End-to-end VC AIS layer management indication generation (LMIG)",
			3: "End-to-end VC RDI LMIG",
			4: "Segment loss of continuity",
			5: "End-to-end loss of continuity",
			6: "CSA",
		},
	}
}

// NewInterworkingVccTerminationPoint (class ID 14) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewInterworkingVccTerminationPoint(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*interworkingvccterminationpointBME, params...)
}
