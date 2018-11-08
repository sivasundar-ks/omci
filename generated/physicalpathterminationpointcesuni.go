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

type PhysicalPathTerminationPointCesUni struct {
	BaseManagedEntity
}

func NewPhysicalPathTerminationPointCesUni(params ...ParamData) (IManagedEntity, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntity{
		Name:     "PhysicalPathTerminationPointCesUni",
		ClassID:  12,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
		},
		AttributeMask: 0,
		Attributes: []IAttribute{
			NewUint16Field("ManagedEntityId", 0, Read),
			NewByteField("ExpectedType", 0, Read|Write),
			NewByteField("SensedType", 0, Read),
			NewByteField("CesLoopbackConfiguration", 0, Read|Write),
			NewByteField("AdministrativeState", 0, Read|Write),
			NewByteField("OperationalState", 0, Read),
			NewByteField("Framing", 0, Read|Write),
			NewByteField("Encoding", 0, Read|Write),
			NewByteField("LineLength", 0, Read|Write),
			NewByteField("Ds1Mode", 0, Read|Write),
			NewByteField("Arc", 0, Read|Write),
			NewByteField("ArcInterval", 0, Read|Write),
			NewByteField("LineType", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &PhysicalPathTerminationPointCesUni{entity}, nil
}
