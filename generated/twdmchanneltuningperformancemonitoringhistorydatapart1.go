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

type TwdmChannelTuningPerformanceMonitoringHistoryDataPart1 struct {
	BaseManagedEntity
}

func NewTwdmChannelTuningPerformanceMonitoringHistoryDataPart1(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		Name:     "TwdmChannelTuningPerformanceMonitoringHistoryDataPart1",
		ClassID:  449,
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
			NewUint32Field("TuningControlRequestsForRxOnlyOrRxAndTx", 0, Read),
			NewUint32Field("TuningControlRequestsForTxOnly", 0, Read),
			NewUint32Field("TuningControlRequestsRejectedIntSfc", 0, Read),
			NewUint32Field("TuningControlRequestsRejectedDsXxx", 0, Read),
			NewUint32Field("TuningControlRequestsRejectedUsXxx", 0, Read),
			NewUint32Field("TuningControlRequestsFulfilledWithOnuReacquiredAtTargetChannel", 0, Read),
			NewUint32Field("TuningControlRequestsFailedDueToTargetDsWavelengthChannelNotFound", 0, Read),
			NewUint32Field("TuningControlRequestsFailedDueToNoFeedbackInTargetDsWavelengthChannel", 0, Read),
			NewUint32Field("TuningControlRequestsResolvedWithOnuReacquiredAtDiscretionaryChannel", 0, Read),
			NewUint32Field("TuningControlRequestsRollbackComDs", 0, Read),
			NewUint32Field("TuningControlRequestsRollbackDsXxx", 0, Read),
			NewUint32Field("TuningControlRequestsRollbackUsXxx", 0, Read),
			NewUint32Field("TuningControlRequestsFailedWithOnuReactivation", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &TwdmChannelTuningPerformanceMonitoringHistoryDataPart1{entity}, nil
}
