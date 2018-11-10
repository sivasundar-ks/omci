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

// Dot1XPerformanceMonitoringHistoryData (class ID #292) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type Dot1XPerformanceMonitoringHistoryData struct {
	BaseManagedEntityDefinition
}

// NewDot1XPerformanceMonitoringHistoryData (class ID 292 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewDot1XPerformanceMonitoringHistoryData(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "Dot1XPerformanceMonitoringHistoryData",
		ClassID:  292,
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
			Uint32Field("EapolFramesReceived", 0, Read),
			Uint32Field("EapolFramesTransmitted", 0, Read),
			Uint32Field("EapolStartFramesReceived", 0, Read),
			Uint32Field("EapolLogoffFramesReceived", 0, Read),
			Uint32Field("InvalidEapolFramesReceived", 0, Read),
			Uint32Field("EapRespIdFramesReceived", 0, Read),
			Uint32Field("EapResponseFramesReceived", 0, Read),
			Uint32Field("EapInitialRequestFramesTransmitted", 0, Read),
			Uint32Field("EapRequestFramesTransmitted", 0, Read),
			Uint32Field("EapLengthErrorFramesReceived", 0, Read),
			Uint32Field("EapSuccessFramesGeneratedAutonomously", 0, Read),
			Uint32Field("EapFailureFramesGeneratedAutonomously", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &Dot1XPerformanceMonitoringHistoryData{entity}, nil
}
