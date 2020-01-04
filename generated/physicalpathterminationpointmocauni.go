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

// PhysicalPathTerminationPointMocaUniClassID is the 16-bit ID for the OMCI
// Managed entity Physical path termination point MoCA UNI
const PhysicalPathTerminationPointMocaUniClassID ClassID = ClassID(162)

var physicalpathterminationpointmocauniBME *ManagedEntityDefinition

// PhysicalPathTerminationPointMocaUni (class ID #162)
//	This ME represents an MoCA UNI, where physical paths terminate and physical path level functions
//	are performed.
//
//	The ONU automatically creates an instance of this ME per port as follows.
//
//	o	When the ONU has MoCA ports built into its factory configuration.
//
//	o	When a cardholder is provisioned to expect a circuit pack of the MoCA type.
//
//	o	When a cardholder provisioned for plug-and-play is equipped with a circuit pack of the MoCA
//	type. Note that the installation of a plug-and-play card may indicate the presence of MoCA ports
//	via equipment ID as well as its type, and indeed may cause the ONU to instantiate a port-mapping
//	package that specifies MoCA ports.
//
//	The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to
//	expect an MoCA circuit pack, nor is it equipped with an MoCA circuit pack.
//
//	Relationships
//		An instance of this ME is associated with each real or pre-provisioned MoCA port.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2-byte
//			number is directly associated with the physical position of the UNI. The first byte is the slot
//			ID (defined in clause 9.1.5). The second byte is the port ID, with the range 1..255. (R)
//			(mandatory) (2-bytes)
//
//		Loopback Configuration
//			Upon ME instantiation, the ONU sets this attribute to 0. (R,-W) (optional) (1-byte)
//
//		Administrative State
//			Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this
//			ME. Administrative state is further described in clause A.1.6. (R,-W) (mandatory) (1-byte)
//
//		Operational State
//			Operational state: This attribute indicates whether the ME is capable of performing its
//			function. Valid values are enabled (0) and disabled (1). (R) (optional) (1-byte)
//
//		Max Frame Size
//			Max frame size: This attribute denotes the maximum frame size allowed across this interface.
//			Upon ME instantiation, the ONU sets this attribute to 1518. (R,-W) (mandatory) (2-bytes)
//
//		Arc
//			ARC:	See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
//		Arc Interval
//			ARC interval: See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
//		Pppoe Filter
//			PPPoE filter: This attribute controls filtering of PPPoE packets on this MoCA port. When its
//			value is 1, all packets other than PPPoE packets are discarded. The default 0 accepts packets of
//			all types. (R,-W) (optional) (1-byte)
//
//		Network Status
//			(R) (mandatory) (1-byte)
//
//		Password
//			Password:	This attribute specifies the MoCA encryption key. It is an ASCII string of 17 decimal
//			digits. Upon ME instantiation, the ONU sets this attribute to 17 null bytes. (R,-W) (mandatory)
//			(17-bytes)
//
//		Privacy Enabled
//			Privacy enabled: This attribute activates (1) link-layer security. The default value 0
//			deactivates it. (R,-W) (mandatory) (1-byte)
//
//		Minimum Bandwidth Alarm Threshold
//			Minimum bandwidth alarm threshold: This attribute specifies the minimum desired PHY link
//			bandwidth between two nodes. If the actual bandwidth is lower, an LL alarm is declared. Valid
//			values are 0 to 0x0410 (260-Mbit/s) in 0.25-Mbit/s increments. The default value is 0x02D0
//			(180-Mbit/s). The value 0 disables the threshold. (R,-W) (optional) (2-bytes)
//
//		Frequency Mask
//			Frequency mask: This attribute is a bit map of the centre frequencies that the interface is
//			permitted to use, where each bit represents a centre frequency. The LSB (b[1]) corresponds to
//			centre frequency 800-MHz. The next significant bit (b[2]) corresponds to centre frequency
//			825-MHz. The 28th bit (b[28]) corresponds to centre frequency 1500-MHz. The four MSBs are not
//			used. (R,-W) (optional) (4-bytes)
//
//		Rf Channel
//			RF channel:	This attribute reports the frequency to which the MoCA interface is currently tuned,
//			in megahertz. (R) (mandatory) (2-bytes)
//
//		Last Operational Frequency
//			Last operational frequency: This attribute reports the frequency to which the MoCA interface was
//			tuned when last operational, in megahertz. (R) (mandatory) (2-bytes)
//
type PhysicalPathTerminationPointMocaUni struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	physicalpathterminationpointmocauniBME = &ManagedEntityDefinition{
		Name:    "PhysicalPathTerminationPointMocaUni",
		ClassID: 162,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0xfffc,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read), false, false, 0),
			1:  ByteField("LoopbackConfiguration", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 1),
			2:  ByteField("AdministrativeState", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, 2),
			3:  ByteField("OperationalState", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), true, true, 3),
			4:  Uint16Field("MaxFrameSize", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, 4),
			5:  ByteField("Arc", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), true, true, 5),
			6:  ByteField("ArcInterval", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 6),
			7:  ByteField("PppoeFilter", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 7),
			8:  ByteField("NetworkStatus", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, 8),
			9:  MultiByteField("Password", OctetsAttributeType, 17, nil, mapset.NewSetWith(Read, Write), false, false, 9),
			10: ByteField("PrivacyEnabled", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, 10),
			11: Uint16Field("MinimumBandwidthAlarmThreshold", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 11),
			12: Uint32Field("FrequencyMask", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 12),
			13: Uint16Field("RfChannel", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, 13),
			14: Uint16Field("LastOperationalFrequency", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, 14),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewPhysicalPathTerminationPointMocaUni (class ID 162) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewPhysicalPathTerminationPointMocaUni(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*physicalpathterminationpointmocauniBME, params...)
}
