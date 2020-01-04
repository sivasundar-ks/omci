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

// ReDownstreamAmplifierClassID is the 16-bit ID for the OMCI
// Managed entity RE downstream amplifier
const ReDownstreamAmplifierClassID ClassID = ClassID(316)

var redownstreamamplifierBME *ManagedEntityDefinition

// ReDownstreamAmplifier (class ID #316)
//	This ME organizes data associated with each OA for downstream data supported by the RE. The
//	management ONU automatically creates one instance of this ME for each downstream OA as follows.
//
//	o	When the RE has mid-span PON RE downstream OA ports built into its factory configuration.
//
//	o	When a cardholder is provisioned to expect a circuit pack of the mid-span PON RE downstream OA
//	type.
//
//	o	When a cardholder provisioned for plug-and-play is equipped with a circuit pack of the midspan
//	PON RE downstream OA type. Note that the installation of a plug-and-play card may indicate the
//	presence of a mid-span PON RE downstream OA via equipment ID as well as its type attribute, and
//	indeed may cause the management ONU to instantiate a port-mapping package to specify the ports
//	precisely.
//
//	The management ONU automatically deletes instances of this ME when a cardholder is neither
//	provisioned to expect a mid-span PON RE downstream OA circuit pack, nor is it equipped with a
//	mid-span PON RE downstream OA circuit pack.
//
//	Relationships
//		An instance of this ME is associated with a downstream OA and with an instance of a circuit
//		pack. If the RE includes OEO regeneration in either direction, the RE downstream amplifier is
//		also associated with an RE ANI-G. Refer to clause-9.14.1 for further discussion.
//
//	Attributes
//		Managed Entity Id
//			NOTE 1 - This ME ID may be identical to that of an RE ANI-G if it shares the same physical slot-
//			port.
//
//		Administrative State
//			NOTE 2- When an RE supports multiple PONs, or protected access to a single PON, its primary
//			ANI-G cannot be completely shut down, due to a loss of the management communications capability.
//			Complete blocking of service and removal of power may nevertheless be appropriate for secondary
//			RE ANI-Gs. Administrative lock suppresses alarms and notifications for both primary and
//			secondary RE ANI-Gs. Administrative lock suppresses alarms and notifications for an RE
//			downstream amplifier, be it either primary or secondary.
//
//		Operational State
//			Operational state: This attribute indicates whether the ME is capable of performing its
//			function. Valid values are enabled (0) and disabled (1). (R) (optional) (1-byte)
//
//		Arc
//			ARC:	See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
//		Arc Interval
//			ARC interval: See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
//		Operational Mode
//			(R,W) (mandatory) (1-byte)
//
//		Input Optical Signal Level
//			Input optical signal level: This attribute reports the current measurement of the input optical
//			signal power of the downstream OA. Its value is a 2s-complement integer referred to 1-mW (i.e.,
//			dBm), with 0.002-dB granularity. (R) (optional) (2-bytes)
//
//		Lower Input Optical Threshold
//			Lower input optical threshold: This attribute specifies the optical level the RE uses to declare
//			the low received optical power alarm. Valid values are -127-dBm (coded as 254) to 0-dBm (coded
//			as 0) in 0.5-dB increments. The default value 0xFF selects the RE's internal policy. (R,-W)
//			(optional) (1-byte)
//
//		Upper Input Optical Threshold
//			Upper input optical threshold: This attribute specifies the optical level the RE uses to declare
//			the high received optical power alarm. Valid values are -127-dBm (coded as 254) to 0-dBm (coded
//			as 0) in 0.5-dB increments. The default value 0xFF selects the RE's internal policy. (R,-W)
//			(optional) (1-byte)
//
//		Output Optical Signal Level
//			Output optical signal level: This attribute reports the current measurement of the mean optical
//			launch power of the downstream OA. Its value is a 2s-complement integer referred to 1-mW (i.e.,
//			dBm), with 0.002-dB granularity. (R) (optional) (2-bytes)
//
//		Lower Output Optical Threshold
//			Lower output optical threshold: This attribute specifies the minimum mean optical launch power
//			that the RE uses to declare the low transmit optical power alarm. Its value is a 2s complement
//			integer referred to 1-mW (i.e., dBm), with 0.5-dB granularity. The default value 0x7F selects
//			the RE's internal policy. (R,-W) (optional) (1-byte)
//
//		Upper Output Optical Threshold
//			Upper output optical threshold: This attribute specifies the maximum mean optical launch power
//			that the RE uses to declare the high transmit optical power alarm. Its value is a 2s complement
//			integer referred to 1-mW (i.e., dBm), with 0.5-dB granularity. The default value 0x7F selects
//			the RE's internal policy. (R,-W) (optional) (1-byte)
//
//		R'S' Splitter Coupling Ratio
//			R'S' splitter coupling ratio: This attribute reports the coupling ratio of the splitter at the
//			R'/S' interface that connects the embedded management ONU and the amplifiers to the OTL. Valid
//			values are 99:1 (coded as 99-decimal) to 1:99 (coded as 1 decimal), where the first value is the
//			value encoded and is the percentage of the optical signal connected to the amplifier. The
//			default value 0xFF indicates that there is no splitter connected to this upstream/downstream
//			amplifier pair. (R) (optional) (1-byte)
//
type ReDownstreamAmplifier struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	redownstreamamplifierBME = &ManagedEntityDefinition{
		Name:    "ReDownstreamAmplifier",
		ClassID: 316,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
			Test,
		),
		AllowedAttributeMask: 0xfff0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read), false, false, 0),
			1:  ByteField("AdministrativeState", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, 1),
			2:  ByteField("OperationalState", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), true, true, 2),
			3:  ByteField("Arc", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), true, true, 3),
			4:  ByteField("ArcInterval", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 4),
			5:  ByteField("OperationalMode", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, 5),
			6:  Uint16Field("InputOpticalSignalLevel", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, true, 6),
			7:  ByteField("LowerInputOpticalThreshold", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 7),
			8:  ByteField("UpperInputOpticalThreshold", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 8),
			9:  Uint16Field("OutputOpticalSignalLevel", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, true, 9),
			10: ByteField("LowerOutputOpticalThreshold", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 10),
			11: ByteField("UpperOutputOpticalThreshold", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, true, 11),
			12: ByteField("R'S'SplitterCouplingRatio", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, true, 12),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewReDownstreamAmplifier (class ID 316) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewReDownstreamAmplifier(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*redownstreamamplifierBME, params...)
}
