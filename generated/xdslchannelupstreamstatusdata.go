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

const XdslChannelUpstreamStatusDataClassId uint16 = 103

var xdslchannelupstreamstatusdataBME *ManagedEntityDefinition

// XdslChannelUpstreamStatusData (class ID #103) defines the basic
// Managed Entity definition that is further extended by types that support
// packet encode/decode and user create managed entities.
type XdslChannelUpstreamStatusData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdslchannelupstreamstatusdataBME = &ManagedEntityDefinition{
		Name:    "XdslChannelUpstreamStatusData",
		ClassID: 103,
		MessageTypes: mapset.NewSetWith(
			Get,
		),
		AllowedAttributeMask: 0XFFE0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, Read, false, false, false, 0),
			1:  ByteField("ActualInterleavingDelay", 0, Read, false, false, false, 1),
			2:  Uint32Field("ActualDataRate", 0, Read, false, false, false, 2),
			3:  Uint32Field("PreviousDataRate", 0, Read, false, false, false, 3),
			4:  ByteField("ActualImpulseNoiseProtection", 0, Read, false, false, false, 4),
			5:  ByteField("ImpulseNoiseProtectionReportingMode", 0, Read, false, false, false, 5),
			6:  ByteField("ActualSizeOfReedSolomonCodeword", 0, Read, false, false, false, 6),
			7:  ByteField("ActualNumberOfReedSolomonRedundancyBytes", 0, Read, false, false, false, 7),
			8:  Uint16Field("ActualNumberOfBitsPerSymbol", 0, Read, false, false, false, 8),
			9:  Uint16Field("ActualInterleavingDepth", 0, Read, false, false, false, 9),
			10: ByteField("ActualInterleavingBlockLength", 0, Read, false, false, false, 10),
			11: ByteField("ActualLatencyPath", 0, Read, false, false, false, 11),
		},
	}
}

// NewXdslChannelUpstreamStatusData (class ID 103 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewXdslChannelUpstreamStatusData(params ...ParamData) (*ManagedEntity, error) {
	return NewManagedEntity(xdslchannelupstreamstatusdataBME, params...)
}
