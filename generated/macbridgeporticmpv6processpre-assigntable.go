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

// MacBridgePortIcmpv6ProcessPreAssignTable (class ID #348) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type MacBridgePortIcmpv6ProcessPreAssignTable struct {
	BaseManagedEntityDefinition
}

// NewMacBridgePortIcmpv6ProcessPreAssignTable (class ID 348 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewMacBridgePortIcmpv6ProcessPreAssignTable(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "MacBridgePortIcmpv6ProcessPreAssignTable",
		ClassID:  348,
		EntityID: eid,
		MessageTypes: []MsgType{
			Get,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: []*AttributeDefinition{
			Uint16Field("ManagedEntityId", 0, Read),
			ByteField("Icmpv6ErrorMessagesProcessing", 0, Read|Write),
			ByteField("Icmpv6InformationalMessagesProcessing", 0, Read|Write),
			ByteField("RouterSolicitationProcessing", 0, Read|Write),
			ByteField("RouterAdvertisementProcessing", 0, Read|Write),
			ByteField("NeighbourSolicitationProcessing", 0, Read|Write),
			ByteField("NeighbourAdvertisementProcessing", 0, Read|Write),
			ByteField("RedirectProcessing", 0, Read|Write),
			ByteField("MulticastListenerQueryProcessing", 0, Read|Write),
			ByteField("UnknownIcmpv6Processing", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &MacBridgePortIcmpv6ProcessPreAssignTable{entity}, nil
}