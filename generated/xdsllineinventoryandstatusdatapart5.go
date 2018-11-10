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

// XdslLineInventoryAndStatusDataPart5 (class ID #325) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type XdslLineInventoryAndStatusDataPart5 struct {
	BaseManagedEntityDefinition
}

// NewXdslLineInventoryAndStatusDataPart5 (class ID 325 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslLineInventoryAndStatusDataPart5(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "XdslLineInventoryAndStatusDataPart5",
		ClassID:  325,
		EntityID: eid,
		MessageTypes: []MsgType{
			Get,
			GetNext,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: []*AttributeDefinition{
			Uint16Field("ManagedEntityId", 0, Read),
			Uint16Field("FextDownstreamSnrMargin", 0, Read),
			Uint16Field("NextDownstreamSnrMargin", 0, Read),
			Uint16Field("FextUpstreamSnrMargin", 0, Read),
			Uint16Field("NextUpstreamSnrMargin", 0, Read),
			Uint32Field("FextDownstreamMaximumAttainableDataRate", 0, Read),
			Uint32Field("NextDownstreamMaximumAttainableDataRate", 0, Read),
			Uint32Field("FextUpstreamMaximumAttainableDataRate", 0, Read),
			Uint32Field("NextUpstreamMaximumAttainableDataRate", 0, Read),
			Uint16Field("FextDownstreamActualPowerSpectralDensity", 0, Read),
			Uint16Field("NextDownstreamActualPowerSpectralDensity", 0, Read),
			Uint16Field("FextUpstreamActualPowerSpectralDensity", 0, Read),
			Uint16Field("NextUpstreamActualPowerSpectralDensity", 0, Read),
			Uint16Field("FextDownstreamActualAggregateTransmitPower", 0, Read),
			Uint16Field("NextDownstreamActualAggregateTransmitPower", 0, Read),
			Uint16Field("FextUpstreamActualAggregateTransmitPower", 0, Read),
			Uint16Field("NextUpstreamActualAggregateTransmitPower", 0, Read),
		},
	}
	entity.computeAttributeMask()
	return &XdslLineInventoryAndStatusDataPart5{entity}, nil
}
