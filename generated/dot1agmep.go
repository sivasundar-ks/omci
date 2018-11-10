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

// Dot1AgMep (class ID #302) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type Dot1AgMep struct {
	BaseManagedEntityDefinition
}

// NewDot1AgMep (class ID 302 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewDot1AgMep(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "Dot1AgMep",
		ClassID:  302,
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
			Uint16Field("Layer2EntityPointer", 0, Read|Write|SetByCreate),
			ByteField("Layer2Type", 0, Read|Write|SetByCreate),
			Uint16Field("MaPointer", 0, Read|Write|SetByCreate),
			Uint16Field("MepId", 0, Read|Write|SetByCreate),
			ByteField("MepControl", 0, Read|Write|SetByCreate),
			Uint16Field("PrimaryVlan", 0, Read|Write|SetByCreate),
			ByteField("AdministrativeState", 0, Read|Write|SetByCreate),
			ByteField("CcmAndLtmPriority", 0, Read|Write|SetByCreate),
			Uint64Field("EgressIdentifier", 0, Read|Write|SetByCreate),
			UnknownField("PeerMepIds", 0, Read|Write),
			ByteField("EthAisControl", 0, Read|Write|SetByCreate),
			ByteField("FaultAlarmThreshold", 0, Read|Write|SetByCreate),
			Uint16Field("AlarmDeclarationSoakTime", 0, Read|Write),
			Uint16Field("AlarmClearSoakTime", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &Dot1AgMep{entity}, nil
}