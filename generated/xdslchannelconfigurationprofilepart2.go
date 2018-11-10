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

// XdslChannelConfigurationProfilePart2 (class ID #412) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type XdslChannelConfigurationProfilePart2 struct {
	BaseManagedEntityDefinition
}

// NewXdslChannelConfigurationProfilePart2 (class ID 412 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslChannelConfigurationProfilePart2(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "XdslChannelConfigurationProfilePart2",
		ClassID:  412,
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
			Uint32Field("MinimumExpectedThroughputForRetransmissionMinetrRtx", 0, Read|Write),
			Uint32Field("MaximumExpectedThroughputForRetransmissionMaxetrRtx", 0, Read|Write),
			Uint32Field("MaximumNetDataRateForRetransmissionMaxndrRtx", 0, Read|Write),
			ByteField("MaximumDelayForRetransmissionDelaymaxRtx", 0, Read|Write),
			ByteField("MinimumDelayForRetransmissionDelayminRtx", 0, Read|Write),
			ByteField("MinimumImpulseNoiseProtectionAgainstSingleHighImpulseNoiseEventShineForRetransmissionInpminShineRtx", 0, Read|Write),
			ByteField("MinimumImpulseNoiseProtectionAgainstShineForRetransmissionForSystemsUsing8625KhzSubcarrierSpacingInpmin8ShineRtx", 0, Read|Write),
			ByteField("ShineratioRtx", 0, Read|Write),
			ByteField("MinimumImpulseNoiseProtectionAgainstReinForRetransmissionInpminReinRtx", 0, Read|Write),
			ByteField("MinimumImpulseNoiseProtectionAgainstReinForRetransmissionForSystemsUsing8625KhzSubcarrierSpacingInpmin8ReinRtx", 0, Read|Write),
			ByteField("ReinInterArrivalTimeForRetransmissionIatReinRtx", 0, Read|Write),
			Uint32Field("TargetNetDataRateTargetNdr", 0, Read|Write),
			Uint32Field("TargetExpectedThroughputForRetransmissionTargetEtr", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &XdslChannelConfigurationProfilePart2{entity}, nil
}