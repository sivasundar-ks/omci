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

type XdslLineConfigurationProfilePart3 struct {
	BaseManagedEntity
}

func NewXdslLineConfigurationProfilePart3(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		Name:     "XdslLineConfigurationProfilePart3",
		ClassID:  106,
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
			NewByteField("LoopDiagnosticsModeForcedLdsf", 0, Read|Write|SetByCreate),
			NewByteField("AutomodeColdStartForced", 0, Read|Write|SetByCreate),
			NewByteField("L2Atpr", 0, Read|Write|SetByCreate),
			NewByteField("L2Atprt", 0, Read|Write|SetByCreate),
			NewByteField("ForceInpDownstream", 0, Read|Write),
			NewByteField("ForceInpUpstream", 0, Read|Write),
			NewByteField("UpdateRequestFlagForNearEndTestParameters", 0, Read|Write),
			NewByteField("UpdateRequestFlagForFarEndTestParameters", 0, Read|Write),
			NewUint16Field("InmInterArrivalTimeOffsetUpstream", 0, Read|Write),
			NewByteField("InmInterArrivalTimeStepUpstream", 0, Read|Write),
			NewByteField("InmClusterContinuationValueUpstream", 0, Read|Write),
			NewByteField("InmEquivalentInpModeUpstream", 0, Read|Write),
			NewUint16Field("InmInterArrivalTimeOffsetDownstream", 0, Read|Write),
			NewByteField("InmInterArrivalTimeStepDownstream", 0, Read|Write),
			NewByteField("InmClusterContinuationValueDownstream", 0, Read|Write),
			NewByteField("InmEquivalentInpModeDownstream", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &XdslLineConfigurationProfilePart3{entity}, nil
}
