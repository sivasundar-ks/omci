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

const PseudowirePerformanceMonitoringHistoryDataClassId uint16 = 285

var pseudowireperformancemonitoringhistorydataBME *ManagedEntityDefinition

// PseudowirePerformanceMonitoringHistoryData (class ID #285) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type PseudowirePerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	pseudowireperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "PseudowirePerformanceMonitoringHistoryData",
		ClassID: 285,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFE,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false, 0),
			1:  ByteField("IntervalEndTime", 0, Read, false, false, false, 1),
			2:  Uint16Field("ThresholdData12Id", 0, Read|SetByCreate|Write, false, false, false, 2),
			3:  Uint32Field("ReceivedPackets", 0, Read, false, false, false, 3),
			4:  Uint32Field("TransmittedPackets", 0, Read, false, false, false, 4),
			5:  Uint32Field("MissingPackets", 0, Read, false, false, false, 5),
			6:  Uint32Field("MisorderedPackets,Usable", 0, Read, false, false, false, 6),
			7:  Uint32Field("MisorderedPacketsDropped", 0, Read, false, false, false, 7),
			8:  Uint32Field("PlayoutBufferUnderrunsOverruns", 0, Read, false, false, false, 8),
			9:  Uint32Field("MalformedPackets", 0, Read, false, false, false, 9),
			10: Uint32Field("StrayPackets", 0, Read, false, false, false, 10),
			11: Uint32Field("RemotePacketLoss", 0, Read, false, false, false, 11),
			12: Uint32Field("TdmLBitPacketsTransmitted", 0, Read, false, false, false, 12),
			13: Uint32Field("Es", 0, Read, false, false, false, 13),
			14: Uint32Field("Ses", 0, Read, false, false, false, 14),
			15: Uint32Field("Uas", 0, Read, false, false, false, 15),
		},
	}
}

// NewPseudowirePerformanceMonitoringHistoryData (class ID 285 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewPseudowirePerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(pseudowireperformancemonitoringhistorydataBME, params...)
}
