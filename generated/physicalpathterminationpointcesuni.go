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

import "github.com/deckarep/golang-set"

const PhysicalPathTerminationPointCesUniClassId uint16 = 12

var physicalpathterminationpointcesuniBME *ManagedEntityDefinition

// PhysicalPathTerminationPointCesUni (class ID #12) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type PhysicalPathTerminationPointCesUni struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	physicalpathterminationpointcesuniBME = &ManagedEntityDefinition{
		Name:    "PhysicalPathTerminationPointCesUni",
		ClassID: 12,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFF0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read, false, false, false),
			1:  ByteField("ExpectedType", 0, Read|Write, false, false, false),
			2:  ByteField("SensedType", 0, Read, true, false, false),
			3:  ByteField("CesLoopbackConfiguration", 0, Read|Write, true, false, false),
			4:  ByteField("AdministrativeState", 0, Read|Write, false, false, false),
			5:  ByteField("OperationalState", 0, Read, true, false, true),
			6:  ByteField("Framing", 0, Read|Write, false, false, true),
			7:  ByteField("Encoding", 0, Read|Write, false, false, false),
			8:  ByteField("LineLength", 0, Read|Write, false, false, true),
			9:  ByteField("Ds1Mode", 0, Read|Write, false, false, true),
			10: ByteField("Arc", 0, Read|Write, true, false, true),
			11: ByteField("ArcInterval", 0, Read|Write, false, false, true),
			12: ByteField("LineType", 0, Read|Write, false, false, false),
		},
	}
}

// NewPhysicalPathTerminationPointCesUni (class ID 12 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewPhysicalPathTerminationPointCesUni(params ...ParamData) (*ManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: physicalpathterminationpointcesuniBME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
