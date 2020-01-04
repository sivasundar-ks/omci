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

// XdslDownstreamRfiBandsProfileClassID is the 16-bit ID for the OMCI
// Managed entity xDSL downstream RFI bands profile
const XdslDownstreamRfiBandsProfileClassID ClassID = ClassID(111)

var xdsldownstreamrfibandsprofileBME *ManagedEntityDefinition

// XdslDownstreamRfiBandsProfile (class ID #111)
//	This ME contains the downstream RFI bands profile for an xDSL UNI. Instances of this ME are
//	created and deleted by the OLT.
//
//	Relationships
//		An instance of this ME may be associated with zero or more instances of the PPTP xDSL UNI part
//		1.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. The value 0 is
//			reserved. (R, setbycreate) (mandatory) (2-bytes)
//
//		Downstream Rfi Bands Table
//			(R,-W) (mandatory for [ITU-T G.992.5], [ITU-T G.993.2]) (5 * N bytes where N is the number of
//			RFI bands)
//
//		Bands Valid
//			(R,-W) (mandatory) (1-byte)
//
type XdslDownstreamRfiBandsProfile struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdsldownstreamrfibandsprofileBME = &ManagedEntityDefinition{
		Name:    "XdslDownstreamRfiBandsProfile",
		ClassID: 111,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetNext,
			Set,
		),
		AllowedAttributeMask: 0xc000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0, mapset.NewSetWith(Read, SetByCreate), false, false, 0),
			1: TableField("DownstreamRfiBandsTable", TableAttributeType, TableInfo{nil, 5}, mapset.NewSetWith(Read, Write), false, false, 1),
			2: ByteField("BandsValid", UnsignedIntegerAttributeType, 0, mapset.NewSetWith(Read, Write), false, false, 2),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewXdslDownstreamRfiBandsProfile (class ID 111) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewXdslDownstreamRfiBandsProfile(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdsldownstreamrfibandsprofileBME, params...)
}
