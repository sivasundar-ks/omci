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
 *              https://github.com/cboling/OMCI-parser
 */
package generated

type XgPonTcPerformanceMonitoringHistoryData struct {
	BaseManagedEntity
}

func NewXgPonTcPerformanceMonitoringHistoryData(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		Name:     "XgPonTcPerformanceMonitoringHistoryData",
		ClassID:  344,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
			Create,
			Delete,
		},
		AttributeMask: 0,
		Attributes: []IAttribute{
			NewUint16Field("ManagedEntityId", 0, Read|SetByCreate),
			NewByteField("IntervalEndTime", 0, Read),
			NewUint16Field("ThresholdData12Id", 0, Read|Write|SetByCreate),
			NewUint32Field("PsbdHecErrorCount", 0, Read),
			NewUint32Field("XgtcHecErrorCount", 0, Read),
			NewUint32Field("UnknownProfileCount", 0, Read),
			NewUint32Field("TransmittedXgPonEncapsulationMethodXgemFrames", 0, Read),
			NewUint32Field("FragmentXgemFrames", 0, Read),
			NewUint32Field("XgemHecLostWordsCount", 0, Read),
			NewUint32Field("XgemKeyErrors", 0, Read),
			NewUint32Field("XgemHecErrorCount", 0, Read),
			NewUint64Field("TransmittedBytesInNonIdleXgemFrames", 0, Read),
			NewUint64Field("ReceivedBytesInNonIdleXgemFrames", 0, Read),
			NewUint32Field("LossOfDownstreamSynchronizationLodsEventCount", 0, Read),
			NewUint32Field("LodsEventRestoredCount", 0, Read),
			NewUint32Field("OnuReactivationByLodsEvents", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &XgPonTcPerformanceMonitoringHistoryData{entity}, nil
}
