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

// PhysicalPathTerminationPointCesUni (class ID #12) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type PhysicalPathTerminationPointCesUni struct {
	BaseManagedEntityDefinition
}

// NewPhysicalPathTerminationPointCesUni (class ID 12 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewPhysicalPathTerminationPointCesUni(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "PhysicalPathTerminationPointCesUni",
		ClassID:  12,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: []*AttributeDefinition{
			Uint16Field("ManagedEntityId", 0, Read),
			ByteField("ExpectedType", 0, Read|Write),
			ByteField("SensedType", 0, Read),
			ByteField("CesLoopbackConfiguration", 0, Read|Write),
			ByteField("AdministrativeState", 0, Read|Write),
			ByteField("OperationalState", 0, Read),
			ByteField("Framing", 0, Read|Write),
			ByteField("Encoding", 0, Read|Write),
			ByteField("LineLength", 0, Read|Write),
			ByteField("Ds1Mode", 0, Read|Write),
			ByteField("Arc", 0, Read|Write),
			ByteField("ArcInterval", 0, Read|Write),
			ByteField("LineType", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &PhysicalPathTerminationPointCesUni{entity}, nil
}
