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

// XdslXtuRPerformanceMonitoringHistoryData (class ID #113) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type XdslXtuRPerformanceMonitoringHistoryData struct {
	BaseManagedEntityDefinition
}

// NewXdslXtuRPerformanceMonitoringHistoryData (class ID 113 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslXtuRPerformanceMonitoringHistoryData(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "XdslXtuRPerformanceMonitoringHistoryData",
		ClassID:  113,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
			Create,
			Delete,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: []*AttributeDefinition{
			Uint16Field("ManagedEntityId", 0, Read|SetByCreate),
			ByteField("IntervalEndTime", 0, Read),
			Uint16Field("ThresholdData12Id", 0, Read|Write|SetByCreate),
			Uint16Field("LossOfFrameSeconds", 0, Read),
			Uint16Field("LossOfSignalSeconds", 0, Read),
			Uint16Field("LossOfPowerSeconds", 0, Read),
			Uint16Field("ErroredSeconds", 0, Read),
			Uint16Field("SeverelyErroredSeconds", 0, Read),
			Uint16Field("FecSeconds", 0, Read),
			Uint16Field("UnavailableSeconds", 0, Read),
			Uint16Field("LeftrDefectSeconds", 0, Read),
			Uint32Field("ErrorFreeBitsCounter", 0, Read),
			Uint32Field("MinimumErrorFreeThroughputMineftr", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &XdslXtuRPerformanceMonitoringHistoryData{entity}, nil
}