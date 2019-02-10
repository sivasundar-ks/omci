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

const XdslChannelConfigurationProfilePart2ClassId uint16 = 412

var xdslchannelconfigurationprofilepart2BME *ManagedEntityDefinition

// XdslChannelConfigurationProfilePart2 (class ID #412) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type XdslChannelConfigurationProfilePart2 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdslchannelconfigurationprofilepart2BME = &ManagedEntityDefinition{
		Name:    "XdslChannelConfigurationProfilePart2",
		ClassID: 412,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFF8,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read|SetByCreate, false, false, false),
			1:  Uint32Field("MinimumExpectedThroughputForRetransmissionMinetrRtx", 0, Read|Write, false, false, false),
			2:  Uint32Field("MaximumExpectedThroughputForRetransmissionMaxetrRtx", 0, Read|Write, false, false, false),
			3:  Uint32Field("MaximumNetDataRateForRetransmissionMaxndrRtx", 0, Read|Write, false, false, false),
			4:  ByteField("MaximumDelayForRetransmissionDelaymaxRtx", 0, Read|Write, false, false, false),
			5:  ByteField("MinimumDelayForRetransmissionDelayminRtx", 0, Read|Write, false, false, false),
			6:  ByteField("MinimumImpulseNoiseProtectionAgainstSingleHighImpulseNoiseEventShineForRetransmissionInpminShineRtx", 0, Read|Write, false, false, false),
			7:  ByteField("MinimumImpulseNoiseProtectionAgainstShineForRetransmissionForSystemsUsing8625KhzSubcarrierSpacingInpmin8ShineRtx", 0, Read|Write, false, false, false),
			8:  ByteField("ShineratioRtx", 0, Read|Write, false, false, false),
			9:  ByteField("MinimumImpulseNoiseProtectionAgainstReinForRetransmissionInpminReinRtx", 0, Read|Write, false, false, false),
			10: ByteField("MinimumImpulseNoiseProtectionAgainstReinForRetransmissionForSystemsUsing8625KhzSubcarrierSpacingInpmin8ReinRtx", 0, Read|Write, false, false, false),
			11: ByteField("ReinInterArrivalTimeForRetransmissionIatReinRtx", 0, Read|Write, false, false, false),
			12: Uint32Field("TargetNetDataRateTargetNdr", 0, Read|Write, false, false, false),
			13: Uint32Field("TargetExpectedThroughputForRetransmissionTargetEtr", 0, Read|Write, false, false, false),
		},
	}
}

// NewXdslChannelConfigurationProfilePart2 (class ID 412 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslChannelConfigurationProfilePart2(params ...ParamData) (*ManagedEntity, error) {
	entity := &ManagedEntity{
		Definition: xdslchannelconfigurationprofilepart2BME,
		Attributes: make(map[string]interface{}),
	}
	if err := entity.setAttributes(params...); err != nil {
		return nil, err
	}
	return entity, nil
}
