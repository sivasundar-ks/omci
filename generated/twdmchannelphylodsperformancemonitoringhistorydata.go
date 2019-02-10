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

const TwdmChannelPhyLodsPerformanceMonitoringHistoryDataClassId uint16 = 444

var twdmchannelphylodsperformancemonitoringhistorydataBME *ManagedEntityDefinition

// TwdmChannelPhyLodsPerformanceMonitoringHistoryData (class ID #444) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type TwdmChannelPhyLodsPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	twdmchannelphylodsperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "TwdmChannelPhyLodsPerformanceMonitoringHistoryData",
		ClassID: 444,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetCurrentData,
			Set,
		),
		AllowedAttributeMask: 0XFFFE,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false),
			1:  ByteField("IntervalEndTime", 0, Read, false, false, false),
			2:  Uint16Field("ThresholdData12Id", 0, Read|SetByCreate|Write, false, false, false),
			3:  Uint64Field("TotalReceivedWordsProtectedByBitInterleavedParity32Bip32", 0, Read, false, false, false),
			4:  Uint32Field("Bip32BitErrorCount", 0, Read, false, false, false),
			5:  Uint32Field("CorrectedPsbdHecErrorCount", 0, Read, false, false, false),
			6:  Uint32Field("UncorrectablePsbdHecErrorCount", 0, Read, false, false, false),
			7:  Uint32Field("CorrectedDownstreamFsHeaderHecErrorCount", 0, Read, false, false, false),
			8:  Uint32Field("UncorrectableDownstreamFsHeaderHecErrorCount", 0, Read, false, false, false),
			9:  Uint32Field("TotalNumberOfLodsEvents", 0, Read, false, false, false),
			10: Uint32Field("LodsEventsRestoredInOperatingTwdmChannel", 0, Read, false, false, false),
			11: Uint32Field("LodsEventsRestoredInProtectionTwdmChannel", 0, Read, false, false, false),
			12: Uint32Field("LodsEventsRestoredInDiscretionaryTwdmChannel", 0, Read, false, false, false),
			13: Uint32Field("LodsEventsResultingInReactivation", 0, Read, false, false, false),
			14: Uint32Field("LodsEventsResultingInReactivationAfterRetuningToProtectionTwdmChannel", 0, Read, false, false, false),
			15: Uint32Field("LodsEventsResultingInReactivationAfterRetuningToDiscretionaryTwdmChannel", 0, Read, false, false, false),
		},
	}
}

// NewTwdmChannelPhyLodsPerformanceMonitoringHistoryData (class ID 444 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewTwdmChannelPhyLodsPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: twdmchannelphylodsperformancemonitoringhistorydataBME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
