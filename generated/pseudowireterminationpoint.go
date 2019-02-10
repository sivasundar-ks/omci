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

const PseudowireTerminationPointClassId uint16 = 282

var pseudowireterminationpointBME *ManagedEntityDefinition

// PseudowireTerminationPoint (class ID #282) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type PseudowireTerminationPoint struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	pseudowireterminationpointBME = &ManagedEntityDefinition{
		Name:    "PseudowireTerminationPoint",
		ClassID: 282,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFE,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false),
			1:  ByteField("UnderlyingTransport", 0, Read|SetByCreate|Write, false, false, false),
			2:  ByteField("ServiceType", 0, Read|SetByCreate|Write, false, false, false),
			3:  ByteField("Signalling", 0, Read|SetByCreate|Write, false, false, false),
			4:  Uint16Field("TdmUniPointer", 0, Read|SetByCreate|Write, false, false, false),
			5:  Uint16Field("NorthSidePointer", 0, Read|SetByCreate|Write, false, false, false),
			6:  Uint16Field("FarEndIpInfo", 0, Read|SetByCreate|Write, false, false, false),
			7:  Uint16Field("PayloadSize", 0, Read|SetByCreate|Write, false, false, false),
			8:  ByteField("PayloadEncapsulationDelay", 0, Read|SetByCreate|Write, false, false, false),
			9:  ByteField("TimingMode", 0, Read|Write, false, false, false),
			10: Uint64Field("TransmitCircuitId", 0, Read|Write, false, false, false),
			11: Uint64Field("ExpectedCircuitId", 0, Read|Write, false, false, false),
			12: Uint64Field("ReceivedCircuitId", 0, Read, false, false, false),
			13: Uint16Field("ExceptionPolicy", 0, Read|Write, false, false, true),
			14: ByteField("Arc", 0, Read|Write, true, false, true),
			15: ByteField("ArcInterval", 0, Read|Write, false, false, true),
		},
	}
}

// NewPseudowireTerminationPoint (class ID 282 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewPseudowireTerminationPoint(params ...ParamData) (*ManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: pseudowireterminationpointBME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
