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

import (
	"../../omci"
)

type VoipLineStatus struct {
	BaseManagedEntity
}

func NewVoipLineStatus(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		name:     "VoipLineStatus",
		classID:  141,
		entityID: eid,
		msgTypes: []omci.MsgType{
			omci.Get,
		},
		attributeList: []omci.IAttribute{
			omci.NewUint16Field("ManagedEntityId", 0, omci.Read),
			omci.NewUint16Field("VoipCodecUsed", 0, omci.Read),
			omci.NewByteField("VoipVoiceServerStatus", 0, omci.Read),
			omci.NewByteField("VoipPortSessionType", 0, omci.Read),
			omci.NewUint16Field("VoipCall1PacketPeriod", 0, omci.Read),
			omci.NewUint16Field("VoipCall2PacketPeriod", 0, omci.Read),
			omci.NewUnknownField("VoipCall1DestAddr", 0, omci.Read),
			omci.NewUnknownField("VoipCall2DestAddr", 0, omci.Read),
			omci.NewByteField("VoipLineState", 0, omci.Read),
			omci.NewByteField("EmergencyCallStatus", 0, omci.Read),
		},
	}
	entity.computeAttributeMask()
	return &VoipLineStatus{entity}, nil
}
