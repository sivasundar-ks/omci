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

const Dot1XPortExtensionPackageClassId uint16 = 290

var dot1xportextensionpackageBME *ManagedEntityDefinition

// Dot1XPortExtensionPackage (class ID #290) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type Dot1XPortExtensionPackage struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	dot1xportextensionpackageBME = &ManagedEntityDefinition{
		Name:    "Dot1XPortExtensionPackage",
		ClassID: 290,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFF0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read, false, false, false, 0),
			1:  ByteField("Dot1XEnable", 0, Read|Write, false, false, false, 1),
			2:  ByteField("ActionRegister", 0, Write, false, false, false, 2),
			3:  ByteField("AuthenticatorPaeState", 0, Read, false, false, true, 3),
			4:  ByteField("BackendAuthenticationState", 0, Read, false, false, true, 4),
			5:  ByteField("AdminControlledDirections", 0, Read|Write, false, false, true, 5),
			6:  ByteField("OperationalControlledDirections", 0, Read, false, false, true, 6),
			7:  ByteField("AuthenticatorControlledPortStatus", 0, Read, false, false, true, 7),
			8:  Uint16Field("QuietPeriod", 0, Read|Write, false, false, true, 8),
			9:  Uint16Field("ServerTimeoutPeriod", 0, Read|Write, false, false, true, 9),
			10: Uint16Field("ReAuthenticationPeriod", 0, Read, false, false, true, 10),
			11: ByteField("ReAuthenticationEnabled", 0, Read, false, false, true, 11),
			12: ByteField("KeyTransmissionEnabled", 0, Read|Write, false, false, true, 12),
		},
	}
}

// NewDot1XPortExtensionPackage (class ID 290 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewDot1XPortExtensionPackage(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(dot1xportextensionpackageBME, params...)
}
