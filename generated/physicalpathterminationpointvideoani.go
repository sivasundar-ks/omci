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

type PhysicalPathTerminationPointVideoAni struct {
	BaseManagedEntity
}

func NewPhysicalPathTerminationPointVideoAni(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		Name:     "PhysicalPathTerminationPointVideoAni",
		ClassID:  90,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
		},
		AttributeMask: 0,
		Attributes: []IAttribute{
			NewUint16Field("ManagedEntityId", 0, Read),
			NewByteField("AdministrativeState", 0, Read|Write),
			NewByteField("OperationalState", 0, Read),
			NewByteField("Arc", 0, Read|Write),
			NewByteField("ArcInterval", 0, Read|Write),
			NewByteField("FrequencyRangeLow", 0, Read),
			NewByteField("FrequencyRangeHigh", 0, Read),
			NewByteField("SignalCapability", 0, Read),
			NewByteField("OpticalSignalLevel", 0, Read),
			NewByteField("PilotSignalLevel", 0, Read),
			NewByteField("SignalLevelMin", 0, Read),
			NewByteField("SignalLevelMax", 0, Read),
			NewUint32Field("PilotFrequency", 0, Read|Write),
			NewByteField("AgcMode", 0, Read|Write),
			NewByteField("AgcSetting", 0, Read|Write),
			NewByteField("VideoLowerOpticalThreshold", 0, Read|Write),
			NewByteField("VideoUpperOpticalThreshold", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &PhysicalPathTerminationPointVideoAni{entity}, nil
}
