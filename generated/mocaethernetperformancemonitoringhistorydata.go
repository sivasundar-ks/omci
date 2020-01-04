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

// MocaEthernetPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity MoCA Ethernet performance monitoring history data
const MocaEthernetPerformanceMonitoringHistoryDataClassID ClassID = ClassID(163)

var mocaethernetperformancemonitoringhistorydataBME *ManagedEntityDefinition

// MocaEthernetPerformanceMonitoringHistoryData (class ID #163)
//	This ME collects PM data for an MoCA Ethernet interface. Instances of this ME are created and
//	deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an instance of the PPTP MoCA UNI ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the PPTP MoCA UNI. (R, setbycreate)
//			(mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs
//			that contains PM threshold values. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Incoming Unicast Packets
//			Incoming unicast packets:	(R) (optional) (4-bytes)
//
//		Incoming Discarded Packets
//			Incoming discarded packets:	(R) (optional) (4-bytes)
//
//		Incoming Errored Packets
//			Incoming errored packets:	(R) (optional) (4-bytes)
//
//		Incoming Unknown Packets
//			Incoming unknown packets:	(R) (optional) (4-bytes)
//
//		Incoming Multicast Packets
//			Incoming multicast packets:	(R) (optional) (4-bytes)
//
//		Incoming Broadcast Packets
//			Incoming broadcast packets:	(R) (optional) (4-bytes)
//
//		Incoming Octets
//			Incoming octets:	(R) (optional) (4-bytes)
//
//		Outgoing Unicast Packets
//			Outgoing unicast packets:	(R) (optional) (4-bytes)
//
//		Outgoing Discarded Packets
//			Outgoing discarded packets:	(R) (optional) (4-bytes)
//
//		Outgoing Errored Packets
//			Outgoing errored packets:	(R) (optional) (4-bytes)
//
//		Outgoing Unknown Packets
//			Outgoing unknown packets:	(R) (optional) (4-bytes)
//
//		Outgoing Multicast Packets
//			Outgoing multicast packets:	(R) (optional) (4-bytes)
//
//		Outgoing Broadcast Packets
//			Outgoing broadcast packets:	(R) (optional) (4-bytes)
//
//		Outgoing Octets
//			Outgoing octets:	(R) (optional) (4-bytes)
//
type MocaEthernetPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	mocaethernetperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "MocaEthernetPerformanceMonitoringHistoryData",
		ClassID: 163,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xffff,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 0),
			1:  ByteField("IntervalEndTime", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, 1),
			2:  Uint16Field("ThresholdData12Id", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 2),
			3:  Uint32Field("IncomingUnicastPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 3),
			4:  Uint32Field("IncomingDiscardedPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 4),
			5:  Uint32Field("IncomingErroredPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 5),
			6:  Uint32Field("IncomingUnknownPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 6),
			7:  Uint32Field("IncomingMulticastPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 7),
			8:  Uint32Field("IncomingBroadcastPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 8),
			9:  Uint32Field("IncomingOctets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 9),
			10: Uint32Field("OutgoingUnicastPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 10),
			11: Uint32Field("OutgoingDiscardedPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 11),
			12: Uint32Field("OutgoingErroredPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 12),
			13: Uint32Field("OutgoingUnknownPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 13),
			14: Uint32Field("OutgoingMulticastPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 14),
			15: Uint32Field("OutgoingBroadcastPackets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 15),
			16: Uint32Field("OutgoingOctets", CounterAttributeType, 0, mapset.NewSetWith(Read), false, true, 16),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewMocaEthernetPerformanceMonitoringHistoryData (class ID 163) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewMocaEthernetPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*mocaethernetperformancemonitoringhistorydataBME, params...)
}
