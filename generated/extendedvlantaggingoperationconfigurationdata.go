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

// ExtendedVlanTaggingOperationConfigurationDataClassID is the 16-bit ID for the OMCI
// Managed entity Extended VLAN tagging operation configuration data
const ExtendedVlanTaggingOperationConfigurationDataClassID ClassID = ClassID(171)

var extendedvlantaggingoperationconfigurationdataBME *ManagedEntityDefinition

// ExtendedVlanTaggingOperationConfigurationData (class ID #171)
//	This ME organizes data associated with VLAN tagging. Regardless of its point of attachment, the
//	specified tagging operations refer to the upstream direction. Instances of this ME are created
//	and deleted by the OLT.
//
//	Relationships
//		Zero or one instance of this ME may exist for an instance of any ME that can terminate or modify
//		an Ethernet stream.////		When this ME is associated with a UNI-side TP, it performs its upstream classification and
//		tagging operations before offering the upstream frame to other filtering, bridging or switching
//		functions. In the downstream direction, the defined inverse operation is the last operation
//		performed on the frame before offering it to the UNI-side termination.////		When this ME is associated with an ANI-side TP, it performs its upstream classification and
//		tagging operations as the last step before transmission to the OLT, after having received the
//		upstream frame from other filtering, bridging or switching functions. In the downstream
//		direction, the defined inverse operation is the first operation performed on the frame before
//		offering it to possible filter, bridge or switch functions.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute provides a unique number for each instance of this ME. (R,
//			setbycreate) (mandatory) (2-bytes)
//
//		Association Type
//			When the extended VLAN tagging ME is associated with the ANI side, it behaves as an upstream
//			egress rule, and as a downstream ingress rule when the downstream mode attribute is equal to 0.
//			When the extended VLAN tagging ME is associated with the UNI side, the extended VLAN tagging ME
//			behaves as an upstream ingress rule, and as a downstream egress rule when the downstream mode
//			attribute is equal to 0.
//
//		Received Frame Vlan Tagging Operation Table Max Size
//			Received frame VLAN tagging operation table max size: This attribute indicates the maximum
//			number of entries that can be set in the received frame VLAN tagging operation table. (R)
//			(mandatory) (2-bytes)
//
//		Input Tpid
//			Input TPID:	This attribute gives the special TPID value for operations on the input (filtering)
//			side of the table. Typical values include 0x88A8 and 0x9100. (R,-W) (mandatory) (2-bytes)
//
//		Output Tpid
//			Output TPID: This attribute gives the special TPID value for operations on the output (tagging)
//			side of the table. Typical values include 0x88A8 and 0x9100. (R,-W) (mandatory) (2-bytes)
//
//		Downstream Mode
//			All other values are reserved. (R, W) (mandatory) (1 byte)
//
//		Received Frame Vlan Tagging Operation Table
//			111	Set TPID-=-output TPID, DEI = 1
//
//		Associated Me Pointer
//			NOTE 5 - When the association type is xDSL, the two MSBs may be used to indicate a bearer
//			channel.
//
//		Dscp To P Bit Mapping
//			NOTE 6 - If certain bits in the DSCP field are to be ignored in the mapping process, the
//			attribute should be provisioned such that all possible values of those bits produce the same
//			P-bit mapping. This can be applied to the case where instead of full DSCP, the operator wishes
//			to adopt the priority mechanism based on IP precedence, which needs only the three MSBs of the
//			DSCP field.
//
type ExtendedVlanTaggingOperationConfigurationData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	extendedvlantaggingoperationconfigurationdataBME = &ManagedEntityDefinition{
		Name:    "ExtendedVlanTaggingOperationConfigurationData",
		ClassID: 171,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetNext,
			Set,
		),
		AllowedAttributeMask: 0xff00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 0),
			1: ByteField("AssociationType", EnumerationAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 1),
			2: Uint16Field("ReceivedFrameVlanTaggingOperationTableMaxSize", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, 2),
			3: Uint16Field("InputTpid", UnsignedIntegerAttributeType, 34984, mapset.NewSetWith(Read, Write), false, false, 3),
			4: Uint16Field("OutputTpid", UnsignedIntegerAttributeType, 34984, mapset.NewSetWith(Read, Write), false, false, 4),
			5: ByteField("DownstreamMode", EnumerationAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, 5),
			6: TableField("ReceivedFrameVlanTaggingOperationTable", TableAttributeType, TableInfo{nil, 16}, mapset.NewSetWith(Read, Write), false, false, 6),
			7: Uint16Field("AssociatedMePointer", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 7),
			8: MultiByteField("DscpToPBitMapping", OctetsAttributeType, 24, toOctets("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"), mapset.NewSetWith(Read, Write), false, true, 8),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewExtendedVlanTaggingOperationConfigurationData (class ID 171) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewExtendedVlanTaggingOperationConfigurationData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*extendedvlantaggingoperationconfigurationdataBME, params...)
}
