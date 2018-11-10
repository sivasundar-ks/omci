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

// FastChannelConfigurationProfile (class ID #432) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type FastChannelConfigurationProfile struct {
	BaseManagedEntityDefinition
}

// NewFastChannelConfigurationProfile (class ID 432 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewFastChannelConfigurationProfile(params ...ParamData) (IManagedEntityDefinition, error) {
	eid := decodeEntityID(params...)
	entity := BaseManagedEntityDefinition{
		Name:     "FastChannelConfigurationProfile",
		ClassID:  432,
		EntityID: eid,
		MessageTypes: []MsgType{
			Set,
			Get,
			Create,
			Delete,
		},
		AllowedAttributeMask: 0,
		AttributeDefinitions: []*AttributeDefinition{
			Uint32Field("MaximumNetDataRateMaxndr", 0, Read|Write),
			Uint32Field("MinimumExpectedThroughputMinetr", 0, Read|Write),
			Uint32Field("MaximumGammaDataRateMaxgdr", 0, Read|Write),
			Uint32Field("MinimumGammaDataRateMingdr", 0, Read|Write),
			Uint32Field("MaximumDelayDelaymax", 0, Read|Write),
			Uint16Field("MinimumImpulseNoiseProtectionAgainstShineInpminShine", 0, Read|Write),
			ByteField("ShineRatioShineratio", 0, Read|Write),
			ByteField("MinimumImpulseNoiseProtectionAgainstReinInpminRein", 0, Read|Write),
			ByteField("ReinInterArrivalTimeIatRein", 0, Read|Write),
			ByteField("MinimumReedSolomonRfecNfecRatioRnratio", 0, Read|Write),
			ByteField("RtxTcTestmodeRtxTestmode", 0, Read|Write),
		},
	}
	entity.computeAttributeMask()
	return &FastChannelConfigurationProfile{entity}, nil
}
