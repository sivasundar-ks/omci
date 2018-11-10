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

// EfmBondingPortPerformanceMonitoringHistoryData (class ID #424) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type EfmBondingPortPerformanceMonitoringHistoryData struct {
	BaseManagedEntityDefinition
}

// NewEfmBondingPortPerformanceMonitoringHistoryData (class ID 424 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewEfmBondingPortPerformanceMonitoringHistoryData(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "EfmBondingPortPerformanceMonitoringHistoryData",
		ClassID:  424,
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
			Uint32Field("RxFrames", 0, Read),
			Uint32Field("TxFrames", 0, Read),
			Uint32Field("RxBytes", 0, Read),
			Uint32Field("TxBytes", 0, Read),
			Uint32Field("TxDiscardedFrames", 0, Read),
			Uint32Field("TxDiscardedBytes", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &EfmBondingPortPerformanceMonitoringHistoryData{entity}, nil
}