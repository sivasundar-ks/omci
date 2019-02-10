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

const XdslLineConfigurationProfilePart2ClassId uint16 = 105

var xdsllineconfigurationprofilepart2BME *ManagedEntityDefinition

// XdslLineConfigurationProfilePart2 (class ID #105) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type XdslLineConfigurationProfilePart2 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdsllineconfigurationprofilepart2BME = &ManagedEntityDefinition{
		Name:    "XdslLineConfigurationProfilePart2",
		ClassID: 105,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFFE,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false),
			1:  Uint16Field("DownstreamMinimumTimeIntervalForUpshiftRateAdaptation", 0, Read|SetByCreate|Write, false, false, true),
			2:  Uint16Field("UpstreamMinimumTimeIntervalForUpshiftRateAdaptation", 0, Read|SetByCreate|Write, false, false, true),
			3:  Uint16Field("DownstreamDownshiftNoiseMargin", 0, Read|SetByCreate|Write, false, false, true),
			4:  Uint16Field("UpstreamDownshiftNoiseMargin", 0, Read|SetByCreate|Write, false, false, true),
			5:  Uint16Field("DownstreamMinimumTimeIntervalForDownshiftRateAdaptation", 0, Read|SetByCreate|Write, false, false, true),
			6:  Uint16Field("UpstreamMinimumTimeIntervalForDownshiftRateAdaptation", 0, Read|SetByCreate|Write, false, false, true),
			7:  ByteField("XtuImpedanceStateForced", 0, Read|SetByCreate|Write, false, false, true),
			8:  ByteField("L0Time", 0, Read|SetByCreate|Write, false, false, false),
			9:  ByteField("L2Time", 0, Read|SetByCreate|Write, false, false, false),
			10: Uint16Field("DownstreamMaximumNominalPowerSpectralDensity", 0, Read|SetByCreate|Write, false, false, false),
			11: Uint16Field("UpstreamMaximumNominalPowerSpectralDensity", 0, Read|SetByCreate|Write, false, false, false),
			12: ByteField("DownstreamMaximumNominalAggregateTransmitPower", 0, Read|SetByCreate|Write, false, false, false),
			13: ByteField("UpstreamMaximumNominalAggregateTransmitPower", 0, Read|SetByCreate|Write, false, false, false),
			14: Uint16Field("UpstreamMaximumAggregateReceivePower", 0, Read, false, false, false),
			15: ByteField("Vdsl2TransmissionSystemEnabling", 0, Read|SetByCreate|Write, false, false, true),
		},
	}
}

// NewXdslLineConfigurationProfilePart2 (class ID 105 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslLineConfigurationProfilePart2(params ...ParamData) (*ManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: xdsllineconfigurationprofilepart2BME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
