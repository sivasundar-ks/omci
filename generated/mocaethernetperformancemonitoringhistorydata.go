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

const MocaEthernetPerformanceMonitoringHistoryDataClassId uint16 = 163

var mocaethernetperformancemonitoringhistorydataBME *ManagedEntityDefinition

// MocaEthernetPerformanceMonitoringHistoryData (class ID #163) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
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
		AllowedAttributeMask: 0XFFFF,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false),
			1:  ByteField("IntervalEndTime", 0, Read, false, false, false),
			2:  Uint16Field("ThresholdData12Id", 0, Read|SetByCreate|Write, false, false, false),
			3:  Uint32Field("IncomingUnicastPackets", 0, Read, false, false, true),
			4:  Uint32Field("IncomingDiscardedPackets", 0, Read, false, false, true),
			5:  Uint32Field("IncomingErroredPackets", 0, Read, false, false, true),
			6:  Uint32Field("IncomingUnknownPackets", 0, Read, false, false, true),
			7:  Uint32Field("IncomingMulticastPackets", 0, Read, false, false, true),
			8:  Uint32Field("IncomingBroadcastPackets", 0, Read, false, false, true),
			9:  Uint32Field("IncomingOctets", 0, Read, false, false, true),
			10: Uint32Field("OutgoingUnicastPackets", 0, Read, false, false, true),
			11: Uint32Field("OutgoingDiscardedPackets", 0, Read, false, false, true),
			12: Uint32Field("OutgoingErroredPackets", 0, Read, false, false, true),
			13: Uint32Field("OutgoingUnknownPackets", 0, Read, false, false, true),
			14: Uint32Field("OutgoingMulticastPackets", 0, Read, false, false, true),
			15: Uint32Field("OutgoingBroadcastPackets", 0, Read, false, false, true),
			16: Uint32Field("OutgoingOctets", 0, Read, false, false, true),
		},
	}
}

// NewMocaEthernetPerformanceMonitoringHistoryData (class ID 163 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewMocaEthernetPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: mocaethernetperformancemonitoringhistorydataBME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
