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

// XdslXtuRChannelPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity xDSL xTU-R channel performance monitoring history data
const XdslXtuRChannelPerformanceMonitoringHistoryDataClassID ClassID = ClassID(115)

var xdslxturchannelperformancemonitoringhistorydataBME *ManagedEntityDefinition

// XdslXtuRChannelPerformanceMonitoringHistoryData (class ID #115)
//	This ME collects PM data of the xTUC to xTUR channel as seen from the xTU-R. Instances of this
//	ME are created and deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an xDSL bearer channel. Several instances may
//		therefore be associated with an xDSL UNI.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. The two MSBs of
//			the first byte are the bearer channel ID. Excluding the first 2-bits of the first byte, the
//			remaining part of the ME ID is identical to that of this ME's parent PPTP xDSL UNI part 1. (R,
//			setbycreate) (mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 ME that
//			contains PM threshold values. Since no threshold value attribute number exceeds 7, a threshold
//			data 2 ME is optional. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Corrected Blocks
//			Corrected blocks: This attribute counts blocks received with errors that were corrected on this
//			channel. (R) (mandatory) (4-bytes)
//
//		Uncorrected Blocks
//			Uncorrected blocks: This attribute counts blocks received with uncorrectable errors on this
//			channel. (R) (mandatory) (4-bytes)
//
//		Transmitted Blocks
//			Transmitted blocks: This attribute counts encoded blocks transmitted on this channel. (R)
//			(mandatory) (4-bytes)
//
//		Received Blocks
//			Received blocks: This attribute counts encoded blocks received on this channel. (R) (mandatory)
//			(4-bytes)
//
//		Code Violations
//			Code violations: This attribute counts FEBE anomalies reported in the downstream bearer channel.
//			If the CRC is applied over multiple bearer channels, then each related FEBE anomaly increments
//			each of the counters related to the individual bearer channels. (R) (mandatory) (2-bytes)
//
//		Forward Error Corrections
//			Forward error corrections: This attribute counts FFEC anomalies reported in the downstream
//			bearer channel. If FEC is applied over multiple bearer channels, each related FFEC anomaly
//			increments each of the counters related to the individual bearer channels. (R) (mandatory)
//			(2-bytes)
//
type XdslXtuRChannelPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdslxturchannelperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "XdslXtuRChannelPerformanceMonitoringHistoryData",
		ClassID: 115,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xff00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 0),
			1: ByteField("IntervalEndTime", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read), false, false, 1),
			2: Uint16Field("ThresholdData12Id", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, 2),
			3: Uint32Field("CorrectedBlocks", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 3),
			4: Uint32Field("UncorrectedBlocks", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 4),
			5: Uint32Field("TransmittedBlocks", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 5),
			6: Uint32Field("ReceivedBlocks", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 6),
			7: Uint16Field("CodeViolations", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 7),
			8: Uint16Field("ForwardErrorCorrections", CounterAttributeType, 0, mapset.NewSetWith(Read), false, false, 8),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewXdslXtuRChannelPerformanceMonitoringHistoryData (class ID 115) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewXdslXtuRChannelPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdslxturchannelperformancemonitoringhistorydataBME, params...)
}
